package types

// DONTCOVER

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cbridge module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
	// this line is used by starport scaffolding # ibc/errors
)

func (e ErrMsg) Error() string {
	return fmt.Sprintf("err code: %d, desc: %s", e.Code, e.Desc)
}

func Error(code ErrCode, desc string, args ...interface{}) *ErrMsg {
	return &ErrMsg{Code: code, Desc: fmt.Sprintf(desc, args...)}
}
