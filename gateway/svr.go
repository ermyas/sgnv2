package gateway

import (
	"context"
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"time"
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
	tr *transactor.Transactor
	ec map[uint64]*ethclient.Client
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
	var chainIds []uint64
	for key := range chainTokenList {
		chainIds = append(chainIds, key)
	}
	chains, err := dal.DB.GetChainInfo(chainIds)
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
	srcToken, found1, err1 := dal.DB.GetTokenBySymbol(tokenSymbol, srcChainId)
	dstToken, found2, err2 := dal.DB.GetTokenBySymbol(tokenSymbol, dstChainId)
	if err1 != nil || !found1 || err2 != nil || !found2 {
		return &webapi.EstimateAmtResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "token not found",
			},
		}, nil
	}
	addr := common.Hex2Addr(request.GetUsrAddr()).String()
	slippage, found, err := dal.DB.GetSlippageSetting(addr)
	if err != nil || !found {
		slippage = 0
	}
	feeInfo, err := cli.GetFee(gs.tr.CliCtx, &types.GetFeeRequest{
		SrcChainId:   srcChainId,
		DstChainId:   dstChainId,
		SrcTokenAddr: srcToken.Token.GetAddress(),
		Amt:          amt,
	})
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
	}
	return &webapi.EstimateAmtResponse{
		EqValueTokenAmt:   eqValueTokenAmt,
		BridgeRate:        float32(bridgeRate),
		Fee:               feeAmt,
		SlippageTolerance: slippage,
	}, nil
}

func (gs *GatewayService) MarkTransfer(ctx context.Context, request *webapi.MarkTransferRequest) (*webapi.MarkTransferResponse, error) {
	transferId := request.GetTransferId()
	addr := common.Hex2Addr(request.GetAddr())
	sendInfo := request.GetSrcSendInfo()
	receivedInfo := request.GetDstMinReceivedInfo()
	dstTransferId := request.GetDstTransferId()
	txHash := request.GetSrcTxHash()
	txType := request.GetType()
	withdrawSeqNum := request.GetWithdrawSeqNum()
	if txType == webapi.TransferType_TRANSFER_TYPE_SEND {
		err := dal.DB.MarkTransferSend(transferId, dstTransferId, addr.String(), sendInfo.GetToken().GetSymbol(),
			sendInfo.GetAmount(), receivedInfo.GetAmount(), txHash, sendInfo.GetChain().GetId(),
			receivedInfo.GetChain().GetId(), gs.f.GetUsdVolume(sendInfo.GetToken(), common.Str2BigInt(sendInfo.GetAmount())))
		if err != nil {
			return &webapi.MarkTransferResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "mark transfer refund failed",
				},
			}, nil
		}
	} else if txType == webapi.TransferType_TRANSFER_TYPE_REFUND {
		err := dal.DB.MarkTransferRefund(transferId, txHash, withdrawSeqNum)
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
	chainTokens, err := dal.DB.GetAllLpChainToken(userAddr)
	if err != nil || len(chainTokens) == 0 {
		return &webapi.GetLPInfoListResponse{}, nil
	}
	var lps []*webapi.LPInfo
	detailList, err := cli.LiquidityDetailList(gs.tr.CliCtx, &types.LiquidityDetailListRequest{
		LpAddr:     userAddr,
		ChainToken: chainTokens,
	})
	if err != nil || detailList == nil || len(detailList.GetLiquidityDetail()) == 0 {
		return &webapi.GetLPInfoListResponse{}, nil
	}
	stakingMap := gs.getUserStaking(ctx, userAddr)
	farmingApyMap := gs.getFarmingApy(ctx)
	data24h := gs.get24hTx()
	for _, detail := range detailList.GetLiquidityDetail() {
		chainId := detail.GetChainId()
		tokenWithAddr := detail.GetToken() // only has addr field
		totalLiquidity := detail.GetTotalLiquidity()
		usrLpFeeEarning := detail.GetUsrLpFeeEarning()
		usrLiquidity := detail.GetUsrLiquidity()
		chain, _, found, err := dal.DB.GetChain(chainId)
		if !found || err != nil {
			continue
		}
		token, found, err := dal.DB.GetTokenByAddr(tokenWithAddr.GetAddress(), chainId)
		if !found || err != nil {
			continue
		}

		data := data24h[chainId][token.Token.GetSymbol()]
		lpFeeEarningApy := 0.0
		volume24h := 0.0
		if data != nil {
			lpFeeEarningApy, _ = new(big.Float).Quo(new(big.Float).SetInt(data.fee), new(big.Float).SetInt(common.Str2BigInt(totalLiquidity))).Float64()
			volume24h = data.volume
		}
		lp := &webapi.LPInfo{
			Chain:                chain,
			Token:                token,
			Liquidity:            gs.f.GetUsdVolume(token.Token, common.Str2BigInt(usrLiquidity)),
			HasFarmingSessions:   stakingMap[chainId][token.Token.GetSymbol()] > 0,
			LpFeeEarning:         gs.f.GetUsdVolume(token.Token, common.Str2BigInt(usrLpFeeEarning)),
			FarmingRewardEarning: 0, // // todo enrich 0 data from farming @aric
			Volume_24H:           volume24h,
			TotalLiquidity:       gs.f.GetUsdVolume(token.Token, common.Str2BigInt(totalLiquidity)),
			LpFeeEarningApy:      lpFeeEarningApy,
			FarmingApy:           farmingApyMap[chainId][token.Token.GetSymbol()],
		}
		lps = append(lps, lp)
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
	token, found, err := dal.DB.GetTokenByAddr(tokenAddr, chainId)
	if !found || err != nil {
		return &webapi.MarkLiquidityResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "token not found in gateway DB",
			},
		}, nil
	}
	txHash := request.GetTxHash()
	err = dal.DB.UpsertLP(addr, token.GetToken().GetSymbol(), token.GetToken().GetAddress(), amt, txHash, chainId, uint64(types.LPHistoryStatus_LP_SUBMITTING), uint64(lpType), seqNum)
	if err != nil {
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
	if transferId != "" {
		// refund transfer
		seqNum, err := gs.initWithdraw(&types.MsgInitWithdraw{
			XferId:  common.Hex2Bytes(transferId),
			Creator: gs.tr.Key.GetAddress().String(),
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
		token, found, err := dal.DB.GetTokenByAddr(tokenAddr, chainId)
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
			Chainid: chainId,
			LpAddr:  common.Hex2Bytes(lp),
			Token:   common.Hex2Bytes(tokenAddr),
			Amount:  common.Hex2Bytes(amt),
			Creator: gs.tr.Key.GetAddress().String(),
		})
		err = dal.DB.UpsertLP(lp, token.Token.Symbol, token.Token.Address, amt, "", chainId, uint64(types.LPHistoryStatus_LP_WAITING_FOR_SGN), uint64(webapi.LPType_LP_TYPE_REMOVE), seqNum)
		if err != nil {
			_ = dal.DB.UpdateLPStatus(seqNum, uint64(types.LPHistoryStatus_LP_FAILED))
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
	resp, err := cli.InitWithdraw(gs.tr, req)
	if resp == nil {
		return 0, fmt.Errorf("can not init withdraw, resp is empty")
	}
	return resp.GetSeqnum(), err
}

// for withdraw only
func (gs *GatewayService) QueryLiquidityStatus(ctx context.Context, request *webapi.QueryLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	seqNum := request.SeqNum
	chainId, txHash, status, found, err := dal.DB.GetLPInfo(seqNum)
	if found && err == nil && status == uint64(types.LPHistoryStatus_LP_SUBMITTING) && txHash != "" {
		ec := gs.ec[chainId]
		if ec == nil {
			gs.initTransactor()
			ec = gs.ec[chainId]
		}
		if ec == nil {
			log.Errorf("no ethClient found for chain:%d", chainId)
			return nil, fmt.Errorf("no ethClient found for chain:%d", chainId)
		}

		receipt, recErr := ec.TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(txHash)))
		if recErr == nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
			log.Warnf("find transfer failed, chain_id %d, hash:%s", chainId, txHash)
			dbErr := dal.DB.UpdateLPStatus(seqNum, uint64(types.LPHistoryStatus_LP_FAILED))
			if dbErr != nil {
				log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", chainId, txHash)
			}
		}

	}

	resp, _ := cli.QueryWithdrawLiquidityStatus(gs.tr.CliCtx, &types.QueryWithdrawLiquidityStatusRequest{
		SeqNum: seqNum,
	})
	if resp.GetStatus() == types.LPHistoryStatus_LP_WAITING_FOR_LP {
		_ = dal.DB.UpdateLPStatus(seqNum, uint64(types.LPHistoryStatus_LP_WAITING_FOR_LP))
	}
	return &types.QueryLiquidityStatusResponse{
		Status: resp.GetStatus(),
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
	err = gs.updateTransferStatusInHistory(ctx, transferList)
	if err != nil {
		log.Warnf("update transfer status failed for user:%s, error:%v", addr, err)
	}
	var transfers []*webapi.TransferHistory
	for _, transfer := range transferList {
		srcChain, srcChainUrl, srcFound, err1 := dal.DB.GetChain(transfer.SrcChainId)
		dstChain, dstChainUrl, dstFound, err2 := dal.DB.GetChain(transfer.SrcChainId)
		if !srcFound || !dstFound || err1 != nil || err2 != nil {
			continue
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
		return &webapi.LPHistoryResponse{}, nil
	}
	gs.UpdateLpStatusInHistory(lpHistory)
	var lps []*webapi.LPHistory
	for _, lp := range lpHistory {
		chain, chainUrl, found, lpErr := dal.DB.GetChain(lp.ChainId)
		if !found || lpErr != nil {
			continue
		}
		token, found, lpErr := dal.DB.GetTokenBySymbol(lp.TokenSymbol, lp.ChainId)
		if !found || lpErr != nil {
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
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
	config.Seal()
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
	resp, err := cli.ChainTokensConfig(gs.tr.CliCtx, &types.ChainTokensConfigRequest{})
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
			//_, found := gs.f.Prices[token.Symbol]
			//if !found {
			//	gs.f.Prices[token.Symbol], err = gs.f.GetUsdPrice(token.Symbol)
			//	if err != nil {
			//		log.Error("get price error", err)
			//	}
			//}
			dbErr := dal.DB.UpsertTokenBaseInfo(token.GetSymbol(), token.GetAddress(), asset.GetContractAddr(), asset.GetMaxAmt(), uint64(chainId), uint64(token.GetDecimal()))
			if dbErr != nil {
				log.Errorf("failed to write token: %v", err)
			}
		}
	}
}

func (gs *GatewayService) UpdateLpStatusInHistory(lpHistory []*dal.LP) {
	for _, lp := range lpHistory {
		if lp.Status == types.LPHistoryStatus_LP_SUBMITTING || lp.Status == types.LPHistoryStatus_LP_WAITING_FOR_SGN {
			resp, err := gs.QueryLiquidityStatus(nil, &webapi.QueryLiquidityStatusRequest{
				SeqNum: lp.SeqNum,
			})
			if err != nil {
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
	transferMap, err := cli.QueryTransferStatus(gs.tr.CliCtx, &types.QueryTransferStatusRequest{
		TransferId: transferIds,
	})
	if err != nil {
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
				gs.initTransactor()
				ec = gs.ec[srcChainId]
			}
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

		if err != nil {
			return err
		}
		if status == types.TransferHistoryStatus_TRANSFER_FAILED ||
			status == types.TransferHistoryStatus_TRANSFER_COMPLETED ||
			status == types.TransferHistoryStatus_TRANSFER_REFUNDED {
			continue
		}
		if transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED ||
			transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND ||
			transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED {
			dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(transferStatusMap[transferId]))
			if dbErr != nil {
				log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, txHash)
			}
		}
		return nil
	}
	return nil
}

func (gs *GatewayService) initTransactor() error {
	rootDir := os.ExpandEnv("$HOME/.sgnd")
	configFilePath := filepath.Join(rootDir, "config", "sgn.toml")
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read in SGN configuration: %w", err)
	}

	encodingConfig := app.MakeEncodingConfig()
	tr, err := transactor.NewTransactor(
		rootDir,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		encodingConfig.Amino,
		encodingConfig.Codec,
		encodingConfig.InterfaceRegistry,
	)
	if err != nil {
		return fmt.Errorf("failed to new transactor: %w", err)
	}
	tr.Run()
	gs.tr = tr

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

func (gs *GatewayService) getUserStaking(ctx context.Context, address string) map[uint64]map[string]int {
	queryClient := farmingtypes.NewQueryClient(gs.tr.CliCtx)
	stakingRes, err := queryClient.StakedPools(
		ctx,
		&farmingtypes.QueryStakedPoolsRequest{
			Address: address,
		},
	)
	stakingPools := make(map[uint64]map[string]int) // map<chain_id, map<token_symbol, FarmingPool>>
	if err == nil {
		for _, pool := range stakingRes.GetPools() {
			staking := make(map[string]int)
			token := pool.GetStakeToken()
			staking[token.Symbol] = len(pool.GetRewardTokenInfos())
			stakingPools[token.ChainId] = staking
		}
	}
	return stakingPools
}

// todo cache this @aric
func (gs *GatewayService) getFarmingApy(ctx context.Context) map[uint64]map[string]float64 {
	queryClient := farmingtypes.NewQueryClient(gs.tr.CliCtx)
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

		totalStakedAmount := pool.TotalStakedAmount
		var totalReward github_com_cosmos_cosmos_sdk_types.Dec
		for _, reward := range pool.GetRewardTokenInfos() {
			totalReward = totalReward.Add(reward.RewardAmountPerBlock)
		}
		apy, _ := totalReward.Quo(totalStakedAmount.Amount).Float64()
		farmingPool[token.GetSymbol()] = apy
		farmingPools[token.GetChainId()] = farmingPool
	}
	return farmingPools
}
