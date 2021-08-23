package types

import (
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrEmptyValidatorAddr = sdk_errors.Register(ModuleName, 2, "empty validator address")
	ErrNoValidatorFound   = sdk_errors.Register(ModuleName, 3, "validator does not exist")
	ErrNoDelegatorFound   = sdk_errors.Register(ModuleName, 4, "delegator does not exist")
	ErrInvalidAddress     = sdk_errors.Register(ModuleName, 5, "invalid address")
)
