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

// Channel wraps an SSPC payment channel.
type Channel struct {
	channel        *l2.Channel
	backend        Backend
	metadata       MetaData
	valueA         *big.Int
	valueB         *big.Int
	round          *big.Int
	sumFunds       *big.Int
	peerSigs       map[*big.Int]*roundStr
	userIsProposer bool
	inactive       bool
}

// NewChannel initializes the channel object.
func NewChannel(backend Backend, addrA, addrB common.Address, valueA, valueB *big.Int) (*Channel, error) {
	channel, err := l2.NewChannel(ChannelAddr, backend)
	if err != nil {
		return nil, err
	}

	metadata := MetaData{
		A: addrA,
		B: addrB,
	}

	return &Channel{
		channel:  channel,
		backend:  backend,
		metadata: metadata,
		valueA:   valueA,
		valueB:   valueB,
		round:    big.NewInt(0),
		sumFunds: new(big.Int).Add(valueA, valueB),
		peerSigs: make(map[*big.Int]*roundStr),
	}, nil
}

// State contains the state of a channel.
type State struct {
	ValueA *big.Int
	ValueB *big.Int
	Round  *big.Int
}

// Current state returns the state of a channel.
func (c *Channel) CurrentState() State {
	return State{
		ValueA: c.valueA,
		ValueB: c.valueB,
		Round:  c.round,
	}
}

// MetaData contains metadata concerning the channel.
type MetaData struct {
	ID [32]byte
	A  common.Address
	B  common.Address
}

// Current state returns the metadata of a channel.
func (c *Channel) MetaData() MetaData {
	return c.metadata
}

// Open proposes to open a channel with the initially provided parameters.
// Open returns the resulting channel ID.
func (c *Channel) Open(auth *bind.TransactOpts) ([32]byte, error) {
	var id [32]byte
	if _, err := crand.Read(id[:]); err != nil {
		return [32]byte{}, err
	}
	auth.Value = c.valueA
	defer func() {
		auth.Value = big.NewInt(0)
	}()

	tx, err := c.channel.Open(auth, id, c.metadata.B, c.valueA, c.valueB)
	if err != nil {
		return [32]byte{}, err
	}
	if err := mustMine(c.backend, tx); err != nil {
		return [32]byte{}, err
	}
	c.userIsProposer = true
	c.metadata.ID = id
	return id, nil
}

// Accept accepts the proposed channel opening with ID id.
func (c *Channel) Accept(auth *bind.TransactOpts, id [32]byte) error {
	auth.Value = c.valueB
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
	c.metadata.ID = id
	return nil
}

// CoopClose sends a cooperative close with the latest channel state.
// The latest channel state must be a valid cooperative close message.
func (c *Channel) CoopClose(auth *bind.TransactOpts) error {
	round, err := c.latestPeerSig()
	if err != nil {
		return err
	}
	tx, err := c.channel.CooperativeClose(auth, c.metadata.ID, round.valueA, round.valueB, round.peerSig)
	if err != nil {
		return err
	}

	return mustMine(c.backend, tx)
}

// CreateCoopClose creates the hash of a coop close which needs to be signed and send to a peer.
// Once a coop close is signed, the channel must not be used to send transactions anymore.
func (c *Channel) CreateCoopClose() ([32]byte, error) {
	hash, err := l2.HashState(c.channel, c.metadata.ID, c.metadata.A, c.metadata.B, c.valueA, c.valueB, l2.CoopCloseRound)
	if err != nil {
		return common.Hash{}, err
	}
	c.inactive = true
	return hash, nil
}

// StartForceClose starts the forceful closure procedure with the latest peer signature.
// It returns the time when the force close is successful.
func (c *Channel) StartForceClose(auth *bind.TransactOpts) (time.Time, error) {
	round, err := c.latestPeerSig()
	if err != nil {
		return time.Time{}, err
	}
	// Setup event watching
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	ch := make(chan *l2.ChannelClosing)
	c.channel.WatchClosing(&bind.WatchOpts{Context: ctx}, ch, [][32]byte{c.metadata.ID})
	// Call challenge
	tx, err := c.channel.Challenge(auth, c.metadata.ID, round.valueA, round.valueB, round.number, round.peerSig)
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
		if event.Round.Cmp(c.round) < 0 {
			return time.Time{}, &ChallengeError{round: event.Round}
		}
		return time.Unix(int64(event.Time), 0), nil
	case <-timer.C:
		return time.Time{}, errors.New("Waiting for event timed out")
	}
}

// DisputeForceClose disputes a force close procedure.
// It returns the time when the force close is finished.
func (c *Channel) DisputeForceClose(auth *bind.TransactOpts, start uint64) (time.Time, error) {
	// Setup event filtering
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	it, err := c.channel.FilterClosing(&bind.FilterOpts{Context: ctx, Start: start}, [][32]byte{c.metadata.ID})
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
	if newest.Round.Cmp(c.round) > 0 {
		return time.Time{}, fmt.Errorf("dispute has higher round than local, local: %v, onchain: %v", c.round, newest.Round)
	}

	// Call challenge
	round, err := c.latestPeerSig()
	if err != nil {
		return time.Time{}, err
	}
	fmt.Printf("Disputed with old state: %v got %v", newest.Round, round.number)
	// Setup event watching
	ctx, cancel = context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	ch := make(chan *l2.ChannelClosing)
	c.channel.WatchClosing(&bind.WatchOpts{Context: ctx}, ch, [][32]byte{c.metadata.ID})

	tx, err := c.channel.DisputeChallenge(auth, c.metadata.ID, round.valueA, round.valueB, round.number, round.peerSig)
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
		if event.Round.Cmp(c.round) < 0 {
			return time.Time{}, &ChallengeError{round: event.Round}
		}
		return time.Unix(int64(event.Time), 0), nil
	case <-timer.C:
		return time.Time{}, errors.New("waiting for event timed out")
	}
}

// FinishForceClose distributes the money after a force close is concluded.
// This can happen both after a StartForceClose or a DisputeForceClose.
func (c *Channel) FinishForceClose(auth *bind.TransactOpts) error {
	tx, err := c.channel.ForceClose(auth, c.metadata.ID)
	if err != nil {
		return err
	}
	return mustMine(c.backend, tx)
}

// SendMoney returns the hash of the state after the value is transfered.
// This hash needs to be signed and send to the peer.
func (c *Channel) SendMoney(value *big.Int) ([32]byte, error) {
	if c.inactive {
		return common.Hash{}, errors.New("cooperative close called, sending money now is impossible")
	}

	// TODO return mimedata for signing
	var (
		valA  *big.Int
		valB  *big.Int
		round *big.Int
	)
	if c.userIsProposer {
		if value.Cmp(c.valueA) > 0 {
			return common.Hash{}, errors.New("not enough funds")
		}
		valA = new(big.Int).Sub(c.valueA, value)
		valB = new(big.Int).Add(c.valueB, value)
		round = new(big.Int).Add(c.round, big.NewInt(1))

	} else {
		if value.Cmp(c.valueB) > 0 {
			return common.Hash{}, errors.New("not enough funds")
		}
		valA = new(big.Int).Add(c.valueA, value)
		valB = new(big.Int).Sub(c.valueB, value)
		round = new(big.Int).Add(c.round, big.NewInt(1))
	}
	hash, err := l2.HashState(c.channel, c.metadata.ID, c.metadata.A, c.metadata.B, valA, valB, round)
	if err != nil {
		return common.Hash{}, err
	}
	c.valueA = valA
	c.valueB = valB
	c.round = round
	return hash, nil
}

// ReceivedMoney updates the local channel state with the new values received from our peer.
func (c *Channel) ReceivedMoney(valueA, valueB, round *big.Int, sig []byte) error {
	if new(big.Int).Add(valueA, valueB).Cmp(c.sumFunds) != 0 {
		return fmt.Errorf("sum of funds not equal, want: %v got %v and %v", c.sumFunds, valueA, valueB)
	}
	if round.Cmp(c.round) <= 0 && c.round.Cmp(l2.CoopCloseRound) != 0 {
		return fmt.Errorf("invalid round: got %v local %v", round, c.round)
	}
	hash, err := l2.HashState(c.channel, c.metadata.ID, c.metadata.A, c.metadata.B, valueA, valueB, round)
	if err != nil {
		return err
	}
	var signer common.Address
	if c.userIsProposer {
		if c.valueA.Cmp(valueA) > 0 {
			return fmt.Errorf("valueA decreased by B, from %v to %v", c.valueA, valueA)
		}
		signer = c.metadata.B
	} else {
		if c.valueB.Cmp(valueB) > 0 {
			return fmt.Errorf("valueB decreased by A, from %v to %v", c.valueB, valueB)
		}
		signer = c.metadata.A
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
	c.round = round
	c.valueA = valueA
	c.valueB = valueB
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
	if c.round.Cmp(big.NewInt(2)) > 0 {
		min = new(big.Int).Sub(c.round, big.NewInt(2))
	}
	for i := c.round; i.Cmp(min) > 0; i.Sub(i, big.NewInt(1)) {
		if c.peerSigs[i] != nil {
			return c.peerSigs[i], nil
		}
	}
	return nil, fmt.Errorf("couldn't find proper signature within round %v", c.round)
}
