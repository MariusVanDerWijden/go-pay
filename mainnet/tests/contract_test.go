package tests

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mariusvanderwijden/go-pay/helpers"
)

func setupChannel(t *testing.T) (skA, skB *ecdsa.PrivateKey, addrA, addrB common.Address, contract *Channel, backend *helpers.SimulatedBackend, openMsg ChannelChannelState) {
	skA, skB, addrA, addrB, contract, backend = setupEnv()
	// A opens a channel
	openMsg = ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 0,
		Round:       big.NewInt(0),
	}
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)
	tx, err := contract.Open(transactorA, openMsg, [32]byte{0})
	if err != nil {
		t.Error()
	}
	backend.Commit()
	if err := helpers.MustMineSuccessfully(backend, tx); err != nil {
		t.Fatal(err)
	}
	return
}

func TestOpen(t *testing.T) {
	setupChannel(t)
}

func TestOpenInvalidSender(t *testing.T) {
	skA, _, addrA, addrB, contract, _ := setupEnv()
	// A opens a channel
	openMsg := ChannelChannelState{
		A:           addrB,
		B:           addrA,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 0,
		Round:       big.NewInt(0),
	}
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)

	if _, err := contract.Open(transactorA, openMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestOpenInvalidValue(t *testing.T) {
	skA, _, addrA, addrB, contract, _ := setupEnv()
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
	transactorA.Value = big.NewInt(999)

	if _, err := contract.Open(transactorA, openMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestOpenInvalidValue2(t *testing.T) {
	skA, _, addrA, addrB, contract, _ := setupEnv()
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
	transactorA.Value = big.NewInt(1001)

	if _, err := contract.Open(transactorA, openMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestOpenChannelIDInUse(t *testing.T) {
	skA, _, addrA, addrB, contract, backend := setupEnv()
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
	tx, err := contract.Open(transactorA, openMsg, [32]byte{0})
	if err != nil {
		t.Error()
	}
	backend.Commit()
	if err := helpers.MustMineSuccessfully(backend, tx); err != nil {
		t.Fatal(err)
	}
	// try to open a channel on the same channelid
	transactorA, _ = bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)

	if _, err := contract.Open(transactorA, openMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestOpenProgressionSet(t *testing.T) {
	skA, _, addrA, addrB, contract, _ := setupEnv()
	// A opens a channel
	openMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 1,
		Round:       big.NewInt(0),
	}
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)

	if _, err := contract.Open(transactorA, openMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestOpenRoundSet(t *testing.T) {
	skA, _, addrA, addrB, contract, _ := setupEnv()
	// A opens a channel
	openMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 0,
		Round:       big.NewInt(1),
	}
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)

	if _, err := contract.Open(transactorA, openMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func setupChannelAndAccept(t *testing.T) (skA, skB *ecdsa.PrivateKey, addrA, addrB common.Address, contract *Channel, backend *helpers.SimulatedBackend, acceptMsg ChannelChannelState) {
	_, skB, addrA, addrB, contract, backend, openMsg := setupChannel(t)
	// B accepts the channel
	acceptMsg = ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 1,
		Round:       big.NewInt(0),
	}
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	tx, err := contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0})
	if err != nil {
		t.Error(err)
	}
	backend.Commit()
	if err := helpers.MustMineSuccessfully(backend, tx); err != nil {
		t.Fatal(err)
	}
	return
}

func TestAccept(t *testing.T) {
	setupChannelAndAccept(t)
}

func TestAcceptInvalidNewState(t *testing.T) {
	_, skB, _, addrB, contract, _, openMsg := setupChannel(t)
	// B accepts the channel
	acceptMsg := ChannelChannelState{
		A:           addrB,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 1,
		Round:       big.NewInt(0),
	}
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidProgression(t *testing.T) {
	_, skB, addrA, addrB, contract, _, openMsg := setupChannel(t)
	// B accepts the channel
	acceptMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 2,
		Round:       big.NewInt(0),
	}
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidRound(t *testing.T) {
	_, skB, addrA, addrB, contract, _, openMsg := setupChannel(t)
	// B accepts the channel
	acceptMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(1000),
		ValueB:      big.NewInt(800),
		Progression: 1,
		Round:       big.NewInt(1),
	}
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidOldstate(t *testing.T) {
	_, skB, addrA, addrB, contract, _, openMsg := setupChannel(t)
	openMsg.Progression = 1
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
	if _, err := contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidNewState2(t *testing.T) {
	_, skB, _, _, contract, _, openMsg := setupChannel(t)
	openMsg.Progression = 1
	// B accepts the channel with openMsg
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, openMsg, openMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidValue(t *testing.T) {
	_, skB, addrA, addrB, contract, _, openMsg := setupChannel(t)
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
	transactorB.Value = big.NewInt(799)
	if _, err := contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidOldValue(t *testing.T) {
	_, skB, addrA, addrB, contract, _, openMsg := setupChannel(t)
	// B accepts the channel
	acceptMsg := ChannelChannelState{
		A:           addrA,
		B:           addrB,
		ValueA:      big.NewInt(100),
		ValueB:      big.NewInt(800),
		Progression: 1,
		Round:       big.NewInt(0),
	}
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, openMsg, acceptMsg, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}
