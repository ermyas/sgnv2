package types

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidInput                       = sdkerrors.Register(ModuleName, 601, "invalid input")
	ErrPoolAlreadyExist                   = sdkerrors.Register(ModuleName, 602, "pool already exists")
	ErrNoFarmingPoolFound                 = sdkerrors.Register(ModuleName, 603, "no farming pool found")
	ErrNoStakeInfoFound                   = sdkerrors.Register(ModuleName, 604, "no stake info found")
	ErrTokenNotExist                      = sdkerrors.Register(ModuleName, 605, "token not exist")
	ErrPoolNotFinished                    = sdkerrors.Register(ModuleName, 606, "pool not finished")
	ErrUnexpectedProposalType             = sdkerrors.Register(ModuleName, 607, "unexpected proposal type")
	ErrInvalidAddress                     = sdkerrors.Register(ModuleName, 608, "invalid address")
	ErrInvalidDenom                       = sdkerrors.Register(ModuleName, 609, "invalid denom")
	ErrSendCoinsFromAccountToModuleFailed = sdkerrors.Register(ModuleName, 610, "send coins from account to module failed")
	ErrUnknownFarmingMsgType              = sdkerrors.Register(ModuleName, 611, "unknown farming msg type")
	ErrUnknownFarmingQueryType            = sdkerrors.Register(ModuleName, 612, "unknown farming query type")
	ErrInvalidInputAmount                 = sdkerrors.Register(ModuleName, 613, "invalid input amount")
	ErrInsufficientAmount                 = sdkerrors.Register(ModuleName, 614, "insufficient amount")
	ErrInvalidStartHeight                 = sdkerrors.Register(ModuleName, 615, "invalid start height")
	ErrPoolNameLength                     = sdkerrors.Register(ModuleName, 616, "invalid pool name length")
	ErrSendCoinsFromModuleToAccountFailed = sdkerrors.Register(ModuleName, 617, "send coins from module to account failed")
	ErrMintCoinsFailed                    = sdkerrors.Register(ModuleName, 618, "mint coins failed")
	ErrBurnCoinsFailed                    = sdkerrors.Register(ModuleName, 619, "burn coins failed")
	ErrInvalidSig                         = sdkerrors.Register(ModuleName, 620, "invalid signature")
	ErrClaimCooldownNotPassed             = sdkerrors.Register(ModuleName, 621, "claim cooldown not passed")
	ErrNoClaimInfoFound                   = sdkerrors.Register(ModuleName, 622, "no claim info found")
)

// WrapErrInvalidInput returns an error when an input parameter is invalid
func WrapErrInvalidInput(msg string) error {
	return sdkerrors.Wrapf(ErrInvalidInput, "invalid input: %s", msg)
}

// WrapErrPoolAlreadyExist returns an error when a pool exist
func WrapErrPoolAlreadyExist(poolName string) error {
	return sdkerrors.Wrapf(ErrPoolAlreadyExist, "farming pool %s already exists", poolName)
}

// WrapErrNoFarmingPoolFound returns an error when a farming pool doesn't exist
func WrapErrNoFarmingPoolFound(poolName string) error {
	return sdkerrors.Wrapf(ErrNoFarmingPoolFound, "farming pool %s does not exist", poolName)
}

// WrapErrNoStakeInfoFound returns an error when an address doesn't have any stake infos
func WrapErrNoStakeInfoFound(addr string, pool string) error {
	return sdkerrors.Wrapf(ErrNoStakeInfoFound, "%s has no stake in pool %s", addr, pool)
}

// WrapErrTokenNotExist returns an error when a token not exists
func WrapErrTokenNotExist(tokenName string) error {
	return sdkerrors.Wrapf(ErrTokenNotExist, "token %s does not exist", tokenName)
}

// WrapErrPoolNotFinished returns an error when the pool is not finished and cannot be removed
func WrapErrPoolNotFinished(poolName string) error {
	return sdkerrors.Wrapf(ErrPoolNotFinished, "cannot remove pool %s which has staked tokens and / or unclaimed rewards", poolName)
}

// WrapErrUnexpectedProposalType returns an error when the proposal type is not supported by the farming module
func WrapErrUnexpectedProposalType(proposalType string) error {
	return sdkerrors.Wrapf(ErrUnexpectedProposalType, "proposal type %s not supported by farming module", proposalType)
}

// WrapErrInvalidAddress returns an error when an invalid address appears
func WrapErrInvalidAddress(addr string) error {
	return sdkerrors.Wrapf(ErrInvalidAddress, "invalid address: %s", addr)
}

// WrapErrInvalidDenom returns an error when it provides an unmatched token name
func WrapErrInvalidDenom(expectedDenom string, actualDenom string) error {
	return sdkerrors.Wrapf(ErrInvalidDenom, "the denom should be %s, not %s", expectedDenom, actualDenom)
}

// WrapErrSendCoinsFromAccountToModuleFailed returns an error when sending from account to module failed
func WrapErrSendCoinsFromAccountToModuleFailed(msg string) error {
	return sdkerrors.Wrapf(ErrSendCoinsFromAccountToModuleFailed, "send coins from account to module failed: %s", msg)
}

// WrapErrUnknownFarmingMsgType returns an error when the message type is unknown
func WrapErrUnknownFarmingMsgType(msg string) error {
	return sdkerrors.Wrapf(ErrUnknownFarmingMsgType, "unknown farming msg type: %s", msg)
}

// WrapErrUnknownFarmingMsgType returns an error when the query type is unknown
func WrapErrUnknownFarmingQueryType(msg string) error {
	return sdkerrors.Wrapf(ErrUnknownFarmingMsgType, "unknown farming query type: %s", msg)
}

// WrapErrInvalidInputAmount returns an error when an input amount is invalid
func WrapErrInvalidInputAmount(amount string) error {
	return sdkerrors.Wrapf(ErrInvalidInputAmount, "invalid input amount: %s", amount)
}

// WrapErrInsufficientAmount returns an error when there is no enough tokens
func WrapErrInsufficientAmount(amount string, inputAmount string) error {
	return sdkerrors.Wrapf(ErrInsufficientAmount, "actual amount %s less than %s", amount, inputAmount)
}

// WrapErrInvalidStartHeight returns an error when the start_height parameter is invalid
func WrapErrInvalidStartHeight(startHeight int64, currentHeight int64) error {
	return sdkerrors.Wrapf(ErrInvalidStartHeight, "start height %d less than current height %d", startHeight, currentHeight)
}

// WrapErrPoolNameLength returns an error when length of pool name is invalid
func WrapErrPoolNameLength(poolName string, got, max int) error {
	return sdkerrors.Wrapf(ErrPoolNameLength, "invalid pool name length for %s, length %d greater than max %d", poolName, got, max)
}

// WrapErrSendCoinsFromModuleToAccountFailed returns an error when sending from module to account failed
func WrapErrSendCoinsFromModuleToAccountFailed(msg string) error {
	return sdkerrors.Wrapf(ErrSendCoinsFromModuleToAccountFailed, "send coins from module to account failed: %s", msg)
}

// WrapErrMintCoinsFailed returns an error when minting coins failed
func WrapErrMintCoinsFailed(msg string) error {
	return sdkerrors.Wrapf(ErrMintCoinsFailed, "mint coins failed: %s", msg)
}

// WrapErrBurnCoinsFailed returns an error when burning coins failed
func WrapErrBurnCoinsFailed(msg string) error {
	return sdkerrors.Wrapf(ErrBurnCoinsFailed, "burn coins failed: %s", msg)
}

// WrapErrInvalidSig returns an error when a signature is invalid
func WrapErrInvalidSig(msg string) error {
	return sdkerrors.Wrapf(ErrInvalidSig, "invalid sig: %s", msg)
}

// WrapErrClaimCooldownNotPassed returns an error if the claim cooldown has not passed
func WrapErrClaimCooldownNotPassed(lastClaimTime time.Time) error {
	return sdkerrors.Wrapf(ErrClaimCooldownNotPassed, "last claim time: %s", lastClaimTime)
}

// WrapErrNoClaimInfoFound returns an error if no claim info is found for an address
func WrapErrNoClaimInfoFound(addr string) error {
	return sdkerrors.Wrapf(ErrNoClaimInfoFound, "no claim info found for: %s", addr)
}
