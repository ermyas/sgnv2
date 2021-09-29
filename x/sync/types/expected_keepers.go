package types

import (
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper expected staking keeper (noalias)
type StakingKeeper interface {
	GetValidator(ctx sdk.Context, ethAddr eth.Addr) (validator stakingtypes.Validator, found bool)
	InitAccount(ctx sdk.Context, accAddress sdk.AccAddress) error
	SetValidatorParams(ctx sdk.Context, val *stakingtypes.Validator, newValidator bool)
	SetValidatorStates(ctx sdk.Context, val *stakingtypes.Validator)
	SetDelegationShares(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr, shares sdk.Int)
	GetBondedValidators(ctx sdk.Context) (validators stakingtypes.Validators)
}
