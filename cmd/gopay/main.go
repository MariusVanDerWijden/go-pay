package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/ethereum/go-ethereum/accounts/external"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	gopay "github.com/mariusvanderwijden/go-pay"
)

var (
	LocalHost = "127.0.0.1"
	Port      = 9580
	peer      net.Conn
	backend   *ethclient.Client
	channel   *gopay.Channel
	signer    *external.ExternalSigner
)

func main() {
	color.Green("Welcome to go-pay")
	promptInit := promptui.Select{
		Label: "Choose one of the following options",
		Items: []string{"Start Listener", "Connect to Peer", "Connect to Blockchain", "Setup External Signer (clef)", "Setup Channel"},
	}

	for {
		index, _, err := promptInit.Run()
		if err != nil {
			color.Red("Goodbye: %v", err)
			return
		}
		err = mainSwitch(index)
		if err != nil {
			color.Red("Error: %v", err)
		}
	}
}

func mainSwitch(index int) error {
	switch index {
	case 0:
		return startListener()
	case 1:
		return connectToPeer()
	case 2:
		return connectToBackend()
	case 3:
		return setupExternalSigner()
	case 4:
		return setupChannel()
	}
	return nil
}

func validateIP(input string) error {
	ip := net.ParseIP(input)
	if ip == nil {
		return errors.New("invalid ip")
	}
	return nil
}

func connectToPeer() error {
	prompt := promptui.Prompt{
		Label:    "IP",
		Validate: validateIP,
	}
	str, err := prompt.Run()
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", str, Port))
	if err != nil {
		return err
	}
	color.Green("Connected to peer: %v", conn.RemoteAddr())
	peer = conn
	return nil
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

func setupExternalSigner() error {
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

func startListener() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", LocalHost, Port))
	if err != nil {
		return err
	}
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			color.Red("Listening failed %v\n", err)
			return
		}
		color.Green("Peer connected: %v", conn.RemoteAddr())
		peer = conn
	}()
	return nil
}

func promptOk() error {
	promptChannel := promptui.Select{
		Label: "Should we proceed",
		Items: []string{"No", "Yes"},
	}
	i, _, err := promptChannel.Run()
	if err != nil || i != 1 {
		return errors.New("we won't proceed")
	}
	return nil
}
