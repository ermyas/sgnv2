package keeper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ERC20DenomSeparator = "@"
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

func (k Keeper) HasERC20Token(ctx sdk.Context, chainId uint64, symbol string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetERC20TokenKey(chainId, symbol))
}

func (k Keeper) DeleteERC20Token(ctx sdk.Context, chainId uint64, symbol string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetERC20TokenKey(chainId, symbol))
}

// DeriveERC20TokenDenom generates denoms of the form symbol@chainId
func DeriveERC20TokenDenom(chainId uint64, symbol string) string {
	return fmt.Sprintf("%s%s%d", symbol, ERC20DenomSeparator, chainId)
}

func ParseERC20TokenDenom(denom string) (chainId uint64, symbol string, err error) {
	splitted := strings.Split(denom, ERC20DenomSeparator)
	if len(splitted) != 2 {
		return 0, "", fmt.Errorf("invalid denom %s", denom)
	}
	chainIdInt64, err := strconv.ParseInt(splitted[1], 10, 64)
	if err != nil {
		return 0, "", err
	}
	return uint64(chainIdInt64), splitted[0], nil
}
