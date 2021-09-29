package types

import (
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper expected staking keeper (noalias)
type StakingKeeper interface {
	GetBondedValidators(ctx sdk.Context) (validators stakingtypes.Validators)
	GetValidatorBySgnAddr(ctx sdk.Context, sgnAddr sdk.AccAddress) (validator stakingtypes.ValidatorI, found bool)
}

type FarmingKeeper interface {
	HasFarmingPool(ctx sdk.Context, poolName string) bool
	Stake(ctx sdk.Context, poolName string, address eth.Addr, amount sdk.Coin, mintStakes bool) error
	Unstake(ctx sdk.Context, poolName string, address eth.Addr, amount sdk.Coin, burnStakes bool) error
}
