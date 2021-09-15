package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgProposeUpdates = "propose_updates"
	TypeMsgVoteUpdates    = "vote_updates"
)

func NewMsgProposeUpdates(updates []*ProposeUpdate, sender string) MsgProposeUpdates {
	return MsgProposeUpdates{
		Updates: updates,
		Sender:  sender,
	}
}

// Route should return the name of the module
func (msg *MsgProposeUpdates) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgProposeUpdates) Type() string { return TypeMsgProposeUpdates }

// ValidateBasic runs stateless checks on the message
func (msg *MsgProposeUpdates) ValidateBasic() error {
	if len(msg.Updates) == 0 {
		return sdkerrors.Wrap(ErrInvalidMsg, "empty update list")
	}
	if msg.Sender == "" {
		return sdkerrors.Wrap(ErrInvalidAddress, msg.Sender)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgProposeUpdates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners defines whose signature is required
func (msg *MsgProposeUpdates) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func NewMsgVoteUpdates(votes []*VoteUpdate, sender string) MsgVoteUpdates {
	return MsgVoteUpdates{
		Votes:  votes,
		Sender: sender,
	}
}

// Route should return the name of the module
func (msg *MsgVoteUpdates) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgVoteUpdates) Type() string { return TypeMsgVoteUpdates }

// ValidateBasic runs stateless checks on the message
func (msg *MsgVoteUpdates) ValidateBasic() error {
	if len(msg.Votes) == 0 {
		return sdkerrors.Wrap(ErrInvalidMsg, "empty vote list")
	}
	if msg.Sender == "" {
		return sdkerrors.Wrap(ErrInvalidAddress, msg.Sender)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgVoteUpdates) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners defines whose signature is required
func (msg *MsgVoteUpdates) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
