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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Accepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"}],\"name\":\"Closing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Open\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"CooperativeClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ForceClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"accept\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Prog\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"disputeChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Prog\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"name\":\"hashState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addrB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1b98eb84": "CooperativeClose(bytes32,uint256,uint256,bytes)",
		"5e8a4616": "ForceClose(bytes32)",
		"e4725ba1": "accept(bytes32)",
		"a7293956": "challenge(bytes32,uint256,uint256,uint128,bytes)",
		"7a7ebd7b": "channels(bytes32)",
		"271d30ca": "disputeChallenge(bytes32,uint256,uint256,uint128,bytes)",
		"11be1997": "disputes(bytes32)",
		"31b5d50d": "hashState(bytes32,(address,address,uint256,uint256,uint8,uint128),uint256,uint256,uint128)",
		"a72d6a48": "open(bytes32,address,uint256,uint256)",
	},
	Bin: "0x6080604052600080546001600160c01b03191667fffffffffffeae80600160c01b0317905534801561003057600080fd5b5061146c806100406000396000f3fe6080604052600436106100865760003560e01c80635e8a4616116100595780635e8a4616146101715780637a7ebd7b14610191578063a72939561461020d578063a72d6a481461022d578063e4725ba11461024057600080fd5b806311be19971461008b5780631b98eb8414610101578063271d30ca1461012357806331b5d50d14610143575b600080fd5b34801561009757600080fd5b506100d46100a63660046110ba565b60026020526000908152604090205467ffffffffffffffff811690600160401b90046001600160a01b031682565b6040805167ffffffffffffffff90931683526001600160a01b039091166020830152015b60405180910390f35b34801561010d57600080fd5b5061012161011c3660046111cc565b610253565b005b34801561012f57600080fd5b5061012161013e366004611226565b610439565b34801561014f57600080fd5b5061016361015e36600461110e565b610732565b6040519081526020016100f8565b34801561017d57600080fd5b5061012161018c3660046110ba565b610796565b34801561019d57600080fd5b506101fb6101ac3660046110ba565b6001602081905260009182526040909120805491810154600282015460038301546004909301546001600160a01b039485169490921692909160ff81169061010090046001600160801b031686565b6040516100f896959493929190611291565b34801561021957600080fd5b50610121610228366004611226565b61091a565b61012161023b3660046110d3565b610b8e565b61012161024e3660046110ba565b610ca8565b83600280600083815260016020526040902060049081015460ff169081111561027e5761027e61140a565b146102a45760405162461bcd60e51b815260040161029b90611350565b60405180910390fd5b8585856102b181836113b0565b600084815260016020526040902060038101546002909101546102d491906113b0565b146102f15760405162461bcd60e51b815260040161029b90611320565b60006102fd8a33610e02565b60008b8152600160208181526040808420815160c08101835281546001600160a01b03908116825294820154909416928401929092526002820154908301526003810154606083015260048082015494955092936103ba938f939291608084019160ff909116908111156103735761037361140a565b60048111156103845761038461140a565b8152600491909101546001600160801b0361010090910481166020909201919091526000548e918e91600160401b900416610732565b9050816001600160a01b03166103d0828a610e90565b6001600160a01b0316146103f65760405162461bcd60e51b815260040161029b906112f5565b6104018b8b8b610f7a565b6040518b907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d7990600090a25050505050505050505050565b84600380600083815260016020526040902060049081015460ff16908111156104645761046461140a565b146104815760405162461bcd60e51b815260040161029b90611350565b86868661048e81836113b0565b600084815260016020526040902060038101546002909101546104b191906113b0565b146104ce5760405162461bcd60e51b815260040161029b90611320565b60006104da8b33610e02565b90506001600160a01b0381166105225760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b21037ba3432b960991b604482015260640161029b565b60008b8152600260205260409020546001600160a01b03828116600160401b90920416146105925760405162461bcd60e51b815260206004820152601a60248201527f696e76616c69642064697370757465206368616c6c656e676572000000000000604482015260640161029b565b60006106828c600160008f81526020019081526020016000206040518060c00160405290816000820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016001820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff16600481111561064c5761064c61140a565b600481111561065d5761065d61140a565b81526004919091015461010090046001600160801b03166020909101528d8d8d610732565b9050816001600160a01b0316610698828a610e90565b6001600160a01b0316146106be5760405162461bcd60e51b815260040161029b906112f5565b60008c9052600160205260405162461bcd60e51b815260040161029b9060208082526021908201527f6469737075746564206368616c6c656e67652077697468206f6c6420737461746040820152606560f81b606082015260800190565b60405180910390a2505050505050505050505050565b835160209485015160408051808801989098526001600160a01b039283168882015291166060870152608086019390935260a08501919091526001600160801b031660c0808501919091528151808503909101815260e09093019052815191012090565b80600380600083815260016020526040902060049081015460ff16908111156107c1576107c161140a565b146107de5760405162461bcd60e51b815260040161029b90611350565b60008381526002602052604090205467ffffffffffffffff1661083a5760405162461bcd60e51b81526020600482015260146024820152736e6f206469737075746520617661696c61626c6560601b604482015260640161029b565b60008054848252600260205260409091205442916108659167ffffffffffffffff91821691166113c8565b67ffffffffffffffff16106108c65760405162461bcd60e51b815260206004820152602160248201527f666f72636520636c6f7365206265666f7265206469737075746520706572696f6044820152601960fa1b606482015260840161029b565b600083815260016020526040902060028101546003909101546108ea918591610f7a565b60405183907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d7990600090a2505050565b84600280600083815260016020526040902060049081015460ff16908111156109455761094561140a565b146109625760405162461bcd60e51b815260040161029b90611350565b86868661096f81836113b0565b6000848152600160205260409020600381015460029091015461099291906113b0565b146109af5760405162461bcd60e51b815260040161029b90611320565b60006109bb8b33610e02565b90506001600160a01b038116610a035760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b21037ba3432b960991b604482015260640161029b565b6000610abd8c600160008f81526020019081526020016000206040518060c00160405290816000820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016001820160009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160028201548152602001600382015481526020016004820160009054906101000a900460ff16600481111561064c5761064c61140a565b9050816001600160a01b0316610ad3828a610e90565b6001600160a01b031614610af95760405162461bcd60e51b815260040161029b906112f5565b60008c81526001602090815260408083206004018054600360ff19909116179055600282529182902080546001600160e01b03191633600160401b0267ffffffffffffffff1916174267ffffffffffffffff1690811790915582516001600160801b038d168152918201528d917f6e129f552c838f9ab74444d21bd821c7694edfcbbd3c2294dbc7bc4829040180910161071c565b83600080600083815260016020526040902060049081015460ff1690811115610bb957610bb961140a565b14610bd65760405162461bcd60e51b815260040161029b90611350565b833414610c255760405162461bcd60e51b815260206004820152601e60248201527f70726f706f736572206d7573742061646420656e6f7567682076616c75650000604482015260640161029b565b60008681526001602081905260408083208054336001600160a01b031991821617825581840180549091166001600160a01b038b161790556002810188905560038101879055600401805460ff19169092179091555187917fddffb592d6434d02d388cf2eb4fbfa796fbfcd09e278d3466e7194dfd3c23a0491a2505050505050565b80600180600083815260016020526040902060049081015460ff1690811115610cd357610cd361140a565b14610cf05760405162461bcd60e51b815260040161029b90611350565b600083815260016020819052604090912001546001600160a01b03163314610d5a5760405162461bcd60e51b815260206004820152601760248201527f6163636570746f72206d7573742062652073656e646572000000000000000000604482015260640161029b565b6000838152600160205260409020600301543414610dba5760405162461bcd60e51b815260206004820152601e60248201527f6163636570746f72206d7573742061646420656e6f7567682076616c75650000604482015260640161029b565b600083815260016020526040808220600401805460ff191660021790555184917f1f3c0697c3ada95f9e84a917995664b76cd8b4ae5de25e77ee111122ae3a00d091a2505050565b6000828152600160205260408120546001600160a01b0383811691161415610e465750600082815260016020819052604090912001546001600160a01b0316610e8a565b6000838152600160205260409020546001600160a01b0383811691161415610e8657506000828152600160205260409020546001600160a01b0316610e8a565b5060005b92915050565b60008151604114610ea357506000610e8a565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115610ee95760009350505050610e8a565b8060ff16601b14158015610f0157508060ff16601c14155b15610f125760009350505050610e8a565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa158015610f65573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b6000838152600160205260408082206004808201805460ff191690911790555490516001600160a01b03909116916108fc851502918591818181858888f150505060008581526001602081905260408083209091015490516001600160a01b03909116935084156108fc0292508491818181858888f15050505050505050565b80356001600160a01b038116811461101157600080fd5b919050565b600082601f83011261102757600080fd5b813567ffffffffffffffff8082111561104257611042611420565b604051601f8301601f19908116603f0116810190828211818310171561106a5761106a611420565b8160405283815286602085880101111561108357600080fd5b836020870160208301376000602085830101528094505050505092915050565b80356001600160801b038116811461101157600080fd5b6000602082840312156110cc57600080fd5b5035919050565b600080600080608085870312156110e957600080fd5b843593506110f960208601610ffa565b93969395505050506040820135916060013590565b600080600080600085870361014081121561112857600080fd5b8635955060c0601f198201121561113e57600080fd5b50611147611387565b61115360208801610ffa565b815261116160408801610ffa565b6020820152606087013560408201526080870135606082015260a08701356005811061118c57600080fd5b608082015261119d60c088016110a3565b60a0820152935060e0860135925061010086013591506111c061012087016110a3565b90509295509295909350565b600080600080608085870312156111e257600080fd5b843593506020850135925060408501359150606085013567ffffffffffffffff81111561120e57600080fd5b61121a87828801611016565b91505092959194509250565b600080600080600060a0868803121561123e57600080fd5b85359450602086013593506040860135925061125c606087016110a3565b9150608086013567ffffffffffffffff81111561127857600080fd5b61128488828901611016565b9150509295509295909350565b6001600160a01b03878116825286166020820152604081018590526060810184905260c08101600584106112d557634e487b7160e01b600052602160045260246000fd5b8360808301526001600160801b03831660a0830152979650505050505050565b602080825260119082015270696e76616c6964207369676e617475726560781b604082015260600190565b6020808252601690820152751a5b9d985b1a59081cdd185d19481c1c9bdd9a59195960521b604082015260600190565b6020808252601b908201527f4368616e6e656c20697320696e20696e76616c69642073746174650000000000604082015260600190565b60405160c0810167ffffffffffffffff811182821017156113aa576113aa611420565b60405290565b600082198211156113c3576113c36113f4565b500190565b600067ffffffffffffffff8083168185168083038211156113eb576113eb6113f4565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea264697066735822122096dea7f3dd88d934a534058f23e31c52e2e807491d486a9a9e9ddc07627b564864736f6c63430008070033",
}

// ChannelABI is the input ABI used to generate the binding from.
// Deprecated: Use ChannelMetaData.ABI instead.
var ChannelABI = ChannelMetaData.ABI

// Deprecated: Use ChannelMetaData.Sigs instead.
// ChannelFuncSigs maps the 4-byte function signature to its string representation.
var ChannelFuncSigs = ChannelMetaData.Sigs

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

// CooperativeClose is a paid mutator transaction binding the contract method 0x1b98eb84.
//
// Solidity: function CooperativeClose(bytes32 id, uint256 valueA, uint256 valueB, bytes sig) returns()
func (_Channel *ChannelTransactor) CooperativeClose(opts *bind.TransactOpts, id [32]byte, valueA *big.Int, valueB *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "CooperativeClose", id, valueA, valueB, sig)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0x1b98eb84.
//
// Solidity: function CooperativeClose(bytes32 id, uint256 valueA, uint256 valueB, bytes sig) returns()
func (_Channel *ChannelSession) CooperativeClose(id [32]byte, valueA *big.Int, valueB *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.CooperativeClose(&_Channel.TransactOpts, id, valueA, valueB, sig)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0x1b98eb84.
//
// Solidity: function CooperativeClose(bytes32 id, uint256 valueA, uint256 valueB, bytes sig) returns()
func (_Channel *ChannelTransactorSession) CooperativeClose(id [32]byte, valueA *big.Int, valueB *big.Int, sig []byte) (*types.Transaction, error) {
	return _Channel.Contract.CooperativeClose(&_Channel.TransactOpts, id, valueA, valueB, sig)
}

// ForceClose is a paid mutator transaction binding the contract method 0x5e8a4616.
//
// Solidity: function ForceClose(bytes32 id) returns()
func (_Channel *ChannelTransactor) ForceClose(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "ForceClose", id)
}

// ForceClose is a paid mutator transaction binding the contract method 0x5e8a4616.
//
// Solidity: function ForceClose(bytes32 id) returns()
func (_Channel *ChannelSession) ForceClose(id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.ForceClose(&_Channel.TransactOpts, id)
}

// ForceClose is a paid mutator transaction binding the contract method 0x5e8a4616.
//
// Solidity: function ForceClose(bytes32 id) returns()
func (_Channel *ChannelTransactorSession) ForceClose(id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.ForceClose(&_Channel.TransactOpts, id)
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

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122076bb7702b1fef6aab5b13e275ad1021f6c07bb98b9c5a57af8bf89dc585960a164736f6c63430008070033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}
