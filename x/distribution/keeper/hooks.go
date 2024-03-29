package keeper

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var _ stakingtypes.StakingHooks = Hooks{}

// Create new distribution hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

// initialize validator distribution record
func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr eth.Addr) {
	val := h.k.stakingKeeper.Validator(ctx, valAddr)
	h.k.initializeValidator(ctx, val)
}

// AfterValidatorRemoved performs clean up after a validator is removed
func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, valAddr eth.Addr) {
	// fetch outstanding
	outstanding := h.k.GetValidatorOutstandingRewardsCoins(ctx, valAddr)

	// force-withdraw commission
	commission := h.k.GetValidatorAccumulatedCommission(ctx, valAddr).Commission
	if !commission.IsZero() {
		// subtract from outstanding
		outstanding = outstanding.Sub(commission)

		// split into integral & remainder
		coins, remainder := commission.TruncateDecimal()

		// remainder to community pool
		feePool := h.k.GetFeePool(ctx)
		feePool.CommunityPool = feePool.CommunityPool.Add(remainder...)
		h.k.SetFeePool(ctx, feePool)

		// add to validator account
		if !coins.IsZero() {
			validator := h.k.stakingKeeper.Validator(ctx, valAddr)
			withdrawAddr := h.k.GetDelegatorWithdrawAddr(ctx, validator.GetEthAddr())
			derivedWithdrawAddr := common.DeriveSdkAccAddressFromEthAddress(types.ModuleName, withdrawAddr)

			if err := h.k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, derivedWithdrawAddr, coins); err != nil {
				panic(err)
			}
		}
	}

	// Add outstanding to community pool
	// The validator is removed only after it has no more delegations.
	// This operation sends only the remaining dust to the community pool.
	feePool := h.k.GetFeePool(ctx)
	feePool.CommunityPool = feePool.CommunityPool.Add(outstanding...)
	h.k.SetFeePool(ctx, feePool)

	// delete outstanding
	h.k.DeleteValidatorOutstandingRewards(ctx, valAddr)

	// remove commission record
	h.k.DeleteValidatorAccumulatedCommission(ctx, valAddr)

	// clear historical rewards
	h.k.DeleteValidatorHistoricalRewards(ctx, valAddr)

	// clear current rewards
	h.k.DeleteValidatorCurrentRewards(ctx, valAddr)
}

// increment period
func (h Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	val := h.k.stakingKeeper.Validator(ctx, valAddr)
	h.k.IncrementValidatorPeriod(ctx, val)
}

// withdraw delegation rewards (which also increments period)
func (h Hooks) BeforeDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	val := h.k.stakingKeeper.Validator(ctx, valAddr)
	del := h.k.stakingKeeper.Delegation(ctx, delAddr, valAddr)

	if _, err := h.k.withdrawDelegationRewards(ctx, val, del); err != nil {
		panic(err)
	}
}

// create new delegation period record
func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) {
	h.k.initializeDelegation(ctx, valAddr, delAddr)
}

func (h Hooks) AfterValidatorBonded(_ sdk.Context, _ eth.Addr)         {}
func (h Hooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ eth.Addr) {}
func (h Hooks) AfterValidatorPowerUpdated(_ sdk.Context, _ eth.Addr)   {}
