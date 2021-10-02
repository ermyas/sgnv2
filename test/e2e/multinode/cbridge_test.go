package multinode

import (
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
	SetupNewSgnEnv(p, false, true)
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
	amts := []*big.Int{big.NewInt(3e18)}
	tc.SetupValidators(t, transactor, amts)
	tc.CbrClient1.SetInitSigners(amts)
	tc.CbrClient2.SetInitSigners(amts)

	log.Infoln("======================== Query ===========================")
	resp, err := cbrcli.QueryChainTokensConfig(transactor.CliCtx, &cbrtypes.ChainTokensConfigRequest{})
	tc.ChkErr(err, "cli Query")
	assert.True(t, len(resp.ChainTokens) > 0)
	log.Infoln("QueryChainTokensConfig resp:", resp.String())

	chainTokens := make([]*cbrtypes.ChainTokenAddrPair, 0)
	chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
		ChainId:   tc.ChainID,
		TokenAddr: tc.CbrClient1.USDTAddr.Hex(),
	})
	chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
		ChainId:   tc.Geth2ChainID,
		TokenAddr: tc.CbrClient2.USDTAddr.Hex(),
	})
	resp2, err := cbrcli.QueryLiquidityDetailList(transactor.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     "0x58b529F9084D7eAA598EB3477Fe36064C5B7bbC1",
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	assert.True(t, len(resp2.LiquidityDetail) > 0)
	log.Infoln("QueryLiquidityDetailList resp:", resp2.String())

	log.Infoln("======================== Add liquidity on chain 1 ===========================")
	addAmt := big.NewInt(5 * 1e10)
	err = tc.CbrClient1.Approve(addAmt)
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
	detail, err := tc.GetWithdrawDetail(transactor, wdSeq)
	tc.ChkErr(err, "client1 get withdrawdetail")
	curss, err := tc.GetCurSortedSigners(transactor, chainId)
	tc.ChkErr(err, "client1 GetCurSortedSigners")
	err = tc.CbrClient1.OnchainWithdraw(detail, curss)
	tc.ChkErr(err, "client1 onchain withdraw")
	// todo: more cases, eg. lp2 withdraw from chain1 after xfer

	log.Infoln("======================== LP claim farming reward on-chain ===========================")
	err = tc.StartClaimAll(transactor, eth.Addr2Hex(tc.CbrClient1.Auth.From))
	tc.ChkErr(err, "client1 start claim all")
	// now sleep and get stuff to send onchain
	time.Sleep(time.Second * 10)
	info, err := tc.GetRewardClaimInfo(transactor, eth.Addr2Hex(tc.CbrClient1.Auth.From))
	tc.ChkErr(err, "client1 get reward info")
	assert.Len(t, info.RewardClaimDetailsList, 1, "must have 1 RewardClaimDetails")
	assert.Len(t, info.RewardClaimDetailsList[0].Signatures, 1, "node0 should sign")
	err = tc.CbrClient1.OnchainClaimRewards(&info.RewardClaimDetailsList[0])
	tc.ChkErr(err, "client1 onchain claim rewards")
}

func cbrSignersTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test cBridge signers ===========================")
	setupCbridge()

	chainId1, chainId2 := tc.ChainID, tc.Geth2ChainID

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	log.Infoln("================== Init bridge signers ======================")
	initSignerPowers := []*big.Int{big.NewInt(1e18)}
	tc.CbrClient1.SetInitSigners(initSignerPowers)
	tc.CbrClient2.SetInitSigners(initSignerPowers)
	tc.Sleep(5)
	expSigners := genSortedSigners(initSignerPowers)
	tc.CheckChainSigners(t, transactor, chainId1, expSigners)
	tc.CheckChainSigners(t, transactor, chainId2, expSigners)

	log.Infoln("================== Add validators ======================")
	tc.AddValidator(t, transactor, 0, big.NewInt(3e18), eth.CommissionRate(0.03))

	tc.AddValidator(t, transactor, 1, big.NewInt(2e18), eth.CommissionRate(0.02))
	expSigners = genSortedSigners([]*big.Int{big.NewInt(3e18), big.NewInt(2e18)})
	tc.CheckLatestSigners(t, transactor, expSigners)
	tc.CheckChainSigners(t, transactor, chainId1, expSigners)
	tc.CheckChainSigners(t, transactor, chainId2, expSigners)

	tc.AddValidator(t, transactor, 2, big.NewInt(4e18), eth.CommissionRate(0.01))
	expSigners = genSortedSigners([]*big.Int{big.NewInt(3e18), big.NewInt(2e18), big.NewInt(4e18)})
	tc.CheckLatestSigners(t, transactor, expSigners)
	tc.CheckChainSigners(t, transactor, chainId1, expSigners)
	tc.CheckChainSigners(t, transactor, chainId2, expSigners)
}

func genSortedSigners(amts []*big.Int) *cbrtypes.SortedSigners {
	ss := new(cbrtypes.SortedSigners)
	for i, amt := range amts {
		ss.Signers = append(ss.Signers,
			&cbrtypes.AddrAmt{
				Addr: tc.ValSignerAddrs[i].Bytes(),
				Amt:  amt.Bytes(),
			})
	}
	return ss
}
