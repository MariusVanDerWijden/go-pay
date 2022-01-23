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

func (b *Backend) setupChannel() error {
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
			err = b.createChannel()
		case 1:
			err = b.acceptCreateChannel()
		case 2:
			return nil
		}
		if err != nil {
			color.Red("Error during %v: %v", res, err)
		}
	}
}

func (b *Backend) createChannel() error {
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
	ch, err := gopay.NewChannel(b.backend, addrA, addrB, valueA, valueB)
	if err != nil {
		return err
	}
	b.channel = ch
	// Sending info to peer
	color.White("Sending opening message to peer")
	enc := gob.NewEncoder(b.peer)
	if err := enc.Encode(OpenChannelMsg{UserA: addrA, UserB: addrB, ValueA: valueA, ValueB: valueB}); err != nil {
		return err
	}
	return b.openChannel()
}

func (b *Backend) acceptCreateChannel() error {
	color.White("Waiting for opening message from peer")
	dec := gob.NewDecoder(b.peer)
	var openMsg OpenChannelMsg
	if err := dec.Decode(&openMsg); err != nil {
		return err
	}
	if err := channelOverview(openMsg.UserA, openMsg.UserB, openMsg.ValueA, openMsg.ValueB, false); err != nil {
		return err
	}
	ch, err := gopay.NewChannel(b.backend, openMsg.UserA, openMsg.UserB, openMsg.ValueA, openMsg.ValueB)
	if err != nil {
		return err
	}
	b.channel = ch
	return b.acceptChannel()
}

func (b *Backend) openChannel() error {
	color.White("Opening channel")
	txOpts := bind.NewClefTransactor(b.signer, accounts.Account{Address: b.channel.A})
	channelID, err := b.channel.Open(txOpts)
	if err != nil {
		return err
	}
	color.Cyan("Opened a new channel with channelID: %v", common.Bytes2Hex(channelID[:]))
	enc := gob.NewEncoder(b.peer)
	if err := enc.Encode(ChannelIDMsg{ID: channelID}); err != nil {
		return err
	}
	return b.useChannel()
}

func (b *Backend) acceptChannel() error {
	color.White("Waiting for peer to open channel")
	dec := gob.NewDecoder(b.peer)
	var idMsg ChannelIDMsg
	if err := dec.Decode(&idMsg); err != nil {
		return err
	}
	txOpts := bind.NewClefTransactor(b.signer, accounts.Account{Address: b.channel.B})
	if err := b.channel.Accept(txOpts, idMsg.ID); err != nil {
		return err
	}
	color.Cyan("Accepted a new channel with channelID: %v", common.Bytes2Hex(idMsg.ID[:]))
	return b.useChannel()
}

func (b *Backend) useChannel() error {
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
			err = b.send()
		case 1:
			err = b.closeChannel()
		case 2:
			return nil
		}
		if err != nil {
			color.Red("Error during %v: %v", res, err)
		}
	}
}

func (b *Backend) send() error {
	prompt := promptui.Prompt{Label: "How much do you want to deposit?"}
	str, err := prompt.Run()
	if err != nil {
		return err
	}
	value, ok := new(big.Int).SetString(str, 0)
	if !ok {
		return errors.New("couldn't parse number")
	}
	hash, err := b.channel.SendMoney(value)
	if err != nil {
		return err
	}
	sig, err := b.signer.SignData(b.signer.Accounts()[0], "hash", hash[:])
	if err != nil {
		return err
	}
	color.Cyan("Sending signature to peer: %v", sig)
	signature := SignatureMsg{
		ValueA:    b.channel.ValueA,
		ValueB:    b.channel.ValueB,
		Round:     b.channel.Round,
		Signature: sig,
	}
	enc := gob.NewEncoder(b.peer)
	if err := enc.Encode(signature); err != nil {
		return err
	}
	return nil
}

func (b *Backend) receive() error {
	dec := gob.NewDecoder(b.peer)
	var sigMsg SignatureMsg
	if err := dec.Decode(&sigMsg); err != nil {
		return err
	}
	err := b.channel.ReceivedSignature(sigMsg.ValueA, sigMsg.ValueB, sigMsg.Round, sigMsg.Signature)
	if err != nil {
		return err
	}
	return nil
}

func (b *Backend) closeChannel() error {
	hash, err := b.channel.CreateCooperativeClose()
	if err != nil {
		return err
	}
	sig, err := b.signer.SignData(b.signer.Accounts()[0], "hash", hash[:])
	if err != nil {
		return err
	}
	color.Cyan("Sending closing signature to peer: %v", sig)
	signature := SignatureMsg{
		ValueA:    b.channel.ValueA,
		ValueB:    b.channel.ValueB,
		Round:     b.channel.Round,
		Signature: sig,
	}
	enc := gob.NewEncoder(b.peer)
	if err := enc.Encode(signature); err != nil {
		return err
	}
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
