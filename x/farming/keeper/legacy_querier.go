package keeper

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	defaultPoolsDisplayedNum = 20
)

// NewLegacyQuerier creates a new legacy querier for farming clients.
func NewLegacyQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryParams:
			return queryParams(ctx, path[1:], req, k, legacyQuerierCdc)
		case types.QueryPool:
			return queryPool(ctx, req, k, legacyQuerierCdc)
		case types.QueryPools:
			return queryPools(ctx, req, k, legacyQuerierCdc)
		case types.QueryEarnings:
			return queryEarnings(ctx, req, k, legacyQuerierCdc)
		case types.QueryStakeInfo:
			return queryStakeInfo(ctx, req, k, legacyQuerierCdc)
		case types.QueryAccountInfo:
			return QueryAccountInfo(ctx, req, k, legacyQuerierCdc)
		case types.QueryAccountsStakedIn:
			return queryAccountsStakedIn(ctx, req, k, legacyQuerierCdc)
		case types.QueryNumPools:
			return queryNumPools(ctx, k, legacyQuerierCdc)
		default:
			return nil, types.WrapErrUnknownFarmingQueryType("failed. unknown farming query endpoint")
		}
	}
}

func queryParams(
	ctx sdk.Context, _ []string, _ abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryPool(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryPoolParams

	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	pool, found := k.GetFarmingPool(ctx, params.PoolName)
	if !found {
		return nil, types.WrapErrNoFarmingPoolFound(params.PoolName)
	}

	updatedPool, _ := k.CalculateAmountEarnedBetween(ctx, pool)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, updatedPool)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// support query by page && limit
func queryPools(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryPoolsParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	pools := k.GetFarmingPools(ctx)
	var updatedPools types.FarmingPools
	for _, pool := range pools {
		updatedPool, _ := k.CalculateAmountEarnedBetween(ctx, pool)
		updatedPools = append(updatedPools, updatedPool)
	}

	if !(params.Page == 1 && params.Limit == 0) {
		start, end := client.Paginate(len(updatedPools), params.Page, params.Limit, defaultPoolsDisplayedNum)
		if start < 0 || end < 0 {
			start, end = 0, 0
		}
		updatedPools = updatedPools[start:end]
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, updatedPools)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryEarnings(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryPoolAccountParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	earnings, sdkErr := k.GetEarnings(ctx, params.PoolName, params.Address)
	if sdkErr != nil {
		return nil, sdkErr
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, earnings)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryStakeInfo(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryPoolAccountParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	stakeInfo, found := k.GetStakeInfo(ctx, params.Address, params.PoolName)
	if !found {
		return nil, types.WrapErrNoStakeInfoFound(params.Address.String(), params.PoolName)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, stakeInfo)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func QueryAccountInfo(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryAccountParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// StakedPools
	poolNames := k.GetFarmingPoolNamesForAccount(ctx, params.Address)
	var updatedPools types.FarmingPools
	for _, poolName := range poolNames {
		pool, found := k.GetFarmingPool(ctx, poolName)
		if !found {
			return nil, types.WrapErrNoStakeInfoFound(params.Address.String(), poolName)
		}
		updatedPool, _ := k.CalculateAmountEarnedBetween(ctx, pool)
		updatedPools = append(updatedPools, updatedPool)
	}

	// EarningsList
	var earningsList []types.Earnings
	for _, poolName := range poolNames {
		earnings, sdkErr := k.GetEarnings(ctx, poolName, params.Address)
		if sdkErr != nil {
			return nil, sdkErr
		}
		earningsList = append(earningsList, earnings)
	}

	// CumulativeRewards
	derivedRewardAccount := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, params.Address)
	cumulativeRewards := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, derivedRewardAccount)...)

	accountInfo := types.AccountInfo{
		StakedPools:             updatedPools,
		EarningsList:            earningsList,
		CumulativeRewardAmounts: cumulativeRewards,
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, accountInfo)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryAccountsStakedIn(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryPoolParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	addrList := k.GetAccountsStakedIn(ctx, params.PoolName)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, addrList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryNumPools(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	numPools := k.GetNumPools(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, numPools)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	return res, nil
}
