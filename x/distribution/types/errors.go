package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/distribution module sentinel errors
var (
	ErrEmptyDelegatorAddr      = sdkerrors.Register(ModuleName, 22, "delegator address is empty")
	ErrEmptyWithdrawAddr       = sdkerrors.Register(ModuleName, 23, "withdraw address is empty")
	ErrEmptyValidatorAddr      = sdkerrors.Register(ModuleName, 24, "validator address is empty")
	ErrEmptyDelegationDistInfo = sdkerrors.Register(ModuleName, 25, "no delegation distribution info")
	ErrNoValidatorDistInfo     = sdkerrors.Register(ModuleName, 26, "no validator distribution info")
	ErrNoValidatorCommission   = sdkerrors.Register(ModuleName, 27, "no validator commission to withdraw")
	ErrSetWithdrawAddrDisabled = sdkerrors.Register(ModuleName, 28, "set withdraw address disabled")
	ErrBadDistribution         = sdkerrors.Register(ModuleName, 29, "community pool does not have sufficient coins to distribute")
	ErrInvalidProposalAmount   = sdkerrors.Register(ModuleName, 30, "invalid community pool spend proposal amount")
	ErrEmptyProposalRecipient  = sdkerrors.Register(ModuleName, 31, "invalid community pool spend proposal recipient")
	ErrNoValidatorExists       = sdkerrors.Register(ModuleName, 32, "validator does not exist")
	ErrNoDelegationExists      = sdkerrors.Register(ModuleName, 33, "delegation does not exist")
	ErrEmptySender             = sdkerrors.Register(ModuleName, 34, "sender address is empty")
	ErrClaimCooldownNotPassed  = sdkerrors.Register(ModuleName, 35, "claim cooldown not passed")
	ErrEmptySignature          = sdkerrors.Register(ModuleName, 36, "empty signature")
)
