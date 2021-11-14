package types

import (
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ValidatorSet expected properties for the set of all validators (noalias)
type ValidatorSet interface {
	// iterate through bonded validators by eth address, execute func for each validator
	IterateBondedValidators(sdk.Context,
		func(index int64, validator ValidatorI) (stop bool))

	Validator(sdk.Context, eth.Addr) ValidatorI                  // get a particular validator by eth address
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) ValidatorI // get a particular validator by consensus address

	// Delegation allows for getting a particular delegation for a given validator
	// and delegator outside the scope of the staking module.
	Delegation(sdk.Context, eth.Addr, eth.Addr) DelegationI
}

// DelegationSet expected properties for the set of all delegations for a particular validator (noalias)
type DelegationSet interface {
	GetValidatorSet() ValidatorSet // validator set for which delegation set is based upon

	// iterate through all delegations from one delegator by validator-AccAddress,
	//   execute func for each validator
	IterateDelegations(ctx sdk.Context, delegator eth.Addr,
		fn func(index int64, delegation DelegationI) (stop bool))
}

// Event Hooks
// These can be utilized to communicate between a staking keeper and another
// keeper which must take particular actions when validators/delegators change
// state. The second keeper must implement this interface, which then the
// staking keeper can call.

// StakingHooks event hooks for staking validator object (noalias)
type StakingHooks interface {
	AfterValidatorCreated(ctx sdk.Context, valAddr eth.Addr) // Must be called when a validator is created
	AfterValidatorRemoved(ctx sdk.Context, valAddr eth.Addr) // Must be called when a validator is deleted

	AfterValidatorBonded(ctx sdk.Context, valAddr eth.Addr)         // Must be called when a validator is bonded
	AfterValidatorBeginUnbonding(ctx sdk.Context, valAddr eth.Addr) // Must be called when a validator begins unbonding
	AfterValidatorPowerUpdated(ctx sdk.Context, valAddr eth.Addr)

	BeforeDelegationCreated(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr)  // Must be called when a delegation is created
	BeforeDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) // Must be called when a delegation's shares are modified
	AfterDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr)
}
