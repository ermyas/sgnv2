package gatewaysvc

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	default_slippage = 5000
)

func (gs *GatewayService) GetTransferStatus(ctx context.Context, request *webapi.GetTransferStatusRequest) (*webapi.GetTransferStatusResponse, error) {
	transfer, found, err := dal.DB.GetTransfer(request.GetTransferId())
	if !found || err != nil {
		return &webapi.GetTransferStatusResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "transfer not found",
			},
		}, nil
	}

	var detail *types.QueryLiquidityStatusResponse
	var wdOnchain []byte
	var sortedSigs [][]byte
	var signers [][]byte
	var powers [][]byte

	var transfers []*dal.Transfer
	transfers = append(transfers, transfer)
	refundReasons, err := gs.updateTransferStatusInHistory(ctx, transfers)
	if err != nil {
		return &webapi.GetTransferStatusResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  err.Error(),
			},
		}, nil
	}
	transfer, found, err = dal.DB.GetTransfer(request.GetTransferId())
	if found && err == nil && (transfer.Status == types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND || transfer.Status == types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED) {
		if transfer.RefundSeqNum > 0 {
			if transfer.Status == types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND && time.Now().Add(-15*time.Minute).After(transfer.UT) {
				tr := gs.TP.GetTransactor()
				gs.signAgainWithdraw(&types.MsgSignAgain{
					DataType: types.SignDataType_WITHDRAW,
					Creator:  tr.Key.GetAddress().String(),
					ReqId:    transfer.RefundSeqNum,
					UserAddr: common.Hex2Addr(transfer.UsrAddr).Bytes(),
				})
				// update db: refresh update_time, so that will sign again after 15 min
				dal.DB.UpdateTransferStatus(transfer.TransferId, uint64(types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND))
			}
			detail, wdOnchain, sortedSigs, signers, powers = gs.getWithdrawInfo(transfer.RefundSeqNum, transfer.SrcChainId, transfer.UsrAddr)
			if detail == nil {
				return &webapi.GetTransferStatusResponse{
					Err: &webapi.ErrMsg{
						Code: webapi.ErrCode_ERROR_CODE_COMMON,
						Msg:  "withdrawInfo not found",
					},
				}, nil
			}
			log.Infof("get lp info for transfer:%s, lp status is :%s", transfer.TransferId, detail.GetStatus().String())
			if detail.GetStatus() == types.WithdrawStatus_WD_WAITING_FOR_LP && transfer.Status != types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED {
				log.Warnf("update transfer:%s by seqNum: %d, from %s, to %s", transfer.TransferId, transfer.RefundSeqNum, transfer.Status, types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED)
				dbErr := dal.DB.UpdateTransferStatus(transfer.TransferId, uint64(types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED))
				if dbErr != nil {
					log.Warnf("UpdateTransferStatus failed, transferId:%s, status:%s", transfer.TransferId, types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED.String())
				}
				transfer.Status = types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED
			}
		} else {
			log.Errorf("transfer seq num not found for transfer:%s", transfer.TransferId)
		}
	}

	blockDelay, found, err := dal.DB.GetChainBlockDelay(transfer.SrcChainId)
	if !found || err != nil {
		blockDelay = 0
	}
	srcTxHash, dstTxHash := gs.getTxHashForTransfer(transfer)
	return &webapi.GetTransferStatusResponse{
		Status:         transfer.Status,
		WdOnchain:      wdOnchain,
		SortedSigs:     sortedSigs,
		Signers:        signers,
		Powers:         powers,
		RefundReason:   refundReasons[transfer.TransferId],
		BlockDelay:     blockDelay,
		SrcBlockTxLink: srcTxHash,
		DstBlockTxLink: dstTxHash,
	}, nil
}

func (gs *GatewayService) EstimateAmt(ctx context.Context, request *webapi.EstimateAmtRequest) (*webapi.EstimateAmtResponse, error) {
	amt := request.GetAmt()
	srcChainId := request.GetSrcChainId()
	dstChainId := request.GetDstChainId()
	tokenSymbol := request.GetTokenSymbol()
	slippage := request.GetSlippageTolerance()
	if slippage == 0 {
		slippage = default_slippage
	}

	if srcChainId == dstChainId {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "src chain can not be the same with dst chain",
			},
		}, nil
	}
	srcToken, found1, err1 := dal.DB.GetTokenBySymbolForTransfer(tokenSymbol, uint64(srcChainId))
	if err1 != nil || !found1 {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_NO_TOKEN_ON_SRC_CHAIN,
				Msg:  "token not support on src chain",
			},
		}, nil
	}
	dstToken, found2, err2 := dal.DB.GetTokenBySymbolForTransfer(tokenSymbol, uint64(dstChainId))
	if err2 != nil || !found2 {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_NO_TOKEN_ON_DST_CHAIN,
				Msg:  "token not support on dst chain",
			},
		}, nil
	}
	addr := common.Hex2Addr(request.GetUsrAddr()).String()

	resp, infoErr := gs.getEstimatedFeeInfo(addr, srcChainId, dstChainId, slippage, srcToken, dstToken, amt, false)
	if infoErr != nil {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  infoErr.Error(),
			},
		}, nil
	} else {
		return resp, nil
	}
}

func (gs *GatewayService) MarkTransfer(ctx context.Context, request *webapi.MarkTransferRequest) (*webapi.MarkTransferResponse, error) {
	transferId := request.GetTransferId()
	addr := common.Hex2Addr(request.GetAddr())
	sendInfo := refineTokenInfo(request.GetSrcSendInfo())
	receivedInfo := refineTokenInfo(request.GetDstMinReceivedInfo())
	txHash := request.GetSrcTxHash()
	txType := request.GetType()
	log.Infof("Mark transfer, transferId: %s, addr:%s, txHash: %s, srcChainId:%d, txType:%d", transferId, addr, txHash, sendInfo.GetChain().GetId(), txType)
	if !utils.CheckMarkTransferParams(transferId, txHash, request.GetAddr(), sendInfo, receivedInfo) {
		return &webapi.MarkTransferResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "params checking failed",
			},
		}, nil
	}
	if txType == webapi.TransferType_TRANSFER_TYPE_SEND {
		srcChainId := uint64(sendInfo.GetChain().GetId())
		dstChainId := uint64(receivedInfo.GetChain().GetId())
		percentage := uint32(400) // default 0.04%
		tr := gs.TP.GetTransactor()
		perc, err := cbrcli.QueryFeePerc(tr.CliCtx, &types.GetFeePercentageRequest{
			SrcChainId: srcChainId,
			DstChainId: dstChainId,
		})
		if err == nil && perc != nil && perc.FeePerc > 0 {
			percentage = perc.FeePerc
		} else {
			log.Warnf("get perc failed:srcChain:%d, dstChain:%d, perc:%+v, err:%+v", srcChainId, dstChainId, perc, err)
		}

		err = dal.DB.MarkTransferSend(transferId, addr.String(), sendInfo.GetToken().GetSymbol(),
			sendInfo.GetAmount(), receivedInfo.GetAmount(), txHash, srcChainId,
			dstChainId, gs.F.GetUsdVolume(sendInfo.GetToken(), common.Str2BigInt(sendInfo.GetAmount())), percentage)
		if err != nil {
			return &webapi.MarkTransferResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "mark transfer refund failed",
				},
			}, nil
		}
	} else if txType == webapi.TransferType_TRANSFER_TYPE_REFUND {
		err := dal.DB.MarkTransferRefund(transferId, txHash)
		if err != nil {
			return &webapi.MarkTransferResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "mark transfer refund failed",
				},
			}, nil
		}
	}
	return &webapi.MarkTransferResponse{
		Err: nil,
	}, nil
}

func (gs *GatewayService) TransferHistory(ctx context.Context, request *webapi.TransferHistoryRequest) (*webapi.TransferHistoryResponse, error) {
	addr := common.Hex2Addr(request.GetAddr()).String()
	endTime := time.Now()
	if request.GetNextPageToken() != "" {
		ts, err := strconv.Atoi(request.GetNextPageToken())
		if err != nil {
			return &webapi.TransferHistoryResponse{}, nil
		}
		endTime = common.TsToTime(uint64(ts))
	}
	transferList, currentPageSize, next, err := dal.DB.PaginateTransferList(addr, endTime, request.GetPageSize())
	if err != nil {
		return &webapi.TransferHistoryResponse{}, nil
	}
	refundReasons, err := gs.updateTransferStatusInHistory(ctx, transferList)
	if err != nil {
		log.Warnf("update transfer status failed for user:%s, error:%v", addr, err)
	}
	var transfers []*webapi.TransferHistory
	for _, transfer := range transferList {
		srcChain, srcChainUrl, srcFound, err1 := dal.DB.GetChain(transfer.SrcChainId)
		dstChain, dstChainUrl, dstFound, err2 := dal.DB.GetChain(transfer.DstChainId)
		if !srcFound || !dstFound || err1 != nil || err2 != nil {
			continue
		}
		if !srcFound {
			srcChain = unknownChain(uint32(transfer.SrcChainId))
		} else {
			srcChain = enrichChainUiInfo(srcChain)
		}
		if !dstFound {
			dstChain = unknownChain(uint32(transfer.DstChainId))
		} else {
			dstChain = enrichChainUiInfo(dstChain)
		}
		srcToken, srcFound, err1 := dal.DB.GetTokenBySymbol(transfer.TokenSymbol, transfer.SrcChainId)
		dstToken, dstFound, err2 := dal.DB.GetTokenBySymbol(transfer.TokenSymbol, transfer.DstChainId)
		if !srcFound || !dstFound || err1 != nil || err2 != nil {
			continue
		}
		srcTxLink := ""
		dstTxLink := ""
		if srcChainUrl != "" && common.IsValidTxHash(transfer.SrcTxHash) {
			srcTxLink = srcChainUrl + transfer.SrcTxHash
		}

		if dstChainUrl != "" && common.IsValidTxHash(transfer.DstTxHash) {
			dstTxLink = dstChainUrl + transfer.DstTxHash
		}

		transfers = append(transfers, &webapi.TransferHistory{
			TransferId: transfer.TransferId,
			SrcSendInfo: &webapi.TransferInfo{
				Chain:  srcChain,
				Token:  srcToken.GetToken(),
				Amount: transfer.SrcAmt,
			},
			DstReceivedInfo: &webapi.TransferInfo{
				Chain:  dstChain,
				Token:  dstToken.GetToken(),
				Amount: transfer.DstAmt,
			},
			Ts:             common.TsMilli(transfer.CT),
			SrcBlockTxLink: srcTxLink,
			DstBlockTxLink: dstTxLink,
			Status:         transfer.Status,
			RefundReason:   refundReasons[transfer.TransferId],
		})
	}
	return &webapi.TransferHistoryResponse{
		History:       transfers,
		NextPageToken: strconv.FormatUint(common.TsMilli(next), 10),
		CurrentSize:   uint64(currentPageSize),
	}, nil
}

// ================================= internal method below =====================================

func refineTokenInfo(token *webapi.TransferInfo) *webapi.TransferInfo {
	t, found, err := dal.DB.GetTokenBySymbol(token.GetToken().GetSymbol(), uint64(token.GetChain().GetId()))
	if !found || err != nil {
		log.Errorf("can not find token in db, token:%s, chain:%d", token.GetToken().GetSymbol(), token.GetChain().GetId())
		return token
	}
	token.Token = t.Token
	return token
}

func (gs *GatewayService) updateTransferStatusInHistory(ctx context.Context, transferList []*dal.Transfer) (map[string]types.XferStatus, error) {
	var transferIds []string
	refundReasons := make(map[string]types.XferStatus)
	for _, transfer := range transferList {
		transferIds = append(transferIds, transfer.TransferId)
	}
	tr := gs.TP.GetTransactor()
	transferMap, err := cbrcli.QueryTransferStatus(tr.CliCtx, &types.QueryTransferStatusRequest{
		TransferId: transferIds,
	})
	if err != nil {
		log.Errorf("updateTransferStatusInHistory when QueryTransferStatus in sgn failed, error: %+v", err)
		return refundReasons, err
	}
	transferStatusMap := transferMap.Status

	for _, transfer := range transferList {
		refundReason := types.XferStatus_UNKNOWN
		transferId := transfer.TransferId
		status := transfer.Status
		srcChainId := transfer.SrcChainId
		txHash := transfer.SrcTxHash
		if status == types.TransferHistoryStatus_TRANSFER_SUBMITTING {
			ec := gs.EC[srcChainId]
			if ec == nil {
				log.Errorf("no ethClient found for chain:%d", srcChainId)
				continue
			}
			receipt, recErr := ec.TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(txHash)))
			if recErr == nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
				log.Warnf("find transfer failed, chain_id %d, hash:%s", srcChainId, txHash)
				dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(types.TransferHistoryStatus_TRANSFER_FAILED))
				if dbErr != nil {
					log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, txHash)
				}
			}
		}

		if status == types.TransferHistoryStatus_TRANSFER_FAILED ||
			status == types.TransferHistoryStatus_TRANSFER_COMPLETED ||
			status == types.TransferHistoryStatus_TRANSFER_REFUNDED {
			continue // finial status, not updated by sgn
		}

		// revert status if onchain refund failed
		if status == types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND {
			chainId := transfer.SrcChainId
			refundTx := transfer.RefundTx
			if refundTx == "" {
				log.Errorf("refund tx is nil for transfer refund:%s", transferId)
			} else {
				ec := gs.EC[chainId]
				if ec == nil {
					log.Errorf("no ethClient found for chain:%d", chainId)
					return nil, fmt.Errorf("no ethClient found for chain:%d", chainId)
				}
				receipt, recErr := ec.TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(refundTx)))
				if recErr == nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
					dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED))
					if dbErr != nil {
						log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, refundTx)
					}
					status = types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED
				}
			}
		}

		// only 3 status below will be updated by sgn query
		if transferStatusMap[transferId].GetGatewayStatus() == types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED ||
			transferStatusMap[transferId].GetGatewayStatus() == types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED ||
			transferStatusMap[transferId].GetGatewayStatus() == types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE {
			if status == types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND || status == types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND {
				continue // user action, not updated by sgn
			}
			if status == types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED && transferStatusMap[transferId].GetGatewayStatus() == types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED {
				continue // waiting for user confirmed but sgn doesn't know, skip
			}
			if status == transferStatusMap[transferId].GetGatewayStatus() {
				//status not change in polling
				continue
			}
			log.Infof("update transfer refund status from sgn, transfer:%s, current is %s, dst is %s", transferId, status.String(), transferStatusMap[transferId].String())
			dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(transferStatusMap[transferId].GetGatewayStatus()))
			if dbErr != nil {
				log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, txHash)
			}
		}
		refundReason = transferStatusMap[transferId].SgnStatus
		refundReasons[transferId] = refundReason
	}
	return refundReasons, nil
}

func (gs *GatewayService) getTxHashForTransfer(transfer *dal.Transfer) (string, string) {
	srcTxHash, dstTxHash := "", ""
	if transfer.Status == types.TransferHistoryStatus_TRANSFER_COMPLETED || transfer.Status == types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE {
		_, url1, found1, err1 := dal.DB.GetChain(transfer.SrcChainId)
		if found1 && err1 == nil && url1 != "" && common.IsValidTxHash(transfer.SrcTxHash) {
			srcTxHash = url1 + transfer.SrcTxHash
		}
		_, url2, found2, err2 := dal.DB.GetChain(transfer.DstChainId)
		if found2 && err2 == nil && url2 != "" && common.IsValidTxHash(transfer.DstTxHash) {
			dstTxHash = url2 + transfer.DstTxHash
		}
	}
	return srcTxHash, dstTxHash
}

func (gs *GatewayService) getEstimatedFeeInfo(addr string, srcChainId, dstChainId, slippage uint32, srcToken, dstToken *webapi.TokenInfo, amt string, useLp bool) (*webapi.EstimateAmtResponse, error) {
	if !utils.IsvalidAmt(amt) {
		return nil, fmt.Errorf("invalid amt, params checking failed")
	}
	tr := gs.TP.GetTransactor()
	getFeeRequest := &types.GetFeeRequest{
		SrcChainId:   uint64(srcChainId),
		DstChainId:   uint64(dstChainId),
		SrcTokenAddr: srcToken.Token.GetAddress(),
		Amt:          amt,
	}
	if useLp {
		getFeeRequest.LpAddr = addr
	}
	feeInfo, err := cbrcli.QueryFee(tr.CliCtx, getFeeRequest)
	if err != nil {
		log.Warnf("cli.QueryFee error, srcChainId:%d, dstChainId:%d, srcTokenAddr:%s, amt:%s, err:%+v", srcChainId, dstChainId, srcToken.Token.GetAddress(), amt, err)
		return nil, err
	}
	if feeInfo == nil {
		return nil, fmt.Errorf("can not estimate fee")
	}
	eqValueTokenAmt := feeInfo.GetEqValueTokenAmt()
	percFee := feeInfo.GetPercFee()
	baseFee := feeInfo.GetBaseFee()
	feeAmt := new(big.Int).Add(common.Str2BigInt(percFee), common.Str2BigInt(baseFee))
	srcVolume := gs.F.GetUsdVolume(srcToken.Token, common.Str2BigInt(amt))
	dstVolume := gs.F.GetUsdVolume(dstToken.Token, common.Str2BigInt(eqValueTokenAmt))
	bridgeRate := 0.0
	if srcVolume > 0.000000001 {
		bridgeRate = dstVolume / srcVolume
	} else {
		return nil, fmt.Errorf("amount should > 0")
	}
	minReceiveVolume := dstVolume*(1-float64(slippage)/1e6) - gs.F.GetUsdVolume(dstToken.Token, feeAmt)
	return &webapi.EstimateAmtResponse{
		EqValueTokenAmt:   eqValueTokenAmt,
		BridgeRate:        float32(bridgeRate),
		PercFee:           percFee,
		BaseFee:           baseFee,
		SlippageTolerance: slippage,
		MaxSlippage:       uint32((srcVolume - minReceiveVolume) * 1e6 / srcVolume),
	}, nil
}
