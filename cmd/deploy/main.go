package main

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/mariusvanderwijden/go-pay/l2"
)

var (
	backend *ethclient.Client
	signer  *external.ExternalSigner
)

func main() {
	color.Green("Welcome to go-pay deployment tool")
	promptInit := promptui.Select{
		Label: "Choose one of the following options",
		Items: []string{"Connect to Blockchain", "Setup External Signer (clef)", "Deploy Contract"},
	}

	for {
		index, _, err := promptInit.Run()
		if err != nil {
			color.Red("Goodbye: %v", err)
			return
		}
		switch index {
		case 0:
			err = connectToBackend()
		case 1:
			err = setupSigner()
		case 2:
			err = deploy()
		}
		if err != nil {
			color.Red("Error: %v", err)
		}
	}
}

func connectToBackend() error {
	prompt := promptui.Prompt{Label: "Blockchain node address"}
	str, err := prompt.Run()
	if err != nil {
		return err
	}
	client, err := ethclient.Dial(str)
	if err != nil {
		return err
	}
	backend = client
	return nil
}

func setupSigner() error {
	prompt := promptui.Prompt{Label: "External Signer Endpoint"}
	str, err := prompt.Run()
	if err != nil {
		return err
	}
	sig, err := external.NewExternalSigner(str)
	if err != nil {
		return err
	}
	signer = sig
	return nil
}

func deploy() error {
	auth := bind.NewClefTransactor(signer, signer.Accounts()[1])
	address, tx, _, err := l2.DeployChannel(auth, backend)
	if err != nil {
		return err
	}
	_, err = bind.WaitDeployed(context.Background(), backend, tx)
	if err != nil {
		return err
	}
	color.Red("Contract deployed at: %v", address)
	return nil
}
