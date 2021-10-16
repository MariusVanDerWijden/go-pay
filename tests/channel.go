// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tests

import (
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

// ChannelABI is the input ABI used to generate the binding from.
const ChannelABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Accepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"}],\"name\":\"Closing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"ID\",\"type\":\"bytes32\"}],\"name\":\"Open\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"oldState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"newState\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"CooperativeClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"oldState\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ForceClose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"oldState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"newState\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"accept\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"oldState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"newState\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"challenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"channels\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"oldState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"newState\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"disputeChallenge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"disputes\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"time\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"state\",\"type\":\"tuple\"}],\"name\":\"hashState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"enumChannel.Progression\",\"name\":\"progression\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"}],\"internalType\":\"structChannel.ChannelState\",\"name\":\"state\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ChannelFuncSigs maps the 4-byte function signature to its string representation.
var ChannelFuncSigs = map[string]string{
	"a48ed78f": "CooperativeClose((address,address,uint256,uint256,uint8,uint128),(address,address,uint256,uint256,uint8,uint128),bytes,bytes32)",
	"8e11a579": "ForceClose((address,address,uint256,uint256,uint8,uint128),bytes32)",
	"e16ac9f8": "accept((address,address,uint256,uint256,uint8,uint128),(address,address,uint256,uint256,uint8,uint128),bytes32)",
	"7614c599": "challenge((address,address,uint256,uint256,uint8,uint128),(address,address,uint256,uint256,uint8,uint128),bytes,bytes32)",
	"7a7ebd7b": "channels(bytes32)",
	"b03a28a8": "disputeChallenge((address,address,uint256,uint256,uint8,uint128),(address,address,uint256,uint256,uint8,uint128),bytes,bytes32)",
	"11be1997": "disputes(bytes32)",
	"b562b709": "hashState((address,address,uint256,uint256,uint8,uint128))",
	"f513def1": "open((address,address,uint256,uint256,uint8,uint128),bytes32)",
}

// ChannelBin is the compiled bytecode used for deploying new contracts.
var ChannelBin = "0x6080604052600080546001600160401b0319166201518017905534801561002557600080fd5b50611719806100356000396000f3fe6080604052600436106100865760003560e01c8063a48ed78f11610059578063a48ed78f14610171578063b03a28a814610191578063b562b709146101a4578063e16ac9f8146101c4578063f513def1146101d757600080fd5b806311be19971461008b5780637614c599146101015780637a7ebd7b146101165780638e11a57914610151575b600080fd5b34801561009757600080fd5b506100d46100a63660046113d1565b60026020526000908152604090205467ffffffffffffffff811690600160401b90046001600160a01b031682565b6040805167ffffffffffffffff90931683526001600160a01b039091166020830152015b60405180910390f35b61011461010f366004611500565b6101ea565b005b34801561012257600080fd5b506101436101313660046113d1565b60016020526000908152604090205481565b6040519081526020016100f8565b34801561015d57600080fd5b5061011461016c3660046113ea565b6104ac565b34801561017d57600080fd5b5061011461018c366004611415565b610660565b61011461019f366004611415565b610892565b3480156101b057600080fd5b506101436101bf366004611480565b610b18565b6101146101d23660046114c0565b610b6d565b6101146101e53660046114a3565b610e90565b8360018082608001516002811115610204576102046116b7565b1461022a5760405162461bcd60e51b815260040161022190611626565b60405180910390fd5b858580600001516001600160a01b031682600001516001600160a01b0316146102655760405162461bcd60e51b8152600401610221906115f6565b80602001516001600160a01b031682602001516001600160a01b03161461029e5760405162461bcd60e51b8152600401610221906115f6565b806040015181606001516102b2919061165d565b826060015183604001516102c6919061165d565b146102e35760405162461bcd60e51b8152600401610221906115f6565b6102ec88610b18565b600086815260016020526040902054146103185760405162461bcd60e51b8152600401610221906115a0565b600061032489336110ad565b90506001600160a01b03811661036c5760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b21037ba3432b960991b6044820152606401610221565b600061037789610b18565b9050816001600160a01b031661038d828a611108565b6001600160a01b0316146103b35760405162461bcd60e51b8152600401610221906115cb565b60008781526002602052604090205467ffffffffffffffff16156104195760405162461bcd60e51b815260206004820152601b60248201527f6469737075746520616c726561647920696e2070726f677265737300000000006044820152606401610221565b6000878152600160209081526040808320849055600282529182902080544267ffffffffffffffff166001600160e01b0319909116811733600160401b021790915560a08c015183516001600160801b0390911681529182015288917f6e129f552c838f9ab74444d21bd821c7694edfcbbd3c2294dbc7bc4829040180910160405180910390a250505050505050505050565b6104bb36839003830183611480565b600180826080015160028111156104d4576104d46116b7565b146104f15760405162461bcd60e51b815260040161022190611626565b6105036101bf36869003860186611480565b6000848152600160205260409020541461052f5760405162461bcd60e51b8152600401610221906115a0565b60008381526002602052604090205467ffffffffffffffff1661058b5760405162461bcd60e51b81526020600482015260146024820152736e6f206469737075746520617661696c61626c6560601b6044820152606401610221565b60008054848252600260205260409091205442916105b69167ffffffffffffffff9182169116611675565b67ffffffffffffffff16106106175760405162461bcd60e51b815260206004820152602160248201527f666f72636520636c6f7365206265666f7265206469737075746520706572696f6044820152601960fa1b6064820152608401610221565b61062f61062936869003860186611480565b846111f2565b60405183907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d7990600090a250505050565b61066f36859003850185611480565b60018082608001516002811115610688576106886116b7565b146106a55760405162461bcd60e51b815260040161022190611626565b6106b436879003870187611480565b8551815187916001600160a01b039182169116146106e45760405162461bcd60e51b8152600401610221906115f6565b80602001516001600160a01b031682602001516001600160a01b03161461071d5760405162461bcd60e51b8152600401610221906115f6565b80604001518160600151610731919061165d565b82606001518360400151610745919061165d565b146107625760405162461bcd60e51b8152600401610221906115f6565b6107746101bf368a90038a018a611480565b600086815260016020526040902054146107a05760405162461bcd60e51b8152600401610221906115a0565b6002876080015160028111156107b8576107b86116b7565b146107fc5760405162461bcd60e51b81526020600482015260146024820152731b995dc81cdd185d19481b9bdd0818db1bdcd95960621b6044820152606401610221565b600061080888336110ad565b9050600061081589610b18565b9050816001600160a01b031661082b828a611108565b6001600160a01b0316146108515760405162461bcd60e51b8152600401610221906115cb565b61085b89886111f2565b60405187907f7b6ac8bce3193cb9464e9070476bf8926e449f5f743f8c7578eea15265467d7990600090a250505050505050505050565b6108a136859003850185611480565b600180826080015160028111156108ba576108ba6116b7565b146108d75760405162461bcd60e51b815260040161022190611626565b6108e636879003870187611480565b8551815187916001600160a01b039182169116146109165760405162461bcd60e51b8152600401610221906115f6565b80602001516001600160a01b031682602001516001600160a01b03161461094f5760405162461bcd60e51b8152600401610221906115f6565b80604001518160600151610963919061165d565b82606001518360400151610977919061165d565b146109945760405162461bcd60e51b8152600401610221906115f6565b6109a66101bf368a90038a018a611480565b600086815260016020526040902054146109d25760405162461bcd60e51b8152600401610221906115a0565b60006109ec6109e6368b90038b018b611480565b336110ad565b6000878152600260205260409020549091506001600160a01b03808316600160401b9092041614610a505760405162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b637b9b2b960911b6044820152606401610221565b6000610a5b89610b18565b9050816001600160a01b0316610a71828a611108565b6001600160a01b031614610a975760405162461bcd60e51b8152600401610221906115cb565b8860a001516001600160801b03168a60a0016020810190610ab89190611521565b6001600160801b0316106104195760405162461bcd60e51b815260206004820152602160248201527f6469737075746564206368616c6c656e67652077697468206f6c6420737461746044820152606560f81b6064820152608401610221565b600080826000015183602001518460400151856060015186608001518760a00151604051602001610b4e9695949392919061153c565b60408051601f1981840301815291905280516020909101209392505050565b8260008082608001516002811115610b8757610b876116b7565b14610ba45760405162461bcd60e51b815260040161022190611626565b8360018082608001516002811115610bbe57610bbe6116b7565b14610bdb5760405162461bcd60e51b815260040161022190611626565b868680600001516001600160a01b031682600001516001600160a01b031614610c165760405162461bcd60e51b8152600401610221906115f6565b80602001516001600160a01b031682602001516001600160a01b031614610c4f5760405162461bcd60e51b8152600401610221906115f6565b80604001518160600151610c63919061165d565b82606001518360400151610c77919061165d565b14610c945760405162461bcd60e51b8152600401610221906115f6565b610c9d89610b18565b60008881526001602052604090205414610cc95760405162461bcd60e51b8152600401610221906115a0565b8760400151896040015114610d2a5760405162461bcd60e51b815260206004820152602160248201527f63616e206e6f7420737465616c206d6f6e65792066726f6d2070726f706f73656044820152603960f91b6064820152608401610221565b88602001516001600160a01b0316336001600160a01b031614610d8f5760405162461bcd60e51b815260206004820152601760248201527f6163636570746f72206d7573742062652073656e6465720000000000000000006044820152606401610221565b88606001513414610de25760405162461bcd60e51b815260206004820152601e60248201527f70726f706f736572206d7573742061646420656e6f7567682076616c756500006044820152606401610221565b60a08801516001600160801b031615610e3d5760405162461bcd60e51b815260206004820152601e60248201527f726f756e642073657420696e206368616e6e656c20616363657074696e6700006044820152606401610221565b6000610e4889610b18565b6000898152600160205260408082208390555191925089917f1f3c0697c3ada95f9e84a917995664b76cd8b4ae5de25e77ee111122ae3a00d09190a250505050505050505050565b81516001600160a01b03163314610ee95760405162461bcd60e51b815260206004820152601760248201527f70726f706f736572206d7573742062652073656e6465720000000000000000006044820152606401610221565b81604001513414610f3c5760405162461bcd60e51b815260206004820152601e60248201527f70726f706f736572206d7573742061646420656e6f7567682076616c756500006044820152606401610221565b60008181526001602052604090205415610f915760405162461bcd60e51b81526020600482015260166024820152756368616e6e656c20616c726561647920696e2075736560501b6044820152606401610221565b60a08201516001600160801b031615610fec5760405162461bcd60e51b815260206004820152601c60248201527f726f756e642073657420696e206368616e6e656c206f70656e696e67000000006044820152606401610221565b600082608001516002811115611004576110046116b7565b146110615760405162461bcd60e51b815260206004820152602760248201527f63616e206f6e6c79206f70656e206368616e6e656c20696e20737461746520706044820152661c9bdc1bdcd95960ca1b6064820152608401610221565b600061106c83610b18565b6000838152600160205260408082208390555191925083917fddffb592d6434d02d388cf2eb4fbfa796fbfcd09e278d3466e7194dfd3c23a049190a2505050565b600082600001516001600160a01b0316826001600160a01b031614156110d857506020820151611102565b82602001516001600160a01b0316826001600160a01b031614156110fe57508151611102565b5060005b92915050565b6000815160411461111b57506000611102565b60208201516040830151606084015160001a7f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156111615760009350505050611102565b8060ff16601b1415801561117957508060ff16601c14155b1561118a5760009350505050611102565b60408051600081526020810180835288905260ff831691810191909152606081018490526080810183905260019060a0016020604051602081039080840390855afa1580156111dd573d6000803e3d6000fd5b5050604051601f190151979650505050505050565b60008181526001602052604080822082905583518482015191516001600160a01b039091169282156108fc02929190818181858888f150505050602083015160608401516040516001600160a01b0390921692506108fc811502916000818181858888f150505050505050565b80356001600160a01b038116811461127657600080fd5b919050565b600082601f83011261128c57600080fd5b813567ffffffffffffffff808211156112a7576112a76116cd565b604051601f8301601f19908116603f011681019082821181831017156112cf576112cf6116cd565b816040528381528660208588010111156112e857600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060c0828403121561131a57600080fd5b50919050565b600060c0828403121561133257600080fd5b60405160c0810181811067ffffffffffffffff82111715611355576113556116cd565b6040529050806113648361125f565b81526113726020840161125f565b6020820152604083013560408201526060830135606082015260808301356003811061139d57600080fd5b60808201526113ae60a084016113ba565b60a08201525092915050565b80356001600160801b038116811461127657600080fd5b6000602082840312156113e357600080fd5b5035919050565b60008060e083850312156113fd57600080fd5b6114078484611308565b9460c0939093013593505050565b6000806000806101c0858703121561142c57600080fd5b6114368686611308565b93506114458660c08701611320565b925061018085013567ffffffffffffffff81111561146257600080fd5b61146e8782880161127b565b949793965093946101a0013593505050565b600060c0828403121561149257600080fd5b61149c8383611320565b9392505050565b60008060e083850312156114b657600080fd5b6114078484611320565b60008060006101a084860312156114d657600080fd5b6114e08585611320565b92506114ef8560c08601611320565b915061018084013590509250925092565b6000806000806101c0858703121561151757600080fd5b6114368686611320565b60006020828403121561153357600080fd5b61149c826113ba565b6001600160a01b03878116825286166020820152604081018590526060810184905260c081016003841061158057634e487b7160e01b600052602160045260246000fd5b8360808301526001600160801b03831660a0830152979650505050505050565b602080825260119082015270696e76616c6964206f6c6420737461746560781b604082015260600190565b602080825260119082015270696e76616c6964207369676e617475726560781b604082015260600190565b6020808252601690820152751a5b9d985b1a59081cdd185d19481c1c9bdd9a59195960521b604082015260600190565b6020808252601b908201527f4368616e6e656c20697320696e20696e76616c69642073746174650000000000604082015260600190565b60008219821115611670576116706116a1565b500190565b600067ffffffffffffffff808316818516808303821115611698576116986116a1565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212202d34c853e842a3a48cc837a64707952416b2c0bd546c8ee97f3469ecafecf48064736f6c63430008070033"

// DeployChannel deploys a new Ethereum contract, binding an instance of Channel to it.
func DeployChannel(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Channel, error) {
	parsed, err := abi.JSON(strings.NewReader(ChannelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChannelBin), backend)
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
// Solidity: function channels(bytes32 ) view returns(bytes32)
func (_Channel *ChannelCaller) Channels(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "channels", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(bytes32)
func (_Channel *ChannelSession) Channels(arg0 [32]byte) ([32]byte, error) {
	return _Channel.Contract.Channels(&_Channel.CallOpts, arg0)
}

// Channels is a free data retrieval call binding the contract method 0x7a7ebd7b.
//
// Solidity: function channels(bytes32 ) view returns(bytes32)
func (_Channel *ChannelCallerSession) Channels(arg0 [32]byte) ([32]byte, error) {
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

// HashState is a free data retrieval call binding the contract method 0xb562b709.
//
// Solidity: function hashState((address,address,uint256,uint256,uint8,uint128) state) pure returns(bytes32)
func (_Channel *ChannelCaller) HashState(opts *bind.CallOpts, state ChannelChannelState) ([32]byte, error) {
	var out []interface{}
	err := _Channel.contract.Call(opts, &out, "hashState", state)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashState is a free data retrieval call binding the contract method 0xb562b709.
//
// Solidity: function hashState((address,address,uint256,uint256,uint8,uint128) state) pure returns(bytes32)
func (_Channel *ChannelSession) HashState(state ChannelChannelState) ([32]byte, error) {
	return _Channel.Contract.HashState(&_Channel.CallOpts, state)
}

// HashState is a free data retrieval call binding the contract method 0xb562b709.
//
// Solidity: function hashState((address,address,uint256,uint256,uint8,uint128) state) pure returns(bytes32)
func (_Channel *ChannelCallerSession) HashState(state ChannelChannelState) ([32]byte, error) {
	return _Channel.Contract.HashState(&_Channel.CallOpts, state)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0xa48ed78f.
//
// Solidity: function CooperativeClose((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) returns()
func (_Channel *ChannelTransactor) CooperativeClose(opts *bind.TransactOpts, oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "CooperativeClose", oldState, newState, sig, id)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0xa48ed78f.
//
// Solidity: function CooperativeClose((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) returns()
func (_Channel *ChannelSession) CooperativeClose(oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.CooperativeClose(&_Channel.TransactOpts, oldState, newState, sig, id)
}

// CooperativeClose is a paid mutator transaction binding the contract method 0xa48ed78f.
//
// Solidity: function CooperativeClose((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) returns()
func (_Channel *ChannelTransactorSession) CooperativeClose(oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.CooperativeClose(&_Channel.TransactOpts, oldState, newState, sig, id)
}

// ForceClose is a paid mutator transaction binding the contract method 0x8e11a579.
//
// Solidity: function ForceClose((address,address,uint256,uint256,uint8,uint128) oldState, bytes32 id) returns()
func (_Channel *ChannelTransactor) ForceClose(opts *bind.TransactOpts, oldState ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "ForceClose", oldState, id)
}

// ForceClose is a paid mutator transaction binding the contract method 0x8e11a579.
//
// Solidity: function ForceClose((address,address,uint256,uint256,uint8,uint128) oldState, bytes32 id) returns()
func (_Channel *ChannelSession) ForceClose(oldState ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.ForceClose(&_Channel.TransactOpts, oldState, id)
}

// ForceClose is a paid mutator transaction binding the contract method 0x8e11a579.
//
// Solidity: function ForceClose((address,address,uint256,uint256,uint8,uint128) oldState, bytes32 id) returns()
func (_Channel *ChannelTransactorSession) ForceClose(oldState ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.ForceClose(&_Channel.TransactOpts, oldState, id)
}

// Accept is a paid mutator transaction binding the contract method 0xe16ac9f8.
//
// Solidity: function accept((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes32 id) payable returns()
func (_Channel *ChannelTransactor) Accept(opts *bind.TransactOpts, oldState ChannelChannelState, newState ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "accept", oldState, newState, id)
}

// Accept is a paid mutator transaction binding the contract method 0xe16ac9f8.
//
// Solidity: function accept((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes32 id) payable returns()
func (_Channel *ChannelSession) Accept(oldState ChannelChannelState, newState ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Accept(&_Channel.TransactOpts, oldState, newState, id)
}

// Accept is a paid mutator transaction binding the contract method 0xe16ac9f8.
//
// Solidity: function accept((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes32 id) payable returns()
func (_Channel *ChannelTransactorSession) Accept(oldState ChannelChannelState, newState ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Accept(&_Channel.TransactOpts, oldState, newState, id)
}

// Challenge is a paid mutator transaction binding the contract method 0x7614c599.
//
// Solidity: function challenge((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) payable returns()
func (_Channel *ChannelTransactor) Challenge(opts *bind.TransactOpts, oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "challenge", oldState, newState, sig, id)
}

// Challenge is a paid mutator transaction binding the contract method 0x7614c599.
//
// Solidity: function challenge((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) payable returns()
func (_Channel *ChannelSession) Challenge(oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Challenge(&_Channel.TransactOpts, oldState, newState, sig, id)
}

// Challenge is a paid mutator transaction binding the contract method 0x7614c599.
//
// Solidity: function challenge((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) payable returns()
func (_Channel *ChannelTransactorSession) Challenge(oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Challenge(&_Channel.TransactOpts, oldState, newState, sig, id)
}

// DisputeChallenge is a paid mutator transaction binding the contract method 0xb03a28a8.
//
// Solidity: function disputeChallenge((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) payable returns()
func (_Channel *ChannelTransactor) DisputeChallenge(opts *bind.TransactOpts, oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "disputeChallenge", oldState, newState, sig, id)
}

// DisputeChallenge is a paid mutator transaction binding the contract method 0xb03a28a8.
//
// Solidity: function disputeChallenge((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) payable returns()
func (_Channel *ChannelSession) DisputeChallenge(oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.DisputeChallenge(&_Channel.TransactOpts, oldState, newState, sig, id)
}

// DisputeChallenge is a paid mutator transaction binding the contract method 0xb03a28a8.
//
// Solidity: function disputeChallenge((address,address,uint256,uint256,uint8,uint128) oldState, (address,address,uint256,uint256,uint8,uint128) newState, bytes sig, bytes32 id) payable returns()
func (_Channel *ChannelTransactorSession) DisputeChallenge(oldState ChannelChannelState, newState ChannelChannelState, sig []byte, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.DisputeChallenge(&_Channel.TransactOpts, oldState, newState, sig, id)
}

// Open is a paid mutator transaction binding the contract method 0xf513def1.
//
// Solidity: function open((address,address,uint256,uint256,uint8,uint128) state, bytes32 id) payable returns()
func (_Channel *ChannelTransactor) Open(opts *bind.TransactOpts, state ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.contract.Transact(opts, "open", state, id)
}

// Open is a paid mutator transaction binding the contract method 0xf513def1.
//
// Solidity: function open((address,address,uint256,uint256,uint8,uint128) state, bytes32 id) payable returns()
func (_Channel *ChannelSession) Open(state ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Open(&_Channel.TransactOpts, state, id)
}

// Open is a paid mutator transaction binding the contract method 0xf513def1.
//
// Solidity: function open((address,address,uint256,uint256,uint8,uint128) state, bytes32 id) payable returns()
func (_Channel *ChannelTransactorSession) Open(state ChannelChannelState, id [32]byte) (*types.Transaction, error) {
	return _Channel.Contract.Open(&_Channel.TransactOpts, state, id)
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

// ECDSAABI is the input ABI used to generate the binding from.
const ECDSAABI = "[]"

// ECDSABin is the compiled bytecode used for deploying new contracts.
var ECDSABin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b0013cb33e8f725049caf38be7c6fa8696b87166ee6baaa6263fbbd142839f3964736f6c63430008070033"

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECDSABin), backend)
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
