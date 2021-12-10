package keeper

import (
	"context"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ChainTokensConfig(c context.Context, req *types.ChainTokensConfigRequest) (*types.ChainTokensConfigResponse, error) {
	return nil, nil
}

func (k Keeper) GetFee(c context.Context, req *types.GetFeeRequest) (*types.GetFeeResponse, error) {
	return nil, nil
}

func (k Keeper) GetFeePercentage(c context.Context, req *types.GetFeePercentageRequest) (*types.GetFeePercentageResponse, error) {
	return nil, nil
}

func (k Keeper) QueryTransferStatus(c context.Context, req *types.QueryTransferStatusRequest) (*types.QueryTransferStatusResponse, error) {
	return nil, nil
}

func (k Keeper) LiquidityDetailList(c context.Context, req *types.LiquidityDetailListRequest) (*types.LiquidityDetailListResponse, error) {
	return nil, nil
}

func (k Keeper) QueryTotalLiquidity(c context.Context, req *types.QueryTotalLiquidityRequest) (*types.QueryTotalLiquidityResponse, error) {
	return nil, nil
}

func (k Keeper) QueryAddLiquidityStatus(c context.Context, req *types.QueryAddLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	return nil, nil
}

func (k Keeper) QueryWithdrawLiquidityStatus(c context.Context, req *types.QueryWithdrawLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	return nil, nil
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
