package gopay

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mariusvanderwijden/go-pay/l2"
	"golang.org/x/net/context"
)

var ChannelAddr = common.HexToAddress("0x01")
var defaultTimeout = time.Minute

type ChallengeError struct {
	round *big.Int
}

func (e *ChallengeError) Error() string {
	return fmt.Sprintf("Challenge is in round %s", e.round)
}

type Backend interface {
	bind.DeployBackend
	bind.ContractBackend
}

type roundStr struct {
	number  *big.Int
	valueA  *big.Int
	valueB  *big.Int
	peerSig []byte
}

type Channel struct {
	channel *l2.Channel
	backend Backend
	ID      [32]byte
	A       common.Address
	B       common.Address
	ValueA  *big.Int
	ValueB  *big.Int
	Round   *big.Int
	// internal fields
	sumFunds       *big.Int
	peerSigs       map[*big.Int]*roundStr
	userIsProposer bool
}

func NewChannel(backend Backend, addrA, addrB common.Address, valueA, valueB *big.Int) (*Channel, error) {
	channel, err := l2.NewChannel(ChannelAddr, backend)
	if err != nil {
		return nil, err
	}
	return &Channel{
		channel:  channel,
		backend:  backend,
		A:        addrA,
		B:        addrB,
		ValueA:   valueA,
		ValueB:   valueB,
		Round:    big.NewInt(0),
		sumFunds: new(big.Int).Add(valueA, valueB),
		peerSigs: make(map[*big.Int]*roundStr),
	}, nil
}

func (c *Channel) Open(auth *bind.TransactOpts) ([32]byte, error) {
	var id [32]byte
	if _, err := crand.Read(id[:]); err != nil {
		return [32]byte{}, err
	}
	auth.Value = c.ValueA
	defer func() {
		auth.Value = big.NewInt(0)
	}()

	tx, err := c.channel.Open(auth, id, c.B, c.ValueA, c.ValueB)
	if err != nil {
		return [32]byte{}, err
	}
	if err := mustMine(c.backend, tx); err != nil {
		return [32]byte{}, err
	}
	c.userIsProposer = true
	c.ID = id
	return id, nil
}

func (c *Channel) Accept(auth *bind.TransactOpts, id [32]byte) error {
	auth.Value = c.ValueB
	defer func() {
		auth.Value = big.NewInt(0)
	}()

	tx, err := c.channel.Accept(auth, id)
	if err != nil {
		return err
	}
	if err := mustMine(c.backend, tx); err != nil {
		return err
	}
	c.userIsProposer = false
	c.ID = id
	return nil
}

func (c *Channel) CoopClose(auth *bind.TransactOpts) error {
	round, err := c.latestPeerSig()
	if err != nil {
		return err
	}
	tx, err := c.channel.CooperativeClose(auth, c.ID, round.valueA, round.valueB, round.peerSig)
	if err != nil {
		return err
	}

	return mustMine(c.backend, tx)
}

func (c *Channel) StartForceClose(auth *bind.TransactOpts, peerSignature []byte) (time.Time, error) {
	// Setup event watching
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	ch := make(chan *l2.ChannelClosing)
	c.channel.WatchClosing(&bind.WatchOpts{Context: ctx}, ch, [][32]byte{c.ID})
	// Call challenge
	tx, err := c.channel.Challenge(auth, c.ID, c.ValueA, c.ValueB, c.Round, peerSignature)
	if err != nil {
		return time.Time{}, err
	}
	if err := mustMine(c.backend, tx); err != nil {
		return time.Time{}, err
	}
	// Check for the event
	timer := time.NewTimer(defaultTimeout)
	select {
	case event := <-ch:
		if event.Round.Cmp(c.Round) < 0 {
			return time.Time{}, &ChallengeError{round: event.Round}
		}
		return time.Unix(int64(event.Time), 0), nil
	case <-timer.C:
		return time.Time{}, errors.New("Waiting for event timed out")
	}
}

func (c *Channel) Dispute(auth *bind.TransactOpts, start uint64) (time.Time, error) {
	// Setup event filtering
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	it, err := c.channel.FilterClosing(&bind.FilterOpts{Context: ctx, Start: start}, [][32]byte{c.ID})
	if err != nil {
		return time.Time{}, err
	}
	var newest *l2.ChannelClosing
	for it.Next() {
		newest = it.Event
	}
	if newest == nil {
		return time.Time{}, errors.New("no dispute event found")
	}
	if newest.Round.Cmp(c.Round) > 0 {
		return time.Time{}, fmt.Errorf("dispute has higher round than local, local: %v, onchain: %v", c.Round, newest.Round)
	}

	fmt.Printf("Disputed with old state: %v got %v", newest.Round, c.Round)
	// Setup event watching
	ctx, cancel = context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	ch := make(chan *l2.ChannelClosing)
	c.channel.WatchClosing(&bind.WatchOpts{Context: ctx}, ch, [][32]byte{c.ID})
	// Call challenge
	round, err := c.latestPeerSig()
	if err != nil {
		return time.Time{}, err
	}
	tx, err := c.channel.DisputeChallenge(auth, c.ID, round.valueA, round.valueB, round.number, round.peerSig)
	if err != nil {
		return time.Time{}, err
	}
	if err := mustMine(c.backend, tx); err != nil {
		return time.Time{}, err
	}
	// Check for the event
	timer := time.NewTimer(defaultTimeout)
	select {
	case event := <-ch:
		if event.Round.Cmp(c.Round) < 0 {
			return time.Time{}, &ChallengeError{round: event.Round}
		}
		return time.Unix(int64(event.Time), 0), nil
	case <-timer.C:
		return time.Time{}, errors.New("waiting for event timed out")
	}
}

func (c *Channel) FinishForceClose(auth *bind.TransactOpts, peerSignature []byte) error {
	tx, err := c.channel.ForceClose(auth, c.ID)
	if err != nil {
		return err
	}
	return mustMine(c.backend, tx)
}

func (c *Channel) SendMoney(value *big.Int) ([32]byte, error) {
	// TODO return mimedata for signing
	var (
		valA  *big.Int
		valB  *big.Int
		round *big.Int
	)
	if c.userIsProposer {
		if value.Cmp(c.ValueA) > 0 {
			return common.Hash{}, errors.New("not enough funds")
		}
		valA = new(big.Int).Sub(c.ValueA, value)
		valB = new(big.Int).Add(c.ValueB, value)
		round = new(big.Int).Add(c.Round, big.NewInt(1))

	} else {
		if value.Cmp(c.ValueB) > 0 {
			return common.Hash{}, errors.New("not enough funds")
		}
		valA = new(big.Int).Add(c.ValueA, value)
		valB = new(big.Int).Sub(c.ValueB, value)
		round = new(big.Int).Add(c.Round, big.NewInt(1))
	}
	hash, err := l2.HashState(c.channel, c.ID, c.A, c.B, valA, valB, round)
	if err != nil {
		return common.Hash{}, err
	}
	c.ValueA = valA
	c.ValueB = valB
	c.Round = round
	return hash, nil
}

func (c *Channel) CreateCooperativeClose() ([32]byte, error) {
	hash, err := l2.HashState(c.channel, c.ID, c.A, c.B, c.ValueA, c.ValueB, l2.CoopCloseRound)
	return hash, err
}

func (c *Channel) ReceivedSignature(valueA, valueB, round *big.Int, sig []byte) error {
	if new(big.Int).Add(valueA, valueB).Cmp(c.sumFunds) != 0 {
		return fmt.Errorf("sum of funds not equal, want: %v got %v and %v", c.sumFunds, valueA, valueB)
	}
	if round.Cmp(c.Round) <= 0 && c.Round.Cmp(l2.CoopCloseRound) != 0 {
		return fmt.Errorf("invalid round: got %v local %v", round, c.Round)
	}
	hash, err := l2.HashState(c.channel, c.ID, c.A, c.B, valueA, valueB, round)
	if err != nil {
		return err
	}
	var signer common.Address
	if c.userIsProposer {
		if c.ValueA.Cmp(valueA) > 0 {
			return fmt.Errorf("valueA decreased by B, from %v to %v", c.ValueA, valueA)
		}
		signer = c.B
	} else {
		if c.ValueB.Cmp(valueB) > 0 {
			return fmt.Errorf("valueB decreased by A, from %v to %v", c.ValueB, valueB)
		}
		signer = c.A
	}
	sig[len(sig)-1] -= 27
	pkBytes, err := crypto.Ecrecover(hash[:], sig)
	if err != nil {
		return err
	}
	pk, err := crypto.UnmarshalPubkey(pkBytes)
	if err != nil {
		return err
	}
	if recovered := crypto.PubkeyToAddress(*pk); signer != recovered {
		return fmt.Errorf("invalid signer recovered, got %v want %v", recovered, signer)
	}
	// Signature looks good, update the state
	sig[len(sig)-1] += 27
	rnd := &roundStr{
		number:  round,
		valueA:  valueA,
		valueB:  valueB,
		peerSig: sig,
	}
	c.Round = round
	c.ValueA = valueA
	c.ValueB = valueB
	c.peerSigs[round] = rnd
	return nil
}

func mustMine(backend Backend, tx *types.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, backend, tx)
	if err != nil {
		return err
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("transaction %v reverted", receipt.TxHash)
	}
	return nil
}

func (c *Channel) latestPeerSig() (*roundStr, error) {
	min := big.NewInt(0)
	if c.Round.Cmp(big.NewInt(2)) > 0 {
		min = new(big.Int).Sub(c.Round, big.NewInt(2))
	}
	for i := c.Round; i.Cmp(min) > 0; i.Sub(i, big.NewInt(1)) {
		if c.peerSigs[i] != nil {
			return c.peerSigs[i], nil
		}
	}
	return nil, fmt.Errorf("couldn't find proper signature within round %v", c.Round)
}
