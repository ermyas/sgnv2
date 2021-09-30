package gateway

import (
	"context"
	"fmt"
	"math/big"
	"path/filepath"
	"strconv"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
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
	tr *transactor.Transactor
	ec map[uint64]*ethclient.Client
}

func (gs *GatewayService) GetTransferStatus(ctx context.Context, request *webapi.GetTransferStatusRequest) (*webapi.GetTransferStatusResponse, error) {
	_, _, _, _, status, found, err := dal.DB.GetTransfer(request.GetTransferId())
	if !found || err != nil {
		return &webapi.GetTransferStatusResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "transfer not found",
			},
		}, nil
	}
	return &webapi.GetTransferStatusResponse{
		Status: cbrtypes.TransferHistoryStatus(status),
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
	dstToken, found2, err2 := dal.DB.GetTokenBySymbol(tokenSymbol, uint64(dstChainId))
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
		slippage = 5000
	}
	feeInfo, err := cbrcli.QueryFee(gs.tr.CliCtx, &cbrtypes.GetFeeRequest{
		SrcChainId:   uint64(srcChainId),
		DstChainId:   uint64(dstChainId),
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
	txHash := request.GetSrcTxHash()
	txType := request.GetType()
	withdrawSeqNum := request.GetWithdrawSeqNum()
	log.Infof("transferId in mark api: %s, bytes:%+v", transferId, common.Hex2Bytes(transferId))
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
	//userAddr := common.Hex2Addr(request.GetAddr()).String()
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
	//detailList, err := cbrcli.QueryLiquidityDetailList(gs.tr.CliCtx, &cbrtypes.LiquidityDetailListRequest{
	//	LpAddr:     userAddr,
	//	ChainToken: chainTokens,
	//})
	//if err != nil || detailList == nil || len(detailList.GetLiquidityDetail()) == 0 {
	//	return &webapi.GetLPInfoListResponse{}, nil
	//}
	//stakingMap := gs.getUserStaking(ctx, userAddr)
	//farmingApyMap := gs.getFarmingApy(ctx)
	//data24h := gs.get24hTx()
	//userDetailMap := make(map[uint64]map[string]*types.LiquidityDetail)
	//for _, detail := range detailList.GetLiquidityDetail() {
	//	chainId := detail.GetChainId()
	//	tokenWithAddr := detail.GetToken() // only has addr field
	//	token, found, err := dal.DB.GetTokenByAddr(tokenWithAddr.GetAddress(), chainId)
	//	if !found || err != nil {
	//		continue
	//	}
	//	detail.Token = token.Token
	//	chainInfo := make(map[string]*types.LiquidityDetail)
	//	chainInfo[token.Token.Symbol] = detail
	//	userDetailMap[chainId] = chainInfo
	//}
	for chainId32, chainToken := range chainTokenInfos {
		chainId := uint64(chainId32)
		for _, token := range chainToken.Token {
			//tokenSymbol := token.Token.Symbol
			totalLiquidity := "0"
			usrLpFeeEarning := "0"
			usrLiquidity := "0"
			//detail, found := userDetailMap[chainId][tokenSymbol]
			//if found {
			//	totalLiquidity = detail.GetTotalLiquidity()
			//	usrLpFeeEarning = detail.GetUsrLpFeeEarning()
			//	usrLiquidity = detail.GetUsrLiquidity()
			//}

			chain, _, found, err := dal.DB.GetChain(chainId)
			if !found || err != nil {
				chain = &webapi.Chain{
					Id:   uint32(chainId),
					Name: "UNKNOWN CHAIN",
					Icon: "",
				}
			}

			//data := data24h[chainId][tokenSymbol]
			lpFeeEarningApy := 0.0
			volume24h := 0.0
			//if data != nil {
			//	lpFeeEarningApy, _ = new(big.Float).Quo(new(big.Float).SetInt(data.fee), new(big.Float).SetInt(common.Str2BigInt(totalLiquidity))).Float64()
			//	volume24h = data.volume
			//}
			lp := &webapi.LPInfo{
				Chain:                chain,
				Token:                token,
				Liquidity:            gs.f.GetUsdVolume(token.Token, common.Str2BigInt(usrLiquidity)),
				HasFarmingSessions:   false, //[chainId][token.Token.GetSymbol()] > 0,
				LpFeeEarning:         gs.f.GetUsdVolume(token.Token, common.Str2BigInt(usrLpFeeEarning)),
				FarmingRewardEarning: 0, // // todo enrich 0 data from farming @aric
				Volume_24H:           volume24h,
				TotalLiquidity:       gs.f.GetUsdVolume(token.Token, common.Str2BigInt(totalLiquidity)),
				LpFeeEarningApy:      lpFeeEarningApy,
				FarmingApy:           0, //farmingApyMap[chainId][token.Token.GetSymbol()],
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
	err = dal.DB.UpsertLP(addr, token.GetToken().GetSymbol(), token.GetToken().GetAddress(), amt, txHash, uint64(chainId), uint64(cbrtypes.LPHistoryStatus_LP_SUBMITTING), uint64(lpType), seqNum)
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
	if transferId != "" {
		// refund transfer
		seqNum, err := gs.initWithdraw(&cbrtypes.MsgInitWithdraw{
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
		seqNum, err := gs.initWithdraw(&cbrtypes.MsgInitWithdraw{
			Chainid: uint64(chainId),
			LpAddr:  common.Hex2Bytes(lp),
			Token:   common.Hex2Bytes(tokenAddr),
			Amount:  common.Hex2Bytes(amt),
			Creator: gs.tr.Key.GetAddress().String(),
		})
		err = dal.DB.UpsertLP(lp, token.Token.Symbol, token.Token.Address, amt, "", uint64(chainId), uint64(cbrtypes.LPHistoryStatus_LP_WAITING_FOR_SGN), uint64(webapi.LPType_LP_TYPE_REMOVE), seqNum)
		if err != nil {
			_ = dal.DB.UpdateLPStatus(seqNum, uint64(cbrtypes.LPHistoryStatus_LP_FAILED))
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

func (gs *GatewayService) initWithdraw(req *cbrtypes.MsgInitWithdraw) (uint64, error) {
	resp, err := cbrcli.InitWithdraw(gs.tr, req)
	if resp == nil {
		return 0, fmt.Errorf("can not init withdraw, resp is empty")
	}
	return resp.GetSeqnum(), err
}

// for withdraw only
func (gs *GatewayService) QueryLiquidityStatus(ctx context.Context, request *webapi.QueryLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	seqNum := request.GetSeqNum()
	chainId := uint64(request.GetChainId())
	lpType := uint64(request.GetType())
	addr := request.GetLpAddr()
	txHash, status, found, err := dal.DB.GetLPInfo(seqNum, lpType, chainId, addr)
	if found && err == nil && status == uint64(cbrtypes.LPHistoryStatus_LP_SUBMITTING) && txHash != "" {
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
			dbErr := dal.DB.UpdateLPStatus(seqNum, uint64(cbrtypes.LPHistoryStatus_LP_FAILED))
			if dbErr != nil {
				log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", chainId, txHash)
			}
		}

	}

	resp, err := cbrcli.QueryWithdrawLiquidityStatus(gs.tr.CliCtx, &cbrtypes.QueryWithdrawLiquidityStatusRequest{
		SeqNum: seqNum,
	})

	curss, err := cbrcli.QueryChainSigners(gs.tr.CliCtx, chainId)
	if resp == nil || err != nil {
		return &types.QueryLiquidityStatusResponse{
			Status:  types.LPHistoryStatus(status),
			Detail:  nil,
			Signers: nil,
		}, err
	} else if resp.GetStatus() == types.LPHistoryStatus_LP_WAITING_FOR_LP {
		_ = dal.DB.UpdateLPStatus(seqNum, uint64(types.LPHistoryStatus_LP_WAITING_FOR_LP))
		resp.Signers = curss.GetSignersBytes()
	} else {
		resp.Status = types.LPHistoryStatus(status)
		resp.Signers = curss.GetSignersBytes()
	}
	return resp, nil
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
	resp, err := cbrcli.QueryChainTokensConfig(gs.tr.CliCtx, &cbrtypes.ChainTokensConfigRequest{})
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
			dbErr := dal.DB.UpsertTokenBaseInfo(token.GetSymbol(), common.Hex2Addr(token.GetAddress()).String(), common.Hex2Addr(asset.GetContractAddr()).String(), asset.GetMaxAmt(), uint64(chainId), uint64(token.GetDecimal()))
			if dbErr != nil {
				log.Errorf("failed to write token: %v", err)
			}
		}
	}
}

func (gs *GatewayService) updateLpStatusInHistory(lpHistory []*dal.LP) {
	for _, lp := range lpHistory {
		if lp.Status == cbrtypes.LPHistoryStatus_LP_SUBMITTING || lp.Status == cbrtypes.LPHistoryStatus_LP_WAITING_FOR_SGN {
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
	transferMap, err := cbrcli.QueryTransferStatus(gs.tr.CliCtx, &cbrtypes.QueryTransferStatusRequest{
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
		if status == cbrtypes.TransferHistoryStatus_TRANSFER_SUBMITTING {
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
				dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(cbrtypes.TransferHistoryStatus_TRANSFER_FAILED))
				if dbErr != nil {
					log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, txHash)
				}
			}
		}

		if status == cbrtypes.TransferHistoryStatus_TRANSFER_FAILED ||
			status == cbrtypes.TransferHistoryStatus_TRANSFER_COMPLETED ||
			status == cbrtypes.TransferHistoryStatus_TRANSFER_REFUNDED {
			continue
		}
		if transferStatusMap[transferId] == cbrtypes.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED ||
			transferStatusMap[transferId] == cbrtypes.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND ||
			transferStatusMap[transferId] == cbrtypes.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED {
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
	if selfStart {
		configFilePath := filepath.Join(rootDir, "config", "sgn.toml")
		viper.SetConfigFile(configFilePath)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read in SGN configuration: %w", err)
		}
	}

	tr, err := transactor.NewTransactor(
		rootDir,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		legacyAmino,
		cdc,
		interfaceRegistry,
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
	dstToken *cbrtypes.Token
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
		var totalReward sdk.Dec
		for _, reward := range pool.GetRewardTokenInfos() {
			totalReward = totalReward.Add(reward.RewardAmountPerBlock)
		}
		apy, _ := totalReward.Quo(totalStakedAmount.Amount).Float64()
		farmingPool[token.GetSymbol()] = apy
		farmingPools[token.GetChainId()] = farmingPool
	}
	return farmingPools
}
