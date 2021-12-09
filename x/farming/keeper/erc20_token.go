package keeper

import (
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetERC20Token(ctx sdk.Context, token commontypes.ERC20Token) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetERC20TokenKey(token.ChainId, token.Symbol), k.cdc.MustMarshal(&token))
}

func (k Keeper) GetERC20Token(ctx sdk.Context, chainId uint64, symbol string) (token commontypes.ERC20Token, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetERC20TokenKey(chainId, symbol))
	if bz == nil {
		return token, false
	}
	k.cdc.MustUnmarshal(bz, &token)
	return token, true
}

func (k Keeper) IterateAllERC20Tokens(
	ctx sdk.Context, handler func(token commontypes.ERC20Token) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ERC20TokenPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var token commontypes.ERC20Token
		k.cdc.MustUnmarshal(iter.Value(), &token)
		if handler(token) {
			break
		}
	}
}

// GetERC20Tokens gets all supported tokens
func (k Keeper) GetERC20Tokens(ctx sdk.Context) (tokens commontypes.ERC20Tokens) {
	k.IterateAllERC20Tokens(ctx, func(token commontypes.ERC20Token) bool {
		tokens = append(tokens, token)
		return false
	})
	return tokens
}

func (k Keeper) HasERC20Token(ctx sdk.Context, chainId uint64, symbol string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetERC20TokenKey(chainId, symbol))
}

func (k Keeper) DeleteERC20Token(ctx sdk.Context, chainId uint64, symbol string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetERC20TokenKey(chainId, symbol))
}
