package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/gov module sentinel errors
var (
	ErrUnknownProposal         = sdkerrors.Register(ModuleName, 501, "unknown proposal")
	ErrInactiveProposal        = sdkerrors.Register(ModuleName, 502, "inactive proposal")
	ErrAlreadyActiveProposal   = sdkerrors.Register(ModuleName, 503, "proposal already active")
	ErrInvalidProposalContent  = sdkerrors.Register(ModuleName, 504, "invalid proposal content")
	ErrInvalidProposalType     = sdkerrors.Register(ModuleName, 505, "invalid proposal type")
	ErrInvalidVote             = sdkerrors.Register(ModuleName, 506, "invalid vote option")
	ErrInvalidGenesis          = sdkerrors.Register(ModuleName, 507, "invalid genesis state")
	ErrNoProposalHandlerExists = sdkerrors.Register(ModuleName, 508, "no handler exists for proposal type")
)
