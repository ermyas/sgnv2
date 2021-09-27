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
	ValAddr            common.Address
	Tokens             *big.Int
	Shares             *big.Int
	UndelegationTokens *big.Int
	Undelegations      []DataTypesUndelegation
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ae8595734093379c5922ffe2978b859a92725cadcce347d08a31bdaaa9fd38be64736f6c63430008070033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"FarmingRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"}],\"name\":\"FarmingRewardContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_rewardsRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signers\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516116a53803806116a583398101604081905261002f916100aa565b6100383361005a565b6000805460ff60a01b1916905560601b6001600160601b0319166080526100da565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bc57600080fd5b81516001600160a01b03811681146100d357600080fd5b9392505050565b60805160601c6115a66100ff60003960008181610186015261026d01526115a66000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063825168ff11610071578063825168ff1461012e5780638456cb59146101415780638da5cb5b146101495780639d4323be1461016e578063ccf2683b14610181578063f2fde38b146101a857600080fd5b80631744092e146100ae5780633f4ba83a146100ec5780635c975abb146100f6578063715018a6146101135780637c81a9da1461011b575b600080fd5b6100d96100bc36600461117c565b600160209081526000928352604080842090915290825290205481565b6040519081526020015b60405180910390f35b6100f46101bb565b005b600054600160a01b900460ff1660405190151581526020016100e3565b6100f46101f8565b6100f46101293660046111fb565b61022c565b6100f461013c3660046111af565b6104d9565b6100f461056b565b6000546001600160a01b03165b6040516001600160a01b0390911681526020016100e3565b6100f461017c3660046111af565b61059d565b6101567f000000000000000000000000000000000000000000000000000000000000000081565b6100f46101b6366004611161565b61062f565b6000546001600160a01b031633146101ee5760405162461bcd60e51b81526004016101e59061142c565b60405180910390fd5b6101f66106ca565b565b6000546001600160a01b031633146102225760405162461bcd60e51b81526004016101e59061142c565b6101f66000610767565b600054600160a01b900460ff16156102565760405162461bcd60e51b81526004016101e590611402565b6040516335b46d1760e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636b68da2e906102ac90899089908990899089908990600401611309565b60006040518083038186803b1580156102c457600080fd5b505afa1580156102d8573d6000803e3d6000fd5b50505050600061031d87878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107b792505050565b9050468160200151146103665760405162461bcd60e51b8152602060048201526011602482015270086d0c2d2dc40928840dad2e6dac2e8c6d607b1b60448201526064016101e5565b6000805b8260400151518110156104915760008360400151828151811061038f5761038f611544565b602002602001015190506000846060015183815181106103b1576103b1611544565b60209081029190910181015186516001600160a01b0390811660009081526001845260408082209287168252919093528220549092506103f190836114ba565b9050801561047b5785516001600160a01b039081166000908152600160208181526040808420948816808552949091529091208490558751909650610437919083610a99565b85516040518281526001600160a01b038086169216907f97e6c3172350795e26977663112f38653689372e771e85bad9fbadb1af0e98b29060200160405180910390a35b5050508080610489906114fd565b91505061036a565b50806104cf5760405162461bcd60e51b815260206004820152600d60248201526c139bc81b995dc81c995dd85c99609a1b60448201526064016101e5565b5050505050505050565b600054600160a01b900460ff16156105035760405162461bcd60e51b81526004016101e590611402565b336105196001600160a01b038416823085610b01565b826001600160a01b0316816001600160a01b03167f40aa1b9a9157bc37a09a78d5a46e53087b82ee0034ebe896d4d1a52f31b333d48460405161055e91815260200190565b60405180910390a3505050565b6000546001600160a01b031633146105955760405162461bcd60e51b81526004016101e59061142c565b6101f6610b3f565b600054600160a01b900460ff166105ed5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016101e5565b6000546001600160a01b031633146106175760405162461bcd60e51b81526004016101e59061142c565b61062b6001600160a01b0383163383610a99565b5050565b6000546001600160a01b031633146106595760405162461bcd60e51b81526004016101e59061142c565b6001600160a01b0381166106be5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101e5565b6106c781610767565b50565b600054600160a01b900460ff1661071a5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016101e5565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6107eb604051806080016040528060006001600160a01b031681526020016000815260200160608152602001606081525090565b6040805180820190915260008082526020820184905261080c826004610ba4565b90508060038151811061082157610821611544565b602002602001015167ffffffffffffffff8111156108415761084161155a565b60405190808252806020026020018201604052801561086a578160200160208202803683370190505b50836040018190525060008160038151811061088857610888611544565b602002602001018181525050806004815181106108a7576108a7611544565b602002602001015167ffffffffffffffff8111156108c7576108c761155a565b6040519080825280602002602001820160405280156108f0578160200160208202803683370190505b50836060018190525060008160048151811061090e5761090e611544565b6020026020010181815250506000805b60208401515184511015610a905761093584610c5e565b909250905081600114156109645761095461094f85610c98565b610d55565b6001600160a01b0316855261091e565b81600214156109885761097e61097985610c98565b610d66565b602086015261091e565b8160031415610a1d5761099d61094f85610c98565b8560400151846003815181106109b5576109b5611544565b6020026020010151815181106109cd576109cd611544565b60200260200101906001600160a01b031690816001600160a01b03168152505082600381518110610a0057610a00611544565b602002602001018051809190610a15906114fd565b90525061091e565b8160041415610a8157610a3261097985610c98565b856060015184600481518110610a4a57610a4a611544565b602002602001015181518110610a6257610a62611544565b60200260200101818152505082600481518110610a0057610a00611544565b610a8b8482610d9d565b61091e565b50505050919050565b6040516001600160a01b038316602482015260448101829052610afc90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610e0f565b505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610b399085906323b872dd60e01b90608401610ac5565b50505050565b600054600160a01b900460ff1615610b695760405162461bcd60e51b81526004016101e590611402565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861074a3390565b8151606090610bb4836001611461565b67ffffffffffffffff811115610bcc57610bcc61155a565b604051908082528060200260200182016040528015610bf5578160200160208202803683370190505b5091506000805b60208601515186511015610c5557610c1386610c5e565b80925081935050506001848381518110610c2f57610c2f611544565b60200260200101818151610c439190611461565b905250610c508682610d9d565b610bfc565b50509092525090565b6000806000610c6c84610ee1565b9050610c79600882611479565b9250806007166005811115610c9057610c9061152e565b915050915091565b60606000610ca583610ee1565b90506000818460000151610cb99190611461565b9050836020015151811115610ccd57600080fd5b8167ffffffffffffffff811115610ce657610ce661155a565b6040519080825280601f01601f191660200182016040528015610d10576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015610d4a578181015183820152610d43602082611461565b9050610d28565b505050935250919050565b6000610d6082610f63565b92915050565b6000602082511115610d7757600080fd5b6020820151905081516020610d8c91906114ba565b610d9790600861149b565b1c919050565b6000816005811115610db157610db161152e565b1415610dc057610afc82610ee1565b6002816005811115610dd457610dd461152e565b14156100a9576000610de583610ee1565b90508083600001818151610df99190611461565b90525060208301515183511115610afc57600080fd5b6000610e64826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610f829092919063ffffffff16565b805190915015610afc5780806020019051810190610e8291906111d9565b610afc5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016101e5565b602080820151825181019091015160009182805b600a811015610f5d5783811a9150610f0e81600761149b565b82607f16901b851794508160801660001415610f4b57610f2f816001611461565b86518790610f3e908390611461565b9052509395945050505050565b80610f55816114fd565b915050610ef5565b50600080fd5b60008151601414610f7357600080fd5b5060200151600160601b900490565b6060610f918484600085610f9b565b90505b9392505050565b606082471015610ffc5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016101e5565b843b61104a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101e5565b600080866001600160a01b0316858760405161106691906112ed565b60006040518083038185875af1925050503d80600081146110a3576040519150601f19603f3d011682016040523d82523d6000602084013e6110a8565b606091505b50915091506110b88282866110c3565b979650505050505050565b606083156110d2575081610f94565b8251156110e25782518084602001fd5b8160405162461bcd60e51b81526004016101e591906113cf565b80356001600160a01b038116811461111357600080fd5b919050565b60008083601f84011261112a57600080fd5b50813567ffffffffffffffff81111561114257600080fd5b60208301915083602082850101111561115a57600080fd5b9250929050565b60006020828403121561117357600080fd5b610f94826110fc565b6000806040838503121561118f57600080fd5b611198836110fc565b91506111a6602084016110fc565b90509250929050565b600080604083850312156111c257600080fd5b6111cb836110fc565b946020939093013593505050565b6000602082840312156111eb57600080fd5b81518015158114610f9457600080fd5b6000806000806000806060878903121561121457600080fd5b863567ffffffffffffffff8082111561122c57600080fd5b6112388a838b01611118565b9098509650602089013591508082111561125157600080fd5b61125d8a838b01611118565b9096509450604089013591508082111561127657600080fd5b818901915089601f83011261128a57600080fd5b81358181111561129957600080fd5b8a60208260051b85010111156112ae57600080fd5b6020830194508093505050509295509295509295565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600082516112ff8184602087016114d1565b9190910192915050565b60608152600061131d60608301888a6112c4565b60208382038185015261133182888a6112c4565b84810360408601528581529150808201600586901b830182018760005b888110156113bd57858303601f190184528135368b9003601e1901811261137457600080fd5b8a01803567ffffffffffffffff81111561138d57600080fd5b8036038c131561139c57600080fd5b6113a985828985016112c4565b95870195945050509084019060010161134e565b50909c9b505050505050505050505050565b60208152600082518060208401526113ee8160408501602087016114d1565b601f01601f19169190910160400192915050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000821982111561147457611474611518565b500190565b60008261149657634e487b7160e01b600052601260045260246000fd5b500490565b60008160001904831182151516156114b5576114b5611518565b500290565b6000828210156114cc576114cc611518565b500390565b60005b838110156114ec5781810151838201526020016114d4565b83811115610b395750506000910152565b600060001982141561151157611511611518565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212204e8caba84c7613488f384ba6033fc8fbf5089cd85820aeab940f1a98386fec1364736f6c63430008070033",
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

// ClaimRewards is a paid mutator transaction binding the contract method 0x7c81a9da.
//
// Solidity: function claimRewards(bytes _rewardsRequest, bytes _signers, bytes[] _sigs) returns()
func (_FarmingRewards *FarmingRewardsTransactor) ClaimRewards(opts *bind.TransactOpts, _rewardsRequest []byte, _signers []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _FarmingRewards.contract.Transact(opts, "claimRewards", _rewardsRequest, _signers, _sigs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x7c81a9da.
//
// Solidity: function claimRewards(bytes _rewardsRequest, bytes _signers, bytes[] _sigs) returns()
func (_FarmingRewards *FarmingRewardsSession) ClaimRewards(_rewardsRequest []byte, _signers []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _FarmingRewards.Contract.ClaimRewards(&_FarmingRewards.TransactOpts, _rewardsRequest, _signers, _sigs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x7c81a9da.
//
// Solidity: function claimRewards(bytes _rewardsRequest, bytes _signers, bytes[] _sigs) returns()
func (_FarmingRewards *FarmingRewardsTransactorSession) ClaimRewards(_rewardsRequest []byte, _signers []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _FarmingRewards.Contract.ClaimRewards(&_FarmingRewards.TransactOpts, _rewardsRequest, _signers, _sigs)
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
	Bin: "0x60e060405234801561001057600080fd5b50604051620014e2380380620014e283398101604081905261003191610058565b6001600160601b0319606093841b811660805291831b821660a05290911b1660c0526100bd565b60008060006060848603121561006d57600080fd5b8351610078816100a5565b6020850151909350610089816100a5565b604085015190925061009a816100a5565b809150509250925092565b6001600160a01b03811681146100ba57600080fd5b50565b60805160601c60a05160601c60c05160601c6113ae62000134600039600081816101ba01526105760152600081816101f401528181610553015281816108ae0152610b1901526000818160e40152818161024b015281816105a901528181610831015281816109880152610a4401526113ae6000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806382d7b4b81161007157806382d7b4b8146101a45780638338f0e5146101ac578063913e77ad146101b5578063934a18ec146101dc578063c6c21e9d146101ef578063e478ed9d1461021657600080fd5b806322da7927146100ae57806325ed6b35146100ca5780634cf088d9146100df578063581c53c51461011e5780637e5fb8f31461013e575b600080fd5b6100b760015481565b6040519081526020015b60405180910390f35b6100dd6100d8366004611041565b610229565b005b6101067f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100c1565b61013161012c366004611011565b6104cd565b6040516100c19190611121565b61019261014c366004610fdf565b6000602081905290815260409020805460018201546002830154600384015460048501546005909501546001600160a01b03909416949293919260ff9182169290911686565b6040516100c1969594939291906110a6565b6100dd6104fb565b6100b760025481565b6101067f000000000000000000000000000000000000000000000000000000000000000081565b6100dd6101ea366004610fdf565b6105a2565b6101067f000000000000000000000000000000000000000000000000000000000000000081565b6100dd610224366004610f92565b61094c565b33600360405163a310624f60e01b81526001600160a01b0383811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a310624f9060240160206040518083038186803b15801561028d57600080fd5b505afa1580156102a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c59190610fc2565b60038111156102d6576102d6611311565b146103285760405162461bcd60e51b815260206004820152601f60248201527f566f746572206973206e6f74206120626f6e6465642076616c696461746f720060448201526064015b60405180910390fd5b60008381526020819052604090206001600582015460ff16600281111561035157610351611311565b146103985760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b604482015260640161031f565b806002015443106103e25760405162461bcd60e51b8152602060048201526014602482015273159bdd1948191958591b1a5b99481c185cdcd95960621b604482015260640161031f565b6001600160a01b038216600090815260068201602052604081205460ff16600381111561041157610411611311565b146104505760405162461bcd60e51b815260206004820152600f60248201526e159bdd195c881a185cc81d9bdd1959608a1b604482015260640161031f565b6001600160a01b03821660009081526006820160205260409020805484919060ff1916600183600381111561048757610487611311565b02179055507f06c7ef6e19454637e93ee60cc680c61fb2ebabb57e58cf36d94141a5036b3d658483856040516104bf93929190611162565b60405180910390a150505050565b6000828152602081815260408083206001600160a01b038516845260060190915290205460ff165b92915050565b6000600254116105425760405162461bcd60e51b8152602060048201526012602482015271139bdd1a1a5b99c81d1bc818dbdb1b1958dd60721b604482015260640161031f565b60025461059b906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016907f000000000000000000000000000000000000000000000000000000000000000090610ba2565b6000600255565b60008060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634021d4d56040518163ffffffff1660e01b815260040160006040518083038186803b15801561060057600080fd5b505afa158015610614573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261063c9190810190610e96565b905060005b81518163ffffffff16101561070b57600161067f86848463ffffffff168151811061066e5761066e611327565b6020026020010151600001516104cd565b600381111561069057610690611311565b14156106c857818163ffffffff16815181106106ae576106ae611327565b602002602001015160200151846106c5919061123b565b93505b818163ffffffff16815181106106e0576106e0611327565b602002602001015160200151836106f7919061123b565b925080610703816112d7565b915050610641565b506000600361071b846002611275565b6107259190611253565b61073090600161123b565b60008681526020819052604090209085101591506001600582015460ff16600281111561075f5761075f611311565b146107a65760405162461bcd60e51b8152602060048201526017602482015276496e76616c69642070726f706f73616c2073746174757360481b604482015260640161031f565b80600201544310156107fa5760405162461bcd60e51b815260206004820152601960248201527f566f746520646561646c696e65206e6f74207265616368656400000000000000604482015260640161031f565b60058101805460ff1916600217905581156108df57600381015460048083015460405163e909156d60e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169363e909156d936108689360ff9092169201611106565b600060405180830381600087803b15801561088257600080fd5b505af1158015610896573d6000803e3d6000fd5b5050825460018401546108da93506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116935090911690610ba2565b6108fb565b8060010154600260008282546108f5919061123b565b90915550505b600381015460048201546040517fd0d659ab2c0f954d2f29cf2e13d8ff2e15e147f3424eb205a079c4caa6bfe1a99261093c928a92879260ff1691906111b6565b60405180910390a1505050505050565b60018054600081815260208190526040902091610969919061123b565b600155604051631042b80b60e21b815233906000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063410ae02c906109bd9084906004016110f8565b60206040518083038186803b1580156109d557600080fd5b505afa1580156109e9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0d9190610ff8565b83546001600160a01b0319166001600160a01b038481169190911785556001808601839055604051631042b80b60e21b81529293507f00000000000000000000000000000000000000000000000000000000000000009091169163410ae02c91610a79916004016110f8565b60206040518083038186803b158015610a9157600080fd5b505afa158015610aa5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac99190610ff8565b610ad3904361123b565b600284015560038301805486919060ff19166001836008811115610af957610af9611311565b02179055506004830184905560058301805460ff19166001179055610b497f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316833084610c0a565b7f4a4d354dbdc4d7b757c1f44b6e074bb6e1afe33f4b9867ce48cfb7004d76f16060018054610b789190611294565b838386600201548989604051610b9396959493929190611185565b60405180910390a15050505050565b6040516001600160a01b038316602482015260448101829052610c0590849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610c48565b505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610c429085906323b872dd60e01b90608401610bce565b50505050565b6000610c9d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610d1a9092919063ffffffff16565b805190915015610c055780806020019051810190610cbb9190610f70565b610c055760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161031f565b6060610d298484600085610d33565b90505b9392505050565b606082471015610d945760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161031f565b843b610de25760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161031f565b600080866001600160a01b03168587604051610dfe919061108a565b60006040518083038185875af1925050503d8060008114610e3b576040519150601f19603f3d011682016040523d82523d6000602084013e610e40565b606091505b5091509150610e50828286610e5d565b925050505b949350505050565b60608315610e6c575081610d2c565b825115610e7c5782518084602001fd5b8160405162461bcd60e51b815260040161031f919061112f565b60006020808385031215610ea957600080fd5b825167ffffffffffffffff80821115610ec157600080fd5b818501915085601f830112610ed557600080fd5b815181811115610ee757610ee761133d565b610ef5848260051b0161120a565b8181528481019250838501600683901b85018601891015610f1557600080fd5b60009450845b83811015610f6257604080838c031215610f33578687fd5b610f3b6111e1565b8351610f4681611353565b8152838901518982015286529487019490910190600101610f1b565b509098975050505050505050565b600060208284031215610f8257600080fd5b81518015158114610d2c57600080fd5b60008060408385031215610fa557600080fd5b823560098110610fb457600080fd5b946020939093013593505050565b600060208284031215610fd457600080fd5b8151610d2c8161136b565b600060208284031215610ff157600080fd5b5035919050565b60006020828403121561100a57600080fd5b5051919050565b6000806040838503121561102457600080fd5b82359150602083013561103681611353565b809150509250929050565b6000806040838503121561105457600080fd5b8235915060208301356110368161136b565b6009811061107657611076611311565b9052565b6004811061107657611076611311565b6000825161109c8184602087016112ab565b9190910192915050565b6001600160a01b0387168152602081018690526040810185905260c081016110d16060830186611066565b836080830152600383106110e7576110e7611311565b8260a0830152979650505050505050565b602081016104f58284611066565b604081016111148285611066565b8260208301529392505050565b602081016104f5828461107a565b602081526000825180602084015261114e8160408501602087016112ab565b601f01601f19169190910160400192915050565b8381526001600160a01b038316602082015260608101610e55604083018461107a565b8681526001600160a01b0386166020820152604081018590526060810184905260c081016110e76080830185611066565b8481528315156020820152608081016111d26040830185611066565b82606083015295945050505050565b6040805190810167ffffffffffffffff811182821017156112045761120461133d565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156112335761123361133d565b604052919050565b6000821982111561124e5761124e6112fb565b500190565b60008261127057634e487b7160e01b600052601260045260246000fd5b500490565b600081600019048311821515161561128f5761128f6112fb565b500290565b6000828210156112a6576112a66112fb565b500390565b60005b838110156112c65781810151838201526020016112ae565b83811115610c425750506000910152565b600063ffffffff808316818114156112f1576112f16112fb565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461136857600080fd5b50565b6004811061136857600080fdfea26469706673582212203048f09e3c2d085f883a690ee7588b49019795c9c21cd18d3b3dfbe6ef2cca2664736f6c63430008070033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signers\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes _signers, bytes[] _sigs) view returns()
func (_ISigsVerifier *ISigsVerifierCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _signers []byte, _sigs [][]byte) error {
	var out []interface{}
	err := _ISigsVerifier.contract.Call(opts, &out, "verifySigs", _msg, _signers, _sigs)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes _signers, bytes[] _sigs) view returns()
func (_ISigsVerifier *ISigsVerifierSession) VerifySigs(_msg []byte, _signers []byte, _sigs [][]byte) error {
	return _ISigsVerifier.Contract.VerifySigs(&_ISigsVerifier.CallOpts, _msg, _signers, _sigs)
}

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes _signers, bytes[] _sigs) view returns()
func (_ISigsVerifier *ISigsVerifierCallerSession) VerifySigs(_msg []byte, _signers []byte, _sigs [][]byte) error {
	return _ISigsVerifier.Contract.VerifySigs(&_ISigsVerifier.CallOpts, _msg, _signers, _sigs)
}

// PbMetaData contains all meta data concerning the Pb contract.
var PbMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f28e6f2aa01bb7e396f4dd76e55cf5279e7c3564e43d35de33741a643e8c91aa64736f6c63430008070033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209a09b605b73c3f5822bb5660ccbf6a217c9a92ca65320659bfef742ef4a7902e64736f6c63430008070033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d2da8b89abfda8e91392170e0f63fc72528abd3a0e9d4093bf5dbcecb7b04ca164736f6c63430008070033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ebb1851f01ea5e49ca6ce246c2b70f96fab0c9316fcb36c52296332833fe037064736f6c63430008070033",
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
	Bin: "0x60a060405234801561001057600080fd5b5060405162001ad438038062001ad4833981016040819052610031916100ac565b61003a3361005c565b6000805460ff60a01b1916905560601b6001600160601b0319166080526100dc565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100be57600080fd5b81516001600160a01b03811681146100d557600080fd5b9392505050565b60805160601c6119b66200011e600039600081816101110152818161056e0152818161060b015281816106b20152818161086a015261096201526119b66000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063c429fe1f11610066578063c429fe1f146101ed578063d0bb93511461020d578063d88ef27114610220578063f2fde38b1461023357600080fd5b80638da5cb5b146101b65780639d4323be146101c7578063b02c43d0146101da57600080fd5b80635c975abb116100c85780635c975abb14610150578063715018a61461016d578063795c2c14146101755780638456cb59146101ae57600080fd5b80633f4ba83a146100ef57806347e7ef24146100f95780634cf088d91461010c575b600080fd5b6100f7610246565b005b6100f76101073660046114c5565b610283565b6101337f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b600054600160a01b900460ff166040519015158152602001610147565b6100f7610398565b6101a061018336600461148c565b600260209081526000928352604080842090915290825290205481565b604051908152602001610147565b6100f76103cc565b6000546001600160a01b0316610133565b6100f76101d53660046114c5565b6103fe565b6101a06101e8366004611610565b610490565b6102006101fb366004611452565b6104b1565b6040516101479190611795565b6100f761021b366004611513565b61054b565b6100f761022e366004611555565b610921565b6100f7610241366004611452565b610b64565b6000546001600160a01b031633146102795760405162461bcd60e51b815260040161027090611802565b60405180910390fd5b610281610bff565b565b600054600160a01b900460ff16156102ad5760405162461bcd60e51b8152600401610270906117d8565b6040516bffffffffffffffffffffffff1933606081811b8316602085015285901b9091166034830152604882018390529060019060680160408051601f198184030181529190528051602091820120825460018101845560009384529190922001556103246001600160a01b038416823085610c9c565b6001805460009161033491611890565b6040805167ffffffffffffffff831681526001600160a01b0385811660208301528716818301526060810186905290519192507f2c0f148b435140de488c1b34647f1511c646f7077e87007bacf22ef9977a16d8919081900360800190a150505050565b6000546001600160a01b031633146103c25760405162461bcd60e51b815260040161027090611802565b6102816000610d0d565b6000546001600160a01b031633146103f65760405162461bcd60e51b815260040161027090611802565b610281610d5d565b600054600160a01b900460ff1661044e5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610270565b6000546001600160a01b031633146104785760405162461bcd60e51b815260040161027090611802565b61048c6001600160a01b0383163383610dc2565b5050565b600181815481106104a057600080fd5b600091825260209091200154905081565b600360205260009081526040902080546104ca906118d3565b80601f01602080910402602001604051908101604052809291908181526020018280546104f6906118d3565b80156105435780601f1061051857610100808354040283529160200191610543565b820191906000526020600020905b81548152906001019060200180831161052657829003601f168201915b505050505081565b604051636d30878360e01b81523360048201819052906000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690636d3087839060240160206040518083038186803b1580156105b057600080fd5b505afa1580156105c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e8919061146f565b6001600160a01b03161461069057604051636d30878360e01b81523360048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690636d3087839060240160206040518083038186803b15801561065557600080fd5b505afa158015610669573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061068d919061146f565b90505b60405163a310624f60e01b81526001600160a01b0382811660048301526000917f00000000000000000000000000000000000000000000000000000000000000009091169063a310624f9060240160206040518083038186803b1580156106f657600080fd5b505afa15801561070a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072e91906115ef565b905060018160038111156107445761074461193f565b1461078a5760405162461bcd60e51b81526020600482015260166024820152752737ba103ab73137b73232b2103b30b634b230ba37b960511b6044820152606401610270565b6001600160a01b038216600090815260036020526040812080546107ad906118d3565b80601f01602080910402602001604051908101604052809291908181526020018280546107d9906118d3565b80156108265780601f106107fb57610100808354040283529160200191610826565b820191906000526020600020905b81548152906001019060200180831161080957829003601f168201915b505050506001600160a01b03851660009081526003602052604090209192506108529190508686611370565b506040516309146f1160e41b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639146f110906108a39086908990899060040161169a565b600060405180830381600087803b1580156108bd57600080fd5b505af11580156108d1573d6000803e3d6000fd5b50505050826001600160a01b03167f8ec5397226cce05bb5f1189621dc680f015802f7f73f89be1a9e89b6af41dcb4828787604051610912939291906117a8565b60405180910390a25050505050565b600054600160a01b900460ff161561094b5760405162461bcd60e51b8152600401610270906117d8565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe9061099d9087908790879087906004016116e4565b60206040518083038186803b1580156109b557600080fd5b505afa1580156109c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ed91906114f1565b506000610a2f85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610df792505050565b80516001600160a01b039081166000908152600260209081526040808320828601519094168352929052818120549183015192935091610a6f9190611890565b905060008111610ac15760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f207769746864726177000000000000006044820152606401610270565b60408083015183516001600160a01b0390811660009081526002602090815284822081880180518516845291529390209190915583519151610b069291169083610dc2565b8151602080840151604080516001600160a01b039485168152939091169183019190915281018290527f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb9060600160405180910390a1505050505050565b6000546001600160a01b03163314610b8e5760405162461bcd60e51b815260040161027090611802565b6001600160a01b038116610bf35760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610270565b610bfc81610d0d565b50565b600054600160a01b900460ff16610c4f5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610270565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040516001600160a01b0380851660248301528316604482015260648101829052610d079085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610ed2565b50505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff1615610d875760405162461bcd60e51b8152600401610270906117d8565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610c7f3390565b6040516001600160a01b038316602482015260448101829052610df290849063a9059cbb60e01b90606401610cd0565b505050565b604080516060810182526000808252602080830182905282840182905283518085019094528184528301849052909190805b60208301515183511015610eca57610e4083610fa4565b90925090508160011415610e6f57610e5f610e5a84610fde565b61109b565b6001600160a01b03168452610e29565b8160021415610e9757610e84610e5a84610fde565b6001600160a01b03166020850152610e29565b8160031415610ebb57610eb1610eac84610fde565b6110ac565b6040850152610e29565b610ec583826110e3565b610e29565b505050919050565b6000610f27826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111559092919063ffffffff16565b805190915015610df25780806020019051810190610f4591906114f1565b610df25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610270565b6000806000610fb28461116e565b9050610fbf60088261184f565b9250806007166005811115610fd657610fd661193f565b915050915091565b60606000610feb8361116e565b90506000818460000151610fff9190611837565b905083602001515181111561101357600080fd5b8167ffffffffffffffff81111561102c5761102c611955565b6040519080825280601f01601f191660200182016040528015611056576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015611090578181015183820152611089602082611837565b905061106e565b505050935250919050565b60006110a6826111f0565b92915050565b60006020825111156110bd57600080fd5b60208201519050815160206110d29190611890565b6110dd906008611871565b1c919050565b60008160058111156110f7576110f761193f565b141561110657610df28261116e565b600281600581111561111a5761111a61193f565b14156100ea57600061112b8361116e565b9050808360000181815161113f9190611837565b90525060208301515183511115610df257600080fd5b6060611164848460008561120f565b90505b9392505050565b602080820151825181019091015160009182805b600a8110156111ea5783811a915061119b816007611871565b82607f16901b8517945081608016600014156111d8576111bc816001611837565b865187906111cb908390611837565b9052509395945050505050565b806111e28161190e565b915050611182565b50600080fd5b6000815160141461120057600080fd5b5060200151600160601b900490565b6060824710156112705760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610270565b843b6112be5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610270565b600080866001600160a01b031685876040516112da919061167e565b60006040518083038185875af1925050503d8060008114611317576040519150601f19603f3d011682016040523d82523d6000602084013e61131c565b606091505b509150915061132c828286611337565b979650505050505050565b60608315611346575081611167565b8251156113565782518084602001fd5b8160405162461bcd60e51b81526004016102709190611795565b82805461137c906118d3565b90600052602060002090601f01602090048101928261139e57600085556113e4565b82601f106113b75782800160ff198235161785556113e4565b828001600101855582156113e4579182015b828111156113e45782358255916020019190600101906113c9565b506113f09291506113f4565b5090565b5b808211156113f057600081556001016113f5565b60008083601f84011261141b57600080fd5b50813567ffffffffffffffff81111561143357600080fd5b60208301915083602082850101111561144b57600080fd5b9250929050565b60006020828403121561146457600080fd5b81356111678161196b565b60006020828403121561148157600080fd5b81516111678161196b565b6000806040838503121561149f57600080fd5b82356114aa8161196b565b915060208301356114ba8161196b565b809150509250929050565b600080604083850312156114d857600080fd5b82356114e38161196b565b946020939093013593505050565b60006020828403121561150357600080fd5b8151801515811461116757600080fd5b6000806020838503121561152657600080fd5b823567ffffffffffffffff81111561153d57600080fd5b61154985828601611409565b90969095509350505050565b6000806000806040858703121561156b57600080fd5b843567ffffffffffffffff8082111561158357600080fd5b61158f88838901611409565b909650945060208701359150808211156115a857600080fd5b818701915087601f8301126115bc57600080fd5b8135818111156115cb57600080fd5b8860208260051b85010111156115e057600080fd5b95989497505060200194505050565b60006020828403121561160157600080fd5b81516004811061116757600080fd5b60006020828403121561162257600080fd5b5035919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6000815180845261166a8160208601602086016118a7565b601f01601f19169290920160200192915050565b600082516116908184602087016118a7565b9190910192915050565b60018060a01b038416815260606020820152600860608201526739b3b716b0b2323960c11b608082015260a0604082015260006116db60a083018486611629565b95945050505050565b6040815260006116f8604083018688611629565b602083820381850152818583528183019050818660051b8401018760005b8881101561178557858303601f190184528135368b9003601e1901811261173c57600080fd5b8a01803567ffffffffffffffff81111561175557600080fd5b8036038c131561176457600080fd5b6117718582898501611629565b958701959450505090840190600101611716565b50909a9950505050505050505050565b6020815260006111676020830184611652565b6040815260006117bb6040830186611652565b82810360208401526117ce818587611629565b9695505050505050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000821982111561184a5761184a611929565b500190565b60008261186c57634e487b7160e01b600052601260045260246000fd5b500490565b600081600019048311821515161561188b5761188b611929565b500290565b6000828210156118a2576118a2611929565b500390565b60005b838110156118c25781810151838201526020016118aa565b83811115610d075750506000910152565b600181811c908216806118e757607f821691505b6020821081141561190857634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561192257611922611929565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610bfc57600080fdfea2646970667358221220af47d94e98bbd95ee972b0a7865d0575fecc3e2e9a2923e366665282c662ce0664736f6c63430008070033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proposalDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxBondedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minValidatorTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_advanceNoticePeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorBondInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valTokens\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"tokenDiff\",\"type\":\"int256\"}],\"name\":\"DelegationUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashAmt\",\"type\":\"uint256\"}],\"name\":\"Slash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashAmtCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"ValidatorNotice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"WhitelistedRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CELER_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bondedValAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectForfeiture\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"completeUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"confirmUnbondedValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forfeiture\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondedValidatorsTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ValidatorTokens[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structDataTypes.DelegatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"}],\"name\":\"getParamValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getQuorumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_checkSelfDelegation\",\"type\":\"bool\"}],\"name\":\"hasMinRequiredTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_commissionRate\",\"type\":\"uint64\"}],\"name\":\"initializeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBondedValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBondBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setGovContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSlashFactor\",\"type\":\"uint256\"}],\"name\":\"setMaxSlashFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataTypes.ParamName\",\"name\":\"_name\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setParamValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setRewardContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_whitelistEnabled\",\"type\":\"bool\"}],\"name\":\"setWhitelistEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerVals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_slashRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newRate\",\"type\":\"uint64\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSelfDelegation\",\"type\":\"uint256\"}],\"name\":\"updateMinSelfDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"updateValidatorSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"valAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"validatorNotice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"bondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"unbondBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162005579380380620055798339810160408190526200003491620001fd565b6200003f33620001ad565b6000805460ff60a01b1916815560609a909a1b6001600160601b031916608052600a6020527f13da86008ba1c6922daee3e07db95305ef49ebced9f5467a0b8613fcc6b343e3989098557fbbc70db1b6c7afd11e79c0fb0051300458f1a3acb8ee9789d9b6b26c61ad9bc7969096557fbff4442b8ed600beeb8e26b1279a0f0d14c6edfaec26d968ee13c86f7d4c2ba8949094557fa856840544dc26124927add067d799967eac11be13e14d82cc281ea46fa39759929092557fe1eb2b2161a492c07c5a334e48012567cba93ec021043f53c1955516a3c5a841557ff35035bc2b01d44bd35a1dcdc552315cffb73da35cfd60570b7b777f98036f9f557f10d9dd018e4cae503383c9f804c1c1603ada5856ee7894375d9b97cd8c8b27db557f22e39f61d1e4986b4f116cea9067f62cc77d74dff1780ae9c8b5166d1dd288295560089091527f2c1fd36ba11b13b555f58753742999069764391f450ca8727fe8a3eeffe67775556200028a565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000806000806000806000806101408b8d0312156200021e57600080fd5b8a516001600160a01b03811681146200023657600080fd5b809a505060208b0151985060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935060e08b015192506101008b015191506101208b015190509295989b9194979a5092959850565b60805160601c61529f620002da6000396000818161087701528181610c6c01528181610ded0152818161175b015281816117d601528181611c8f0152818161202a0152612e45015261529f6000f3fe60806040526004361061031e5760003560e01c80636b68da2e116101ab57806390e360f8116100f7578063b4f7fa3411610095578063eb505dd51161006f578063eb505dd5146109a8578063eecefef8146109d5578063f2fde38b14610a02578063fa52c7d814610a2257600080fd5b8063b4f7fa341461092f578063c8f9f9841461094f578063e909156d1461098857600080fd5b8063960dc08a116100d1578063960dc08a146108655780639b19251a14610899578063a310624f146108c9578063acc62ccf1461090f57600080fd5b806390e360f8146107f55780639146f1101461082557806392bb243c1461084557600080fd5b806382d7b4b8116101645780638456cb591161013e5780638456cb591461078d57806389f9aab5146107a25780638a74d5fe146107b75780638da5cb5b146107d757600080fd5b806382d7b4b81461074c5780638338f0e51461076157806383cfb3181461077757600080fd5b80636b68da2e146106815780636d308783146106a15780636ea69d62146106d7578063715018a6146106f757806371bc02161461070c5780637a50dbd21461072c57600080fd5b80634021d4d51161026a57806351508f0a116102235780635c975abb116101fd5780635c975abb1461060c5780635e593eff1461062b57806365d5d4201461064b57806368706e541461066157600080fd5b806351508f0a146105b257806351fb012d146105d2578063525eba21146105ec57600080fd5b80634021d4d5146104f0578063410ae02c14610512578063473849bd1461053257806347abfdbf1461055257806349955e39146105725780634d99dd161461059257600080fd5b8063291d9549116102d7578063386c024a116102b1578063386c024a1461045d5780633985c4e6146104725780633af32abf146104925780633f4ba83a146104db57600080fd5b8063291d9549146103f05780632fa4d12b1461041057806336f1635f1461044857600080fd5b8063026e402b1461032a578063052d9e7e1461034c57806310154bad1461036c578063145aa1161461038c5780631a203257146103ac5780631cfe4f0b146103cc57600080fd5b3661032557005b600080fd5b34801561033657600080fd5b5061034a610345366004614a29565b610ac3565b005b34801561035857600080fd5b5061034a610367366004614a8f565b610cf5565b34801561037857600080fd5b5061034a610387366004614924565b610d33565b34801561039857600080fd5b5061034a6103a7366004614c7b565b610d66565b3480156103b857600080fd5b5061034a6103c7366004614c7b565b610e14565b3480156103d857600080fd5b506005545b6040519081526020015b60405180910390f35b3480156103fc57600080fd5b5061034a61040b366004614924565b610e6c565b34801561041c57600080fd5b50600b54610430906001600160a01b031681565b6040516001600160a01b0390911681526020016103e7565b34801561045457600080fd5b5061034a610e9f565b34801561046957600080fd5b506103dd611204565b34801561047e57600080fd5b5061034a61048d366004614ac9565b611230565b34801561049e57600080fd5b506104cb6104ad366004614924565b6001600160a01b031660009081526001602052604090205460ff1690565b60405190151581526020016103e7565b3480156104e757600080fd5b5061034a6118db565b3480156104fc57600080fd5b5061050561190f565b6040516103e79190614d42565b34801561051e57600080fd5b506103dd61052d366004614c44565b611a0f565b34801561053e57600080fd5b5061034a61054d366004614924565b611a4e565b34801561055e57600080fd5b506104cb61056d366004614972565b611d0e565b34801561057e57600080fd5b5061034a61058d366004614c94565b611dc8565b34801561059e57600080fd5b5061034a6105ad366004614a29565b611ed9565b3480156105be57600080fd5b5061034a6105cd366004614924565b612291565b3480156105de57600080fd5b506002546104cb9060ff1681565b3480156105f857600080fd5b5061034a610607366004614a53565b612336565b34801561061857600080fd5b50600054600160a01b900460ff166104cb565b34801561063757600080fd5b5061034a610646366004614c7b565b61277b565b34801561065757600080fd5b506103dd60035481565b34801561066d57600080fd5b5061034a61067c366004614924565b612947565b34801561068d57600080fd5b5061034a61069c366004614b34565b6129ec565b3480156106ad57600080fd5b506104306106bc366004614924565b6008602052600090815260409020546001600160a01b031681565b3480156106e357600080fd5b50600c54610430906001600160a01b031681565b34801561070357600080fd5b5061034a612a7e565b34801561071857600080fd5b5061034a610727366004614924565b612ab2565b34801561073857600080fd5b5061034a610747366004614924565b612bef565b34801561075857600080fd5b5061034a612dea565b34801561076d57600080fd5b506103dd600d5481565b34801561078357600080fd5b506103dd60045481565b34801561079957600080fd5b5061034a612e76565b3480156107ae57600080fd5b506006546103dd565b3480156107c357600080fd5b506104cb6107d2366004614bcd565b612ea8565b3480156107e357600080fd5b506000546001600160a01b0316610430565b34801561080157600080fd5b506104cb610810366004614c7b565b60096020526000908152604090205460ff1681565b34801561083157600080fd5b5061034a6108403660046149a9565b613081565b34801561085157600080fd5b50610430610860366004614c7b565b61310d565b34801561087157600080fd5b506104307f000000000000000000000000000000000000000000000000000000000000000081565b3480156108a557600080fd5b506104cb6108b4366004614924565b60016020526000908152604090205460ff1681565b3480156108d557600080fd5b506109026108e4366004614924565b6001600160a01b031660009081526007602052604090205460ff1690565b6040516103e79190614d9a565b34801561091b57600080fd5b5061043061092a366004614c7b565b613137565b34801561093b57600080fd5b506104cb61094a366004614924565b613147565b34801561095b57600080fd5b506103dd61096a366004614924565b6001600160a01b031660009081526007602052604090206001015490565b34801561099457600080fd5b5061034a6109a3366004614c5f565b61317f565b3480156109b457600080fd5b506103dd6109c3366004614c44565b600a6020526000908152604090205481565b3480156109e157600080fd5b506109f56109f036600461493f565b613216565b6040516103e79190614fe4565b348015610a0e57600080fd5b5061034a610a1d366004614924565b613422565b348015610a2e57600080fd5b50610aad610a3d366004614924565b6007602081905260009182526040909120805460018201546002830154600384015460048501546006860154959096015460ff8516966101009095046001600160a01b03169593949293919291906001600160401b0380821691600160401b8104821691600160801b909104168a565b6040516103e79a99989796959493929190614da8565b600054600160a01b900460ff1615610af65760405162461bcd60e51b8152600401610aed90614ebc565b60405180910390fd5b33670de0b6b3a7640000821015610b4f5760405162461bcd60e51b815260206004820152601860248201527f4d696e696d616c20616d6f756e7420697320312043454c5200000000000000006044820152606401610aed565b6001600160a01b038316600090815260076020526040812090815460ff166003811115610b7e57610b7e6151e3565b1415610b9c5760405162461bcd60e51b8152600401610aed90614f79565b6000610bb184836001015484600201546134ba565b6001600160a01b0384166000908152600584016020526040812080549293509183918391610be09084906150a3565b9250508190555081836002016000828254610bfb91906150a3565b9250508190555084836001016000828254610c1691906150a3565b9091555060039050835460ff166003811115610c3457610c346151e3565b1415610c5f578460036000828254610c4c91906150a3565b90915550506001830154610c5f906134e7565b610c946001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168530886135e8565b6001830154815460408051928352602083019190915281018690526001600160a01b0380861691908816907f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea906060015b60405180910390a3505050505050565b6000546001600160a01b03163314610d1f5760405162461bcd60e51b8152600401610aed90614f44565b6002805460ff191682151517905550565b50565b6000546001600160a01b03163314610d5d5760405162461bcd60e51b8152600401610aed90614f44565b610d3081613659565b600054600160a01b900460ff16610db65760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610aed565b6000546001600160a01b03163314610de05760405162461bcd60e51b8152600401610aed90614f44565b610d306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163383613714565b6000546001600160a01b03163314610e3e5760405162461bcd60e51b8152600401610aed90614f44565b6008600052600a6020527f2c1fd36ba11b13b555f58753742999069764391f450ca8727fe8a3eeffe6777555565b6000546001600160a01b03163314610e965760405162461bcd60e51b8152600401610aed90614f44565b610d3081613744565b336000818152600860205260409020546001600160a01b031615610ed85750336000908152600860205260409020546001600160a01b03165b6001600160a01b03811660009081526007602052604090206001815460ff166003811115610f0857610f086151e3565b1480610f2957506002815460ff166003811115610f2757610f276151e3565b145b610f755760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642076616c696461746f722073746174757300000000000000006044820152606401610aed565b60078101546001600160401b0316431015610fcb5760405162461bcd60e51b8152602060048201526016602482015275109bdb9908189b1bd8dac81b9bdd081c995858da195960521b6044820152606401610aed565b60045443101561101d5760405162461bcd60e51b815260206004820152601b60248201527f546f6f206672657175656e742076616c696461746f7220626f6e6400000000006044820152606401610aed565b6007600052600a6020527f22e39f61d1e4986b4f116cea9067f62cc77d74dff1780ae9c8b5166d1dd288295461105390436150a3565b600455611061826001611d0e565b6110a35760405162461bcd60e51b81526020600482015260136024820152724e6f742068617665206d696e20746f6b656e7360681b6044820152606401610aed565b6003600052600a6020527fa856840544dc26124927add067d799967eac11be13e14d82cc281ea46fa39759546006548111156110e7576110e2836137ef565b505050565b6000196000805b8381101561119c5782600760006006848154811061110e5761110e61520f565b60009182526020808320909101546001600160a01b03168352820192909252604001902060010154101561118a5780915060076000600683815481106111565761115661520f565b60009182526020808320909101546001600160a01b0316835282019290925260400190206001015492508261118a5761119c565b8061119481615171565b9150506110ee565b50818460010154116111e65760405162461bcd60e51b8152602060048201526013602482015272496e73756666696369656e7420746f6b656e7360681b6044820152606401610aed565b6111f08582613843565b6111fd84600101546134e7565b5050505050565b600060038054600261121691906150dd565b61122091906150bb565b61122b9060016150a3565b905090565b600054600160a01b900460ff161561125a5760405162461bcd60e51b8152600401610aed90614ebc565b6112a084848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107d29250859150869050615138565b5060006112e285858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506138be92505050565b905080606001516001600160401b031642106113305760405162461bcd60e51b815260206004820152600d60248201526c14db185cda08195e1c1a5c9959609a1b6044820152606401610aed565b620f424081604001516001600160401b031611156113875760405162461bcd60e51b815260206004820152601460248201527324b73b30b634b21039b630b9b4103330b1ba37b960611b6044820152606401610aed565b6008600052600a6020527f2c1fd36ba11b13b555f58753742999069764391f450ca8727fe8a3eeffe677755460408201516001600160401b0316111561140f5760405162461bcd60e51b815260206004820152601760248201527f457863656564206d617820736c61736820666163746f720000000000000000006044820152606401610aed565b6020808201516001600160401b031660009081526009909152604090205460ff16156114705760405162461bcd60e51b815260206004820152601060248201526f5573656420736c617368206e6f6e636560801b6044820152606401610aed565b6020808201516001600160401b03166000908152600982526040808220805460ff1916600190811790915584516001600160a01b03811684526007909452912090815460ff1660038111156114c7576114c76151e3565b141561150b5760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881d5b989bdd5b991959606a1b6044820152606401610aed565b6000620f424084604001516001600160401b0316836001015461152e91906150dd565b61153891906150bb565b90508082600101600082825461154e91906150fc565b9091555060039050825460ff16600381111561156c5761156c6151e3565b141561160957806003600082825461158491906150fc565b909155505060808401516001600160401b03161515806115ac57506115aa836001611d0e565b155b15611609576115ba83613b2b565b60808401516001600160401b0316156116095760808401516115e5906001600160401b0316436150a3565b60078301805467ffffffffffffffff19166001600160401b03929092169190911790555b60006001600160a01b0316836001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea8460010154600085611650906151b0565b6040805193845260208401929092529082015260600160405180910390a36000620f424085604001516001600160401b0316846003015461169191906150dd565b61169b91906150bb565b9050808360030160008282546116b191906150fc565b909155506116c1905081836150a3565b91506000805b8660a001515181101561185b5760008760a0015182815181106116ec576116ec61520f565b602002602001015190508481602001518461170791906150a3565b111561171d5761171783866150fc565b60208201525b60208101511561184857602081015161173690846150a3565b81519093506001600160a01b03166117c1576020810151611783906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016903390613714565b60208082015160405190815233917fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3910160405180910390a2611848565b805160208201516117fc916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691613714565b80600001516001600160a01b03167fb1375221b23a15d2f6887c7dbdc6745a07d9a5245076d51fb41879590ebbd2a3826020015160405161183f91815260200190565b60405180910390a25b508061185381615171565b9150506116c7565b5061186681846150fc565b600d600082825461187791906150a3565b9091555050602080870151604080516001600160401b0390921682529181018590526001600160a01b038716917f10863f35bc5db9fda133333468bf7b1ceaaa88cb4263c061f890f97b79bf9008910160405180910390a250505050505050505050565b6000546001600160a01b031633146119055760405162461bcd60e51b8152600401610aed90614f44565b61190d613c88565b565b6005546060906000906001600160401b0381111561192f5761192f615225565b60405190808252806020026020018201604052801561197457816020015b604080518082019091526000808252602082015281526020019060019003908161194d5790505b50905060005b600654811015611a09576000600682815481106119995761199961520f565b60009182526020808320909101546040805180820182526001600160a01b0390921680835280855260078452932060010154918101919091528451919250908490849081106119ea576119ea61520f565b6020026020010181905250508080611a0190615171565b91505061197a565b50919050565b6000600a6000836008811115611a2757611a276151e3565b6008811115611a3857611a386151e3565b8152602001908152602001600020549050919050565b6001600160a01b03811660009081526007602052604081203391815460ff166003811115611a7e57611a7e6151e3565b1415611a9c5760405162461bcd60e51b8152600401610aed90614f79565b6001600160a01b03821660009081526005820160209081526040822060028352600a9091527fbff4442b8ed600beeb8e26b1279a0f0d14c6edfaec26d968ee13c86f7d4c2ba85483549192909160019060ff166003811115611b0057611b006151e3565b60028501549114915063ffffffff1660005b600285015463ffffffff64010000000090910481169083161015611bc4578280611b62575063ffffffff821660009081526001808701602052604090912001544390611b5f9086906150a3565b11155b15611bad5763ffffffff82166000908152600186016020526040902054611b8990826150a3565b63ffffffff8316600090815260018088016020526040822082815501559050611bb2565b611bc4565b81611bbc8161518c565b925050611b12565b60028501805463ffffffff191663ffffffff841617905580611c365760405162461bcd60e51b815260206004820152602560248201527f4e6f20756e64656c65676174696f6e20726561647920746f20626520636f6d706044820152641b195d195960da1b6064820152608401610aed565b6000611c4b8288600301548960040154613d25565b905081876004016000828254611c6191906150fc565b9250508190555080876003016000828254611c7c91906150fc565b90915550611cb690506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168983613714565b876001600160a01b0316896001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c83604051611cfb91815260200190565b60405180910390a3505050505050505050565b6001600160a01b03821660009081526007602090815260408220600181015460048452600a9092527fe1eb2b2161a492c07c5a334e48012567cba93ec021043f53c1955516a3c5a84154909190811015611d6d57600092505050611dc2565b8315611dbb576001600160a01b03851660009081526005830160205260408120546002840154611d9f91908490613d25565b90508260060154811015611db95760009350505050611dc2565b505b6001925050505b92915050565b33600081815260076020526040812090815460ff166003811115611dee57611dee6151e3565b1415611e0c5760405162461bcd60e51b8152600401610aed90614f79565b612710836001600160401b03161115611e5a5760405162461bcd60e51b815260206004820152601060248201526f496e76616c6964206e6577207261746560801b6044820152606401610aed565b60078101805467ffffffffffffffff60801b1916600160801b6001600160401b038616908102919091179091556040805160208101929092526001600160a01b0384169160008051602061524a833981519152910160408051601f1981840301815290829052611ecc91600090614fb0565b60405180910390a2505050565b33670de0b6b3a7640000821015611f325760405162461bcd60e51b815260206004820152601960248201527f4d696e696d616c20616d6f756e742069732031207368617265000000000000006044820152606401610aed565b6001600160a01b038316600090815260076020526040812090815460ff166003811115611f6157611f616151e3565b1415611f7f5760405162461bcd60e51b8152600401610aed90614f79565b6000611f948483600101548460020154613d25565b6001600160a01b0384166000908152600584016020526040812080549293509186918391611fc39084906150fc565b9250508190555084836002016000828254611fde91906150fc565b9250508190555081836001016000828254611ff991906150fc565b9091555060019050835460ff166003811115612017576120176151e3565b1415612096576120516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168584613714565b836001600160a01b0316866001600160a01b03167f4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c84604051610ce591815260200190565b6003835460ff1660038111156120ae576120ae6151e3565b14156120f85781600360008282546120c691906150fc565b925050819055506120eb86876001600160a01b0316866001600160a01b031614611d0e565b6120f8576120f886613b2b565b6002810154600a9061211b9063ffffffff80821691640100000000900416615113565b63ffffffff161061216e5760405162461bcd60e51b815260206004820152601f60248201527f457863656564206d617820756e64656c65676174696f6e20656e7472696573006044820152606401610aed565b600061218383856003015486600401546134ba565b90508084600401600082825461219991906150a3565b92505081905550828460030160008282546121b491906150a3565b909155505060028201805463ffffffff640100000000918290048116600090815260018087016020526040909120858155439181019190915583549093929004169060046122018361518c565b91906101000a81548163ffffffff021916908363ffffffff16021790555050856001600160a01b0316886001600160a01b03167f2e289e5a72f8e92e344eb866e0a32621f332835d2df2cf1f76e5a345b23cf1ea8760010154866000015488612269906151b0565b6040805193845260208401929092529082015260600160405180910390a35050505050505050565b6000546001600160a01b031633146122bb5760405162461bcd60e51b8152600401610aed90614f44565b600c546001600160a01b0316156123145760405162461bcd60e51b815260206004820152601b60248201527f72657761726420636f6e747261637420616c72656164792073657400000000006044820152606401610aed565b600c80546001600160a01b0319166001600160a01b0392909216919091179055565b600054600160a01b900460ff16156123605760405162461bcd60e51b8152600401610aed90614ebc565b60025460ff16156123ca573360009081526001602052604090205460ff166123ca5760405162461bcd60e51b815260206004820152601960248201527f43616c6c6572206973206e6f742077686974656c6973746564000000000000006044820152606401610aed565b33600081815260076020526040812090815460ff1660038111156123f0576123f06151e3565b1461243d5760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f7220697320696e697469616c697a656400000000000000006044820152606401610aed565b6001600160a01b03851660009081526007602052604081205460ff16600381111561246a5761246a6151e3565b146124b35760405162461bcd60e51b815260206004820152601960248201527829b4b3b732b91034b99037ba3432b9103b30b634b230ba37b960391b6044820152606401610aed565b6001600160a01b03828116600090815260086020526040902054161561251b5760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206973206f74686572207369676e6572000000000000006044820152606401610aed565b6001600160a01b0385811660009081526008602052604090205416156125795760405162461bcd60e51b815260206004820152601360248201527214da59db995c88185b1c9958591e481d5cd959606a1b6044820152606401610aed565b612710836001600160401b031611156125d45760405162461bcd60e51b815260206004820152601760248201527f496e76616c696420636f6d6d697373696f6e20726174650000000000000000006044820152606401610aed565b6005600052600a6020527ff35035bc2b01d44bd35a1dcdc552315cffb73da35cfd60570b7b777f98036f9f5484101561264f5760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610aed565b80546001600160a81b03191660ff196101006001600160a01b038881169182029290921692909217600190811784556006840187905560078401805467ffffffffffffffff60801b1916600160801b6001600160401b038916021790556005805491820190557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180546001600160a01b0319908116928616928317909155600092835260086020526040909220805490921617905561270f8285610ac3565b604080516001600160a01b0387811660208301529181018690526001600160401b03851660608201529083169060008051602061524a8339815191529060800160408051601f198184030181529082905261276c91600090614f16565b60405180910390a25050505050565b33600081815260076020526040812090815460ff1660038111156127a1576127a16151e3565b14156127bf5760405162461bcd60e51b8152600401610aed90614f79565b6005600052600a6020527ff35035bc2b01d44bd35a1dcdc552315cffb73da35cfd60570b7b777f98036f9f5483101561283a5760405162461bcd60e51b815260206004820181905260248201527f496e73756666696369656e74206d696e2073656c662064656c65676174696f6e6044820152606401610aed565b80600601548310156128fc576003815460ff16600381111561285e5761285e6151e3565b14156128a25760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881a5cc8189bdb991959606a1b6044820152606401610aed565b6006600052600a6020527f10d9dd018e4cae503383c9f804c1c1603ada5856ee7894375d9b97cd8c8b27db546128d890436150a3565b60078201805467ffffffffffffffff19166001600160401b03929092169190911790555b6006810183905560408051602081018590526001600160a01b0384169160008051602061524a833981519152910160408051601f1981840301815290829052611ecc91600090614e68565b6000546001600160a01b031633146129715760405162461bcd60e51b8152600401610aed90614f44565b600b546001600160a01b0316156129ca5760405162461bcd60e51b815260206004820152601860248201527f676f7620636f6e747261637420616c72656164792073657400000000000000006044820152606401610aed565b600b80546001600160a01b0319166001600160a01b0392909216919091179055565b612a3286868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506107d29250859150869050615138565b612a765760405162461bcd60e51b81526020600482015260156024820152744661696c656420746f20766572696679207369677360581b6044820152606401610aed565b505050505050565b6000546001600160a01b03163314612aa85760405162461bcd60e51b8152600401610aed90614f44565b61190d6000613d3e565b6001600160a01b03811660009081526007602052604090206002815460ff166003811115612ae257612ae26151e3565b14612b2f5760405162461bcd60e51b815260206004820152601760248201527f56616c696461746f72206e6f7420756e626f6e64696e670000000000000000006044820152606401610aed565b6007810154600160401b90046001600160401b0316431015612b935760405162461bcd60e51b815260206004820152601860248201527f556e626f6e6420626c6f636b206e6f74207265616368656400000000000000006044820152606401610aed565b805460ff19166001908117825560078201805467ffffffffffffffff60401b191690555b6040516001600160a01b038416907fd5e59fa85493a77fb57f6bf9080f2f71fde9de0eadc62b27b43b6263f3f1f59a90600090a35050565b33600081815260076020526040812090815460ff166003811115612c1557612c156151e3565b1415612c635760405162461bcd60e51b815260206004820152601960248201527f56616c696461746f72206e6f7420696e697469616c697a6564000000000000006044820152606401610aed565b6001600160a01b038381166000908152600860205260409020541615612cc15760405162461bcd60e51b815260206004820152601360248201527214da59db995c88185b1c9958591e481d5cd959606a1b6044820152606401610aed565b816001600160a01b0316836001600160a01b031614612d50576001600160a01b03831660009081526007602052604081205460ff166003811115612d0757612d076151e3565b14612d505760405162461bcd60e51b815260206004820152601960248201527829b4b3b732b91034b99037ba3432b9103b30b634b230ba37b960391b6044820152606401610aed565b8054610100908190046001600160a01b03908116600090815260086020908152604080832080546001600160a01b03199081169091558654610100600160a81b031916898616968702178755858452928190208054909316938716938417909255815190810193909352909160008051602061524a833981519152910160408051601f1981840301815290829052611ecc91600090614ee6565b6000600d5411612e315760405162461bcd60e51b8152602060048201526012602482015271139bdd1a1a5b99c81d1bc818dbdb1b1958dd60721b6044820152606401610aed565b600c54600d54612e6f916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811692911690613714565b6000600d55565b6000546001600160a01b03163314612ea05760405162461bcd60e51b8152600401610aed90614f44565b61190d613d8e565b600080612f0984805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905060008080612f17611204565b905060005b8651811015613043576000612f53888381518110612f3c57612f3c61520f565b602002602001015187613df390919063ffffffff16565b9050836001600160a01b0316816001600160a01b031611612fb65760405162461bcd60e51b815260206004820152601e60248201527f5369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610aed565b6001600160a01b038082166000908152600860209081526040808320549093168252600790522090935083906003815460ff166003811115612ffa57612ffa6151e3565b14613006575050613031565b600181015461301590876150a3565b955083861061302e576001975050505050505050611dc2565b50505b8061303b81615171565b915050612f1c565b5060405162461bcd60e51b8152602060048201526012602482015271145d5bdc9d5b481b9bdd081c995858da195960721b6044820152606401610aed565b6001600160a01b038516600090815260076020526040812090815460ff1660038111156130b0576130b06151e3565b14156130ce5760405162461bcd60e51b8152600401610aed90614f79565b856001600160a01b031660008051602061524a83398151915286868686336040516130fd959493929190614e13565b60405180910390a2505050505050565b6005818154811061311d57600080fd5b6000918252602090912001546001600160a01b0316905081565b6006818154811061311d57600080fd5b600060036001600160a01b03831660009081526007602052604090205460ff166003811115613178576131786151e3565b1492915050565b600b546001600160a01b031633146131d95760405162461bcd60e51b815260206004820152601a60248201527f43616c6c6572206973206e6f7420676f7620636f6e74726163740000000000006044820152606401610aed565b80600a60008460088111156131f0576131f06151e3565b6008811115613201576132016151e3565b81526020810191909152604001600020555050565b6132516040518060a0016040528060006001600160a01b03168152602001600081526020016000815260200160008152602001606081525090565b6001600160a01b0380841660009081526007602090815260408083209386168352600584019091528120805460018401546002850154929392613295929190613d25565b600283015490915060009081906132bd9063ffffffff80821691640100000000900416615113565b63ffffffff1690506000816001600160401b038111156132df576132df615225565b60405190808252806020026020018201604052801561332457816020015b60408051808201909152600080825260208201528152602001906001900390816132fd5790505b50905060005b828110156133cc576002860154600187019060009061334f9063ffffffff16846150a3565b81526020019081526020016000206040518060400160405290816000820154815260200160018201548152505082828151811061338e5761338e61520f565b60209081029190910181019190915260008281526001880190915260409020546133b890856150a3565b9350806133c481615171565b91505061332a565b5060006133e28488600301548960040154613d25565b6040805160a0810182526001600160a01b03909c168c5260208c01969096529554948a019490945260608901949094525050506080850152509192915050565b6000546001600160a01b0316331461344c5760405162461bcd60e51b8152600401610aed90614f44565b6001600160a01b0381166134b15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610aed565b610d3081613d3e565b6000826134c85750826134e0565b826134d383866150dd565b6134dd91906150bb565b90505b9392505050565b60065460028114806134f95750806003145b1561356f57613506611204565b821061356b5760405162461bcd60e51b815260206004820152602e60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526d2071756f72756d20746f6b656e7360901b6064820152608401610aed565b5050565b600381111561356b576003805461358691906150bb565b821061356b5760405162461bcd60e51b815260206004820152602b60248201527f53696e676c652076616c696461746f722073686f756c64206e6f74206861766560448201526a20312f3320746f6b656e7360a81b6064820152608401610aed565b6040516001600160a01b03808516602483015283166044820152606481018290526136539085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613e97565b50505050565b6001600160a01b03811660009081526001602052604090205460ff16156136b85760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481dda1a5d195b1a5cdd1959606a1b6044820152606401610aed565b6001600160a01b038116600081815260016020818152604092839020805460ff191690921790915590519182527fee1504a83b6d4a361f4c1dc78ab59bfa30d6a3b6612c403e86bb01ef2984295f91015b60405180910390a150565b6040516001600160a01b0383166024820152604481018290526110e290849063a9059cbb60e01b9060640161361c565b6001600160a01b03811660009081526001602052604090205460ff1661379e5760405162461bcd60e51b815260206004820152600f60248201526e139bdd081dda1a5d195b1a5cdd1959608a1b6044820152606401610aed565b6001600160a01b038116600081815260016020908152604091829020805460ff1916905590519182527f270d9b30cf5b0793bbfd54c9d5b94aeb49462b8148399000265144a8722da6b69101613709565b600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b0319166001600160a01b038316179055610d3081613f69565b613873600682815481106138595761385961520f565b6000918252602090912001546001600160a01b0316613fc9565b81600682815481106138875761388761520f565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555061356b82613f69565b6040805160c0810182526000808252602080830182905282840182905260608084018390526080840183905260a08401528351808501909452818452830184905290919061390d82600661406e565b9050806006815181106139225761392261520f565b60200260200101516001600160401b0381111561394157613941615225565b60405190808252806020026020018201604052801561398657816020015b604080518082019091526000808252602082015281526020019060019003908161395f5790505b508360a001819052506000816006815181106139a4576139a461520f565b6020026020010181815250506000805b60208401515184511015613b22576139cb84614127565b909250905081600114156139fa576139ea6139e585614161565b61421d565b6001600160a01b031685526139b4565b8160021415613a1f57613a0c84614228565b6001600160401b031660208601526139b4565b8160031415613a4457613a3184614228565b6001600160401b031660408601526139b4565b8160041415613a6957613a5684614228565b6001600160401b031660608601526139b4565b8160051415613a8e57613a7b84614228565b6001600160401b031660808601526139b4565b8160061415613b1357613aa8613aa385614161565b6142aa565b8560a0015184600681518110613ac057613ac061520f565b602002602001015181518110613ad857613ad861520f565b602002602001018190525082600681518110613af657613af661520f565b602002602001018051809190613b0b90615171565b9052506139b4565b613b1d8482614351565b6139b4565b50505050919050565b600654600090613b3d906001906150fc565b905060005b600654811015613c4857826001600160a01b031660068281548110613b6957613b6961520f565b6000918252602090912001546001600160a01b03161415613c365781811015613bfa5760068281548110613b9f57613b9f61520f565b600091825260209091200154600680546001600160a01b039092169183908110613bcb57613bcb61520f565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b6006805480613c0b57613c0b6151f9565b600082815260209020810160001990810180546001600160a01b03191690550190556110e283613fc9565b80613c4081615171565b915050613b42565b5060405162461bcd60e51b81526020600482015260146024820152732737ba103137b73232b2103b30b634b230ba37b960611b6044820152606401610aed565b600054600160a01b900460ff16613cd85760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610aed565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600081613d335750826134e0565b816134d384866150dd565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff1615613db85760405162461bcd60e51b8152600401610aed90614ebc565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613d083390565b6000815160411415613e275760208201516040830151606084015160001a613e1d868285856143c3565b9350505050611dc2565b815160401415613e4f5760208201516040830151613e4685838361456c565b92505050611dc2565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610aed565b6000613eec826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166145969092919063ffffffff16565b8051909150156110e25780806020019051810190613f0a9190614aac565b6110e25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610aed565b6001600160a01b03811660009081526007602081905260408220805460ff191660039081178255918101805467ffffffffffffffff60401b19169055600181015482549193909291613fbc9084906150a3565b9091555060039050612bb7565b6001600160a01b03811660009081526007602090815260408220805460ff191660029081178255909252600a90527fbff4442b8ed600beeb8e26b1279a0f0d14c6edfaec26d968ee13c86f7d4c2ba85461402390436150a3565b8160070160086101000a8154816001600160401b0302191690836001600160401b0316021790555080600101546003600082825461406191906150fc565b9091555060029050612bb7565b815160609061407e8360016150a3565b6001600160401b0381111561409557614095615225565b6040519080825280602002602001820160405280156140be578160200160208202803683370190505b5091506000805b6020860151518651101561411e576140dc86614127565b809250819350505060018483815181106140f8576140f861520f565b6020026020010181815161410c91906150a3565b9052506141198682614351565b6140c5565b50509092525090565b600080600061413584614228565b90506141426008826150bb565b9250806007166005811115614159576141596151e3565b915050915091565b6060600061416e83614228565b9050600081846000015161418291906150a3565b905083602001515181111561419657600080fd5b816001600160401b038111156141ae576141ae615225565b6040519080825280601f01601f1916602001820160405280156141d8576020820181803683370190505b50602080860151865192955091818601919083010160005b8581101561421257818101518382015261420b6020826150a3565b90506141f0565b505050935250919050565b6000611dc2826145a5565b602080820151825181019091015160009182805b600a8110156142a45783811a91506142558160076150dd565b82607f16901b851794508160801660001415614292576142768160016150a3565b865187906142859083906150a3565b9052509395945050505050565b8061429c81615171565b91505061423c565b50600080fd5b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b60208301515183511015614349576142ec83614127565b90925090508160011415614316576143066139e584614161565b6001600160a01b031684526142d5565b816002141561433a5761433061432b84614161565b6145c4565b60208501526142d5565b6143448382614351565b6142d5565b505050919050565b6000816005811115614365576143656151e3565b1415614374576110e282614228565b6002816005811115614388576143886151e3565b141561032557600061439983614228565b905080836000018181516143ad91906150a3565b905250602083015151835111156110e257600080fd5b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156144405760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610aed565b8360ff16601b148061445557508360ff16601c145b6144ac5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610aed565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015614500573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166145635760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610aed565b95945050505050565b60006001600160ff1b03821660ff83901c601b0161458c868287856143c3565b9695505050505050565b60606134dd84846000856145fb565b600081516014146145b557600080fd5b5060200151600160601b900490565b60006020825111156145d557600080fd5b60208201519050815160206145ea91906150fc565b6145f59060086150dd565b1c919050565b60608247101561465c5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610aed565b843b6146aa5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610aed565b600080866001600160a01b031685876040516146c69190614d26565b60006040518083038185875af1925050503d8060008114614703576040519150601f19603f3d011682016040523d82523d6000602084013e614708565b606091505b5091509150614718828286614723565b979650505050505050565b606083156147325750816134e0565b8251156147425782518084602001fd5b8160405162461bcd60e51b8152600401610aed9190614e55565b60006001600160401b038084111561477657614776615225565b8360051b6020614787818301615073565b86815293508084018583810189101561479f57600080fd5b60009350835b888110156147da578135868111156147bb578586fd5b6147c78b828b0161488f565b84525091830191908301906001016147a5565b5050505050509392505050565b80356001600160a01b03811681146147fe57600080fd5b919050565b60008083601f84011261481557600080fd5b5081356001600160401b0381111561482c57600080fd5b6020830191508360208260051b850101111561484757600080fd5b9250929050565b60008083601f84011261486057600080fd5b5081356001600160401b0381111561487757600080fd5b60208301915083602082850101111561484757600080fd5b600082601f8301126148a057600080fd5b81356001600160401b038111156148b9576148b9615225565b6148cc601f8201601f1916602001615073565b8181528460208386010111156148e157600080fd5b816020850160208301376000918101602001919091529392505050565b8035600981106147fe57600080fd5b80356001600160401b03811681146147fe57600080fd5b60006020828403121561493657600080fd5b6134e0826147e7565b6000806040838503121561495257600080fd5b61495b836147e7565b9150614969602084016147e7565b90509250929050565b6000806040838503121561498557600080fd5b61498e836147e7565b9150602083013561499e8161523b565b809150509250929050565b6000806000806000606086880312156149c157600080fd5b6149ca866147e7565b945060208601356001600160401b03808211156149e657600080fd5b6149f289838a0161484e565b90965094506040880135915080821115614a0b57600080fd5b50614a188882890161484e565b969995985093965092949392505050565b60008060408385031215614a3c57600080fd5b614a45836147e7565b946020939093013593505050565b600080600060608486031215614a6857600080fd5b614a71846147e7565b925060208401359150614a866040850161490d565b90509250925092565b600060208284031215614aa157600080fd5b81356134e08161523b565b600060208284031215614abe57600080fd5b81516134e08161523b565b60008060008060408587031215614adf57600080fd5b84356001600160401b0380821115614af657600080fd5b614b028883890161484e565b90965094506020870135915080821115614b1b57600080fd5b50614b2887828801614803565b95989497509550505050565b60008060008060008060608789031215614b4d57600080fd5b86356001600160401b0380821115614b6457600080fd5b614b708a838b0161484e565b90985096506020890135915080821115614b8957600080fd5b614b958a838b0161484e565b90965094506040890135915080821115614bae57600080fd5b50614bbb89828a01614803565b979a9699509497509295939492505050565b60008060408385031215614be057600080fd5b82356001600160401b0380821115614bf757600080fd5b614c038683870161488f565b93506020850135915080821115614c1957600080fd5b508301601f81018513614c2b57600080fd5b614c3a8582356020840161475c565b9150509250929050565b600060208284031215614c5657600080fd5b6134e0826148fe565b60008060408385031215614c7257600080fd5b614a45836148fe565b600060208284031215614c8d57600080fd5b5035919050565b600060208284031215614ca657600080fd5b6134e08261490d565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008151808452614cf0816020860160208601615145565b601f01601f19169290920160200192915050565b60048110614d2257634e487b7160e01b600052602160045260246000fd5b9052565b60008251614d38818460208701615145565b9190910192915050565b602080825282518282018190526000919060409081850190868401855b82811015614d8d57815180516001600160a01b03168552860151868501529284019290850190600101614d5f565b5091979650505050505050565b60208101611dc28284614d04565b6101408101614db7828d614d04565b6001600160a01b039a909a16602082015260408101989098526060880196909652608087019490945260a086019290925260c08501526001600160401b0390811660e08501529081166101008401521661012090910152919050565b606081526000614e27606083018789614caf565b8281036020840152614e3a818688614caf565b91505060018060a01b03831660408301529695505050505050565b6020815260006134e06020830184614cd8565b60608152601360608201527236b4b716b9b2b63316b232b632b3b0ba34b7b760691b608082015260a060208201526000614ea560a0830185614cd8565b905060018060a01b03831660408301529392505050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b60608152600660608201526539b4b3b732b960d11b608082015260a060208201526000614ea560a0830185614cd8565b6060815260046060820152631a5b9a5d60e21b608082015260a060208201526000614ea560a0830185614cd8565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601c908201527f56616c696461746f72206973206e6f7420696e697469616c697a656400000000604082015260600190565b60608152600a60608201526931b7b6b6b4b9b9b4b7b760b11b608082015260a060208201526000614ea560a0830185614cd8565b6000602080835260c0830160018060a01b03855116828501528185015160408181870152808701516060870152606087015160808701526080870151915060a08087015282825180855260e0880191508584019450600093505b80841015615067578451805183528601518683015293850193600193909301929082019061503e565b50979650505050505050565b604051601f8201601f191681016001600160401b038111828210171561509b5761509b615225565b604052919050565b600082198211156150b6576150b66151cd565b500190565b6000826150d857634e487b7160e01b600052601260045260246000fd5b500490565b60008160001904831182151516156150f7576150f76151cd565b500290565b60008282101561510e5761510e6151cd565b500390565b600063ffffffff83811690831681811015615130576151306151cd565b039392505050565b60006134e036848461475c565b60005b83811015615160578181015183820152602001615148565b838111156136535750506000910152565b6000600019821415615185576151856151cd565b5060010190565b600063ffffffff808316818114156151a6576151a66151cd565b6001019392505050565b6000600160ff1b8214156151c6576151c66151cd565b5060000390565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b8015158114610d3057600080fdfe3683b59f352bc42833c21c736ba7631d3e35fed49723ebac8298d4e0f36e512ca2646970667358221220254819c8490545041bfce1034339c4e2c0aee5128f90c1c4ea4b44e9ca196a9264736f6c63430008070033",
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
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
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
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
func (_Staking *StakingSession) GetDelegatorInfo(_valAddr common.Address, _delAddr common.Address) (DataTypesDelegatorInfo, error) {
	return _Staking.Contract.GetDelegatorInfo(&_Staking.CallOpts, _valAddr, _delAddr)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0xeecefef8.
//
// Solidity: function getDelegatorInfo(address _valAddr, address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[]))
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

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes , bytes[] _sigs) view returns()
func (_Staking *StakingCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, arg1 []byte, _sigs [][]byte) error {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "verifySigs", _msg, arg1, _sigs)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes , bytes[] _sigs) view returns()
func (_Staking *StakingSession) VerifySigs(_msg []byte, arg1 []byte, _sigs [][]byte) error {
	return _Staking.Contract.VerifySigs(&_Staking.CallOpts, _msg, arg1, _sigs)
}

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes , bytes[] _sigs) view returns()
func (_Staking *StakingCallerSession) VerifySigs(_msg []byte, arg1 []byte, _sigs [][]byte) error {
	return _Staking.Contract.VerifySigs(&_Staking.CallOpts, _msg, arg1, _sigs)
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
	Bin: "0x60a060405234801561001057600080fd5b506040516113e93803806113e983398101604081905261002f916100aa565b6100383361005a565b6000805460ff60a01b1916905560601b6001600160601b0319166080526100da565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bc57600080fd5b81516001600160a01b03811681146100d357600080fd5b9392505050565b60805160601c6112d66101136000396000818160e3015281816101f40152818161035c0152818161057101526106e701526112d66000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063715018a611610071578063715018a61461013f5780638456cb59146101475780638da5cb5b1461014f57806396db0fef14610160578063f2fde38b1461018e578063f8df0dc5146101a157600080fd5b80630a300b09146100ae578063145aa116146100c35780633f4ba83a146100d65780634cf088d9146100de5780635c975abb14610122575b600080fd5b6100c16100bc366004610ff1565b6101b4565b005b6100c16100d1366004610ff1565b6102db565b6100c16103fe565b6101057f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b600054600160a01b900460ff166040519015158152602001610119565b6100c1610432565b6100c1610466565b6000546001600160a01b0316610105565b61018061016e366004610ed0565b60016020526000908152604090205481565b604051908152602001610119565b6100c161019c366004610ed0565b610498565b6100c16101af366004610f0f565b610530565b600054600160a01b900460ff16156101e75760405162461bcd60e51b81526004016101de90611133565b60405180910390fd5b60003390506102948130847f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561024b57600080fd5b505afa15801561025f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102839190610fd4565b6001600160a01b031692919061078e565b806001600160a01b03167ff67017a05194c0853be9169be60cad9fa6e75d34b6b507a7a4261510e19c6d79836040516102cf91815260200190565b60405180910390a25050565b600054600160a01b900460ff1661032b5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016101de565b6000546001600160a01b031633146103555760405162461bcd60e51b81526004016101de9061115d565b6103fb33827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156103b357600080fd5b505afa1580156103c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103eb9190610fd4565b6001600160a01b031691906107ff565b50565b6000546001600160a01b031633146104285760405162461bcd60e51b81526004016101de9061115d565b610430610834565b565b6000546001600160a01b0316331461045c5760405162461bcd60e51b81526004016101de9061115d565b61043060006108d1565b6000546001600160a01b031633146104905760405162461bcd60e51b81526004016101de9061115d565b610430610921565b6000546001600160a01b031633146104c25760405162461bcd60e51b81526004016101de9061115d565b6001600160a01b0381166105275760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016101de565b6103fb816108d1565b600054600160a01b900460ff161561055a5760405162461bcd60e51b81526004016101de90611133565b60405163453a6aff60e11b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638a74d5fe906105ac90879087908790879060040161104f565b60206040518083038186803b1580156105c457600080fd5b505afa1580156105d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105fc9190610eed565b50600061063e85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061098692505050565b60208082015182516001600160a01b03166000908152600190925260408220549293509161066c90836111eb565b9050600081116106ae5760405162461bcd60e51b815260206004820152600d60248201526c139bc81b995dc81c995dd85c99609a1b60448201526064016101de565b816001600085600001516001600160a01b03166001600160a01b031681526020019081526020016000208190555061073e8360000151827f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663960dc08a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156103b357600080fd5b82600001516001600160a01b03167f6dd401e61ba732582a5eba3d54ccc3afb3609cd3ac1a166d1d36f75fc0aedcda8260405161077d91815260200190565b60405180910390a250505050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526107f99085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610a32565b50505050565b6040516001600160a01b03831660248201526044810182905261082f90849063a9059cbb60e01b906064016107c2565b505050565b600054600160a01b900460ff166108845760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016101de565b6000805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600054600160a01b900460ff161561094b5760405162461bcd60e51b81526004016101de90611133565b6000805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586108b43390565b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b60208301515183511015610a2a576109c883610b04565b909250905081600114156109f7576109e76109e284610b3e565b610bfb565b6001600160a01b031684526109b1565b8160021415610a1b57610a11610a0c84610b3e565b610c0c565b60208501526109b1565b610a258382610c43565b6109b1565b505050919050565b6000610a87826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610cb59092919063ffffffff16565b80519091501561082f5780806020019051810190610aa59190610eed565b61082f5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016101de565b6000806000610b1284610cce565b9050610b1f6008826111aa565b9250806007166005811115610b3657610b3661125f565b915050915091565b60606000610b4b83610cce565b90506000818460000151610b5f9190611192565b9050836020015151811115610b7357600080fd5b8167ffffffffffffffff811115610b8c57610b8c611275565b6040519080825280601f01601f191660200182016040528015610bb6576020820181803683370190505b50602080860151865192955091818601919083010160005b85811015610bf0578181015183820152610be9602082611192565b9050610bce565b505050935250919050565b6000610c0682610d50565b92915050565b6000602082511115610c1d57600080fd5b6020820151905081516020610c3291906111eb565b610c3d9060086111cc565b1c919050565b6000816005811115610c5757610c5761125f565b1415610c665761082f82610cce565b6002816005811115610c7a57610c7a61125f565b14156100a9576000610c8b83610cce565b90508083600001818151610c9f9190611192565b9052506020830151518351111561082f57600080fd5b6060610cc48484600085610d6f565b90505b9392505050565b602080820151825181019091015160009182805b600a811015610d4a5783811a9150610cfb8160076111cc565b82607f16901b851794508160801660001415610d3857610d1c816001611192565b86518790610d2b908390611192565b9052509395945050505050565b80610d428161122e565b915050610ce2565b50600080fd5b60008151601414610d6057600080fd5b5060200151600160601b900490565b606082471015610dd05760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016101de565b843b610e1e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101de565b600080866001600160a01b03168587604051610e3a9190611033565b60006040518083038185875af1925050503d8060008114610e77576040519150601f19603f3d011682016040523d82523d6000602084013e610e7c565b606091505b5091509150610e8c828286610e97565b979650505050505050565b60608315610ea6575081610cc7565b825115610eb65782518084602001fd5b8160405162461bcd60e51b81526004016101de9190611100565b600060208284031215610ee257600080fd5b8135610cc78161128b565b600060208284031215610eff57600080fd5b81518015158114610cc757600080fd5b60008060008060408587031215610f2557600080fd5b843567ffffffffffffffff80821115610f3d57600080fd5b818701915087601f830112610f5157600080fd5b813581811115610f6057600080fd5b886020828501011115610f7257600080fd5b602092830196509450908601359080821115610f8d57600080fd5b818701915087601f830112610fa157600080fd5b813581811115610fb057600080fd5b8860208260051b8501011115610fc557600080fd5b95989497505060200194505050565b600060208284031215610fe657600080fd5b8151610cc78161128b565b60006020828403121561100357600080fd5b5035919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008251611045818460208701611202565b9190910192915050565b60408152600061106360408301868861100a565b602083820381850152818583528183019050818660051b8401018760005b888110156110f057858303601f190184528135368b9003601e190181126110a757600080fd5b8a01803567ffffffffffffffff8111156110c057600080fd5b8036038c13156110cf57600080fd5b6110dc858289850161100a565b958701959450505090840190600101611081565b50909a9950505050505050505050565b602081526000825180602084015261111f816040850160208701611202565b601f01601f19169190910160400192915050565b60208082526010908201526f14185d5cd8589b194e881c185d5cd95960821b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600082198211156111a5576111a5611249565b500190565b6000826111c757634e487b7160e01b600052601260045260246000fd5b500490565b60008160001904831182151516156111e6576111e6611249565b500290565b6000828210156111fd576111fd611249565b500390565b60005b8381101561121d578181015183820152602001611205565b838111156107f95750506000910152565b600060001982141561124257611242611249565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146103fb57600080fdfea26469706673582212205f3f7fcf10d47cddd4778ad43cb24be65e99c95581c3c9acb6d055612abfac7b64736f6c63430008070033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractStaking\",\"name\":\"_staking\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getBondedValidatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"internalType\":\"structDataTypes.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"undelegationTokens\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Undelegation[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structDataTypes.DelegatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delAddr\",\"type\":\"address\"}],\"name\":\"getDelegatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"internalType\":\"structDataTypes.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"valAddr\",\"type\":\"address\"},{\"internalType\":\"enumDataTypes.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSelfDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"commissionRate\",\"type\":\"uint64\"}],\"internalType\":\"structDataTypes.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_valAddr\",\"type\":\"address\"}],\"name\":\"shouldBondValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161172938038061172983398101604081905261002f91610044565b60601b6001600160601b031916608052610074565b60006020828403121561005657600080fd5b81516001600160a01b038116811461006d57600080fd5b9392505050565b60805160601c6116366100f36000396000818160bb01528181610199015281816102b00152818161035c015281816105bd015281816106d4015281816107c8015281816108c40152818161096d01528181610a2901528181610c0501528181610d1c01528181610db101528181610ecd0152610f7301526116366000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638dc2336d1161005b5780638dc2336d1461012a578063c6fc1ed614610140578063d87ffe9114610168578063e9fe6b0b1461017057600080fd5b8063313019bb1461008d5780634cf088d9146100b657806366ab5d28146100f55780638a11d7c91461010a575b600080fd5b6100a061009b3660046110c8565b610193565b6040516100ad919061137f565b60405180910390f35b6100dd7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100ad565b6100fd6105b7565b6040516100ad919061144f565b61011d6101183660046110c8565b610794565b6040516100ad91906114ab565b6101326108bf565b6040519081526020016100ad565b61015361014e3660046110c8565b610b5a565b604080519283526020830191909152016100ad565b6100fd610bff565b61018361017e3660046110c8565b610d89565b60405190151581526020016100ad565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631cfe4f0b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101f057600080fd5b505afa158015610204573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022891906112f9565b90506000816001600160401b03811115610244576102446115d2565b60405190808252806020026020018201604052801561027d57816020015b61026a611024565b8152602001906001900390816102625790505b5090506000805b838163ffffffff16101561047c576040516324aec90f60e21b815263ffffffff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906392bb243c9060240160206040518083038186803b1580156102fa57600080fd5b505afa15801561030e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033291906110ec565b604051631dd9dfdf60e31b81526001600160a01b03808316600483015289811660248301529192507f00000000000000000000000000000000000000000000000000000000000000009091169063eecefef89060440160006040518083038186803b1580156103a057600080fd5b505afa1580156103b4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103dc91908101906111cd565b848363ffffffff16815181106103f4576103f46115bc565b6020026020010181905250838263ffffffff1681518110610417576104176115bc565b60200260200101516040015160001415806104565750838263ffffffff1681518110610445576104456115bc565b602002602001015160600151600014155b1561046957826104658161156c565b9350505b50806104748161156c565b915050610284565b5060008163ffffffff166001600160401b0381111561049d5761049d6115d2565b6040519080825280602002602001820160405280156104d657816020015b6104c3611024565b8152602001906001900390816104bb5790505b5090506000805b858163ffffffff1610156105ab57848163ffffffff1681518110610503576105036115bc565b60200260200101516040015160001415806105425750848163ffffffff1681518110610531576105316115bc565b602002602001015160600151600014155b1561059957848163ffffffff168151811061055f5761055f6115bc565b6020026020010151838363ffffffff168151811061057f5761057f6115bc565b602002602001018190525081806105959061156c565b9250505b806105a38161156c565b9150506104dd565b50909695505050505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316631cfe4f0b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561061457600080fd5b505afa158015610628573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064c91906112f9565b90506000816001600160401b03811115610668576106686115d2565b6040519080825280602002602001820160405280156106a157816020015b61068e61105c565b8152602001906001900390816106865790505b50905060005b828163ffffffff16101561078d576040516324aec90f60e21b815263ffffffff82166004820152610757907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316906392bb243c906024015b60206040518083038186803b15801561071f57600080fd5b505afa158015610733573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011891906110ec565b828263ffffffff168151811061076f5761076f6115bc565b602002602001018190525080806107859061156c565b9150506106a7565b5092915050565b61079c61105c565b604051631f4a58fb60e31b81526001600160a01b038381166004830152600091829182918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063fa52c7d8906024016101406040518083038186803b15801561080b57600080fd5b505afa15801561081f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610843919061112b565b995050509750505095509550955095506040518060e00160405280896001600160a01b0316815260200187600381111561087f5761087f6115a6565b8152602001866001600160a01b03168152602001858152602001848152602001838152602001826001600160401b03168152509650505050505050919050565b6000807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166389f9aab56040518163ffffffff1660e01b815260040160206040518083038186803b15801561091b57600080fd5b505afa15801561092f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095391906112f9565b60405163eb505dd560e01b81529091506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063eb505dd5906109a390600390600401611491565b60206040518083038186803b1580156109bb57600080fd5b505afa1580156109cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f391906112f9565b811015610a0257600091505090565b60001960015b8281101561078d5760405163acc62ccf60e01b8152600481018290526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063c8f9f98490829063acc62ccf9060240160206040518083038186803b158015610a7b57600080fd5b505afa158015610a8f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab391906110ec565b6040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260240160206040518083038186803b158015610af257600080fd5b505afa158015610b06573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b2a91906112f9565b905082811015610b475791508180610b4757600094505050505090565b5080610b5281611551565b915050610a08565b6000806000610b6884610193565b905060008060005b83518163ffffffff161015610bf357838163ffffffff1681518110610b9757610b976115bc565b60200260200101516020015183610bae9190611539565b9250838163ffffffff1681518110610bc857610bc86115bc565b60200260200101516060015182610bdf9190611539565b915080610beb8161156c565b915050610b70565b50909590945092505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166389f9aab56040518163ffffffff1660e01b815260040160206040518083038186803b158015610c5c57600080fd5b505afa158015610c70573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9491906112f9565b90506000816001600160401b03811115610cb057610cb06115d2565b604051908082528060200260200182016040528015610ce957816020015b610cd661105c565b815260200190600190039081610cce5790505b50905060005b828163ffffffff16101561078d5760405163acc62ccf60e01b815263ffffffff82166004820152610d53907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063acc62ccf90602401610707565b828263ffffffff1681518110610d6b57610d6b6115bc565b60200260200101819052508080610d819061156c565b915050610cef565b604051631f4a58fb60e31b81526001600160a01b0382811660048301526000918291829182917f0000000000000000000000000000000000000000000000000000000000000000169063fa52c7d8906024016101406040518083038186803b158015610df457600080fd5b505afa158015610e08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2c919061112b565b5050975050505050935050925060006003811115610e4c57610e4c6115a6565b836003811115610e5e57610e5e6115a6565b1480610e7b57506003836003811115610e7957610e796115a6565b145b15610e8b57506000949350505050565b806001600160401b0316431015610ea757506000949350505050565b6040516347abfdbf60e01b81526001600160a01b038681166004830152600160248301527f000000000000000000000000000000000000000000000000000000000000000016906347abfdbf9060440160206040518083038186803b158015610f0f57600080fd5b505afa158015610f23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f479190611109565b610f5657506000949350505050565b610f5e6108bf565b8211610f6f57506000949350505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166383cfb3186040518163ffffffff1660e01b815260040160206040518083038186803b158015610fca57600080fd5b505afa158015610fde573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100291906112f9565b9050804310156110185750600095945050505050565b50600195945050505050565b6040518060a0016040528060006001600160a01b03168152602001600081526020016000815260200160008152602001606081525090565b6040805160e0810190915260008082526020820190815260200160006001600160a01b0316815260200160008152602001600081526020016000815260200160006001600160401b031681525090565b80516001600160401b03811681146110c357600080fd5b919050565b6000602082840312156110da57600080fd5b81356110e5816115e8565b9392505050565b6000602082840312156110fe57600080fd5b81516110e5816115e8565b60006020828403121561111b57600080fd5b815180151581146110e557600080fd5b6000806000806000806000806000806101408b8d03121561114b57600080fd5b8a516004811061115a57600080fd5b60208c0151909a5061116b816115e8565b8099505060408b0151975060608b0151965060808b0151955060a08b0151945060c08b0151935061119e60e08c016110ac565b92506111ad6101008c016110ac565b91506111bc6101208c016110ac565b90509295989b9194979a5092959850565b600060208083850312156111e057600080fd5b82516001600160401b03808211156111f757600080fd5b9084019060a0828703121561120b57600080fd5b6112136114bf565b825161121e816115e8565b81528284015184820152604080840151818301526060808501519083015260808401518381111561124e57600080fd5b80850194505087601f85011261126357600080fd5b835183811115611275576112756115d2565b611283868260051b01611509565b8181528681019450858701600683901b870188018b10156112a357600080fd5b60009650865b838110156112e55784828d0312156112bf578788fd5b6112c76114e7565b82518152898301518a820152875295880195908401906001016112a9565b505060808401525090979650505050505050565b60006020828403121561130b57600080fd5b5051919050565b80516001600160a01b039081168352602082015160048110611336576113366115a6565b8060208501525080604083015116604084015250606081015160608301526080810151608083015260a081015160a08301526001600160401b0360c08201511660c08301525050565b60006020808301818452808551808352604092508286019150828160051b8701018488016000805b8481101561144057898403603f19018652825180516001600160a01b03168552888101518986015287810151888601526060808201519086015260809081015160a0918601829052805191860182905289019060c086019084905b8082101561142b578351805184528c01518c840152928b0192918a019160019190910190611402565b505096890196945050918701916001016113a7565b50919998505050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156105ab5761147e838551611312565b9284019260e0929092019160010161146b565b60208101600983106114a5576114a56115a6565b91905290565b60e081016114b98284611312565b92915050565b60405160a081016001600160401b03811182821017156114e1576114e16115d2565b60405290565b604080519081016001600160401b03811182821017156114e1576114e16115d2565b604051601f8201601f191681016001600160401b0381118282101715611531576115316115d2565b604052919050565b6000821982111561154c5761154c611590565b500190565b600060001982141561156557611565611590565b5060010190565b600063ffffffff8083168181141561158657611586611590565b6001019392505050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146115fd57600080fd5b5056fea26469706673582212209c48e3265fc97c794269bb00656a2f9993e28913313659d75767f5d5fe805ebc64736f6c63430008070033",
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
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[])[])
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
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[])[])
func (_Viewer *ViewerSession) GetDelegatorInfos(_delAddr common.Address) ([]DataTypesDelegatorInfo, error) {
	return _Viewer.Contract.GetDelegatorInfos(&_Viewer.CallOpts, _delAddr)
}

// GetDelegatorInfos is a free data retrieval call binding the contract method 0x313019bb.
//
// Solidity: function getDelegatorInfos(address _delAddr) view returns((address,uint256,uint256,uint256,(uint256,uint256)[])[])
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
