package multinode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/transactor"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupBridgeTest() {
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
	SetupNewSgnEnv(p, true, false, false, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestBridge(t *testing.T) {
	t.Run("e2e-bridge", func(t *testing.T) {
		t.Run("bridgeTest", bridgeTest)
	})
}

// Test pegbridge
func bridgeTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("============ Test Bridge (Both cBridge and pegBridge) =============")
	setupBridgeTest()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	prepareValidators(t, transactor)

	tc.RunAllAndWait(
		func() {
			cbrTest(t, transactor)
		},
		func() {
			pbrTest(t, transactor)
		},
	)
}

func prepareValidators(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("================== Prepare validators start =================")

	log.Infoln("================== Setup validators ======================")
	// Make the stake amounts more realistic to test precision handling when distributing fee share
	vAmts := []*big.Int{
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
	}
	tc.SetupValidators(t, transactor, vAmts)

	log.Infoln("================== Setup bridge signers ======================")
	tc.CbrChain1.SetInitSigners(vAmts)
	tc.CbrChain2.SetInitSigners(vAmts)

	log.Infoln("================== Delegate from delegator 0 to all validators ======================")
	valAddrs := []eth.Addr{tc.ValEthAddrs[0], tc.ValEthAddrs[1], tc.ValEthAddrs[2]}
	dAmts := []*big.Int{
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
	}
	tc.MultiDelegate(tc.DelAuths[0], valAddrs, dAmts)
	for i := 0; i < 3; i++ {
		expDel := &stakingtypes.Delegation{
			DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[0]),
			ValidatorAddress: eth.Addr2Hex(tc.ValEthAddrs[i]),
			Shares:           sdk.NewIntFromBigInt(dAmts[i]),
		}
		tc.CheckDelegation(t, transactor, expDel)
	}

	log.Infoln("================== Prepare validators done =================")
}
