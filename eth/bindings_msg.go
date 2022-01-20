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

// CrossChainSwapSwapInfo is an auto generated low-level Go binding around an user-defined struct.
type CrossChainSwapSwapInfo struct {
	WantToken      common.Address
	User           common.Address
	SendBack       bool
	CbrMaxSlippage uint32
}

// MessageBusReceiverRouteInfo is an auto generated low-level Go binding around an user-defined struct.
type MessageBusReceiverRouteInfo struct {
	Sender     common.Address
	Receiver   common.Address
	SrcChainId uint64
}

// MessageBusReceiverTransferInfo is an auto generated low-level Go binding around an user-defined struct.
type MessageBusReceiverTransferInfo struct {
	T          uint8
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Seqnum     uint64
	SrcChainId uint64
	RefId      [32]byte
}

// TransferSwapSwapInfo is an auto generated low-level Go binding around an user-defined struct.
type TransferSwapSwapInfo struct {
	Path       []common.Address
	Dex        common.Address
	Deadline   *big.Int
	MinRecvAmt *big.Int
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c5836de53890d7f64a6b52ffe62306e803487ca8dc87ca8b3e06e3fa088190f564736f6c63430008090033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BatchTransferMetaData contains all meta data concerning the BatchTransfer contract.
var BatchTransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"enumMessageSenderLib.BridgeType\",\"name\":\"_bridgeType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"enumBatchTransfer.TransferStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002099380380620020998339810160408190526200003491620000b5565b6200003f3362000065565b600180546001600160a01b0319166001600160a01b0392909216919091179055620000e7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215620000c857600080fd5b81516001600160a01b0381168114620000e057600080fd5b9392505050565b611fa280620000f76000396000f3fe6080604052600436106100b15760003560e01c80638da5cb5b11610069578063ce35dd9a1161004e578063ce35dd9a146101bd578063f00f39ce146101d0578063f2fde38b146101e357600080fd5b80638da5cb5b1461016b578063a1a227fa1461019d57600080fd5b806320bff8931161009a57806320bff893146100f1578063547cad121461013657806364d39f421461015857600080fd5b80631599d265146100b657806320be95f2146100de575b600080fd5b6100c96100c4366004611720565b610203565b60405190151581526020015b60405180910390f35b6100c96100ec366004611782565b610384565b3480156100fd57600080fd5b5061012861010c36600461180b565b6002602052600090815260409020805460019091015460ff1682565b6040516100d5929190611860565b34801561014257600080fd5b50610156610151366004611874565b6103e9565b005b6101566101663660046118f1565b610481565b34801561017757600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016100d5565b3480156101a957600080fd5b50600154610185906001600160a01b031681565b6100c96101cb3660046119d2565b610723565b6100c96101de3660046119d2565b6108cf565b3480156101ef57600080fd5b506101566101fe366004611874565b6109bd565b6001546000906001600160a01b031633146102655760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b60008280602001905181019061027b9190611a51565b6040516bffffffffffffffffffffffff19606088901b1660208201526001600160c01b031960c087901b166034820152909150603c0160408051601f198184030181529181528151602092830120835167ffffffffffffffff16600090815260029093529120541461032f5760405162461bcd60e51b815260206004820152600f60248201527f696e76616c6964206d6573736167650000000000000000000000000000000000604482015260640161025c565b602080820151825167ffffffffffffffff16600090815260029283905260409020600190810180549293909260ff19169190849081111561037257610372611828565b021790555060019150505b9392505050565b6001546000906001600160a01b031633146103e15760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161025c565b949350505050565b336103fc6000546001600160a01b031690565b6001600160a01b0316146104525760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161025c565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b3332146104d05760405162461bcd60e51b815260206004820152600760248201527f4e6f7420454f4100000000000000000000000000000000000000000000000000604482015260640161025c565b6000805b82811015610514578383828181106104ee576104ee611ab4565b90506020020135826105009190611ae0565b91508061050c81611af8565b9150506104d4565b5060018054819060149061053a908390600160a01b900467ffffffffffffffff16611b13565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405180604001604052808c8a6040516020016105a892919060609290921b6bffffffffffffffffffffffff1916825260c01b6001600160c01b0319166014820152601c0190565b60408051601f1981840301815291905280516020918201208252016000905260018054600160a01b900467ffffffffffffffff1660009081526002602081815260409092208451815591840151828401805493949193909260ff199091169190849081111561061957610619611828565b02179055506106369150506001600160a01b038b1633308c610aae565b60408051608081018252600154600160a01b900467ffffffffffffffff16815281516020878102808301820190945287825260009381840192918a918a918291908501908490808284376000920191909152505050908252506040805160208781028281018201909352878252928301929091889188918291850190849080828437600092019190915250505090825250336020918201526040516106dc929101611b7a565b60405160208183030381529060405290506107148c8c8c8c600160149054906101000a900467ffffffffffffffff168d878e34610b4c565b50505050505050505050505050565b6001546000906001600160a01b031633146107805760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161025c565b6000828060200190518101906107969190611cba565b90506000805b82602001515181101561083e57610801836020015182815181106107c2576107c2611ab4565b6020026020010151846040015183815181106107e0576107e0611ab4565b60200260200101518a6001600160a01b0316610b839092919063ffffffff16565b8260400151818151811061081757610817611ab4565b60200260200101518261082a9190611ae0565b91508061083681611af8565b91505061079c565b50600061084b8288611dcd565b90508187111561086f57606083015161086f906001600160a01b038a169083610b83565b60408051808201909152835167ffffffffffffffff16815260009060208101600190526040516108a29190602001611de4565b60405160208183030381529060405290506108bf8a888334610bb8565b5060019998505050505050505050565b6001546000906001600160a01b0316331461092c5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161025c565b6000828060200190518101906109429190611cba565b606081015190915061095f906001600160a01b0388169087610b83565b60408051808201909152815167ffffffffffffffff16815260009060208101600290526040516109929190602001611de4565b60405160208183030381529060405290506109af88868334610bb8565b506001979650505050505050565b336109d06000546001600160a01b031690565b6001600160a01b031614610a265760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161025c565b6001600160a01b038116610aa25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161025c565b610aab81610bd4565b50565b6040516001600160a01b0380851660248301528316604482015260648101829052610b469085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610c31565b50505050565b6000610b758a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b610d16565b9a9950505050505050505050565b6040516001600160a01b038316602482015260448101829052610bb390849063a9059cbb60e01b90606401610ae2565b505050565b600154610b46908590859085906001600160a01b031685610de6565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610c86826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610e519092919063ffffffff16565b805190915015610bb35780806020019051810190610ca49190611e10565b610bb35760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161025c565b60006001846003811115610d2c57610d2c611828565b1415610d4a57610d438b8b8b8b8b8b8b8a8a610e60565b9050610b75565b6002846003811115610d5e57610d5e611828565b1415610d7457610d438b8b8b8b8b8a898961106e565b6003846003811115610d8857610d88611828565b1415610d9e57610d438b8b8b8b8b8a8989611288565b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f7274656400000000000000604482015260640161025c565b604051634f9e72ad60e11b81526001600160a01b03831690639f3ce55a908390610e1890899089908990600401611e8a565b6000604051808303818588803b158015610e3157600080fd5b505af1158015610e45573d6000803e3d6000fd5b50505050505050505050565b60606103e184846000856113dc565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b158015610e9c57600080fd5b505afa158015610eb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed49190611ec5565b9050610eea6001600160a01b038b16828b61151b565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b158015610f6057600080fd5b505af1158015610f74573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b815260040161102c959493929190611ee2565b6000604051808303818588803b15801561104557600080fd5b505af1158015611059573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b600080836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b1580156110aa57600080fd5b505afa1580156110be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e29190611ec5565b90506110f86001600160a01b038a16828a61151b565b6040516308d18d8960e21b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c821660648401528816608483015282169063234636249060a401600060405180830381600087803b15801561116257600080fd5b505af1158015611176573d6000803e3d6000fd5b505050506000308a8a8a8e8b466040516020016111fa9796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858d8b86868c6040518763ffffffff1660e01b8152600401611247959493929190611ee2565b6000604051808303818588803b15801561126057600080fd5b505af1158015611274573d6000803e3d6000fd5b50939e9d5050505050505050505050505050565b600080836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156112c457600080fd5b505afa1580156112d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112fc9190611ec5565b604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff891660648301529192509082169063de790c7e90608401600060405180830381600087803b15801561136257600080fd5b505af1158015611376573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528d811b82166034840152604883018d90528e901b1660688201526001600160c01b031960c08a811b8216607c84015246901b16608482015260009250608c0190506111fa565b6060824710156114545760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161025c565b843b6114a25760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161025c565b600080866001600160a01b031685876040516114be9190611f24565b60006040518083038185875af1925050503d80600081146114fb576040519150601f19603f3d011682016040523d82523d6000602084013e611500565b606091505b50915091506115108282866115dc565b979650505050505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b15801561156757600080fd5b505afa15801561157b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061159f9190611f40565b6115a99190611ae0565b6040516001600160a01b038516602482015260448101829052909150610b4690859063095ea7b360e01b90606401610ae2565b606083156115eb57508161037d565b8251156115fb5782518084602001fd5b8160405162461bcd60e51b815260040161025c9190611f59565b6001600160a01b0381168114610aab57600080fd5b67ffffffffffffffff81168114610aab57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561167957611679611640565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156116a8576116a8611640565b604052919050565b600082601f8301126116c157600080fd5b813567ffffffffffffffff8111156116db576116db611640565b6116ee601f8201601f191660200161167f565b81815284602083860101111561170357600080fd5b816020850160208301376000918101602001919091529392505050565b60008060006060848603121561173557600080fd5b833561174081611615565b925060208401356117508161162a565b9150604084013567ffffffffffffffff81111561176c57600080fd5b611778868287016116b0565b9150509250925092565b6000806000806060858703121561179857600080fd5b84356117a381611615565b935060208501359250604085013567ffffffffffffffff808211156117c757600080fd5b818701915087601f8301126117db57600080fd5b8135818111156117ea57600080fd5b8860208285010111156117fc57600080fd5b95989497505060200194505050565b60006020828403121561181d57600080fd5b813561037d8161162a565b634e487b7160e01b600052602160045260246000fd5b6003811061185c57634e487b7160e01b600052602160045260246000fd5b9052565b8281526040810161037d602083018461183e565b60006020828403121561188657600080fd5b813561037d81611615565b8035600481106118a057600080fd5b919050565b60008083601f8401126118b757600080fd5b50813567ffffffffffffffff8111156118cf57600080fd5b6020830191508360208260051b85010111156118ea57600080fd5b9250929050565b6000806000806000806000806000806101008b8d03121561191157600080fd5b8a3561191c81611615565b995060208b013561192c81611615565b985060408b0135975060608b01356119438161162a565b965060808b013563ffffffff8116811461195c57600080fd5b955061196a60a08c01611891565b945060c08b013567ffffffffffffffff8082111561198757600080fd5b6119938e838f016118a5565b909650945060e08d01359150808211156119ac57600080fd5b506119b98d828e016118a5565b915080935050809150509295989b9194979a5092959850565b600080600080600060a086880312156119ea57600080fd5b85356119f581611615565b94506020860135611a0581611615565b9350604086013592506060860135611a1c8161162a565b9150608086013567ffffffffffffffff811115611a3857600080fd5b611a44888289016116b0565b9150509295509295909350565b600060408284031215611a6357600080fd5b6040516040810181811067ffffffffffffffff82111715611a8657611a86611640565b6040528251611a948161162a565b8152602083015160038110611aa857600080fd5b60208201529392505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115611af357611af3611aca565b500190565b6000600019821415611b0c57611b0c611aca565b5060010190565b600067ffffffffffffffff808316818516808303821115611b3657611b36611aca565b01949350505050565b600081518084526020808501945080840160005b83811015611b6f57815187529582019590820190600101611b53565b509495945050505050565b6020808252825167ffffffffffffffff16828201528281015160806040840152805160a0840181905260009291820190839060c08601905b80831015611bdb5783516001600160a01b03168252928401926001929092019190840190611bb2565b506040870151868203601f190160608801529350611bf98185611b3f565b93505050506060840151611c1860808501826001600160a01b03169052565b509392505050565b600067ffffffffffffffff821115611c3a57611c3a611640565b5060051b60200190565b80516118a081611615565b600082601f830112611c6057600080fd5b81516020611c75611c7083611c20565b61167f565b82815260059290921b84018101918181019086841115611c9457600080fd5b8286015b84811015611caf5780518352918301918301611c98565b509695505050505050565b60006020808385031215611ccd57600080fd5b825167ffffffffffffffff80821115611ce557600080fd5b9084019060808287031215611cf957600080fd5b611d01611656565b8251611d0c8161162a565b81528284015182811115611d1f57600080fd5b8301601f81018813611d3057600080fd5b8051611d3e611c7082611c20565b81815260059190911b8201860190868101908a831115611d5d57600080fd5b928701925b82841015611d84578351611d7581611615565b82529287019290870190611d62565b8088860152505050506040830151935081841115611da157600080fd5b611dad87858501611c4f565b6040820152611dbe60608401611c44565b60608201529695505050505050565b600082821015611ddf57611ddf611aca565b500390565b815167ffffffffffffffff1681526020808301516040830191611e099084018261183e565b5092915050565b600060208284031215611e2257600080fd5b8151801515811461037d57600080fd5b60005b83811015611e4d578181015183820152602001611e35565b83811115610b465750506000910152565b60008151808452611e76816020860160208601611e32565b601f01601f19169290920160200192915050565b6001600160a01b038416815267ffffffffffffffff83166020820152606060408201526000611ebc6060830184611e5e565b95945050505050565b600060208284031215611ed757600080fd5b815161037d81611615565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a0608083015261151060a0830184611e5e565b60008251611f36818460208701611e32565b9190910192915050565b600060208284031215611f5257600080fd5b5051919050565b60208152600061037d6020830184611e5e56fea2646970667358221220b418243afaa332f25d2e0bcc930c3335f31062b7d7383b10c3627c115bb44ff064736f6c63430008090033",
}

// BatchTransferABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchTransferMetaData.ABI instead.
var BatchTransferABI = BatchTransferMetaData.ABI

// BatchTransferBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchTransferMetaData.Bin instead.
var BatchTransferBin = BatchTransferMetaData.Bin

// DeployBatchTransfer deploys a new Ethereum contract, binding an instance of BatchTransfer to it.
func DeployBatchTransfer(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address) (common.Address, *types.Transaction, *BatchTransfer, error) {
	parsed, err := BatchTransferMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchTransferBin), backend, _messageBus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchTransfer{BatchTransferCaller: BatchTransferCaller{contract: contract}, BatchTransferTransactor: BatchTransferTransactor{contract: contract}, BatchTransferFilterer: BatchTransferFilterer{contract: contract}}, nil
}

// BatchTransfer is an auto generated Go binding around an Ethereum contract.
type BatchTransfer struct {
	BatchTransferCaller     // Read-only binding to the contract
	BatchTransferTransactor // Write-only binding to the contract
	BatchTransferFilterer   // Log filterer for contract events
}

// BatchTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchTransferSession struct {
	Contract     *BatchTransfer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchTransferCallerSession struct {
	Contract *BatchTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BatchTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchTransferTransactorSession struct {
	Contract     *BatchTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BatchTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchTransferRaw struct {
	Contract *BatchTransfer // Generic contract binding to access the raw methods on
}

// BatchTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchTransferCallerRaw struct {
	Contract *BatchTransferCaller // Generic read-only contract binding to access the raw methods on
}

// BatchTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchTransferTransactorRaw struct {
	Contract *BatchTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchTransfer creates a new instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransfer(address common.Address, backend bind.ContractBackend) (*BatchTransfer, error) {
	contract, err := bindBatchTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchTransfer{BatchTransferCaller: BatchTransferCaller{contract: contract}, BatchTransferTransactor: BatchTransferTransactor{contract: contract}, BatchTransferFilterer: BatchTransferFilterer{contract: contract}}, nil
}

// NewBatchTransferCaller creates a new read-only instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransferCaller(address common.Address, caller bind.ContractCaller) (*BatchTransferCaller, error) {
	contract, err := bindBatchTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTransferCaller{contract: contract}, nil
}

// NewBatchTransferTransactor creates a new write-only instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchTransferTransactor, error) {
	contract, err := bindBatchTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchTransferTransactor{contract: contract}, nil
}

// NewBatchTransferFilterer creates a new log filterer instance of BatchTransfer, bound to a specific deployed contract.
func NewBatchTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchTransferFilterer, error) {
	contract, err := bindBatchTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchTransferFilterer{contract: contract}, nil
}

// bindBatchTransfer binds a generic wrapper to an already deployed contract.
func bindBatchTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BatchTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTransfer *BatchTransferRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchTransfer.Contract.BatchTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTransfer *BatchTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTransfer.Contract.BatchTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTransfer *BatchTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTransfer.Contract.BatchTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchTransfer *BatchTransferCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchTransfer *BatchTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchTransfer *BatchTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchTransfer.Contract.contract.Transact(opts, method, params...)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_BatchTransfer *BatchTransferCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchTransfer.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_BatchTransfer *BatchTransferSession) MessageBus() (common.Address, error) {
	return _BatchTransfer.Contract.MessageBus(&_BatchTransfer.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_BatchTransfer *BatchTransferCallerSession) MessageBus() (common.Address, error) {
	return _BatchTransfer.Contract.MessageBus(&_BatchTransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchTransfer *BatchTransferCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BatchTransfer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchTransfer *BatchTransferSession) Owner() (common.Address, error) {
	return _BatchTransfer.Contract.Owner(&_BatchTransfer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BatchTransfer *BatchTransferCallerSession) Owner() (common.Address, error) {
	return _BatchTransfer.Contract.Owner(&_BatchTransfer.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x20bff893.
//
// Solidity: function status(uint64 ) view returns(bytes32 h, uint8 status)
func (_BatchTransfer *BatchTransferCaller) Status(opts *bind.CallOpts, arg0 uint64) (struct {
	H      [32]byte
	Status uint8
}, error) {
	var out []interface{}
	err := _BatchTransfer.contract.Call(opts, &out, "status", arg0)

	outstruct := new(struct {
		H      [32]byte
		Status uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.H = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// Status is a free data retrieval call binding the contract method 0x20bff893.
//
// Solidity: function status(uint64 ) view returns(bytes32 h, uint8 status)
func (_BatchTransfer *BatchTransferSession) Status(arg0 uint64) (struct {
	H      [32]byte
	Status uint8
}, error) {
	return _BatchTransfer.Contract.Status(&_BatchTransfer.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x20bff893.
//
// Solidity: function status(uint64 ) view returns(bytes32 h, uint8 status)
func (_BatchTransfer *BatchTransferCallerSession) Status(arg0 uint64) (struct {
	H      [32]byte
	Status uint8
}, error) {
	return _BatchTransfer.Contract.Status(&_BatchTransfer.CallOpts, arg0)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x64d39f42.
//
// Solidity: function batchTransfer(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint32 _maxSlippage, uint8 _bridgeType, address[] _accounts, uint256[] _amounts) payable returns()
func (_BatchTransfer *BatchTransferTransactor) BatchTransfer(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _maxSlippage uint32, _bridgeType uint8, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "batchTransfer", _receiver, _token, _amount, _dstChainId, _maxSlippage, _bridgeType, _accounts, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x64d39f42.
//
// Solidity: function batchTransfer(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint32 _maxSlippage, uint8 _bridgeType, address[] _accounts, uint256[] _amounts) payable returns()
func (_BatchTransfer *BatchTransferSession) BatchTransfer(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _maxSlippage uint32, _bridgeType uint8, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTransfer.Contract.BatchTransfer(&_BatchTransfer.TransactOpts, _receiver, _token, _amount, _dstChainId, _maxSlippage, _bridgeType, _accounts, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x64d39f42.
//
// Solidity: function batchTransfer(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint32 _maxSlippage, uint8 _bridgeType, address[] _accounts, uint256[] _amounts) payable returns()
func (_BatchTransfer *BatchTransferTransactorSession) BatchTransfer(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _maxSlippage uint32, _bridgeType uint8, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _BatchTransfer.Contract.BatchTransfer(&_BatchTransfer.TransactOpts, _receiver, _token, _amount, _dstChainId, _maxSlippage, _bridgeType, _accounts, _amounts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessage(&_BatchTransfer.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactorSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessage(&_BatchTransfer.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessageWithTransfer(&_BatchTransfer.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessageWithTransfer(&_BatchTransfer.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessageWithTransferFallback(&_BatchTransfer.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessageWithTransferFallback(&_BatchTransfer.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessageWithTransferRefund(&_BatchTransfer.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_BatchTransfer *BatchTransferTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _BatchTransfer.Contract.ExecuteMessageWithTransferRefund(&_BatchTransfer.TransactOpts, _token, _amount, _message)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_BatchTransfer *BatchTransferTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_BatchTransfer *BatchTransferSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _BatchTransfer.Contract.SetMessageBus(&_BatchTransfer.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_BatchTransfer *BatchTransferTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _BatchTransfer.Contract.SetMessageBus(&_BatchTransfer.TransactOpts, _messageBus)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchTransfer *BatchTransferTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BatchTransfer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchTransfer *BatchTransferSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchTransfer.Contract.TransferOwnership(&_BatchTransfer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BatchTransfer *BatchTransferTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BatchTransfer.Contract.TransferOwnership(&_BatchTransfer.TransactOpts, newOwner)
}

// BatchTransferOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BatchTransfer contract.
type BatchTransferOwnershipTransferredIterator struct {
	Event *BatchTransferOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BatchTransferOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BatchTransferOwnershipTransferred)
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
		it.Event = new(BatchTransferOwnershipTransferred)
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
func (it *BatchTransferOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BatchTransferOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BatchTransferOwnershipTransferred represents a OwnershipTransferred event raised by the BatchTransfer contract.
type BatchTransferOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchTransfer *BatchTransferFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BatchTransferOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchTransfer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BatchTransferOwnershipTransferredIterator{contract: _BatchTransfer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BatchTransfer *BatchTransferFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BatchTransferOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BatchTransfer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BatchTransferOwnershipTransferred)
				if err := _BatchTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BatchTransfer *BatchTransferFilterer) ParseOwnershipTransferred(log types.Log) (*BatchTransferOwnershipTransferred, error) {
	event := new(BatchTransferOwnershipTransferred)
	if err := _BatchTransfer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossChainSwapMetaData contains all meta data concerning the CrossChainSwap contract.
var CrossChainSwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dex_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"wantToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"sendBack\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"cbrMaxSlippage\",\"type\":\"uint32\"}],\"internalType\":\"structCrossChainSwap.SwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"}],\"name\":\"startCrossChainSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ec638038062001ec68339810160408190526200003491620000b5565b6200003f3362000065565b600280546001600160a01b0319166001600160a01b0392909216919091179055620000e7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215620000c857600080fd5b81516001600160a01b0381168114620000e057600080fd5b9392505050565b611dcf80620000f76000396000f3fe6080604052600436106100b15760003560e01c8063a1a227fa11610069578063f00f39ce1161004e578063f00f39ce14610197578063f2fde38b146101aa578063f88954dc146101ca57600080fd5b8063a1a227fa14610164578063ce35dd9a1461018457600080fd5b8063547cad121161009a578063547cad12146100ec578063692058c21461010e5780638da5cb5b1461014657600080fd5b80631599d265146100b657806320be95f2146100de575b600080fd5b6100c96100c43660046116e5565b6101dd565b60405190151581526020015b60405180910390f35b6100c96100c4366004611748565b3480156100f857600080fd5b5061010c61010736600461178c565b610247565b005b34801561011a57600080fd5b5060025461012e906001600160a01b031681565b6040516001600160a01b0390911681526020016100d5565b34801561015257600080fd5b506000546001600160a01b031661012e565b34801561017057600080fd5b5060015461012e906001600160a01b031681565b6100c96101923660046117f0565b6102df565b6100c96101a53660046118c5565b610646565b3480156101b657600080fd5b5061010c6101c536600461178c565b6106ad565b61010c6101d8366004611948565b61079e565b6001546000906001600160a01b0316331461023f5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b949350505050565b3361025a6000546001600160a01b031690565b6001600160a01b0316146102b05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610236565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6001546000906001600160a01b0316331461033c5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610236565b60008280602001905181019061035291906119d7565b60025460405163095ea7b360e01b81526001600160a01b0391821660048201526024810188905291925087169063095ea7b390604401602060405180830381600087803b1580156103a257600080fd5b505af11580156103b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103da9190611a5c565b50604080516002808252606082018352600092602083019080368337019050509050868160008151811061041057610410611a79565b60200260200101906001600160a01b031690816001600160a01b03168152505081600001518160018151811061044857610448611a79565b60200260200101906001600160a01b031690816001600160a01b0316815250508160400151156105a1576001600260148282829054906101000a900467ffffffffffffffff166104989190611aa5565b825467ffffffffffffffff9182166101009390930a9283029190920219909116179055506002546040516338ed173960e01b81526000916001600160a01b0316906338ed1739906104f7908a9085908790309060001990600401611ad1565b600060405180830381600087803b15801561051157600080fd5b505af1158015610525573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261054d9190810190611b42565b905061059b836020015184600001518360018151811061056f5761056f611a79565b602002602001015189600260149054906101000a900467ffffffffffffffff1688606001516001610871565b50610638565b60025460208301516040516338ed173960e01b81526001600160a01b03909216916338ed1739916105e0918a9160009187919060001990600401611ad1565b600060405180830381600087803b1580156105fa57600080fd5b505af115801561060e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106369190810190611b42565b505b506001979650505050505050565b6001546000906001600160a01b031633146106a35760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610236565b9695505050505050565b336106c06000546001600160a01b031690565b6001600160a01b0316146107165760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610236565b6001600160a01b0381166107925760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610236565b61079b81610a4c565b50565b6001600260148282829054906101000a900467ffffffffffffffff166107c49190611aa5565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061080a333085876001600160a01b0316610aa9909392919063ffffffff16565b60008160405160200161081d9190611be8565b604051602081830303815290604052905061086886868686600260149054906101000a900467ffffffffffffffff1687606001602081019061085f9190611c4f565b87600134610b47565b50505050505050565b6000600182600381111561088757610887611c6c565b141561091a57600160009054906101000a90046001600160a01b03166001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b505afa1580156108ef573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109139190611c82565b9050610a32565b600282600381111561092e5761092e611c6c565b141561098257600160009054906101000a90046001600160a01b03166001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b600382600381111561099657610996611c6c565b14156109ea57600160009054906101000a90046001600160a01b03166001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f72746564000000000000006044820152606401610236565b610a428888888888888888610b7e565b5050505050505050565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040516001600160a01b0380851660248301528316604482015260648101829052610b419085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610d2a565b50505050565b6000610b708a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b610e14565b9a9950505050505050505050565b6001826003811115610b9257610b92611c6c565b1415610c4057610bac6001600160a01b0388168288610e9c565b60405163a5977fbb60e01b81526001600160a01b03898116600483015288811660248301526044820188905267ffffffffffffffff80881660648401528616608483015263ffffffff851660a483015282169063a5977fbb9060c4015b600060405180830381600087803b158015610c2357600080fd5b505af1158015610c37573d6000803e3d6000fd5b50505050610a42565b6002826003811115610c5457610c54611c6c565b1415610cc357610c6e6001600160a01b0388168288610e9c565b6040516308d18d8960e21b81526001600160a01b0388811660048301526024820188905267ffffffffffffffff80881660448401528a821660648401528616608483015282169063234636249060a401610c09565b6003826003811115610cd757610cd7611c6c565b14156109ea57604051636f3c863f60e11b81526001600160a01b03888116600483015260248201889052898116604483015267ffffffffffffffff8616606483015282169063de790c7e90608401610c09565b6000610d7f826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610f5d9092919063ffffffff16565b805190915015610e0f5780806020019051810190610d9d9190611a5c565b610e0f5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610236565b505050565b60006001846003811115610e2a57610e2a611c6c565b1415610e4857610e418b8b8b8b8b8b8b8a8a610f76565b9050610b70565b6002846003811115610e5c57610e5c611c6c565b1415610e7257610e418b8b8b8b8b8a8989611184565b6003846003811115610e8657610e86611c6c565b14156109ea57610e418b8b8b8b8b8a898961139e565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b158015610ee857600080fd5b505afa158015610efc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f209190611c9f565b610f2a9190611cb8565b6040516001600160a01b038516602482015260448101829052909150610b4190859063095ea7b360e01b90606401610add565b6060610f6c84846000856114f2565b90505b9392505050565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b158015610fb257600080fd5b505afa158015610fc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fea9190611c82565b90506110006001600160a01b038b16828b610e9c565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b15801561107657600080fd5b505af115801561108a573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b8152600401611142959493929190611d28565b6000604051808303818588803b15801561115b57600080fd5b505af115801561116f573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b600080836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b1580156111c057600080fd5b505afa1580156111d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111f89190611c82565b905061120e6001600160a01b038a16828a610e9c565b6040516308d18d8960e21b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c821660648401528816608483015282169063234636249060a401600060405180830381600087803b15801561127857600080fd5b505af115801561128c573d6000803e3d6000fd5b505050506000308a8a8a8e8b466040516020016113109796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858d8b86868c6040518763ffffffff1660e01b815260040161135d959493929190611d28565b6000604051808303818588803b15801561137657600080fd5b505af115801561138a573d6000803e3d6000fd5b50939e9d5050505050505050505050505050565b600080836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156113da57600080fd5b505afa1580156113ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114129190611c82565b604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff891660648301529192509082169063de790c7e90608401600060405180830381600087803b15801561147857600080fd5b505af115801561148c573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528d811b82166034840152604883018d90528e901b1660688201526001600160c01b031960c08a811b8216607c84015246901b16608482015260009250608c019050611310565b60608247101561156a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610236565b843b6115b85760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610236565b600080866001600160a01b031685876040516115d49190611d6a565b60006040518083038185875af1925050503d8060008114611611576040519150601f19603f3d011682016040523d82523d6000602084013e611616565b606091505b5091509150611626828286611631565b979650505050505050565b60608315611640575081610f6f565b8251156116505782518084602001fd5b8160405162461bcd60e51b81526004016102369190611d86565b6001600160a01b038116811461079b57600080fd5b803567ffffffffffffffff8116811461169757600080fd5b919050565b60008083601f8401126116ae57600080fd5b50813567ffffffffffffffff8111156116c657600080fd5b6020830191508360208285010111156116de57600080fd5b9250929050565b600080600080606085870312156116fb57600080fd5b84356117068161166a565b93506117146020860161167f565b9250604085013567ffffffffffffffff81111561173057600080fd5b61173c8782880161169c565b95989497509550505050565b6000806000806060858703121561175e57600080fd5b84356117698161166a565b935060208501359250604085013567ffffffffffffffff81111561173057600080fd5b60006020828403121561179e57600080fd5b8135610f6f8161166a565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156117e8576117e86117a9565b604052919050565b600080600080600060a0868803121561180857600080fd5b85356118138161166a565b94506020868101356118248161166a565b9450604087013593506118396060880161167f565b9250608087013567ffffffffffffffff8082111561185657600080fd5b818901915089601f83011261186a57600080fd5b81358181111561187c5761187c6117a9565b61188e601f8201601f191685016117bf565b91508082528a848285010111156118a457600080fd5b80848401858401376000848284010152508093505050509295509295909350565b60008060008060008060a087890312156118de57600080fd5b86356118e98161166a565b955060208701356118f98161166a565b94506040870135935061190e6060880161167f565b9250608087013567ffffffffffffffff81111561192a57600080fd5b61193689828a0161169c565b979a9699509497509295939492505050565b600080600080600085870361010081121561196257600080fd5b863561196d8161166a565b9550602087013561197d8161166a565b9450604087013593506119926060880161167f565b92506080607f19820112156119a657600080fd5b506080860190509295509295909350565b801515811461079b57600080fd5b63ffffffff8116811461079b57600080fd5b6000608082840312156119e957600080fd5b6040516080810181811067ffffffffffffffff82111715611a0c57611a0c6117a9565b6040528251611a1a8161166a565b81526020830151611a2a8161166a565b60208201526040830151611a3d816119b7565b60408201526060830151611a50816119c5565b60608201529392505050565b600060208284031215611a6e57600080fd5b8151610f6f816119b7565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600067ffffffffffffffff808316818516808303821115611ac857611ac8611a8f565b01949350505050565b600060a082018783526020878185015260a0604085015281875180845260c086019150828901935060005b81811015611b215784516001600160a01b031683529383019391830191600101611afc565b50506001600160a01b03969096166060850152505050608001529392505050565b60006020808385031215611b5557600080fd5b825167ffffffffffffffff80821115611b6d57600080fd5b818501915085601f830112611b8157600080fd5b815181811115611b9357611b936117a9565b8060051b9150611ba48483016117bf565b8181529183018401918481019088841115611bbe57600080fd5b938501935b83851015611bdc57845182529385019390850190611bc3565b98975050505050505050565b608081018235611bf78161166a565b6001600160a01b039081168352602084013590611c138261166a565b1660208301526040830135611c27816119b7565b151560408301526060830135611c3c816119c5565b63ffffffff811660608401525092915050565b600060208284031215611c6157600080fd5b8135610f6f816119c5565b634e487b7160e01b600052602160045260246000fd5b600060208284031215611c9457600080fd5b8151610f6f8161166a565b600060208284031215611cb157600080fd5b5051919050565b60008219821115611ccb57611ccb611a8f565b500190565b60005b83811015611ceb578181015183820152602001611cd3565b83811115610b415750506000910152565b60008151808452611d14816020860160208601611cd0565b601f01601f19169290920160200192915050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a0608083015261162660a0830184611cfc565b60008251611d7c818460208701611cd0565b9190910192915050565b602081526000610f6f6020830184611cfc56fea26469706673582212201ee13e8b5621b3c46e580142e470d5dd3bcffef95f9484dccd43f78738a46b6764736f6c63430008090033",
}

// CrossChainSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossChainSwapMetaData.ABI instead.
var CrossChainSwapABI = CrossChainSwapMetaData.ABI

// CrossChainSwapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrossChainSwapMetaData.Bin instead.
var CrossChainSwapBin = CrossChainSwapMetaData.Bin

// DeployCrossChainSwap deploys a new Ethereum contract, binding an instance of CrossChainSwap to it.
func DeployCrossChainSwap(auth *bind.TransactOpts, backend bind.ContractBackend, dex_ common.Address) (common.Address, *types.Transaction, *CrossChainSwap, error) {
	parsed, err := CrossChainSwapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrossChainSwapBin), backend, dex_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrossChainSwap{CrossChainSwapCaller: CrossChainSwapCaller{contract: contract}, CrossChainSwapTransactor: CrossChainSwapTransactor{contract: contract}, CrossChainSwapFilterer: CrossChainSwapFilterer{contract: contract}}, nil
}

// CrossChainSwap is an auto generated Go binding around an Ethereum contract.
type CrossChainSwap struct {
	CrossChainSwapCaller     // Read-only binding to the contract
	CrossChainSwapTransactor // Write-only binding to the contract
	CrossChainSwapFilterer   // Log filterer for contract events
}

// CrossChainSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossChainSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossChainSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossChainSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossChainSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossChainSwapSession struct {
	Contract     *CrossChainSwap   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossChainSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossChainSwapCallerSession struct {
	Contract *CrossChainSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CrossChainSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossChainSwapTransactorSession struct {
	Contract     *CrossChainSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CrossChainSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossChainSwapRaw struct {
	Contract *CrossChainSwap // Generic contract binding to access the raw methods on
}

// CrossChainSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossChainSwapCallerRaw struct {
	Contract *CrossChainSwapCaller // Generic read-only contract binding to access the raw methods on
}

// CrossChainSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossChainSwapTransactorRaw struct {
	Contract *CrossChainSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossChainSwap creates a new instance of CrossChainSwap, bound to a specific deployed contract.
func NewCrossChainSwap(address common.Address, backend bind.ContractBackend) (*CrossChainSwap, error) {
	contract, err := bindCrossChainSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossChainSwap{CrossChainSwapCaller: CrossChainSwapCaller{contract: contract}, CrossChainSwapTransactor: CrossChainSwapTransactor{contract: contract}, CrossChainSwapFilterer: CrossChainSwapFilterer{contract: contract}}, nil
}

// NewCrossChainSwapCaller creates a new read-only instance of CrossChainSwap, bound to a specific deployed contract.
func NewCrossChainSwapCaller(address common.Address, caller bind.ContractCaller) (*CrossChainSwapCaller, error) {
	contract, err := bindCrossChainSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainSwapCaller{contract: contract}, nil
}

// NewCrossChainSwapTransactor creates a new write-only instance of CrossChainSwap, bound to a specific deployed contract.
func NewCrossChainSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossChainSwapTransactor, error) {
	contract, err := bindCrossChainSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossChainSwapTransactor{contract: contract}, nil
}

// NewCrossChainSwapFilterer creates a new log filterer instance of CrossChainSwap, bound to a specific deployed contract.
func NewCrossChainSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossChainSwapFilterer, error) {
	contract, err := bindCrossChainSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossChainSwapFilterer{contract: contract}, nil
}

// bindCrossChainSwap binds a generic wrapper to an already deployed contract.
func bindCrossChainSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossChainSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainSwap *CrossChainSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainSwap.Contract.CrossChainSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainSwap *CrossChainSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.CrossChainSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainSwap *CrossChainSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.CrossChainSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossChainSwap *CrossChainSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossChainSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossChainSwap *CrossChainSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossChainSwap *CrossChainSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.contract.Transact(opts, method, params...)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_CrossChainSwap *CrossChainSwapCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainSwap.contract.Call(opts, &out, "dex")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_CrossChainSwap *CrossChainSwapSession) Dex() (common.Address, error) {
	return _CrossChainSwap.Contract.Dex(&_CrossChainSwap.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_CrossChainSwap *CrossChainSwapCallerSession) Dex() (common.Address, error) {
	return _CrossChainSwap.Contract.Dex(&_CrossChainSwap.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_CrossChainSwap *CrossChainSwapCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainSwap.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_CrossChainSwap *CrossChainSwapSession) MessageBus() (common.Address, error) {
	return _CrossChainSwap.Contract.MessageBus(&_CrossChainSwap.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_CrossChainSwap *CrossChainSwapCallerSession) MessageBus() (common.Address, error) {
	return _CrossChainSwap.Contract.MessageBus(&_CrossChainSwap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossChainSwap *CrossChainSwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossChainSwap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossChainSwap *CrossChainSwapSession) Owner() (common.Address, error) {
	return _CrossChainSwap.Contract.Owner(&_CrossChainSwap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrossChainSwap *CrossChainSwapCallerSession) Owner() (common.Address, error) {
	return _CrossChainSwap.Contract.Owner(&_CrossChainSwap.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessage(&_CrossChainSwap.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactorSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessage(&_CrossChainSwap.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address , address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, arg0 common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.contract.Transact(opts, "executeMessageWithTransfer", arg0, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address , address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapSession) ExecuteMessageWithTransfer(arg0 common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessageWithTransfer(&_CrossChainSwap.TransactOpts, arg0, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address , address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactorSession) ExecuteMessageWithTransfer(arg0 common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessageWithTransfer(&_CrossChainSwap.TransactOpts, arg0, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessageWithTransferFallback(&_CrossChainSwap.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessageWithTransferFallback(&_CrossChainSwap.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessageWithTransferRefund(&_CrossChainSwap.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_CrossChainSwap *CrossChainSwapTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.ExecuteMessageWithTransferRefund(&_CrossChainSwap.TransactOpts, _token, _amount, _message)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_CrossChainSwap *CrossChainSwapTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _CrossChainSwap.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_CrossChainSwap *CrossChainSwapSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.SetMessageBus(&_CrossChainSwap.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_CrossChainSwap *CrossChainSwapTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.SetMessageBus(&_CrossChainSwap.TransactOpts, _messageBus)
}

// StartCrossChainSwap is a paid mutator transaction binding the contract method 0xf88954dc.
//
// Solidity: function startCrossChainSwap(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, (address,address,bool,uint32) swapInfo) payable returns()
func (_CrossChainSwap *CrossChainSwapTransactor) StartCrossChainSwap(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, swapInfo CrossChainSwapSwapInfo) (*types.Transaction, error) {
	return _CrossChainSwap.contract.Transact(opts, "startCrossChainSwap", _receiver, _token, _amount, _dstChainId, swapInfo)
}

// StartCrossChainSwap is a paid mutator transaction binding the contract method 0xf88954dc.
//
// Solidity: function startCrossChainSwap(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, (address,address,bool,uint32) swapInfo) payable returns()
func (_CrossChainSwap *CrossChainSwapSession) StartCrossChainSwap(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, swapInfo CrossChainSwapSwapInfo) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.StartCrossChainSwap(&_CrossChainSwap.TransactOpts, _receiver, _token, _amount, _dstChainId, swapInfo)
}

// StartCrossChainSwap is a paid mutator transaction binding the contract method 0xf88954dc.
//
// Solidity: function startCrossChainSwap(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, (address,address,bool,uint32) swapInfo) payable returns()
func (_CrossChainSwap *CrossChainSwapTransactorSession) StartCrossChainSwap(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, swapInfo CrossChainSwapSwapInfo) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.StartCrossChainSwap(&_CrossChainSwap.TransactOpts, _receiver, _token, _amount, _dstChainId, swapInfo)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossChainSwap *CrossChainSwapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CrossChainSwap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossChainSwap *CrossChainSwapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.TransferOwnership(&_CrossChainSwap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CrossChainSwap *CrossChainSwapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CrossChainSwap.Contract.TransferOwnership(&_CrossChainSwap.TransactOpts, newOwner)
}

// CrossChainSwapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CrossChainSwap contract.
type CrossChainSwapOwnershipTransferredIterator struct {
	Event *CrossChainSwapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrossChainSwapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossChainSwapOwnershipTransferred)
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
		it.Event = new(CrossChainSwapOwnershipTransferred)
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
func (it *CrossChainSwapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossChainSwapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossChainSwapOwnershipTransferred represents a OwnershipTransferred event raised by the CrossChainSwap contract.
type CrossChainSwapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossChainSwap *CrossChainSwapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrossChainSwapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossChainSwap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrossChainSwapOwnershipTransferredIterator{contract: _CrossChainSwap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CrossChainSwap *CrossChainSwapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrossChainSwapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CrossChainSwap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossChainSwapOwnershipTransferred)
				if err := _CrossChainSwap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CrossChainSwap *CrossChainSwapFilterer) ParseOwnershipTransferred(log types.Log) (*CrossChainSwapOwnershipTransferred, error) {
	event := new(CrossChainSwapOwnershipTransferred)
	if err := _CrossChainSwap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMessageReceiverAppMetaData contains all meta data concerning the IMessageReceiverApp contract.
var IMessageReceiverAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IMessageReceiverAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageReceiverAppMetaData.ABI instead.
var IMessageReceiverAppABI = IMessageReceiverAppMetaData.ABI

// IMessageReceiverApp is an auto generated Go binding around an Ethereum contract.
type IMessageReceiverApp struct {
	IMessageReceiverAppCaller     // Read-only binding to the contract
	IMessageReceiverAppTransactor // Write-only binding to the contract
	IMessageReceiverAppFilterer   // Log filterer for contract events
}

// IMessageReceiverAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageReceiverAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageReceiverAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageReceiverAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageReceiverAppSession struct {
	Contract     *IMessageReceiverApp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageReceiverAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageReceiverAppCallerSession struct {
	Contract *IMessageReceiverAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageReceiverAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageReceiverAppTransactorSession struct {
	Contract     *IMessageReceiverAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageReceiverAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageReceiverAppRaw struct {
	Contract *IMessageReceiverApp // Generic contract binding to access the raw methods on
}

// IMessageReceiverAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageReceiverAppCallerRaw struct {
	Contract *IMessageReceiverAppCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageReceiverAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageReceiverAppTransactorRaw struct {
	Contract *IMessageReceiverAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageReceiverApp creates a new instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverApp(address common.Address, backend bind.ContractBackend) (*IMessageReceiverApp, error) {
	contract, err := bindIMessageReceiverApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverApp{IMessageReceiverAppCaller: IMessageReceiverAppCaller{contract: contract}, IMessageReceiverAppTransactor: IMessageReceiverAppTransactor{contract: contract}, IMessageReceiverAppFilterer: IMessageReceiverAppFilterer{contract: contract}}, nil
}

// NewIMessageReceiverAppCaller creates a new read-only instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppCaller(address common.Address, caller bind.ContractCaller) (*IMessageReceiverAppCaller, error) {
	contract, err := bindIMessageReceiverApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppCaller{contract: contract}, nil
}

// NewIMessageReceiverAppTransactor creates a new write-only instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageReceiverAppTransactor, error) {
	contract, err := bindIMessageReceiverApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppTransactor{contract: contract}, nil
}

// NewIMessageReceiverAppFilterer creates a new log filterer instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageReceiverAppFilterer, error) {
	contract, err := bindIMessageReceiverApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppFilterer{contract: contract}, nil
}

// bindIMessageReceiverApp binds a generic wrapper to an already deployed contract.
func bindIMessageReceiverApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageReceiverAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageReceiverApp *IMessageReceiverAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageReceiverApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageReceiverApp *IMessageReceiverAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageReceiverApp *IMessageReceiverAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.contract.Transact(opts, method, params...)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_IMessageReceiverApp.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_IMessageReceiverApp.TransactOpts, _token, _amount, _message)
}

// ISwapTokenMetaData contains all meta data concerning the ISwapToken contract.
var ISwapTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ISwapTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapTokenMetaData.ABI instead.
var ISwapTokenABI = ISwapTokenMetaData.ABI

// ISwapToken is an auto generated Go binding around an Ethereum contract.
type ISwapToken struct {
	ISwapTokenCaller     // Read-only binding to the contract
	ISwapTokenTransactor // Write-only binding to the contract
	ISwapTokenFilterer   // Log filterer for contract events
}

// ISwapTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapTokenSession struct {
	Contract     *ISwapToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapTokenCallerSession struct {
	Contract *ISwapTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ISwapTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapTokenTransactorSession struct {
	Contract     *ISwapTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ISwapTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapTokenRaw struct {
	Contract *ISwapToken // Generic contract binding to access the raw methods on
}

// ISwapTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapTokenCallerRaw struct {
	Contract *ISwapTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapTokenTransactorRaw struct {
	Contract *ISwapTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapToken creates a new instance of ISwapToken, bound to a specific deployed contract.
func NewISwapToken(address common.Address, backend bind.ContractBackend) (*ISwapToken, error) {
	contract, err := bindISwapToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapToken{ISwapTokenCaller: ISwapTokenCaller{contract: contract}, ISwapTokenTransactor: ISwapTokenTransactor{contract: contract}, ISwapTokenFilterer: ISwapTokenFilterer{contract: contract}}, nil
}

// NewISwapTokenCaller creates a new read-only instance of ISwapToken, bound to a specific deployed contract.
func NewISwapTokenCaller(address common.Address, caller bind.ContractCaller) (*ISwapTokenCaller, error) {
	contract, err := bindISwapToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapTokenCaller{contract: contract}, nil
}

// NewISwapTokenTransactor creates a new write-only instance of ISwapToken, bound to a specific deployed contract.
func NewISwapTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapTokenTransactor, error) {
	contract, err := bindISwapToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapTokenTransactor{contract: contract}, nil
}

// NewISwapTokenFilterer creates a new log filterer instance of ISwapToken, bound to a specific deployed contract.
func NewISwapTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapTokenFilterer, error) {
	contract, err := bindISwapToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapTokenFilterer{contract: contract}, nil
}

// bindISwapToken binds a generic wrapper to an already deployed contract.
func bindISwapToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapToken *ISwapTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapToken.Contract.ISwapTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapToken *ISwapTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapToken.Contract.ISwapTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapToken *ISwapTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapToken.Contract.ISwapTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapToken *ISwapTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapToken *ISwapTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapToken *ISwapTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapToken.Contract.contract.Transact(opts, method, params...)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 , uint256 , address[] , address , uint256 ) returns(uint256[])
func (_ISwapToken *ISwapTokenTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _ISwapToken.contract.Transact(opts, "swapExactTokensForTokens", arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 , uint256 , address[] , address , uint256 ) returns(uint256[])
func (_ISwapToken *ISwapTokenSession) SwapExactTokensForTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _ISwapToken.Contract.SwapExactTokensForTokens(&_ISwapToken.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 , uint256 , address[] , address , uint256 ) returns(uint256[])
func (_ISwapToken *ISwapTokenTransactorSession) SwapExactTokensForTokens(arg0 *big.Int, arg1 *big.Int, arg2 []common.Address, arg3 common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _ISwapToken.Contract.SwapExactTokensForTokens(&_ISwapToken.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// MessageBusMetaData contains all meta data concerning the MessageBus contract.
var MessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"}],\"internalType\":\"structMessageBusReceiver.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeeBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200279e3803806200279e8339810160408190526200003491620000fa565b82828286620000433362000091565b6001600160a01b03908116608052600580546001600160a01b031990811695831695909517905560068054851693821693909317909255600780549093169116179055506200016292505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0381168114620000f757600080fd5b50565b600080600080608085870312156200011157600080fd5b84516200011e81620000e1565b60208601519094506200013181620000e1565b60408601519093506200014481620000e1565b60608601519092506200015781620000e1565b939692955090935050565b6080516126196200018560003960008181610384015261064401526126196000f3fe6080604052600436106101805760003560e01c80638da5cb5b116100d6578063cd2abd661161007f578063e2c1ed2511610059578063e2c1ed2514610423578063f2fde38b14610443578063f60bbe2a1461046357600080fd5b8063cd2abd66146103a6578063d8257d17146103e3578063dfa2dbaf1461040357600080fd5b80639f3ce55a116100b05780639f3ce55a1461034c578063a22322131461035f578063ccf2683b1461037257600080fd5b80638da5cb5b146102f857806395e911a8146103165780639b05a7751461032c57600080fd5b80635335dca2116101385780635b3e5f50116101125780635b3e5f5014610280578063654317bf146102ad57806382980dc4146102c057600080fd5b80635335dca21461021a578063588be02b1461024d578063588df4161461026d57600080fd5b8063184b955911610169578063184b9559146101c75780632ff4c411146101e75780634289fbb31461020757600080fd5b806303cbfe661461018557806306c28bd6146101a7575b600080fd5b34801561019157600080fd5b506101a56101a0366004611c19565b610479565b005b3480156101b357600080fd5b506101a56101c2366004611c34565b610509565b3480156101d357600080fd5b506101a56101e2366004611c4d565b610577565b3480156101f357600080fd5b506101a5610202366004611cdc565b61058f565b6101a5610215366004611dd2565b6107e5565b34801561022657600080fd5b5061023a610235366004611e4a565b61088c565b6040519081526020015b60405180910390f35b34801561025957600080fd5b506101a5610268366004611c19565b6108b2565b6101a561027b366004611e8c565b61093d565b34801561028c57600080fd5b5061023a61029b366004611c19565b60036020526000908152604090205481565b6101a56102bb366004611f7d565b610b6a565b3480156102cc57600080fd5b506005546102e0906001600160a01b031681565b6040516001600160a01b039091168152602001610244565b34801561030457600080fd5b506000546001600160a01b03166102e0565b34801561032257600080fd5b5061023a60015481565b34801561033857600080fd5b506101a5610347366004611c19565b610d65565b6101a561035a366004612044565b610df0565b6101a561036d366004611e8c565b610e91565b34801561037e57600080fd5b506102e07f000000000000000000000000000000000000000000000000000000000000000081565b3480156103b257600080fd5b506103d66103c1366004611c34565b60046020526000908152604090205460ff1681565b60405161024491906120c8565b3480156103ef57600080fd5b506007546102e0906001600160a01b031681565b34801561040f57600080fd5b506006546102e0906001600160a01b031681565b34801561042f57600080fd5b506101a561043e366004611c34565b61105a565b34801561044f57600080fd5b506101a561045e366004611c19565b6110c8565b34801561046f57600080fd5b5061023a60025481565b3361048c6000546001600160a01b031690565b6001600160a01b0316146104e75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b3361051c6000546001600160a01b031690565b6001600160a01b0316146105725760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104de565b600155565b61057f6111b9565b61058a83838361121d565b505050565b600046306040516020016105e592919091825260601b6bffffffffffffffffffffffff191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6bffffffffffffffffffffffff19168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc229161068b918b908b908b908b908b908b906078016121c5565b60006040518083038186803b1580156106a357600080fd5b505afa1580156106b7573d6000803e3d6000fd5b505050506001600160a01b0389166000908152600360205260408120546106de908a6122d5565b9050600081116107305760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f2077697468647261770000000000000060448201526064016104de565b60008a6001600160a01b03168261c35090604051600060405180830381858888f193505050503d8060008114610782576040519150601f19603f3d011682016040523d82523d6000602084013e610787565b606091505b50509050806107d85760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f207769746864726177206665650000000000000000000060448201526064016104de565b5050505050505050505050565b60006107f1838361088c565b9050803410156108365760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016104de565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f668888888888883460405161087b97969594939291906122ec565b60405180910390a250505050505050565b60025460009061089c9083612339565b6001546108a99190612358565b90505b92915050565b336108c56000546001600160a01b031690565b6001600160a01b03161461091b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104de565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b6000610948886112b5565b90506000808281526004602052604090205460ff16600381111561096e5761096e61209e565b146109bb5760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c72656164792065786563757465640000000000000060448201526064016104de565b60004630604051602001610a1192919091825260601b6bffffffffffffffffffffffff191660208201527f4d657373616765576974685472616e73666572526566756e64000000000000006034820152604d0190565b604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e604051602001610a649493929190612370565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610a9b97969594939291906121c5565b60006040518083038186803b158015610ab357600080fd5b505afa158015610ac7573d6000803e3d6000fd5b50505050600080610ad98b8e8e61190f565b90508015610aea5760019150610aef565b600291505b6000848152600460205260409020805483919060ff19166001836003811115610b1a57610b1a61209e565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e760008584604051610b5393929190612391565b60405180910390a150505050505050505050505050565b6000610b77888b8b611a46565b90506000808281526004602052604090205460ff166003811115610b9d57610b9d61209e565b14610bea5760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c7265616479206578656375746564000000000000000060448201526064016104de565b60004630604051602001610c4092919091825260601b6bffffffffffffffffffffffff191660208201527f4d657373616765000000000000000000000000000000000000000000000000006034820152603b0190565b60408051601f1981840301815282825280516020918201206005549184018190528383018690528251808503840181526060850193849052633416de1160e11b90935293506001600160a01b03169163682dbc2291610cad918c908c908c908c908c908c906064016121c5565b60006040518083038186803b158015610cc557600080fd5b505afa158015610cd9573d6000803e3d6000fd5b50505050600080610ceb8b8e8e611aad565b90508015610cfc5760019150610d01565b600291505b6000848152600460205260409020805483919060ff19166001836003811115610d2c57610d2c61209e565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e760018584604051610b5393929190612391565b33610d786000546001600160a01b031690565b6001600160a01b031614610dce5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104de565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6000610dfc838361088c565b905080341015610e415760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016104de565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e48686868634604051610e829594939291906123c3565b60405180910390a25050505050565b6000610e9c886112b5565b90506000808281526004602052604090205460ff166003811115610ec257610ec261209e565b14610f0f5760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c72656164792065786563757465640000000000000060448201526064016104de565b60004630604051602001610f6592919091825260601b6bffffffffffffffffffffffff191660208201527f4d657373616765576974685472616e7366657200000000000000000000000000603482015260470190565b604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e604051602001610fb89493929190612370565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610fef97969594939291906121c5565b60006040518083038186803b15801561100757600080fd5b505afa15801561101b573d6000803e3d6000fd5b5050505060008061102d8b8e8e611b05565b9050801561103e5760019150610aef565b6110498b8e8e611b77565b90508015610aea5760039150610aef565b3361106d6000546001600160a01b031690565b6001600160a01b0316146110c35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104de565b600255565b336110db6000546001600160a01b031690565b6001600160a01b0316146111315760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104de565b6001600160a01b0381166111ad5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104de565b6111b681611bad565b50565b6000546001600160a01b0316156112125760405162461bcd60e51b815260206004820152601160248201527f6f776e657220616c72656164792073657400000000000000000000000000000060448201526064016104de565b61121b33611bad565b565b6005546001600160a01b0316156112765760405162461bcd60e51b815260206004820152601b60248201527f6c697175696469747942726964676520616c726561647920736574000000000060448201526064016104de565b600580546001600160a01b039485166001600160a01b031991821617909155600680549385169382169390931790925560078054919093169116179055565b6000808060016112c860208601866123fe565b60048111156112d9576112d961209e565b1415611479576112ef6040850160208601611c19565b6112ff6060860160408701611c19565b61130f6080870160608801611c19565b608087013561132460e0890160c08a0161241f565b6040516bffffffffffffffffffffffff19606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600554633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b9060240160206040518083038186803b1580156113eb57600080fd5b505afa1580156113ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114239190612449565b15156001146114745760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f742065786973740000000000000000000060448201526064016104de565b6118da565b600261148860208601866123fe565b60048111156114995761149961209e565b141561160b57466114b060c0860160a0870161241f565b6114c06060870160408801611c19565b6114d06080880160608901611c19565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526bffffffffffffffffffffffff19606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600554631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab4289060240160206040518083038186803b15801561158257600080fd5b505afa158015611596573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ba9190612449565b15156001146114745760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f742065786973740000000000000060448201526064016104de565b600361161a60208601866123fe565b600481111561162b5761162b61209e565b14806116545750600461164160208601866123fe565b60048111156116525761165261209e565b145b156118da576116696060850160408601611c19565b6116796080860160608701611c19565b608086013561168e6040880160208901611c19565b61169e60e0890160c08a0161241f565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f1981840301815291905280516020909101209150600361172360208601866123fe565b60048111156117345761173461209e565b141561180a57506006546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e647259060240160206040518083038186803b15801561178157600080fd5b505afa158015611795573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117b99190612449565b15156001146114745760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f74206578697374000000000000000000000060448201526064016104de565b506007546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e647259060240160206040518083038186803b15801561185157600080fd5b505afa158015611865573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118899190612449565b15156001146118da5760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f742065786973740000000000000060448201526064016104de565b600081836040516020016118f093929190612482565b6040516020818303038152906040528051906020012092505050919050565b600080806119236060870160408801611c19565b6001600160a01b03163463105f4af960e11b61194560808a0160608b01611c19565b8960800135898960405160240161195f94939291906124b3565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516119ca91906124e6565b60006040518083038185875af1925050503d8060008114611a07576040519150601f19603f3d011682016040523d82523d6000602084013e611a0c565b606091505b50915091508115611a3857600081806020019051810190611a2d9190612449565b9350611a3f92505050565b6000925050505b9392505050565b60006001611a576020860186611c19565b611a676040870160208801611c19565b611a77606088016040890161241f565b8686604051602001611a8e96959493929190612502565b6040516020818303038152906040528051906020012090509392505050565b60008080611ac16040870160208801611c19565b6001600160a01b031634631599d26560e01b611ae060208a018a611c19565b611af060608b0160408c0161241f565b898960405160240161195f9493929190612561565b60008080611b196060870160408801611c19565b6001600160a01b03163463671aeecd60e11b611b3b60408a0160208b01611c19565b611b4b60808b0160608c01611c19565b60808b0135611b6060e08d0160c08e0161241f565b8b8b60405160240161195f96959493929190612594565b60008080611b8b6060870160408801611c19565b6001600160a01b0316346378079ce760e11b611b3b60408a0160208b01611c19565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b0381168114611c1457600080fd5b919050565b600060208284031215611c2b57600080fd5b6108a982611bfd565b600060208284031215611c4657600080fd5b5035919050565b600080600060608486031215611c6257600080fd5b611c6b84611bfd565b9250611c7960208501611bfd565b9150611c8760408501611bfd565b90509250925092565b60008083601f840112611ca257600080fd5b50813567ffffffffffffffff811115611cba57600080fd5b6020830191508360208260051b8501011115611cd557600080fd5b9250929050565b60008060008060008060008060a0898b031215611cf857600080fd5b611d0189611bfd565b975060208901359650604089013567ffffffffffffffff80821115611d2557600080fd5b611d318c838d01611c90565b909850965060608b0135915080821115611d4a57600080fd5b611d568c838d01611c90565b909650945060808b0135915080821115611d6f57600080fd5b50611d7c8b828c01611c90565b999c989b5096995094979396929594505050565b60008083601f840112611da257600080fd5b50813567ffffffffffffffff811115611dba57600080fd5b602083019150836020828501011115611cd557600080fd5b60008060008060008060a08789031215611deb57600080fd5b611df487611bfd565b955060208701359450611e0960408801611bfd565b935060608701359250608087013567ffffffffffffffff811115611e2c57600080fd5b611e3889828a01611d90565b979a9699509497509295939492505050565b60008060208385031215611e5d57600080fd5b823567ffffffffffffffff811115611e7457600080fd5b611e8085828601611d90565b90969095509350505050565b6000806000806000806000806000898b03610180811215611eac57600080fd5b8a3567ffffffffffffffff80821115611ec457600080fd5b611ed08e838f01611d90565b909c509a508a9150610100601f1984011215611eeb57600080fd5b60208d0199506101208d0135925080831115611f0657600080fd5b611f128e848f01611c90565b90995097506101408d0135925088915080831115611f2f57600080fd5b611f3b8e848f01611c90565b90975095506101608d0135925086915080831115611f5857600080fd5b5050611f668c828d01611c90565b915080935050809150509295985092959850929598565b6000806000806000806000806000898b0360e0811215611f9c57600080fd5b8a3567ffffffffffffffff80821115611fb457600080fd5b611fc08e838f01611d90565b909c509a508a91506060601f1984011215611fda57600080fd5b60208d01995060808d0135925080831115611ff457600080fd5b6120008e848f01611c90565b909950975060a08d013592508891508083111561201c57600080fd5b6120288e848f01611c90565b909750955060c08d0135925086915080831115611f5857600080fd5b6000806000806060858703121561205a57600080fd5b61206385611bfd565b935060208501359250604085013567ffffffffffffffff81111561208657600080fd5b61209287828801611d90565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b600481106120c4576120c461209e565b9052565b602081016108ac82846120b4565b60005b838110156120f15781810151838201526020016120d9565b83811115612100576000848401525b50505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8183526000602080850194508260005b8581101561216b576001600160a01b0361215883611bfd565b168752958201959082019060010161213f565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156121a857600080fd5b8260051b8083602087013760009401602001938452509192915050565b60808152600088518060808401526121e48160a0850160208d016120d6565b601f01601f1916820182810360a090810160208501528101889052600588901b810160c09081019082018a60005b8b8110156122855784840360bf190183528135368e9003601e1901811261223857600080fd5b8d01803567ffffffffffffffff81111561225157600080fd5b8036038f131561226057600080fd5b61226e868260208501612106565b955050506020928301929190910190600101612212565b505050838103604085015261229b81888a61212f565b91505082810360608401526122b1818587612176565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156122e7576122e76122bf565b500390565b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c0608083015261232560c083018587612106565b90508260a083015298975050505050505050565b6000816000190483118215151615612353576123536122bf565b500290565b6000821982111561236b5761236b6122bf565b500190565b84815283602082015281836040830137600091016040019081529392505050565b60608101600285106123a5576123a561209e565b8482528360208301526123bb60408301846120b4565b949350505050565b6001600160a01b03861681528460208201526080604082015260006123ec608083018587612106565b90508260608301529695505050505050565b60006020828403121561241057600080fd5b813560058110611a3f57600080fd5b60006020828403121561243157600080fd5b813567ffffffffffffffff81168114611a3f57600080fd5b60006020828403121561245b57600080fd5b81518015158114611a3f57600080fd5b6002811061247b5761247b61209e565b60f81b9052565b61248c818561246b565b60609290921b6bffffffffffffffffffffffff191660018301526015820152603501919050565b6001600160a01b03851681528360208201526060604082015260006124dc606083018486612106565b9695505050505050565b600082516124f88184602087016120d6565b9190910192915050565b61250c818861246b565b60006bffffffffffffffffffffffff19808860601b166001840152808760601b166015840152506001600160c01b03198560c01b16602983015282846031840137506000910160310190815295945050505050565b6001600160a01b038516815267ffffffffffffffff841660208201526060604082015260006124dc606083018486612106565b60006001600160a01b03808916835280881660208401525085604083015267ffffffffffffffff8516606083015260a060808301526125d760a083018486612106565b9897505050505050505056fea2646970667358221220dac74e5ae7de5e0176f2ac64dc066345aebed33156065fd76f868e40e9a5286e64736f6c63430008090033",
}

// MessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusMetaData.ABI instead.
var MessageBusABI = MessageBusMetaData.ABI

// MessageBusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusMetaData.Bin instead.
var MessageBusBin = MessageBusMetaData.Bin

// DeployMessageBus deploys a new Ethereum contract, binding an instance of MessageBus to it.
func DeployMessageBus(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address) (common.Address, *types.Transaction, *MessageBus, error) {
	parsed, err := MessageBusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusBin), backend, _sigsVerifier, _liquidityBridge, _pegBridge, _pegVault)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBus{MessageBusCaller: MessageBusCaller{contract: contract}, MessageBusTransactor: MessageBusTransactor{contract: contract}, MessageBusFilterer: MessageBusFilterer{contract: contract}}, nil
}

// MessageBus is an auto generated Go binding around an Ethereum contract.
type MessageBus struct {
	MessageBusCaller     // Read-only binding to the contract
	MessageBusTransactor // Write-only binding to the contract
	MessageBusFilterer   // Log filterer for contract events
}

// MessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusSession struct {
	Contract     *MessageBus       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusCallerSession struct {
	Contract *MessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusTransactorSession struct {
	Contract     *MessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusRaw struct {
	Contract *MessageBus // Generic contract binding to access the raw methods on
}

// MessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusCallerRaw struct {
	Contract *MessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusTransactorRaw struct {
	Contract *MessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBus creates a new instance of MessageBus, bound to a specific deployed contract.
func NewMessageBus(address common.Address, backend bind.ContractBackend) (*MessageBus, error) {
	contract, err := bindMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBus{MessageBusCaller: MessageBusCaller{contract: contract}, MessageBusTransactor: MessageBusTransactor{contract: contract}, MessageBusFilterer: MessageBusFilterer{contract: contract}}, nil
}

// NewMessageBusCaller creates a new read-only instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusCaller(address common.Address, caller bind.ContractCaller) (*MessageBusCaller, error) {
	contract, err := bindMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusCaller{contract: contract}, nil
}

// NewMessageBusTransactor creates a new write-only instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusTransactor, error) {
	contract, err := bindMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusTransactor{contract: contract}, nil
}

// NewMessageBusFilterer creates a new log filterer instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusFilterer, error) {
	contract, err := bindMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusFilterer{contract: contract}, nil
}

// bindMessageBus binds a generic wrapper to an already deployed contract.
func bindMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBus *MessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBus.Contract.MessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBus *MessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBus.Contract.MessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBus *MessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBus.Contract.MessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBus *MessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBus *MessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBus *MessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBus.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBus.Contract.CalcFee(&_MessageBus.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBus.Contract.CalcFee(&_MessageBus.CallOpts, _message)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusCaller) ExecutedMessages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBus.Contract.ExecutedMessages(&_MessageBus.CallOpts, arg0)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusCallerSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBus.Contract.ExecutedMessages(&_MessageBus.CallOpts, arg0)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusSession) FeeBase() (*big.Int, error) {
	return _MessageBus.Contract.FeeBase(&_MessageBus.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) FeeBase() (*big.Int, error) {
	return _MessageBus.Contract.FeeBase(&_MessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusSession) FeePerByte() (*big.Int, error) {
	return _MessageBus.Contract.FeePerByte(&_MessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) FeePerByte() (*big.Int, error) {
	return _MessageBus.Contract.FeePerByte(&_MessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusSession) LiquidityBridge() (common.Address, error) {
	return _MessageBus.Contract.LiquidityBridge(&_MessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusCallerSession) LiquidityBridge() (common.Address, error) {
	return _MessageBus.Contract.LiquidityBridge(&_MessageBus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusSession) Owner() (common.Address, error) {
	return _MessageBus.Contract.Owner(&_MessageBus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusCallerSession) Owner() (common.Address, error) {
	return _MessageBus.Contract.Owner(&_MessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusSession) PegBridge() (common.Address, error) {
	return _MessageBus.Contract.PegBridge(&_MessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegBridge() (common.Address, error) {
	return _MessageBus.Contract.PegBridge(&_MessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusSession) PegVault() (common.Address, error) {
	return _MessageBus.Contract.PegVault(&_MessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegVault() (common.Address, error) {
	return _MessageBus.Contract.PegVault(&_MessageBus.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusSession) SigsVerifier() (common.Address, error) {
	return _MessageBus.Contract.SigsVerifier(&_MessageBus.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusCallerSession) SigsVerifier() (common.Address, error) {
	return _MessageBus.Contract.SigsVerifier(&_MessageBus.CallOpts)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusCaller) WithdrawnFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "withdrawnFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBus.Contract.WithdrawnFees(&_MessageBus.CallOpts, arg0)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusCallerSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBus.Contract.WithdrawnFees(&_MessageBus.CallOpts, arg0)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MessageBusReceiverRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessage(_message []byte, _route MessageBusReceiverRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessage(_message []byte, _route MessageBusReceiverRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessageWithTransfer(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransfer(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransfer(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransferRefund(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransferRefund(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault) returns()
func (_MessageBus *MessageBusTransactor) Init(opts *bind.TransactOpts, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "init", _liquidityBridge, _pegBridge, _pegVault)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault) returns()
func (_MessageBus *MessageBusSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault) returns()
func (_MessageBus *MessageBusTransactorSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactorSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessageWithTransfer(&_MessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessageWithTransfer(&_MessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactor) SetFeeBase(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setFeeBase", _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeeBase(&_MessageBus.TransactOpts, _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactorSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeeBase(&_MessageBus.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactor) SetFeePerByte(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setFeePerByte", _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeePerByte(&_MessageBus.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactorSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeePerByte(&_MessageBus.TransactOpts, _fee)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetLiquidityBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setLiquidityBridge", _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetLiquidityBridge(&_MessageBus.TransactOpts, _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetLiquidityBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegBridge", _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegVault", _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVault(&_MessageBus.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVault(&_MessageBus.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferOwnership(&_MessageBus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferOwnership(&_MessageBus.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.WithdrawFee(&_MessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.WithdrawFee(&_MessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// MessageBusExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBus contract.
type MessageBusExecutedIterator struct {
	Event *MessageBusExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusExecuted)
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
		it.Event = new(MessageBusExecuted)
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
func (it *MessageBusExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusExecuted represents a Executed event raised by the MessageBus contract.
type MessageBusExecuted struct {
	MsgType uint8
	Id      [32]byte
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e7.
//
// Solidity: event Executed(uint8 msgType, bytes32 id, uint8 status)
func (_MessageBus *MessageBusFilterer) FilterExecuted(opts *bind.FilterOpts) (*MessageBusExecutedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &MessageBusExecutedIterator{contract: _MessageBus.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e7.
//
// Solidity: event Executed(uint8 msgType, bytes32 id, uint8 status)
func (_MessageBus *MessageBusFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusExecuted) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusExecuted)
				if err := _MessageBus.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e7.
//
// Solidity: event Executed(uint8 msgType, bytes32 id, uint8 status)
func (_MessageBus *MessageBusFilterer) ParseExecuted(log types.Log) (*MessageBusExecuted, error) {
	event := new(MessageBusExecuted)
	if err := _MessageBus.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the MessageBus contract.
type MessageBusMessageIterator struct {
	Event *MessageBusMessage // Event containing the contract specifics and raw log

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
func (it *MessageBusMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusMessage)
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
		it.Event = new(MessageBusMessage)
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
func (it *MessageBusMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusMessage represents a Message event raised by the MessageBus contract.
type MessageBusMessage struct {
	Sender     common.Address
	Receiver   common.Address
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*MessageBusMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusMessageIterator{contract: _MessageBus.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *MessageBusMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusMessage)
				if err := _MessageBus.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) ParseMessage(log types.Log) (*MessageBusMessage, error) {
	event := new(MessageBusMessage)
	if err := _MessageBus.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusMessageWithTransferIterator is returned from FilterMessageWithTransfer and is used to iterate over the raw logs and unpacked data for MessageWithTransfer events raised by the MessageBus contract.
type MessageBusMessageWithTransferIterator struct {
	Event *MessageBusMessageWithTransfer // Event containing the contract specifics and raw log

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
func (it *MessageBusMessageWithTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusMessageWithTransfer)
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
		it.Event = new(MessageBusMessageWithTransfer)
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
func (it *MessageBusMessageWithTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusMessageWithTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusMessageWithTransfer represents a MessageWithTransfer event raised by the MessageBus contract.
type MessageBusMessageWithTransfer struct {
	Sender        common.Address
	Receiver      common.Address
	DstChainId    *big.Int
	Bridge        common.Address
	SrcTransferId [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageWithTransfer is a free log retrieval operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) FilterMessageWithTransfer(opts *bind.FilterOpts, sender []common.Address) (*MessageBusMessageWithTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusMessageWithTransferIterator{contract: _MessageBus.contract, event: "MessageWithTransfer", logs: logs, sub: sub}, nil
}

// WatchMessageWithTransfer is a free log subscription operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) WatchMessageWithTransfer(opts *bind.WatchOpts, sink chan<- *MessageBusMessageWithTransfer, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusMessageWithTransfer)
				if err := _MessageBus.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
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

// ParseMessageWithTransfer is a log parse operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) ParseMessageWithTransfer(log types.Log) (*MessageBusMessageWithTransfer, error) {
	event := new(MessageBusMessageWithTransfer)
	if err := _MessageBus.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBus contract.
type MessageBusOwnershipTransferredIterator struct {
	Event *MessageBusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusOwnershipTransferred)
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
		it.Event = new(MessageBusOwnershipTransferred)
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
func (it *MessageBusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBus contract.
type MessageBusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBus *MessageBusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusOwnershipTransferredIterator{contract: _MessageBus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBus *MessageBusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusOwnershipTransferred)
				if err := _MessageBus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBus *MessageBusFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusOwnershipTransferred, error) {
	event := new(MessageBusOwnershipTransferred)
	if err := _MessageBus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusAddressMetaData contains all meta data concerning the MessageBusAddress contract.
var MessageBusAddressMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MessageBusAddressABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusAddressMetaData.ABI instead.
var MessageBusAddressABI = MessageBusAddressMetaData.ABI

// MessageBusAddress is an auto generated Go binding around an Ethereum contract.
type MessageBusAddress struct {
	MessageBusAddressCaller     // Read-only binding to the contract
	MessageBusAddressTransactor // Write-only binding to the contract
	MessageBusAddressFilterer   // Log filterer for contract events
}

// MessageBusAddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusAddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusAddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusAddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusAddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusAddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusAddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusAddressSession struct {
	Contract     *MessageBusAddress // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MessageBusAddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusAddressCallerSession struct {
	Contract *MessageBusAddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MessageBusAddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusAddressTransactorSession struct {
	Contract     *MessageBusAddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MessageBusAddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusAddressRaw struct {
	Contract *MessageBusAddress // Generic contract binding to access the raw methods on
}

// MessageBusAddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusAddressCallerRaw struct {
	Contract *MessageBusAddressCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusAddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusAddressTransactorRaw struct {
	Contract *MessageBusAddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusAddress creates a new instance of MessageBusAddress, bound to a specific deployed contract.
func NewMessageBusAddress(address common.Address, backend bind.ContractBackend) (*MessageBusAddress, error) {
	contract, err := bindMessageBusAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusAddress{MessageBusAddressCaller: MessageBusAddressCaller{contract: contract}, MessageBusAddressTransactor: MessageBusAddressTransactor{contract: contract}, MessageBusAddressFilterer: MessageBusAddressFilterer{contract: contract}}, nil
}

// NewMessageBusAddressCaller creates a new read-only instance of MessageBusAddress, bound to a specific deployed contract.
func NewMessageBusAddressCaller(address common.Address, caller bind.ContractCaller) (*MessageBusAddressCaller, error) {
	contract, err := bindMessageBusAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusAddressCaller{contract: contract}, nil
}

// NewMessageBusAddressTransactor creates a new write-only instance of MessageBusAddress, bound to a specific deployed contract.
func NewMessageBusAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusAddressTransactor, error) {
	contract, err := bindMessageBusAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusAddressTransactor{contract: contract}, nil
}

// NewMessageBusAddressFilterer creates a new log filterer instance of MessageBusAddress, bound to a specific deployed contract.
func NewMessageBusAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusAddressFilterer, error) {
	contract, err := bindMessageBusAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusAddressFilterer{contract: contract}, nil
}

// bindMessageBusAddress binds a generic wrapper to an already deployed contract.
func bindMessageBusAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusAddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusAddress *MessageBusAddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusAddress.Contract.MessageBusAddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusAddress *MessageBusAddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.MessageBusAddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusAddress *MessageBusAddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.MessageBusAddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusAddress *MessageBusAddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusAddress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusAddress *MessageBusAddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusAddress *MessageBusAddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.contract.Transact(opts, method, params...)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageBusAddress *MessageBusAddressCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusAddress.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageBusAddress *MessageBusAddressSession) MessageBus() (common.Address, error) {
	return _MessageBusAddress.Contract.MessageBus(&_MessageBusAddress.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageBusAddress *MessageBusAddressCallerSession) MessageBus() (common.Address, error) {
	return _MessageBusAddress.Contract.MessageBus(&_MessageBusAddress.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusAddress *MessageBusAddressCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusAddress.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusAddress *MessageBusAddressSession) Owner() (common.Address, error) {
	return _MessageBusAddress.Contract.Owner(&_MessageBusAddress.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusAddress *MessageBusAddressCallerSession) Owner() (common.Address, error) {
	return _MessageBusAddress.Contract.Owner(&_MessageBusAddress.CallOpts)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageBusAddress *MessageBusAddressTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _MessageBusAddress.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageBusAddress *MessageBusAddressSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.SetMessageBus(&_MessageBusAddress.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageBusAddress *MessageBusAddressTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.SetMessageBus(&_MessageBusAddress.TransactOpts, _messageBus)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusAddress *MessageBusAddressTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusAddress.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusAddress *MessageBusAddressSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.TransferOwnership(&_MessageBusAddress.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusAddress *MessageBusAddressTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusAddress.Contract.TransferOwnership(&_MessageBusAddress.TransactOpts, newOwner)
}

// MessageBusAddressOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusAddress contract.
type MessageBusAddressOwnershipTransferredIterator struct {
	Event *MessageBusAddressOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusAddressOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusAddressOwnershipTransferred)
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
		it.Event = new(MessageBusAddressOwnershipTransferred)
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
func (it *MessageBusAddressOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusAddressOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusAddressOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusAddress contract.
type MessageBusAddressOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusAddress *MessageBusAddressFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusAddressOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusAddress.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusAddressOwnershipTransferredIterator{contract: _MessageBusAddress.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusAddress *MessageBusAddressFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusAddressOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusAddress.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusAddressOwnershipTransferred)
				if err := _MessageBusAddress.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusAddress *MessageBusAddressFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusAddressOwnershipTransferred, error) {
	event := new(MessageBusAddressOwnershipTransferred)
	if err := _MessageBusAddress.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverMetaData contains all meta data concerning the MessageBusReceiver contract.
var MessageBusReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"}],\"internalType\":\"structMessageBusReceiver.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ca738038062001ca78339810160408190526200003491620000ef565b6200003f3362000082565b600280546001600160a01b039485166001600160a01b03199182161790915560038054938516938216939093179092556004805491909316911617905562000139565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000ea57600080fd5b919050565b6000806000606084860312156200010557600080fd5b6200011084620000d2565b92506200012060208501620000d2565b91506200013060408501620000d2565b90509250925092565b611b5e80620001496000396000f3fe6080604052600436106100c75760003560e01c80639b05a77511610074578063d8257d171161004e578063d8257d17146101ff578063dfa2dbaf1461021f578063f2fde38b1461023f57600080fd5b80639b05a7751461018f578063a2232213146101af578063cd2abd66146101c257600080fd5b8063654317bf116100a5578063654317bf1461012157806382980dc4146101345780638da5cb5b1461017157600080fd5b806303cbfe66146100cc578063588be02b146100ee578063588df4161461010e575b600080fd5b3480156100d857600080fd5b506100ec6100e736600461144f565b61025f565b005b3480156100fa57600080fd5b506100ec61010936600461144f565b6102ef565b6100ec61011c3660046114f8565b61037a565b6100ec61012f3660046115e9565b6105aa565b34801561014057600080fd5b50600254610154906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561017d57600080fd5b506000546001600160a01b0316610154565b34801561019b57600080fd5b506100ec6101aa36600461144f565b6107a8565b6100ec6101bd3660046114f8565b610833565b3480156101ce57600080fd5b506101f26101dd3660046116b0565b60016020526000908152604090205460ff1681565b60405161016891906116f3565b34801561020b57600080fd5b50600454610154906001600160a01b031681565b34801561022b57600080fd5b50600354610154906001600160a01b031681565b34801561024b57600080fd5b506100ec61025a36600461144f565b6109fc565b336102726000546001600160a01b031690565b6001600160a01b0316146102cd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b336103026000546001600160a01b031690565b6001600160a01b0316146103585760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c4565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600061038588610aed565b90506000808281526001602052604090205460ff1660038111156103ab576103ab6116c9565b146103f85760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c72656164792065786563757465640000000000000060448201526064016102c4565b6000463060405160200161044e92919091825260601b6bffffffffffffffffffffffff191660208201527f4d657373616765576974685472616e73666572526566756e64000000000000006034820152604d0190565b604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e6040516020016104a19493929190611707565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b81526004016104d89796959493929190611817565b60006040518083038186803b1580156104f057600080fd5b505afa158015610504573d6000803e3d6000fd5b505050506000806105168b8e8e611145565b90508015610527576001915061052c565b600291505b60008481526001602081905260409091208054849260ff199091169083600381111561055a5761055a6116c9565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e76000858460405161059393929190611911565b60405180910390a150505050505050505050505050565b60006105b7888b8b61127c565b90506000808281526001602052604090205460ff1660038111156105dd576105dd6116c9565b1461062a5760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c7265616479206578656375746564000000000000000060448201526064016102c4565b6000463060405160200161068092919091825260601b6bffffffffffffffffffffffff191660208201527f4d657373616765000000000000000000000000000000000000000000000000006034820152603b0190565b60408051601f1981840301815282825280516020918201206002549184018190528383018690528251808503840181526060850193849052633416de1160e11b90935293506001600160a01b03169163682dbc22916106ed918c908c908c908c908c908c90606401611817565b60006040518083038186803b15801561070557600080fd5b505afa158015610719573d6000803e3d6000fd5b5050505060008061072b8b8e8e6112e3565b9050801561073c5760019150610741565b600291505b60008481526001602081905260409091208054849260ff199091169083600381111561076f5761076f6116c9565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e76001858460405161059393929190611911565b336107bb6000546001600160a01b031690565b6001600160a01b0316146108115760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c4565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b600061083e88610aed565b90506000808281526001602052604090205460ff166003811115610864576108646116c9565b146108b15760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c72656164792065786563757465640000000000000060448201526064016102c4565b6000463060405160200161090792919091825260601b6bffffffffffffffffffffffff191660208201527f4d657373616765576974685472616e7366657200000000000000000000000000603482015260470190565b604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e60405160200161095a9493929190611707565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b81526004016109919796959493929190611817565b60006040518083038186803b1580156109a957600080fd5b505afa1580156109bd573d6000803e3d6000fd5b505050506000806109cf8b8e8e61133b565b905080156109e0576001915061052c565b6109eb8b8e8e6113ad565b90508015610527576003915061052c565b33610a0f6000546001600160a01b031690565b6001600160a01b031614610a655760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102c4565b6001600160a01b038116610ae15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102c4565b610aea816113e3565b50565b600080806001610b006020860186611943565b6004811115610b1157610b116116c9565b1415610cb157610b27604085016020860161144f565b610b37606086016040870161144f565b610b47608087016060880161144f565b6080870135610b5c60e0890160c08a01611964565b6040516bffffffffffffffffffffffff19606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600254633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b9060240160206040518083038186803b158015610c2357600080fd5b505afa158015610c37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5b919061198e565b1515600114610cac5760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f742065786973740000000000000000000060448201526064016102c4565b611110565b6002610cc06020860186611943565b6004811115610cd157610cd16116c9565b1415610e435746610ce860c0860160a08701611964565b610cf8606087016040880161144f565b610d08608088016060890161144f565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526bffffffffffffffffffffffff19606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600254631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab4289060240160206040518083038186803b158015610dba57600080fd5b505afa158015610dce573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df2919061198e565b1515600114610cac5760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f742065786973740000000000000060448201526064016102c4565b6003610e526020860186611943565b6004811115610e6357610e636116c9565b1480610e8c57506004610e796020860186611943565b6004811115610e8a57610e8a6116c9565b145b1561111057610ea1606085016040860161144f565b610eb1608086016060870161144f565b6080860135610ec6604088016020890161144f565b610ed660e0890160c08a01611964565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f19818403018152919052805160209091012091506003610f5b6020860186611943565b6004811115610f6c57610f6c6116c9565b141561104257506003546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e647259060240160206040518083038186803b158015610fb957600080fd5b505afa158015610fcd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff1919061198e565b1515600114610cac5760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f74206578697374000000000000000000000060448201526064016102c4565b50600480546040516301e6472560e01b81529182018390526001600160a01b03169081906301e647259060240160206040518083038186803b15801561108757600080fd5b505afa15801561109b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110bf919061198e565b15156001146111105760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f742065786973740000000000000060448201526064016102c4565b60008183604051602001611126939291906119c7565b6040516020818303038152906040528051906020012092505050919050565b60008080611159606087016040880161144f565b6001600160a01b03163463105f4af960e11b61117b60808a0160608b0161144f565b8960800135898960405160240161119594939291906119f8565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516112009190611a2b565b60006040518083038185875af1925050503d806000811461123d576040519150601f19603f3d011682016040523d82523d6000602084013e611242565b606091505b5091509150811561126e57600081806020019051810190611263919061198e565b935061127592505050565b6000925050505b9392505050565b6000600161128d602086018661144f565b61129d604087016020880161144f565b6112ad6060880160408901611964565b86866040516020016112c496959493929190611a47565b6040516020818303038152906040528051906020012090509392505050565b600080806112f7604087016020880161144f565b6001600160a01b031634631599d26560e01b61131660208a018a61144f565b61132660608b0160408c01611964565b89896040516024016111959493929190611aa6565b6000808061134f606087016040880161144f565b6001600160a01b03163463671aeecd60e11b61137160408a0160208b0161144f565b61138160808b0160608c0161144f565b60808b013561139660e08d0160c08e01611964565b8b8b60405160240161119596959493929190611ad9565b600080806113c1606087016040880161144f565b6001600160a01b0316346378079ce760e11b61137160408a0160208b0161144f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b038116811461144a57600080fd5b919050565b60006020828403121561146157600080fd5b61127582611433565b60008083601f84011261147c57600080fd5b50813567ffffffffffffffff81111561149457600080fd5b6020830191508360208285010111156114ac57600080fd5b9250929050565b60008083601f8401126114c557600080fd5b50813567ffffffffffffffff8111156114dd57600080fd5b6020830191508360208260051b85010111156114ac57600080fd5b6000806000806000806000806000898b0361018081121561151857600080fd5b8a3567ffffffffffffffff8082111561153057600080fd5b61153c8e838f0161146a565b909c509a508a9150610100601f198401121561155757600080fd5b60208d0199506101208d013592508083111561157257600080fd5b61157e8e848f016114b3565b90995097506101408d013592508891508083111561159b57600080fd5b6115a78e848f016114b3565b90975095506101608d01359250869150808311156115c457600080fd5b50506115d28c828d016114b3565b915080935050809150509295985092959850929598565b6000806000806000806000806000898b0360e081121561160857600080fd5b8a3567ffffffffffffffff8082111561162057600080fd5b61162c8e838f0161146a565b909c509a508a91506060601f198401121561164657600080fd5b60208d01995060808d013592508083111561166057600080fd5b61166c8e848f016114b3565b909950975060a08d013592508891508083111561168857600080fd5b6116948e848f016114b3565b909750955060c08d01359250869150808311156115c457600080fd5b6000602082840312156116c257600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b600481106116ef576116ef6116c9565b9052565b6020810161170182846116df565b92915050565b84815283602082015281836040830137600091016040019081529392505050565b60005b8381101561174357818101518382015260200161172b565b83811115611752576000848401525b50505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8183526000602080850194508260005b858110156117bd576001600160a01b036117aa83611433565b1687529582019590820190600101611791565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156117fa57600080fd5b8260051b8083602087013760009401602001938452509192915050565b60808152600088518060808401526118368160a0850160208d01611728565b601f01601f1916820182810360a090810160208501528101889052600588901b810160c09081019082018a60005b8b8110156118d75784840360bf190183528135368e9003601e1901811261188a57600080fd5b8d01803567ffffffffffffffff8111156118a357600080fd5b8036038f13156118b257600080fd5b6118c0868260208501611758565b955050506020928301929190910190600101611864565b50505083810360408501526118ed81888a611781565b91505082810360608401526119038185876117c8565b9a9950505050505050505050565b6060810160028510611925576119256116c9565b84825283602083015261193b60408301846116df565b949350505050565b60006020828403121561195557600080fd5b81356005811061127557600080fd5b60006020828403121561197657600080fd5b813567ffffffffffffffff8116811461127557600080fd5b6000602082840312156119a057600080fd5b8151801515811461127557600080fd5b600281106119c0576119c06116c9565b60f81b9052565b6119d181856119b0565b60609290921b6bffffffffffffffffffffffff191660018301526015820152603501919050565b6001600160a01b0385168152836020820152606060408201526000611a21606083018486611758565b9695505050505050565b60008251611a3d818460208701611728565b9190910192915050565b611a5181886119b0565b60006bffffffffffffffffffffffff19808860601b166001840152808760601b166015840152506001600160c01b03198560c01b16602983015282846031840137506000910160310190815295945050505050565b6001600160a01b038516815267ffffffffffffffff84166020820152606060408201526000611a21606083018486611758565b60006001600160a01b03808916835280881660208401525085604083015267ffffffffffffffff8516606083015260a06080830152611b1c60a083018486611758565b9897505050505050505056fea26469706673582212207ac7a08540a769a9b2a51adbb5bb59aa0a398dc56368a014783b9d9f092c99bc64736f6c63430008090033",
}

// MessageBusReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusReceiverMetaData.ABI instead.
var MessageBusReceiverABI = MessageBusReceiverMetaData.ABI

// MessageBusReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusReceiverMetaData.Bin instead.
var MessageBusReceiverBin = MessageBusReceiverMetaData.Bin

// DeployMessageBusReceiver deploys a new Ethereum contract, binding an instance of MessageBusReceiver to it.
func DeployMessageBusReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address) (common.Address, *types.Transaction, *MessageBusReceiver, error) {
	parsed, err := MessageBusReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusReceiverBin), backend, _liquidityBridge, _pegBridge, _pegVault)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusReceiver{MessageBusReceiverCaller: MessageBusReceiverCaller{contract: contract}, MessageBusReceiverTransactor: MessageBusReceiverTransactor{contract: contract}, MessageBusReceiverFilterer: MessageBusReceiverFilterer{contract: contract}}, nil
}

// MessageBusReceiver is an auto generated Go binding around an Ethereum contract.
type MessageBusReceiver struct {
	MessageBusReceiverCaller     // Read-only binding to the contract
	MessageBusReceiverTransactor // Write-only binding to the contract
	MessageBusReceiverFilterer   // Log filterer for contract events
}

// MessageBusReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusReceiverSession struct {
	Contract     *MessageBusReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessageBusReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusReceiverCallerSession struct {
	Contract *MessageBusReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MessageBusReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusReceiverTransactorSession struct {
	Contract     *MessageBusReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MessageBusReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusReceiverRaw struct {
	Contract *MessageBusReceiver // Generic contract binding to access the raw methods on
}

// MessageBusReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusReceiverCallerRaw struct {
	Contract *MessageBusReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusReceiverTransactorRaw struct {
	Contract *MessageBusReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusReceiver creates a new instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiver(address common.Address, backend bind.ContractBackend) (*MessageBusReceiver, error) {
	contract, err := bindMessageBusReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiver{MessageBusReceiverCaller: MessageBusReceiverCaller{contract: contract}, MessageBusReceiverTransactor: MessageBusReceiverTransactor{contract: contract}, MessageBusReceiverFilterer: MessageBusReceiverFilterer{contract: contract}}, nil
}

// NewMessageBusReceiverCaller creates a new read-only instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverCaller(address common.Address, caller bind.ContractCaller) (*MessageBusReceiverCaller, error) {
	contract, err := bindMessageBusReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverCaller{contract: contract}, nil
}

// NewMessageBusReceiverTransactor creates a new write-only instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusReceiverTransactor, error) {
	contract, err := bindMessageBusReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverTransactor{contract: contract}, nil
}

// NewMessageBusReceiverFilterer creates a new log filterer instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusReceiverFilterer, error) {
	contract, err := bindMessageBusReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverFilterer{contract: contract}, nil
}

// bindMessageBusReceiver binds a generic wrapper to an already deployed contract.
func bindMessageBusReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiver *MessageBusReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiver.Contract.MessageBusReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiver *MessageBusReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.MessageBusReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiver *MessageBusReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.MessageBusReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiver *MessageBusReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiver *MessageBusReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiver *MessageBusReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.contract.Transact(opts, method, params...)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverCaller) ExecutedMessages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBusReceiver.Contract.ExecutedMessages(&_MessageBusReceiver.CallOpts, arg0)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBusReceiver.Contract.ExecutedMessages(&_MessageBusReceiver.CallOpts, arg0)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) LiquidityBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.LiquidityBridge(&_MessageBusReceiver.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) LiquidityBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.LiquidityBridge(&_MessageBusReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) Owner() (common.Address, error) {
	return _MessageBusReceiver.Contract.Owner(&_MessageBusReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) Owner() (common.Address, error) {
	return _MessageBusReceiver.Contract.Owner(&_MessageBusReceiver.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridge(&_MessageBusReceiver.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridge(&_MessageBusReceiver.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegVault() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVault(&_MessageBusReceiver.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegVault() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVault(&_MessageBusReceiver.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MessageBusReceiverRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessage(_message []byte, _route MessageBusReceiverRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessage(_message []byte, _route MessageBusReceiverRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessageWithTransfer(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransfer(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransfer(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransferRefund(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MessageBusReceiverTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransferRefund(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetLiquidityBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setLiquidityBridge", _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetLiquidityBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetLiquidityBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegBridge", _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegVault", _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVault(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVault(&_MessageBusReceiver.TransactOpts, _addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferOwnership(&_MessageBusReceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferOwnership(&_MessageBusReceiver.TransactOpts, newOwner)
}

// MessageBusReceiverExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBusReceiver contract.
type MessageBusReceiverExecutedIterator struct {
	Event *MessageBusReceiverExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverExecuted)
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
		it.Event = new(MessageBusReceiverExecuted)
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
func (it *MessageBusReceiverExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverExecuted represents a Executed event raised by the MessageBusReceiver contract.
type MessageBusReceiverExecuted struct {
	MsgType uint8
	Id      [32]byte
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e7.
//
// Solidity: event Executed(uint8 msgType, bytes32 id, uint8 status)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterExecuted(opts *bind.FilterOpts) (*MessageBusReceiverExecutedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverExecutedIterator{contract: _MessageBusReceiver.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e7.
//
// Solidity: event Executed(uint8 msgType, bytes32 id, uint8 status)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverExecuted) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverExecuted)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e7.
//
// Solidity: event Executed(uint8 msgType, bytes32 id, uint8 status)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseExecuted(log types.Log) (*MessageBusReceiverExecuted, error) {
	event := new(MessageBusReceiverExecuted)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusReceiver contract.
type MessageBusReceiverOwnershipTransferredIterator struct {
	Event *MessageBusReceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverOwnershipTransferred)
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
		it.Event = new(MessageBusReceiverOwnershipTransferred)
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
func (it *MessageBusReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusReceiver contract.
type MessageBusReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverOwnershipTransferredIterator{contract: _MessageBusReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverOwnershipTransferred)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusReceiverOwnershipTransferred, error) {
	event := new(MessageBusReceiverOwnershipTransferred)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMetaData contains all meta data concerning the MessageBusSender contract.
var MessageBusSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeeBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051610f01380380610f0183398101604081905261002f91610099565b61003833610049565b6001600160a01b03166080526100c9565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ab57600080fd5b81516001600160a01b03811681146100c257600080fd5b9392505050565b608051610e166100eb600039600081816101ee015261038e0152610e166000f3fe6080604052600436106100c75760003560e01c806395e911a811610074578063e2c1ed251161004e578063e2c1ed2514610210578063f2fde38b14610230578063f60bbe2a1461025057600080fd5b806395e911a8146101b35780639f3ce55a146101c9578063ccf2683b146101dc57600080fd5b80635335dca2116100a55780635335dca2146101215780635b3e5f50146101545780638da5cb5b1461018157600080fd5b806306c28bd6146100cc5780632ff4c411146100ee5780634289fbb31461010e575b600080fd5b3480156100d857600080fd5b506100ec6100e7366004610862565b610266565b005b3480156100fa57600080fd5b506100ec6101093660046108e3565b6102d9565b6100ec61011c3660046109d9565b61052f565b34801561012d57600080fd5b5061014161013c366004610a51565b6105d6565b6040519081526020015b60405180910390f35b34801561016057600080fd5b5061014161016f366004610a93565b60036020526000908152604090205481565b34801561018d57600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161014b565b3480156101bf57600080fd5b5061014160015481565b6100ec6101d7366004610aae565b6105fa565b3480156101e857600080fd5b5061019b7f000000000000000000000000000000000000000000000000000000000000000081565b34801561021c57600080fd5b506100ec61022b366004610862565b61069b565b34801561023c57600080fd5b506100ec61024b366004610a93565b610709565b34801561025c57600080fd5b5061014160025481565b336102796000546001600160a01b031690565b6001600160a01b0316146102d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600155565b6000463060405160200161032f92919091825260601b6bffffffffffffffffffffffff191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6bffffffffffffffffffffffff19168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc22916103d5918b908b908b908b908b908b90607801610c58565b60006040518083038186803b1580156103ed57600080fd5b505afa158015610401573d6000803e3d6000fd5b505050506001600160a01b038916600090815260036020526040812054610428908a610d0a565b90506000811161047a5760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f2077697468647261770000000000000060448201526064016102cb565b60008a6001600160a01b03168261c35090604051600060405180830381858888f193505050503d80600081146104cc576040519150601f19603f3d011682016040523d82523d6000602084013e6104d1565b606091505b50509050806105225760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f207769746864726177206665650000000000000000000060448201526064016102cb565b5050505050505050505050565b600061053b83836105d6565b9050803410156105805760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102cb565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66888888888888346040516105c59796959493929190610d21565b60405180910390a250505050505050565b6002546000906105e69083610d6e565b6001546105f39190610d8d565b9392505050565b600061060683836105d6565b90508034101561064b5760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102cb565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4868686863460405161068c959493929190610da5565b60405180910390a25050505050565b336106ae6000546001600160a01b031690565b6001600160a01b0316146107045760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102cb565b600255565b3361071c6000546001600160a01b031690565b6001600160a01b0316146107725760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102cb565b6001600160a01b0381166107ee5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102cb565b6107f7816107fa565b50565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561087457600080fd5b5035919050565b80356001600160a01b038116811461089257600080fd5b919050565b60008083601f8401126108a957600080fd5b50813567ffffffffffffffff8111156108c157600080fd5b6020830191508360208260051b85010111156108dc57600080fd5b9250929050565b60008060008060008060008060a0898b0312156108ff57600080fd5b6109088961087b565b975060208901359650604089013567ffffffffffffffff8082111561092c57600080fd5b6109388c838d01610897565b909850965060608b013591508082111561095157600080fd5b61095d8c838d01610897565b909650945060808b013591508082111561097657600080fd5b506109838b828c01610897565b999c989b5096995094979396929594505050565b60008083601f8401126109a957600080fd5b50813567ffffffffffffffff8111156109c157600080fd5b6020830191508360208285010111156108dc57600080fd5b60008060008060008060a087890312156109f257600080fd5b6109fb8761087b565b955060208701359450610a106040880161087b565b935060608701359250608087013567ffffffffffffffff811115610a3357600080fd5b610a3f89828a01610997565b979a9699509497509295939492505050565b60008060208385031215610a6457600080fd5b823567ffffffffffffffff811115610a7b57600080fd5b610a8785828601610997565b90969095509350505050565b600060208284031215610aa557600080fd5b6105f38261087b565b60008060008060608587031215610ac457600080fd5b610acd8561087b565b935060208501359250604085013567ffffffffffffffff811115610af057600080fd5b610afc87828801610997565b95989497509550505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b87811015610bb55782840389528135601e19883603018112610b6c57600080fd5b8701803567ffffffffffffffff811115610b8557600080fd5b803603891315610b9457600080fd5b610ba18682898501610b08565b9a87019a9550505090840190600101610b4b565b5091979650505050505050565b8183526000602080850194508260005b85811015610bfe576001600160a01b03610beb8361087b565b1687529582019590820190600101610bd2565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610c3b57600080fd5b8260051b8083602087013760009401602001938452509192915050565b608081526000885180608084015260005b81811015610c86576020818c0181015160a0868401015201610c69565b81811115610c9857600060a083860101525b50601f01601f1916820182810360a09081016020850152610cbc908201898b610b31565b90508281036040840152610cd1818789610bc2565b90508281036060840152610ce6818587610c09565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610d1c57610d1c610cf4565b500390565b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c06080830152610d5a60c083018587610b08565b90508260a083015298975050505050505050565b6000816000190483118215151615610d8857610d88610cf4565b500290565b60008219821115610da057610da0610cf4565b500190565b6001600160a01b0386168152846020820152608060408201526000610dce608083018587610b08565b9050826060830152969550505050505056fea26469706673582212205ee2f8d80eb931bf0778bb098971416c53ed900ada3b875053c76243e948f8fe64736f6c63430008090033",
}

// MessageBusSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusSenderMetaData.ABI instead.
var MessageBusSenderABI = MessageBusSenderMetaData.ABI

// MessageBusSenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusSenderMetaData.Bin instead.
var MessageBusSenderBin = MessageBusSenderMetaData.Bin

// DeployMessageBusSender deploys a new Ethereum contract, binding an instance of MessageBusSender to it.
func DeployMessageBusSender(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address) (common.Address, *types.Transaction, *MessageBusSender, error) {
	parsed, err := MessageBusSenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusSenderBin), backend, _sigsVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusSender{MessageBusSenderCaller: MessageBusSenderCaller{contract: contract}, MessageBusSenderTransactor: MessageBusSenderTransactor{contract: contract}, MessageBusSenderFilterer: MessageBusSenderFilterer{contract: contract}}, nil
}

// MessageBusSender is an auto generated Go binding around an Ethereum contract.
type MessageBusSender struct {
	MessageBusSenderCaller     // Read-only binding to the contract
	MessageBusSenderTransactor // Write-only binding to the contract
	MessageBusSenderFilterer   // Log filterer for contract events
}

// MessageBusSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusSenderSession struct {
	Contract     *MessageBusSender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageBusSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusSenderCallerSession struct {
	Contract *MessageBusSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MessageBusSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusSenderTransactorSession struct {
	Contract     *MessageBusSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MessageBusSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusSenderRaw struct {
	Contract *MessageBusSender // Generic contract binding to access the raw methods on
}

// MessageBusSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusSenderCallerRaw struct {
	Contract *MessageBusSenderCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusSenderTransactorRaw struct {
	Contract *MessageBusSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusSender creates a new instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSender(address common.Address, backend bind.ContractBackend) (*MessageBusSender, error) {
	contract, err := bindMessageBusSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusSender{MessageBusSenderCaller: MessageBusSenderCaller{contract: contract}, MessageBusSenderTransactor: MessageBusSenderTransactor{contract: contract}, MessageBusSenderFilterer: MessageBusSenderFilterer{contract: contract}}, nil
}

// NewMessageBusSenderCaller creates a new read-only instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderCaller(address common.Address, caller bind.ContractCaller) (*MessageBusSenderCaller, error) {
	contract, err := bindMessageBusSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderCaller{contract: contract}, nil
}

// NewMessageBusSenderTransactor creates a new write-only instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusSenderTransactor, error) {
	contract, err := bindMessageBusSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderTransactor{contract: contract}, nil
}

// NewMessageBusSenderFilterer creates a new log filterer instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusSenderFilterer, error) {
	contract, err := bindMessageBusSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFilterer{contract: contract}, nil
}

// bindMessageBusSender binds a generic wrapper to an already deployed contract.
func bindMessageBusSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSender *MessageBusSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSender.Contract.MessageBusSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSender *MessageBusSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSender.Contract.MessageBusSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSender *MessageBusSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSender.Contract.MessageBusSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSender *MessageBusSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSender *MessageBusSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSender *MessageBusSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSender.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBusSender.Contract.CalcFee(&_MessageBusSender.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBusSender.Contract.CalcFee(&_MessageBusSender.CallOpts, _message)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) FeeBase() (*big.Int, error) {
	return _MessageBusSender.Contract.FeeBase(&_MessageBusSender.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) FeeBase() (*big.Int, error) {
	return _MessageBusSender.Contract.FeeBase(&_MessageBusSender.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) FeePerByte() (*big.Int, error) {
	return _MessageBusSender.Contract.FeePerByte(&_MessageBusSender.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) FeePerByte() (*big.Int, error) {
	return _MessageBusSender.Contract.FeePerByte(&_MessageBusSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderSession) Owner() (common.Address, error) {
	return _MessageBusSender.Contract.Owner(&_MessageBusSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderCallerSession) Owner() (common.Address, error) {
	return _MessageBusSender.Contract.Owner(&_MessageBusSender.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderSession) SigsVerifier() (common.Address, error) {
	return _MessageBusSender.Contract.SigsVerifier(&_MessageBusSender.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderCallerSession) SigsVerifier() (common.Address, error) {
	return _MessageBusSender.Contract.SigsVerifier(&_MessageBusSender.CallOpts)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) WithdrawnFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "withdrawnFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBusSender.Contract.WithdrawnFees(&_MessageBusSender.CallOpts, arg0)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBusSender.Contract.WithdrawnFees(&_MessageBusSender.CallOpts, arg0)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactor) SendMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessageWithTransfer(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessageWithTransfer(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactor) SetFeeBase(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "setFeeBase", _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeeBase(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeeBase(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactor) SetFeePerByte(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "setFeePerByte", _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeePerByte(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeePerByte(&_MessageBusSender.TransactOpts, _fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.Contract.TransferOwnership(&_MessageBusSender.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.Contract.TransferOwnership(&_MessageBusSender.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.WithdrawFee(&_MessageBusSender.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.WithdrawFee(&_MessageBusSender.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// MessageBusSenderMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the MessageBusSender contract.
type MessageBusSenderMessageIterator struct {
	Event *MessageBusSenderMessage // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderMessage)
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
		it.Event = new(MessageBusSenderMessage)
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
func (it *MessageBusSenderMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderMessage represents a Message event raised by the MessageBusSender contract.
type MessageBusSenderMessage struct {
	Sender     common.Address
	Receiver   common.Address
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*MessageBusSenderMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderMessageIterator{contract: _MessageBusSender.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *MessageBusSenderMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderMessage)
				if err := _MessageBusSender.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) ParseMessage(log types.Log) (*MessageBusSenderMessage, error) {
	event := new(MessageBusSenderMessage)
	if err := _MessageBusSender.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMessageWithTransferIterator is returned from FilterMessageWithTransfer and is used to iterate over the raw logs and unpacked data for MessageWithTransfer events raised by the MessageBusSender contract.
type MessageBusSenderMessageWithTransferIterator struct {
	Event *MessageBusSenderMessageWithTransfer // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderMessageWithTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderMessageWithTransfer)
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
		it.Event = new(MessageBusSenderMessageWithTransfer)
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
func (it *MessageBusSenderMessageWithTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderMessageWithTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderMessageWithTransfer represents a MessageWithTransfer event raised by the MessageBusSender contract.
type MessageBusSenderMessageWithTransfer struct {
	Sender        common.Address
	Receiver      common.Address
	DstChainId    *big.Int
	Bridge        common.Address
	SrcTransferId [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageWithTransfer is a free log retrieval operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) FilterMessageWithTransfer(opts *bind.FilterOpts, sender []common.Address) (*MessageBusSenderMessageWithTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderMessageWithTransferIterator{contract: _MessageBusSender.contract, event: "MessageWithTransfer", logs: logs, sub: sub}, nil
}

// WatchMessageWithTransfer is a free log subscription operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) WatchMessageWithTransfer(opts *bind.WatchOpts, sink chan<- *MessageBusSenderMessageWithTransfer, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderMessageWithTransfer)
				if err := _MessageBusSender.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
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

// ParseMessageWithTransfer is a log parse operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) ParseMessageWithTransfer(log types.Log) (*MessageBusSenderMessageWithTransfer, error) {
	event := new(MessageBusSenderMessageWithTransfer)
	if err := _MessageBusSender.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusSender contract.
type MessageBusSenderOwnershipTransferredIterator struct {
	Event *MessageBusSenderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderOwnershipTransferred)
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
		it.Event = new(MessageBusSenderOwnershipTransferred)
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
func (it *MessageBusSenderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusSender contract.
type MessageBusSenderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSender *MessageBusSenderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusSenderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderOwnershipTransferredIterator{contract: _MessageBusSender.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSender *MessageBusSenderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusSenderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderOwnershipTransferred)
				if err := _MessageBusSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageBusSender *MessageBusSenderFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusSenderOwnershipTransferred, error) {
	event := new(MessageBusSenderOwnershipTransferred)
	if err := _MessageBusSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAppMetaData contains all meta data concerning the MessageReceiverApp contract.
var MessageReceiverAppMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MessageReceiverAppABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageReceiverAppMetaData.ABI instead.
var MessageReceiverAppABI = MessageReceiverAppMetaData.ABI

// MessageReceiverApp is an auto generated Go binding around an Ethereum contract.
type MessageReceiverApp struct {
	MessageReceiverAppCaller     // Read-only binding to the contract
	MessageReceiverAppTransactor // Write-only binding to the contract
	MessageReceiverAppFilterer   // Log filterer for contract events
}

// MessageReceiverAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageReceiverAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageReceiverAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageReceiverAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageReceiverAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageReceiverAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageReceiverAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageReceiverAppSession struct {
	Contract     *MessageReceiverApp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessageReceiverAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageReceiverAppCallerSession struct {
	Contract *MessageReceiverAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MessageReceiverAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageReceiverAppTransactorSession struct {
	Contract     *MessageReceiverAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MessageReceiverAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageReceiverAppRaw struct {
	Contract *MessageReceiverApp // Generic contract binding to access the raw methods on
}

// MessageReceiverAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageReceiverAppCallerRaw struct {
	Contract *MessageReceiverAppCaller // Generic read-only contract binding to access the raw methods on
}

// MessageReceiverAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageReceiverAppTransactorRaw struct {
	Contract *MessageReceiverAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageReceiverApp creates a new instance of MessageReceiverApp, bound to a specific deployed contract.
func NewMessageReceiverApp(address common.Address, backend bind.ContractBackend) (*MessageReceiverApp, error) {
	contract, err := bindMessageReceiverApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverApp{MessageReceiverAppCaller: MessageReceiverAppCaller{contract: contract}, MessageReceiverAppTransactor: MessageReceiverAppTransactor{contract: contract}, MessageReceiverAppFilterer: MessageReceiverAppFilterer{contract: contract}}, nil
}

// NewMessageReceiverAppCaller creates a new read-only instance of MessageReceiverApp, bound to a specific deployed contract.
func NewMessageReceiverAppCaller(address common.Address, caller bind.ContractCaller) (*MessageReceiverAppCaller, error) {
	contract, err := bindMessageReceiverApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAppCaller{contract: contract}, nil
}

// NewMessageReceiverAppTransactor creates a new write-only instance of MessageReceiverApp, bound to a specific deployed contract.
func NewMessageReceiverAppTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageReceiverAppTransactor, error) {
	contract, err := bindMessageReceiverApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAppTransactor{contract: contract}, nil
}

// NewMessageReceiverAppFilterer creates a new log filterer instance of MessageReceiverApp, bound to a specific deployed contract.
func NewMessageReceiverAppFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageReceiverAppFilterer, error) {
	contract, err := bindMessageReceiverApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAppFilterer{contract: contract}, nil
}

// bindMessageReceiverApp binds a generic wrapper to an already deployed contract.
func bindMessageReceiverApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageReceiverAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageReceiverApp *MessageReceiverAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageReceiverApp.Contract.MessageReceiverAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageReceiverApp *MessageReceiverAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.MessageReceiverAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageReceiverApp *MessageReceiverAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.MessageReceiverAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageReceiverApp *MessageReceiverAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageReceiverApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageReceiverApp *MessageReceiverAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageReceiverApp *MessageReceiverAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.contract.Transact(opts, method, params...)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageReceiverApp *MessageReceiverAppCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageReceiverApp.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageReceiverApp *MessageReceiverAppSession) MessageBus() (common.Address, error) {
	return _MessageReceiverApp.Contract.MessageBus(&_MessageReceiverApp.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageReceiverApp *MessageReceiverAppCallerSession) MessageBus() (common.Address, error) {
	return _MessageReceiverApp.Contract.MessageBus(&_MessageReceiverApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageReceiverApp *MessageReceiverAppCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageReceiverApp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageReceiverApp *MessageReceiverAppSession) Owner() (common.Address, error) {
	return _MessageReceiverApp.Contract.Owner(&_MessageReceiverApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageReceiverApp *MessageReceiverAppCallerSession) Owner() (common.Address, error) {
	return _MessageReceiverApp.Contract.Owner(&_MessageReceiverApp.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessage(&_MessageReceiverApp.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactorSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessage(&_MessageReceiverApp.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_MessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_MessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_MessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_MessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_MessageReceiverApp.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_MessageReceiverApp *MessageReceiverAppTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_MessageReceiverApp.TransactOpts, _token, _amount, _message)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageReceiverApp *MessageReceiverAppTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _MessageReceiverApp.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageReceiverApp *MessageReceiverAppSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.SetMessageBus(&_MessageReceiverApp.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageReceiverApp *MessageReceiverAppTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.SetMessageBus(&_MessageReceiverApp.TransactOpts, _messageBus)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageReceiverApp *MessageReceiverAppTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageReceiverApp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageReceiverApp *MessageReceiverAppSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.TransferOwnership(&_MessageReceiverApp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageReceiverApp *MessageReceiverAppTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageReceiverApp.Contract.TransferOwnership(&_MessageReceiverApp.TransactOpts, newOwner)
}

// MessageReceiverAppOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageReceiverApp contract.
type MessageReceiverAppOwnershipTransferredIterator struct {
	Event *MessageReceiverAppOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAppOwnershipTransferred)
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
		it.Event = new(MessageReceiverAppOwnershipTransferred)
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
func (it *MessageReceiverAppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAppOwnershipTransferred represents a OwnershipTransferred event raised by the MessageReceiverApp contract.
type MessageReceiverAppOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageReceiverApp *MessageReceiverAppFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageReceiverAppOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageReceiverApp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAppOwnershipTransferredIterator{contract: _MessageReceiverApp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageReceiverApp *MessageReceiverAppFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageReceiverAppOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageReceiverApp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAppOwnershipTransferred)
				if err := _MessageReceiverApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageReceiverApp *MessageReceiverAppFilterer) ParseOwnershipTransferred(log types.Log) (*MessageReceiverAppOwnershipTransferred, error) {
	event := new(MessageReceiverAppOwnershipTransferred)
	if err := _MessageReceiverApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageSenderAppMetaData contains all meta data concerning the MessageSenderApp contract.
var MessageSenderAppMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MessageSenderAppABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageSenderAppMetaData.ABI instead.
var MessageSenderAppABI = MessageSenderAppMetaData.ABI

// MessageSenderApp is an auto generated Go binding around an Ethereum contract.
type MessageSenderApp struct {
	MessageSenderAppCaller     // Read-only binding to the contract
	MessageSenderAppTransactor // Write-only binding to the contract
	MessageSenderAppFilterer   // Log filterer for contract events
}

// MessageSenderAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageSenderAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSenderAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageSenderAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSenderAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageSenderAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSenderAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageSenderAppSession struct {
	Contract     *MessageSenderApp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageSenderAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageSenderAppCallerSession struct {
	Contract *MessageSenderAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MessageSenderAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageSenderAppTransactorSession struct {
	Contract     *MessageSenderAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MessageSenderAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageSenderAppRaw struct {
	Contract *MessageSenderApp // Generic contract binding to access the raw methods on
}

// MessageSenderAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageSenderAppCallerRaw struct {
	Contract *MessageSenderAppCaller // Generic read-only contract binding to access the raw methods on
}

// MessageSenderAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageSenderAppTransactorRaw struct {
	Contract *MessageSenderAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageSenderApp creates a new instance of MessageSenderApp, bound to a specific deployed contract.
func NewMessageSenderApp(address common.Address, backend bind.ContractBackend) (*MessageSenderApp, error) {
	contract, err := bindMessageSenderApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageSenderApp{MessageSenderAppCaller: MessageSenderAppCaller{contract: contract}, MessageSenderAppTransactor: MessageSenderAppTransactor{contract: contract}, MessageSenderAppFilterer: MessageSenderAppFilterer{contract: contract}}, nil
}

// NewMessageSenderAppCaller creates a new read-only instance of MessageSenderApp, bound to a specific deployed contract.
func NewMessageSenderAppCaller(address common.Address, caller bind.ContractCaller) (*MessageSenderAppCaller, error) {
	contract, err := bindMessageSenderApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageSenderAppCaller{contract: contract}, nil
}

// NewMessageSenderAppTransactor creates a new write-only instance of MessageSenderApp, bound to a specific deployed contract.
func NewMessageSenderAppTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageSenderAppTransactor, error) {
	contract, err := bindMessageSenderApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageSenderAppTransactor{contract: contract}, nil
}

// NewMessageSenderAppFilterer creates a new log filterer instance of MessageSenderApp, bound to a specific deployed contract.
func NewMessageSenderAppFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageSenderAppFilterer, error) {
	contract, err := bindMessageSenderApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageSenderAppFilterer{contract: contract}, nil
}

// bindMessageSenderApp binds a generic wrapper to an already deployed contract.
func bindMessageSenderApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageSenderAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageSenderApp *MessageSenderAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageSenderApp.Contract.MessageSenderAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageSenderApp *MessageSenderAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.MessageSenderAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageSenderApp *MessageSenderAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.MessageSenderAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageSenderApp *MessageSenderAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageSenderApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageSenderApp *MessageSenderAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageSenderApp *MessageSenderAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.contract.Transact(opts, method, params...)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageSenderApp *MessageSenderAppCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageSenderApp.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageSenderApp *MessageSenderAppSession) MessageBus() (common.Address, error) {
	return _MessageSenderApp.Contract.MessageBus(&_MessageSenderApp.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageSenderApp *MessageSenderAppCallerSession) MessageBus() (common.Address, error) {
	return _MessageSenderApp.Contract.MessageBus(&_MessageSenderApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageSenderApp *MessageSenderAppCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageSenderApp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageSenderApp *MessageSenderAppSession) Owner() (common.Address, error) {
	return _MessageSenderApp.Contract.Owner(&_MessageSenderApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageSenderApp *MessageSenderAppCallerSession) Owner() (common.Address, error) {
	return _MessageSenderApp.Contract.Owner(&_MessageSenderApp.CallOpts)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageSenderApp *MessageSenderAppTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _MessageSenderApp.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageSenderApp *MessageSenderAppSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.SetMessageBus(&_MessageSenderApp.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageSenderApp *MessageSenderAppTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.SetMessageBus(&_MessageSenderApp.TransactOpts, _messageBus)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageSenderApp *MessageSenderAppTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageSenderApp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageSenderApp *MessageSenderAppSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.TransferOwnership(&_MessageSenderApp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageSenderApp *MessageSenderAppTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageSenderApp.Contract.TransferOwnership(&_MessageSenderApp.TransactOpts, newOwner)
}

// MessageSenderAppOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageSenderApp contract.
type MessageSenderAppOwnershipTransferredIterator struct {
	Event *MessageSenderAppOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageSenderAppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageSenderAppOwnershipTransferred)
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
		it.Event = new(MessageSenderAppOwnershipTransferred)
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
func (it *MessageSenderAppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageSenderAppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageSenderAppOwnershipTransferred represents a OwnershipTransferred event raised by the MessageSenderApp contract.
type MessageSenderAppOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageSenderApp *MessageSenderAppFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageSenderAppOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageSenderApp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageSenderAppOwnershipTransferredIterator{contract: _MessageSenderApp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageSenderApp *MessageSenderAppFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageSenderAppOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageSenderApp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageSenderAppOwnershipTransferred)
				if err := _MessageSenderApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageSenderApp *MessageSenderAppFilterer) ParseOwnershipTransferred(log types.Log) (*MessageSenderAppOwnershipTransferred, error) {
	event := new(MessageSenderAppOwnershipTransferred)
	if err := _MessageSenderApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageSenderLibMetaData contains all meta data concerning the MessageSenderLib contract.
var MessageSenderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201b4107c7975ddd46d35da7d0f73842aff687b071f5e9694ec6cdc7700242681164736f6c63430008090033",
}

// MessageSenderLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageSenderLibMetaData.ABI instead.
var MessageSenderLibABI = MessageSenderLibMetaData.ABI

// MessageSenderLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageSenderLibMetaData.Bin instead.
var MessageSenderLibBin = MessageSenderLibMetaData.Bin

// DeployMessageSenderLib deploys a new Ethereum contract, binding an instance of MessageSenderLib to it.
func DeployMessageSenderLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageSenderLib, error) {
	parsed, err := MessageSenderLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageSenderLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageSenderLib{MessageSenderLibCaller: MessageSenderLibCaller{contract: contract}, MessageSenderLibTransactor: MessageSenderLibTransactor{contract: contract}, MessageSenderLibFilterer: MessageSenderLibFilterer{contract: contract}}, nil
}

// MessageSenderLib is an auto generated Go binding around an Ethereum contract.
type MessageSenderLib struct {
	MessageSenderLibCaller     // Read-only binding to the contract
	MessageSenderLibTransactor // Write-only binding to the contract
	MessageSenderLibFilterer   // Log filterer for contract events
}

// MessageSenderLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageSenderLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSenderLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageSenderLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSenderLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageSenderLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSenderLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageSenderLibSession struct {
	Contract     *MessageSenderLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageSenderLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageSenderLibCallerSession struct {
	Contract *MessageSenderLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MessageSenderLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageSenderLibTransactorSession struct {
	Contract     *MessageSenderLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MessageSenderLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageSenderLibRaw struct {
	Contract *MessageSenderLib // Generic contract binding to access the raw methods on
}

// MessageSenderLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageSenderLibCallerRaw struct {
	Contract *MessageSenderLibCaller // Generic read-only contract binding to access the raw methods on
}

// MessageSenderLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageSenderLibTransactorRaw struct {
	Contract *MessageSenderLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageSenderLib creates a new instance of MessageSenderLib, bound to a specific deployed contract.
func NewMessageSenderLib(address common.Address, backend bind.ContractBackend) (*MessageSenderLib, error) {
	contract, err := bindMessageSenderLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageSenderLib{MessageSenderLibCaller: MessageSenderLibCaller{contract: contract}, MessageSenderLibTransactor: MessageSenderLibTransactor{contract: contract}, MessageSenderLibFilterer: MessageSenderLibFilterer{contract: contract}}, nil
}

// NewMessageSenderLibCaller creates a new read-only instance of MessageSenderLib, bound to a specific deployed contract.
func NewMessageSenderLibCaller(address common.Address, caller bind.ContractCaller) (*MessageSenderLibCaller, error) {
	contract, err := bindMessageSenderLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageSenderLibCaller{contract: contract}, nil
}

// NewMessageSenderLibTransactor creates a new write-only instance of MessageSenderLib, bound to a specific deployed contract.
func NewMessageSenderLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageSenderLibTransactor, error) {
	contract, err := bindMessageSenderLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageSenderLibTransactor{contract: contract}, nil
}

// NewMessageSenderLibFilterer creates a new log filterer instance of MessageSenderLib, bound to a specific deployed contract.
func NewMessageSenderLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageSenderLibFilterer, error) {
	contract, err := bindMessageSenderLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageSenderLibFilterer{contract: contract}, nil
}

// bindMessageSenderLib binds a generic wrapper to an already deployed contract.
func bindMessageSenderLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageSenderLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageSenderLib *MessageSenderLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageSenderLib.Contract.MessageSenderLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageSenderLib *MessageSenderLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageSenderLib.Contract.MessageSenderLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageSenderLib *MessageSenderLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageSenderLib.Contract.MessageSenderLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageSenderLib *MessageSenderLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageSenderLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageSenderLib *MessageSenderLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageSenderLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageSenderLib *MessageSenderLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageSenderLib.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207136978c197da711d07ae3e00fb4dd4cea3514cbe1ac10035076120c675c7d4164736f6c63430008090033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// TestRefundMetaData contains all meta data concerning the TestRefund contract.
var TestRefundMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"enumMessageSenderLib.BridgeType\",\"name\":\"_bridgeType\",\"type\":\"uint8\"}],\"name\":\"sendWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161151838038061151883398101604081905261002f916100ad565b6100383361005d565b600180546001600160a01b0319166001600160a01b03929092169190911790556100dd565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bf57600080fd5b81516001600160a01b03811681146100d657600080fd5b9392505050565b61142c806100ec6000396000f3fe6080604052600436106100965760003560e01c80638da5cb5b11610069578063ce35dd9a1161004e578063ce35dd9a1461015d578063f00f39ce1461015d578063f2fde38b1461017057600080fd5b80638da5cb5b1461010b578063a1a227fa1461013d57600080fd5b80631599d2651461009b57806320be95f2146100c3578063547cad12146100d65780637ee1bd4a146100f8575b600080fd5b6100ae6100a9366004611076565b610190565b60405190151581526020015b60405180910390f35b6100ae6100d13660046110d9565b6101fd565b3480156100e257600080fd5b506100f66100f136600461111d565b6102c7565b005b6100f661010636600461113a565b61035f565b34801561011757600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016100ba565b34801561014957600080fd5b50600154610125906001600160a01b031681565b6100ae61016b3660046111d0565b6103b9565b34801561017c57600080fd5b506100f661018b36600461111d565b610420565b6001546000906001600160a01b031633146101f25760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b506001949350505050565b6001546000906001600160a01b0316331461025a5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016101e9565b7fee93b3d90bf7307985573224b309958a34f05c4bc6e541d3c2a9f903651cb3b68585858560405161028f9493929190611253565b60405180910390a160006102a58385018561111d565b90506102bb6001600160a01b0387168287610511565b50600195945050505050565b336102da6000546001600160a01b031690565b6001600160a01b0316146103305760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101e9565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6103746001600160a01b0387163330886105a6565b604080516001600160a01b03891660208201526000910160405160208183030381529060405290506103ae888888888888878960006105e4565b505050505050505050565b6001546000906001600160a01b031633146104165760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016101e9565b9695505050505050565b336104336000546001600160a01b031690565b6001600160a01b0316146104895760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101e9565b6001600160a01b0381166105055760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016101e9565b61050e8161061b565b50565b6040516001600160a01b0383166024820152604481018290526105a190849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610678565b505050565b6040516001600160a01b03808516602483015283166044820152606481018290526105de9085906323b872dd60e01b9060840161053d565b50505050565b600061060d8a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b61075d565b9a9950505050505050505050565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006106cd826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661082d9092919063ffffffff16565b8051909150156105a157808060200190518101906106eb9190611299565b6105a15760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101e9565b60006001846003811115610773576107736112bb565b14156107915761078a8b8b8b8b8b8b8b8a8a610846565b905061060d565b60028460038111156107a5576107a56112bb565b14156107bb5761078a8b8b8b8b8b8a8989610a54565b60038460038111156107cf576107cf6112bb565b14156107e55761078a8b8b8b8b8b8a8989610c6e565b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f727465640000000000000060448201526064016101e9565b606061083c8484600085610dc2565b90505b9392505050565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b15801561088257600080fd5b505afa158015610896573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ba91906112d1565b90506108d06001600160a01b038b16828b610f01565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b15801561094657600080fd5b505af115801561095a573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b8152600401610a12959493929190611346565b6000604051808303818588803b158015610a2b57600080fd5b505af1158015610a3f573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b600080836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b158015610a9057600080fd5b505afa158015610aa4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac891906112d1565b9050610ade6001600160a01b038a16828a610f01565b6040516308d18d8960e21b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c821660648401528816608483015282169063234636249060a401600060405180830381600087803b158015610b4857600080fd5b505af1158015610b5c573d6000803e3d6000fd5b505050506000308a8a8a8e8b46604051602001610be09796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858d8b86868c6040518763ffffffff1660e01b8152600401610c2d959493929190611346565b6000604051808303818588803b158015610c4657600080fd5b505af1158015610c5a573d6000803e3d6000fd5b50939e9d5050505050505050505050505050565b600080836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b158015610caa57600080fd5b505afa158015610cbe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce291906112d1565b604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff891660648301529192509082169063de790c7e90608401600060405180830381600087803b158015610d4857600080fd5b505af1158015610d5c573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528d811b82166034840152604883018d90528e901b1660688201526001600160c01b031960c08a811b8216607c84015246901b16608482015260009250608c019050610be0565b606082471015610e3a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101e9565b843b610e885760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101e9565b600080866001600160a01b03168587604051610ea49190611388565b60006040518083038185875af1925050503d8060008114610ee1576040519150601f19603f3d011682016040523d82523d6000602084013e610ee6565b606091505b5091509150610ef6828286610fc2565b979650505050505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b158015610f4d57600080fd5b505afa158015610f61573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8591906113a4565b610f8f91906113bd565b6040516001600160a01b0385166024820152604481018290529091506105de90859063095ea7b360e01b9060640161053d565b60608315610fd157508161083f565b825115610fe15782518084602001fd5b8160405162461bcd60e51b81526004016101e991906113e3565b6001600160a01b038116811461050e57600080fd5b803567ffffffffffffffff8116811461102857600080fd5b919050565b60008083601f84011261103f57600080fd5b50813567ffffffffffffffff81111561105757600080fd5b60208301915083602082850101111561106f57600080fd5b9250929050565b6000806000806060858703121561108c57600080fd5b843561109781610ffb565b93506110a560208601611010565b9250604085013567ffffffffffffffff8111156110c157600080fd5b6110cd8782880161102d565b95989497509550505050565b600080600080606085870312156110ef57600080fd5b84356110fa81610ffb565b935060208501359250604085013567ffffffffffffffff8111156110c157600080fd5b60006020828403121561112f57600080fd5b813561083f81610ffb565b600080600080600080600060e0888a03121561115557600080fd5b873561116081610ffb565b9650602088013561117081610ffb565b95506040880135945061118560608901611010565b935061119360808901611010565b925060a088013563ffffffff811681146111ac57600080fd5b915060c0880135600481106111c057600080fd5b8091505092959891949750929550565b60008060008060008060a087890312156111e957600080fd5b86356111f481610ffb565b9550602087013561120481610ffb565b94506040870135935061121960608801611010565b9250608087013567ffffffffffffffff81111561123557600080fd5b61124189828a0161102d565b979a9699509497509295939492505050565b6001600160a01b038516815283602082015260606040820152816060820152818360808301376000818301608090810191909152601f909201601f191601019392505050565b6000602082840312156112ab57600080fd5b8151801515811461083f57600080fd5b634e487b7160e01b600052602160045260246000fd5b6000602082840312156112e357600080fd5b815161083f81610ffb565b60005b838110156113095781810151838201526020016112f1565b838111156105de5750506000910152565b600081518084526113328160208601602086016112ee565b601f01601f19169290920160200192915050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a06080830152610ef660a083018461131a565b6000825161139a8184602087016112ee565b9190910192915050565b6000602082840312156113b657600080fd5b5051919050565b600082198211156113de57634e487b7160e01b600052601160045260246000fd5b500190565b60208152600061083f602083018461131a56fea2646970667358221220322403764c39058f5547cb9933565593b7ba480924aa81c814e48bf140eddeaf64736f6c63430008090033",
}

// TestRefundABI is the input ABI used to generate the binding from.
// Deprecated: Use TestRefundMetaData.ABI instead.
var TestRefundABI = TestRefundMetaData.ABI

// TestRefundBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestRefundMetaData.Bin instead.
var TestRefundBin = TestRefundMetaData.Bin

// DeployTestRefund deploys a new Ethereum contract, binding an instance of TestRefund to it.
func DeployTestRefund(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address) (common.Address, *types.Transaction, *TestRefund, error) {
	parsed, err := TestRefundMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestRefundBin), backend, _messageBus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestRefund{TestRefundCaller: TestRefundCaller{contract: contract}, TestRefundTransactor: TestRefundTransactor{contract: contract}, TestRefundFilterer: TestRefundFilterer{contract: contract}}, nil
}

// TestRefund is an auto generated Go binding around an Ethereum contract.
type TestRefund struct {
	TestRefundCaller     // Read-only binding to the contract
	TestRefundTransactor // Write-only binding to the contract
	TestRefundFilterer   // Log filterer for contract events
}

// TestRefundCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestRefundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRefundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestRefundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRefundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestRefundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestRefundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestRefundSession struct {
	Contract     *TestRefund       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRefundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestRefundCallerSession struct {
	Contract *TestRefundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestRefundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestRefundTransactorSession struct {
	Contract     *TestRefundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestRefundRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRefundRaw struct {
	Contract *TestRefund // Generic contract binding to access the raw methods on
}

// TestRefundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestRefundCallerRaw struct {
	Contract *TestRefundCaller // Generic read-only contract binding to access the raw methods on
}

// TestRefundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestRefundTransactorRaw struct {
	Contract *TestRefundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestRefund creates a new instance of TestRefund, bound to a specific deployed contract.
func NewTestRefund(address common.Address, backend bind.ContractBackend) (*TestRefund, error) {
	contract, err := bindTestRefund(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestRefund{TestRefundCaller: TestRefundCaller{contract: contract}, TestRefundTransactor: TestRefundTransactor{contract: contract}, TestRefundFilterer: TestRefundFilterer{contract: contract}}, nil
}

// NewTestRefundCaller creates a new read-only instance of TestRefund, bound to a specific deployed contract.
func NewTestRefundCaller(address common.Address, caller bind.ContractCaller) (*TestRefundCaller, error) {
	contract, err := bindTestRefund(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestRefundCaller{contract: contract}, nil
}

// NewTestRefundTransactor creates a new write-only instance of TestRefund, bound to a specific deployed contract.
func NewTestRefundTransactor(address common.Address, transactor bind.ContractTransactor) (*TestRefundTransactor, error) {
	contract, err := bindTestRefund(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestRefundTransactor{contract: contract}, nil
}

// NewTestRefundFilterer creates a new log filterer instance of TestRefund, bound to a specific deployed contract.
func NewTestRefundFilterer(address common.Address, filterer bind.ContractFilterer) (*TestRefundFilterer, error) {
	contract, err := bindTestRefund(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestRefundFilterer{contract: contract}, nil
}

// bindTestRefund binds a generic wrapper to an already deployed contract.
func bindTestRefund(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestRefundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRefund *TestRefundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestRefund.Contract.TestRefundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRefund *TestRefundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRefund.Contract.TestRefundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRefund *TestRefundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRefund.Contract.TestRefundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestRefund *TestRefundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestRefund.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestRefund *TestRefundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestRefund.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestRefund *TestRefundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestRefund.Contract.contract.Transact(opts, method, params...)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TestRefund *TestRefundCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestRefund.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TestRefund *TestRefundSession) MessageBus() (common.Address, error) {
	return _TestRefund.Contract.MessageBus(&_TestRefund.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TestRefund *TestRefundCallerSession) MessageBus() (common.Address, error) {
	return _TestRefund.Contract.MessageBus(&_TestRefund.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestRefund *TestRefundCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestRefund.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestRefund *TestRefundSession) Owner() (common.Address, error) {
	return _TestRefund.Contract.Owner(&_TestRefund.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TestRefund *TestRefundCallerSession) Owner() (common.Address, error) {
	return _TestRefund.Contract.Owner(&_TestRefund.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 , bytes ) payable returns(bool)
func (_TestRefund *TestRefundTransactor) ExecuteMessage(opts *bind.TransactOpts, arg0 common.Address, arg1 uint64, arg2 []byte) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "executeMessage", arg0, arg1, arg2)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 , bytes ) payable returns(bool)
func (_TestRefund *TestRefundSession) ExecuteMessage(arg0 common.Address, arg1 uint64, arg2 []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessage(&_TestRefund.TransactOpts, arg0, arg1, arg2)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 , bytes ) payable returns(bool)
func (_TestRefund *TestRefundTransactorSession) ExecuteMessage(arg0 common.Address, arg1 uint64, arg2 []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessage(&_TestRefund.TransactOpts, arg0, arg1, arg2)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessageWithTransfer(&_TestRefund.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessageWithTransfer(&_TestRefund.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessageWithTransferFallback(&_TestRefund.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessageWithTransferFallback(&_TestRefund.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessageWithTransferRefund(&_TestRefund.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessageWithTransferRefund(&_TestRefund.TransactOpts, _token, _amount, _message)
}

// SendWithTransfer is a paid mutator transaction binding the contract method 0x7ee1bd4a.
//
// Solidity: function sendWithTransfer(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeType) payable returns()
func (_TestRefund *TestRefundTransactor) SendWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeType uint8) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "sendWithTransfer", _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeType)
}

// SendWithTransfer is a paid mutator transaction binding the contract method 0x7ee1bd4a.
//
// Solidity: function sendWithTransfer(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeType) payable returns()
func (_TestRefund *TestRefundSession) SendWithTransfer(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeType uint8) (*types.Transaction, error) {
	return _TestRefund.Contract.SendWithTransfer(&_TestRefund.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeType)
}

// SendWithTransfer is a paid mutator transaction binding the contract method 0x7ee1bd4a.
//
// Solidity: function sendWithTransfer(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage, uint8 _bridgeType) payable returns()
func (_TestRefund *TestRefundTransactorSession) SendWithTransfer(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32, _bridgeType uint8) (*types.Transaction, error) {
	return _TestRefund.Contract.SendWithTransfer(&_TestRefund.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage, _bridgeType)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TestRefund *TestRefundTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TestRefund *TestRefundSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TestRefund.Contract.SetMessageBus(&_TestRefund.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TestRefund *TestRefundTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TestRefund.Contract.SetMessageBus(&_TestRefund.TransactOpts, _messageBus)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestRefund *TestRefundTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestRefund *TestRefundSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestRefund.Contract.TransferOwnership(&_TestRefund.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TestRefund *TestRefundTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TestRefund.Contract.TransferOwnership(&_TestRefund.TransactOpts, newOwner)
}

// TestRefundOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TestRefund contract.
type TestRefundOwnershipTransferredIterator struct {
	Event *TestRefundOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TestRefundOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestRefundOwnershipTransferred)
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
		it.Event = new(TestRefundOwnershipTransferred)
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
func (it *TestRefundOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestRefundOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestRefundOwnershipTransferred represents a OwnershipTransferred event raised by the TestRefund contract.
type TestRefundOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestRefund *TestRefundFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TestRefundOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestRefund.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TestRefundOwnershipTransferredIterator{contract: _TestRefund.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TestRefund *TestRefundFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TestRefundOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TestRefund.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestRefundOwnershipTransferred)
				if err := _TestRefund.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TestRefund *TestRefundFilterer) ParseOwnershipTransferred(log types.Log) (*TestRefundOwnershipTransferred, error) {
	event := new(TestRefundOwnershipTransferred)
	if err := _TestRefund.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestRefundRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the TestRefund contract.
type TestRefundRefundedIterator struct {
	Event *TestRefundRefunded // Event containing the contract specifics and raw log

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
func (it *TestRefundRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestRefundRefunded)
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
		it.Event = new(TestRefundRefunded)
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
func (it *TestRefundRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestRefundRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestRefundRefunded represents a Refunded event raised by the TestRefund contract.
type TestRefundRefunded struct {
	Token   common.Address
	Amount  *big.Int
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xee93b3d90bf7307985573224b309958a34f05c4bc6e541d3c2a9f903651cb3b6.
//
// Solidity: event Refunded(address token, uint256 amount, bytes message)
func (_TestRefund *TestRefundFilterer) FilterRefunded(opts *bind.FilterOpts) (*TestRefundRefundedIterator, error) {

	logs, sub, err := _TestRefund.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &TestRefundRefundedIterator{contract: _TestRefund.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xee93b3d90bf7307985573224b309958a34f05c4bc6e541d3c2a9f903651cb3b6.
//
// Solidity: event Refunded(address token, uint256 amount, bytes message)
func (_TestRefund *TestRefundFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *TestRefundRefunded) (event.Subscription, error) {

	logs, sub, err := _TestRefund.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestRefundRefunded)
				if err := _TestRefund.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0xee93b3d90bf7307985573224b309958a34f05c4bc6e541d3c2a9f903651cb3b6.
//
// Solidity: event Refunded(address token, uint256 amount, bytes message)
func (_TestRefund *TestRefundFilterer) ParseRefunded(log types.Log) (*TestRefundRefunded, error) {
	event := new(TestRefundRefunded)
	if err := _TestRefund.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferMessageMetaData contains all meta data concerning the TransferMessage contract.
var TransferMessageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"transferMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161098838038061098883398101604081905261002f916100ad565b6100383361005d565b600180546001600160a01b0319166001600160a01b03929092169190911790556100dd565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bf57600080fd5b81516001600160a01b03811681146100d657600080fd5b9392505050565b61089c806100ec6000396000f3fe6080604052600436106100965760003560e01c80638da5cb5b11610069578063ce35dd9a1161004e578063ce35dd9a1461015d578063f00f39ce1461015d578063f2fde38b1461017057600080fd5b80638da5cb5b1461010b578063a1a227fa1461013d57600080fd5b80631599d2651461009b57806320be95f2146100c3578063547cad12146100d6578063867fd811146100f8575b600080fd5b6100ae6100a93660046105c3565b610190565b60405190151581526020015b60405180910390f35b6100ae6100d1366004610624565b6101fe565b3480156100e257600080fd5b506100f66100f1366004610666565b61025b565b005b6100f661010636600461069e565b6102f3565b34801561011757600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016100ba565b34801561014957600080fd5b50600154610125906001600160a01b031681565b6100ae61016b366004610770565b610304565b34801561017c57600080fd5b506100f661018b366004610666565b61036b565b6001546000906001600160a01b031633146101f25760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b5060015b949350505050565b6001546000906001600160a01b031633146101f65760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016101e9565b3361026e6000546001600160a01b031690565b6001600160a01b0316146102c45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101e9565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6102ff8383833461045c565b505050565b6001546000906001600160a01b031633146103615760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016101e9565b9695505050505050565b3361037e6000546001600160a01b031690565b6001600160a01b0316146103d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101e9565b6001600160a01b0381166104505760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016101e9565b6104598161047e565b50565b600154610478908590859085906001600160a01b0316856104db565b50505050565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b604051634f9e72ad60e11b81526001600160a01b03831690639f3ce55a90839061050d908990899089906004016107ef565b6000604051808303818588803b15801561052657600080fd5b505af115801561053a573d6000803e3d6000fd5b50505050505050505050565b80356001600160a01b038116811461055d57600080fd5b919050565b803567ffffffffffffffff8116811461055d57600080fd5b60008083601f84011261058c57600080fd5b50813567ffffffffffffffff8111156105a457600080fd5b6020830191508360208285010111156105bc57600080fd5b9250929050565b600080600080606085870312156105d957600080fd5b6105e285610546565b93506105f060208601610562565b9250604085013567ffffffffffffffff81111561060c57600080fd5b6106188782880161057a565b95989497509550505050565b6000806000806060858703121561063a57600080fd5b61064385610546565b935060208501359250604085013567ffffffffffffffff81111561060c57600080fd5b60006020828403121561067857600080fd5b61068182610546565b9392505050565b634e487b7160e01b600052604160045260246000fd5b6000806000606084860312156106b357600080fd5b6106bc84610546565b92506106ca60208501610562565b9150604084013567ffffffffffffffff808211156106e757600080fd5b818601915086601f8301126106fb57600080fd5b81358181111561070d5761070d610688565b604051601f8201601f19908116603f0116810190838211818310171561073557610735610688565b8160405282815289602084870101111561074e57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060008060008060a0878903121561078957600080fd5b61079287610546565b95506107a060208801610546565b9450604087013593506107b560608801610562565b9250608087013567ffffffffffffffff8111156107d157600080fd5b6107dd89828a0161057a565b979a9699509497509295939492505050565b6001600160a01b03841681526000602067ffffffffffffffff85168184015260606040840152835180606085015260005b8181101561083c57858101830151858201608001528201610820565b8181111561084e576000608083870101525b50601f01601f1916929092016080019594505050505056fea2646970667358221220851a26e00ef775381bddfc6dd94b503732be58417237fd3804da7a5c597204b364736f6c63430008090033",
}

// TransferMessageABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferMessageMetaData.ABI instead.
var TransferMessageABI = TransferMessageMetaData.ABI

// TransferMessageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferMessageMetaData.Bin instead.
var TransferMessageBin = TransferMessageMetaData.Bin

// DeployTransferMessage deploys a new Ethereum contract, binding an instance of TransferMessage to it.
func DeployTransferMessage(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address) (common.Address, *types.Transaction, *TransferMessage, error) {
	parsed, err := TransferMessageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferMessageBin), backend, _messageBus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferMessage{TransferMessageCaller: TransferMessageCaller{contract: contract}, TransferMessageTransactor: TransferMessageTransactor{contract: contract}, TransferMessageFilterer: TransferMessageFilterer{contract: contract}}, nil
}

// TransferMessage is an auto generated Go binding around an Ethereum contract.
type TransferMessage struct {
	TransferMessageCaller     // Read-only binding to the contract
	TransferMessageTransactor // Write-only binding to the contract
	TransferMessageFilterer   // Log filterer for contract events
}

// TransferMessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferMessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferMessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferMessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferMessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferMessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferMessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferMessageSession struct {
	Contract     *TransferMessage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferMessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferMessageCallerSession struct {
	Contract *TransferMessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TransferMessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferMessageTransactorSession struct {
	Contract     *TransferMessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TransferMessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferMessageRaw struct {
	Contract *TransferMessage // Generic contract binding to access the raw methods on
}

// TransferMessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferMessageCallerRaw struct {
	Contract *TransferMessageCaller // Generic read-only contract binding to access the raw methods on
}

// TransferMessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferMessageTransactorRaw struct {
	Contract *TransferMessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferMessage creates a new instance of TransferMessage, bound to a specific deployed contract.
func NewTransferMessage(address common.Address, backend bind.ContractBackend) (*TransferMessage, error) {
	contract, err := bindTransferMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferMessage{TransferMessageCaller: TransferMessageCaller{contract: contract}, TransferMessageTransactor: TransferMessageTransactor{contract: contract}, TransferMessageFilterer: TransferMessageFilterer{contract: contract}}, nil
}

// NewTransferMessageCaller creates a new read-only instance of TransferMessage, bound to a specific deployed contract.
func NewTransferMessageCaller(address common.Address, caller bind.ContractCaller) (*TransferMessageCaller, error) {
	contract, err := bindTransferMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferMessageCaller{contract: contract}, nil
}

// NewTransferMessageTransactor creates a new write-only instance of TransferMessage, bound to a specific deployed contract.
func NewTransferMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferMessageTransactor, error) {
	contract, err := bindTransferMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferMessageTransactor{contract: contract}, nil
}

// NewTransferMessageFilterer creates a new log filterer instance of TransferMessage, bound to a specific deployed contract.
func NewTransferMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferMessageFilterer, error) {
	contract, err := bindTransferMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferMessageFilterer{contract: contract}, nil
}

// bindTransferMessage binds a generic wrapper to an already deployed contract.
func bindTransferMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferMessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferMessage *TransferMessageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferMessage.Contract.TransferMessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferMessage *TransferMessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferMessage.Contract.TransferMessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferMessage *TransferMessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferMessage.Contract.TransferMessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferMessage *TransferMessageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferMessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferMessage *TransferMessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferMessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferMessage *TransferMessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferMessage.Contract.contract.Transact(opts, method, params...)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TransferMessage *TransferMessageCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferMessage.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TransferMessage *TransferMessageSession) MessageBus() (common.Address, error) {
	return _TransferMessage.Contract.MessageBus(&_TransferMessage.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TransferMessage *TransferMessageCallerSession) MessageBus() (common.Address, error) {
	return _TransferMessage.Contract.MessageBus(&_TransferMessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferMessage *TransferMessageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferMessage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferMessage *TransferMessageSession) Owner() (common.Address, error) {
	return _TransferMessage.Contract.Owner(&_TransferMessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferMessage *TransferMessageCallerSession) Owner() (common.Address, error) {
	return _TransferMessage.Contract.Owner(&_TransferMessage.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 , bytes ) payable returns(bool)
func (_TransferMessage *TransferMessageTransactor) ExecuteMessage(opts *bind.TransactOpts, arg0 common.Address, arg1 uint64, arg2 []byte) (*types.Transaction, error) {
	return _TransferMessage.contract.Transact(opts, "executeMessage", arg0, arg1, arg2)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 , bytes ) payable returns(bool)
func (_TransferMessage *TransferMessageSession) ExecuteMessage(arg0 common.Address, arg1 uint64, arg2 []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessage(&_TransferMessage.TransactOpts, arg0, arg1, arg2)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address , uint64 , bytes ) payable returns(bool)
func (_TransferMessage *TransferMessageTransactorSession) ExecuteMessage(arg0 common.Address, arg1 uint64, arg2 []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessage(&_TransferMessage.TransactOpts, arg0, arg1, arg2)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessageWithTransfer(&_TransferMessage.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessageWithTransfer(&_TransferMessage.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessageWithTransferFallback(&_TransferMessage.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessageWithTransferFallback(&_TransferMessage.TransactOpts, _sender, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessageWithTransferRefund(&_TransferMessage.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TransferMessage *TransferMessageTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.ExecuteMessageWithTransferRefund(&_TransferMessage.TransactOpts, _token, _amount, _message)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TransferMessage *TransferMessageTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _TransferMessage.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TransferMessage *TransferMessageSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TransferMessage.Contract.SetMessageBus(&_TransferMessage.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TransferMessage *TransferMessageTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TransferMessage.Contract.SetMessageBus(&_TransferMessage.TransactOpts, _messageBus)
}

// TransferMessage is a paid mutator transaction binding the contract method 0x867fd811.
//
// Solidity: function transferMessage(address _receiver, uint64 _dstChainId, bytes _message) payable returns()
func (_TransferMessage *TransferMessageTransactor) TransferMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.contract.Transact(opts, "transferMessage", _receiver, _dstChainId, _message)
}

// TransferMessage is a paid mutator transaction binding the contract method 0x867fd811.
//
// Solidity: function transferMessage(address _receiver, uint64 _dstChainId, bytes _message) payable returns()
func (_TransferMessage *TransferMessageSession) TransferMessage(_receiver common.Address, _dstChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.TransferMessage(&_TransferMessage.TransactOpts, _receiver, _dstChainId, _message)
}

// TransferMessage is a paid mutator transaction binding the contract method 0x867fd811.
//
// Solidity: function transferMessage(address _receiver, uint64 _dstChainId, bytes _message) payable returns()
func (_TransferMessage *TransferMessageTransactorSession) TransferMessage(_receiver common.Address, _dstChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferMessage.Contract.TransferMessage(&_TransferMessage.TransactOpts, _receiver, _dstChainId, _message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferMessage *TransferMessageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TransferMessage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferMessage *TransferMessageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferMessage.Contract.TransferOwnership(&_TransferMessage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferMessage *TransferMessageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferMessage.Contract.TransferOwnership(&_TransferMessage.TransactOpts, newOwner)
}

// TransferMessageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TransferMessage contract.
type TransferMessageOwnershipTransferredIterator struct {
	Event *TransferMessageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TransferMessageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferMessageOwnershipTransferred)
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
		it.Event = new(TransferMessageOwnershipTransferred)
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
func (it *TransferMessageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferMessageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferMessageOwnershipTransferred represents a OwnershipTransferred event raised by the TransferMessage contract.
type TransferMessageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferMessage *TransferMessageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TransferMessageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferMessage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TransferMessageOwnershipTransferredIterator{contract: _TransferMessage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferMessage *TransferMessageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TransferMessageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferMessage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferMessageOwnershipTransferred)
				if err := _TransferMessage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TransferMessage *TransferMessageFilterer) ParseOwnershipTransferred(log types.Log) (*TransferMessageOwnershipTransferred, error) {
	event := new(TransferMessageOwnershipTransferred)
	if err := _TransferMessage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferSwapMetaData contains all meta data concerning the TransferSwap contract.
var TransferSwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_supportedDex\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nativeWrap\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"DirectSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumTransferSwap.SwapStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"SwapRequestDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstToken\",\"type\":\"address\"}],\"name\":\"SwapRequestSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minSwapAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minSwapAmount\",\"type\":\"uint256\"}],\"name\":\"setMinSwapAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeWrap\",\"type\":\"address\"}],\"name\":\"setNativeWrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dex\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSupportedDex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRecvAmt\",\"type\":\"uint256\"}],\"internalType\":\"structTransferSwap.SwapInfo\",\"name\":\"_srcSwap\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRecvAmt\",\"type\":\"uint256\"}],\"internalType\":\"structTransferSwap.SwapInfo\",\"name\":\"_dstSwap\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_maxBridgeSlippage\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"transferWithSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRecvAmt\",\"type\":\"uint256\"}],\"internalType\":\"structTransferSwap.SwapInfo\",\"name\":\"_srcSwap\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"dex\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRecvAmt\",\"type\":\"uint256\"}],\"internalType\":\"structTransferSwap.SwapInfo\",\"name\":\"_dstSwap\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_maxBridgeSlippage\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_nativeOut\",\"type\":\"bool\"}],\"name\":\"transferWithSwapNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002ba638038062002ba68339810160408190526200003491620000f9565b6200003f336200008c565b600180546001600160a01b039485166001600160a01b03199182161782559284166000908152600360205260409020805460ff191690911790556004805491909316911617905562000143565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000f457600080fd5b919050565b6000806000606084860312156200010f57600080fd5b6200011a84620000dc565b92506200012a60208501620000dc565b91506200013a60408501620000dc565b90509250925092565b612a5380620001536000396000f3fe6080604052600436106100ec5760003560e01c80635fdb3d8d1161008a578063ce35dd9a11610059578063ce35dd9a14610267578063d9b796de1461027a578063f00f39ce1461029a578063f2fde38b146102ad57600080fd5b80635fdb3d8d146102035780638da5cb5b14610216578063a1a227fa14610234578063a8a7e29c1461025457600080fd5b8063457bfa2f116100c6578063457bfa2f14610169578063496dbfd9146101a1578063547cad12146101c35780635b5a66a7146101e357600080fd5b80631599d265146100f857806320be95f21461012057806334e181251461012e57600080fd5b366100f357005b600080fd5b61010b610106366004611f32565b6102cd565b60405190151581526020015b60405180910390f35b61010b610106366004611f97565b34801561013a57600080fd5b5061015b610149366004611fdb565b60026020526000908152604090205481565b604051908152602001610117565b34801561017557600080fd5b50600454610189906001600160a01b031681565b6040516001600160a01b039091168152602001610117565b3480156101ad57600080fd5b506101c16101bc366004611ff8565b610337565b005b3480156101cf57600080fd5b506101c16101de366004611fdb565b6103bc565b3480156101ef57600080fd5b506101c16101fe366004611fdb565b610454565b6101c161021136600461205e565b6104ec565b34801561022257600080fd5b506000546001600160a01b0316610189565b34801561024057600080fd5b50600154610189906001600160a01b031681565b6101c1610262366004612129565b61068d565b61010b610275366004612251565b610738565b34801561028657600080fd5b506101c1610295366004612328565b61099f565b61010b6102a8366004612251565b610a33565b3480156102b957600080fd5b506101c16102c8366004611fdb565b610b07565b6001546000906001600160a01b0316331461032f5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b949350505050565b3361034a6000546001600160a01b031690565b6001600160a01b0316146103a05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001600160a01b03909116600090815260026020526040902055565b336103cf6000546001600160a01b031690565b6001600160a01b0316146104255760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b336104676000546001600160a01b031690565b6001600160a01b0316146104bd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6004805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b3332146105255760405162461bcd60e51b81526020600482015260076024820152664e6f7420454f4160c81b6044820152606401610326565b863410156105755760405162461bcd60e51b815260206004820152601360248201527f416d6f756e7420696e73756666696369656e74000000000000000000000000006044820152606401610326565b6004546001600160a01b031661058b8680612361565b600081811061059c5761059c6123ab565b90506020020160208101906105b19190611fdb565b6001600160a01b0316146105f85760405162461bcd60e51b815260206004820152600e60248201526d0e8ded6cadc40dad2e6dac2e8c6d60931b6044820152606401610326565b6004805460408051630d0e30db60e41b815290516001600160a01b039092169263d0e30db0928b92808301926000929182900301818588803b15801561063d57600080fd5b505af1158015610651573d6000803e3d6000fd5b505050505061068388888888610666906123e5565b61066f896123e5565b8888888f3461067e91906124c7565b610bf8565b5050505050505050565b3332146106c65760405162461bcd60e51b81526020600482015260076024820152664e6f7420454f4160c81b6044820152606401610326565b61070d3330886106d68880612361565b60008181106106e7576106e76123ab565b90506020020160208101906106fc9190611fdb565b6001600160a01b0316929190610e5e565b61072f87878761071c886123e5565b610725886123e5565b8787600034610bf8565b50505050505050565b6001546000906001600160a01b031633146107955760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610326565b6000828060200190518101906107ab91906124ff565b8051518051919250906000906107c3576107c36123ab565b60200260200101516001600160a01b0316866001600160a01b0316146108775760405162461bcd60e51b815260206004820152604a60248201527f6272696467656420746f6b656e206d757374206265207468652073616d65206160448201527f732074686520666972737420746f6b656e20696e2064657374696e6174696f6e60648201527f2073776170207061746800000000000000000000000000000000000000000000608482015260a401610326565b60006108898260200151864687610efc565b8251515190915060009060019081101561091c5783516001906108ac908a610f35565b9350905080156108fd5784515180516108f491906108cc906001906124c7565b815181106108dc576108dc6123ab565b60200260200101518487602001518860600151611082565b60019150610916565b61090e8a8a87602001516000611082565b889250600391505b50610954565b835151805161094d9190600090610935576109356123ab565b60200260200101518986602001518760600151611082565b5086905060015b7fccbb695db2dfb79e3f69bcf724d11b850c79b0ae358e62d60ed94044f38a2b9c83838360405161098793929190612685565b60405180910390a15060019998505050505050505050565b336109b26000546001600160a01b031690565b6001600160a01b031614610a085760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001600160a01b03919091166000908152600360205260409020805460ff1916911515919091179055565b6001546000906001600160a01b03163314610a905760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610326565b600082806020019051810190610aa691906124ff565b90506000610aba8260200151864687610efc565b90507fccbb695db2dfb79e3f69bcf724d11b850c79b0ae358e62d60ed94044f38a2b9c8160006002604051610af193929190612685565b60405180910390a1506000979650505050505050565b33610b1a6000546001600160a01b031690565b6001600160a01b031614610b705760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001600160a01b038116610bec5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610326565b610bf5816111f4565b50565b855151610c475760405162461bcd60e51b815260206004820152601360248201527f656d7074792073726320737761702070617468000000000000000000000000006044820152606401610326565b8551805160009190610c5b906001906124c7565b81518110610c6b57610c6b6123ab565b60200260200101519050600260008860000151600081518110610c9057610c906123ab565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548911610d2d5760405162461bcd60e51b815260206004820152602b60248201527f616d6f756e74206d7573742062652067726561746572207468616e206d696e2060448201527f7377617020616d6f756e740000000000000000000000000000000000000000006064820152608401610326565b865151469060011080610d5457508067ffffffffffffffff168967ffffffffffffffff1614155b610da05760405162461bcd60e51b815260206004820152601360248201527f6e6f6f70206973206e6f7420616c6c6f776564000000000000000000000000006044820152606401610326565b8751518a9060011015610e0c576001610db98a8d610f35565b9250905080610e0a5760405162461bcd60e51b815260206004820152600f60248201527f7372632073776170206661696c656400000000000000000000000000000000006044820152606401610326565b505b8167ffffffffffffffff168a67ffffffffffffffff161415610e3c57610e378c8c848c8a8887611251565b610e50565b610e508c8c848d8d8d8d8d8d8d8d8c611330565b505050505050505050505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610ef69085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526114c6565b50505050565b600084848484604051602001610f1594939291906126cc565b604051602081830303815290604052805190602001209050949350505050565b6020808301516001600160a01b03166000908152600390915260408120548190819060ff16610f695760009250905061107b565b610fa98560200151858760000151600081518110610f8957610f896123ab565b60200260200101516001600160a01b03166115b09092919063ffffffff16565b84602001516001600160a01b03166338ed17398587606001518860000151308a604001516040518663ffffffff1660e01b8152600401610fed95949392919061276a565b600060405180830381600087803b15801561100757600080fd5b505af192505050801561103c57506040513d6000823e601f3d908101601f1916820160405261103991908101906127a6565b60015b61104b5760009250905061107b565b6001816001835161105c91906124c7565b8151811061106c5761106c6123ab565b60200260200101519350935050505b9250929050565b80156111e0576004546001600160a01b038581169116146110d65760405162461bcd60e51b815260206004820152600e60248201526d0e8ded6cadc40dad2e6dac2e8c6d60931b6044820152606401610326565b60048054604051632e1a7d4d60e01b81529182018590526001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561111a57600080fd5b505af115801561112e573d6000803e3d6000fd5b505050506000826001600160a01b03168461c35090604051600060405180830381858888f193505050503d8060008114611184576040519150601f19603f3d011682016040523d82523d6000602084013e611189565b606091505b50509050806111da5760405162461bcd60e51b815260206004820152601560248201527f6661696c656420746f2073656e64206e617469766500000000000000000000006044820152606401610326565b50610ef6565b610ef66001600160a01b0385168385611671565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6112656001600160a01b0383168883611671565b60003386898688604051602001611280959493929190612874565b6040516020818303038152906040528051906020012090507fdbd876e57420a75850bd1ce36de349d111ab4ff35e2c2379af8acc750376917481878988600001516000815181106112d3576112d36123ab565b6020908102919091018101516040805195865267ffffffffffffffff90941691850191909152918301526001600160a01b03908116606083015260808201859052851660a082015260c00160405180910390a15050505050505050565b86515161137f5760405162461bcd60e51b815260206004820152601360248201527f656d7074792064737420737761702070617468000000000000000000000000006044820152606401610326565b60006040518060800160405280898152602001336001600160a01b031681526020018767ffffffffffffffff1681526020018615158152506040516020016113c791906128ba565b604051602081830303815290604052905060006113e6338d8d85610efc565b90506113fa8e85858e8b8d8860018d6116a1565b507f600b63623f19b18f8f0e7825ec78b3d9779d8124c76ad75b58aac7167f04f8d2818c8f8d60000151600081518110611436576114366123ab565b60209081029190910101518d518051611451906001906124c7565b81518110611461576114616123ab565b60200260200101516040516114ae95949392919094855267ffffffffffffffff93909316602085015260408401919091526001600160a01b03908116606084015216608082015260a00190565b60405180910390a15050505050505050505050505050565b600061151b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116d89092919063ffffffff16565b8051909150156115ab57808060200190518101906115399190612915565b6115ab5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610326565b505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b1580156115fc57600080fd5b505afa158015611610573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116349190612932565b61163e919061294b565b6040516001600160a01b038516602482015260448101829052909150610ef690859063095ea7b360e01b90606401610e92565b6040516001600160a01b0383166024820152604481018290526115ab90849063a9059cbb60e01b90606401610e92565b60006116ca8a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b6116f1565b9a9950505050505050505050565b60606116e784846000856117c1565b90505b9392505050565b600060018460038111156117075761170761264d565b14156117255761171e8b8b8b8b8b8b8b8a8a611900565b90506116ca565b60028460038111156117395761173961264d565b141561174f5761171e8b8b8b8b8b8a8989611b0e565b60038460038111156117635761176361264d565b14156117795761171e8b8b8b8b8b8a8989611d28565b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f72746564000000000000006044820152606401610326565b6060824710156118395760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610326565b843b6118875760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610326565b600080866001600160a01b031685876040516118a39190612963565b60006040518083038185875af1925050503d80600081146118e0576040519150601f19603f3d011682016040523d82523d6000602084013e6118e5565b606091505b50915091506118f5828286611e7c565b979650505050505050565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b15801561193c57600080fd5b505afa158015611950573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611974919061297f565b905061198a6001600160a01b038b16828b6115b0565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b158015611a0057600080fd5b505af1158015611a14573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b8152600401611acc9594939291906129c8565b6000604051808303818588803b158015611ae557600080fd5b505af1158015611af9573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b600080836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b158015611b4a57600080fd5b505afa158015611b5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b82919061297f565b9050611b986001600160a01b038a16828a6115b0565b6040516308d18d8960e21b81526001600160a01b038a81166004830152602482018a905267ffffffffffffffff808a1660448401528c821660648401528816608483015282169063234636249060a401600060405180830381600087803b158015611c0257600080fd5b505af1158015611c16573d6000803e3d6000fd5b505050506000308a8a8a8e8b46604051602001611c9a9796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858d8b86868c6040518763ffffffff1660e01b8152600401611ce79594939291906129c8565b6000604051808303818588803b158015611d0057600080fd5b505af1158015611d14573d6000803e3d6000fd5b50939e9d5050505050505050505050505050565b600080836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b158015611d6457600080fd5b505afa158015611d78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d9c919061297f565b604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff891660648301529192509082169063de790c7e90608401600060405180830381600087803b158015611e0257600080fd5b505af1158015611e16573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528d811b82166034840152604883018d90528e901b1660688201526001600160c01b031960c08a811b8216607c84015246901b16608482015260009250608c019050611c9a565b60608315611e8b5750816116ea565b825115611e9b5782518084602001fd5b8160405162461bcd60e51b81526004016103269190612a0a565b6001600160a01b0381168114610bf557600080fd5b8035611ed581611eb5565b919050565b67ffffffffffffffff81168114610bf557600080fd5b60008083601f840112611f0257600080fd5b50813567ffffffffffffffff811115611f1a57600080fd5b60208301915083602082850101111561107b57600080fd5b60008060008060608587031215611f4857600080fd5b8435611f5381611eb5565b93506020850135611f6381611eda565b9250604085013567ffffffffffffffff811115611f7f57600080fd5b611f8b87828801611ef0565b95989497509550505050565b60008060008060608587031215611fad57600080fd5b8435611fb881611eb5565b935060208501359250604085013567ffffffffffffffff811115611f7f57600080fd5b600060208284031215611fed57600080fd5b81356116ea81611eb5565b6000806040838503121561200b57600080fd5b823561201681611eb5565b946020939093013593505050565b60006080828403121561203657600080fd5b50919050565b803563ffffffff81168114611ed557600080fd5b8015158114610bf557600080fd5b600080600080600080600080610100898b03121561207b57600080fd5b883561208681611eb5565b975060208901359650604089013561209d81611eda565b9550606089013567ffffffffffffffff808211156120ba57600080fd5b6120c68c838d01612024565b965060808b01359150808211156120dc57600080fd5b506120e98b828c01612024565b9450506120f860a08a0161203c565b925060c089013561210881611eda565b915060e089013561211881612050565b809150509295985092959890939650565b600080600080600080600060e0888a03121561214457600080fd5b873561214f81611eb5565b965060208801359550604088013561216681611eda565b9450606088013567ffffffffffffffff8082111561218357600080fd5b61218f8b838c01612024565b955060808a01359150808211156121a557600080fd5b506121b28a828b01612024565b9350506121c160a0890161203c565b915060c08801356121d181611eda565b8091505092959891949750929550565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561221a5761221a6121e1565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715612249576122496121e1565b604052919050565b600080600080600060a0868803121561226957600080fd5b853561227481611eb5565b945060208681013561228581611eb5565b945060408701359350606087013561229c81611eda565b9250608087013567ffffffffffffffff808211156122b957600080fd5b818901915089601f8301126122cd57600080fd5b8135818111156122df576122df6121e1565b6122f1601f8201601f19168501612220565b91508082528a8482850101111561230757600080fd5b80848401858401376000848284010152508093505050509295509295909350565b6000806040838503121561233b57600080fd5b823561234681611eb5565b9150602083013561235681612050565b809150509250929050565b6000808335601e1984360301811261237857600080fd5b83018035915067ffffffffffffffff82111561239357600080fd5b6020019150600581901b360382131561107b57600080fd5b634e487b7160e01b600052603260045260246000fd5b600067ffffffffffffffff8211156123db576123db6121e1565b5060051b60200190565b6000608082360312156123f757600080fd5b6123ff6121f7565b823567ffffffffffffffff81111561241657600080fd5b830136601f82011261242757600080fd5b8035602061243c612437836123c1565b612220565b82815260059290921b8301810191818101903684111561245b57600080fd5b938201935b8385101561248257843561247381611eb5565b82529382019390820190612460565b855250612490868201611eca565b90840152505060408381013590820152606092830135928101929092525090565b634e487b7160e01b600052601160045260246000fd5b6000828210156124d9576124d96124b1565b500390565b8051611ed581611eb5565b8051611ed581611eda565b8051611ed581612050565b6000602080838503121561251257600080fd5b825167ffffffffffffffff8082111561252a57600080fd5b908401906080828703121561253e57600080fd5b6125466121f7565b82518281111561255557600080fd5b83016080818903121561256757600080fd5b61256f6121f7565b81518481111561257e57600080fd5b82019350601f8401891361259157600080fd5b835161259f612437826123c1565b81815260059190911b8501870190878101908b8311156125be57600080fd5b958801955b828710156125e55786516125d681611eb5565b825295880195908801906125c3565b8352506125f590508287016124de565b868201526040820151604082015260608201516060820152808352505061261d8484016124de565b8482015261262d604084016124e9565b604082015261263e606084016124f4565b60608201529695505050505050565b634e487b7160e01b600052602160045260246000fd5b6004811061268157634e487b7160e01b600052602160045260246000fd5b9052565b838152602081018390526060810161032f6040830184612663565b60005b838110156126bb5781810151838201526020016126a3565b83811115610ef65750506000910152565b6bffffffffffffffffffffffff198560601b16815260006001600160c01b0319808660c01b166014840152808560c01b16601c8401525082516127168160248501602087016126a0565b9190910160240195945050505050565b600081518084526020808501945080840160005b8381101561275f5781516001600160a01b03168752958201959082019060010161273a565b509495945050505050565b85815284602082015260a06040820152600061278960a0830186612726565b6001600160a01b0394909416606083015250608001529392505050565b600060208083850312156127b957600080fd5b825167ffffffffffffffff8111156127d057600080fd5b8301601f810185136127e157600080fd5b80516127ef612437826123c1565b81815260059190911b8201830190838101908783111561280e57600080fd5b928401925b828410156118f557835182529284019290840190612813565b60008151608084526128416080850182612726565b90506001600160a01b03602084015116602085015260408301516040850152606083015160608501528091505092915050565b60006001600160a01b03808816835267ffffffffffffffff808816602085015281871660408501528086166060850152505060a060808301526118f560a083018461282c565b6020815260008251608060208401526128d660a084018261282c565b90506001600160a01b03602085015116604084015267ffffffffffffffff60408501511660608401526060840151151560808401528091505092915050565b60006020828403121561292757600080fd5b81516116ea81612050565b60006020828403121561294457600080fd5b5051919050565b6000821982111561295e5761295e6124b1565b500190565b600082516129758184602087016126a0565b9190910192915050565b60006020828403121561299157600080fd5b81516116ea81611eb5565b600081518084526129b48160208601602086016126a0565b601f01601f19169290920160200192915050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a060808301526118f560a083018461299c565b6020815260006116ea602083018461299c56fea264697066735822122016ea99c48855c219231875f9ba42ed95ec50501402dd86d275a48aa89f429bd664736f6c63430008090033",
}

// TransferSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferSwapMetaData.ABI instead.
var TransferSwapABI = TransferSwapMetaData.ABI

// TransferSwapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TransferSwapMetaData.Bin instead.
var TransferSwapBin = TransferSwapMetaData.Bin

// DeployTransferSwap deploys a new Ethereum contract, binding an instance of TransferSwap to it.
func DeployTransferSwap(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address, _supportedDex common.Address, _nativeWrap common.Address) (common.Address, *types.Transaction, *TransferSwap, error) {
	parsed, err := TransferSwapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TransferSwapBin), backend, _messageBus, _supportedDex, _nativeWrap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TransferSwap{TransferSwapCaller: TransferSwapCaller{contract: contract}, TransferSwapTransactor: TransferSwapTransactor{contract: contract}, TransferSwapFilterer: TransferSwapFilterer{contract: contract}}, nil
}

// TransferSwap is an auto generated Go binding around an Ethereum contract.
type TransferSwap struct {
	TransferSwapCaller     // Read-only binding to the contract
	TransferSwapTransactor // Write-only binding to the contract
	TransferSwapFilterer   // Log filterer for contract events
}

// TransferSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferSwapSession struct {
	Contract     *TransferSwap     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferSwapCallerSession struct {
	Contract *TransferSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TransferSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferSwapTransactorSession struct {
	Contract     *TransferSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TransferSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferSwapRaw struct {
	Contract *TransferSwap // Generic contract binding to access the raw methods on
}

// TransferSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferSwapCallerRaw struct {
	Contract *TransferSwapCaller // Generic read-only contract binding to access the raw methods on
}

// TransferSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferSwapTransactorRaw struct {
	Contract *TransferSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferSwap creates a new instance of TransferSwap, bound to a specific deployed contract.
func NewTransferSwap(address common.Address, backend bind.ContractBackend) (*TransferSwap, error) {
	contract, err := bindTransferSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferSwap{TransferSwapCaller: TransferSwapCaller{contract: contract}, TransferSwapTransactor: TransferSwapTransactor{contract: contract}, TransferSwapFilterer: TransferSwapFilterer{contract: contract}}, nil
}

// NewTransferSwapCaller creates a new read-only instance of TransferSwap, bound to a specific deployed contract.
func NewTransferSwapCaller(address common.Address, caller bind.ContractCaller) (*TransferSwapCaller, error) {
	contract, err := bindTransferSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferSwapCaller{contract: contract}, nil
}

// NewTransferSwapTransactor creates a new write-only instance of TransferSwap, bound to a specific deployed contract.
func NewTransferSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferSwapTransactor, error) {
	contract, err := bindTransferSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferSwapTransactor{contract: contract}, nil
}

// NewTransferSwapFilterer creates a new log filterer instance of TransferSwap, bound to a specific deployed contract.
func NewTransferSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferSwapFilterer, error) {
	contract, err := bindTransferSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferSwapFilterer{contract: contract}, nil
}

// bindTransferSwap binds a generic wrapper to an already deployed contract.
func bindTransferSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferSwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferSwap *TransferSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferSwap.Contract.TransferSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferSwap *TransferSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferSwap *TransferSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferSwap *TransferSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferSwap *TransferSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferSwap *TransferSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferSwap.Contract.contract.Transact(opts, method, params...)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TransferSwap *TransferSwapCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferSwap.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TransferSwap *TransferSwapSession) MessageBus() (common.Address, error) {
	return _TransferSwap.Contract.MessageBus(&_TransferSwap.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_TransferSwap *TransferSwapCallerSession) MessageBus() (common.Address, error) {
	return _TransferSwap.Contract.MessageBus(&_TransferSwap.CallOpts)
}

// MinSwapAmounts is a free data retrieval call binding the contract method 0x34e18125.
//
// Solidity: function minSwapAmounts(address ) view returns(uint256)
func (_TransferSwap *TransferSwapCaller) MinSwapAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TransferSwap.contract.Call(opts, &out, "minSwapAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSwapAmounts is a free data retrieval call binding the contract method 0x34e18125.
//
// Solidity: function minSwapAmounts(address ) view returns(uint256)
func (_TransferSwap *TransferSwapSession) MinSwapAmounts(arg0 common.Address) (*big.Int, error) {
	return _TransferSwap.Contract.MinSwapAmounts(&_TransferSwap.CallOpts, arg0)
}

// MinSwapAmounts is a free data retrieval call binding the contract method 0x34e18125.
//
// Solidity: function minSwapAmounts(address ) view returns(uint256)
func (_TransferSwap *TransferSwapCallerSession) MinSwapAmounts(arg0 common.Address) (*big.Int, error) {
	return _TransferSwap.Contract.MinSwapAmounts(&_TransferSwap.CallOpts, arg0)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_TransferSwap *TransferSwapCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferSwap.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_TransferSwap *TransferSwapSession) NativeWrap() (common.Address, error) {
	return _TransferSwap.Contract.NativeWrap(&_TransferSwap.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_TransferSwap *TransferSwapCallerSession) NativeWrap() (common.Address, error) {
	return _TransferSwap.Contract.NativeWrap(&_TransferSwap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferSwap *TransferSwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferSwap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferSwap *TransferSwapSession) Owner() (common.Address, error) {
	return _TransferSwap.Contract.Owner(&_TransferSwap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferSwap *TransferSwapCallerSession) Owner() (common.Address, error) {
	return _TransferSwap.Contract.Owner(&_TransferSwap.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessage(&_TransferSwap.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactorSession) ExecuteMessage(_sender common.Address, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessage(&_TransferSwap.TransactOpts, _sender, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address , address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, arg0 common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "executeMessageWithTransfer", arg0, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address , address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapSession) ExecuteMessageWithTransfer(arg0 common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessageWithTransfer(&_TransferSwap.TransactOpts, arg0, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xce35dd9a.
//
// Solidity: function executeMessageWithTransfer(address , address _token, uint256 _amount, uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactorSession) ExecuteMessageWithTransfer(arg0 common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessageWithTransfer(&_TransferSwap.TransactOpts, arg0, _token, _amount, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address , address , uint256 , uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "executeMessageWithTransferFallback", arg0, arg1, arg2, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address , address , uint256 , uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapSession) ExecuteMessageWithTransferFallback(arg0 common.Address, arg1 common.Address, arg2 *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessageWithTransferFallback(&_TransferSwap.TransactOpts, arg0, arg1, arg2, _srcChainId, _message)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0xf00f39ce.
//
// Solidity: function executeMessageWithTransferFallback(address , address , uint256 , uint64 _srcChainId, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactorSession) ExecuteMessageWithTransferFallback(arg0 common.Address, arg1 common.Address, arg2 *big.Int, _srcChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessageWithTransferFallback(&_TransferSwap.TransactOpts, arg0, arg1, arg2, _srcChainId, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessageWithTransferRefund(&_TransferSwap.TransactOpts, _token, _amount, _message)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x20be95f2.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message) payable returns(bool)
func (_TransferSwap *TransferSwapTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte) (*types.Transaction, error) {
	return _TransferSwap.Contract.ExecuteMessageWithTransferRefund(&_TransferSwap.TransactOpts, _token, _amount, _message)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TransferSwap *TransferSwapTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TransferSwap *TransferSwapSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetMessageBus(&_TransferSwap.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_TransferSwap *TransferSwapTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetMessageBus(&_TransferSwap.TransactOpts, _messageBus)
}

// SetMinSwapAmount is a paid mutator transaction binding the contract method 0x496dbfd9.
//
// Solidity: function setMinSwapAmount(address _token, uint256 _minSwapAmount) returns()
func (_TransferSwap *TransferSwapTransactor) SetMinSwapAmount(opts *bind.TransactOpts, _token common.Address, _minSwapAmount *big.Int) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "setMinSwapAmount", _token, _minSwapAmount)
}

// SetMinSwapAmount is a paid mutator transaction binding the contract method 0x496dbfd9.
//
// Solidity: function setMinSwapAmount(address _token, uint256 _minSwapAmount) returns()
func (_TransferSwap *TransferSwapSession) SetMinSwapAmount(_token common.Address, _minSwapAmount *big.Int) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetMinSwapAmount(&_TransferSwap.TransactOpts, _token, _minSwapAmount)
}

// SetMinSwapAmount is a paid mutator transaction binding the contract method 0x496dbfd9.
//
// Solidity: function setMinSwapAmount(address _token, uint256 _minSwapAmount) returns()
func (_TransferSwap *TransferSwapTransactorSession) SetMinSwapAmount(_token common.Address, _minSwapAmount *big.Int) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetMinSwapAmount(&_TransferSwap.TransactOpts, _token, _minSwapAmount)
}

// SetNativeWrap is a paid mutator transaction binding the contract method 0x5b5a66a7.
//
// Solidity: function setNativeWrap(address _nativeWrap) returns()
func (_TransferSwap *TransferSwapTransactor) SetNativeWrap(opts *bind.TransactOpts, _nativeWrap common.Address) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "setNativeWrap", _nativeWrap)
}

// SetNativeWrap is a paid mutator transaction binding the contract method 0x5b5a66a7.
//
// Solidity: function setNativeWrap(address _nativeWrap) returns()
func (_TransferSwap *TransferSwapSession) SetNativeWrap(_nativeWrap common.Address) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetNativeWrap(&_TransferSwap.TransactOpts, _nativeWrap)
}

// SetNativeWrap is a paid mutator transaction binding the contract method 0x5b5a66a7.
//
// Solidity: function setNativeWrap(address _nativeWrap) returns()
func (_TransferSwap *TransferSwapTransactorSession) SetNativeWrap(_nativeWrap common.Address) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetNativeWrap(&_TransferSwap.TransactOpts, _nativeWrap)
}

// SetSupportedDex is a paid mutator transaction binding the contract method 0xd9b796de.
//
// Solidity: function setSupportedDex(address _dex, bool _enabled) returns()
func (_TransferSwap *TransferSwapTransactor) SetSupportedDex(opts *bind.TransactOpts, _dex common.Address, _enabled bool) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "setSupportedDex", _dex, _enabled)
}

// SetSupportedDex is a paid mutator transaction binding the contract method 0xd9b796de.
//
// Solidity: function setSupportedDex(address _dex, bool _enabled) returns()
func (_TransferSwap *TransferSwapSession) SetSupportedDex(_dex common.Address, _enabled bool) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetSupportedDex(&_TransferSwap.TransactOpts, _dex, _enabled)
}

// SetSupportedDex is a paid mutator transaction binding the contract method 0xd9b796de.
//
// Solidity: function setSupportedDex(address _dex, bool _enabled) returns()
func (_TransferSwap *TransferSwapTransactorSession) SetSupportedDex(_dex common.Address, _enabled bool) (*types.Transaction, error) {
	return _TransferSwap.Contract.SetSupportedDex(&_TransferSwap.TransactOpts, _dex, _enabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferSwap *TransferSwapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferSwap *TransferSwapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferOwnership(&_TransferSwap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferSwap *TransferSwapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferOwnership(&_TransferSwap.TransactOpts, newOwner)
}

// TransferWithSwap is a paid mutator transaction binding the contract method 0xa8a7e29c.
//
// Solidity: function transferWithSwap(address _receiver, uint256 _amountIn, uint64 _dstChainId, (address[],address,uint256,uint256) _srcSwap, (address[],address,uint256,uint256) _dstSwap, uint32 _maxBridgeSlippage, uint64 _nonce) payable returns()
func (_TransferSwap *TransferSwapTransactor) TransferWithSwap(opts *bind.TransactOpts, _receiver common.Address, _amountIn *big.Int, _dstChainId uint64, _srcSwap TransferSwapSwapInfo, _dstSwap TransferSwapSwapInfo, _maxBridgeSlippage uint32, _nonce uint64) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "transferWithSwap", _receiver, _amountIn, _dstChainId, _srcSwap, _dstSwap, _maxBridgeSlippage, _nonce)
}

// TransferWithSwap is a paid mutator transaction binding the contract method 0xa8a7e29c.
//
// Solidity: function transferWithSwap(address _receiver, uint256 _amountIn, uint64 _dstChainId, (address[],address,uint256,uint256) _srcSwap, (address[],address,uint256,uint256) _dstSwap, uint32 _maxBridgeSlippage, uint64 _nonce) payable returns()
func (_TransferSwap *TransferSwapSession) TransferWithSwap(_receiver common.Address, _amountIn *big.Int, _dstChainId uint64, _srcSwap TransferSwapSwapInfo, _dstSwap TransferSwapSwapInfo, _maxBridgeSlippage uint32, _nonce uint64) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferWithSwap(&_TransferSwap.TransactOpts, _receiver, _amountIn, _dstChainId, _srcSwap, _dstSwap, _maxBridgeSlippage, _nonce)
}

// TransferWithSwap is a paid mutator transaction binding the contract method 0xa8a7e29c.
//
// Solidity: function transferWithSwap(address _receiver, uint256 _amountIn, uint64 _dstChainId, (address[],address,uint256,uint256) _srcSwap, (address[],address,uint256,uint256) _dstSwap, uint32 _maxBridgeSlippage, uint64 _nonce) payable returns()
func (_TransferSwap *TransferSwapTransactorSession) TransferWithSwap(_receiver common.Address, _amountIn *big.Int, _dstChainId uint64, _srcSwap TransferSwapSwapInfo, _dstSwap TransferSwapSwapInfo, _maxBridgeSlippage uint32, _nonce uint64) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferWithSwap(&_TransferSwap.TransactOpts, _receiver, _amountIn, _dstChainId, _srcSwap, _dstSwap, _maxBridgeSlippage, _nonce)
}

// TransferWithSwapNative is a paid mutator transaction binding the contract method 0x5fdb3d8d.
//
// Solidity: function transferWithSwapNative(address _receiver, uint256 _amountIn, uint64 _dstChainId, (address[],address,uint256,uint256) _srcSwap, (address[],address,uint256,uint256) _dstSwap, uint32 _maxBridgeSlippage, uint64 _nonce, bool _nativeOut) payable returns()
func (_TransferSwap *TransferSwapTransactor) TransferWithSwapNative(opts *bind.TransactOpts, _receiver common.Address, _amountIn *big.Int, _dstChainId uint64, _srcSwap TransferSwapSwapInfo, _dstSwap TransferSwapSwapInfo, _maxBridgeSlippage uint32, _nonce uint64, _nativeOut bool) (*types.Transaction, error) {
	return _TransferSwap.contract.Transact(opts, "transferWithSwapNative", _receiver, _amountIn, _dstChainId, _srcSwap, _dstSwap, _maxBridgeSlippage, _nonce, _nativeOut)
}

// TransferWithSwapNative is a paid mutator transaction binding the contract method 0x5fdb3d8d.
//
// Solidity: function transferWithSwapNative(address _receiver, uint256 _amountIn, uint64 _dstChainId, (address[],address,uint256,uint256) _srcSwap, (address[],address,uint256,uint256) _dstSwap, uint32 _maxBridgeSlippage, uint64 _nonce, bool _nativeOut) payable returns()
func (_TransferSwap *TransferSwapSession) TransferWithSwapNative(_receiver common.Address, _amountIn *big.Int, _dstChainId uint64, _srcSwap TransferSwapSwapInfo, _dstSwap TransferSwapSwapInfo, _maxBridgeSlippage uint32, _nonce uint64, _nativeOut bool) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferWithSwapNative(&_TransferSwap.TransactOpts, _receiver, _amountIn, _dstChainId, _srcSwap, _dstSwap, _maxBridgeSlippage, _nonce, _nativeOut)
}

// TransferWithSwapNative is a paid mutator transaction binding the contract method 0x5fdb3d8d.
//
// Solidity: function transferWithSwapNative(address _receiver, uint256 _amountIn, uint64 _dstChainId, (address[],address,uint256,uint256) _srcSwap, (address[],address,uint256,uint256) _dstSwap, uint32 _maxBridgeSlippage, uint64 _nonce, bool _nativeOut) payable returns()
func (_TransferSwap *TransferSwapTransactorSession) TransferWithSwapNative(_receiver common.Address, _amountIn *big.Int, _dstChainId uint64, _srcSwap TransferSwapSwapInfo, _dstSwap TransferSwapSwapInfo, _maxBridgeSlippage uint32, _nonce uint64, _nativeOut bool) (*types.Transaction, error) {
	return _TransferSwap.Contract.TransferWithSwapNative(&_TransferSwap.TransactOpts, _receiver, _amountIn, _dstChainId, _srcSwap, _dstSwap, _maxBridgeSlippage, _nonce, _nativeOut)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransferSwap *TransferSwapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferSwap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransferSwap *TransferSwapSession) Receive() (*types.Transaction, error) {
	return _TransferSwap.Contract.Receive(&_TransferSwap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TransferSwap *TransferSwapTransactorSession) Receive() (*types.Transaction, error) {
	return _TransferSwap.Contract.Receive(&_TransferSwap.TransactOpts)
}

// TransferSwapDirectSwapIterator is returned from FilterDirectSwap and is used to iterate over the raw logs and unpacked data for DirectSwap events raised by the TransferSwap contract.
type TransferSwapDirectSwapIterator struct {
	Event *TransferSwapDirectSwap // Event containing the contract specifics and raw log

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
func (it *TransferSwapDirectSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferSwapDirectSwap)
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
		it.Event = new(TransferSwapDirectSwap)
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
func (it *TransferSwapDirectSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferSwapDirectSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferSwapDirectSwap represents a DirectSwap event raised by the TransferSwap contract.
type TransferSwapDirectSwap struct {
	Id         [32]byte
	SrcChainId uint64
	AmountIn   *big.Int
	TokenIn    common.Address
	AmountOut  *big.Int
	TokenOut   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDirectSwap is a free log retrieval operation binding the contract event 0xdbd876e57420a75850bd1ce36de349d111ab4ff35e2c2379af8acc7503769174.
//
// Solidity: event DirectSwap(bytes32 id, uint64 srcChainId, uint256 amountIn, address tokenIn, uint256 amountOut, address tokenOut)
func (_TransferSwap *TransferSwapFilterer) FilterDirectSwap(opts *bind.FilterOpts) (*TransferSwapDirectSwapIterator, error) {

	logs, sub, err := _TransferSwap.contract.FilterLogs(opts, "DirectSwap")
	if err != nil {
		return nil, err
	}
	return &TransferSwapDirectSwapIterator{contract: _TransferSwap.contract, event: "DirectSwap", logs: logs, sub: sub}, nil
}

// WatchDirectSwap is a free log subscription operation binding the contract event 0xdbd876e57420a75850bd1ce36de349d111ab4ff35e2c2379af8acc7503769174.
//
// Solidity: event DirectSwap(bytes32 id, uint64 srcChainId, uint256 amountIn, address tokenIn, uint256 amountOut, address tokenOut)
func (_TransferSwap *TransferSwapFilterer) WatchDirectSwap(opts *bind.WatchOpts, sink chan<- *TransferSwapDirectSwap) (event.Subscription, error) {

	logs, sub, err := _TransferSwap.contract.WatchLogs(opts, "DirectSwap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferSwapDirectSwap)
				if err := _TransferSwap.contract.UnpackLog(event, "DirectSwap", log); err != nil {
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

// ParseDirectSwap is a log parse operation binding the contract event 0xdbd876e57420a75850bd1ce36de349d111ab4ff35e2c2379af8acc7503769174.
//
// Solidity: event DirectSwap(bytes32 id, uint64 srcChainId, uint256 amountIn, address tokenIn, uint256 amountOut, address tokenOut)
func (_TransferSwap *TransferSwapFilterer) ParseDirectSwap(log types.Log) (*TransferSwapDirectSwap, error) {
	event := new(TransferSwapDirectSwap)
	if err := _TransferSwap.contract.UnpackLog(event, "DirectSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferSwapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TransferSwap contract.
type TransferSwapOwnershipTransferredIterator struct {
	Event *TransferSwapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TransferSwapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferSwapOwnershipTransferred)
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
		it.Event = new(TransferSwapOwnershipTransferred)
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
func (it *TransferSwapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferSwapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferSwapOwnershipTransferred represents a OwnershipTransferred event raised by the TransferSwap contract.
type TransferSwapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferSwap *TransferSwapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TransferSwapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferSwap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TransferSwapOwnershipTransferredIterator{contract: _TransferSwap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferSwap *TransferSwapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TransferSwapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferSwap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferSwapOwnershipTransferred)
				if err := _TransferSwap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TransferSwap *TransferSwapFilterer) ParseOwnershipTransferred(log types.Log) (*TransferSwapOwnershipTransferred, error) {
	event := new(TransferSwapOwnershipTransferred)
	if err := _TransferSwap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferSwapSwapRequestDoneIterator is returned from FilterSwapRequestDone and is used to iterate over the raw logs and unpacked data for SwapRequestDone events raised by the TransferSwap contract.
type TransferSwapSwapRequestDoneIterator struct {
	Event *TransferSwapSwapRequestDone // Event containing the contract specifics and raw log

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
func (it *TransferSwapSwapRequestDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferSwapSwapRequestDone)
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
		it.Event = new(TransferSwapSwapRequestDone)
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
func (it *TransferSwapSwapRequestDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferSwapSwapRequestDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferSwapSwapRequestDone represents a SwapRequestDone event raised by the TransferSwap contract.
type TransferSwapSwapRequestDone struct {
	Id        [32]byte
	DstAmount *big.Int
	Status    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapRequestDone is a free log retrieval operation binding the contract event 0xccbb695db2dfb79e3f69bcf724d11b850c79b0ae358e62d60ed94044f38a2b9c.
//
// Solidity: event SwapRequestDone(bytes32 id, uint256 dstAmount, uint8 status)
func (_TransferSwap *TransferSwapFilterer) FilterSwapRequestDone(opts *bind.FilterOpts) (*TransferSwapSwapRequestDoneIterator, error) {

	logs, sub, err := _TransferSwap.contract.FilterLogs(opts, "SwapRequestDone")
	if err != nil {
		return nil, err
	}
	return &TransferSwapSwapRequestDoneIterator{contract: _TransferSwap.contract, event: "SwapRequestDone", logs: logs, sub: sub}, nil
}

// WatchSwapRequestDone is a free log subscription operation binding the contract event 0xccbb695db2dfb79e3f69bcf724d11b850c79b0ae358e62d60ed94044f38a2b9c.
//
// Solidity: event SwapRequestDone(bytes32 id, uint256 dstAmount, uint8 status)
func (_TransferSwap *TransferSwapFilterer) WatchSwapRequestDone(opts *bind.WatchOpts, sink chan<- *TransferSwapSwapRequestDone) (event.Subscription, error) {

	logs, sub, err := _TransferSwap.contract.WatchLogs(opts, "SwapRequestDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferSwapSwapRequestDone)
				if err := _TransferSwap.contract.UnpackLog(event, "SwapRequestDone", log); err != nil {
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

// ParseSwapRequestDone is a log parse operation binding the contract event 0xccbb695db2dfb79e3f69bcf724d11b850c79b0ae358e62d60ed94044f38a2b9c.
//
// Solidity: event SwapRequestDone(bytes32 id, uint256 dstAmount, uint8 status)
func (_TransferSwap *TransferSwapFilterer) ParseSwapRequestDone(log types.Log) (*TransferSwapSwapRequestDone, error) {
	event := new(TransferSwapSwapRequestDone)
	if err := _TransferSwap.contract.UnpackLog(event, "SwapRequestDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TransferSwapSwapRequestSentIterator is returned from FilterSwapRequestSent and is used to iterate over the raw logs and unpacked data for SwapRequestSent events raised by the TransferSwap contract.
type TransferSwapSwapRequestSentIterator struct {
	Event *TransferSwapSwapRequestSent // Event containing the contract specifics and raw log

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
func (it *TransferSwapSwapRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferSwapSwapRequestSent)
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
		it.Event = new(TransferSwapSwapRequestSent)
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
func (it *TransferSwapSwapRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferSwapSwapRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferSwapSwapRequestSent represents a SwapRequestSent event raised by the TransferSwap contract.
type TransferSwapSwapRequestSent struct {
	Id         [32]byte
	DstChainId uint64
	SrcAmount  *big.Int
	SrcToken   common.Address
	DstToken   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapRequestSent is a free log retrieval operation binding the contract event 0x600b63623f19b18f8f0e7825ec78b3d9779d8124c76ad75b58aac7167f04f8d2.
//
// Solidity: event SwapRequestSent(bytes32 id, uint64 dstChainId, uint256 srcAmount, address srcToken, address dstToken)
func (_TransferSwap *TransferSwapFilterer) FilterSwapRequestSent(opts *bind.FilterOpts) (*TransferSwapSwapRequestSentIterator, error) {

	logs, sub, err := _TransferSwap.contract.FilterLogs(opts, "SwapRequestSent")
	if err != nil {
		return nil, err
	}
	return &TransferSwapSwapRequestSentIterator{contract: _TransferSwap.contract, event: "SwapRequestSent", logs: logs, sub: sub}, nil
}

// WatchSwapRequestSent is a free log subscription operation binding the contract event 0x600b63623f19b18f8f0e7825ec78b3d9779d8124c76ad75b58aac7167f04f8d2.
//
// Solidity: event SwapRequestSent(bytes32 id, uint64 dstChainId, uint256 srcAmount, address srcToken, address dstToken)
func (_TransferSwap *TransferSwapFilterer) WatchSwapRequestSent(opts *bind.WatchOpts, sink chan<- *TransferSwapSwapRequestSent) (event.Subscription, error) {

	logs, sub, err := _TransferSwap.contract.WatchLogs(opts, "SwapRequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferSwapSwapRequestSent)
				if err := _TransferSwap.contract.UnpackLog(event, "SwapRequestSent", log); err != nil {
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

// ParseSwapRequestSent is a log parse operation binding the contract event 0x600b63623f19b18f8f0e7825ec78b3d9779d8124c76ad75b58aac7167f04f8d2.
//
// Solidity: event SwapRequestSent(bytes32 id, uint64 dstChainId, uint256 srcAmount, address srcToken, address dstToken)
func (_TransferSwap *TransferSwapFilterer) ParseSwapRequestSent(log types.Log) (*TransferSwapSwapRequestSent, error) {
	event := new(TransferSwapSwapRequestSent)
	if err := _TransferSwap.contract.UnpackLog(event, "SwapRequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
