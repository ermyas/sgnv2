package keeper

import (
	"bytes"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
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

	// when signers changed, remember to update gas cost
	gasCostParam := GetRelayGasCostParam(store, s.ChainId)
	var gasCost uint32
	if gasCostParam == nil {
		gasCost = 0
	}
	gasCost = gasCostParam.GetCostBase() + gasCostParam.GetPerValidator()*uint32(len(s.GetSortedSigners())) +
		gasCostParam.GetPerSig()*types.MinSigsForQuorum(s.GetSortedSigners())
	setUint32(store, types.CfgKeyChain2EstimateRelayGasCost(s.ChainId), gasCost)
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

func (k Keeper) UpdateLatestSigners(ctx sdk.Context, force bool) bool {
	latestSigners, found := k.GetLatestSigners(ctx)
	if found && !force {
		duration := k.GetSignerUpdateDuration(ctx)
		if latestSigners.GetUpdateTime().Add(duration).After(ctx.BlockHeader().Time) {
			return false
		}
	}

	vals := k.stakingKeeper.GetBondedValidators(ctx)
	newSigners := &types.LatestSigners{}
	for _, v := range vals {
		signer := &types.Signer{
			Addr:  v.GetSignerAddr().Bytes(),
			Power: v.Tokens.BigInt().Bytes(),
		}
		newSigners.SortedSigners = append(newSigners.SortedSigners, signer)
	}
	newSigners.Sort()
	newSigners.GenerateSignersBytes()

	if bytes.Equal(latestSigners.GetSignersBytes(), newSigners.SignersBytes) {
		return false
	}

	log.Infoln("Update latest signers:", newSigners.String())
	newSigners.UpdateTime = ctx.BlockHeader().Time
	k.SetLatestSigners(ctx, newSigners)
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDataToSign,
		sdk.NewAttribute(types.AttributeKeyType, types.SignDataType_SIGNERS.String()),
		sdk.NewAttribute(types.AttributeKeyData, eth.Bytes2Hex(newSigners.SignersBytes)),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	return true
}
