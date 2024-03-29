package keeper

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryParams(c context.Context, request *types.EmptyRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: &params}, nil
}

func (k Keeper) QueryConfig(c context.Context, request *types.EmptyRequest) (*types.QueryConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	config := k.GetCbrConfig(ctx)
	return &types.QueryConfigResponse{CbrConfig: &config}, nil
}

func (k Keeper) QueryDebugAny(c context.Context, request *types.QueryDebugAnyRequest) (*types.QueryDebugAnyResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	kv := ctx.KVStore(k.storeKey)
	return &types.QueryDebugAnyResponse{Data: kv.Get(request.Key)}, nil
}

func (k Keeper) QueryRelay(c context.Context, request *types.QueryRelayRequest) (*types.QueryRelayResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	relay, found := k.GetXferRelay(ctx, eth.Bytes2Hash(request.XrefId))
	if !found {
		return nil, sdkerrors.ErrKeyNotFound.Wrap("relay does not exist")
	}
	return &types.QueryRelayResponse{
		XferRelay: relay,
	}, nil
}

func (k Keeper) QueryRefund(c context.Context, request *types.QueryRefundRequest) (*types.QueryRefundResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	kv := ctx.KVStore(k.storeKey)
	wdOnchain := GetXferRefund(kv, eth.Bytes2Hash(request.XrefId))
	if wdOnchain == nil {
		return nil, sdkerrors.ErrKeyNotFound.Wrap("refund does not exist")
	}
	return &types.QueryRefundResponse{
		WdOnchain: wdOnchain,
	}, nil
}

func (k Keeper) QueryChainSigners(c context.Context, request *types.QueryChainSignersRequest) (*types.QueryChainSignersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	chainSigners, found := k.GetChainSigners(ctx, request.ChainId)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound.Wrap(fmt.Sprintf("chain %d has no signers", request.ChainId))
	}
	return &types.QueryChainSignersResponse{ChainSigners: &chainSigners}, nil
}
func (k Keeper) QueryLatestSigners(c context.Context, request *types.EmptyRequest) (*types.QueryLatestSignersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	latestSigners, found := k.GetLatestSigners(ctx)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound.Wrap("no current signers")
	}
	return &types.QueryLatestSignersResponse{LatestSigners: &latestSigners}, nil
}

func (k Keeper) QueryCheckChainTokenValid(c context.Context, request *types.CheckChainTokenValidRequest) (*types.CheckChainTokenValidResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	src := &ChainIdTokenAddr{
		ChId:      request.SrcChainId,
		TokenAddr: eth.Hex2Addr(request.SrcTokenAddr),
	}
	assetSym := GetAssetSymbol(ctx.KVStore(k.storeKey), src)
	srcToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, request.SrcChainId)
	destToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, request.DestChainId)

	resp := &types.CheckChainTokenValidResponse{
		Valid: srcToken != nil && !srcToken.GetXferDisabled() && destToken != nil && !destToken.GetXferDisabled(),
	}
	return resp, nil
}

func (k Keeper) QueryChkLiqSum(c context.Context, req *types.CheckLiqSumRequest) (*types.CheckLiqSumResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	kv := ctx.KVStore(k.storeKey)
	chtok := &ChainIdTokenAddr{
		ChId:      req.ChainId,
		TokenAddr: eth.Hex2Addr(req.TokenAddr),
	}
	resp := &types.CheckLiqSumResponse{
		Liqsum:  GetLiq(kv, chtok).String(),
		Sumiter: GetLiqIterSum(kv, chtok).String(),
	}
	return resp, nil
}

func (k Keeper) ChainTokensConfig(c context.Context, request *types.ChainTokensConfigRequest) (resp *types.ChainTokensConfigResponse, err error) {
	ctx := sdk.UnwrapSDKContext(c)
	var mcc []*common.OneChainConfig
	err = viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		log.Fatalln("fail to load multichain configs err:", err)
	}
	mccMap := make(map[uint64]*common.OneChainConfig)
	for _, occ := range mcc {
		mccMap[occ.ChainID] = occ
	}

	mca := k.GetCbrConfig(ctx)
	if len(mca.Assets) == 0 {
		log.Warnln("no chain assets configured yet!")
	}
	chainTokens := make(map[string]*types.Tokens)
	for _, a := range mca.Assets {
		occ, ok := mccMap[a.ChainId]
		if !ok {
			log.Warnf("chain %d is not configured in multichain, assume it's off shelf", a.ChainId)
			occ = &common.OneChainConfig{}
		}

		chid := strconv.FormatUint(a.ChainId, 10)
		assets, ok := chainTokens[chid]
		if !ok {
			assets = &types.Tokens{
				Tokens:       make([]*types.Token, 0),
				ContractAddr: occ.CBridge,
				BlockDelay:   uint32(occ.BlkDelay),
			}
			chainTokens[chid] = assets
		}
		assets.Tokens = append(assets.Tokens, &types.Token{
			Symbol:       a.Symbol,
			Address:      a.Addr,
			Decimal:      int32(a.Decimal),
			XferDisabled: a.XferDisabled,
		},
		)
	}
	resp = &types.ChainTokensConfigResponse{
		ChainTokens: chainTokens,
	}

	return resp, nil
}

func (k Keeper) GetFee(c context.Context, request *types.GetFeeRequest) (*types.GetFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	src := &ChainIdTokenAddr{
		ChId:      request.SrcChainId,
		TokenAddr: eth.Hex2Addr(request.SrcTokenAddr),
	}
	assetSym := GetAssetSymbol(ctx.KVStore(k.storeKey), src)
	srcToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, request.SrcChainId)
	destToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, request.DstChainId)
	destTokenAddr := eth.Hex2Addr(destToken.Addr)
	dest := &ChainIdTokenAddr{
		ChId:      request.DstChainId,
		TokenAddr: destTokenAddr,
	}
	srcAmt, _ := big.NewInt(0).SetString(request.Amt, 10)
	kv := ctx.KVStore(k.storeKey)
	destAmt, err := CalcEqualOnDestChain(kv, &ChainIdTokenDecimal{
		ChainIdTokenAddr: src,
		Decimal:          srcToken.Decimal,
	}, &ChainIdTokenDecimal{
		ChainIdTokenAddr: dest,
		Decimal:          destToken.Decimal,
	}, srcAmt, eth.Hex2Addr(request.LpAddr))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	resp := &types.GetFeeResponse{
		EqValueTokenAmt: destAmt.String(),
		PercFee:         CalcPercFee(kv, src, dest, destAmt).String(),
		Decimal:         uint64(destToken.Decimal),
	}
	if request.LpAddr == "" {
		resp.BaseFee = CalcBaseFee(kv, assetSym, dest.ChId, dest.ChId).String()
	}
	return resp, nil
}

func (k Keeper) GetFeePercentage(c context.Context, request *types.GetFeePercentageRequest) (*types.GetFeePercentageResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	kv := ctx.KVStore(k.storeKey)
	feePerc := GetFeePerc(kv, request.SrcChainId, request.DstChainId, request.Symbol) // fee percent * 1e6
	resp := &types.GetFeePercentageResponse{FeePerc: feePerc}

	return resp, nil
}

func (k Keeper) QueryTransferStatus(c context.Context, request *types.QueryTransferStatusRequest) (*types.QueryTransferStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	status := make(map[string]*types.TransferStatus)

	for _, xferId := range request.TransferId {
		xferStatus := GetEvSendStatus(ctx.KVStore(k.storeKey), eth.Hex2Hash(xferId))
		var xferHistoryStatus types.TransferHistoryStatus
		switch xferStatus {
		case types.XferStatus_UNKNOWN,
			types.XferStatus_BAD_ADDRESS,
			types.XferStatus_BAD_TOKEN,
			types.XferStatus_XFER_DELAYED:
			// Just map to UNKNOWN history status
			xferHistoryStatus = types.TransferHistoryStatus_TRANSFER_UNKNOWN
		case types.XferStatus_OK_TO_RELAY:
			xferHistoryStatus = types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE
		case types.XferStatus_SUCCESS:
			xferHistoryStatus = types.TransferHistoryStatus_TRANSFER_COMPLETED
		case types.XferStatus_BAD_LIQUIDITY,
			types.XferStatus_BAD_SLIPPAGE,
			types.XferStatus_BAD_XFER_DISABLED,
			types.XferStatus_BAD_DEST_CHAIN,
			types.XferStatus_EXCEED_MAX_OUT_AMOUNT:
			xferHistoryStatus = types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED
		case types.XferStatus_REFUND_REQUESTED:
			xferHistoryStatus = types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND
		case types.XferStatus_REFUND_DONE:
			xferHistoryStatus = types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED
		default:
			log.Errorln("unknown status:", xferStatus)
			xferHistoryStatus = types.TransferHistoryStatus_TRANSFER_UNKNOWN
		}
		status[xferId] = &types.TransferStatus{
			GatewayStatus: xferHistoryStatus,
			SgnStatus:     xferStatus,
		}
	}

	resp := &types.QueryTransferStatusResponse{
		Status: status,
	}

	return resp, nil
}

func (k Keeper) LiquidityDetailList(c context.Context, request *types.LiquidityDetailListRequest) (*types.LiquidityDetailListResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	ldList := make([]*types.LiquidityDetail, 0)
	for _, pair := range request.ChainToken {
		tokenEthAddr := eth.Hex2Addr(pair.TokenAddr)
		lpEthAddr := eth.Hex2Addr(request.LpAddr)
		ldList = append(ldList, &types.LiquidityDetail{
			ChainId: pair.ChainId,
			Token: &types.Token{
				Address: pair.TokenAddr,
			},
			UsrLiquidity:    GetLPBalance(ctx.KVStore(k.storeKey), pair.ChainId, tokenEthAddr, lpEthAddr).String(),
			UsrLpFeeEarning: GetLPFee(ctx.KVStore(k.storeKey), pair.ChainId, tokenEthAddr, lpEthAddr).String(),
			TotalLiquidity: GetLiq(ctx.KVStore(k.storeKey), &ChainIdTokenAddr{
				ChId:      pair.ChainId,
				TokenAddr: tokenEthAddr,
			}).String(),
		})
	}

	resp := &types.LiquidityDetailListResponse{
		LiquidityDetail: ldList,
	}
	return resp, nil
}

func (k Keeper) QueryTotalLiquidity(c context.Context, request *types.QueryTotalLiquidityRequest) (*types.QueryTotalLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	totalLiquidity := GetLiq(ctx.KVStore(k.storeKey), &ChainIdTokenAddr{
		ChId:      request.GetChainId(),
		TokenAddr: eth.Hex2Addr(request.GetTokenAddr()),
	}).String()

	resp := &types.QueryTotalLiquidityResponse{TotalLiq: totalLiquidity}
	return resp, nil
}

func (k Keeper) QueryAddLiquidityStatus(c context.Context, request *types.QueryAddLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var status types.WithdrawStatus
	if HasEvLiqAdd(ctx.KVStore(k.storeKey), request.ChainId, request.SeqNum) {
		status = types.WithdrawStatus_WD_COMPLETED
	} else {
		status = types.WithdrawStatus_WD_WAITING_FOR_SGN
	}

	resp := &types.QueryLiquidityStatusResponse{
		Status: status,
	}
	return resp, nil
}

func (k Keeper) QueryWithdrawLiquidityStatus(c context.Context, request *types.QueryWithdrawLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	var status types.WithdrawStatus
	wd := GetWithdrawDetail(ctx.KVStore(k.storeKey), eth.Hex2Addr(request.UsrAddr), request.SeqNum)
	if wd == nil {
		return nil, sdkerrors.ErrKeyNotFound.Wrap(fmt.Sprintf("withdraw not exist, usr:%s seq: %d", request.UsrAddr, request.SeqNum))
	}

	if wd.Completed {
		status = types.WithdrawStatus_WD_COMPLETED
	} else {
		wdOnchain := new(types.WithdrawOnchain)
		wdOnchain.Unmarshal(wd.WdOnchain)
		chainSigners, _ := k.GetChainSigners(ctx, wdOnchain.Chainid)
		pass, _ := types.ValidateSigQuorum(wd.GetSortedSigs(), chainSigners.GetSortedSigners())
		if pass {
			status = types.WithdrawStatus_WD_WAITING_FOR_LP
		} else {
			status = types.WithdrawStatus_WD_WAITING_FOR_SGN
		}
	}

	resp := &types.QueryLiquidityStatusResponse{
		Status: status,
		Detail: wd,
	}
	return resp, nil
}

func (k Keeper) QueryLiquidity(c context.Context, request *types.QueryLiquidityRequest) (*types.QueryLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	return &types.QueryLiquidityResponse{
		LpBalance: GetLPsBalanceAtChain(store, request.GetChainId(), eth.Hex2Addr(request.GetTokenAddr())),
	}, nil
}

func (k Keeper) QueryLPs(c context.Context, req *types.QueryLPsRequest) (*types.QueryLPsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	lps, err := GetLPs(store, &ChainIdTokenAddr{req.ChainId, eth.Hex2Addr(req.TokenAddr)})
	if err != nil {
		log.Errorln(err)
		return nil, status.Error(codes.Internal, "invalid key")
	}
	addrs := make([]string, 0)
	for _, lp := range lps {
		addrs = append(addrs, lp.String())
	}
	return &types.QueryLPsResponse{Lps: addrs}, nil
}

func (k Keeper) QueryAssets(c context.Context, request *types.EmptyRequest) (*types.QueryAssetsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	assets := make([]*types.ChainAsset, 0)
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, []byte("cfg-sym2info-"))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		assetRaw := iter.Value()
		asset := new(types.ChainAsset)
		asset.Unmarshal(assetRaw)
		assets = append(assets, asset)
	}
	return &types.QueryAssetsResponse{Assets: assets}, nil
}

func (k Keeper) QueryAssetPrice(c context.Context, request *types.QueryAssetPriceRequest) (*types.QueryAssetPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	price, extraPower := GetAssetUsdPrice(store, request.GetSymbol())
	return &types.QueryAssetPriceResponse{Price: price, ExtraPower10: extraPower}, nil
}

func (k Keeper) QueryLPOrigin(c context.Context, request *types.QueryLPOriginRequest) (*types.QueryLPOriginResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	chainId := GetLPOrigin(store, eth.Hex2Addr(request.UsrAddr))
	return &types.QueryLPOriginResponse{ChainId: chainId}, nil
}

func (k Keeper) QueryAssetsSymbols(c context.Context, request *types.QueryAssetsSymbolsRequest) (*types.QueryAssetsSymbolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	symbols := make([]string, 0)
	for _, chainToken := range request.ChainTokens {
		symbol := GetAssetSymbol(store, &ChainIdTokenAddr{ChId: chainToken.ChainId, TokenAddr: eth.Hex2Addr(chainToken.TokenAddr)})
		symbols = append(symbols, symbol)
	}
	return &types.QueryAssetsSymbolsResponse{Symbols: symbols}, nil
}

func (k Keeper) QueryAssetsInfos(c context.Context, request *types.QueryAssetsInfosRequest) (*types.QueryAssetsInfosResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	assets := make([]*types.ChainAsset, 0)
	if len(request.Symbols) != len(request.ChainIds) {
		return nil, fmt.Errorf("length mismatch")
	}
	for i := range request.Symbols {
		asset := GetAssetInfo(store, request.Symbols[i], request.ChainIds[i])
		assets = append(assets, asset)
	}
	return &types.QueryAssetsInfosResponse{Assets: assets}, nil
}
