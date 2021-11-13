package keeper

import (
	"errors"
	"fmt"

	"github.com/celer-network/sgn-v2/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case types.QuerySlash:
			return querySlash(ctx, req, keeper, legacyQuerierCdc)
		case types.QuerySlashes:
			return querySlashes(ctx, req, keeper, legacyQuerierCdc)
		case types.QueryParameters:
			return queryParameters(ctx, keeper, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Unknown slash query endpoint")
		}
	}
}

func querySlash(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QuerySlashParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to parse params: %s", err)
	}

	slash, found := keeper.GetSlash(ctx, params.Nonce)
	if !found {
		return nil, errors.New("Slash does not exist")
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, slash)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())

	}

	return res, nil
}

func querySlashes(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	slashes := keeper.GetSlashes(ctx)
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, slashes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())

	}

	return res, nil
}

func queryParameters(ctx sdk.Context, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
