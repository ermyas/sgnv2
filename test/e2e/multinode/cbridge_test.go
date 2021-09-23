package multinode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/relayer"
	tc "github.com/celer-network/sgn-v2/test/common"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
)

var (
	Cbr1Addr     ethcommon.Address
	Cbr1Contract *eth.BridgeContract
	Cbr2Addr     ethcommon.Address
	Cbr2Contract *eth.BridgeContract
)

func setupCbridge(t *testing.T) {
	log.Infoln("Set up another mainchain for bridge")
	SetupMainchain2ForBridge()

	log.Infoln("Set up new sgn env")
	p := &tc.ContractParams{
		CelrAddr:              tc.CelrAddr,
		ProposalDeposit:       big.NewInt(1e18),
		VotePeriod:            big.NewInt(5),
		UnbondingPeriod:       big.NewInt(5),
		MaxBondedValidators:   big.NewInt(3),
		MinValidatorTokens:    big.NewInt(1e18),
		MinSelfDelegation:     big.NewInt(1e18),
		AdvanceNoticePeriod:   big.NewInt(1),
		ValidatorBondInterval: big.NewInt(0),
		MaxSlashFactor:        big.NewInt(1e5),
	}
	SetupNewSgnEnv(p, false)
	tc.SleepWithLog(10, "sgn being ready")

	transactor := tc.NewTestTransactor(
		t,
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	amt1 := big.NewInt(3e18)
	amt2 := big.NewInt(2e18)
	amt3 := big.NewInt(2e18)
	amts := []*big.Int{amt1, amt2, amt3}
	SetupValidators(t, transactor, amts)

	validators, err := stakingcli.QueryValidators(transactor.CliCtx)
	tc.ChkErr(err, "failed to query validators contract")
	signers, _ := proto.Marshal(relayer.GetSortedSigners(validators))
	Cbr1Addr, Cbr1Contract = tc.DeployBridgeContract(tc.EthClient, tc.EtherBaseAuth, signers)
	Cbr2Addr, Cbr2Contract = tc.DeployBridgeContract(tc.EthClient2, tc.EtherBaseAuth2, signers)
}

func TestCbridge(t *testing.T) {
	t.Run("e2e-cbridge", func(t *testing.T) {
		t.Run("cbridgeTest", cbridgeTest)
	})
}

// Test cbridge
func cbridgeTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test cbridge ===========================")

	setupCbridge(t)
}
