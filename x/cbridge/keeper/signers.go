package keeper

import (
	"bytes"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetChainSigners(ctx sdk.Context, chainId uint64) (signers types.ChainSigners, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetChainSignersKey(chainId))
	if value == nil {
		return signers, false
	}
	signers = types.MustUnmarshalChainSigners(k.cdc, value)
	return signers, true
}

func (k Keeper) SetChainSigners(ctx sdk.Context, s *types.ChainSigners) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetChainSignersKey(s.ChainId), types.MustMarshalChainSigners(k.cdc, s))
}

func (k Keeper) GetLatestSigners(ctx sdk.Context) (signers types.LatestSigners, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.LatestSignersKey)
	if value == nil {
		return signers, false
	}
	signers = types.MustUnmarshalLatestSigners(k.cdc, value)
	return signers, true
}

func (k Keeper) SetLatestSigners(ctx sdk.Context, s *types.LatestSigners) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LatestSignersKey, types.MustMarshalLatestSigners(k.cdc, s))
}

func (k Keeper) UpdateLatestSigners(ctx sdk.Context, force bool) {
	latestSigners, found := k.GetLatestSigners(ctx)
	if found && !force {
		duraion := k.GetSignerUpdateDuraion(ctx)
		if latestSigners.GetUpdateTime().Add(duraion).Before(ctx.BlockHeader().Time) {
			return
		}
	}

	vals := k.stakingKeeper.GetBondedValidators(ctx)
	newSigners := &types.LatestSigners{}
	for _, v := range vals {
		signer := &types.AddrAmt{
			Addr: v.GetEthAddr().Bytes(),
			Amt:  v.Tokens.BigInt().Bytes(),
		}
		newSigners.Signers.Signers = append(newSigners.Signers.Signers, signer)
	}
	newSigners.Signers.Sort()
	newSigners.GenerateSignersBytes()

	if bytes.Compare(latestSigners.GetSignersBytes(), newSigners.SignersBytes) == 0 {
		return
	}

	log.Infoln("Update latest signers:", newSigners.Signers.String())
	newSigners.UpdateTime = ctx.BlockHeader().Time
	k.SetLatestSigners(ctx, newSigners)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventToSign,
		sdk.NewAttribute(types.EvAttrType, types.SignDataType_SIGNERS.String()),
		sdk.NewAttribute(types.EvAttrData, string(newSigners.SignersBytes))))
}
