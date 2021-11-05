package common

import (
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	etherBaseKs  = EnvDir + "/keystore/etherbase.json"
	ChainID      = uint64(883)
	Geth2ChainID = uint64(884)

	EthClient     *ethclient.Client
	EtherBaseAuth *bind.TransactOpts
	ValAuths      []*bind.TransactOpts
	SignerAuths   []*bind.TransactOpts
	DelAuths      []*bind.TransactOpts

	Contracts    *eth.Contracts
	CelrAddr     eth.Addr
	CelrContract *eth.Erc20

	ValSgnAddrs []sdk.AccAddress

	CbrChain1, CbrChain2 *CbrChain
)

type CbrChain struct {
	ChainId    uint64
	Ec         *ethclient.Client
	Auth       *bind.TransactOpts // etherbase auth
	Users      []*TestEthClient
	Delegators []*TestEthClient
	// contract addr
	CbrAddr, USDTAddr eth.Addr
	CbrContract       *eth.BridgeContract
	USDTContract      *eth.Erc20
}

type TestEthClient struct {
	Address eth.Addr
	Auth    *bind.TransactOpts
	Signer  ethutils.Signer
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
