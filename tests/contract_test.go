package tests

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestWalkthrough(t *testing.T) {
	backend, faucetSK := getSimBackend()
	contract, _, err := deployChannel(backend, faucetSK)
	if err != nil {
		t.Error(err)
	}
	skA, _ := crypto.GenerateKey()
	skB, _ := crypto.GenerateKey()
	addrA := crypto.PubkeyToAddress(skA.PublicKey)
	addrB := crypto.PubkeyToAddress(skB.PublicKey)
	fund(backend, faucetSK, addrA)
	fund(backend, faucetSK, addrB)

	// A opens a channel
	openMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 0,
		Round:       big.NewInt(0),
	}
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)
	if _, err = contract.Open(transactorA, openMsg, [32]byte{0}); err != nil {
		t.Error()
	}
	backend.Commit()

	// B accepts the channel
	acceptMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 1,
		Round:       big.NewInt(0),
	}
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err = contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0}); err != nil {
		t.Error()
	}
	backend.Commit()

	// A cooperative closes the channel
	cooperativeCloseMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(700),
		ValueB:      big.NewInt(1100),
		Progression: 2,
		Round:       big.NewInt(0),
	}
	transactorA, _ = bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	_, signature, err := signChannelState(cooperativeCloseMsg, *contract, skB)
	if err != nil {
		t.Error(err)
	}
	tx, err := contract.CooperativeClose(transactorA, acceptMsg, cooperativeCloseMsg, signature, [32]byte{0})
	if err != nil {
		t.Error(err)
	}
	backend.Commit()
	if err = mustMineSuccessfully(backend, tx); err != nil {
		t.Error(err)
	}
}
