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

// DataTypesDelegatorInfo is an auto generated low-level Go binding around an user-defined struct.
type DataTypesDelegatorInfo struct {
	ValAddr                        common.Address
	Tokens                         *big.Int
	Shares                         *big.Int
	Undelegations                  []DataTypesUndelegation
	UndelegationTokens             *big.Int
	WithdrawableUndelegationTokens *big.Int
}

// DataTypesUndelegation is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUndelegation struct {
	Shares        *big.Int
	CreationBlock *big.Int
}

// DataTypesValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type DataTypesValidatorInfo struct {
	ValAddr           common.Address
	Status            uint8
	Signer            common.Address
	Tokens            *big.Int
	Shares            *big.Int
	MinSelfDelegation *big.Int
	CommissionRate    uint64
}

// DataTypesValidatorTokens is an auto generated low-level Go binding around an user-defined struct.
type DataTypesValidatorTokens struct {
	ValAddr common.Address
	Tokens  *big.Int
}

// DataTypesMetaData contains all meta data concerning the DataTypes contract.
var DataTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122020de24d75dd60db0d54cc9739bf5e58a2c6562341e8ea4ae9a99dcc5c9a000ed64736f6c63430008090033",
}

// DataTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use DataTypesMetaData.ABI instead.
var DataTypesABI = DataTypesMetaData.ABI

// DataTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataTypesMetaData.Bin instead.
var DataTypesBin = DataTypesMetaData.Bin

// DeployDataTypes deploys a new Ethereum contract, binding an instance of DataTypes to it.
func DeployDataTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataTypes, error) {
	parsed, err := DataTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataTypes{DataTypesCaller: DataTypesCaller{contract: contract}, DataTypesTransactor: DataTypesTransactor{contract: contract}, DataTypesFilterer: DataTypesFilterer{contract: contract}}, nil
}

// DataTypes is an auto generated Go binding around an Ethereum contract.
type DataTypes struct {
	DataTypesCaller     // Read-only binding to the contract
	DataTypesTransactor // Write-only binding to the contract
	DataTypesFilterer   // Log filterer for contract events
}

// DataTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataTypesSession struct {
	Contract     *DataTypes        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataTypesCallerSession struct {
	Contract *DataTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DataTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTypesTransactorSession struct {
	Contract     *DataTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DataTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataTypesRaw struct {
	Contract *DataTypes // Generic contract binding to access the raw methods on
}

// DataTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataTypesCallerRaw struct {
	Contract *DataTypesCaller // Generic read-only contract binding to access the raw methods on
}

// DataTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTypesTransactorRaw struct {
	Contract *DataTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataTypes creates a new instance of DataTypes, bound to a specific deployed contract.
func NewDataTypes(address common.Address, backend bind.ContractBackend) (*DataTypes, error) {
	contract, err := bindDataTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataTypes{DataTypesCaller: DataTypesCaller{contract: contract}, DataTypesTransactor: DataTypesTransactor{contract: contract}, DataTypesFilterer: DataTypesFilterer{contract: contract}}, nil
}

// NewDataTypesCaller creates a new read-only instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesCaller(address common.Address, caller bind.ContractCaller) (*DataTypesCaller, error) {
	contract, err := bindDataTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypesCaller{contract: contract}, nil
}

// NewDataTypesTransactor creates a new write-only instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTypesTransactor, error) {
	contract, err := bindDataTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypesTransactor{contract: contract}, nil
}

// NewDataTypesFilterer creates a new log filterer instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*DataTypesFilterer, error) {
	contract, err := bindDataTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataTypesFilterer{contract: contract}, nil
}

// bindDataTypes binds a generic wrapper to an already deployed contract.
func bindDataTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypes *DataTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTypes.Contract.DataTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypes *DataTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTypes.Contract.DataTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypes *DataTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTypes.Contract.DataTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypes *DataTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypes *DataTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypes *DataTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTypes.Contract.contract.Transact(opts, method, params...)
}

// FarmingRewardsMetaData contains all meta data concerning the FarmingRewards contract.
var FarmingRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"FarmingRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"}],\"name\":\"FarmingRewardContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardsRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001d4138038062001d41833981016040819052620000349162000182565b6200003f3362000069565b6000805460ff60a01b191690556200005733620000b9565b6001600160a01b0316608052620001b4565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff1615620001275760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640160405180910390fd5b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910160405180910390a150565b6000602082840312156200019557600080fd5b81516001600160a01b0381168114620001ad57600080fd5b9392505050565b608051611b6a620001d76000396000818161025f01526103cc0152611b6a6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806380f51c12116100975780638da5cb5b116100665780638da5cb5b146102225780639d4323be14610247578063ccf2683b1461025a578063f2fde38b1461028157600080fd5b806380f51c12146101d1578063825168ff146101f457806382dc1ec4146102075780638456cb591461021a57600080fd5b80636b2c0f55116100d35780636b2c0f551461019b5780636b5d21e9146101ae5780636ef8d66d146101c1578063715018a6146101c957600080fd5b80631744092e146101055780633f4ba83a1461014357806346fbf68e1461014d5780635c975abb14610189575b600080fd5b610130610113366004611660565b600260209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b61014b610294565b005b61017961015b366004611693565b6001600160a01b031660009081526001602052604090205460ff1690565b604051901515815260200161013a565b600054600160a01b900460ff16610179565b61014b6101a9366004611693565b610302565b61014b6101bc3660046116fa565b610368565b61014b61065a565b61014b610663565b6101796101df366004611693565b60016020526000908152604090205460ff1681565b61014b6102023660046117e9565b6106c7565b61014b610215366004611693565b61077c565b61014b6107df565b6000546001600160a01b03165b6040516001600160a01b03909116815260200161013a565b61014b6102553660046117e9565b610846565b61022f7f000000000000000000000000000000000000000000000000000000000000000081565b61014b61028f366004611693565b610911565b3360009081526001602052604090205460ff166102f85760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064015b60405180910390fd5b6103006109f0565b565b6000546001600160a01b0316331461035c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102ef565b61036581610a96565b50565b600054600160a01b900460ff16156103b55760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102ef565b604051633416de1160e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063682dbc229061040f908b908b908b908b908b908b908b908b906004016118d2565b60006040518083038186803b15801561042757600080fd5b505afa15801561043b573d6000803e3d6000fd5b50505050600061048089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b5692505050565b9050468160200151146104d55760405162461bcd60e51b815260206004820152601160248201527f436861696e204944206d69736d6174636800000000000000000000000000000060448201526064016102ef565b6000805b826040015151811015610600576000836040015182815181106104fe576104fe6119b4565b60200260200101519050600084606001518381518110610520576105206119b4565b60209081029190910181015186516001600160a01b03908116600090815260028452604080822092871682529190935282205490925061056090836119e0565b905080156105ea5785516001600160a01b0390811660009081526002602090815260408083209387168084529390915290208390558651600196506105a6919083610e38565b85516040518281526001600160a01b038086169216907f97e6c3172350795e26977663112f38653689372e771e85bad9fbadb1af0e98b29060200160405180910390a35b50505080806105f8906119f7565b9150506104d9565b508061064e5760405162461bcd60e51b815260206004820152600d60248201527f4e6f206e6577207265776172640000000000000000000000000000000000000060448201526064016102ef565b50505050505050505050565b61030033610a96565b6000546001600160a01b031633146106bd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102ef565b6103006000610ecd565b600054600160a01b900460ff16156107145760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102ef565b3361072a6001600160a01b038416823085610f35565b826001600160a01b0316816001600160a01b03167f40aa1b9a9157bc37a09a78d5a46e53087b82ee0034ebe896d4d1a52f31b333d48460405161076f91815260200190565b60405180910390a3505050565b6000546001600160a01b031633146107d65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102ef565b61036581610f73565b3360009081526001602052604090205460ff1661083e5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016102ef565b610300611031565b600054600160a01b900460ff1661089f5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102ef565b6000546001600160a01b031633146108f95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102ef565b61090d6001600160a01b0383163383610e38565b5050565b6000546001600160a01b0316331461096b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102ef565b6001600160a01b0381166109e75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102ef565b61036581610ecd565b600054600160a01b900460ff16610a495760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102ef565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b03811660009081526001602052604090205460ff16610afe5760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f7420706175736572000000000000000000000060448201526064016102ef565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e91015b60405180910390a150565b610b8a604051806080016040528060006001600160a01b031681526020016000815260200160608152602001606081525090565b60408051808201909152600080825260208201849052610bab8260046110b9565b905080600381518110610bc057610bc06119b4565b602002602001015167ffffffffffffffff811115610be057610be0611a12565b604051908082528060200260200182016040528015610c09578160200160208202803683370190505b508360400181905250600081600381518110610c2757610c276119b4565b60200260200101818152505080600481518110610c4657610c466119b4565b602002602001015167ffffffffffffffff811115610c6657610c66611a12565b604051908082528060200260200182016040528015610c8f578160200160208202803683370190505b508360600181905250600081600481518110610cad57610cad6119b4565b6020026020010181815250506000805b60208401515184511015610e2f57610cd484611173565b90925090508160011415610d0357610cf3610cee856111ad565b61126a565b6001600160a01b03168552610cbd565b8160021415610d2757610d1d610d18856111ad565b61127b565b6020860152610cbd565b8160031415610dbc57610d3c610cee856111ad565b856040015184600381518110610d5457610d546119b4565b602002602001015181518110610d6c57610d6c6119b4565b60200260200101906001600160a01b031690816001600160a01b03168152505082600381518110610d9f57610d9f6119b4565b602002602001018051809190610db4906119f7565b905250610cbd565b8160041415610e2057610dd1610d18856111ad565b856060015184600481518110610de957610de96119b4565b602002602001015181518110610e0157610e016119b4565b60200260200101818152505082600481518110610d9f57610d9f6119b4565b610e2a84826112b2565b610cbd565b50505050919050565b6040516001600160a01b038316602482015260448101829052610ec890849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611324565b505050565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040516001600160a01b0380851660248301528316604482015260648101829052610f6d9085906323b872dd60e01b90608401610e64565b50505050565b6001600160a01b03811660009081526001602052604090205460ff1615610fdc5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064016102ef565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101610b4b565b600054600160a01b900460ff161561107e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102ef565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610a793390565b81516060906110c9836001611a28565b67ffffffffffffffff8111156110e1576110e1611a12565b60405190808252806020026020018201604052801561110a578160200160208202803683370190505b5091506000805b6020860151518651101561116a5761112886611173565b80925081935050506001848381518110611144576111446119b4565b602002602001018181516111589190611a28565b90525061116586826112b2565b611111565b50509092525090565b600080600061118184611409565b905061118e600882611a40565b92508060071660058111156111a5576111a5611a62565b915050915091565b606060006111ba83611409565b905060008184600001516111ce9190611a28565b90508360200151518111156111e257600080fd5b8167ffffffffffffffff8111156111fb576111fb611a12565b6040519080825280601f01601f191660200182016040528015611225576020820181803683370190505b50602080860151865192955091818601919083010160005b8581101561125f578181015183820152611258602082611a28565b905061123d565b505050935250919050565b60006112758261148b565b92915050565b600060208251111561128c57600080fd5b60208201519050815160206112a191906119e0565b6112ac906008611a78565b1c919050565b60008160058111156112c6576112c6611a62565b14156112d557610ec882611409565b60028160058111156112e9576112e9611a62565b14156101005760006112fa83611409565b9050808360000181815161130e9190611a28565b90525060208301515183511115610ec857600080fd5b6000611379826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166114b39092919063ffffffff16565b805190915015610ec857808060200190518101906113979190611a97565b610ec85760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102ef565b602080820151825181019091015160009182805b600a8110156114855783811a9150611436816007611a78565b82607f16901b85179450816080166000141561147357611457816001611a28565b86518790611466908390611a28565b9052509395945050505050565b8061147d816119f7565b91505061141d565b50600080fd5b6000815160141461149b57600080fd5b50602001516c01000000000000000000000000900490565b60606114c284846000856114cc565b90505b9392505050565b6060824710156115445760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102ef565b843b6115925760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102ef565b600080866001600160a01b031685876040516115ae9190611ae5565b60006040518083038185875af1925050503d80600081146115eb576040519150601f19603f3d011682016040523d82523d6000602084013e6115f0565b606091505b509150915061160082828661160b565b979650505050505050565b6060831561161a5750816114c5565b82511561162a5782518084602001fd5b8160405162461bcd60e51b81526004016102ef9190611b01565b80356001600160a01b038116811461165b57600080fd5b919050565b6000806040838503121561167357600080fd5b61167c83611644565b915061168a60208401611644565b90509250929050565b6000602082840312156116a557600080fd5b6114c582611644565b60008083601f8401126116c057600080fd5b50813567ffffffffffffffff8111156116d857600080fd5b6020830191508360208260051b85010111156116f357600080fd5b9250929050565b6000806000806000806000806080898b03121561171657600080fd5b883567ffffffffffffffff8082111561172e57600080fd5b818b0191508b601f83011261174257600080fd5b81358181111561175157600080fd5b8c602082850101111561176357600080fd5b60209283019a509850908a0135908082111561177e57600080fd5b61178a8c838d016116ae565b909850965060408b01359150808211156117a357600080fd5b6117af8c838d016116ae565b909650945060608b01359150808211156117c857600080fd5b506117d58b828c016116ae565b999c989b5096995094979396929594505050565b600080604083850312156117fc57600080fd5b61180583611644565b946020939093013593505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8183526000602080850194508260005b85811015611878576001600160a01b0361186583611644565b168752958201959082019060010161184c565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156118b557600080fd5b8260051b8083602087013760009401602001938452509192915050565b6080815260006118e6608083018a8c611813565b82810360208401528088825260208201905060208960051b8301018a60005b8b81101561197957848303601f190184528135368e9003601e1901811261192b57600080fd5b8d01803567ffffffffffffffff81111561194457600080fd5b8036038f131561195357600080fd5b611961858260208501611813565b60209687019690955093909301925050600101611905565b5050848103604086015261198e81898b61183c565b9250505082810360608401526119a5818587611883565b9b9a5050505050505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000828210156119f2576119f26119ca565b500390565b6000600019821415611a0b57611a0b6119ca565b5060010190565b634e487b7160e01b600052604160045260246000fd5b60008219821115611a3b57611a3b6119ca565b500190565b600082611a5d57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b6000816000190483118215151615611a9257611a926119ca565b500290565b600060208284031215611aa957600080fd5b815180151581146114c557600080fd5b60005b83811015611ad4578181015183820152602001611abc565b83811115610f6d5750506000910152565b60008251611af7818460208701611ab9565b9190910192915050565b6020815260008251806020840152611b20816040850160208701611ab9565b601f01601f1916919091016040019291505056fea264697066735822122084f689d6315770aba8c017e4fd002c8b8ca9bf2cffb70d0693aa8123ae2b93f964736f6c63430008090033",
}

// FarmingRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use FarmingRewardsMetaData.ABI instead.
var FarmingRewardsABI = FarmingRewardsMetaData.ABI

// FarmingRewardsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FarmingRewardsMetaData.Bin instead.
var FarmingRewardsBin = FarmingRewardsMetaData.Bin

// DeployFarmingRewards deploys a new Ethereum contract, binding an instance of FarmingRewards to it.
func DeployFarmingRewards(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address) (common.Address, *types.Transaction, *FarmingRewards, error) {
	parsed, err := FarmingRewardsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FarmingRewardsBin), backend, _sigsVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FarmingRewards{FarmingRewardsCaller: FarmingRewardsCaller{contract: contract}, FarmingRewardsTransactor: FarmingRewardsTransactor{contract: contract}, FarmingRewardsFilterer: FarmingRewardsFilterer{contract: contract}}, nil
}

// FarmingRewards is an auto generated Go binding around an Ethereum contract.
type FarmingRewards struct {
	FarmingRewardsCaller     // Read-only binding to the contract
	FarmingRewardsTransactor // Write-only binding to the contract
	FarmingRewardsFilterer   // Log filterer for contract events
}

// FarmingRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarmingRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarmingRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarmingRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarmingRewardsSession struct {
	Contract     *FarmingRewards   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmingRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarmingRewardsCallerSession struct {
	Contract *FarmingRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FarmingRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarmingRewardsTransactorSession struct {
	Contract     *FarmingRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FarmingRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarmingRewardsRaw struct {
	Contract *FarmingRewards // Generic contract binding to access the raw methods on
}

// FarmingRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarmingRewardsCallerRaw struct {
	Contract *FarmingRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// FarmingRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarmingRewardsTransactorRaw struct {
	Contract *FarmingRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarmingRewards creates a new instance of FarmingRewards, bound to a specific deployed contract.
func NewFarmingRewards(address common.Address, backend bind.ContractBackend) (*FarmingRewards, error) {
	contract, err := bindFarmingRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FarmingRewards{FarmingRewardsCaller: FarmingRewardsCaller{contract: contract}, FarmingRewardsTransactor: FarmingRewardsTransactor{contract: contract}, FarmingRewardsFilterer: FarmingRewardsFilterer{contract: contract}}, nil
}

// NewFarmingRewardsCaller creates a new read-only instance of FarmingRewards, bound to a specific deployed contract.
func NewFarmingRewardsCaller(address common.Address, caller bind.ContractCaller) (*FarmingRewardsCaller, error) {
	contract, err := bindFarmingRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsCaller{contract: contract}, nil
}

// NewFarmingRewardsTransactor creates a new write-only instance of FarmingRewards, bound to a specific deployed contract.
func NewFarmingRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*FarmingRewardsTransactor, error) {
	contract, err := bindFarmingRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsTransactor{contract: contract}, nil
}

// NewFarmingRewardsFilterer creates a new log filterer instance of FarmingRewards, bound to a specific deployed contract.
func NewFarmingRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*FarmingRewardsFilterer, error) {
	contract, err := bindFarmingRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsFilterer{contract: contract}, nil
}

// bindFarmingRewards binds a generic wrapper to an already deployed contract.
func bindFarmingRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FarmingRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmingRewards *FarmingRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmingRewards.Contract.FarmingRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmingRewards *FarmingRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingRewards.Contract.FarmingRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmingRewards *FarmingRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmingRewards.Contract.FarmingRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmingRewards *FarmingRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmingRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmingRewards *FarmingRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmingRewards *FarmingRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmingRewards.Contract.contract.Transact(opts, method, params...)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x1744092e.
//
// Solidity: function claimedRewardAmounts(address , address ) view returns(uint256)
func (_FarmingRewards *FarmingRewardsCaller) ClaimedRewardAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FarmingRewards.contract.Call(opts, &out, "claimedRewardAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x1744092e.
//
// Solidity: function claimedRewardAmounts(address , address ) view returns(uint256)
func (_FarmingRewards *FarmingRewardsSession) ClaimedRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FarmingRewards.Contract.ClaimedRewardAmounts(&_FarmingRewards.CallOpts, arg0, arg1)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x1744092e.
//
// Solidity: function claimedRewardAmounts(address , address ) view returns(uint256)
func (_FarmingRewards *FarmingRewardsCallerSession) ClaimedRewardAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _FarmingRewards.Contract.ClaimedRewardAmounts(&_FarmingRewards.CallOpts, arg0, arg1)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_FarmingRewards *FarmingRewardsCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _FarmingRewards.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_FarmingRewards *FarmingRewardsSession) IsPauser(account common.Address) (bool, error) {
	return _FarmingRewards.Contract.IsPauser(&_FarmingRewards.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_FarmingRewards *FarmingRewardsCallerSession) IsPauser(account common.Address) (bool, error) {
	return _FarmingRewards.Contract.IsPauser(&_FarmingRewards.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarmingRewards *FarmingRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmingRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarmingRewards *FarmingRewardsSession) Owner() (common.Address, error) {
	return _FarmingRewards.Contract.Owner(&_FarmingRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FarmingRewards *FarmingRewardsCallerSession) Owner() (common.Address, error) {
	return _FarmingRewards.Contract.Owner(&_FarmingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarmingRewards *FarmingRewardsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FarmingRewards.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarmingRewards *FarmingRewardsSession) Paused() (bool, error) {
	return _FarmingRewards.Contract.Paused(&_FarmingRewards.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_FarmingRewards *FarmingRewardsCallerSession) Paused() (bool, error) {
	return _FarmingRewards.Contract.Paused(&_FarmingRewards.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_FarmingRewards *FarmingRewardsCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FarmingRewards.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_FarmingRewards *FarmingRewardsSession) Pausers(arg0 common.Address) (bool, error) {
	return _FarmingRewards.Contract.Pausers(&_FarmingRewards.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_FarmingRewards *FarmingRewardsCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _FarmingRewards.Contract.Pausers(&_FarmingRewards.CallOpts, arg0)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_FarmingRewards *FarmingRewardsCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmingRewards.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_FarmingRewards *FarmingRewardsSession) SigsVerifier() (common.Address, error) {
	return _FarmingRewards.Contract.SigsVerifier(&_FarmingRewards.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_FarmingRewards *FarmingRewardsCallerSession) SigsVerifier() (common.Address, error) {
	return _FarmingRewards.Contract.SigsVerifier(&_FarmingRewards.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_FarmingRewards *FarmingRewardsTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_FarmingRewards *FarmingRewardsSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _FarmingRewards.Contract.AddPauser(&_FarmingRewards.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _FarmingRewards.Contract.AddPauser(&_FarmingRewards.TransactOpts, account)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6b5d21e9.
//
// Solidity: function claimRewards(bytes _rewardsRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_FarmingRewards *FarmingRewardsTransactor) ClaimRewards(opts *bind.TransactOpts, _rewardsRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "claimRewards", _rewardsRequest, _sigs, _signers, _powers)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6b5d21e9.
//
// Solidity: function claimRewards(bytes _rewardsRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_FarmingRewards *FarmingRewardsSession) ClaimRewards(_rewardsRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _FarmingRewards.Contract.ClaimRewards(&_FarmingRewards.TransactOpts, _rewardsRequest, _sigs, _signers, _powers)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x6b5d21e9.
//
// Solidity: function claimRewards(bytes _rewardsRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) ClaimRewards(_rewardsRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _FarmingRewards.Contract.ClaimRewards(&_FarmingRewards.TransactOpts, _rewardsRequest, _sigs, _signers, _powers)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x825168ff.
//
// Solidity: function contributeToRewardPool(address _token, uint256 _amount) returns()
func (_FarmingRewards *FarmingRewardsTransactor) ContributeToRewardPool(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "contributeToRewardPool", _token, _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x825168ff.
//
// Solidity: function contributeToRewardPool(address _token, uint256 _amount) returns()
func (_FarmingRewards *FarmingRewardsSession) ContributeToRewardPool(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FarmingRewards.Contract.ContributeToRewardPool(&_FarmingRewards.TransactOpts, _token, _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x825168ff.
//
// Solidity: function contributeToRewardPool(address _token, uint256 _amount) returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) ContributeToRewardPool(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FarmingRewards.Contract.ContributeToRewardPool(&_FarmingRewards.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_FarmingRewards *FarmingRewardsTransactor) DrainToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "drainToken", _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_FarmingRewards *FarmingRewardsSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FarmingRewards.Contract.DrainToken(&_FarmingRewards.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _FarmingRewards.Contract.DrainToken(&_FarmingRewards.TransactOpts, _token, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarmingRewards *FarmingRewardsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarmingRewards *FarmingRewardsSession) Pause() (*types.Transaction, error) {
	return _FarmingRewards.Contract.Pause(&_FarmingRewards.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) Pause() (*types.Transaction, error) {
	return _FarmingRewards.Contract.Pause(&_FarmingRewards.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_FarmingRewards *FarmingRewardsTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_FarmingRewards *FarmingRewardsSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _FarmingRewards.Contract.RemovePauser(&_FarmingRewards.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _FarmingRewards.Contract.RemovePauser(&_FarmingRewards.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarmingRewards *FarmingRewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarmingRewards *FarmingRewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _FarmingRewards.Contract.RenounceOwnership(&_FarmingRewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FarmingRewards.Contract.RenounceOwnership(&_FarmingRewards.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_FarmingRewards *FarmingRewardsTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_FarmingRewards *FarmingRewardsSession) RenouncePauser() (*types.Transaction, error) {
	return _FarmingRewards.Contract.RenouncePauser(&_FarmingRewards.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _FarmingRewards.Contract.RenouncePauser(&_FarmingRewards.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarmingRewards *FarmingRewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarmingRewards *FarmingRewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarmingRewards.Contract.TransferOwnership(&_FarmingRewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FarmingRewards.Contract.TransferOwnership(&_FarmingRewards.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarmingRewards *FarmingRewardsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarmingRewards *FarmingRewardsSession) Unpause() (*types.Transaction, error) {
	return _FarmingRewards.Contract.Unpause(&_FarmingRewards.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) Unpause() (*types.Transaction, error) {
	return _FarmingRewards.Contract.Unpause(&_FarmingRewards.TransactOpts)
}

// FarmingRewardsFarmingRewardClaimedIterator is returned from FilterFarmingRewardClaimed and is used to iterate over the raw logs and unpacked data for FarmingRewardClaimed events raised by the FarmingRewards contract.
type FarmingRewardsFarmingRewardClaimedIterator struct {
	Event *FarmingRewardsFarmingRewardClaimed // Event containing the contract specifics and raw log

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
func (it *FarmingRewardsFarmingRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRewardsFarmingRewardClaimed)
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
		it.Event = new(FarmingRewardsFarmingRewardClaimed)
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
func (it *FarmingRewardsFarmingRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRewardsFarmingRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRewardsFarmingRewardClaimed represents a FarmingRewardClaimed event raised by the FarmingRewards contract.
type FarmingRewardsFarmingRewardClaimed struct {
	Recipient common.Address
	Token     common.Address
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFarmingRewardClaimed is a free log retrieval operation binding the contract event 0x97e6c3172350795e26977663112f38653689372e771e85bad9fbadb1af0e98b2.
//
// Solidity: event FarmingRewardClaimed(address indexed recipient, address indexed token, uint256 reward)
func (_FarmingRewards *FarmingRewardsFilterer) FilterFarmingRewardClaimed(opts *bind.FilterOpts, recipient []common.Address, token []common.Address) (*FarmingRewardsFarmingRewardClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FarmingRewards.contract.FilterLogs(opts, "FarmingRewardClaimed", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsFarmingRewardClaimedIterator{contract: _FarmingRewards.contract, event: "FarmingRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchFarmingRewardClaimed is a free log subscription operation binding the contract event 0x97e6c3172350795e26977663112f38653689372e771e85bad9fbadb1af0e98b2.
//
// Solidity: event FarmingRewardClaimed(address indexed recipient, address indexed token, uint256 reward)
func (_FarmingRewards *FarmingRewardsFilterer) WatchFarmingRewardClaimed(opts *bind.WatchOpts, sink chan<- *FarmingRewardsFarmingRewardClaimed, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FarmingRewards.contract.WatchLogs(opts, "FarmingRewardClaimed", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRewardsFarmingRewardClaimed)
				if err := _FarmingRewards.contract.UnpackLog(event, "FarmingRewardClaimed", log); err != nil {
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

// ParseFarmingRewardClaimed is a log parse operation binding the contract event 0x97e6c3172350795e26977663112f38653689372e771e85bad9fbadb1af0e98b2.
//
// Solidity: event FarmingRewardClaimed(address indexed recipient, address indexed token, uint256 reward)
func (_FarmingRewards *FarmingRewardsFilterer) ParseFarmingRewardClaimed(log types.Log) (*FarmingRewardsFarmingRewardClaimed, error) {
	event := new(FarmingRewardsFarmingRewardClaimed)
	if err := _FarmingRewards.contract.UnpackLog(event, "FarmingRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRewardsFarmingRewardContributedIterator is returned from FilterFarmingRewardContributed and is used to iterate over the raw logs and unpacked data for FarmingRewardContributed events raised by the FarmingRewards contract.
type FarmingRewardsFarmingRewardContributedIterator struct {
	Event *FarmingRewardsFarmingRewardContributed // Event containing the contract specifics and raw log

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
func (it *FarmingRewardsFarmingRewardContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRewardsFarmingRewardContributed)
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
		it.Event = new(FarmingRewardsFarmingRewardContributed)
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
func (it *FarmingRewardsFarmingRewardContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRewardsFarmingRewardContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRewardsFarmingRewardContributed represents a FarmingRewardContributed event raised by the FarmingRewards contract.
type FarmingRewardsFarmingRewardContributed struct {
	Contributor  common.Address
	Token        common.Address
	Contribution *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFarmingRewardContributed is a free log retrieval operation binding the contract event 0x40aa1b9a9157bc37a09a78d5a46e53087b82ee0034ebe896d4d1a52f31b333d4.
//
// Solidity: event FarmingRewardContributed(address indexed contributor, address indexed token, uint256 contribution)
func (_FarmingRewards *FarmingRewardsFilterer) FilterFarmingRewardContributed(opts *bind.FilterOpts, contributor []common.Address, token []common.Address) (*FarmingRewardsFarmingRewardContributedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FarmingRewards.contract.FilterLogs(opts, "FarmingRewardContributed", contributorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsFarmingRewardContributedIterator{contract: _FarmingRewards.contract, event: "FarmingRewardContributed", logs: logs, sub: sub}, nil
}

// WatchFarmingRewardContributed is a free log subscription operation binding the contract event 0x40aa1b9a9157bc37a09a78d5a46e53087b82ee0034ebe896d4d1a52f31b333d4.
//
// Solidity: event FarmingRewardContributed(address indexed contributor, address indexed token, uint256 contribution)
func (_FarmingRewards *FarmingRewardsFilterer) WatchFarmingRewardContributed(opts *bind.WatchOpts, sink chan<- *FarmingRewardsFarmingRewardContributed, contributor []common.Address, token []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FarmingRewards.contract.WatchLogs(opts, "FarmingRewardContributed", contributorRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRewardsFarmingRewardContributed)
				if err := _FarmingRewards.contract.UnpackLog(event, "FarmingRewardContributed", log); err != nil {
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

// ParseFarmingRewardContributed is a log parse operation binding the contract event 0x40aa1b9a9157bc37a09a78d5a46e53087b82ee0034ebe896d4d1a52f31b333d4.
//
// Solidity: event FarmingRewardContributed(address indexed contributor, address indexed token, uint256 contribution)
func (_FarmingRewards *FarmingRewardsFilterer) ParseFarmingRewardContributed(log types.Log) (*FarmingRewardsFarmingRewardContributed, error) {
	event := new(FarmingRewardsFarmingRewardContributed)
	if err := _FarmingRewards.contract.UnpackLog(event, "FarmingRewardContributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FarmingRewards contract.
type FarmingRewardsOwnershipTransferredIterator struct {
	Event *FarmingRewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FarmingRewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRewardsOwnershipTransferred)
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
		it.Event = new(FarmingRewardsOwnershipTransferred)
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
func (it *FarmingRewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRewardsOwnershipTransferred represents a OwnershipTransferred event raised by the FarmingRewards contract.
type FarmingRewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarmingRewards *FarmingRewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FarmingRewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarmingRewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsOwnershipTransferredIterator{contract: _FarmingRewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FarmingRewards *FarmingRewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FarmingRewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FarmingRewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRewardsOwnershipTransferred)
				if err := _FarmingRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FarmingRewards *FarmingRewardsFilterer) ParseOwnershipTransferred(log types.Log) (*FarmingRewardsOwnershipTransferred, error) {
	event := new(FarmingRewardsOwnershipTransferred)
	if err := _FarmingRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRewardsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the FarmingRewards contract.
type FarmingRewardsPausedIterator struct {
	Event *FarmingRewardsPaused // Event containing the contract specifics and raw log

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
func (it *FarmingRewardsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRewardsPaused)
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
		it.Event = new(FarmingRewardsPaused)
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
func (it *FarmingRewardsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRewardsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRewardsPaused represents a Paused event raised by the FarmingRewards contract.
type FarmingRewardsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarmingRewards *FarmingRewardsFilterer) FilterPaused(opts *bind.FilterOpts) (*FarmingRewardsPausedIterator, error) {

	logs, sub, err := _FarmingRewards.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsPausedIterator{contract: _FarmingRewards.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarmingRewards *FarmingRewardsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FarmingRewardsPaused) (event.Subscription, error) {

	logs, sub, err := _FarmingRewards.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRewardsPaused)
				if err := _FarmingRewards.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_FarmingRewards *FarmingRewardsFilterer) ParsePaused(log types.Log) (*FarmingRewardsPaused, error) {
	event := new(FarmingRewardsPaused)
	if err := _FarmingRewards.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRewardsPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the FarmingRewards contract.
type FarmingRewardsPauserAddedIterator struct {
	Event *FarmingRewardsPauserAdded // Event containing the contract specifics and raw log

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
func (it *FarmingRewardsPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRewardsPauserAdded)
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
		it.Event = new(FarmingRewardsPauserAdded)
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
func (it *FarmingRewardsPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRewardsPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRewardsPauserAdded represents a PauserAdded event raised by the FarmingRewards contract.
type FarmingRewardsPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_FarmingRewards *FarmingRewardsFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*FarmingRewardsPauserAddedIterator, error) {

	logs, sub, err := _FarmingRewards.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsPauserAddedIterator{contract: _FarmingRewards.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_FarmingRewards *FarmingRewardsFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *FarmingRewardsPauserAdded) (event.Subscription, error) {

	logs, sub, err := _FarmingRewards.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRewardsPauserAdded)
				if err := _FarmingRewards.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_FarmingRewards *FarmingRewardsFilterer) ParsePauserAdded(log types.Log) (*FarmingRewardsPauserAdded, error) {
	event := new(FarmingRewardsPauserAdded)
	if err := _FarmingRewards.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRewardsPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the FarmingRewards contract.
type FarmingRewardsPauserRemovedIterator struct {
	Event *FarmingRewardsPauserRemoved // Event containing the contract specifics and raw log

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
func (it *FarmingRewardsPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRewardsPauserRemoved)
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
		it.Event = new(FarmingRewardsPauserRemoved)
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
func (it *FarmingRewardsPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRewardsPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRewardsPauserRemoved represents a PauserRemoved event raised by the FarmingRewards contract.
type FarmingRewardsPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_FarmingRewards *FarmingRewardsFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*FarmingRewardsPauserRemovedIterator, error) {

	logs, sub, err := _FarmingRewards.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsPauserRemovedIterator{contract: _FarmingRewards.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_FarmingRewards *FarmingRewardsFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *FarmingRewardsPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _FarmingRewards.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRewardsPauserRemoved)
				if err := _FarmingRewards.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_FarmingRewards *FarmingRewardsFilterer) ParsePauserRemoved(log types.Log) (*FarmingRewardsPauserRemoved, error) {
	event := new(FarmingRewardsPauserRemoved)
	if err := _FarmingRewards.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRewardsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the FarmingRewards contract.
type FarmingRewardsUnpausedIterator struct {
	Event *FarmingRewardsUnpaused // Event containing the contract specifics and raw log

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
func (it *FarmingRewardsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRewardsUnpaused)
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
		it.Event = new(FarmingRewardsUnpaused)
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
func (it *FarmingRewardsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRewardsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRewardsUnpaused represents a Unpaused event raised by the FarmingRewards contract.
type FarmingRewardsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarmingRewards *FarmingRewardsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FarmingRewardsUnpausedIterator, error) {

	logs, sub, err := _FarmingRewards.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsUnpausedIterator{contract: _FarmingRewards.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarmingRewards *FarmingRewardsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FarmingRewardsUnpaused) (event.Subscription, error) {

	logs, sub, err := _FarmingRewards.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRewardsUnpaused)
				if err := _FarmingRewards.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_FarmingRewards *FarmingRewardsFilterer) ParseUnpaused(log types.Log) (*FarmingRewardsUnpaused, error) {
	event := new(FarmingRewardsUnpaused)
	if err := _FarmingRewards.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernMetaData contains all meta data concerning the Govern contract.
var GovernMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collector\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"passed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ConfirmParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"CreateParamProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGovern.VoteOption\",\"name\":\"vote\",\"type\":\"uint8\"}],\"name\":\"VoteParam\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"celerToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectForfeiture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"confirmParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"createParamProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forfeiture\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getParamProposalVote\",\"outputs\":[{\"internalType\":\"enumGovern.VoteOption\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextParamProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paramProposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"voteDeadline\",\"type\":\"uint256\"},{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumGovern.VoteOption\",\"name\":\"_vote\",\"type\":\"uint8\"}],\"name\":\"voteParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b50604051620015f7380380620015f783398101604081905262000034916200006b565b6001600160a01b0392831660805290821660a0521660c052620000bf565b6001600160a01b03811681146200006857600080fd5b50565b6000806000606084860312156200008157600080fd5b83516200008e8162000052565b6020850151909350620000a18162000052565b6040850151909250620000b48162000052565b809150509250925092565b60805160a05160c0516114c96200012e600039600081816101da0152610620015260008181610214015281816105fd0152818161095e0152610be10152600081816101040152818161026b01528181610653015281816108e101528181610a380152610b0c01526114c96000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806382d7b4b811610081578063934a18ec1161005b578063934a18ec146101fc578063c6c21e9d1461020f578063e478ed9d1461023657600080fd5b806382d7b4b8146101c45780638338f0e5146101cc578063913e77ad146101d557600080fd5b80634cf088d9116100b25780634cf088d9146100ff578063581c53c51461013e5780637e5fb8f31461015e57600080fd5b806322da7927146100ce57806325ed6b35146100ea575b600080fd5b6100d760015481565b6040519081526020015b60405180910390f35b6100fd6100f8366004610fc5565b610249565b005b6101267f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100e1565b61015161014c36600461100a565b61056c565b6040516100e19190611059565b6101b261016c366004611067565b6000602081905290815260409020805460018201546002830154600384015460048501546005909501546001600160a01b03909416949293919260ff9182169290911686565b6040516100e196959493929190611090565b6100fd61059a565b6100d760025481565b6101267f000000000000000000000000000000000000000000000000000000000000000081565b6100fd61020a366004611067565b61064c565b6101267f000000000000000000000000000000000000000000000000000000000000000081565b6100fd6102443660046110e2565b6109fc565b33600360405163a310624f60e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a310624f9060240160206040518083038186803b1580156102ad57600080fd5b505afa1580156102c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e59190611112565b60038111156102f6576102f661102f565b146103485760405162461bcd60e51b815260206004820152601f60248201527f566f746572206973206e6f74206120626f6e6465642076616c696461746f720060448201526064015b60405180910390fd5b60008381526020819052604090206001600582015460ff1660028111156103715761037161102f565b146103be5760405162461bcd60e51b815260206004820152601760248201527f496e76616c69642070726f706f73616c20737461747573000000000000000000604482015260640161033f565b806002015443106104115760405162461bcd60e51b815260206004820152601460248201527f566f746520646561646c696e6520706173736564000000000000000000000000604482015260640161033f565b6001600160a01b038216600090815260068201602052604081205460ff1660038111156104405761044061102f565b1461048d5760405162461bcd60e51b815260206004820152600f60248201527f566f7465722068617320766f7465640000000000000000000000000000000000604482015260640161033f565b60008360038111156104a1576104a161102f565b14156104ef5760405162461bcd60e51b815260206004820152600c60248201527f496e76616c696420766f74650000000000000000000000000000000000000000604482015260640161033f565b6001600160a01b03821660009081526006820160205260409020805484919060ff191660018360038111156105265761052661102f565b02179055507f06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d6584838560405161055e9392919061112f565b60405180910390a150505050565b6000828152602081815260408083206001600160a01b038516845260060190915290205460ff165b92915050565b6000600254116105ec5760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20636f6c6c6563740000000000000000000000000000604482015260640161033f565b600254610645906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016907f000000000000000000000000000000000000000000000000000000000000000090610c6a565b6000600255565b60008060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634021d4d56040518163ffffffff1660e01b815260040160006040518083038186803b1580156106aa57600080fd5b505afa1580156106be573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106e691908101906111c2565b905060005b81518163ffffffff1610156107b557600161072986848463ffffffff168151811061071857610718611298565b60200260200101516000015161056c565b600381111561073a5761073a61102f565b141561077257818163ffffffff168151811061075857610758611298565b6020026020010151602001518461076f91906112c4565b93505b818163ffffffff168151811061078a5761078a611298565b602002602001015160200151836107a191906112c4565b9250806107ad816112dc565b9150506106eb565b50600060036107c5846002611300565b6107cf919061131f565b6107da9060016112c4565b60008681526020819052604090209085101591506001600582015460ff1660028111156108095761080961102f565b146108565760405162461bcd60e51b815260206004820152601760248201527f496e76616c69642070726f706f73616c20737461747573000000000000000000604482015260640161033f565b80600201544310156108aa5760405162461bcd60e51b815260206004820152601960248201527f566f746520646561646c696e65206e6f74207265616368656400000000000000604482015260640161033f565b60058101805460ff19166002179055811561098f57600381015460048083015460405163e909156d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169363e909156d936109189360ff9092169201611341565b600060405180830381600087803b15801561093257600080fd5b505af1158015610946573d6000803e3d6000fd5b50508254600184015461098a93506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116935090911690610c6a565b6109ab565b8060010154600260008282546109a591906112c4565b90915550505b600381015460048201546040517fd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9926109ec928a92879260ff16919061135c565b60405180910390a1505050505050565b60018054600081815260208190526040902091610a1991906112c4565b600155604051631042b80b60e21b815233906000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063410ae02c90610a6d908490600401611387565b60206040518083038186803b158015610a8557600080fd5b505afa158015610a99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abd9190611395565b83547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038481169190911785556001808601839055604051631042b80b60e21b81529293507f00000000000000000000000000000000000000000000000000000000000000009091169163410ae02c91610b4191600401611387565b60206040518083038186803b158015610b5957600080fd5b505afa158015610b6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b919190611395565b610b9b90436112c4565b600284015560038301805486919060ff19166001836008811115610bc157610bc161102f565b02179055506004830184905560058301805460ff19166001179055610c117f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316833084610cff565b7f4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f16060018054610c4091906113ae565b838386600201548989604051610c5b969594939291906113c5565b60405180910390a15050505050565b6040516001600160a01b038316602482015260448101829052610cfa90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610d3d565b505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610d379085906323b872dd60e01b90608401610c96565b50505050565b6000610d92826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610e229092919063ffffffff16565b805190915015610cfa5780806020019051810190610db091906113f6565b610cfa5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161033f565b6060610e318484600085610e3b565b90505b9392505050565b606082471015610eb35760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161033f565b843b610f015760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161033f565b600080866001600160a01b03168587604051610f1d9190611444565b60006040518083038185875af1925050503d8060008114610f5a576040519150601f19603f3d011682016040523d82523d6000602084013e610f5f565b606091505b5091509150610f6f828286610f7c565b925050505b949350505050565b60608315610f8b575081610e34565b825115610f9b5782518084602001fd5b8160405162461bcd60e51b815260040161033f9190611460565b60048110610fc257600080fd5b50565b60008060408385031215610fd857600080fd5b823591506020830135610fea81610fb5565b809150509250929050565b6001600160a01b0381168114610fc257600080fd5b6000806040838503121561101d57600080fd5b823591506020830135610fea81610ff5565b634e487b7160e01b600052602160045260246000fd5b600481106110555761105561102f565b9052565b602081016105948284611045565b60006020828403121561107957600080fd5b5035919050565b600981106110555761105561102f565b6001600160a01b0387168152602081018690526040810185905260c081016110bb6060830186611080565b836080830152600383106110d1576110d161102f565b8260a0830152979650505050505050565b600080604083850312156110f557600080fd5b82356009811061110457600080fd5b946020939093013593505050565b60006020828403121561112457600080fd5b8151610e3481610fb5565b8381526001600160a01b038316602082015260608101610f746040830184611045565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff8111828210171561118b5761118b611152565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156111ba576111ba611152565b604052919050565b600060208083850312156111d557600080fd5b825167ffffffffffffffff808211156111ed57600080fd5b818501915085601f83011261120157600080fd5b81518181111561121357611213611152565b611221848260051b01611191565b818152848101925060069190911b83018401908782111561124157600080fd5b928401925b8184101561128d576040848903121561125f5760008081fd5b611267611168565b845161127281610ff5565b81528486015186820152835260409093019291840191611246565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156112d7576112d76112ae565b500190565b600063ffffffff808316818114156112f6576112f66112ae565b6001019392505050565b600081600019048311821515161561131a5761131a6112ae565b500290565b60008261133c57634e487b7160e01b600052601260045260246000fd5b500490565b6040810161134f8285611080565b8260208301529392505050565b8481528315156020820152608081016113786040830185611080565b82606083015295945050505050565b602081016105948284611080565b6000602082840312156113a757600080fd5b5051919050565b6000828210156113c0576113c06112ae565b500390565b8681526001600160a01b0386166020820152604081018590526060810184905260c081016110d16080830185611080565b60006020828403121561140857600080fd5b81518015158114610e3457600080fd5b60005b8381101561143357818101518382015260200161141b565b83811115610d375750506000910152565b60008251611456818460208701611418565b9190910192915050565b602081526000825180602084015261147f816040850160208701611418565b601f01601f1916919091016040019291505056fea2646970667358221220f34a4b9bcde6d6a96c47c5dd0fec757acf39b72ca883c099df7274e38687eb1664736f6c63430008090033",
}

// GovernABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernMetaData.ABI instead.
var GovernABI = GovernMetaData.ABI

// GovernBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GovernMetaData.Bin instead.
var GovernBin = GovernMetaData.Bin

// DeployGovern deploys a new Ethereum contract, binding an instance of Govern to it.
func DeployGovern(auth *bind.TransactOpts, backend bind.ContractBackend, _staking common.Address, _celerTokenAddress common.Address, _collector common.Address) (common.Address, *types.Transaction, *Govern, error) {
	parsed, err := GovernMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GovernBin), backend, _staking, _celerTokenAddress, _collector)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Govern{GovernCaller: GovernCaller{contract: contract}, GovernTransactor: GovernTransactor{contract: contract}, GovernFilterer: GovernFilterer{contract: contract}}, nil
}

// Govern is an auto generated Go binding around an Ethereum contract.
type Govern struct {
	GovernCaller     // Read-only binding to the contract
	GovernTransactor // Write-only binding to the contract
	GovernFilterer   // Log filterer for contract events
}

// GovernCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernSession struct {
	Contract     *Govern           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernCallerSession struct {
	Contract *GovernCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GovernTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernTransactorSession struct {
	Contract     *GovernTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernRaw struct {
	Contract *Govern // Generic contract binding to access the raw methods on
}

// GovernCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernCallerRaw struct {
	Contract *GovernCaller // Generic read-only contract binding to access the raw methods on
}

// GovernTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernTransactorRaw struct {
	Contract *GovernTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovern creates a new instance of Govern, bound to a specific deployed contract.
func NewGovern(address common.Address, backend bind.ContractBackend) (*Govern, error) {
	contract, err := bindGovern(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Govern{GovernCaller: GovernCaller{contract: contract}, GovernTransactor: GovernTransactor{contract: contract}, GovernFilterer: GovernFilterer{contract: contract}}, nil
}

// NewGovernCaller creates a new read-only instance of Govern, bound to a specific deployed contract.
func NewGovernCaller(address common.Address, caller bind.ContractCaller) (*GovernCaller, error) {
	contract, err := bindGovern(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernCaller{contract: contract}, nil
}

// NewGovernTransactor creates a new write-only instance of Govern, bound to a specific deployed contract.
func NewGovernTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernTransactor, error) {
	contract, err := bindGovern(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernTransactor{contract: contract}, nil
}

// NewGovernFilterer creates a new log filterer instance of Govern, bound to a specific deployed contract.
func NewGovernFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernFilterer, error) {
	contract, err := bindGovern(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernFilterer{contract: contract}, nil
}

// bindGovern binds a generic wrapper to an already deployed contract.
func bindGovern(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Govern *GovernRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Govern.Contract.GovernCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Govern *GovernRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govern.Contract.GovernTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Govern *GovernRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Govern.Contract.GovernTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Govern *GovernCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Govern.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Govern *GovernTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govern.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Govern *GovernTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Govern.Contract.contract.Transact(opts, method, params...)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Govern *GovernCaller) CelerToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "celerToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Govern *GovernSession) CelerToken() (common.Address, error) {
	return _Govern.Contract.CelerToken(&_Govern.CallOpts)
}

// CelerToken is a free data retrieval call binding the contract method 0xc6c21e9d.
//
// Solidity: function celerToken() view returns(address)
func (_Govern *GovernCallerSession) CelerToken() (common.Address, error) {
	return _Govern.Contract.CelerToken(&_Govern.CallOpts)
}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Govern *GovernCaller) Collector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "collector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Govern *GovernSession) Collector() (common.Address, error) {
	return _Govern.Contract.Collector(&_Govern.CallOpts)
}

// Collector is a free data retrieval call binding the contract method 0x913e77ad.
//
// Solidity: function collector() view returns(address)
func (_Govern *GovernCallerSession) Collector() (common.Address, error) {
	return _Govern.Contract.Collector(&_Govern.CallOpts)
}

// Forfeiture is a free data retrieval call binding the contract method 0x8338f0e5.
//
// Solidity: function forfeiture() view returns(uint256)
func (_Govern *GovernCaller) Forfeiture(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "forfeiture")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Forfeiture is a free data retrieval call binding the contract method 0x8338f0e5.
//
// Solidity: function forfeiture() view returns(uint256)
func (_Govern *GovernSession) Forfeiture() (*big.Int, error) {
	return _Govern.Contract.Forfeiture(&_Govern.CallOpts)
}

// Forfeiture is a free data retrieval call binding the contract method 0x8338f0e5.
//
// Solidity: function forfeiture() view returns(uint256)
func (_Govern *GovernCallerSession) Forfeiture() (*big.Int, error) {
	return _Govern.Contract.Forfeiture(&_Govern.CallOpts)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Govern *GovernCaller) GetParamProposalVote(opts *bind.CallOpts, _proposalId *big.Int, _voter common.Address) (uint8, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "getParamProposalVote", _proposalId, _voter)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Govern *GovernSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _Govern.Contract.GetParamProposalVote(&_Govern.CallOpts, _proposalId, _voter)
}

// GetParamProposalVote is a free data retrieval call binding the contract method 0x581c53c5.
//
// Solidity: function getParamProposalVote(uint256 _proposalId, address _voter) view returns(uint8)
func (_Govern *GovernCallerSession) GetParamProposalVote(_proposalId *big.Int, _voter common.Address) (uint8, error) {
	return _Govern.Contract.GetParamProposalVote(&_Govern.CallOpts, _proposalId, _voter)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Govern *GovernCaller) NextParamProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "nextParamProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Govern *GovernSession) NextParamProposalId() (*big.Int, error) {
	return _Govern.Contract.NextParamProposalId(&_Govern.CallOpts)
}

// NextParamProposalId is a free data retrieval call binding the contract method 0x22da7927.
//
// Solidity: function nextParamProposalId() view returns(uint256)
func (_Govern *GovernCallerSession) NextParamProposalId() (*big.Int, error) {
	return _Govern.Contract.NextParamProposalId(&_Govern.CallOpts)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Govern *GovernCaller) ParamProposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "paramProposals", arg0)

	outstruct := new(struct {
		Proposer     common.Address
		Deposit      *big.Int
		VoteDeadline *big.Int
		Name         uint8
		NewValue     *big.Int
		Status       uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deposit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VoteDeadline = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.NewValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Govern *GovernSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Govern.Contract.ParamProposals(&_Govern.CallOpts, arg0)
}

// ParamProposals is a free data retrieval call binding the contract method 0x7e5fb8f3.
//
// Solidity: function paramProposals(uint256 ) view returns(address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue, uint8 status)
func (_Govern *GovernCallerSession) ParamProposals(arg0 *big.Int) (struct {
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Status       uint8
}, error) {
	return _Govern.Contract.ParamProposals(&_Govern.CallOpts, arg0)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Govern *GovernCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Govern.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Govern *GovernSession) Staking() (common.Address, error) {
	return _Govern.Contract.Staking(&_Govern.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Govern *GovernCallerSession) Staking() (common.Address, error) {
	return _Govern.Contract.Staking(&_Govern.CallOpts)
}

// CollectForfeiture is a paid mutator transaction binding the contract method 0x82d7b4b8.
//
// Solidity: function collectForfeiture() returns()
func (_Govern *GovernTransactor) CollectForfeiture(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "collectForfeiture")
}

// CollectForfeiture is a paid mutator transaction binding the contract method 0x82d7b4b8.
//
// Solidity: function collectForfeiture() returns()
func (_Govern *GovernSession) CollectForfeiture() (*types.Transaction, error) {
	return _Govern.Contract.CollectForfeiture(&_Govern.TransactOpts)
}

// CollectForfeiture is a paid mutator transaction binding the contract method 0x82d7b4b8.
//
// Solidity: function collectForfeiture() returns()
func (_Govern *GovernTransactorSession) CollectForfeiture() (*types.Transaction, error) {
	return _Govern.Contract.CollectForfeiture(&_Govern.TransactOpts)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_Govern *GovernTransactor) ConfirmParamProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "confirmParamProposal", _proposalId)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_Govern *GovernSession) ConfirmParamProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.ConfirmParamProposal(&_Govern.TransactOpts, _proposalId)
}

// ConfirmParamProposal is a paid mutator transaction binding the contract method 0x934a18ec.
//
// Solidity: function confirmParamProposal(uint256 _proposalId) returns()
func (_Govern *GovernTransactorSession) ConfirmParamProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.ConfirmParamProposal(&_Govern.TransactOpts, _proposalId)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Govern *GovernTransactor) CreateParamProposal(opts *bind.TransactOpts, _name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "createParamProposal", _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Govern *GovernSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.CreateParamProposal(&_Govern.TransactOpts, _name, _value)
}

// CreateParamProposal is a paid mutator transaction binding the contract method 0xe478ed9d.
//
// Solidity: function createParamProposal(uint8 _name, uint256 _value) returns()
func (_Govern *GovernTransactorSession) CreateParamProposal(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Govern.Contract.CreateParamProposal(&_Govern.TransactOpts, _name, _value)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_Govern *GovernTransactor) VoteParam(opts *bind.TransactOpts, _proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _Govern.contract.Transact(opts, "voteParam", _proposalId, _vote)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_Govern *GovernSession) VoteParam(_proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _Govern.Contract.VoteParam(&_Govern.TransactOpts, _proposalId, _vote)
}

// VoteParam is a paid mutator transaction binding the contract method 0x25ed6b35.
//
// Solidity: function voteParam(uint256 _proposalId, uint8 _vote) returns()
func (_Govern *GovernTransactorSession) VoteParam(_proposalId *big.Int, _vote uint8) (*types.Transaction, error) {
	return _Govern.Contract.VoteParam(&_Govern.TransactOpts, _proposalId, _vote)
}

// GovernConfirmParamProposalIterator is returned from FilterConfirmParamProposal and is used to iterate over the raw logs and unpacked data for ConfirmParamProposal events raised by the Govern contract.
type GovernConfirmParamProposalIterator struct {
	Event *GovernConfirmParamProposal // Event containing the contract specifics and raw log

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
func (it *GovernConfirmParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernConfirmParamProposal)
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
		it.Event = new(GovernConfirmParamProposal)
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
func (it *GovernConfirmParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernConfirmParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernConfirmParamProposal represents a ConfirmParamProposal event raised by the Govern contract.
type GovernConfirmParamProposal struct {
	ProposalId *big.Int
	Passed     bool
	Name       uint8
	NewValue   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfirmParamProposal is a free log retrieval operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) FilterConfirmParamProposal(opts *bind.FilterOpts) (*GovernConfirmParamProposalIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return &GovernConfirmParamProposalIterator{contract: _Govern.contract, event: "ConfirmParamProposal", logs: logs, sub: sub}, nil
}

// WatchConfirmParamProposal is a free log subscription operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) WatchConfirmParamProposal(opts *bind.WatchOpts, sink chan<- *GovernConfirmParamProposal) (event.Subscription, error) {

	logs, sub, err := _Govern.contract.WatchLogs(opts, "ConfirmParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernConfirmParamProposal)
				if err := _Govern.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
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

// ParseConfirmParamProposal is a log parse operation binding the contract event 0xd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a9.
//
// Solidity: event ConfirmParamProposal(uint256 proposalId, bool passed, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) ParseConfirmParamProposal(log types.Log) (*GovernConfirmParamProposal, error) {
	event := new(GovernConfirmParamProposal)
	if err := _Govern.contract.UnpackLog(event, "ConfirmParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernCreateParamProposalIterator is returned from FilterCreateParamProposal and is used to iterate over the raw logs and unpacked data for CreateParamProposal events raised by the Govern contract.
type GovernCreateParamProposalIterator struct {
	Event *GovernCreateParamProposal // Event containing the contract specifics and raw log

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
func (it *GovernCreateParamProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernCreateParamProposal)
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
		it.Event = new(GovernCreateParamProposal)
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
func (it *GovernCreateParamProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernCreateParamProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernCreateParamProposal represents a CreateParamProposal event raised by the Govern contract.
type GovernCreateParamProposal struct {
	ProposalId   *big.Int
	Proposer     common.Address
	Deposit      *big.Int
	VoteDeadline *big.Int
	Name         uint8
	NewValue     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateParamProposal is a free log retrieval operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) FilterCreateParamProposal(opts *bind.FilterOpts) (*GovernCreateParamProposalIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return &GovernCreateParamProposalIterator{contract: _Govern.contract, event: "CreateParamProposal", logs: logs, sub: sub}, nil
}

// WatchCreateParamProposal is a free log subscription operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) WatchCreateParamProposal(opts *bind.WatchOpts, sink chan<- *GovernCreateParamProposal) (event.Subscription, error) {

	logs, sub, err := _Govern.contract.WatchLogs(opts, "CreateParamProposal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernCreateParamProposal)
				if err := _Govern.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
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

// ParseCreateParamProposal is a log parse operation binding the contract event 0x4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f160.
//
// Solidity: event CreateParamProposal(uint256 proposalId, address proposer, uint256 deposit, uint256 voteDeadline, uint8 name, uint256 newValue)
func (_Govern *GovernFilterer) ParseCreateParamProposal(log types.Log) (*GovernCreateParamProposal, error) {
	event := new(GovernCreateParamProposal)
	if err := _Govern.contract.UnpackLog(event, "CreateParamProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernVoteParamIterator is returned from FilterVoteParam and is used to iterate over the raw logs and unpacked data for VoteParam events raised by the Govern contract.
type GovernVoteParamIterator struct {
	Event *GovernVoteParam // Event containing the contract specifics and raw log

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
func (it *GovernVoteParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernVoteParam)
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
		it.Event = new(GovernVoteParam)
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
func (it *GovernVoteParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernVoteParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernVoteParam represents a VoteParam event raised by the Govern contract.
type GovernVoteParam struct {
	ProposalId *big.Int
	Voter      common.Address
	Vote       uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteParam is a free log retrieval operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Govern *GovernFilterer) FilterVoteParam(opts *bind.FilterOpts) (*GovernVoteParamIterator, error) {

	logs, sub, err := _Govern.contract.FilterLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return &GovernVoteParamIterator{contract: _Govern.contract, event: "VoteParam", logs: logs, sub: sub}, nil
}

// WatchVoteParam is a free log subscription operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Govern *GovernFilterer) WatchVoteParam(opts *bind.WatchOpts, sink chan<- *GovernVoteParam) (event.Subscription, error) {

	logs, sub, err := _Govern.contract.WatchLogs(opts, "VoteParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernVoteParam)
				if err := _Govern.contract.UnpackLog(event, "VoteParam", log); err != nil {
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

// ParseVoteParam is a log parse operation binding the contract event 0x06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d65.
//
// Solidity: event VoteParam(uint256 proposalId, address voter, uint8 vote)
func (_Govern *GovernFilterer) ParseVoteParam(log types.Log) (*GovernVoteParam, error) {
	event := new(GovernVoteParam)
	if err := _Govern.contract.UnpackLog(event, "VoteParam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISigsVerifierMetaData contains all meta data concerning the ISigsVerifier contract.
var ISigsVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISigsVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use ISigsVerifierMetaData.ABI instead.
var ISigsVerifierABI = ISigsVerifierMetaData.ABI

// ISigsVerifier is an auto generated Go binding around an Ethereum contract.
type ISigsVerifier struct {
	ISigsVerifierCaller     // Read-only binding to the contract
	ISigsVerifierTransactor // Write-only binding to the contract
	ISigsVerifierFilterer   // Log filterer for contract events
}

// ISigsVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISigsVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISigsVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISigsVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISigsVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISigsVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISigsVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISigsVerifierSession struct {
	Contract     *ISigsVerifier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISigsVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISigsVerifierCallerSession struct {
	Contract *ISigsVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISigsVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISigsVerifierTransactorSession struct {
	Contract     *ISigsVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISigsVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISigsVerifierRaw struct {
	Contract *ISigsVerifier // Generic contract binding to access the raw methods on
}

// ISigsVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISigsVerifierCallerRaw struct {
	Contract *ISigsVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// ISigsVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISigsVerifierTransactorRaw struct {
	Contract *ISigsVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISigsVerifier creates a new instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifier(address common.Address, backend bind.ContractBackend) (*ISigsVerifier, error) {
	contract, err := bindISigsVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifier{ISigsVerifierCaller: ISigsVerifierCaller{contract: contract}, ISigsVerifierTransactor: ISigsVerifierTransactor{contract: contract}, ISigsVerifierFilterer: ISigsVerifierFilterer{contract: contract}}, nil
}

// NewISigsVerifierCaller creates a new read-only instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifierCaller(address common.Address, caller bind.ContractCaller) (*ISigsVerifierCaller, error) {
	contract, err := bindISigsVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifierCaller{contract: contract}, nil
}

// NewISigsVerifierTransactor creates a new write-only instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*ISigsVerifierTransactor, error) {
	contract, err := bindISigsVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifierTransactor{contract: contract}, nil
}

// NewISigsVerifierFilterer creates a new log filterer instance of ISigsVerifier, bound to a specific deployed contract.
func NewISigsVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*ISigsVerifierFilterer, error) {
	contract, err := bindISigsVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISigsVerifierFilterer{contract: contract}, nil
}

// bindISigsVerifier binds a generic wrapper to an already deployed contract.
func bindISigsVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISigsVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISigsVerifier *ISigsVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISigsVerifier.Contract.ISigsVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISigsVerifier *ISigsVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.ISigsVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISigsVerifier *ISigsVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.ISigsVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISigsVerifier *ISigsVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISigsVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISigsVerifier *ISigsVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISigsVerifier *ISigsVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISigsVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ISigsVerifier *ISigsVerifierCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	var out []interface{}
	err := _ISigsVerifier.contract.Call(opts, &out, "verifySigs", _msg, _sigs, _signers, _powers)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ISigsVerifier *ISigsVerifierSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _ISigsVerifier.Contract.VerifySigs(&_ISigsVerifier.CallOpts, _msg, _sigs, _signers, _powers)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_ISigsVerifier *ISigsVerifierCallerSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _ISigsVerifier.Contract.VerifySigs(&_ISigsVerifier.CallOpts, _msg, _sigs, _signers, _powers)
}

// PauserMetaData contains all meta data concerning the Pauser contract.
var PauserMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PauserABI is the input ABI used to generate the binding from.
// Deprecated: Use PauserMetaData.ABI instead.
var PauserABI = PauserMetaData.ABI

// Pauser is an auto generated Go binding around an Ethereum contract.
type Pauser struct {
	PauserCaller     // Read-only binding to the contract
	PauserTransactor // Write-only binding to the contract
	PauserFilterer   // Log filterer for contract events
}

// PauserCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauserSession struct {
	Contract     *Pauser           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauserCallerSession struct {
	Contract *PauserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PauserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauserTransactorSession struct {
	Contract     *PauserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauserRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauserRaw struct {
	Contract *Pauser // Generic contract binding to access the raw methods on
}

// PauserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauserCallerRaw struct {
	Contract *PauserCaller // Generic read-only contract binding to access the raw methods on
}

// PauserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauserTransactorRaw struct {
	Contract *PauserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauser creates a new instance of Pauser, bound to a specific deployed contract.
func NewPauser(address common.Address, backend bind.ContractBackend) (*Pauser, error) {
	contract, err := bindPauser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pauser{PauserCaller: PauserCaller{contract: contract}, PauserTransactor: PauserTransactor{contract: contract}, PauserFilterer: PauserFilterer{contract: contract}}, nil
}

// NewPauserCaller creates a new read-only instance of Pauser, bound to a specific deployed contract.
func NewPauserCaller(address common.Address, caller bind.ContractCaller) (*PauserCaller, error) {
	contract, err := bindPauser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauserCaller{contract: contract}, nil
}

// NewPauserTransactor creates a new write-only instance of Pauser, bound to a specific deployed contract.
func NewPauserTransactor(address common.Address, transactor bind.ContractTransactor) (*PauserTransactor, error) {
	contract, err := bindPauser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauserTransactor{contract: contract}, nil
}

// NewPauserFilterer creates a new log filterer instance of Pauser, bound to a specific deployed contract.
func NewPauserFilterer(address common.Address, filterer bind.ContractFilterer) (*PauserFilterer, error) {
	contract, err := bindPauser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauserFilterer{contract: contract}, nil
}

// bindPauser binds a generic wrapper to an already deployed contract.
func bindPauser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PauserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pauser *PauserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pauser.Contract.PauserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pauser *PauserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pauser.Contract.PauserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pauser *PauserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pauser.Contract.PauserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pauser *PauserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pauser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pauser *PauserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pauser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pauser *PauserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pauser.Contract.contract.Transact(opts, method, params...)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Pauser *PauserCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pauser.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Pauser *PauserSession) IsPauser(account common.Address) (bool, error) {
	return _Pauser.Contract.IsPauser(&_Pauser.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Pauser *PauserCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Pauser.Contract.IsPauser(&_Pauser.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pauser *PauserCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pauser.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pauser *PauserSession) Owner() (common.Address, error) {
	return _Pauser.Contract.Owner(&_Pauser.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pauser *PauserCallerSession) Owner() (common.Address, error) {
	return _Pauser.Contract.Owner(&_Pauser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pauser *PauserCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pauser.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pauser *PauserSession) Paused() (bool, error) {
	return _Pauser.Contract.Paused(&_Pauser.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pauser *PauserCallerSession) Paused() (bool, error) {
	return _Pauser.Contract.Paused(&_Pauser.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Pauser *PauserCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Pauser.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Pauser *PauserSession) Pausers(arg0 common.Address) (bool, error) {
	return _Pauser.Contract.Pausers(&_Pauser.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Pauser *PauserCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Pauser.Contract.Pausers(&_Pauser.CallOpts, arg0)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Pauser *PauserTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Pauser.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Pauser *PauserSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Pauser.Contract.AddPauser(&_Pauser.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Pauser *PauserTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Pauser.Contract.AddPauser(&_Pauser.TransactOpts, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pauser *PauserTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pauser.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pauser *PauserSession) Pause() (*types.Transaction, error) {
	return _Pauser.Contract.Pause(&_Pauser.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pauser *PauserTransactorSession) Pause() (*types.Transaction, error) {
	return _Pauser.Contract.Pause(&_Pauser.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Pauser *PauserTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Pauser.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Pauser *PauserSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Pauser.Contract.RemovePauser(&_Pauser.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Pauser *PauserTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Pauser.Contract.RemovePauser(&_Pauser.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pauser *PauserTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pauser.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pauser *PauserSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pauser.Contract.RenounceOwnership(&_Pauser.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pauser *PauserTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pauser.Contract.RenounceOwnership(&_Pauser.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pauser *PauserTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pauser.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pauser *PauserSession) RenouncePauser() (*types.Transaction, error) {
	return _Pauser.Contract.RenouncePauser(&_Pauser.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Pauser *PauserTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Pauser.Contract.RenouncePauser(&_Pauser.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pauser *PauserTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pauser.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pauser *PauserSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pauser.Contract.TransferOwnership(&_Pauser.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pauser *PauserTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pauser.Contract.TransferOwnership(&_Pauser.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pauser *PauserTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pauser.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pauser *PauserSession) Unpause() (*types.Transaction, error) {
	return _Pauser.Contract.Unpause(&_Pauser.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pauser *PauserTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pauser.Contract.Unpause(&_Pauser.TransactOpts)
}

// PauserOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pauser contract.
type PauserOwnershipTransferredIterator struct {
	Event *PauserOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PauserOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserOwnershipTransferred)
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
		it.Event = new(PauserOwnershipTransferred)
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
func (it *PauserOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserOwnershipTransferred represents a OwnershipTransferred event raised by the Pauser contract.
type PauserOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pauser *PauserFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PauserOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pauser.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PauserOwnershipTransferredIterator{contract: _Pauser.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pauser *PauserFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PauserOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pauser.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserOwnershipTransferred)
				if err := _Pauser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pauser *PauserFilterer) ParseOwnershipTransferred(log types.Log) (*PauserOwnershipTransferred, error) {
	event := new(PauserOwnershipTransferred)
	if err := _Pauser.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauserPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pauser contract.
type PauserPausedIterator struct {
	Event *PauserPaused // Event containing the contract specifics and raw log

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
func (it *PauserPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserPaused)
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
		it.Event = new(PauserPaused)
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
func (it *PauserPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserPaused represents a Paused event raised by the Pauser contract.
type PauserPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pauser *PauserFilterer) FilterPaused(opts *bind.FilterOpts) (*PauserPausedIterator, error) {

	logs, sub, err := _Pauser.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PauserPausedIterator{contract: _Pauser.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pauser *PauserFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PauserPaused) (event.Subscription, error) {

	logs, sub, err := _Pauser.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserPaused)
				if err := _Pauser.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pauser *PauserFilterer) ParsePaused(log types.Log) (*PauserPaused, error) {
	event := new(PauserPaused)
	if err := _Pauser.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauserPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Pauser contract.
type PauserPauserAddedIterator struct {
	Event *PauserPauserAdded // Event containing the contract specifics and raw log

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
func (it *PauserPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserPauserAdded)
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
		it.Event = new(PauserPauserAdded)
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
func (it *PauserPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserPauserAdded represents a PauserAdded event raised by the Pauser contract.
type PauserPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Pauser *PauserFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*PauserPauserAddedIterator, error) {

	logs, sub, err := _Pauser.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &PauserPauserAddedIterator{contract: _Pauser.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Pauser *PauserFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *PauserPauserAdded) (event.Subscription, error) {

	logs, sub, err := _Pauser.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserPauserAdded)
				if err := _Pauser.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Pauser *PauserFilterer) ParsePauserAdded(log types.Log) (*PauserPauserAdded, error) {
	event := new(PauserPauserAdded)
	if err := _Pauser.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauserPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Pauser contract.
type PauserPauserRemovedIterator struct {
	Event *PauserPauserRemoved // Event containing the contract specifics and raw log

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
func (it *PauserPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserPauserRemoved)
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
		it.Event = new(PauserPauserRemoved)
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
func (it *PauserPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserPauserRemoved represents a PauserRemoved event raised by the Pauser contract.
type PauserPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Pauser *PauserFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*PauserPauserRemovedIterator, error) {

	logs, sub, err := _Pauser.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &PauserPauserRemovedIterator{contract: _Pauser.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Pauser *PauserFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *PauserPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Pauser.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserPauserRemoved)
				if err := _Pauser.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Pauser *PauserFilterer) ParsePauserRemoved(log types.Log) (*PauserPauserRemoved, error) {
	event := new(PauserPauserRemoved)
	if err := _Pauser.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauserUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pauser contract.
type PauserUnpausedIterator struct {
	Event *PauserUnpaused // Event containing the contract specifics and raw log

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
func (it *PauserUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauserUnpaused)
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
		it.Event = new(PauserUnpaused)
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
func (it *PauserUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauserUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauserUnpaused represents a Unpaused event raised by the Pauser contract.
type PauserUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pauser *PauserFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PauserUnpausedIterator, error) {

	logs, sub, err := _Pauser.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PauserUnpausedIterator{contract: _Pauser.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pauser *PauserFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PauserUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pauser.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauserUnpaused)
				if err := _Pauser.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pauser *PauserFilterer) ParseUnpaused(log types.Log) (*PauserUnpaused, error) {
	event := new(PauserUnpaused)
	if err := _Pauser.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PbMetaData contains all meta data concerning the Pb contract.
var PbMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201984af984727ce1952a7ba24449ec8d204dd0db8bc1b71bf9b6abb9f2955fe5464736f6c63430008090033",
}

// PbABI is the input ABI used to generate the binding from.
// Deprecated: Use PbMetaData.ABI instead.
var PbABI = PbMetaData.ABI

// PbBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PbMetaData.Bin instead.
var PbBin = PbMetaData.Bin

// DeployPb deploys a new Ethereum contract, binding an instance of Pb to it.
func DeployPb(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pb, error) {
	parsed, err := PbMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PbBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// Pb is an auto generated Go binding around an Ethereum contract.
type Pb struct {
	PbCaller     // Read-only binding to the contract
	PbTransactor // Write-only binding to the contract
	PbFilterer   // Log filterer for contract events
}

// PbCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbSession struct {
	Contract     *Pb               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbCallerSession struct {
	Contract *PbCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbTransactorSession struct {
	Contract     *PbTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbRaw struct {
	Contract *Pb // Generic contract binding to access the raw methods on
}

// PbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbCallerRaw struct {
	Contract *PbCaller // Generic read-only contract binding to access the raw methods on
}

// PbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbTransactorRaw struct {
	Contract *PbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPb creates a new instance of Pb, bound to a specific deployed contract.
func NewPb(address common.Address, backend bind.ContractBackend) (*Pb, error) {
	contract, err := bindPb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pb{PbCaller: PbCaller{contract: contract}, PbTransactor: PbTransactor{contract: contract}, PbFilterer: PbFilterer{contract: contract}}, nil
}

// NewPbCaller creates a new read-only instance of Pb, bound to a specific deployed contract.
func NewPbCaller(address common.Address, caller bind.ContractCaller) (*PbCaller, error) {
	contract, err := bindPb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbCaller{contract: contract}, nil
}

// NewPbTransactor creates a new write-only instance of Pb, bound to a specific deployed contract.
func NewPbTransactor(address common.Address, transactor bind.ContractTransactor) (*PbTransactor, error) {
	contract, err := bindPb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbTransactor{contract: contract}, nil
}

// NewPbFilterer creates a new log filterer instance of Pb, bound to a specific deployed contract.
func NewPbFilterer(address common.Address, filterer bind.ContractFilterer) (*PbFilterer, error) {
	contract, err := bindPb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbFilterer{contract: contract}, nil
}

// bindPb binds a generic wrapper to an already deployed contract.
func bindPb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.PbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.PbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pb *PbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pb *PbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pb *PbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pb.Contract.contract.Transact(opts, method, params...)
}

// PbFarmingMetaData contains all meta data concerning the PbFarming contract.
var PbFarmingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c7e14f59ced87e58a12a97a843a19631ff5c5c834644f77fedb6792066a7d31564736f6c63430008090033",
}

// PbFarmingABI is the input ABI used to generate the binding from.
// Deprecated: Use PbFarmingMetaData.ABI instead.
var PbFarmingABI = PbFarmingMetaData.ABI

// PbFarmingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PbFarmingMetaData.Bin instead.
var PbFarmingBin = PbFarmingMetaData.Bin

// DeployPbFarming deploys a new Ethereum contract, binding an instance of PbFarming to it.
func DeployPbFarming(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbFarming, error) {
	parsed, err := PbFarmingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PbFarmingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbFarming{PbFarmingCaller: PbFarmingCaller{contract: contract}, PbFarmingTransactor: PbFarmingTransactor{contract: contract}, PbFarmingFilterer: PbFarmingFilterer{contract: contract}}, nil
}

// PbFarming is an auto generated Go binding around an Ethereum contract.
type PbFarming struct {
	PbFarmingCaller     // Read-only binding to the contract
	PbFarmingTransactor // Write-only binding to the contract
	PbFarmingFilterer   // Log filterer for contract events
}

// PbFarmingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbFarmingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbFarmingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbFarmingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbFarmingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbFarmingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbFarmingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbFarmingSession struct {
	Contract     *PbFarming        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbFarmingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbFarmingCallerSession struct {
	Contract *PbFarmingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PbFarmingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbFarmingTransactorSession struct {
	Contract     *PbFarmingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PbFarmingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbFarmingRaw struct {
	Contract *PbFarming // Generic contract binding to access the raw methods on
}

// PbFarmingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbFarmingCallerRaw struct {
	Contract *PbFarmingCaller // Generic read-only contract binding to access the raw methods on
}

// PbFarmingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbFarmingTransactorRaw struct {
	Contract *PbFarmingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbFarming creates a new instance of PbFarming, bound to a specific deployed contract.
func NewPbFarming(address common.Address, backend bind.ContractBackend) (*PbFarming, error) {
	contract, err := bindPbFarming(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbFarming{PbFarmingCaller: PbFarmingCaller{contract: contract}, PbFarmingTransactor: PbFarmingTransactor{contract: contract}, PbFarmingFilterer: PbFarmingFilterer{contract: contract}}, nil
}

// NewPbFarmingCaller creates a new read-only instance of PbFarming, bound to a specific deployed contract.
func NewPbFarmingCaller(address common.Address, caller bind.ContractCaller) (*PbFarmingCaller, error) {
	contract, err := bindPbFarming(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbFarmingCaller{contract: contract}, nil
}

// NewPbFarmingTransactor creates a new write-only instance of PbFarming, bound to a specific deployed contract.
func NewPbFarmingTransactor(address common.Address, transactor bind.ContractTransactor) (*PbFarmingTransactor, error) {
	contract, err := bindPbFarming(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbFarmingTransactor{contract: contract}, nil
}

// NewPbFarmingFilterer creates a new log filterer instance of PbFarming, bound to a specific deployed contract.
func NewPbFarmingFilterer(address common.Address, filterer bind.ContractFilterer) (*PbFarmingFilterer, error) {
	contract, err := bindPbFarming(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbFarmingFilterer{contract: contract}, nil
}

// bindPbFarming binds a generic wrapper to an already deployed contract.
func bindPbFarming(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbFarmingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbFarming *PbFarmingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbFarming.Contract.PbFarmingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbFarming *PbFarmingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbFarming.Contract.PbFarmingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbFarming *PbFarmingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbFarming.Contract.PbFarmingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbFarming *PbFarmingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbFarming.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbFarming *PbFarmingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbFarming.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbFarming *PbFarmingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbFarming.Contract.contract.Transact(opts, method, params...)
}

// PbSgnMetaData contains all meta data concerning the PbSgn contract.
var PbSgnMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220629aad710fbb1a68c12181855483ec9c03c45727ffcbeb8995f154c0e5e9cf2f64736f6c63430008090033",
}

// PbSgnABI is the input ABI used to generate the binding from.
// Deprecated: Use PbSgnMetaData.ABI instead.
var PbSgnABI = PbSgnMetaData.ABI

// PbSgnBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PbSgnMetaData.Bin instead.
var PbSgnBin = PbSgnMetaData.Bin

// DeployPbSgn deploys a new Ethereum contract, binding an instance of PbSgn to it.
func DeployPbSgn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbSgn, error) {
	parsed, err := PbSgnMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PbSgnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbSgn{PbSgnCaller: PbSgnCaller{contract: contract}, PbSgnTransactor: PbSgnTransactor{contract: contract}, PbSgnFilterer: PbSgnFilterer{contract: contract}}, nil
}

// PbSgn is an auto generated Go binding around an Ethereum contract.
type PbSgn struct {
	PbSgnCaller     // Read-only binding to the contract
	PbSgnTransactor // Write-only binding to the contract
	PbSgnFilterer   // Log filterer for contract events
}

// PbSgnCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbSgnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSgnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbSgnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSgnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbSgnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbSgnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbSgnSession struct {
	Contract     *PbSgn            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbSgnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbSgnCallerSession struct {
	Contract *PbSgnCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PbSgnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbSgnTransactorSession struct {
	Contract     *PbSgnTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbSgnRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbSgnRaw struct {
	Contract *PbSgn // Generic contract binding to access the raw methods on
}

// PbSgnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbSgnCallerRaw struct {
	Contract *PbSgnCaller // Generic read-only contract binding to access the raw methods on
}

// PbSgnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbSgnTransactorRaw struct {
	Contract *PbSgnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbSgn creates a new instance of PbSgn, bound to a specific deployed contract.
func NewPbSgn(address common.Address, backend bind.ContractBackend) (*PbSgn, error) {
	contract, err := bindPbSgn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbSgn{PbSgnCaller: PbSgnCaller{contract: contract}, PbSgnTransactor: PbSgnTransactor{contract: contract}, PbSgnFilterer: PbSgnFilterer{contract: contract}}, nil
}

// NewPbSgnCaller creates a new read-only instance of PbSgn, bound to a specific deployed contract.
func NewPbSgnCaller(address common.Address, caller bind.ContractCaller) (*PbSgnCaller, error) {
	contract, err := bindPbSgn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbSgnCaller{contract: contract}, nil
}

// NewPbSgnTransactor creates a new write-only instance of PbSgn, bound to a specific deployed contract.
func NewPbSgnTransactor(address common.Address, transactor bind.ContractTransactor) (*PbSgnTransactor, error) {
	contract, err := bindPbSgn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbSgnTransactor{contract: contract}, nil
}

// NewPbSgnFilterer creates a new log filterer instance of PbSgn, bound to a specific deployed contract.
func NewPbSgnFilterer(address common.Address, filterer bind.ContractFilterer) (*PbSgnFilterer, error) {
	contract, err := bindPbSgn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbSgnFilterer{contract: contract}, nil
}

// bindPbSgn binds a generic wrapper to an already deployed contract.
func bindPbSgn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbSgnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbSgn *PbSgnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbSgn.Contract.PbSgnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbSgn *PbSgnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbSgn.Contract.PbSgnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbSgn *PbSgnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbSgn.Contract.PbSgnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbSgn *PbSgnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbSgn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbSgn *PbSgnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbSgn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbSgn *PbSgnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbSgn.Contract.contract.Transact(opts, method, params...)
}

// PbStakingMetaData contains all meta data concerning the PbStaking contract.
var PbStakingMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200cfd62c2c5d8729c76b07936ec6f1c4b62c8773121d0558a488ca84fb9519de264736f6c63430008090033",
}

// PbStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use PbStakingMetaData.ABI instead.
var PbStakingABI = PbStakingMetaData.ABI

// PbStakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PbStakingMetaData.Bin instead.
var PbStakingBin = PbStakingMetaData.Bin

// DeployPbStaking deploys a new Ethereum contract, binding an instance of PbStaking to it.
func DeployPbStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PbStaking, error) {
	parsed, err := PbStakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PbStakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PbStaking{PbStakingCaller: PbStakingCaller{contract: contract}, PbStakingTransactor: PbStakingTransactor{contract: contract}, PbStakingFilterer: PbStakingFilterer{contract: contract}}, nil
}

// PbStaking is an auto generated Go binding around an Ethereum contract.
type PbStaking struct {
	PbStakingCaller     // Read-only binding to the contract
	PbStakingTransactor // Write-only binding to the contract
	PbStakingFilterer   // Log filterer for contract events
}

// PbStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PbStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PbStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PbStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PbStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PbStakingSession struct {
	Contract     *PbStaking        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PbStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PbStakingCallerSession struct {
	Contract *PbStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PbStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PbStakingTransactorSession struct {
	Contract     *PbStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PbStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PbStakingRaw struct {
	Contract *PbStaking // Generic contract binding to access the raw methods on
}

// PbStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PbStakingCallerRaw struct {
	Contract *PbStakingCaller // Generic read-only contract binding to access the raw methods on
}

// PbStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PbStakingTransactorRaw struct {
	Contract *PbStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPbStaking creates a new instance of PbStaking, bound to a specific deployed contract.
func NewPbStaking(address common.Address, backend bind.ContractBackend) (*PbStaking, error) {
	contract, err := bindPbStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PbStaking{PbStakingCaller: PbStakingCaller{contract: contract}, PbStakingTransactor: PbStakingTransactor{contract: contract}, PbStakingFilterer: PbStakingFilterer{contract: contract}}, nil
}

// NewPbStakingCaller creates a new read-only instance of PbStaking, bound to a specific deployed contract.
func NewPbStakingCaller(address common.Address, caller bind.ContractCaller) (*PbStakingCaller, error) {
	contract, err := bindPbStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PbStakingCaller{contract: contract}, nil
}

// NewPbStakingTransactor creates a new write-only instance of PbStaking, bound to a specific deployed contract.
func NewPbStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*PbStakingTransactor, error) {
	contract, err := bindPbStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PbStakingTransactor{contract: contract}, nil
}

// NewPbStakingFilterer creates a new log filterer instance of PbStaking, bound to a specific deployed contract.
func NewPbStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*PbStakingFilterer, error) {
	contract, err := bindPbStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PbStakingFilterer{contract: contract}, nil
}

// bindPbStaking binds a generic wrapper to an already deployed contract.
func bindPbStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PbStakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbStaking *PbStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbStaking.Contract.PbStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbStaking *PbStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbStaking.Contract.PbStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbStaking *PbStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbStaking.Contract.PbStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PbStaking *PbStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PbStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PbStaking *PbStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PbStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PbStaking *PbStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PbStaking.Contract.contract.Transact(opts, method, params...)
}

// SGNMetaData contains all meta data concerning the SGN contract.
var SGNMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddr\",\"type\":\"bytes\"}],\"name\":\"SgnAddrUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sgnAddrs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sgnAddr\",\"type\":\"bytes\"}],\"name\":\"updateSgnAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_withdrawalRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnAmts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200209f3803806200209f833981016040819052620000349162000182565b6200003f3362000069565b6000805460ff60a01b191690556200005733620000b9565b6001600160a01b0316608052620001b4565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff1615620001275760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640160405180910390fd5b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910160405180910390a150565b6000602082840312156200019557600080fd5b81516001600160a01b0381168114620001ad57600080fd5b9392505050565b608051611eac620001f3600039600081816101b90152818161081d015281816108ba0152818161096101528181610b200152610c3b0152611eac6000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c806380f51c12116100cd578063b02c43d011610081578063d0bb935111610066578063d0bb9351146102f6578063d88ef27114610309578063f2fde38b1461031c57600080fd5b8063b02c43d0146102c3578063c429fe1f146102d657600080fd5b80638456cb59116100b25780638456cb59146102975780638da5cb5b1461029f5780639d4323be146102b057600080fd5b806380f51c121461026157806382dc1ec41461028457600080fd5b80635c975abb116101245780636ef8d66d116101095780636ef8d66d14610218578063715018a614610220578063795c2c141461022857600080fd5b80635c975abb146101f35780636b2c0f551461020557600080fd5b80633f4ba83a1461015657806346fbf68e1461016057806347e7ef24146101a15780634cf088d9146101b4575b600080fd5b61015e61032f565b005b61018c61016e36600461195d565b6001600160a01b031660009081526001602052604090205460ff1690565b60405190151581526020015b60405180910390f35b61015e6101af36600461197a565b61039d565b6101db7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610198565b600054600160a01b900460ff1661018c565b61015e61021336600461195d565b6104d7565b61015e61053d565b61015e610546565b6102536102363660046119a6565b600360209081526000928352604080842090915290825290205481565b604051908152602001610198565b61018c61026f36600461195d565b60016020526000908152604090205460ff1681565b61015e61029236600461195d565b6105aa565b61015e61060d565b6000546001600160a01b03166101db565b61015e6102be36600461197a565b610674565b6102536102d13660046119df565b61073f565b6102e96102e436600461195d565b610760565b6040516101989190611a50565b61015e610304366004611aac565b6107fa565b61015e610317366004611aee565b610bd7565b61015e61032a36600461195d565b610e3d565b3360009081526001602052604090205460ff166103935760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064015b60405180910390fd5b61039b610f1c565b565b600054600160a01b900460ff16156103ea5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161038a565b6040516bffffffffffffffffffffffff1933606081811b8316602085015285901b9091166034830152604882018390529060029060680160408051601f198184030181529190528051602091820120825460018101845560009384529190922001556104616001600160a01b038416823085610fc2565b60025460009061047390600190611b9e565b6040805167ffffffffffffffff831681526001600160a01b0385811660208301528716818301526060810186905290519192507f2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8919081900360800190a150505050565b6000546001600160a01b031633146105315760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038a565b61053a81611060565b50565b61039b33611060565b6000546001600160a01b031633146105a05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038a565b61039b6000611120565b6000546001600160a01b031633146106045760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038a565b61053a81611188565b3360009081526001602052604090205460ff1661066c5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f7420706175736572000000000000000000000000604482015260640161038a565b61039b611246565b600054600160a01b900460ff166106cd5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161038a565b6000546001600160a01b031633146107275760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038a565b61073b6001600160a01b03831633836112ce565b5050565b6002818154811061074f57600080fd5b600091825260209091200154905081565b6004602052600090815260409020805461077990611bb5565b80601f01602080910402602001604051908101604052809291908181526020018280546107a590611bb5565b80156107f25780601f106107c7576101008083540402835291602001916107f2565b820191906000526020600020905b8154815290600101906020018083116107d557829003601f168201915b505050505081565b604051636d30878360e01b81523360048201819052906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636d3087839060240160206040518083038186803b15801561085f57600080fd5b505afa158015610873573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108979190611bf0565b6001600160a01b03161461093f57604051636d30878360e01b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636d3087839060240160206040518083038186803b15801561090457600080fd5b505afa158015610918573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093c9190611bf0565b90505b60405163a310624f60e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a310624f9060240160206040518083038186803b1580156109a557600080fd5b505afa1580156109b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109dd9190611c0d565b905060018160038111156109f3576109f3611c2e565b14610a405760405162461bcd60e51b815260206004820152601660248201527f4e6f7420756e626f6e6465642076616c696461746f7200000000000000000000604482015260640161038a565b6001600160a01b03821660009081526004602052604081208054610a6390611bb5565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8f90611bb5565b8015610adc5780601f10610ab157610100808354040283529160200191610adc565b820191906000526020600020905b815481529060010190602001808311610abf57829003601f168201915b505050506001600160a01b0385166000908152600460205260409020919250610b0891905086866118af565b506040516309146f1160e41b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639146f11090610b5990869089908990600401611c6d565b600060405180830381600087803b158015610b7357600080fd5b505af1158015610b87573d6000803e3d6000fd5b50505050826001600160a01b03167f8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4828787604051610bc893929190611ccd565b60405180910390a25050505050565b600054600160a01b900460ff1615610c245760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161038a565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe90610c76908790879087908790600401611cfd565b60206040518083038186803b158015610c8e57600080fd5b505afa158015610ca2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cc69190611dae565b506000610d0885858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061130392505050565b80516001600160a01b039081166000908152600360209081526040808320828601519094168352929052818120549183015192935091610d489190611b9e565b905060008111610d9a5760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f20776974686472617700000000000000604482015260640161038a565b60408083015183516001600160a01b0390811660009081526003602090815284822081880180518516845291529390209190915583519151610ddf92911690836112ce565b8151602080840151604080516001600160a01b039485168152939091169183019190915281018290527f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9060600160405180910390a1505050505050565b6000546001600160a01b03163314610e975760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161038a565b6001600160a01b038116610f135760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161038a565b61053a81611120565b600054600160a01b900460ff16610f755760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161038a565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040516001600160a01b038085166024830152831660448201526064810182905261105a9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526113de565b50505050565b6001600160a01b03811660009081526001602052604090205460ff166110c85760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f74207061757365720000000000000000000000604482015260640161038a565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e91015b60405180910390a150565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff16156111f15760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640161038a565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101611115565b600054600160a01b900460ff16156112935760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161038a565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610fa53390565b6040516001600160a01b0383166024820152604481018290526112fe90849063a9059cbb60e01b90606401610ff6565b505050565b604080516060810182526000808252602080830182905282840182905283518085019094528184528301849052909190805b602083015151835110156113d65761134c836114c3565b9092509050816001141561137b5761136b611366846114fd565b6115ba565b6001600160a01b03168452611335565b81600214156113a357611390611366846114fd565b6001600160a01b03166020850152611335565b81600314156113c7576113bd6113b8846114fd565b6115cb565b6040850152611335565b6113d18382611602565b611335565b505050919050565b6000611433826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116749092919063ffffffff16565b8051909150156112fe57808060200190518101906114519190611dae565b6112fe5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161038a565b60008060006114d18461168d565b90506114de600882611dd0565b92508060071660058111156114f5576114f5611c2e565b915050915091565b6060600061150a8361168d565b9050600081846000015161151e9190611df2565b905083602001515181111561153257600080fd5b8167ffffffffffffffff81111561154b5761154b611e0a565b6040519080825280601f01601f191660200182016040528015611575576020820181803683370190505b50602080860151865192955091818601919083010160005b858110156115af5781810151838201526115a8602082611df2565b905061158d565b505050935250919050565b60006115c58261170f565b92915050565b60006020825111156115dc57600080fd5b60208201519050815160206115f19190611b9e565b6115fc906008611e20565b1c919050565b600081600581111561161657611616611c2e565b1415611625576112fe8261168d565b600281600581111561163957611639611c2e565b141561015157600061164a8361168d565b9050808360000181815161165e9190611df2565b905250602083015151835111156112fe57600080fd5b60606116838484600085611737565b90505b9392505050565b602080820151825181019091015160009182805b600a8110156117095783811a91506116ba816007611e20565b82607f16901b8517945081608016600014156116f7576116db816001611df2565b865187906116ea908390611df2565b9052509395945050505050565b8061170181611e3f565b9150506116a1565b50600080fd5b6000815160141461171f57600080fd5b50602001516c01000000000000000000000000900490565b6060824710156117af5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161038a565b843b6117fd5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161038a565b600080866001600160a01b031685876040516118199190611e5a565b60006040518083038185875af1925050503d8060008114611856576040519150601f19603f3d011682016040523d82523d6000602084013e61185b565b606091505b509150915061186b828286611876565b979650505050505050565b60608315611885575081611686565b8251156118955782518084602001fd5b8160405162461bcd60e51b815260040161038a9190611a50565b8280546118bb90611bb5565b90600052602060002090601f0160209004810192826118dd5760008555611923565b82601f106118f65782800160ff19823516178555611923565b82800160010185558215611923579182015b82811115611923578235825591602001919060010190611908565b5061192f929150611933565b5090565b5b8082111561192f5760008155600101611934565b6001600160a01b038116811461053a57600080fd5b60006020828403121561196f57600080fd5b813561168681611948565b6000806040838503121561198d57600080fd5b823561199881611948565b946020939093013593505050565b600080604083850312156119b957600080fd5b82356119c481611948565b915060208301356119d481611948565b809150509250929050565b6000602082840312156119f157600080fd5b5035919050565b60005b83811015611a135781810151838201526020016119fb565b8381111561105a5750506000910152565b60008151808452611a3c8160208601602086016119f8565b601f01601f19169290920160200192915050565b6020815260006116866020830184611a24565b60008083601f840112611a7557600080fd5b50813567ffffffffffffffff811115611a8d57600080fd5b602083019150836020828501011115611aa557600080fd5b9250929050565b60008060208385031215611abf57600080fd5b823567ffffffffffffffff811115611ad657600080fd5b611ae285828601611a63565b90969095509350505050565b60008060008060408587031215611b0457600080fd5b843567ffffffffffffffff80821115611b1c57600080fd5b611b2888838901611a63565b90965094506020870135915080821115611b4157600080fd5b818701915087601f830112611b5557600080fd5b813581811115611b6457600080fd5b8860208260051b8501011115611b7957600080fd5b95989497505060200194505050565b634e487b7160e01b600052601160045260246000fd5b600082821015611bb057611bb0611b88565b500390565b600181811c90821680611bc957607f821691505b60208210811415611bea57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215611c0257600080fd5b815161168681611948565b600060208284031215611c1f57600080fd5b81516004811061168657600080fd5b634e487b7160e01b600052602160045260246000fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b038416815260606020820152600860608201527f73676e2d61646472000000000000000000000000000000000000000000000000608082015260a060408201526000611cc460a083018486611c44565b95945050505050565b604081526000611ce06040830186611a24565b8281036020840152611cf3818587611c44565b9695505050505050565b604081526000611d11604083018688611c44565b602083820381850152818583528183019050818660051b8401018760005b88811015611d9e57858303601f190184528135368b9003601e19018112611d5557600080fd5b8a01803567ffffffffffffffff811115611d6e57600080fd5b8036038c1315611d7d57600080fd5b611d8a8582898501611c44565b958701959450505090840190600101611d2f565b50909a9950505050505050505050565b600060208284031215611dc057600080fd5b8151801515811461168657600080fd5b600082611ded57634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115611e0557611e05611b88565b500190565b634e487b7160e01b600052604160045260246000fd5b6000816000190483118215151615611e3a57611e3a611b88565b500290565b6000600019821415611e5357611e53611b88565b5060010190565b60008251611e6c8184602087016119f8565b919091019291505056fea2646970667358221220b13720f4feb1e8b4755d0653163b776bac046e5f7c106c6c94af3a05fc265f4364736f6c63430008090033",
}

// SGNABI is the input ABI used to generate the binding from.
// Deprecated: Use SGNMetaData.ABI instead.
var SGNABI = SGNMetaData.ABI

// SGNBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SGNMetaData.Bin instead.
var SGNBin = SGNMetaData.Bin

// DeploySGN deploys a new Ethereum contract, binding an instance of SGN to it.
func DeploySGN(auth *bind.TransactOpts, backend bind.ContractBackend, _staking common.Address) (common.Address, *types.Transaction, *SGN, error) {
	parsed, err := SGNMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SGNBin), backend, _staking)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SGN{SGNCaller: SGNCaller{contract: contract}, SGNTransactor: SGNTransactor{contract: contract}, SGNFilterer: SGNFilterer{contract: contract}}, nil
}

// SGN is an auto generated Go binding around an Ethereum contract.
type SGN struct {
	SGNCaller     // Read-only binding to the contract
	SGNTransactor // Write-only binding to the contract
	SGNFilterer   // Log filterer for contract events
}

// SGNCaller is an auto generated read-only Go binding around an Ethereum contract.
type SGNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SGNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGNFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SGNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SGNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SGNSession struct {
	Contract     *SGN              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SGNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SGNCallerSession struct {
	Contract *SGNCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SGNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SGNTransactorSession struct {
	Contract     *SGNTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SGNRaw is an auto generated low-level Go binding around an Ethereum contract.
type SGNRaw struct {
	Contract *SGN // Generic contract binding to access the raw methods on
}

// SGNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SGNCallerRaw struct {
	Contract *SGNCaller // Generic read-only contract binding to access the raw methods on
}

// SGNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SGNTransactorRaw struct {
	Contract *SGNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSGN creates a new instance of SGN, bound to a specific deployed contract.
func NewSGN(address common.Address, backend bind.ContractBackend) (*SGN, error) {
	contract, err := bindSGN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SGN{SGNCaller: SGNCaller{contract: contract}, SGNTransactor: SGNTransactor{contract: contract}, SGNFilterer: SGNFilterer{contract: contract}}, nil
}

// NewSGNCaller creates a new read-only instance of SGN, bound to a specific deployed contract.
func NewSGNCaller(address common.Address, caller bind.ContractCaller) (*SGNCaller, error) {
	contract, err := bindSGN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SGNCaller{contract: contract}, nil
}

// NewSGNTransactor creates a new write-only instance of SGN, bound to a specific deployed contract.
func NewSGNTransactor(address common.Address, transactor bind.ContractTransactor) (*SGNTransactor, error) {
	contract, err := bindSGN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SGNTransactor{contract: contract}, nil
}

// NewSGNFilterer creates a new log filterer instance of SGN, bound to a specific deployed contract.
func NewSGNFilterer(address common.Address, filterer bind.ContractFilterer) (*SGNFilterer, error) {
	contract, err := bindSGN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SGNFilterer{contract: contract}, nil
}

// bindSGN binds a generic wrapper to an already deployed contract.
func bindSGN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SGNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGN *SGNRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGN.Contract.SGNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGN *SGNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.Contract.SGNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGN *SGNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGN.Contract.SGNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SGN *SGNCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SGN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SGN *SGNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SGN *SGNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SGN.Contract.contract.Transact(opts, method, params...)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_SGN *SGNCaller) Deposits(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "deposits", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_SGN *SGNSession) Deposits(arg0 *big.Int) ([32]byte, error) {
	return _SGN.Contract.Deposits(&_SGN.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xb02c43d0.
//
// Solidity: function deposits(uint256 ) view returns(bytes32)
func (_SGN *SGNCallerSession) Deposits(arg0 *big.Int) ([32]byte, error) {
	return _SGN.Contract.Deposits(&_SGN.CallOpts, arg0)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_SGN *SGNCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_SGN *SGNSession) IsPauser(account common.Address) (bool, error) {
	return _SGN.Contract.IsPauser(&_SGN.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_SGN *SGNCallerSession) IsPauser(account common.Address) (bool, error) {
	return _SGN.Contract.IsPauser(&_SGN.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGN *SGNCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGN *SGNSession) Owner() (common.Address, error) {
	return _SGN.Contract.Owner(&_SGN.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SGN *SGNCallerSession) Owner() (common.Address, error) {
	return _SGN.Contract.Owner(&_SGN.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGN *SGNCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGN *SGNSession) Paused() (bool, error) {
	return _SGN.Contract.Paused(&_SGN.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SGN *SGNCallerSession) Paused() (bool, error) {
	return _SGN.Contract.Paused(&_SGN.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_SGN *SGNCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_SGN *SGNSession) Pausers(arg0 common.Address) (bool, error) {
	return _SGN.Contract.Pausers(&_SGN.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_SGN *SGNCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _SGN.Contract.Pausers(&_SGN.CallOpts, arg0)
}

// SgnAddrs is a free data retrieval call binding the contract method 0xc429fe1f.
//
// Solidity: function sgnAddrs(address ) view returns(bytes)
func (_SGN *SGNCaller) SgnAddrs(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "sgnAddrs", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SgnAddrs is a free data retrieval call binding the contract method 0xc429fe1f.
//
// Solidity: function sgnAddrs(address ) view returns(bytes)
func (_SGN *SGNSession) SgnAddrs(arg0 common.Address) ([]byte, error) {
	return _SGN.Contract.SgnAddrs(&_SGN.CallOpts, arg0)
}

// SgnAddrs is a free data retrieval call binding the contract method 0xc429fe1f.
//
// Solidity: function sgnAddrs(address ) view returns(bytes)
func (_SGN *SGNCallerSession) SgnAddrs(arg0 common.Address) ([]byte, error) {
	return _SGN.Contract.SgnAddrs(&_SGN.CallOpts, arg0)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_SGN *SGNCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_SGN *SGNSession) Staking() (common.Address, error) {
	return _SGN.Contract.Staking(&_SGN.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_SGN *SGNCallerSession) Staking() (common.Address, error) {
	return _SGN.Contract.Staking(&_SGN.CallOpts)
}

// WithdrawnAmts is a free data retrieval call binding the contract method 0x795c2c14.
//
// Solidity: function withdrawnAmts(address , address ) view returns(uint256)
func (_SGN *SGNCaller) WithdrawnAmts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SGN.contract.Call(opts, &out, "withdrawnAmts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnAmts is a free data retrieval call binding the contract method 0x795c2c14.
//
// Solidity: function withdrawnAmts(address , address ) view returns(uint256)
func (_SGN *SGNSession) WithdrawnAmts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SGN.Contract.WithdrawnAmts(&_SGN.CallOpts, arg0, arg1)
}

// WithdrawnAmts is a free data retrieval call binding the contract method 0x795c2c14.
//
// Solidity: function withdrawnAmts(address , address ) view returns(uint256)
func (_SGN *SGNCallerSession) WithdrawnAmts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SGN.Contract.WithdrawnAmts(&_SGN.CallOpts, arg0, arg1)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_SGN *SGNTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_SGN *SGNSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _SGN.Contract.AddPauser(&_SGN.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_SGN *SGNTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _SGN.Contract.AddPauser(&_SGN.TransactOpts, account)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_SGN *SGNSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.Deposit(&_SGN.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.Deposit(&_SGN.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactor) DrainToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "drainToken", _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_SGN *SGNSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.DrainToken(&_SGN.TransactOpts, _token, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x9d4323be.
//
// Solidity: function drainToken(address _token, uint256 _amount) returns()
func (_SGN *SGNTransactorSession) DrainToken(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SGN.Contract.DrainToken(&_SGN.TransactOpts, _token, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SGN *SGNTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SGN *SGNSession) Pause() (*types.Transaction, error) {
	return _SGN.Contract.Pause(&_SGN.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SGN *SGNTransactorSession) Pause() (*types.Transaction, error) {
	return _SGN.Contract.Pause(&_SGN.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_SGN *SGNTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_SGN *SGNSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _SGN.Contract.RemovePauser(&_SGN.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_SGN *SGNTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _SGN.Contract.RemovePauser(&_SGN.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGN *SGNTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGN *SGNSession) RenounceOwnership() (*types.Transaction, error) {
	return _SGN.Contract.RenounceOwnership(&_SGN.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SGN *SGNTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SGN.Contract.RenounceOwnership(&_SGN.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_SGN *SGNTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_SGN *SGNSession) RenouncePauser() (*types.Transaction, error) {
	return _SGN.Contract.RenouncePauser(&_SGN.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_SGN *SGNTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _SGN.Contract.RenouncePauser(&_SGN.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGN *SGNTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGN *SGNSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SGN.Contract.TransferOwnership(&_SGN.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SGN *SGNTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SGN.Contract.TransferOwnership(&_SGN.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SGN *SGNTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SGN *SGNSession) Unpause() (*types.Transaction, error) {
	return _SGN.Contract.Unpause(&_SGN.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SGN *SGNTransactorSession) Unpause() (*types.Transaction, error) {
	return _SGN.Contract.Unpause(&_SGN.TransactOpts)
}

// UpdateSgnAddr is a paid mutator transaction binding the contract method 0xd0bb9351.
//
// Solidity: function updateSgnAddr(bytes _sgnAddr) returns()
func (_SGN *SGNTransactor) UpdateSgnAddr(opts *bind.TransactOpts, _sgnAddr []byte) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "updateSgnAddr", _sgnAddr)
}

// UpdateSgnAddr is a paid mutator transaction binding the contract method 0xd0bb9351.
//
// Solidity: function updateSgnAddr(bytes _sgnAddr) returns()
func (_SGN *SGNSession) UpdateSgnAddr(_sgnAddr []byte) (*types.Transaction, error) {
	return _SGN.Contract.UpdateSgnAddr(&_SGN.TransactOpts, _sgnAddr)
}

// UpdateSgnAddr is a paid mutator transaction binding the contract method 0xd0bb9351.
//
// Solidity: function updateSgnAddr(bytes _sgnAddr) returns()
func (_SGN *SGNTransactorSession) UpdateSgnAddr(_sgnAddr []byte) (*types.Transaction, error) {
	return _SGN.Contract.UpdateSgnAddr(&_SGN.TransactOpts, _sgnAddr)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd88ef271.
//
// Solidity: function withdraw(bytes _withdrawalRequest, bytes[] _sigs) returns()
func (_SGN *SGNTransactor) Withdraw(opts *bind.TransactOpts, _withdrawalRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _SGN.contract.Transact(opts, "withdraw", _withdrawalRequest, _sigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd88ef271.
//
// Solidity: function withdraw(bytes _withdrawalRequest, bytes[] _sigs) returns()
func (_SGN *SGNSession) Withdraw(_withdrawalRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _SGN.Contract.Withdraw(&_SGN.TransactOpts, _withdrawalRequest, _sigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd88ef271.
//
// Solidity: function withdraw(bytes _withdrawalRequest, bytes[] _sigs) returns()
func (_SGN *SGNTransactorSession) Withdraw(_withdrawalRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _SGN.Contract.Withdraw(&_SGN.TransactOpts, _withdrawalRequest, _sigs)
}

// SGNDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SGN contract.
type SGNDepositIterator struct {
	Event *SGNDeposit // Event containing the contract specifics and raw log

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
func (it *SGNDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNDeposit)
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
		it.Event = new(SGNDeposit)
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
func (it *SGNDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNDeposit represents a Deposit event raised by the SGN contract.
type SGNDeposit struct {
	DepositId *big.Int
	Account   common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8.
//
// Solidity: event Deposit(uint256 depositId, address account, address token, uint256 amount)
func (_SGN *SGNFilterer) FilterDeposit(opts *bind.FilterOpts) (*SGNDepositIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &SGNDepositIterator{contract: _SGN.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8.
//
// Solidity: event Deposit(uint256 depositId, address account, address token, uint256 amount)
func (_SGN *SGNFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SGNDeposit) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNDeposit)
				if err := _SGN.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8.
//
// Solidity: event Deposit(uint256 depositId, address account, address token, uint256 amount)
func (_SGN *SGNFilterer) ParseDeposit(log types.Log) (*SGNDeposit, error) {
	event := new(SGNDeposit)
	if err := _SGN.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SGN contract.
type SGNOwnershipTransferredIterator struct {
	Event *SGNOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SGNOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNOwnershipTransferred)
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
		it.Event = new(SGNOwnershipTransferred)
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
func (it *SGNOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNOwnershipTransferred represents a OwnershipTransferred event raised by the SGN contract.
type SGNOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SGN *SGNFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SGNOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SGN.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SGNOwnershipTransferredIterator{contract: _SGN.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SGN *SGNFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SGNOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SGN.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNOwnershipTransferred)
				if err := _SGN.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SGN *SGNFilterer) ParseOwnershipTransferred(log types.Log) (*SGNOwnershipTransferred, error) {
	event := new(SGNOwnershipTransferred)
	if err := _SGN.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SGN contract.
type SGNPausedIterator struct {
	Event *SGNPaused // Event containing the contract specifics and raw log

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
func (it *SGNPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNPaused)
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
		it.Event = new(SGNPaused)
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
func (it *SGNPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNPaused represents a Paused event raised by the SGN contract.
type SGNPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGN *SGNFilterer) FilterPaused(opts *bind.FilterOpts) (*SGNPausedIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SGNPausedIterator{contract: _SGN.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGN *SGNFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SGNPaused) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNPaused)
				if err := _SGN.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SGN *SGNFilterer) ParsePaused(log types.Log) (*SGNPaused, error) {
	event := new(SGNPaused)
	if err := _SGN.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the SGN contract.
type SGNPauserAddedIterator struct {
	Event *SGNPauserAdded // Event containing the contract specifics and raw log

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
func (it *SGNPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNPauserAdded)
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
		it.Event = new(SGNPauserAdded)
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
func (it *SGNPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNPauserAdded represents a PauserAdded event raised by the SGN contract.
type SGNPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_SGN *SGNFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*SGNPauserAddedIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &SGNPauserAddedIterator{contract: _SGN.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_SGN *SGNFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *SGNPauserAdded) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNPauserAdded)
				if err := _SGN.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_SGN *SGNFilterer) ParsePauserAdded(log types.Log) (*SGNPauserAdded, error) {
	event := new(SGNPauserAdded)
	if err := _SGN.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the SGN contract.
type SGNPauserRemovedIterator struct {
	Event *SGNPauserRemoved // Event containing the contract specifics and raw log

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
func (it *SGNPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNPauserRemoved)
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
		it.Event = new(SGNPauserRemoved)
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
func (it *SGNPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNPauserRemoved represents a PauserRemoved event raised by the SGN contract.
type SGNPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_SGN *SGNFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*SGNPauserRemovedIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &SGNPauserRemovedIterator{contract: _SGN.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_SGN *SGNFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *SGNPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNPauserRemoved)
				if err := _SGN.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_SGN *SGNFilterer) ParsePauserRemoved(log types.Log) (*SGNPauserRemoved, error) {
	event := new(SGNPauserRemoved)
	if err := _SGN.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNSgnAddrUpdateIterator is returned from FilterSgnAddrUpdate and is used to iterate over the raw logs and unpacked data for SgnAddrUpdate events raised by the SGN contract.
type SGNSgnAddrUpdateIterator struct {
	Event *SGNSgnAddrUpdate // Event containing the contract specifics and raw log

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
func (it *SGNSgnAddrUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNSgnAddrUpdate)
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
		it.Event = new(SGNSgnAddrUpdate)
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
func (it *SGNSgnAddrUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNSgnAddrUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNSgnAddrUpdate represents a SgnAddrUpdate event raised by the SGN contract.
type SGNSgnAddrUpdate struct {
	ValAddr common.Address
	OldAddr []byte
	NewAddr []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSgnAddrUpdate is a free log retrieval operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes oldAddr, bytes newAddr)
func (_SGN *SGNFilterer) FilterSgnAddrUpdate(opts *bind.FilterOpts, valAddr []common.Address) (*SGNSgnAddrUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _SGN.contract.FilterLogs(opts, "SgnAddrUpdate", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &SGNSgnAddrUpdateIterator{contract: _SGN.contract, event: "SgnAddrUpdate", logs: logs, sub: sub}, nil
}

// WatchSgnAddrUpdate is a free log subscription operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes oldAddr, bytes newAddr)
func (_SGN *SGNFilterer) WatchSgnAddrUpdate(opts *bind.WatchOpts, sink chan<- *SGNSgnAddrUpdate, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _SGN.contract.WatchLogs(opts, "SgnAddrUpdate", valAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNSgnAddrUpdate)
				if err := _SGN.contract.UnpackLog(event, "SgnAddrUpdate", log); err != nil {
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

// ParseSgnAddrUpdate is a log parse operation binding the contract event 0x8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4.
//
// Solidity: event SgnAddrUpdate(address indexed valAddr, bytes oldAddr, bytes newAddr)
func (_SGN *SGNFilterer) ParseSgnAddrUpdate(log types.Log) (*SGNSgnAddrUpdate, error) {
	event := new(SGNSgnAddrUpdate)
	if err := _SGN.contract.UnpackLog(event, "SgnAddrUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SGN contract.
type SGNUnpausedIterator struct {
	Event *SGNUnpaused // Event containing the contract specifics and raw log

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
func (it *SGNUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNUnpaused)
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
		it.Event = new(SGNUnpaused)
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
func (it *SGNUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNUnpaused represents a Unpaused event raised by the SGN contract.
type SGNUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGN *SGNFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SGNUnpausedIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SGNUnpausedIterator{contract: _SGN.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGN *SGNFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SGNUnpaused) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNUnpaused)
				if err := _SGN.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SGN *SGNFilterer) ParseUnpaused(log types.Log) (*SGNUnpaused, error) {
	event := new(SGNUnpaused)
	if err := _SGN.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SGNWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SGN contract.
type SGNWithdrawIterator struct {
	Event *SGNWithdraw // Event containing the contract specifics and raw log

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
func (it *SGNWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SGNWithdraw)
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
		it.Event = new(SGNWithdraw)
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
func (it *SGNWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SGNWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SGNWithdraw represents a Withdraw event raised by the SGN contract.
type SGNWithdraw struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address account, address token, uint256 amount)
func (_SGN *SGNFilterer) FilterWithdraw(opts *bind.FilterOpts) (*SGNWithdrawIterator, error) {

	logs, sub, err := _SGN.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SGNWithdrawIterator{contract: _SGN.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address account, address token, uint256 amount)
func (_SGN *SGNFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SGNWithdraw) (event.Subscription, error) {

	logs, sub, err := _SGN.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SGNWithdraw)
				if err := _SGN.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address account, address token, uint256 amount)
func (_SGN *SGNFilterer) ParseWithdraw(log types.Log) (*SGNWithdraw, error) {
	event := new(SGNWithdraw)
	if err := _SGN.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorBondInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"tokenDiff\",\"type\":\"int256\"}],\"name\":\"DelegationUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmt\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashAmtCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ValidatorNotice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CELER_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bondedValAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectForfeiture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"completeUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"confirmUnbondedValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forfeiture\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorsTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ValidatorTokens[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawableUndelegationTokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.DelegatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"}],\"name\":\"getParamValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_checkSelfDelegation\",\"type\":\"bool\"}],\"name\":\"hasMinRequiredTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_commissionRate\",\"type\":\"uint64\"}],\"name\":\"initializeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBondedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBondBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setGovContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"name\":\"setMaxSlashFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setParamValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRewardContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerVals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegateShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"undelegateTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newRate\",\"type\":\"uint64\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"}],\"name\":\"updateMinSelfDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"updateValidatorSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"validatorNotice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"bondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"unbondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162006207380380620062078339810160408190526200003491620002cd565b6200003f33620001b4565b6000805460ff60a01b19169055620000573362000204565b6001600160a01b0399909916608052600b6020527fdf7de25b7f1fd6d0b5205f0e18f1f35bd7b8d84cce336588d184533ce43a6f76979097557f72c6bfb7988af3a1efa6568f02a999bc52252641c659d85961ca3d372b57d5cf959095557fa50eece07c7db1631545c0069bd8f5f54d5935e215d59097edf258a44ba91634939093557f64c15cc42be7899b001f818cf4433057002112c418d1d3a67cd5cb453051d33e919091557f12d0c11577e2f0950f57c455c117796550b79f444811db8ba2f69c57b646c784557febae6141bae5521e99e0a8d610356b0f501fea54980b59c84841db43ba7204f4557f0387e9d1203691d8e3362a7e4c6723de358a4010d7f31ecbec3fbfc61d1c75fc557ff5559028dc9ba50d75343c779b2f75e13a84a14662932fc67a486f263ca31a965560086000527f71f482bdabd1ea844d62c952b094e632959690d7448ca2aab34034ec98569358556200035a565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff1615620002725760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640160405180910390fd5b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910160405180910390a150565b6000806000806000806000806000806101408b8d031215620002ee57600080fd5b8a516001600160a01b03811681146200030657600080fd5b809a505060208b0151985060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935060e08b015192506101008b015191506101208b015190509295989b9194979a5092959850565b608051615e60620003a7600039600081816109a701528181610e08015281816110a501528181611bcb01528181611c460152818161217e01528181613149015261470f0152615e606000f3fe6080604052600436106103905760003560e01c80636ea69d62116101dc57806390e360f811610102578063b4f7fa34116100a0578063eb505dd51161006f578063eb505dd514610af8578063eecefef814610b25578063f2fde38b14610b52578063fa52c7d814610b7257600080fd5b8063b4f7fa3414610a5f578063c8f9f98414610a7f578063dcfdc1e114610ab8578063e909156d14610ad857600080fd5b8063960dc08a116100dc578063960dc08a146109955780639b19251a146109c9578063a310624f146109f9578063acc62ccf14610a3f57600080fd5b806390e360f8146109255780639146f1101461095557806392bb243c1461097557600080fd5b806382dc1ec41161017a57806388d996e81161014957806388d996e8146108b257806389f9aab5146108d25780638a74d5fe146108e75780638da5cb5b1461090757600080fd5b806382dc1ec4146108515780638338f0e51461087157806383cfb318146108875780638456cb591461089d57600080fd5b806371bc0216116101b657806371bc0216146107cc5780637a50dbd2146107ec57806380f51c121461080c57806382d7b4b81461083c57600080fd5b80636ea69d62146107825780636ef8d66d146107a2578063715018a6146107b757600080fd5b8063410ae02c116102c1578063525eba211161025f578063682dbc221161022e578063682dbc22146106ec57806368706e541461070c5780636b2c0f551461072c5780636d3087831461074c57600080fd5b8063525eba21146106775780635c975abb146106975780635e593eff146106b657806365d5d420146106d657600080fd5b806347abfdbf1161029b57806347abfdbf146105fd57806349955e391461061d57806351508f0a1461063d57806351fb012d1461065d57600080fd5b8063410ae02c1461058457806346fbf68e146105a4578063473849bd146105dd57600080fd5b80632fa4d12b1161032e5780633985c4e6116103085780633985c4e6146104e45780633af32abf146105045780633f4ba83a1461054d5780634021d4d51461056257600080fd5b80632fa4d12b1461048257806336f1635f146104ba578063386c024a146104cf57600080fd5b8063145aa1161161036a578063145aa116146103fe5780631a2032571461041e5780631cfe4f0b1461043e578063291d95491461046257600080fd5b8063026e402b1461039c578063052d9e7e146103be57806310154bad146103de57600080fd5b3661039757005b600080fd5b3480156103a857600080fd5b506103bc6103b7366004615360565b610c10565b005b3480156103ca57600080fd5b506103bc6103d9366004615398565b610e90565b3480156103ea57600080fd5b506103bc6103f93660046153b5565b610eeb565b34801561040a57600080fd5b506103bc6104193660046153d0565b610ff7565b34801561042a57600080fd5b506103bc6104393660046153d0565b6110cf565b34801561044a57600080fd5b506006545b6040519081526020015b60405180910390f35b34801561046e57600080fd5b506103bc61047d3660046153b5565b611145565b34801561048e57600080fd5b50600c546104a2906001600160a01b031681565b6040516001600160a01b039091168152602001610459565b3480156104c657600080fd5b506103bc611246565b3480156104db57600080fd5b5061044f6115d4565b3480156104f057600080fd5b506103bc6104ff366004615477565b611601565b34801561051057600080fd5b5061053d61051f3660046153b5565b6001600160a01b031660009081526002602052604090205460ff1690565b6040519015158152602001610459565b34801561055957600080fd5b506103bc611d4c565b34801561056e57600080fd5b50610577611db5565b60405161045991906154e3565b34801561059057600080fd5b5061044f61059f36600461554a565b611eb6565b3480156105b057600080fd5b5061053d6105bf3660046153b5565b6001600160a01b031660009081526001602052604090205460ff1690565b3480156105e957600080fd5b506103bc6105f83660046153b5565b611ef5565b34801561060957600080fd5b5061053d610618366004615565565b6121fd565b34801561062957600080fd5b506103bc6106383660046155b4565b6122b7565b34801561064957600080fd5b506103bc6106583660046153b5565b612419565b34801561066957600080fd5b5060035461053d9060ff1681565b34801561068357600080fd5b506103bc6106923660046155cf565b6124dc565b3480156106a357600080fd5b50600054600160a01b900460ff1661053d565b3480156106c257600080fd5b506103bc6106d13660046153d0565b61297e565b3480156106e257600080fd5b5061044f60045481565b3480156106f857600080fd5b506103bc6107073660046156c2565b612b97565b34801561071857600080fd5b506103bc6107273660046153b5565b612bfa565b34801561073857600080fd5b506103bc6107473660046153b5565b612cbd565b34801561075857600080fd5b506104a26107673660046153b5565b6009602052600090815260409020546001600160a01b031681565b34801561078e57600080fd5b50600d546104a2906001600160a01b031681565b3480156107ae57600080fd5b506103bc612d0e565b3480156107c357600080fd5b506103bc612d17565b3480156107d857600080fd5b506103bc6107e73660046153b5565b612d69565b3480156107f857600080fd5b506103bc6108073660046153b5565b612eb1565b34801561081857600080fd5b5061053d6108273660046153b5565b60016020526000908152604090205460ff1681565b34801561084857600080fd5b506103bc6130e3565b34801561085d57600080fd5b506103bc61086c3660046153b5565b61317a565b34801561087d57600080fd5b5061044f600e5481565b34801561089357600080fd5b5061044f60055481565b3480156108a957600080fd5b506103bc6131cb565b3480156108be57600080fd5b506103bc6108cd366004615360565b613232565b3480156108de57600080fd5b5060075461044f565b3480156108f357600080fd5b5061053d610902366004615806565b613330565b34801561091357600080fd5b506000546001600160a01b03166104a2565b34801561093157600080fd5b5061053d6109403660046153d0565b600a6020526000908152604090205460ff1681565b34801561096157600080fd5b506103bc61097036600461587e565b613514565b34801561098157600080fd5b506104a26109903660046153d0565b6135e2565b3480156109a157600080fd5b506104a27f000000000000000000000000000000000000000000000000000000000000000081565b3480156109d557600080fd5b5061053d6109e43660046153b5565b60026020526000908152604090205460ff1681565b348015610a0557600080fd5b50610a32610a143660046153b5565b6001600160a01b031660009081526008602052604090205460ff1690565b6040516104599190615937565b348015610a4b57600080fd5b506104a2610a5a3660046153d0565b61360c565b348015610a6b57600080fd5b5061053d610a7a3660046153b5565b61361c565b348015610a8b57600080fd5b5061044f610a9a3660046153b5565b6001600160a01b031660009081526008602052604090206001015490565b348015610ac457600080fd5b506103bc610ad3366004615360565b613654565b348015610ae457600080fd5b506103bc610af3366004615945565b61374c565b348015610b0457600080fd5b5061044f610b1336600461554a565b600b6020526000908152604090205481565b348015610b3157600080fd5b50610b45610b40366004615961565b61384f565b6040516104599190615994565b348015610b5e57600080fd5b506103bc610b6d3660046153b5565b613b48565b348015610b7e57600080fd5b50610bfa610b8d3660046153b5565b600860205260009081526040902080546001820154600283015460038401546004850154600686015460079096015460ff8616966101009096046001600160a01b0316959067ffffffffffffffff80821691680100000000000000008104821691600160801b909104168a565b6040516104599a99989796959493929190615a33565b600054600160a01b900460ff1615610c625760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064015b60405180910390fd5b33670de0b6b3a7640000821015610cbb5760405162461bcd60e51b815260206004820152601860248201527f4d696e696d616c20616d6f756e7420697320312043454c5200000000000000006044820152606401610c59565b6001600160a01b038316600090815260086020526040812090815460ff166003811115610cea57610cea6158ff565b1415610d385760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610c59565b6000610d4d8483600101548460020154613c15565b6001600160a01b0384166000908152600584016020526040812080549293509183918391610d7c908490615ab5565b9250508190555081836002016000828254610d979190615ab5565b9250508190555084836001016000828254610db29190615ab5565b9091555060039050835460ff166003811115610dd057610dd06158ff565b1415610dfb578460046000828254610de89190615ab5565b90915550506001830154610dfb90613c42565b610e306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016853088613d65565b6001830154815460408051928352602083019190915281018690526001600160a01b0380861691908816907f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea9060600160405180910390a3505050505050565b6000546001600160a01b03163314610ed85760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6003805460ff1916911515919091179055565b6000546001600160a01b03163314610f335760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6001600160a01b03811660009081526002602052604090205460ff1615610f9c5760405162461bcd60e51b815260206004820152601360248201527f416c72656164792077686974656c6973746564000000000000000000000000006044820152606401610c59565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f91015b60405180910390a150565b600054600160a01b900460ff166110505760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c59565b6000546001600160a01b031633146110985760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6110cc6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163383613dfd565b50565b6000546001600160a01b031633146111175760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6008600052600b6020527f71f482bdabd1ea844d62c952b094e632959690d7448ca2aab34034ec9856935855565b6000546001600160a01b0316331461118d5760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6001600160a01b03811660009081526002602052604090205460ff166111f55760405162461bcd60e51b815260206004820152600f60248201527f4e6f742077686974656c697374656400000000000000000000000000000000006044820152606401610c59565b6001600160a01b038116600081815260026020908152604091829020805460ff1916905590519182527f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b69101610fec565b336000818152600960205260409020546001600160a01b03161561127f5750336000908152600960205260409020546001600160a01b03165b6001600160a01b03811660009081526008602052604090206001815460ff1660038111156112af576112af6158ff565b14806112d057506002815460ff1660038111156112ce576112ce6158ff565b145b61131c5760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610c59565b600781015467ffffffffffffffff1643101561137a5760405162461bcd60e51b815260206004820152601660248201527f426f6e6420626c6f636b206e6f742072656163686564000000000000000000006044820152606401610c59565b6005544310156113cc5760405162461bcd60e51b815260206004820152601b60248201527f546f6f206672657175656e742076616c696461746f7220626f6e6400000000006044820152606401610c59565b6007600052600b6020527ff5559028dc9ba50d75343c779b2f75e13a84a14662932fc67a486f263ca31a96546114029043615ab5565b6005556114108260016121fd565b61145c5760405162461bcd60e51b815260206004820152601360248201527f4e6f742068617665206d696e20746f6b656e73000000000000000000000000006044820152606401610c59565b6003600052600b6020527f64c15cc42be7899b001f818cf4433057002112c418d1d3a67cd5cb453051d33e546007548111156114ad5761149b83613e2d565b6114a88260010154613c42565b505050565b6000196000805b83811015611562578260086000600784815481106114d4576114d4615acd565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154101561155057809150600860006007838154811061151c5761151c615acd565b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015492508261155057611562565b8061155a81615ae3565b9150506114b4565b50818460010154116115b65760405162461bcd60e51b815260206004820152601360248201527f496e73756666696369656e7420746f6b656e73000000000000000000000000006044820152606401610c59565b6115c08582613e81565b6115cd8460010154613c42565b5050505050565b6000600360045460026115e79190615afe565b6115f19190615b1d565b6115fc906001615ab5565b905090565b600054600160a01b900460ff161561164e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c59565b61169484848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506109029250859150869050615b3f565b5060006116d685858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613efc92505050565b9050806060015167ffffffffffffffff1642106117355760405162461bcd60e51b815260206004820152600d60248201527f536c6173682065787069726564000000000000000000000000000000000000006044820152606401610c59565b620f4240816040015167ffffffffffffffff1611156117965760405162461bcd60e51b815260206004820152601460248201527f496e76616c696420736c61736820666163746f720000000000000000000000006044820152606401610c59565b6008600052600b6020527f71f482bdabd1ea844d62c952b094e632959690d7448ca2aab34034ec9856935854604082015167ffffffffffffffff16111561181f5760405162461bcd60e51b815260206004820152601760248201527f457863656564206d617820736c61736820666163746f720000000000000000006044820152606401610c59565b60208082015167ffffffffffffffff166000908152600a909152604090205460ff161561188e5760405162461bcd60e51b815260206004820152601060248201527f5573656420736c617368206e6f6e6365000000000000000000000000000000006044820152606401610c59565b60208082015167ffffffffffffffff166000908152600a82526040808220805460ff1916600117905583516001600160a01b0381168352600890935290206003815460ff1660038111156118e4576118e46158ff565b148061190557506002815460ff166003811115611903576119036158ff565b145b6119515760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610c59565b6000620f4240846040015167ffffffffffffffff1683600101546119759190615afe565b61197f9190615b1d565b9050808260010160008282546119959190615b4c565b9091555060039050825460ff1660038111156119b3576119b36158ff565b1415611a025780600460008282546119cb9190615b4c565b9091555050608084015167ffffffffffffffff161515806119f457506119f28360016121fd565b155b15611a0257611a028361416e565b6002825460ff166003811115611a1a57611a1a6158ff565b148015611a3557506000846080015167ffffffffffffffff16115b15611a78576080840151611a539067ffffffffffffffff1643615ab5565b60078301805467ffffffffffffffff191667ffffffffffffffff929092169190911790555b60006001600160a01b0316836001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea8460010154600085611abf90615b63565b6040805193845260208401929092529082015260600160405180910390a36000620f4240856040015167ffffffffffffffff168460030154611b019190615afe565b611b0b9190615b1d565b905080836003016000828254611b219190615b4c565b90915550611b3190508183615ab5565b91506000805b8660a0015151811015611ccb5760008760a001518281518110611b5c57611b5c615acd565b6020026020010151905084816020015184611b779190615ab5565b1115611b8d57611b878386615b4c565b60208201525b602081015115611cb8576020810151611ba69084615ab5565b81519093506001600160a01b0316611c31576020810151611bf3906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016903390613dfd565b60208082015160405190815233917fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3910160405180910390a2611cb8565b80516020820151611c6c916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691613dfd565b80600001516001600160a01b03167fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a38260200151604051611caf91815260200190565b60405180910390a25b5080611cc381615ae3565b915050611b37565b50611cd68184615b4c565b600e6000828254611ce79190615ab5565b90915550506020808701516040805167ffffffffffffffff90921682529181018590526001600160a01b038716917f10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008910160405180910390a250505050505050505050565b3360009081526001602052604090205460ff16611dab5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610c59565b611db36142d4565b565b60075460609060009067ffffffffffffffff811115611dd657611dd661560b565b604051908082528060200260200182016040528015611e1b57816020015b6040805180820190915260008082526020820152815260200190600190039081611df45790505b50905060005b600754811015611eb057600060078281548110611e4057611e40615acd565b60009182526020808320909101546040805180820182526001600160a01b039092168083528085526008845293206001015491810191909152845191925090849084908110611e9157611e91615acd565b6020026020010181905250508080611ea890615ae3565b915050611e21565b50919050565b6000600b6000836008811115611ece57611ece6158ff565b6008811115611edf57611edf6158ff565b8152602001908152602001600020549050919050565b6001600160a01b03811660009081526008602052604081203391815460ff166003811115611f2557611f256158ff565b1415611f735760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610c59565b6001600160a01b03821660009081526005820160209081526040822060028352600b9091527fa50eece07c7db1631545c0069bd8f5f54d5935e215d59097edf258a44ba916345483549192909160019060ff166003811115611fd757611fd76158ff565b60028501549114915063ffffffff1660005b600285015463ffffffff6401000000009091048116908316101561209b578280612039575063ffffffff821660009081526001808701602052604090912001544390612036908690615ab5565b11155b156120845763ffffffff821660009081526001860160205260409020546120609082615ab5565b63ffffffff8316600090815260018088016020526040822082815501559050612089565b61209b565b8161209381615b80565b925050611fe9565b60028501805463ffffffff191663ffffffff8416179055806121255760405162461bcd60e51b815260206004820152602560248201527f4e6f20756e64656c65676174696f6e20726561647920746f20626520636f6d7060448201527f6c657465640000000000000000000000000000000000000000000000000000006064820152608401610c59565b600061213a828860030154896004015461437a565b9050818760040160008282546121509190615b4c565b925050819055508087600301600082825461216b9190615b4c565b909155506121a590506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168983613dfd565b876001600160a01b0316896001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c836040516121ea91815260200190565b60405180910390a3505050505050505050565b6001600160a01b03821660009081526008602090815260408220600181015460048452600b9092527f12d0c11577e2f0950f57c455c117796550b79f444811db8ba2f69c57b646c7845490919081101561225c576000925050506122b1565b83156122aa576001600160a01b0385166000908152600583016020526040812054600284015461228e9190849061437a565b905082600601548110156122a857600093505050506122b1565b505b6001925050505b92915050565b33600081815260086020526040812090815460ff1660038111156122dd576122dd6158ff565b141561232b5760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610c59565b6127108367ffffffffffffffff1611156123875760405162461bcd60e51b815260206004820152601060248201527f496e76616c6964206e65772072617465000000000000000000000000000000006044820152606401610c59565b60078101805467ffffffffffffffff60801b1916600160801b67ffffffffffffffff8616908102919091179091556040805160208101929092526001600160a01b038416917f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c910160408051601f198184030181529082905261240c91600090615bfc565b60405180910390a2505050565b6000546001600160a01b031633146124615760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b600d546001600160a01b0316156124ba5760405162461bcd60e51b815260206004820152601b60248201527f72657761726420636f6e747261637420616c72656164792073657400000000006044820152606401610c59565b600d80546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a01b900460ff16156125295760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c59565b60035460ff1615612593573360009081526002602052604090205460ff166125935760405162461bcd60e51b815260206004820152601960248201527f43616c6c6572206973206e6f742077686974656c6973746564000000000000006044820152606401610c59565b33600081815260086020526040812090815460ff1660038111156125b9576125b96158ff565b146126065760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f7220697320696e697469616c697a656400000000000000006044820152606401610c59565b6001600160a01b03851660009081526008602052604081205460ff166003811115612633576126336158ff565b146126805760405162461bcd60e51b815260206004820152601960248201527f5369676e6572206973206f746865722076616c696461746f72000000000000006044820152606401610c59565b6001600160a01b0382811660009081526009602052604090205416156126e85760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206973206f74686572207369676e6572000000000000006044820152606401610c59565b6001600160a01b0385811660009081526009602052604090205416156127505760405162461bcd60e51b815260206004820152601360248201527f5369676e657220616c72656164792075736564000000000000000000000000006044820152606401610c59565b6127108367ffffffffffffffff1611156127ac5760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636f6d6d697373696f6e20726174650000000000000000006044820152606401610c59565b6005600052600b6020527febae6141bae5521e99e0a8d610356b0f501fea54980b59c84841db43ba7204f4548410156128275760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610c59565b80547fffffffffffffffffffffff0000000000000000000000000000000000000000001660ff196101006001600160a01b03888116918202929092169290921760019081178455600680850188905560078501805467ffffffffffffffff60801b1916600160801b67ffffffffffffffff8a1602179055805491820190557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b031990811692861692831790915560009283526009602052604090922080549092161790556128ff8285610c10565b604080516001600160a01b03878116602083015291810186905267ffffffffffffffff85166060820152908316907f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c9060800160408051601f198184030181529082905261296f91600090615c5b565b60405180910390a25050505050565b33600081815260086020526040812090815460ff1660038111156129a4576129a46158ff565b14156129f25760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610c59565b6005600052600b6020527febae6141bae5521e99e0a8d610356b0f501fea54980b59c84841db43ba7204f454831015612a6d5760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610c59565b8060060154831015612b3a576003815460ff166003811115612a9157612a916158ff565b1415612adf5760405162461bcd60e51b815260206004820152601360248201527f56616c696461746f7220697320626f6e646564000000000000000000000000006044820152606401610c59565b6006600052600b6020527f0387e9d1203691d8e3362a7e4c6723de358a4010d7f31ecbec3fbfc61d1c75fc54612b159043615ab5565b60078201805467ffffffffffffffff191667ffffffffffffffff929092169190911790555b6006810183905560408051602081018590526001600160a01b038416917f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c910160408051601f198184030181529082905261240c91600090615c89565b612ba5876109028789615b3f565b612bf15760405162461bcd60e51b815260206004820152601560248201527f4661696c656420746f20766572696679207369677300000000000000000000006044820152606401610c59565b50505050505050565b6000546001600160a01b03163314612c425760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b600c546001600160a01b031615612c9b5760405162461bcd60e51b815260206004820152601860248201527f676f7620636f6e747261637420616c72656164792073657400000000000000006044820152606401610c59565b600c80546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314612d055760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6110cc81614393565b611db333614393565b6000546001600160a01b03163314612d5f5760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b611db3600061444c565b6001600160a01b03811660009081526008602052604090206002815460ff166003811115612d9957612d996158ff565b14612de65760405162461bcd60e51b815260206004820152601760248201527f56616c696461746f72206e6f7420756e626f6e64696e670000000000000000006044820152606401610c59565b600781015468010000000000000000900467ffffffffffffffff16431015612e505760405162461bcd60e51b815260206004820152601860248201527f556e626f6e6420626c6f636b206e6f74207265616368656400000000000000006044820152606401610c59565b805460ff1916600190811782556007820180546fffffffffffffffff0000000000000000191690555b6040516001600160a01b038416907fd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a90600090a35050565b33600081815260086020526040812090815460ff166003811115612ed757612ed76158ff565b1415612f255760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206e6f7420696e697469616c697a6564000000000000006044820152606401610c59565b6001600160a01b038381166000908152600960205260409020541615612f8d5760405162461bcd60e51b815260206004820152601360248201527f5369676e657220616c72656164792075736564000000000000000000000000006044820152606401610c59565b816001600160a01b0316836001600160a01b031614613020576001600160a01b03831660009081526008602052604081205460ff166003811115612fd357612fd36158ff565b146130205760405162461bcd60e51b815260206004820152601960248201527f5369676e6572206973206f746865722076616c696461746f72000000000000006044820152606401610c59565b8054610100908190046001600160a01b03908116600090815260096020908152604080832080546001600160a01b031990811690915586547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1689861696870217875585845292819020805490931693871693841790925581519081019390935290917f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c910160408051601f198184030181529082905261240c91600090615cd0565b6000600e54116131355760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20636f6c6c65637400000000000000000000000000006044820152606401610c59565b600d54600e54613173916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811692911690613dfd565b6000600e55565b6000546001600160a01b031633146131c25760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6110cc8161449c565b3360009081526001602052604090205460ff1661322a5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f74207061757365720000000000000000000000006044820152606401610c59565b611db361455a565b670de0b6b3a764000081101561328a5760405162461bcd60e51b815260206004820152601960248201527f4d696e696d616c20616d6f756e742069732031207368617265000000000000006044820152606401610c59565b6001600160a01b038216600090815260086020526040812090815460ff1660038111156132b9576132b96158ff565b14156133075760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610c59565b600061331c838360010154846002015461437a565b905061332a828583866145e2565b50505050565b60008061339184805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506000808061339f6115d4565b905060005b86518110156134cb5760006133db8883815181106133c4576133c4615acd565b60200260200101518761498590919063ffffffff16565b9050836001600160a01b0316816001600160a01b03161161343e5760405162461bcd60e51b815260206004820152601e60248201527f5369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610c59565b6001600160a01b038082166000908152600960209081526040808320549093168252600890522090935083906003815460ff166003811115613482576134826158ff565b1461348e5750506134b9565b600181015461349d9087615ab5565b95508386106134b65760019750505050505050506122b1565b50505b806134c381615ae3565b9150506133a4565b5060405162461bcd60e51b815260206004820152601260248201527f51756f72756d206e6f74207265616368656400000000000000000000000000006044820152606401610c59565b6001600160a01b038516600090815260086020526040812090815460ff166003811115613543576135436158ff565b14156135915760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610c59565b856001600160a01b03167f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c86868686336040516135d2959493929190615d40565b60405180910390a2505050505050565b600681815481106135f257600080fd5b6000918252602090912001546001600160a01b0316905081565b600781815481106135f257600080fd5b600060036001600160a01b03831660009081526008602052604090205460ff16600381111561364d5761364d6158ff565b1492915050565b670de0b6b3a76400008110156136ac5760405162461bcd60e51b815260206004820152601860248201527f4d696e696d616c20616d6f756e7420697320312043454c5200000000000000006044820152606401610c59565b6001600160a01b038216600090815260086020526040812090815460ff1660038111156136db576136db6158ff565b14156137295760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610c59565b600061373e8383600101548460020154613c15565b905061332a828585846145e2565b600c546001600160a01b031633146137a65760405162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f7420676f7620636f6e74726163740000000000006044820152606401610c59565b60038260088111156137ba576137ba6158ff565b1415613812576007548110156138125760405162461bcd60e51b815260206004820152600d60248201527f696e76616c69642076616c7565000000000000000000000000000000000000006044820152606401610c59565b80600b6000846008811115613829576138296158ff565b600881111561383a5761383a6158ff565b81526020810191909152604001600020555050565b6138916040518060c0016040528060006001600160a01b0316815260200160008152602001600081526020016060815260200160008152602001600081525090565b6001600160a01b03808416600090815260086020908152604080832093861683526005840190915281208054600184015460028501549293926138d592919061437a565b60026000908152600b6020527fa50eece07c7db1631545c0069bd8f5f54d5935e215d59097edf258a44ba916345485549293509091829190829060019060ff166003811115613926576139266158ff565b60028801549114915060009061394d9063ffffffff80821691640100000000900416615d83565b63ffffffff16905060008167ffffffffffffffff8111156139705761397061560b565b6040519080825280602002602001820160405280156139b557816020015b604080518082019091526000808252602082015281526020019060019003908161398e5790505b50905060005b82811015613acc57600289015460018a01906000906139e09063ffffffff1684615ab5565b815260200190815260200160002060405180604001604052908160008201548152602001600182015481525050828281518110613a1f57613a1f615acd565b6020026020010181905250818181518110613a3c57613a3c615acd565b60200260200101516000015187613a539190615ab5565b96508380613a8957504385838381518110613a7057613a70615acd565b602002602001015160200151613a869190615ab5565b11155b15613aba57818181518110613aa057613aa0615acd565b60200260200101516000015186613ab79190615ab5565b95505b80613ac481615ae3565b9150506139bb565b506000613ae2878b600301548c6004015461437a565b90506000613af9878c600301548d6004015461437a565b90506040518060c001604052808f6001600160a01b031681526020018a81526020018b600001548152602001848152602001838152602001828152509b50505050505050505050505092915050565b6000546001600160a01b03163314613b905760405162461bcd60e51b81526020600482018190526024820152600080516020615e0b8339815191526044820152606401610c59565b6001600160a01b038116613c0c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610c59565b6110cc8161444c565b600082613c23575082613c3b565b82613c2e8386615afe565b613c389190615b1d565b90505b9392505050565b6007546002811480613c545750806003145b15613cd957613c616115d4565b8210613cd55760405162461bcd60e51b815260206004820152602e60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201527f2071756f72756d20746f6b656e730000000000000000000000000000000000006064820152608401610c59565b5050565b6003811115613cd5576003600454613cf19190615b1d565b8210613cd55760405162461bcd60e51b815260206004820152602b60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201527f20312f3320746f6b656e730000000000000000000000000000000000000000006064820152608401610c59565b6040516001600160a01b038085166024830152831660448201526064810182905261332a9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152614a29565b6040516001600160a01b0383166024820152604481018290526114a890849063a9059cbb60e01b90606401613d99565b600780546001810182556000919091527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b0319166001600160a01b0383161790556110cc81614b0e565b613eb160078281548110613e9757613e97615acd565b6000918252602090912001546001600160a01b0316614b74565b8160078281548110613ec557613ec5615acd565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613cd582614b0e565b6040805160c0810182526000808252602080830182905282840182905260608084018390526080840183905260a084015283518085019094528184528301849052909190613f4b826006614c1b565b905080600681518110613f6057613f60615acd565b602002602001015167ffffffffffffffff811115613f8057613f8061560b565b604051908082528060200260200182016040528015613fc557816020015b6040805180820190915260008082526020820152815260200190600190039081613f9e5790505b508360a00181905250600081600681518110613fe357613fe3615acd565b6020026020010181815250506000805b602084015151845110156141655761400a84614cd5565b909250905081600114156140395761402961402485614d0f565b614dcc565b6001600160a01b03168552613ff3565b816002141561405f5761404b84614dd7565b67ffffffffffffffff166020860152613ff3565b81600314156140855761407184614dd7565b67ffffffffffffffff166040860152613ff3565b81600414156140ab5761409784614dd7565b67ffffffffffffffff166060860152613ff3565b81600514156140d1576140bd84614dd7565b67ffffffffffffffff166080860152613ff3565b8160061415614156576140eb6140e685614d0f565b614e59565b8560a001518460068151811061410357614103615acd565b60200260200101518151811061411b5761411b615acd565b60200260200101819052508260068151811061413957614139615acd565b60200260200101805180919061414e90615ae3565b905250613ff3565b6141608482614f00565b613ff3565b50505050919050565b60075460009061418090600190615b4c565b905060005b60075481101561428b57826001600160a01b0316600782815481106141ac576141ac615acd565b6000918252602090912001546001600160a01b03161415614279578181101561423d57600782815481106141e2576141e2615acd565b600091825260209091200154600780546001600160a01b03909216918390811061420e5761420e615acd565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600780548061424e5761424e615da8565b600082815260209020810160001990810180546001600160a01b03191690550190556114a883614b74565b8061428381615ae3565b915050614185565b5060405162461bcd60e51b815260206004820152601460248201527f4e6f7420626f6e6465642076616c696461746f720000000000000000000000006044820152606401610c59565b600054600160a01b900460ff1661432d5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610c59565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600081614388575082613c3b565b81613c2e8486615afe565b6001600160a01b03811660009081526001602052604090205460ff166143fb5760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f742070617573657200000000000000000000006044820152606401610c59565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e9101610fec565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff16156145055760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c726561647920706175736572000000000000006044820152606401610c59565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101610fec565b600054600160a01b900460ff16156145a75760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c59565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861435d3390565b3360008181526005860160205260408120805490918491839190614607908490615b4c565b92505081905550828660020160008282546146229190615b4c565b925050819055508386600101600082825461463d9190615b4c565b9091555050600286015460018701541480159061465c57508054600210155b15614681578054600287018054600090614677908490615b4c565b9091555050600081555b8054158061469857508054670de0b6b3a764000011155b6146e45760405162461bcd60e51b815260206004820152601b60248201527f6e6f7420656e6f7567682072656d61696e696e672073686172657300000000006044820152606401610c59565b6001865460ff1660038111156146fc576146fc6158ff565b141561478a576147366001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168386613dfd565b816001600160a01b0316856001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c8660405161477b91815260200190565b60405180910390a3505061332a565b6003865460ff1660038111156147a2576147a26158ff565b14156147ec5783600460008282546147ba9190615b4c565b925050819055506147df85866001600160a01b0316846001600160a01b0316146121fd565b6147ec576147ec8561416e565b6002810154600a9061480f9063ffffffff80821691640100000000900416615d83565b63ffffffff16106148625760405162461bcd60e51b815260206004820152601f60248201527f457863656564206d617820756e64656c65676174696f6e20656e7472696573006044820152606401610c59565b60006148778588600301548960040154613c15565b90508087600401600082825461488d9190615ab5565b92505081905550848760030160008282546148a89190615ab5565b909155505060028201805463ffffffff640100000000918290048116600090815260018087016020526040909120858155439181019190915583549093929004169060046148f583615b80565b91906101000a81548163ffffffff021916908363ffffffff16021790555050836001600160a01b0316876001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea8a6001015486600001548a61495d90615b63565b6040805193845260208401929092529082015260600160405180910390a35050505050505050565b60008151604114156149b95760208201516040830151606084015160001a6149af86828585614f72565b93505050506122b1565b8151604014156149e157602082015160408301516149d885838361511b565b925050506122b1565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610c59565b6000614a7e826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661515e9092919063ffffffff16565b8051909150156114a85780806020019051810190614a9c9190615dbe565b6114a85760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610c59565b6001600160a01b0381166000908152600860205260408120805460ff191660031781556007810180546fffffffffffffffff00000000000000001916905560018101546004805492939192909190614b67908490615ab5565b9091555060039050612e79565b6001600160a01b03811660009081526008602090815260408220805460ff191660029081178255909252600b90527fa50eece07c7db1631545c0069bd8f5f54d5935e215d59097edf258a44ba9163454614bce9043615ab5565b8160070160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806001015460046000828254614c0e9190615b4c565b9091555060029050612e79565b8151606090614c2b836001615ab5565b67ffffffffffffffff811115614c4357614c4361560b565b604051908082528060200260200182016040528015614c6c578160200160208202803683370190505b5091506000805b60208601515186511015614ccc57614c8a86614cd5565b80925081935050506001848381518110614ca657614ca6615acd565b60200260200101818151614cba9190615ab5565b905250614cc78682614f00565b614c73565b50509092525090565b6000806000614ce384614dd7565b9050614cf0600882615b1d565b9250806007166005811115614d0757614d076158ff565b915050915091565b60606000614d1c83614dd7565b90506000818460000151614d309190615ab5565b9050836020015151811115614d4457600080fd5b8167ffffffffffffffff811115614d5d57614d5d61560b565b6040519080825280601f01601f191660200182016040528015614d87576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015614dc1578181015183820152614dba602082615ab5565b9050614d9f565b505050935250919050565b60006122b18261516d565b602080820151825181019091015160009182805b600a811015614e535783811a9150614e04816007615afe565b82607f16901b851794508160801660001415614e4157614e25816001615ab5565b86518790614e34908390615ab5565b9052509395945050505050565b80614e4b81615ae3565b915050614deb565b50600080fd5b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b60208301515183511015614ef857614e9b83614cd5565b90925090508160011415614ec557614eb561402484614d0f565b6001600160a01b03168452614e84565b8160021415614ee957614edf614eda84614d0f565b615195565b6020850152614e84565b614ef38382614f00565b614e84565b505050919050565b6000816005811115614f1457614f146158ff565b1415614f23576114a882614dd7565b6002816005811115614f3757614f376158ff565b1415610397576000614f4883614dd7565b90508083600001818151614f5c9190615ab5565b905250602083015151835111156114a857600080fd5b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115614fef5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610c59565b8360ff16601b148061500457508360ff16601c145b61505b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610c59565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa1580156150af573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166151125760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610c59565b95945050505050565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821660ff83901c601b0161515486828785614f72565b9695505050505050565b6060613c3884846000856151cc565b6000815160141461517d57600080fd5b50602001516c01000000000000000000000000900490565b60006020825111156151a657600080fd5b60208201519050815160206151bb9190615b4c565b6151c6906008615afe565b1c919050565b6060824710156152445760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610c59565b843b6152925760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610c59565b600080866001600160a01b031685876040516152ae9190615ddb565b60006040518083038185875af1925050503d80600081146152eb576040519150601f19603f3d011682016040523d82523d6000602084013e6152f0565b606091505b509150915061530082828661530b565b979650505050505050565b6060831561531a575081613c3b565b82511561532a5782518084602001fd5b8160405162461bcd60e51b8152600401610c599190615df7565b80356001600160a01b038116811461535b57600080fd5b919050565b6000806040838503121561537357600080fd5b61537c83615344565b946020939093013593505050565b80151581146110cc57600080fd5b6000602082840312156153aa57600080fd5b8135613c3b8161538a565b6000602082840312156153c757600080fd5b613c3b82615344565b6000602082840312156153e257600080fd5b5035919050565b60008083601f8401126153fb57600080fd5b50813567ffffffffffffffff81111561541357600080fd5b60208301915083602082850101111561542b57600080fd5b9250929050565b60008083601f84011261544457600080fd5b50813567ffffffffffffffff81111561545c57600080fd5b6020830191508360208260051b850101111561542b57600080fd5b6000806000806040858703121561548d57600080fd5b843567ffffffffffffffff808211156154a557600080fd5b6154b1888389016153e9565b909650945060208701359150808211156154ca57600080fd5b506154d787828801615432565b95989497509550505050565b602080825282518282018190526000919060409081850190868401855b8281101561552e57815180516001600160a01b03168552860151868501529284019290850190600101615500565b5091979650505050505050565b80356009811061535b57600080fd5b60006020828403121561555c57600080fd5b613c3b8261553b565b6000806040838503121561557857600080fd5b61558183615344565b915060208301356155918161538a565b809150509250929050565b803567ffffffffffffffff8116811461535b57600080fd5b6000602082840312156155c657600080fd5b613c3b8261559c565b6000806000606084860312156155e457600080fd5b6155ed84615344565b9250602084013591506156026040850161559c565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561564a5761564a61560b565b604052919050565b600082601f83011261566357600080fd5b813567ffffffffffffffff81111561567d5761567d61560b565b615690601f8201601f1916602001615621565b8181528460208386010111156156a557600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060006080888a0312156156dd57600080fd5b873567ffffffffffffffff808211156156f557600080fd5b6157018b838c01615652565b985060208a013591508082111561571757600080fd5b6157238b838c01615432565b909850965060408a013591508082111561573c57600080fd5b6157488b838c01615432565b909650945060608a013591508082111561576157600080fd5b5061576e8a828b01615432565b989b979a50959850939692959293505050565b600067ffffffffffffffff8084111561579c5761579c61560b565b8360051b60206157ad818301615621565b868152935090840190808401878311156157c657600080fd5b855b838110156157fa578035858111156157e05760008081fd5b6157ec8a828a01615652565b8352509082019082016157c8565b50505050509392505050565b6000806040838503121561581957600080fd5b823567ffffffffffffffff8082111561583157600080fd5b61583d86838701615652565b9350602085013591508082111561585357600080fd5b508301601f8101851361586557600080fd5b61587485823560208401615781565b9150509250929050565b60008060008060006060868803121561589657600080fd5b61589f86615344565b9450602086013567ffffffffffffffff808211156158bc57600080fd5b6158c889838a016153e9565b909650945060408801359150808211156158e157600080fd5b506158ee888289016153e9565b969995985093965092949392505050565b634e487b7160e01b600052602160045260246000fd5b6004811061593357634e487b7160e01b600052602160045260246000fd5b9052565b602081016122b18284615915565b6000806040838503121561595857600080fd5b61537c8361553b565b6000806040838503121561597457600080fd5b61597d83615344565b915061598b60208401615344565b90509250929050565b6000602080835260e083016001600160a01b038551168285015281850151604081818701528087015160608701526060870151915060c06080870152828251808552610100880191508584019450600093505b80841015615a1057845180518352860151868301529385019360019390930192908201906159e7565b50608088015160a088015260a088015160c0880152809550505050505092915050565b6101408101615a42828d615915565b6001600160a01b039a909a16602082015260408101989098526060880196909652608087019490945260a086019290925260c085015267ffffffffffffffff90811660e08501529081166101008401521661012090910152919050565b634e487b7160e01b600052601160045260246000fd5b60008219821115615ac857615ac8615a9f565b500190565b634e487b7160e01b600052603260045260246000fd5b6000600019821415615af757615af7615a9f565b5060010190565b6000816000190483118215151615615b1857615b18615a9f565b500290565b600082615b3a57634e487b7160e01b600052601260045260246000fd5b500490565b6000613c3b368484615781565b600082821015615b5e57615b5e615a9f565b500390565b6000600160ff1b821415615b7957615b79615a9f565b5060000390565b600063ffffffff80831681811415615b9a57615b9a615a9f565b6001019392505050565b60005b83811015615bbf578181015183820152602001615ba7565b8381111561332a5750506000910152565b60008151808452615be8816020860160208601615ba4565b601f01601f19169290920160200192915050565b60608152600a60608201527f636f6d6d697373696f6e00000000000000000000000000000000000000000000608082015260a060208201526000615c4360a0830185615bd0565b90506001600160a01b03831660408301529392505050565b6060815260046060820152631a5b9a5d60e21b608082015260a060208201526000615c4360a0830185615bd0565b60608152601360608201527f6d696e2d73656c662d64656c65676174696f6e00000000000000000000000000608082015260a060208201526000615c4360a0830185615bd0565b60608152600660608201527f7369676e65720000000000000000000000000000000000000000000000000000608082015260a060208201526000615c4360a0830185615bd0565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b606081526000615d54606083018789615d17565b8281036020840152615d67818688615d17565b9150506001600160a01b03831660408301529695505050505050565b600063ffffffff83811690831681811015615da057615da0615a9f565b039392505050565b634e487b7160e01b600052603160045260246000fd5b600060208284031215615dd057600080fd5b8151613c3b8161538a565b60008251615ded818460208701615ba4565b9190910192915050565b602081526000613c3b6020830184615bd056fe4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a2646970667358221220055a212ea83a13bbad8b1499da4f37b85e4d9473e37c1bd7179b2ff36c821b4164736f6c63430008090033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address, _proposalDeposit *big.Int, _votingPeriod *big.Int, _unbondingPeriod *big.Int, _maxBondedValidators *big.Int, _minValidatorTokens *big.Int, _minSelfDelegation *big.Int, _advanceNoticePeriod *big.Int, _validatorBondInterval *big.Int, _maxSlashFactor *big.Int) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, _celerTokenAddress, _proposalDeposit, _votingPeriod, _unbondingPeriod, _maxBondedValidators, _minValidatorTokens, _minSelfDelegation, _advanceNoticePeriod, _validatorBondInterval, _maxSlashFactor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// CELERTOKEN is a free data retrieval call binding the contract method 0x960dc08a.
//
// Solidity: function CELER_TOKEN() view returns(address)
func (_Staking *StakingCaller) CELERTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "CELER_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CELERTOKEN is a free data retrieval call binding the contract method 0x960dc08a.
//
// Solidity: function CELER_TOKEN() view returns(address)
func (_Staking *StakingSession) CELERTOKEN() (common.Address, error) {
	return _Staking.Contract.CELERTOKEN(&_Staking.CallOpts)
}

// CELERTOKEN is a free data retrieval call binding the contract method 0x960dc08a.
//
// Solidity: function CELER_TOKEN() view returns(address)
func (_Staking *StakingCallerSession) CELERTOKEN() (common.Address, error) {
	return _Staking.Contract.CELERTOKEN(&_Staking.CallOpts)
}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_Staking *StakingCaller) BondedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "bondedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_Staking *StakingSession) BondedTokens() (*big.Int, error) {
	return _Staking.Contract.BondedTokens(&_Staking.CallOpts)
}

// BondedTokens is a free data retrieval call binding the contract method 0x65d5d420.
//
// Solidity: function bondedTokens() view returns(uint256)
func (_Staking *StakingCallerSession) BondedTokens() (*big.Int, error) {
	return _Staking.Contract.BondedTokens(&_Staking.CallOpts)
}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_Staking *StakingCaller) BondedValAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "bondedValAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_Staking *StakingSession) BondedValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.BondedValAddrs(&_Staking.CallOpts, arg0)
}

// BondedValAddrs is a free data retrieval call binding the contract method 0xacc62ccf.
//
// Solidity: function bondedValAddrs(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) BondedValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.BondedValAddrs(&_Staking.CallOpts, arg0)
}

// Forfeiture is a free data retrieval call binding the contract method 0x8338f0e5.
//
// Solidity: function forfeiture() view returns(uint256)
func (_Staking *StakingCaller) Forfeiture(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "forfeiture")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Forfeiture is a free data retrieval call binding the contract method 0x8338f0e5.
//
// Solidity: function forfeiture() view returns(uint256)
func (_Staking *StakingSession) Forfeiture() (*big.Int, error) {
	return _Staking.Contract.Forfeiture(&_Staking.CallOpts)
}

// Forfeiture is a free data retrieval call binding the contract method 0x8338f0e5.
//
// Solidity: function forfeiture() view returns(uint256)
func (_Staking *StakingCallerSession) Forfeiture() (*big.Int, error) {
	return _Staking.Contract.Forfeiture(&_Staking.CallOpts)
}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_Staking *StakingCaller) GetBondedValidatorNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getBondedValidatorNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_Staking *StakingSession) GetBondedValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetBondedValidatorNum(&_Staking.CallOpts)
}

// GetBondedValidatorNum is a free data retrieval call binding the contract method 0x89f9aab5.
//
// Solidity: function getBondedValidatorNum() view returns(uint256)
func (_Staking *StakingCallerSession) GetBondedValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetBondedValidatorNum(&_Staking.CallOpts)
}

// GetBondedValidatorsTokens is a free data retrieval call binding the contract method 0x4021d4d5.
//
// Solidity: function getBondedValidatorsTokens() view returns((address,uint256)[])
func (_Staking *StakingCaller) GetBondedValidatorsTokens(opts *bind.CallOpts) ([]DataTypesValidatorTokens, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getBondedValidatorsTokens")

	if err != nil {
		return *new([]DataTypesValidatorTokens), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesValidatorTokens)).(*[]DataTypesValidatorTokens)

	return out0, err

}

// GetBondedValidatorsTokens is a free data retrieval call binding the contract method 0x4021d4d5.
//
// Solidity: function getBondedValidatorsTokens() view returns((address,uint256)[])
func (_Staking *StakingSession) GetBondedValidatorsTokens() ([]DataTypesValidatorTokens, error) {
	return _Staking.Contract.GetBondedValidatorsTokens(&_Staking.CallOpts)
}

// GetBondedValidatorsTokens is a free data retrieval call binding the contract method 0x4021d4d5.
//
// Solidity: function getBondedValidatorsTokens() view returns((address,uint256)[])
func (_Staking *StakingCallerSession) GetBondedValidatorsTokens() ([]DataTypesValidatorTokens, error) {
	return _Staking.Contract.GetBondedValidatorsTokens(&_Staking.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,(uint256,uint256)[],uint256,uint256))
func (_Staking *StakingCaller) GetDelegatorInfo(opts *bind.CallOpts, _valAddr common.Address, _delAddr common.Address) (DataTypesDelegatorInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDelegatorInfo", _valAddr, _delAddr)

	if err != nil {
		return *new(DataTypesDelegatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesDelegatorInfo)).(*DataTypesDelegatorInfo)

	return out0, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,(uint256,uint256)[],uint256,uint256))
func (_Staking *StakingSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (DataTypesDelegatorInfo, error) {
	return _Staking.Contract.GetDelegatorInfo(&_Staking.CallOpts, _valAddr, _delAddr)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,(uint256,uint256)[],uint256,uint256))
func (_Staking *StakingCallerSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (DataTypesDelegatorInfo, error) {
	return _Staking.Contract.GetDelegatorInfo(&_Staking.CallOpts, _valAddr, _delAddr)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Staking *StakingCaller) GetParamValue(opts *bind.CallOpts, _name uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getParamValue", _name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Staking *StakingSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _Staking.Contract.GetParamValue(&_Staking.CallOpts, _name)
}

// GetParamValue is a free data retrieval call binding the contract method 0x410ae02c.
//
// Solidity: function getParamValue(uint8 _name) view returns(uint256)
func (_Staking *StakingCallerSession) GetParamValue(_name uint8) (*big.Int, error) {
	return _Staking.Contract.GetParamValue(&_Staking.CallOpts, _name)
}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_Staking *StakingCaller) GetQuorumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getQuorumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_Staking *StakingSession) GetQuorumTokens() (*big.Int, error) {
	return _Staking.Contract.GetQuorumTokens(&_Staking.CallOpts)
}

// GetQuorumTokens is a free data retrieval call binding the contract method 0x386c024a.
//
// Solidity: function getQuorumTokens() view returns(uint256)
func (_Staking *StakingCallerSession) GetQuorumTokens() (*big.Int, error) {
	return _Staking.Contract.GetQuorumTokens(&_Staking.CallOpts)
}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_Staking *StakingCaller) GetValidatorNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_Staking *StakingSession) GetValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetValidatorNum(&_Staking.CallOpts)
}

// GetValidatorNum is a free data retrieval call binding the contract method 0x1cfe4f0b.
//
// Solidity: function getValidatorNum() view returns(uint256)
func (_Staking *StakingCallerSession) GetValidatorNum() (*big.Int, error) {
	return _Staking.Contract.GetValidatorNum(&_Staking.CallOpts)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_Staking *StakingCaller) GetValidatorStatus(opts *bind.CallOpts, _valAddr common.Address) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorStatus", _valAddr)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_Staking *StakingSession) GetValidatorStatus(_valAddr common.Address) (uint8, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, _valAddr)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _valAddr) view returns(uint8)
func (_Staking *StakingCallerSession) GetValidatorStatus(_valAddr common.Address) (uint8, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, _valAddr)
}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_Staking *StakingCaller) GetValidatorTokens(opts *bind.CallOpts, _valAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorTokens", _valAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_Staking *StakingSession) GetValidatorTokens(_valAddr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetValidatorTokens(&_Staking.CallOpts, _valAddr)
}

// GetValidatorTokens is a free data retrieval call binding the contract method 0xc8f9f984.
//
// Solidity: function getValidatorTokens(address _valAddr) view returns(uint256)
func (_Staking *StakingCallerSession) GetValidatorTokens(_valAddr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetValidatorTokens(&_Staking.CallOpts, _valAddr)
}

// GovContract is a free data retrieval call binding the contract method 0x2fa4d12b.
//
// Solidity: function govContract() view returns(address)
func (_Staking *StakingCaller) GovContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "govContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovContract is a free data retrieval call binding the contract method 0x2fa4d12b.
//
// Solidity: function govContract() view returns(address)
func (_Staking *StakingSession) GovContract() (common.Address, error) {
	return _Staking.Contract.GovContract(&_Staking.CallOpts)
}

// GovContract is a free data retrieval call binding the contract method 0x2fa4d12b.
//
// Solidity: function govContract() view returns(address)
func (_Staking *StakingCallerSession) GovContract() (common.Address, error) {
	return _Staking.Contract.GovContract(&_Staking.CallOpts)
}

// HasMinRequiredTokens is a free data retrieval call binding the contract method 0x47abfdbf.
//
// Solidity: function hasMinRequiredTokens(address _valAddr, bool _checkSelfDelegation) view returns(bool)
func (_Staking *StakingCaller) HasMinRequiredTokens(opts *bind.CallOpts, _valAddr common.Address, _checkSelfDelegation bool) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasMinRequiredTokens", _valAddr, _checkSelfDelegation)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMinRequiredTokens is a free data retrieval call binding the contract method 0x47abfdbf.
//
// Solidity: function hasMinRequiredTokens(address _valAddr, bool _checkSelfDelegation) view returns(bool)
func (_Staking *StakingSession) HasMinRequiredTokens(_valAddr common.Address, _checkSelfDelegation bool) (bool, error) {
	return _Staking.Contract.HasMinRequiredTokens(&_Staking.CallOpts, _valAddr, _checkSelfDelegation)
}

// HasMinRequiredTokens is a free data retrieval call binding the contract method 0x47abfdbf.
//
// Solidity: function hasMinRequiredTokens(address _valAddr, bool _checkSelfDelegation) view returns(bool)
func (_Staking *StakingCallerSession) HasMinRequiredTokens(_valAddr common.Address, _checkSelfDelegation bool) (bool, error) {
	return _Staking.Contract.HasMinRequiredTokens(&_Staking.CallOpts, _valAddr, _checkSelfDelegation)
}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_Staking *StakingCaller) IsBondedValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isBondedValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_Staking *StakingSession) IsBondedValidator(_addr common.Address) (bool, error) {
	return _Staking.Contract.IsBondedValidator(&_Staking.CallOpts, _addr)
}

// IsBondedValidator is a free data retrieval call binding the contract method 0xb4f7fa34.
//
// Solidity: function isBondedValidator(address _addr) view returns(bool)
func (_Staking *StakingCallerSession) IsBondedValidator(_addr common.Address) (bool, error) {
	return _Staking.Contract.IsBondedValidator(&_Staking.CallOpts, _addr)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Staking *StakingCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Staking *StakingSession) IsPauser(account common.Address) (bool, error) {
	return _Staking.Contract.IsPauser(&_Staking.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Staking *StakingCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Staking.Contract.IsPauser(&_Staking.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Staking *StakingCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Staking *StakingSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Staking.Contract.IsWhitelisted(&_Staking.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Staking *StakingCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Staking.Contract.IsWhitelisted(&_Staking.CallOpts, account)
}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_Staking *StakingCaller) NextBondBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "nextBondBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_Staking *StakingSession) NextBondBlock() (*big.Int, error) {
	return _Staking.Contract.NextBondBlock(&_Staking.CallOpts)
}

// NextBondBlock is a free data retrieval call binding the contract method 0x83cfb318.
//
// Solidity: function nextBondBlock() view returns(uint256)
func (_Staking *StakingCallerSession) NextBondBlock() (*big.Int, error) {
	return _Staking.Contract.NextBondBlock(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Staking *StakingCaller) Params(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "params", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Staking *StakingSession) Params(arg0 uint8) (*big.Int, error) {
	return _Staking.Contract.Params(&_Staking.CallOpts, arg0)
}

// Params is a free data retrieval call binding the contract method 0xeb505dd5.
//
// Solidity: function params(uint8 ) view returns(uint256)
func (_Staking *StakingCallerSession) Params(arg0 uint8) (*big.Int, error) {
	return _Staking.Contract.Params(&_Staking.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCallerSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Staking *StakingCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Staking *StakingSession) Pausers(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Pausers(&_Staking.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Staking *StakingCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Pausers(&_Staking.CallOpts, arg0)
}

// RewardContract is a free data retrieval call binding the contract method 0x6ea69d62.
//
// Solidity: function rewardContract() view returns(address)
func (_Staking *StakingCaller) RewardContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardContract is a free data retrieval call binding the contract method 0x6ea69d62.
//
// Solidity: function rewardContract() view returns(address)
func (_Staking *StakingSession) RewardContract() (common.Address, error) {
	return _Staking.Contract.RewardContract(&_Staking.CallOpts)
}

// RewardContract is a free data retrieval call binding the contract method 0x6ea69d62.
//
// Solidity: function rewardContract() view returns(address)
func (_Staking *StakingCallerSession) RewardContract() (common.Address, error) {
	return _Staking.Contract.RewardContract(&_Staking.CallOpts)
}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_Staking *StakingCaller) SignerVals(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "signerVals", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_Staking *StakingSession) SignerVals(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.SignerVals(&_Staking.CallOpts, arg0)
}

// SignerVals is a free data retrieval call binding the contract method 0x6d308783.
//
// Solidity: function signerVals(address ) view returns(address)
func (_Staking *StakingCallerSession) SignerVals(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.SignerVals(&_Staking.CallOpts, arg0)
}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_Staking *StakingCaller) SlashNonces(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "slashNonces", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_Staking *StakingSession) SlashNonces(arg0 *big.Int) (bool, error) {
	return _Staking.Contract.SlashNonces(&_Staking.CallOpts, arg0)
}

// SlashNonces is a free data retrieval call binding the contract method 0x90e360f8.
//
// Solidity: function slashNonces(uint256 ) view returns(bool)
func (_Staking *StakingCallerSession) SlashNonces(arg0 *big.Int) (bool, error) {
	return _Staking.Contract.SlashNonces(&_Staking.CallOpts, arg0)
}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_Staking *StakingCaller) ValAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "valAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_Staking *StakingSession) ValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.ValAddrs(&_Staking.CallOpts, arg0)
}

// ValAddrs is a free data retrieval call binding the contract method 0x92bb243c.
//
// Solidity: function valAddrs(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) ValAddrs(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.ValAddrs(&_Staking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 minSelfDelegation, uint64 bondBlock, uint64 unbondBlock, uint64 commissionRate)
func (_Staking *StakingCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	MinSelfDelegation  *big.Int
	BondBlock          uint64
	UnbondBlock        uint64
	CommissionRate     uint64
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Status             uint8
		Signer             common.Address
		Tokens             *big.Int
		Shares             *big.Int
		UndelegationTokens *big.Int
		UndelegationShares *big.Int
		MinSelfDelegation  *big.Int
		BondBlock          uint64
		UnbondBlock        uint64
		CommissionRate     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Signer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Tokens = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Shares = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UndelegationTokens = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UndelegationShares = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MinSelfDelegation = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.BondBlock = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.UnbondBlock = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.CommissionRate = *abi.ConvertType(out[9], new(uint64)).(*uint64)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 minSelfDelegation, uint64 bondBlock, uint64 unbondBlock, uint64 commissionRate)
func (_Staking *StakingSession) Validators(arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	MinSelfDelegation  *big.Int
	BondBlock          uint64
	UnbondBlock        uint64
	CommissionRate     uint64
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint8 status, address signer, uint256 tokens, uint256 shares, uint256 undelegationTokens, uint256 undelegationShares, uint256 minSelfDelegation, uint64 bondBlock, uint64 unbondBlock, uint64 commissionRate)
func (_Staking *StakingCallerSession) Validators(arg0 common.Address) (struct {
	Status             uint8
	Signer             common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	UndelegationShares *big.Int
	MinSelfDelegation  *big.Int
	BondBlock          uint64
	UnbondBlock        uint64
	CommissionRate     uint64
}, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_Staking *StakingCaller) VerifySignatures(opts *bind.CallOpts, _msg []byte, _sigs [][]byte) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "verifySignatures", _msg, _sigs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_Staking *StakingSession) VerifySignatures(_msg []byte, _sigs [][]byte) (bool, error) {
	return _Staking.Contract.VerifySignatures(&_Staking.CallOpts, _msg, _sigs)
}

// VerifySignatures is a free data retrieval call binding the contract method 0x8a74d5fe.
//
// Solidity: function verifySignatures(bytes _msg, bytes[] _sigs) view returns(bool)
func (_Staking *StakingCallerSession) VerifySignatures(_msg []byte, _sigs [][]byte) (bool, error) {
	return _Staking.Contract.VerifySignatures(&_Staking.CallOpts, _msg, _sigs)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] , uint256[] ) view returns()
func (_Staking *StakingCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _sigs [][]byte, arg2 []common.Address, arg3 []*big.Int) error {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "verifySigs", _msg, _sigs, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] , uint256[] ) view returns()
func (_Staking *StakingSession) VerifySigs(_msg []byte, _sigs [][]byte, arg2 []common.Address, arg3 []*big.Int) error {
	return _Staking.Contract.VerifySigs(&_Staking.CallOpts, _msg, _sigs, arg2, arg3)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] , uint256[] ) view returns()
func (_Staking *StakingCallerSession) VerifySigs(_msg []byte, _sigs [][]byte, arg2 []common.Address, arg3 []*big.Int) error {
	return _Staking.Contract.VerifySigs(&_Staking.CallOpts, _msg, _sigs, arg2, arg3)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Whitelist(&_Staking.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Whitelist(&_Staking.CallOpts, arg0)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Staking *StakingCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Staking *StakingSession) WhitelistEnabled() (bool, error) {
	return _Staking.Contract.WhitelistEnabled(&_Staking.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Staking *StakingCallerSession) WhitelistEnabled() (bool, error) {
	return _Staking.Contract.WhitelistEnabled(&_Staking.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Staking *StakingTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Staking *StakingSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddPauser(&_Staking.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Staking *StakingTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddPauser(&_Staking.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Staking *StakingTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Staking *StakingSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddWhitelisted(&_Staking.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Staking *StakingTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddWhitelisted(&_Staking.TransactOpts, account)
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_Staking *StakingTransactor) BondValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "bondValidator")
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_Staking *StakingSession) BondValidator() (*types.Transaction, error) {
	return _Staking.Contract.BondValidator(&_Staking.TransactOpts)
}

// BondValidator is a paid mutator transaction binding the contract method 0x36f1635f.
//
// Solidity: function bondValidator() returns()
func (_Staking *StakingTransactorSession) BondValidator() (*types.Transaction, error) {
	return _Staking.Contract.BondValidator(&_Staking.TransactOpts)
}

// CollectForfeiture is a paid mutator transaction binding the contract method 0x82d7b4b8.
//
// Solidity: function collectForfeiture() returns()
func (_Staking *StakingTransactor) CollectForfeiture(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "collectForfeiture")
}

// CollectForfeiture is a paid mutator transaction binding the contract method 0x82d7b4b8.
//
// Solidity: function collectForfeiture() returns()
func (_Staking *StakingSession) CollectForfeiture() (*types.Transaction, error) {
	return _Staking.Contract.CollectForfeiture(&_Staking.TransactOpts)
}

// CollectForfeiture is a paid mutator transaction binding the contract method 0x82d7b4b8.
//
// Solidity: function collectForfeiture() returns()
func (_Staking *StakingTransactorSession) CollectForfeiture() (*types.Transaction, error) {
	return _Staking.Contract.CollectForfeiture(&_Staking.TransactOpts)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_Staking *StakingTransactor) CompleteUndelegate(opts *bind.TransactOpts, _valAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "completeUndelegate", _valAddr)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_Staking *StakingSession) CompleteUndelegate(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.CompleteUndelegate(&_Staking.TransactOpts, _valAddr)
}

// CompleteUndelegate is a paid mutator transaction binding the contract method 0x473849bd.
//
// Solidity: function completeUndelegate(address _valAddr) returns()
func (_Staking *StakingTransactorSession) CompleteUndelegate(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.CompleteUndelegate(&_Staking.TransactOpts, _valAddr)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_Staking *StakingTransactor) ConfirmUnbondedValidator(opts *bind.TransactOpts, _valAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "confirmUnbondedValidator", _valAddr)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_Staking *StakingSession) ConfirmUnbondedValidator(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ConfirmUnbondedValidator(&_Staking.TransactOpts, _valAddr)
}

// ConfirmUnbondedValidator is a paid mutator transaction binding the contract method 0x71bc0216.
//
// Solidity: function confirmUnbondedValidator(address _valAddr) returns()
func (_Staking *StakingTransactorSession) ConfirmUnbondedValidator(_valAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ConfirmUnbondedValidator(&_Staking.TransactOpts, _valAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, _valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", _valAddr, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingSession) Delegate(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, _valAddr, _tokens)
}

// Delegate is a paid mutator transaction binding the contract method 0x026e402b.
//
// Solidity: function delegate(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingTransactorSession) Delegate(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, _valAddr, _tokens)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_Staking *StakingTransactor) DrainToken(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "drainToken", _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_Staking *StakingSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DrainToken(&_Staking.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DrainToken(&_Staking.TransactOpts, _amount)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x525eba21.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint64 _commissionRate) returns()
func (_Staking *StakingTransactor) InitializeValidator(opts *bind.TransactOpts, _signer common.Address, _minSelfDelegation *big.Int, _commissionRate uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initializeValidator", _signer, _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x525eba21.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint64 _commissionRate) returns()
func (_Staking *StakingSession) InitializeValidator(_signer common.Address, _minSelfDelegation *big.Int, _commissionRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.InitializeValidator(&_Staking.TransactOpts, _signer, _minSelfDelegation, _commissionRate)
}

// InitializeValidator is a paid mutator transaction binding the contract method 0x525eba21.
//
// Solidity: function initializeValidator(address _signer, uint256 _minSelfDelegation, uint64 _commissionRate) returns()
func (_Staking *StakingTransactorSession) InitializeValidator(_signer common.Address, _minSelfDelegation *big.Int, _commissionRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.InitializeValidator(&_Staking.TransactOpts, _signer, _minSelfDelegation, _commissionRate)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Staking *StakingTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Staking *StakingSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemovePauser(&_Staking.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Staking *StakingTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemovePauser(&_Staking.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Staking *StakingTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Staking *StakingSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveWhitelisted(&_Staking.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Staking *StakingTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveWhitelisted(&_Staking.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Staking *StakingTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Staking *StakingSession) RenouncePauser() (*types.Transaction, error) {
	return _Staking.Contract.RenouncePauser(&_Staking.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Staking *StakingTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Staking.Contract.RenouncePauser(&_Staking.TransactOpts)
}

// SetGovContract is a paid mutator transaction binding the contract method 0x68706e54.
//
// Solidity: function setGovContract(address _addr) returns()
func (_Staking *StakingTransactor) SetGovContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setGovContract", _addr)
}

// SetGovContract is a paid mutator transaction binding the contract method 0x68706e54.
//
// Solidity: function setGovContract(address _addr) returns()
func (_Staking *StakingSession) SetGovContract(_addr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetGovContract(&_Staking.TransactOpts, _addr)
}

// SetGovContract is a paid mutator transaction binding the contract method 0x68706e54.
//
// Solidity: function setGovContract(address _addr) returns()
func (_Staking *StakingTransactorSession) SetGovContract(_addr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetGovContract(&_Staking.TransactOpts, _addr)
}

// SetMaxSlashFactor is a paid mutator transaction binding the contract method 0x1a203257.
//
// Solidity: function setMaxSlashFactor(uint256 _maxSlashFactor) returns()
func (_Staking *StakingTransactor) SetMaxSlashFactor(opts *bind.TransactOpts, _maxSlashFactor *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMaxSlashFactor", _maxSlashFactor)
}

// SetMaxSlashFactor is a paid mutator transaction binding the contract method 0x1a203257.
//
// Solidity: function setMaxSlashFactor(uint256 _maxSlashFactor) returns()
func (_Staking *StakingSession) SetMaxSlashFactor(_maxSlashFactor *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxSlashFactor(&_Staking.TransactOpts, _maxSlashFactor)
}

// SetMaxSlashFactor is a paid mutator transaction binding the contract method 0x1a203257.
//
// Solidity: function setMaxSlashFactor(uint256 _maxSlashFactor) returns()
func (_Staking *StakingTransactorSession) SetMaxSlashFactor(_maxSlashFactor *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxSlashFactor(&_Staking.TransactOpts, _maxSlashFactor)
}

// SetParamValue is a paid mutator transaction binding the contract method 0xe909156d.
//
// Solidity: function setParamValue(uint8 _name, uint256 _value) returns()
func (_Staking *StakingTransactor) SetParamValue(opts *bind.TransactOpts, _name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setParamValue", _name, _value)
}

// SetParamValue is a paid mutator transaction binding the contract method 0xe909156d.
//
// Solidity: function setParamValue(uint8 _name, uint256 _value) returns()
func (_Staking *StakingSession) SetParamValue(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetParamValue(&_Staking.TransactOpts, _name, _value)
}

// SetParamValue is a paid mutator transaction binding the contract method 0xe909156d.
//
// Solidity: function setParamValue(uint8 _name, uint256 _value) returns()
func (_Staking *StakingTransactorSession) SetParamValue(_name uint8, _value *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetParamValue(&_Staking.TransactOpts, _name, _value)
}

// SetRewardContract is a paid mutator transaction binding the contract method 0x51508f0a.
//
// Solidity: function setRewardContract(address _addr) returns()
func (_Staking *StakingTransactor) SetRewardContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setRewardContract", _addr)
}

// SetRewardContract is a paid mutator transaction binding the contract method 0x51508f0a.
//
// Solidity: function setRewardContract(address _addr) returns()
func (_Staking *StakingSession) SetRewardContract(_addr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetRewardContract(&_Staking.TransactOpts, _addr)
}

// SetRewardContract is a paid mutator transaction binding the contract method 0x51508f0a.
//
// Solidity: function setRewardContract(address _addr) returns()
func (_Staking *StakingTransactorSession) SetRewardContract(_addr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetRewardContract(&_Staking.TransactOpts, _addr)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Staking *StakingTransactor) SetWhitelistEnabled(opts *bind.TransactOpts, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setWhitelistEnabled", _whitelistEnabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Staking *StakingSession) SetWhitelistEnabled(_whitelistEnabled bool) (*types.Transaction, error) {
	return _Staking.Contract.SetWhitelistEnabled(&_Staking.TransactOpts, _whitelistEnabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Staking *StakingTransactorSession) SetWhitelistEnabled(_whitelistEnabled bool) (*types.Transaction, error) {
	return _Staking.Contract.SetWhitelistEnabled(&_Staking.TransactOpts, _whitelistEnabled)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, _slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", _slashRequest, _sigs)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_Staking *StakingSession) Slash(_slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _slashRequest, _sigs)
}

// Slash is a paid mutator transaction binding the contract method 0x3985c4e6.
//
// Solidity: function slash(bytes _slashRequest, bytes[] _sigs) returns()
func (_Staking *StakingTransactorSession) Slash(_slashRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _slashRequest, _sigs)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// UndelegateShares is a paid mutator transaction binding the contract method 0x88d996e8.
//
// Solidity: function undelegateShares(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingTransactor) UndelegateShares(opts *bind.TransactOpts, _valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegateShares", _valAddr, _shares)
}

// UndelegateShares is a paid mutator transaction binding the contract method 0x88d996e8.
//
// Solidity: function undelegateShares(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingSession) UndelegateShares(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UndelegateShares(&_Staking.TransactOpts, _valAddr, _shares)
}

// UndelegateShares is a paid mutator transaction binding the contract method 0x88d996e8.
//
// Solidity: function undelegateShares(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingTransactorSession) UndelegateShares(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UndelegateShares(&_Staking.TransactOpts, _valAddr, _shares)
}

// UndelegateTokens is a paid mutator transaction binding the contract method 0xdcfdc1e1.
//
// Solidity: function undelegateTokens(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingTransactor) UndelegateTokens(opts *bind.TransactOpts, _valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegateTokens", _valAddr, _tokens)
}

// UndelegateTokens is a paid mutator transaction binding the contract method 0xdcfdc1e1.
//
// Solidity: function undelegateTokens(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingSession) UndelegateTokens(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UndelegateTokens(&_Staking.TransactOpts, _valAddr, _tokens)
}

// UndelegateTokens is a paid mutator transaction binding the contract method 0xdcfdc1e1.
//
// Solidity: function undelegateTokens(address _valAddr, uint256 _tokens) returns()
func (_Staking *StakingTransactorSession) UndelegateTokens(_valAddr common.Address, _tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UndelegateTokens(&_Staking.TransactOpts, _valAddr, _tokens)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x49955e39.
//
// Solidity: function updateCommissionRate(uint64 _newRate) returns()
func (_Staking *StakingTransactor) UpdateCommissionRate(opts *bind.TransactOpts, _newRate uint64) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateCommissionRate", _newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x49955e39.
//
// Solidity: function updateCommissionRate(uint64 _newRate) returns()
func (_Staking *StakingSession) UpdateCommissionRate(_newRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.UpdateCommissionRate(&_Staking.TransactOpts, _newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x49955e39.
//
// Solidity: function updateCommissionRate(uint64 _newRate) returns()
func (_Staking *StakingTransactorSession) UpdateCommissionRate(_newRate uint64) (*types.Transaction, error) {
	return _Staking.Contract.UpdateCommissionRate(&_Staking.TransactOpts, _newRate)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_Staking *StakingTransactor) UpdateMinSelfDelegation(opts *bind.TransactOpts, _minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateMinSelfDelegation", _minSelfDelegation)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_Staking *StakingSession) UpdateMinSelfDelegation(_minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateMinSelfDelegation(&_Staking.TransactOpts, _minSelfDelegation)
}

// UpdateMinSelfDelegation is a paid mutator transaction binding the contract method 0x5e593eff.
//
// Solidity: function updateMinSelfDelegation(uint256 _minSelfDelegation) returns()
func (_Staking *StakingTransactorSession) UpdateMinSelfDelegation(_minSelfDelegation *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateMinSelfDelegation(&_Staking.TransactOpts, _minSelfDelegation)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_Staking *StakingTransactor) UpdateValidatorSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateValidatorSigner", _signer)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_Staking *StakingSession) UpdateValidatorSigner(_signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateValidatorSigner(&_Staking.TransactOpts, _signer)
}

// UpdateValidatorSigner is a paid mutator transaction binding the contract method 0x7a50dbd2.
//
// Solidity: function updateValidatorSigner(address _signer) returns()
func (_Staking *StakingTransactorSession) UpdateValidatorSigner(_signer common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateValidatorSigner(&_Staking.TransactOpts, _signer)
}

// ValidatorNotice is a paid mutator transaction binding the contract method 0x9146f110.
//
// Solidity: function validatorNotice(address _valAddr, string _key, bytes _data) returns()
func (_Staking *StakingTransactor) ValidatorNotice(opts *bind.TransactOpts, _valAddr common.Address, _key string, _data []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "validatorNotice", _valAddr, _key, _data)
}

// ValidatorNotice is a paid mutator transaction binding the contract method 0x9146f110.
//
// Solidity: function validatorNotice(address _valAddr, string _key, bytes _data) returns()
func (_Staking *StakingSession) ValidatorNotice(_valAddr common.Address, _key string, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.ValidatorNotice(&_Staking.TransactOpts, _valAddr, _key, _data)
}

// ValidatorNotice is a paid mutator transaction binding the contract method 0x9146f110.
//
// Solidity: function validatorNotice(address _valAddr, string _key, bytes _data) returns()
func (_Staking *StakingTransactorSession) ValidatorNotice(_valAddr common.Address, _key string, _data []byte) (*types.Transaction, error) {
	return _Staking.Contract.ValidatorNotice(&_Staking.TransactOpts, _valAddr, _key, _data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingDelegationUpdateIterator is returned from FilterDelegationUpdate and is used to iterate over the raw logs and unpacked data for DelegationUpdate events raised by the Staking contract.
type StakingDelegationUpdateIterator struct {
	Event *StakingDelegationUpdate // Event containing the contract specifics and raw log

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
func (it *StakingDelegationUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegationUpdate)
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
		it.Event = new(StakingDelegationUpdate)
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
func (it *StakingDelegationUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegationUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegationUpdate represents a DelegationUpdate event raised by the Staking contract.
type StakingDelegationUpdate struct {
	ValAddr   common.Address
	DelAddr   common.Address
	ValTokens *big.Int
	DelShares *big.Int
	TokenDiff *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegationUpdate is a free log retrieval operation binding the contract event 0x2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea.
//
// Solidity: event DelegationUpdate(address indexed valAddr, address indexed delAddr, uint256 valTokens, uint256 delShares, int256 tokenDiff)
func (_Staking *StakingFilterer) FilterDelegationUpdate(opts *bind.FilterOpts, valAddr []common.Address, delAddr []common.Address) (*StakingDelegationUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DelegationUpdate", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegationUpdateIterator{contract: _Staking.contract, event: "DelegationUpdate", logs: logs, sub: sub}, nil
}

// WatchDelegationUpdate is a free log subscription operation binding the contract event 0x2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea.
//
// Solidity: event DelegationUpdate(address indexed valAddr, address indexed delAddr, uint256 valTokens, uint256 delShares, int256 tokenDiff)
func (_Staking *StakingFilterer) WatchDelegationUpdate(opts *bind.WatchOpts, sink chan<- *StakingDelegationUpdate, valAddr []common.Address, delAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DelegationUpdate", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegationUpdate)
				if err := _Staking.contract.UnpackLog(event, "DelegationUpdate", log); err != nil {
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

// ParseDelegationUpdate is a log parse operation binding the contract event 0x2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea.
//
// Solidity: event DelegationUpdate(address indexed valAddr, address indexed delAddr, uint256 valTokens, uint256 delShares, int256 tokenDiff)
func (_Staking *StakingFilterer) ParseDelegationUpdate(log types.Log) (*StakingDelegationUpdate, error) {
	event := new(StakingDelegationUpdate)
	if err := _Staking.contract.UnpackLog(event, "DelegationUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Staking contract.
type StakingPausedIterator struct {
	Event *StakingPaused // Event containing the contract specifics and raw log

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
func (it *StakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPaused)
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
		it.Event = new(StakingPaused)
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
func (it *StakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPaused represents a Paused event raised by the Staking contract.
type StakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingPausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingPausedIterator{contract: _Staking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingPaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPaused)
				if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) ParsePaused(log types.Log) (*StakingPaused, error) {
	event := new(StakingPaused)
	if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Staking contract.
type StakingPauserAddedIterator struct {
	Event *StakingPauserAdded // Event containing the contract specifics and raw log

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
func (it *StakingPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPauserAdded)
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
		it.Event = new(StakingPauserAdded)
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
func (it *StakingPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPauserAdded represents a PauserAdded event raised by the Staking contract.
type StakingPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Staking *StakingFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*StakingPauserAddedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &StakingPauserAddedIterator{contract: _Staking.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Staking *StakingFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *StakingPauserAdded) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPauserAdded)
				if err := _Staking.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Staking *StakingFilterer) ParsePauserAdded(log types.Log) (*StakingPauserAdded, error) {
	event := new(StakingPauserAdded)
	if err := _Staking.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Staking contract.
type StakingPauserRemovedIterator struct {
	Event *StakingPauserRemoved // Event containing the contract specifics and raw log

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
func (it *StakingPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPauserRemoved)
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
		it.Event = new(StakingPauserRemoved)
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
func (it *StakingPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPauserRemoved represents a PauserRemoved event raised by the Staking contract.
type StakingPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Staking *StakingFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*StakingPauserRemovedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &StakingPauserRemovedIterator{contract: _Staking.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Staking *StakingFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *StakingPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPauserRemoved)
				if err := _Staking.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Staking *StakingFilterer) ParsePauserRemoved(log types.Log) (*StakingPauserRemoved, error) {
	event := new(StakingPauserRemoved)
	if err := _Staking.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashIterator is returned from FilterSlash and is used to iterate over the raw logs and unpacked data for Slash events raised by the Staking contract.
type StakingSlashIterator struct {
	Event *StakingSlash // Event containing the contract specifics and raw log

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
func (it *StakingSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlash)
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
		it.Event = new(StakingSlash)
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
func (it *StakingSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlash represents a Slash event raised by the Staking contract.
type StakingSlash struct {
	ValAddr  common.Address
	Nonce    uint64
	SlashAmt *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlash is a free log retrieval operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_Staking *StakingFilterer) FilterSlash(opts *bind.FilterOpts, valAddr []common.Address) (*StakingSlashIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Slash", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashIterator{contract: _Staking.contract, event: "Slash", logs: logs, sub: sub}, nil
}

// WatchSlash is a free log subscription operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_Staking *StakingFilterer) WatchSlash(opts *bind.WatchOpts, sink chan<- *StakingSlash, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Slash", valAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlash)
				if err := _Staking.contract.UnpackLog(event, "Slash", log); err != nil {
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

// ParseSlash is a log parse operation binding the contract event 0x10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008.
//
// Solidity: event Slash(address indexed valAddr, uint64 nonce, uint256 slashAmt)
func (_Staking *StakingFilterer) ParseSlash(log types.Log) (*StakingSlash, error) {
	event := new(StakingSlash)
	if err := _Staking.contract.UnpackLog(event, "Slash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashAmtCollectedIterator is returned from FilterSlashAmtCollected and is used to iterate over the raw logs and unpacked data for SlashAmtCollected events raised by the Staking contract.
type StakingSlashAmtCollectedIterator struct {
	Event *StakingSlashAmtCollected // Event containing the contract specifics and raw log

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
func (it *StakingSlashAmtCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlashAmtCollected)
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
		it.Event = new(StakingSlashAmtCollected)
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
func (it *StakingSlashAmtCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashAmtCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlashAmtCollected represents a SlashAmtCollected event raised by the Staking contract.
type StakingSlashAmtCollected struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashAmtCollected is a free log retrieval operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_Staking *StakingFilterer) FilterSlashAmtCollected(opts *bind.FilterOpts, recipient []common.Address) (*StakingSlashAmtCollectedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SlashAmtCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashAmtCollectedIterator{contract: _Staking.contract, event: "SlashAmtCollected", logs: logs, sub: sub}, nil
}

// WatchSlashAmtCollected is a free log subscription operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_Staking *StakingFilterer) WatchSlashAmtCollected(opts *bind.WatchOpts, sink chan<- *StakingSlashAmtCollected, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SlashAmtCollected", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlashAmtCollected)
				if err := _Staking.contract.UnpackLog(event, "SlashAmtCollected", log); err != nil {
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

// ParseSlashAmtCollected is a log parse operation binding the contract event 0xb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3.
//
// Solidity: event SlashAmtCollected(address indexed recipient, uint256 amount)
func (_Staking *StakingFilterer) ParseSlashAmtCollected(log types.Log) (*StakingSlashAmtCollected, error) {
	event := new(StakingSlashAmtCollected)
	if err := _Staking.contract.UnpackLog(event, "SlashAmtCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Staking contract.
type StakingUndelegatedIterator struct {
	Event *StakingUndelegated // Event containing the contract specifics and raw log

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
func (it *StakingUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUndelegated)
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
		it.Event = new(StakingUndelegated)
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
func (it *StakingUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUndelegated represents a Undelegated event raised by the Staking contract.
type StakingUndelegated struct {
	ValAddr common.Address
	DelAddr common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_Staking *StakingFilterer) FilterUndelegated(opts *bind.FilterOpts, valAddr []common.Address, delAddr []common.Address) (*StakingUndelegatedIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Undelegated", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingUndelegatedIterator{contract: _Staking.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_Staking *StakingFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *StakingUndelegated, valAddr []common.Address, delAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var delAddrRule []interface{}
	for _, delAddrItem := range delAddr {
		delAddrRule = append(delAddrRule, delAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Undelegated", valAddrRule, delAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUndelegated)
				if err := _Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed valAddr, address indexed delAddr, uint256 amount)
func (_Staking *StakingFilterer) ParseUndelegated(log types.Log) (*StakingUndelegated, error) {
	event := new(StakingUndelegated)
	if err := _Staking.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Staking contract.
type StakingUnpausedIterator struct {
	Event *StakingUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnpaused)
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
		it.Event = new(StakingUnpaused)
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
func (it *StakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnpaused represents a Unpaused event raised by the Staking contract.
type StakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingUnpausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingUnpausedIterator{contract: _Staking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnpaused)
				if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) ParseUnpaused(log types.Log) (*StakingUnpaused, error) {
	event := new(StakingUnpaused)
	if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorNoticeIterator is returned from FilterValidatorNotice and is used to iterate over the raw logs and unpacked data for ValidatorNotice events raised by the Staking contract.
type StakingValidatorNoticeIterator struct {
	Event *StakingValidatorNotice // Event containing the contract specifics and raw log

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
func (it *StakingValidatorNoticeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorNotice)
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
		it.Event = new(StakingValidatorNotice)
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
func (it *StakingValidatorNoticeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorNoticeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorNotice represents a ValidatorNotice event raised by the Staking contract.
type StakingValidatorNotice struct {
	ValAddr common.Address
	Key     string
	Data    []byte
	From    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorNotice is a free log retrieval operation binding the contract event 0x3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c.
//
// Solidity: event ValidatorNotice(address indexed valAddr, string key, bytes data, address from)
func (_Staking *StakingFilterer) FilterValidatorNotice(opts *bind.FilterOpts, valAddr []common.Address) (*StakingValidatorNoticeIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorNotice", valAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorNoticeIterator{contract: _Staking.contract, event: "ValidatorNotice", logs: logs, sub: sub}, nil
}

// WatchValidatorNotice is a free log subscription operation binding the contract event 0x3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c.
//
// Solidity: event ValidatorNotice(address indexed valAddr, string key, bytes data, address from)
func (_Staking *StakingFilterer) WatchValidatorNotice(opts *bind.WatchOpts, sink chan<- *StakingValidatorNotice, valAddr []common.Address) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorNotice", valAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorNotice)
				if err := _Staking.contract.UnpackLog(event, "ValidatorNotice", log); err != nil {
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

// ParseValidatorNotice is a log parse operation binding the contract event 0x3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c.
//
// Solidity: event ValidatorNotice(address indexed valAddr, string key, bytes data, address from)
func (_Staking *StakingFilterer) ParseValidatorNotice(log types.Log) (*StakingValidatorNotice, error) {
	event := new(StakingValidatorNotice)
	if err := _Staking.contract.UnpackLog(event, "ValidatorNotice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorStatusUpdateIterator is returned from FilterValidatorStatusUpdate and is used to iterate over the raw logs and unpacked data for ValidatorStatusUpdate events raised by the Staking contract.
type StakingValidatorStatusUpdateIterator struct {
	Event *StakingValidatorStatusUpdate // Event containing the contract specifics and raw log

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
func (it *StakingValidatorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorStatusUpdate)
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
		it.Event = new(StakingValidatorStatusUpdate)
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
func (it *StakingValidatorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorStatusUpdate represents a ValidatorStatusUpdate event raised by the Staking contract.
type StakingValidatorStatusUpdate struct {
	ValAddr common.Address
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusUpdate is a free log retrieval operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_Staking *StakingFilterer) FilterValidatorStatusUpdate(opts *bind.FilterOpts, valAddr []common.Address, status []uint8) (*StakingValidatorStatusUpdateIterator, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorStatusUpdate", valAddrRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorStatusUpdateIterator{contract: _Staking.contract, event: "ValidatorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusUpdate is a free log subscription operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_Staking *StakingFilterer) WatchValidatorStatusUpdate(opts *bind.WatchOpts, sink chan<- *StakingValidatorStatusUpdate, valAddr []common.Address, status []uint8) (event.Subscription, error) {

	var valAddrRule []interface{}
	for _, valAddrItem := range valAddr {
		valAddrRule = append(valAddrRule, valAddrItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorStatusUpdate", valAddrRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorStatusUpdate)
				if err := _Staking.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
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

// ParseValidatorStatusUpdate is a log parse operation binding the contract event 0xd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a.
//
// Solidity: event ValidatorStatusUpdate(address indexed valAddr, uint8 indexed status)
func (_Staking *StakingFilterer) ParseValidatorStatusUpdate(log types.Log) (*StakingValidatorStatusUpdate, error) {
	event := new(StakingValidatorStatusUpdate)
	if err := _Staking.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the Staking contract.
type StakingWhitelistedAddedIterator struct {
	Event *StakingWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *StakingWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhitelistedAdded)
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
		it.Event = new(StakingWhitelistedAdded)
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
func (it *StakingWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhitelistedAdded represents a WhitelistedAdded event raised by the Staking contract.
type StakingWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Staking *StakingFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts) (*StakingWhitelistedAddedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return &StakingWhitelistedAddedIterator{contract: _Staking.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Staking *StakingFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *StakingWhitelistedAdded) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhitelistedAdded)
				if err := _Staking.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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

// ParseWhitelistedAdded is a log parse operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Staking *StakingFilterer) ParseWhitelistedAdded(log types.Log) (*StakingWhitelistedAdded, error) {
	event := new(StakingWhitelistedAdded)
	if err := _Staking.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the Staking contract.
type StakingWhitelistedRemovedIterator struct {
	Event *StakingWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *StakingWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhitelistedRemoved)
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
		it.Event = new(StakingWhitelistedRemoved)
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
func (it *StakingWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhitelistedRemoved represents a WhitelistedRemoved event raised by the Staking contract.
type StakingWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Staking *StakingFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts) (*StakingWhitelistedRemovedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return &StakingWhitelistedRemovedIterator{contract: _Staking.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Staking *StakingFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *StakingWhitelistedRemoved) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhitelistedRemoved)
				if err := _Staking.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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

// ParseWhitelistedRemoved is a log parse operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Staking *StakingFilterer) ParseWhitelistedRemoved(log types.Log) (*StakingWhitelistedRemoved, error) {
	event := new(StakingWhitelistedRemoved)
	if err := _Staking.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardMetaData contains all meta data concerning the StakingReward contract.
var StakingRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"StakingRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"}],\"name\":\"StakingRewardContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b506040516200199938038062001999833981016040819052620000349162000182565b6200003f3362000069565b6000805460ff60a01b191690556200005733620000b9565b6001600160a01b0316608052620001b4565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff1615620001275760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640160405180910390fd5b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910160405180910390a150565b6000602082840312156200019557600080fd5b81516001600160a01b0381168114620001ad57600080fd5b9392505050565b6080516117ad620001ec6000396000818161017b015281816102ec0152818161048d0152818161087501526109fb01526117ad6000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a6116100975780638da5cb5b116100665780638da5cb5b1461022857806396db0fef14610239578063f2fde38b14610267578063f8df0dc51461027a57600080fd5b8063715018a6146101e257806380f51c12146101ea57806382dc1ec41461020d5780638456cb591461022057600080fd5b80634cf088d9116100d35780634cf088d9146101765780635c975abb146101b55780636b2c0f55146101c75780636ef8d66d146101da57600080fd5b80630a300b0914610105578063145aa1161461011a5780633f4ba83a1461012d57806346fbf68e14610135575b600080fd5b610118610113366004611406565b61028d565b005b610118610128366004611406565b6103d3565b61011861052f565b610161610143366004611434565b6001600160a01b031660009081526001602052604090205460ff1690565b60405190151581526020015b60405180910390f35b61019d7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161016d565b600054600160a01b900460ff16610161565b6101186101d5366004611434565b610598565b6101186105fb565b610118610604565b6101616101f8366004611434565b60016020526000908152604090205460ff1681565b61011861021b366004611434565b610668565b6101186106cb565b6000546001600160a01b031661019d565b610259610247366004611434565b60026020526000908152604090205481565b60405190815260200161016d565b610118610275366004611434565b610732565b610118610288366004611451565b610811565b600054600160a01b900460ff16156102df5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064015b60405180910390fd5b600033905061038c8130847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561034357600080fd5b505afa158015610357573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061037b9190611516565b6001600160a01b0316929190610aa2565b806001600160a01b03167ff67017a05194c0853be9169be60cad9fa6e75d34b6b507a7a4261510e19c6d79836040516103c791815260200190565b60405180910390a25050565b600054600160a01b900460ff1661042c5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102d6565b6000546001600160a01b031633146104865760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d6565b61052c33827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156104e457600080fd5b505afa1580156104f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061051c9190611516565b6001600160a01b03169190610b40565b50565b3360009081526001602052604090205460ff1661058e5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016102d6565b610596610b75565b565b6000546001600160a01b031633146105f25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d6565b61052c81610c1b565b61059633610c1b565b6000546001600160a01b0316331461065e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d6565b6105966000610cdb565b6000546001600160a01b031633146106c25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d6565b61052c81610d43565b3360009081526001602052604090205460ff1661072a5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016102d6565b610596610e01565b6000546001600160a01b0316331461078c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102d6565b6001600160a01b0381166108085760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102d6565b61052c81610cdb565b600054600160a01b900460ff161561085e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102d6565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe906108b090879087908790879060040161155c565b60206040518083038186803b1580156108c857600080fd5b505afa1580156108dc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610900919061160d565b50600061094285858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610e8992505050565b60208082015182516001600160a01b0316600090815260029092526040822054929350916109709083611645565b9050600081116109c25760405162461bcd60e51b815260206004820152600d60248201527f4e6f206e6577207265776172640000000000000000000000000000000000000060448201526064016102d6565b816002600085600001516001600160a01b03166001600160a01b0316815260200190815260200160002081905550610a528360000151827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156104e457600080fd5b82600001516001600160a01b03167f6dd401e61ba732582a5eba3d54ccc3afb3609cd3ac1a166d1d36f75fc0aedcda82604051610a9191815260200190565b60405180910390a250505050505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610b3a9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610f35565b50505050565b6040516001600160a01b038316602482015260448101829052610b7090849063a9059cbb60e01b90606401610ad6565b505050565b600054600160a01b900460ff16610bce5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102d6565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b03811660009081526001602052604090205460ff16610c835760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f7420706175736572000000000000000000000060448201526064016102d6565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e91015b60405180910390a150565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526001602052604090205460ff1615610dac5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064016102d6565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101610cd0565b600054600160a01b900460ff1615610e4e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102d6565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610bfe3390565b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b60208301515183511015610f2d57610ecb8361101a565b90925090508160011415610efa57610eea610ee584611054565b611111565b6001600160a01b03168452610eb4565b8160021415610f1e57610f14610f0f84611054565b611122565b6020850152610eb4565b610f288382611159565b610eb4565b505050919050565b6000610f8a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111cb9092919063ffffffff16565b805190915015610b705780806020019051810190610fa8919061160d565b610b705760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102d6565b6000806000611028846111e4565b905061103560088261165c565b925080600716600581111561104c5761104c61167e565b915050915091565b60606000611061836111e4565b905060008184600001516110759190611694565b905083602001515181111561108957600080fd5b8167ffffffffffffffff8111156110a2576110a26116ac565b6040519080825280601f01601f1916602001820160405280156110cc576020820181803683370190505b50602080860151865192955091818601919083010160005b858110156111065781810151838201526110ff602082611694565b90506110e4565b505050935250919050565b600061111c82611266565b92915050565b600060208251111561113357600080fd5b60208201519050815160206111489190611645565b6111539060086116c2565b1c919050565b600081600581111561116d5761116d61167e565b141561117c57610b70826111e4565b60028160058111156111905761119061167e565b14156101005760006111a1836111e4565b905080836000018181516111b59190611694565b90525060208301515183511115610b7057600080fd5b60606111da848460008561128e565b90505b9392505050565b602080820151825181019091015160009182805b600a8110156112605783811a91506112118160076116c2565b82607f16901b85179450816080166000141561124e57611232816001611694565b86518790611241908390611694565b9052509395945050505050565b80611258816116e1565b9150506111f8565b50600080fd5b6000815160141461127657600080fd5b50602001516c01000000000000000000000000900490565b6060824710156113065760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102d6565b843b6113545760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102d6565b600080866001600160a01b031685876040516113709190611728565b60006040518083038185875af1925050503d80600081146113ad576040519150601f19603f3d011682016040523d82523d6000602084013e6113b2565b606091505b50915091506113c28282866113cd565b979650505050505050565b606083156113dc5750816111dd565b8251156113ec5782518084602001fd5b8160405162461bcd60e51b81526004016102d69190611744565b60006020828403121561141857600080fd5b5035919050565b6001600160a01b038116811461052c57600080fd5b60006020828403121561144657600080fd5b81356111dd8161141f565b6000806000806040858703121561146757600080fd5b843567ffffffffffffffff8082111561147f57600080fd5b818701915087601f83011261149357600080fd5b8135818111156114a257600080fd5b8860208285010111156114b457600080fd5b6020928301965094509086013590808211156114cf57600080fd5b818701915087601f8301126114e357600080fd5b8135818111156114f257600080fd5b8860208260051b850101111561150757600080fd5b95989497505060200194505050565b60006020828403121561152857600080fd5b81516111dd8161141f565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000611570604083018688611533565b602083820381850152818583528183019050818660051b8401018760005b888110156115fd57858303601f190184528135368b9003601e190181126115b457600080fd5b8a01803567ffffffffffffffff8111156115cd57600080fd5b8036038c13156115dc57600080fd5b6115e98582898501611533565b95870195945050509084019060010161158e565b50909a9950505050505050505050565b60006020828403121561161f57600080fd5b815180151581146111dd57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000828210156116575761165761162f565b500390565b60008261167957634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b600082198211156116a7576116a761162f565b500190565b634e487b7160e01b600052604160045260246000fd5b60008160001904831182151516156116dc576116dc61162f565b500290565b60006000198214156116f5576116f561162f565b5060010190565b60005b838110156117175781810151838201526020016116ff565b83811115610b3a5750506000910152565b6000825161173a8184602087016116fc565b9190910192915050565b60208152600082518060208401526117638160408501602087016116fc565b601f01601f1916919091016040019291505056fea26469706673582212205b21c9c62f655398f9281c942242859c4f108c2b3d8f5f91bc803a62ecf2f35764736f6c63430008090033",
}

// StakingRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingRewardMetaData.ABI instead.
var StakingRewardABI = StakingRewardMetaData.ABI

// StakingRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingRewardMetaData.Bin instead.
var StakingRewardBin = StakingRewardMetaData.Bin

// DeployStakingReward deploys a new Ethereum contract, binding an instance of StakingReward to it.
func DeployStakingReward(auth *bind.TransactOpts, backend bind.ContractBackend, _staking common.Address) (common.Address, *types.Transaction, *StakingReward, error) {
	parsed, err := StakingRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingRewardBin), backend, _staking)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakingReward{StakingRewardCaller: StakingRewardCaller{contract: contract}, StakingRewardTransactor: StakingRewardTransactor{contract: contract}, StakingRewardFilterer: StakingRewardFilterer{contract: contract}}, nil
}

// StakingReward is an auto generated Go binding around an Ethereum contract.
type StakingReward struct {
	StakingRewardCaller     // Read-only binding to the contract
	StakingRewardTransactor // Write-only binding to the contract
	StakingRewardFilterer   // Log filterer for contract events
}

// StakingRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingRewardSession struct {
	Contract     *StakingReward    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingRewardCallerSession struct {
	Contract *StakingRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingRewardTransactorSession struct {
	Contract     *StakingRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRewardRaw struct {
	Contract *StakingReward // Generic contract binding to access the raw methods on
}

// StakingRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingRewardCallerRaw struct {
	Contract *StakingRewardCaller // Generic read-only contract binding to access the raw methods on
}

// StakingRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingRewardTransactorRaw struct {
	Contract *StakingRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingReward creates a new instance of StakingReward, bound to a specific deployed contract.
func NewStakingReward(address common.Address, backend bind.ContractBackend) (*StakingReward, error) {
	contract, err := bindStakingReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingReward{StakingRewardCaller: StakingRewardCaller{contract: contract}, StakingRewardTransactor: StakingRewardTransactor{contract: contract}, StakingRewardFilterer: StakingRewardFilterer{contract: contract}}, nil
}

// NewStakingRewardCaller creates a new read-only instance of StakingReward, bound to a specific deployed contract.
func NewStakingRewardCaller(address common.Address, caller bind.ContractCaller) (*StakingRewardCaller, error) {
	contract, err := bindStakingReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingRewardCaller{contract: contract}, nil
}

// NewStakingRewardTransactor creates a new write-only instance of StakingReward, bound to a specific deployed contract.
func NewStakingRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingRewardTransactor, error) {
	contract, err := bindStakingReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingRewardTransactor{contract: contract}, nil
}

// NewStakingRewardFilterer creates a new log filterer instance of StakingReward, bound to a specific deployed contract.
func NewStakingRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingRewardFilterer, error) {
	contract, err := bindStakingReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingRewardFilterer{contract: contract}, nil
}

// bindStakingReward binds a generic wrapper to an already deployed contract.
func bindStakingReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingReward *StakingRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingReward.Contract.StakingRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingReward *StakingRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReward.Contract.StakingRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingReward *StakingRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingReward.Contract.StakingRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingReward *StakingRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingReward *StakingRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingReward *StakingRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingReward.Contract.contract.Transact(opts, method, params...)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x96db0fef.
//
// Solidity: function claimedRewardAmounts(address ) view returns(uint256)
func (_StakingReward *StakingRewardCaller) ClaimedRewardAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingReward.contract.Call(opts, &out, "claimedRewardAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x96db0fef.
//
// Solidity: function claimedRewardAmounts(address ) view returns(uint256)
func (_StakingReward *StakingRewardSession) ClaimedRewardAmounts(arg0 common.Address) (*big.Int, error) {
	return _StakingReward.Contract.ClaimedRewardAmounts(&_StakingReward.CallOpts, arg0)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0x96db0fef.
//
// Solidity: function claimedRewardAmounts(address ) view returns(uint256)
func (_StakingReward *StakingRewardCallerSession) ClaimedRewardAmounts(arg0 common.Address) (*big.Int, error) {
	return _StakingReward.Contract.ClaimedRewardAmounts(&_StakingReward.CallOpts, arg0)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_StakingReward *StakingRewardCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingReward.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_StakingReward *StakingRewardSession) IsPauser(account common.Address) (bool, error) {
	return _StakingReward.Contract.IsPauser(&_StakingReward.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_StakingReward *StakingRewardCallerSession) IsPauser(account common.Address) (bool, error) {
	return _StakingReward.Contract.IsPauser(&_StakingReward.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingReward *StakingRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingReward *StakingRewardSession) Owner() (common.Address, error) {
	return _StakingReward.Contract.Owner(&_StakingReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingReward *StakingRewardCallerSession) Owner() (common.Address, error) {
	return _StakingReward.Contract.Owner(&_StakingReward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingReward *StakingRewardCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingReward.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingReward *StakingRewardSession) Paused() (bool, error) {
	return _StakingReward.Contract.Paused(&_StakingReward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingReward *StakingRewardCallerSession) Paused() (bool, error) {
	return _StakingReward.Contract.Paused(&_StakingReward.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_StakingReward *StakingRewardCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StakingReward.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_StakingReward *StakingRewardSession) Pausers(arg0 common.Address) (bool, error) {
	return _StakingReward.Contract.Pausers(&_StakingReward.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_StakingReward *StakingRewardCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _StakingReward.Contract.Pausers(&_StakingReward.CallOpts, arg0)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_StakingReward *StakingRewardCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingReward.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_StakingReward *StakingRewardSession) Staking() (common.Address, error) {
	return _StakingReward.Contract.Staking(&_StakingReward.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_StakingReward *StakingRewardCallerSession) Staking() (common.Address, error) {
	return _StakingReward.Contract.Staking(&_StakingReward.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_StakingReward *StakingRewardTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_StakingReward *StakingRewardSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _StakingReward.Contract.AddPauser(&_StakingReward.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_StakingReward *StakingRewardTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _StakingReward.Contract.AddPauser(&_StakingReward.TransactOpts, account)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_StakingReward *StakingRewardTransactor) ClaimReward(opts *bind.TransactOpts, _rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "claimReward", _rewardRequest, _sigs)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_StakingReward *StakingRewardSession) ClaimReward(_rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _StakingReward.Contract.ClaimReward(&_StakingReward.TransactOpts, _rewardRequest, _sigs)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xf8df0dc5.
//
// Solidity: function claimReward(bytes _rewardRequest, bytes[] _sigs) returns()
func (_StakingReward *StakingRewardTransactorSession) ClaimReward(_rewardRequest []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _StakingReward.Contract.ClaimReward(&_StakingReward.TransactOpts, _rewardRequest, _sigs)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_StakingReward *StakingRewardTransactor) ContributeToRewardPool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "contributeToRewardPool", _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_StakingReward *StakingRewardSession) ContributeToRewardPool(_amount *big.Int) (*types.Transaction, error) {
	return _StakingReward.Contract.ContributeToRewardPool(&_StakingReward.TransactOpts, _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_StakingReward *StakingRewardTransactorSession) ContributeToRewardPool(_amount *big.Int) (*types.Transaction, error) {
	return _StakingReward.Contract.ContributeToRewardPool(&_StakingReward.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_StakingReward *StakingRewardTransactor) DrainToken(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "drainToken", _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_StakingReward *StakingRewardSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _StakingReward.Contract.DrainToken(&_StakingReward.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_StakingReward *StakingRewardTransactorSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _StakingReward.Contract.DrainToken(&_StakingReward.TransactOpts, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingReward *StakingRewardTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingReward *StakingRewardSession) Pause() (*types.Transaction, error) {
	return _StakingReward.Contract.Pause(&_StakingReward.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingReward *StakingRewardTransactorSession) Pause() (*types.Transaction, error) {
	return _StakingReward.Contract.Pause(&_StakingReward.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_StakingReward *StakingRewardTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_StakingReward *StakingRewardSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _StakingReward.Contract.RemovePauser(&_StakingReward.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_StakingReward *StakingRewardTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _StakingReward.Contract.RemovePauser(&_StakingReward.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingReward *StakingRewardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingReward *StakingRewardSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingReward.Contract.RenounceOwnership(&_StakingReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingReward *StakingRewardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingReward.Contract.RenounceOwnership(&_StakingReward.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_StakingReward *StakingRewardTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_StakingReward *StakingRewardSession) RenouncePauser() (*types.Transaction, error) {
	return _StakingReward.Contract.RenouncePauser(&_StakingReward.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_StakingReward *StakingRewardTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _StakingReward.Contract.RenouncePauser(&_StakingReward.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingReward *StakingRewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingReward *StakingRewardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingReward.Contract.TransferOwnership(&_StakingReward.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingReward *StakingRewardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingReward.Contract.TransferOwnership(&_StakingReward.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingReward *StakingRewardTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingReward.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingReward *StakingRewardSession) Unpause() (*types.Transaction, error) {
	return _StakingReward.Contract.Unpause(&_StakingReward.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingReward *StakingRewardTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakingReward.Contract.Unpause(&_StakingReward.TransactOpts)
}

// StakingRewardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingReward contract.
type StakingRewardOwnershipTransferredIterator struct {
	Event *StakingRewardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingRewardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardOwnershipTransferred)
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
		it.Event = new(StakingRewardOwnershipTransferred)
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
func (it *StakingRewardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardOwnershipTransferred represents a OwnershipTransferred event raised by the StakingReward contract.
type StakingRewardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingReward *StakingRewardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingRewardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingReward.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardOwnershipTransferredIterator{contract: _StakingReward.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingReward *StakingRewardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingRewardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingReward.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardOwnershipTransferred)
				if err := _StakingReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingReward *StakingRewardFilterer) ParseOwnershipTransferred(log types.Log) (*StakingRewardOwnershipTransferred, error) {
	event := new(StakingRewardOwnershipTransferred)
	if err := _StakingReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingReward contract.
type StakingRewardPausedIterator struct {
	Event *StakingRewardPaused // Event containing the contract specifics and raw log

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
func (it *StakingRewardPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardPaused)
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
		it.Event = new(StakingRewardPaused)
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
func (it *StakingRewardPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardPaused represents a Paused event raised by the StakingReward contract.
type StakingRewardPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingReward *StakingRewardFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingRewardPausedIterator, error) {

	logs, sub, err := _StakingReward.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingRewardPausedIterator{contract: _StakingReward.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingReward *StakingRewardFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingRewardPaused) (event.Subscription, error) {

	logs, sub, err := _StakingReward.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardPaused)
				if err := _StakingReward.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingReward *StakingRewardFilterer) ParsePaused(log types.Log) (*StakingRewardPaused, error) {
	event := new(StakingRewardPaused)
	if err := _StakingReward.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the StakingReward contract.
type StakingRewardPauserAddedIterator struct {
	Event *StakingRewardPauserAdded // Event containing the contract specifics and raw log

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
func (it *StakingRewardPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardPauserAdded)
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
		it.Event = new(StakingRewardPauserAdded)
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
func (it *StakingRewardPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardPauserAdded represents a PauserAdded event raised by the StakingReward contract.
type StakingRewardPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_StakingReward *StakingRewardFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*StakingRewardPauserAddedIterator, error) {

	logs, sub, err := _StakingReward.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &StakingRewardPauserAddedIterator{contract: _StakingReward.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_StakingReward *StakingRewardFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *StakingRewardPauserAdded) (event.Subscription, error) {

	logs, sub, err := _StakingReward.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardPauserAdded)
				if err := _StakingReward.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_StakingReward *StakingRewardFilterer) ParsePauserAdded(log types.Log) (*StakingRewardPauserAdded, error) {
	event := new(StakingRewardPauserAdded)
	if err := _StakingReward.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the StakingReward contract.
type StakingRewardPauserRemovedIterator struct {
	Event *StakingRewardPauserRemoved // Event containing the contract specifics and raw log

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
func (it *StakingRewardPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardPauserRemoved)
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
		it.Event = new(StakingRewardPauserRemoved)
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
func (it *StakingRewardPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardPauserRemoved represents a PauserRemoved event raised by the StakingReward contract.
type StakingRewardPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_StakingReward *StakingRewardFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*StakingRewardPauserRemovedIterator, error) {

	logs, sub, err := _StakingReward.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &StakingRewardPauserRemovedIterator{contract: _StakingReward.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_StakingReward *StakingRewardFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *StakingRewardPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _StakingReward.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardPauserRemoved)
				if err := _StakingReward.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_StakingReward *StakingRewardFilterer) ParsePauserRemoved(log types.Log) (*StakingRewardPauserRemoved, error) {
	event := new(StakingRewardPauserRemoved)
	if err := _StakingReward.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardStakingRewardClaimedIterator is returned from FilterStakingRewardClaimed and is used to iterate over the raw logs and unpacked data for StakingRewardClaimed events raised by the StakingReward contract.
type StakingRewardStakingRewardClaimedIterator struct {
	Event *StakingRewardStakingRewardClaimed // Event containing the contract specifics and raw log

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
func (it *StakingRewardStakingRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardStakingRewardClaimed)
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
		it.Event = new(StakingRewardStakingRewardClaimed)
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
func (it *StakingRewardStakingRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardStakingRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardStakingRewardClaimed represents a StakingRewardClaimed event raised by the StakingReward contract.
type StakingRewardStakingRewardClaimed struct {
	Recipient common.Address
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardClaimed is a free log retrieval operation binding the contract event 0x6dd401e61ba732582a5eba3d54ccc3afb3609cd3ac1a166d1d36f75fc0aedcda.
//
// Solidity: event StakingRewardClaimed(address indexed recipient, uint256 reward)
func (_StakingReward *StakingRewardFilterer) FilterStakingRewardClaimed(opts *bind.FilterOpts, recipient []common.Address) (*StakingRewardStakingRewardClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakingReward.contract.FilterLogs(opts, "StakingRewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardStakingRewardClaimedIterator{contract: _StakingReward.contract, event: "StakingRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardClaimed is a free log subscription operation binding the contract event 0x6dd401e61ba732582a5eba3d54ccc3afb3609cd3ac1a166d1d36f75fc0aedcda.
//
// Solidity: event StakingRewardClaimed(address indexed recipient, uint256 reward)
func (_StakingReward *StakingRewardFilterer) WatchStakingRewardClaimed(opts *bind.WatchOpts, sink chan<- *StakingRewardStakingRewardClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _StakingReward.contract.WatchLogs(opts, "StakingRewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardStakingRewardClaimed)
				if err := _StakingReward.contract.UnpackLog(event, "StakingRewardClaimed", log); err != nil {
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

// ParseStakingRewardClaimed is a log parse operation binding the contract event 0x6dd401e61ba732582a5eba3d54ccc3afb3609cd3ac1a166d1d36f75fc0aedcda.
//
// Solidity: event StakingRewardClaimed(address indexed recipient, uint256 reward)
func (_StakingReward *StakingRewardFilterer) ParseStakingRewardClaimed(log types.Log) (*StakingRewardStakingRewardClaimed, error) {
	event := new(StakingRewardStakingRewardClaimed)
	if err := _StakingReward.contract.UnpackLog(event, "StakingRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardStakingRewardContributedIterator is returned from FilterStakingRewardContributed and is used to iterate over the raw logs and unpacked data for StakingRewardContributed events raised by the StakingReward contract.
type StakingRewardStakingRewardContributedIterator struct {
	Event *StakingRewardStakingRewardContributed // Event containing the contract specifics and raw log

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
func (it *StakingRewardStakingRewardContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardStakingRewardContributed)
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
		it.Event = new(StakingRewardStakingRewardContributed)
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
func (it *StakingRewardStakingRewardContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardStakingRewardContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardStakingRewardContributed represents a StakingRewardContributed event raised by the StakingReward contract.
type StakingRewardStakingRewardContributed struct {
	Contributor  common.Address
	Contribution *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardContributed is a free log retrieval operation binding the contract event 0xf67017a05194c0853be9169be60cad9fa6e75d34b6b507a7a4261510e19c6d79.
//
// Solidity: event StakingRewardContributed(address indexed contributor, uint256 contribution)
func (_StakingReward *StakingRewardFilterer) FilterStakingRewardContributed(opts *bind.FilterOpts, contributor []common.Address) (*StakingRewardStakingRewardContributedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _StakingReward.contract.FilterLogs(opts, "StakingRewardContributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardStakingRewardContributedIterator{contract: _StakingReward.contract, event: "StakingRewardContributed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardContributed is a free log subscription operation binding the contract event 0xf67017a05194c0853be9169be60cad9fa6e75d34b6b507a7a4261510e19c6d79.
//
// Solidity: event StakingRewardContributed(address indexed contributor, uint256 contribution)
func (_StakingReward *StakingRewardFilterer) WatchStakingRewardContributed(opts *bind.WatchOpts, sink chan<- *StakingRewardStakingRewardContributed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _StakingReward.contract.WatchLogs(opts, "StakingRewardContributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardStakingRewardContributed)
				if err := _StakingReward.contract.UnpackLog(event, "StakingRewardContributed", log); err != nil {
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

// ParseStakingRewardContributed is a log parse operation binding the contract event 0xf67017a05194c0853be9169be60cad9fa6e75d34b6b507a7a4261510e19c6d79.
//
// Solidity: event StakingRewardContributed(address indexed contributor, uint256 contribution)
func (_StakingReward *StakingRewardFilterer) ParseStakingRewardContributed(log types.Log) (*StakingRewardStakingRewardContributed, error) {
	event := new(StakingRewardStakingRewardContributed)
	if err := _StakingReward.contract.UnpackLog(event, "StakingRewardContributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingReward contract.
type StakingRewardUnpausedIterator struct {
	Event *StakingRewardUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingRewardUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardUnpaused)
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
		it.Event = new(StakingRewardUnpaused)
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
func (it *StakingRewardUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardUnpaused represents a Unpaused event raised by the StakingReward contract.
type StakingRewardUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingReward *StakingRewardFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingRewardUnpausedIterator, error) {

	logs, sub, err := _StakingReward.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingRewardUnpausedIterator{contract: _StakingReward.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingReward *StakingRewardFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingRewardUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingReward.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardUnpaused)
				if err := _StakingReward.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingReward *StakingRewardFilterer) ParseUnpaused(log types.Log) (*StakingRewardUnpaused, error) {
	event := new(StakingRewardUnpaused)
	if err := _StakingReward.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ViewerMetaData contains all meta data concerning the Viewer contract.
var ViewerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getBondedValidatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"internalType\":\"structDataTypes.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawableUndelegationTokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.DelegatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"internalType\":\"structDataTypes.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"internalType\":\"structDataTypes.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"shouldBondValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516117c73803806117c783398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516116db6100ec6000396000818160bb01528181610199015281816102eb0152818161039701528181610633015281816107790152818161089e0152818161099b01528181610a4401528181610b0001528181610cf401528181610e3a01528181610ecf01528181610fec015261109201526116db6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638dc2336d1161005b5780638dc2336d1461012a578063c6fc1ed614610140578063d87ffe9114610168578063e9fe6b0b1461017057600080fd5b8063313019bb1461008d5780634cf088d9146100b657806366ab5d28146100f55780638a11d7c91461010a575b600080fd5b6100a061009b36600461115b565b610193565b6040516100ad919061117f565b60405180910390f35b6100dd7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ad565b6100fd61062d565b6040516100ad91906112d7565b61011d61011836600461115b565b610839565b6040516100ad9190611319565b610132610996565b6040519081526020016100ad565b61015361014e36600461115b565b610c49565b604080519283526020830191909152016100ad565b6100fd610cee565b61018361017e36600461115b565b610ea7565b60405190151581526020016100ad565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631cfe4f0b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101f057600080fd5b505afa158015610204573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610228919061132d565b905060008167ffffffffffffffff81111561024557610245611346565b6040519080825280602002602001820160405280156102b857816020015b6102a56040518060c0016040528060006001600160a01b0316815260200160008152602001600081526020016060815260200160008152602001600081525090565b8152602001906001900390816102635790505b5090506000805b838163ffffffff1610156104b7576040516324aec90f60e21b815263ffffffff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906392bb243c9060240160206040518083038186803b15801561033557600080fd5b505afa158015610349573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036d919061135c565b604051631dd9dfdf60e31b81526001600160a01b03808316600483015289811660248301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063eecefef89060440160006040518083038186803b1580156103db57600080fd5b505afa1580156103ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261041791908101906113f6565b848363ffffffff168151811061042f5761042f611527565b6020026020010181905250838263ffffffff168151811061045257610452611527565b60200260200101516040015160001415806104915750838263ffffffff168151811061048057610480611527565b602002602001015160800151600014155b156104a457826104a081611553565b9350505b50806104af81611553565b9150506102bf565b5060008163ffffffff1667ffffffffffffffff8111156104d9576104d9611346565b60405190808252806020026020018201604052801561054c57816020015b6105396040518060c0016040528060006001600160a01b0316815260200160008152602001600081526020016060815260200160008152602001600081525090565b8152602001906001900390816104f75790505b5090506000805b858163ffffffff16101561062157848163ffffffff168151811061057957610579611527565b60200260200101516040015160001415806105b85750848163ffffffff16815181106105a7576105a7611527565b602002602001015160800151600014155b1561060f57848163ffffffff16815181106105d5576105d5611527565b6020026020010151838363ffffffff16815181106105f5576105f5611527565b6020026020010181905250818061060b90611553565b9250505b8061061981611553565b915050610553565b50909695505050505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631cfe4f0b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561068a57600080fd5b505afa15801561069e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c2919061132d565b905060008167ffffffffffffffff8111156106df576106df611346565b60405190808252806020026020018201604052801561074657816020015b6040805160e08101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282526000199092019101816106fd5790505b50905060005b828163ffffffff161015610832576040516324aec90f60e21b815263ffffffff821660048201526107fc907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906392bb243c906024015b60206040518083038186803b1580156107c457600080fd5b505afa1580156107d8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610118919061135c565b828263ffffffff168151811061081457610814611527565b6020026020010181905250808061082a90611553565b91505061074c565b5092915050565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152604051631f4a58fb60e31b81526001600160a01b038381166004830152600091829182918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063fa52c7d8906024016101406040518083038186803b1580156108e157600080fd5b505afa1580156108f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109199190611594565b995050509750505095509550955095506040518060e00160405280896001600160a01b0316815260200187600381111561095557610955611254565b8152602001866001600160a01b031681526020018581526020018481526020018381526020018267ffffffffffffffff168152509650505050505050919050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166389f9aab56040518163ffffffff1660e01b815260040160206040518083038186803b1580156109f257600080fd5b505afa158015610a06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2a919061132d565b60405163eb505dd560e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb505dd590610a7a90600390600401611636565b60206040518083038186803b158015610a9257600080fd5b505afa158015610aa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aca919061132d565b811015610ad957600091505090565b60001960005b828110156108325760405163acc62ccf60e01b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063c8f9f98490829063acc62ccf9060240160206040518083038186803b158015610b5257600080fd5b505afa158015610b66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8a919061135c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015260240160206040518083038186803b158015610be157600080fd5b505afa158015610bf5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c19919061132d565b905082811015610c365791508180610c3657600094505050505090565b5080610c4181611650565b915050610adf565b6000806000610c5784610193565b905060008060005b83518163ffffffff161015610ce257838163ffffffff1681518110610c8657610c86611527565b60200260200101516020015183610c9d919061166b565b9250838163ffffffff1681518110610cb757610cb7611527565b60200260200101516080015182610cce919061166b565b915080610cda81611553565b915050610c5f565b50909590945092505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166389f9aab56040518163ffffffff1660e01b815260040160206040518083038186803b158015610d4b57600080fd5b505afa158015610d5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d83919061132d565b905060008167ffffffffffffffff811115610da057610da0611346565b604051908082528060200260200182016040528015610e0757816020015b6040805160e08101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c08201528252600019909201910181610dbe5790505b50905060005b828163ffffffff1610156108325760405163acc62ccf60e01b815263ffffffff82166004820152610e71907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063acc62ccf906024016107ac565b828263ffffffff1681518110610e8957610e89611527565b60200260200101819052508080610e9f90611553565b915050610e0d565b604051631f4a58fb60e31b81526001600160a01b0382811660048301526000918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063fa52c7d8906024016101406040518083038186803b158015610f1257600080fd5b505afa158015610f26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4a9190611594565b5050975050505050935050925060006003811115610f6a57610f6a611254565b836003811115610f7c57610f7c611254565b1480610f9957506003836003811115610f9757610f97611254565b145b15610fa957506000949350505050565b8067ffffffffffffffff16431015610fc657506000949350505050565b6040516347abfdbf60e01b81526001600160a01b038681166004830152600160248301527f000000000000000000000000000000000000000000000000000000000000000016906347abfdbf9060440160206040518083038186803b15801561102e57600080fd5b505afa158015611042573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110669190611683565b61107557506000949350505050565b61107d610996565b821161108e57506000949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166383cfb3186040518163ffffffff1660e01b815260040160206040518083038186803b1580156110e957600080fd5b505afa1580156110fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611121919061132d565b9050804310156111375750600095945050505050565b50600195945050505050565b6001600160a01b038116811461115857600080fd5b50565b60006020828403121561116d57600080fd5b813561117881611143565b9392505050565b60006020808301818452808551808352604092508286019150828160051b8701018488016000805b8481101561124557898403603f19018652825180516001600160a01b031685528881015189860152878101518886015260608082015160c0918701829052805191870182905260e0870191908b0190855b8181101561121d578251805185528d01518d850152928b0192918c01916001016111f8565b5050506080828101519087015260a091820151919095015294870194918701916001016111a7565b50919998505050505050505050565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b0380825116835260208201516004811061128d5761128d611254565b8060208501525080604083015116604084015250606081015160608301526080810151608083015260a081015160a083015267ffffffffffffffff60c08201511660c08301525050565b6020808252825182820181905260009190848201906040850190845b818110156106215761130683855161126a565b9284019260e092909201916001016112f3565b60e08101611327828461126a565b92915050565b60006020828403121561133f57600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561136e57600080fd5b815161117881611143565b60405160c0810167ffffffffffffffff8111828210171561139c5761139c611346565b60405290565b6040805190810167ffffffffffffffff8111828210171561139c5761139c611346565b604051601f8201601f1916810167ffffffffffffffff811182821017156113ee576113ee611346565b604052919050565b6000602080838503121561140957600080fd5b825167ffffffffffffffff8082111561142157600080fd5b9084019060c0828703121561143557600080fd5b61143d611379565b825161144881611143565b815282840151848201526040808401518183015260608401518381111561146e57600080fd5b8401601f8101891361147f57600080fd5b80518481111561149157611491611346565b61149f878260051b016113c5565b818152878101955060069190911b82018701908a8211156114bf57600080fd5b918701915b818310156114ff5783838c0312156114dc5760008081fd5b6114e46113a2565b835181528884015189820152865294870194918301916114c4565b60608501525050506080838101519082015260a0928301519281019290925250949350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168181141561156d5761156d61153d565b6001019392505050565b805167ffffffffffffffff8116811461158f57600080fd5b919050565b6000806000806000806000806000806101408b8d0312156115b457600080fd5b8a51600481106115c357600080fd5b60208c0151909a506115d481611143565b8099505060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935061160760e08c01611577565b92506116166101008c01611577565b91506116256101208c01611577565b90509295989b9194979a5092959850565b602081016009831061164a5761164a611254565b91905290565b60006000198214156116645761166461153d565b5060010190565b6000821982111561167e5761167e61153d565b500190565b60006020828403121561169557600080fd5b8151801515811461117857600080fdfea264697066735822122085a78e0919a496fbd233f10407bbd7764e6441f58ad996972fb639656b4f687164736f6c63430008090033",
}

// ViewerABI is the input ABI used to generate the binding from.
// Deprecated: Use ViewerMetaData.ABI instead.
var ViewerABI = ViewerMetaData.ABI

// ViewerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ViewerMetaData.Bin instead.
var ViewerBin = ViewerMetaData.Bin

// DeployViewer deploys a new Ethereum contract, binding an instance of Viewer to it.
func DeployViewer(auth *bind.TransactOpts, backend bind.ContractBackend, _staking common.Address) (common.Address, *types.Transaction, *Viewer, error) {
	parsed, err := ViewerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ViewerBin), backend, _staking)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Viewer{ViewerCaller: ViewerCaller{contract: contract}, ViewerTransactor: ViewerTransactor{contract: contract}, ViewerFilterer: ViewerFilterer{contract: contract}}, nil
}

// Viewer is an auto generated Go binding around an Ethereum contract.
type Viewer struct {
	ViewerCaller     // Read-only binding to the contract
	ViewerTransactor // Write-only binding to the contract
	ViewerFilterer   // Log filterer for contract events
}

// ViewerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ViewerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ViewerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ViewerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ViewerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ViewerSession struct {
	Contract     *Viewer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViewerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ViewerCallerSession struct {
	Contract *ViewerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ViewerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ViewerTransactorSession struct {
	Contract     *ViewerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ViewerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ViewerRaw struct {
	Contract *Viewer // Generic contract binding to access the raw methods on
}

// ViewerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ViewerCallerRaw struct {
	Contract *ViewerCaller // Generic read-only contract binding to access the raw methods on
}

// ViewerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ViewerTransactorRaw struct {
	Contract *ViewerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewViewer creates a new instance of Viewer, bound to a specific deployed contract.
func NewViewer(address common.Address, backend bind.ContractBackend) (*Viewer, error) {
	contract, err := bindViewer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Viewer{ViewerCaller: ViewerCaller{contract: contract}, ViewerTransactor: ViewerTransactor{contract: contract}, ViewerFilterer: ViewerFilterer{contract: contract}}, nil
}

// NewViewerCaller creates a new read-only instance of Viewer, bound to a specific deployed contract.
func NewViewerCaller(address common.Address, caller bind.ContractCaller) (*ViewerCaller, error) {
	contract, err := bindViewer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ViewerCaller{contract: contract}, nil
}

// NewViewerTransactor creates a new write-only instance of Viewer, bound to a specific deployed contract.
func NewViewerTransactor(address common.Address, transactor bind.ContractTransactor) (*ViewerTransactor, error) {
	contract, err := bindViewer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ViewerTransactor{contract: contract}, nil
}

// NewViewerFilterer creates a new log filterer instance of Viewer, bound to a specific deployed contract.
func NewViewerFilterer(address common.Address, filterer bind.ContractFilterer) (*ViewerFilterer, error) {
	contract, err := bindViewer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ViewerFilterer{contract: contract}, nil
}

// bindViewer binds a generic wrapper to an already deployed contract.
func bindViewer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ViewerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Viewer *ViewerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Viewer.Contract.ViewerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Viewer *ViewerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Viewer.Contract.ViewerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Viewer *ViewerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Viewer.Contract.ViewerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Viewer *ViewerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Viewer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Viewer *ViewerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Viewer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Viewer *ViewerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Viewer.Contract.contract.Transact(opts, method, params...)
}

// GetBondedValidatorInfos is a free data retrieval call binding the contract method 0xd87ffe91.
//
// Solidity: function getBondedValidatorInfos() view returns((address,uint8,address,uint256,uint256,uint256,uint64)[])
func (_Viewer *ViewerCaller) GetBondedValidatorInfos(opts *bind.CallOpts) ([]DataTypesValidatorInfo, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "getBondedValidatorInfos")

	if err != nil {
		return *new([]DataTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesValidatorInfo)).(*[]DataTypesValidatorInfo)

	return out0, err

}

// GetBondedValidatorInfos is a free data retrieval call binding the contract method 0xd87ffe91.
//
// Solidity: function getBondedValidatorInfos() view returns((address,uint8,address,uint256,uint256,uint256,uint64)[])
func (_Viewer *ViewerSession) GetBondedValidatorInfos() ([]DataTypesValidatorInfo, error) {
	return _Viewer.Contract.GetBondedValidatorInfos(&_Viewer.CallOpts)
}

// GetBondedValidatorInfos is a free data retrieval call binding the contract method 0xd87ffe91.
//
// Solidity: function getBondedValidatorInfos() view returns((address,uint8,address,uint256,uint256,uint256,uint64)[])
func (_Viewer *ViewerCallerSession) GetBondedValidatorInfos() ([]DataTypesValidatorInfo, error) {
	return _Viewer.Contract.GetBondedValidatorInfos(&_Viewer.CallOpts)
}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,(uint256,uint256)[],uint256,uint256)[])
func (_Viewer *ViewerCaller) GetDelegatorInfos(opts *bind.CallOpts, _delAddr common.Address) ([]DataTypesDelegatorInfo, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "getDelegatorInfos", _delAddr)

	if err != nil {
		return *new([]DataTypesDelegatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesDelegatorInfo)).(*[]DataTypesDelegatorInfo)

	return out0, err

}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,(uint256,uint256)[],uint256,uint256)[])
func (_Viewer *ViewerSession) GetDelegatorInfos(_delAddr common.Address) ([]DataTypesDelegatorInfo, error) {
	return _Viewer.Contract.GetDelegatorInfos(&_Viewer.CallOpts, _delAddr)
}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,(uint256,uint256)[],uint256,uint256)[])
func (_Viewer *ViewerCallerSession) GetDelegatorInfos(_delAddr common.Address) ([]DataTypesDelegatorInfo, error) {
	return _Viewer.Contract.GetDelegatorInfos(&_Viewer.CallOpts, _delAddr)
}

// GetDelegatorTokens is a free data retrieval call binding the contract method 0xc6fc1ed6.
//
// Solidity: function getDelegatorTokens(address _delAddr) view returns(uint256, uint256)
func (_Viewer *ViewerCaller) GetDelegatorTokens(opts *bind.CallOpts, _delAddr common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "getDelegatorTokens", _delAddr)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDelegatorTokens is a free data retrieval call binding the contract method 0xc6fc1ed6.
//
// Solidity: function getDelegatorTokens(address _delAddr) view returns(uint256, uint256)
func (_Viewer *ViewerSession) GetDelegatorTokens(_delAddr common.Address) (*big.Int, *big.Int, error) {
	return _Viewer.Contract.GetDelegatorTokens(&_Viewer.CallOpts, _delAddr)
}

// GetDelegatorTokens is a free data retrieval call binding the contract method 0xc6fc1ed6.
//
// Solidity: function getDelegatorTokens(address _delAddr) view returns(uint256, uint256)
func (_Viewer *ViewerCallerSession) GetDelegatorTokens(_delAddr common.Address) (*big.Int, *big.Int, error) {
	return _Viewer.Contract.GetDelegatorTokens(&_Viewer.CallOpts, _delAddr)
}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_Viewer *ViewerCaller) GetMinValidatorTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "getMinValidatorTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_Viewer *ViewerSession) GetMinValidatorTokens() (*big.Int, error) {
	return _Viewer.Contract.GetMinValidatorTokens(&_Viewer.CallOpts)
}

// GetMinValidatorTokens is a free data retrieval call binding the contract method 0x8dc2336d.
//
// Solidity: function getMinValidatorTokens() view returns(uint256)
func (_Viewer *ViewerCallerSession) GetMinValidatorTokens() (*big.Int, error) {
	return _Viewer.Contract.GetMinValidatorTokens(&_Viewer.CallOpts)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _valAddr) view returns((address,uint8,address,uint256,uint256,uint256,uint64))
func (_Viewer *ViewerCaller) GetValidatorInfo(opts *bind.CallOpts, _valAddr common.Address) (DataTypesValidatorInfo, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "getValidatorInfo", _valAddr)

	if err != nil {
		return *new(DataTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesValidatorInfo)).(*DataTypesValidatorInfo)

	return out0, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _valAddr) view returns((address,uint8,address,uint256,uint256,uint256,uint64))
func (_Viewer *ViewerSession) GetValidatorInfo(_valAddr common.Address) (DataTypesValidatorInfo, error) {
	return _Viewer.Contract.GetValidatorInfo(&_Viewer.CallOpts, _valAddr)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _valAddr) view returns((address,uint8,address,uint256,uint256,uint256,uint64))
func (_Viewer *ViewerCallerSession) GetValidatorInfo(_valAddr common.Address) (DataTypesValidatorInfo, error) {
	return _Viewer.Contract.GetValidatorInfo(&_Viewer.CallOpts, _valAddr)
}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x66ab5d28.
//
// Solidity: function getValidatorInfos() view returns((address,uint8,address,uint256,uint256,uint256,uint64)[])
func (_Viewer *ViewerCaller) GetValidatorInfos(opts *bind.CallOpts) ([]DataTypesValidatorInfo, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "getValidatorInfos")

	if err != nil {
		return *new([]DataTypesValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesValidatorInfo)).(*[]DataTypesValidatorInfo)

	return out0, err

}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x66ab5d28.
//
// Solidity: function getValidatorInfos() view returns((address,uint8,address,uint256,uint256,uint256,uint64)[])
func (_Viewer *ViewerSession) GetValidatorInfos() ([]DataTypesValidatorInfo, error) {
	return _Viewer.Contract.GetValidatorInfos(&_Viewer.CallOpts)
}

// GetValidatorInfos is a free data retrieval call binding the contract method 0x66ab5d28.
//
// Solidity: function getValidatorInfos() view returns((address,uint8,address,uint256,uint256,uint256,uint64)[])
func (_Viewer *ViewerCallerSession) GetValidatorInfos() ([]DataTypesValidatorInfo, error) {
	return _Viewer.Contract.GetValidatorInfos(&_Viewer.CallOpts)
}

// ShouldBondValidator is a free data retrieval call binding the contract method 0xe9fe6b0b.
//
// Solidity: function shouldBondValidator(address _valAddr) view returns(bool)
func (_Viewer *ViewerCaller) ShouldBondValidator(opts *bind.CallOpts, _valAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "shouldBondValidator", _valAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldBondValidator is a free data retrieval call binding the contract method 0xe9fe6b0b.
//
// Solidity: function shouldBondValidator(address _valAddr) view returns(bool)
func (_Viewer *ViewerSession) ShouldBondValidator(_valAddr common.Address) (bool, error) {
	return _Viewer.Contract.ShouldBondValidator(&_Viewer.CallOpts, _valAddr)
}

// ShouldBondValidator is a free data retrieval call binding the contract method 0xe9fe6b0b.
//
// Solidity: function shouldBondValidator(address _valAddr) view returns(bool)
func (_Viewer *ViewerCallerSession) ShouldBondValidator(_valAddr common.Address) (bool, error) {
	return _Viewer.Contract.ShouldBondValidator(&_Viewer.CallOpts, _valAddr)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Viewer *ViewerCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Viewer.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Viewer *ViewerSession) Staking() (common.Address, error) {
	return _Viewer.Contract.Staking(&_Viewer.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_Viewer *ViewerCallerSession) Staking() (common.Address, error) {
	return _Viewer.Contract.Staking(&_Viewer.CallOpts)
}

// WhitelistMetaData contains all meta data concerning the Whitelist contract.
var WhitelistMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WhitelistABI is the input ABI used to generate the binding from.
// Deprecated: Use WhitelistMetaData.ABI instead.
var WhitelistABI = WhitelistMetaData.ABI

// Whitelist is an auto generated Go binding around an Ethereum contract.
type Whitelist struct {
	WhitelistCaller     // Read-only binding to the contract
	WhitelistTransactor // Write-only binding to the contract
	WhitelistFilterer   // Log filterer for contract events
}

// WhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistSession struct {
	Contract     *Whitelist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistCallerSession struct {
	Contract *WhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistTransactorSession struct {
	Contract     *WhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistRaw struct {
	Contract *Whitelist // Generic contract binding to access the raw methods on
}

// WhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistCallerRaw struct {
	Contract *WhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistTransactorRaw struct {
	Contract *WhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelist creates a new instance of Whitelist, bound to a specific deployed contract.
func NewWhitelist(address common.Address, backend bind.ContractBackend) (*Whitelist, error) {
	contract, err := bindWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelist{WhitelistCaller: WhitelistCaller{contract: contract}, WhitelistTransactor: WhitelistTransactor{contract: contract}, WhitelistFilterer: WhitelistFilterer{contract: contract}}, nil
}

// NewWhitelistCaller creates a new read-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistCaller(address common.Address, caller bind.ContractCaller) (*WhitelistCaller, error) {
	contract, err := bindWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistCaller{contract: contract}, nil
}

// NewWhitelistTransactor creates a new write-only instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistTransactor, error) {
	contract, err := bindWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistTransactor{contract: contract}, nil
}

// NewWhitelistFilterer creates a new log filterer instance of Whitelist, bound to a specific deployed contract.
func NewWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistFilterer, error) {
	contract, err := bindWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistFilterer{contract: contract}, nil
}

// bindWhitelist binds a generic wrapper to an already deployed contract.
func bindWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.WhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.WhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelist *WhitelistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Whitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelist *WhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelist *WhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelist.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Whitelist *WhitelistCaller) IsWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "isWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Whitelist *WhitelistSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Whitelist.Contract.IsWhitelisted(&_Whitelist.CallOpts, account)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(address account) view returns(bool)
func (_Whitelist *WhitelistCallerSession) IsWhitelisted(account common.Address) (bool, error) {
	return _Whitelist.Contract.IsWhitelisted(&_Whitelist.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Whitelist *WhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Whitelist *WhitelistSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Whitelist *WhitelistCallerSession) Owner() (common.Address, error) {
	return _Whitelist.Contract.Owner(&_Whitelist.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Whitelist *WhitelistCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Whitelist *WhitelistSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Whitelist(&_Whitelist.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Whitelist *WhitelistCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelist.Contract.Whitelist(&_Whitelist.CallOpts, arg0)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Whitelist *WhitelistCaller) WhitelistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Whitelist.contract.Call(opts, &out, "whitelistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Whitelist *WhitelistSession) WhitelistEnabled() (bool, error) {
	return _Whitelist.Contract.WhitelistEnabled(&_Whitelist.CallOpts)
}

// WhitelistEnabled is a free data retrieval call binding the contract method 0x51fb012d.
//
// Solidity: function whitelistEnabled() view returns(bool)
func (_Whitelist *WhitelistCallerSession) WhitelistEnabled() (bool, error) {
	return _Whitelist.Contract.WhitelistEnabled(&_Whitelist.CallOpts)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Whitelist *WhitelistTransactor) AddWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "addWhitelisted", account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Whitelist *WhitelistSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.AddWhitelisted(&_Whitelist.TransactOpts, account)
}

// AddWhitelisted is a paid mutator transaction binding the contract method 0x10154bad.
//
// Solidity: function addWhitelisted(address account) returns()
func (_Whitelist *WhitelistTransactorSession) AddWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.AddWhitelisted(&_Whitelist.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Whitelist *WhitelistTransactor) RemoveWhitelisted(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "removeWhitelisted", account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Whitelist *WhitelistSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.RemoveWhitelisted(&_Whitelist.TransactOpts, account)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0x291d9549.
//
// Solidity: function removeWhitelisted(address account) returns()
func (_Whitelist *WhitelistTransactorSession) RemoveWhitelisted(account common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.RemoveWhitelisted(&_Whitelist.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Whitelist *WhitelistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Whitelist *WhitelistSession) RenounceOwnership() (*types.Transaction, error) {
	return _Whitelist.Contract.RenounceOwnership(&_Whitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Whitelist *WhitelistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Whitelist.Contract.RenounceOwnership(&_Whitelist.TransactOpts)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Whitelist *WhitelistTransactor) SetWhitelistEnabled(opts *bind.TransactOpts, _whitelistEnabled bool) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "setWhitelistEnabled", _whitelistEnabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Whitelist *WhitelistSession) SetWhitelistEnabled(_whitelistEnabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetWhitelistEnabled(&_Whitelist.TransactOpts, _whitelistEnabled)
}

// SetWhitelistEnabled is a paid mutator transaction binding the contract method 0x052d9e7e.
//
// Solidity: function setWhitelistEnabled(bool _whitelistEnabled) returns()
func (_Whitelist *WhitelistTransactorSession) SetWhitelistEnabled(_whitelistEnabled bool) (*types.Transaction, error) {
	return _Whitelist.Contract.SetWhitelistEnabled(&_Whitelist.TransactOpts, _whitelistEnabled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Whitelist *WhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Whitelist *WhitelistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.TransferOwnership(&_Whitelist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Whitelist *WhitelistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Whitelist.Contract.TransferOwnership(&_Whitelist.TransactOpts, newOwner)
}

// WhitelistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Whitelist contract.
type WhitelistOwnershipTransferredIterator struct {
	Event *WhitelistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WhitelistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistOwnershipTransferred)
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
		it.Event = new(WhitelistOwnershipTransferred)
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
func (it *WhitelistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistOwnershipTransferred represents a OwnershipTransferred event raised by the Whitelist contract.
type WhitelistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Whitelist *WhitelistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WhitelistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistOwnershipTransferredIterator{contract: _Whitelist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Whitelist *WhitelistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WhitelistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistOwnershipTransferred)
				if err := _Whitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Whitelist *WhitelistFilterer) ParseOwnershipTransferred(log types.Log) (*WhitelistOwnershipTransferred, error) {
	event := new(WhitelistOwnershipTransferred)
	if err := _Whitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistWhitelistedAddedIterator is returned from FilterWhitelistedAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAdded events raised by the Whitelist contract.
type WhitelistWhitelistedAddedIterator struct {
	Event *WhitelistWhitelistedAdded // Event containing the contract specifics and raw log

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
func (it *WhitelistWhitelistedAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistWhitelistedAdded)
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
		it.Event = new(WhitelistWhitelistedAdded)
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
func (it *WhitelistWhitelistedAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistWhitelistedAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistWhitelistedAdded represents a WhitelistedAdded event raised by the Whitelist contract.
type WhitelistWhitelistedAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAdded is a free log retrieval operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Whitelist *WhitelistFilterer) FilterWhitelistedAdded(opts *bind.FilterOpts) (*WhitelistWhitelistedAddedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return &WhitelistWhitelistedAddedIterator{contract: _Whitelist.contract, event: "WhitelistedAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAdded is a free log subscription operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Whitelist *WhitelistFilterer) WatchWhitelistedAdded(opts *bind.WatchOpts, sink chan<- *WhitelistWhitelistedAdded) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "WhitelistedAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistWhitelistedAdded)
				if err := _Whitelist.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
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

// ParseWhitelistedAdded is a log parse operation binding the contract event 0xee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f.
//
// Solidity: event WhitelistedAdded(address account)
func (_Whitelist *WhitelistFilterer) ParseWhitelistedAdded(log types.Log) (*WhitelistWhitelistedAdded, error) {
	event := new(WhitelistWhitelistedAdded)
	if err := _Whitelist.contract.UnpackLog(event, "WhitelistedAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistWhitelistedRemovedIterator is returned from FilterWhitelistedRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedRemoved events raised by the Whitelist contract.
type WhitelistWhitelistedRemovedIterator struct {
	Event *WhitelistWhitelistedRemoved // Event containing the contract specifics and raw log

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
func (it *WhitelistWhitelistedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistWhitelistedRemoved)
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
		it.Event = new(WhitelistWhitelistedRemoved)
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
func (it *WhitelistWhitelistedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistWhitelistedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistWhitelistedRemoved represents a WhitelistedRemoved event raised by the Whitelist contract.
type WhitelistWhitelistedRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedRemoved is a free log retrieval operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Whitelist *WhitelistFilterer) FilterWhitelistedRemoved(opts *bind.FilterOpts) (*WhitelistWhitelistedRemovedIterator, error) {

	logs, sub, err := _Whitelist.contract.FilterLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return &WhitelistWhitelistedRemovedIterator{contract: _Whitelist.contract, event: "WhitelistedRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedRemoved is a free log subscription operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Whitelist *WhitelistFilterer) WatchWhitelistedRemoved(opts *bind.WatchOpts, sink chan<- *WhitelistWhitelistedRemoved) (event.Subscription, error) {

	logs, sub, err := _Whitelist.contract.WatchLogs(opts, "WhitelistedRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistWhitelistedRemoved)
				if err := _Whitelist.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
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

// ParseWhitelistedRemoved is a log parse operation binding the contract event 0x270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b6.
//
// Solidity: event WhitelistedRemoved(address account)
func (_Whitelist *WhitelistFilterer) ParseWhitelistedRemoved(log types.Log) (*WhitelistWhitelistedRemoved, error) {
	event := new(WhitelistWhitelistedRemoved)
	if err := _Whitelist.contract.UnpackLog(event, "WhitelistedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
