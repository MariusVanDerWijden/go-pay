package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Channels struct {
	ID [][32]byte
}

type Channel struct {
	ID           [32]byte
	States       map[*big.Int]State
	AddrA        common.Address
	AddrB        common.Address
	CurrentRound *big.Int
}

type LightChannel struct {
	ID           [32]byte
	AddrA        common.Address
	AddrB        common.Address
	CurrentState *big.Int
	ValueA       *big.Int
	ValueB       *big.Int
}

type State struct {
	ValueA *big.Int
	ValueB *big.Int
}

type UpdateChannel struct {
	ID     [32]byte
	AddrA  common.Address
	AddrB  common.Address
	Round  *big.Int
	ValueA *big.Int
	ValueB *big.Int
}
