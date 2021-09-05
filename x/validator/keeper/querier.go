package keeper

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryValidator:
			return queryValidator(ctx, req, k, legacyQuerierCdc)
		case types.QueryValidators:
			return queryValidators(ctx, k, legacyQuerierCdc)
		case types.QueryDelegator:
			return queryDelegator(ctx, req, k, legacyQuerierCdc)
		case types.QueryDelegators:
			return queryDelegators(ctx, req, k, legacyQuerierCdc)
		case types.QuerySgnAccount:
			return querySgnAccountExist(ctx, req, k, legacyQuerierCdc)
		case types.QuerySyncer:
			return querySyncer(ctx, k, legacyQuerierCdc)
		case types.QueryParams:
			return queryParams(ctx, k, legacyQuerierCdc)
		default:
			return nil, sdk_errors.Wrap(sdk_errors.ErrUnknownRequest, "Unknown validator query endpoint")
		}
	}
}

func queryValidator(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	}

	validator, found := k.GetValidator(ctx, params.EthAddress)
	if !found {
		return nil, types.ErrValidatorNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validator)
	if err != nil {
		log.Error(err)
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryValidators(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	validators := k.GetAllValidators(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validators)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryDelegator(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryDelegatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	}

	delegator, found := k.GetDelegator(ctx, params.ValAddress, params.DelAddress)
	if !found {
		return nil, types.ErrDelegatorNotFound
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, delegator)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// query all delegators of a given validator
func queryDelegators(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryValidatorParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	}

	delegators := k.GetAllDelegators(ctx, params.EthAddress)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, delegators)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func querySgnAccountExist(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QuerySgnAccountParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONUnmarshal, err.Error())
	}

	acctAddr, err := types.SdkAccAddrFromSgnBech32(params.SgnAddress)
	if err != nil {
		return nil, sdk_errors.Wrapf(types.ErrInvalidAddress, "%s", params.SgnAddress)
	}
	account := k.sdkAccountKeeper.GetAccount(ctx, acctAddr)
	if account == nil {
		return nil, sdk_errors.Wrapf(types.ErrSgnAccounNotFound, "%s", params.SgnAddress)
	}

	return []byte{1}, nil
}

func querySyncer(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	syncer := k.GetSyncer(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, syncer)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryParams(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdk_errors.Wrap(sdk_errors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
