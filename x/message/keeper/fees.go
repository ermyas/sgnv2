package keeper

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintFee(ctx sdk.Context, chainId uint64, amount *big.Int) error {
	messageBus, found := k.GetMessageBus(ctx, chainId)
	if !found {
		return errors.New("MessageBus not found")
	}
	denom := fmt.Sprintf("%s%s/%d", types.MessageFeeDenomPrefix, messageBus.FeeTokenSymbol, chainId)
	coin := sdk.NewCoin(denom, sdk.NewIntFromBigInt(amount))
	if err := k.bankKeeper.MintCoins(ctx, k.feeCollectorName, sdk.NewCoins(coin)); err != nil {
		return err
	}
	log.Debugf("minted message fee %s, chainId %d", amount.String(), chainId)
	return nil
}

// SetFeeClaimInfo sets a fee claim info.
func (k Keeper) SetFeeClaimInfo(ctx sdk.Context, info types.FeeClaimInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.GetFeeClaimInfoKey(eth.Hex2Addr(info.Recipient)),
		k.cdc.MustMarshal(&info))
}

// GetFeeClaimInfo gets the fee claim info for an Ethereum address.
func (k Keeper) GetFeeClaimInfo(ctx sdk.Context, addr eth.Addr) (
	info types.FeeClaimInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	rewardKey := types.GetFeeClaimInfoKey(addr)
	bz := store.Get(rewardKey)
	if bz == nil {
		return info, false
	}
	k.cdc.MustUnmarshal(bz, &info)
	return info, true
}

// HasFeeClaimInfo returns whether a fee claim info exists for an Ethereum address.
func (k Keeper) HasFeeClaimInfo(ctx sdk.Context, addr eth.Addr) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeClaimInfoKey(addr))
}
