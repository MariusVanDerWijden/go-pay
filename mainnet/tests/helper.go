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

func signChannelState(msg ChannelChannelState, channel Channel, sk *ecdsa.PrivateKey) ([32]byte, []byte, error) {
	hash, err := channel.HashState(&bind.CallOpts{}, msg)
	if err != nil {
		return [32]byte{}, []byte{}, err
	}
	sig, err := crypto.Sign(hash[:], sk)
	// The ECDSA library requires v to be 27 or 28, secp returns either 0 or 1
	sig[len(sig)-1] += 27
	return hash, sig, err
}
