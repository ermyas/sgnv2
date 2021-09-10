package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrUnknownUpdate         = sdkerrors.Register(ModuleName, 1, "unknown update")
	ErrDoubleVote            = sdkerrors.Register(ModuleName, 2, "double vote")
	ErrInvalidUpdateType     = sdkerrors.Register(ModuleName, 3, "invalid update type")
	ErrInvalidUpdateData     = sdkerrors.Register(ModuleName, 4, "invalid update data")
	ErrInvalidGenesis        = sdkerrors.Register(ModuleName, 5, "invalid genesis state")
	ErrInvalidMsg            = sdkerrors.Register(ModuleName, 6, "invalid request")
	ErrInvalidAddress        = sdkerrors.Register(ModuleName, 7, "invalid address")
	ErrPendingUpdateNotFound = sdkerrors.Register(ModuleName, 8, "pending update not found")
)
