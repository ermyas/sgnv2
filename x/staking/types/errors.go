package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyValidatorAddr = sdkerrors.Register(ModuleName, 101, "empty validator address")
	ErrValidatorNotFound  = sdkerrors.Register(ModuleName, 102, "validator not found")
	ErrDelegationNotFound = sdkerrors.Register(ModuleName, 103, "delegation not found")
	ErrSgnAccountNotFound = sdkerrors.Register(ModuleName, 104, "sgn account not found")
	ErrInvalidAddress     = sdkerrors.Register(ModuleName, 105, "invalid address")
	ErrInvalidType        = sdkerrors.Register(ModuleName, 106, "invalid type")
)
