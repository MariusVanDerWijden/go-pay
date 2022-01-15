package main

import (
	"encoding/gob"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	gopay "github.com/mariusvanderwijden/go-pay"
)

func setupChannel() error {
	promptChannel := promptui.Select{
		Label: "Choose one of the following options",
		Items: []string{"Create Channel (proposer)", "Accept Channel Creation (acceptor)", "Exit"},
	}
	for {
		index, res, err := promptChannel.Run()
		if err != nil {
			return err
		}
		switch index {
		case 0:
			err = createChannel()
		case 1:
			err = acceptCreateChannel()
		case 2:
			return nil
		}
		if err != nil {
			color.Red("Error during %v: %v", res, err)
		}
	}
}

func createChannel() error {
	prompt := promptui.Prompt{Label: "GoPay contract address"}
	str, err := prompt.Run()
	if err != nil {
		return err
	}
	gopay.ChannelAddr = common.HexToAddress(str)
	color.Yellow("Contract address: %v", gopay.ChannelAddr)
	// UserA
	prompt = promptui.Prompt{Label: "What's your address?"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	addrA := common.HexToAddress(str)
	prompt = promptui.Prompt{Label: "How much do you want to deposit?"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	valueA, ok := new(big.Int).SetString(str, 0)
	if !ok {
		return errors.New("could not understand value")
	}
	// UserB
	prompt = promptui.Prompt{Label: "What's the other persons address?"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	addrB := common.HexToAddress(str)
	prompt = promptui.Prompt{Label: "How much should they deposit?"}
	str, err = prompt.Run()
	if err != nil {
		return err
	}
	valueB, ok := new(big.Int).SetString(str, 0)
	if !ok {
		return errors.New("could not understand value")
	}
	// Double check and create
	if err := channelOverview(addrA, addrB, valueA, valueB, true); err != nil {
		return err
	}
	ch, err := gopay.NewChannel(backend, addrA, addrB, valueA, valueB)
	if err != nil {
		return err
	}
	channel = ch
	// Sending info to peer
	color.White("Sending opening message to peer")
	enc := gob.NewEncoder(peer)
	if err := enc.Encode(OpenChannelMsg{UserA: addrA, UserB: addrB, ValueA: valueA, ValueB: valueB}); err != nil {
		return err
	}
	return openChannel()
}

func acceptCreateChannel() error {
	color.White("Waiting for opening message from peer")
	dec := gob.NewDecoder(peer)
	var openMsg OpenChannelMsg
	if err := dec.Decode(&openMsg); err != nil {
		return err
	}
	if err := channelOverview(openMsg.UserA, openMsg.UserB, openMsg.ValueA, openMsg.ValueB, false); err != nil {
		return err
	}
	ch, err := gopay.NewChannel(backend, openMsg.UserA, openMsg.UserB, openMsg.ValueA, openMsg.ValueB)
	if err != nil {
		return err
	}
	channel = ch
	return acceptChannel()
}

func openChannel() error {
	color.White("Opening channel")
	txOpts := bind.NewClefTransactor(signer, accounts.Account{Address: channel.A})
	channelID, err := channel.Open(txOpts)
	if err != nil {
		return err
	}
	color.Cyan("Opened a new channel with channelID: %v", common.Bytes2Hex(channelID[:]))
	enc := gob.NewEncoder(peer)
	if err := enc.Encode(ChannelIDMsg{ID: channelID}); err != nil {
		return err
	}
	return useChannel()
}

func acceptChannel() error {
	color.White("Waiting for peer to open channel")
	dec := gob.NewDecoder(peer)
	var idMsg ChannelIDMsg
	if err := dec.Decode(&idMsg); err != nil {
		return err
	}
	txOpts := bind.NewClefTransactor(signer, accounts.Account{Address: channel.B})
	if err := channel.Accept(txOpts, idMsg.ID); err != nil {
		return err
	}
	color.Cyan("Accepted a new channel with channelID: %v", common.Bytes2Hex(idMsg.ID[:]))
	return useChannel()
}

func useChannel() error {
	promptChannel := promptui.Select{
		Label: "Choose one of the following options",
		Items: []string{"Create Channel (proposer)", "Accept Channel Creation (acceptor)", "Exit"},
	}
	for {
		index, res, err := promptChannel.Run()
		if err != nil {
			return err
		}
		switch index {
		case 0:
			err = send()
		case 1:
			err = closeChannel()
		case 2:
			return nil
		}
		if err != nil {
			color.Red("Error during %v: %v", res, err)
		}
	}
}

func send() error {
	return nil
}

func closeChannel() error {
	return nil
}

func channelOverview(addrA, addrB common.Address, valueA, valueB *big.Int, a bool) error {
	color.Cyan("Please double check the following values")
	if a {
		color.Cyan("(You should be A)")
	} else {
		color.Cyan("(You should be B)")
	}
	color.Cyan("---------------------")
	color.Cyan("AddrA: %v", addrA)
	color.Cyan("ValueA: %v", valueA)
	color.Cyan("AddrB: %v", addrB)
	color.Cyan("ValueB: %v", valueB)
	return promptOk()
}
