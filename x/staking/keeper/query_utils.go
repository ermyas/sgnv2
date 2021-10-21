package keeper

import (
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDelegatorValidators returns all validators that a delegator is bonded to.
// If maxRetrieve is supplied, the respective amount will be returned.
func (k Keeper) GetDelegatorValidators(
	ctx sdk.Context, delegatorAddr eth.Addr, maxRetrieve uint32,
) types.Validators {
	validators := make([]types.Validator, maxRetrieve)

	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := types.GetDelegationsKey(delegatorAddr)

	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey) // smallest to largest
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())

		validator, found := k.GetValidator(ctx, delegation.GetValidatorAddr())
		if !found {
			panic(types.ErrValidatorNotFound)
		}

		validators[i] = validator
		i++
	}

	return validators[:i] // trim
}

// GetDelegatorValidator returns a validator that a delegator is bonded to
func (k Keeper) GetDelegatorValidator(
	ctx sdk.Context, delegatorAddr eth.Addr, validatorAddr eth.Addr,
) (validator types.Validator, err error) {
	delegation, found := k.GetDelegation(ctx, delegatorAddr, validatorAddr)
	if !found {
		return validator, types.ErrDelegationNotFound
	}

	validator, found = k.GetValidator(ctx, delegation.GetValidatorAddr())
	if !found {
		panic(types.ErrValidatorNotFound)
	}

	return validator, nil
}

// GetAllDelegatorDelegations returns all delegations for a delegator
func (k Keeper) GetAllDelegatorDelegations(ctx sdk.Context, delegator eth.Addr) []types.Delegation {
	delegations := make([]types.Delegation, 0)

	store := ctx.KVStore(k.storeKey)
	delegatorPrefixKey := types.GetDelegationsKey(delegator)

	iterator := sdk.KVStorePrefixIterator(store, delegatorPrefixKey) // smallest to largest
	defer iterator.Close()

	i := 0

	for ; iterator.Valid(); iterator.Next() {
		delegation := types.MustUnmarshalDelegation(k.cdc, iterator.Value())
		delegations = append(delegations, delegation)
		i++
	}

	return delegations
}
