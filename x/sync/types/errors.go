package types

import (
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrUnknownUpdate         = sdk_errors.Register(ModuleName, 1, "unknown update")
	ErrDoubleVote            = sdk_errors.Register(ModuleName, 2, "double vote")
	ErrInvalidUpdateType     = sdk_errors.Register(ModuleName, 3, "invalid update type")
	ErrInvalidUpdateData     = sdk_errors.Register(ModuleName, 4, "invalid update data")
	ErrInvalidGenesis        = sdk_errors.Register(ModuleName, 5, "invalid genesis state")
	ErrInvalidMsg            = sdk_errors.Register(ModuleName, 6, "invalid request")
	ErrInvalidAddress        = sdk_errors.Register(ModuleName, 7, "invalid address")
	ErrPendingUpdateNotFound = sdk_errors.Register(ModuleName, 8, "pending update not found")
)
