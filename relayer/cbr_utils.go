package relayer

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *CbrOneChain) setCurss(ss []*cbrtypes.Signer) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs, c.curss.powers = cbrtypes.SignersToEthArrays(ss)
}

func (c *CbrOneChain) setCurssByEvent(e *eth.BridgeSignersUpdated) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs = make([]eth.Addr, len(e.Signers))
	c.curss.powers = make([]*big.Int, len(e.Powers))
	for i, addr := range e.Signers {
		c.curss.addrs[i] = addr
		c.curss.powers[i] = e.Powers[i]
	}
}

func (c *CbrOneChain) getCurss() currentSigners {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.curss
}

// each event's key is name-blkNum-index, value is json marshaled elog
func (c *CbrOneChain) saveEvent(name string, elog ethtypes.Log) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	key := fmt.Sprintf("%s-%d-%d", name, elog.BlockNumber, elog.Index)
	val, _ := json.Marshal(elog)
	return c.db.Set([]byte(key), val)
}

func (c *CbrOneChain) delEvent(name string, blknum, idx uint64) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.db.Delete([]byte(fmt.Sprintf("%s-%d-%d", name, blknum, idx)))
}

func (c *CbrOneChain) getTokenFromDB(tokenAddr string) (*webapi.TokenInfo, bool) {
	token, found, err := dal.GetTokenByAddr(tokenAddr, c.chainid)
	if err != nil || !found {
		return nil, false
	}
	return token, true
}

func getFeePerc(srcChainId, dstChainId uint64) uint32 {
	perc := uint32(0)
	tr := CurRelayerInstance.Transactor
	if tr != nil {
		_perc, err := cbrcli.QueryFeePerc(tr.CliCtx, &cbrtypes.GetFeePercentageRequest{
			SrcChainId: srcChainId,
			DstChainId: dstChainId,
		})
		if _perc == nil || err != nil {
			log.Warnf("get fee perc failed, srcChainId:%d, dsChainId:%d, will record 0 in db", srcChainId, dstChainId)
		} else {
			perc = _perc.FeePerc
		}
	}
	return perc
}

func getEstimatedAmt(srcChainId, dstChainId uint64, srcToken *webapi.TokenInfo, amt string) (string, error) {
	if !utils.IsValidAmt(amt) {
		return "0", fmt.Errorf("invalid amt, params checking failed")
	}
	tr := CurRelayerInstance.Transactor

	getFeeRequest := &cbrtypes.GetFeeRequest{
		SrcChainId:   srcChainId,
		DstChainId:   dstChainId,
		SrcTokenAddr: srcToken.Token.GetAddress(),
		Amt:          amt,
	}
	feeInfo, err := cbrcli.QueryFee(tr.CliCtx, getFeeRequest)
	if err != nil {
		log.Warnf("cli.QueryFee error, srcChainId:%d, dstChainId:%d, srcTokenAddr:%s, amt:%s, err:%+v", srcChainId, dstChainId, srcToken.Token.GetAddress(), amt, err)
		return "0", err
	}
	if feeInfo == nil {
		return "0", fmt.Errorf("can not estimate fee")
	}
	eqValueTokenAmt := feeInfo.GetEqValueTokenAmt()
	percFee := feeInfo.GetPercFee()
	baseFee := feeInfo.GetBaseFee()
	feeAmt := new(big.Int).Add(common.Str2BigInt(percFee), common.Str2BigInt(baseFee))
	estimateReceivedAmt := new(big.Int).Sub(common.Str2BigInt(eqValueTokenAmt), feeAmt)
	return estimateReceivedAmt.String(), nil
}

func GatewayOnSend(transferId, usrAddr, tokenAddr, amt, sendTxHash string, srcChainId, dsChainId uint64) error {
	if dal.DB == nil {
		return nil
	}
	srcToken, _, _ := dal.DB.GetTokenByAddr(tokenAddr, srcChainId)
	if srcToken == nil {
		return nil
	}
	estimatedAmt, err := getEstimatedAmt(srcChainId, dsChainId, srcToken, amt)
	if err != nil {
		return nil
	}
	return dal.UpsertTransferOnSend(transferId, usrAddr, srcToken, amt, estimatedAmt, sendTxHash, srcChainId, dsChainId, getFeePerc(srcChainId, dsChainId))
}

func GatewayOnRelay(transferId, txHash, dstTransferId, amt string) error {
	if dal.DB == nil {
		return nil
	}
	err := dal.TransferCompleted(transferId, txHash, dstTransferId, amt)
	if err != nil {
		dal.AddFeeRebateFee(transferId)
	}
	return err
}

func GatewayOnLiqAdd(lpAddr, tokenAddr, amt, txHash string, chainId uint64, seqNum, nonce uint64) error {
	if dal.DB == nil {
		return nil
	}
	token, _, _ := dal.DB.GetTokenByAddr(tokenAddr, chainId)
	if token == nil {
		return nil
	}
	status := cbrtypes.WithdrawStatus_WD_WAITING_FOR_SGN
	lpType := webapi.LPType_LP_TYPE_ADD
	return dal.UpsertLPForLiqAdd(lpAddr, token.GetToken().GetSymbol(), token.GetToken().GetAddress(), amt, txHash, chainId, uint64(status), uint64(lpType), seqNum, nonce)
}

func GatewayOnLiqWithdraw(id, tx string, chid, seq uint64, addr string) {
	if dal.DB == nil {
		return
	}
	_, isDelayed, err := dal.DB.GetDelayedOp(id)
	if err != nil {
		log.Warnf("Unable to fetch record from delayed_ops, id %s err %s", id, err.Error())
	}
	txhash := tx
	if isDelayed {
		txhash = ""
	}
	/*
		the "refund" kind of withdrawal
	*/
	transferId, found, err := dal.GetTransferByRefundSeqNum(chid, seq, addr)
	if err != nil {
		log.Warnf("error when get transfer, usr:%s chainId:%d, seqNum:%d, err:%+v", addr, chid, seq, err)
	}
	if found {
		toStatus := uint64(cbrtypes.TransferHistoryStatus_TRANSFER_REFUNDED)
		if isDelayed {
			toStatus = uint64(cbrtypes.TransferHistoryStatus_TRANSFER_DELAYED)
			// update delayed operation type so that when receiving the DelayedTransferExecuted we know that it's a refund not a withdrawal
			dal.DB.UpdateDelayedOpType(id, dal.DelayedOpRefund)
		}
		// save refund_id so if we later receive DelayedTransferExecuted, the handler can find this record
		err := dal.UpdateTransferForRefund(transferId, toStatus, id, txhash)
		if err != nil {
			log.Warnf("db when UpdateTransferStatus to TRANSFER_REFUNDED, transferId:%s, err:%+v", transferId, err)
		}
		return
	}
	/*
		liquidity withdrawal
	*/
	toStatus := uint64(cbrtypes.WithdrawStatus_WD_COMPLETED)
	if isDelayed {
		toStatus = uint64(cbrtypes.WithdrawStatus_WD_DELAYED)
		// update delayed operation type so that when receiving the DelayedTransferExecuted we know that it's a withdrawal not a refund
		dal.DB.UpdateDelayedOpType(id, dal.DelayedOpWithdraw)
	}
	logmsg := fmt.Sprintf("cannot process WithdrawDone with id %s, chid %d, seq %d, addr %s:", id, chid, seq, addr)
	l, found, err := dal.DB.GetLPInfo(seq, uint64(webapi.LPType_LP_TYPE_REMOVE), chid, addr)
	if err != nil {
		log.Errorln(logmsg, err.Error())
		return
	}
	if !found {
		log.Errorln(logmsg, "record not found in lp")
		return
	}
	// calculate withdraw id
	wdid := utils.GenWithdrawId(chid, seq, l.UsrAddr, l.TokenAddr, l.Amt)
	err = dal.UpdateLP(chid, seq, toStatus, addr, wdid.Hex(), txhash)
	if err != nil {
		log.Errorln(logmsg, err)
	}
}

func GatewayOnDelayXferAdd(xferId, txHash string) error {
	if dal.DB == nil {
		return nil
	}
	return dal.DelayXferAdd(xferId, txHash)
}

func GatewayOnDelayXferExec(xferId, txHash string) error {
	if dal.DB == nil {
		return nil
	}
	return dal.DelayXferExec(xferId, txHash)
}
