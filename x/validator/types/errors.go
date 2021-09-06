package types

import (
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyValidatorAddr = sdk_errors.Register(ModuleName, 2, "empty validator address")
	ErrValidatorNotFound  = sdk_errors.Register(ModuleName, 3, "validator not found")
	ErrDelegatorNotFound  = sdk_errors.Register(ModuleName, 4, "delegator not found")
	ErrSgnAccounNotFound  = sdk_errors.Register(ModuleName, 5, "sgn account not found")
	ErrInvalidAddress     = sdk_errors.Register(ModuleName, 6, "invalid address")
	ErrInvalidType        = sdk_errors.Register(ModuleName, 7, "invalid type")
)
