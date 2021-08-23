package types

import (
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidMsg     = sdk_errors.Register(ModuleName, 1, "invalid request")
	ErrInvalidAddress = sdk_errors.Register(ModuleName, 2, "invalid address")
)
