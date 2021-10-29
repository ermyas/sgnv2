package multinode

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/stretchr/testify/assert"
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
	SetupNewSgnEnv(p, true, false, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestCbridge(t *testing.T) {
	t.Run("e2e-cbridge", func(t *testing.T) {
		t.Run("cbridgeTest", cbridgeTest)
		//t.Run("cbrSignersTest", cbrSignersTest)
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
	amts := []*big.Int{big.NewInt(2e18), big.NewInt(2e18), big.NewInt(2e18)}
	tc.SetupValidators(t, transactor, amts)
	tc.CbrChain1.SetInitSigners(amts)
	tc.CbrChain2.SetInitSigners(amts)
	expSigners := genSortedSigners(amts)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	log.Infoln("======================== Query ===========================")
	resp, err := cbrcli.QueryChainTokensConfig(transactor.CliCtx, &cbrtypes.ChainTokensConfigRequest{})
	tc.ChkErr(err, "cli Query")
	assert.True(t, len(resp.ChainTokens) > 0)
	log.Infoln("QueryChainTokensConfig resp:", resp.String())

	chainTokens := make([]*cbrtypes.ChainTokenAddrPair, 0)
	chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
		ChainId:   tc.CbrChain1.ChainId,
		TokenAddr: tc.CbrChain1.USDTAddr.Hex(),
	})
	chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
		ChainId:   tc.CbrChain2.ChainId,
		TokenAddr: tc.CbrChain2.USDTAddr.Hex(),
	})
	res, err := cbrcli.QueryLiquidityDetailList(transactor.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     eth.ZeroAddrHex,
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	assert.True(t, len(res.LiquidityDetail) > 0)
	log.Infoln("QueryLiquidityDetailList resp:", res.String())

	log.Infoln("======================== Add liquidity on chain 1 ===========================")
	addAmt := big.NewInt(5 * 1e10)
	var i uint64
	for i = 0; i < 2; i++ {
		err = tc.CbrChain1.Approve(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 approve", i))
		err = tc.CbrChain1.AddLiq(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 addliq", i))
		tc.CheckAddLiquidityStatus(transactor, tc.CbrChain1.ChainId, i+1)
	}

	log.Infoln("======================== Add liquidity on chain 2 ===========================")
	for i = 0; i < 2; i++ {
		err = tc.CbrChain2.Approve(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain2 approve", i))
		err = tc.CbrChain2.AddLiq(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain2 addliq", i))
		tc.CheckAddLiquidityStatus(transactor, tc.CbrChain1.ChainId, i+1)
	}
	res, err = cbrcli.QueryLiquidityDetailList(transactor.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     tc.ClientEthAddrs[0].Hex(),
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	log.Infoln("QueryLiquidityDetailList resp:", res.String())

	log.Infoln("======================== Xfer ===========================")
	xferAmt := big.NewInt(1e10)
	err = tc.CbrChain1.Approve(0, xferAmt)
	tc.ChkErr(err, "u0 chain1 approve")
	xferId, err := tc.CbrChain1.Send(0, xferAmt, tc.CbrChain2.ChainId, 1)
	tc.ChkErr(err, "u0 chain1 send")
	tc.CheckXfer(transactor, xferId[:])

	log.Infoln("======================== LP withdraw liquidity ===========================")
	reqid := uint64(time.Now().Unix())
	wdLq1 := tc.CbrChain1.GetWithdrawLq(20000000) // withdraw 20%
	wdLq2 := tc.CbrChain2.GetWithdrawLq(10000000) // withdraw 10%
	err = tc.CbrChain1.StartWithdraw(transactor, reqid, 0, wdLq1, wdLq2)
	tc.ChkErr(err, "u0 chain1 start withdraw")
	log.Infoln("withdraw reqid:", reqid)
	detail := tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Users[0].Address, reqid, 3)
	curss, err := tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw")
	// todo: more cases, eg. lp2 withdraw from chain1 after xfer

	log.Infoln("======================== LP claim farming reward on-chain ===========================")
	err = tc.StartClaimFarmingRewards(transactor, 0)
	tc.ChkErr(err, "u0 start claim all farming rewards")
	info := tc.GetFarmingRewardClaimInfoWithSigs(transactor, 0, 3)
	err = tc.OnchainClaimFarmingRewards(&info.RewardClaimDetailsList[0])
	tc.ChkErr(err, "u0 onchain claim farming rewards")
}

func cbrSignersTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test cBridge signers ===========================")
	setupCbridge()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	log.Infoln("================== Init bridge signers ======================")
	initSignerPowers := []*big.Int{big.NewInt(1e18)}
	tc.CbrChain1.SetInitSigners(initSignerPowers)
	tc.CbrChain2.SetInitSigners(initSignerPowers)
	tc.Sleep(5)
	expSigners := genSortedSigners(initSignerPowers)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	log.Infoln("================== Add validators ======================")
	tc.AddValidator(t, transactor, 0, big.NewInt(3e18), eth.CommissionRate(0.03))

	tc.AddValidator(t, transactor, 1, big.NewInt(2e18), eth.CommissionRate(0.02))
	expSigners = genSortedSigners([]*big.Int{big.NewInt(3e18), big.NewInt(2e18)})
	tc.CheckLatestSigners(t, transactor, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	tc.AddValidator(t, transactor, 2, big.NewInt(4e18), eth.CommissionRate(0.01))
	expSigners = genSortedSigners([]*big.Int{big.NewInt(3e18), big.NewInt(2e18), big.NewInt(4e18)})
	tc.CheckLatestSigners(t, transactor, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)
}

func genSortedSigners(amts []*big.Int) []*cbrtypes.Signer {
	var ss []*cbrtypes.Signer
	for i, amt := range amts {
		ss = append(ss,
			&cbrtypes.Signer{
				Addr:  tc.ValSignerAddrs[i].Bytes(),
				Power: amt.Bytes(),
			})
	}
	return ss
}
