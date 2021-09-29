package multinode

import (
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	tc "github.com/celer-network/sgn-v2/test/common"
)

func setupCbridge() {
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
	SetupNewSgnEnv(p, false, true)
	tc.SleepWithLog(10, "sgn being ready")
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
	setupCbridge()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	log.Infoln("================== Setup validators and bridge signers ======================")
	amts := []*big.Int{big.NewInt(3e18)}
	SetupValidators(transactor, amts)
	tc.CbrClient1.SetInitSigners(amts)
	tc.CbrClient2.SetInitSigners(amts)

	log.Infoln("======================== Add liquidity on chain 1 ===========================")
	addAmt := big.NewInt(5 * 1e10)
	err := tc.CbrClient1.Approve(addAmt)
	tc.ChkErr(err, "client1 approve")
	err = tc.CbrClient1.AddLiq(addAmt)
	tc.ChkErr(err, "client1 addliq")
	tc.CheckAddLiquidityStatus(transactor, tc.ChainID, 1)

	log.Infoln("======================== Add liquidity on chain 2 ===========================")
	err = tc.CbrClient2.Approve(addAmt)
	tc.ChkErr(err, "client2 approve")
	err = tc.CbrClient2.AddLiq(addAmt)
	tc.ChkErr(err, "client2 addliq")
	tc.CheckAddLiquidityStatus(transactor, tc.Geth2ChainID, 1)

	log.Infoln("======================== Xfer ===========================")
	xferAmt := big.NewInt(1e10)
	err = tc.CbrClient1.Approve(xferAmt)
	tc.ChkErr(err, "client1 approve")
	xferId, err := tc.CbrClient1.Send(xferAmt, tc.ValEthAddrs[0], tc.Geth2ChainID, 1)
	tc.ChkErr(err, "client1 send")
	tc.CheckXfer(transactor, xferId[:])

	log.Infoln("======================== LP withdraw liquidity ===========================")
	chainId := uint64(883)
	wdSeq, err := tc.CbrClient1.StartWithdraw(transactor, chainId, big.NewInt(1e10))
	tc.ChkErr(err, "client1 start withdraw")
	log.Info("withdraw seqnum: ", wdSeq)
	// now sleep and get stuff to send onchain
	time.Sleep(time.Second * 10)
	detail, err := tc.CbrClient1.GetWithdrawDetail(transactor, wdSeq)
	tc.ChkErr(err, "client1 get withdrawdetail")
	curss, err := tc.CbrClient1.GetCurSortedSigners(transactor, chainId)
	tc.ChkErr(err, "client1 GetCurSortedSigners")
	err = tc.CbrClient1.OnchainWithdraw(detail, curss)
	tc.ChkErr(err, "client1 onchain withdraw")
	// todo: more cases, eg. lp2 withdraw from chain1 after xfer
}
