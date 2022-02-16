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

// IMessageBusRouteInfo is an auto generated low-level Go binding around an user-defined struct.
type IMessageBusRouteInfo struct {
	Sender     common.Address
	Receiver   common.Address
	SrcChainId uint64
}

// IMessageBusTransferInfo is an auto generated low-level Go binding around an user-defined struct.
type IMessageBusTransferInfo struct {
	T          uint8
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Seqnum     uint64
	SrcChainId uint64
	RefId      [32]byte
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

// BatchTransferMetaData contains all meta data concerning the BatchTransfer contract.
var BatchTransferMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"enumMessageSenderLib.BridgeType\",\"name\":\"_bridgeType\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"enumBatchTransfer.TransferStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200248c3803806200248c8339810160408190526200003491620000b5565b6200003f3362000065565b600180546001600160a01b0319166001600160a01b0392909216919091179055620000e7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215620000c857600080fd5b81516001600160a01b0381168114620000e057600080fd5b9392505050565b61239580620000f76000396000f3fe6080604052600436106100b15760003560e01c80638da5cb5b11610069578063ce35dd9a1161004e578063ce35dd9a146101bd578063f00f39ce146101d0578063f2fde38b146101e357600080fd5b80638da5cb5b1461016b578063a1a227fa1461019d57600080fd5b806320bff8931161009a57806320bff893146100f1578063547cad121461013657806364d39f421461015857600080fd5b80631599d265146100b657806320be95f2146100de575b600080fd5b6100c96100c43660046119b9565b610203565b60405190151581526020015b60405180910390f35b6100c96100ec366004611a1b565b610384565b3480156100fd57600080fd5b5061012861010c366004611aa4565b6002602052600090815260409020805460019091015460ff1682565b6040516100d5929190611af9565b34801561014257600080fd5b50610156610151366004611b0d565b610418565b005b610156610166366004611b85565b6104b0565b34801561017757600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016100d5565b3480156101a957600080fd5b50600154610185906001600160a01b031681565b6100c96101cb366004611c66565b610752565b6100c96101de366004611c66565b6108fe565b3480156101ef57600080fd5b506101566101fe366004611b0d565b6109ec565b6001546000906001600160a01b031633146102655760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b60008280602001905181019061027b9190611ce5565b6040516bffffffffffffffffffffffff19606088901b1660208201526001600160c01b031960c087901b166034820152909150603c0160408051601f198184030181529181528151602092830120835167ffffffffffffffff16600090815260029093529120541461032f5760405162461bcd60e51b815260206004820152600f60248201527f696e76616c6964206d6573736167650000000000000000000000000000000000604482015260640161025c565b602080820151825167ffffffffffffffff16600090815260029283905260409020600190810180549293909260ff19169190849081111561037257610372611ac1565b021790555060019150505b9392505050565b6001546000906001600160a01b031633146103e15760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161025c565b60006103ef83850185611dd7565b606081015190915061040c906001600160a01b0388169087610add565b50600195945050505050565b3361042b6000546001600160a01b031690565b6001600160a01b0316146104815760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161025c565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b3332146104ff5760405162461bcd60e51b815260206004820152600760248201527f4e6f7420454f4100000000000000000000000000000000000000000000000000604482015260640161025c565b6000805b828110156105435783838281811061051d5761051d611eea565b905060200201358261052f9190611f16565b91508061053b81611f2e565b915050610503565b50600180548190601490610569908390600160a01b900467ffffffffffffffff16611f49565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060405180604001604052808c8a6040516020016105d792919060609290921b6bffffffffffffffffffffffff1916825260c01b6001600160c01b0319166014820152601c0190565b60408051601f1981840301815291905280516020918201208252016000905260018054600160a01b900467ffffffffffffffff1660009081526002602081815260409092208451815591840151828401805493949193909260ff199091169190849081111561064857610648611ac1565b02179055506106659150506001600160a01b038b1633308c610b72565b60408051608081018252600154600160a01b900467ffffffffffffffff16815281516020878102808301820190945287825260009381840192918a918a9182919085019084908082843760009201919091525050509082525060408051602087810282810182019093528782529283019290918891889182918501908490808284376000920191909152505050908252503360209182015260405161070b929101611fb0565b60405160208183030381529060405290506107438c8c8c8c600160149054906101000a900467ffffffffffffffff168d878e34610bb0565b50505050505050505050505050565b6001546000906001600160a01b031633146107af5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161025c565b6000828060200190518101906107c591906120bc565b90506000805b82602001515181101561086d57610830836020015182815181106107f1576107f1611eea565b60200260200101518460400151838151811061080f5761080f611eea565b60200260200101518a6001600160a01b0316610add9092919063ffffffff16565b8260400151818151811061084657610846611eea565b6020026020010151826108599190611f16565b91508061086581611f2e565b9150506107cb565b50600061087a82886121c0565b90508187111561089e57606083015161089e906001600160a01b038a169083610add565b60408051808201909152835167ffffffffffffffff16815260009060208101600190526040516108d191906020016121d7565b60405160208183030381529060405290506108ee8a888334610be7565b5060019998505050505050505050565b6001546000906001600160a01b0316331461095b5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161025c565b60008280602001905181019061097191906120bc565b606081015190915061098e906001600160a01b0388169087610add565b60408051808201909152815167ffffffffffffffff16815260009060208101600290526040516109c191906020016121d7565b60405160208183030381529060405290506109de88868334610be7565b506001979650505050505050565b336109ff6000546001600160a01b031690565b6001600160a01b031614610a555760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161025c565b6001600160a01b038116610ad15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161025c565b610ada81610c03565b50565b6040516001600160a01b038316602482015260448101829052610b6d90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610c60565b505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610baa9085906323b872dd60e01b90608401610b09565b50505050565b6000610bd98a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b610d45565b9a9950505050505050505050565b600154610baa908590859085906001600160a01b031685610e4f565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610cb5826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610eba9092919063ffffffff16565b805190915015610b6d5780806020019051810190610cd39190612203565b610b6d5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161025c565b60006001846005811115610d5b57610d5b611ac1565b1415610d7957610d728b8b8b8b8b8b8b8a8a610ed1565b9050610bd9565b6002846005811115610d8d57610d8d611ac1565b1480610daa57506004846005811115610da857610da8611ac1565b145b15610dc057610d72848c8c8c8c8c8b8a8a6110df565b6003846005811115610dd457610dd4611ac1565b1480610df157506005846005811115610def57610def611ac1565b145b15610e0757610d72848c8c8c8c8c8b8a8a611411565b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f7274656400000000000000604482015260640161025c565b604051634f9e72ad60e11b81526001600160a01b03831690639f3ce55a908390610e819089908990899060040161227d565b6000604051808303818588803b158015610e9a57600080fd5b505af1158015610eae573d6000803e3d6000fd5b50505050505050505050565b6060610ec98484600085611665565b949350505050565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b158015610f0d57600080fd5b505afa158015610f21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4591906122b8565b9050610f5b6001600160a01b038b16828b6117a4565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b158015610fd157600080fd5b505af1158015610fe5573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b815260040161109d9594939291906122d5565b6000604051808303818588803b1580156110b657600080fd5b505af11580156110ca573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b60008060028b60058111156110f6576110f6611ac1565b141561117457836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b15801561113557600080fd5b505afa158015611149573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061116d91906122b8565b90506111e8565b836001600160a01b031663c66a9c5a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111ad57600080fd5b505afa1580156111c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e591906122b8565b90505b6111fc6001600160a01b038a16828a6117a4565b600060028c600581111561121257611212611ac1565b1415611335576040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a401600060405180830381600087803b15801561128257600080fd5b505af1158015611296573d6000803e3d6000fd5b50505050308a8a8a8e8b466040516020016113189796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b6040516020818303038152906040528051906020012090506113db565b6040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a4015b602060405180830381600087803b1580156113a057600080fd5b505af11580156113b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d89190612317565b90505b604051634289fbb360e01b81526001600160a01b03861690634289fbb390869061109d908f908d90889088908e906004016122d5565b60008060038b600581111561142857611428611ac1565b14156114a657836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561146757600080fd5b505afa15801561147b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061149f91906122b8565b905061151a565b836001600160a01b03166395b12c276040518163ffffffff1660e01b815260040160206040518083038186803b1580156114df57600080fd5b505afa1580156114f3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061151791906122b8565b90505b600060038c600581111561153057611530611ac1565b141561161057604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff8916606483015283169063de790c7e90608401600060405180830381600087803b15801561159857600080fd5b505af11580156115ac573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528e811b82166034840152604883018e90528f901b1660688201526001600160c01b031960c08b811b8216607c84015246901b166084820152608c0191506113189050565b60405163a002930160e01b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063a00293019060a401611386565b6060824710156116dd5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161025c565b843b61172b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161025c565b600080866001600160a01b031685876040516117479190612330565b60006040518083038185875af1925050503d8060008114611784576040519150601f19603f3d011682016040523d82523d6000602084013e611789565b606091505b5091509150611799828286611865565b979650505050505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b1580156117f057600080fd5b505afa158015611804573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118289190612317565b6118329190611f16565b6040516001600160a01b038516602482015260448101829052909150610baa90859063095ea7b360e01b90606401610b09565b6060831561187457508161037d565b8251156118845782518084602001fd5b8160405162461bcd60e51b815260040161025c919061234c565b6001600160a01b0381168114610ada57600080fd5b80356118be8161189e565b919050565b67ffffffffffffffff81168114610ada57600080fd5b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715611912576119126118d9565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715611941576119416118d9565b604052919050565b600082601f83011261195a57600080fd5b813567ffffffffffffffff811115611974576119746118d9565b611987601f8201601f1916602001611918565b81815284602083860101111561199c57600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000606084860312156119ce57600080fd5b83356119d98161189e565b925060208401356119e9816118c3565b9150604084013567ffffffffffffffff811115611a0557600080fd5b611a1186828701611949565b9150509250925092565b60008060008060608587031215611a3157600080fd5b8435611a3c8161189e565b935060208501359250604085013567ffffffffffffffff80821115611a6057600080fd5b818701915087601f830112611a7457600080fd5b813581811115611a8357600080fd5b886020828501011115611a9557600080fd5b95989497505060200194505050565b600060208284031215611ab657600080fd5b813561037d816118c3565b634e487b7160e01b600052602160045260246000fd5b60038110611af557634e487b7160e01b600052602160045260246000fd5b9052565b8281526040810161037d6020830184611ad7565b600060208284031215611b1f57600080fd5b813561037d8161189e565b8035600681106118be57600080fd5b60008083601f840112611b4b57600080fd5b50813567ffffffffffffffff811115611b6357600080fd5b6020830191508360208260051b8501011115611b7e57600080fd5b9250929050565b6000806000806000806000806000806101008b8d031215611ba557600080fd5b8a35611bb08161189e565b995060208b0135611bc08161189e565b985060408b0135975060608b0135611bd7816118c3565b965060808b013563ffffffff81168114611bf057600080fd5b9550611bfe60a08c01611b2a565b945060c08b013567ffffffffffffffff80821115611c1b57600080fd5b611c278e838f01611b39565b909650945060e08d0135915080821115611c4057600080fd5b50611c4d8d828e01611b39565b915080935050809150509295989b9194979a5092959850565b600080600080600060a08688031215611c7e57600080fd5b8535611c898161189e565b94506020860135611c998161189e565b9350604086013592506060860135611cb0816118c3565b9150608086013567ffffffffffffffff811115611ccc57600080fd5b611cd888828901611949565b9150509295509295909350565b600060408284031215611cf757600080fd5b6040516040810181811067ffffffffffffffff82111715611d1a57611d1a6118d9565b6040528251611d28816118c3565b8152602083015160038110611d3c57600080fd5b60208201529392505050565b600067ffffffffffffffff821115611d6257611d626118d9565b5060051b60200190565b600082601f830112611d7d57600080fd5b81356020611d92611d8d83611d48565b611918565b82815260059290921b84018101918181019086841115611db157600080fd5b8286015b84811015611dcc5780358352918301918301611db5565b509695505050505050565b60006020808385031215611dea57600080fd5b823567ffffffffffffffff80821115611e0257600080fd5b9084019060808287031215611e1657600080fd5b611e1e6118ef565b8235611e29816118c3565b81528284013582811115611e3c57600080fd5b8301601f81018813611e4d57600080fd5b8035611e5b611d8d82611d48565b81815260059190911b8201860190868101908a831115611e7a57600080fd5b928701925b82841015611ea1578335611e928161189e565b82529287019290870190611e7f565b8088860152505050506040830135935081841115611ebe57600080fd5b611eca87858501611d6c565b6040820152611edb606084016118b3565b60608201529695505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008219821115611f2957611f29611f00565b500190565b6000600019821415611f4257611f42611f00565b5060010190565b600067ffffffffffffffff808316818516808303821115611f6c57611f6c611f00565b01949350505050565b600081518084526020808501945080840160005b83811015611fa557815187529582019590820190600101611f89565b509495945050505050565b6020808252825167ffffffffffffffff16828201528281015160806040840152805160a0840181905260009291820190839060c08601905b808310156120115783516001600160a01b03168252928401926001929092019190840190611fe8565b506040870151868203601f19016060880152935061202f8185611f75565b9350505050606084015161204e60808501826001600160a01b03169052565b509392505050565b80516118be8161189e565b600082601f83011261207257600080fd5b81516020612082611d8d83611d48565b82815260059290921b840181019181810190868411156120a157600080fd5b8286015b84811015611dcc57805183529183019183016120a5565b600060208083850312156120cf57600080fd5b825167ffffffffffffffff808211156120e757600080fd5b90840190608082870312156120fb57600080fd5b6121036118ef565b825161210e816118c3565b8152828401518281111561212157600080fd5b8301601f8101881361213257600080fd5b8051612140611d8d82611d48565b81815260059190911b8201860190868101908a83111561215f57600080fd5b928701925b828410156121865783516121778161189e565b82529287019290870190612164565b80888601525050505060408301519350818411156121a357600080fd5b6121af87858501612061565b6040820152611edb60608401612056565b6000828210156121d2576121d2611f00565b500390565b815167ffffffffffffffff16815260208083015160408301916121fc90840182611ad7565b5092915050565b60006020828403121561221557600080fd5b8151801515811461037d57600080fd5b60005b83811015612240578181015183820152602001612228565b83811115610baa5750506000910152565b60008151808452612269816020860160208601612225565b601f01601f19169290920160200192915050565b6001600160a01b038416815267ffffffffffffffff831660208201526060604082015260006122af6060830184612251565b95945050505050565b6000602082840312156122ca57600080fd5b815161037d8161189e565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a0608083015261179960a0830184612251565b60006020828403121561232957600080fd5b5051919050565b60008251612342818460208701612225565b9190910192915050565b60208152600061037d602083018461225156fea2646970667358221220fd660c6ee46381bb70be9a1a1196ce4fa7e56f72ef3fff6302c86a2ac650560764736f6c63430008090033",
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
	Bin: "0x60806040523480156200001157600080fd5b506040516200232e3803806200232e8339810160408190526200003491620000b5565b6200003f3362000065565b600280546001600160a01b0319166001600160a01b0392909216919091179055620000e7565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215620000c857600080fd5b81516001600160a01b0381168114620000e057600080fd5b9392505050565b61223780620000f76000396000f3fe6080604052600436106100b15760003560e01c8063a1a227fa11610069578063f00f39ce1161004e578063f00f39ce14610197578063f2fde38b146101aa578063f88954dc146101ca57600080fd5b8063a1a227fa14610164578063ce35dd9a1461018457600080fd5b8063547cad121161009a578063547cad12146100ec578063692058c21461010e5780638da5cb5b1461014657600080fd5b80631599d265146100b657806320be95f2146100de575b600080fd5b6100c96100c4366004611b4d565b6101dd565b60405190151581526020015b60405180910390f35b6100c96100c4366004611bb0565b3480156100f857600080fd5b5061010c610107366004611bf4565b610247565b005b34801561011a57600080fd5b5060025461012e906001600160a01b031681565b6040516001600160a01b0390911681526020016100d5565b34801561015257600080fd5b506000546001600160a01b031661012e565b34801561017057600080fd5b5060015461012e906001600160a01b031681565b6100c9610192366004611c58565b6102df565b6100c96101a5366004611d2d565b610646565b3480156101b657600080fd5b5061010c6101c5366004611bf4565b6106ad565b61010c6101d8366004611db0565b61079e565b6001546000906001600160a01b0316331461023f5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b949350505050565b3361025a6000546001600160a01b031690565b6001600160a01b0316146102b05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610236565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6001546000906001600160a01b0316331461033c5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610236565b6000828060200190518101906103529190611e3f565b60025460405163095ea7b360e01b81526001600160a01b0391821660048201526024810188905291925087169063095ea7b390604401602060405180830381600087803b1580156103a257600080fd5b505af11580156103b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103da9190611ec4565b50604080516002808252606082018352600092602083019080368337019050509050868160008151811061041057610410611ee1565b60200260200101906001600160a01b031690816001600160a01b03168152505081600001518160018151811061044857610448611ee1565b60200260200101906001600160a01b031690816001600160a01b0316815250508160400151156105a1576001600260148282829054906101000a900467ffffffffffffffff166104989190611f0d565b825467ffffffffffffffff9182166101009390930a9283029190920219909116179055506002546040516338ed173960e01b81526000916001600160a01b0316906338ed1739906104f7908a9085908790309060001990600401611f39565b600060405180830381600087803b15801561051157600080fd5b505af1158015610525573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261054d9190810190611faa565b905061059b836020015184600001518360018151811061056f5761056f611ee1565b602002602001015189600260149054906101000a900467ffffffffffffffff1688606001516001610871565b50610638565b60025460208301516040516338ed173960e01b81526001600160a01b03909216916338ed1739916105e0918a9160009187919060001990600401611f39565b600060405180830381600087803b1580156105fa57600080fd5b505af115801561060e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106369190810190611faa565b505b506001979650505050505050565b6001546000906001600160a01b031633146106a35760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610236565b9695505050505050565b336106c06000546001600160a01b031690565b6001600160a01b0316146107165760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610236565b6001600160a01b0381166107925760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610236565b61079b81610b1c565b50565b6001600260148282829054906101000a900467ffffffffffffffff166107c49190611f0d565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555061080a333085876001600160a01b0316610b79909392919063ffffffff16565b60008160405160200161081d9190612050565b604051602081830303815290604052905061086886868686600260149054906101000a900467ffffffffffffffff1687606001602081019061085f91906120b7565b87600134610c17565b50505050505050565b60006001826005811115610887576108876120d4565b141561091a57600160009054906101000a90046001600160a01b03166001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b505afa1580156108ef573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091391906120ea565b9050610b02565b600282600581111561092e5761092e6120d4565b141561098257600160009054906101000a90046001600160a01b03166001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b6003826005811115610996576109966120d4565b14156109ea57600160009054906101000a90046001600160a01b03166001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b60048260058111156109fe576109fe6120d4565b1415610a5257600160009054906101000a90046001600160a01b03166001600160a01b031663c66a9c5a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b6005826005811115610a6657610a666120d4565b1415610aba57600160009054906101000a90046001600160a01b03166001600160a01b03166395b12c276040518163ffffffff1660e01b815260040160206040518083038186803b1580156108db57600080fd5b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f72746564000000000000006044820152606401610236565b610b128888888888888888610c4e565b5050505050505050565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6040516001600160a01b0380851660248301528316604482015260648101829052610c119085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610f40565b50505050565b6000610c408a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b61102a565b9a9950505050505050505050565b6001826005811115610c6257610c626120d4565b1415610d1057610c7c6001600160a01b03881682886110ec565b60405163a5977fbb60e01b81526001600160a01b03898116600483015288811660248301526044820188905267ffffffffffffffff80881660648401528616608483015263ffffffff851660a483015282169063a5977fbb9060c4015b600060405180830381600087803b158015610cf357600080fd5b505af1158015610d07573d6000803e3d6000fd5b50505050610b12565b6002826005811115610d2457610d246120d4565b1415610d9357610d3e6001600160a01b03881682886110ec565b6040516308d18d8960e21b81526001600160a01b0388811660048301526024820188905267ffffffffffffffff80881660448401528a821660648401528616608483015282169063234636249060a401610cd9565b6003826005811115610da757610da76120d4565b1415610dfa57604051636f3c863f60e11b81526001600160a01b03888116600483015260248201889052898116604483015267ffffffffffffffff8616606483015282169063de790c7e90608401610cd9565b6004826005811115610e0e57610e0e6120d4565b1415610ed157610e286001600160a01b03881682886110ec565b6040516308d18d8960e21b81526001600160a01b0388811660048301526024820188905267ffffffffffffffff80881660448401528a821660648401528616608483015282169063234636249060a4015b602060405180830381600087803b158015610e9357600080fd5b505af1158015610ea7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecb9190612107565b50610b12565b6005826005811115610ee557610ee56120d4565b1415610aba5760405163a002930160e01b81526001600160a01b0388811660048301526024820188905267ffffffffffffffff80881660448401528a821660648401528616608483015282169063a00293019060a401610e79565b6000610f95826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166111ad9092919063ffffffff16565b8051909150156110255780806020019051810190610fb39190611ec4565b6110255760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610236565b505050565b60006001846005811115611040576110406120d4565b141561105e576110578b8b8b8b8b8b8b8a8a6111c6565b9050610c40565b6002846005811115611072576110726120d4565b148061108f5750600484600581111561108d5761108d6120d4565b145b156110a557611057848c8c8c8c8c8b8a8a6113d4565b60038460058111156110b9576110b96120d4565b14806110d6575060058460058111156110d4576110d46120d4565b145b15610aba57611057848c8c8c8c8c8b8a8a611706565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b15801561113857600080fd5b505afa15801561114c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111709190612107565b61117a9190612120565b6040516001600160a01b038516602482015260448101829052909150610c1190859063095ea7b360e01b90606401610bad565b60606111bc848460008561195a565b90505b9392505050565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b15801561120257600080fd5b505afa158015611216573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123a91906120ea565b90506112506001600160a01b038b16828b6110ec565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b1580156112c657600080fd5b505af11580156112da573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b8152600401611392959493929190612190565b6000604051808303818588803b1580156113ab57600080fd5b505af11580156113bf573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b60008060028b60058111156113eb576113eb6120d4565b141561146957836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b15801561142a57600080fd5b505afa15801561143e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061146291906120ea565b90506114dd565b836001600160a01b031663c66a9c5a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156114a257600080fd5b505afa1580156114b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114da91906120ea565b90505b6114f16001600160a01b038a16828a6110ec565b600060028c6005811115611507576115076120d4565b141561162a576040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a401600060405180830381600087803b15801561157757600080fd5b505af115801561158b573d6000803e3d6000fd5b50505050308a8a8a8e8b4660405160200161160d9796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b6040516020818303038152906040528051906020012090506116d0565b6040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a4015b602060405180830381600087803b15801561169557600080fd5b505af11580156116a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116cd9190612107565b90505b604051634289fbb360e01b81526001600160a01b03861690634289fbb3908690611392908f908d90889088908e90600401612190565b60008060038b600581111561171d5761171d6120d4565b141561179b57836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561175c57600080fd5b505afa158015611770573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061179491906120ea565b905061180f565b836001600160a01b03166395b12c276040518163ffffffff1660e01b815260040160206040518083038186803b1580156117d457600080fd5b505afa1580156117e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061180c91906120ea565b90505b600060038c6005811115611825576118256120d4565b141561190557604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff8916606483015283169063de790c7e90608401600060405180830381600087803b15801561188d57600080fd5b505af11580156118a1573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528e811b82166034840152604883018e90528f901b1660688201526001600160c01b031960c08b811b8216607c84015246901b166084820152608c01915061160d9050565b60405163a002930160e01b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063a00293019060a40161167b565b6060824710156119d25760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610236565b843b611a205760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610236565b600080866001600160a01b03168587604051611a3c91906121d2565b60006040518083038185875af1925050503d8060008114611a79576040519150601f19603f3d011682016040523d82523d6000602084013e611a7e565b606091505b5091509150611a8e828286611a99565b979650505050505050565b60608315611aa85750816111bf565b825115611ab85782518084602001fd5b8160405162461bcd60e51b815260040161023691906121ee565b6001600160a01b038116811461079b57600080fd5b803567ffffffffffffffff81168114611aff57600080fd5b919050565b60008083601f840112611b1657600080fd5b50813567ffffffffffffffff811115611b2e57600080fd5b602083019150836020828501011115611b4657600080fd5b9250929050565b60008060008060608587031215611b6357600080fd5b8435611b6e81611ad2565b9350611b7c60208601611ae7565b9250604085013567ffffffffffffffff811115611b9857600080fd5b611ba487828801611b04565b95989497509550505050565b60008060008060608587031215611bc657600080fd5b8435611bd181611ad2565b935060208501359250604085013567ffffffffffffffff811115611b9857600080fd5b600060208284031215611c0657600080fd5b81356111bf81611ad2565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611c5057611c50611c11565b604052919050565b600080600080600060a08688031215611c7057600080fd5b8535611c7b81611ad2565b9450602086810135611c8c81611ad2565b945060408701359350611ca160608801611ae7565b9250608087013567ffffffffffffffff80821115611cbe57600080fd5b818901915089601f830112611cd257600080fd5b813581811115611ce457611ce4611c11565b611cf6601f8201601f19168501611c27565b91508082528a84828501011115611d0c57600080fd5b80848401858401376000848284010152508093505050509295509295909350565b60008060008060008060a08789031215611d4657600080fd5b8635611d5181611ad2565b95506020870135611d6181611ad2565b945060408701359350611d7660608801611ae7565b9250608087013567ffffffffffffffff811115611d9257600080fd5b611d9e89828a01611b04565b979a9699509497509295939492505050565b6000806000806000858703610100811215611dca57600080fd5b8635611dd581611ad2565b95506020870135611de581611ad2565b945060408701359350611dfa60608801611ae7565b92506080607f1982011215611e0e57600080fd5b506080860190509295509295909350565b801515811461079b57600080fd5b63ffffffff8116811461079b57600080fd5b600060808284031215611e5157600080fd5b6040516080810181811067ffffffffffffffff82111715611e7457611e74611c11565b6040528251611e8281611ad2565b81526020830151611e9281611ad2565b60208201526040830151611ea581611e1f565b60408201526060830151611eb881611e2d565b60608201529392505050565b600060208284031215611ed657600080fd5b81516111bf81611e1f565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600067ffffffffffffffff808316818516808303821115611f3057611f30611ef7565b01949350505050565b600060a082018783526020878185015260a0604085015281875180845260c086019150828901935060005b81811015611f895784516001600160a01b031683529383019391830191600101611f64565b50506001600160a01b03969096166060850152505050608001529392505050565b60006020808385031215611fbd57600080fd5b825167ffffffffffffffff80821115611fd557600080fd5b818501915085601f830112611fe957600080fd5b815181811115611ffb57611ffb611c11565b8060051b915061200c848301611c27565b818152918301840191848101908884111561202657600080fd5b938501935b838510156120445784518252938501939085019061202b565b98975050505050505050565b60808101823561205f81611ad2565b6001600160a01b03908116835260208401359061207b82611ad2565b166020830152604083013561208f81611e1f565b1515604083015260608301356120a481611e2d565b63ffffffff811660608401525092915050565b6000602082840312156120c957600080fd5b81356111bf81611e2d565b634e487b7160e01b600052602160045260246000fd5b6000602082840312156120fc57600080fd5b81516111bf81611ad2565b60006020828403121561211957600080fd5b5051919050565b6000821982111561213357612133611ef7565b500190565b60005b8381101561215357818101518382015260200161213b565b83811115610c115750506000910152565b6000815180845261217c816020860160208601612138565b601f01601f19169290920160200192915050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a06080830152611a8e60a0830184612164565b600082516121e4818460208701612138565b9190910192915050565b6020815260006111bf602083018461216456fea264697066735822122035bb2c07e5c4047c9a40654a028d5a93a8e94dec394e8624a9b594634f015acb64736f6c63430008090033",
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

// IMessageBusMetaData contains all meta data concerning the IMessageBus contract.
var IMessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"}],\"internalType\":\"structIMessageBus.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumIMessageBus.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structIMessageBus.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumIMessageBus.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structIMessageBus.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageBusMetaData.ABI instead.
var IMessageBusABI = IMessageBusMetaData.ABI

// IMessageBus is an auto generated Go binding around an Ethereum contract.
type IMessageBus struct {
	IMessageBusCaller     // Read-only binding to the contract
	IMessageBusTransactor // Write-only binding to the contract
	IMessageBusFilterer   // Log filterer for contract events
}

// IMessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageBusSession struct {
	Contract     *IMessageBus      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageBusCallerSession struct {
	Contract *IMessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageBusTransactorSession struct {
	Contract     *IMessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageBusRaw struct {
	Contract *IMessageBus // Generic contract binding to access the raw methods on
}

// IMessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageBusCallerRaw struct {
	Contract *IMessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageBusTransactorRaw struct {
	Contract *IMessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageBus creates a new instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBus(address common.Address, backend bind.ContractBackend) (*IMessageBus, error) {
	contract, err := bindIMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageBus{IMessageBusCaller: IMessageBusCaller{contract: contract}, IMessageBusTransactor: IMessageBusTransactor{contract: contract}, IMessageBusFilterer: IMessageBusFilterer{contract: contract}}, nil
}

// NewIMessageBusCaller creates a new read-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusCaller(address common.Address, caller bind.ContractCaller) (*IMessageBusCaller, error) {
	contract, err := bindIMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusCaller{contract: contract}, nil
}

// NewIMessageBusTransactor creates a new write-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageBusTransactor, error) {
	contract, err := bindIMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusTransactor{contract: contract}, nil
}

// NewIMessageBusFilterer creates a new log filterer instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageBusFilterer, error) {
	contract, err := bindIMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageBusFilterer{contract: contract}, nil
}

// bindIMessageBus binds a generic wrapper to an already deployed contract.
func bindIMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageBusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.IMessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusSession) CalcFee(_message []byte) (*big.Int, error) {
	return _IMessageBus.Contract.CalcFee(&_IMessageBus.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_IMessageBus *IMessageBusCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _IMessageBus.Contract.CalcFee(&_IMessageBus.CallOpts, _message)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusSession) LiquidityBridge() (common.Address, error) {
	return _IMessageBus.Contract.LiquidityBridge(&_IMessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) LiquidityBridge() (common.Address, error) {
	return _IMessageBus.Contract.LiquidityBridge(&_IMessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusSession) PegBridge() (common.Address, error) {
	return _IMessageBus.Contract.PegBridge(&_IMessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegBridge() (common.Address, error) {
	return _IMessageBus.Contract.PegBridge(&_IMessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusSession) PegBridgeV2() (common.Address, error) {
	return _IMessageBus.Contract.PegBridgeV2(&_IMessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegBridgeV2() (common.Address, error) {
	return _IMessageBus.Contract.PegBridgeV2(&_IMessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusSession) PegVault() (common.Address, error) {
	return _IMessageBus.Contract.PegVault(&_IMessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegVault() (common.Address, error) {
	return _IMessageBus.Contract.PegVault(&_IMessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IMessageBus.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusSession) PegVaultV2() (common.Address, error) {
	return _IMessageBus.Contract.PegVaultV2(&_IMessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_IMessageBus *IMessageBusCallerSession) PegVaultV2() (common.Address, error) {
	return _IMessageBus.Contract.PegVaultV2(&_IMessageBus.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route IMessageBusRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessage(_message []byte, _route IMessageBusRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x654317bf.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessage(_message []byte, _route IMessageBusRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer IMessageBusTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessageWithTransfer(_message []byte, _transfer IMessageBusTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransfer(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0xa2232213.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer IMessageBusTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransfer(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer IMessageBusTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer IMessageBusTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransferRefund(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x588df416.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer IMessageBusTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessageWithTransferRefund(&_IMessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessage(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessageWithTransfer(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessageWithTransfer(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_IMessageBus *IMessageBusTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"}],\"internalType\":\"structMessageBusReceiver.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeeBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridgeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002c8038038062002c8083398101604081905262000034916200011d565b84848484848a6200004533620000b4565b6001600160a01b03908116608052600580546001600160a01b0319908116978316979097179055600680548716958216959095179094556007805486169385169390931790925560088054851691841691909117905560098054909316911617905550620001b1945050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146200011a57600080fd5b50565b60008060008060008060c087890312156200013757600080fd5b8651620001448162000104565b6020880151909650620001578162000104565b60408801519095506200016a8162000104565b60608801519094506200017d8162000104565b6080880151909350620001908162000104565b60a0880151909250620001a38162000104565b809150509295509295509295565b608051612aac620001d46000396000818161041001526106ce0152612aac6000f3fe6080604052600436106101ac5760003560e01c806395b12c27116100ec578063cd2abd661161008a578063e2c1ed2511610064578063e2c1ed25146104af578063f2fde38b146104cf578063f60bbe2a146104ef578063f83b0fb91461050557600080fd5b8063cd2abd6614610432578063d8257d171461046f578063dfa2dbaf1461048f57600080fd5b80639f3ce55a116100c65780639f3ce55a146103b8578063a2232213146103cb578063c66a9c5a146103de578063ccf2683b146103fe57600080fd5b806395b12c271461036257806395e911a8146103825780639b05a7751461039857600080fd5b8063588be02b11610159578063654317bf11610133578063654317bf146102d957806382980dc4146102ec57806382efd502146103245780638da5cb5b1461034457600080fd5b8063588be02b14610279578063588df416146102995780635b3e5f50146102ac57600080fd5b8063359ef75b1161018a578063359ef75b146102135780634289fbb3146102335780635335dca21461024657600080fd5b806303cbfe66146101b157806306c28bd6146101d35780632ff4c411146101f3575b600080fd5b3480156101bd57600080fd5b506101d16101cc36600461208f565b610525565b005b3480156101df57600080fd5b506101d16101ee3660046120aa565b6105b5565b3480156101ff57600080fd5b506101d161020e36600461210f565b610623565b34801561021f57600080fd5b506101d161022e3660046121c3565b61086f565b6101d161024136600461226a565b61088b565b34801561025257600080fd5b506102666102613660046122e2565b610932565b6040519081526020015b60405180910390f35b34801561028557600080fd5b506101d161029436600461208f565b610958565b6101d16102a7366004612324565b6109e3565b3480156102b857600080fd5b506102666102c736600461208f565b60036020526000908152604090205481565b6101d16102e7366004612415565b610c26565b3480156102f857600080fd5b5060055461030c906001600160a01b031681565b6040516001600160a01b039091168152602001610270565b34801561033057600080fd5b506101d161033f36600461208f565b610e35565b34801561035057600080fd5b506000546001600160a01b031661030c565b34801561036e57600080fd5b5060085461030c906001600160a01b031681565b34801561038e57600080fd5b5061026660015481565b3480156103a457600080fd5b506101d16103b336600461208f565b610ec0565b6101d16103c63660046124dc565b610f4b565b6101d16103d9366004612324565b610fec565b3480156103ea57600080fd5b5060095461030c906001600160a01b031681565b34801561040a57600080fd5b5061030c7f000000000000000000000000000000000000000000000000000000000000000081565b34801561043e57600080fd5b5061046261044d3660046120aa565b60046020526000908152604090205460ff1681565b6040516102709190612560565b34801561047b57600080fd5b5060075461030c906001600160a01b031681565b34801561049b57600080fd5b5060065461030c906001600160a01b031681565b3480156104bb57600080fd5b506101d16104ca3660046120aa565b6111c6565b3480156104db57600080fd5b506101d16104ea36600461208f565b611234565b3480156104fb57600080fd5b5061026660025481565b34801561051157600080fd5b506101d161052036600461208f565b611325565b336105386000546001600160a01b031690565b6001600160a01b0316146105935760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b336105c86000546001600160a01b031690565b6001600160a01b03161461061e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161058a565b600155565b6000463060405160200161067492919091825260601b6001600160601b03191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6001600160601b0319168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc2291610715918b908b908b908b908b908b9060780161265d565b60006040518083038186803b15801561072d57600080fd5b505afa158015610741573d6000803e3d6000fd5b505050506001600160a01b038916600090815260036020526040812054610768908a61276d565b9050600081116107ba5760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f20776974686472617700000000000000604482015260640161058a565b60008a6001600160a01b03168261c35090604051600060405180830381858888f193505050503d806000811461080c576040519150601f19603f3d011682016040523d82523d6000602084013e610811565b606091505b50509050806108625760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f2077697468647261772066656500000000000000000000604482015260640161058a565b5050505050505050505050565b6108776113b0565b6108848585858585611414565b5050505050565b60006108978383610932565b9050803410156108dc5760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b604482015260640161058a565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66888888888888346040516109219796959493929190612784565b60405180910390a250505050505050565b60025460009061094290836127d1565b60015461094f91906127f0565b90505b92915050565b3361096b6000546001600160a01b031690565b6001600160a01b0316146109c15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161058a565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b60006109ee886114cc565b90506000808281526004602081905260409091205460ff1690811115610a1657610a16612536565b14610a635760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c726561647920657865637574656400000000000000604482015260640161058a565b6000818152600460208181526040808420805460ff1916909317909255815146918101919091526001600160601b03193060601b16918101919091527f4d657373616765576974685472616e73666572526566756e64000000000000006054820152606d01604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e604051602001610b1b9493929190612808565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610b52979695949392919061265d565b60006040518083038186803b158015610b6a57600080fd5b505afa158015610b7e573d6000803e3d6000fd5b50505050600080610b908b8e8e611d85565b90508015610ba15760019150610ba6565b600291505b60008481526004602081905260409091208054849260ff19909116906001908490811115610bd657610bd6612536565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e760008584604051610c0f93929190612829565b60405180910390a150505050505050505050505050565b6000610c33888b8b611ebc565b90506000808281526004602081905260409091205460ff1690811115610c5b57610c5b612536565b14610ca85760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c72656164792065786563757465640000000000000000604482015260640161058a565b600081815260046020818152604092839020805460ff1916909217909155815146818301523060601b6001600160601b031916818401527f4d6573736167650000000000000000000000000000000000000000000000000060548201528251603b818303018152605b820184528051920191909120600554607b8301829052609b8084018690528451808503909101815260bb840194859052633416de1160e11b90945290926001600160a01b039091169163682dbc2291610d78918c908c908c908c908c908c9060bf0161265d565b60006040518083038186803b158015610d9057600080fd5b505afa158015610da4573d6000803e3d6000fd5b50505050600080610db68b8e8e611f23565b90508015610dc75760019150610dcc565b600291505b60008481526004602081905260409091208054849260ff19909116906001908490811115610dfc57610dfc612536565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e760018584604051610c0f93929190612829565b33610e486000546001600160a01b031690565b6001600160a01b031614610e9e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161058a565b600880546001600160a01b0319166001600160a01b0392909216919091179055565b33610ed36000546001600160a01b031690565b6001600160a01b031614610f295760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161058a565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6000610f578383610932565b905080341015610f9c5760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b604482015260640161058a565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e48686868634604051610fdd95949392919061285b565b60405180910390a25050505050565b6000610ff7886114cc565b90506000808281526004602081905260409091205460ff169081111561101f5761101f612536565b1461106c5760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c726561647920657865637574656400000000000000604482015260640161058a565b6000818152600460208181526040808420805460ff1916909317909255815146918101919091526001600160601b03193060601b16918101919091527f4d657373616765576974685472616e73666572000000000000000000000000006054820152606701604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e6040516020016111249493929190612808565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b815260040161115b979695949392919061265d565b60006040518083038186803b15801561117357600080fd5b505afa158015611187573d6000803e3d6000fd5b505050506000806111998b8e8e611f7b565b905080156111aa5760019150610ba6565b6111b58b8e8e611fed565b90508015610ba15760039150610ba6565b336111d96000546001600160a01b031690565b6001600160a01b03161461122f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161058a565b600255565b336112476000546001600160a01b031690565b6001600160a01b03161461129d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161058a565b6001600160a01b0381166113195760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161058a565b61132281612023565b50565b336113386000546001600160a01b031690565b6001600160a01b03161461138e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161058a565b600980546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b0316156114095760405162461bcd60e51b815260206004820152601160248201527f6f776e657220616c726561647920736574000000000000000000000000000000604482015260640161058a565b61141233612023565b565b6005546001600160a01b03161561146d5760405162461bcd60e51b815260206004820152601b60248201527f6c697175696469747942726964676520616c7265616479207365740000000000604482015260640161058a565b600580546001600160a01b03199081166001600160a01b03978816179091556006805482169587169590951790945560078054851693861693909317909255600880548416918516919091179055600980549092169216919091179055565b6000808060016114df6020860186612896565b60068111156114f0576114f0612536565b141561168b57611506604085016020860161208f565b611516606086016040870161208f565b611526608087016060880161208f565b608087013561153b60e0890160c08a016128b7565b6040516001600160601b0319606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600554633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b9060240160206040518083038186803b1580156115fd57600080fd5b505afa158015611611573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163591906128e1565b15156001146116865760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f7420657869737400000000000000000000604482015260640161058a565b611d50565b600261169a6020860186612896565b60068111156116ab576116ab612536565b141561181857466116c260c0860160a087016128b7565b6116d2606087016040880161208f565b6116e2608088016060890161208f565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526001600160601b0319606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600554631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab4289060240160206040518083038186803b15801561178f57600080fd5b505afa1580156117a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117c791906128e1565b15156001146116865760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f7420657869737400000000000000604482015260640161058a565b60036118276020860186612896565b600681111561183857611838612536565b14806118615750600461184e6020860186612896565b600681111561185f5761185f612536565b145b15611ae357611876606085016040860161208f565b611886608086016060870161208f565b608086013561189b604088016020890161208f565b6118ab60e0890160c08a016128b7565b604051606095861b6001600160601b0319908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f1981840301815291905280516020909101209150600361192b6020860186612896565b600681111561193c5761193c612536565b1415611a1357506006546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e64725906024015b60206040518083038186803b15801561198a57600080fd5b505afa15801561199e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c291906128e1565b15156001146116865760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f742065786973740000000000000000000000604482015260640161058a565b506007546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e647259060240160206040518083038186803b158015611a5a57600080fd5b505afa158015611a6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a9291906128e1565b15156001146116865760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f7420657869737400000000000000604482015260640161058a565b6005611af26020860186612896565b6006811115611b0357611b03612536565b1480611b2c57506006611b196020860186612896565b6006811115611b2a57611b2a612536565b145b15611d50576005611b406020860186612896565b6006811115611b5157611b51612536565b1415611b6957506006546001600160a01b0316611b77565b506007546001600160a01b03165b611b87606085016040860161208f565b611b97608086016060870161208f565b6080860135611bac604088016020890161208f565b611bbc60e0890160c08a016128b7565b604051606095861b6001600160601b0319908116602083015294861b85166034820152604881019390935290841b8316606883015260c01b6001600160c01b031916607c82015260e087013560848201529183901b1660a482015260b80160408051601f19818403018152919052805160209091012091506005611c436020860186612896565b6006811115611c5457611c54612536565b1415611c87576040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401611972565b6040516301e6472560e01b8152600481018390526001600160a01b038216906301e647259060240160206040518083038186803b158015611cc757600080fd5b505afa158015611cdb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cff91906128e1565b1515600114611d505760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f7420657869737400000000000000604482015260640161058a565b60008183604051602001611d669392919061291a565b6040516020818303038152906040528051906020012092505050919050565b60008080611d99606087016040880161208f565b6001600160a01b03163463105f4af960e11b611dbb60808a0160608b0161208f565b89608001358989604051602401611dd59493929190612946565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909416939093179092529051611e409190612979565b60006040518083038185875af1925050503d8060008114611e7d576040519150601f19603f3d011682016040523d82523d6000602084013e611e82565b606091505b50915091508115611eae57600081806020019051810190611ea391906128e1565b9350611eb592505050565b6000925050505b9392505050565b60006001611ecd602086018661208f565b611edd604087016020880161208f565b611eed60608801604089016128b7565b8686604051602001611f0496959493929190612995565b6040516020818303038152906040528051906020012090509392505050565b60008080611f37604087016020880161208f565b6001600160a01b031634631599d26560e01b611f5660208a018a61208f565b611f6660608b0160408c016128b7565b8989604051602401611dd594939291906129f4565b60008080611f8f606087016040880161208f565b6001600160a01b03163463671aeecd60e11b611fb160408a0160208b0161208f565b611fc160808b0160608c0161208f565b60808b0135611fd660e08d0160c08e016128b7565b8b8b604051602401611dd596959493929190612a27565b60008080612001606087016040880161208f565b6001600160a01b0316346378079ce760e11b611fb160408a0160208b0161208f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b038116811461208a57600080fd5b919050565b6000602082840312156120a157600080fd5b61094f82612073565b6000602082840312156120bc57600080fd5b5035919050565b60008083601f8401126120d557600080fd5b50813567ffffffffffffffff8111156120ed57600080fd5b6020830191508360208260051b850101111561210857600080fd5b9250929050565b60008060008060008060008060a0898b03121561212b57600080fd5b61213489612073565b975060208901359650604089013567ffffffffffffffff8082111561215857600080fd5b6121648c838d016120c3565b909850965060608b013591508082111561217d57600080fd5b6121898c838d016120c3565b909650945060808b01359150808211156121a257600080fd5b506121af8b828c016120c3565b999c989b5096995094979396929594505050565b600080600080600060a086880312156121db57600080fd5b6121e486612073565b94506121f260208701612073565b935061220060408701612073565b925061220e60608701612073565b915061221c60808701612073565b90509295509295909350565b60008083601f84011261223a57600080fd5b50813567ffffffffffffffff81111561225257600080fd5b60208301915083602082850101111561210857600080fd5b60008060008060008060a0878903121561228357600080fd5b61228c87612073565b9550602087013594506122a160408801612073565b935060608701359250608087013567ffffffffffffffff8111156122c457600080fd5b6122d089828a01612228565b979a9699509497509295939492505050565b600080602083850312156122f557600080fd5b823567ffffffffffffffff81111561230c57600080fd5b61231885828601612228565b90969095509350505050565b6000806000806000806000806000898b0361018081121561234457600080fd5b8a3567ffffffffffffffff8082111561235c57600080fd5b6123688e838f01612228565b909c509a508a9150610100601f198401121561238357600080fd5b60208d0199506101208d013592508083111561239e57600080fd5b6123aa8e848f016120c3565b90995097506101408d01359250889150808311156123c757600080fd5b6123d38e848f016120c3565b90975095506101608d01359250869150808311156123f057600080fd5b50506123fe8c828d016120c3565b915080935050809150509295985092959850929598565b6000806000806000806000806000898b0360e081121561243457600080fd5b8a3567ffffffffffffffff8082111561244c57600080fd5b6124588e838f01612228565b909c509a508a91506060601f198401121561247257600080fd5b60208d01995060808d013592508083111561248c57600080fd5b6124988e848f016120c3565b909950975060a08d01359250889150808311156124b457600080fd5b6124c08e848f016120c3565b909750955060c08d01359250869150808311156123f057600080fd5b600080600080606085870312156124f257600080fd5b6124fb85612073565b935060208501359250604085013567ffffffffffffffff81111561251e57600080fd5b61252a87828801612228565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b6005811061255c5761255c612536565b9052565b60208101610952828461254c565b60005b83811015612589578181015183820152602001612571565b83811115612598576000848401525b50505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8183526000602080850194508260005b85811015612603576001600160a01b036125f083612073565b16875295820195908201906001016125d7565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561264057600080fd5b8260051b8083602087013760009401602001938452509192915050565b608081526000885180608084015261267c8160a0850160208d0161256e565b601f01601f1916820182810360a090810160208501528101889052600588901b810160c09081019082018a60005b8b81101561271d5784840360bf190183528135368e9003601e190181126126d057600080fd5b8d01803567ffffffffffffffff8111156126e957600080fd5b8036038f13156126f857600080fd5b61270686826020850161259e565b9550505060209283019291909101906001016126aa565b505050838103604085015261273381888a6125c7565b915050828103606084015261274981858761260e565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b60008282101561277f5761277f612757565b500390565b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c060808301526127bd60c08301858761259e565b90508260a083015298975050505050505050565b60008160001904831182151516156127eb576127eb612757565b500290565b6000821982111561280357612803612757565b500190565b84815283602082015281836040830137600091016040019081529392505050565b606081016002851061283d5761283d612536565b848252836020830152612853604083018461254c565b949350505050565b6001600160a01b038616815284602082015260806040820152600061288460808301858761259e565b90508260608301529695505050505050565b6000602082840312156128a857600080fd5b813560078110611eb557600080fd5b6000602082840312156128c957600080fd5b813567ffffffffffffffff81168114611eb557600080fd5b6000602082840312156128f357600080fd5b81518015158114611eb557600080fd5b6002811061291357612913612536565b60f81b9052565b6129248185612903565b60609290921b6001600160601b03191660018301526015820152603501919050565b6001600160a01b038516815283602082015260606040820152600061296f60608301848661259e565b9695505050505050565b6000825161298b81846020870161256e565b9190910192915050565b61299f8188612903565b60006bffffffffffffffffffffffff19808860601b166001840152808760601b166015840152506001600160c01b03198560c01b16602983015282846031840137506000910160310190815295945050505050565b6001600160a01b038516815267ffffffffffffffff8416602082015260606040820152600061296f60608301848661259e565b60006001600160a01b03808916835280881660208401525085604083015267ffffffffffffffff8516606083015260a06080830152612a6a60a08301848661259e565b9897505050505050505056fea26469706673582212205824a872bcdcf6b6f809243142a84e8aba584a5c98fb75516eb1146245fcc64964736f6c63430008090033",
}

// MessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusMetaData.ABI instead.
var MessageBusABI = MessageBusMetaData.ABI

// MessageBusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusMetaData.Bin instead.
var MessageBusBin = MessageBusMetaData.Bin

// DeployMessageBus deploys a new Ethereum contract, binding an instance of MessageBus to it.
func DeployMessageBus(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (common.Address, *types.Transaction, *MessageBus, error) {
	parsed, err := MessageBusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusBin), backend, _sigsVerifier, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
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

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusSession) PegBridgeV2() (common.Address, error) {
	return _MessageBus.Contract.PegBridgeV2(&_MessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegBridgeV2() (common.Address, error) {
	return _MessageBus.Contract.PegBridgeV2(&_MessageBus.CallOpts)
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

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusSession) PegVaultV2() (common.Address, error) {
	return _MessageBus.Contract.PegVaultV2(&_MessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegVaultV2() (common.Address, error) {
	return _MessageBus.Contract.PegVaultV2(&_MessageBus.CallOpts)
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

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusTransactor) Init(opts *bind.TransactOpts, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "init", _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusTransactorSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
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

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegBridgeV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegBridgeV2", _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridgeV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridgeV2(&_MessageBus.TransactOpts, _addr)
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

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegVaultV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegVaultV2", _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVaultV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVaultV2(&_MessageBus.TransactOpts, _addr)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"}],\"internalType\":\"structMessageBusReceiver.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMessageBusReceiver.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"}],\"internalType\":\"structMessageBusReceiver.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMessageBusReceiver.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridgeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620021523803806200215283398101604081905262000034916200010f565b6200003f33620000a2565b600280546001600160a01b03199081166001600160a01b039788161790915560038054821695871695909517909455600480548516938616939093179092556005805484169185169190911790556006805490921692169190911790556200017f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146200010a57600080fd5b919050565b600080600080600060a086880312156200012857600080fd5b6200013386620000f2565b94506200014360208701620000f2565b93506200015360408701620000f2565b92506200016360608701620000f2565b91506200017360808701620000f2565b90509295509295909350565b611fc3806200018f6000396000f3fe6080604052600436106100f35760003560e01c80639b05a7751161008a578063d8257d1711610059578063d8257d171461028b578063dfa2dbaf146102ab578063f2fde38b146102cb578063f83b0fb9146102eb57600080fd5b80639b05a775146101fb578063a22322131461021b578063c66a9c5a1461022e578063cd2abd661461024e57600080fd5b806382980dc4116100c657806382980dc41461016057806382efd5021461019d5780638da5cb5b146101bd57806395b12c27146101db57600080fd5b806303cbfe66146100f8578063588be02b1461011a578063588df4161461013a578063654317bf1461014d575b600080fd5b34801561010457600080fd5b506101186101133660046118b4565b61030b565b005b34801561012657600080fd5b506101186101353660046118b4565b61039b565b61011861014836600461195d565b610426565b61011861015b366004611a4e565b610666565b34801561016c57600080fd5b50600254610180906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101a957600080fd5b506101186101b83660046118b4565b610874565b3480156101c957600080fd5b506000546001600160a01b0316610180565b3480156101e757600080fd5b50600554610180906001600160a01b031681565b34801561020757600080fd5b506101186102163660046118b4565b6108ff565b61011861022936600461195d565b61098a565b34801561023a57600080fd5b50600654610180906001600160a01b031681565b34801561025a57600080fd5b5061027e610269366004611b15565b60016020526000908152604090205460ff1681565b6040516101949190611b58565b34801561029757600080fd5b50600454610180906001600160a01b031681565b3480156102b757600080fd5b50600354610180906001600160a01b031681565b3480156102d757600080fd5b506101186102e63660046118b4565b610b63565b3480156102f757600080fd5b506101186103063660046118b4565b610c54565b3361031e6000546001600160a01b031690565b6001600160a01b0316146103795760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600380546001600160a01b0319166001600160a01b0392909216919091179055565b336103ae6000546001600160a01b031690565b6001600160a01b0316146104045760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610370565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600061043188610cdf565b90506000808281526001602052604090205460ff16600481111561045757610457611b2e565b146104a45760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c7265616479206578656375746564000000000000006044820152606401610370565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765576974685472616e73666572526566756e64000000000000006054820152606d01604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e60405160200161055d9493929190611b6c565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b81526004016105949796959493929190611c7c565b60006040518083038186803b1580156105ac57600080fd5b505afa1580156105c0573d6000803e3d6000fd5b505050506000806105d28b8e8e6115aa565b905080156105e357600191506105e8565b600291505b60008481526001602081905260409091208054849260ff199091169083600481111561061657610616611b2e565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e76000858460405161064f93929190611d76565b60405180910390a150505050505050505050505050565b6000610673888b8b6116e1565b90506000808281526001602052604090205460ff16600481111561069957610699611b2e565b146106e65760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c726561647920657865637574656400000000000000006044820152606401610370565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765000000000000000000000000000000000000000000000000006054820152605b0160408051601f1981840301815282825280516020918201206002549184018190528383018690528251808503840181526060850193849052633416de1160e11b90935293506001600160a01b03169163682dbc22916107b9918c908c908c908c908c908c90606401611c7c565b60006040518083038186803b1580156107d157600080fd5b505afa1580156107e5573d6000803e3d6000fd5b505050506000806107f78b8e8e611748565b90508015610808576001915061080d565b600291505b60008481526001602081905260409091208054849260ff199091169083600481111561083b5761083b611b2e565b02179055507f29122f2c841ca2c3b2feefc4c23e90755d735d8e5b84f307151532e0f1ad62e76001858460405161064f93929190611d76565b336108876000546001600160a01b031690565b6001600160a01b0316146108dd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610370565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b336109126000546001600160a01b031690565b6001600160a01b0316146109685760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610370565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b600061099588610cdf565b90506000808281526001602052604090205460ff1660048111156109bb576109bb611b2e565b14610a085760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c7265616479206578656375746564000000000000006044820152606401610370565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765576974685472616e73666572000000000000000000000000006054820152606701604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e604051602001610ac19493929190611b6c565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610af89796959493929190611c7c565b60006040518083038186803b158015610b1057600080fd5b505afa158015610b24573d6000803e3d6000fd5b50505050600080610b368b8e8e6117a0565b90508015610b4757600191506105e8565b610b528b8e8e611812565b905080156105e357600391506105e8565b33610b766000546001600160a01b031690565b6001600160a01b031614610bcc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610370565b6001600160a01b038116610c485760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610370565b610c5181611848565b50565b33610c676000546001600160a01b031690565b6001600160a01b031614610cbd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610370565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b600080806001610cf26020860186611da8565b6006811115610d0357610d03611b2e565b1415610ea357610d1960408501602086016118b4565b610d2960608601604087016118b4565b610d3960808701606088016118b4565b6080870135610d4e60e0890160c08a01611dc9565b6040516bffffffffffffffffffffffff19606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600254633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b9060240160206040518083038186803b158015610e1557600080fd5b505afa158015610e29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e4d9190611df3565b1515600114610e9e5760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f74206578697374000000000000000000006044820152606401610370565b611575565b6002610eb26020860186611da8565b6006811115610ec357610ec3611b2e565b14156110355746610eda60c0860160a08701611dc9565b610eea60608701604088016118b4565b610efa60808801606089016118b4565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526bffffffffffffffffffffffff19606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600254631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab4289060240160206040518083038186803b158015610fac57600080fd5b505afa158015610fc0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fe49190611df3565b1515600114610e9e5760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f74206578697374000000000000006044820152606401610370565b60036110446020860186611da8565b600681111561105557611055611b2e565b148061107e5750600461106b6020860186611da8565b600681111561107c5761107c611b2e565b145b156113035761109360608501604086016118b4565b6110a360808601606087016118b4565b60808601356110b860408801602089016118b4565b6110c860e0890160c08a01611dc9565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f1981840301815291905280516020909101209150600361114d6020860186611da8565b600681111561115e5761115e611b2e565b141561123557506003546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e64725906024015b60206040518083038186803b1580156111ac57600080fd5b505afa1580156111c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e49190611df3565b1515600114610e9e5760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f7420657869737400000000000000000000006044820152606401610370565b50600480546040516301e6472560e01b81529182018390526001600160a01b03169081906301e647259060240160206040518083038186803b15801561127a57600080fd5b505afa15801561128e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b29190611df3565b1515600114610e9e5760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f74206578697374000000000000006044820152606401610370565b60056113126020860186611da8565b600681111561132357611323611b2e565b148061134c575060066113396020860186611da8565b600681111561134a5761134a611b2e565b145b156115755760056113606020860186611da8565b600681111561137157611371611b2e565b141561138957506003546001600160a01b0316611397565b506004546001600160a01b03165b6113a760608501604086016118b4565b6113b760808601606087016118b4565b60808601356113cc60408801602089016118b4565b6113dc60e0890160c08a01611dc9565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b85166034820152604881019390935290841b8316606883015260c01b6001600160c01b031916607c82015260e087013560848201529183901b1660a482015260b80160408051601f198184030181529190528051602090910120915060056114686020860186611da8565b600681111561147957611479611b2e565b14156114ac576040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401611194565b6040516301e6472560e01b8152600481018390526001600160a01b038216906301e647259060240160206040518083038186803b1580156114ec57600080fd5b505afa158015611500573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115249190611df3565b15156001146115755760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f74206578697374000000000000006044820152606401610370565b6000818360405160200161158b93929190611e2c565b6040516020818303038152906040528051906020012092505050919050565b600080806115be60608701604088016118b4565b6001600160a01b03163463105f4af960e11b6115e060808a0160608b016118b4565b896080013589896040516024016115fa9493929190611e5d565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009094169390931790925290516116659190611e90565b60006040518083038185875af1925050503d80600081146116a2576040519150601f19603f3d011682016040523d82523d6000602084013e6116a7565b606091505b509150915081156116d3576000818060200190518101906116c89190611df3565b93506116da92505050565b6000925050505b9392505050565b600060016116f260208601866118b4565b61170260408701602088016118b4565b6117126060880160408901611dc9565b868660405160200161172996959493929190611eac565b6040516020818303038152906040528051906020012090509392505050565b6000808061175c60408701602088016118b4565b6001600160a01b031634631599d26560e01b61177b60208a018a6118b4565b61178b60608b0160408c01611dc9565b89896040516024016115fa9493929190611f0b565b600080806117b460608701604088016118b4565b6001600160a01b03163463671aeecd60e11b6117d660408a0160208b016118b4565b6117e660808b0160608c016118b4565b60808b01356117fb60e08d0160c08e01611dc9565b8b8b6040516024016115fa96959493929190611f3e565b6000808061182660608701604088016118b4565b6001600160a01b0316346378079ce760e11b6117d660408a0160208b016118b4565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80356001600160a01b03811681146118af57600080fd5b919050565b6000602082840312156118c657600080fd5b6116da82611898565b60008083601f8401126118e157600080fd5b50813567ffffffffffffffff8111156118f957600080fd5b60208301915083602082850101111561191157600080fd5b9250929050565b60008083601f84011261192a57600080fd5b50813567ffffffffffffffff81111561194257600080fd5b6020830191508360208260051b850101111561191157600080fd5b6000806000806000806000806000898b0361018081121561197d57600080fd5b8a3567ffffffffffffffff8082111561199557600080fd5b6119a18e838f016118cf565b909c509a508a9150610100601f19840112156119bc57600080fd5b60208d0199506101208d01359250808311156119d757600080fd5b6119e38e848f01611918565b90995097506101408d0135925088915080831115611a0057600080fd5b611a0c8e848f01611918565b90975095506101608d0135925086915080831115611a2957600080fd5b5050611a378c828d01611918565b915080935050809150509295985092959850929598565b6000806000806000806000806000898b0360e0811215611a6d57600080fd5b8a3567ffffffffffffffff80821115611a8557600080fd5b611a918e838f016118cf565b909c509a508a91506060601f1984011215611aab57600080fd5b60208d01995060808d0135925080831115611ac557600080fd5b611ad18e848f01611918565b909950975060a08d0135925088915080831115611aed57600080fd5b611af98e848f01611918565b909750955060c08d0135925086915080831115611a2957600080fd5b600060208284031215611b2757600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b60058110611b5457611b54611b2e565b9052565b60208101611b668284611b44565b92915050565b84815283602082015281836040830137600091016040019081529392505050565b60005b83811015611ba8578181015183820152602001611b90565b83811115611bb7576000848401525b50505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8183526000602080850194508260005b85811015611c22576001600160a01b03611c0f83611898565b1687529582019590820190600101611bf6565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115611c5f57600080fd5b8260051b8083602087013760009401602001938452509192915050565b6080815260008851806080840152611c9b8160a0850160208d01611b8d565b601f01601f1916820182810360a090810160208501528101889052600588901b810160c09081019082018a60005b8b811015611d3c5784840360bf190183528135368e9003601e19018112611cef57600080fd5b8d01803567ffffffffffffffff811115611d0857600080fd5b8036038f1315611d1757600080fd5b611d25868260208501611bbd565b955050506020928301929190910190600101611cc9565b5050508381036040850152611d5281888a611be6565b9150508281036060840152611d68818587611c2d565b9a9950505050505050505050565b6060810160028510611d8a57611d8a611b2e565b848252836020830152611da06040830184611b44565b949350505050565b600060208284031215611dba57600080fd5b8135600781106116da57600080fd5b600060208284031215611ddb57600080fd5b813567ffffffffffffffff811681146116da57600080fd5b600060208284031215611e0557600080fd5b815180151581146116da57600080fd5b60028110611e2557611e25611b2e565b60f81b9052565b611e368185611e15565b60609290921b6bffffffffffffffffffffffff191660018301526015820152603501919050565b6001600160a01b0385168152836020820152606060408201526000611e86606083018486611bbd565b9695505050505050565b60008251611ea2818460208701611b8d565b9190910192915050565b611eb68188611e15565b60006bffffffffffffffffffffffff19808860601b166001840152808760601b166015840152506001600160c01b03198560c01b16602983015282846031840137506000910160310190815295945050505050565b6001600160a01b038516815267ffffffffffffffff84166020820152606060408201526000611e86606083018486611bbd565b60006001600160a01b03808916835280881660208401525085604083015267ffffffffffffffff8516606083015260a06080830152611f8160a083018486611bbd565b9897505050505050505056fea2646970667358221220e43e5cf7060996dc241baa74cbe0454bcc503eefae99d1134fefc6d8f659139c64736f6c63430008090033",
}

// MessageBusReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusReceiverMetaData.ABI instead.
var MessageBusReceiverABI = MessageBusReceiverMetaData.ABI

// MessageBusReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusReceiverMetaData.Bin instead.
var MessageBusReceiverBin = MessageBusReceiverMetaData.Bin

// DeployMessageBusReceiver deploys a new Ethereum contract, binding an instance of MessageBusReceiver to it.
func DeployMessageBusReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (common.Address, *types.Transaction, *MessageBusReceiver, error) {
	parsed, err := MessageBusReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusReceiverBin), backend, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
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

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegBridgeV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridgeV2(&_MessageBusReceiver.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegBridgeV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridgeV2(&_MessageBusReceiver.CallOpts)
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

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegVaultV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVaultV2(&_MessageBusReceiver.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegVaultV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVaultV2(&_MessageBusReceiver.CallOpts)
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

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegBridgeV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegBridgeV2", _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridgeV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridgeV2(&_MessageBusReceiver.TransactOpts, _addr)
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

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegVaultV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegVaultV2", _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVaultV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVaultV2(&_MessageBusReceiver.TransactOpts, _addr)
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
	Bin: "0x60a060405234801561001057600080fd5b50604051610f01380380610f0183398101604081905261002f91610099565b61003833610049565b6001600160a01b03166080526100c9565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ab57600080fd5b81516001600160a01b03811681146100c257600080fd5b9392505050565b608051610e166100eb600039600081816101ee015261038e0152610e166000f3fe6080604052600436106100c75760003560e01c806395e911a811610074578063e2c1ed251161004e578063e2c1ed2514610210578063f2fde38b14610230578063f60bbe2a1461025057600080fd5b806395e911a8146101b35780639f3ce55a146101c9578063ccf2683b146101dc57600080fd5b80635335dca2116100a55780635335dca2146101215780635b3e5f50146101545780638da5cb5b1461018157600080fd5b806306c28bd6146100cc5780632ff4c411146100ee5780634289fbb31461010e575b600080fd5b3480156100d857600080fd5b506100ec6100e7366004610862565b610266565b005b3480156100fa57600080fd5b506100ec6101093660046108e3565b6102d9565b6100ec61011c3660046109d9565b61052f565b34801561012d57600080fd5b5061014161013c366004610a51565b6105d6565b6040519081526020015b60405180910390f35b34801561016057600080fd5b5061014161016f366004610a93565b60036020526000908152604090205481565b34801561018d57600080fd5b506000546001600160a01b03165b6040516001600160a01b03909116815260200161014b565b3480156101bf57600080fd5b5061014160015481565b6100ec6101d7366004610aae565b6105fa565b3480156101e857600080fd5b5061019b7f000000000000000000000000000000000000000000000000000000000000000081565b34801561021c57600080fd5b506100ec61022b366004610862565b61069b565b34801561023c57600080fd5b506100ec61024b366004610a93565b610709565b34801561025c57600080fd5b5061014160025481565b336102796000546001600160a01b031690565b6001600160a01b0316146102d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600155565b6000463060405160200161032f92919091825260601b6bffffffffffffffffffffffff191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6bffffffffffffffffffffffff19168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc22916103d5918b908b908b908b908b908b90607801610c58565b60006040518083038186803b1580156103ed57600080fd5b505afa158015610401573d6000803e3d6000fd5b505050506001600160a01b038916600090815260036020526040812054610428908a610d0a565b90506000811161047a5760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f2077697468647261770000000000000060448201526064016102cb565b60008a6001600160a01b03168261c35090604051600060405180830381858888f193505050503d80600081146104cc576040519150601f19603f3d011682016040523d82523d6000602084013e6104d1565b606091505b50509050806105225760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f207769746864726177206665650000000000000000000060448201526064016102cb565b5050505050505050505050565b600061053b83836105d6565b9050803410156105805760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102cb565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66888888888888346040516105c59796959493929190610d21565b60405180910390a250505050505050565b6002546000906105e69083610d6e565b6001546105f39190610d8d565b9392505050565b600061060683836105d6565b90508034101561064b5760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102cb565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4868686863460405161068c959493929190610da5565b60405180910390a25050505050565b336106ae6000546001600160a01b031690565b6001600160a01b0316146107045760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102cb565b600255565b3361071c6000546001600160a01b031690565b6001600160a01b0316146107725760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102cb565b6001600160a01b0381166107ee5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102cb565b6107f7816107fa565b50565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561087457600080fd5b5035919050565b80356001600160a01b038116811461089257600080fd5b919050565b60008083601f8401126108a957600080fd5b50813567ffffffffffffffff8111156108c157600080fd5b6020830191508360208260051b85010111156108dc57600080fd5b9250929050565b60008060008060008060008060a0898b0312156108ff57600080fd5b6109088961087b565b975060208901359650604089013567ffffffffffffffff8082111561092c57600080fd5b6109388c838d01610897565b909850965060608b013591508082111561095157600080fd5b61095d8c838d01610897565b909650945060808b013591508082111561097657600080fd5b506109838b828c01610897565b999c989b5096995094979396929594505050565b60008083601f8401126109a957600080fd5b50813567ffffffffffffffff8111156109c157600080fd5b6020830191508360208285010111156108dc57600080fd5b60008060008060008060a087890312156109f257600080fd5b6109fb8761087b565b955060208701359450610a106040880161087b565b935060608701359250608087013567ffffffffffffffff811115610a3357600080fd5b610a3f89828a01610997565b979a9699509497509295939492505050565b60008060208385031215610a6457600080fd5b823567ffffffffffffffff811115610a7b57600080fd5b610a8785828601610997565b90969095509350505050565b600060208284031215610aa557600080fd5b6105f38261087b565b60008060008060608587031215610ac457600080fd5b610acd8561087b565b935060208501359250604085013567ffffffffffffffff811115610af057600080fd5b610afc87828801610997565b95989497509550505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b87811015610bb55782840389528135601e19883603018112610b6c57600080fd5b8701803567ffffffffffffffff811115610b8557600080fd5b803603891315610b9457600080fd5b610ba18682898501610b08565b9a87019a9550505090840190600101610b4b565b5091979650505050505050565b8183526000602080850194508260005b85811015610bfe576001600160a01b03610beb8361087b565b1687529582019590820190600101610bd2565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610c3b57600080fd5b8260051b8083602087013760009401602001938452509192915050565b608081526000885180608084015260005b81811015610c86576020818c0181015160a0868401015201610c69565b81811115610c9857600060a083860101525b50601f01601f1916820182810360a09081016020850152610cbc908201898b610b31565b90508281036040840152610cd1818789610bc2565b90508281036060840152610ce6818587610c09565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610d1c57610d1c610cf4565b500390565b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c06080830152610d5a60c083018587610b08565b90508260a083015298975050505050505050565b6000816000190483118215151615610d8857610d88610cf4565b500290565b60008219821115610da057610da0610cf4565b500190565b6001600160a01b0386168152846020820152608060408201526000610dce608083018587610b08565b9050826060830152969550505050505056fea2646970667358221220b3fee23e3a2090776399623a9c00c5305ab41fe9e275ec9027c25415e3691bba64736f6c63430008090033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208431c0c4ee52059b31268994c57744d74084b4770e7587b9fc576d98d1167c3e64736f6c63430008090033",
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

// TestRefundMetaData contains all meta data concerning the TestRefund contract.
var TestRefundMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"}],\"name\":\"MessageReceivedWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"},{\"internalType\":\"enumMessageSenderLib.BridgeType\",\"name\":\"_bridgeType\",\"type\":\"uint8\"}],\"name\":\"sendWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405162001bb338038062001bb3833981016040819052610031916100af565b61003a3361005f565b600180546001600160a01b0319166001600160a01b03929092169190911790556100df565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100c157600080fd5b81516001600160a01b03811681146100d857600080fd5b9392505050565b611ac480620000ef6000396000f3fe6080604052600436106100b15760003560e01c80638da5cb5b11610069578063ce35dd9a1161004e578063ce35dd9a1461018b578063f00f39ce1461019e578063f2fde38b146101b157600080fd5b80638da5cb5b14610139578063a1a227fa1461016b57600080fd5b8063547cad121161009a578063547cad12146100f157806355b9f7f7146101135780637ee1bd4a1461012657600080fd5b80631599d265146100b657806320be95f2146100de575b600080fd5b6100c96100c43660046114d3565b6101d1565b60405190151581526020015b60405180910390f35b6100c96100ec366004611536565b610289565b3480156100fd57600080fd5b5061011161010c36600461157a565b610353565b005b6101116101213660046115a5565b6103eb565b6101116101343660046115ee565b61041e565b34801561014557600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016100d5565b34801561017757600080fd5b50600154610153906001600160a01b031681565b6100c961019936600461169a565b61046e565b6100c96101ac36600461178b565b610543565b3480156101bd57600080fd5b506101116101cc36600461157a565b6105aa565b6001546000906001600160a01b031633146102335760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b60006102418385018561180e565b90507f0993eabb2ce4a554387afd1dd46cad1404aaf34ad41bba845836cb43105862a8868686866040516102789493929190611854565b60405180910390a195945050505050565b6001546000906001600160a01b031633146102e65760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161022a565b60006102f48385018561157a565b905061030a6001600160a01b038716828761069b565b7fee93b3d90bf7307985573224b309958a34f05c4bc6e541d3c2a9f903651cb3b68686868660405161033f9493929190611887565b60405180910390a150600195945050505050565b336103666000546001600160a01b031690565b6001600160a01b0316146103bc5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161022a565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60408051821515602082015260009101604051602081830303815290604052905061041884848334610730565b50505050565b6104336001600160a01b03871633308861074c565b60408051336020820152600091016040516020818303038152906040529050610463888888888888878934610784565b505050505050505050565b6001546000906001600160a01b031633146104cb5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161022a565b6000828060200190518101906104e191906118b0565b90506104f76001600160a01b038716828761069b565b7f6d93ac0161970e4a38513a159b2a5a056f09ea9e98b35d7a259513d95ea565448686858a8860405161052e959493929190611925565b60405180910390a15060019695505050505050565b6001546000906001600160a01b031633146105a05760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d6573736167652062757300000000000000604482015260640161022a565b9695505050505050565b336105bd6000546001600160a01b031690565b6001600160a01b0316146106135760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161022a565b6001600160a01b03811661068f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161022a565b610698816107bb565b50565b6040516001600160a01b03831660248201526044810182905261072b90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152610818565b505050565b600154610418908590859085906001600160a01b0316856108fd565b6040516001600160a01b03808516602483015283166044820152606481018290526104189085906323b872dd60e01b906084016106c7565b60006107ad8a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b610968565b9a9950505050505050505050565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600061086d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610a729092919063ffffffff16565b80519091501561072b578080602001905181019061088b9190611970565b61072b5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161022a565b604051634f9e72ad60e11b81526001600160a01b03831690639f3ce55a90839061092f9089908990899060040161198d565b6000604051808303818588803b15801561094857600080fd5b505af115801561095c573d6000803e3d6000fd5b50505050505050505050565b6000600184600581111561097e5761097e6119c8565b141561099c576109958b8b8b8b8b8b8b8a8a610a8b565b90506107ad565b60028460058111156109b0576109b06119c8565b14806109cd575060048460058111156109cb576109cb6119c8565b145b156109e357610995848c8c8c8c8c8b8a8a610c99565b60038460058111156109f7576109f76119c8565b1480610a1457506005846005811115610a1257610a126119c8565b145b15610a2a57610995848c8c8c8c8c8b8a8a610fcb565b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f7274656400000000000000604482015260640161022a565b6060610a81848460008561121f565b90505b9392505050565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b158015610ac757600080fd5b505afa158015610adb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aff91906118b0565b9050610b156001600160a01b038b16828b61135e565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b158015610b8b57600080fd5b505af1158015610b9f573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b8152600401610c579594939291906119de565b6000604051808303818588803b158015610c7057600080fd5b505af1158015610c84573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b60008060028b6005811115610cb057610cb06119c8565b1415610d2e57836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b158015610cef57600080fd5b505afa158015610d03573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2791906118b0565b9050610da2565b836001600160a01b031663c66a9c5a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d6757600080fd5b505afa158015610d7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9f91906118b0565b90505b610db66001600160a01b038a16828a61135e565b600060028c6005811115610dcc57610dcc6119c8565b1415610eef576040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a401600060405180830381600087803b158015610e3c57600080fd5b505af1158015610e50573d6000803e3d6000fd5b50505050308a8a8a8e8b46604051602001610ed29796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b604051602081830303815290604052805190602001209050610f95565b6040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a4015b602060405180830381600087803b158015610f5a57600080fd5b505af1158015610f6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f929190611a20565b90505b604051634289fbb360e01b81526001600160a01b03861690634289fbb3908690610c57908f908d90889088908e906004016119de565b60008060038b6005811115610fe257610fe26119c8565b141561106057836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561102157600080fd5b505afa158015611035573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105991906118b0565b90506110d4565b836001600160a01b03166395b12c276040518163ffffffff1660e01b815260040160206040518083038186803b15801561109957600080fd5b505afa1580156110ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d191906118b0565b90505b600060038c60058111156110ea576110ea6119c8565b14156111ca57604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff8916606483015283169063de790c7e90608401600060405180830381600087803b15801561115257600080fd5b505af1158015611166573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528e811b82166034840152604883018e90528f901b1660688201526001600160c01b031960c08b811b8216607c84015246901b166084820152608c019150610ed29050565b60405163a002930160e01b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063a00293019060a401610f40565b6060824710156112975760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161022a565b843b6112e55760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161022a565b600080866001600160a01b031685876040516113019190611a39565b60006040518083038185875af1925050503d806000811461133e576040519150601f19603f3d011682016040523d82523d6000602084013e611343565b606091505b509150915061135382828661141f565b979650505050505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b1580156113aa57600080fd5b505afa1580156113be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113e29190611a20565b6113ec9190611a55565b6040516001600160a01b03851660248201526044810182905290915061041890859063095ea7b360e01b906064016106c7565b6060831561142e575081610a84565b82511561143e5782518084602001fd5b8160405162461bcd60e51b815260040161022a9190611a7b565b6001600160a01b038116811461069857600080fd5b803567ffffffffffffffff8116811461148557600080fd5b919050565b60008083601f84011261149c57600080fd5b50813567ffffffffffffffff8111156114b457600080fd5b6020830191508360208285010111156114cc57600080fd5b9250929050565b600080600080606085870312156114e957600080fd5b84356114f481611458565b93506115026020860161146d565b9250604085013567ffffffffffffffff81111561151e57600080fd5b61152a8782880161148a565b95989497509550505050565b6000806000806060858703121561154c57600080fd5b843561155781611458565b935060208501359250604085013567ffffffffffffffff81111561151e57600080fd5b60006020828403121561158c57600080fd5b8135610a8481611458565b801515811461069857600080fd5b6000806000606084860312156115ba57600080fd5b83356115c581611458565b92506115d36020850161146d565b915060408401356115e381611597565b809150509250925092565b600080600080600080600060e0888a03121561160957600080fd5b873561161481611458565b9650602088013561162481611458565b9550604088013594506116396060890161146d565b93506116476080890161146d565b925060a088013563ffffffff8116811461166057600080fd5b915060c08801356006811061167457600080fd5b8091505092959891949750929550565b634e487b7160e01b600052604160045260246000fd5b600080600080600060a086880312156116b257600080fd5b85356116bd81611458565b945060208601356116cd81611458565b9350604086013592506116e26060870161146d565b9150608086013567ffffffffffffffff808211156116ff57600080fd5b818801915088601f83011261171357600080fd5b81358181111561172557611725611684565b604051601f8201601f19908116603f0116810190838211818310171561174d5761174d611684565b816040528281528b602084870101111561176657600080fd5b8260208601602083013760006020848301015280955050505050509295509295909350565b60008060008060008060a087890312156117a457600080fd5b86356117af81611458565b955060208701356117bf81611458565b9450604087013593506117d46060880161146d565b9250608087013567ffffffffffffffff8111156117f057600080fd5b6117fc89828a0161148a565b979a9699509497509295939492505050565b60006020828403121561182057600080fd5b8135610a8481611597565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b038516815267ffffffffffffffff841660208201526060604082015260006105a060608301848661182b565b6001600160a01b03851681528360208201526060604082015260006105a060608301848661182b565b6000602082840312156118c257600080fd5b8151610a8481611458565b60005b838110156118e85781810151838201526020016118d0565b838111156104185750506000910152565b600081518084526119118160208601602086016118cd565b601f01601f19169290920160200192915050565b60006001600160a01b03808816835286602084015260a0604084015261194e60a08401876118f9565b941660608301525067ffffffffffffffff919091166080909101529392505050565b60006020828403121561198257600080fd5b8151610a8481611597565b6001600160a01b038416815267ffffffffffffffff831660208201526060604082015260006119bf60608301846118f9565b95945050505050565b634e487b7160e01b600052602160045260246000fd5b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a0608083015261135360a08301846118f9565b600060208284031215611a3257600080fd5b5051919050565b60008251611a4b8184602087016118cd565b9190910192915050565b60008219821115611a7657634e487b7160e01b600052601160045260246000fd5b500190565b602081526000610a8460208301846118f956fea264697066735822122057fcb751977ea25cb9a80b5cc22dce00dbde98555459dac801f94049b2ce76e764736f6c63430008090033",
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
// Solidity: function executeMessage(address _receiver, uint64 _dstChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactor) ExecuteMessage(opts *bind.TransactOpts, _receiver common.Address, _dstChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "executeMessage", _receiver, _dstChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _receiver, uint64 _dstChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundSession) ExecuteMessage(_receiver common.Address, _dstChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessage(&_TestRefund.TransactOpts, _receiver, _dstChainId, _message)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x1599d265.
//
// Solidity: function executeMessage(address _receiver, uint64 _dstChainId, bytes _message) payable returns(bool)
func (_TestRefund *TestRefundTransactorSession) ExecuteMessage(_receiver common.Address, _dstChainId uint64, _message []byte) (*types.Transaction, error) {
	return _TestRefund.Contract.ExecuteMessage(&_TestRefund.TransactOpts, _receiver, _dstChainId, _message)
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

// Send is a paid mutator transaction binding the contract method 0x55b9f7f7.
//
// Solidity: function send(address _receiver, uint64 _dstChainId, bool _success) payable returns()
func (_TestRefund *TestRefundTransactor) Send(opts *bind.TransactOpts, _receiver common.Address, _dstChainId uint64, _success bool) (*types.Transaction, error) {
	return _TestRefund.contract.Transact(opts, "send", _receiver, _dstChainId, _success)
}

// Send is a paid mutator transaction binding the contract method 0x55b9f7f7.
//
// Solidity: function send(address _receiver, uint64 _dstChainId, bool _success) payable returns()
func (_TestRefund *TestRefundSession) Send(_receiver common.Address, _dstChainId uint64, _success bool) (*types.Transaction, error) {
	return _TestRefund.Contract.Send(&_TestRefund.TransactOpts, _receiver, _dstChainId, _success)
}

// Send is a paid mutator transaction binding the contract method 0x55b9f7f7.
//
// Solidity: function send(address _receiver, uint64 _dstChainId, bool _success) payable returns()
func (_TestRefund *TestRefundTransactorSession) Send(_receiver common.Address, _dstChainId uint64, _success bool) (*types.Transaction, error) {
	return _TestRefund.Contract.Send(&_TestRefund.TransactOpts, _receiver, _dstChainId, _success)
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

// TestRefundMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the TestRefund contract.
type TestRefundMessageReceivedIterator struct {
	Event *TestRefundMessageReceived // Event containing the contract specifics and raw log

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
func (it *TestRefundMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestRefundMessageReceived)
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
		it.Event = new(TestRefundMessageReceived)
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
func (it *TestRefundMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestRefundMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestRefundMessageReceived represents a MessageReceived event raised by the TestRefund contract.
type TestRefundMessageReceived struct {
	Receiver   common.Address
	DstChainId uint64
	Message    []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x0993eabb2ce4a554387afd1dd46cad1404aaf34ad41bba845836cb43105862a8.
//
// Solidity: event MessageReceived(address receiver, uint64 dstChainId, bytes message)
func (_TestRefund *TestRefundFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*TestRefundMessageReceivedIterator, error) {

	logs, sub, err := _TestRefund.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &TestRefundMessageReceivedIterator{contract: _TestRefund.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x0993eabb2ce4a554387afd1dd46cad1404aaf34ad41bba845836cb43105862a8.
//
// Solidity: event MessageReceived(address receiver, uint64 dstChainId, bytes message)
func (_TestRefund *TestRefundFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *TestRefundMessageReceived) (event.Subscription, error) {

	logs, sub, err := _TestRefund.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestRefundMessageReceived)
				if err := _TestRefund.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x0993eabb2ce4a554387afd1dd46cad1404aaf34ad41bba845836cb43105862a8.
//
// Solidity: event MessageReceived(address receiver, uint64 dstChainId, bytes message)
func (_TestRefund *TestRefundFilterer) ParseMessageReceived(log types.Log) (*TestRefundMessageReceived, error) {
	event := new(TestRefundMessageReceived)
	if err := _TestRefund.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestRefundMessageReceivedWithTransferIterator is returned from FilterMessageReceivedWithTransfer and is used to iterate over the raw logs and unpacked data for MessageReceivedWithTransfer events raised by the TestRefund contract.
type TestRefundMessageReceivedWithTransferIterator struct {
	Event *TestRefundMessageReceivedWithTransfer // Event containing the contract specifics and raw log

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
func (it *TestRefundMessageReceivedWithTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestRefundMessageReceivedWithTransfer)
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
		it.Event = new(TestRefundMessageReceivedWithTransfer)
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
func (it *TestRefundMessageReceivedWithTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestRefundMessageReceivedWithTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestRefundMessageReceivedWithTransfer represents a MessageReceivedWithTransfer event raised by the TestRefund contract.
type TestRefundMessageReceivedWithTransfer struct {
	Token      common.Address
	Amount     *big.Int
	Message    []byte
	Sender     common.Address
	SrcChainId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageReceivedWithTransfer is a free log retrieval operation binding the contract event 0x6d93ac0161970e4a38513a159b2a5a056f09ea9e98b35d7a259513d95ea56544.
//
// Solidity: event MessageReceivedWithTransfer(address token, uint256 amount, bytes message, address sender, uint64 srcChainId)
func (_TestRefund *TestRefundFilterer) FilterMessageReceivedWithTransfer(opts *bind.FilterOpts) (*TestRefundMessageReceivedWithTransferIterator, error) {

	logs, sub, err := _TestRefund.contract.FilterLogs(opts, "MessageReceivedWithTransfer")
	if err != nil {
		return nil, err
	}
	return &TestRefundMessageReceivedWithTransferIterator{contract: _TestRefund.contract, event: "MessageReceivedWithTransfer", logs: logs, sub: sub}, nil
}

// WatchMessageReceivedWithTransfer is a free log subscription operation binding the contract event 0x6d93ac0161970e4a38513a159b2a5a056f09ea9e98b35d7a259513d95ea56544.
//
// Solidity: event MessageReceivedWithTransfer(address token, uint256 amount, bytes message, address sender, uint64 srcChainId)
func (_TestRefund *TestRefundFilterer) WatchMessageReceivedWithTransfer(opts *bind.WatchOpts, sink chan<- *TestRefundMessageReceivedWithTransfer) (event.Subscription, error) {

	logs, sub, err := _TestRefund.contract.WatchLogs(opts, "MessageReceivedWithTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestRefundMessageReceivedWithTransfer)
				if err := _TestRefund.contract.UnpackLog(event, "MessageReceivedWithTransfer", log); err != nil {
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

// ParseMessageReceivedWithTransfer is a log parse operation binding the contract event 0x6d93ac0161970e4a38513a159b2a5a056f09ea9e98b35d7a259513d95ea56544.
//
// Solidity: event MessageReceivedWithTransfer(address token, uint256 amount, bytes message, address sender, uint64 srcChainId)
func (_TestRefund *TestRefundFilterer) ParseMessageReceivedWithTransfer(log types.Log) (*TestRefundMessageReceivedWithTransfer, error) {
	event := new(TestRefundMessageReceivedWithTransfer)
	if err := _TestRefund.contract.UnpackLog(event, "MessageReceivedWithTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Bin: "0x608060405234801561001057600080fd5b5060405161098838038061098883398101604081905261002f916100ad565b6100383361005d565b600180546001600160a01b0319166001600160a01b03929092169190911790556100dd565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100bf57600080fd5b81516001600160a01b03811681146100d657600080fd5b9392505050565b61089c806100ec6000396000f3fe6080604052600436106100965760003560e01c80638da5cb5b11610069578063ce35dd9a1161004e578063ce35dd9a1461015d578063f00f39ce1461015d578063f2fde38b1461017057600080fd5b80638da5cb5b1461010b578063a1a227fa1461013d57600080fd5b80631599d2651461009b57806320be95f2146100c3578063547cad12146100d6578063867fd811146100f8575b600080fd5b6100ae6100a93660046105c3565b610190565b60405190151581526020015b60405180910390f35b6100ae6100d1366004610624565b6101fe565b3480156100e257600080fd5b506100f66100f1366004610666565b61025b565b005b6100f661010636600461069e565b6102f3565b34801561011757600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016100ba565b34801561014957600080fd5b50600154610125906001600160a01b031681565b6100ae61016b366004610770565b610304565b34801561017c57600080fd5b506100f661018b366004610666565b61036b565b6001546000906001600160a01b031633146101f25760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b5060015b949350505050565b6001546000906001600160a01b031633146101f65760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016101e9565b3361026e6000546001600160a01b031690565b6001600160a01b0316146102c45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101e9565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b6102ff8383833461045c565b505050565b6001546000906001600160a01b031633146103615760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016101e9565b9695505050505050565b3361037e6000546001600160a01b031690565b6001600160a01b0316146103d45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016101e9565b6001600160a01b0381166104505760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016101e9565b6104598161047e565b50565b600154610478908590859085906001600160a01b0316856104db565b50505050565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b604051634f9e72ad60e11b81526001600160a01b03831690639f3ce55a90839061050d908990899089906004016107ef565b6000604051808303818588803b15801561052657600080fd5b505af115801561053a573d6000803e3d6000fd5b50505050505050505050565b80356001600160a01b038116811461055d57600080fd5b919050565b803567ffffffffffffffff8116811461055d57600080fd5b60008083601f84011261058c57600080fd5b50813567ffffffffffffffff8111156105a457600080fd5b6020830191508360208285010111156105bc57600080fd5b9250929050565b600080600080606085870312156105d957600080fd5b6105e285610546565b93506105f060208601610562565b9250604085013567ffffffffffffffff81111561060c57600080fd5b6106188782880161057a565b95989497509550505050565b6000806000806060858703121561063a57600080fd5b61064385610546565b935060208501359250604085013567ffffffffffffffff81111561060c57600080fd5b60006020828403121561067857600080fd5b61068182610546565b9392505050565b634e487b7160e01b600052604160045260246000fd5b6000806000606084860312156106b357600080fd5b6106bc84610546565b92506106ca60208501610562565b9150604084013567ffffffffffffffff808211156106e757600080fd5b818601915086601f8301126106fb57600080fd5b81358181111561070d5761070d610688565b604051601f8201601f19908116603f0116810190838211818310171561073557610735610688565b8160405282815289602084870101111561074e57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060008060008060a0878903121561078957600080fd5b61079287610546565b95506107a060208801610546565b9450604087013593506107b560608801610562565b9250608087013567ffffffffffffffff8111156107d157600080fd5b6107dd89828a0161057a565b979a9699509497509295939492505050565b6001600160a01b03841681526000602067ffffffffffffffff85168184015260606040840152835180606085015260005b8181101561083c57858101830151858201608001528201610820565b8181111561084e576000608083870101525b50601f01601f1916929092016080019594505050505056fea2646970667358221220ddc1d0e096d16419124e8619cf4ca3eb4dc86001364b3e55dedc7a2cab72a1e864736f6c63430008090033",
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
	Bin: "0x60806040523480156200001157600080fd5b5060405162002df838038062002df88339810160408190526200003491620000f9565b6200003f336200008c565b600180546001600160a01b039485166001600160a01b03199182161782559284166000908152600360205260409020805460ff191690911790556004805491909316911617905562000143565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000f457600080fd5b919050565b6000806000606084860312156200010f57600080fd5b6200011a84620000dc565b92506200012a60208501620000dc565b91506200013a60408501620000dc565b90509250925092565b612ca580620001536000396000f3fe6080604052600436106100ec5760003560e01c80635fdb3d8d1161008a578063ce35dd9a11610059578063ce35dd9a14610267578063d9b796de1461027a578063f00f39ce1461029a578063f2fde38b146102ad57600080fd5b80635fdb3d8d146102035780638da5cb5b14610216578063a1a227fa14610234578063a8a7e29c1461025457600080fd5b8063457bfa2f116100c6578063457bfa2f14610169578063496dbfd9146101a1578063547cad12146101c35780635b5a66a7146101e357600080fd5b80631599d265146100f857806320be95f21461012057806334e181251461012e57600080fd5b366100f357005b600080fd5b61010b610106366004612184565b6102cd565b60405190151581526020015b60405180910390f35b61010b6101063660046121e9565b34801561013a57600080fd5b5061015b61014936600461222d565b60026020526000908152604090205481565b604051908152602001610117565b34801561017557600080fd5b50600454610189906001600160a01b031681565b6040516001600160a01b039091168152602001610117565b3480156101ad57600080fd5b506101c16101bc36600461224a565b610337565b005b3480156101cf57600080fd5b506101c16101de36600461222d565b6103bc565b3480156101ef57600080fd5b506101c16101fe36600461222d565b610454565b6101c16102113660046122b0565b6104ec565b34801561022257600080fd5b506000546001600160a01b0316610189565b34801561024057600080fd5b50600154610189906001600160a01b031681565b6101c161026236600461237b565b61068d565b61010b6102753660046124a3565b610738565b34801561028657600080fd5b506101c161029536600461257a565b61099f565b61010b6102a83660046124a3565b610a33565b3480156102b957600080fd5b506101c16102c836600461222d565b610b07565b6001546000906001600160a01b0316331461032f5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b949350505050565b3361034a6000546001600160a01b031690565b6001600160a01b0316146103a05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001600160a01b03909116600090815260026020526040902055565b336103cf6000546001600160a01b031690565b6001600160a01b0316146104255760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b336104676000546001600160a01b031690565b6001600160a01b0316146104bd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6004805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b3332146105255760405162461bcd60e51b81526020600482015260076024820152664e6f7420454f4160c81b6044820152606401610326565b863410156105755760405162461bcd60e51b815260206004820152601360248201527f416d6f756e7420696e73756666696369656e74000000000000000000000000006044820152606401610326565b6004546001600160a01b031661058b86806125b3565b600081811061059c5761059c6125fd565b90506020020160208101906105b1919061222d565b6001600160a01b0316146105f85760405162461bcd60e51b815260206004820152600e60248201526d0e8ded6cadc40dad2e6dac2e8c6d60931b6044820152606401610326565b6004805460408051630d0e30db60e41b815290516001600160a01b039092169263d0e30db0928b92808301926000929182900301818588803b15801561063d57600080fd5b505af1158015610651573d6000803e3d6000fd5b50505050506106838888888861066690612637565b61066f89612637565b8888888f3461067e9190612719565b610bf8565b5050505050505050565b3332146106c65760405162461bcd60e51b81526020600482015260076024820152664e6f7420454f4160c81b6044820152606401610326565b61070d3330886106d688806125b3565b60008181106106e7576106e76125fd565b90506020020160208101906106fc919061222d565b6001600160a01b0316929190610e5e565b61072f87878761071c88612637565b61072588612637565b8787600034610bf8565b50505050505050565b6001546000906001600160a01b031633146107955760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610326565b6000828060200190518101906107ab9190612751565b8051518051919250906000906107c3576107c36125fd565b60200260200101516001600160a01b0316866001600160a01b0316146108775760405162461bcd60e51b815260206004820152604a60248201527f6272696467656420746f6b656e206d757374206265207468652073616d65206160448201527f732074686520666972737420746f6b656e20696e2064657374696e6174696f6e60648201527f2073776170207061746800000000000000000000000000000000000000000000608482015260a401610326565b60006108898260200151864687610efc565b8251515190915060009060019081101561091c5783516001906108ac908a610f35565b9350905080156108fd5784515180516108f491906108cc90600190612719565b815181106108dc576108dc6125fd565b60200260200101518487602001518860600151611082565b60019150610916565b61090e8a8a87602001516000611082565b889250600391505b50610954565b835151805161094d9190600090610935576109356125fd565b60200260200101518986602001518760600151611082565b5086905060015b7fccbb695db2dfb79e3f69bcf724d11b850c79b0ae358e62d60ed94044f38a2b9c838383604051610987939291906128d7565b60405180910390a15060019998505050505050505050565b336109b26000546001600160a01b031690565b6001600160a01b031614610a085760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001600160a01b03919091166000908152600360205260409020805460ff1916911515919091179055565b6001546000906001600160a01b03163314610a905760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610326565b600082806020019051810190610aa69190612751565b90506000610aba8260200151864687610efc565b90507fccbb695db2dfb79e3f69bcf724d11b850c79b0ae358e62d60ed94044f38a2b9c8160006002604051610af1939291906128d7565b60405180910390a1506000979650505050505050565b33610b1a6000546001600160a01b031690565b6001600160a01b031614610b705760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610326565b6001600160a01b038116610bec5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610326565b610bf5816111f4565b50565b855151610c475760405162461bcd60e51b815260206004820152601360248201527f656d7074792073726320737761702070617468000000000000000000000000006044820152606401610326565b8551805160009190610c5b90600190612719565b81518110610c6b57610c6b6125fd565b60200260200101519050600260008860000151600081518110610c9057610c906125fd565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020548911610d2d5760405162461bcd60e51b815260206004820152602b60248201527f616d6f756e74206d7573742062652067726561746572207468616e206d696e2060448201527f7377617020616d6f756e740000000000000000000000000000000000000000006064820152608401610326565b865151469060011080610d5457508067ffffffffffffffff168967ffffffffffffffff1614155b610da05760405162461bcd60e51b815260206004820152601360248201527f6e6f6f70206973206e6f7420616c6c6f776564000000000000000000000000006044820152606401610326565b8751518a9060011015610e0c576001610db98a8d610f35565b9250905080610e0a5760405162461bcd60e51b815260206004820152600f60248201527f7372632073776170206661696c656400000000000000000000000000000000006044820152606401610326565b505b8167ffffffffffffffff168a67ffffffffffffffff161415610e3c57610e378c8c848c8a8887611251565b610e50565b610e508c8c848d8d8d8d8d8d8d8d8c611330565b505050505050505050505050565b6040516001600160a01b0380851660248301528316604482015260648101829052610ef69085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526114c6565b50505050565b600084848484604051602001610f15949392919061291e565b604051602081830303815290604052805190602001209050949350505050565b6020808301516001600160a01b03166000908152600390915260408120548190819060ff16610f695760009250905061107b565b610fa98560200151858760000151600081518110610f8957610f896125fd565b60200260200101516001600160a01b03166115b09092919063ffffffff16565b84602001516001600160a01b03166338ed17398587606001518860000151308a604001516040518663ffffffff1660e01b8152600401610fed9594939291906129bc565b600060405180830381600087803b15801561100757600080fd5b505af192505050801561103c57506040513d6000823e601f3d908101601f1916820160405261103991908101906129f8565b60015b61104b5760009250905061107b565b6001816001835161105c9190612719565b8151811061106c5761106c6125fd565b60200260200101519350935050505b9250929050565b80156111e0576004546001600160a01b038581169116146110d65760405162461bcd60e51b815260206004820152600e60248201526d0e8ded6cadc40dad2e6dac2e8c6d60931b6044820152606401610326565b60048054604051632e1a7d4d60e01b81529182018590526001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561111a57600080fd5b505af115801561112e573d6000803e3d6000fd5b505050506000826001600160a01b03168461c35090604051600060405180830381858888f193505050503d8060008114611184576040519150601f19603f3d011682016040523d82523d6000602084013e611189565b606091505b50509050806111da5760405162461bcd60e51b815260206004820152601560248201527f6661696c656420746f2073656e64206e617469766500000000000000000000006044820152606401610326565b50610ef6565b610ef66001600160a01b0385168385611671565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6112656001600160a01b0383168883611671565b60003386898688604051602001611280959493929190612ac6565b6040516020818303038152906040528051906020012090507fdbd876e57420a75850bd1ce36de349d111ab4ff35e2c2379af8acc750376917481878988600001516000815181106112d3576112d36125fd565b6020908102919091018101516040805195865267ffffffffffffffff90941691850191909152918301526001600160a01b03908116606083015260808201859052851660a082015260c00160405180910390a15050505050505050565b86515161137f5760405162461bcd60e51b815260206004820152601360248201527f656d7074792064737420737761702070617468000000000000000000000000006044820152606401610326565b60006040518060800160405280898152602001336001600160a01b031681526020018767ffffffffffffffff1681526020018615158152506040516020016113c79190612b0c565b604051602081830303815290604052905060006113e6338d8d85610efc565b90506113fa8e85858e8b8d8860018d6116a1565b507f600b63623f19b18f8f0e7825ec78b3d9779d8124c76ad75b58aac7167f04f8d2818c8f8d60000151600081518110611436576114366125fd565b60209081029190910101518d51805161145190600190612719565b81518110611461576114616125fd565b60200260200101516040516114ae95949392919094855267ffffffffffffffff93909316602085015260408401919091526001600160a01b03908116606084015216608082015260a00190565b60405180910390a15050505050505050505050505050565b600061151b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116d89092919063ffffffff16565b8051909150156115ab57808060200190518101906115399190612b67565b6115ab5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610326565b505050565b604051636eb1769f60e11b81523060048201526001600160a01b038381166024830152600091839186169063dd62ed3e9060440160206040518083038186803b1580156115fc57600080fd5b505afa158015611610573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116349190612b84565b61163e9190612b9d565b6040516001600160a01b038516602482015260448101829052909150610ef690859063095ea7b360e01b90606401610e92565b6040516001600160a01b0383166024820152604481018290526115ab90849063a9059cbb60e01b90606401610e92565b60006116ca8a8a8a8a8a8a8a8a600160009054906101000a90046001600160a01b03168b6116f1565b9a9950505050505050505050565b60606116e784846000856117fb565b90505b9392505050565b600060018460058111156117075761170761289f565b14156117255761171e8b8b8b8b8b8b8b8a8a61193a565b90506116ca565b60028460058111156117395761173961289f565b1480611756575060048460058111156117545761175461289f565b145b1561176c5761171e848c8c8c8c8c8b8a8a611b48565b60038460058111156117805761178061289f565b148061179d5750600584600581111561179b5761179b61289f565b145b156117b35761171e848c8c8c8c8c8b8a8a611e7a565b60405162461bcd60e51b815260206004820152601960248201527f6272696467652074797065206e6f7420737570706f72746564000000000000006044820152606401610326565b6060824710156118735760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610326565b843b6118c15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610326565b600080866001600160a01b031685876040516118dd9190612bb5565b60006040518083038185875af1925050503d806000811461191a576040519150601f19603f3d011682016040523d82523d6000602084013e61191f565b606091505b509150915061192f8282866120ce565b979650505050505050565b600080836001600160a01b03166382980dc46040518163ffffffff1660e01b815260040160206040518083038186803b15801561197657600080fd5b505afa15801561198a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119ae9190612bd1565b90506119c46001600160a01b038b16828b6115b0565b60405163a5977fbb60e01b81526001600160a01b038c811660048301528b81166024830152604482018b905267ffffffffffffffff808b1660648401528916608483015263ffffffff881660a483015282169063a5977fbb9060c401600060405180830381600087803b158015611a3a57600080fd5b505af1158015611a4e573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528f811b821660348401528e901b166048820152605c81018c90526001600160c01b031960c08c811b8216607c8401528b811b8216608484015246901b16608c820152600092506094019050604051602081830303815290604052805190602001209050846001600160a01b0316634289fbb3858e8c86868c6040518763ffffffff1660e01b8152600401611b06959493929190612c1a565b6000604051808303818588803b158015611b1f57600080fd5b505af1158015611b33573d6000803e3d6000fd5b50939f9e505050505050505050505050505050565b60008060028b6005811115611b5f57611b5f61289f565b1415611bdd57836001600160a01b031663d8257d176040518163ffffffff1660e01b815260040160206040518083038186803b158015611b9e57600080fd5b505afa158015611bb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd69190612bd1565b9050611c51565b836001600160a01b031663c66a9c5a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c1657600080fd5b505afa158015611c2a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c4e9190612bd1565b90505b611c656001600160a01b038a16828a6115b0565b600060028c6005811115611c7b57611c7b61289f565b1415611d9e576040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a401600060405180830381600087803b158015611ceb57600080fd5b505af1158015611cff573d6000803e3d6000fd5b50505050308a8a8a8e8b46604051602001611d819796959493929190606097881b6bffffffffffffffffffffffff19908116825296881b87166014820152602881019590955260c093841b6001600160c01b031990811660488701529290961b909416605084015292811b831660648301529290921b16606c82015260740190565b604051602081830303815290604052805190602001209050611e44565b6040516308d18d8960e21b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063234636249060a4015b602060405180830381600087803b158015611e0957600080fd5b505af1158015611e1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e419190612b84565b90505b604051634289fbb360e01b81526001600160a01b03861690634289fbb3908690611b06908f908d90889088908e90600401612c1a565b60008060038b6005811115611e9157611e9161289f565b1415611f0f57836001600160a01b031663dfa2dbaf6040518163ffffffff1660e01b815260040160206040518083038186803b158015611ed057600080fd5b505afa158015611ee4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f089190612bd1565b9050611f83565b836001600160a01b03166395b12c276040518163ffffffff1660e01b815260040160206040518083038186803b158015611f4857600080fd5b505afa158015611f5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f809190612bd1565b90505b600060038c6005811115611f9957611f9961289f565b141561207957604051636f3c863f60e11b81526001600160a01b038b81166004830152602482018b90528c8116604483015267ffffffffffffffff8916606483015283169063de790c7e90608401600060405180830381600087803b15801561200157600080fd5b505af1158015612015573d6000803e3d6000fd5b50506040516bffffffffffffffffffffffff1930606090811b821660208401528e811b82166034840152604883018e90528f901b1660688201526001600160c01b031960c08b811b8216607c84015246901b166084820152608c019150611d819050565b60405163a002930160e01b81526001600160a01b038b81166004830152602482018b905267ffffffffffffffff808b1660448401528d821660648401528916608483015283169063a00293019060a401611def565b606083156120dd5750816116ea565b8251156120ed5782518084602001fd5b8160405162461bcd60e51b81526004016103269190612c5c565b6001600160a01b0381168114610bf557600080fd5b803561212781612107565b919050565b67ffffffffffffffff81168114610bf557600080fd5b60008083601f84011261215457600080fd5b50813567ffffffffffffffff81111561216c57600080fd5b60208301915083602082850101111561107b57600080fd5b6000806000806060858703121561219a57600080fd5b84356121a581612107565b935060208501356121b58161212c565b9250604085013567ffffffffffffffff8111156121d157600080fd5b6121dd87828801612142565b95989497509550505050565b600080600080606085870312156121ff57600080fd5b843561220a81612107565b935060208501359250604085013567ffffffffffffffff8111156121d157600080fd5b60006020828403121561223f57600080fd5b81356116ea81612107565b6000806040838503121561225d57600080fd5b823561226881612107565b946020939093013593505050565b60006080828403121561228857600080fd5b50919050565b803563ffffffff8116811461212757600080fd5b8015158114610bf557600080fd5b600080600080600080600080610100898b0312156122cd57600080fd5b88356122d881612107565b97506020890135965060408901356122ef8161212c565b9550606089013567ffffffffffffffff8082111561230c57600080fd5b6123188c838d01612276565b965060808b013591508082111561232e57600080fd5b5061233b8b828c01612276565b94505061234a60a08a0161228e565b925060c089013561235a8161212c565b915060e089013561236a816122a2565b809150509295985092959890939650565b600080600080600080600060e0888a03121561239657600080fd5b87356123a181612107565b96506020880135955060408801356123b88161212c565b9450606088013567ffffffffffffffff808211156123d557600080fd5b6123e18b838c01612276565b955060808a01359150808211156123f757600080fd5b506124048a828b01612276565b93505061241360a0890161228e565b915060c08801356124238161212c565b8091505092959891949750929550565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff8111828210171561246c5761246c612433565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561249b5761249b612433565b604052919050565b600080600080600060a086880312156124bb57600080fd5b85356124c681612107565b94506020868101356124d781612107565b94506040870135935060608701356124ee8161212c565b9250608087013567ffffffffffffffff8082111561250b57600080fd5b818901915089601f83011261251f57600080fd5b81358181111561253157612531612433565b612543601f8201601f19168501612472565b91508082528a8482850101111561255957600080fd5b80848401858401376000848284010152508093505050509295509295909350565b6000806040838503121561258d57600080fd5b823561259881612107565b915060208301356125a8816122a2565b809150509250929050565b6000808335601e198436030181126125ca57600080fd5b83018035915067ffffffffffffffff8211156125e557600080fd5b6020019150600581901b360382131561107b57600080fd5b634e487b7160e01b600052603260045260246000fd5b600067ffffffffffffffff82111561262d5761262d612433565b5060051b60200190565b60006080823603121561264957600080fd5b612651612449565b823567ffffffffffffffff81111561266857600080fd5b830136601f82011261267957600080fd5b8035602061268e61268983612613565b612472565b82815260059290921b830181019181810190368411156126ad57600080fd5b938201935b838510156126d45784356126c581612107565b825293820193908201906126b2565b8552506126e286820161211c565b90840152505060408381013590820152606092830135928101929092525090565b634e487b7160e01b600052601160045260246000fd5b60008282101561272b5761272b612703565b500390565b805161212781612107565b80516121278161212c565b8051612127816122a2565b6000602080838503121561276457600080fd5b825167ffffffffffffffff8082111561277c57600080fd5b908401906080828703121561279057600080fd5b612798612449565b8251828111156127a757600080fd5b8301608081890312156127b957600080fd5b6127c1612449565b8151848111156127d057600080fd5b82019350601f840189136127e357600080fd5b83516127f161268982612613565b81815260059190911b8501870190878101908b83111561281057600080fd5b958801955b8287101561283757865161282881612107565b82529588019590880190612815565b8352506128479050828701612730565b868201526040820151604082015260608201516060820152808352505061286f848401612730565b8482015261287f6040840161273b565b604082015261289060608401612746565b60608201529695505050505050565b634e487b7160e01b600052602160045260246000fd5b600481106128d357634e487b7160e01b600052602160045260246000fd5b9052565b838152602081018390526060810161032f60408301846128b5565b60005b8381101561290d5781810151838201526020016128f5565b83811115610ef65750506000910152565b6bffffffffffffffffffffffff198560601b16815260006001600160c01b0319808660c01b166014840152808560c01b16601c8401525082516129688160248501602087016128f2565b9190910160240195945050505050565b600081518084526020808501945080840160005b838110156129b15781516001600160a01b03168752958201959082019060010161298c565b509495945050505050565b85815284602082015260a0604082015260006129db60a0830186612978565b6001600160a01b0394909416606083015250608001529392505050565b60006020808385031215612a0b57600080fd5b825167ffffffffffffffff811115612a2257600080fd5b8301601f81018513612a3357600080fd5b8051612a4161268982612613565b81815260059190911b82018301908381019087831115612a6057600080fd5b928401925b8284101561192f57835182529284019290840190612a65565b6000815160808452612a936080850182612978565b90506001600160a01b03602084015116602085015260408301516040850152606083015160608501528091505092915050565b60006001600160a01b03808816835267ffffffffffffffff808816602085015281871660408501528086166060850152505060a0608083015261192f60a0830184612a7e565b602081526000825160806020840152612b2860a0840182612a7e565b90506001600160a01b03602085015116604084015267ffffffffffffffff60408501511660608401526060840151151560808401528091505092915050565b600060208284031215612b7957600080fd5b81516116ea816122a2565b600060208284031215612b9657600080fd5b5051919050565b60008219821115612bb057612bb0612703565b500190565b60008251612bc78184602087016128f2565b9190910192915050565b600060208284031215612be357600080fd5b81516116ea81612107565b60008151808452612c068160208601602086016128f2565b601f01601f19169290920160200192915050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525083606083015260a0608083015261192f60a0830184612bee565b6020815260006116ea6020830184612bee56fea2646970667358221220c2b1cd84a0fe75a2c66dc0525e850a920af1b4dc40158f059ade486d0d68f29a64736f6c63430008090033",
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
