package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetRewardClaimInfo sets a reward claim info
func (k Keeper) SetRewardClaimInfo(ctx sdk.Context, info types.RewardClaimInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.GetRewardClaimInfoKey(eth.Hex2Addr(info.Recipient)),
		k.cdc.MustMarshal(&info))
}

// GetRewardClaimInfo gets the reward claim info for an Ethereum address
func (k Keeper) GetRewardClaimInfo(ctx sdk.Context, addr eth.Addr) (
	info types.RewardClaimInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	rewardKey := types.GetRewardClaimInfoKey(addr)
	bz := store.Get(rewardKey)
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

// HasRewardClaimInfo returns whether a reward claim info exists for an Ethereum address
func (k Keeper) HasRewardClaimInfo(ctx sdk.Context, addr eth.Addr) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetRewardClaimInfoKey(addr))
}
