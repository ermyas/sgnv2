package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStakingRewardClaimInfo sets a staking reward claim info
func (k Keeper) SetStakingRewardClaimInfo(ctx sdk.Context, info types.StakingRewardClaimInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.GetStakingRewardClaimInfoKey(eth.Hex2Addr(info.Recipient)),
		k.cdc.MustMarshal(&info))
}

// GetStakingRewardClaimInfo gets the staking reward claim info for an Ethereum address
func (k Keeper) GetStakingRewardClaimInfo(ctx sdk.Context, addr eth.Addr) (
	info types.StakingRewardClaimInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	rewardKey := types.GetStakingRewardClaimInfoKey(addr)
	bz := store.Get(rewardKey)
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

// HasStakingRewardClaimInfo returns whether a staking reward claim info exists for an Ethereum address
func (k Keeper) HasStakingRewardClaimInfo(ctx sdk.Context, addr eth.Addr) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetStakingRewardClaimInfoKey(addr))
}
