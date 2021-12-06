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

// IncentiveEventsRewardMetaData contains all meta data concerning the IncentiveEventsReward contract.
var IncentiveEventsRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_celerTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"IncentiveRewardClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contribution\",\"type\":\"uint256\"}],\"name\":\"IncentiveRewardContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CELER_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimDeadlines\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_sig\",\"type\":\"bytes\"}],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimedRewardAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"contributeToRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"drainToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eventId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"setClaimDeadline\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161163538038061163583398101604081905261002f916100ac565b6000805460ff1916905561004233610053565b6001600160a01b03166080526100dc565b600080546001600160a01b03838116610100818102610100600160a81b0319851617855560405193049190911692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35050565b6000602082840312156100be57600080fd5b81516001600160a01b03811681146100d557600080fd5b9392505050565b60805161152961010c600039600081816101ef015281816102bb015281816103e9015261075d01526115296000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636c19e78311610097578063960dc08a11610066578063960dc08a146101ea578063a5e2c49014610211578063b8d2ce7f1461023c578063f2fde38b1461024f57600080fd5b80636c19e783146101b1578063715018a6146101c45780638456cb59146101cc5780638da5cb5b146101d457600080fd5b80632b5f3ece116100d35780632b5f3ece146101525780633f4ba83a146101805780635c975abb1461018857806361519f421461019e57600080fd5b80630a300b09146100fa578063145aa1161461010f578063238ac93314610122575b600080fd5b61010d6101083660046112e9565b610262565b005b61010d61011d3660046112e9565b61032a565b600154610135906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6101726101603660046112e9565b60036020526000908152604090205481565b604051908152602001610149565b61010d610413565b60005460ff166040519015158152602001610149565b61010d6101ac36600461131e565b61047d565b61010d6101bf3660046113b2565b6107d2565b61010d61086c565b61010d6108d6565b60005461010090046001600160a01b0316610135565b6101357f000000000000000000000000000000000000000000000000000000000000000081565b61017261021f3660046113cd565b600260209081526000928352604080842090915290825290205481565b61010d61024a3660046113f9565b61093e565b61010d61025d3660046113b2565b6109b0565b60005460ff16156102ad5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064015b60405180910390fd5b336102e36001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016823085610a95565b806001600160a01b03167fae26f78c6f6f1b85c2d212321268551fd9253c29c002447a2d5fd92743134e838360405161031e91815260200190565b60405180910390a25050565b60005460ff1661037c5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102a4565b6000546001600160a01b036101009091041633146103dc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a4565b6104106001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163383610b33565b50565b6000546001600160a01b036101009091041633146104735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a4565b61047b610b68565b565b60005460ff16156104c35760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102a4565b6000848152600360205260409020548061051f5760405162461bcd60e51b815260206004820152600f60248201527f496e76616c6964206576656e744964000000000000000000000000000000000060448201526064016102a4565b8042111561056f5760405162461bcd60e51b815260206004820152600d60248201527f436c61696d20657870697265640000000000000000000000000000000000000060448201526064016102a4565b6040516bffffffffffffffffffffffff19606088901b166020820152603481018690526054810185905260009061060c90607401604051602081830303815290604052805190602001206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050600061065285858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508693925050610c049050565b6001549091506001600160a01b038083169116146106b25760405162461bcd60e51b815260206004820152600b60248201527f496e76616c69642073696700000000000000000000000000000000000000000060448201526064016102a4565b60008781526002602090815260408083206001600160a01b038c1684529091528120546106df908861141b565b9050600081116107315760405162461bcd60e51b815260206004820152600d60248201527f4e6f206e6577207265776172640000000000000000000000000000000000000060448201526064016102a4565b60008881526002602090815260408083206001600160a01b03808e1685529252909120889055610784907f0000000000000000000000000000000000000000000000000000000000000000168a83610b33565b886001600160a01b03167f249ef50521b37da6ff1df71d2ffd77cba3ab9bb155803daa7ae14a8870f532d4826040516107bf91815260200190565b60405180910390a2505050505050505050565b6000546001600160a01b036101009091041633146108325760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a4565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6000546001600160a01b036101009091041633146108cc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a4565b61047b6000610c28565b6000546001600160a01b036101009091041633146109365760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a4565b61047b610c98565b6000546001600160a01b0361010090910416331461099e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a4565b60009182526003602052604090912055565b6000546001600160a01b03610100909104163314610a105760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a4565b6001600160a01b038116610a8c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102a4565b61041081610c28565b6040516001600160a01b0380851660248301528316604482015260648101829052610b2d9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610d13565b50505050565b6040516001600160a01b038316602482015260448101829052610b6390849063a9059cbb60e01b90606401610ac9565b505050565b60005460ff16610bba5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016102a4565b6000805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6000806000610c138585610df8565b91509150610c2081610e68565b509392505050565b600080546001600160a01b038381166101008181027fffffffffffffffffffffff0000000000000000000000000000000000000000ff851617855560405193049190911692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a35050565b60005460ff1615610cde5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016102a4565b6000805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610be73390565b6000610d68826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110239092919063ffffffff16565b805190915015610b635780806020019051810190610d869190611440565b610b635760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102a4565b600080825160411415610e2f5760208301516040840151606085015160001a610e238782858561103c565b94509450505050610e61565b825160401415610e595760208301516040840151610e4e868383611129565b935093505050610e61565b506000905060025b9250929050565b6000816004811115610e7c57610e7c611462565b1415610e855750565b6001816004811115610e9957610e99611462565b1415610ee75760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016102a4565b6002816004811115610efb57610efb611462565b1415610f495760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016102a4565b6003816004811115610f5d57610f5d611462565b1415610fb65760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016102a4565b6004816004811115610fca57610fca611462565b14156104105760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016102a4565b60606110328484600085611171565b90505b9392505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156110735750600090506003611120565b8460ff16601b1415801561108b57508460ff16601c14155b1561109c5750600090506004611120565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156110f0573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661111957600060019250925050611120565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831660ff84901c601b016111638782888561103c565b935093505050935093915050565b6060824710156111e95760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102a4565b843b6112375760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102a4565b600080866001600160a01b0316858760405161125391906114a4565b60006040518083038185875af1925050503d8060008114611290576040519150601f19603f3d011682016040523d82523d6000602084013e611295565b606091505b50915091506112a58282866112b0565b979650505050505050565b606083156112bf575081611035565b8251156112cf5782518084602001fd5b8160405162461bcd60e51b81526004016102a491906114c0565b6000602082840312156112fb57600080fd5b5035919050565b80356001600160a01b038116811461131957600080fd5b919050565b60008060008060006080868803121561133657600080fd5b61133f86611302565b94506020860135935060408601359250606086013567ffffffffffffffff8082111561136a57600080fd5b818801915088601f83011261137e57600080fd5b81358181111561138d57600080fd5b89602082850101111561139f57600080fd5b9699959850939650602001949392505050565b6000602082840312156113c457600080fd5b61103582611302565b600080604083850312156113e057600080fd5b823591506113f060208401611302565b90509250929050565b6000806040838503121561140c57600080fd5b50508035926020909101359150565b60008282101561143b57634e487b7160e01b600052601160045260246000fd5b500390565b60006020828403121561145257600080fd5b8151801515811461103557600080fd5b634e487b7160e01b600052602160045260246000fd5b60005b8381101561149357818101518382015260200161147b565b83811115610b2d5750506000910152565b600082516114b6818460208701611478565b9190910192915050565b60208152600082518060208401526114df816040850160208701611478565b601f01601f1916919091016040019291505056fea26469706673582212201a963febe1385e6a079a72ca6176ebe4164edbc697b0035690c4a96e2126a56a64736f6c634300080a0033",
}

// IncentiveEventsRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use IncentiveEventsRewardMetaData.ABI instead.
var IncentiveEventsRewardABI = IncentiveEventsRewardMetaData.ABI

// IncentiveEventsRewardBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IncentiveEventsRewardMetaData.Bin instead.
var IncentiveEventsRewardBin = IncentiveEventsRewardMetaData.Bin

// DeployIncentiveEventsReward deploys a new Ethereum contract, binding an instance of IncentiveEventsReward to it.
func DeployIncentiveEventsReward(auth *bind.TransactOpts, backend bind.ContractBackend, _celerTokenAddress common.Address) (common.Address, *types.Transaction, *IncentiveEventsReward, error) {
	parsed, err := IncentiveEventsRewardMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IncentiveEventsRewardBin), backend, _celerTokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IncentiveEventsReward{IncentiveEventsRewardCaller: IncentiveEventsRewardCaller{contract: contract}, IncentiveEventsRewardTransactor: IncentiveEventsRewardTransactor{contract: contract}, IncentiveEventsRewardFilterer: IncentiveEventsRewardFilterer{contract: contract}}, nil
}

// IncentiveEventsReward is an auto generated Go binding around an Ethereum contract.
type IncentiveEventsReward struct {
	IncentiveEventsRewardCaller     // Read-only binding to the contract
	IncentiveEventsRewardTransactor // Write-only binding to the contract
	IncentiveEventsRewardFilterer   // Log filterer for contract events
}

// IncentiveEventsRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncentiveEventsRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentiveEventsRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncentiveEventsRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentiveEventsRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncentiveEventsRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncentiveEventsRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncentiveEventsRewardSession struct {
	Contract     *IncentiveEventsReward // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IncentiveEventsRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncentiveEventsRewardCallerSession struct {
	Contract *IncentiveEventsRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IncentiveEventsRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncentiveEventsRewardTransactorSession struct {
	Contract     *IncentiveEventsRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IncentiveEventsRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncentiveEventsRewardRaw struct {
	Contract *IncentiveEventsReward // Generic contract binding to access the raw methods on
}

// IncentiveEventsRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncentiveEventsRewardCallerRaw struct {
	Contract *IncentiveEventsRewardCaller // Generic read-only contract binding to access the raw methods on
}

// IncentiveEventsRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncentiveEventsRewardTransactorRaw struct {
	Contract *IncentiveEventsRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncentiveEventsReward creates a new instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsReward(address common.Address, backend bind.ContractBackend) (*IncentiveEventsReward, error) {
	contract, err := bindIncentiveEventsReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsReward{IncentiveEventsRewardCaller: IncentiveEventsRewardCaller{contract: contract}, IncentiveEventsRewardTransactor: IncentiveEventsRewardTransactor{contract: contract}, IncentiveEventsRewardFilterer: IncentiveEventsRewardFilterer{contract: contract}}, nil
}

// NewIncentiveEventsRewardCaller creates a new read-only instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsRewardCaller(address common.Address, caller bind.ContractCaller) (*IncentiveEventsRewardCaller, error) {
	contract, err := bindIncentiveEventsReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardCaller{contract: contract}, nil
}

// NewIncentiveEventsRewardTransactor creates a new write-only instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*IncentiveEventsRewardTransactor, error) {
	contract, err := bindIncentiveEventsReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardTransactor{contract: contract}, nil
}

// NewIncentiveEventsRewardFilterer creates a new log filterer instance of IncentiveEventsReward, bound to a specific deployed contract.
func NewIncentiveEventsRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*IncentiveEventsRewardFilterer, error) {
	contract, err := bindIncentiveEventsReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardFilterer{contract: contract}, nil
}

// bindIncentiveEventsReward binds a generic wrapper to an already deployed contract.
func bindIncentiveEventsReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncentiveEventsRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncentiveEventsReward *IncentiveEventsRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncentiveEventsReward.Contract.IncentiveEventsRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncentiveEventsReward *IncentiveEventsRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.IncentiveEventsRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncentiveEventsReward *IncentiveEventsRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.IncentiveEventsRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IncentiveEventsReward *IncentiveEventsRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IncentiveEventsReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.contract.Transact(opts, method, params...)
}

// CELERTOKEN is a free data retrieval call binding the contract method 0x960dc08a.
//
// Solidity: function CELER_TOKEN() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) CELERTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "CELER_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CELERTOKEN is a free data retrieval call binding the contract method 0x960dc08a.
//
// Solidity: function CELER_TOKEN() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) CELERTOKEN() (common.Address, error) {
	return _IncentiveEventsReward.Contract.CELERTOKEN(&_IncentiveEventsReward.CallOpts)
}

// CELERTOKEN is a free data retrieval call binding the contract method 0x960dc08a.
//
// Solidity: function CELER_TOKEN() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) CELERTOKEN() (common.Address, error) {
	return _IncentiveEventsReward.Contract.CELERTOKEN(&_IncentiveEventsReward.CallOpts)
}

// ClaimDeadlines is a free data retrieval call binding the contract method 0x2b5f3ece.
//
// Solidity: function claimDeadlines(uint256 ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) ClaimDeadlines(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "claimDeadlines", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimDeadlines is a free data retrieval call binding the contract method 0x2b5f3ece.
//
// Solidity: function claimDeadlines(uint256 ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ClaimDeadlines(arg0 *big.Int) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimDeadlines(&_IncentiveEventsReward.CallOpts, arg0)
}

// ClaimDeadlines is a free data retrieval call binding the contract method 0x2b5f3ece.
//
// Solidity: function claimDeadlines(uint256 ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) ClaimDeadlines(arg0 *big.Int) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimDeadlines(&_IncentiveEventsReward.CallOpts, arg0)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0xa5e2c490.
//
// Solidity: function claimedRewardAmounts(uint256 , address ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) ClaimedRewardAmounts(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "claimedRewardAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0xa5e2c490.
//
// Solidity: function claimedRewardAmounts(uint256 , address ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ClaimedRewardAmounts(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimedRewardAmounts(&_IncentiveEventsReward.CallOpts, arg0, arg1)
}

// ClaimedRewardAmounts is a free data retrieval call binding the contract method 0xa5e2c490.
//
// Solidity: function claimedRewardAmounts(uint256 , address ) view returns(uint256)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) ClaimedRewardAmounts(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _IncentiveEventsReward.Contract.ClaimedRewardAmounts(&_IncentiveEventsReward.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Owner() (common.Address, error) {
	return _IncentiveEventsReward.Contract.Owner(&_IncentiveEventsReward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) Owner() (common.Address, error) {
	return _IncentiveEventsReward.Contract.Owner(&_IncentiveEventsReward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Paused() (bool, error) {
	return _IncentiveEventsReward.Contract.Paused(&_IncentiveEventsReward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) Paused() (bool, error) {
	return _IncentiveEventsReward.Contract.Paused(&_IncentiveEventsReward.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IncentiveEventsReward.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Signer() (common.Address, error) {
	return _IncentiveEventsReward.Contract.Signer(&_IncentiveEventsReward.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_IncentiveEventsReward *IncentiveEventsRewardCallerSession) Signer() (common.Address, error) {
	return _IncentiveEventsReward.Contract.Signer(&_IncentiveEventsReward.CallOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x61519f42.
//
// Solidity: function claimReward(address _recipient, uint256 _eventId, uint256 _rewardAmount, bytes _sig) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) ClaimReward(opts *bind.TransactOpts, _recipient common.Address, _eventId *big.Int, _rewardAmount *big.Int, _sig []byte) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "claimReward", _recipient, _eventId, _rewardAmount, _sig)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x61519f42.
//
// Solidity: function claimReward(address _recipient, uint256 _eventId, uint256 _rewardAmount, bytes _sig) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ClaimReward(_recipient common.Address, _eventId *big.Int, _rewardAmount *big.Int, _sig []byte) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ClaimReward(&_IncentiveEventsReward.TransactOpts, _recipient, _eventId, _rewardAmount, _sig)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x61519f42.
//
// Solidity: function claimReward(address _recipient, uint256 _eventId, uint256 _rewardAmount, bytes _sig) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) ClaimReward(_recipient common.Address, _eventId *big.Int, _rewardAmount *big.Int, _sig []byte) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ClaimReward(&_IncentiveEventsReward.TransactOpts, _recipient, _eventId, _rewardAmount, _sig)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) ContributeToRewardPool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "contributeToRewardPool", _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) ContributeToRewardPool(_amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ContributeToRewardPool(&_IncentiveEventsReward.TransactOpts, _amount)
}

// ContributeToRewardPool is a paid mutator transaction binding the contract method 0x0a300b09.
//
// Solidity: function contributeToRewardPool(uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) ContributeToRewardPool(_amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.ContributeToRewardPool(&_IncentiveEventsReward.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) DrainToken(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "drainToken", _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.DrainToken(&_IncentiveEventsReward.TransactOpts, _amount)
}

// DrainToken is a paid mutator transaction binding the contract method 0x145aa116.
//
// Solidity: function drainToken(uint256 _amount) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) DrainToken(_amount *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.DrainToken(&_IncentiveEventsReward.TransactOpts, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Pause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Pause(&_IncentiveEventsReward.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) Pause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Pause(&_IncentiveEventsReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) RenounceOwnership() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.RenounceOwnership(&_IncentiveEventsReward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.RenounceOwnership(&_IncentiveEventsReward.TransactOpts)
}

// SetClaimDeadline is a paid mutator transaction binding the contract method 0xb8d2ce7f.
//
// Solidity: function setClaimDeadline(uint256 _eventId, uint256 _deadline) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) SetClaimDeadline(opts *bind.TransactOpts, _eventId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "setClaimDeadline", _eventId, _deadline)
}

// SetClaimDeadline is a paid mutator transaction binding the contract method 0xb8d2ce7f.
//
// Solidity: function setClaimDeadline(uint256 _eventId, uint256 _deadline) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) SetClaimDeadline(_eventId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetClaimDeadline(&_IncentiveEventsReward.TransactOpts, _eventId, _deadline)
}

// SetClaimDeadline is a paid mutator transaction binding the contract method 0xb8d2ce7f.
//
// Solidity: function setClaimDeadline(uint256 _eventId, uint256 _deadline) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) SetClaimDeadline(_eventId *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetClaimDeadline(&_IncentiveEventsReward.TransactOpts, _eventId, _deadline)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) SetSigner(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "setSigner", _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetSigner(&_IncentiveEventsReward.TransactOpts, _signer)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _signer) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) SetSigner(_signer common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.SetSigner(&_IncentiveEventsReward.TransactOpts, _signer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.TransferOwnership(&_IncentiveEventsReward.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.TransferOwnership(&_IncentiveEventsReward.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IncentiveEventsReward.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardSession) Unpause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Unpause(&_IncentiveEventsReward.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IncentiveEventsReward *IncentiveEventsRewardTransactorSession) Unpause() (*types.Transaction, error) {
	return _IncentiveEventsReward.Contract.Unpause(&_IncentiveEventsReward.TransactOpts)
}

// IncentiveEventsRewardIncentiveRewardClaimedIterator is returned from FilterIncentiveRewardClaimed and is used to iterate over the raw logs and unpacked data for IncentiveRewardClaimed events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardClaimedIterator struct {
	Event *IncentiveEventsRewardIncentiveRewardClaimed // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardIncentiveRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardIncentiveRewardClaimed)
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
		it.Event = new(IncentiveEventsRewardIncentiveRewardClaimed)
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
func (it *IncentiveEventsRewardIncentiveRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardIncentiveRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardIncentiveRewardClaimed represents a IncentiveRewardClaimed event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardClaimed struct {
	Recipient common.Address
	Reward    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncentiveRewardClaimed is a free log retrieval operation binding the contract event 0x249ef50521b37da6ff1df71d2ffd77cba3ab9bb155803daa7ae14a8870f532d4.
//
// Solidity: event IncentiveRewardClaimed(address indexed recipient, uint256 reward)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterIncentiveRewardClaimed(opts *bind.FilterOpts, recipient []common.Address) (*IncentiveEventsRewardIncentiveRewardClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "IncentiveRewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardIncentiveRewardClaimedIterator{contract: _IncentiveEventsReward.contract, event: "IncentiveRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchIncentiveRewardClaimed is a free log subscription operation binding the contract event 0x249ef50521b37da6ff1df71d2ffd77cba3ab9bb155803daa7ae14a8870f532d4.
//
// Solidity: event IncentiveRewardClaimed(address indexed recipient, uint256 reward)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchIncentiveRewardClaimed(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardIncentiveRewardClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "IncentiveRewardClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardIncentiveRewardClaimed)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardClaimed", log); err != nil {
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

// ParseIncentiveRewardClaimed is a log parse operation binding the contract event 0x249ef50521b37da6ff1df71d2ffd77cba3ab9bb155803daa7ae14a8870f532d4.
//
// Solidity: event IncentiveRewardClaimed(address indexed recipient, uint256 reward)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseIncentiveRewardClaimed(log types.Log) (*IncentiveEventsRewardIncentiveRewardClaimed, error) {
	event := new(IncentiveEventsRewardIncentiveRewardClaimed)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardIncentiveRewardContributedIterator is returned from FilterIncentiveRewardContributed and is used to iterate over the raw logs and unpacked data for IncentiveRewardContributed events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardContributedIterator struct {
	Event *IncentiveEventsRewardIncentiveRewardContributed // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardIncentiveRewardContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardIncentiveRewardContributed)
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
		it.Event = new(IncentiveEventsRewardIncentiveRewardContributed)
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
func (it *IncentiveEventsRewardIncentiveRewardContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardIncentiveRewardContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardIncentiveRewardContributed represents a IncentiveRewardContributed event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardIncentiveRewardContributed struct {
	Contributor  common.Address
	Contribution *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIncentiveRewardContributed is a free log retrieval operation binding the contract event 0xae26f78c6f6f1b85c2d212321268551fd9253c29c002447a2d5fd92743134e83.
//
// Solidity: event IncentiveRewardContributed(address indexed contributor, uint256 contribution)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterIncentiveRewardContributed(opts *bind.FilterOpts, contributor []common.Address) (*IncentiveEventsRewardIncentiveRewardContributedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "IncentiveRewardContributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardIncentiveRewardContributedIterator{contract: _IncentiveEventsReward.contract, event: "IncentiveRewardContributed", logs: logs, sub: sub}, nil
}

// WatchIncentiveRewardContributed is a free log subscription operation binding the contract event 0xae26f78c6f6f1b85c2d212321268551fd9253c29c002447a2d5fd92743134e83.
//
// Solidity: event IncentiveRewardContributed(address indexed contributor, uint256 contribution)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchIncentiveRewardContributed(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardIncentiveRewardContributed, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "IncentiveRewardContributed", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardIncentiveRewardContributed)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardContributed", log); err != nil {
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

// ParseIncentiveRewardContributed is a log parse operation binding the contract event 0xae26f78c6f6f1b85c2d212321268551fd9253c29c002447a2d5fd92743134e83.
//
// Solidity: event IncentiveRewardContributed(address indexed contributor, uint256 contribution)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseIncentiveRewardContributed(log types.Log) (*IncentiveEventsRewardIncentiveRewardContributed, error) {
	event := new(IncentiveEventsRewardIncentiveRewardContributed)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "IncentiveRewardContributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardOwnershipTransferredIterator struct {
	Event *IncentiveEventsRewardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardOwnershipTransferred)
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
		it.Event = new(IncentiveEventsRewardOwnershipTransferred)
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
func (it *IncentiveEventsRewardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardOwnershipTransferred represents a OwnershipTransferred event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IncentiveEventsRewardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardOwnershipTransferredIterator{contract: _IncentiveEventsReward.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardOwnershipTransferred)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseOwnershipTransferred(log types.Log) (*IncentiveEventsRewardOwnershipTransferred, error) {
	event := new(IncentiveEventsRewardOwnershipTransferred)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardPausedIterator struct {
	Event *IncentiveEventsRewardPaused // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardPaused)
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
		it.Event = new(IncentiveEventsRewardPaused)
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
func (it *IncentiveEventsRewardPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardPaused represents a Paused event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterPaused(opts *bind.FilterOpts) (*IncentiveEventsRewardPausedIterator, error) {

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardPausedIterator{contract: _IncentiveEventsReward.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardPaused) (event.Subscription, error) {

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardPaused)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParsePaused(log types.Log) (*IncentiveEventsRewardPaused, error) {
	event := new(IncentiveEventsRewardPaused)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncentiveEventsRewardUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardUnpausedIterator struct {
	Event *IncentiveEventsRewardUnpaused // Event containing the contract specifics and raw log

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
func (it *IncentiveEventsRewardUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IncentiveEventsRewardUnpaused)
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
		it.Event = new(IncentiveEventsRewardUnpaused)
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
func (it *IncentiveEventsRewardUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IncentiveEventsRewardUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IncentiveEventsRewardUnpaused represents a Unpaused event raised by the IncentiveEventsReward contract.
type IncentiveEventsRewardUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IncentiveEventsRewardUnpausedIterator, error) {

	logs, sub, err := _IncentiveEventsReward.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IncentiveEventsRewardUnpausedIterator{contract: _IncentiveEventsReward.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IncentiveEventsRewardUnpaused) (event.Subscription, error) {

	logs, sub, err := _IncentiveEventsReward.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IncentiveEventsRewardUnpaused)
				if err := _IncentiveEventsReward.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IncentiveEventsReward *IncentiveEventsRewardFilterer) ParseUnpaused(log types.Log) (*IncentiveEventsRewardUnpaused, error) {
	event := new(IncentiveEventsRewardUnpaused)
	if err := _IncentiveEventsReward.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
