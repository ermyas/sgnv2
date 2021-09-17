package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_params "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter namespace
const (
	DefaultParamspace = types.ModuleName
)

// ParamTable for validator module
func ParamKeyTable() sdk_params.KeyTable {
	return sdk_params.NewKeyTable().RegisterParamSet(&types.Params{})
}

// MaxValidatorDiff - max validator add
func (k Keeper) MultiChainAsset(ctx sdk.Context) (res types.MultiChainAssetParam) {
	k.paramstore.Get(ctx, types.KeyMultiChainAsset, &res)
	return
	/*return types.MultiChainAssetParam{
		ChainAsset: []types.ChainAsset{
			{
				ChainId: 1,
				TokenAddr: "xx",
				TokenSymbol: "yy",
				Decimal: 18,
				MaxFeeAmount: "12123",
			},
		},
	}*/
}

func (k Keeper) PowerReduction(ctx sdk.Context) sdk.Int {
	return sdk.DefaultPowerReduction
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(k.MultiChainAsset(ctx))
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// utils to deal with asset, chid and address

// given chid and token address, return which asset eg. USDT
// empty string if not found
func (k Keeper) GetAssetSymbol(chaddr *ChainIdTokenAddr) string {
	return ""
}

// given asset symbol, return token address for chid, zero address if not found
func (k Keeper) GetTokenAddr(sym string, chid uint64) eth.Addr {
	return eth.ZeroAddr // TODO: impl
}
