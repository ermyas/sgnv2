package common

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	distrcli "github.com/celer-network/sgn-v2/x/distribution/client/cli"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	farmingcli "github.com/celer-network/sgn-v2/x/farming/client/cli"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
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
		if err == nil && resp.Status == cbrtypes.LiqStatus_LIQ_COMPLETED {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryAddLiquidityStatus")
	if resp.Status != cbrtypes.LiqStatus_LIQ_COMPLETED {
		log.Fatalln("incorrect status")
	}
}

func CheckXfer(transactor *transactor.Transactor, xferId []byte) {
	var resp *cbrtypes.QueryTransferStatusResponse
	var err error
	var prevXferStatus cbrtypes.TransferHistoryStatus
	xferIdStr := common.Bytes2Hex(xferId)
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

func CheckChainSigners(t *testing.T, transactor *transactor.Transactor, chainId uint64, expSigners []*cbrtypes.Signer) {
	var err error
	var signers *cbrtypes.ChainSigners
	for retry := 0; retry < RetryLimit; retry++ {
		signers, err = cbrcli.QueryChainSigners(transactor.CliCtx, chainId)
		if err == nil && signers != nil && sameSortedSigenrs(signers.GetSortedSigners(), expSigners) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryChainSigners")
	log.Infof("Query sgn and get chain %d signers: %s", chainId, signers.String())
	assert.True(t, sameSortedSigenrs(signers.GetSortedSigners(), expSigners),
		"expected signers should be: "+cbrtypes.PrintSigners(expSigners))
}

func CheckLatestSigners(t *testing.T, transactor *transactor.Transactor, expSigners []*cbrtypes.Signer) {
	var err error
	var signers *cbrtypes.LatestSigners
	for retry := 0; retry < RetryLimit; retry++ {
		signers, err = cbrcli.QueryLatestSigners(transactor.CliCtx)
		if err == nil && signers != nil && sameSortedSigenrs(signers.GetSortedSigners(), expSigners) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryLatestSigners")
	log.Infof("Query sgn and get latest signers: %s", signers.String())
	assert.True(t, sameSortedSigenrs(signers.GetSortedSigners(), expSigners),
		"expected signers should be: "+cbrtypes.PrintSigners(expSigners))
}

func sameSortedSigenrs(ss1, ss2 []*cbrtypes.Signer) bool {
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

func (c *CbrChain) StartWithdrawClaimFeeShare(transactor *transactor.Transactor, reqid, uid uint64, wdLq *cbrtypes.WithdrawLq) error {
	// NOTE: Only support single wdLq for now
	withdrawReq := &cbrtypes.WithdrawReq{
		Withdraws:    []*cbrtypes.WithdrawLq{wdLq},
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

func GetCBridgeFeeShareInfo(transactor *transactor.Transactor, delId uint64) (*distrtypes.CBridgeFeeShareInfo, error) {
	return distrcli.QueryCBridgeFeeShareInfo(context.Background(), transactor.CliCtx, eth.Addr2Hex(DelEthAddrs[delId]))
}
