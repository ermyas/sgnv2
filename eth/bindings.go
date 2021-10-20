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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"FarmingRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"}],\"name\":\"FarmingRewardContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardsRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161195238038061195283398101604081905261002f916100a6565b61003833610056565b6000805460ff60a01b191690556001600160a01b03166080526100d6565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100b857600080fd5b81516001600160a01b03811681146100cf57600080fd5b9392505050565b60805161185a6100f8600039600081816101a601526102a8015261185a6000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063825168ff116100815780639d4323be1161005b5780639d4323be1461018e578063ccf2683b146101a1578063f2fde38b146101c857600080fd5b8063825168ff1461014e5780638456cb59146101615780638da5cb5b1461016957600080fd5b80635c975abb116100b25780635c975abb146101165780636b5d21e914610133578063715018a61461014657600080fd5b80631744092e146100ce5780633f4ba83a1461010c575b600080fd5b6100f96100dc366004611350565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6101146101db565b005b600054600160a01b900460ff166040519015158152602001610103565b6101146101413660046113cf565b610244565b610114610536565b61011461015c3660046114be565b61059a565b61011461064f565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610103565b61011461019c3660046114be565b6106b1565b6101767f000000000000000000000000000000000000000000000000000000000000000081565b6101146101d63660046114e8565b61077c565b6000546001600160a01b0316331461023a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b61024261085e565b565b600054600160a01b900460ff16156102915760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610231565b604051633416de1160e11b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063682dbc22906102eb908b908b908b908b908b908b908b908b906004016115c2565b60006040518083038186803b15801561030357600080fd5b505afa158015610317573d6000803e3d6000fd5b50505050600061035c89898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061090492505050565b9050468160200151146103b15760405162461bcd60e51b815260206004820152601160248201527f436861696e204944206d69736d617463680000000000000000000000000000006044820152606401610231565b6000805b8260400151518110156104dc576000836040015182815181106103da576103da6116a4565b602002602001015190506000846060015183815181106103fc576103fc6116a4565b60209081029190910181015186516001600160a01b03908116600090815260018452604080822092871682529190935282205490925061043c90836116d0565b905080156104c65785516001600160a01b039081166000908152600160208181526040808420948816808552949091529091208490558751909650610482919083610be6565b85516040518281526001600160a01b038086169216907f97e6c3172350795e26977663112f38653689372e771e85bad9fbadb1af0e98b29060200160405180910390a35b50505080806104d4906116e7565b9150506103b5565b508061052a5760405162461bcd60e51b815260206004820152600d60248201527f4e6f206e657720726577617264000000000000000000000000000000000000006044820152606401610231565b50505050505050505050565b6000546001600160a01b031633146105905760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610231565b6102426000610c7b565b600054600160a01b900460ff16156105e75760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610231565b336105fd6001600160a01b038416823085610ce3565b826001600160a01b0316816001600160a01b03167f40aa1b9a9157bc37a09a78d5a46e53087b82ee0034ebe896d4d1a52f31b333d48460405161064291815260200190565b60405180910390a3505050565b6000546001600160a01b031633146106a95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610231565b610242610d21565b600054600160a01b900460ff1661070a5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610231565b6000546001600160a01b031633146107645760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610231565b6107786001600160a01b0383163383610be6565b5050565b6000546001600160a01b031633146107d65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610231565b6001600160a01b0381166108525760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610231565b61085b81610c7b565b50565b600054600160a01b900460ff166108b75760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610231565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b610938604051806080016040528060006001600160a01b031681526020016000815260200160608152602001606081525090565b60408051808201909152600080825260208201849052610959826004610da9565b90508060038151811061096e5761096e6116a4565b602002602001015167ffffffffffffffff81111561098e5761098e611702565b6040519080825280602002602001820160405280156109b7578160200160208202803683370190505b5083604001819052506000816003815181106109d5576109d56116a4565b602002602001018181525050806004815181106109f4576109f46116a4565b602002602001015167ffffffffffffffff811115610a1457610a14611702565b604051908082528060200260200182016040528015610a3d578160200160208202803683370190505b508360600181905250600081600481518110610a5b57610a5b6116a4565b6020026020010181815250506000805b60208401515184511015610bdd57610a8284610e63565b90925090508160011415610ab157610aa1610a9c85610e9d565b610f5a565b6001600160a01b03168552610a6b565b8160021415610ad557610acb610ac685610e9d565b610f6b565b6020860152610a6b565b8160031415610b6a57610aea610a9c85610e9d565b856040015184600381518110610b0257610b026116a4565b602002602001015181518110610b1a57610b1a6116a4565b60200260200101906001600160a01b031690816001600160a01b03168152505082600381518110610b4d57610b4d6116a4565b602002602001018051809190610b62906116e7565b905250610a6b565b8160041415610bce57610b7f610ac685610e9d565b856060015184600481518110610b9757610b976116a4565b602002602001015181518110610baf57610baf6116a4565b60200260200101818152505082600481518110610b4d57610b4d6116a4565b610bd88482610fa2565b610a6b565b50505050919050565b6040516001600160a01b038316602482015260448101829052610c7690849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611014565b505050565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040516001600160a01b0380851660248301528316604482015260648101829052610d1b9085906323b872dd60e01b90608401610c12565b50505050565b600054600160a01b900460ff1615610d6e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610231565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586108e73390565b8151606090610db9836001611718565b67ffffffffffffffff811115610dd157610dd1611702565b604051908082528060200260200182016040528015610dfa578160200160208202803683370190505b5091506000805b60208601515186511015610e5a57610e1886610e63565b80925081935050506001848381518110610e3457610e346116a4565b60200260200101818151610e489190611718565b905250610e558682610fa2565b610e01565b50509092525090565b6000806000610e71846110f9565b9050610e7e600882611730565b9250806007166005811115610e9557610e95611752565b915050915091565b60606000610eaa836110f9565b90506000818460000151610ebe9190611718565b9050836020015151811115610ed257600080fd5b8167ffffffffffffffff811115610eeb57610eeb611702565b6040519080825280601f01601f191660200182016040528015610f15576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015610f4f578181015183820152610f48602082611718565b9050610f2d565b505050935250919050565b6000610f658261117b565b92915050565b6000602082511115610f7c57600080fd5b6020820151905081516020610f9191906116d0565b610f9c906008611768565b1c919050565b6000816005811115610fb657610fb6611752565b1415610fc557610c76826110f9565b6002816005811115610fd957610fd9611752565b14156100c9576000610fea836110f9565b90508083600001818151610ffe9190611718565b90525060208301515183511115610c7657600080fd5b6000611069826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111a39092919063ffffffff16565b805190915015610c7657808060200190518101906110879190611787565b610c765760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610231565b602080820151825181019091015160009182805b600a8110156111755783811a9150611126816007611768565b82607f16901b85179450816080166000141561116357611147816001611718565b86518790611156908390611718565b9052509395945050505050565b8061116d816116e7565b91505061110d565b50600080fd5b6000815160141461118b57600080fd5b50602001516c01000000000000000000000000900490565b60606111b284846000856111bc565b90505b9392505050565b6060824710156112345760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610231565b843b6112825760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610231565b600080866001600160a01b0316858760405161129e91906117d5565b60006040518083038185875af1925050503d80600081146112db576040519150601f19603f3d011682016040523d82523d6000602084013e6112e0565b606091505b50915091506112f08282866112fb565b979650505050505050565b6060831561130a5750816111b5565b82511561131a5782518084602001fd5b8160405162461bcd60e51b815260040161023191906117f1565b80356001600160a01b038116811461134b57600080fd5b919050565b6000806040838503121561136357600080fd5b61136c83611334565b915061137a60208401611334565b90509250929050565b60008083601f84011261139557600080fd5b50813567ffffffffffffffff8111156113ad57600080fd5b6020830191508360208260051b85010111156113c857600080fd5b9250929050565b6000806000806000806000806080898b0312156113eb57600080fd5b883567ffffffffffffffff8082111561140357600080fd5b818b0191508b601f83011261141757600080fd5b81358181111561142657600080fd5b8c602082850101111561143857600080fd5b60209283019a509850908a0135908082111561145357600080fd5b61145f8c838d01611383565b909850965060408b013591508082111561147857600080fd5b6114848c838d01611383565b909650945060608b013591508082111561149d57600080fd5b506114aa8b828c01611383565b999c989b5096995094979396929594505050565b600080604083850312156114d157600080fd5b6114da83611334565b946020939093013593505050565b6000602082840312156114fa57600080fd5b6111b582611334565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8183526000602080850194508260005b85811015611568576001600160a01b0361155583611334565b168752958201959082019060010161153c565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8311156115a557600080fd5b8260051b8083602087013760009401602001938452509192915050565b6080815260006115d6608083018a8c611503565b82810360208401528088825260208201905060208960051b8301018a60005b8b81101561166957848303601f190184528135368e9003601e1901811261161b57600080fd5b8d01803567ffffffffffffffff81111561163457600080fd5b8036038f131561164357600080fd5b611651858260208501611503565b602096870196909550939093019250506001016115f5565b5050848103604086015261167e81898b61152c565b925050508281036060840152611695818587611573565b9b9a5050505050505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000828210156116e2576116e26116ba565b500390565b60006000198214156116fb576116fb6116ba565b5060010190565b634e487b7160e01b600052604160045260246000fd5b6000821982111561172b5761172b6116ba565b500190565b60008261174d57634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b6000816000190483118215151615611782576117826116ba565b500290565b60006020828403121561179957600080fd5b815180151581146111b557600080fd5b60005b838110156117c45781810151838201526020016117ac565b83811115610d1b5750506000910152565b600082516117e78184602087016117a9565b9190910192915050565b60208152600082518060208401526118108160408501602087016117a9565b601f01601f1916919091016040019291505056fea2646970667358221220d70cb74fa731f64e7491ab1bf6e0f9a6499c5672c10c6957fef6126eac966b3164736f6c63430008090033",
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
	Bin: "0x60e06040523480156200001157600080fd5b50604051620015953803806200159583398101604081905262000034916200006b565b6001600160a01b0392831660805290821660a0521660c052620000bf565b6001600160a01b03811681146200006857600080fd5b50565b6000806000606084860312156200008157600080fd5b83516200008e8162000052565b6020850151909350620000a18162000052565b6040850151909250620000b48162000052565b809150509250925092565b60805160a05160c0516114676200012e600039600081816101da01526105be0152600081816102140152818161059b015281816108fc0152610b7f0152600081816101040152818161026b015281816105f10152818161087f015281816109d60152610aaa01526114676000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806382d7b4b811610081578063934a18ec1161005b578063934a18ec146101fc578063c6c21e9d1461020f578063e478ed9d1461023657600080fd5b806382d7b4b8146101c45780638338f0e5146101cc578063913e77ad146101d557600080fd5b80634cf088d9116100b25780634cf088d9146100ff578063581c53c51461013e5780637e5fb8f31461015e57600080fd5b806322da7927146100ce57806325ed6b35146100ea575b600080fd5b6100d760015481565b6040519081526020015b60405180910390f35b6100fd6100f8366004610f63565b610249565b005b6101267f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100e1565b61015161014c366004610fa8565b61050a565b6040516100e19190610ff7565b6101b261016c366004611005565b6000602081905290815260409020805460018201546002830154600384015460048501546005909501546001600160a01b03909416949293919260ff9182169290911686565b6040516100e19695949392919061102e565b6100fd610538565b6100d760025481565b6101267f000000000000000000000000000000000000000000000000000000000000000081565b6100fd61020a366004611005565b6105ea565b6101267f000000000000000000000000000000000000000000000000000000000000000081565b6100fd610244366004611080565b61099a565b33600360405163a310624f60e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a310624f9060240160206040518083038186803b1580156102ad57600080fd5b505afa1580156102c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e591906110b0565b60038111156102f6576102f6610fcd565b146103485760405162461bcd60e51b815260206004820152601f60248201527f566f746572206973206e6f74206120626f6e6465642076616c696461746f720060448201526064015b60405180910390fd5b60008381526020819052604090206001600582015460ff16600281111561037157610371610fcd565b146103be5760405162461bcd60e51b815260206004820152601760248201527f496e76616c69642070726f706f73616c20737461747573000000000000000000604482015260640161033f565b806002015443106104115760405162461bcd60e51b815260206004820152601460248201527f566f746520646561646c696e6520706173736564000000000000000000000000604482015260640161033f565b6001600160a01b038216600090815260068201602052604081205460ff16600381111561044057610440610fcd565b1461048d5760405162461bcd60e51b815260206004820152600f60248201527f566f7465722068617320766f7465640000000000000000000000000000000000604482015260640161033f565b6001600160a01b03821660009081526006820160205260409020805484919060ff191660018360038111156104c4576104c4610fcd565b02179055507f06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d658483856040516104fc939291906110cd565b60405180910390a150505050565b6000828152602081815260408083206001600160a01b038516845260060190915290205460ff165b92915050565b60006002541161058a5760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20636f6c6c6563740000000000000000000000000000604482015260640161033f565b6002546105e3906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016907f000000000000000000000000000000000000000000000000000000000000000090610c08565b6000600255565b60008060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634021d4d56040518163ffffffff1660e01b815260040160006040518083038186803b15801561064857600080fd5b505afa15801561065c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106849190810190611160565b905060005b81518163ffffffff1610156107535760016106c786848463ffffffff16815181106106b6576106b6611236565b60200260200101516000015161050a565b60038111156106d8576106d8610fcd565b141561071057818163ffffffff16815181106106f6576106f6611236565b6020026020010151602001518461070d9190611262565b93505b818163ffffffff168151811061072857610728611236565b6020026020010151602001518361073f9190611262565b92508061074b8161127a565b915050610689565b506000600361076384600261129e565b61076d91906112bd565b610778906001611262565b60008681526020819052604090209085101591506001600582015460ff1660028111156107a7576107a7610fcd565b146107f45760405162461bcd60e51b815260206004820152601760248201527f496e76616c69642070726f706f73616c20737461747573000000000000000000604482015260640161033f565b80600201544310156108485760405162461bcd60e51b815260206004820152601960248201527f566f746520646561646c696e65206e6f74207265616368656400000000000000604482015260640161033f565b60058101805460ff19166002179055811561092d57600381015460048083015460405163e909156d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169363e909156d936108b69360ff90921692016112df565b600060405180830381600087803b1580156108d057600080fd5b505af11580156108e4573d6000803e3d6000fd5b50508254600184015461092893506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116935090911690610c08565b610949565b8060010154600260008282546109439190611262565b90915550505b600381015460048201546040517fd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a99261098a928a92879260ff1691906112fa565b60405180910390a1505050505050565b600180546000818152602081905260409020916109b79190611262565b600155604051631042b80b60e21b815233906000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063410ae02c90610a0b908490600401611325565b60206040518083038186803b158015610a2357600080fd5b505afa158015610a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5b9190611333565b83547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038481169190911785556001808601839055604051631042b80b60e21b81529293507f00000000000000000000000000000000000000000000000000000000000000009091169163410ae02c91610adf91600401611325565b60206040518083038186803b158015610af757600080fd5b505afa158015610b0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2f9190611333565b610b399043611262565b600284015560038301805486919060ff19166001836008811115610b5f57610b5f610fcd565b02179055506004830184905560058301805460ff19166001179055610baf7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316833084610c9d565b7f4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f16060018054610bde919061134c565b838386600201548989604051610bf996959493929190611363565b60405180910390a15050505050565b6040516001600160a01b038316602482015260448101829052610c9890849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610cdb565b505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610cd59085906323b872dd60e01b90608401610c34565b50505050565b6000610d30826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610dc09092919063ffffffff16565b805190915015610c985780806020019051810190610d4e9190611394565b610c985760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161033f565b6060610dcf8484600085610dd9565b90505b9392505050565b606082471015610e515760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161033f565b843b610e9f5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161033f565b600080866001600160a01b03168587604051610ebb91906113e2565b60006040518083038185875af1925050503d8060008114610ef8576040519150601f19603f3d011682016040523d82523d6000602084013e610efd565b606091505b5091509150610f0d828286610f1a565b925050505b949350505050565b60608315610f29575081610dd2565b825115610f395782518084602001fd5b8160405162461bcd60e51b815260040161033f91906113fe565b60048110610f6057600080fd5b50565b60008060408385031215610f7657600080fd5b823591506020830135610f8881610f53565b809150509250929050565b6001600160a01b0381168114610f6057600080fd5b60008060408385031215610fbb57600080fd5b823591506020830135610f8881610f93565b634e487b7160e01b600052602160045260246000fd5b60048110610ff357610ff3610fcd565b9052565b602081016105328284610fe3565b60006020828403121561101757600080fd5b5035919050565b60098110610ff357610ff3610fcd565b6001600160a01b0387168152602081018690526040810185905260c08101611059606083018661101e565b8360808301526003831061106f5761106f610fcd565b8260a0830152979650505050505050565b6000806040838503121561109357600080fd5b8235600981106110a257600080fd5b946020939093013593505050565b6000602082840312156110c257600080fd5b8151610dd281610f53565b8381526001600160a01b038316602082015260608101610f126040830184610fe3565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715611129576111296110f0565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611158576111586110f0565b604052919050565b6000602080838503121561117357600080fd5b825167ffffffffffffffff8082111561118b57600080fd5b818501915085601f83011261119f57600080fd5b8151818111156111b1576111b16110f0565b6111bf848260051b0161112f565b818152848101925060069190911b8301840190878211156111df57600080fd5b928401925b8184101561122b57604084890312156111fd5760008081fd5b611205611106565b845161121081610f93565b815284860151868201528352604090930192918401916111e4565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156112755761127561124c565b500190565b600063ffffffff808316818114156112945761129461124c565b6001019392505050565b60008160001904831182151516156112b8576112b861124c565b500290565b6000826112da57634e487b7160e01b600052601260045260246000fd5b500490565b604081016112ed828561101e565b8260208301529392505050565b848152831515602082015260808101611316604083018561101e565b82606083015295945050505050565b60208101610532828461101e565b60006020828403121561134557600080fd5b5051919050565b60008282101561135e5761135e61124c565b500390565b8681526001600160a01b0386166020820152604081018590526060810184905260c0810161106f608083018561101e565b6000602082840312156113a657600080fd5b81518015158114610dd257600080fd5b60005b838110156113d15781810151838201526020016113b9565b83811115610cd55750506000910152565b600082516113f48184602087016113b6565b9190910192915050565b602081526000825180602084015261141d8160408501602087016113b6565b601f01601f1916919091016040019291505056fea264697066735822122080a34fe6ac80ae11d09e3a9073531d50f6a4169f34eac14792f820deda42ab7a64736f6c63430008090033",
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

// PbMetaData contains all meta data concerning the Pb contract.
var PbMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208c86508021ecf303aa39a3a2e9c5e099853f51f65d46fd898a4ab3951fe5175d64736f6c63430008090033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f8198a61ca72d4ddaf10ce7e4140662ac5148289fb3bc4b8ad7408cf2666b0de64736f6c63430008090033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209f74846f2188d1ca01ee12d307953d2ee6133c819f842c679dee7638f80f4c2164736f6c63430008090033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205295a6dd7508517f6a2d4c32ce66ff99ba87c14b1cbce02412a164c3052fcb3a64736f6c63430008090033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"oldAddr\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newAddr\",\"type\":\"bytes\"}],\"name\":\"SgnAddrUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sgnAddrs\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sgnAddr\",\"type\":\"bytes\"}],\"name\":\"updateSgnAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_withdrawalRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnAmts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405162001c8138038062001c81833981016040819052610031916100a8565b61003a33610058565b6000805460ff60a01b191690556001600160a01b03166080526100d8565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ba57600080fd5b81516001600160a01b03811681146100d157600080fd5b9392505050565b608051611b6a620001176000396000818161011101528181610656015281816106f30152818161079a015281816109590152610a740152611b6a6000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063c429fe1f11610066578063c429fe1f146101ed578063d0bb93511461020d578063d88ef27114610220578063f2fde38b1461023357600080fd5b80638da5cb5b146101b65780639d4323be146101c7578063b02c43d0146101da57600080fd5b80635c975abb116100c85780635c975abb14610150578063715018a61461016d578063795c2c14146101755780638456cb59146101ae57600080fd5b80633f4ba83a146100ef57806347e7ef24146100f95780634cf088d91461010c575b600080fd5b6100f7610246565b005b6100f761010736600461161b565b6102af565b6101337f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b600054600160a01b900460ff166040519015158152602001610147565b6100f76103e7565b6101a0610183366004611647565b600260209081526000928352604080842090915290825290205481565b604051908152602001610147565b6100f761044b565b6000546001600160a01b0316610133565b6100f76101d536600461161b565b6104ad565b6101a06101e8366004611680565b610578565b6102006101fb366004611699565b610599565b604051610147919061170e565b6100f761021b36600461176a565b610633565b6100f761022e3660046117ac565b610a10565b6100f7610241366004611699565b610c76565b6000546001600160a01b031633146102a55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6102ad610d58565b565b600054600160a01b900460ff16156102fc5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161029c565b6040516bffffffffffffffffffffffff1933606081811b8316602085015285901b9091166034830152604882018390529060019060680160408051601f198184030181529190528051602091820120825460018101845560009384529190922001556103736001600160a01b038416823085610dfe565b600180546000916103839161185c565b6040805167ffffffffffffffff831681526001600160a01b0385811660208301528716818301526060810186905290519192507f2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8919081900360800190a150505050565b6000546001600160a01b031633146104415760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161029c565b6102ad6000610e9c565b6000546001600160a01b031633146104a55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161029c565b6102ad610f04565b600054600160a01b900460ff166105065760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161029c565b6000546001600160a01b031633146105605760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161029c565b6105746001600160a01b0383163383610f8c565b5050565b6001818154811061058857600080fd5b600091825260209091200154905081565b600360205260009081526040902080546105b290611873565b80601f01602080910402602001604051908101604052809291908181526020018280546105de90611873565b801561062b5780601f106106005761010080835404028352916020019161062b565b820191906000526020600020905b81548152906001019060200180831161060e57829003601f168201915b505050505081565b604051636d30878360e01b81523360048201819052906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636d3087839060240160206040518083038186803b15801561069857600080fd5b505afa1580156106ac573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106d091906118ae565b6001600160a01b03161461077857604051636d30878360e01b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636d3087839060240160206040518083038186803b15801561073d57600080fd5b505afa158015610751573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077591906118ae565b90505b60405163a310624f60e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a310624f9060240160206040518083038186803b1580156107de57600080fd5b505afa1580156107f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061081691906118cb565b9050600181600381111561082c5761082c6118ec565b146108795760405162461bcd60e51b815260206004820152601660248201527f4e6f7420756e626f6e6465642076616c696461746f7200000000000000000000604482015260640161029c565b6001600160a01b0382166000908152600360205260408120805461089c90611873565b80601f01602080910402602001604051908101604052809291908181526020018280546108c890611873565b80156109155780601f106108ea57610100808354040283529160200191610915565b820191906000526020600020905b8154815290600101906020018083116108f857829003601f168201915b505050506001600160a01b0385166000908152600360205260409020919250610941919050868661156d565b506040516309146f1160e41b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639146f110906109929086908990899060040161192b565b600060405180830381600087803b1580156109ac57600080fd5b505af11580156109c0573d6000803e3d6000fd5b50505050826001600160a01b03167f8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4828787604051610a019392919061198b565b60405180910390a25050505050565b600054600160a01b900460ff1615610a5d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161029c565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe90610aaf9087908790879087906004016119bb565b60206040518083038186803b158015610ac757600080fd5b505afa158015610adb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aff9190611a6c565b506000610b4185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610fc192505050565b80516001600160a01b039081166000908152600260209081526040808320828601519094168352929052818120549183015192935091610b81919061185c565b905060008111610bd35760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f20776974686472617700000000000000604482015260640161029c565b60408083015183516001600160a01b0390811660009081526002602090815284822081880180518516845291529390209190915583519151610c189291169083610f8c565b8151602080840151604080516001600160a01b039485168152939091169183019190915281018290527f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9060600160405180910390a1505050505050565b6000546001600160a01b03163314610cd05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161029c565b6001600160a01b038116610d4c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161029c565b610d5581610e9c565b50565b600054600160a01b900460ff16610db15760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161029c565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040516001600160a01b0380851660248301528316604482015260648101829052610e969085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261109c565b50505050565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff1615610f515760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161029c565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610de13390565b6040516001600160a01b038316602482015260448101829052610fbc90849063a9059cbb60e01b90606401610e32565b505050565b604080516060810182526000808252602080830182905282840182905283518085019094528184528301849052909190805b602083015151835110156110945761100a83611181565b9092509050816001141561103957611029611024846111bb565b611278565b6001600160a01b03168452610ff3565b81600214156110615761104e611024846111bb565b6001600160a01b03166020850152610ff3565b81600314156110855761107b611076846111bb565b611289565b6040850152610ff3565b61108f83826112c0565b610ff3565b505050919050565b60006110f1826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166113329092919063ffffffff16565b805190915015610fbc578080602001905181019061110f9190611a6c565b610fbc5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161029c565b600080600061118f8461134b565b905061119c600882611a8e565b92508060071660058111156111b3576111b36118ec565b915050915091565b606060006111c88361134b565b905060008184600001516111dc9190611ab0565b90508360200151518111156111f057600080fd5b8167ffffffffffffffff81111561120957611209611ac8565b6040519080825280601f01601f191660200182016040528015611233576020820181803683370190505b50602080860151865192955091818601919083010160005b8581101561126d578181015183820152611266602082611ab0565b905061124b565b505050935250919050565b6000611283826113cd565b92915050565b600060208251111561129a57600080fd5b60208201519050815160206112af919061185c565b6112ba906008611ade565b1c919050565b60008160058111156112d4576112d46118ec565b14156112e357610fbc8261134b565b60028160058111156112f7576112f76118ec565b14156100ea5760006113088361134b565b9050808360000181815161131c9190611ab0565b90525060208301515183511115610fbc57600080fd5b606061134184846000856113f5565b90505b9392505050565b602080820151825181019091015160009182805b600a8110156113c75783811a9150611378816007611ade565b82607f16901b8517945081608016600014156113b557611399816001611ab0565b865187906113a8908390611ab0565b9052509395945050505050565b806113bf81611afd565b91505061135f565b50600080fd5b600081516014146113dd57600080fd5b50602001516c01000000000000000000000000900490565b60608247101561146d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161029c565b843b6114bb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161029c565b600080866001600160a01b031685876040516114d79190611b18565b60006040518083038185875af1925050503d8060008114611514576040519150601f19603f3d011682016040523d82523d6000602084013e611519565b606091505b5091509150611529828286611534565b979650505050505050565b60608315611543575081611344565b8251156115535782518084602001fd5b8160405162461bcd60e51b815260040161029c919061170e565b82805461157990611873565b90600052602060002090601f01602090048101928261159b57600085556115e1565b82601f106115b45782800160ff198235161785556115e1565b828001600101855582156115e1579182015b828111156115e15782358255916020019190600101906115c6565b506115ed9291506115f1565b5090565b5b808211156115ed57600081556001016115f2565b6001600160a01b0381168114610d5557600080fd5b6000806040838503121561162e57600080fd5b823561163981611606565b946020939093013593505050565b6000806040838503121561165a57600080fd5b823561166581611606565b9150602083013561167581611606565b809150509250929050565b60006020828403121561169257600080fd5b5035919050565b6000602082840312156116ab57600080fd5b813561134481611606565b60005b838110156116d15781810151838201526020016116b9565b83811115610e965750506000910152565b600081518084526116fa8160208601602086016116b6565b601f01601f19169290920160200192915050565b60208152600061134460208301846116e2565b60008083601f84011261173357600080fd5b50813567ffffffffffffffff81111561174b57600080fd5b60208301915083602082850101111561176357600080fd5b9250929050565b6000806020838503121561177d57600080fd5b823567ffffffffffffffff81111561179457600080fd5b6117a085828601611721565b90969095509350505050565b600080600080604085870312156117c257600080fd5b843567ffffffffffffffff808211156117da57600080fd5b6117e688838901611721565b909650945060208701359150808211156117ff57600080fd5b818701915087601f83011261181357600080fd5b81358181111561182257600080fd5b8860208260051b850101111561183757600080fd5b95989497505060200194505050565b634e487b7160e01b600052601160045260246000fd5b60008282101561186e5761186e611846565b500390565b600181811c9082168061188757607f821691505b602082108114156118a857634e487b7160e01b600052602260045260246000fd5b50919050565b6000602082840312156118c057600080fd5b815161134481611606565b6000602082840312156118dd57600080fd5b81516004811061134457600080fd5b634e487b7160e01b600052602160045260246000fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b038416815260606020820152600860608201527f73676e2d61646472000000000000000000000000000000000000000000000000608082015260a06040820152600061198260a083018486611902565b95945050505050565b60408152600061199e60408301866116e2565b82810360208401526119b1818587611902565b9695505050505050565b6040815260006119cf604083018688611902565b602083820381850152818583528183019050818660051b8401018760005b88811015611a5c57858303601f190184528135368b9003601e19018112611a1357600080fd5b8a01803567ffffffffffffffff811115611a2c57600080fd5b8036038c1315611a3b57600080fd5b611a488582898501611902565b9587019594505050908401906001016119ed565b50909a9950505050505050505050565b600060208284031215611a7e57600080fd5b8151801515811461134457600080fd5b600082611aab57634e487b7160e01b600052601260045260246000fd5b500490565b60008219821115611ac357611ac3611846565b500190565b634e487b7160e01b600052604160045260246000fd5b6000816000190483118215151615611af857611af8611846565b500290565b6000600019821415611b1157611b11611846565b5060010190565b60008251611b2a8184602087016116b6565b919091019291505056fea26469706673582212208424e101f39f349a1385216aba0d89a89f0a8affc310806026e99fcedd0b800164736f6c63430008090033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorBondInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"tokenDiff\",\"type\":\"int256\"}],\"name\":\"DelegationUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmt\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashAmtCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ValidatorNotice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CELER_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bondedValAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectForfeiture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"completeUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"confirmUnbondedValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forfeiture\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorsTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ValidatorTokens[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawableUndelegationTokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.DelegatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"}],\"name\":\"getParamValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_checkSelfDelegation\",\"type\":\"bool\"}],\"name\":\"hasMinRequiredTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_commissionRate\",\"type\":\"uint64\"}],\"name\":\"initializeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBondedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBondBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setGovContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"name\":\"setMaxSlashFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setParamValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRewardContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerVals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newRate\",\"type\":\"uint64\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"}],\"name\":\"updateMinSelfDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"updateValidatorSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"validatorNotice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"bondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"unbondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162005b9438038062005b948339810160408190526200003491620001f9565b6200003f33620001a9565b6000805460ff60a01b191681556001600160a01b039a909a16608052600a6020527f13da86008ba1c6922daee3e07db95305ef49ebced9f5467a0b8613fcc6b343e3989098557fbbc70db1b6c7afd11e79c0fb0051300458f1a3acb8ee9789d9b6b26c61ad9bc7969096557fbff4442b8ed600beeb8e26b1279a0f0d14c6edfaec26d968ee13c86f7d4c2ba8949094557fa856840544dc26124927add067d799967eac11be13e14d82cc281ea46fa39759929092557fe1eb2b2161a492c07c5a334e48012567cba93ec021043f53c1955516a3c5a841557ff35035bc2b01d44bd35a1dcdc552315cffb73da35cfd60570b7b777f98036f9f557f10d9dd018e4cae503383c9f804c1c1603ada5856ee7894375d9b97cd8c8b27db557f22e39f61d1e4986b4f116cea9067f62cc77d74dff1780ae9c8b5166d1dd288295560089091527f2c1fd36ba11b13b555f58753742999069764391f450ca8727fe8a3eeffe677755562000286565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000806000806000806000806101408b8d0312156200021a57600080fd5b8a516001600160a01b03811681146200023257600080fd5b809a505060208b0151985060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935060e08b015192506101008b015191506101208b015190509295989b9194979a5092959850565b6080516158c1620002d3600039600081816108a701528181610cf101528181610ed5015281816119160152818161199101528181611eb2015281816122ce015261320c01526158c16000f3fe60806040526004361061034e5760003560e01c806368706e54116101bb57806390e360f8116100f7578063b4f7fa3411610095578063eb505dd51161006f578063eb505dd5146109d8578063eecefef814610a05578063f2fde38b14610a32578063fa52c7d814610a5257600080fd5b8063b4f7fa341461095f578063c8f9f9841461097f578063e909156d146109b857600080fd5b8063960dc08a116100d1578063960dc08a146108955780639b19251a146108c9578063a310624f146108f9578063acc62ccf1461093f57600080fd5b806390e360f8146108255780639146f1101461085557806392bb243c1461087557600080fd5b806382d7b4b8116101645780638456cb591161013e5780638456cb59146107bd57806389f9aab5146107d25780638a74d5fe146107e75780638da5cb5b1461080757600080fd5b806382d7b4b81461077c5780638338f0e51461079157806383cfb318146107a757600080fd5b8063715018a611610195578063715018a61461072757806371bc02161461073c5780637a50dbd21461075c57600080fd5b806368706e54146106b15780636d308783146106d15780636ea69d621461070757600080fd5b80634021d4d51161028a57806351508f0a116102335780635c975abb1161020d5780635c975abb1461063c5780635e593eff1461065b57806365d5d4201461067b578063682dbc221461069157600080fd5b806351508f0a146105e257806351fb012d14610602578063525eba211461061c57600080fd5b806347abfdbf1161026457806347abfdbf1461058257806349955e39146105a25780634d99dd16146105c257600080fd5b80634021d4d514610520578063410ae02c14610542578063473849bd1461056257600080fd5b8063291d9549116102f7578063386c024a116102d1578063386c024a1461048d5780633985c4e6146104a25780633af32abf146104c25780633f4ba83a1461050b57600080fd5b8063291d9549146104205780632fa4d12b1461044057806336f1635f1461047857600080fd5b8063145aa11611610328578063145aa116146103bc5780631a203257146103dc5780631cfe4f0b146103fc57600080fd5b8063026e402b1461035a578063052d9e7e1461037c57806310154bad1461039c57600080fd5b3661035557005b600080fd5b34801561036657600080fd5b5061037a610375366004614dc1565b610af9565b005b34801561038857600080fd5b5061037a610397366004614df9565b610d7a565b3480156103a857600080fd5b5061037a6103b7366004614e16565b610dd6565b3480156103c857600080fd5b5061037a6103d7366004614e31565b610e27565b3480156103e857600080fd5b5061037a6103f7366004614e31565b610efc565b34801561040857600080fd5b506005545b6040519081526020015b60405180910390f35b34801561042c57600080fd5b5061037a61043b366004614e16565b610f72565b34801561044c57600080fd5b50600b54610460906001600160a01b031681565b6040516001600160a01b039091168152602001610417565b34801561048457600080fd5b5061037a610fc3565b34801561049957600080fd5b5061040d611344565b3480156104ae57600080fd5b5061037a6104bd366004614ed8565b611370565b3480156104ce57600080fd5b506104fb6104dd366004614e16565b6001600160a01b031660009081526001602052604090205460ff1690565b6040519015158152602001610417565b34801561051757600080fd5b5061037a611a97565b34801561052c57600080fd5b50610535611ae9565b6040516104179190614f44565b34801561054e57600080fd5b5061040d61055d366004614fab565b611bea565b34801561056e57600080fd5b5061037a61057d366004614e16565b611c29565b34801561058e57600080fd5b506104fb61059d366004614fc6565b611f31565b3480156105ae57600080fd5b5061037a6105bd366004615015565b611feb565b3480156105ce57600080fd5b5061037a6105dd366004614dc1565b61214d565b3480156105ee57600080fd5b5061037a6105fd366004614e16565b612535565b34801561060e57600080fd5b506002546104fb9060ff1681565b34801561062857600080fd5b5061037a610637366004615030565b6125f8565b34801561064857600080fd5b50600054600160a01b900460ff166104fb565b34801561066757600080fd5b5061037a610676366004614e31565b612a9b565b34801561068757600080fd5b5061040d60035481565b34801561069d57600080fd5b5061037a6106ac366004615123565b612cb4565b3480156106bd57600080fd5b5061037a6106cc366004614e16565b612d17565b3480156106dd57600080fd5b506104606106ec366004614e16565b6008602052600090815260409020546001600160a01b031681565b34801561071357600080fd5b50600c54610460906001600160a01b031681565b34801561073357600080fd5b5061037a612dda565b34801561074857600080fd5b5061037a610757366004614e16565b612e2c565b34801561076857600080fd5b5061037a610777366004614e16565b612f74565b34801561078857600080fd5b5061037a6131a6565b34801561079d57600080fd5b5061040d600d5481565b3480156107b357600080fd5b5061040d60045481565b3480156107c957600080fd5b5061037a61323d565b3480156107de57600080fd5b5060065461040d565b3480156107f357600080fd5b506104fb610802366004615267565b61328d565b34801561081357600080fd5b506000546001600160a01b0316610460565b34801561083157600080fd5b506104fb610840366004614e31565b60096020526000908152604090205460ff1681565b34801561086157600080fd5b5061037a6108703660046152df565b613471565b34801561088157600080fd5b50610460610890366004614e31565b61353f565b3480156108a157600080fd5b506104607f000000000000000000000000000000000000000000000000000000000000000081565b3480156108d557600080fd5b506104fb6108e4366004614e16565b60016020526000908152604090205460ff1681565b34801561090557600080fd5b50610932610914366004614e16565b6001600160a01b031660009081526007602052604090205460ff1690565b6040516104179190615398565b34801561094b57600080fd5b5061046061095a366004614e31565b613569565b34801561096b57600080fd5b506104fb61097a366004614e16565b613579565b34801561098b57600080fd5b5061040d61099a366004614e16565b6001600160a01b031660009081526007602052604090206001015490565b3480156109c457600080fd5b5061037a6109d33660046153a6565b6135b1565b3480156109e457600080fd5b5061040d6109f3366004614fab565b600a6020526000908152604090205481565b348015610a1157600080fd5b50610a25610a203660046153c2565b613648565b60405161041791906153f5565b348015610a3e57600080fd5b5061037a610a4d366004614e16565b613941565b348015610a5e57600080fd5b50610ae3610a6d366004614e16565b6007602081905260009182526040909120805460018201546002830154600384015460048501546006860154959096015460ff8516966101009095046001600160a01b031695939492939192919067ffffffffffffffff80821691680100000000000000008104821691600160801b909104168a565b6040516104179a99989796959493929190615494565b600054600160a01b900460ff1615610b4b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064015b60405180910390fd5b33670de0b6b3a7640000821015610ba45760405162461bcd60e51b815260206004820152601860248201527f4d696e696d616c20616d6f756e7420697320312043454c5200000000000000006044820152606401610b42565b6001600160a01b038316600090815260076020526040812090815460ff166003811115610bd357610bd3615360565b1415610c215760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610b42565b6000610c368483600101548460020154613a0e565b6001600160a01b0384166000908152600584016020526040812080549293509183918391610c65908490615516565b9250508190555081836002016000828254610c809190615516565b9250508190555084836001016000828254610c9b9190615516565b9091555060039050835460ff166003811115610cb957610cb9615360565b1415610ce4578460036000828254610cd19190615516565b90915550506001830154610ce490613a3b565b610d196001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016853088613b5d565b6001830154815460408051928352602083019190915281018690526001600160a01b0380861691908816907f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea906060015b60405180910390a3505050505050565b6000546001600160a01b03163314610dc25760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b6002805460ff191682151517905550565b50565b6000546001600160a01b03163314610e1e5760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b610dd381613bfb565b600054600160a01b900460ff16610e805760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610b42565b6000546001600160a01b03163314610ec85760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b610dd36001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163383613cc0565b6000546001600160a01b03163314610f445760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b6008600052600a6020527f2c1fd36ba11b13b555f58753742999069764391f450ca8727fe8a3eeffe6777555565b6000546001600160a01b03163314610fba5760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b610dd381613cf0565b336000818152600860205260409020546001600160a01b031615610ffc5750336000908152600860205260409020546001600160a01b03165b6001600160a01b03811660009081526007602052604090206001815460ff16600381111561102c5761102c615360565b148061104d57506002815460ff16600381111561104b5761104b615360565b145b6110995760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610b42565b600781015467ffffffffffffffff164310156110f75760405162461bcd60e51b815260206004820152601660248201527f426f6e6420626c6f636b206e6f742072656163686564000000000000000000006044820152606401610b42565b6004544310156111495760405162461bcd60e51b815260206004820152601b60248201527f546f6f206672657175656e742076616c696461746f7220626f6e6400000000006044820152606401610b42565b6007600052600a6020527f22e39f61d1e4986b4f116cea9067f62cc77d74dff1780ae9c8b5166d1dd288295461117f9043615516565b60045561118d826001611f31565b6111d95760405162461bcd60e51b815260206004820152601360248201527f4e6f742068617665206d696e20746f6b656e73000000000000000000000000006044820152606401610b42565b6003600052600a6020527fa856840544dc26124927add067d799967eac11be13e14d82cc281ea46fa397595460065481111561121d5761121883613da9565b505050565b6000196000805b838110156112d2578260076000600684815481106112445761124461552e565b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015410156112c057809150600760006006838154811061128c5761128c61552e565b60009182526020808320909101546001600160a01b031683528201929092526040019020600101549250826112c0576112d2565b806112ca81615544565b915050611224565b50818460010154116113265760405162461bcd60e51b815260206004820152601360248201527f496e73756666696369656e7420746f6b656e73000000000000000000000000006044820152606401610b42565b6113308582613dfd565b61133d8460010154613a3b565b5050505050565b6000600380546002611356919061555f565b611360919061557e565b61136b906001615516565b905090565b600054600160a01b900460ff16156113bd5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b42565b61140384848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061080292508591508690506155a0565b50600061144585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250613e7892505050565b9050806060015167ffffffffffffffff1642106114a45760405162461bcd60e51b815260206004820152600d60248201527f536c6173682065787069726564000000000000000000000000000000000000006044820152606401610b42565b620f4240816040015167ffffffffffffffff1611156115055760405162461bcd60e51b815260206004820152601460248201527f496e76616c696420736c61736820666163746f720000000000000000000000006044820152606401610b42565b6008600052600a6020527f2c1fd36ba11b13b555f58753742999069764391f450ca8727fe8a3eeffe6777554604082015167ffffffffffffffff16111561158e5760405162461bcd60e51b815260206004820152601760248201527f457863656564206d617820736c61736820666163746f720000000000000000006044820152606401610b42565b60208082015167ffffffffffffffff1660009081526009909152604090205460ff16156115fd5760405162461bcd60e51b815260206004820152601060248201527f5573656420736c617368206e6f6e6365000000000000000000000000000000006044820152606401610b42565b60208082015167ffffffffffffffff166000908152600982526040808220805460ff1916600117905583516001600160a01b0381168352600790935290206003815460ff16600381111561165357611653615360565b148061167457506002815460ff16600381111561167257611672615360565b145b6116c05760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610b42565b6000620f4240846040015167ffffffffffffffff1683600101546116e4919061555f565b6116ee919061557e565b90508082600101600082825461170491906155ad565b9091555060039050825460ff16600381111561172257611722615360565b14156117c357806003600082825461173a91906155ad565b9091555050608084015167ffffffffffffffff161515806117635750611761836001611f31565b155b156117c357611771836140ea565b608084015167ffffffffffffffff16156117c357608084015161179e9067ffffffffffffffff1643615516565b60078301805467ffffffffffffffff191667ffffffffffffffff929092169190911790555b60006001600160a01b0316836001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea846001015460008561180a906155c4565b6040805193845260208401929092529082015260600160405180910390a36000620f4240856040015167ffffffffffffffff16846003015461184c919061555f565b611856919061557e565b90508083600301600082825461186c91906155ad565b9091555061187c90508183615516565b91506000805b8660a0015151811015611a165760008760a0015182815181106118a7576118a761552e565b60200260200101519050848160200151846118c29190615516565b11156118d8576118d283866155ad565b60208201525b602081015115611a035760208101516118f19084615516565b81519093506001600160a01b031661197c57602081015161193e906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016903390613cc0565b60208082015160405190815233917fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3910160405180910390a2611a03565b805160208201516119b7916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691613cc0565b80600001516001600160a01b03167fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a382602001516040516119fa91815260200190565b60405180910390a25b5080611a0e81615544565b915050611882565b50611a2181846155ad565b600d6000828254611a329190615516565b90915550506020808701516040805167ffffffffffffffff90921682529181018590526001600160a01b038716917f10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008910160405180910390a250505050505050505050565b6000546001600160a01b03163314611adf5760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b611ae7614250565b565b60055460609060009067ffffffffffffffff811115611b0a57611b0a61506c565b604051908082528060200260200182016040528015611b4f57816020015b6040805180820190915260008082526020820152815260200190600190039081611b285790505b50905060005b600654811015611be457600060068281548110611b7457611b7461552e565b60009182526020808320909101546040805180820182526001600160a01b039092168083528085526007845293206001015491810191909152845191925090849084908110611bc557611bc561552e565b6020026020010181905250508080611bdc90615544565b915050611b55565b50919050565b6000600a6000836008811115611c0257611c02615360565b6008811115611c1357611c13615360565b8152602001908152602001600020549050919050565b6001600160a01b03811660009081526007602052604081203391815460ff166003811115611c5957611c59615360565b1415611ca75760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610b42565b6001600160a01b03821660009081526005820160209081526040822060028352600a9091527fbff4442b8ed600beeb8e26b1279a0f0d14c6edfaec26d968ee13c86f7d4c2ba85483549192909160019060ff166003811115611d0b57611d0b615360565b60028501549114915063ffffffff1660005b600285015463ffffffff64010000000090910481169083161015611dcf578280611d6d575063ffffffff821660009081526001808701602052604090912001544390611d6a908690615516565b11155b15611db85763ffffffff82166000908152600186016020526040902054611d949082615516565b63ffffffff8316600090815260018088016020526040822082815501559050611dbd565b611dcf565b81611dc7816155e1565b925050611d1d565b60028501805463ffffffff191663ffffffff841617905580611e595760405162461bcd60e51b815260206004820152602560248201527f4e6f20756e64656c65676174696f6e20726561647920746f20626520636f6d7060448201527f6c657465640000000000000000000000000000000000000000000000000000006064820152608401610b42565b6000611e6e82886003015489600401546142f6565b905081876004016000828254611e8491906155ad565b9250508190555080876003016000828254611e9f91906155ad565b90915550611ed990506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168983613cc0565b876001600160a01b0316896001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c83604051611f1e91815260200190565b60405180910390a3505050505050505050565b6001600160a01b03821660009081526007602090815260408220600181015460048452600a9092527fe1eb2b2161a492c07c5a334e48012567cba93ec021043f53c1955516a3c5a84154909190811015611f9057600092505050611fe5565b8315611fde576001600160a01b03851660009081526005830160205260408120546002840154611fc2919084906142f6565b90508260060154811015611fdc5760009350505050611fe5565b505b6001925050505b92915050565b33600081815260076020526040812090815460ff16600381111561201157612011615360565b141561205f5760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610b42565b6127108367ffffffffffffffff1611156120bb5760405162461bcd60e51b815260206004820152601060248201527f496e76616c6964206e65772072617465000000000000000000000000000000006044820152606401610b42565b60078101805467ffffffffffffffff60801b1916600160801b67ffffffffffffffff8616908102919091179091556040805160208101929092526001600160a01b038416917f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c910160408051601f19818403018152908290526121409160009061565d565b60405180910390a2505050565b33670de0b6b3a76400008210156121a65760405162461bcd60e51b815260206004820152601960248201527f4d696e696d616c20616d6f756e742069732031207368617265000000000000006044820152606401610b42565b6001600160a01b038316600090815260076020526040812090815460ff1660038111156121d5576121d5615360565b14156122235760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610b42565b600061223884836001015484600201546142f6565b6001600160a01b03841660009081526005840160205260408120805492935091869183916122679084906155ad565b925050819055508483600201600082825461228291906155ad565b925050819055508183600101600082825461229d91906155ad565b9091555060019050835460ff1660038111156122bb576122bb615360565b141561233a576122f56001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168584613cc0565b836001600160a01b0316866001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c84604051610d6a91815260200190565b6003835460ff16600381111561235257612352615360565b141561239c57816003600082825461236a91906155ad565b9250508190555061238f86876001600160a01b0316866001600160a01b031614611f31565b61239c5761239c866140ea565b6002810154600a906123bf9063ffffffff808216916401000000009004166156bc565b63ffffffff16106124125760405162461bcd60e51b815260206004820152601f60248201527f457863656564206d617820756e64656c65676174696f6e20656e7472696573006044820152606401610b42565b60006124278385600301548660040154613a0e565b90508084600401600082825461243d9190615516565b92505081905550828460030160008282546124589190615516565b909155505060028201805463ffffffff640100000000918290048116600090815260018087016020526040909120858155439181019190915583549093929004169060046124a5836155e1565b91906101000a81548163ffffffff021916908363ffffffff16021790555050856001600160a01b0316886001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea876001015486600001548861250d906155c4565b6040805193845260208401929092529082015260600160405180910390a35050505050505050565b6000546001600160a01b0316331461257d5760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b600c546001600160a01b0316156125d65760405162461bcd60e51b815260206004820152601b60248201527f72657761726420636f6e747261637420616c72656164792073657400000000006044820152606401610b42565b600c80546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a01b900460ff16156126455760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b42565b60025460ff16156126af573360009081526001602052604090205460ff166126af5760405162461bcd60e51b815260206004820152601960248201527f43616c6c6572206973206e6f742077686974656c6973746564000000000000006044820152606401610b42565b33600081815260076020526040812090815460ff1660038111156126d5576126d5615360565b146127225760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f7220697320696e697469616c697a656400000000000000006044820152606401610b42565b6001600160a01b03851660009081526007602052604081205460ff16600381111561274f5761274f615360565b1461279c5760405162461bcd60e51b815260206004820152601960248201527f5369676e6572206973206f746865722076616c696461746f72000000000000006044820152606401610b42565b6001600160a01b0382811660009081526008602052604090205416156128045760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206973206f74686572207369676e6572000000000000006044820152606401610b42565b6001600160a01b03858116600090815260086020526040902054161561286c5760405162461bcd60e51b815260206004820152601360248201527f5369676e657220616c72656164792075736564000000000000000000000000006044820152606401610b42565b6127108367ffffffffffffffff1611156128c85760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636f6d6d697373696f6e20726174650000000000000000006044820152606401610b42565b6005600052600a6020527ff35035bc2b01d44bd35a1dcdc552315cffb73da35cfd60570b7b777f98036f9f548410156129435760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610b42565b80547fffffffffffffffffffffff0000000000000000000000000000000000000000001660ff196101006001600160a01b038881169182029290921692909217600190811784556006840187905560078401805467ffffffffffffffff60801b1916600160801b67ffffffffffffffff8916021790556005805491820190557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b03199081169286169283179091556000928352600860205260409092208054909216179055612a1c8285610af9565b604080516001600160a01b03878116602083015291810186905267ffffffffffffffff85166060820152908316907f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c9060800160408051601f1981840301815290829052612a8c916000906156e1565b60405180910390a25050505050565b33600081815260076020526040812090815460ff166003811115612ac157612ac1615360565b1415612b0f5760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610b42565b6005600052600a6020527ff35035bc2b01d44bd35a1dcdc552315cffb73da35cfd60570b7b777f98036f9f54831015612b8a5760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610b42565b8060060154831015612c57576003815460ff166003811115612bae57612bae615360565b1415612bfc5760405162461bcd60e51b815260206004820152601360248201527f56616c696461746f7220697320626f6e646564000000000000000000000000006044820152606401610b42565b6006600052600a6020527f10d9dd018e4cae503383c9f804c1c1603ada5856ee7894375d9b97cd8c8b27db54612c329043615516565b60078201805467ffffffffffffffff191667ffffffffffffffff929092169190911790555b6006810183905560408051602081018590526001600160a01b038416917f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c910160408051601f19818403018152908290526121409160009061570f565b612cc28761080287896155a0565b612d0e5760405162461bcd60e51b815260206004820152601560248201527f4661696c656420746f20766572696679207369677300000000000000000000006044820152606401610b42565b50505050505050565b6000546001600160a01b03163314612d5f5760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b600b546001600160a01b031615612db85760405162461bcd60e51b815260206004820152601860248201527f676f7620636f6e747261637420616c72656164792073657400000000000000006044820152606401610b42565b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b03163314612e225760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b611ae7600061430f565b6001600160a01b03811660009081526007602052604090206002815460ff166003811115612e5c57612e5c615360565b14612ea95760405162461bcd60e51b815260206004820152601760248201527f56616c696461746f72206e6f7420756e626f6e64696e670000000000000000006044820152606401610b42565b600781015468010000000000000000900467ffffffffffffffff16431015612f135760405162461bcd60e51b815260206004820152601860248201527f556e626f6e6420626c6f636b206e6f74207265616368656400000000000000006044820152606401610b42565b805460ff1916600190811782556007820180546fffffffffffffffff0000000000000000191690555b6040516001600160a01b038416907fd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a90600090a35050565b33600081815260076020526040812090815460ff166003811115612f9a57612f9a615360565b1415612fe85760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206e6f7420696e697469616c697a6564000000000000006044820152606401610b42565b6001600160a01b0383811660009081526008602052604090205416156130505760405162461bcd60e51b815260206004820152601360248201527f5369676e657220616c72656164792075736564000000000000000000000000006044820152606401610b42565b816001600160a01b0316836001600160a01b0316146130e3576001600160a01b03831660009081526007602052604081205460ff16600381111561309657613096615360565b146130e35760405162461bcd60e51b815260206004820152601960248201527f5369676e6572206973206f746865722076616c696461746f72000000000000006044820152606401610b42565b8054610100908190046001600160a01b03908116600090815260086020908152604080832080546001600160a01b031990811690915586547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1689861696870217875585845292819020805490931693871693841790925581519081019390935290917f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c910160408051601f198184030181529082905261214091600090615756565b6000600d54116131f85760405162461bcd60e51b815260206004820152601260248201527f4e6f7468696e6720746f20636f6c6c65637400000000000000000000000000006044820152606401610b42565b600c54600d54613236916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811692911690613cc0565b6000600d55565b6000546001600160a01b031633146132855760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b611ae761435f565b6000806132ee84805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600080806132fc611344565b905060005b86518110156134285760006133388883815181106133215761332161552e565b6020026020010151876143e790919063ffffffff16565b9050836001600160a01b0316816001600160a01b03161161339b5760405162461bcd60e51b815260206004820152601e60248201527f5369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610b42565b6001600160a01b038082166000908152600860209081526040808320549093168252600790522090935083906003815460ff1660038111156133df576133df615360565b146133eb575050613416565b60018101546133fa9087615516565b9550838610613413576001975050505050505050611fe5565b50505b8061342081615544565b915050613301565b5060405162461bcd60e51b815260206004820152601260248201527f51756f72756d206e6f74207265616368656400000000000000000000000000006044820152606401610b42565b6001600160a01b038516600090815260076020526040812090815460ff1660038111156134a0576134a0615360565b14156134ee5760405162461bcd60e51b815260206004820152601c60248201527f56616c696461746f72206973206e6f7420696e697469616c697a6564000000006044820152606401610b42565b856001600160a01b03167f3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512c868686863360405161352f9594939291906157c6565b60405180910390a2505050505050565b6005818154811061354f57600080fd5b6000918252602090912001546001600160a01b0316905081565b6006818154811061354f57600080fd5b600060036001600160a01b03831660009081526007602052604090205460ff1660038111156135aa576135aa615360565b1492915050565b600b546001600160a01b0316331461360b5760405162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f7420676f7620636f6e74726163740000000000006044820152606401610b42565b80600a600084600881111561362257613622615360565b600881111561363357613633615360565b81526020810191909152604001600020555050565b61368a6040518060c0016040528060006001600160a01b0316815260200160008152602001600081526020016060815260200160008152602001600081525090565b6001600160a01b03808416600090815260076020908152604080832093861683526005840190915281208054600184015460028501549293926136ce9291906142f6565b60026000908152600a6020527fbff4442b8ed600beeb8e26b1279a0f0d14c6edfaec26d968ee13c86f7d4c2ba85485549293509091829190829060019060ff16600381111561371f5761371f615360565b6002880154911491506000906137469063ffffffff808216916401000000009004166156bc565b63ffffffff16905060008167ffffffffffffffff8111156137695761376961506c565b6040519080825280602002602001820160405280156137ae57816020015b60408051808201909152600080825260208201528152602001906001900390816137875790505b50905060005b828110156138c557600289015460018a01906000906137d99063ffffffff1684615516565b8152602001908152602001600020604051806040016040529081600082015481526020016001820154815250508282815181106138185761381861552e565b60200260200101819052508181815181106138355761383561552e565b6020026020010151600001518761384c9190615516565b96508380613882575043858383815181106138695761386961552e565b60200260200101516020015161387f9190615516565b11155b156138b3578181815181106138995761389961552e565b602002602001015160000151866138b09190615516565b95505b806138bd81615544565b9150506137b4565b5060006138db878b600301548c600401546142f6565b905060006138f2878c600301548d600401546142f6565b90506040518060c001604052808f6001600160a01b031681526020018a81526020018b600001548152602001848152602001838152602001828152509b50505050505050505050505092915050565b6000546001600160a01b031633146139895760405162461bcd60e51b8152602060048201819052602482015260008051602061586c8339815191526044820152606401610b42565b6001600160a01b038116613a055760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610b42565b610dd38161430f565b600082613a1c575082613a34565b82613a27838661555f565b613a31919061557e565b90505b9392505050565b6006546002811480613a4d5750806003145b15613ad257613a5a611344565b8210613ace5760405162461bcd60e51b815260206004820152602e60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201527f2071756f72756d20746f6b656e730000000000000000000000000000000000006064820152608401610b42565b5050565b6003811115613ace5760038054613ae9919061557e565b8210613ace5760405162461bcd60e51b815260206004820152602b60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201527f20312f3320746f6b656e730000000000000000000000000000000000000000006064820152608401610b42565b6040516001600160a01b0380851660248301528316604482015260648101829052613bf59085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261448b565b50505050565b6001600160a01b03811660009081526001602052604090205460ff1615613c645760405162461bcd60e51b815260206004820152601360248201527f416c72656164792077686974656c6973746564000000000000000000000000006044820152606401610b42565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f91015b60405180910390a150565b6040516001600160a01b03831660248201526044810182905261121890849063a9059cbb60e01b90606401613b91565b6001600160a01b03811660009081526001602052604090205460ff16613d585760405162461bcd60e51b815260206004820152600f60248201527f4e6f742077686974656c697374656400000000000000000000000000000000006044820152606401610b42565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b69101613cb5565b600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b0319166001600160a01b038316179055610dd381614570565b613e2d60068281548110613e1357613e1361552e565b6000918252602090912001546001600160a01b03166145d5565b8160068281548110613e4157613e4161552e565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550613ace82614570565b6040805160c0810182526000808252602080830182905282840182905260608084018390526080840183905260a084015283518085019094528184528301849052909190613ec782600661467c565b905080600681518110613edc57613edc61552e565b602002602001015167ffffffffffffffff811115613efc57613efc61506c565b604051908082528060200260200182016040528015613f4157816020015b6040805180820190915260008082526020820152815260200190600190039081613f1a5790505b508360a00181905250600081600681518110613f5f57613f5f61552e565b6020026020010181815250506000805b602084015151845110156140e157613f8684614736565b90925090508160011415613fb557613fa5613fa085614770565b61482d565b6001600160a01b03168552613f6f565b8160021415613fdb57613fc784614838565b67ffffffffffffffff166020860152613f6f565b816003141561400157613fed84614838565b67ffffffffffffffff166040860152613f6f565b81600414156140275761401384614838565b67ffffffffffffffff166060860152613f6f565b816005141561404d5761403984614838565b67ffffffffffffffff166080860152613f6f565b81600614156140d25761406761406285614770565b6148ba565b8560a001518460068151811061407f5761407f61552e565b6020026020010151815181106140975761409761552e565b6020026020010181905250826006815181106140b5576140b561552e565b6020026020010180518091906140ca90615544565b905250613f6f565b6140dc8482614961565b613f6f565b50505050919050565b6006546000906140fc906001906155ad565b905060005b60065481101561420757826001600160a01b0316600682815481106141285761412861552e565b6000918252602090912001546001600160a01b031614156141f557818110156141b9576006828154811061415e5761415e61552e565b600091825260209091200154600680546001600160a01b03909216918390811061418a5761418a61552e565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60068054806141ca576141ca615809565b600082815260209020810160001990810180546001600160a01b0319169055019055611218836145d5565b806141ff81615544565b915050614101565b5060405162461bcd60e51b815260206004820152601460248201527f4e6f7420626f6e6465642076616c696461746f720000000000000000000000006044820152606401610b42565b600054600160a01b900460ff166142a95760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610b42565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600081614304575082613a34565b81613a27848661555f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff16156143ac5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610b42565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586142d93390565b600081516041141561441b5760208201516040830151606084015160001a614411868285856149d3565b9350505050611fe5565b815160401415614443576020820151604083015161443a858383614b7c565b92505050611fe5565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610b42565b60006144e0826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614bbf9092919063ffffffff16565b80519091501561121857808060200190518101906144fe919061581f565b6112185760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610b42565b6001600160a01b03811660009081526007602081905260408220805460ff19166003908117825591810180546fffffffffffffffff0000000000000000191690556001810154825491939092916145c8908490615516565b9091555060039050612f3c565b6001600160a01b03811660009081526007602090815260408220805460ff191660029081178255909252600a90527fbff4442b8ed600beeb8e26b1279a0f0d14c6edfaec26d968ee13c86f7d4c2ba85461462f9043615516565b8160070160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080600101546003600082825461466f91906155ad565b9091555060029050612f3c565b815160609061468c836001615516565b67ffffffffffffffff8111156146a4576146a461506c565b6040519080825280602002602001820160405280156146cd578160200160208202803683370190505b5091506000805b6020860151518651101561472d576146eb86614736565b809250819350505060018483815181106147075761470761552e565b6020026020010181815161471b9190615516565b9052506147288682614961565b6146d4565b50509092525090565b600080600061474484614838565b905061475160088261557e565b925080600716600581111561476857614768615360565b915050915091565b6060600061477d83614838565b905060008184600001516147919190615516565b90508360200151518111156147a557600080fd5b8167ffffffffffffffff8111156147be576147be61506c565b6040519080825280601f01601f1916602001820160405280156147e8576020820181803683370190505b50602080860151865192955091818601919083010160005b8581101561482257818101518382015261481b602082615516565b9050614800565b505050935250919050565b6000611fe582614bce565b602080820151825181019091015160009182805b600a8110156148b45783811a915061486581600761555f565b82607f16901b8517945081608016600014156148a257614886816001615516565b86518790614895908390615516565b9052509395945050505050565b806148ac81615544565b91505061484c565b50600080fd5b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b60208301515183511015614959576148fc83614736565b9092509050816001141561492657614916613fa084614770565b6001600160a01b031684526148e5565b816002141561494a5761494061493b84614770565b614bf6565b60208501526148e5565b6149548382614961565b6148e5565b505050919050565b600081600581111561497557614975615360565b14156149845761121882614838565b600281600581111561499857614998615360565b14156103555760006149a983614838565b905080836000018181516149bd9190615516565b9052506020830151518351111561121857600080fd5b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115614a505760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610b42565b8360ff16601b1480614a6557508360ff16601c145b614abc5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610b42565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015614b10573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116614b735760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610b42565b95945050505050565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821660ff83901c601b01614bb5868287856149d3565b9695505050505050565b6060613a318484600085614c2d565b60008151601414614bde57600080fd5b50602001516c01000000000000000000000000900490565b6000602082511115614c0757600080fd5b6020820151905081516020614c1c91906155ad565b614c2790600861555f565b1c919050565b606082471015614ca55760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610b42565b843b614cf35760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610b42565b600080866001600160a01b03168587604051614d0f919061583c565b60006040518083038185875af1925050503d8060008114614d4c576040519150601f19603f3d011682016040523d82523d6000602084013e614d51565b606091505b5091509150614d61828286614d6c565b979650505050505050565b60608315614d7b575081613a34565b825115614d8b5782518084602001fd5b8160405162461bcd60e51b8152600401610b429190615858565b80356001600160a01b0381168114614dbc57600080fd5b919050565b60008060408385031215614dd457600080fd5b614ddd83614da5565b946020939093013593505050565b8015158114610dd357600080fd5b600060208284031215614e0b57600080fd5b8135613a3481614deb565b600060208284031215614e2857600080fd5b613a3482614da5565b600060208284031215614e4357600080fd5b5035919050565b60008083601f840112614e5c57600080fd5b50813567ffffffffffffffff811115614e7457600080fd5b602083019150836020828501011115614e8c57600080fd5b9250929050565b60008083601f840112614ea557600080fd5b50813567ffffffffffffffff811115614ebd57600080fd5b6020830191508360208260051b8501011115614e8c57600080fd5b60008060008060408587031215614eee57600080fd5b843567ffffffffffffffff80821115614f0657600080fd5b614f1288838901614e4a565b90965094506020870135915080821115614f2b57600080fd5b50614f3887828801614e93565b95989497509550505050565b602080825282518282018190526000919060409081850190868401855b82811015614f8f57815180516001600160a01b03168552860151868501529284019290850190600101614f61565b5091979650505050505050565b803560098110614dbc57600080fd5b600060208284031215614fbd57600080fd5b613a3482614f9c565b60008060408385031215614fd957600080fd5b614fe283614da5565b91506020830135614ff281614deb565b809150509250929050565b803567ffffffffffffffff81168114614dbc57600080fd5b60006020828403121561502757600080fd5b613a3482614ffd565b60008060006060848603121561504557600080fd5b61504e84614da5565b92506020840135915061506360408501614ffd565b90509250925092565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156150ab576150ab61506c565b604052919050565b600082601f8301126150c457600080fd5b813567ffffffffffffffff8111156150de576150de61506c565b6150f1601f8201601f1916602001615082565b81815284602083860101111561510657600080fd5b816020850160208301376000918101602001919091529392505050565b60008060008060008060006080888a03121561513e57600080fd5b873567ffffffffffffffff8082111561515657600080fd5b6151628b838c016150b3565b985060208a013591508082111561517857600080fd5b6151848b838c01614e93565b909850965060408a013591508082111561519d57600080fd5b6151a98b838c01614e93565b909650945060608a01359150808211156151c257600080fd5b506151cf8a828b01614e93565b989b979a50959850939692959293505050565b600067ffffffffffffffff808411156151fd576151fd61506c565b8360051b602061520e818301615082565b8681529350908401908084018783111561522757600080fd5b855b8381101561525b578035858111156152415760008081fd5b61524d8a828a016150b3565b835250908201908201615229565b50505050509392505050565b6000806040838503121561527a57600080fd5b823567ffffffffffffffff8082111561529257600080fd5b61529e868387016150b3565b935060208501359150808211156152b457600080fd5b508301601f810185136152c657600080fd5b6152d5858235602084016151e2565b9150509250929050565b6000806000806000606086880312156152f757600080fd5b61530086614da5565b9450602086013567ffffffffffffffff8082111561531d57600080fd5b61532989838a01614e4a565b9096509450604088013591508082111561534257600080fd5b5061534f88828901614e4a565b969995985093965092949392505050565b634e487b7160e01b600052602160045260246000fd5b6004811061539457634e487b7160e01b600052602160045260246000fd5b9052565b60208101611fe58284615376565b600080604083850312156153b957600080fd5b614ddd83614f9c565b600080604083850312156153d557600080fd5b6153de83614da5565b91506153ec60208401614da5565b90509250929050565b6000602080835260e083016001600160a01b038551168285015281850151604081818701528087015160608701526060870151915060c06080870152828251808552610100880191508584019450600093505b808410156154715784518051835286015186830152938501936001939093019290820190615448565b50608088015160a088015260a088015160c0880152809550505050505092915050565b61014081016154a3828d615376565b6001600160a01b039a909a16602082015260408101989098526060880196909652608087019490945260a086019290925260c085015267ffffffffffffffff90811660e08501529081166101008401521661012090910152919050565b634e487b7160e01b600052601160045260246000fd5b6000821982111561552957615529615500565b500190565b634e487b7160e01b600052603260045260246000fd5b600060001982141561555857615558615500565b5060010190565b600081600019048311821515161561557957615579615500565b500290565b60008261559b57634e487b7160e01b600052601260045260246000fd5b500490565b6000613a343684846151e2565b6000828210156155bf576155bf615500565b500390565b6000600160ff1b8214156155da576155da615500565b5060000390565b600063ffffffff808316818114156155fb576155fb615500565b6001019392505050565b60005b83811015615620578181015183820152602001615608565b83811115613bf55750506000910152565b60008151808452615649816020860160208601615605565b601f01601f19169290920160200192915050565b60608152600a60608201527f636f6d6d697373696f6e00000000000000000000000000000000000000000000608082015260a0602082015260006156a460a0830185615631565b90506001600160a01b03831660408301529392505050565b600063ffffffff838116908316818110156156d9576156d9615500565b039392505050565b6060815260046060820152631a5b9a5d60e21b608082015260a0602082015260006156a460a0830185615631565b60608152601360608201527f6d696e2d73656c662d64656c65676174696f6e00000000000000000000000000608082015260a0602082015260006156a460a0830185615631565b60608152600660608201527f7369676e65720000000000000000000000000000000000000000000000000000608082015260a0602082015260006156a460a0830185615631565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6060815260006157da60608301878961579d565b82810360208401526157ed81868861579d565b9150506001600160a01b03831660408301529695505050505050565b634e487b7160e01b600052603160045260246000fd5b60006020828403121561583157600080fd5b8151613a3481614deb565b6000825161584e818460208701615605565b9190910192915050565b602081526000613a34602083018461563156fe4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a2646970667358221220a466f5b90af7d6b6ed2cc9f63c5557a2f5f4b333eccbc63dfb889e8d19f98c5864736f6c63430008090033",
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

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingTransactor) Undelegate(opts *bind.TransactOpts, _valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "undelegate", _valAddr, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingSession) Undelegate(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, _valAddr, _shares)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _valAddr, uint256 _shares) returns()
func (_Staking *StakingTransactorSession) Undelegate(_valAddr common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Undelegate(&_Staking.TransactOpts, _valAddr, _shares)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"StakingRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"}],\"name\":\"StakingRewardContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516115aa3803806115aa83398101604081905261002f916100a6565b61003833610056565b6000805460ff60a01b191690556001600160a01b03166080526100d6565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100b857600080fd5b81516001600160a01b03811681146100cf57600080fd5b9392505050565b60805161149d61010d6000396000818161010301528181610233015281816103d4015281816106e30152610869015261149d6000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063715018a61161008157806396db0fef1161005b57806396db0fef14610180578063f2fde38b146101ae578063f8df0dc5146101c157600080fd5b8063715018a61461015f5780638456cb59146101675780638da5cb5b1461016f57600080fd5b80633f4ba83a116100b25780633f4ba83a146100f65780634cf088d9146100fe5780635c975abb1461014257600080fd5b80630a300b09146100ce578063145aa116146100e3575b600080fd5b6100e16100dc3660046110f6565b6101d4565b005b6100e16100f13660046110f6565b61031a565b6100e1610476565b6101257f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b600054600160a01b900460ff166040519015158152602001610139565b6100e16104da565b6100e161053e565b6000546001600160a01b0316610125565b6101a061018e366004611124565b60016020526000908152604090205481565b604051908152602001610139565b6100e16101bc366004611124565b6105a0565b6100e16101cf366004611141565b61067f565b600054600160a01b900460ff16156102265760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064015b60405180910390fd5b60003390506102d38130847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561028a57600080fd5b505afa15801561029e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c29190611206565b6001600160a01b0316929190610910565b806001600160a01b03167ff67017a05194c0853be9169be60cad9fa6e75d34b6b507a7a4261510e19c6d798360405161030e91815260200190565b60405180910390a25050565b600054600160a01b900460ff166103735760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161021d565b6000546001600160a01b031633146103cd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161021d565b61047333827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561042b57600080fd5b505afa15801561043f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104639190611206565b6001600160a01b031691906109ae565b50565b6000546001600160a01b031633146104d05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161021d565b6104d86109e3565b565b6000546001600160a01b031633146105345760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161021d565b6104d86000610a89565b6000546001600160a01b031633146105985760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161021d565b6104d8610af1565b6000546001600160a01b031633146105fa5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161021d565b6001600160a01b0381166106765760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161021d565b61047381610a89565b600054600160a01b900460ff16156106cc5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161021d565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe9061071e90879087908790879060040161124c565b60206040518083038186803b15801561073657600080fd5b505afa15801561074a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076e91906112fd565b5060006107b085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b7992505050565b60208082015182516001600160a01b0316600090815260019092526040822054929350916107de9083611335565b9050600081116108305760405162461bcd60e51b815260206004820152600d60248201527f4e6f206e65772072657761726400000000000000000000000000000000000000604482015260640161021d565b816001600085600001516001600160a01b03166001600160a01b03168152602001908152602001600020819055506108c08360000151827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561042b57600080fd5b82600001516001600160a01b03167f6dd401e61ba732582a5eba3d54ccc3afb3609cd3ac1a166d1d36f75fc0aedcda826040516108ff91815260200190565b60405180910390a250505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526109a89085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610c25565b50505050565b6040516001600160a01b0383166024820152604481018290526109de90849063a9059cbb60e01b90606401610944565b505050565b600054600160a01b900460ff16610a3c5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015260640161021d565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff1615610b3e5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161021d565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610a6c3390565b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b60208301515183511015610c1d57610bbb83610d0a565b90925090508160011415610bea57610bda610bd584610d44565b610e01565b6001600160a01b03168452610ba4565b8160021415610c0e57610c04610bff84610d44565b610e12565b6020850152610ba4565b610c188382610e49565b610ba4565b505050919050565b6000610c7a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610ebb9092919063ffffffff16565b8051909150156109de5780806020019051810190610c9891906112fd565b6109de5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161021d565b6000806000610d1884610ed4565b9050610d2560088261134c565b9250806007166005811115610d3c57610d3c61136e565b915050915091565b60606000610d5183610ed4565b90506000818460000151610d659190611384565b9050836020015151811115610d7957600080fd5b8167ffffffffffffffff811115610d9257610d9261139c565b6040519080825280601f01601f191660200182016040528015610dbc576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015610df6578181015183820152610def602082611384565b9050610dd4565b505050935250919050565b6000610e0c82610f56565b92915050565b6000602082511115610e2357600080fd5b6020820151905081516020610e389190611335565b610e439060086113b2565b1c919050565b6000816005811115610e5d57610e5d61136e565b1415610e6c576109de82610ed4565b6002816005811115610e8057610e8061136e565b14156100c9576000610e9183610ed4565b90508083600001818151610ea59190611384565b905250602083015151835111156109de57600080fd5b6060610eca8484600085610f7e565b90505b9392505050565b602080820151825181019091015160009182805b600a811015610f505783811a9150610f018160076113b2565b82607f16901b851794508160801660001415610f3e57610f22816001611384565b86518790610f31908390611384565b9052509395945050505050565b80610f48816113d1565b915050610ee8565b50600080fd5b60008151601414610f6657600080fd5b50602001516c01000000000000000000000000900490565b606082471015610ff65760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161021d565b843b6110445760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161021d565b600080866001600160a01b031685876040516110609190611418565b60006040518083038185875af1925050503d806000811461109d576040519150601f19603f3d011682016040523d82523d6000602084013e6110a2565b606091505b50915091506110b28282866110bd565b979650505050505050565b606083156110cc575081610ecd565b8251156110dc5782518084602001fd5b8160405162461bcd60e51b815260040161021d9190611434565b60006020828403121561110857600080fd5b5035919050565b6001600160a01b038116811461047357600080fd5b60006020828403121561113657600080fd5b8135610ecd8161110f565b6000806000806040858703121561115757600080fd5b843567ffffffffffffffff8082111561116f57600080fd5b818701915087601f83011261118357600080fd5b81358181111561119257600080fd5b8860208285010111156111a457600080fd5b6020928301965094509086013590808211156111bf57600080fd5b818701915087601f8301126111d357600080fd5b8135818111156111e257600080fd5b8860208260051b85010111156111f757600080fd5b95989497505060200194505050565b60006020828403121561121857600080fd5b8151610ecd8161110f565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000611260604083018688611223565b602083820381850152818583528183019050818660051b8401018760005b888110156112ed57858303601f190184528135368b9003601e190181126112a457600080fd5b8a01803567ffffffffffffffff8111156112bd57600080fd5b8036038c13156112cc57600080fd5b6112d98582898501611223565b95870195945050509084019060010161127e565b50909a9950505050505050505050565b60006020828403121561130f57600080fd5b81518015158114610ecd57600080fd5b634e487b7160e01b600052601160045260246000fd5b6000828210156113475761134761131f565b500390565b60008261136957634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b600082198211156113975761139761131f565b500190565b634e487b7160e01b600052604160045260246000fd5b60008160001904831182151516156113cc576113cc61131f565b500290565b60006000198214156113e5576113e561131f565b5060010190565b60005b838110156114075781810151838201526020016113ef565b838111156109a85750506000910152565b6000825161142a8184602087016113ec565b9190910192915050565b60208152600082518060208401526114538160408501602087016113ec565b601f01601f1916919091016040019291505056fea2646970667358221220bd63b9130944326c2cfc811203e3b2fe871fd59ea61e972f4bf0cc4384343ec564736f6c63430008090033",
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
	Bin: "0x60a060405234801561001057600080fd5b506040516117c73803806117c783398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b6080516116db6100ec6000396000818160bb01528181610199015281816102eb0152818161039701528181610633015281816107790152818161089e0152818161099b01528181610a4401528181610b0001528181610cf401528181610e3a01528181610ecf01528181610fec015261109201526116db6000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638dc2336d1161005b5780638dc2336d1461012a578063c6fc1ed614610140578063d87ffe9114610168578063e9fe6b0b1461017057600080fd5b8063313019bb1461008d5780634cf088d9146100b657806366ab5d28146100f55780638a11d7c91461010a575b600080fd5b6100a061009b36600461115b565b610193565b6040516100ad919061117f565b60405180910390f35b6100dd7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ad565b6100fd61062d565b6040516100ad91906112d7565b61011d61011836600461115b565b610839565b6040516100ad9190611319565b610132610996565b6040519081526020016100ad565b61015361014e36600461115b565b610c49565b604080519283526020830191909152016100ad565b6100fd610cee565b61018361017e36600461115b565b610ea7565b60405190151581526020016100ad565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631cfe4f0b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101f057600080fd5b505afa158015610204573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610228919061132d565b905060008167ffffffffffffffff81111561024557610245611346565b6040519080825280602002602001820160405280156102b857816020015b6102a56040518060c0016040528060006001600160a01b0316815260200160008152602001600081526020016060815260200160008152602001600081525090565b8152602001906001900390816102635790505b5090506000805b838163ffffffff1610156104b7576040516324aec90f60e21b815263ffffffff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906392bb243c9060240160206040518083038186803b15801561033557600080fd5b505afa158015610349573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036d919061135c565b604051631dd9dfdf60e31b81526001600160a01b03808316600483015289811660248301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063eecefef89060440160006040518083038186803b1580156103db57600080fd5b505afa1580156103ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261041791908101906113f6565b848363ffffffff168151811061042f5761042f611527565b6020026020010181905250838263ffffffff168151811061045257610452611527565b60200260200101516040015160001415806104915750838263ffffffff168151811061048057610480611527565b602002602001015160800151600014155b156104a457826104a081611553565b9350505b50806104af81611553565b9150506102bf565b5060008163ffffffff1667ffffffffffffffff8111156104d9576104d9611346565b60405190808252806020026020018201604052801561054c57816020015b6105396040518060c0016040528060006001600160a01b0316815260200160008152602001600081526020016060815260200160008152602001600081525090565b8152602001906001900390816104f75790505b5090506000805b858163ffffffff16101561062157848163ffffffff168151811061057957610579611527565b60200260200101516040015160001415806105b85750848163ffffffff16815181106105a7576105a7611527565b602002602001015160800151600014155b1561060f57848163ffffffff16815181106105d5576105d5611527565b6020026020010151838363ffffffff16815181106105f5576105f5611527565b6020026020010181905250818061060b90611553565b9250505b8061061981611553565b915050610553565b50909695505050505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631cfe4f0b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561068a57600080fd5b505afa15801561069e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c2919061132d565b905060008167ffffffffffffffff8111156106df576106df611346565b60405190808252806020026020018201604052801561074657816020015b6040805160e08101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c082015282526000199092019101816106fd5790505b50905060005b828163ffffffff161015610832576040516324aec90f60e21b815263ffffffff821660048201526107fc907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906392bb243c906024015b60206040518083038186803b1580156107c457600080fd5b505afa1580156107d8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610118919061135c565b828263ffffffff168151811061081457610814611527565b6020026020010181905250808061082a90611553565b91505061074c565b5092915050565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152604051631f4a58fb60e31b81526001600160a01b038381166004830152600091829182918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063fa52c7d8906024016101406040518083038186803b1580156108e157600080fd5b505afa1580156108f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109199190611594565b995050509750505095509550955095506040518060e00160405280896001600160a01b0316815260200187600381111561095557610955611254565b8152602001866001600160a01b031681526020018581526020018481526020018381526020018267ffffffffffffffff168152509650505050505050919050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166389f9aab56040518163ffffffff1660e01b815260040160206040518083038186803b1580156109f257600080fd5b505afa158015610a06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2a919061132d565b60405163eb505dd560e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb505dd590610a7a90600390600401611636565b60206040518083038186803b158015610a9257600080fd5b505afa158015610aa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aca919061132d565b811015610ad957600091505090565b60001960005b828110156108325760405163acc62ccf60e01b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063c8f9f98490829063acc62ccf9060240160206040518083038186803b158015610b5257600080fd5b505afa158015610b66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8a919061135c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b03909116600482015260240160206040518083038186803b158015610be157600080fd5b505afa158015610bf5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c19919061132d565b905082811015610c365791508180610c3657600094505050505090565b5080610c4181611650565b915050610adf565b6000806000610c5784610193565b905060008060005b83518163ffffffff161015610ce257838163ffffffff1681518110610c8657610c86611527565b60200260200101516020015183610c9d919061166b565b9250838163ffffffff1681518110610cb757610cb7611527565b60200260200101516080015182610cce919061166b565b915080610cda81611553565b915050610c5f565b50909590945092505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166389f9aab56040518163ffffffff1660e01b815260040160206040518083038186803b158015610d4b57600080fd5b505afa158015610d5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d83919061132d565b905060008167ffffffffffffffff811115610da057610da0611346565b604051908082528060200260200182016040528015610e0757816020015b6040805160e08101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c08201528252600019909201910181610dbe5790505b50905060005b828163ffffffff1610156108325760405163acc62ccf60e01b815263ffffffff82166004820152610e71907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063acc62ccf906024016107ac565b828263ffffffff1681518110610e8957610e89611527565b60200260200101819052508080610e9f90611553565b915050610e0d565b604051631f4a58fb60e31b81526001600160a01b0382811660048301526000918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063fa52c7d8906024016101406040518083038186803b158015610f1257600080fd5b505afa158015610f26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4a9190611594565b5050975050505050935050925060006003811115610f6a57610f6a611254565b836003811115610f7c57610f7c611254565b1480610f9957506003836003811115610f9757610f97611254565b145b15610fa957506000949350505050565b8067ffffffffffffffff16431015610fc657506000949350505050565b6040516347abfdbf60e01b81526001600160a01b038681166004830152600160248301527f000000000000000000000000000000000000000000000000000000000000000016906347abfdbf9060440160206040518083038186803b15801561102e57600080fd5b505afa158015611042573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110669190611683565b61107557506000949350505050565b61107d610996565b821161108e57506000949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166383cfb3186040518163ffffffff1660e01b815260040160206040518083038186803b1580156110e957600080fd5b505afa1580156110fd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611121919061132d565b9050804310156111375750600095945050505050565b50600195945050505050565b6001600160a01b038116811461115857600080fd5b50565b60006020828403121561116d57600080fd5b813561117881611143565b9392505050565b60006020808301818452808551808352604092508286019150828160051b8701018488016000805b8481101561124557898403603f19018652825180516001600160a01b031685528881015189860152878101518886015260608082015160c0918701829052805191870182905260e0870191908b0190855b8181101561121d578251805185528d01518d850152928b0192918c01916001016111f8565b5050506080828101519087015260a091820151919095015294870194918701916001016111a7565b50919998505050505050505050565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b0380825116835260208201516004811061128d5761128d611254565b8060208501525080604083015116604084015250606081015160608301526080810151608083015260a081015160a083015267ffffffffffffffff60c08201511660c08301525050565b6020808252825182820181905260009190848201906040850190845b818110156106215761130683855161126a565b9284019260e092909201916001016112f3565b60e08101611327828461126a565b92915050565b60006020828403121561133f57600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561136e57600080fd5b815161117881611143565b60405160c0810167ffffffffffffffff8111828210171561139c5761139c611346565b60405290565b6040805190810167ffffffffffffffff8111828210171561139c5761139c611346565b604051601f8201601f1916810167ffffffffffffffff811182821017156113ee576113ee611346565b604052919050565b6000602080838503121561140957600080fd5b825167ffffffffffffffff8082111561142157600080fd5b9084019060c0828703121561143557600080fd5b61143d611379565b825161144881611143565b815282840151848201526040808401518183015260608401518381111561146e57600080fd5b8401601f8101891361147f57600080fd5b80518481111561149157611491611346565b61149f878260051b016113c5565b818152878101955060069190911b82018701908a8211156114bf57600080fd5b918701915b818310156114ff5783838c0312156114dc5760008081fd5b6114e46113a2565b835181528884015189820152865294870194918301916114c4565b60608501525050506080838101519082015260a0928301519281019290925250949350505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600063ffffffff8083168181141561156d5761156d61153d565b6001019392505050565b805167ffffffffffffffff8116811461158f57600080fd5b919050565b6000806000806000806000806000806101408b8d0312156115b457600080fd5b8a51600481106115c357600080fd5b60208c0151909a506115d481611143565b8099505060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935061160760e08c01611577565b92506116166101008c01611577565b91506116256101208c01611577565b90509295989b9194979a5092959850565b602081016009831061164a5761164a611254565b91905290565b60006000198214156116645761166461153d565b5060010190565b6000821982111561167e5761167e61153d565b500190565b60006020828403121561169557600080fd5b8151801515811461117857600080fdfea2646970667358221220c4f2152cc6969e96940471c16684ea05409bc2ba9387cee133ff186d78c5d04464736f6c63430008090033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
