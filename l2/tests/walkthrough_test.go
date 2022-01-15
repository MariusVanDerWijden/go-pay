package tests

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	gopay "github.com/mariusvanderwijden/go-pay"
)

func setupEnv() (skA, skB *ecdsa.PrivateKey, addrA, addrB common.Address, contract *Channel, backend *backends.SimulatedBackend) {
	backend, faucetSK := gopay.GetSimBackend()
	contract, _, err := deployChannel(backend, faucetSK)
	if err != nil {
		panic(err)
	}
	skA, _ = crypto.GenerateKey()
	skB, _ = crypto.GenerateKey()
	addrA = crypto.PubkeyToAddress(skA.PublicKey)
	addrB = crypto.PubkeyToAddress(skB.PublicKey)
	gopay.Fund(backend, faucetSK, addrA)
	gopay.Fund(backend, faucetSK, addrB)
	return
}

// A -> open
// B -> open
// A -> cooperative close
func TestWalkthrough(t *testing.T) {
	skA, skB, addrA, addrB, contract, backend := setupEnv()
	channelID := [32]byte{0}
	// A opens a channel
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)
	if _, err := contract.Open(transactorA, channelID, addrB, big.NewInt(1000), big.NewInt(800)); err != nil {
		t.Error()
	}
	backend.Commit()

	// B accepts the channel
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, channelID); err != nil {
		t.Error()
	}
	backend.Commit()

	// A cooperative closes the channel
	ccState := State{
		id:     channelID,
		a:      addrA,
		b:      addrB,
		valueA: big.NewInt(700),
		valueB: big.NewInt(1100),
		round:  coopCloseRound,
	}
	transactorA, _ = bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	_, signature, err := signChannelState(ccState, contract, skB)
	if err != nil {
		t.Error(err)
	}
	tx, err := contract.CooperativeClose(transactorA, channelID, ccState.valueA, ccState.valueB, signature)
	if err != nil {
		t.Error(err)
	}
	backend.Commit()
	if err = gopay.MustMineSuccessfully(backend, tx); err != nil {
		t.Error(err)
	}
}

// A -> open
// B -> open
// A -> dispute
// A -> forceclose
func TestForceClose(t *testing.T) {
	skA, skB, addrA, addrB, contract, backend := setupEnv()
	channelID := [32]byte{0}
	// A opens a channel
	transactorA, _ := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	transactorA.Value = big.NewInt(1000)
	if _, err := contract.Open(transactorA, channelID, addrB, big.NewInt(1000), big.NewInt(800)); err != nil {
		t.Error()
	}
	backend.Commit()

	// B accepts the channel
	transactorB, _ := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	transactorB.Value = big.NewInt(800)
	if _, err := contract.Accept(transactorB, channelID); err != nil {
		t.Error()
	}
	backend.Commit()

	// A cooperative closes the channel
	ccState := State{
		id:     channelID,
		a:      addrA,
		b:      addrB,
		valueA: big.NewInt(700),
		valueB: big.NewInt(1100),
		round:  big.NewInt(100),
	}
	transactorA, _ = bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	_, signature, err := signChannelState(ccState, contract, skB)
	if err != nil {
		t.Error(err)
	}
	tx, err := contract.Challenge(transactorA, channelID, ccState.valueA, ccState.valueB, ccState.round, signature)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()
	if err = gopay.MustMineSuccessfully(backend, tx); err != nil {
		t.Error(err)
	}
	backend.AdjustTime(25 * time.Hour)
	backend.Commit()
	// A force closes the channel
	transactorA, _ = bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	tx, err = contract.ForceClose(transactorA, channelID)
	if err != nil {
		t.Fatal(err)
	}
	backend.Commit()
	if err = gopay.MustMineSuccessfully(backend, tx); err != nil {
		t.Fatal(err)
	}
}
