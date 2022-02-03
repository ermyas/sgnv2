// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// BridgeTestTokenMetaData contains all meta data concerning the BridgeTestToken contract.
var BridgeTestTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"}],\"name\":\"BridgeSupplyCapUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"TokenSwapCapUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridges\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_swapCap\",\"type\":\"uint256\"}],\"name\":\"setBridgeTokenSwapCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"swapBridgeForCanonical\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"swapCanonicalForBridge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapSupplies\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"updateBridgeSupplyCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001a0038038062001a00833981016040819052620000349162000252565b8282828282816003908051906020019062000051929190620000df565b50805162000067906004906020840190620000df565b5050506200007b336200008d60201b60201c565b60ff1660805250620003149350505050565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b828054620000ed90620002d7565b90600052602060002090601f0160209004810192826200011157600085556200015c565b82601f106200012c57805160ff19168380011785556200015c565b828001600101855582156200015c579182015b828111156200015c5782518255916020019190600101906200013f565b506200016a9291506200016e565b5090565b5b808211156200016a57600081556001016200016f565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001ad57600080fd5b81516001600160401b0380821115620001ca57620001ca62000185565b604051601f8301601f19908116603f01168101908282118183101715620001f557620001f562000185565b816040528381526020925086838588010111156200021257600080fd5b600091505b8382101562000236578582018301518183018401529082019062000217565b83821115620002485760008385830101525b9695505050505050565b6000806000606084860312156200026857600080fd5b83516001600160401b03808211156200028057600080fd5b6200028e878388016200019b565b94506020860151915080821115620002a557600080fd5b50620002b4868287016200019b565b925050604084015160ff81168114620002cc57600080fd5b809150509250925092565b600181811c90821680620002ec57607f821691505b602082108114156200030e57634e487b7160e01b600052602260045260246000fd5b50919050565b6080516116d062000330600039600061020b01526116d06000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c8063893d20e8116100d8578063a9059cbb1161008c578063dd62ed3e11610066578063dd62ed3e14610381578063f2fde38b146103ba578063f716932e146103cd57600080fd5b8063a9059cbb14610334578063cd04611914610347578063ced67f0c1461035a57600080fd5b806395d89b41116100bd57806395d89b41146103065780639dc29fac1461030e578063a457c2d71461032157600080fd5b8063893d20e8146102d55780638da5cb5b146102f557600080fd5b8063313ce5671161013a57806340c10f191161011457806340c10f19146102845780634ce2f71a1461029757806370a08231146102ac57600080fd5b8063313ce5671461020457806334faea1b14610235578063395093511461027157600080fd5b806313f1a4a41161016b57806313f1a4a4146101c857806318160ddd146101e957806323b872dd146101f157600080fd5b806306fdde0314610187578063095ea7b3146101a5575b600080fd5b61018f6103e0565b60405161019c91906114d9565b60405180910390f35b6101b86101b3366004611528565b610472565b604051901515815260200161019c565b6101db6101d6366004611528565b610488565b60405190815260200161019c565b6002546101db565b6101b86101ff366004611552565b610533565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019c565b61025c61024336600461158e565b6007602052600090815260409020805460019091015482565b6040805192835260208301919091520161019c565b6101b861027f366004611528565b6105f4565b6101b8610292366004611528565b610630565b6102aa6102a5366004611528565b610700565b005b6101db6102ba36600461158e565b6001600160a01b031660009081526020819052604090205490565b6102dd6107c2565b6040516001600160a01b03909116815260200161019c565b6005546001600160a01b03166102dd565b61018f6107db565b6101b861031c366004611528565b6107ea565b6101b861032f366004611528565b61085c565b6101b8610342366004611528565b610903565b6102aa610355366004611528565b610910565b61025c61036836600461158e565b6006602052600090815260409020805460019091015482565b6101db61038f3660046115a9565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6102aa6103c836600461158e565b6109ca565b6101db6103db366004611528565b610abb565b6060600380546103ef906115dc565b80601f016020809104026020016040519081016040528092919081815260200182805461041b906115dc565b80156104685780601f1061043d57610100808354040283529160200191610468565b820191906000526020600020905b81548152906001019060200180831161044b57829003601f168201915b5050505050905090565b600061047f338484610bb9565b50600192915050565b6001600160a01b038216600090815260076020526040812080546104f35760405162461bcd60e51b815260206004820152601460248201527f696e76616c69642062726964676520746f6b656e00000000000000000000000060448201526064015b60405180910390fd5b82816001016000828254610507919061162d565b9091555061051790503384610cde565b61052b6001600160a01b0385163385610e29565b509092915050565b6000610540848484610eb9565b6001600160a01b0384166000908152600160209081526040808320338452909152902054828110156105da5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206160448201527f6c6c6f77616e636500000000000000000000000000000000000000000000000060648201526084016104ea565b6105e78533858403610bb9565b60019150505b9392505050565b3360008181526001602090815260408083206001600160a01b0387168452909152812054909161047f91859061062b908690611644565b610bb9565b336000908152600660205260408120805461067e5760405162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b60448201526064016104ea565b828160010160008282546106929190611644565b90915550508054600182015411156106ec5760405162461bcd60e51b815260206004820152601960248201527f657863656564732062726964676520737570706c79206361700000000000000060448201526064016104ea565b6106f684846110b8565b5060019392505050565b336107136005546001600160a01b031690565b6001600160a01b0316146107695760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104ea565b6001600160a01b038216600081815260066020908152604091829020849055815192835282018390527f59e1e4348943de408b89af8ab71e502ea722dd41efd1ff4a3548c60e83e91c6091015b60405180910390a15050565b60006107d66005546001600160a01b031690565b905090565b6060600480546103ef906115dc565b33600090815260066020526040812080546108385760405162461bcd60e51b815260206004820152600e60248201526d34b73b30b634b21031b0b63632b960911b60448201526064016104ea565b8281600101600082825461084c919061162d565b909155506106f690508484610cde565b3360009081526001602090815260408083206001600160a01b0386168452909152812054828110156108f65760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016104ea565b6106f63385858403610bb9565b600061047f338484610eb9565b336109236005546001600160a01b031690565b6001600160a01b0316146109795760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104ea565b6001600160a01b038216600081815260076020908152604091829020849055815192835282018390527f51c7b3899924578d835c066303e3f765c25fea17d7b18840cd109a90f5c5601f91016107b6565b336109dd6005546001600160a01b031690565b6001600160a01b031614610a335760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104ea565b6001600160a01b038116610aaf5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104ea565b610ab881611197565b50565b6001600160a01b03821660009081526007602052604081208054610b215760405162461bcd60e51b815260206004820152601460248201527f696e76616c69642062726964676520746f6b656e00000000000000000000000060448201526064016104ea565b80546001820154610b33908590611644565b10610b805760405162461bcd60e51b815260206004820152600f60248201527f657863656564207377617020636170000000000000000000000000000000000060448201526064016104ea565b82816001016000828254610b949190611644565b90915550610ba4905033846110b8565b61052b6001600160a01b038516333086611201565b6001600160a01b038316610c1b5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016104ea565b6001600160a01b038216610c7c5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016104ea565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b6001600160a01b038216610d3e5760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b60648201526084016104ea565b6001600160a01b03821660009081526020819052604090205481811015610db25760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b60648201526084016104ea565b6001600160a01b0383166000908152602081905260408120838303905560028054849290610de190849061162d565b90915550506040518281526000906001600160a01b038516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610cd1565b505050565b6040516001600160a01b038316602482015260448101829052610e2490849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611239565b6001600160a01b038316610f355760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016104ea565b6001600160a01b038216610f975760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016104ea565b6001600160a01b038316600090815260208190526040902054818110156110265760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016104ea565b6001600160a01b0380851660009081526020819052604080822085850390559185168152908120805484929061105d908490611644565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516110a991815260200190565b60405180910390a35b50505050565b6001600160a01b03821661110e5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016104ea565b80600260008282546111209190611644565b90915550506001600160a01b0382166000908152602081905260408120805483929061114d908490611644565b90915550506040518181526001600160a01b038316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b600580546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6040516001600160a01b03808516602483015283166044820152606481018290526110b29085906323b872dd60e01b90608401610e55565b600061128e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661131e9092919063ffffffff16565b805190915015610e2457808060200190518101906112ac919061165c565b610e245760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016104ea565b606061132d8484600085611335565b949350505050565b6060824710156113ad5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016104ea565b843b6113fb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104ea565b600080866001600160a01b03168587604051611417919061167e565b60006040518083038185875af1925050503d8060008114611454576040519150601f19603f3d011682016040523d82523d6000602084013e611459565b606091505b5091509150611469828286611474565b979650505050505050565b606083156114835750816105ed565b8251156114935782518084602001fd5b8160405162461bcd60e51b81526004016104ea91906114d9565b60005b838110156114c85781810151838201526020016114b0565b838111156110b25750506000910152565b60208152600082518060208401526114f88160408501602087016114ad565b601f01601f19169190910160400192915050565b80356001600160a01b038116811461152357600080fd5b919050565b6000806040838503121561153b57600080fd5b6115448361150c565b946020939093013593505050565b60008060006060848603121561156757600080fd5b6115708461150c565b925061157e6020850161150c565b9150604084013590509250925092565b6000602082840312156115a057600080fd5b6105ed8261150c565b600080604083850312156115bc57600080fd5b6115c58361150c565b91506115d36020840161150c565b90509250929050565b600181811c908216806115f057607f821691505b6020821081141561161157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008282101561163f5761163f611617565b500390565b6000821982111561165757611657611617565b500190565b60006020828403121561166e57600080fd5b815180151581146105ed57600080fd5b600082516116908184602087016114ad565b919091019291505056fea26469706673582212207d062e703c7ffea6b7f07e2e9c85ff8bc8dadd0c6c03e1810be6a8aba599ed2e64736f6c63430008090033",
}

// BridgeTestTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTestTokenMetaData.ABI instead.
var BridgeTestTokenABI = BridgeTestTokenMetaData.ABI

// BridgeTestTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTestTokenMetaData.Bin instead.
var BridgeTestTokenBin = BridgeTestTokenMetaData.Bin

// DeployBridgeTestToken deploys a new Ethereum contract, binding an instance of BridgeTestToken to it.
func DeployBridgeTestToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *BridgeTestToken, error) {
	parsed, err := BridgeTestTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTestTokenBin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeTestToken{BridgeTestTokenCaller: BridgeTestTokenCaller{contract: contract}, BridgeTestTokenTransactor: BridgeTestTokenTransactor{contract: contract}, BridgeTestTokenFilterer: BridgeTestTokenFilterer{contract: contract}}, nil
}

// BridgeTestToken is an auto generated Go binding around an Ethereum contract.
type BridgeTestToken struct {
	BridgeTestTokenCaller     // Read-only binding to the contract
	BridgeTestTokenTransactor // Write-only binding to the contract
	BridgeTestTokenFilterer   // Log filterer for contract events
}

// BridgeTestTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTestTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTestTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTestTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTestTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTestTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTestTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTestTokenSession struct {
	Contract     *BridgeTestToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTestTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTestTokenCallerSession struct {
	Contract *BridgeTestTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BridgeTestTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTestTokenTransactorSession struct {
	Contract     *BridgeTestTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BridgeTestTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTestTokenRaw struct {
	Contract *BridgeTestToken // Generic contract binding to access the raw methods on
}

// BridgeTestTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTestTokenCallerRaw struct {
	Contract *BridgeTestTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTestTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTestTokenTransactorRaw struct {
	Contract *BridgeTestTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTestToken creates a new instance of BridgeTestToken, bound to a specific deployed contract.
func NewBridgeTestToken(address common.Address, backend bind.ContractBackend) (*BridgeTestToken, error) {
	contract, err := bindBridgeTestToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTestToken{BridgeTestTokenCaller: BridgeTestTokenCaller{contract: contract}, BridgeTestTokenTransactor: BridgeTestTokenTransactor{contract: contract}, BridgeTestTokenFilterer: BridgeTestTokenFilterer{contract: contract}}, nil
}

// NewBridgeTestTokenCaller creates a new read-only instance of BridgeTestToken, bound to a specific deployed contract.
func NewBridgeTestTokenCaller(address common.Address, caller bind.ContractCaller) (*BridgeTestTokenCaller, error) {
	contract, err := bindBridgeTestToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenCaller{contract: contract}, nil
}

// NewBridgeTestTokenTransactor creates a new write-only instance of BridgeTestToken, bound to a specific deployed contract.
func NewBridgeTestTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTestTokenTransactor, error) {
	contract, err := bindBridgeTestToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenTransactor{contract: contract}, nil
}

// NewBridgeTestTokenFilterer creates a new log filterer instance of BridgeTestToken, bound to a specific deployed contract.
func NewBridgeTestTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTestTokenFilterer, error) {
	contract, err := bindBridgeTestToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenFilterer{contract: contract}, nil
}

// bindBridgeTestToken binds a generic wrapper to an already deployed contract.
func bindBridgeTestToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTestTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTestToken *BridgeTestTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTestToken.Contract.BridgeTestTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTestToken *BridgeTestTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.BridgeTestTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTestToken *BridgeTestTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.BridgeTestTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTestToken *BridgeTestTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTestToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTestToken *BridgeTestTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTestToken *BridgeTestTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeTestToken.Contract.Allowance(&_BridgeTestToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeTestToken.Contract.Allowance(&_BridgeTestToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BridgeTestToken.Contract.BalanceOf(&_BridgeTestToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BridgeTestToken.Contract.BalanceOf(&_BridgeTestToken.CallOpts, account)
}

// Bridges is a free data retrieval call binding the contract method 0xced67f0c.
//
// Solidity: function bridges(address ) view returns(uint256 cap, uint256 total)
func (_BridgeTestToken *BridgeTestTokenCaller) Bridges(opts *bind.CallOpts, arg0 common.Address) (struct {
	Cap   *big.Int
	Total *big.Int
}, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "bridges", arg0)

	outstruct := new(struct {
		Cap   *big.Int
		Total *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Bridges is a free data retrieval call binding the contract method 0xced67f0c.
//
// Solidity: function bridges(address ) view returns(uint256 cap, uint256 total)
func (_BridgeTestToken *BridgeTestTokenSession) Bridges(arg0 common.Address) (struct {
	Cap   *big.Int
	Total *big.Int
}, error) {
	return _BridgeTestToken.Contract.Bridges(&_BridgeTestToken.CallOpts, arg0)
}

// Bridges is a free data retrieval call binding the contract method 0xced67f0c.
//
// Solidity: function bridges(address ) view returns(uint256 cap, uint256 total)
func (_BridgeTestToken *BridgeTestTokenCallerSession) Bridges(arg0 common.Address) (struct {
	Cap   *big.Int
	Total *big.Int
}, error) {
	return _BridgeTestToken.Contract.Bridges(&_BridgeTestToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BridgeTestToken *BridgeTestTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BridgeTestToken *BridgeTestTokenSession) Decimals() (uint8, error) {
	return _BridgeTestToken.Contract.Decimals(&_BridgeTestToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BridgeTestToken *BridgeTestTokenCallerSession) Decimals() (uint8, error) {
	return _BridgeTestToken.Contract.Decimals(&_BridgeTestToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BridgeTestToken *BridgeTestTokenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BridgeTestToken *BridgeTestTokenSession) GetOwner() (common.Address, error) {
	return _BridgeTestToken.Contract.GetOwner(&_BridgeTestToken.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BridgeTestToken *BridgeTestTokenCallerSession) GetOwner() (common.Address, error) {
	return _BridgeTestToken.Contract.GetOwner(&_BridgeTestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeTestToken *BridgeTestTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeTestToken *BridgeTestTokenSession) Name() (string, error) {
	return _BridgeTestToken.Contract.Name(&_BridgeTestToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BridgeTestToken *BridgeTestTokenCallerSession) Name() (string, error) {
	return _BridgeTestToken.Contract.Name(&_BridgeTestToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTestToken *BridgeTestTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTestToken *BridgeTestTokenSession) Owner() (common.Address, error) {
	return _BridgeTestToken.Contract.Owner(&_BridgeTestToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeTestToken *BridgeTestTokenCallerSession) Owner() (common.Address, error) {
	return _BridgeTestToken.Contract.Owner(&_BridgeTestToken.CallOpts)
}

// SwapSupplies is a free data retrieval call binding the contract method 0x34faea1b.
//
// Solidity: function swapSupplies(address ) view returns(uint256 cap, uint256 total)
func (_BridgeTestToken *BridgeTestTokenCaller) SwapSupplies(opts *bind.CallOpts, arg0 common.Address) (struct {
	Cap   *big.Int
	Total *big.Int
}, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "swapSupplies", arg0)

	outstruct := new(struct {
		Cap   *big.Int
		Total *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SwapSupplies is a free data retrieval call binding the contract method 0x34faea1b.
//
// Solidity: function swapSupplies(address ) view returns(uint256 cap, uint256 total)
func (_BridgeTestToken *BridgeTestTokenSession) SwapSupplies(arg0 common.Address) (struct {
	Cap   *big.Int
	Total *big.Int
}, error) {
	return _BridgeTestToken.Contract.SwapSupplies(&_BridgeTestToken.CallOpts, arg0)
}

// SwapSupplies is a free data retrieval call binding the contract method 0x34faea1b.
//
// Solidity: function swapSupplies(address ) view returns(uint256 cap, uint256 total)
func (_BridgeTestToken *BridgeTestTokenCallerSession) SwapSupplies(arg0 common.Address) (struct {
	Cap   *big.Int
	Total *big.Int
}, error) {
	return _BridgeTestToken.Contract.SwapSupplies(&_BridgeTestToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeTestToken *BridgeTestTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeTestToken *BridgeTestTokenSession) Symbol() (string, error) {
	return _BridgeTestToken.Contract.Symbol(&_BridgeTestToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BridgeTestToken *BridgeTestTokenCallerSession) Symbol() (string, error) {
	return _BridgeTestToken.Contract.Symbol(&_BridgeTestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BridgeTestToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenSession) TotalSupply() (*big.Int, error) {
	return _BridgeTestToken.Contract.TotalSupply(&_BridgeTestToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BridgeTestToken *BridgeTestTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BridgeTestToken.Contract.TotalSupply(&_BridgeTestToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Approve(&_BridgeTestToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Approve(&_BridgeTestToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Burn(&_BridgeTestToken.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Burn(&_BridgeTestToken.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BridgeTestToken *BridgeTestTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.DecreaseAllowance(&_BridgeTestToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.DecreaseAllowance(&_BridgeTestToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BridgeTestToken *BridgeTestTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.IncreaseAllowance(&_BridgeTestToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.IncreaseAllowance(&_BridgeTestToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Mint(&_BridgeTestToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Mint(&_BridgeTestToken.TransactOpts, _to, _amount)
}

// SetBridgeTokenSwapCap is a paid mutator transaction binding the contract method 0xcd046119.
//
// Solidity: function setBridgeTokenSwapCap(address _bridgeToken, uint256 _swapCap) returns()
func (_BridgeTestToken *BridgeTestTokenTransactor) SetBridgeTokenSwapCap(opts *bind.TransactOpts, _bridgeToken common.Address, _swapCap *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "setBridgeTokenSwapCap", _bridgeToken, _swapCap)
}

// SetBridgeTokenSwapCap is a paid mutator transaction binding the contract method 0xcd046119.
//
// Solidity: function setBridgeTokenSwapCap(address _bridgeToken, uint256 _swapCap) returns()
func (_BridgeTestToken *BridgeTestTokenSession) SetBridgeTokenSwapCap(_bridgeToken common.Address, _swapCap *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.SetBridgeTokenSwapCap(&_BridgeTestToken.TransactOpts, _bridgeToken, _swapCap)
}

// SetBridgeTokenSwapCap is a paid mutator transaction binding the contract method 0xcd046119.
//
// Solidity: function setBridgeTokenSwapCap(address _bridgeToken, uint256 _swapCap) returns()
func (_BridgeTestToken *BridgeTestTokenTransactorSession) SetBridgeTokenSwapCap(_bridgeToken common.Address, _swapCap *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.SetBridgeTokenSwapCap(&_BridgeTestToken.TransactOpts, _bridgeToken, _swapCap)
}

// SwapBridgeForCanonical is a paid mutator transaction binding the contract method 0xf716932e.
//
// Solidity: function swapBridgeForCanonical(address _bridgeToken, uint256 _amount) returns(uint256)
func (_BridgeTestToken *BridgeTestTokenTransactor) SwapBridgeForCanonical(opts *bind.TransactOpts, _bridgeToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "swapBridgeForCanonical", _bridgeToken, _amount)
}

// SwapBridgeForCanonical is a paid mutator transaction binding the contract method 0xf716932e.
//
// Solidity: function swapBridgeForCanonical(address _bridgeToken, uint256 _amount) returns(uint256)
func (_BridgeTestToken *BridgeTestTokenSession) SwapBridgeForCanonical(_bridgeToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.SwapBridgeForCanonical(&_BridgeTestToken.TransactOpts, _bridgeToken, _amount)
}

// SwapBridgeForCanonical is a paid mutator transaction binding the contract method 0xf716932e.
//
// Solidity: function swapBridgeForCanonical(address _bridgeToken, uint256 _amount) returns(uint256)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) SwapBridgeForCanonical(_bridgeToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.SwapBridgeForCanonical(&_BridgeTestToken.TransactOpts, _bridgeToken, _amount)
}

// SwapCanonicalForBridge is a paid mutator transaction binding the contract method 0x13f1a4a4.
//
// Solidity: function swapCanonicalForBridge(address _bridgeToken, uint256 _amount) returns(uint256)
func (_BridgeTestToken *BridgeTestTokenTransactor) SwapCanonicalForBridge(opts *bind.TransactOpts, _bridgeToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "swapCanonicalForBridge", _bridgeToken, _amount)
}

// SwapCanonicalForBridge is a paid mutator transaction binding the contract method 0x13f1a4a4.
//
// Solidity: function swapCanonicalForBridge(address _bridgeToken, uint256 _amount) returns(uint256)
func (_BridgeTestToken *BridgeTestTokenSession) SwapCanonicalForBridge(_bridgeToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.SwapCanonicalForBridge(&_BridgeTestToken.TransactOpts, _bridgeToken, _amount)
}

// SwapCanonicalForBridge is a paid mutator transaction binding the contract method 0x13f1a4a4.
//
// Solidity: function swapCanonicalForBridge(address _bridgeToken, uint256 _amount) returns(uint256)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) SwapCanonicalForBridge(_bridgeToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.SwapCanonicalForBridge(&_BridgeTestToken.TransactOpts, _bridgeToken, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Transfer(&_BridgeTestToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.Transfer(&_BridgeTestToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.TransferFrom(&_BridgeTestToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BridgeTestToken *BridgeTestTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.TransferFrom(&_BridgeTestToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTestToken *BridgeTestTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTestToken *BridgeTestTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.TransferOwnership(&_BridgeTestToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeTestToken *BridgeTestTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.TransferOwnership(&_BridgeTestToken.TransactOpts, newOwner)
}

// UpdateBridgeSupplyCap is a paid mutator transaction binding the contract method 0x4ce2f71a.
//
// Solidity: function updateBridgeSupplyCap(address _bridge, uint256 _cap) returns()
func (_BridgeTestToken *BridgeTestTokenTransactor) UpdateBridgeSupplyCap(opts *bind.TransactOpts, _bridge common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.contract.Transact(opts, "updateBridgeSupplyCap", _bridge, _cap)
}

// UpdateBridgeSupplyCap is a paid mutator transaction binding the contract method 0x4ce2f71a.
//
// Solidity: function updateBridgeSupplyCap(address _bridge, uint256 _cap) returns()
func (_BridgeTestToken *BridgeTestTokenSession) UpdateBridgeSupplyCap(_bridge common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.UpdateBridgeSupplyCap(&_BridgeTestToken.TransactOpts, _bridge, _cap)
}

// UpdateBridgeSupplyCap is a paid mutator transaction binding the contract method 0x4ce2f71a.
//
// Solidity: function updateBridgeSupplyCap(address _bridge, uint256 _cap) returns()
func (_BridgeTestToken *BridgeTestTokenTransactorSession) UpdateBridgeSupplyCap(_bridge common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _BridgeTestToken.Contract.UpdateBridgeSupplyCap(&_BridgeTestToken.TransactOpts, _bridge, _cap)
}

// BridgeTestTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BridgeTestToken contract.
type BridgeTestTokenApprovalIterator struct {
	Event *BridgeTestTokenApproval // Event containing the contract specifics and raw log

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
func (it *BridgeTestTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTestTokenApproval)
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
		it.Event = new(BridgeTestTokenApproval)
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
func (it *BridgeTestTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTestTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTestTokenApproval represents a Approval event raised by the BridgeTestToken contract.
type BridgeTestTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeTestToken *BridgeTestTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BridgeTestTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeTestToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenApprovalIterator{contract: _BridgeTestToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeTestToken *BridgeTestTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BridgeTestTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BridgeTestToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTestTokenApproval)
				if err := _BridgeTestToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BridgeTestToken *BridgeTestTokenFilterer) ParseApproval(log types.Log) (*BridgeTestTokenApproval, error) {
	event := new(BridgeTestTokenApproval)
	if err := _BridgeTestToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTestTokenBridgeSupplyCapUpdatedIterator is returned from FilterBridgeSupplyCapUpdated and is used to iterate over the raw logs and unpacked data for BridgeSupplyCapUpdated events raised by the BridgeTestToken contract.
type BridgeTestTokenBridgeSupplyCapUpdatedIterator struct {
	Event *BridgeTestTokenBridgeSupplyCapUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTestTokenBridgeSupplyCapUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTestTokenBridgeSupplyCapUpdated)
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
		it.Event = new(BridgeTestTokenBridgeSupplyCapUpdated)
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
func (it *BridgeTestTokenBridgeSupplyCapUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTestTokenBridgeSupplyCapUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTestTokenBridgeSupplyCapUpdated represents a BridgeSupplyCapUpdated event raised by the BridgeTestToken contract.
type BridgeTestTokenBridgeSupplyCapUpdated struct {
	Bridge    common.Address
	SupplyCap *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgeSupplyCapUpdated is a free log retrieval operation binding the contract event 0x59e1e4348943de408b89af8ab71e502ea722dd41efd1ff4a3548c60e83e91c60.
//
// Solidity: event BridgeSupplyCapUpdated(address bridge, uint256 supplyCap)
func (_BridgeTestToken *BridgeTestTokenFilterer) FilterBridgeSupplyCapUpdated(opts *bind.FilterOpts) (*BridgeTestTokenBridgeSupplyCapUpdatedIterator, error) {

	logs, sub, err := _BridgeTestToken.contract.FilterLogs(opts, "BridgeSupplyCapUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenBridgeSupplyCapUpdatedIterator{contract: _BridgeTestToken.contract, event: "BridgeSupplyCapUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeSupplyCapUpdated is a free log subscription operation binding the contract event 0x59e1e4348943de408b89af8ab71e502ea722dd41efd1ff4a3548c60e83e91c60.
//
// Solidity: event BridgeSupplyCapUpdated(address bridge, uint256 supplyCap)
func (_BridgeTestToken *BridgeTestTokenFilterer) WatchBridgeSupplyCapUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTestTokenBridgeSupplyCapUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTestToken.contract.WatchLogs(opts, "BridgeSupplyCapUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTestTokenBridgeSupplyCapUpdated)
				if err := _BridgeTestToken.contract.UnpackLog(event, "BridgeSupplyCapUpdated", log); err != nil {
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

// ParseBridgeSupplyCapUpdated is a log parse operation binding the contract event 0x59e1e4348943de408b89af8ab71e502ea722dd41efd1ff4a3548c60e83e91c60.
//
// Solidity: event BridgeSupplyCapUpdated(address bridge, uint256 supplyCap)
func (_BridgeTestToken *BridgeTestTokenFilterer) ParseBridgeSupplyCapUpdated(log types.Log) (*BridgeTestTokenBridgeSupplyCapUpdated, error) {
	event := new(BridgeTestTokenBridgeSupplyCapUpdated)
	if err := _BridgeTestToken.contract.UnpackLog(event, "BridgeSupplyCapUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTestTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeTestToken contract.
type BridgeTestTokenOwnershipTransferredIterator struct {
	Event *BridgeTestTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeTestTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTestTokenOwnershipTransferred)
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
		it.Event = new(BridgeTestTokenOwnershipTransferred)
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
func (it *BridgeTestTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTestTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTestTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeTestToken contract.
type BridgeTestTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTestToken *BridgeTestTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeTestTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeTestToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenOwnershipTransferredIterator{contract: _BridgeTestToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTestToken *BridgeTestTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeTestTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeTestToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTestTokenOwnershipTransferred)
				if err := _BridgeTestToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeTestToken *BridgeTestTokenFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeTestTokenOwnershipTransferred, error) {
	event := new(BridgeTestTokenOwnershipTransferred)
	if err := _BridgeTestToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTestTokenTokenSwapCapUpdatedIterator is returned from FilterTokenSwapCapUpdated and is used to iterate over the raw logs and unpacked data for TokenSwapCapUpdated events raised by the BridgeTestToken contract.
type BridgeTestTokenTokenSwapCapUpdatedIterator struct {
	Event *BridgeTestTokenTokenSwapCapUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeTestTokenTokenSwapCapUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTestTokenTokenSwapCapUpdated)
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
		it.Event = new(BridgeTestTokenTokenSwapCapUpdated)
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
func (it *BridgeTestTokenTokenSwapCapUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTestTokenTokenSwapCapUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTestTokenTokenSwapCapUpdated represents a TokenSwapCapUpdated event raised by the BridgeTestToken contract.
type BridgeTestTokenTokenSwapCapUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenSwapCapUpdated is a free log retrieval operation binding the contract event 0x51c7b3899924578d835c066303e3f765c25fea17d7b18840cd109a90f5c5601f.
//
// Solidity: event TokenSwapCapUpdated(address token, uint256 cap)
func (_BridgeTestToken *BridgeTestTokenFilterer) FilterTokenSwapCapUpdated(opts *bind.FilterOpts) (*BridgeTestTokenTokenSwapCapUpdatedIterator, error) {

	logs, sub, err := _BridgeTestToken.contract.FilterLogs(opts, "TokenSwapCapUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenTokenSwapCapUpdatedIterator{contract: _BridgeTestToken.contract, event: "TokenSwapCapUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenSwapCapUpdated is a free log subscription operation binding the contract event 0x51c7b3899924578d835c066303e3f765c25fea17d7b18840cd109a90f5c5601f.
//
// Solidity: event TokenSwapCapUpdated(address token, uint256 cap)
func (_BridgeTestToken *BridgeTestTokenFilterer) WatchTokenSwapCapUpdated(opts *bind.WatchOpts, sink chan<- *BridgeTestTokenTokenSwapCapUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeTestToken.contract.WatchLogs(opts, "TokenSwapCapUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTestTokenTokenSwapCapUpdated)
				if err := _BridgeTestToken.contract.UnpackLog(event, "TokenSwapCapUpdated", log); err != nil {
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

// ParseTokenSwapCapUpdated is a log parse operation binding the contract event 0x51c7b3899924578d835c066303e3f765c25fea17d7b18840cd109a90f5c5601f.
//
// Solidity: event TokenSwapCapUpdated(address token, uint256 cap)
func (_BridgeTestToken *BridgeTestTokenFilterer) ParseTokenSwapCapUpdated(log types.Log) (*BridgeTestTokenTokenSwapCapUpdated, error) {
	event := new(BridgeTestTokenTokenSwapCapUpdated)
	if err := _BridgeTestToken.contract.UnpackLog(event, "TokenSwapCapUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTestTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BridgeTestToken contract.
type BridgeTestTokenTransferIterator struct {
	Event *BridgeTestTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BridgeTestTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTestTokenTransfer)
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
		it.Event = new(BridgeTestTokenTransfer)
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
func (it *BridgeTestTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTestTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTestTokenTransfer represents a Transfer event raised by the BridgeTestToken contract.
type BridgeTestTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeTestToken *BridgeTestTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BridgeTestTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTestToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTestTokenTransferIterator{contract: _BridgeTestToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeTestToken *BridgeTestTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BridgeTestTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BridgeTestToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTestTokenTransfer)
				if err := _BridgeTestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BridgeTestToken *BridgeTestTokenFilterer) ParseTransfer(log types.Log) (*BridgeTestTokenTransfer, error) {
	event := new(BridgeTestTokenTransfer)
	if err := _BridgeTestToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
