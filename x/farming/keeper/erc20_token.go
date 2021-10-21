package keeper

import (
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetERC20Token(ctx sdk.Context, token types.ERC20Token) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetERC20TokenKey(token.ChainId, token.Symbol), k.cdc.MustMarshal(&token))
}

func (k Keeper) GetERC20Token(ctx sdk.Context, chainId uint64, symbol string) (token types.ERC20Token, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetERC20TokenKey(chainId, symbol))
	if bz == nil {
		return token, false
	}
	k.cdc.MustUnmarshal(bz, &token)
	return token, true
}

// GetERC20Tokens gets all supported tokens
func (k Keeper) GetERC20Tokens(ctx sdk.Context) (tokens types.ERC20Tokens) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ERC20TokenPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var token types.ERC20Token
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		tokens = append(tokens, token)
	}

	return
}

func (k Keeper) HasERC20Token(ctx sdk.Context, chainId uint64, symbol string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetERC20TokenKey(chainId, symbol))
}

func (k Keeper) DeleteERC20Token(ctx sdk.Context, chainId uint64, symbol string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetERC20TokenKey(chainId, symbol))
}
