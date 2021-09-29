package types

import (
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper expected staking keeper (noalias)
type StakingKeeper interface {
	GetBondedValidators(ctx sdk.Context) (validators stakingtypes.Validators)
	GetValidatorBySgnAddr(ctx sdk.Context, sgnAddr sdk.AccAddress) (validator stakingtypes.ValidatorI, found bool)
}
