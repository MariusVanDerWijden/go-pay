package l2

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mariusvanderwijden/go-pay/helpers"
)

func setupChannel(t *testing.T) (skA, skB *ecdsa.PrivateKey, addrA, addrB common.Address, contract *Channel, backend *helpers.SimulatedBackend) {
	skA, skB, addrA, addrB, contract, backend = setupEnv()
	// A opens a channel
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)
	tx, err := contract.Open(transactorA, [32]byte{0}, addrB, big.NewInt(1000), big.NewInt(800))
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

func TestOpenInvalidValue(t *testing.T) {
	skA, _, _, addrB, contract, _ := setupEnv()
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(999)

	_, err := contract.Open(transactorA, [32]byte{0}, addrB, big.NewInt(1000), big.NewInt(800))
	if err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestOpenInvalidValue2(t *testing.T) {
	skA, _, _, addrB, contract, _ := setupEnv()
	// A opens a channel
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1001)

	_, err := contract.Open(transactorA, [32]byte{0}, addrB, big.NewInt(1000), big.NewInt(800))
	if err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestOpenChannelIDInUse(t *testing.T) {
	skA, _, _, addrB, contract, backend := setupEnv()
	// a opens a channels
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)
	tx, err := contract.Open(transactorA, [32]byte{0}, addrB, big.NewInt(1000), big.NewInt(800))
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

	_, err = contract.Open(transactorA, [32]byte{0}, addrB, big.NewInt(1000), big.NewInt(800))
	if err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func setupChannelAndAccept(t *testing.T) (skA, skB *ecdsa.PrivateKey, addrA, addrB common.Address, contract *Channel, backend *helpers.SimulatedBackend) {
	skA, skB, addrA, addrB, contract, backend = setupChannel(t)
	// B accepts the channel
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	tx, err := contract.Accept(transactorB, [32]byte{0})
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

func TestAcceptInvalidValue(t *testing.T) {
	_, skB, _, _, contract, _ := setupChannel(t)
	// B accepts the channel
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(799)
	if _, err := contract.Accept(transactorB, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidSender(t *testing.T) {
	skA, _, _, _, contract, _ := setupChannel(t)
	// a accepts the channel twice
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorA, [32]byte{0}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestAcceptInvalidChannel(t *testing.T) {
	_, skB, _, _, contract, _ := setupChannel(t)
	// B accepts the channel
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, [32]byte{1}); err == nil {
		t.Fatal("Expected error, got no error")
	}
}

func TestCooperativeCloseA(t *testing.T) {
	skA, skB, addrA, addrB, contract, backend := setupChannelAndAccept(t)
	state := State{
		id:     [32]byte{0},
		a:      addrA,
		b:      addrB,
		valueA: big.NewInt(1100),
		valueB: big.NewInt(700),
		round:  CoopCloseRound,
	}
	// A tries to cooperative close with their own sk
	_, sigA, err := SignChannelState(state, contract, skA)
	if err != nil {
		t.Fatal(err)
	}
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	_, err = contract.CooperativeClose(transactorA, [32]byte{0}, state.valueA, state.valueB, sigA)
	if err == nil {
		t.Fatal("expected error")
	}

	_, sigB, err := SignChannelState(state, contract, skB)
	if err != nil {
		t.Fatal(err)
	}
	tx, err := contract.CooperativeClose(transactorA, [32]byte{0}, state.valueA, state.valueB, sigB)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()
	if err := helpers.MustMineSuccessfully(backend, tx); err != nil {
		t.Fatal(err)
	}
}

func TestCooperativeCloseB(t *testing.T) {
	skA, skB, addrA, addrB, contract, backend := setupChannelAndAccept(t)
	state := State{
		id:     [32]byte{0},
		a:      addrA,
		b:      addrB,
		valueA: big.NewInt(1100),
		valueB: big.NewInt(700),
		round:  CoopCloseRound,
	}
	// B tries to cooperative close with their own sk
	_, sigB, err := SignChannelState(state, contract, skB)
	if err != nil {
		t.Fatal(err)
	}
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	_, err = contract.CooperativeClose(transactorB, [32]byte{0}, state.valueA, state.valueB, sigB)
	if err == nil {
		t.Fatal("expected error")
	}

	_, sigA, err := SignChannelState(state, contract, skA)
	if err != nil {
		t.Fatal(err)
	}
	tx, err := contract.CooperativeClose(transactorB, [32]byte{0}, state.valueA, state.valueB, sigA)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()
	if err := helpers.MustMineSuccessfully(backend, tx); err != nil {
		t.Fatal(err)
	}
}
