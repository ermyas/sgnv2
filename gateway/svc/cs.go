package gatewaysvc

import (
	"context"
	"fmt"
	"time"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	OnChainTime = 15 * time.Minute

	NormalMsg     = "normal case, open history and operate according to tips, or report it to eng team if has more problems"
	ToolMsg       = "try to use miss event tools, if problem not fixed after using tools report it to eng team"
	WaitingMsg    = "too short time after user operation, keep waiting for a few minutes"
	CheckInputMsg = "can not find any result, check your input txHash and chain. If input is correct, waiting for 15 min. if you have waited longer than 15 min, report it to eng team"
	ReportMsg     = "unknown issues, report to eng"
)

func (gs *GatewayService) GetInfoByTxHash(ctx context.Context, request *webapi.GetInfoByTxHashRequest) (*webapi.GetInfoByTxHashResponse, error) {
	return gs.checkCaseStatus(request.GetType(), request.GetTxHash(), request.GetChainId()), nil
}

func (gs *GatewayService) FixEventMiss(ctx context.Context, request *webapi.FixEventMissRequest) (*webapi.FixEventMissResponse, error) {
	request.GetTxHash()
	switch request.GetType() {
	// transfer related cases
	case webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY:

	case webapi.UserCaseStatus_CC_TRANSFER_SUBMITTING:
		// update数据库waiting for sgn， 更新状态
		// 根据状态决定是否走下一条
	case webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_SGN_CONFIRMATION:
	case webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_FUND_RELEASE:
	// add related cases
	case webapi.UserCaseStatus_CC_ADD_NO_HISTORY:
	case webapi.UserCaseStatus_CC_ADD_SUBMITTING:
	case webapi.UserCaseStatus_CC_ADD_WAITING_FOR_SGN:
	//withdraw related cases
	case webapi.UserCaseStatus_CC_WAITING_FOR_LP:
	case webapi.UserCaseStatus_CC_WITHDRAW_SUBMITTING:
	case webapi.UserCaseStatus_CC_WITHDRAW_WAITING_FOR_SGN:
	case webapi.UserCaseStatus_CC_TRANSFER_REQUESTING_REFUND:
	case webapi.UserCaseStatus_CC_TRANSFER_CONFIRMING_YOUR_REFUND:
	}
	return &webapi.FixEventMissResponse{}, nil
}

func (gs *GatewayService) checkCaseStatus(status webapi.CSType, txHash string, chainId uint32) *webapi.GetInfoByTxHashResponse {
	switch status {
	case webapi.CSType_CT_TX:
		return diagnosisTx(txHash, chainId)
	case webapi.CSType_CT_LP_ADD:
		lpAddr, err := gs.getAddrFromHash(txHash, uint64(chainId))
		if err != nil {
			return &webapi.GetInfoByTxHashResponse{
				Memo: "can not find lp addr from txHash and chainId",
			}
		}
		return diagnosisLp(txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_ADD)
	case webapi.CSType_CT_LP_RM:
		lpAddr, err := gs.getAddrFromHash(txHash, uint64(chainId))
		if err != nil {
			return &webapi.GetInfoByTxHashResponse{
				Memo: "can not find lp addr from txHash and chainId",
			}
		}
		return diagnosisLp(txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_REMOVE)
	}
	return &webapi.GetInfoByTxHashResponse{}
}

func diagnosisTx(txHash string, chainId uint32) *webapi.GetInfoByTxHashResponse {
	resp := &webapi.GetInfoByTxHashResponse{
		Operation: webapi.CSOperation_CA_NORMAL,
		Memo:      NormalMsg,
	}
	tx, txFound, dbErr := dal.DB.GetTransferBySrcTxHash(txHash, chainId)
	if txFound && dbErr == nil {
		caseStatus := mapTxStatus2CaseStatus(tx.Status)
		if tx.UT.Add(OnChainTime).After(time.Now()) {
			if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_FUND_RELEASE {
				// todo resign?
				resp = newInfoResponse(webapi.CSOperation_CA_USE_RESUMBIT_TOOL, ToolMsg, caseStatus)
			} else if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_SUBMITTING ||
				caseStatus == webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_SGN_CONFIRMATION ||
				caseStatus == webapi.UserCaseStatus_CC_TRANSFER_CONFIRMING_YOUR_REFUND {
				resp = newInfoResponse(webapi.CSOperation_CA_USE_RESYNC_TOOL, ToolMsg, caseStatus)
			} else if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_REQUESTING_REFUND {
				// todo whether report to eng?? note: it will auto signAgain after 15 min
				resp = newInfoResponse(webapi.CSOperation_CA_REPORT, ReportMsg, caseStatus)
			} else {
				resp = newInfoResponse(webapi.CSOperation_CA_NORMAL, NormalMsg, caseStatus)
			}
		} else {
			resp = newInfoResponse(webapi.CSOperation_CA_WAITING, WaitingMsg, caseStatus)
		}
		resp.Info = fmt.Sprintf("transferId:%s, status:%s, updateTime:%s", tx.TransferId, tx.Status.String(), tx.UT.String())
	} else {
		resp = newInfoResponse(webapi.CSOperation_CA_MORE_INFO_NEEDED, CheckInputMsg, webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY)
		return nil
	}
	return resp
}

func diagnosisLp(txHash, lpAddr string, chainId uint32, lpType webapi.LPType) *webapi.GetInfoByTxHashResponse {
	resp := &webapi.GetInfoByTxHashResponse{
		Operation: webapi.CSOperation_CA_NORMAL,
		Memo:      NormalMsg,
	}

	seqNum, status, ut, lpFound, dbErr := dal.DB.GetLPInfoByHash(uint64(lpType), uint64(chainId), lpAddr, txHash)
	if lpFound && dbErr == nil {
		caseStatus := mapLpStatus2CaseStatus(types.WithdrawStatus(status), lpType)
		if ut.Add(OnChainTime).After(time.Now()) {
			if caseStatus == webapi.UserCaseStatus_CC_WAITING_FOR_LP {
				resp = newInfoResponse(webapi.CSOperation_CA_NORMAL, NormalMsg, caseStatus)
			} else if caseStatus == webapi.UserCaseStatus_CC_ADD_SUBMITTING ||
				caseStatus == webapi.UserCaseStatus_CC_ADD_WAITING_FOR_SGN ||
				caseStatus == webapi.UserCaseStatus_CC_WITHDRAW_SUBMITTING {
				resp = newInfoResponse(webapi.CSOperation_CA_USE_RESYNC_TOOL, ToolMsg, caseStatus)
			} else if caseStatus == webapi.UserCaseStatus_CC_WITHDRAW_WAITING_FOR_SGN {
				// todo whether report to eng?? note: it will auto signAgain after 15 min
				resp = newInfoResponse(webapi.CSOperation_CA_REPORT, ReportMsg, caseStatus)
			}
		} else {
			resp = newInfoResponse(webapi.CSOperation_CA_WAITING, WaitingMsg, caseStatus)
		}
		resp.Info = fmt.Sprintf("seqNum:%d, status:%s, updateTime:%s", seqNum, types.WithdrawStatus(status).String(), ut.String())
	} else {
		resp = newInfoResponse(webapi.CSOperation_CA_MORE_INFO_NEEDED, CheckInputMsg, webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY)
		return nil
	}
	return resp
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
	return gs.EC[chainId].TransactionByHash(context.Background(), eth.Hex2Hash(txHash))
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
