package types

import (
	"errors"

	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _, _, _, _ sdk.Msg = &MsgSignMessage{}, &MsgTriggerSignMessage{}, &MsgClaimAllFees{}, &MsgSignFees{}

func NewMsgSignMessage(sender sdk.AccAddress) *MsgSignMessage {
	return &MsgSignMessage{
		Sender: sender.String(),
	}
}

func (msg *MsgSignMessage) Route() string {
	return RouterKey
}

func (msg *MsgSignMessage) Type() string {
	return "SignMessage"
}

func (msg *MsgSignMessage) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSignMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgTriggerSignMessage(sender sdk.AccAddress) *MsgTriggerSignMessage {
	return &MsgTriggerSignMessage{
		Sender: sender.String(),
	}
}

func (msg *MsgTriggerSignMessage) Route() string {
	return RouterKey
}

func (msg *MsgTriggerSignMessage) Type() string {
	return "TriggerSignMessage"
}

func (msg *MsgTriggerSignMessage) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgTriggerSignMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerSignMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgClaimAllFees(delAddr eth.Addr, sender sdk.AccAddress) *MsgClaimAllFees {
	return &MsgClaimAllFees{
		DelegatorAddress: eth.Addr2Hex(delAddr),
		Sender:           sender.String(),
	}
}

func (msg *MsgClaimAllFees) Route() string {
	return RouterKey
}

func (msg *MsgClaimAllFees) Type() string {
	return "ClaimAllFees"
}

func (msg *MsgClaimAllFees) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgClaimAllFees) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimAllFees) ValidateBasic() error {
	if msg.DelegatorAddress == "" {
		return errors.New("invalid address")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgSignFees(address eth.Addr, sender sdk.AccAddress, signatureDetailsList []SignatureDetails) *MsgSignFees {
	return &MsgSignFees{
		Address:              eth.Addr2Hex(address),
		Sender:               sender.String(),
		SignatureDetailsList: signatureDetailsList,
	}
}

func (msg *MsgSignFees) Route() string {
	return RouterKey
}

func (msg *MsgSignFees) Type() string {
	return "SignFees"
}

func (msg *MsgSignFees) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSignFees) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignFees) ValidateBasic() error {
	if msg.Address == "" {
		return errors.New("invalid address")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	if len(msg.SignatureDetailsList) == 0 {
		return errors.New("empty signature details list")
	}
	return nil
}
