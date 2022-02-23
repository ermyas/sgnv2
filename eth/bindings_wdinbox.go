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

// ContractAsLPMetaData contains all meta data concerning the ContractAsLP contract.
var ContractAsLPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_wdSeq\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_toChain\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"_fromChains\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_ratios\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_slippages\",\"type\":\"uint32[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620016e3380380620016e38339810160408190526200003491620001c5565b600160005562000044336200008e565b6001805460ff60a01b191690556200005c33620000e0565b600380546001600160a01b039384166001600160a01b03199182161790915560048054929093169116179055620001fd565b600180546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03811660009081526002602052604090205460ff16156200014e5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640160405180910390fd5b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910160405180910390a150565b80516001600160a01b0381168114620001c057600080fd5b919050565b60008060408385031215620001d957600080fd5b620001e483620001a8565b9150620001f460208401620001a8565b90509250929050565b6114d6806200020d6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806380f51c1211610097578063a485529611610066578063a48552961461020e578063e78cea9214610221578063f2fde38b14610234578063fb0e722b1461024757600080fd5b806380f51c12146101ab57806382dc1ec4146101ce5780638456cb59146101e15780638da5cb5b146101e957600080fd5b80635c975abb116100d35780635c975abb146101765780636b2c0f55146101885780636ef8d66d1461019b578063715018a6146101a357600080fd5b80633f4ba83a1461010557806346fbf68e1461010f57806347e7ef24146101505780635668870014610163575b600080fd5b61010d61025a565b005b61013b61011d3660046110cf565b6001600160a01b031660009081526002602052604090205460ff1690565b60405190151581526020015b60405180910390f35b61010d61015e3660046110ea565b6102c8565b61010d6101713660046110ea565b61042e565b600154600160a01b900460ff1661013b565b61010d6101963660046110cf565b61061f565b61010d610685565b61010d61068e565b61013b6101b93660046110cf565b60026020526000908152604090205460ff1681565b61010d6101dc3660046110cf565b6106f2565b61010d610755565b6001546001600160a01b03165b6040516001600160a01b039091168152602001610147565b61010d61021c366004611178565b6107bc565b6003546101f6906001600160a01b031681565b61010d6102423660046110cf565b6108e5565b6004546101f6906001600160a01b031681565b3360009081526002602052604090205460ff166102be5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064015b60405180910390fd5b6102c66109c4565b565b6002600054141561031b5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016102b5565b6002600055600154600160a01b900460ff161561036d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102b5565b6001546001600160a01b031633146103c75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b5565b6103dc6001600160a01b038316333084610a6a565b604080513381526001600160a01b03841660208201529081018290527f8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a79060600160405180910390a150506001600055565b600154600160a01b900460ff161561047b5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102b5565b6001546001600160a01b031633146104d55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b5565b6040516370a0823160e01b815230600482015281906001600160a01b038416906370a082319060240160206040518083038186803b15801561051657600080fd5b505afa15801561052a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061054e9190611275565b101561059c5760405162461bcd60e51b815260206004820152601460248201527f696e73756666696369656e742062616c616e636500000000000000000000000060448201526064016102b5565b6003546105b6906001600160a01b03848116911683610b08565b6003546040516256688760e81b81526001600160a01b0384811660048301526024820184905290911690635668870090604401600060405180830381600087803b15801561060357600080fd5b505af1158015610617573d6000803e3d6000fd5b505050505050565b6001546001600160a01b031633146106795760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b5565b61068281610bc9565b50565b6102c633610bc9565b6001546001600160a01b031633146106e85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b5565b6102c66000610c89565b6001546001600160a01b0316331461074c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b5565b61068281610cf3565b3360009081526002602052604090205460ff166107b45760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016102b5565b6102c6610db0565b600154600160a01b900460ff16156108095760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102b5565b6001546001600160a01b031633146108635760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b5565b60048054604051635242a94b60e11b81526001600160a01b039091169163a4855296916108a6918f918f918f918f918f918f918f918f918f918f918f9101611321565b600060405180830381600087803b1580156108c057600080fd5b505af11580156108d4573d6000803e3d6000fd5b505050505050505050505050505050565b6001546001600160a01b0316331461093f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102b5565b6001600160a01b0381166109bb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102b5565b61068281610c89565b600154600160a01b900460ff16610a1d5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102b5565b6001805460ff60a01b191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6040516001600160a01b0380851660248301528316604482015260648101829052610b029085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610e38565b50505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b158015610b5457600080fd5b505afa158015610b68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8c9190611275565b610b9691906113dd565b6040516001600160a01b038516602482015260448101829052909150610b0290859063095ea7b360e01b90606401610a9e565b6001600160a01b03811660009081526002602052604090205460ff16610c315760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f7420706175736572000000000000000000000060448201526064016102b5565b6001600160a01b038116600081815260026020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e91015b60405180910390a150565b600180546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b03811660009081526002602052604090205460ff1615610d5c5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064016102b5565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f89101610c7e565b600154600160a01b900460ff1615610dfd5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102b5565b6001805460ff60a01b1916600160a01b1790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610a4d3390565b6000610e8d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610f229092919063ffffffff16565b805190915015610f1d5780806020019051810190610eab9190611403565b610f1d5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102b5565b505050565b6060610f318484600085610f3b565b90505b9392505050565b606082471015610fb35760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102b5565b843b6110015760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102b5565b600080866001600160a01b0316858760405161101d9190611451565b60006040518083038185875af1925050503d806000811461105a576040519150601f19603f3d011682016040523d82523d6000602084013e61105f565b606091505b509150915061106f82828661107a565b979650505050505050565b60608315611089575081610f34565b8251156110995782518084602001fd5b8160405162461bcd60e51b81526004016102b5919061146d565b80356001600160a01b03811681146110ca57600080fd5b919050565b6000602082840312156110e157600080fd5b610f34826110b3565b600080604083850312156110fd57600080fd5b611106836110b3565b946020939093013593505050565b803567ffffffffffffffff811681146110ca57600080fd5b60008083601f84011261113e57600080fd5b50813567ffffffffffffffff81111561115657600080fd5b6020830191508360208260051b850101111561117157600080fd5b9250929050565b600080600080600080600080600080600060e08c8e03121561119957600080fd5b6111a28c611114565b9a506111b060208d016110b3565b99506111be60408d01611114565b985067ffffffffffffffff8060608e013511156111da57600080fd5b6111ea8e60608f01358f0161112c565b909950975060808d013581101561120057600080fd5b6112108e60808f01358f0161112c565b909750955060a08d013581101561122657600080fd5b6112368e60a08f01358f0161112c565b909550935060c08d013581101561124c57600080fd5b5061125d8d60c08e01358e0161112c565b81935080925050509295989b509295989b9093969950565b60006020828403121561128757600080fd5b5051919050565b8183526000602080850194508260005b858110156112ca576001600160a01b036112b7836110b3565b168752958201959082019060010161129e565b509495945050505050565b818352600060208085019450826000805b8681101561131557823563ffffffff8116808214611302578384fd5b89525096830196918301916001016112e6565b50959695505050505050565b600060e0820167ffffffffffffffff808f16845260206001600160a01b038f1681860152818e16604086015260e06060860152828c8452610100860190508d935060005b8d81101561138a578361137786611114565b1682529382019390820190600101611365565b50858103608087015261139e818c8e61128e565b935050505082810360a08401526113b68187896112d5565b905082810360c08401526113cb8185876112d5565b9e9d5050505050505050505050505050565b600082198211156113fe57634e487b7160e01b600052601160045260246000fd5b500190565b60006020828403121561141557600080fd5b81518015158114610f3457600080fd5b60005b83811015611440578181015183820152602001611428565b83811115610b025750506000910152565b60008251611463818460208701611425565b9190910192915050565b602081526000825180602084015261148c816040850160208701611425565b601f01601f1916919091016040019291505056fea26469706673582212206affffaf4fce1c22f5aeac323dac40f0135a9ce8888e0463b6d154740f9ce3f264736f6c63430008090033",
}

// ContractAsLPABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractAsLPMetaData.ABI instead.
var ContractAsLPABI = ContractAsLPMetaData.ABI

// ContractAsLPBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractAsLPMetaData.Bin instead.
var ContractAsLPBin = ContractAsLPMetaData.Bin

// DeployContractAsLP deploys a new Ethereum contract, binding an instance of ContractAsLP to it.
func DeployContractAsLP(auth *bind.TransactOpts, backend bind.ContractBackend, _bridge common.Address, _inbox common.Address) (common.Address, *types.Transaction, *ContractAsLP, error) {
	parsed, err := ContractAsLPMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractAsLPBin), backend, _bridge, _inbox)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractAsLP{ContractAsLPCaller: ContractAsLPCaller{contract: contract}, ContractAsLPTransactor: ContractAsLPTransactor{contract: contract}, ContractAsLPFilterer: ContractAsLPFilterer{contract: contract}}, nil
}

// ContractAsLP is an auto generated Go binding around an Ethereum contract.
type ContractAsLP struct {
	ContractAsLPCaller     // Read-only binding to the contract
	ContractAsLPTransactor // Write-only binding to the contract
	ContractAsLPFilterer   // Log filterer for contract events
}

// ContractAsLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractAsLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAsLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractAsLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAsLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractAsLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractAsLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractAsLPSession struct {
	Contract     *ContractAsLP     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractAsLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractAsLPCallerSession struct {
	Contract *ContractAsLPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ContractAsLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractAsLPTransactorSession struct {
	Contract     *ContractAsLPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractAsLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractAsLPRaw struct {
	Contract *ContractAsLP // Generic contract binding to access the raw methods on
}

// ContractAsLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractAsLPCallerRaw struct {
	Contract *ContractAsLPCaller // Generic read-only contract binding to access the raw methods on
}

// ContractAsLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractAsLPTransactorRaw struct {
	Contract *ContractAsLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractAsLP creates a new instance of ContractAsLP, bound to a specific deployed contract.
func NewContractAsLP(address common.Address, backend bind.ContractBackend) (*ContractAsLP, error) {
	contract, err := bindContractAsLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractAsLP{ContractAsLPCaller: ContractAsLPCaller{contract: contract}, ContractAsLPTransactor: ContractAsLPTransactor{contract: contract}, ContractAsLPFilterer: ContractAsLPFilterer{contract: contract}}, nil
}

// NewContractAsLPCaller creates a new read-only instance of ContractAsLP, bound to a specific deployed contract.
func NewContractAsLPCaller(address common.Address, caller bind.ContractCaller) (*ContractAsLPCaller, error) {
	contract, err := bindContractAsLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAsLPCaller{contract: contract}, nil
}

// NewContractAsLPTransactor creates a new write-only instance of ContractAsLP, bound to a specific deployed contract.
func NewContractAsLPTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractAsLPTransactor, error) {
	contract, err := bindContractAsLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractAsLPTransactor{contract: contract}, nil
}

// NewContractAsLPFilterer creates a new log filterer instance of ContractAsLP, bound to a specific deployed contract.
func NewContractAsLPFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractAsLPFilterer, error) {
	contract, err := bindContractAsLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractAsLPFilterer{contract: contract}, nil
}

// bindContractAsLP binds a generic wrapper to an already deployed contract.
func bindContractAsLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractAsLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAsLP *ContractAsLPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAsLP.Contract.ContractAsLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAsLP *ContractAsLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAsLP.Contract.ContractAsLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAsLP *ContractAsLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAsLP.Contract.ContractAsLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractAsLP *ContractAsLPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractAsLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractAsLP *ContractAsLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAsLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractAsLP *ContractAsLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractAsLP.Contract.contract.Transact(opts, method, params...)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ContractAsLP *ContractAsLPCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAsLP.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ContractAsLP *ContractAsLPSession) Bridge() (common.Address, error) {
	return _ContractAsLP.Contract.Bridge(&_ContractAsLP.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_ContractAsLP *ContractAsLPCallerSession) Bridge() (common.Address, error) {
	return _ContractAsLP.Contract.Bridge(&_ContractAsLP.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ContractAsLP *ContractAsLPCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAsLP.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ContractAsLP *ContractAsLPSession) Inbox() (common.Address, error) {
	return _ContractAsLP.Contract.Inbox(&_ContractAsLP.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_ContractAsLP *ContractAsLPCallerSession) Inbox() (common.Address, error) {
	return _ContractAsLP.Contract.Inbox(&_ContractAsLP.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ContractAsLP *ContractAsLPCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAsLP.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ContractAsLP *ContractAsLPSession) IsPauser(account common.Address) (bool, error) {
	return _ContractAsLP.Contract.IsPauser(&_ContractAsLP.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_ContractAsLP *ContractAsLPCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ContractAsLP.Contract.IsPauser(&_ContractAsLP.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAsLP *ContractAsLPCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractAsLP.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAsLP *ContractAsLPSession) Owner() (common.Address, error) {
	return _ContractAsLP.Contract.Owner(&_ContractAsLP.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractAsLP *ContractAsLPCallerSession) Owner() (common.Address, error) {
	return _ContractAsLP.Contract.Owner(&_ContractAsLP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractAsLP *ContractAsLPCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractAsLP.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractAsLP *ContractAsLPSession) Paused() (bool, error) {
	return _ContractAsLP.Contract.Paused(&_ContractAsLP.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ContractAsLP *ContractAsLPCallerSession) Paused() (bool, error) {
	return _ContractAsLP.Contract.Paused(&_ContractAsLP.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_ContractAsLP *ContractAsLPCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ContractAsLP.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_ContractAsLP *ContractAsLPSession) Pausers(arg0 common.Address) (bool, error) {
	return _ContractAsLP.Contract.Pausers(&_ContractAsLP.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_ContractAsLP *ContractAsLPCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _ContractAsLP.Contract.Pausers(&_ContractAsLP.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_ContractAsLP *ContractAsLPTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "addLiquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_ContractAsLP *ContractAsLPSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAsLP.Contract.AddLiquidity(&_ContractAsLP.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_ContractAsLP *ContractAsLPTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAsLP.Contract.AddLiquidity(&_ContractAsLP.TransactOpts, _token, _amount)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ContractAsLP *ContractAsLPTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ContractAsLP *ContractAsLPSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ContractAsLP.Contract.AddPauser(&_ContractAsLP.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ContractAsLP *ContractAsLPTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ContractAsLP.Contract.AddPauser(&_ContractAsLP.TransactOpts, account)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_ContractAsLP *ContractAsLPTransactor) Deposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "deposit", _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_ContractAsLP *ContractAsLPSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAsLP.Contract.Deposit(&_ContractAsLP.TransactOpts, _token, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _token, uint256 _amount) returns()
func (_ContractAsLP *ContractAsLPTransactorSession) Deposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ContractAsLP.Contract.Deposit(&_ContractAsLP.TransactOpts, _token, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractAsLP *ContractAsLPTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractAsLP *ContractAsLPSession) Pause() (*types.Transaction, error) {
	return _ContractAsLP.Contract.Pause(&_ContractAsLP.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ContractAsLP *ContractAsLPTransactorSession) Pause() (*types.Transaction, error) {
	return _ContractAsLP.Contract.Pause(&_ContractAsLP.TransactOpts)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_ContractAsLP *ContractAsLPTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_ContractAsLP *ContractAsLPSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _ContractAsLP.Contract.RemovePauser(&_ContractAsLP.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_ContractAsLP *ContractAsLPTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _ContractAsLP.Contract.RemovePauser(&_ContractAsLP.TransactOpts, account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAsLP *ContractAsLPTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAsLP *ContractAsLPSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAsLP.Contract.RenounceOwnership(&_ContractAsLP.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractAsLP *ContractAsLPTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractAsLP.Contract.RenounceOwnership(&_ContractAsLP.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ContractAsLP *ContractAsLPTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ContractAsLP *ContractAsLPSession) RenouncePauser() (*types.Transaction, error) {
	return _ContractAsLP.Contract.RenouncePauser(&_ContractAsLP.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ContractAsLP *ContractAsLPTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ContractAsLP.Contract.RenouncePauser(&_ContractAsLP.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAsLP *ContractAsLPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAsLP *ContractAsLPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAsLP.Contract.TransferOwnership(&_ContractAsLP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractAsLP *ContractAsLPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractAsLP.Contract.TransferOwnership(&_ContractAsLP.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractAsLP *ContractAsLPTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractAsLP *ContractAsLPSession) Unpause() (*types.Transaction, error) {
	return _ContractAsLP.Contract.Unpause(&_ContractAsLP.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ContractAsLP *ContractAsLPTransactorSession) Unpause() (*types.Transaction, error) {
	return _ContractAsLP.Contract.Unpause(&_ContractAsLP.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa4855296.
//
// Solidity: function withdraw(uint64 _wdSeq, address _receiver, uint64 _toChain, uint64[] _fromChains, address[] _tokens, uint32[] _ratios, uint32[] _slippages) returns()
func (_ContractAsLP *ContractAsLPTransactor) Withdraw(opts *bind.TransactOpts, _wdSeq uint64, _receiver common.Address, _toChain uint64, _fromChains []uint64, _tokens []common.Address, _ratios []uint32, _slippages []uint32) (*types.Transaction, error) {
	return _ContractAsLP.contract.Transact(opts, "withdraw", _wdSeq, _receiver, _toChain, _fromChains, _tokens, _ratios, _slippages)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa4855296.
//
// Solidity: function withdraw(uint64 _wdSeq, address _receiver, uint64 _toChain, uint64[] _fromChains, address[] _tokens, uint32[] _ratios, uint32[] _slippages) returns()
func (_ContractAsLP *ContractAsLPSession) Withdraw(_wdSeq uint64, _receiver common.Address, _toChain uint64, _fromChains []uint64, _tokens []common.Address, _ratios []uint32, _slippages []uint32) (*types.Transaction, error) {
	return _ContractAsLP.Contract.Withdraw(&_ContractAsLP.TransactOpts, _wdSeq, _receiver, _toChain, _fromChains, _tokens, _ratios, _slippages)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa4855296.
//
// Solidity: function withdraw(uint64 _wdSeq, address _receiver, uint64 _toChain, uint64[] _fromChains, address[] _tokens, uint32[] _ratios, uint32[] _slippages) returns()
func (_ContractAsLP *ContractAsLPTransactorSession) Withdraw(_wdSeq uint64, _receiver common.Address, _toChain uint64, _fromChains []uint64, _tokens []common.Address, _ratios []uint32, _slippages []uint32) (*types.Transaction, error) {
	return _ContractAsLP.Contract.Withdraw(&_ContractAsLP.TransactOpts, _wdSeq, _receiver, _toChain, _fromChains, _tokens, _ratios, _slippages)
}

// ContractAsLPDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the ContractAsLP contract.
type ContractAsLPDepositedIterator struct {
	Event *ContractAsLPDeposited // Event containing the contract specifics and raw log

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
func (it *ContractAsLPDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAsLPDeposited)
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
		it.Event = new(ContractAsLPDeposited)
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
func (it *ContractAsLPDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAsLPDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAsLPDeposited represents a Deposited event raised by the ContractAsLP contract.
type ContractAsLPDeposited struct {
	Depositor common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address depositor, address token, uint256 amount)
func (_ContractAsLP *ContractAsLPFilterer) FilterDeposited(opts *bind.FilterOpts) (*ContractAsLPDepositedIterator, error) {

	logs, sub, err := _ContractAsLP.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &ContractAsLPDepositedIterator{contract: _ContractAsLP.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address depositor, address token, uint256 amount)
func (_ContractAsLP *ContractAsLPFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *ContractAsLPDeposited) (event.Subscription, error) {

	logs, sub, err := _ContractAsLP.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAsLPDeposited)
				if err := _ContractAsLP.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x8752a472e571a816aea92eec8dae9baf628e840f4929fbcc2d155e6233ff68a7.
//
// Solidity: event Deposited(address depositor, address token, uint256 amount)
func (_ContractAsLP *ContractAsLPFilterer) ParseDeposited(log types.Log) (*ContractAsLPDeposited, error) {
	event := new(ContractAsLPDeposited)
	if err := _ContractAsLP.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAsLPOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractAsLP contract.
type ContractAsLPOwnershipTransferredIterator struct {
	Event *ContractAsLPOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractAsLPOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAsLPOwnershipTransferred)
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
		it.Event = new(ContractAsLPOwnershipTransferred)
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
func (it *ContractAsLPOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAsLPOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAsLPOwnershipTransferred represents a OwnershipTransferred event raised by the ContractAsLP contract.
type ContractAsLPOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAsLP *ContractAsLPFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractAsLPOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAsLP.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractAsLPOwnershipTransferredIterator{contract: _ContractAsLP.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractAsLP *ContractAsLPFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractAsLPOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractAsLP.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAsLPOwnershipTransferred)
				if err := _ContractAsLP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractAsLP *ContractAsLPFilterer) ParseOwnershipTransferred(log types.Log) (*ContractAsLPOwnershipTransferred, error) {
	event := new(ContractAsLPOwnershipTransferred)
	if err := _ContractAsLP.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAsLPPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractAsLP contract.
type ContractAsLPPausedIterator struct {
	Event *ContractAsLPPaused // Event containing the contract specifics and raw log

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
func (it *ContractAsLPPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAsLPPaused)
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
		it.Event = new(ContractAsLPPaused)
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
func (it *ContractAsLPPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAsLPPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAsLPPaused represents a Paused event raised by the ContractAsLP contract.
type ContractAsLPPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractAsLP *ContractAsLPFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractAsLPPausedIterator, error) {

	logs, sub, err := _ContractAsLP.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractAsLPPausedIterator{contract: _ContractAsLP.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ContractAsLP *ContractAsLPFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractAsLPPaused) (event.Subscription, error) {

	logs, sub, err := _ContractAsLP.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAsLPPaused)
				if err := _ContractAsLP.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ContractAsLP *ContractAsLPFilterer) ParsePaused(log types.Log) (*ContractAsLPPaused, error) {
	event := new(ContractAsLPPaused)
	if err := _ContractAsLP.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAsLPPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ContractAsLP contract.
type ContractAsLPPauserAddedIterator struct {
	Event *ContractAsLPPauserAdded // Event containing the contract specifics and raw log

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
func (it *ContractAsLPPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAsLPPauserAdded)
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
		it.Event = new(ContractAsLPPauserAdded)
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
func (it *ContractAsLPPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAsLPPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAsLPPauserAdded represents a PauserAdded event raised by the ContractAsLP contract.
type ContractAsLPPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_ContractAsLP *ContractAsLPFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*ContractAsLPPauserAddedIterator, error) {

	logs, sub, err := _ContractAsLP.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &ContractAsLPPauserAddedIterator{contract: _ContractAsLP.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_ContractAsLP *ContractAsLPFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ContractAsLPPauserAdded) (event.Subscription, error) {

	logs, sub, err := _ContractAsLP.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAsLPPauserAdded)
				if err := _ContractAsLP.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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
func (_ContractAsLP *ContractAsLPFilterer) ParsePauserAdded(log types.Log) (*ContractAsLPPauserAdded, error) {
	event := new(ContractAsLPPauserAdded)
	if err := _ContractAsLP.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAsLPPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ContractAsLP contract.
type ContractAsLPPauserRemovedIterator struct {
	Event *ContractAsLPPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ContractAsLPPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAsLPPauserRemoved)
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
		it.Event = new(ContractAsLPPauserRemoved)
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
func (it *ContractAsLPPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAsLPPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAsLPPauserRemoved represents a PauserRemoved event raised by the ContractAsLP contract.
type ContractAsLPPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_ContractAsLP *ContractAsLPFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*ContractAsLPPauserRemovedIterator, error) {

	logs, sub, err := _ContractAsLP.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &ContractAsLPPauserRemovedIterator{contract: _ContractAsLP.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_ContractAsLP *ContractAsLPFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ContractAsLPPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _ContractAsLP.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAsLPPauserRemoved)
				if err := _ContractAsLP.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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
func (_ContractAsLP *ContractAsLPFilterer) ParsePauserRemoved(log types.Log) (*ContractAsLPPauserRemoved, error) {
	event := new(ContractAsLPPauserRemoved)
	if err := _ContractAsLP.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAsLPUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractAsLP contract.
type ContractAsLPUnpausedIterator struct {
	Event *ContractAsLPUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractAsLPUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAsLPUnpaused)
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
		it.Event = new(ContractAsLPUnpaused)
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
func (it *ContractAsLPUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAsLPUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAsLPUnpaused represents a Unpaused event raised by the ContractAsLP contract.
type ContractAsLPUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractAsLP *ContractAsLPFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractAsLPUnpausedIterator, error) {

	logs, sub, err := _ContractAsLP.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractAsLPUnpausedIterator{contract: _ContractAsLP.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ContractAsLP *ContractAsLPFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractAsLPUnpaused) (event.Subscription, error) {

	logs, sub, err := _ContractAsLP.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAsLPUnpaused)
				if err := _ContractAsLP.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ContractAsLP *ContractAsLPFilterer) ParseUnpaused(log types.Log) (*ContractAsLPUnpaused, error) {
	event := new(ContractAsLPUnpaused)
	if err := _ContractAsLP.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawInboxMetaData contains all meta data concerning the WithdrawInbox contract.
var WithdrawInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"toChain\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"fromChains\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint32[]\",\"name\":\"ratios\",\"type\":\"uint32[]\"},{\"indexed\":false,\"internalType\":\"uint32[]\",\"name\":\"slippages\",\"type\":\"uint32[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"minimalMaxSlippage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minimalMaxSlippage\",\"type\":\"uint32\"}],\"name\":\"setMinimalMaxSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validityPeriod\",\"type\":\"uint256\"}],\"name\":\"setValidityPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validityPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_wdSeq\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_toChain\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"_fromChains\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_ratios\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"_slippages\",\"type\":\"uint32[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a33610025565b611c20600155610075565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6109f4806100846000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806387e647ed1161005b57806387e647ed146100df5780638da5cb5b146100f2578063a48552961461010d578063f2fde38b1461012057600080fd5b80631e6c3850146100825780632fd1b0a41461009e57806348234126146100ca575b600080fd5b61008b60015481565b6040519081526020015b60405180910390f35b6000546100b590600160a01b900463ffffffff1681565b60405163ffffffff9091168152602001610095565b6100dd6100d8366004610641565b610133565b005b6100dd6100ed366004610663565b6101df565b6000546040516001600160a01b039091168152602001610095565b6100dd61011b3660046106f7565b61024d565b6100dd61012e3660046107f4565b6104cf565b336101466000546001600160a01b031690565b6001600160a01b0316146101a15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6000805463ffffffff909216600160a01b027fffffffffffffffff00000000ffffffffffffffffffffffffffffffffffffffff909216919091179055565b336101f26000546001600160a01b031690565b6001600160a01b0316146102485760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610198565b600155565b8661029a5760405162461bcd60e51b815260206004820152601860248201527f656d707479207769746864726177616c207265717565737400000000000000006044820152606401610198565b84871480156102a857508287145b80156102b357508087145b6102ff5760405162461bcd60e51b815260206004820152600f60248201527f6c656e677468206d69736d6174636800000000000000000000000000000000006044820152606401610198565b60005b8381101561045f57600085858381811061031e5761031e61080f565b90506020020160208101906103339190610641565b63ffffffff1611801561037657506305f5e1008585838181106103585761035861080f565b905060200201602081019061036d9190610641565b63ffffffff1611155b6103c25760405162461bcd60e51b815260206004820152600d60248201527f696e76616c696420726174696f000000000000000000000000000000000000006044820152606401610198565b600054600160a01b900463ffffffff168383838181106103e4576103e461080f565b90506020020160208101906103f99190610641565b63ffffffff16101561044d5760405162461bcd60e51b815260206004820152601260248201527f736c69707061676520746f6f20736d616c6c00000000000000000000000000006044820152606401610198565b806104578161083b565b915050610302565b506000600154426104709190610856565b90507f7e2b24139224d852dd26bdb9f06f8136f7a1c9227a386d815a4ed8f1b8d7cc958c338d8d8d8d8d8d8d8d8d8d8d6040516104b99d9c9b9a999897969594939291906108ee565b60405180910390a1505050505050505050505050565b336104e26000546001600160a01b031690565b6001600160a01b0316146105385760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610198565b6001600160a01b0381166105b45760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610198565b6105bd816105c0565b50565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b803563ffffffff8116811461063c57600080fd5b919050565b60006020828403121561065357600080fd5b61065c82610628565b9392505050565b60006020828403121561067557600080fd5b5035919050565b803567ffffffffffffffff8116811461063c57600080fd5b80356001600160a01b038116811461063c57600080fd5b60008083601f8401126106bd57600080fd5b50813567ffffffffffffffff8111156106d557600080fd5b6020830191508360208260051b85010111156106f057600080fd5b9250929050565b600080600080600080600080600080600060e08c8e03121561071857600080fd5b6107218c61067c565b9a5061072f60208d01610694565b995061073d60408d0161067c565b985067ffffffffffffffff8060608e0135111561075957600080fd5b6107698e60608f01358f016106ab565b909950975060808d013581101561077f57600080fd5b61078f8e60808f01358f016106ab565b909750955060a08d01358110156107a557600080fd5b6107b58e60a08f01358f016106ab565b909550935060c08d01358110156107cb57600080fd5b506107dc8d60c08e01358e016106ab565b81935080925050509295989b509295989b9093969950565b60006020828403121561080657600080fd5b61065c82610694565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561084f5761084f610825565b5060010190565b6000821982111561086957610869610825565b500190565b8183526000602080850194508260005b858110156108aa576001600160a01b0361089783610694565b168752958201959082019060010161087e565b509495945050505050565b8183526000602080850194508260005b858110156108aa5763ffffffff6108db83610628565b16875295820195908201906001016108c5565b67ffffffffffffffff8e811682526001600160a01b038e811660208401528d1660408301528b166060820152610120608082018190528101899052600061014082018b825b8c8110156109635767ffffffffffffffff61094d8361067c565b1683526020928301929190910190600101610933565b505082810360a0840152610978818a8c61086e565b905082810360c084015261098d81888a6108b5565b905082810360e08401526109a28186886108b5565b915050826101008301529e9d505050505050505050505050505056fea26469706673582212208e28c74959766df1b0d2525418cf050f48ded6e6eb99fded0a61a9b5353b4f8764736f6c63430008090033",
}

// WithdrawInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawInboxMetaData.ABI instead.
var WithdrawInboxABI = WithdrawInboxMetaData.ABI

// WithdrawInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WithdrawInboxMetaData.Bin instead.
var WithdrawInboxBin = WithdrawInboxMetaData.Bin

// DeployWithdrawInbox deploys a new Ethereum contract, binding an instance of WithdrawInbox to it.
func DeployWithdrawInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WithdrawInbox, error) {
	parsed, err := WithdrawInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WithdrawInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WithdrawInbox{WithdrawInboxCaller: WithdrawInboxCaller{contract: contract}, WithdrawInboxTransactor: WithdrawInboxTransactor{contract: contract}, WithdrawInboxFilterer: WithdrawInboxFilterer{contract: contract}}, nil
}

// WithdrawInbox is an auto generated Go binding around an Ethereum contract.
type WithdrawInbox struct {
	WithdrawInboxCaller     // Read-only binding to the contract
	WithdrawInboxTransactor // Write-only binding to the contract
	WithdrawInboxFilterer   // Log filterer for contract events
}

// WithdrawInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawInboxSession struct {
	Contract     *WithdrawInbox    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawInboxCallerSession struct {
	Contract *WithdrawInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WithdrawInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawInboxTransactorSession struct {
	Contract     *WithdrawInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WithdrawInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawInboxRaw struct {
	Contract *WithdrawInbox // Generic contract binding to access the raw methods on
}

// WithdrawInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawInboxCallerRaw struct {
	Contract *WithdrawInboxCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawInboxTransactorRaw struct {
	Contract *WithdrawInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawInbox creates a new instance of WithdrawInbox, bound to a specific deployed contract.
func NewWithdrawInbox(address common.Address, backend bind.ContractBackend) (*WithdrawInbox, error) {
	contract, err := bindWithdrawInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawInbox{WithdrawInboxCaller: WithdrawInboxCaller{contract: contract}, WithdrawInboxTransactor: WithdrawInboxTransactor{contract: contract}, WithdrawInboxFilterer: WithdrawInboxFilterer{contract: contract}}, nil
}

// NewWithdrawInboxCaller creates a new read-only instance of WithdrawInbox, bound to a specific deployed contract.
func NewWithdrawInboxCaller(address common.Address, caller bind.ContractCaller) (*WithdrawInboxCaller, error) {
	contract, err := bindWithdrawInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawInboxCaller{contract: contract}, nil
}

// NewWithdrawInboxTransactor creates a new write-only instance of WithdrawInbox, bound to a specific deployed contract.
func NewWithdrawInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawInboxTransactor, error) {
	contract, err := bindWithdrawInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawInboxTransactor{contract: contract}, nil
}

// NewWithdrawInboxFilterer creates a new log filterer instance of WithdrawInbox, bound to a specific deployed contract.
func NewWithdrawInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawInboxFilterer, error) {
	contract, err := bindWithdrawInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawInboxFilterer{contract: contract}, nil
}

// bindWithdrawInbox binds a generic wrapper to an already deployed contract.
func bindWithdrawInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawInbox *WithdrawInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawInbox.Contract.WithdrawInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawInbox *WithdrawInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.WithdrawInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawInbox *WithdrawInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.WithdrawInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawInbox *WithdrawInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawInbox *WithdrawInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawInbox *WithdrawInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.contract.Transact(opts, method, params...)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_WithdrawInbox *WithdrawInboxCaller) MinimalMaxSlippage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _WithdrawInbox.contract.Call(opts, &out, "minimalMaxSlippage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_WithdrawInbox *WithdrawInboxSession) MinimalMaxSlippage() (uint32, error) {
	return _WithdrawInbox.Contract.MinimalMaxSlippage(&_WithdrawInbox.CallOpts)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_WithdrawInbox *WithdrawInboxCallerSession) MinimalMaxSlippage() (uint32, error) {
	return _WithdrawInbox.Contract.MinimalMaxSlippage(&_WithdrawInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawInbox *WithdrawInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawInbox *WithdrawInboxSession) Owner() (common.Address, error) {
	return _WithdrawInbox.Contract.Owner(&_WithdrawInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WithdrawInbox *WithdrawInboxCallerSession) Owner() (common.Address, error) {
	return _WithdrawInbox.Contract.Owner(&_WithdrawInbox.CallOpts)
}

// ValidityPeriod is a free data retrieval call binding the contract method 0x1e6c3850.
//
// Solidity: function validityPeriod() view returns(uint256)
func (_WithdrawInbox *WithdrawInboxCaller) ValidityPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawInbox.contract.Call(opts, &out, "validityPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidityPeriod is a free data retrieval call binding the contract method 0x1e6c3850.
//
// Solidity: function validityPeriod() view returns(uint256)
func (_WithdrawInbox *WithdrawInboxSession) ValidityPeriod() (*big.Int, error) {
	return _WithdrawInbox.Contract.ValidityPeriod(&_WithdrawInbox.CallOpts)
}

// ValidityPeriod is a free data retrieval call binding the contract method 0x1e6c3850.
//
// Solidity: function validityPeriod() view returns(uint256)
func (_WithdrawInbox *WithdrawInboxCallerSession) ValidityPeriod() (*big.Int, error) {
	return _WithdrawInbox.Contract.ValidityPeriod(&_WithdrawInbox.CallOpts)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_WithdrawInbox *WithdrawInboxTransactor) SetMinimalMaxSlippage(opts *bind.TransactOpts, _minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _WithdrawInbox.contract.Transact(opts, "setMinimalMaxSlippage", _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_WithdrawInbox *WithdrawInboxSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.SetMinimalMaxSlippage(&_WithdrawInbox.TransactOpts, _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_WithdrawInbox *WithdrawInboxTransactorSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.SetMinimalMaxSlippage(&_WithdrawInbox.TransactOpts, _minimalMaxSlippage)
}

// SetValidityPeriod is a paid mutator transaction binding the contract method 0x87e647ed.
//
// Solidity: function setValidityPeriod(uint256 _validityPeriod) returns()
func (_WithdrawInbox *WithdrawInboxTransactor) SetValidityPeriod(opts *bind.TransactOpts, _validityPeriod *big.Int) (*types.Transaction, error) {
	return _WithdrawInbox.contract.Transact(opts, "setValidityPeriod", _validityPeriod)
}

// SetValidityPeriod is a paid mutator transaction binding the contract method 0x87e647ed.
//
// Solidity: function setValidityPeriod(uint256 _validityPeriod) returns()
func (_WithdrawInbox *WithdrawInboxSession) SetValidityPeriod(_validityPeriod *big.Int) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.SetValidityPeriod(&_WithdrawInbox.TransactOpts, _validityPeriod)
}

// SetValidityPeriod is a paid mutator transaction binding the contract method 0x87e647ed.
//
// Solidity: function setValidityPeriod(uint256 _validityPeriod) returns()
func (_WithdrawInbox *WithdrawInboxTransactorSession) SetValidityPeriod(_validityPeriod *big.Int) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.SetValidityPeriod(&_WithdrawInbox.TransactOpts, _validityPeriod)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawInbox *WithdrawInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawInbox *WithdrawInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.TransferOwnership(&_WithdrawInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WithdrawInbox *WithdrawInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.TransferOwnership(&_WithdrawInbox.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa4855296.
//
// Solidity: function withdraw(uint64 _wdSeq, address _receiver, uint64 _toChain, uint64[] _fromChains, address[] _tokens, uint32[] _ratios, uint32[] _slippages) returns()
func (_WithdrawInbox *WithdrawInboxTransactor) Withdraw(opts *bind.TransactOpts, _wdSeq uint64, _receiver common.Address, _toChain uint64, _fromChains []uint64, _tokens []common.Address, _ratios []uint32, _slippages []uint32) (*types.Transaction, error) {
	return _WithdrawInbox.contract.Transact(opts, "withdraw", _wdSeq, _receiver, _toChain, _fromChains, _tokens, _ratios, _slippages)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa4855296.
//
// Solidity: function withdraw(uint64 _wdSeq, address _receiver, uint64 _toChain, uint64[] _fromChains, address[] _tokens, uint32[] _ratios, uint32[] _slippages) returns()
func (_WithdrawInbox *WithdrawInboxSession) Withdraw(_wdSeq uint64, _receiver common.Address, _toChain uint64, _fromChains []uint64, _tokens []common.Address, _ratios []uint32, _slippages []uint32) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.Withdraw(&_WithdrawInbox.TransactOpts, _wdSeq, _receiver, _toChain, _fromChains, _tokens, _ratios, _slippages)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa4855296.
//
// Solidity: function withdraw(uint64 _wdSeq, address _receiver, uint64 _toChain, uint64[] _fromChains, address[] _tokens, uint32[] _ratios, uint32[] _slippages) returns()
func (_WithdrawInbox *WithdrawInboxTransactorSession) Withdraw(_wdSeq uint64, _receiver common.Address, _toChain uint64, _fromChains []uint64, _tokens []common.Address, _ratios []uint32, _slippages []uint32) (*types.Transaction, error) {
	return _WithdrawInbox.Contract.Withdraw(&_WithdrawInbox.TransactOpts, _wdSeq, _receiver, _toChain, _fromChains, _tokens, _ratios, _slippages)
}

// WithdrawInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WithdrawInbox contract.
type WithdrawInboxOwnershipTransferredIterator struct {
	Event *WithdrawInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WithdrawInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawInboxOwnershipTransferred)
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
		it.Event = new(WithdrawInboxOwnershipTransferred)
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
func (it *WithdrawInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawInboxOwnershipTransferred represents a OwnershipTransferred event raised by the WithdrawInbox contract.
type WithdrawInboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WithdrawInbox *WithdrawInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WithdrawInboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WithdrawInbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawInboxOwnershipTransferredIterator{contract: _WithdrawInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WithdrawInbox *WithdrawInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WithdrawInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WithdrawInbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawInboxOwnershipTransferred)
				if err := _WithdrawInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WithdrawInbox *WithdrawInboxFilterer) ParseOwnershipTransferred(log types.Log) (*WithdrawInboxOwnershipTransferred, error) {
	event := new(WithdrawInboxOwnershipTransferred)
	if err := _WithdrawInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawInboxWithdrawalRequestIterator is returned from FilterWithdrawalRequest and is used to iterate over the raw logs and unpacked data for WithdrawalRequest events raised by the WithdrawInbox contract.
type WithdrawInboxWithdrawalRequestIterator struct {
	Event *WithdrawInboxWithdrawalRequest // Event containing the contract specifics and raw log

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
func (it *WithdrawInboxWithdrawalRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawInboxWithdrawalRequest)
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
		it.Event = new(WithdrawInboxWithdrawalRequest)
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
func (it *WithdrawInboxWithdrawalRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawInboxWithdrawalRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawInboxWithdrawalRequest represents a WithdrawalRequest event raised by the WithdrawInbox contract.
type WithdrawInboxWithdrawalRequest struct {
	SeqNum     uint64
	Sender     common.Address
	Receiver   common.Address
	ToChain    uint64
	FromChains []uint64
	Tokens     []common.Address
	Ratios     []uint32
	Slippages  []uint32
	Deadline   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequest is a free log retrieval operation binding the contract event 0x7e2b24139224d852dd26bdb9f06f8136f7a1c9227a386d815a4ed8f1b8d7cc95.
//
// Solidity: event WithdrawalRequest(uint64 seqNum, address sender, address receiver, uint64 toChain, uint64[] fromChains, address[] tokens, uint32[] ratios, uint32[] slippages, uint256 deadline)
func (_WithdrawInbox *WithdrawInboxFilterer) FilterWithdrawalRequest(opts *bind.FilterOpts) (*WithdrawInboxWithdrawalRequestIterator, error) {

	logs, sub, err := _WithdrawInbox.contract.FilterLogs(opts, "WithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return &WithdrawInboxWithdrawalRequestIterator{contract: _WithdrawInbox.contract, event: "WithdrawalRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequest is a free log subscription operation binding the contract event 0x7e2b24139224d852dd26bdb9f06f8136f7a1c9227a386d815a4ed8f1b8d7cc95.
//
// Solidity: event WithdrawalRequest(uint64 seqNum, address sender, address receiver, uint64 toChain, uint64[] fromChains, address[] tokens, uint32[] ratios, uint32[] slippages, uint256 deadline)
func (_WithdrawInbox *WithdrawInboxFilterer) WatchWithdrawalRequest(opts *bind.WatchOpts, sink chan<- *WithdrawInboxWithdrawalRequest) (event.Subscription, error) {

	logs, sub, err := _WithdrawInbox.contract.WatchLogs(opts, "WithdrawalRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawInboxWithdrawalRequest)
				if err := _WithdrawInbox.contract.UnpackLog(event, "WithdrawalRequest", log); err != nil {
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

// ParseWithdrawalRequest is a log parse operation binding the contract event 0x7e2b24139224d852dd26bdb9f06f8136f7a1c9227a386d815a4ed8f1b8d7cc95.
//
// Solidity: event WithdrawalRequest(uint64 seqNum, address sender, address receiver, uint64 toChain, uint64[] fromChains, address[] tokens, uint32[] ratios, uint32[] slippages, uint256 deadline)
func (_WithdrawInbox *WithdrawInboxFilterer) ParseWithdrawalRequest(log types.Log) (*WithdrawInboxWithdrawalRequest, error) {
	event := new(WithdrawInboxWithdrawalRequest)
	if err := _WithdrawInbox.contract.UnpackLog(event, "WithdrawalRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
