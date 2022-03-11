package common

import (
	"math/big"

	"github.com/celer-network/cbridge-flow/signer"
	flowutils "github.com/celer-network/cbridge-flow/utils"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	etherBaseKs = EnvDir + "/keystore/etherbase.json"
	flowBaseKs  = EnvDir + "/keystore/flowbase.json"

	EthClient     *ethclient.Client
	EtherBaseAuth *bind.TransactOpts
	ValAuths      []*bind.TransactOpts
	DelAuths      []*bind.TransactOpts

	Contracts    *eth.Contracts
	CelrAddr     eth.Addr
	CelrContract *eth.Erc20

	ValSgnAddrs []sdk.AccAddress

	CbrChain1, CbrChain2, CbrChain3 *CbrChain

	FlowServiceAccountClient  *flowutils.FlowCbrClient
	FlowContractAccountClient *flowutils.FlowCbrClient
	FlowUserAccountClient     *flowutils.FlowCbrClient
	FlowServiceAccountSigner  *signer.FlowSigner
)

type CbrChain struct {
	ChainId          uint64
	Ec               *ethclient.Client
	Auth             *bind.TransactOpts   // etherbase auth
	Transactor       *ethutils.Transactor // etherbase transactor
	Users            []*TestEthClient
	Validators       []*TestEthClient
	ValidatorSigners []*TestEthClient
	Delegators       []*TestEthClient

	// contract addr
	CbrContract      *eth.BridgeContract
	CbrAddr          eth.Addr
	WdInboxContract  *eth.WdInboxContract
	CLPContract      *eth.CLPContract
	WdiAddr, CLPAddr eth.Addr

	PegBridgeContract *eth.PegBridgeContract
	PegVaultContract  *eth.PegVaultContract
	PegBridgeAddr     eth.Addr
	PegVaultAddr      eth.Addr

	PegBridgeV2Contract *eth.PegBridgeV2Contract
	PegVaultV2Contract  *eth.PegVaultV2Contract
	PegBridgeV2Addr     eth.Addr
	PegVaultV2Addr      eth.Addr

	USDTContract, UNIContract *eth.BridgeTestToken
	USDTAddr, UNIAddr         eth.Addr

	MessageBusContract      *eth.MessageBus
	MessageBusAddr          eth.Addr
	BatchTransferContract   *eth.BatchTransfer
	BatchTransferAddr       eth.Addr
	TransferMessageContract *eth.TransferMessage
	TransferMessageAddr     eth.Addr
	TestRefundContract      *eth.TestRefund
	TestRefundAddr          eth.Addr
}

type TestEthClient struct {
	Address    eth.Addr
	Auth       *bind.TransactOpts
	Signer     ethutils.Signer
	Transactor *ethutils.Transactor
}

func (c *TestEthClient) SignMsg(data []byte) []byte {
	ret, _ := c.Signer.SignEthMessage(data)
	return ret
}

type ContractParams struct {
	CelrAddr              eth.Addr
	ProposalDeposit       *big.Int
	VotePeriod            *big.Int
	UnbondingPeriod       *big.Int
	MaxBondedValidators   *big.Int
	MinValidatorTokens    *big.Int
	MinSelfDelegation     *big.Int
	AdvanceNoticePeriod   *big.Int
	ValidatorBondInterval *big.Int
	MaxSlashFactor        *big.Int

	// TODO: Remove from here
	StartGateway bool
}

// Killable is object that has Kill() func
type Killable interface {
	Kill() error
}

func TearDown(tokill []Killable) {
	log.Info("Tear down Killables ing...")
	for _, p := range tokill {
		ChkErr(p.Kill(), "kill process error")
	}
}

func ChkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
