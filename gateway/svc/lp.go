package gatewaysvc

import (
	"context"
	"fmt"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"time"
)

type txData struct {
	volume   float64
	fee      *big.Int
	dstToken *types.Token
}

func (gs *GatewayService) MarkLiquidity(ctx context.Context, request *webapi.MarkLiquidityRequest) (*webapi.MarkLiquidityResponse, error) {
	lpType := request.GetType()
	chainId := request.GetChainId()
	amt := request.GetAmt()
	addr := common.Hex2Addr(request.GetLpAddr()).String()
	tokenAddr := common.Hex2Addr(request.GetTokenAddr()).String()
	log.Infof("Liquidity in mark api request:%+v", request)
	token, found, err := dal.DB.GetTokenByAddr(tokenAddr, uint64(chainId))
	if !found || err != nil {
		return &webapi.MarkLiquidityResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "token not found in gateway DB",
			},
		}, nil
	}
	txHash := request.GetTxHash()
	if lpType == webapi.LPType_LP_TYPE_ADD {
		err = dal.DB.UpsertLPWithTx(addr, token.GetToken().GetSymbol(), token.GetToken().GetAddress(), amt, txHash, uint64(chainId), uint64(types.LPHistoryStatus_LP_SUBMITTING), uint64(lpType), 0)
	} else if lpType == webapi.LPType_LP_TYPE_REMOVE {
		seqNum := request.GetSeqNum()
		err = dal.DB.UpsertLPWithSeqNum(addr, token.GetToken().GetSymbol(), token.GetToken().GetAddress(), amt, txHash, uint64(chainId), uint64(types.LPHistoryStatus_LP_SUBMITTING), uint64(lpType), seqNum)
	}
	if err == nil {
		return &webapi.MarkLiquidityResponse{}, nil
	} else {
		return &webapi.MarkLiquidityResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "update data err",
			},
		}, nil
	}
}

func (gs *GatewayService) WithdrawLiquidity(ctx context.Context, request *webapi.WithdrawLiquidityRequest) (*webapi.WithdrawLiquidityResponse, error) {
	log.Debugf("WithdrawLiquidity req:%+v", request)
	transferId := request.GetTransferId()
	tr := gs.TP.GetTransactor()
	if transferId != "" {
		// refund transfer
		transfer, tFound, err := dal.DB.GetTransfer(transferId)
		if !tFound || err != nil {
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "transfer not found",
				},
			}, nil
		}
		seqNum := request.Reqid
		receiver := common.Hex2Addr(transfer.UsrAddr).Bytes()
		if transfer.RefundSeqNum > 0 {
			// for sign again test only, not normal case
			log.Debugf("signAgain for transfer:%s, seqNum:%d", transferId, transfer.RefundSeqNum)
			seqNum, err = gs.signAgainWithdraw(&types.MsgSignAgain{
				Creator:  tr.Key.GetAddress().String(),
				ReqId:    seqNum,
				UserAddr: receiver,
			})
		} else {
			err = dal.DB.MarkTransferRequestingRefund(transferId, seqNum)
			if err != nil {
				return &webapi.WithdrawLiquidityResponse{
					Err: &webapi.ErrMsg{
						Code: webapi.ErrCode_ERROR_CODE_COMMON,
						Msg:  "db error when mark refund",
					},
				}, nil
			}
			seqNum, err = gs.initWithdraw(&types.MsgInitWithdraw{
				XferId:  common.Hex2Bytes(transferId),
				Creator: tr.Key.GetAddress().String(),
				ReqId:   request.Reqid,
				UserSig: request.Sig,
				LpAddr:  receiver,
			})
		}

		if err != nil {
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  err.Error(),
				},
			}, nil
		}
		return &webapi.WithdrawLiquidityResponse{
			SeqNum: seqNum,
		}, nil
	} else {
		// remove liquidity
		amt := request.GetAmount()
		chainId := request.GetChainId()
		tokenAddr := common.Hex2Addr(request.GetTokenAddr()).String()
		token, found, err := dal.DB.GetTokenByAddr(tokenAddr, uint64(chainId))
		if !found || err != nil {
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "token not found in gateway DB",
				},
			}, nil
		}
		lp := common.Hex2Addr(request.GetReceiverAddr()).String()
		seqNum := request.Reqid
		if dal.DB.HasSeqNumUsedForWithdraw(seqNum, lp) {
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "invalid seq num, it has been used for current lp",
				},
			}, nil
		}
		err = dal.DB.UpsertLPWithSeqNum(lp, token.Token.Symbol, token.Token.Address, amt, strconv.Itoa(int(seqNum)), uint64(chainId), uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN), uint64(webapi.LPType_LP_TYPE_REMOVE), seqNum)
		if err != nil {
			_ = dal.DB.UpdateLPStatusForWithdraw(uint64(chainId), seqNum, uint64(types.LPHistoryStatus_LP_FAILED), lp)
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "db error when mark withdraw",
				},
			}, nil
		}
		seqNum, err = gs.initWithdraw(&types.MsgInitWithdraw{
			Chainid: uint64(chainId),
			LpAddr:  common.Hex2Bytes(lp),
			Token:   common.Hex2Bytes(tokenAddr),
			Amount:  common.Str2BigInt(amt).Bytes(),
			Creator: tr.Key.GetAddress().String(),
			ReqId:   seqNum,
			UserSig: request.Sig,
		})
		if err != nil {
			_ = dal.DB.UpdateLPStatusForWithdraw(uint64(chainId), seqNum, uint64(types.LPHistoryStatus_LP_FAILED), lp)
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  err.Error(),
				},
			}, nil
		}
		return &webapi.WithdrawLiquidityResponse{
			SeqNum: seqNum,
		}, nil
	}
}

func (gs *GatewayService) QueryLiquidityStatus(ctx context.Context, request *webapi.QueryLiquidityStatusRequest) (*webapi.QueryLiquidityStatusResponse, error) {
	seqNum := request.GetSeqNum()
	chainId := uint64(request.GetChainId())
	lpType := uint64(request.GetType())
	addr := common.Hex2Addr(request.GetLpAddr())
	tr := gs.TP.GetTransactor()
	txHash, status, lpUpdateTime, found, err := dal.DB.GetLPInfo(seqNum, lpType, chainId, addr.String())
	if found && err == nil && status == uint64(types.LPHistoryStatus_LP_SUBMITTING) && common.IsValidTxHash(txHash) {
		ec := gs.EC[chainId]
		if ec == nil {
			log.Errorf("no ethClient found for chain:%d", chainId)
			return nil, fmt.Errorf("no ethClient found for chain:%d", chainId)
		}

		receipt, recErr := ec.TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(txHash)))
		if recErr == nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
			log.Warnf("find transfer failed, chain_id %d, hash:%s", chainId, txHash)
			if lpType == uint64(webapi.LPType_LP_TYPE_ADD) {
				dbErr := dal.DB.UpdateLPStatus(seqNum, lpType, chainId, addr.String(), uint64(types.LPHistoryStatus_LP_FAILED))
				if dbErr != nil {
					log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", chainId, txHash)
				} else {
					status = uint64(types.LPHistoryStatus_LP_FAILED)
				}
			} else if lpType == uint64(webapi.LPType_LP_TYPE_REMOVE) {
				dbErr := dal.DB.UpdateLPStatus(seqNum, lpType, chainId, addr.String(), uint64(types.LPHistoryStatus_LP_WAITING_FOR_LP))
				if dbErr != nil {
					log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", chainId, txHash)
				} else {
					status = uint64(types.LPHistoryStatus_LP_WAITING_FOR_LP)
				}
			}
		}
	}

	link := ""
	if common.IsValidTxHash(txHash) {
		_, chainUrl, chainFound, chainErr := dal.DB.GetChain(chainId)
		if chainFound && chainErr == nil && chainUrl != "" {
			link = chainUrl + txHash
		}
	}

	if found && lpType == uint64(webapi.LPType_LP_TYPE_ADD) { // add type
		if status == uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN) {
			resp, err2 := cbrcli.QueryAddLiquidityStatus(tr.CliCtx, &types.QueryAddLiquidityStatusRequest{
				ChainId: chainId,
				SeqNum:  seqNum,
			})
			if resp != nil && err2 == nil {
				_ = dal.DB.UpdateLPStatus(seqNum, lpType, chainId, addr.String(), uint64(resp.Status))
				status = uint64(resp.Status)
			}
		}

		return &webapi.QueryLiquidityStatusResponse{
			Status:      types.LPHistoryStatus(status),
			WdOnchain:   nil,
			Signers:     nil,
			SortedSigs:  nil,
			BlockTxLink: link,
		}, nil
	} else if found && lpType == uint64(webapi.LPType_LP_TYPE_REMOVE) { // withdraw type
		resp := &webapi.QueryLiquidityStatusResponse{
			Status:      types.LPHistoryStatus(status),
			WdOnchain:   nil,
			Signers:     nil,
			SortedSigs:  nil,
			BlockTxLink: link,
		}

		if status == uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN) || status == uint64(types.LPHistoryStatus_LP_WAITING_FOR_LP) {
			if status == uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN) && time.Now().Add(-15*time.Minute).After(lpUpdateTime) {
				seqNum, err = gs.signAgainWithdraw(&types.MsgSignAgain{
					Creator:  tr.Key.GetAddress().String(),
					ReqId:    seqNum,
					UserAddr: addr.Bytes(),
				})
				if err != nil {
					// sign again failed, we will mark this tx as WAITING_FOR_SGN again and will check again after 15 min
					_ = dal.DB.UpdateLPStatusForWithdraw(chainId, seqNum, uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN), addr.String())
				}
			} else {
				detail, wdOnchain, sortedSigs, signers, powers := gs.getWithdrawInfo(seqNum, chainId, addr.String())
				resp.WdOnchain = wdOnchain
				resp.SortedSigs = sortedSigs
				resp.Signers = signers
				resp.Powers = powers
				if detail != nil && status == uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN) && detail.GetStatus() != resp.Status {
					dberr := dal.DB.UpdateLPStatusForWithdraw(chainId, seqNum, uint64(detail.Status), addr.String())
					if dberr != nil {
						log.Errorf("db error:%+v", dberr)
					}
					resp.Status = detail.GetStatus()
				}
			}
		}
		return resp, nil
	}

	return &webapi.QueryLiquidityStatusResponse{
		Status:      types.LPHistoryStatus(status),
		WdOnchain:   nil,
		Signers:     nil,
		SortedSigs:  nil,
		BlockTxLink: link,
	}, nil
}

func (gs *GatewayService) LPHistory(ctx context.Context, request *webapi.LPHistoryRequest) (*webapi.LPHistoryResponse, error) {
	addr := common.Hex2Addr(request.GetAddr()).String()
	endTime := time.Now()
	if request.GetNextPageToken() != "" {
		ts, err := strconv.Atoi(request.GetNextPageToken())
		if err != nil {
			return &webapi.LPHistoryResponse{}, nil
		}
		endTime = common.TsToTime(uint64(ts))
	}
	lpHistory, currentPageSize, next, err := dal.DB.PaginateLpHistory(addr, endTime, request.GetPageSize())
	if err != nil {
		log.Error("db error", err)
		return &webapi.LPHistoryResponse{}, nil
	}
	gs.updateLpStatusInHistory(lpHistory)
	var lps []*webapi.LPHistory
	for _, lp := range lpHistory {
		chain, chainUrl, found, lpErr := dal.DB.GetChain(lp.ChainId)
		if lpErr != nil {
			log.Errorf("chain not found: %d", lp.ChainId)
			continue
		}
		if !found {
			chain = unknownChain(uint32(lp.ChainId))
		}
		token, found, lpErr := dal.DB.GetTokenBySymbol(lp.TokenSymbol, lp.ChainId)
		if !found || lpErr != nil {
			log.Errorf("token not found for token: %s, on chain: %d", lp.TokenSymbol, lp.ChainId)
			continue
		}
		txLink := ""
		if common.IsValidTxHash(lp.TxHash) {
			txLink = chainUrl + lp.TxHash
		}

		lps = append(lps, &webapi.LPHistory{
			Chain:       chain,
			Token:       token,
			Amount:      lp.Amt,
			Ts:          common.TsMilli(lp.Ct),
			BlockTxLink: txLink,
			Status:      lp.Status,
			Type:        lp.LpType,
			SeqNum:      lp.SeqNum,
		})
	}
	return &webapi.LPHistoryResponse{
		History:       lps,
		NextPageToken: strconv.FormatUint(common.TsMilli(next), 10),
		CurrentSize:   uint64(currentPageSize),
	}, nil
}

// ================================= internal method below =====================================

func (gs *GatewayService) initWithdraw(req *types.MsgInitWithdraw) (uint64, error) {
	tr := gs.TP.GetTransactor()
	log.Debugf("init withdraw, req:%+v", req)
	err := checkSig(req.GetReqId(), req.GetUserSig(), common.Bytes2Addr(req.GetLpAddr()))
	if err != nil {
		log.Errorf("checkSig err:%+v", err)
		return 0, err
	}
	_, err = cbrcli.InitWithdraw(tr, req)
	return req.ReqId, err
}

func (gs *GatewayService) signAgainWithdraw(req *types.MsgSignAgain) (uint64, error) {
	tr := gs.TP.GetTransactor()
	log.Debugf("sign again, req:%+v", req)
	_, err := cbrcli.SignAgain(tr, req)
	return req.ReqId, err
}

func (gs *GatewayService) getWithdrawInfo(seqNum, chainId uint64, usrAddr string) (*types.QueryLiquidityStatusResponse, []byte, [][]byte, [][]byte, [][]byte) {
	tr := gs.TP.GetTransactor()
	detail, err2 := cbrcli.QueryWithdrawLiquidityStatus(tr.CliCtx, &types.QueryWithdrawLiquidityStatusRequest{
		SeqNum:  seqNum,
		UsrAddr: usrAddr,
	})
	var wdOnchain []byte
	var signers [][]byte
	var powers [][]byte
	var sortedSigs [][]byte
	if detail != nil && err2 == nil {
		wdOnchain = detail.GetDetail().GetWdOnchain()
		sortedSigs = detail.GetDetail().GetSortedSigsBytes()
		curss, signErr := cbrcli.QueryChainSigners(tr.CliCtx, chainId)
		if signErr != nil {
			log.Warnf("QueryChainSigners error:%+v", signErr)
		} else {
			ss, ps := types.SignersToEthArrays(curss.GetSortedSigners())
			for i, s := range ss {
				signers = append(signers, s.Bytes())
				powers = append(powers, ps[i].Bytes())
			}
		}
	} else {
		log.Warnf("QueryWithdrawLiquidityStatus error for detail, error%+v", err2)
	}
	return detail, wdOnchain, sortedSigs, signers, powers
}

func (gs *GatewayService) updateLpStatusInHistory(lpHistory []*dal.LP) {
	for _, lp := range lpHistory {
		if lp.Status == types.LPHistoryStatus_LP_SUBMITTING || lp.Status == types.LPHistoryStatus_LP_WAITING_FOR_SGN {
			resp, err := gs.QueryLiquidityStatus(context.Background(), &webapi.QueryLiquidityStatusRequest{
				SeqNum:  lp.SeqNum,
				LpAddr:  lp.Addr,
				ChainId: uint32(lp.ChainId),
				Type:    lp.LpType,
			})
			if err != nil {
				log.Warn("updateLpStatusInHistory error", err)
				continue
			}
			lp.Status = resp.GetStatus()
		}
	}
}

func checkSig(reqId uint64, sig []byte, addr common.Addr) error {
	signAddr, err := ethutils.RecoverSigner(eth.ToPadBytes(reqId), sig)
	if err != nil {
		return err
	}
	if signAddr != addr {
		return fmt.Errorf("error sig addr, sigAddr:%s, usrAddr:%s", signAddr.String(), addr.String())
	}
	return nil
}