package types

import (
	"github.com/celer-network/sgn-v2/eth"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) types.ModuleAccountI

	// TODO remove with genesis 2-phases refactor https://github.com/cosmos/cosmos-sdk/issues/2862
	SetModuleAccount(sdk.Context, types.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule string, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

// StakingKeeper expected staking keeper (noalias)
type StakingKeeper interface {
	Validator(sdk.Context, eth.Addr) stakingtypes.ValidatorI                  // get a particular validator by operator address
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI // get a particular validator by consensus address

	// Delegation allows for getting a particular delegation for a given validator
	// and delegator outside the scope of the staking module.
	Delegation(sdk.Context, eth.Addr, eth.Addr) stakingtypes.DelegationI

	IterateDelegations(ctx sdk.Context, delegator eth.Addr,
		fn func(index int64, delegation stakingtypes.DelegationI) (stop bool))
}

// StakingHooks event hooks for staking validator object (noalias)
type StakingHooks interface {
	AfterValidatorCreated(ctx sdk.Context, valAddr eth.Addr)                           // Must be called when a validator is created
	AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr eth.Addr) // Must be called when a validator is deleted

	BeforeDelegationCreated(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr)        // Must be called when a delegation is created
	BeforeDelegationSharesModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr) // Must be called when a delegation's shares are modified
	AfterDelegationModified(ctx sdk.Context, delAddr eth.Addr, valAddr eth.Addr)
	BeforeValidatorSlashed(ctx sdk.Context, valAddr eth.Addr, fraction sdk.Dec)
}
