package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgOnchainEvent{}

func NewMsgOnchainEvent(creator string, chainid uint64, evt string, elog []byte) *MsgOnchainEvent {
	return &MsgOnchainEvent{
		Creator: creator,
		Chainid: chainid,
		Evtype:  evt,
		Elog:    elog,
	}
}

func (msg *MsgOnchainManyEvents) Route() string {
	return RouterKey
}

func (msg *MsgOnchainManyEvents) Type() string {
	return "OnchainManyEvents"
}

func (msg *MsgOnchainManyEvents) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOnchainManyEvents) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOnchainManyEvents) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgOnchainEvent) Route() string {
	return RouterKey
}

func (msg *MsgOnchainEvent) Type() string {
	return "OnchainEvent"
}

func (msg *MsgOnchainEvent) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOnchainEvent) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOnchainEvent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
