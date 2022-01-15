package gopay

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mariusvanderwijden/go-pay/helpers"
	"github.com/mariusvanderwijden/go-pay/l2"
)

func setupEnv() (skA, skB *ecdsa.PrivateKey, addrA, addrB common.Address, bknd Backend) {
	backend, faucetSK := helpers.GetSimBackend()
	txopts, err := bind.NewKeyedTransactorWithChainID(faucetSK, big.NewInt(1337))
	if err != nil {
		panic(err)
	}
	addr, _, _, err := l2.DeployChannel(txopts, backend)
	if err != nil {
		panic(err)
	}
	skA, _ = crypto.GenerateKey()
	skB, _ = crypto.GenerateKey()
	addrA = crypto.PubkeyToAddress(skA.PublicKey)
	addrB = crypto.PubkeyToAddress(skB.PublicKey)
	helpers.Fund(backend, faucetSK, addrA)
	helpers.Fund(backend, faucetSK, addrB)
	ChannelAddr = addr
	return skA, skB, addrA, addrB, backend
}

func signHash(hash [32]byte, sk *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(hash[:], sk)
	// The ECDSA library requires v to be 27 or 28, secp returns either 0 or 1
	sig[len(sig)-1] += 27
	return sig, err
}

func TestWalkthrough(t *testing.T) {
	skA, skB, addrA, addrB, backend := setupEnv()
	fmt.Printf("A: %v\n", addrA)
	fmt.Printf("B: %v\n", addrB)
	txOptsA, err := bind.NewKeyedTransactorWithChainID(skA, big.NewInt(1337))
	if err != nil {
		panic(err)
	}
	txOptsB, err := bind.NewKeyedTransactorWithChainID(skB, big.NewInt(1337))
	if err != nil {
		panic(err)
	}
	valueA := big.NewInt(1000)
	valueB := big.NewInt(800)
	// Create a channel object for a
	channel, err := NewChannel(backend, addrA, addrB, valueA, valueB)
	if err != nil {
		t.Fatal(err)
	}
	// Create a channel object for b
	channelB, err := NewChannel(backend, addrA, addrB, valueA, valueB)
	if err != nil {
		t.Fatal(err)
	}
	// Open a channel
	id, err := channel.Open(txOptsA)
	if err != nil {
		t.Fatal(err)
	}
	// B confirms the channel
	if err := channelB.Accept(txOptsB, id); err != nil {
		t.Fatal(err)
	}
	// A sends some funds
	hashA, err := channel.SendMoney(big.NewInt(100))
	if err != nil {
		t.Fatal(err)
	}
	sigA, err := signHash(hashA, skA)
	if err != nil {
		t.Fatal(err)
	}
	// A sends valueA, valueB, round and sig to B
	// B inserts the values into its object
	if err := channelB.ReceivedSignature(channel.ValueA, channel.ValueB, channel.Round, sigA); err != nil {
		t.Fatal(err)
	}
	// B signs cooperative close
	hash, err := channelB.CreateCooperativeClose()
	if err != nil {
		t.Fatal(err)
	}
	sig, err := signHash(hash, skB)
	if err != nil {
		t.Fatal(err)
	}
	// B sends valueA, valueB, round and sig to A
	// A inserts the values into its object
	if err := channel.ReceivedSignature(channelB.ValueA, channelB.ValueB, l2.CoopCloseRound, sig); err != nil {
		t.Fatal(err)
	}
	// A sends coop close
	if err := channel.CoopClose(txOptsA); err != nil {
		t.Fatal(err)
	}
}
