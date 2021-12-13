package onchain

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

func GatewayOnSend(transferId, usrAddr, tokenAddr, amt, sendTxHash string, srcChainId, dsChainId uint64, bridgeType int) error {
	token, found, dbErr := dal.DB.GetTokenByAddr(tokenAddr, srcChainId)
	if token == nil || !found || dbErr != nil {
		log.Errorf("token from send event not found in db, addr:%s, chainId:%d", tokenAddr, srcChainId)
		return nil
	}
	estimatedAmt := "0"
	var getEstimatedAmtErr error
	switch bridgeType {
	case dal.BridgeTypeSendRelay:
		estimatedAmt, getEstimatedAmtErr = getEstimatedAmt(srcChainId, dsChainId, token, amt)
		if getEstimatedAmtErr != nil {
			log.Warnf("estimateAmt on send for transferId:%s failed, err:%s", transferId, getEstimatedAmtErr.Error())
			estimatedAmt = "0"
		}
		break
	case dal.BridgeTypeDepositMint, dal.BridgeTypeBurnWithDraw:
		estimatedAmt, getEstimatedAmtErr = getEstimatedPeggedAmt(srcChainId, dsChainId, token.GetToken().GetSymbol(), amt)
		if getEstimatedAmtErr != nil {
			log.Warnf("estimateAmt on send for transferId:%s failed, err:%s", transferId, getEstimatedAmtErr.Error())
			estimatedAmt = "0"
		}
		break
	}
	volume, getVolumeErr := dal.DB.GetUsdVolume(token.GetToken().GetSymbol(), srcChainId, common.Str2BigInt(amt))
	if getVolumeErr != nil {
		log.Warnf("find invalid token volume, symbol:%s, chainId:%d, we set volume to 0 first", token.GetToken().GetSymbol(), srcChainId)
		// continue to save 0 volume in db
	}
	return dal.DB.UpsertTransferOnSend(transferId, usrAddr, token, amt, estimatedAmt, sendTxHash, srcChainId, dsChainId, volume, getFeePerc(srcChainId, dsChainId, token.GetToken().GetSymbol()), bridgeType)
}

func GatewayOnRelay(c *ethclient.Client, transferId, txHash, dstTransferId, amt, usrAddr, tokenAddr string, srcChainId, dstChainId uint64, bridgeType int) error {
	_, isDelayed, err := dal.DB.GetDelayedOp(dstTransferId)
	if err != nil {
		return err
	}
	if isDelayed {
		dal.DB.UpdateDelayedOpType(dstTransferId, dal.DelayedOpTransfer)
	}
	token, found, dbErr := dal.DB.GetTokenByAddr(tokenAddr, dstChainId)
	if token == nil || !found || dbErr != nil {
		log.Errorf("token from relay event not found in db, addr:%s, chainId:%d", tokenAddr, srcChainId)
		return nil
	}
	err = dal.DB.UpsertTransferOnRelay(transferId, dstTransferId, usrAddr, token, amt, txHash, srcChainId, dstChainId, isDelayed, bridgeType)
	if err == nil {
		dal.DB.AddFeeRebateFee(transferId)
		sendGasOnArrival(c, transferId)
	}
	return err
}

func GatewayOnLiqWithdraw(id, tx string, chid, seq uint64, addr string) {
	_, isDelayed, err := dal.DB.GetDelayedOp(id)
	if err != nil {
		log.Warnf("Unable to fetch record from delayed_ops, id %s err %s", id, err.Error())
	}
	/*
		the "refund" kind of withdrawal
	*/
	transferId, found, err := dal.DB.GetTransferByRefundSeqNum(chid, seq, addr)
	if err != nil {
		log.Warnf("error when get transfer, usr:%s chainId:%d, seqNum:%d, err:%+v", addr, chid, seq, err)
	}
	if found {
		toStatus := uint64(types.TransferHistoryStatus_TRANSFER_REFUNDED)
		if isDelayed {
			toStatus = uint64(types.TransferHistoryStatus_TRANSFER_DELAYED)
			// update delayed operation type so that when receiving the DelayedTransferExecuted we know that it's a refund not a withdrawal
			dal.DB.UpdateDelayedOpType(id, dal.DelayedOpRefund)
		}
		// save refund_id so if we later receive DelayedTransferExecuted, the handler can find this record
		err := dal.DB.UpdateTransferForRefund(transferId, toStatus, id, tx)
		if err != nil {
			log.Warnf("db when UpdateTransferStatus to TRANSFER_REFUNDED, transferId:%s, err:%+v", transferId, err)
		}
		return
	}
	/*
		liquidity withdrawal
	*/
	toStatus := uint64(types.WithdrawStatus_WD_COMPLETED)
	if isDelayed {
		toStatus = uint64(types.WithdrawStatus_WD_DELAYED)
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
	err = dal.DB.UpdateLP(chid, seq, toStatus, addr, wdid.Hex(), tx)
	if err != nil {
		log.Errorln(logmsg, err)
	}
}

func GatewayOnLiqAdd(lpAddr, tokenAddr, amt, txHash string, chainId uint64, seqNum, nonce uint64) error {
	token, found, dbErr := dal.DB.GetTokenByAddr(tokenAddr, chainId)
	if token == nil || !found || dbErr != nil {
		log.Errorf("token from LiqAdd event not found in db, addr:%s, chainId:%d", tokenAddr, chainId)
		return nil
	}
	status := types.WithdrawStatus_WD_WAITING_FOR_SGN
	lpType := webapi.LPType_LP_TYPE_ADD
	volume, getVolumeErr := dal.DB.GetUsdVolume(token.GetToken().GetSymbol(), chainId, common.Str2BigInt(amt))
	if getVolumeErr != nil {
		log.Warnf("find invalid token volume, symbol:%s, chainId:%d, we set volume to 0 first", token, chainId)
		// continue to save 0 volume in db
	}
	return dal.DB.UpsertLPWithTx(lpAddr, token.GetToken().GetSymbol(), tokenAddr, amt, txHash, chainId,
		uint64(status), uint64(lpType), seqNum, nonce, volume)
}

func GatewayOnDelayXferAdd(id, txHash string) error {
	t, err := bestEffortChecks(id, txHash)
	if err != nil {
		return err
	}
	// if DelayedTransferAdded event precedes Relay and WithdrawDone, which is expected, insert a record in delayed_op.
	// the arg t is delayed op type, if DelayedTransferAdded precedes the normal events, t should be Unknown and should
	// be set when receiving normal events. but if normal events precedes DelayedTransferAdded, then t also must be set
	// here so that when the delayed operations are executed, the executor knows which table it should look for record
	// to update
	err = dal.DB.InsertDelayedOp(id, txHash, t)
	if err != nil {
		return err
	}
	return nil
}

func bestEffortChecks(id, txHash string) (dal.DelayedOpType, error) {
	// The best effort checks are meant to correct the finalized state of the records of COMPLETED
	// to DELAYED in case DelayedTransferAdded event does not arrive before the corresponding events
	// this is at the cost of some additional DB queries but considering there won't be a lot of
	// delay events, it should be fine
	// Best effort: check transfer table to make sure "Relay" did not arrive first
	_, found, err := dal.DB.GetTransferByDstTransferId(id)
	if err != nil {
		return dal.DelayedOpUnknown, err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than Relay, id %s txhash %s", id, txHash)
		dal.DB.UpdateTransferStatusByDstTransferId(id, types.TransferHistoryStatus_TRANSFER_DELAYED, "")
		return dal.DelayedOpTransfer, nil
	}
	// Best effort: check transfer table to make sure the "WithdrawDone" for refund did not arrive first
	found, err = dal.DB.ExistsTransferWithRefundId(id)
	if err != nil {
		return dal.DelayedOpUnknown, err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than WithdrawDone(refund), id %s txhash %s", id, txHash)
		dal.DB.UpdateTransferStatusByRefundId(id, types.TransferHistoryStatus_TRANSFER_DELAYED, "")
		return dal.DelayedOpRefund, nil
	}
	// Best effort: check lp table to make sure "WithdrawDone" did not arrive first
	found, err = dal.DB.ExistsLPInfoWithWithdrawId(id)
	if err != nil {
		return dal.DelayedOpUnknown, err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than WithdrawDone(withdraw), id %s txhash %s", id, txHash)
		dal.DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_DELAYED, "")
		return dal.DelayedOpWithdraw, nil
	}
	return dal.DelayedOpUnknown, nil
}

func GatewayOnDelayXferExec(id string, txHash string) error {
	t, found, err := dal.DB.GetDelayedOp(id)
	if err != nil {
		return err
	}
	if !found {
		logmsg := fmt.Sprintf("Got DelayedTransferExecuted but no delayed_op record found with id %s. the DelayedTransferAdded event was probably lost.", id)
		// if no record found, it probably means the DelayedTransferAdded event was lost
		// in this case we do best effort to look for records in lp and transfer table and finalize it's state to COMPLETED
		_, found, err := dal.DB.GetTransferByDstTransferId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating transfer status to COMPLETED")
			dal.DB.UpdateTransferStatusByDstTransferId(id, types.TransferHistoryStatus_TRANSFER_COMPLETED, txHash)
			return nil
		}
		// Best effort: check transfer table to make sure the "WithdrawDone" for refund did not arrive first
		found, err = dal.DB.ExistsTransferWithRefundId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating transfer status to REFUNDED")
			dal.DB.UpdateTransferStatusByRefundId(id, types.TransferHistoryStatus_TRANSFER_REFUNDED, txHash)
			return nil
		}
		// Best effort: check lp table to make sure "WithdrawDone" did not arrive first
		found, err = dal.DB.ExistsLPInfoWithWithdrawId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating withdraw status to COMPLETED")
			dal.DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_COMPLETED, txHash)
			return nil
		}
	}

	if t == uint64(dal.DelayedOpTransfer) || t == uint64(dal.DelayedOpRefund) {
		_, found, err := dal.DB.GetTransferByDstTransferId(id)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("cannot process DelayedTransferExec with id %s, type %d: record not found in transfer table", id, t)
		}
		var toStatus types.TransferHistoryStatus
		if t == uint64(dal.DelayedOpTransfer) {
			toStatus = types.TransferHistoryStatus_TRANSFER_COMPLETED
		} else {
			toStatus = types.TransferHistoryStatus_TRANSFER_REFUNDED
		}
		err = dal.DB.UpdateTransferStatusByDstTransferId(id, toStatus, txHash)
		if err == nil {
			log.Infof("handled DelayedTransferExecuted, id %s status %d", id, toStatus)
			return nil
		}
	} else if t == uint64(dal.DelayedOpWithdraw) {
		found, err = dal.DB.ExistsLPInfoWithWithdrawId(id)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("cannot process DelayedTransferExec with id %s, type %d: record not found in transfer table", id, t)
		}
		err := dal.DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_COMPLETED, txHash)
		if err == nil {
			log.Infof("handled DelayedTransferExecuted, id %s status %d", id, types.WithdrawStatus_WD_COMPLETED)
			return nil
		}
	} else {
		return fmt.Errorf("cannot process DelayedTransferExecuted with id %s: the fetched record has an unknown type %d", id, t)
	}
	return nil
}

func sendGasOnArrival(c *ethclient.Client, transferId string) {
	transfer, found, err := dal.DB.GetTransfer(transferId)
	if err != nil || !found {
		log.Errorln("can't find transfer info at gateway, ", transferId, err)
		return
	}
	chain, _ := dal.GetChainCache(transfer.DstChainId)
	if transfer.TokenSymbol == "WETH" || chain.GetDropGasAmt() == "0" {
		return
	}
	dropGasAmt, found := big.NewInt(0).SetString(chain.GetDropGasAmt(), 10)
	if !found {
		return
	}
	userAddr := eth.Hex2Addr(transfer.UsrAddr)
	var ksBytes []byte
	ksBytes, err = ioutil.ReadFile(viper.GetString(common.FlagGatewayIncentiveRewardsKeystore))
	if err != nil {
		log.Errorln("fail to get FlagGatewayIncentiveRewardsKeystore ", err)
		return
	}
	ksAddrStr, err := eth.GetAddressFromKeystore(ksBytes)
	if err != nil {
		log.Errorln("fail to get GetAddressFromKeystore ", err)
		return
	}
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(ksBytes)), viper.GetString(common.FlagGatewayIncentiveRewardsPassphrase), big.NewInt(int64(transfer.DstChainId)))
	if err != nil {
		log.Errorln("fail to get NewTransactorWithChainID ", err)
		return
	}
	auth.Value = dropGasAmt
	ctx := context.Background()
	acctAddr := eth.Hex2Addr(ksAddrStr)
	var gasLimit uint64 = 21000
	var rawTx *ethtypes.Transaction
	head, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Errorln("fail to get HeaderByNumber ", err)
		return
	}
	nonce, err := c.PendingNonceAt(ctx, acctAddr)
	if err != nil {
		log.Errorln("fail to get PendingNonceAt ", err)
		return
	}
	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		log.Errorln("fail to get SuggestGasPrice ", err)
		return
	}
	if head.BaseFee != nil {
		// eip 1559, new dynamic tx, per spec we should do
		// maxPriorityFeePerGas: eth_gasPrice - base_fee or just use the eth_maxPriorityFeePerGas rpc
		// maxFeePerGas: maxPriorityFeePerGas + 2 * base_fee = eth_gasPrice + base_fee
		// note if we calculate sendamt based on maxFeePerGas, it will leave one base_fee*gas residual
		// assume maxPriorityFee is way smaller than base fee, we could do following:
		// GasTipCap := eth_maxPriorityFeePerGas and GasFeeCap := eth_gasPrice + GasTipCap
		// but the risk is if eth becomes busy, our tx may pending for a long time. as here our gas is only 21K, we are ok w/ base_fee*gas residual
		gasFeeCap := new(big.Int).Add(gasPrice, head.BaseFee)
		rawTx = ethtypes.NewTx(&ethtypes.DynamicFeeTx{
			Nonce:     nonce,
			To:        &userAddr,
			Gas:       21000,
			GasTipCap: new(big.Int).Sub(gasPrice, head.BaseFee),
			GasFeeCap: gasFeeCap,
			Value:     auth.Value,
		})
	} else {
		rawTx = ethtypes.NewTx(&ethtypes.LegacyTx{
			Nonce:    nonce,
			To:       &userAddr,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Value:    auth.Value,
		})
	}
	tx, err := auth.Signer(acctAddr, rawTx)
	if err != nil {
		log.Errorln("fail to Signer ", err)
		return
	}

	err = c.SendTransaction(ctx, tx)
	if err != nil {
		log.Errorln("fail to send Gas On Arrival on chain ", transfer.DstChainId, " amt:", dropGasAmt.String(), err)
		return
	}
	_, err = ethutils.WaitMined(context.Background(), c, tx, ethutils.WithBlockDelay(1), ethutils.WithPollingInterval(time.Second*5))
	if err != nil {
		log.Errorf("send gas on arrival to %x on chain %d dropGasAmt %s, WaitMined err %v", userAddr, transfer.DstChainId, dropGasAmt.String(), err)
		return
	}
	log.Infoln("send gas on arrival to ", userAddr, " on chain ", transfer.DstChainId, " dropGasAmt:", dropGasAmt)
}
