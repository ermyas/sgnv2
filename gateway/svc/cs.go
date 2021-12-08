package gatewaysvc

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/onchain"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/ops"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

const (
	OnChainTime = 15 * time.Minute

	NormalMsg     = "normal case, open history and operate according to tips, or report it to eng team if has more problems"
	ToolMsg       = "try to use fix tools(click fix button) and search again, if problem not fixed after using tools report it to eng team"
	WaitingMsg    = "too short time after user operation, keep waiting for a few minutes"
	CheckInputMsg = "can not find any result, check your input txHash and chain. If input is correct, waiting for 15 min. if you have waited longer than 15 min, report it to eng team"
)

func (gs *GatewayService) GetInfoByTxHash(ctx context.Context, request *webapi.GetInfoByTxHashRequest) (*webapi.GetInfoByTxHashResponse, error) {
	if !checkSigner(common.Hex2Addr(request.GetAddr()).Bytes(), request.GetSig()) {
		return &webapi.GetInfoByTxHashResponse{
			Info: "invalid operator",
		}, nil
	}
	if !gs.checkTxExits(ctx, request.GetTxHash(), request.GetChainId()) {
		return &webapi.GetInfoByTxHashResponse{
			Info: "can not find tx or error tx status, please check again",
		}, nil
	}
	return gs.checkCaseStatus(ctx, request.GetType(), request.GetTxHash(), request.GetChainId(), request.GetType()), nil
}

func (gs *GatewayService) FixEventMiss(ctx context.Context, request *webapi.FixEventMissRequest) (*webapi.FixEventMissResponse, error) {
	if !checkSigner(common.Hex2Addr(request.GetAddr()).Bytes(), request.GetSig()) {
		return &webapi.FixEventMissResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "invalid operator",
			},
		}, nil
	}
	if !gs.checkTxExits(ctx, request.GetTxHash(), request.GetChainId()) {
		return &webapi.FixEventMissResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "can not find tx or error tx status, please check again",
			},
		}, nil
	}
	txHash := request.GetTxHash()
	chainId := request.GetChainId()
	status := request.GetType()
	switch status {
	case webapi.CSType_CT_TX:
		err := gs.fixTx(ctx, txHash, chainId, status)
		if err != nil {
			return &webapi.FixEventMissResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  err.Error(),
				},
			}, nil
		}
	case webapi.CSType_CT_LP_ADD:
		lpAddr, err := gs.getAddrFromHash(ctx, txHash, uint64(chainId))
		if err == nil {
			err = gs.fixLp(ctx, txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_ADD, status)
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
		lpAddr, err := gs.getAddrFromHash(ctx, txHash, uint64(chainId))
		if err == nil {
			err = gs.fixLp(ctx, txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_REMOVE, status)
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

func (gs *GatewayService) checkCaseStatus(ctx context.Context, status webapi.CSType, txHash string, chainId uint32, csType webapi.CSType) *webapi.GetInfoByTxHashResponse {
	switch status {
	case webapi.CSType_CT_TX:
		return gs.diagnosisTx(ctx, txHash, chainId, csType)
	case webapi.CSType_CT_LP_ADD:
		lpAddr, err := gs.getAddrFromHash(ctx, txHash, uint64(chainId))
		if err != nil {
			return &webapi.GetInfoByTxHashResponse{
				Info: "can not find lp addr from txHash and chainId",
			}
		}
		return gs.diagnosisLp(ctx, txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_ADD, csType)
	case webapi.CSType_CT_LP_RM:
		lpAddr, err := gs.getAddrFromHash(ctx, txHash, uint64(chainId))
		if err != nil {
			return &webapi.GetInfoByTxHashResponse{
				Info: "can not find lp addr from txHash and chainId",
			}
		}
		return gs.diagnosisLp(ctx, txHash, lpAddr, chainId, webapi.LPType_LP_TYPE_REMOVE, csType)
	}
	return &webapi.GetInfoByTxHashResponse{}
}

func (gs *GatewayService) diagnosisTx(ctx context.Context, txHash string, chainId uint32, csType webapi.CSType) *webapi.GetInfoByTxHashResponse {
	resp := &webapi.GetInfoByTxHashResponse{
		Operation: webapi.CSOperation_CA_NORMAL,
		Memo:      NormalMsg,
	}
	tx0, txFound, dbErr := dal.DB.GetTransferBySrcTxHash(txHash, chainId)
	if txFound && dbErr == nil {
		_, _ = gs.GetTransferStatus(ctx, &webapi.GetTransferStatusRequest{TransferId: tx0.TransferId})
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
		srcToken, _, _ := dal.DB.GetTokenBySymbol(tx.TokenSymbol, tx.SrcChainId)
		dstToken, _, _ := dal.DB.GetTokenBySymbol(tx.TokenSymbol, tx.SrcChainId)
		srcAmt := rmAmtDec(tx.SrcAmt, int(srcToken.GetToken().Decimal))
		dstAmt := rmAmtDec(tx.DstAmt, int(dstToken.GetToken().Decimal))
		resp.Info = fmt.Sprintf(
			"transferId: %s, \n"+
				"token: %s, \n"+
				"dstChainId: %d, \n"+
				"status: %s, \n"+
				"addr: %s, \n"+
				"updateTime: %s, \n"+
				"createTime: %s, \n"+
				"srcAmt: %.6f, \n"+
				"dstAmt: %.6f, \n"+
				"refundTx: %s, \n"+
				"refundSeqNum: %d",
			tx.TransferId, tx.TokenSymbol, tx.DstChainId, tx.Status.String(), tx.UsrAddr, tx.UT.String(), tx.CT.String(), srcAmt, dstAmt, tx.RefundTx, tx.RefundSeqNum)
	} else {
		if gs.isTxEventMissFixable(ctx, txHash, chainId, csType) {
			resp = newInfoResponse(webapi.CSOperation_CA_USE_RESYNC_TOOL, ToolMsg, webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY)
			resp.Info = "no history found in backend, but you can try to fix it with fix button blow. After that, click 'search' again to check fix result"
		} else {
			resp = newInfoResponse(webapi.CSOperation_CA_MORE_INFO_NEEDED, CheckInputMsg, webapi.UserCaseStatus_CC_TRANSFER_NO_HISTORY)
		}
	}
	return resp
}

func (gs *GatewayService) isTxEventMissFixable(ctx context.Context, txHash string, chainId uint32, csType webapi.CSType) bool {
	eventName := getEventName(csType)
	if eventName == "" {
		return false
	}
	cli := gs.Chains.GetEthClient(uint64(chainId))
	if cli == nil {
		return false
	}
	elog, err := onchain.GetCbrLog(uint64(chainId), txHash, eventName)
	if err != nil || elog == nil {
		return false
	}

	chain, _, _, _ := dal.DB.GetChain(uint64(chainId))
	if chain != nil {
		if elog.Address != eth.Hex2Addr(chain.ContractAddr) {
			log.Warnf("isTxEventMissFixable failed, chain id check failed, chainId:%d, addrOnChain:%s, addInTx:%s", chainId, chain.ContractAddr, elog.Address.String())
			return false
		}
	} else {
		return false
	}
	parser, err := eth.NewBridgeFilterer(eth.ZeroAddr, nil)

	if err != nil {
		return false
	}
	if csType == webapi.CSType_CT_TX {
		_, parseErr := parser.ParseSend(*elog)
		if parseErr != nil {
			return false
		}
	} else if csType == webapi.CSType_CT_LP_ADD {
		_, parseErr := parser.ParseLiquidityAdded(*elog)
		if parseErr != nil {
			return false
		}
	} else if csType == webapi.CSType_CT_LP_RM {
		_, parseErr := parser.ParseWithdrawDone(*elog)
		if parseErr != nil {
			return false
		}
	}

	receipt, err := cli.TransactionReceipt(ctx, eth.Hex2Hash(txHash))
	if receipt == nil || err != nil {
		return false
	}
	block, err2 := cli.BlockByHash(ctx, receipt.BlockHash)
	if block != nil && err2 == nil {
		blockTime := common.TsSecToTime(block.Time())
		blockNum := block.NumberU64()
		latestBlockNum, err3 := cli.BlockNumber(ctx)
		if err3 == nil {
			confirmation := latestBlockNum - blockNum
			log.Infof("confirmation:%d, blockTime:%s", confirmation, blockTime.String())
			return confirmation > 50 && blockTime.Add(OnChainTime).Before(time.Now())
		}
	}
	return false
}

func getEventName(csType webapi.CSType) string {
	eventName := ""
	if csType == webapi.CSType_CT_TX {
		eventName = "Send"
	} else if csType == webapi.CSType_CT_LP_ADD {
		eventName = "LiquidityAdded"
	} else if csType == webapi.CSType_CT_LP_RM {
		eventName = "WithdrawDone"
	}
	return eventName
}

func (gs *GatewayService) fixTxEventMiss(ctx context.Context, txHash string, chainId uint32, csType webapi.CSType) error {
	eventName := getEventName(csType)
	if eventName == "" {
		return fmt.Errorf("error cs type, only tx or lp_add is valid")
	}

	elog, err := onchain.GetCbrLog(uint64(chainId), txHash, eventName)
	if elog == nil || err != nil {
		log.Warnf("GetCbrLog failed,chainId:%d, txHash:%s, eventName:%s, elog:%+v, err:%+v", uint64(chainId), txHash, eventName, elog, err)
		return err
	}
	parser, err := eth.NewBridgeFilterer(eth.ZeroAddr, nil)

	if parser == nil || err != nil {
		log.Warnf("get parser failed, parser:%+v, err:%+v", parser, err)
		return err
	}

	if csType == webapi.CSType_CT_TX {
		ev, parseErr := parser.ParseSend(*elog)

		if parseErr != nil {
			return parseErr
		}
		if ev == nil {
			return fmt.Errorf("parse failed from elog:%+v", elog)
		}

		err = onchain.GatewayOnSend(common.Hash(ev.TransferId).String(), ev.Sender.String(), ev.Token.String(), ev.Amount.String(), txHash, uint64(chainId), ev.DstChainId)
		if err != nil {
			return err
		}
		transfer, txFound, dbErr := dal.DB.GetTransfer(common.Hash(ev.TransferId).String())
		if txFound && transfer != nil && dbErr == nil {
			var transfers []*dal.Transfer
			transfers = append(transfers, transfer)
			gs.updateTransferStatusInHistory(ctx, transfers)
		}
		transfer, _, _ = dal.DB.GetTransfer(transfer.TransferId)
		if transfer != nil {
			if transfer.Status == types.TransferHistoryStatus_TRANSFER_COMPLETED || transfer.Status == types.TransferHistoryStatus_TRANSFER_DELAYED {
				tr := onchain.SGNTransactors.GetTransactor()
				relay, cliErr := cbrcli.QueryRelay(tr.CliCtx, eth.Hex2Hash(transfer.TransferId).Bytes())
				if cliErr != nil {
					return cliErr
				}
				relayOnChain := new(types.RelayOnChain)
				err = relayOnChain.Unmarshal(relay.Relay)
				if err != nil {
					return err
				}
				err = onchain.GatewayOnRelay(gs.Chains.GetEthClient(uint64(chainId)), transfer.TransferId, "", relayOnChain.GetRelayOnChainTransferId().String(), new(big.Int).SetBytes(relayOnChain.GetAmount()).String())
				if err != nil {
					return err
				}
			}
		}
	} else if csType == webapi.CSType_CT_LP_ADD {
		ev, parseErr := parser.ParseLiquidityAdded(*elog)

		if parseErr != nil {
			return parseErr
		}

		if ev == nil {
			return fmt.Errorf("parse failed from elog:%+v", elog)
		}

		token, found, dbErr := dal.DB.GetTokenByAddr(ev.Token.String(), uint64(chainId))
		if !found || dbErr != nil {
			return fmt.Errorf("token not found:%s, on chain%d", ev.Token.String(), chainId)
		}

		cli := gs.Chains.GetEthClient(uint64(chainId))
		if cli == nil {
			return fmt.Errorf("ec for chain:%d not found", chainId)
		}

		tx, _, txErr := cli.TransactionByHash(ctx, eth.Hex2Hash(txHash))

		if tx != nil && txErr == nil {
			err = onchain.GatewayOnLiqAdd(ev.Provider.String(), token.Token.Address, ev.Amount.String(), txHash, uint64(chainId), ev.Seqnum, tx.Nonce())
			if err != nil {
				return err
			}
			lpHistory, _, _, _ := dal.DB.PaginateLpHistory(ev.Provider.String(), time.Now(), 1000)
			if lpHistory != nil && len(lpHistory) > 0 {
				gs.updateLpStatusInHistory(lpHistory)
			}
		} else {
			return fmt.Errorf("get nonce failed,chainId:%d, TxHash:%s, err: %s", chainId, txHash, txErr)
		}
	} else if csType == webapi.CSType_CT_LP_RM {
		ev, parseErr := parser.ParseWithdrawDone(*elog)
		if parseErr != nil {
			return parseErr
		}
		idstr := common.Hash(ev.WithdrawId).String()
		onchain.GatewayOnLiqWithdraw(idstr, elog.TxHash.String(), uint64(chainId), ev.Seqnum, ev.Receiver.String())
		lpHistory, _, _, _ := dal.DB.PaginateLpHistory(ev.Receiver.String(), time.Now(), 1000)
		if lpHistory != nil && len(lpHistory) > 0 {
			gs.updateLpStatusInHistory(lpHistory)
		}
	}
	return nil
}

func rmAmtDec(amt string, decimal int) float64 {
	f, _ := new(big.Float).Quo(new(big.Float).SetInt(common.Str2BigInt(amt)), big.NewFloat(math.Pow10(decimal))).Float64()
	return f
}

func (gs *GatewayService) diagnosisLp(ctx context.Context, txHash, lpAddr string, chainId uint32, lpType webapi.LPType, csType webapi.CSType) *webapi.GetInfoByTxHashResponse {
	resp := &webapi.GetInfoByTxHashResponse{
		Operation: webapi.CSOperation_CA_NORMAL,
		Memo:      NormalMsg,
	}
	seqNum0, _, ut, lpFound, dbErr := dal.DB.GetLPInfoByHash(uint64(lpType), uint64(chainId), lpAddr, txHash)
	if lpFound && dbErr == nil {
		_, _ = gs.QueryLiquidityStatus(ctx, &webapi.QueryLiquidityStatusRequest{
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
				if gs.isTxEventMissFixable(ctx, txHash, chainId, csType) {
					resp = newInfoResponse(webapi.CSOperation_CA_USE_RESYNC_TOOL, ToolMsg, webapi.UserCaseStatus_CC_ADD_NO_HISTORY)
				} else {
					resp = newInfoResponse(webapi.CSOperation_CA_NORMAL, NormalMsg, caseStatus)
				}
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
		methodType, tokenSymbol, amt := dal.DB.GetCsInfoByHash(uint64(lpType), uint64(chainId), lpAddr, txHash)
		token, _, _ := dal.DB.GetTokenBySymbol(tokenSymbol, uint64(chainId))
		amtF := rmAmtDec(amt, int(token.GetToken().Decimal))
		t := ""
		if methodType == webapi.WithdrawMethodType_WD_METHOD_TYPE_ONE_RM {
			t = "common remove liquidity"
		} else if methodType == webapi.WithdrawMethodType_WD_METHOD_TYPE_ALL_IN_ONE {
			t = "single chain remove liquidity"
		}
		resp.Info = fmt.Sprintf(
			"seqNum: %d, \n"+
				"status: %s, \n"+
				"addr: %s, \n"+
				"token: %s, \n"+
				"amt: %.6f, \n"+
				"updateTime: %s\n",
			seqNum, types.WithdrawStatus(status).String(), lpAddr, tokenSymbol, amtF, ut.String())
		if lpType == webapi.LPType_LP_TYPE_REMOVE {
			resp.Info = resp.Info + fmt.Sprintf("type: %s", t)
		}
	} else {
		resp = newInfoResponse(webapi.CSOperation_CA_MORE_INFO_NEEDED, CheckInputMsg, webapi.UserCaseStatus_CC_ADD_NO_HISTORY)
	}
	return resp
}

func (gs *GatewayService) fixTx(ctx context.Context, txHash string, chainId uint32, csType webapi.CSType) error {
	tx, txFound, dbErr := dal.DB.GetTransferBySrcTxHash(txHash, chainId)
	if txFound && dbErr == nil {
		caseStatus := mapTxStatus2CaseStatus(tx.Status)
		if tx.UT.Add(OnChainTime).Before(time.Now()) {
			if caseStatus == webapi.UserCaseStatus_CC_TRANSFER_WAITING_FOR_FUND_RELEASE || caseStatus == webapi.UserCaseStatus_CC_TRANSFER_REQUESTING_REFUND {
				log.Infof("cs fix tx by resign, txHash:%s, chainId:%d, txId:%s", txHash, chainId, tx.TransferId)
				dal.DB.UpdateTransferStatus(tx.TransferId, uint64(tx.Status))
				_, err := gs.signAgainWithdraw(&types.MsgSignAgain{
					DataType: types.SignDataType_RELAY,
					Creator:  onchain.SGNTransactors.GetTransactor().Key.GetAddress().String(),
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
					err = ops.SyncCbrEvent(onchain.SGNTransactors.GetTransactor().CliCtx, uint64(chainId), txHash, types.CbrEventSend)
				} else {
					err = ops.SyncCbrEvent(onchain.SGNTransactors.GetTransactor().CliCtx, uint64(chainId), tx.DstTxHash, types.CbrEventRelay)
				}
				if err != nil {
					return err
				}

			}
		} else {
			return fmt.Errorf("frequence limited, please operate after until:%s", tx.UT.Add(OnChainTime).String())
		}
	} else if gs.isTxEventMissFixable(ctx, txHash, chainId, csType) {
		return gs.fixTxEventMiss(ctx, txHash, chainId, webapi.CSType_CT_TX)
	}
	return nil
}

func (gs *GatewayService) fixLp(ctx context.Context, txHash, lpAddr string, chainId uint32, lpType webapi.LPType, csType webapi.CSType) error {
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
					Creator:  onchain.SGNTransactors.GetTransactor().Key.GetAddress().String(),
					ReqId:    seqNum,
					UserAddr: eth.Hex2Addr(lpAddr).Bytes(),
				})
				if err != nil {
					return err
				}
			} else if caseStatus == webapi.UserCaseStatus_CC_WAITING_FOR_LP && gs.isTxEventMissFixable(ctx, txHash, chainId, csType) {
				return gs.fixTxEventMiss(ctx, txHash, chainId, csType)
			} else if caseStatus == webapi.UserCaseStatus_CC_ADD_SUBMITTING ||
				caseStatus == webapi.UserCaseStatus_CC_ADD_WAITING_FOR_SGN ||
				caseStatus == webapi.UserCaseStatus_CC_WITHDRAW_SUBMITTING {
				log.Infof("cs fix lp by resync, txHash:%s, chainId:%d, lpAddr:%s, lpType:%s", txHash, chainId, lpAddr, lpType.String())
				// refresh update time
				dal.DB.UpdateLPStatus(seqNum, uint64(lpType), uint64(chainId), lpAddr, status)
				var err error
				if lpType == webapi.LPType_LP_TYPE_ADD {
					err = ops.SyncCbrEvent(onchain.SGNTransactors.GetTransactor().CliCtx, uint64(chainId), txHash, types.CbrEventLiqAdd)
				} else if lpType == webapi.LPType_LP_TYPE_REMOVE {
					err = ops.SyncCbrEvent(onchain.SGNTransactors.GetTransactor().CliCtx, uint64(chainId), txHash, types.CbrEventWithdraw)
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
	} else if lpType == webapi.LPType_LP_TYPE_ADD && gs.isTxEventMissFixable(ctx, txHash, chainId, csType) {
		return gs.fixTxEventMiss(ctx, txHash, chainId, csType)
	}
	return nil
}

func (gs *GatewayService) getAddrFromHash(ctx context.Context, txHash string, chainId uint64) (string, error) {
	tx, _, err := gs.getTransactionByHash(ctx, txHash, chainId)
	if err != nil {
		return "", err
	}
	sender, err := ethtypes.Sender(ethtypes.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return "", err
	}
	return sender.String(), nil
}

func (gs *GatewayService) getTransactionByHash(ctx context.Context, txHash string, chainId uint64) (*ethtypes.Transaction, bool, error) {
	ec := gs.Chains.GetEthClient(chainId)
	if ec == nil {
		return nil, false, fmt.Errorf("eth client not found for chainId:%d", chainId)
	}
	return ec.TransactionByHash(ctx, eth.Hex2Hash(txHash))
}

func (gs *GatewayService) checkTxExits(ctx context.Context, txHash string, chainId uint32) bool {
	client := gs.Chains.GetEthClient(uint64(chainId))
	if client == nil {
		return false
	}
	receipt, recErr := client.TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(txHash)))
	if receipt == nil || recErr != nil || receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return false
	}
	return true
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
