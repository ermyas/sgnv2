package common

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
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
	msgcli "github.com/celer-network/sgn-v2/x/message/client/cli"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	pegbrcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func QueryTotalLiquidity(transactor *transactor.Transactor, chainId uint64, token eth.Addr) (*big.Int, error) {
	res, err := cbrcli.QueryTotalLiquidity(transactor.CliCtx, &cbrtypes.QueryTotalLiquidityRequest{ChainId: chainId, TokenAddr: token.Hex()})
	if err != nil {
		return nil, err
	}
	liq, ok := new(big.Int).SetString(res.TotalLiq, 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert liq to big int")
	}
	return liq, nil
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

func (c *CbrChain) StartDelegatorWithdrawClaimCbrFeeShare(
	transactor *transactor.Transactor, reqid, uid uint64, wdLqs []*cbrtypes.WithdrawLq) error {
	return c.startWithdrawClaimCbrFeeShare(transactor, reqid, c.Delegators[uid], wdLqs)
}

func (c *CbrChain) StartValidatorWithdrawClaimCbrFeeShare(
	transactor *transactor.Transactor, reqid, uid uint64, wdLqs []*cbrtypes.WithdrawLq) error {
	return c.startWithdrawClaimCbrFeeShare(transactor, reqid, c.Validators[uid], wdLqs)
}

func (c *CbrChain) startWithdrawClaimCbrFeeShare(
	transactor *transactor.Transactor, reqid uint64, client *TestEthClient, wdLqs []*cbrtypes.WithdrawLq) error {
	withdrawReq := &cbrtypes.WithdrawReq{
		Withdraws:    wdLqs,
		ExitChainId:  c.ChainId,
		ReqId:        reqid,
		WithdrawType: cbrtypes.ClaimFeeShare,
	}
	wdBytes, _ := withdrawReq.Marshal()

	_, err := cbrcli.InitWithdraw(transactor, &cbrtypes.MsgInitWithdraw{
		WithdrawReq: wdBytes,
		UserSig:     client.SignMsg(wdBytes),
		Creator:     transactor.Key.GetAddress().String(),
	})
	return err
}

func StartValidatorMultiWithdrawClaimCbrFeeShares(vid, reqid uint64, wdLqs []*cbrtypes.WithdrawLq) error {
	txr := NewTestTransactor(
		SgnHomes[vid],
		SgnChainID,
		SgnNodeURI,
		ValSgnAddrStrs[vid],
		SgnPassphrase,
	)

	var msgs []sdk.Msg
	for _, wd := range wdLqs {
		withdrawReq := &cbrtypes.WithdrawReq{
			Withdraws:    []*cbrtypes.WithdrawLq{wd},
			ExitChainId:  wd.FromChainId,
			ReqId:        reqid,
			WithdrawType: cbrtypes.ValidatorClaimFeeShare,
		}
		reqid += 1
		wdBytes, _ := withdrawReq.Marshal()
		msg := &cbrtypes.MsgInitWithdraw{
			WithdrawReq: wdBytes,
			Creator:     txr.Key.GetAddress().String(),
		}
		msgs = append(msgs, msg)
	}
	_, err := txr.SendTxMsgsWaitMined(msgs)
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
	_, err := transactor.LockSendTx(
		&farmingtypes.MsgClaimAllRewards{
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

func (c *CbrChain) StartDelegatorClaimPegBridgeFee(
	transactor *transactor.Transactor, uid uint64, chainId uint64, tokenAddress eth.Addr, nonce uint64) error {
	return c.startClaimPegBridgeFee(transactor, c.Delegators[uid], chainId, tokenAddress, nonce)
}

func (c *CbrChain) StartValidatorClaimPegBridgeFee(
	transactor *transactor.Transactor, uid uint64, chainId uint64, tokenAddress eth.Addr, nonce uint64) error {
	return c.startClaimPegBridgeFee(transactor, c.Validators[uid], chainId, tokenAddress, nonce)
}

func (c *CbrChain) startClaimPegBridgeFee(
	transactor *transactor.Transactor, client *TestEthClient, chainId uint64, tokenAddress eth.Addr, nonce uint64) error {
	msg := &pegbrtypes.MsgClaimFee{
		DelegatorAddress: client.Address.Hex(),
		ChainId:          chainId,
		TokenAddress:     eth.Addr2Hex(tokenAddress),
		Nonce:            nonce,
		Sender:           transactor.Key.GetAddress().String(),
	}
	signature := client.SignMsg(msg.EncodeDataToSignByDelegator())
	msg.Signature = signature

	_, err := pegbrcli.InitClaimFee(transactor, msg)
	return err
}

func StartValidatorSelfClaimPegbrFee(vid uint64, chainId uint64, tokenAddress eth.Addr, nonce uint64) error {
	txr := NewTestTransactor(
		SgnHomes[vid],
		SgnChainID,
		SgnNodeURI,
		ValSgnAddrStrs[vid],
		SgnPassphrase,
	)
	msg := &pegbrtypes.MsgClaimFee{
		DelegatorAddress: "",
		ChainId:          chainId,
		TokenAddress:     eth.Addr2Hex(tokenAddress),
		Nonce:            nonce,
		Signature:        []byte{},
		Sender:           txr.Key.GetAddress().String(),
		IsValidator:      true,
	}
	_, err := txr.LockSendTx(msg)
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

func WaitForMessageExecuted(transactor *transactor.Transactor, expectedStatus msgtypes.ExecutionStatus) {
	var err error
	var resp *msgtypes.QueryExecutionContextsResponse
	log.Infoln("finding active message id...")
	for retry := 0; retry < RetryLimit; retry++ {
		resp, err = msgcli.QueryExecutionContexts(transactor.CliCtx, &msgtypes.QueryExecutionContextsRequest{})
		if err == nil && len(resp.ExecutionContexts) == 1 {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryExecutionContexts")
	if len(resp.ExecutionContexts) == 0 {
		log.Fatalf("QueryExecutionContexts expected more than 1 result, actual %d result", len(resp.ExecutionContexts))
	}

	msgId := eth.Bytes2Hash(resp.ExecutionContexts[0].MessageId)
	var expected bool
	log.Infoln("checking message", msgId)
	for retry := 0; retry < RetryLimit; retry++ {
		message, err := msgcli.QueryMessage(transactor.CliCtx, msgId.Hex())
		if err == nil && message.ExecutionStatus == expectedStatus {
			expected = true
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryMessage")
	if !expected {
		log.Fatal("message check failed")
	}
	log.Infof("message executed with expected status:%s", expectedStatus.String())
}

func CheckTotalSupply(
	transactor *transactor.Transactor,
	peggedChainId uint64,
	peggedAddress eth.Addr,
	expected string,
) {
	_, total, err := pegbrcli.QuerySupplyInfo(transactor.CliCtx, peggedChainId, peggedAddress)
	ChkErr(err, "failed to query supply info")
	if total == expected {
		return
	}
	log.Fatalf("Total supply for %s on chain %d is %s, expected %s",
		eth.Addr2Hex(peggedAddress), peggedChainId, total, expected)
}

func GetSupplyCap(
	transactor *transactor.Transactor,
	peggedChainId uint64,
	peggedAddress eth.Addr,
) string {
	cap, _, err := pegbrcli.QuerySupplyInfo(transactor.CliCtx, peggedChainId, peggedAddress)
	ChkErr(err, "failed to query supply info")
	return cap
}

func WaitPbrDepositWithEmptyMintId(transactor *transactor.Transactor, depositId string) error {
	var err error
	log.Infoln("wait for deposit with empty mint id", depositId)
	for retry := 0; retry < RetryLimit; retry++ {
		depositInfo, err := pegbrcli.QueryDepositInfo(transactor.CliCtx, depositId)
		if err == nil {
			if len(depositInfo.MintId) == 0 {
				return nil
			} else {
				return fmt.Errorf("depositInfo found but with non-empty mintId: %s", eth.Bytes2Hex(depositInfo.MintId))
			}
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryDepositInfo")
	return nil
}

func WaitPbrBurnWithEmptyWithdrawId(transactor *transactor.Transactor, burnId string) error {
	var err error
	log.Infoln("wait for burn with empty withdraw id", burnId)
	for retry := 0; retry < RetryLimit; retry++ {
		burnInfo, err := pegbrcli.QueryBurnInfo(transactor.CliCtx, burnId)
		if err == nil {
			if len(burnInfo.WithdrawId) == 0 {
				return nil
			} else {
				return fmt.Errorf("burnInfo found but with non-empty withdrawId: %s", eth.Bytes2Hex(burnInfo.WithdrawId))
			}
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryBurnInfo")
	return nil
}

func StartClaimPegbridgeRefund(transactor *transactor.Transactor, refId string) error {
	msg := &pegbrtypes.MsgClaimRefund{
		RefId:  refId,
		Sender: transactor.Key.GetAddress().String(),
	}
	_, err := pegbrcli.InitClaimRefund(transactor, msg)
	return err
}

func FakeStartClaimPegbridgeRefund(transactor *transactor.Transactor, depositId string) error {
	err := StartClaimPegbridgeRefund(transactor, depositId)
	if err != nil {
		return nil
	}
	return fmt.Errorf("claim refund initiated with success")
}

func GetRefundWithdrawInfoWithSigs(
	transactor *transactor.Transactor, depositId string, expSigNum int) (
	withdrawId string, withdrawInfo *pegbrtypes.WithdrawInfo) {
	var err error
	// query refundClaimInfo, in order to get withdrawId
	for retry := 0; retry < RetryLimit; retry++ {
		withdrawId, err = pegbrcli.QueryRefundClaimInfo(
			transactor.CliCtx,
			depositId,
		)
		if err == nil {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryRefundClaimInfo")
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

func GetRefundMintInfoWithSigs(
	transactor *transactor.Transactor, burnId string, expSigNum int) (
	mintId string, mintInfo *pegbrtypes.MintInfo) {
	var err error
	// query refundClaimInfo, in order to get mintId
	for retry := 0; retry < RetryLimit; retry++ {
		mintId, err = pegbrcli.QueryRefundClaimInfo(
			transactor.CliCtx,
			burnId,
		)
		if err == nil {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryRefundClaimInfo")
	// Then wait for MintInfo with enough signatures
	var mtInfo pegbrtypes.MintInfo
	for retry := 0; retry < RetryLimit; retry++ {
		mtInfo, err = pegbrcli.QueryMintInfo(
			transactor.CliCtx,
			mintId,
		)
		if err == nil && len(mtInfo.Signatures) == expSigNum {
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to QueryMintInfo")
	mintInfo = &mtInfo
	if len(mintInfo.Signatures) != expSigNum {
		log.Fatalf("QueryMintInfo expected sigNum %d, actual %d", expSigNum, len(mintInfo.Signatures))
	}
	return mintId, mintInfo
}
