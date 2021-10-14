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
	"sort"
	"strconv"
	"time"
)

type txData struct {
	volume   float64
	fee      *big.Int
	dstToken *types.Token
}

func (gs *GatewayService) GetLPInfoList(ctx context.Context, request *webapi.GetLPInfoListRequest) (*webapi.GetLPInfoListResponse, error) {
	userAddr := common.Hex2Addr(request.GetAddr()).String()
	chainTokenInfos, err := dal.DB.GetChainTokenList()
	if err != nil || len(chainTokenInfos) == 0 {
		return &webapi.GetLPInfoListResponse{}, nil
	}
	var chainTokens []*types.ChainTokenAddrPair
	for chainId, tokens := range chainTokenInfos {
		for _, tokenInfo := range tokens.Token {
			chainTokens = append(chainTokens, &types.ChainTokenAddrPair{
				ChainId:   uint64(chainId),
				TokenAddr: tokenInfo.GetToken().Address,
			})
		}
	}

	var lps []*webapi.LPInfo
	tr := gs.TP.GetTransactor()
	detailList, err := cbrcli.QueryLiquidityDetailList(tr.CliCtx, &types.LiquidityDetailListRequest{
		LpAddr:     userAddr,
		ChainToken: chainTokens,
	})

	if err != nil || detailList == nil || len(detailList.GetLiquidityDetail()) == 0 {
		return &webapi.GetLPInfoListResponse{}, nil
	}
	farmingApyMap := gs.getFarmingApy(ctx)
	data24h := get24hTx()
	userDetailMap := make(map[uint64]map[string]*types.LiquidityDetail)
	for _, detail := range detailList.GetLiquidityDetail() {
		chainId := detail.GetChainId()
		tokenWithAddr := detail.GetToken() // only has addr field
		token, found, dbErr := dal.DB.GetTokenByAddr(common.Hex2Addr(tokenWithAddr.GetAddress()).String(), chainId)
		if !found || dbErr != nil {
			log.Debugf("data, token not found in lp list, token addr:%s, chainId:%d", tokenWithAddr.GetAddress(), chainId)
			continue
		}
		detail.Token = token.Token
		chainInfo, found := userDetailMap[chainId]
		if !found {
			chainInfo = make(map[string]*types.LiquidityDetail)
		}
		chainInfo[token.Token.Symbol] = detail
		userDetailMap[chainId] = chainInfo
	}
	for chainId32, chainToken := range chainTokenInfos {
		chainId := uint64(chainId32)
		for _, token := range chainToken.Token {
			tokenSymbol := token.Token.Symbol
			totalLiquidity := "0"
			usrLpFeeEarning := "0"
			usrLiquidity := "0"
			detail, found := userDetailMap[chainId][tokenSymbol]
			if found {
				totalLiquidity = detail.GetTotalLiquidity()
				usrLpFeeEarning = detail.GetUsrLpFeeEarning()
				usrLiquidity = detail.GetUsrLiquidity()
			}

			enrichUnknownToken(token)
			chain, _, found, dbErr := dal.DB.GetChain(chainId)
			if !found || dbErr != nil {
				chain = unknownChain(chainId32)
			}

			data := data24h[chainId][tokenSymbol]
			lpFeeEarningApy := 0.0
			volume24h := 0.0
			if data != nil {
				if common.Str2BigInt(totalLiquidity).Cmp(new(big.Int).SetInt64(0)) > 0 {
					lpFeeEarningApy, _ = new(big.Float).Quo(new(big.Float).SetInt(data.fee), new(big.Float).SetInt(common.Str2BigInt(totalLiquidity))).Float64()
				}
				volume24h = data.volume
			}
			farmingApy, hasSession := farmingApyMap[chainId][token.Token.GetSymbol()]
			lp := &webapi.LPInfo{
				Chain:              chain,
				Token:              token,
				Liquidity:          gs.F.GetUsdVolume(token.Token, common.Str2BigInt(usrLiquidity)),
				LiquidityAmt:       usrLiquidity,
				HasFarmingSessions: hasSession,
				LpFeeEarning:       gs.F.GetUsdVolume(token.Token, common.Str2BigInt(usrLpFeeEarning)),
				Volume_24H:         volume24h,
				TotalLiquidity:     gs.F.GetUsdVolume(token.Token, common.Str2BigInt(totalLiquidity)),
				TotalLiquidityAmt:  totalLiquidity,
				LpFeeEarningApy:    lpFeeEarningApy,
				FarmingApy:         farmingApy,
			}
			lps = append(lps, lp)
		}
	}
	sort.SliceStable(lps, func(i, j int) bool {
		if lps[i].HasFarmingSessions {
			if lps[j].HasFarmingSessions {
				return lps[i].GetVolume_24H() < lps[j].GetVolume_24H()
			} else {
				return false
			}
		} else {
			if lps[j].HasFarmingSessions {
				return true
			} else {
				return lps[i].GetVolume_24H() < lps[j].GetVolume_24H()
			}
		}
	})
	return &webapi.GetLPInfoListResponse{
		LpInfo: lps,
	}, nil
}

func (gs *GatewayService) MarkLiquidity(ctx context.Context, request *webapi.MarkLiquidityRequest) (*webapi.MarkLiquidityResponse, error) {
	lpType := request.GetType()
	chainId := request.GetChainId()
	amt := request.GetAmt()
	addr := common.Hex2Addr(request.GetLpAddr()).String()
	seqNum := request.GetSeqNum()
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
	err = dal.DB.UpsertLP(addr, token.GetToken().GetSymbol(), token.GetToken().GetAddress(), amt, txHash, uint64(chainId), uint64(types.LPHistoryStatus_LP_SUBMITTING), uint64(lpType), seqNum)
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
		err = dal.DB.UpsertLP(lp, token.Token.Symbol, token.Token.Address, amt, "", uint64(chainId), uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN), uint64(webapi.LPType_LP_TYPE_REMOVE), seqNum)
		if err != nil {
			_ = dal.DB.UpdateLPStatusForWithdraw(seqNum, uint64(types.LPHistoryStatus_LP_FAILED))
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "db error when mark refund",
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
			_ = dal.DB.UpdateLPStatusForWithdraw(seqNum, uint64(types.LPHistoryStatus_LP_FAILED))
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
	if found && err == nil && status == uint64(types.LPHistoryStatus_LP_SUBMITTING) && txHash != "" {
		ec := gs.EC[chainId]
		if ec == nil {
			log.Errorf("no ethClient found for chain:%d", chainId)
			return nil, fmt.Errorf("no ethClient found for chain:%d", chainId)
		}

		receipt, recErr := ec.TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(txHash)))
		if recErr == nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
			log.Warnf("find transfer failed, chain_id %d, hash:%s", chainId, txHash)
			dbErr := dal.DB.UpdateLPStatus(seqNum, lpType, chainId, addr.String(), uint64(types.LPHistoryStatus_LP_FAILED))
			if dbErr != nil {
				log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", chainId, txHash)
			} else {
				status = uint64(types.LPHistoryStatus_LP_FAILED)
			}
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
				return &webapi.QueryLiquidityStatusResponse{
					Status:     resp.Status,
					WdOnchain:  nil,
					Signers:    nil,
					SortedSigs: nil,
				}, nil
			}
		}
	} else if found && lpType == uint64(webapi.LPType_LP_TYPE_REMOVE) { // withdraw type
		resp := &webapi.QueryLiquidityStatusResponse{
			Status:     types.LPHistoryStatus(status),
			WdOnchain:  nil,
			Signers:    nil,
			SortedSigs: nil,
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
					_ = dal.DB.UpdateLPStatusForWithdraw(seqNum, uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN))
				}
			} else {
				detail, wdOnchain, sortedSigs, signers, powers := gs.getWithdrawInfo(seqNum, chainId, addr.String())
				resp.WdOnchain = wdOnchain
				resp.SortedSigs = sortedSigs
				resp.Signers = signers
				resp.Powers = powers
				if detail != nil && status == uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN) && detail.GetStatus() != resp.Status {
					_ = dal.DB.UpdateLPStatusForWithdraw(seqNum, uint64(detail.Status))
					resp.Status = detail.GetStatus()
				}
			}
		}
		return resp, nil
	}

	return &webapi.QueryLiquidityStatusResponse{
		Status:     types.LPHistoryStatus(status),
		WdOnchain:  nil,
		Signers:    nil,
		SortedSigs: nil,
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
		if lp.TxHash != "" {
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

// todo cache this @aric
func get24hTx() map[uint64]map[string]*txData {
	txs, err := dal.DB.Get24hTx()
	resp := make(map[uint64]map[string]*txData) // map<chain_id, map<token_symbol, txData>>
	if err == nil {
		for _, tx := range txs {
			tokenSymbol := tx.TokenSymbol
			dstToken, found, dbErr := dal.DB.GetTokenBySymbol(tokenSymbol, tx.DstChainId)
			if !found || dbErr != nil {
				continue
			}
			dstChainId := tx.DstChainId
			data, found := resp[dstChainId]
			if !found || data == nil {
				data = make(map[string]*txData)
			}
			d, found := data[tokenSymbol]
			if !found || d == nil {
				d = &txData{
					volume:   0,
					fee:      new(big.Int),
					dstToken: dstToken.Token,
				}
			}
			d.fee = new(big.Int).Add(d.fee, common.Str2BigInt(tx.DstAmt))
			d.volume += tx.Volume
			data[tokenSymbol] = d
			resp[tx.DstChainId] = data
		}
	}
	return resp
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
