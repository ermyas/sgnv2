package dal

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"math/big"
)

// THIS FILE IS TO BE DELETED WHEN GATEWAY OFFICIALLY SEPARATES FROM SGN

func UpdateTransferForRefund(transferId string, status uint64, refundId, refundTx string) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateTransferForRefund(transferId, status, refundId, refundTx)
	}
}

// UpsertTransferOnSend update api
func UpsertTransferOnSend(transferId, usrAddr string, token *webapi.TokenInfo, amt, estimatedAmt, sendTxHash string, srcChainId, dsChainId uint64, perc uint32) error {
	if DB == nil {
		return nil
	} else {
		volume, getVolumeErr := DB.GetUsdVolume(token.GetToken().GetSymbol(), srcChainId, common.Str2BigInt(amt))
		if getVolumeErr != nil {
			log.Warnf("find invalid token volume, symbol:%s, chainId:%d, we set volume to 0 first", token.GetToken().GetSymbol(), srcChainId)
			// continue to save 0 volume in db
		}
		return DB.UpsertTransferOnSend(transferId, usrAddr, token, amt, estimatedAmt, sendTxHash, srcChainId, dsChainId, volume, perc)
	}
}

// TransferCompleted update api
func TransferCompleted(transferId, txHash, dstTransferId, amt string) error {
	if DB == nil {
		return nil
	}
	_, isDelayed, err := DB.GetDelayedOp(dstTransferId)
	if err != nil {
		return err
	}
	if isDelayed {
		DB.UpdateDelayedOpType(dstTransferId, DelayedOpTransfer)
	}
	return DB.TransferCompleted(transferId, txHash, dstTransferId, amt, isDelayed)
}

func DelayXferAdd(id, txHash string) error {
	if DB == nil {
		return nil
	}

	t, err := bestEffortChecks(id, txHash)
	if err != nil {
		return err
	}
	// if DelayedTransferAdded event precedes Relay and WithdrawDone, which is expected, insert a record in delayed_op.
	// the arg t is delayed op type, if DelayedTransferAdded precedes the normal events, t should be Unknown and should
	// be set when receiving normal events. but if normal events precedes DelayedTransferAdded, then t also must be set
	// here so that when the delayed operations are executed, the executor knows which table it should look for record
	// to update
	err = DB.InsertDelayedOp(id, txHash, t)
	if err != nil {
		return err
	}
	return nil
}

func bestEffortChecks(id, txHash string) (DelayedOpType, error) {
	// The best effort checks are meant to correct the finalized state of the records of COMPLETED
	// to DELAYED in case DelayedTransferAdded event does not arrive before the corresponding events
	// this is at the cost of some additional DB queries but considering there won't be a lot of
	// delay events, it should be fine
	// Best effort: check transfer table to make sure "Relay" did not arrive first
	_, found, err := DB.GetTransferByDstTransferId(id)
	if err != nil {
		return DelayedOpUnknown, err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than Relay, id %s txhash %s", id, txHash)
		DB.UpdateTransferStatusByDstTransferId(id, types.TransferHistoryStatus_TRANSFER_DELAYED, txHash)
		return DelayedOpTransfer, nil
	}
	// Best effort: check transfer table to make sure the "WithdrawDone" for refund did not arrive first
	found, err = DB.ExistsTransferWithRefundId(id)
	if err != nil {
		return DelayedOpUnknown, err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than WithdrawDone(refund), id %s txhash %s", id, txHash)
		DB.UpdateTransferStatusByRefundId(id, types.TransferHistoryStatus_TRANSFER_DELAYED, txHash)
		return DelayedOpRefund, nil
	}
	// Best effort: check lp table to make sure "WithdrawDone" did not arrive first
	found, err = DB.ExistsLPInfoWithWithdrawId(id)
	if err != nil {
		return DelayedOpUnknown, err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than WithdrawDone(withdraw), id %s txhash %s", id, txHash)
		DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_DELAYED, txHash)
		return DelayedOpWithdraw, nil
	}
	return DelayedOpUnknown, nil
}

// id is dst_transfer_id or withdraw_id
func DelayXferExec(id, txHash string) error {
	if DB == nil {
		return nil
	}

	t, found, err := DB.GetDelayedOp(id)
	if err != nil {
		return err
	}
	if !found {
		logmsg := fmt.Sprintf("Got DelayedTransferExecuted but no delayed_op record found with id %s. the DelayedTransferAdded event was probably lost.", id)
		// if no record found, it probably means the DelayedTransferAdded event was lost
		// in this case we do best effort to look for records in lp and transfer table and finalize it's state to COMPLETED
		_, found, err := DB.GetTransferByDstTransferId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating transfer status to COMPLETED")
			DB.UpdateTransferStatusByDstTransferId(id, types.TransferHistoryStatus_TRANSFER_COMPLETED, txHash)
			return nil
		}
		// Best effort: check transfer table to make sure the "WithdrawDone" for refund did not arrive first
		found, err = DB.ExistsTransferWithRefundId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating transfer status to REFUNDED")
			DB.UpdateTransferStatusByRefundId(id, types.TransferHistoryStatus_TRANSFER_REFUNDED, txHash)
			return nil
		}
		// Best effort: check lp table to make sure "WithdrawDone" did not arrive first
		found, err = DB.ExistsLPInfoWithWithdrawId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating withdraw status to COMPLETED")
			DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_COMPLETED, txHash)
			return nil
		}
	}

	if t == uint64(DelayedOpTransfer) || t == uint64(DelayedOpRefund) {
		_, found, err := DB.GetTransferByDstTransferId(id)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("cannot process DelayedTransferExec with id %s, type %d: record not found in transfer table", id, t)
		}
		var toStatus types.TransferHistoryStatus
		if t == uint64(DelayedOpTransfer) {
			toStatus = types.TransferHistoryStatus_TRANSFER_COMPLETED
		} else {
			toStatus = types.TransferHistoryStatus_TRANSFER_REFUNDED
		}
		err = DB.UpdateTransferStatusByDstTransferId(id, toStatus, txHash)
		if err == nil {
			log.Infof("handled DelayedTransferExecuted, id %s status %d", id, toStatus)
			return nil
		}
	} else if t == uint64(DelayedOpWithdraw) {
		found, err = DB.ExistsLPInfoWithWithdrawId(id)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("cannot process DelayedTransferExec with id %s, type %d: record not found in transfer table", id, t)
		}
		err := DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_COMPLETED, txHash)
		if err == nil {
			log.Infof("handled DelayedTransferExecuted, id %s status %d", id, types.WithdrawStatus_WD_COMPLETED)
			return nil
		}
	} else {
		return fmt.Errorf("cannot process DelayedTransferExecuted with id %s: the fetched record has an unknown type %d", id, t)
	}
	return nil
}

// GetTokenByAddr query api
func GetTokenByAddr(addr string, chainId uint64) (*webapi.TokenInfo, bool, error) {
	if DB == nil {
		return nil, false, noDBErrorForQuery()
	} else {
		return DB.GetTokenByAddr(addr, chainId)
	}
}

// UpsertLPForLiqAdd update api
func UpsertLPForLiqAdd(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum, nonce uint64) error {
	if DB == nil {
		return nil
	} else {
		volume, getVolumeErr := DB.GetUsdVolume(tokenSymbol, chainId, common.Str2BigInt(amt))
		if getVolumeErr != nil {
			log.Warnf("find invalid token volume, symbol:%s, chainId:%d, we set volume to 0 first", tokenSymbol, chainId)
			// continue to save 0 volume in db
		}
		return DB.UpsertLPWithTx(usrAddr, tokenSymbol, tokenAddr, amt, txHash, chainId, status, lpType, seqNum, nonce, volume)
	}
}

// GetTransferByRefundSeqNum query api
func GetTransferByRefundSeqNum(chainId, seqNum uint64, usrAddr string) (string, bool, error) {
	if DB == nil {
		return "", false, noDBErrorForQuery()
	} else {
		return DB.GetTransferByRefundSeqNum(chainId, seqNum, usrAddr)
	}
}

func GetGasOnArrival(transferId string) (dropGasAmt *big.Int, userAddr *common.Addr, chainId uint64, needDrop bool, err error) {
	if DB == nil {
		return nil, nil, 0, false, noDBErrorForQuery()
	} else {
		transfer, b, err := DB.GetTransfer(transferId)
		if err != nil || !b {
			log.Errorln("can't find transfer info at gateway, ", transferId)
			return nil, nil, 0, false, err
		}
		chain, _ := GetChainCache(transfer.DstChainId)
		if transfer.TokenSymbol == "WETH" {
			return nil, nil, 0, false, nil
		}
		dropGasAmt, b := big.NewInt(0).SetString(chain.GetDropGasAmt(), 10)
		if !b {
			return nil, nil, 0, false, nil
		}
		usrAddr := common.Hex2Addr(transfer.UsrAddr)
		return dropGasAmt, &usrAddr, transfer.DstChainId, true, nil
	}
}

// UpdateLP update api
func UpdateLP(chainId, seqNum, status uint64, addr, wdid, tx string) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateLP(chainId, seqNum, status, addr, wdid, tx)
	}
}

// common error
func noDBErrorForQuery() error {
	return fmt.Errorf("no gateway DB support")
}
