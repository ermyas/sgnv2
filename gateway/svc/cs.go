package gatewaysvc

import (
	"context"
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/ops"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	OnChainTime = 15 * time.Minute

	NormalMsg     = "normal case, open history and operate according to tips, or report it to eng team if has more problems"
	ToolMsg       = "try to use miss event tools, if problem not fixed after using tools report it to eng team"
	WaitingMsg    = "too short time after user operation, keep waiting for a few minutes"
	CheckInputMsg = "can not find any result, check your input txHash and chain. If input is correct, waiting for 15 min. if you have waited longer than 15 min, report it to eng team"
)

func (gs *GatewayService) GetInfoByTxHash(ctx context.Context, request *webapi.GetInfoByTxHashRequest) (*webapi.GetInfoByTxHashResponse, error) {
	return gs.checkCaseStatus(request.GetType(), request.GetTxHash(), request.GetChainId()), nil
}

func (gs *GatewayService) FixEventMiss(ctx context.Context, request *webapi.FixEventMissRequest) (*webapi.FixEventMissResponse, error) {
	txHash := request.GetTxHash()
	chainId := request.GetChainId()
	status := request.GetType()
	switch status {
	case webapi.CSType_CT_TX:
		err := gs.fixTx(txHash, chainId)
		if err != nil {
			return &webapi.FixEventMissResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  err.Error(),
				},
			}, nil
		}
	case webapi.CSType_CT_LP_ADD:
		lpAddr, err := gs.getAddrFromHash(txHash, uint64(chainId))
		if err == nil {
			err = gs.fixLp(txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_REMOVE)
		}
		if err != nil {
			return &webapi.FixEventMissResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  err.Error(),
				},
			}, nil
		}
	case webapi.CSType_CT_LP_RM:
		lpAddr, err := gs.getAddrFromHash(txHash, uint64(chainId))
		if err == nil {
			err = gs.fixLp(txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_REMOVE)
		}

		if err != nil {
			return &webapi.FixEventMissResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  err.Error(),
				},
			}, nil
		}
	}
	return &webapi.FixEventMissResponse{}, nil
}

func (gs *GatewayService) checkCaseStatus(status webapi.CSType, txHash string, chainId uint32) *webapi.GetInfoByTxHashResponse {
	switch status {
	case webapi.CSType_CT_TX:
		return gs.diagnosisTx(txHash, chainId)
	case webapi.CSType_CT_LP_ADD:
		lpAddr, err := gs.getAddrFromHash(txHash, uint64(chainId))
		if err != nil {
			return &webapi.GetInfoByTxHashResponse{
				Memo: "can not find lp addr from txHash and chainId",
			}
		}
		return gs.diagnosisLp(txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_ADD)
	case webapi.CSType_CT_LP_RM:
		lpAddr, err := gs.getAddrFromHash(txHash, uint64(chainId))
		if err != nil {
			return &webapi.GetInfoByTxHashResponse{
				Memo: "can not find lp addr from txHash and chainId",
			}
		}
		return gs.diagnosisLp(txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_REMOVE)
	}
	return &webapi.GetInfoByTxHashResponse{}
}

func (gs *GatewayService) diagnosisTx(txHash string, chainId uint32) *webapi.GetInfoByTxHashResponse {
	resp := &webapi.GetInfoByTxHashResponse{
		Operation: webapi.CSOperation_CA_NORMAL,
		Memo:      NormalMsg,
	}
	tx0, txFound, dbErr := dal.DB.GetTransferBySrcTxHash(txHash, chainId)
	if txFound && dbErr == nil {
		_, _ = gs.GetTransferStatus(context.Background(), &webapi.GetTransferStatusRequest{TransferId: tx0.TransferId})
		tx, _, _ := dal.DB.GetTransferBySrcTxHash(txHash, chainId)
		caseStatus := mapTxStatus2CaseStatus(tx.Status)
		if tx.Status == types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED ||
			tx.Status == types.TransferHistoryStatus_TRANSFER_FAILED ||
			tx.Status == types.TransferHistoryStatus_TRANSFER_COMPLETED {
			resp = newInfoResponse(webapi.CSOperation_CA_NORMAL, NormalMsg, caseStatus)
		} else if tx0.UT.Add(OnChainTime).Before(time.Now()) {
			if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_FUND_RELEASE || caseStatus == webapi.UserCaseStatus_CC_TRANSFER_REQUESTING_REFUND {
				resp = newInfoResponse(webapi.CSOperation_CA_USE_RESIGN_TOOL, ToolMsg, caseStatus)
			} else if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_SUBMITTING ||
				caseStatus == webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_SGN_CONFIRMATION ||
				caseStatus == webapi.UserCaseStatus_CC_TRANSFER_CONFIRMING_YOUR_REFUND {
				resp = newInfoResponse(webapi.CSOperation_CA_USE_RESYNC_TOOL, ToolMsg, caseStatus)
			}
		} else {
			resp = newInfoResponse(webapi.CSOperation_CA_WAITING, WaitingMsg, caseStatus)
		}
		resp.Info = fmt.Sprintf("transferId:%s, status:%s, addr:%s, updateTime:%s, createTime:%s,srcAmt:%s, dstAmt:%s,, refundTx:%s, refundSeqNum:%d", tx.TransferId, tx.Status.String(), tx.UsrAddr, tx.UT.String(), tx.CT.String(), tx.SrcAmt, tx.DstAmt, tx.RefundTx, tx.RefundSeqNum)
	} else {
		resp = newInfoResponse(webapi.CSOperation_CA_MORE_INFO_NEEDED, CheckInputMsg, webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY)
	}
	return resp
}

func (gs *GatewayService) diagnosisLp(txHash, lpAddr string, chainId uint32, lpType webapi.LPType) *webapi.GetInfoByTxHashResponse {
	resp := &webapi.GetInfoByTxHashResponse{
		Operation: webapi.CSOperation_CA_NORMAL,
		Memo:      NormalMsg,
	}
	seqNum0, _, ut, lpFound, dbErr := dal.DB.GetLPInfoByHash(uint64(lpType), uint64(chainId), lpAddr, txHash)
	if lpFound && dbErr == nil {
		_, _ = gs.QueryLiquidityStatus(context.Background(), &webapi.QueryLiquidityStatusRequest{
			SeqNum:  seqNum0,
			TxHash:  txHash,
			LpAddr:  lpAddr,
			ChainId: chainId,
			Type:    lpType,
		})
		seqNum, status, _, _, _ := dal.DB.GetLPInfoByHash(uint64(lpType), uint64(chainId), lpAddr, txHash)
		caseStatus := mapLpStatus2CaseStatus(types.WithdrawStatus(status), lpType)
		if ut.Add(OnChainTime).Before(time.Now()) {
			if caseStatus == webapi.UserCaseStatus_CC_WAITING_FOR_LP {
				resp = newInfoResponse(webapi.CSOperation_CA_NORMAL, NormalMsg, caseStatus)
			} else if caseStatus == webapi.UserCaseStatus_CC_WITHDRAW_WAITING_FOR_SGN {
				resp = newInfoResponse(webapi.CSOperation_CA_USE_RESIGN_TOOL, ToolMsg, caseStatus)
			} else if caseStatus == webapi.UserCaseStatus_CC_ADD_SUBMITTING ||
				caseStatus == webapi.UserCaseStatus_CC_ADD_WAITING_FOR_SGN ||
				caseStatus == webapi.UserCaseStatus_CC_WITHDRAW_SUBMITTING {
				resp = newInfoResponse(webapi.CSOperation_CA_USE_RESYNC_TOOL, ToolMsg, caseStatus)
			}
		} else {
			resp = newInfoResponse(webapi.CSOperation_CA_WAITING, WaitingMsg, caseStatus)
		}
		resp.Info = fmt.Sprintf("seqNum:%d, status:%s,addr:%s, updateTime:%s", seqNum, types.WithdrawStatus(status).String(), lpAddr, ut.String())
	} else {
		resp = newInfoResponse(webapi.CSOperation_CA_MORE_INFO_NEEDED, CheckInputMsg, webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY)
	}
	return resp
}

func (gs *GatewayService) fixTx(txHash string, chainId uint32) error {
	tx, txFound, dbErr := dal.DB.GetTransferBySrcTxHash(txHash, chainId)
	if txFound && dbErr == nil {
		caseStatus := mapTxStatus2CaseStatus(tx.Status)
		if tx.UT.Add(OnChainTime).Before(time.Now()) {
			if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_FUND_RELEASE || caseStatus == webapi.UserCaseStatus_CC_TRANSFER_REQUESTING_REFUND {
				log.Infof("cs fix tx by resign, txHash:%s, chainId:%d, txId:%s", txHash, chainId, tx.TransferId)
				dal.DB.UpdateTransferStatus(tx.TransferId, uint64(tx.Status))
				_, err := gs.signAgainWithdraw(&types.MsgSignAgain{
					DataType: types.SignDataType_RELAY,
					Creator:  gs.TP.GetTransactor().Key.GetAddress().String(),
					XferId:   eth.Hex2Hash(tx.TransferId).Bytes(),
				})
				if err != nil {
					return err
				}
			} else if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_SUBMITTING ||
				caseStatus == webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_SGN_CONFIRMATION ||
				caseStatus == webapi.UserCaseStatus_CC_TRANSFER_CONFIRMING_YOUR_REFUND {
				log.Infof("cs fix tx by resync, txHash:%s, chainId:%d", txHash, chainId)
				// refresh update time
				dal.DB.UpdateTransferStatus(tx.TransferId, uint64(tx.Status))
				var err error
				if tx.DstTxHash == "" {
					err = ops.SyncCbrEvent(gs.TP.GetTransactor().CliCtx, uint64(chainId), txHash, types.CbrEventSend)
				} else {
					err = ops.SyncCbrEvent(gs.TP.GetTransactor().CliCtx, uint64(chainId), tx.DstTxHash, types.CbrEventRelay)
				}
				if err != nil {
					return err
				}

			}
		} else {
			return fmt.Errorf("frequence limited, please operate after until:%s", tx.UT.Add(OnChainTime).String())
		}
	}
	return nil
}

func (gs *GatewayService) fixLp(txHash, lpAddr string, chainId uint32, lpType webapi.LPType) error {
	seqNum, status, ut, lpFound, dbErr := dal.DB.GetLPInfoByHash(uint64(lpType), uint64(chainId), lpAddr, txHash)
	if lpFound && dbErr == nil {
		caseStatus := mapLpStatus2CaseStatus(types.WithdrawStatus(status), lpType)
		if ut.Add(OnChainTime).Before(time.Now()) {
			if caseStatus == webapi.UserCaseStatus_CC_WITHDRAW_WAITING_FOR_SGN {
				log.Infof("cs fix lp by resign, ReqId:%d, UserAddr:%s, chainId:%d, lpType:%s", seqNum, lpAddr, chainId, lpType.String())
				// refresh update time
				dal.DB.UpdateLPStatus(seqNum, uint64(lpType), uint64(chainId), lpAddr, status)
				_, err := gs.signAgainWithdraw(&types.MsgSignAgain{
					DataType: types.SignDataType_WITHDRAW,
					Creator:  gs.TP.GetTransactor().Key.GetAddress().String(),
					ReqId:    seqNum,
					UserAddr: eth.Hex2Addr(lpAddr).Bytes(),
				})
				if err != nil {
					return err
				}
			} else if caseStatus == webapi.UserCaseStatus_CC_ADD_SUBMITTING ||
				caseStatus == webapi.UserCaseStatus_CC_ADD_WAITING_FOR_SGN ||
				caseStatus == webapi.UserCaseStatus_CC_WITHDRAW_SUBMITTING {
				log.Infof("cs fix lp by resync, txHash:%s, chainId:%d, lpAddr:%s, lpType:%s", txHash, chainId, lpAddr, lpType.String())
				// refresh update time
				dal.DB.UpdateLPStatus(seqNum, uint64(lpType), uint64(chainId), lpAddr, status)
				var err error
				if lpType == webapi.LPType_LP_TYPE_ADD {
					err = ops.SyncCbrEvent(gs.TP.GetTransactor().CliCtx, uint64(chainId), txHash, types.CbrEventLiqAdd)
				} else if lpType == webapi.LPType_LP_TYPE_REMOVE {
					err = ops.SyncCbrEvent(gs.TP.GetTransactor().CliCtx, uint64(chainId), txHash, types.CbrEventWithdraw)
				} else {
					err = fmt.Errorf("unknown lp type:%s", lpType.String())
				}
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("frequence limited, please operate after until:%s", ut.Add(OnChainTime).String())
		}
	}
	return nil
}

func (gs *GatewayService) getAddrFromHash(txHash string, chainId uint64) (string, error) {
	tx, _, err := gs.getTransactionByHash(txHash, chainId)
	if err != nil {
		return "", err
	}
	sender, err := ethtypes.Sender(ethtypes.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return "", err
	}
	return sender.String(), nil
}

func (gs *GatewayService) getTransactionByHash(txHash string, chainId uint64) (*ethtypes.Transaction, bool, error) {
	ec := gs.EC[chainId]
	if ec == nil {
		return nil, false, fmt.Errorf("eth client not found for chainId:%d", chainId)
	}
	return ec.TransactionByHash(context.Background(), eth.Hex2Hash(txHash))
}

func mapLpStatus2CaseStatus(status types.WithdrawStatus, lpType webapi.LPType) webapi.UserCaseStatus {
	switch status {
	case types.WithdrawStatus_WD_UNKNOWN:
		if lpType == webapi.LPType_LP_TYPE_ADD {
			return webapi.UserCaseStatus_CC_ADD_NO_HISTORY
		} else {
			return webapi.UserCaseStatus_CC_UNKNOWN
		}
	case types.WithdrawStatus_WD_WAITING_FOR_SGN:
		if lpType == webapi.LPType_LP_TYPE_ADD {
			return webapi.UserCaseStatus_CC_ADD_WAITING_FOR_SGN
		} else {
			return webapi.UserCaseStatus_CC_WITHDRAW_WAITING_FOR_SGN
		}
	case types.WithdrawStatus_WD_SUBMITTING:
		if lpType == webapi.LPType_LP_TYPE_ADD {
			return webapi.UserCaseStatus_CC_ADD_SUBMITTING
		} else {
			return webapi.UserCaseStatus_CC_WITHDRAW_SUBMITTING
		}
	case types.WithdrawStatus_WD_WAITING_FOR_LP:
		return webapi.UserCaseStatus_CC_WAITING_FOR_LP
	default:
		return webapi.UserCaseStatus_CC_UNKNOWN
	}
}
func mapTxStatus2CaseStatus(status types.TransferHistoryStatus) webapi.UserCaseStatus {
	switch status {
	case types.TransferHistoryStatus_TRANSFER_UNKNOWN:
		return webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY
	case types.TransferHistoryStatus_TRANSFER_SUBMITTING:
		return webapi.UserCaseStatus_CC_TRANSFER_SUBMITTING
	case types.TransferHistoryStatus_TRANSFER_WAITING_FOR_SGN_CONFIRMATION:
		return webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_SGN_CONFIRMATION
	case types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE:
		return webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_FUND_RELEASE
	case types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND:
		return webapi.UserCaseStatus_CC_TRANSFER_REQUESTING_REFUND
	case types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND:
		return webapi.UserCaseStatus_CC_TRANSFER_CONFIRMING_YOUR_REFUND
	default:
		return webapi.UserCaseStatus_CC_UNKNOWN
	}
}

func newInfoResponse(operation webapi.CSOperation, memo string, status webapi.UserCaseStatus) *webapi.GetInfoByTxHashResponse {
	return &webapi.GetInfoByTxHashResponse{
		Operation: operation,
		Memo:      memo,
		Status:    status,
	}
}
