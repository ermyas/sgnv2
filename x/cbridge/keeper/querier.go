package keeper

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParams:
			return queryParams(ctx, k, legacyQuerierCdc)
		case types.QueryRelay:
			return queryRelay(ctx, req, k, legacyQuerierCdc)
		case types.QueryChainTokensConfig:
			return queryChainTokensConfig(ctx, req, k, legacyQuerierCdc)
		case types.QueryFee:
			return queryFee(ctx, req, k, legacyQuerierCdc)
		case types.QueryFeePerc:
			return queryFeePerc(ctx, req, k, legacyQuerierCdc)
		case types.QueryTransferStatus:
			return queryTransferStatus(ctx, req, k, legacyQuerierCdc)
		case types.QueryLiquidityDetailList:
			return queryLiquidityDetailList(ctx, req, k, legacyQuerierCdc)
		case types.QueryAddLiquidityStatus:
			return queryAddLiquidityStatus(ctx, req, k, legacyQuerierCdc)
		case types.QueryWithdrawLiquidityStatus:
			return queryWithdrawLiquidityStatus(ctx, req, k, legacyQuerierCdc)
		case types.QueryChainSigners:
			return queryChainSigners(ctx, req, k, legacyQuerierCdc)
		case types.QueryLatestSigners:
			return queryLatestSigners(ctx, k, legacyQuerierCdc)
		case types.QueryDebugAny:
			return queryDebugAny(ctx, req, k)
		case types.QueryCheckChainTokenValid:
			return queryCheckTokenValid(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Unknown cbridge query endpoint")
		}
	}
}

// req.data is key, return value raw, expect caller to decode properly
func queryDebugAny(ctx sdk.Context, req abci.RequestQuery, k Keeper) ([]byte, error) {
	kv := ctx.KVStore(k.storeKey)
	return kv.Get(req.Data), nil
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	cfg := k.GetCbrConfig(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, cfg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryRelay(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryRelayParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	var xferId [32]byte
	copy(xferId[:], params.XrefId)
	relay := GetXferRelay(ctx.KVStore(k.storeKey), xferId, k.cdc)
	if relay == nil {
		return nil, errors.New("relay does not exist")
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, relay)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryChainTokensConfig(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var mcc []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &mcc)
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
			log.Errorf("chain with Id %d is not configured", a.ChainId)
			return nil, fmt.Errorf("chain with Id %d is not configured", a.ChainId)
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
	resp := types.ChainTokensConfigResponse{
		ChainTokens: chainTokens,
	}
	res, err := k.cdc.Marshal(&resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func queryFeePerc(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.GetFeePercentageRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		log.Errorf("failed to parse params: %s", err)
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	kv := ctx.KVStore(k.storeKey)

	feePerc := GetFeePerc(kv, params.SrcChainId, params.DstChainId) // fee percent * 1e6
	resp := types.GetFeePercentageResponse{FeePerc: feePerc}
	res, err := k.cdc.Marshal(&resp)

	if err != nil {
		log.Errorf("failed to marshal response: %s", err)
		return nil, err
	}

	return res, nil
}

func queryFee(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.GetFeeRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		log.Errorf("failed to parse params: %s", err)
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	src := &ChainIdTokenAddr{
		ChId:      params.SrcChainId,
		TokenAddr: eth.Hex2Addr(params.SrcTokenAddr),
	}
	assetSym := GetAssetSymbol(ctx.KVStore(k.storeKey), src)
	srcToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, params.SrcChainId)
	destToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, params.DstChainId)
	destTokenAddr := eth.Hex2Addr(destToken.Addr)
	dest := &ChainIdTokenAddr{
		ChId:      params.DstChainId,
		TokenAddr: destTokenAddr,
	}
	srcAmt, _ := big.NewInt(0).SetString(params.Amt, 10)
	kv := ctx.KVStore(k.storeKey)
	destAmt := CalcEqualOnDestChain(kv, &ChainIdTokenDecimal{
		ChainIdTokenAddr: src,
		Decimal:          srcToken.Decimal,
	}, &ChainIdTokenDecimal{
		ChainIdTokenAddr: dest,
		Decimal:          destToken.Decimal,
	}, srcAmt, eth.Hex2Addr(params.LpAddr))

	resp := types.GetFeeResponse{
		EqValueTokenAmt: destAmt.String(),
		PercFee:         CalcPercFee(kv, src, dest, destAmt).String(),
		BaseFee:         CalcBaseFee(kv, assetSym, dest.ChId).String(),
		Decimal:         uint64(destToken.Decimal),
	}
	res, err := k.cdc.Marshal(&resp)
	if err != nil {
		log.Errorf("failed to marshal response: %s", err)
		return nil, err
	}

	return res, nil
}

func queryCheckTokenValid(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.CheckChainTokenValidRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		log.Errorf("failed to parse params: %s", err)
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	src := &ChainIdTokenAddr{
		ChId:      params.SrcChainId,
		TokenAddr: eth.Hex2Addr(params.SrcTokenAddr),
	}
	assetSym := GetAssetSymbol(ctx.KVStore(k.storeKey), src)
	srcToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, params.SrcChainId)
	destToken := GetAssetInfo(ctx.KVStore(k.storeKey), assetSym, params.DestChainId)

	resp := types.CheckChainTokenValidResponse{
		Valid: srcToken != nil && !srcToken.GetXferDisabled() && destToken != nil && !destToken.GetXferDisabled(),
	}
	res, err := k.cdc.Marshal(&resp)
	if err != nil {
		log.Errorf("failed to marshal response: %s", err)
		return nil, err
	}

	return res, nil
}

func queryTransferStatus(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryTransferStatusRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	status := make(map[string]*types.TransferStatus)

	for _, xferId := range params.TransferId {
		xferStatus := GetEvSendStatus(ctx.KVStore(k.storeKey), eth.Bytes2Hash(common.Hex2Bytes(xferId)))
		switch xferStatus {
		case types.XferStatus_UNKNOWN:
			status[xferId] = &types.TransferStatus{
				GatewayStatus: types.TransferHistoryStatus_TRANSFER_UNKNOWN,
				SgnStatus:     xferStatus,
			}
		case types.XferStatus_OK_TO_RELAY:
			status[xferId] = &types.TransferStatus{
				GatewayStatus: types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE,
				SgnStatus:     xferStatus,
			}
		case types.XferStatus_SUCCESS:
			status[xferId] = &types.TransferStatus{
				GatewayStatus: types.TransferHistoryStatus_TRANSFER_COMPLETED,
				SgnStatus:     xferStatus,
			}
		case types.XferStatus_BAD_LIQUIDITY, types.XferStatus_BAD_SLIPPAGE, types.XferStatus_BAD_XFER_DISABLED, types.XferStatus_BAD_DEST_CHAIN:
			status[xferId] = &types.TransferStatus{
				GatewayStatus: types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED,
				SgnStatus:     xferStatus,
			}
		case types.XferStatus_REFUND_REQUESTED:
			status[xferId] = &types.TransferStatus{
				GatewayStatus: types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND,
				SgnStatus:     xferStatus,
			}
		case types.XferStatus_REFUND_DONE:
			status[xferId] = &types.TransferStatus{
				GatewayStatus: types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED,
				SgnStatus:     xferStatus,
			}
		default:
			log.Errorln("unknown status:", xferStatus)
			status[xferId] = &types.TransferStatus{
				GatewayStatus: types.TransferHistoryStatus_TRANSFER_UNKNOWN,
				SgnStatus:     xferStatus,
			}
		}
	}

	resp := types.QueryTransferStatusResponse{
		Status: status,
	}
	res, err := k.cdc.Marshal(&resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func queryLiquidityDetailList(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.LiquidityDetailListRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	ldList := make([]*types.LiquidityDetail, 0)
	for _, pair := range params.ChainToken {
		tokenEthAddr := eth.Hex2Addr(pair.TokenAddr)
		lpEthAddr := eth.Hex2Addr(params.LpAddr)
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
			//TODO
			//LpFeeRate: ,
		})
	}

	resp := types.LiquidityDetailListResponse{
		LiquidityDetail: ldList,
	}
	res, err := k.cdc.Marshal(&resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func queryAddLiquidityStatus(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAddLiquidityStatusRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	var status types.LPHistoryStatus
	if HasEvLiqAdd(ctx.KVStore(k.storeKey), params.ChainId, params.SeqNum) {
		status = types.LPHistoryStatus_LP_COMPLETED
	} else {
		status = types.LPHistoryStatus_LP_WAITING_FOR_SGN
	}

	resp := types.QueryLiquidityStatusResponse{
		Status: status,
	}
	res, err := k.cdc.Marshal(&resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func queryWithdrawLiquidityStatus(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryWithdrawLiquidityStatusRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	var status types.LPHistoryStatus
	wd := GetWithdrawDetail(ctx.KVStore(k.storeKey), eth.Hex2Addr(params.UsrAddr), params.SeqNum)
	if wd == nil {
		return nil, fmt.Errorf("withdraw not exist, usr:%s seq: %d", params.UsrAddr, params.SeqNum)
	}

	if wd.Completed {
		status = types.LPHistoryStatus_LP_COMPLETED
	} else {
		wdOnchain := new(types.WithdrawOnchain)
		wdOnchain.Unmarshal(wd.WdOnchain)
		chainSigners, _ := k.GetChainSigners(ctx, wdOnchain.Chainid)
		if types.ValidateSigQuorum(wd.GetSortedSigs(), chainSigners.GetSortedSigners()) {
			status = types.LPHistoryStatus_LP_WAITING_FOR_LP
		} else {
			status = types.LPHistoryStatus_LP_WAITING_FOR_SGN
		}
	}

	resp := types.QueryLiquidityStatusResponse{
		Status: status,
		Detail: wd,
	}
	res, err := k.cdc.Marshal(&resp)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func queryChainSigners(
	ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryChainSignersParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}
	chainSigners, found := k.GetChainSigners(ctx, params.ChainId)
	if !found {
		return nil, types.ErrRecordNotFound
	}
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, chainSigners)
	if err != nil {
		log.Error(err)
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryLatestSigners(
	ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	latestSigners, found := k.GetLatestSigners(ctx)
	if !found {
		return nil, types.ErrRecordNotFound
	}
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, latestSigners)
	if err != nil {
		log.Error(err)
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
