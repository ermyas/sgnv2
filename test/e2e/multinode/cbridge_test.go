package multinode

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
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
	numVals := len(amts)
	tc.SetupValidators(t, transactor, amts)
	tc.CbrChain1.SetInitSigners(amts)
	tc.CbrChain2.SetInitSigners(amts)
	expSigners := genSortedSigners([]eth.Addr{tc.ValSignerAddrs[0], tc.ValSignerAddrs[1], tc.ValSignerAddrs[2]}, amts)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	log.Infoln("================== Delegate from delegator 0 to all validators ======================")
	for i := 0; i < numVals; i++ {
		tc.Delegate(tc.DelAuths[0], tc.ValEthAddrs[i], amts[i])
	}
	for i := 0; i < numVals; i++ {
		expDel := &stakingtypes.Delegation{
			DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[0]),
			ValidatorAddress: eth.Addr2Hex(tc.ValEthAddrs[i]),
			Shares:           sdk.NewIntFromBigInt(amts[i]),
		}
		tc.CheckDelegation(t, transactor, expDel)
	}

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
		tc.CheckAddLiquidityStatus(transactor, tc.CbrChain2.ChainId, i+1)
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

	res, err = cbrcli.QueryLiquidityDetailList(transactor.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     tc.ClientEthAddrs[0].Hex(),
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	log.Infoln("QueryLiquidityDetailList resp:", res.String())
	res, err = cbrcli.QueryLiquidityDetailList(transactor.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     tc.ClientEthAddrs[1].Hex(),
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	log.Infoln("QueryLiquidityDetailList resp:", res.String())

	log.Infoln("======================== LP withdraw liquidity ===========================")
	reqid := uint64(time.Now().Unix())
	wdLq1 := tc.CbrChain1.GetWithdrawLq(20000000) // withdraw 20%
	wdLq2 := tc.CbrChain2.GetWithdrawLq(10000000) // withdraw 10%
	err = tc.CbrChain1.StartWithdrawRemoveLiquidity(transactor, reqid, 0, wdLq1, wdLq2)
	tc.ChkErr(err, "u0 chain1 start withdraw")
	log.Infoln("withdraw reqid:", reqid)
	detail := tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Users[0].Address, reqid, 3)
	curss, err := tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw")

	res, err = cbrcli.QueryLiquidityDetailList(transactor.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     tc.ClientEthAddrs[0].Hex(),
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	log.Infoln("QueryLiquidityDetailList resp:", res.String())

	log.Infoln("======================== Xfer back ===========================")
	err = tc.CbrChain2.Approve(0, xferAmt)
	tc.ChkErr(err, "u0 chain2 approve")
	xferId, err = tc.CbrChain2.Send(0, xferAmt, tc.CbrChain1.ChainId, 1)
	tc.ChkErr(err, "u0 chain2 send")
	tc.CheckXfer(transactor, xferId[:])

	// todo: more cases, eg. lp2 withdraw from chain1 after xfer

	log.Infoln("======================== LP claim farming reward on-chain ===========================")
	err = tc.StartClaimFarmingRewards(transactor, 0)
	tc.ChkErr(err, "u0 start claim all farming rewards")
	info := tc.GetFarmingRewardClaimInfoWithSigs(transactor, 0, 3)
	err = tc.OnchainClaimFarmingRewards(&info.RewardClaimDetailsList[0])
	tc.ChkErr(err, "u0 onchain claim farming rewards")

	log.Infoln("======================== Delegator 0 claim fee share ===========================")
	feeShareInfo, err := tc.GetCBridgeFeeShareInfo(transactor, 0)
	tc.ChkErr(err, "del0 get fee share info before claim")
	log.Infoln("feeShareInfo before claim:", feeShareInfo)
	assert.Equal(t, 2, len(feeShareInfo.ClaimableFeeAmounts), "Should have 2 fees")

	reqid = uint64(time.Now().Unix())
	feeShareWdLq := &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain1.ChainId,
		TokenAddr:   tc.CbrChain1.USDTAddr.Hex(),
		Ratio:       100000000, // Only support 100% for now
		// MaxSlippage unsupported for now
	}
	err = tc.CbrChain1.StartWithdrawClaimFeeShare(transactor, reqid, 0, feeShareWdLq)
	tc.ChkErr(err, "del0 chain1 start claim fee share")
	log.Infoln("claim fee share withdraw reqid:", reqid)
	detail = tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Delegators[0].Address, reqid, 3)
	curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw fee share")

	feeShareInfo, err = tc.GetCBridgeFeeShareInfo(transactor, 0)
	tc.ChkErr(err, "del0 get fee share info after claim")
	log.Infoln("feeShareInfo after claim:", feeShareInfo)
	assert.Equal(t, 1, len(feeShareInfo.ClaimableFeeAmounts), "Should have 1 fee")
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
	expSigners := genSortedSigners([]eth.Addr{tc.ValSignerAddrs[0]}, initSignerPowers)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	log.Infoln("================== Add validators ======================")
	tc.AddValidator(t, transactor, 0, big.NewInt(3e18), eth.CommissionRate(0.03))

	tc.AddValidator(t, transactor, 1, big.NewInt(2e18), eth.CommissionRate(0.02))
	expSigners = genSortedSigners(
		[]eth.Addr{tc.ValSignerAddrs[0], tc.ValSignerAddrs[1]},
		[]*big.Int{big.NewInt(3e18), big.NewInt(2e18)})
	tc.CheckLatestSigners(t, transactor, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	tc.AddValidator(t, transactor, 2, big.NewInt(4e18), eth.CommissionRate(0.01))
	expSigners = genSortedSigners(
		[]eth.Addr{tc.ValSignerAddrs[0], tc.ValSignerAddrs[1], tc.ValSignerAddrs[2]},
		[]*big.Int{big.NewInt(3e18), big.NewInt(2e18), big.NewInt(4e18)})
	tc.CheckLatestSigners(t, transactor, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	log.Infoln("============= Update validator signer =================")
	ShutdownNode(0)
	configPath := "../../../docker-volumes/node1/sgnd/config/sgn.toml"
	configFileViper := viper.New()
	configFileViper.SetConfigFile(configPath)
	err := configFileViper.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")
	configFileViper.Set(common.FlagEthSignerKeystore, "./keys/vethks1.json")
	err = configFileViper.WriteConfig()
	tc.ChkErr(err, "Failed to write config")
	BringupNode(0)

	tc.UpdateValidatorSigner(tc.ValAuths[1], tc.ValEthAddrs[1])
	expSigners = genSortedSigners(
		[]eth.Addr{tc.ValEthAddrs[1], tc.ValSignerAddrs[0], tc.ValSignerAddrs[2]},
		[]*big.Int{big.NewInt(2e18), big.NewInt(3e18), big.NewInt(4e18)})
	tc.CheckLatestSigners(t, transactor, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)
}

func genSortedSigners(addrs []eth.Addr, amts []*big.Int) []*cbrtypes.Signer {
	var ss []*cbrtypes.Signer
	for i, amt := range amts {
		ss = append(ss,
			&cbrtypes.Signer{
				Addr:  addrs[i].Bytes(),
				Power: amt.Bytes(),
			})
	}
	return ss
}
