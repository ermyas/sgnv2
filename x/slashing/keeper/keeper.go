package keeper

import (
	"encoding/binary"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/slashing/types"
	stakingkeeper "github.com/celer-network/sgn-v2/x/staking/keeper"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/tendermint/tendermint/crypto"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey      sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc           codec.BinaryCodec
	StakingKeeper stakingkeeper.Keeper
	paramstore    paramtypes.Subspace
}

// NewKeeper creates new instances of the slash Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.BinaryCodec, stakingKeeper stakingkeeper.Keeper, paramstore paramtypes.Subspace) Keeper {
	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		StakingKeeper: stakingKeeper,
		paramstore:    paramstore,
	}
}

// HandleDoubleSign handles a validator signing two blocks at the same height.
func (k Keeper) HandleDoubleSign(ctx sdk.Context, addr crypto.Address, blkTime time.Time) {
	consAddr := sdk.ConsAddress(addr)
	validator, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		log.Errorf("Cannot find validator %s", consAddr)
		return
	}

	log.Infof("Confirmed double sign from %s %s", validator.EthAddress, consAddr)
	k.Slash(ctx, types.AttributeValueDoubleSign, validator, k.SlashFactorDoubleSign(ctx), nil, blkTime) //collector and syncer reward will be done in next version
}

// HandleValidatorSignature handles a validator signature, must be called once per validator per block.
func (k Keeper) HandleValidatorSignature(ctx sdk.Context, addr crypto.Address, signed bool, blkTime time.Time) {
	height := ctx.BlockHeight()
	consAddr := sdk.ConsAddress(addr)
	validator, found := k.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		log.Errorf("Cannot find validator %s", consAddr)
		return
	}

	signInfo, found := k.GetValidatorSigningInfo(ctx, consAddr)
	if !found {
		signInfo = slashingtypes.NewValidatorSigningInfo(
			consAddr,
			height,
			0,
			time.Unix(0, 0),
			false,
			0,
		)
	}

	// this is a relative index, so it counts blocks the validator *should* have signed
	// will use the 0-value default signing info if not present, except for start height
	signedBlocksWindow := k.SignedBlocksWindow(ctx)
	index := signInfo.IndexOffset % signedBlocksWindow
	signInfo.IndexOffset++

	// Update signed block bit array & counter
	// This counter just tracks the sum of the bit array
	// That way we avoid needing to read/write the whole array each time
	previous := k.GetValidatorMissedBlockBitArray(ctx, consAddr, index)
	missed := !signed
	switch {
	case !previous && missed:
		// Array value has changed from not missed to missed, increment counter
		k.SetValidatorMissedBlockBitArray(ctx, consAddr, index, true)
		signInfo.MissedBlocksCounter++
	case previous && !missed:
		// Array value has changed from missed to not missed, decrement counter
		k.SetValidatorMissedBlockBitArray(ctx, consAddr, index, false)
		signInfo.MissedBlocksCounter--
	default:
		// Array value at this index has not changed, no need to update counter
	}

	minHeight := signInfo.StartHeight + signedBlocksWindow
	maxMissed := signedBlocksWindow - k.MinSignedPerWindow(ctx).MulInt64(signedBlocksWindow).RoundInt64()

	// if we are past the minimum height and the validator has missed too many blocks, slash them
	if height > minHeight && signInfo.MissedBlocksCounter > maxMissed {
		// Downtime confirmed: slash the validator
		log.Infof("Validator %s %s past min height of %d and above max miss threshold of %d",
			validator.EthAddress, consAddr, minHeight, maxMissed)

		// We need to reset the counter & array so that the validator won't be immediately slashed for downtime upon rebonding.
		signInfo.MissedBlocksCounter = 0
		signInfo.IndexOffset = 0
		k.ClearValidatorMissedBlockBitArray(ctx, consAddr)
		k.Slash(ctx, types.AttributeValueMissingSignature, validator, k.SlashFactorDowntime(ctx), nil, blkTime) //collector and syncer reward will be done in next version
	}

	k.SetValidatorSigningInfo(ctx, signInfo)
}

// Slash a validator for an infraction
// Find the contributing stake and burn the specified slashFactor of it
func (k Keeper) Slash(
	ctx sdk.Context, reason string, failedValidator stakingtypes.Validator, slashFactor uint64,
	collectors []*types.AcctAmtPair, blkTime time.Time) {

	_, found := k.StakingKeeper.GetValidator(ctx, failedValidator.GetEthAddress())
	if !found {
		log.Errorln("cannot find profile for the failed validator, eth addr: ", failedValidator.EthAddress)
		return
	}

	enableSlash := k.EnableSlash(ctx)
	slashNonce := k.GetSlashNonce(ctx)
	slashExpireTime := uint64(blkTime.Unix()) + k.SlashTimeout(ctx)

	slash := types.NewSlash(
		slashNonce, failedValidator.GetEthAddr(), slashFactor, k.JailPeriod(ctx), slashExpireTime, reason, collectors)
	log.Warnf("Slash validator: %x, reason: %s, nonce: %d, enabled: %t",
		failedValidator.GetEthAddress(), reason, slash.SlashOnChain.Nonce, enableSlash)

	if enableSlash {
		slash.GenerateSlashBytes()
		k.SetSlash(ctx, slash)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSlash,
				sdk.NewAttribute(types.AttributeKeyNonce, sdk.NewUint(slash.SlashOnChain.Nonce).String()),
				sdk.NewAttribute(types.AttributeKeyReason, reason),
			),
		)

		slashNonce += 1
		k.SetSlashNonce(ctx, slashNonce)
	}
}

// Get the next Slash nonce
func (k Keeper) GetSlashNonce(ctx sdk.Context) (nonce uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SlashNonceKey)
	if bz != nil {
		nonce = binary.BigEndian.Uint64(bz)
	}

	return
}

// Set the slash nonce
func (k Keeper) SetSlashNonce(ctx sdk.Context, nonce uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SlashNonceKey, sdk.Uint64ToBigEndian(nonce))
}

// Get the entire Slash metadata for a nonce
func (k Keeper) GetSlash(ctx sdk.Context, nonce uint64) (slash types.Slash, found bool) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.GetSlashKey(nonce)) {
		return slash, false
	}

	value := store.Get(types.GetSlashKey(nonce))
	k.cdc.MustUnmarshal(value, &slash)
	return slash, true
}

// Set the entire slash metadata for a nonce
func (k Keeper) SetSlash(ctx sdk.Context, slash types.Slash) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetSlashKey(slash.SlashOnChain.Nonce), k.cdc.MustMarshal(&slash))
}

// IterateSlashes iterates over the stored slashes
func (k Keeper) IterateSlashes(ctx sdk.Context,
	handler func(slash types.Slash) (stop bool)) {

	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.SlashKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var slash types.Slash
		k.cdc.MustUnmarshal(iter.Value(), &slash)
		if handler(slash) {
			break
		}
	}
}

// GetSlashes returns all the slashes from store
func (keeper Keeper) GetSlashes(ctx sdk.Context) (slashes []types.Slash) {
	keeper.IterateSlashes(ctx, func(slash types.Slash) bool {
		slashes = append(slashes, slash)
		return false
	})
	return
}

// Stored by *validator consensus* address (not account address)
func (k Keeper) GetValidatorSigningInfo(ctx sdk.Context, address sdk.ConsAddress) (info slashingtypes.ValidatorSigningInfo, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(slashingtypes.ValidatorSigningInfoKey(address))
	if bz == nil {
		found = false
		return
	}

	k.cdc.MustUnmarshalLengthPrefixed(bz, &info)
	found = true
	return
}

// Stored by *validator consensus* address (not account address)
func (k Keeper) SetValidatorSigningInfo(ctx sdk.Context, info slashingtypes.ValidatorSigningInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalLengthPrefixed(&info)
	consAddr, _ := sdk.ConsAddressFromBech32(info.Address)
	store.Set(slashingtypes.ValidatorSigningInfoKey(consAddr), bz)
}

// Stored by *validator consensus* address (not account address)
func (k Keeper) GetValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress, index int64) (missed bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(slashingtypes.ValidatorMissedBlockBitArrayKey(address, index))
	if bz == nil {
		// lazy: treat empty key as not missed
		missed = false
		return
	}

	types.ModuleCdc.LegacyAmino.MustUnmarshalLengthPrefixed(bz, &missed)
	return
}

// Stored by *validator consensus* address (not account address)
func (k Keeper) SetValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress, index int64, missed bool) {
	store := ctx.KVStore(k.storeKey)
	bz := types.ModuleCdc.LegacyAmino.MustMarshalLengthPrefixed(missed)
	store.Set(slashingtypes.ValidatorMissedBlockBitArrayKey(address, index), bz)
}

// Stored by *validator consensus* address (not account address)
func (k Keeper) ClearValidatorMissedBlockBitArray(ctx sdk.Context, address sdk.ConsAddress) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, slashingtypes.ValidatorMissedBlockBitArrayPrefixKey(address))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}
