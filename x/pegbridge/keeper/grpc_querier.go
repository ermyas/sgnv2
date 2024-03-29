package keeper

import (
	"context"
	"errors"
	"math/big"

	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

// Params queries params of pegbridge module
func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	var params types.Params
	k.paramSpace.GetParamSet(ctx, &params)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) Config(c context.Context, req *types.QueryConfigRequest) (*types.PegConfig, error) {
	ctx := sdk.UnwrapSDKContext(c)
	config := k.GetConfig(ctx)
	return &config, nil
}

func (k Keeper) OrigPeggedPairs(
	c context.Context, req *types.QueryOrigPeggedPairsRequest) (*types.QueryOrigPeggedPairsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	pairs := make([]types.OrigPeggedPair, 0)

	orig := req.GetOrig()
	pegged := req.GetPegged()
	if orig != nil {
		// If pegged chain ID specified, return single pair
		if pegged.GetChainId() != 0 {
			pair, found := k.GetOrigPeggedPair(ctx, orig.ChainId, orig.Address, pegged.ChainId)
			if found {
				pairs = append(pairs, pair)
			}
			return &types.QueryOrigPeggedPairsResponse{Pairs: pairs}, nil
		}
		// If pegged chain ID not specified, return all pairs from origin
		k.IterateOrigPeggedPairsByOrig(ctx, orig.GetChainId(), orig.GetAddress(), func(pair types.OrigPeggedPair) bool {
			pairs = append(pairs, pair)
			return false
		})
		return &types.QueryOrigPeggedPairsResponse{Pairs: pairs}, nil
	}

	// If orig not specified but pegged specified, return single pair
	if pegged != nil {
		var pegAddr interface{}
		if commontypes.IsFlowChain(pegged.GetChainId()) {
			pegAddr = pegged.GetAddress() // use string directly, as token string has non-hex char
		} else {
			// evm
			pegAddr = eth.Hex2Addr(pegged.GetAddress())
		}
		pair, found := k.GetOrigPeggedPairByPegged(ctx, pegged.GetChainId(), pegAddr)
		if found {
			pairs = append(pairs, pair)
		}
		return &types.QueryOrigPeggedPairsResponse{Pairs: pairs}, nil
	}

	// Else, return all pairs
	k.IterateAllOrigPeggedPairs(ctx, func(pair types.OrigPeggedPair) bool {
		pairs = append(pairs, pair)
		return false
	})
	return &types.QueryOrigPeggedPairsResponse{Pairs: pairs}, nil
}

func (k Keeper) EstimatedAmountFees(
	c context.Context, req *types.QueryEstimatedAmountFeesRequest) (*types.QueryEstimatedAmountFeesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	requestAmount, valid := new(big.Int).SetString(req.RequestAmount, 10)
	if !valid {
		return nil, errors.New("invalid request amount")
	}
	// Use stored pair info to estimate fees
	reqPair := req.GetPair()
	pair, found := k.GetOrigPeggedPair(ctx, reqPair.Orig.ChainId, reqPair.Orig.Address, reqPair.Pegged.ChainId)
	if !found {
		return nil, errors.New("invalid pegged pair")
	}
	receiveAmount, baseFee, percFee := k.CalcAmountAndFees(ctx, pair, requestAmount, req.Mint)
	return &types.QueryEstimatedAmountFeesResponse{
		ReceiveAmount: receiveAmount.String(),
		BaseFee:       baseFee.String(),
		PercentageFee: percFee.String(),
	}, nil
}

func (k Keeper) DepositInfo(c context.Context, req *types.QueryDepositInfoRequest) (*types.QueryDepositInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	depositId := eth.Hex2Hash(req.DepositId)
	info, found := k.GetDepositInfo(ctx, depositId)
	if !found {
		return nil, types.WrapErrNoInfoFound(depositId)
	}
	return &types.QueryDepositInfoResponse{DepositInfo: info}, nil
}

func (k Keeper) WithdrawInfo(c context.Context, req *types.QueryWithdrawInfoRequest) (*types.QueryWithdrawInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	withdrawId := eth.Hex2Hash(req.WithdrawId)
	info, found := k.GetWithdrawInfo(ctx, withdrawId)
	if !found {
		return nil, types.WrapErrNoInfoFound(withdrawId)
	}
	return &types.QueryWithdrawInfoResponse{WithdrawInfo: info}, nil
}

func (k Keeper) MintInfo(c context.Context, req *types.QueryMintInfoRequest) (*types.QueryMintInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	mintId := eth.Hex2Hash(req.MintId)
	info, found := k.GetMintInfo(ctx, mintId)
	if !found {
		return nil, types.WrapErrNoInfoFound(mintId)
	}
	return &types.QueryMintInfoResponse{MintInfo: info}, nil
}

func (k Keeper) BurnInfo(c context.Context, req *types.QueryBurnInfoRequest) (*types.QueryBurnInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	burnId := eth.Hex2Hash(req.BurnId)
	info, found := k.GetBurnInfo(ctx, burnId)
	if !found {
		return nil, types.WrapErrNoInfoFound(burnId)
	}
	return &types.QueryBurnInfoResponse{BurnInfo: info}, nil
}

func (k Keeper) FeeClaimInfo(c context.Context, req *types.QueryFeeClaimInfoRequest) (*types.QueryFeeClaimInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	info, found := k.GetFeeClaimInfo(ctx, eth.Hex2Addr(req.Address), req.Nonce)
	if !found {
		return nil, errors.New("fee claim info not found")
	}
	return &types.QueryFeeClaimInfoResponse{FeeClaimInfo: info}, nil
}

func (k Keeper) SupplyInfo(c context.Context, req *types.QuerySupplyInfoRequest) (*types.QuerySupplyInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	pair, found := k.GetOrigPeggedPairByPegged(ctx, req.PeggedChainId, eth.Hex2Addr(req.PeggedAddress))
	if !found {
		return nil, errors.New("orig-pegged pair not found")
	}
	total, found := k.GetTotalSupply(ctx, req.PeggedChainId, eth.Hex2Addr(req.PeggedAddress))
	if !found {
		total = new(big.Int).SetInt64(0)
	}
	return &types.QuerySupplyInfoResponse{Total: total.String(), Cap: pair.SupplyCap}, nil
}

func (k Keeper) RefundClaimInfo(c context.Context, req *types.QueryRefundClaimInfoRequest) (*types.QueryRefundClaimInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	info, found := k.GetRefundClaimInfo(ctx, eth.Hex2Hash(req.DepositId))
	if !found {
		return nil, errors.New("refund claim info not found")
	}
	return &types.QueryRefundClaimInfoResponse{WithdrawId: info.Hex()}, nil
}
