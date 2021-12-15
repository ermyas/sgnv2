package common

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	distrcli "github.com/celer-network/sgn-v2/x/distribution/client/cli"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	farmingcli "github.com/celer-network/sgn-v2/x/farming/client/cli"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	pegbrcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/stretchr/testify/assert"
)

func CheckAddLiquidityStatus(transactor *transactor.Transactor, chainId, seqNum uint64) {
	var resp *cbrtypes.QueryLiquidityStatusResponse
	var err error
	for retry := 0; retry < RetryLimit*2; retry++ {
		resp, err = cbrcli.QueryAddLiquidityStatus(transactor.CliCtx, &cbrtypes.QueryAddLiquidityStatusRequest{
			ChainId: chainId,
			SeqNum:  seqNum,
		})
		if err == nil && resp.Status == cbrtypes.WithdrawStatus_WD_COMPLETED {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryAddLiquidityStatus")
	if resp.Status != cbrtypes.WithdrawStatus_WD_COMPLETED {
		log.Fatalln("incorrect status")
	}
}

func CheckXfer(transactor *transactor.Transactor, xferId []byte) {
	var resp *cbrtypes.QueryTransferStatusResponse
	var err error
	var prevXferStatus cbrtypes.TransferHistoryStatus
	xferIdStr := eth.Bytes2Hex(xferId)
	for retry := 0; retry < RetryLimit*2; retry++ {
		resp, err = cbrcli.QueryTransferStatus(transactor.CliCtx, &cbrtypes.QueryTransferStatusRequest{
			TransferId: []string{xferIdStr},
		})
		curStatus, ok := resp.Status[xferIdStr]
		if ok && curStatus.GatewayStatus != prevXferStatus {
			log.Infof("xfer status changed from %s to %s", prevXferStatus.String(), curStatus.String())
			prevXferStatus = curStatus.GatewayStatus
		}
		if err == nil && resp.Status[xferIdStr].GatewayStatus == cbrtypes.TransferHistoryStatus_TRANSFER_COMPLETED {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryTransferStatus")
	if resp.Status[xferIdStr].GatewayStatus != cbrtypes.TransferHistoryStatus_TRANSFER_COMPLETED {
		log.Fatalln("incorrect status")
	}
}

func WaitPbrDeposit(transactor *transactor.Transactor, depositId string) *pegbrtypes.DepositInfo {
	var err error
	log.Infoln("waiting for deposit", depositId)
	for retry := 0; retry < RetryLimit*2; retry++ {
		resp, err := pegbrcli.QueryDepositInfo(transactor.CliCtx, depositId)
		if err == nil {
			return &resp
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryDepositInfo")
	return nil
}

func CheckPbrWithdraw(transactor *transactor.Transactor, withdrawId string) *pegbrtypes.WithdrawInfo {
	var err error
	var expected bool
	var resp pegbrtypes.WithdrawInfo
	log.Infoln("checking withdraw Id", withdrawId)
	for retry := 0; retry < RetryLimit*2; retry++ {
		resp, err = pegbrcli.QueryWithdrawInfo(transactor.CliCtx, withdrawId)
		if err == nil && resp.Success {
			expected = true
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryWithdrawInfo")
	if !expected {
		log.Fatal("CheckPbrWithdraw check failed")
	}
	return &resp
}

func CheckPbrMint(transactor *transactor.Transactor, mintId string) *pegbrtypes.MintInfo {
	var err error
	var expected bool
	var resp pegbrtypes.MintInfo
	log.Infoln("checking mint Id", mintId)
	for retry := 0; retry < RetryLimit*2; retry++ {
		resp, err = pegbrcli.QueryMintInfo(transactor.CliCtx, mintId)
		if err == nil && resp.Success {
			expected = true
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryMintInfo")
	if !expected {
		log.Fatal("CheckPbrMint check failed")
	}
	return &resp
}

func WaitPbrBurn(transactor *transactor.Transactor, burnId string) *pegbrtypes.BurnInfo {
	var err error
	log.Infoln("waiting for burn", burnId)
	for retry := 0; retry < RetryLimit*2; retry++ {
		resp, err := pegbrcli.QueryBurnInfo(transactor.CliCtx, burnId)
		if err == nil {
			return &resp
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryBurnInfo")
	return nil
}

func CheckChainSigners(t *testing.T, transactor *transactor.Transactor, chainId uint64, expSigners []*cbrtypes.Signer) {
	var err error
	var signers *cbrtypes.ChainSigners
	for retry := 0; retry < RetryLimit; retry++ {
		signers, err = cbrcli.QueryChainSigners(transactor.CliCtx, chainId)
		if err == nil && signers != nil && sameSortedSigners(signers.GetSortedSigners(), expSigners) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryChainSigners")
	log.Infof("Query sgn and get chain %d signers: %s", chainId, signers.String())
	assert.True(t, sameSortedSigners(signers.GetSortedSigners(), expSigners),
		"expected signers should be: "+cbrtypes.PrintSigners(expSigners))
}

func CheckLatestSigners(t *testing.T, transactor *transactor.Transactor, expSigners []*cbrtypes.Signer) {
	var err error
	var signers *cbrtypes.LatestSigners
	for retry := 0; retry < RetryLimit; retry++ {
		signers, err = cbrcli.QueryLatestSigners(transactor.CliCtx)
		if err == nil && signers != nil && sameSortedSigners(signers.GetSortedSigners(), expSigners) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryLatestSigners")
	log.Infof("Query sgn and get latest signers: %s", signers.String())
	assert.True(t, sameSortedSigners(signers.GetSortedSigners(), expSigners),
		"expected signers should be: "+cbrtypes.PrintSigners(expSigners))
}

func sameSortedSigners(ss1, ss2 []*cbrtypes.Signer) bool {
	if len(ss1) != len(ss2) {
		return false
	}
	for i, s1 := range ss1 {
		if !bytes.Equal(s1.Addr, ss2[i].Addr) {
			return false
		}
		if !bytes.Equal(s1.Power, ss2[i].Power) {
			return false
		}
	}
	return true
}

func (c *CbrChain) GetWithdrawLq(ratio uint32) *cbrtypes.WithdrawLq {
	return &cbrtypes.WithdrawLq{
		FromChainId: c.ChainId,
		TokenAddr:   c.USDTAddr.Hex(),
		Ratio:       ratio,
		MaxSlippage: 50000,
	}
}

func (c *CbrChain) StartWithdrawRemoveLiquidity(transactor *transactor.Transactor, reqid, uid uint64, wdLqs ...*cbrtypes.WithdrawLq) error {
	withdrawReq := &cbrtypes.WithdrawReq{
		Withdraws:    wdLqs,
		ExitChainId:  c.ChainId,
		ReqId:        reqid,
		WithdrawType: cbrtypes.RemoveLiquidity,
	}
	wdBytes, _ := withdrawReq.Marshal()

	_, err := cbrcli.InitWithdraw(transactor, &cbrtypes.MsgInitWithdraw{
		WithdrawReq: wdBytes,
		UserSig:     c.Users[uid].SignMsg(wdBytes),
		Creator:     transactor.Key.GetAddress().String(),
	})
	return err
}

func (c *CbrChain) StartWithdrawClaimCbrFeeShare(transactor *transactor.Transactor, reqid, uid uint64, wdLqs []*cbrtypes.WithdrawLq) error {
	// NOTE: Only support single wdLq for now
	withdrawReq := &cbrtypes.WithdrawReq{
		Withdraws:    wdLqs,
		ExitChainId:  c.ChainId,
		ReqId:        reqid,
		WithdrawType: cbrtypes.ClaimFeeShare,
	}
	wdBytes, _ := withdrawReq.Marshal()

	_, err := cbrcli.InitWithdraw(transactor, &cbrtypes.MsgInitWithdraw{
		WithdrawReq: wdBytes,
		UserSig:     c.Delegators[uid].SignMsg(wdBytes),
		Creator:     transactor.Key.GetAddress().String(),
	})
	return err
}

func GetWithdrawDetailWithSigs(transactor *transactor.Transactor, usraddr eth.Addr, reqid uint64, expSigNum int) *cbrtypes.WithdrawDetail {
	var resp *cbrtypes.QueryLiquidityStatusResponse
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		resp, err = cbrcli.QueryWithdrawLiquidityStatus(
			transactor.CliCtx,
			&cbrtypes.QueryWithdrawLiquidityStatusRequest{
				SeqNum:  reqid,
				UsrAddr: eth.Addr2Hex(usraddr),
			})

		if err == nil && len(resp.GetDetail().GetSortedSigs()) == expSigNum {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to GetWithdrawDetail")
	if len(resp.GetDetail().GetSortedSigs()) != expSigNum {
		log.Fatalf("GetWithdrawDetail expected sigNum %d, actual %d", expSigNum, len(resp.GetDetail().GetSortedSigs()))
	}
	log.Infoln("Query sgn and get usr", usraddr.String(), "wdseq", reqid, "status:", resp.Status)
	return resp.Detail
}

func GetWithdrawDetail(transactor *transactor.Transactor, wdseq uint64) (*cbrtypes.WithdrawDetail, error) {
	resp, err := cbrcli.QueryWithdrawLiquidityStatus(
		transactor.CliCtx,
		&cbrtypes.QueryWithdrawLiquidityStatusRequest{
			SeqNum: wdseq,
		})
	if err != nil {
		return nil, err
	}
	log.Infoln("Query sgn and get wdseq", wdseq, "status:", resp.Status)
	return resp.Detail, err
}

func GetCurSortedSigners(transactor *transactor.Transactor, chid uint64) ([]*cbrtypes.Signer, error) {
	cs, err := cbrcli.QueryChainSigners(transactor.CliCtx, chid)
	if err != nil {
		return nil, err
	}
	return cs.SortedSigners, nil
}

// StartClaimFarmingRewards sends MsgClaimAll to the farming module
func StartClaimFarmingRewards(transactor *transactor.Transactor, uid uint64) error {
	_, err := farmingcli.ClaimAllRewards(transactor, &farmingtypes.MsgClaimAllRewards{
		Address: eth.Addr2Hex(ClientEthAddrs[uid]),
		Sender:  transactor.Key.GetAddress().String(),
	})
	return err
}

func GetFarmingRewardClaimInfoWithSigs(
	transactor *transactor.Transactor, uid uint64, expSigNum int) *farmingtypes.RewardClaimInfo {
	var info *farmingtypes.RewardClaimInfo
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		info, err = GetFarmingRewardClaimInfo(transactor, uid)
		if err == nil && len(info.RewardClaimDetailsList) == 1 &&
			len(info.RewardClaimDetailsList[0].Signatures) == expSigNum {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to GetFarmingRewardClaimInfo")
	list := info.GetRewardClaimDetailsList()
	if len(list) != 1 {
		log.Fatalf("GetFarmingRewardClaimInfo details list len %d", len(list))
	}
	if len(list[0].Signatures) != expSigNum {
		log.Fatalf("GetFarmingRewardClaimInfo sigs num %d, expected %d", len(list[0].Signatures), expSigNum)
	}
	return info
}

func GetFarmingRewardClaimInfo(transactor *transactor.Transactor, uid uint64) (*farmingtypes.RewardClaimInfo, error) {
	return farmingcli.QueryRewardClaimInfo(context.Background(), transactor.CliCtx, eth.Addr2Hex(ClientEthAddrs[uid]))
}

// StartClaimStakingReward sends MsgClaimAllStakingReward to the distribution module
func StartClaimStakingReward(transactor *transactor.Transactor, uid uint64) error {
	_, err := distrcli.ClaimAllStakingReward(transactor, &distrtypes.MsgClaimAllStakingReward{
		DelegatorAddress: eth.Addr2Hex(DelEthAddrs[uid]),
		Sender:           transactor.Key.GetAddress().String(),
	})
	return err
}

func GetStakingRewardClaimInfoWithSigs(
	transactor *transactor.Transactor, uid uint64, expSigNum int) *distrtypes.StakingRewardClaimInfo {
	var info *distrtypes.StakingRewardClaimInfo
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		info, err = GetStakingRewardClaimInfo(transactor, uid)
		if err == nil && len(info.Signatures) == expSigNum {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to GetStakingRewardClaimInfo")
	if len(info.Signatures) != expSigNum {
		log.Fatalf("GetStakingRewardClaimInfo sigs num %d, expected %d", len(info.Signatures), expSigNum)
	}
	return info
}

func GetStakingRewardClaimInfo(
	transactor *transactor.Transactor, uid uint64) (*distrtypes.StakingRewardClaimInfo, error) {
	return distrcli.QueryStakingRewardClaimInfo(
		context.Background(), transactor.CliCtx, eth.Addr2Hex(DelEthAddrs[uid]))
}

func GetCBridgeFeeShareInfo(transactor *transactor.Transactor, delId uint64) (*distrtypes.ClaimableFeesInfo, error) {
	return distrcli.QueryCBridgeFeeShareInfo(context.Background(), transactor.CliCtx, eth.Addr2Hex(DelEthAddrs[delId]))
}

func GetPegBridgeFeesInfo(transactor *transactor.Transactor, delId uint64) (*distrtypes.ClaimableFeesInfo, error) {
	return distrcli.QueryPegBridgeFeesInfo(context.Background(), transactor.CliCtx, eth.Addr2Hex(DelEthAddrs[delId]))
}

func (c *CbrChain) StartClaimPegBridgeFee(
	transactor *transactor.Transactor, uid uint64, chainId uint64, tokenAddress eth.Addr, nonce uint64) error {
	delegator := c.Delegators[uid]
	msg := &pegbrtypes.MsgClaimFee{
		DelegatorAddress: delegator.Address.Hex(),
		ChainId:          chainId,
		TokenAddress:     eth.Addr2Hex(tokenAddress),
		Nonce:            nonce,
		Sender:           transactor.Key.GetAddress().String(),
	}
	signature := delegator.SignMsg(msg.EncodeDataToSignByDelegator())
	msg.Signature = signature

	_, err := pegbrcli.InitClaimFee(transactor, msg)
	return err
}

func GetPegBridgeFeeClaimWithdrawInfoWithSigs(
	transactor *transactor.Transactor, delAddr eth.Addr, nonce uint64, expSigNum int) (
	withdrawId string, withdrawInfo *pegbrtypes.WithdrawInfo) {
	var feeClaimInfo pegbrtypes.FeeClaimInfo
	var err error
	// First wait for FeeClaimInfo
	for retry := 0; retry < RetryLimit; retry++ {
		feeClaimInfo, err = pegbrcli.QueryFeeClaimInfo(
			transactor.CliCtx,
			delAddr,
			nonce,
		)
		if err == nil {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryFeeClaimInfo")
	withdrawId = eth.Bytes2Hex(feeClaimInfo.WithdrawId)
	// Then wait for WithdrawInfo with enough signatures
	var wdInfo pegbrtypes.WithdrawInfo
	for retry := 0; retry < RetryLimit; retry++ {
		wdInfo, err = pegbrcli.QueryWithdrawInfo(
			transactor.CliCtx,
			withdrawId,
		)
		if err == nil && len(wdInfo.Signatures) == expSigNum {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryWithdrawInfo")
	withdrawInfo = &wdInfo
	if len(withdrawInfo.Signatures) != expSigNum {
		log.Fatalf("QueryWithdrawInfo expected sigNum %d, actual %d", expSigNum, len(withdrawInfo.Signatures))
	}
	return withdrawId, withdrawInfo
}
