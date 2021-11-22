package dal

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

// wrapper for other package usage, out of gateway
// for query api, if DB is nil, will a common err
// for update api, will return error==nil if DB is nil

// UpdateTransferStatus update api
func UpdateTransferStatus(transferId string, status uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateTransferStatus(transferId, status)
	}
}

func UpdateTransferForRefund(transferId string, status uint64, refundId string) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateTransferForRefund(transferId, status, refundId)
	}
}

// UpsertTransferOnSend update api
func UpsertTransferOnSend(transferId, usrAddr, tokenSymbol, amt, sendTxHash string, srcChainId, dsChainId uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpsertTransferOnSend(transferId, usrAddr, tokenSymbol, amt, sendTxHash, srcChainId, dsChainId)
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

	// The best effort checks are meant to correct the finalized state of the records of COMPLETED
	// to DELAYED in case DelayedTransferAdded event does not arrive before the corresponding events
	// this is at the cost of some additional DB queries but considering there won't be a lot of
	// delay events, it should be fine

	t := DelayedOpUnknown
	// Best effort: check transfer table to make sure "Relay" did not arrive first
	_, found, err := DB.GetTransferByDstTransferId(id)
	if err != nil {
		return err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than Relay, id %s txhash %s", id, txHash)
		DB.UpdateTransferStatusByDstTransferId(id, types.TransferHistoryStatus_TRANSFER_DELAYED)
		t = DelayedOpTransfer
		return nil
	}
	// Best effort: check transfer table to make sure the "WithdrawDone" for refund did not arrive first
	found, err = DB.ExistsTransferWithRefundId(id)
	if err != nil {
		return err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than WithdrawDone(refund), id %s txhash %s", id, txHash)
		DB.UpdateTransferStatusByRefundId(id, types.TransferHistoryStatus_TRANSFER_DELAYED)
		t = DelayedOpRefund
		return nil
	}
	// Best effort: check lp table to make sure "WithdrawDone" did not arrive first
	found, err = DB.ExistsLPInfoWithWithdrawId(id)
	if err != nil {
		return err
	}
	if found {
		log.Warnf("DelayedTransferAdded arrives later than WithdrawDone(withdraw), id %s txhash %s", id, txHash)
		DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_DELAYED)
		t = DelayedOpWithdraw
		return nil
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

// id is dst_transfer_id or withdraw_id
func DelayXferExec(id string) error {
	if DB == nil {
		return nil
	}

	defer func() {
		err := DB.DeleteDelayedOp(id)
		if err != nil {
			log.Errorf("Could not delete delayed_op record of id %s: %s", id, err)
		}
	}()

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
			DB.UpdateTransferStatusByDstTransferId(id, types.TransferHistoryStatus_TRANSFER_COMPLETED)
			return nil
		}
		// Best effort: check transfer table to make sure the "WithdrawDone" for refund did not arrive first
		found, err = DB.ExistsTransferWithRefundId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating transfer status to REFUNDED")
			DB.UpdateTransferStatusByRefundId(id, types.TransferHistoryStatus_TRANSFER_REFUNDED)
			return nil
		}
		// Best effort: check lp table to make sure "WithdrawDone" did not arrive first
		found, err = DB.ExistsLPInfoWithWithdrawId(id)
		if err != nil {
			return err
		}
		if found {
			log.Warnln(logmsg, "Updating withdraw status to COMPLETED")
			DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_COMPLETED)
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
		err = DB.UpdateTransferStatusByDstTransferId(id, toStatus)
		if err == nil {
			log.Infof("handled DelayedTransferExecuted, id %s status %d", id, toStatus)
		}
	} else if t == uint64(DelayedOpWithdraw) {
		found, err = DB.ExistsLPInfoWithWithdrawId(id)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("cannot process DelayedTransferExec with id %s, type %d: record not found in transfer table", id, t)
		}
		err := DB.UpdateLPStatusByWithdrawId(id, types.WithdrawStatus_WD_COMPLETED)
		if err == nil {
			log.Infof("handled DelayedTransferExecuted, id %s status %d", id, types.WithdrawStatus_WD_COMPLETED)
		}
	} else {
		return fmt.Errorf("cannot process DelayedTransferExecuted with id %s: the fetched record has an unknown type %d", id, t)
	}

	return fmt.Errorf("DelayXferExec: id %s not found in either transfers or lp table", id)
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
func UpsertLPForLiqAdd(usrAddr, tokenSymbol, tokenAddr, amt, txHash string, chainId, status, lpType, seqNum uint64) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpsertLPWithTx(usrAddr, tokenSymbol, tokenAddr, amt, txHash, chainId, status, lpType, seqNum)
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

// UpdateLPStatusForWithdraw update api
func UpdateLPStatusForWithdraw(chainId, seqNum, status uint64, addr string) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateLPStatusForWithdraw(chainId, seqNum, status, addr)
	}
}

// UpdateLPStatusForWithdraw update api
func UpdateLP(chainId, seqNum, status uint64, addr, wdid string) error {
	if DB == nil {
		return nil
	} else {
		return DB.UpdateLP(chainId, seqNum, status, addr, wdid)
	}
}

// common error
func noDBErrorForQuery() error {
	return fmt.Errorf("no gateway DB support")
}
