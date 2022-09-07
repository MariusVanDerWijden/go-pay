package main

import "github.com/BurntSushi/toml"

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
