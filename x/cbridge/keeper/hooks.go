package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var _ stakingtypes.StakingHooks = Hooks{}

func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (h Hooks) AfterValidatorBonded(ctx sdk.Context, valAddr eth.Addr) {
	h.k.UpdateLatestSigners(ctx, true)
}

func (h Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, valAddr eth.Addr) {
	h.k.UpdateLatestSigners(ctx, true)
}

func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	h.k.UpdateLatestSigners(ctx, false)
}

func (h Hooks) AfterValidatorCreated(_ sdk.Context, _ eth.Addr)                {}
func (h Hooks) AfterValidatorRemoved(_ sdk.Context, valAddr eth.Addr)          {}
func (h Hooks) BeforeDelegationCreated(_ sdk.Context, _ eth.Addr, _ eth.Addr)  {}
func (h Hooks) BeforeDelegationModified(_ sdk.Context, _ eth.Addr, _ eth.Addr) {}
