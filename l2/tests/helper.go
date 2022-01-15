package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	coopCloseRound, _ = new(big.Int).SetString("0xffffffffffffffffffffffffffffffff", 0)
)

type State struct {
	id     [32]byte
	a      common.Address
	b      common.Address
	valueA *big.Int
	valueB *big.Int
	round  *big.Int
}

func deployChannel(backend *backends.SimulatedBackend, sk *ecdsa.PrivateKey) (*Channel, common.Address, error) {
	transactor, _ := bind.NewKeyedTransactorWithChainID(sk, big.NewInt(1337))
	addr, tx, contract, err := DeployChannel(transactor, backend)
	if err != nil {
		return nil, common.Address{}, err
	}
	backend.Commit()
	if _, err = bind.WaitDeployed(context.Background(), backend, tx); err != nil {
		return nil, common.Address{}, err
	}
	return contract, addr, nil
}

func signChannelState(msg State, channel *Channel, sk *ecdsa.PrivateKey) ([32]byte, []byte, error) {
	state := ChannelChannelState{A: msg.a, B: msg.b, ValueA: common.Big0, ValueB: common.Big0, Progression: 0, Round: common.Big0}
	hash, err := channel.HashState(&bind.CallOpts{}, msg.id, state, msg.valueA, msg.valueB, msg.round)
	if err != nil {
		return [32]byte{}, []byte{}, err
	}
	sig, err := crypto.Sign(hash[:], sk)
	// The ECDSA library requires v to be 27 or 28, secp returns either 0 or 1
	sig[len(sig)-1] += 27
	return hash, sig, err
}
