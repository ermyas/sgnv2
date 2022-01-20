package types

import (
	time "time"

	"github.com/celer-network/sgn-v2/eth"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/message module sentinel errors
var (
	ErrNoMessageFound         = sdkerrors.Register(ModuleName, 8100, "no message found")
	ErrNoMessageBusFound      = sdkerrors.Register(ModuleName, 8101, "no MessageBus found")
	ErrNoTransferFound        = sdkerrors.Register(ModuleName, 8102, "no transfer found")
	ErrClaimCooldownNotPassed = sdkerrors.Register(ModuleName, 8103, "claim cooldown not passed")
	ErrNoClaimInfoFound       = sdkerrors.Register(ModuleName, 8104, "no claim info found")
)

// WrapErrNoMessageFound returns an error if no message is found for a message ID
func WrapErrNoMessageFound(messageId eth.Hash) error {
	return sdkerrors.Wrapf(ErrNoMessageFound, "%x", messageId)
}

// WrapErrNoMessageBusFound returns an error if no MessageBus contract is found for a chainId
func WrapErrNoMessageBusFound(chainId uint64) error {
	return sdkerrors.Wrapf(ErrNoMessageBusFound, "chainId: %d", chainId)
}

// WrapErrNoTransferFound returns an error if no transfer is found for a message ID
func WrapErrNoTransferFound(messageId eth.Hash) error {
	return sdkerrors.Wrapf(ErrNoTransferFound, "%x", messageId)
}

// WrapErrClaimCooldownNotPassed returns an error if the claim cooldown has not passed
func WrapErrClaimCooldownNotPassed(lastClaimTime time.Time) error {
	return sdkerrors.Wrapf(ErrClaimCooldownNotPassed, "last claim time: %s", lastClaimTime)
}

// WrapErrNoClaimInfoFound returns an error if no claim info is found for an address
func WrapErrNoClaimInfoFound(addr string) error {
	return sdkerrors.Wrapf(ErrNoClaimInfoFound, "no claim info found for: %s", addr)
}
