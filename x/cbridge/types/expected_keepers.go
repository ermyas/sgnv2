package types

import (
	"github.com/celer-network/sgn-v2/eth"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

// StakingKeeper expected staking keeper (noalias)
type StakingKeeper interface {
	GetBondedValidators(ctx sdk.Context) (validators stakingtypes.Validators)
	GetValidatorBySgnAddr(ctx sdk.Context, sgnAddr sdk.AccAddress) (validator stakingtypes.ValidatorI, found bool)
}

type FarmingKeeper interface {
	HasFarmingPool(ctx sdk.Context, poolName string) bool
	GetStakeInfo(ctx sdk.Context, addr eth.Addr, poolName string) (info farmingtypes.StakeInfo, found bool)
	Stake(ctx sdk.Context, poolName string, address eth.Addr, amount sdk.Coin) error
	Unstake(ctx sdk.Context, poolName string, address eth.Addr, amount sdk.Coin) error
}

type DistributionKeeper interface {
	ClaimCBridgeFeeShare(ctx sdk.Context, delAddr eth.Addr) error
	GetWithdrawableBalance(ctx sdk.Context, delAddr eth.Addr, coin sdk.Coin) sdk.Coin
}
