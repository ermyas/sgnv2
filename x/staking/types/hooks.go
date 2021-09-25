package types

import (
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// combine multiple staking hooks, all hook functions are run in array sequence
type MultiStakingHooks []StakingHooks

func NewMultiStakingHooks(hooks ...StakingHooks) MultiStakingHooks {
	return hooks
}

func (h MultiStakingHooks) AfterValidatorCreated(ctx sdk.Context, valAddr eth.Addr) {
	for i := range h {
		h[i].AfterValidatorCreated(ctx, valAddr)
	}
}
func (h MultiStakingHooks) AfterValidatorRemoved(ctx sdk.Context, valAddr eth.Addr) {
	for i := range h {
		h[i].AfterValidatorRemoved(ctx, valAddr)
	}
}
func (h MultiStakingHooks) AfterValidatorBonded(ctx sdk.Context, valAddr eth.Addr) {
	for i := range h {
		h[i].AfterValidatorBonded(ctx, valAddr)
	}
}
func (h MultiStakingHooks) AfterValidatorBeginUnbonding(ctx sdk.Context, valAddr eth.Addr) {
	for i := range h {
		h[i].AfterValidatorBeginUnbonding(ctx, valAddr)
	}
}
func (h MultiStakingHooks) BeforeDelegationCreated(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	for i := range h {
		h[i].BeforeDelegationCreated(ctx, delAddr, valAddr)
	}
}
func (h MultiStakingHooks) BeforeDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	for i := range h {
		h[i].BeforeDelegationModified(ctx, delAddr, valAddr)
	}
}
func (h MultiStakingHooks) AfterDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	for i := range h {
		h[i].AfterDelegationModified(ctx, delAddr, valAddr)
	}
}
