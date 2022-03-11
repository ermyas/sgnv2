package keeper

import (
	"fmt"
	"runtime/debug"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ApplyUpdate(ctx sdk.Context, update *types.PendingUpdate) (applied bool) {
	// Gracefully handle any panic when applying updates
	defer func() {
		if r := recover(); r != nil {
			applied = false
			log.Errorf("panic when applying update %d: %s, %s", update.Id, r, string(debug.Stack()))
		}
	}()

	cacheCtx, writeCache := ctx.CacheContext()
	var err error
	switch update.Type {
	case types.DataType_ValidatorSgnAddr:
		applied, err = k.applyValidatorSgnAddr(cacheCtx, update)
	case types.DataType_ValidatorParams:
		applied, err = k.applyValidatorParams(cacheCtx, update)
	case types.DataType_ValidatorStates:
		applied, err = k.applyValidatorStates(cacheCtx, update)
	case types.DataType_DelegatorShares:
		applied, err = k.applyDelegatorShares(cacheCtx, update)
	case types.DataType_CbrOnchainEvent:
		applied, err = k.cbrKeeper.ApplyEvent(cacheCtx, update.Data)
	case types.DataType_CbrUpdateCbrPrice:
		applied, err = k.cbrKeeper.ApplyUpdateCbrPrice(cacheCtx, update.Data)
	case types.DataType_PegbrOnChainEvent:
		applied, err = k.pegbrKeeper.ApplyEvent(cacheCtx, update.Data)
	case types.DataType_MsgbrOnChainEvent:
		applied, err = k.msgbrKeeper.ApplyEvent(cacheCtx, update.Data)
	}

	if err != nil {
		log.Errorln("Apply update err:", err)
	}

	if applied {
		// The cached context is created with a new EventManager. However, since
		// the application was successful, we want to track/keep
		// any events emitted, so we re-emit to "merge" the events into the
		// original Context's EventManager.
		ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())

		// write state to the underlying multi-store
		writeCache()
	}

	return applied
}

func (k Keeper) applyValidatorSgnAddr(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator sgn addr %s", v.String())
	// TODO: remove previous sgnaddr account
	acct, err := sdk.AccAddressFromBech32(v.SgnAddress)
	if err != nil {
		return false, err
	}
	err = k.stakingKeeper.InitAccount(ctx, acct)
	if err != nil {
		return false, err
	}
	return true, nil
}

// TODO: allow sgnaddr/consaddr update
func (k Keeper) applyValidatorParams(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	if update.Proposer != v.SgnAddress {
		return false, fmt.Errorf("validator %s not msg sender: %s", v.EthAddress, update.Proposer)
	}
	if v.ConsensusPubkey == nil {
		return false, fmt.Errorf("empty consensus pub key")
	}
	consAddr, err := v.GetConsAddr()
	if err != nil {
		return false, fmt.Errorf("validator %s failed to get consensus addr, err %s", v.EthAddress, err)
	}
	log.Infof("Apply validator params %s", v.String())
	var prevSigner string
	val, found := k.stakingKeeper.GetValidator(ctx, eth.Hex2Addr(v.EthAddress))
	if !found {
		val = *stakingtypes.NewValidator(v.EthAddress, v.EthSigner, v.SgnAddress)
		val.Description = v.Description
	} else {
		if v.SgnAddress != val.SgnAddress {
			return false, fmt.Errorf("update of sgnaddr is not supported: %s %s %s", v.EthAddress, v.SgnAddress, val.SgnAddress)
		}
		storedConsAddr, err := val.GetConsAddr()
		if err != nil {
			return false, fmt.Errorf("validator %s failed to get stored consensus addr, err %s", v.EthAddress, err)
		}
		if !consAddr.Equals(storedConsAddr) {
			return false, fmt.Errorf("update of consaddr is not supported: %s %s %s", v.EthAddress, consAddr, storedConsAddr)
		}
		if val.Status == stakingtypes.Bonded && val.EthSigner != eth.FormatAddrHex(v.EthSigner) {
			prevSigner = val.EthSigner
		}
		val.EthSigner = eth.FormatAddrHex(v.EthSigner)
		val.SgnAddress = v.SgnAddress
	}
	val.ConsensusPubkey = v.ConsensusPubkey
	val.CommissionRate = v.CommissionRate
	k.stakingKeeper.SetValidatorParams(ctx, &val, !found)
	if prevSigner != "" {
		log.Infof("Update bonded validator %s signer from %s to %s", val.EthAddress, prevSigner, val.EthSigner)
		k.cbrKeeper.UpdateLatestSigners(ctx, true)
	}
	// TODO: gas coins
	return true, nil
}

func (k Keeper) applyValidatorStates(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	v, err := stakingtypes.UnmarshalValidator(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply validator states %s", v.String())
	val, found := k.stakingKeeper.GetValidator(ctx, eth.Hex2Addr(v.EthAddress))
	if !found {
		return false, fmt.Errorf("validator %s not found", val.EthAddress)
	}
	val.Status = v.Status
	val.Tokens = v.Tokens
	val.DelegatorShares = v.DelegatorShares

	k.stakingKeeper.SetValidatorStates(ctx, &val)
	return true, nil
}

func (k Keeper) applyDelegatorShares(ctx sdk.Context, update *types.PendingUpdate) (bool, error) {
	d, err := stakingtypes.UnmarshalDelegation(k.cdc, update.Data)
	if err != nil {
		return false, err
	}
	log.Infof("Apply delegator shares, valAddr %s delAddr %s shares %s", d.ValidatorAddress, d.DelegatorAddress, d.Shares)
	valAddr := eth.Hex2Addr(d.ValidatorAddress)
	val, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if found && !val.DelegatorShares.Equal(sdk.ZeroInt()) {
		log.Infof("Apply delegator shares, validator found with DelegatorShares %s", val.DelegatorShares)
		err = k.stakingKeeper.SetDelegationShares(
			ctx,
			eth.Hex2Addr(d.DelegatorAddress),
			eth.Hex2Addr(d.ValidatorAddress),
			d.Shares)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	if !found {
		log.Info("Apply delegator shares, validator not found, deferring")
	} else {
		log.Infof("Apply delegator shares, validator %s has zero DelegatorShares, deferring", valAddr)
	}
	// If validator not found or its DelegatorShares is zero, the delegation is processed before we have done the initial sync of the
	// validator, just try again later.
	return false, nil
}
