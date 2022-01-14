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
	// Make the stake amounts more realistic to test precision handling when distributing fee share
	vAmts := []*big.Int{
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
	}
	vAddrs := []eth.Addr{
		tc.ValEthAddrs[0], tc.ValEthAddrs[1], tc.ValEthAddrs[2],
	}
	err := tc.FundAddrsErc20(tc.CelrAddr, vAddrs, vAmts[0], tc.EthClient, tc.EtherBaseAuth)
	tc.ChkErr(err, "fund validator accounts")
	numVals := len(vAmts)
	tc.SetupValidators(t, transactor, vAmts)
	tc.CbrChain1.SetInitSigners(vAmts)
	tc.CbrChain2.SetInitSigners(vAmts)
	expSigners := genSortedSigners([]eth.Addr{tc.ValSignerAddrs[0], tc.ValSignerAddrs[1], tc.ValSignerAddrs[2]}, vAmts)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	log.Infoln("================== Delegate from delegator 0 to all validators ======================")
	dAmts := []*big.Int{
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
	}
	dAddrs := []eth.Addr{tc.DelEthAddrs[0]}
	err = tc.FundAddrsErc20(tc.CelrAddr, dAddrs, new(big.Int).Mul(big.NewInt(3), dAmts[0]), tc.EthClient, tc.EtherBaseAuth)
	tc.ChkErr(err, "und delegator account")
	for i := 0; i < numVals; i++ {
		tc.Delegate(tc.DelAuths[0], tc.ValEthAddrs[i], dAmts[i])
	}
	for i := 0; i < numVals; i++ {
		expDel := &stakingtypes.Delegation{
			DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[0]),
			ValidatorAddress: eth.Addr2Hex(tc.ValEthAddrs[i]),
			Shares:           sdk.NewIntFromBigInt(dAmts[i]),
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
	addAmt := big.NewInt(50000 * 1e6)
	var i uint64
	for i = 0; i < 2; i++ {
		err = tc.CbrChain1.ApproveUSDT(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 approve", i))
		err = tc.CbrChain1.AddLiq(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 addliq", i))
		tc.CheckAddLiquidityStatus(transactor, tc.CbrChain1.ChainId, i+1)
	}

	log.Infoln("======================== Add liquidity on chain 2 ===========================")
	for i = 0; i < 2; i++ {
		err = tc.CbrChain2.ApproveUSDT(i, addAmt)
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
	xferAmt := big.NewInt(10000 * 1e6)
	err = tc.CbrChain1.ApproveUSDT(0, xferAmt)
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
	err = tc.CbrChain1.OnchainCbrWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw")

	res, err = cbrcli.QueryLiquidityDetailList(transactor.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     tc.ClientEthAddrs[0].Hex(),
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	log.Infoln("QueryLiquidityDetailList resp:", res.String())

	log.Infoln("======================== Xfer back ===========================")
	err = tc.CbrChain2.ApproveUSDT(0, xferAmt)
	tc.ChkErr(err, "u0 chain2 approve")
	xferId, err = tc.CbrChain2.Send(0, xferAmt, tc.CbrChain1.ChainId, 1)
	tc.ChkErr(err, "u0 chain2 send")
	tc.CheckXfer(transactor, xferId[:])

	// todo: more cases, eg. lp2 withdraw from chain1 after xfer

	log.Infoln("======================== LP claim farming reward on-chain ===========================")
	err = tc.StartClaimFarmingRewards(transactor, 0)
	tc.ChkErr(err, "u0 start claim all farming rewards")
	info := tc.GetFarmingRewardClaimInfoWithSigs(transactor, 0, 3)
	assert.Equal(t, len(info.RewardClaimDetailsList), 1)
	rewardClaimDetail := info.RewardClaimDetailsList[0]
	log.Infoln("rewardClaimDetail.CumulativeRewardAmounts", rewardClaimDetail.CumulativeRewardAmounts)
	assert.Equal(t, tc.CbrChain1.ChainId, rewardClaimDetail.ChainId)
	assert.Equal(t, 2, len(rewardClaimDetail.CumulativeRewardAmounts))
	reward0 := rewardClaimDetail.CumulativeRewardAmounts[0]
	reward1 := rewardClaimDetail.CumulativeRewardAmounts[1]
	assert.Equal(t, fmt.Sprintf("CELR/%d", tc.CbrChain1.ChainId), reward0.Denom)
	assert.Equal(t, fmt.Sprintf("USDT/%d", tc.CbrChain1.ChainId), reward1.Denom)
	// TODO: Check reward amounts are reasonable

	err = tc.OnchainClaimFarmingRewards(&info.RewardClaimDetailsList[0])
	tc.ChkErr(err, "u0 onchain claim farming rewards")

	log.Infoln("======================== Delegator 0 claim fee share chain 1 ===========================")
	feeShareInfo, err := tc.GetCBridgeFeeShareInfo(transactor, 0)
	tc.ChkErr(err, "del0 get fee share info before claim")
	log.Infoln("feeShareInfo.ClaimableFeeAmounts before claim", feeShareInfo.ClaimableFeeAmounts)
	assert.Equal(t, 2, len(feeShareInfo.ClaimableFeeAmounts), "Should have 2 fees")
	fee0 := feeShareInfo.ClaimableFeeAmounts[0]
	fee1 := feeShareInfo.ClaimableFeeAmounts[1]
	assert.Equal(t, fmt.Sprintf("CBF-USDT/%d", tc.CbrChain1.ChainId), fee0.Denom)
	assert.Equal(t, fmt.Sprintf("CBF-USDT/%d", tc.CbrChain2.ChainId), fee1.Denom)
	assert.True(t, fee0.Amount.GT(sdk.NewDec(1e5)))
	assert.True(t, fee0.Amount.LT(sdk.NewDec(2e5)))
	assert.True(t, fee1.Amount.GT(sdk.NewDec(1e5)))
	assert.True(t, fee1.Amount.LT(sdk.NewDec(2e5)))

	reqid = uint64(time.Now().Unix())
	feeShareWdLq := &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain1.ChainId,
		TokenAddr:   tc.CbrChain1.USDTAddr.Hex(),
		Ratio:       100000000, // Only support 100% for now
		// MaxSlippage unsupported for now
	}
	err = tc.CbrChain1.StartDelegatorWithdrawClaimCbrFeeShare(transactor, reqid, 0, []*cbrtypes.WithdrawLq{feeShareWdLq})
	tc.ChkErr(err, "del0 chain1 start claim fee share")
	log.Infoln("claim fee share withdraw reqid:", reqid)
	detail = tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Delegators[0].Address, reqid, 3)
	curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainCbrWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw fee share")

	feeShareInfo, err = tc.GetCBridgeFeeShareInfo(transactor, 0)
	tc.ChkErr(err, "del0 get fee share info after claim")
	log.Infoln("feeShareInfo.ClaimableFeeAmounts after claim", feeShareInfo.ClaimableFeeAmounts)
	assert.Equal(t, 1, len(feeShareInfo.ClaimableFeeAmounts), "Should have 1 fee")
	fee0 = feeShareInfo.ClaimableFeeAmounts[0]
	assert.Equal(t, fmt.Sprintf("CBF-USDT/%d", tc.CbrChain2.ChainId), fee0.Denom)
	assert.True(t, fee0.Amount.GT(sdk.NewDec(1e5)))
	assert.True(t, fee0.Amount.LT(sdk.NewDec(2e5)))

	// transfer from chain 2 to 1 again to generate fee for testing single delegator reward claim
	log.Infoln("======================== Xfer back 2 ===========================")
	xferAmt = big.NewInt(10000 * 1e6)
	err = tc.CbrChain2.ApproveUSDT(0, xferAmt)
	tc.ChkErr(err, "u0 chain2 approve")
	xferId, err = tc.CbrChain2.Send(0, xferAmt, tc.CbrChain1.ChainId, 2)
	tc.ChkErr(err, "u0 chain2 send")
	tc.CheckXfer(transactor, xferId[:])

	log.Infoln("======================== Delegator 0 claim fee share all chains ===========================")
	feeShareInfo, err = tc.GetCBridgeFeeShareInfo(transactor, 0)
	tc.ChkErr(err, "del0 get fee share info before claim")
	log.Infoln("feeShareInfo.ClaimableFeeAmounts before claim", feeShareInfo.ClaimableFeeAmounts)
	assert.Equal(t, 2, len(feeShareInfo.ClaimableFeeAmounts), "Should have 2 fees")
	fee0 = feeShareInfo.ClaimableFeeAmounts[0]
	fee1 = feeShareInfo.ClaimableFeeAmounts[1]
	assert.Equal(t, fmt.Sprintf("CBF-USDT/%d", tc.CbrChain1.ChainId), fee0.Denom)
	assert.Equal(t, fmt.Sprintf("CBF-USDT/%d", tc.CbrChain2.ChainId), fee1.Denom)
	// count in the fact that due to one additional xfer from chain 2 to chain 1, the slippage causes fee to be smaller than 1e5
	assert.True(t, fee0.Amount.GT(sdk.NewDec(90000)))
	assert.True(t, fee0.Amount.LT(sdk.NewDec(1e5)))
	assert.True(t, fee1.Amount.GT(sdk.NewDec(1e5)))
	assert.True(t, fee1.Amount.LT(sdk.NewDec(2e5)))

	reqid = uint64(time.Now().Unix())
	feeShareWdLq1 := &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain1.ChainId,
		TokenAddr:   tc.CbrChain1.USDTAddr.Hex(),
		MaxSlippage: 1000000, // 100%
	}
	feeShareWdLq2 := &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain2.ChainId,
		TokenAddr:   tc.CbrChain2.USDTAddr.Hex(),
		MaxSlippage: 1000000, // 100%
	}
	err = tc.CbrChain1.StartDelegatorWithdrawClaimCbrFeeShare(transactor, reqid, 0, []*cbrtypes.WithdrawLq{feeShareWdLq1, feeShareWdLq2})
	tc.ChkErr(err, "del0 chain1 start claim fee share")
	log.Infoln("claim fee share withdraw reqid:", reqid)
	detail = tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Delegators[0].Address, reqid, 3)
	curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainCbrWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw fee share")

	feeShareInfo, err = tc.GetCBridgeFeeShareInfo(transactor, 0)
	tc.ChkErr(err, "del0 get fee share info after claim")
	log.Infoln("feeShareInfo.ClaimableFeeAmounts after claim", feeShareInfo.ClaimableFeeAmounts)
	// expect an extra fee generated through the transfer entailed by the single chain claim
	assert.Equal(t, 1, len(feeShareInfo.ClaimableFeeAmounts), "Should have 1 fee")
	assert.True(t, feeShareInfo.ClaimableFeeAmounts[0].Amount.LT(sdk.NewDec(20)))

	log.Infoln("======================== Validator 0 claim fee share all chains ===========================")
	reqid = uint64(time.Now().Unix())
	feeShareWdLq1 = &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain1.ChainId,
		TokenAddr:   tc.CbrChain1.USDTAddr.Hex(),
		MaxSlippage: 1000000, // 100%
	}
	feeShareWdLq2 = &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain2.ChainId,
		TokenAddr:   tc.CbrChain2.USDTAddr.Hex(),
		MaxSlippage: 1000000, // 100%
	}
	err = tc.CbrChain1.StartValidatorWithdrawClaimCbrFeeShare(transactor, reqid, 0, []*cbrtypes.WithdrawLq{feeShareWdLq1, feeShareWdLq2})
	tc.ChkErr(err, "val0 chain1 start claim fee share")
	log.Infoln("claim fee share withdraw reqid:", reqid)
	detail = tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Validators[0].Address, reqid, 3)
	curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainCbrWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw fee share")

	log.Infoln("======================== Validator 1 claim fee share all chains ===========================")
	startReqId := uint64(time.Now().Unix())
	feeShareWdLq1 = &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain1.ChainId,
		TokenAddr:   tc.CbrChain1.USDTAddr.Hex(),
	}
	feeShareWdLq2 = &cbrtypes.WithdrawLq{
		FromChainId: tc.CbrChain2.ChainId,
		TokenAddr:   tc.CbrChain2.USDTAddr.Hex(),
	}
	err = tc.StartValidatorMultiWithdrawClaimCbrFeeShares(1, startReqId, []*cbrtypes.WithdrawLq{feeShareWdLq1, feeShareWdLq2})
	tc.ChkErr(err, "val1 start claim fee share on all chains")

	log.Infoln("claim fee share withdraw reqid:", startReqId)
	detail = tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Validators[1].Address, startReqId, 3)
	curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainCbrWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw fee share")

	log.Infoln("claim fee share withdraw reqid:", startReqId+1)
	detail = tc.GetWithdrawDetailWithSigs(transactor, tc.CbrChain1.Validators[1].Address, startReqId+1, 3)
	curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain2.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain2.OnchainCbrWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw fee share")
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
