package gateway

import (
	"context"
	"fmt"
	"math/big"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingkp "github.com/celer-network/sgn-v2/x/farming/keeper"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
)

var (
	selfStart         bool
	rootDir           string
	legacyAmino       *codec.LegacyAmino
	cdc               codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry
)

// Close the database DAL.
func (gs *GatewayService) Close() {
	if dal.DB == nil {
		return
	}
	dal.DB.Close()
	dal.DB = nil
}

type GatewayConfig struct {
}

type GatewayService struct {
	f  *fee.TokenPriceCache
	tp *transactor.TransactorPool
	ec map[uint64]*ethclient.Client
}

func (gs *GatewayService) GetTransferStatus(ctx context.Context, request *webapi.GetTransferStatusRequest) (*webapi.GetTransferStatusResponse, error) {
	transfer, found, err := dal.DB.GetTransfer(request.GetTransferId())
	var detail *types.QueryLiquidityStatusResponse
	var wdOnchain []byte
	var sortedSigs [][]byte
	var signers [][]byte
	var powers [][]byte
	if found && err == nil {
		var transfers []*dal.Transfer
		transfers = append(transfers, transfer)
		err = gs.updateTransferStatusInHistory(ctx, transfers)
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
				detail, wdOnchain, sortedSigs, signers, powers = gs.getWithdrawInfo(transfer.RefundSeqNum, transfer.SrcChainId)
				if detail == nil {
					return &webapi.GetTransferStatusResponse{
						Err: &webapi.ErrMsg{
							Code: webapi.ErrCode_ERROR_CODE_COMMON,
							Msg:  "withdrawInfo not found",
						},
					}, nil
				}
				log.Debugf("get lp info for transfer, status is :%+v", detail.GetStatus())
				if detail.GetStatus() == types.LPHistoryStatus_LP_WAITING_FOR_LP && transfer.Status != types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED {
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
	}
	if !found || err != nil {
		return &webapi.GetTransferStatusResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "transfer not found",
			},
		}, nil
	}
	return &webapi.GetTransferStatusResponse{
		Status:     transfer.Status,
		WdOnchain:  wdOnchain,
		SortedSigs: sortedSigs,
		Signers:    signers,
		Powers:     powers,
	}, nil
}

func (gs *GatewayService) SetAdvancedInfo(ctx context.Context, request *webapi.SetAdvancedInfoRequest) (*webapi.SetAdvancedInfoResponse, error) {
	addr := common.Hex2Addr(request.GetAddr()).String()
	err := dal.DB.UpsertSlippageSetting(addr, request.GetSlippageTolerance())
	if err == nil {
		return &webapi.SetAdvancedInfoResponse{}, nil
	} else {
		return &webapi.SetAdvancedInfoResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "update setting failed",
			},
		}, nil
	}
}

func (gs *GatewayService) GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error) {
	chainTokenList, err := dal.DB.GetChainTokenList()
	if err != nil {
		return &webapi.GetTransferConfigsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get chain_token failed",
			},
			Chains:     nil,
			ChainToken: nil,
		}, nil
	}
	var chainIds []uint32
	for key := range chainTokenList {
		chainIds = append(chainIds, key)
	}
	chains, err := dal.DB.GetChainInfo(chainIds)
	chainFound := make(map[uint32]bool)
	for _, chain := range chains {
		chainFound[chain.Id] = true
	}
	for chainId, tokens := range chainTokenList {
		_, found := chainFound[chainId]
		if !found {
			chains = append(chains, unknownChain(chainId))
		}
		for _, token := range tokens.Token {
			enrichUnknownToken(token)
		}
	}
	if err != nil {
		return &webapi.GetTransferConfigsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get chain info failed",
			},
			Chains:     chains,
			ChainToken: chainTokenList,
		}, nil
	}
	return &webapi.GetTransferConfigsResponse{
		Err:        nil,
		Chains:     chains,
		ChainToken: chainTokenList,
	}, nil
}

func (gs *GatewayService) EstimateAmt(ctx context.Context, request *webapi.EstimateAmtRequest) (*webapi.EstimateAmtResponse, error) {
	amt := request.GetAmt()
	srcChainId := request.GetSrcChainId()
	dstChainId := request.GetDstChainId()
	tokenSymbol := request.GetTokenSymbol()
	srcToken, found1, err1 := dal.DB.GetTokenBySymbol(tokenSymbol, uint64(srcChainId))
	if err1 != nil || !found1 {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "token not found",
			},
		}, nil
	}
	dstToken, found2, err2 := dal.DB.GetTokenBySymbol(tokenSymbol, uint64(dstChainId))
	if err2 != nil || !found2 {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_NO_TOKEN_ON_DST_CHAIN,
				Msg:  "token not support on dst chain",
			},
		}, nil
	}

	addr := common.Hex2Addr(request.GetUsrAddr()).String()
	slippage, found, err := dal.DB.GetSlippageSetting(addr)
	if err != nil || !found {
		slippage = 5000
	}
	tr := gs.tp.GetTransactor()
	feeInfo, err := cbrcli.QueryFee(tr.CliCtx, &types.GetFeeRequest{
		SrcChainId:   uint64(srcChainId),
		DstChainId:   uint64(dstChainId),
		SrcTokenAddr: srcToken.Token.GetAddress(),
		Amt:          amt,
	})
	if err != nil {
		log.Warnf("cli.QueryFee error:%+v", err)
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  err.Error(),
			},
		}, nil
	}
	if feeInfo == nil {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "can not estimate fee",
			},
		}, nil
	}
	eqValueTokenAmt := feeInfo.GetEqValueTokenAmt()
	feeAmt := feeInfo.GetFee()

	srcVolume := gs.f.GetUsdVolume(srcToken.Token, common.Str2BigInt(amt))
	dstVolume := gs.f.GetUsdVolume(dstToken.Token, common.Str2BigInt(eqValueTokenAmt))
	bridgeRate := 0.0
	if srcVolume > 0.000000001 {
		bridgeRate = dstVolume / srcVolume
	} else {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "amount should > 0",
			},
		}, nil
	}
	minReceiveVolume := dstVolume*(1-float64(slippage)/1e6) - gs.f.GetUsdVolume(dstToken.Token, common.Str2BigInt(feeAmt))
	return &webapi.EstimateAmtResponse{
		EqValueTokenAmt:   eqValueTokenAmt,
		BridgeRate:        float32(bridgeRate),
		Fee:               feeAmt,
		SlippageTolerance: slippage,
		MaxSlippage:       uint32((srcVolume - minReceiveVolume) * 1e6 / srcVolume),
	}, nil
}

func (gs *GatewayService) MarkTransfer(ctx context.Context, request *webapi.MarkTransferRequest) (*webapi.MarkTransferResponse, error) {
	transferId := request.GetTransferId()
	addr := common.Hex2Addr(request.GetAddr())
	sendInfo := request.GetSrcSendInfo()
	receivedInfo := request.GetDstMinReceivedInfo()
	txHash := request.GetSrcTxHash()
	txType := request.GetType()
	log.Infof("transferId in mark api: %s, bytes:%+v, request: %+v", transferId, common.Hex2Bytes(transferId), request)
	if txType == webapi.TransferType_TRANSFER_TYPE_SEND {
		err := dal.DB.MarkTransferSend(transferId, addr.String(), sendInfo.GetToken().GetSymbol(),
			sendInfo.GetAmount(), receivedInfo.GetAmount(), txHash, uint64(sendInfo.GetChain().GetId()),
			uint64(receivedInfo.GetChain().GetId()), gs.f.GetUsdVolume(sendInfo.GetToken(), common.Str2BigInt(sendInfo.GetAmount())))
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
	tr := gs.tp.GetTransactor()
	detailList, err := cbrcli.QueryLiquidityDetailList(tr.CliCtx, &types.LiquidityDetailListRequest{
		LpAddr:     userAddr,
		ChainToken: chainTokens,
	})

	if err != nil || detailList == nil || len(detailList.GetLiquidityDetail()) == 0 {
		return &webapi.GetLPInfoListResponse{}, nil
	}
	stakingMap := gs.getUserStaking(ctx, userAddr)
	farmingEarningMap := gs.getUserFarmingCumulativeEarning(ctx, userAddr)
	farmingApyMap := gs.getFarmingApy(ctx)
	data24h := gs.get24hTx()
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
			staking := stakingMap[chainId][token.Token.GetSymbol()]
			hasSession := false
			currentEarning := 0.0
			if staking != nil {
				hasSession = staking.staking > 0
				currentEarning = staking.earning
			}
			lp := &webapi.LPInfo{
				Chain:                chain,
				Token:                token,
				Liquidity:            gs.f.GetUsdVolume(token.Token, common.Str2BigInt(usrLiquidity)),
				HasFarmingSessions:   hasSession,
				LpFeeEarning:         gs.f.GetUsdVolume(token.Token, common.Str2BigInt(usrLpFeeEarning)),
				FarmingRewardEarning: farmingEarningMap[chainId][token.Token.GetSymbol()] + currentEarning,
				Volume_24H:           volume24h,
				TotalLiquidity:       gs.f.GetUsdVolume(token.Token, common.Str2BigInt(totalLiquidity)),
				LpFeeEarningApy:      lpFeeEarningApy,
				FarmingApy:           farmingApyMap[chainId][token.Token.GetSymbol()],
			}
			lps = append(lps, lp)
		}
	}
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
	transferId := request.GetTransferId()
	tr := gs.tp.GetTransactor()
	if transferId != "" {
		// refund transfer
		seqNum, err := gs.initWithdraw(&types.MsgInitWithdraw{
			XferId:  common.Hex2Bytes(transferId),
			Creator: tr.Key.GetAddress().String(),
		})
		if err != nil {
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  err.Error(),
				},
			}, nil
		}
		err = dal.DB.MarkTransferRequestingRefund(transferId, seqNum)
		if err != nil {
			return &webapi.WithdrawLiquidityResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "db error when mark refund",
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
		seqNum, err := gs.initWithdraw(&types.MsgInitWithdraw{
			Chainid: uint64(chainId),
			LpAddr:  common.Hex2Bytes(lp),
			Token:   common.Hex2Bytes(tokenAddr),
			Amount:  common.Str2BigInt(amt).Bytes(),
			Creator: tr.Key.GetAddress().String(),
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
		return &webapi.WithdrawLiquidityResponse{
			SeqNum: seqNum,
		}, nil
	}
}

func (gs *GatewayService) initWithdraw(req *types.MsgInitWithdraw) (uint64, error) {
	tr := gs.tp.GetTransactor()
	resp, err := cbrcli.InitWithdraw(tr, req)
	if resp == nil {
		return 0, err
	}
	return resp.GetSeqnum(), err
}

func (gs *GatewayService) QueryLiquidityStatus(ctx context.Context, request *webapi.QueryLiquidityStatusRequest) (*webapi.QueryLiquidityStatusResponse, error) {
	seqNum := request.GetSeqNum()
	chainId := uint64(request.GetChainId())
	lpType := uint64(request.GetType())
	addr := request.GetLpAddr()
	tr := gs.tp.GetTransactor()
	txHash, status, found, err := dal.DB.GetLPInfo(seqNum, lpType, chainId, addr)
	if found && err == nil && status == uint64(types.LPHistoryStatus_LP_SUBMITTING) && txHash != "" {
		ec := gs.ec[chainId]
		if ec == nil {
			log.Errorf("no ethClient found for chain:%d", chainId)
			return nil, fmt.Errorf("no ethClient found for chain:%d", chainId)
		}

		receipt, recErr := ec.TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(txHash)))
		if recErr == nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
			log.Warnf("find transfer failed, chain_id %d, hash:%s", chainId, txHash)
			dbErr := dal.DB.UpdateLPStatus(seqNum, lpType, chainId, addr, uint64(types.LPHistoryStatus_LP_FAILED))
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
				_ = dal.DB.UpdateLPStatus(seqNum, lpType, chainId, addr, uint64(resp.Status))
				return &webapi.QueryLiquidityStatusResponse{
					Status:     types.LPHistoryStatus(status),
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
			detail, wdOnchain, sortedSigs, signers, powers := gs.getWithdrawInfo(seqNum, chainId)
			resp.WdOnchain = wdOnchain
			resp.SortedSigs = sortedSigs
			resp.Signers = signers
			resp.Powers = powers
			if status == uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN) && detail.GetStatus() != resp.Status {
				_ = dal.DB.UpdateLPStatusForWithdraw(seqNum, uint64(detail.Status))
				resp.Status = detail.GetStatus()
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

func (gs *GatewayService) getWithdrawInfo(seqNum, chainId uint64) (*types.QueryLiquidityStatusResponse, []byte, [][]byte, [][]byte, [][]byte) {
	tr := gs.tp.GetTransactor()
	detail, err2 := cbrcli.QueryWithdrawLiquidityStatus(tr.CliCtx, &types.QueryWithdrawLiquidityStatusRequest{
		SeqNum: seqNum,
	})
	var wdOnchain []byte
	var signers [][]byte
	var powers [][]byte
	var sortedSigs [][]byte
	if detail != nil && err2 == nil {
		wdOnchain = detail.GetDetail().GetWdOnchain()
	} else {
		log.Warnf("QueryWithdrawLiquidityStatus error for detail, error%+v", err2)
	}
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
	return detail, wdOnchain, sortedSigs, signers, powers
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
	err = gs.updateTransferStatusInHistory(ctx, transferList)
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
			srcChain = &webapi.Chain{
				Id:   uint32(transfer.SrcChainId),
				Name: "UNKNOWN NAME",
				Icon: "",
			}
		}
		if !dstFound {
			dstChain = &webapi.Chain{
				Id:   uint32(transfer.DstChainId),
				Name: "UNKNOWN NAME",
				Icon: "",
			}
		}
		srcToken, srcFound, err1 := dal.DB.GetTokenBySymbol(transfer.TokenSymbol, transfer.SrcChainId)
		dstToken, dstFound, err2 := dal.DB.GetTokenBySymbol(transfer.TokenSymbol, transfer.DstChainId)
		if !srcFound || !dstFound || err1 != nil || err2 != nil {
			continue
		}
		srcTxLink := ""
		dstTxLink := ""
		if transfer.SrcTxHash != "" {
			srcTxLink = srcChainUrl + transfer.SrcTxHash
		}

		if transfer.DstTxHash != "" {
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
		})
	}
	return &webapi.TransferHistoryResponse{
		History:       transfers,
		NextPageToken: strconv.FormatUint(common.TsMilli(next), 10),
		CurrentSize:   uint64(currentPageSize),
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
			chain = &webapi.Chain{
				Id:   uint32(lp.ChainId),
				Name: "UNKNOWN NAME",
				Icon: "",
			}
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

func NewGatewayService(dbUrl string) (*GatewayService, error) {
	if selfStart {
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
		config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
		config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
		config.Seal()
	}
	// Make a private config copy.
	_db, err := dal.NewDAL("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", dbUrl), 10)
	if err != nil {
		return nil, err
	}

	dal.DB = _db
	gateway := &GatewayService{}

	return gateway, nil
}

// StartTokenPricePolling starts a loop with the given interval and 3s stdev for polling price
func (gs *GatewayService) StartChainTokenPolling(interval time.Duration) {
	gs.pollChainToken() // make sure run at least once before return
	polledInside := false
	go func() {
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			if polledInside {
				gs.pollChainToken()
			}
			polledInside = true
		}
	}()
}

func (gs *GatewayService) pollChainToken() {
	tr := gs.tp.GetTransactor()
	resp, err := cbrcli.QueryChainTokensConfig(tr.CliCtx, &types.ChainTokensConfigRequest{})
	if err != nil {
		log.Errorln("we will use mocked chain tokens failed to load basic token info:", err)
	}
	chainTokens := resp.GetChainTokens()
	for chainIdStr, assets := range chainTokens {
		chainId, convErr := strconv.Atoi(chainIdStr)
		if convErr != nil {
			log.Errorf("error chain id found:%s", chainIdStr)
			continue
		}
		for _, asset := range assets.Assets {
			token := asset.GetToken()
			dbErr := dal.DB.UpsertTokenBaseInfo(token.GetSymbol(), common.Hex2Addr(token.GetAddress()).String(), common.Hex2Addr(asset.GetContractAddr()).String(), uint64(chainId), uint64(token.GetDecimal()))
			if dbErr != nil {
				log.Errorf("failed to write token: %v", err)
			}
		}
	}
}

func (gs *GatewayService) updateLpStatusInHistory(lpHistory []*dal.LP) {
	for _, lp := range lpHistory {
		if lp.Status == types.LPHistoryStatus_LP_SUBMITTING || lp.Status == types.LPHistoryStatus_LP_WAITING_FOR_SGN {
			resp, err := gs.QueryLiquidityStatus(nil, &webapi.QueryLiquidityStatusRequest{
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

func (gs *GatewayService) updateTransferStatusInHistory(ctx context.Context, transferList []*dal.Transfer) error {
	var transferIds []string
	for _, transfer := range transferList {
		transferIds = append(transferIds, transfer.TransferId)
	}
	tr := gs.tp.GetTransactor()
	transferMap, err := cbrcli.QueryTransferStatus(tr.CliCtx, &types.QueryTransferStatusRequest{
		TransferId: transferIds,
	})
	if err != nil {
		log.Errorf("updateTransferStatusInHistory when QueryTransferStatus in sgn failed, error: %+v", err)
		return err
	}
	transferStatusMap := transferMap.Status

	for _, transfer := range transferList {
		transferId := transfer.TransferId
		status := transfer.Status
		srcChainId := transfer.SrcChainId
		txHash := transfer.SrcTxHash
		if status == types.TransferHistoryStatus_TRANSFER_SUBMITTING {
			ec := gs.ec[srcChainId]
			if ec == nil {
				log.Errorf("no ethClient found for chain:%d", srcChainId)
				return fmt.Errorf("no ethClient found for chain:%d", srcChainId)
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
		if transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED ||
			transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED {
			if status == types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND || status == types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND {
				continue // user action, not updated by sgn
			}
			if status == transferStatusMap[transferId] {
				log.Debugf("status not change in polling for transfer:%s, status:%s", transfer.TransferId, status)
				continue
			}
			log.Debugf("update transfer refund status from sgn, current is %s, dst is %s", status.String(), transferStatusMap[transferId].String())
			dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(transferStatusMap[transferId]))
			if dbErr != nil {
				log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, txHash)
			}
		}
		return nil
	}
	return nil
}

func (gs *GatewayService) initTransactors() error {
	if selfStart {
		cbrCfgFile := filepath.Join(rootDir, "config", "cbridge.toml")
		viper.SetConfigFile(cbrCfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read in cbridge configuration: %w", err)
		}
		configFilePath := filepath.Join(rootDir, "config", "sgn.toml")
		viper.SetConfigFile(configFilePath)
		if err := viper.MergeInConfig(); err != nil {
			return fmt.Errorf("failed to read in SGN configuration: %w", err)
		}
	}

	tp := transactor.NewTransactorPool(rootDir, viper.GetString(common.FlagSgnChainId), legacyAmino, cdc, interfaceRegistry)
	err := tp.AddTransactors(
		viper.GetString(common.FlagSgnNodeURI), viper.GetString(common.FlagSgnPassphrase), viper.GetStringSlice(common.FlagSgnTransactors))
	if err != nil {
		return fmt.Errorf("failed to add transactors: %w", err)
	}
	gs.tp = tp

	var mcc []*common.OneChainConfig
	err = viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		return fmt.Errorf("failed to new mcc: %w", err)
	}
	e := make(map[uint64]*ethclient.Client)
	for _, m := range mcc {
		ec, ecErr := ethclient.Dial(m.Gateway)
		if ecErr == nil {
			e[m.ChainID] = ec
		}
	}
	gs.ec = e

	return nil
}

// todo cache this @aric
type txData struct {
	volume   float64
	fee      *big.Int
	dstToken *types.Token
}

func (gs *GatewayService) get24hTx() map[uint64]map[string]*txData {
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

type stakingInfo struct {
	staking int
	earning float64
}

func (gs *GatewayService) getUserStaking(ctx context.Context, address string) map[uint64]map[string]*stakingInfo {
	tr := gs.tp.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	stakingRes, err := queryClient.StakedPools(
		ctx,
		&farmingtypes.QueryStakedPoolsRequest{
			Address: address,
		},
	)
	log.Debugf("farming stakingRes:%+v", stakingRes)

	stakingPools := make(map[uint64]map[string]*stakingInfo) // map<chain_id, map<token_symbol, FarmingPool>>
	if err == nil {
		for _, pool := range stakingRes.GetPools() {
			erc20Token := pool.GetStakeToken()

			currentRes, earningErr := queryClient.Earnings(
				ctx,
				&farmingtypes.QueryEarningsRequest{
					PoolName: pool.GetName(),
					Address:  address,
				},
			)
			if earningErr != nil {
				log.Errorf("earningErr:%+v", earningErr)
				continue
			}
			log.Debugf("farming current earning reqAddr:%s, reqPool:%s, Res:%+v", address, pool.GetName(), currentRes)
			currentEarnings := currentRes.GetEarnings()
			earning := 0.0
			tokenSymbol := getSymbolFromFarmingToken(erc20Token.GetSymbol())
			for _, reward := range currentEarnings.GetRewardAmounts() {
				amt, parseErr := gs.getInfoFromFarmingReward(reward)
				if parseErr != nil {
					continue
				}
				earning += amt
				log.Debugf("chain:%d, token:%s, earning%.2f", erc20Token.GetChainId(), tokenSymbol, amt)
			}

			// return info
			staking, found := stakingPools[erc20Token.ChainId]
			if !found {
				staking = make(map[string]*stakingInfo)
			}
			staking[tokenSymbol] = &stakingInfo{
				staking: len(pool.GetRewardTokenInfos()),
				earning: earning,
			}
			stakingPools[erc20Token.ChainId] = staking
		}
	}
	return stakingPools
}

func (gs *GatewayService) getInfoFromFarmingReward(reward sdk.DecCoin) (float64, error) {
	chainId, tokenSymbol, parseErr := farmingkp.ParseERC20TokenDenom(reward.GetDenom())
	if parseErr != nil {
		log.Errorf("parse token denom error, denom:%s, err:%+v", reward.GetDenom(), parseErr)
	}
	tokenSymbol = getSymbolFromFarmingToken(tokenSymbol)
	token, found, dbErr := dal.DB.GetTokenBySymbol(tokenSymbol, chainId)
	if !found || dbErr != nil {
		return 0, dbErr
	}
	return gs.f.GetUsdVolume(token.Token, common.Str2BigInt(reward.Amount.String())), parseErr
}

func (gs *GatewayService) getUserFarmingCumulativeEarning(ctx context.Context, address string) map[uint64]map[string]float64 {
	tr := gs.tp.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.RewardClaimInfo(
		ctx,
		&farmingtypes.QueryRewardClaimInfoRequest{
			Address: address,
		},
	)
	log.Debugf("farming earningRes:%+v", res)
	earnings := make(map[uint64]map[string]float64) // map<chain_id, map<token_symbol, earning>>
	if res == nil || err != nil {
		return earnings
	}
	rewardClaimInfo := res.GetRewardClaimInfo()
	for _, detail := range rewardClaimInfo.GetRewardClaimDetailsList() {
		chainId := detail.GetChainId()
		earning, found := earnings[chainId]
		if !found {
			earning = make(map[string]float64)
		}
		for _, reward := range detail.GetCumulativeRewardAmounts() {
			_, tokenSymbol, parseErr := farmingkp.ParseERC20TokenDenom(reward.GetDenom())
			if parseErr != nil {
				log.Errorf("parse token denom error, denom:%s, err:%+v", reward.GetDenom(), parseErr)
				continue
			}
			amt, parseErr := reward.Amount.Float64()
			if parseErr != nil {
				log.Errorf("parse reward amt error, amt:%s, err:%+v", reward.Amount.String(), parseErr)
				continue
			}
			earning[tokenSymbol] += amt
		}
		earnings[chainId] = earning
	}
	return earnings
}

// todo cache this @aric
func (gs *GatewayService) getFarmingApy(ctx context.Context) map[uint64]map[string]float64 {
	tr := gs.tp.GetTransactor()
	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	res, err := queryClient.Pools(
		ctx,
		&farmingtypes.QueryPoolsRequest{},
	)
	if err != nil {
		return nil
	}
	farmingPools := make(map[uint64]map[string]float64) // map<chain_id, map<token_symbol, FarmingPool>>
	for _, pool := range res.GetPools() {
		farmingPool := make(map[string]float64)
		token := pool.GetStakeToken()
		tokenSymbol := getSymbolFromFarmingToken(token.GetSymbol())
		totalStakedAmount := pool.TotalStakedAmount
		if totalStakedAmount.Amount.Equal(sdk.ZeroDec()) {
			log.Debugf("farming totalStakedAmount is 0 on chain:%d, token: %s", token.GetChainId(), tokenSymbol)
			farmingPool[tokenSymbol] = 0.0
		} else {
			totalReward := sdk.ZeroDec()
			for _, reward := range pool.GetRewardTokenInfos() {
				totalReward = totalReward.Add(reward.RewardAmountPerBlock)
			}

			// apy=totalReward/totalStakedAmount
			apy, _ := new(big.Float).Quo(new(big.Float).SetInt(totalReward.BigInt()), new(big.Float).SetInt(totalStakedAmount.Amount.BigInt())).Float64()
			farmingPool[tokenSymbol] = apy
		}
		farmingPools[token.GetChainId()] = farmingPool
	}
	return farmingPools
}

func getSymbolFromFarmingToken(token string) string {
	return strings.Replace(token, "CB-", "", 1)
}

func unknownChain(chainId uint32) *webapi.Chain {
	return &webapi.Chain{
		Id:   chainId,
		Name: fmt.Sprintf("Chain-%d", chainId),
		Icon: "https://cbridge.celer.network/ETH.png",
	}
}

func enrichUnknownToken(token *webapi.TokenInfo) {
	if token.GetName() == "" {
		token.Name = token.Token.GetSymbol()
		token.Icon = "https://get.celer.app/cbridge-icons/ETH.png"
	}
}
