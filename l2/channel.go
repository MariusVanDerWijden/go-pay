// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChannelChannelState is an auto generated low-level Go binding around an user-defined struct.
type ChannelChannelState struct {
	A           common.Address
	B           common.Address
	ValueA      *big.Int
	ValueB      *big.Int
	Progression uint8
	Round       *big.Int
}

// ChannelMetaData contains all meta data concerning the Channel contract.
var ChannelMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Accepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"}],\"name\":\"Closing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Open\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"accept\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Prog\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"cooperativeClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"disputeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"forceClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Prog\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"name\":\"hashState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addrB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x6080604052620151806000806101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506fffffffffffffffffffffffffffffffff600060086101000a8154816fffffffffffffffffffffffffffffffff02191690836fffffffffffffffffffffffffffffffff16021790555034801561008457600080fd5b5061279f806100946000396000f3fe6080604052600436106100865760003560e01c806331b5d50d1161005957806331b5d50d146101445780637a7ebd7b14610181578063a7293956146101c3578063a72d6a48146101ec578063e4725ba11461020857610086565b8063098d419d1461008b57806311be1997146100b4578063267656cc146100f2578063271d30ca1461011b575b600080fd5b34801561009757600080fd5b506100b260048036038101906100ad9190611c68565b610224565b005b3480156100c057600080fd5b506100db60048036038101906100d69190611b56565b610592565b6040516100e9929190612217565b60405180910390f35b3480156100fe57600080fd5b5061011960048036038101906101149190611b56565b6105ea565b005b34801561012757600080fd5b50610142600480360381019061013d9190611ceb565b6107ec565b005b34801561015057600080fd5b5061016b60048036038101906101669190611bea565b610d0d565b6040516101789190611fcd565b60405180910390f35b34801561018d57600080fd5b506101a860048036038101906101a39190611b56565b610d58565b6040516101ba96959493929190611f6c565b60405180910390f35b3480156101cf57600080fd5b506101ea60048036038101906101e59190611ceb565b610dfd565b005b61020660048036038101906102019190611b83565b611221565b005b610222600480360381019061021d9190611b56565b611442565b005b83600280600481111561023a57610239612452565b5b6001600084815260200190815260200160002060040160009054906101000a900460ff1660048111156102705761026f612452565b5b146102b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a79061212e565b60405180910390fd5b85858580826102bf91906122a7565b600160008581526020019081526020016000206003015460016000868152602001908152602001600020600201546102f791906122a7565b14610337576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032e906120ce565b60405180910390fd5b60006103438a3361163b565b905060006104d48b600160008e81526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff16600481111561044e5761044d612452565b5b60048111156104605761045f612452565b5b81526020016004820160019054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815250508c8c600060089054906101000a90046fffffffffffffffffffffffffffffffff16610d0d565b90508173ffffffffffffffffffffffffffffffffffffffff166104f7828a6117d5565b73ffffffffffffffffffffffffffffffffffffffff161461054d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610544906120ae565b60405180910390fd5b6105588b8b8b6118c1565b8a7f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d7960405160405180910390a25050505050505050505050565b60026020528060005260406000206000915090508060000160009054906101000a900467ffffffffffffffff16908060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905082565b806003806004811115610600576105ff612452565b5b6001600084815260200190815260200160002060040160009054906101000a900460ff16600481111561063657610635612452565b5b14610676576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066d9061212e565b60405180910390fd5b60006002600085815260200190815260200160002060000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1614156106ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e59061214e565b60405180910390fd5b4260008054906101000a900467ffffffffffffffff166002600086815260200190815260200160002060000160009054906101000a900467ffffffffffffffff1661073991906122fd565b67ffffffffffffffff1610610783576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077a906121ce565b60405180910390fd5b6107ba83600160008681526020019081526020016000206002015460016000878152602001908152602001600020600301546118c1565b827f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d7960405160405180910390a2505050565b84600380600481111561080257610801612452565b5b6001600084815260200190815260200160002060040160009054906101000a900460ff16600481111561083857610837612452565b5b14610878576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086f9061212e565b60405180910390fd5b868686808261088791906122a7565b600160008581526020019081526020016000206003015460016000868152602001908152602001600020600201546108bf91906122a7565b146108ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f6906120ce565b60405180910390fd5b866fffffffffffffffffffffffffffffffff16600160008c815260200190815260200160002060040160019054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff1610610997576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098e906121ae565b60405180910390fd5b60006109a38b3361163b565b90508073ffffffffffffffffffffffffffffffffffffffff16600260008d815260200190815260200160002060000160089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a40906120ee565b60405180910390fd5b6000610bba8c600160008f81526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff166004811115610b5257610b51612452565b5b6004811115610b6457610b63612452565b5b81526020016004820160019054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815250508d8d8d610d0d565b90508173ffffffffffffffffffffffffffffffffffffffff16610bdd828a6117d5565b73ffffffffffffffffffffffffffffffffffffffff1614610c33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2a906120ae565b60405180910390fd5b33600260008e815260200190815260200160002060000160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600260008e815260200190815260200160002060000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508b7f6e129f552c838f9ab74444d21bd821c7694edfcbbd3c2294dbc7bc48290401808a42604051610cf79291906121ee565b60405180910390a2505050505050505050505050565b6000808686600001518760200151878787604051602001610d3396959493929190611fe8565b6040516020818303038152906040529050808051906020012091505095945050505050565b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030154908060040160009054906101000a900460ff16908060040160019054906101000a90046fffffffffffffffffffffffffffffffff16905086565b846002806004811115610e1357610e12612452565b5b6001600084815260200190815260200160002060040160009054906101000a900460ff166004811115610e4957610e48612452565b5b14610e89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e809061212e565b60405180910390fd5b8686868082610e9891906122a7565b60016000858152602001908152602001600020600301546001600086815260200190815260200160002060020154610ed091906122a7565b14610f10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f07906120ce565b60405180910390fd5b6000610f1c8b3361163b565b9050600061108f8c600160008f81526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff16600481111561102757611026612452565b5b600481111561103957611038612452565b5b81526020016004820160019054906101000a90046fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff166fffffffffffffffffffffffffffffffff16815250508d8d8d610d0d565b90508173ffffffffffffffffffffffffffffffffffffffff166110b2828a6117d5565b73ffffffffffffffffffffffffffffffffffffffff1614611108576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110ff906120ae565b60405180910390fd5b6003600160008e815260200190815260200160002060040160006101000a81548160ff0219169083600481111561114257611141612452565b5b021790555033600260008e815260200190815260200160002060000160086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555042600260008e815260200190815260200160002060000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508b7f6e129f552c838f9ab74444d21bd821c7694edfcbbd3c2294dbc7bc48290401808a4260405161120b9291906121ee565b60405180910390a2505050505050505050505050565b83600080600481111561123757611236612452565b5b6001600084815260200190815260200160002060040160009054906101000a900460ff16600481111561126d5761126c612452565b5b146112ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a49061212e565b60405180910390fd5b8334146112ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112e69061208e565b60405180910390fd5b336001600088815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846001600088815260200190815260200160002060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836001600088815260200190815260200160002060020181905550826001600088815260200190815260200160002060030181905550600180600088815260200190815260200160002060040160006101000a81548160ff0219169083600481111561140857611407612452565b5b0217905550857fddffb592d6434d02d388cf2eb4fbfa796fbfcd09e278d3466e7194dfd3c23a0460405160405180910390a2505050505050565b80600180600481111561145857611457612452565b5b6001600084815260200190815260200160002060040160009054906101000a900460ff16600481111561148e5761148d612452565b5b146114ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c59061212e565b60405180910390fd5b6001600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611572576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115699061216e565b60405180910390fd5b600160008481526020019081526020016000206003015434146115ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115c19061210e565b60405180910390fd5b60026001600085815260200190815260200160002060040160006101000a81548160ff0219169083600481111561160457611603612452565b5b0217905550827f1f3c0697c3ada95f9e84a917995664b76cd8b4ae5de25e77ee111122ae3a00d060405160405180910390a2505050565b60006001600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156116e5576001600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506117cf565b6001600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561178d576001600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506117cf565b60006117ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117c59061218e565b60405180910390fd5b5b92915050565b600060418251146117e957600090506118bb565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c111561183d57600093505050506118bb565b601b8160ff16141580156118555750601c8160ff1614155b1561186657600093505050506118bb565b600186828585604051600081526020016040526040516118899493929190612049565b6020604051602081039080840390855afa1580156118ab573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60046001600085815260200190815260200160002060040160006101000a81548160ff021916908360048111156118fb576118fa612452565b5b02179055506001600084815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050506001600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505050505050565b60006119f06119eb84612265565b612240565b905082815260208101848484011115611a0c57611a0b6124ba565b5b611a178482856123e3565b509392505050565b600081359050611a2e816126fd565b92915050565b600081359050611a4381612714565b92915050565b600082601f830112611a5e57611a5d6124b0565b5b8135611a6e8482602086016119dd565b91505092915050565b600081359050611a868161272b565b92915050565b600060c08284031215611aa257611aa16124b5565b5b611aac60c0612240565b90506000611abc84828501611a1f565b6000830152506020611ad084828501611a1f565b6020830152506040611ae484828501611b41565b6040830152506060611af884828501611b41565b6060830152506080611b0c84828501611a77565b60808301525060a0611b2084828501611b2c565b60a08301525092915050565b600081359050611b3b8161273b565b92915050565b600081359050611b5081612752565b92915050565b600060208284031215611b6c57611b6b6124c4565b5b6000611b7a84828501611a34565b91505092915050565b60008060008060808587031215611b9d57611b9c6124c4565b5b6000611bab87828801611a34565b9450506020611bbc87828801611a1f565b9350506040611bcd87828801611b41565b9250506060611bde87828801611b41565b91505092959194509250565b60008060008060006101408688031215611c0757611c066124c4565b5b6000611c1588828901611a34565b9550506020611c2688828901611a8c565b94505060e0611c3788828901611b41565b935050610100611c4988828901611b41565b925050610120611c5b88828901611b2c565b9150509295509295909350565b60008060008060808587031215611c8257611c816124c4565b5b6000611c9087828801611a34565b9450506020611ca187828801611b41565b9350506040611cb287828801611b41565b925050606085013567ffffffffffffffff811115611cd357611cd26124bf565b5b611cdf87828801611a49565b91505092959194509250565b600080600080600060a08688031215611d0757611d066124c4565b5b6000611d1588828901611a34565b9550506020611d2688828901611b41565b9450506040611d3788828901611b41565b9350506060611d4888828901611b2c565b925050608086013567ffffffffffffffff811115611d6957611d686124bf565b5b611d7588828901611a49565b9150509295509295909350565b611d8b8161233b565b82525050565b611d9a8161234d565b82525050565b611da9816123d1565b82525050565b6000611dbc601e83612296565b9150611dc7826124da565b602082019050919050565b6000611ddf601183612296565b9150611dea82612503565b602082019050919050565b6000611e02601683612296565b9150611e0d8261252c565b602082019050919050565b6000611e25601a83612296565b9150611e3082612555565b602082019050919050565b6000611e48601e83612296565b9150611e538261257e565b602082019050919050565b6000611e6b601b83612296565b9150611e76826125a7565b602082019050919050565b6000611e8e601483612296565b9150611e99826125d0565b602082019050919050565b6000611eb1601783612296565b9150611ebc826125f9565b602082019050919050565b6000611ed4600d83612296565b9150611edf82612622565b602082019050919050565b6000611ef7602183612296565b9150611f028261264b565b604082019050919050565b6000611f1a602183612296565b9150611f258261269a565b604082019050919050565b611f398161236a565b82525050565b611f48816123a6565b82525050565b611f57816123b0565b82525050565b611f66816123c4565b82525050565b600060c082019050611f816000830189611d82565b611f8e6020830188611d82565b611f9b6040830187611f3f565b611fa86060830186611f3f565b611fb56080830185611da0565b611fc260a0830184611f30565b979650505050505050565b6000602082019050611fe26000830184611d91565b92915050565b600060c082019050611ffd6000830189611d91565b61200a6020830188611d82565b6120176040830187611d82565b6120246060830186611f3f565b6120316080830185611f3f565b61203e60a0830184611f30565b979650505050505050565b600060808201905061205e6000830187611d91565b61206b6020830186611f5d565b6120786040830185611d91565b6120856060830184611d91565b95945050505050565b600060208201905081810360008301526120a781611daf565b9050919050565b600060208201905081810360008301526120c781611dd2565b9050919050565b600060208201905081810360008301526120e781611df5565b9050919050565b6000602082019050818103600083015261210781611e18565b9050919050565b6000602082019050818103600083015261212781611e3b565b9050919050565b6000602082019050818103600083015261214781611e5e565b9050919050565b6000602082019050818103600083015261216781611e81565b9050919050565b6000602082019050818103600083015261218781611ea4565b9050919050565b600060208201905081810360008301526121a781611ec7565b9050919050565b600060208201905081810360008301526121c781611eea565b9050919050565b600060208201905081810360008301526121e781611f0d565b9050919050565b60006040820190506122036000830185611f30565b6122106020830184611f4e565b9392505050565b600060408201905061222c6000830185611f4e565b6122396020830184611d82565b9392505050565b600061224a61225b565b905061225682826123f2565b919050565b6000604051905090565b600067ffffffffffffffff8211156122805761227f612481565b5b612289826124c9565b9050602081019050919050565b600082825260208201905092915050565b60006122b2826123a6565b91506122bd836123a6565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156122f2576122f1612423565b5b828201905092915050565b6000612308826123b0565b9150612313836123b0565b92508267ffffffffffffffff038211156123305761232f612423565b5b828201905092915050565b600061234682612386565b9050919050565b6000819050919050565b6000819050612365826126e9565b919050565b60006fffffffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60006123dc82612357565b9050919050565b82818337600083830152505050565b6123fb826124c9565b810181811067ffffffffffffffff8211171561241a57612419612481565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f70726f706f736572206d7573742061646420656e6f7567682076616c75650000600082015250565b7f696e76616c6964207369676e6174757265000000000000000000000000000000600082015250565b7f696e76616c69642073746174652070726f766964656400000000000000000000600082015250565b7f696e76616c69642064697370757465206368616c6c656e676572000000000000600082015250565b7f6163636570746f72206d7573742061646420656e6f7567682076616c75650000600082015250565b7f4368616e6e656c20697320696e20696e76616c69642073746174650000000000600082015250565b7f6e6f206469737075746520617661696c61626c65000000000000000000000000600082015250565b7f6163636570746f72206d7573742062652073656e646572000000000000000000600082015250565b7f696e76616c6964206f7468657200000000000000000000000000000000000000600082015250565b7f6469737075746564206368616c6c656e67652077697468206f6c64207374617460008201527f6500000000000000000000000000000000000000000000000000000000000000602082015250565b7f666f72636520636c6f7365206265666f7265206469737075746520706572696f60008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b600581106126fa576126f9612452565b5b50565b6127068161233b565b811461271157600080fd5b50565b61271d8161234d565b811461272857600080fd5b50565b6005811061273857600080fd5b50565b6127448161236a565b811461274f57600080fd5b50565b61275b816123a6565b811461276657600080fd5b5056fea26469706673582212203a7716e6aaa5ade19d1300e290b3eba8a93119b3041f3ea9c7c32ab873c18e4c64736f6c63430008070033",
}

// ChannelABI is the input ABI used to generate the binding from.
// Deprecated: Use ChannelMetaData.ABI instead.
var ChannelABI = ChannelMetaData.ABI

// ChannelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChannelMetaData.Bin instead.
var ChannelBin = ChannelMetaData.Bin

// DeployChannel deploys a new Ethereum contract, binding an instance of Channel to it.
func DeployChannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Channel, error) {
	parsed, err := ChannelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChannelBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// Channel is an auto generated Go binding around an Ethereum contract.
type Channel struct {
	ChannelCaller     // Read-only binding to the contract
	ChannelTransactor // Write-only binding to the contract
	ChannelFilterer   // Log filterer for contract events
}

// ChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChannelSession struct {
	Contract     *Channel          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChannelCallerSession struct {
	Contract *ChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChannelTransactorSession struct {
	Contract     *ChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChannelRaw struct {
	Contract *Channel // Generic contract binding to access the raw methods on
}

// ChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChannelCallerRaw struct {
	Contract *ChannelCaller // Generic read-only contract binding to access the raw methods on
}

// ChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChannelTransactorRaw struct {
	Contract *ChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChannel creates a new instance of Channel, bound to a specific deployed contract.
func NewChannel(address common.Address, backend bind.ContractBackend) (*Channel, error) {
	contract, err := bindChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Channel{ChannelCaller: ChannelCaller{contract: contract}, ChannelTransactor: ChannelTransactor{contract: contract}, ChannelFilterer: ChannelFilterer{contract: contract}}, nil
}

// NewChannelCaller creates a new read-only instance of Channel, bound to a specific deployed contract.
func NewChannelCaller(address common.Address, caller bind.ContractCaller) (*ChannelCaller, error) {
	contract, err := bindChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelCaller{contract: contract}, nil
}

// NewChannelTransactor creates a new write-only instance of Channel, bound to a specific deployed contract.
func NewChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*ChannelTransactor, error) {
	contract, err := bindChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChannelTransactor{contract: contract}, nil
}

// NewChannelFilterer creates a new log filterer instance of Channel, bound to a specific deployed contract.
func NewChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*ChannelFilterer, error) {
	contract, err := bindChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChannelFilterer{contract: contract}, nil
}

// bindChannel binds a generic wrapper to an already deployed contract.
func bindChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.ChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.ChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Channel *ChannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Channel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Channel *ChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Channel *ChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Channel.Contract.contract.Transact(opts, method, params...)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address a, address b, uint256 valueA, uint256 valueB, uint8 progression, uint128 round)
func (_Channel *ChannelCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) (struct {
	A           common.Address
	B           common.Address
	ValueA      *big.Int
	ValueB      *big.Int
	Progression uint8
	Round       *big.Int
}, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "channels", arg0)

	outstruct := new(struct {
		A           common.Address
		B           common.Address
		ValueA      *big.Int
		ValueB      *big.Int
		Progression uint8
		Round       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.B = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ValueA = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValueB = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Progression = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Round = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address a, address b, uint256 valueA, uint256 valueB, uint8 progression, uint128 round)
func (_Channel *ChannelSession) Channels(arg0 [32]byte) (struct {
	A           common.Address
	B           common.Address
	ValueA      *big.Int
	ValueB      *big.Int
	Progression uint8
	Round       *big.Int
}, error) {
	return _Channel.Contract.Channels(&_Channel.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(address a, address b, uint256 valueA, uint256 valueB, uint8 progression, uint128 round)
func (_Channel *ChannelCallerSession) Channels(arg0 [32]byte) (struct {
	A           common.Address
	B           common.Address
	ValueA      *big.Int
	ValueB      *big.Int
	Progression uint8
	Round       *big.Int
}, error) {
	return _Channel.Contract.Channels(&_Channel.CallOpts, arg0)
}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 time, address closer)
func (_Channel *ChannelCaller) Disputes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Time   uint64
	Closer common.Address
}, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "disputes", arg0)

	outstruct := new(struct {
		Time   uint64
		Closer common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Time = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Closer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 time, address closer)
func (_Channel *ChannelSession) Disputes(arg0 [32]byte) (struct {
	Time   uint64
	Closer common.Address
}, error) {
	return _Channel.Contract.Disputes(&_Channel.CallOpts, arg0)
}

// Disputes is a free data retrieval call binding the contract method 0x11be1997.
//
// Solidity: function disputes(bytes32 ) view returns(uint64 time, address closer)
func (_Channel *ChannelCallerSession) Disputes(arg0 [32]byte) (struct {
	Time   uint64
	Closer common.Address
}, error) {
	return _Channel.Contract.Disputes(&_Channel.CallOpts, arg0)
}

// HashState is a free data retrieval call binding the contract method 0x31b5d50d.
//
// Solidity: function hashState(bytes32 id, (address,address,uint256,uint256,uint8,uint128) state, uint256 valueA, uint256 valueB, uint128 round) pure returns(bytes32)
func (_Channel *ChannelCaller) HashState(opts *bind.CallOpts, id [32]byte, state ChannelChannelState, valueA *big.Int, valueB *big.Int, round *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "hashState", id, state, valueA, valueB, round)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashState is a free data retrieval call binding the contract method 0x31b5d50d.
//
// Solidity: function hashState(bytes32 id, (address,address,uint256,uint256,uint8,uint128) state, uint256 valueA, uint256 valueB, uint128 round) pure returns(bytes32)
func (_Channel *ChannelSession) HashState(id [32]byte, state ChannelChannelState, valueA *big.Int, valueB *big.Int, round *big.Int) ([32]byte, error) {
	return _Channel.Contract.HashState(&_Channel.CallOpts, id, state, valueA, valueB, round)
}

// HashState is a free data retrieval call binding the contract method 0x31b5d50d.
//
// Solidity: function hashState(bytes32 id, (address,address,uint256,uint256,uint8,uint128) state, uint256 valueA, uint256 valueB, uint128 round) pure returns(bytes32)
func (_Channel *ChannelCallerSession) HashState(id [32]byte, state ChannelChannelState, valueA *big.Int, valueB *big.Int, round *big.Int) ([32]byte, error) {
	return _Channel.Contract.HashState(&_Channel.CallOpts, id, state, valueA, valueB, round)
}

// Accept is a paid mutator transaction binding the contract method 0xe4725ba1.
//
// Solidity: function accept(bytes32 id) payable returns()
func (_Channel *ChannelTransactor) Accept(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "accept", id)
}

// Accept is a paid mutator transaction binding the contract method 0xe4725ba1.
//
// Solidity: function accept(bytes32 id) payable returns()
func (_Channel *ChannelSession) Accept(id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Accept(&_Channel.TransactOpts, id)
}

// Accept is a paid mutator transaction binding the contract method 0xe4725ba1.
//
// Solidity: function accept(bytes32 id) payable returns()
func (_Channel *ChannelTransactorSession) Accept(id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Accept(&_Channel.TransactOpts, id)
}

// Challenge is a paid mutator transaction binding the contract method 0xa7293956.
//
// Solidity: function challenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes sig) returns()
func (_Channel *ChannelTransactor) Challenge(opts *bind.TransactOpts, id [32]byte, valueA *big.Int, valueB *big.Int, round *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "challenge", id, valueA, valueB, round, sig)
}

// Challenge is a paid mutator transaction binding the contract method 0xa7293956.
//
// Solidity: function challenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes sig) returns()
func (_Channel *ChannelSession) Challenge(id [32]byte, valueA *big.Int, valueB *big.Int, round *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.Challenge(&_Channel.TransactOpts, id, valueA, valueB, round, sig)
}

// Challenge is a paid mutator transaction binding the contract method 0xa7293956.
//
// Solidity: function challenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes sig) returns()
func (_Channel *ChannelTransactorSession) Challenge(id [32]byte, valueA *big.Int, valueB *big.Int, round *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.Challenge(&_Channel.TransactOpts, id, valueA, valueB, round, sig)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0x098d419d.
//
// Solidity: function cooperativeClose(bytes32 id, uint256 valueA, uint256 valueB, bytes sig) returns()
func (_Channel *ChannelTransactor) CooperativeClose(opts *bind.TransactOpts, id [32]byte, valueA *big.Int, valueB *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "cooperativeClose", id, valueA, valueB, sig)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0x098d419d.
//
// Solidity: function cooperativeClose(bytes32 id, uint256 valueA, uint256 valueB, bytes sig) returns()
func (_Channel *ChannelSession) CooperativeClose(id [32]byte, valueA *big.Int, valueB *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.CooperativeClose(&_Channel.TransactOpts, id, valueA, valueB, sig)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0x098d419d.
//
// Solidity: function cooperativeClose(bytes32 id, uint256 valueA, uint256 valueB, bytes sig) returns()
func (_Channel *ChannelTransactorSession) CooperativeClose(id [32]byte, valueA *big.Int, valueB *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.CooperativeClose(&_Channel.TransactOpts, id, valueA, valueB, sig)
}

// DisputeChallenge is a paid mutator transaction binding the contract method 0x271d30ca.
//
// Solidity: function disputeChallenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes sig) returns()
func (_Channel *ChannelTransactor) DisputeChallenge(opts *bind.TransactOpts, id [32]byte, valueA *big.Int, valueB *big.Int, round *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "disputeChallenge", id, valueA, valueB, round, sig)
}

// DisputeChallenge is a paid mutator transaction binding the contract method 0x271d30ca.
//
// Solidity: function disputeChallenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes sig) returns()
func (_Channel *ChannelSession) DisputeChallenge(id [32]byte, valueA *big.Int, valueB *big.Int, round *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.DisputeChallenge(&_Channel.TransactOpts, id, valueA, valueB, round, sig)
}

// DisputeChallenge is a paid mutator transaction binding the contract method 0x271d30ca.
//
// Solidity: function disputeChallenge(bytes32 id, uint256 valueA, uint256 valueB, uint128 round, bytes sig) returns()
func (_Channel *ChannelTransactorSession) DisputeChallenge(id [32]byte, valueA *big.Int, valueB *big.Int, round *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.DisputeChallenge(&_Channel.TransactOpts, id, valueA, valueB, round, sig)
}

// ForceClose is a paid mutator transaction binding the contract method 0x267656cc.
//
// Solidity: function forceClose(bytes32 id) returns()
func (_Channel *ChannelTransactor) ForceClose(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "forceClose", id)
}

// ForceClose is a paid mutator transaction binding the contract method 0x267656cc.
//
// Solidity: function forceClose(bytes32 id) returns()
func (_Channel *ChannelSession) ForceClose(id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.ForceClose(&_Channel.TransactOpts, id)
}

// ForceClose is a paid mutator transaction binding the contract method 0x267656cc.
//
// Solidity: function forceClose(bytes32 id) returns()
func (_Channel *ChannelTransactorSession) ForceClose(id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.ForceClose(&_Channel.TransactOpts, id)
}

// Open is a paid mutator transaction binding the contract method 0xa72d6a48.
//
// Solidity: function open(bytes32 id, address addrB, uint256 valueA, uint256 valueB) payable returns()
func (_Channel *ChannelTransactor) Open(opts *bind.TransactOpts, id [32]byte, addrB common.Address, valueA *big.Int, valueB *big.Int) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "open", id, addrB, valueA, valueB)
}

// Open is a paid mutator transaction binding the contract method 0xa72d6a48.
//
// Solidity: function open(bytes32 id, address addrB, uint256 valueA, uint256 valueB) payable returns()
func (_Channel *ChannelSession) Open(id [32]byte, addrB common.Address, valueA *big.Int, valueB *big.Int) (*types.Transaction, error) {
	return _Channel.Contract.Open(&_Channel.TransactOpts, id, addrB, valueA, valueB)
}

// Open is a paid mutator transaction binding the contract method 0xa72d6a48.
//
// Solidity: function open(bytes32 id, address addrB, uint256 valueA, uint256 valueB) payable returns()
func (_Channel *ChannelTransactorSession) Open(id [32]byte, addrB common.Address, valueA *big.Int, valueB *big.Int) (*types.Transaction, error) {
	return _Channel.Contract.Open(&_Channel.TransactOpts, id, addrB, valueA, valueB)
}

// ChannelAcceptedIterator is returned from FilterAccepted and is used to iterate over the raw logs and unpacked data for Accepted events raised by the Channel contract.
type ChannelAcceptedIterator struct {
	Event *ChannelAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChannelAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChannelAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChannelAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelAccepted represents a Accepted event raised by the Channel contract.
type ChannelAccepted struct {
	ID  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccepted is a free log retrieval operation binding the contract event 0x1f3c0697c3ada95f9e84a917995664b76cd8b4ae5de25e77ee111122ae3a00d0.
//
// Solidity: event Accepted(bytes32 indexed ID)
func (_Channel *ChannelFilterer) FilterAccepted(opts *bind.FilterOpts, ID [][32]byte) (*ChannelAcceptedIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.FilterLogs(opts, "Accepted", IDRule)
	if err != nil {
		return nil, err
	}
	return &ChannelAcceptedIterator{contract: _Channel.contract, event: "Accepted", logs: logs, sub: sub}, nil
}

// WatchAccepted is a free log subscription operation binding the contract event 0x1f3c0697c3ada95f9e84a917995664b76cd8b4ae5de25e77ee111122ae3a00d0.
//
// Solidity: event Accepted(bytes32 indexed ID)
func (_Channel *ChannelFilterer) WatchAccepted(opts *bind.WatchOpts, sink chan<- *ChannelAccepted, ID [][32]byte) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.WatchLogs(opts, "Accepted", IDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelAccepted)
				if err := _Channel.contract.UnpackLog(event, "Accepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccepted is a log parse operation binding the contract event 0x1f3c0697c3ada95f9e84a917995664b76cd8b4ae5de25e77ee111122ae3a00d0.
//
// Solidity: event Accepted(bytes32 indexed ID)
func (_Channel *ChannelFilterer) ParseAccepted(log types.Log) (*ChannelAccepted, error) {
	event := new(ChannelAccepted)
	if err := _Channel.contract.UnpackLog(event, "Accepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChannelClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the Channel contract.
type ChannelClosedIterator struct {
	Event *ChannelClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChannelClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChannelClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChannelClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelClosed represents a Closed event raised by the Channel contract.
type ChannelClosed struct {
	ID  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79.
//
// Solidity: event Closed(bytes32 indexed ID)
func (_Channel *ChannelFilterer) FilterClosed(opts *bind.FilterOpts, ID [][32]byte) (*ChannelClosedIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.FilterLogs(opts, "Closed", IDRule)
	if err != nil {
		return nil, err
	}
	return &ChannelClosedIterator{contract: _Channel.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79.
//
// Solidity: event Closed(bytes32 indexed ID)
func (_Channel *ChannelFilterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *ChannelClosed, ID [][32]byte) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.WatchLogs(opts, "Closed", IDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelClosed)
				if err := _Channel.contract.UnpackLog(event, "Closed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosed is a log parse operation binding the contract event 0x7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d79.
//
// Solidity: event Closed(bytes32 indexed ID)
func (_Channel *ChannelFilterer) ParseClosed(log types.Log) (*ChannelClosed, error) {
	event := new(ChannelClosed)
	if err := _Channel.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChannelClosingIterator is returned from FilterClosing and is used to iterate over the raw logs and unpacked data for Closing events raised by the Channel contract.
type ChannelClosingIterator struct {
	Event *ChannelClosing // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChannelClosingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelClosing)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChannelClosing)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChannelClosingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelClosingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelClosing represents a Closing event raised by the Channel contract.
type ChannelClosing struct {
	ID    [32]byte
	Round *big.Int
	Time  uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterClosing is a free log retrieval operation binding the contract event 0x6e129f552c838f9ab74444d21bd821c7694edfcbbd3c2294dbc7bc4829040180.
//
// Solidity: event Closing(bytes32 indexed ID, uint128 round, uint64 time)
func (_Channel *ChannelFilterer) FilterClosing(opts *bind.FilterOpts, ID [][32]byte) (*ChannelClosingIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.FilterLogs(opts, "Closing", IDRule)
	if err != nil {
		return nil, err
	}
	return &ChannelClosingIterator{contract: _Channel.contract, event: "Closing", logs: logs, sub: sub}, nil
}

// WatchClosing is a free log subscription operation binding the contract event 0x6e129f552c838f9ab74444d21bd821c7694edfcbbd3c2294dbc7bc4829040180.
//
// Solidity: event Closing(bytes32 indexed ID, uint128 round, uint64 time)
func (_Channel *ChannelFilterer) WatchClosing(opts *bind.WatchOpts, sink chan<- *ChannelClosing, ID [][32]byte) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.WatchLogs(opts, "Closing", IDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelClosing)
				if err := _Channel.contract.UnpackLog(event, "Closing", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClosing is a log parse operation binding the contract event 0x6e129f552c838f9ab74444d21bd821c7694edfcbbd3c2294dbc7bc4829040180.
//
// Solidity: event Closing(bytes32 indexed ID, uint128 round, uint64 time)
func (_Channel *ChannelFilterer) ParseClosing(log types.Log) (*ChannelClosing, error) {
	event := new(ChannelClosing)
	if err := _Channel.contract.UnpackLog(event, "Closing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChannelOpenIterator is returned from FilterOpen and is used to iterate over the raw logs and unpacked data for Open events raised by the Channel contract.
type ChannelOpenIterator struct {
	Event *ChannelOpen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChannelOpenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChannelOpen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChannelOpen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChannelOpenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChannelOpenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChannelOpen represents a Open event raised by the Channel contract.
type ChannelOpen struct {
	ID  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOpen is a free log retrieval operation binding the contract event 0xddffb592d6434d02d388cf2eb4fbfa796fbfcd09e278d3466e7194dfd3c23a04.
//
// Solidity: event Open(bytes32 indexed ID)
func (_Channel *ChannelFilterer) FilterOpen(opts *bind.FilterOpts, ID [][32]byte) (*ChannelOpenIterator, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.FilterLogs(opts, "Open", IDRule)
	if err != nil {
		return nil, err
	}
	return &ChannelOpenIterator{contract: _Channel.contract, event: "Open", logs: logs, sub: sub}, nil
}

// WatchOpen is a free log subscription operation binding the contract event 0xddffb592d6434d02d388cf2eb4fbfa796fbfcd09e278d3466e7194dfd3c23a04.
//
// Solidity: event Open(bytes32 indexed ID)
func (_Channel *ChannelFilterer) WatchOpen(opts *bind.WatchOpts, sink chan<- *ChannelOpen, ID [][32]byte) (event.Subscription, error) {

	var IDRule []interface{}
	for _, IDItem := range ID {
		IDRule = append(IDRule, IDItem)
	}

	logs, sub, err := _Channel.contract.WatchLogs(opts, "Open", IDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChannelOpen)
				if err := _Channel.contract.UnpackLog(event, "Open", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOpen is a log parse operation binding the contract event 0xddffb592d6434d02d388cf2eb4fbfa796fbfcd09e278d3466e7194dfd3c23a04.
//
// Solidity: event Open(bytes32 indexed ID)
func (_Channel *ChannelFilterer) ParseOpen(log types.Log) (*ChannelOpen, error) {
	event := new(ChannelOpen)
	if err := _Channel.contract.UnpackLog(event, "Open", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
