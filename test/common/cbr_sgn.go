package common

import (
	"bytes"
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingcli "github.com/celer-network/sgn-v2/x/farming/client/cli"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/gogo/protobuf/proto"
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
		if err == nil && resp.Status == cbrtypes.LPHistoryStatus_LP_COMPLETED {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryAddLiquidityStatus")
	if resp.Status != cbrtypes.LPHistoryStatus_LP_COMPLETED {
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
		if ok && curStatus != prevXferStatus {
			log.Infof("xfer status changed from %s to %s", prevXferStatus.String(), curStatus.String())
			prevXferStatus = curStatus
		}
		if err == nil && resp.Status[xferIdStr] == cbrtypes.TransferHistoryStatus_TRANSFER_COMPLETED {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryTransferStatus")
	if resp.Status[xferIdStr] != cbrtypes.TransferHistoryStatus_TRANSFER_COMPLETED {
		log.Fatalln("incorrect status")
	}
}

func CheckChainSigners(t *testing.T, transactor *transactor.Transactor, chainId uint64, expSigners *cbrtypes.SortedSigners) {
	var err error
	var signers *cbrtypes.ChainSigners
	for retry := 0; retry < RetryLimit; retry++ {
		signers, err = cbrcli.QueryChainSigners(transactor.CliCtx, chainId)
		if err == nil && signers != nil && sameSortedSigners(signers.GetCurrSigners(), expSigners) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryChainSigners")
	log.Infof("Query sgn and get chain %d signers: %s", chainId, signers.String())
	assert.True(t, sameSortedSigners(signers.GetCurrSigners(), expSigners), "expected signers should be: "+expSigners.String())
}

func CheckLatestSigners(t *testing.T, transactor *transactor.Transactor, expSigners *cbrtypes.SortedSigners) {
	var err error
	var signers *cbrtypes.LatestSigners
	for retry := 0; retry < RetryLimit; retry++ {
		signers, err = cbrcli.QueryLatestSigners(transactor.CliCtx)
		if err == nil && signers != nil && sameSortedSigners(signers.GetSigners(), expSigners) {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryLatestSigners")
	log.Infof("Query sgn and get latest signers: %s", signers.String())
	assert.True(t, sameSortedSigners(signers.GetSigners(), expSigners), "expected signers should be: "+expSigners.String())
}

func sameSortedSigners(ss1, ss2 *cbrtypes.SortedSigners) bool {
	b1, _ := proto.Marshal(ss1)
	b2, _ := proto.Marshal(ss2)
	return bytes.Compare(b1, b2) == 0
}

// call initwithdraw and return withdraw seqnum
func (c *CbrChain) StartWithdraw(transactor *transactor.Transactor, uid uint64, amt *big.Int) (uint64, error) {
	resp, err := cbrcli.InitWithdraw(transactor, &cbrtypes.MsgInitWithdraw{
		Chainid: c.ChainId,
		LpAddr:  c.Users[uid].Address.Bytes(),
		Token:   c.USDTAddr.Bytes(),
		Amount:  amt.Bytes(),
		Creator: transactor.Key.GetAddress().String(),
	})
	if err != nil {
		return 0, err
	}
	if resp.Errmsg != nil {
		return 0, errors.New(resp.Errmsg.String())
	}
	return resp.Seqnum, nil
}

func GetWithdrawDetailWithSigs(transactor *transactor.Transactor, wdseq uint64, expSigNum int) *cbrtypes.WithdrawDetail {
	var resp *cbrtypes.QueryLiquidityStatusResponse
	var err error
	for retry := 0; retry < RetryLimit; retry++ {
		resp, err = cbrcli.QueryWithdrawLiquidityStatus(
			transactor.CliCtx,
			&cbrtypes.QueryWithdrawLiquidityStatusRequest{
				SeqNum: wdseq,
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
	log.Infoln("wdseq", wdseq, "status:", resp.Status)
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
	log.Infoln("wdseq", wdseq, "status:", resp.Status)
	return resp.Detail, err
}

func GetCurSortedSigners(transactor *transactor.Transactor, chid uint64) ([]byte, error) {
	signers, err := cbrcli.QueryChainSigners(transactor.CliCtx, chid)
	if err != nil {
		return nil, err
	}
	return signers.SignersBytes, nil
}

// call claim-all
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
