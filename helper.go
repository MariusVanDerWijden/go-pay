package gopay

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetSimBackend() (*backends.SimulatedBackend, *ecdsa.PrivateKey) {
	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	faucetAddr := crypto.PubkeyToAddress(sk.PublicKey)
	addr := map[common.Address]core.GenesisAccount{
		common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
		common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
		common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
		common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
		common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
		common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
		common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
		common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
		faucetAddr:                       {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
	}
	alloc := core.GenesisAlloc(addr)
	return backends.NewSimulatedBackend(alloc, 80000000), sk
}

func Fund(backend *backends.SimulatedBackend, sk *ecdsa.PrivateKey, recipient common.Address) error {
	sender := crypto.PubkeyToAddress(sk.PublicKey)
	nonce, err := backend.PendingNonceAt(context.Background(), sender)
	if err != nil {
		panic(err)
	}
	gp := big.NewInt(1769259102)
	tx := types.NewTransaction(nonce, recipient, big.NewInt(100000000000000000), 21000, gp, nil)
	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1337)), sk)
	if err := backend.SendTransaction(context.Background(), signedTx); err != nil {
		return err
	}
	backend.Commit()
	_, err = bind.WaitMined(context.Background(), backend, signedTx)
	return err
}

func MustMineSuccessfully(backend *backends.SimulatedBackend, tx *types.Transaction) error {
	receipt, err := bind.WaitMined(context.Background(), backend, tx)
	if err != nil {
		return err
	}
	if receipt.Status == types.ReceiptStatusFailed {
		return errors.New("mined but failed")
	}
	return nil
}
