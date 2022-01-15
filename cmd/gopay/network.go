package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OpenChannelMsg struct {
	UserA  common.Address
	UserB  common.Address
	ValueA *big.Int
	ValueB *big.Int
}

type ChannelIDMsg struct {
	ID [32]byte
}
