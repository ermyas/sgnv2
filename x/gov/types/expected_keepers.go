package types

import (
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// ParamSubspace defines the expected Subspace interface for parameters (noalias)
type ParamSubspace interface {
	Get(ctx sdk.Context, key []byte, ptr interface{})
	Set(ctx sdk.Context, key []byte, param interface{})
}

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

// StakingKeeper expected staking keeper (noalias)
type StakingKeeper interface {
	// iterate through bonded validators by operator address, execute func for each validator
	IterateBondedValidators(sdk.Context,
		func(index int64, validator stakingtypes.ValidatorI) (stop bool))

	Validator(sdk.Context, eth.Addr) stakingtypes.ValidatorI                  // get a particular validator by ETH address
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI // get a particular validator by consensus address

	Delegation(sdk.Context, eth.Addr, eth.Addr) stakingtypes.DelegationI

	IterateDelegations(ctx sdk.Context, delegator eth.Addr,
		fn func(index int64, delegation stakingtypes.DelegationI) (stop bool))

	// SGN-specific
	GetValidatorBySgnAddr(ctx sdk.Context, sgnAddr sdk.AccAddress) (validator stakingtypes.ValidatorI, found bool)
}
