package main

import (
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	gopay "github.com/mariusvanderwijden/go-pay"
)

type Config struct {
	BlockchainNode  string
	ExternalSigner  string
	ContractAddress string
}

func readConfig(path string) Config {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic(err)
	}
	return conf
}

func setupNodeFromConfig(b *Backend, config string) error {
	color.Green("Reading config from: %v", config)
	c := readConfig(config)
	// setup contract
	gopay.ChannelAddr = common.HexToAddress(c.ContractAddress)
	// setup blockchain
	client, err := ethclient.Dial(c.BlockchainNode)
	if err != nil {
		return err
	}
	b.backend = client
	// setup signer
	sig, err := external.NewExternalSigner(c.ExternalSigner)
	if err != nil {
		return err
	}
	accs := sig.Accounts()
	color.Green("Accounts: ")
	for i, acc := range accs {
		color.Green("%v: %v", i, acc.Address)
	}
	b.signer = sig
	return nil
}
