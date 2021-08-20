package validator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QuerySyncer:
			return querySyncer(ctx, req, keeper)
		case QueryDelegator:
			return queryDelegator(ctx, req, keeper)
		case QueryValidator:
			return queryValidator(ctx, req, keeper)
		case QueryValidators:
			return queryValidators(ctx, req, keeper)
		case QueryDelegators:
			return queryDelegators(ctx, req, keeper)
		case QueryParameters:
			return queryParameters(ctx, keeper)
		default:
			return nil, sdk_errors.Wrap(sdk_errors.ErrUnknownRequest, "Unknown validator query endpoint")
		}
	}
}

func querySyncer(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	return nil, nil
}

func queryDelegator(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	return nil, nil
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	return nil, nil
}

func queryValidators(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	return nil, nil
}

func queryDelegators(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	return nil, nil
}

func queryParameters(ctx sdk.Context, k Keeper) ([]byte, error) {
	return nil, nil
}
