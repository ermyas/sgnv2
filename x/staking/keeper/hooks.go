package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Implements StakingHooks interface
var _ types.StakingHooks = Keeper{}

// AfterValidatorCreated - call hook if registered
func (k Keeper) AfterValidatorCreated(ctx sdk.Context, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.AfterValidatorCreated(ctx, valAddr)
	}
}

// AfterValidatorRemoved - call hook if registered
func (k Keeper) AfterValidatorRemoved(ctx sdk.Context, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.AfterValidatorRemoved(ctx, valAddr)
	}
}

// AfterValidatorBonded - call hook if registered
func (k Keeper) AfterValidatorBonded(ctx sdk.Context, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.AfterValidatorBonded(ctx, valAddr)
	}
}

// AfterValidatorBeginUnbonding - call hook if registered
func (k Keeper) AfterValidatorBeginUnbonding(ctx sdk.Context, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.AfterValidatorBeginUnbonding(ctx, valAddr)
	}
}

// AfterValidatorPowerUpdated - call hook if registered
func (k Keeper) AfterValidatorPowerUpdated(ctx sdk.Context, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.AfterValidatorPowerUpdated(ctx, valAddr)
	}
}

// BeforeDelegationCreated - call hook if registered
func (k Keeper) BeforeDelegationCreated(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.BeforeDelegationCreated(ctx, delAddr, valAddr)
	}
}

// BeforeDelegationModified - call hook if registered
func (k Keeper) BeforeDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.BeforeDelegationModified(ctx, delAddr, valAddr)
	}
}

// AfterDelegationModified - call hook if registered
func (k Keeper) AfterDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	if k.hooks != nil {
		k.hooks.AfterDelegationModified(ctx, delAddr, valAddr)
	}
}
