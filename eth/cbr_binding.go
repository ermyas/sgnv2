// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signers\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"}],\"name\":\"Relay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"curSigners\",\"type\":\"bytes\"}],\"name\":\"SignersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainid\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawDone\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"add_liquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_relayRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_curss\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"minsend\",\"type\":\"uint256[]\"}],\"name\":\"setMinSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"minslip\",\"type\":\"uint32\"}],\"name\":\"setMinSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_newss\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_curss\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_curss\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_wdmsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_curss\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdraws\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x60806040523480156200001157600080fd5b5060405162002609380380620026098339810160408190526200003491620000a2565b80516020820120600055620000493362000050565b5062000194565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006020808385031215620000b657600080fd5b82516001600160401b0380821115620000ce57600080fd5b818501915085601f830112620000e357600080fd5b815181811115620000f857620000f86200017e565b604051601f8201601f19908116603f011681019083821181831017156200012357620001236200017e565b8160405282815288868487010111156200013c57600080fd5b600093505b8284101562000160578484018601518185018701529285019262000141565b82841115620001725760008684830101525b98975050505050505050565b634e487b7160e01b600052604160045260246000fd5b61246580620001a46000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806382e00bfe1161008c578063d61ca5ae11610066578063d61ca5ae146101d1578063e09ab428146101e4578063f2fde38b14610207578063f8b30d7d1461021a57600080fd5b806382e00bfe146101905780638da5cb5b146101a3578063a5977fbb146101be57600080fd5b8063264e8893116100c8578063264e88931461012a5780633c64f04b1461013d5780636b68da2e14610175578063715018a61461018857600080fd5b80630857036f146100ef57806308992741146101045780631b24e7d014610117575b600080fd5b6101026100fd366004612138565b610248565b005b610102610112366004612092565b6103cf565b6101026101253660046121d1565b61047b565b610102610138366004612068565b6104c1565b61016061014b36600461211f565b60046020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b610102610183366004612138565b61058a565b61010261096f565b61010261019e366004612138565b6109a5565b6003546040516001600160a01b03909116815260200161016c565b6101026101cc366004611ffb565b610c20565b6101026101df366004612138565b610e46565b6101606101f236600461211f565b60026020526000908152604090205460ff1681565b610102610215366004611fe0565b611093565b61023a610228366004611fe0565b60056020526000908152604090205481565b60405190815260200161016c565b61025686868686868661058a565b600061029787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112e92505050565b90506000805b82515181101561037157816001600160a01b0316836000015182815181106102c7576102c7612403565b6020026020010151600001516001600160a01b03161161033c5760405162461bcd60e51b815260206004820152602560248201527f7369676e65722061646472657373206e6f7420696e20617363656e64696e672060448201526437b93232b960d91b60648201526084015b60405180910390fd5b825180518290811061035057610350612403565b60200260200101516000015191508080610369906123bc565b91505061029d565b5087876040516103829291906121ec565b6040519081900381206000557f81a95f093ac2e19ae9492d575f2580ed21ef667a6abccf751ce1eca7a9792705906103bd908a908a90612218565b60405180910390a15050505050505050565b6003546001600160a01b031633146103f95760405162461bcd60e51b81526004016103339061227a565b60005b838110156104745782828281811061041657610416612403565b905060200201356005600087878581811061043357610433612403565b90506020020160208101906104489190611fe0565b6001600160a01b031681526020810191909152604001600020558061046c816123bc565b9150506103fc565b5050505050565b6003546001600160a01b031633146104a55760405162461bcd60e51b81526004016103339061227a565b6006805463ffffffff191663ffffffff92909216919091179055565b6001805481906000906104de9083906001600160401b031661230d565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550610522333083856001600160a01b03166112bb909392919063ffffffff16565b600154604080516001600160401b034681168252909216602083015233908201526001600160a01b0383166060820152608081018290527fb37cfee98c7de6250a011396d78d9ffa228a81a381d7c46feefecfb40190eb959060a00160405180910390a15050565b838360405161059a9291906121ec565b6040518091039020600054146105f25760405162461bcd60e51b815260206004820152601860248201527f6d69736d617463682063757272656e74207369676e65727300000000000000006044820152606401610333565b600061063385858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061112e92505050565b90506000805b82515181101561068257825180518290811061065757610657612403565b6020026020010151602001518261066e91906122f5565b91508061067a816123bc565b915050610639565b5060006106f389896040516106989291906121ec565b60405180910390206040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905060008080805b878110156108fa5760006107698a8a8481811061071a5761071a612403565b905060200281019061072c91906122af565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508a9392505061132c9050565b9050836001600160a01b0316816001600160a01b0316116107cc5760405162461bcd60e51b815260206004820152601e60248201527f5369676e657273206e6f7420696e20617363656e64696e67206f7264657200006044820152606401610333565b8093505b87518051849081106107e4576107e4612403565b6020026020010151600001516001600160a01b0316816001600160a01b03161115610880576108146001846122f5565b885151909350831061087b5760405162461bcd60e51b815260206004820152602a60248201527f7369676e6572206e6f7420666f756e6420696e2063757272656e7420736f72746044820152696564207369676e65727360b01b6064820152608401610333565b6107d0565b875180518490811061089457610894612403565b6020026020010151600001516001600160a01b0316816001600160a01b031614156108e75787518051849081106108cd576108cd612403565b602002602001015160200151856108e491906122f5565b94505b50806108f2816123bc565b9150506106fb565b50600361090886600261235a565b6109129190612338565b61091d9060016122f5565b8310156109615760405162461bcd60e51b8152602060048201526012602482015271145d5bdc9d5b481b9bdd081c995858da195960721b6044820152606401610333565b505050505050505050505050565b6003546001600160a01b031633146109995760405162461bcd60e51b81526004016103339061227a565b6109a360006113d6565b565b6109b386868686868661058a565b60006109f487878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061142892505050565b9050468160a001516001600160401b031614610a4a5760405162461bcd60e51b81526020600482015260156024820152740c8e6e840c6d0c2d2dc92c840dcdee840dac2e8c6d605b1b6044820152606401610333565b8051602080830151604080850151606080870151608088015160a089015160e08a015160c0808c015197516001600160601b03199b871b8c169a81019a909a5297851b8a1660348a01529490931b9097166048870152605c8601526001600160c01b031995841b8616607c860152831b85166084850152608c840152901b90911660ac82015260009060b40160408051601f1981840301815291815281516020928301206000818152600490935291205490915060ff1615610b405760405162461bcd60e51b815260206004820152600f60248201526e7472616e736665722065786973747360881b6044820152606401610333565b600081815260046020908152604091829020805460ff19166001179055830151606084015191840151610b7f926001600160a01b0390911691906115e2565b7f6884a1d53b522ff69516b22f060b84db9968a409ec4438fc017fa1f4ed1fbee881836000015184602001518560400151866060015187608001518860c001518960e001516040516103bd9897969594939291909788526001600160a01b039687166020890152948616604088015292909416606086015260808501526001600160401b0392831660a085015290911660c083015260e08201526101000190565b6001600160a01b0385166000908152600560205260409020548411610c7a5760405162461bcd60e51b815260206004820152601060248201526f185b5bdd5b9d081d1bdbc81cdb585b1b60821b6044820152606401610333565b60065463ffffffff90811690821611610cce5760405162461bcd60e51b81526020600482015260166024820152751b585e081cdb1a5c1c1859d9481d1bdbc81cdb585b1b60521b6044820152606401610333565b6040516001600160601b031933606090811b8216602084015288811b8216603484015287901b166048820152605c81018590526001600160c01b031960c085811b8216607c84015284901b16608482015246608c82015260009060ac0160408051601f1981840301815291815281516020928301206000818152600490935291205490915060ff1615610d955760405162461bcd60e51b815260206004820152600f60248201526e7472616e736665722065786973747360881b6044820152606401610333565b6000818152600460205260409020805460ff19166001179055610dc36001600160a01b0387163330886112bb565b604080518281523360208201526001600160a01b038981168284015288166060820152608081018790526001600160401b0386811660a0830152851660c082015263ffffffff841660e082015290517f89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01918190036101000190a150505050505050565b610e5486868686868661058a565b6000610e9587878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061161792505050565b80519091506001600160401b03164614610ee85760405162461bcd60e51b81526020600482015260146024820152730c8e6e840c6d0c2d2dc92c840dad2e6dac2e8c6d60631b6044820152606401610333565b8051602080830151604080850151606086015160808701519251600096610f5696909594910160c095861b6001600160c01b031990811682529490951b9093166008850152606091821b6001600160601b03199081166010860152911b166024830152603882015260580190565b60408051601f1981840301815291815281516020928301206000818152600290935291205490915060ff1615610fce5760405162461bcd60e51b815260206004820152601a60248201527f776974686472617720616c7265616479207375636365656465640000000000006044820152606401610333565b60008181526002602052604090819020805460ff191660011790558201516080830151606084015161100b926001600160a01b03909116916115e2565b7f4661dd1e14044bc1a2e7df9289b91e8281a88bc46e5e8f000b2c2020544a5f4481836000015184602001518560400151866060015187608001516040516103bd969594939291909586526001600160401b0394851660208701529290931660408501526001600160a01b03908116606085015291909116608083015260a082015260c00190565b6003546001600160a01b031633146110bd5760405162461bcd60e51b81526004016103339061227a565b6001600160a01b0381166111225760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610333565b61112b816113d6565b50565b6040805160208082018352606082528251808401909352600080845290830184905290919061115e826001611738565b90508060018151811061117357611173612403565b60200260200101516001600160401b0381111561119257611192612419565b6040519080825280602002602001820160405280156111d757816020015b60408051808201909152600080825260208201528152602001906001900390816111b05790505b5083528051600090829060019081106111f2576111f2612403565b6020026020010181815250506000805b602084015151845110156112b257611219846117f1565b909250905081600114156112a3576112386112338561182b565b6118e7565b85600001518460018151811061125057611250612403565b60200260200101518151811061126857611268612403565b60200260200101819052508260018151811061128657611286612403565b60200260200101805180919061129b906123bc565b905250611202565b6112ad8482611981565b611202565b50505050919050565b6040516001600160a01b03808516602483015283166044820152606481018290526113269085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526119f3565b50505050565b60008151604114156113605760208201516040830151606084015160001a61135686828585611ac5565b93505050506113d0565b815160401415611388576020820151604083015161137f858383611c6e565b925050506113d0565b60405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610333565b92915050565b600380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b604080516101008101825260008082526020808301829052828401829052606083018290526080830182905260a0830182905260c0830182905260e0830182905283518085019094528184528301849052909190805b602083015151835110156115da57611495836117f1565b909250905081600114156114c4576114b46114af8461182b565b611c9a565b6001600160a01b0316845261147e565b81600214156114ec576114d96114af8461182b565b6001600160a01b0316602085015261147e565b8160031415611514576115016114af8461182b565b6001600160a01b0316604085015261147e565b81600414156115385761152e6115298461182b565b611ca5565b606085015261147e565b816005141561155d5761154a83611cdc565b6001600160401b0316608085015261147e565b81600614156115825761156f83611cdc565b6001600160401b031660a085015261147e565b81600714156115a75761159483611cdc565b6001600160401b031660c085015261147e565b81600814156115cb576115c16115bc8461182b565b611d5e565b60e085015261147e565b6115d58382611981565b61147e565b505050919050565b6040516001600160a01b03831660248201526044810182905261161290849063a9059cbb60e01b906064016112ef565b505050565b6040805160a08101825260008082526020808301829052828401829052606083018290526080830182905283518085019094528184528301849052909190805b602083015151835110156115da5761166e836117f1565b909250905081600114156116955761168583611cdc565b6001600160401b03168452611657565b81600214156116ba576116a783611cdc565b6001600160401b03166020850152611657565b81600314156116e2576116cf6114af8461182b565b6001600160a01b03166040850152611657565b816004141561170a576116f76114af8461182b565b6001600160a01b03166060850152611657565b81600514156117295761171f6115298461182b565b6080850152611657565b6117338382611981565b611657565b81516060906117488360016122f5565b6001600160401b0381111561175f5761175f612419565b604051908082528060200260200182016040528015611788578160200160208202803683370190505b5091506000805b602086015151865110156117e8576117a6866117f1565b809250819350505060018483815181106117c2576117c2612403565b602002602001018181516117d691906122f5565b9052506117e38682611981565b61178f565b50509092525090565b60008060006117ff84611cdc565b905061180c600882612338565b9250806007166005811115611823576118236123ed565b915050915091565b6060600061183883611cdc565b9050600081846000015161184c91906122f5565b905083602001515181111561186057600080fd5b816001600160401b0381111561187857611878612419565b6040519080825280601f01601f1916602001820160405280156118a2576020820181803683370190505b50602080860151865192955091818601919083010160005b858110156118dc5781810151838201526118d56020826122f5565b90506118ba565b505050935250919050565b6040805180820182526000808252602080830182905283518085019094528184528301849052909190805b602083015151835110156115da57611929836117f1565b90925090508160011415611953576119436114af8461182b565b6001600160a01b03168452611912565b8160021415611972576119686115298461182b565b6020850152611912565b61197c8382611981565b611912565b6000816005811115611995576119956123ed565b14156119a45761161282611cdc565b60028160058111156119b8576119b86123ed565b14156100ea5760006119c983611cdc565b905080836000018181516119dd91906122f5565b9052506020830151518351111561161257600080fd5b6000611a48826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611d769092919063ffffffff16565b8051909150156116125780806020019051810190611a6691906120fd565b6116125760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610333565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0821115611b425760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610333565b8360ff16601b1480611b5757508360ff16601c145b611bae5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b6064820152608401610333565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa158015611c02573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611c655760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610333565b95945050505050565b60006001600160ff1b03821660ff83901c601b01611c8e86828785611ac5565b925050505b9392505050565b60006113d082611d8d565b6000602082511115611cb657600080fd5b6020820151905081516020611ccb9190612379565b611cd690600861235a565b1c919050565b602080820151825181019091015160009182805b600a811015611d585783811a9150611d0981600761235a565b82607f16901b851794508160801660001415611d4657611d2a8160016122f5565b86518790611d399083906122f5565b9052509395945050505050565b80611d50816123bc565b915050611cf0565b50600080fd5b60008151602014611d6e57600080fd5b506020015190565b6060611d858484600085611dac565b949350505050565b60008151601414611d9d57600080fd5b5060200151600160601b900490565b606082471015611e0d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610333565b843b611e5b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610333565b600080866001600160a01b03168587604051611e7791906121fc565b60006040518083038185875af1925050503d8060008114611eb4576040519150601f19603f3d011682016040523d82523d6000602084013e611eb9565b606091505b5091509150611ec9828286611ed4565b979650505050505050565b60608315611ee3575081611c93565b825115611ef35782518084602001fd5b8160405162461bcd60e51b81526004016103339190612247565b80356001600160a01b0381168114611f2457600080fd5b919050565b60008083601f840112611f3b57600080fd5b5081356001600160401b03811115611f5257600080fd5b6020830191508360208260051b8501011115611f6d57600080fd5b9250929050565b60008083601f840112611f8657600080fd5b5081356001600160401b03811115611f9d57600080fd5b602083019150836020828501011115611f6d57600080fd5b803563ffffffff81168114611f2457600080fd5b80356001600160401b0381168114611f2457600080fd5b600060208284031215611ff257600080fd5b611c9382611f0d565b60008060008060008060c0878903121561201457600080fd5b61201d87611f0d565b955061202b60208801611f0d565b94506040870135935061204060608801611fc9565b925061204e60808801611fc9565b915061205c60a08801611fb5565b90509295509295509295565b6000806040838503121561207b57600080fd5b61208483611f0d565b946020939093013593505050565b600080600080604085870312156120a857600080fd5b84356001600160401b03808211156120bf57600080fd5b6120cb88838901611f29565b909650945060208701359150808211156120e457600080fd5b506120f187828801611f29565b95989497509550505050565b60006020828403121561210f57600080fd5b81518015158114611c9357600080fd5b60006020828403121561213157600080fd5b5035919050565b6000806000806000806060878903121561215157600080fd5b86356001600160401b038082111561216857600080fd5b6121748a838b01611f74565b9098509650602089013591508082111561218d57600080fd5b6121998a838b01611f74565b909650945060408901359150808211156121b257600080fd5b506121bf89828a01611f29565b979a9699509497509295939492505050565b6000602082840312156121e357600080fd5b611c9382611fb5565b8183823760009101908152919050565b6000825161220e818460208701612390565b9190910192915050565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6020815260008251806020840152612266816040850160208701612390565b601f01601f19169190910160400192915050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6000808335601e198436030181126122c657600080fd5b8301803591506001600160401b038211156122e057600080fd5b602001915036819003821315611f6d57600080fd5b60008219821115612308576123086123d7565b500190565b60006001600160401b0380831681851680830382111561232f5761232f6123d7565b01949350505050565b60008261235557634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615612374576123746123d7565b500290565b60008282101561238b5761238b6123d7565b500390565b60005b838110156123ab578181015183820152602001612393565b838111156113265750506000910152565b60006000198214156123d0576123d06123d7565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220725c69e68566f0b5151d2b2a5cb5dabcd8c8b6d31bf553b3bbf2fbbd36dbf67364736f6c63430008070033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _signers []byte) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, _signers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeCaller) MinSend(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "minSend", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinSend(&_Bridge.CallOpts, arg0)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.MinSend(&_Bridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "transfers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Transfers(&_Bridge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Transfers(&_Bridge.CallOpts, arg0)
}

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes _curss, bytes[] _sigs) view returns()
func (_Bridge *BridgeCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _curss []byte, _sigs [][]byte) error {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "verifySigs", _msg, _curss, _sigs)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes _curss, bytes[] _sigs) view returns()
func (_Bridge *BridgeSession) VerifySigs(_msg []byte, _curss []byte, _sigs [][]byte) error {
	return _Bridge.Contract.VerifySigs(&_Bridge.CallOpts, _msg, _curss, _sigs)
}

// VerifySigs is a free data retrieval call binding the contract method 0x6b68da2e.
//
// Solidity: function verifySigs(bytes _msg, bytes _curss, bytes[] _sigs) view returns()
func (_Bridge *BridgeCallerSession) VerifySigs(_msg []byte, _curss []byte, _sigs [][]byte) error {
	return _Bridge.Contract.VerifySigs(&_Bridge.CallOpts, _msg, _curss, _sigs)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeCaller) Withdraws(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "withdraws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Withdraws(&_Bridge.CallOpts, arg0)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Bridge *BridgeCallerSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Bridge.Contract.Withdraws(&_Bridge.CallOpts, arg0)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x264e8893.
//
// Solidity: function add_liquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "add_liquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x264e8893.
//
// Solidity: function add_liquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddLiquidity(&_Bridge.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x264e8893.
//
// Solidity: function add_liquidity(address _token, uint256 _amount) returns()
func (_Bridge *BridgeTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AddLiquidity(&_Bridge.TransactOpts, _token, _amount)
}

// Relay is a paid mutator transaction binding the contract method 0x82e00bfe.
//
// Solidity: function relay(bytes _relayRequest, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeTransactor) Relay(opts *bind.TransactOpts, _relayRequest []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "relay", _relayRequest, _curss, _sigs)
}

// Relay is a paid mutator transaction binding the contract method 0x82e00bfe.
//
// Solidity: function relay(bytes _relayRequest, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeSession) Relay(_relayRequest []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Relay(&_Bridge.TransactOpts, _relayRequest, _curss, _sigs)
}

// Relay is a paid mutator transaction binding the contract method 0x82e00bfe.
//
// Solidity: function relay(bytes _relayRequest, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeTransactorSession) Relay(_relayRequest []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Relay(&_Bridge.TransactOpts, _relayRequest, _curss, _sigs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeTransactor) Send(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "send", _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.Send(&_Bridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Bridge *BridgeTransactorSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Bridge.Contract.Send(&_Bridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] tokens, uint256[] minsend) returns()
func (_Bridge *BridgeTransactor) SetMinSend(opts *bind.TransactOpts, tokens []common.Address, minsend []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinSend", tokens, minsend)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] tokens, uint256[] minsend) returns()
func (_Bridge *BridgeSession) SetMinSend(tokens []common.Address, minsend []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSend(&_Bridge.TransactOpts, tokens, minsend)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] tokens, uint256[] minsend) returns()
func (_Bridge *BridgeTransactorSession) SetMinSend(tokens []common.Address, minsend []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSend(&_Bridge.TransactOpts, tokens, minsend)
}

// SetMinSlippage is a paid mutator transaction binding the contract method 0x1b24e7d0.
//
// Solidity: function setMinSlippage(uint32 minslip) returns()
func (_Bridge *BridgeTransactor) SetMinSlippage(opts *bind.TransactOpts, minslip uint32) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setMinSlippage", minslip)
}

// SetMinSlippage is a paid mutator transaction binding the contract method 0x1b24e7d0.
//
// Solidity: function setMinSlippage(uint32 minslip) returns()
func (_Bridge *BridgeSession) SetMinSlippage(minslip uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSlippage(&_Bridge.TransactOpts, minslip)
}

// SetMinSlippage is a paid mutator transaction binding the contract method 0x1b24e7d0.
//
// Solidity: function setMinSlippage(uint32 minslip) returns()
func (_Bridge *BridgeTransactorSession) SetMinSlippage(minslip uint32) (*types.Transaction, error) {
	return _Bridge.Contract.SetMinSlippage(&_Bridge.TransactOpts, minslip)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x0857036f.
//
// Solidity: function update(bytes _newss, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeTransactor) Update(opts *bind.TransactOpts, _newss []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "update", _newss, _curss, _sigs)
}

// Update is a paid mutator transaction binding the contract method 0x0857036f.
//
// Solidity: function update(bytes _newss, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeSession) Update(_newss []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Update(&_Bridge.TransactOpts, _newss, _curss, _sigs)
}

// Update is a paid mutator transaction binding the contract method 0x0857036f.
//
// Solidity: function update(bytes _newss, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeTransactorSession) Update(_newss []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Update(&_Bridge.TransactOpts, _newss, _curss, _sigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd61ca5ae.
//
// Solidity: function withdraw(bytes _wdmsg, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, _wdmsg []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", _wdmsg, _curss, _sigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd61ca5ae.
//
// Solidity: function withdraw(bytes _wdmsg, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeSession) Withdraw(_wdmsg []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _wdmsg, _curss, _sigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd61ca5ae.
//
// Solidity: function withdraw(bytes _wdmsg, bytes _curss, bytes[] _sigs) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(_wdmsg []byte, _curss []byte, _sigs [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, _wdmsg, _curss, _sigs)
}

// BridgeLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Bridge contract.
type BridgeLiquidityAddedIterator struct {
	Event *BridgeLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *BridgeLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLiquidityAdded)
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
		it.Event = new(BridgeLiquidityAdded)
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
func (it *BridgeLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLiquidityAdded represents a LiquidityAdded event raised by the Bridge contract.
type BridgeLiquidityAdded struct {
	ChainId  uint64
	Seqnum   uint64
	Provider common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xb37cfee98c7de6250a011396d78d9ffa228a81a381d7c46feefecfb40190eb95.
//
// Solidity: event LiquidityAdded(uint64 chainId, uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*BridgeLiquidityAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeLiquidityAddedIterator{contract: _Bridge.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xb37cfee98c7de6250a011396d78d9ffa228a81a381d7c46feefecfb40190eb95.
//
// Solidity: event LiquidityAdded(uint64 chainId, uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *BridgeLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLiquidityAdded)
				if err := _Bridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xb37cfee98c7de6250a011396d78d9ffa228a81a381d7c46feefecfb40190eb95.
//
// Solidity: event LiquidityAdded(uint64 chainId, uint64 seqnum, address provider, address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseLiquidityAdded(log types.Log) (*BridgeLiquidityAdded, error) {
	event := new(BridgeLiquidityAdded)
	if err := _Bridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
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
		it.Event = new(BridgeOwnershipTransferred)
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
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayIterator is returned from FilterRelay and is used to iterate over the raw logs and unpacked data for Relay events raised by the Bridge contract.
type BridgeRelayIterator struct {
	Event *BridgeRelay // Event containing the contract specifics and raw log

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
func (it *BridgeRelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelay)
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
		it.Event = new(BridgeRelay)
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
func (it *BridgeRelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelay represents a Relay event raised by the Bridge contract.
type BridgeRelay struct {
	TransferId    [32]byte
	Sender        common.Address
	Receiver      common.Address
	Token         common.Address
	Amount        *big.Int
	SrcChainId    uint64
	Nonce         uint64
	SrcTransferId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRelay is a free log retrieval operation binding the contract event 0x6884a1d53b522ff69516b22f060b84db9968a409ec4438fc017fa1f4ed1fbee8.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, uint64 nonce, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) FilterRelay(opts *bind.FilterOpts) (*BridgeRelayIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return &BridgeRelayIterator{contract: _Bridge.contract, event: "Relay", logs: logs, sub: sub}, nil
}

// WatchRelay is a free log subscription operation binding the contract event 0x6884a1d53b522ff69516b22f060b84db9968a409ec4438fc017fa1f4ed1fbee8.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, uint64 nonce, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) WatchRelay(opts *bind.WatchOpts, sink chan<- *BridgeRelay) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelay)
				if err := _Bridge.contract.UnpackLog(event, "Relay", log); err != nil {
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

// ParseRelay is a log parse operation binding the contract event 0x6884a1d53b522ff69516b22f060b84db9968a409ec4438fc017fa1f4ed1fbee8.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, uint64 nonce, bytes32 srcTransferId)
func (_Bridge *BridgeFilterer) ParseRelay(log types.Log) (*BridgeRelay, error) {
	event := new(BridgeRelay)
	if err := _Bridge.contract.UnpackLog(event, "Relay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the Bridge contract.
type BridgeSendIterator struct {
	Event *BridgeSend // Event containing the contract specifics and raw log

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
func (it *BridgeSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSend)
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
		it.Event = new(BridgeSend)
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
func (it *BridgeSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSend represents a Send event raised by the Bridge contract.
type BridgeSend struct {
	TransferId  [32]byte
	Sender      common.Address
	Receiver    common.Address
	Token       common.Address
	Amount      *big.Int
	DstChainId  uint64
	Nonce       uint64
	MaxSlippage uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) FilterSend(opts *bind.FilterOpts) (*BridgeSendIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &BridgeSendIterator{contract: _Bridge.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *BridgeSend) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSend)
				if err := _Bridge.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Bridge *BridgeFilterer) ParseSend(log types.Log) (*BridgeSend, error) {
	event := new(BridgeSend)
	if err := _Bridge.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeSignersUpdatedIterator is returned from FilterSignersUpdated and is used to iterate over the raw logs and unpacked data for SignersUpdated events raised by the Bridge contract.
type BridgeSignersUpdatedIterator struct {
	Event *BridgeSignersUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeSignersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeSignersUpdated)
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
		it.Event = new(BridgeSignersUpdated)
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
func (it *BridgeSignersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeSignersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeSignersUpdated represents a SignersUpdated event raised by the Bridge contract.
type BridgeSignersUpdated struct {
	CurSigners []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSignersUpdated is a free log retrieval operation binding the contract event 0x81a95f093ac2e19ae9492d575f2580ed21ef667a6abccf751ce1eca7a9792705.
//
// Solidity: event SignersUpdated(bytes curSigners)
func (_Bridge *BridgeFilterer) FilterSignersUpdated(opts *bind.FilterOpts) (*BridgeSignersUpdatedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeSignersUpdatedIterator{contract: _Bridge.contract, event: "SignersUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersUpdated is a free log subscription operation binding the contract event 0x81a95f093ac2e19ae9492d575f2580ed21ef667a6abccf751ce1eca7a9792705.
//
// Solidity: event SignersUpdated(bytes curSigners)
func (_Bridge *BridgeFilterer) WatchSignersUpdated(opts *bind.WatchOpts, sink chan<- *BridgeSignersUpdated) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeSignersUpdated)
				if err := _Bridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
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

// ParseSignersUpdated is a log parse operation binding the contract event 0x81a95f093ac2e19ae9492d575f2580ed21ef667a6abccf751ce1eca7a9792705.
//
// Solidity: event SignersUpdated(bytes curSigners)
func (_Bridge *BridgeFilterer) ParseSignersUpdated(log types.Log) (*BridgeSignersUpdated, error) {
	event := new(BridgeSignersUpdated)
	if err := _Bridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeWithdrawDoneIterator is returned from FilterWithdrawDone and is used to iterate over the raw logs and unpacked data for WithdrawDone events raised by the Bridge contract.
type BridgeWithdrawDoneIterator struct {
	Event *BridgeWithdrawDone // Event containing the contract specifics and raw log

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
func (it *BridgeWithdrawDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeWithdrawDone)
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
		it.Event = new(BridgeWithdrawDone)
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
func (it *BridgeWithdrawDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeWithdrawDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeWithdrawDone represents a WithdrawDone event raised by the Bridge contract.
type BridgeWithdrawDone struct {
	WithdrawId [32]byte
	Chainid    uint64
	Seqnum     uint64
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDone is a free log retrieval operation binding the contract event 0x4661dd1e14044bc1a2e7df9289b91e8281a88bc46e5e8f000b2c2020544a5f44.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 chainid, uint64 seqnum, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) FilterWithdrawDone(opts *bind.FilterOpts) (*BridgeWithdrawDoneIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return &BridgeWithdrawDoneIterator{contract: _Bridge.contract, event: "WithdrawDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawDone is a free log subscription operation binding the contract event 0x4661dd1e14044bc1a2e7df9289b91e8281a88bc46e5e8f000b2c2020544a5f44.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 chainid, uint64 seqnum, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) WatchWithdrawDone(opts *bind.WatchOpts, sink chan<- *BridgeWithdrawDone) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeWithdrawDone)
				if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
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

// ParseWithdrawDone is a log parse operation binding the contract event 0x4661dd1e14044bc1a2e7df9289b91e8281a88bc46e5e8f000b2c2020544a5f44.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 chainid, uint64 seqnum, address receiver, address token, uint256 amount)
func (_Bridge *BridgeFilterer) ParseWithdrawDone(log types.Log) (*BridgeWithdrawDone, error) {
	event := new(BridgeWithdrawDone)
	if err := _Bridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
