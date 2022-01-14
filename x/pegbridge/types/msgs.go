package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _, _, _, _, _ sdk.Msg = &MsgSignMint{}, &MsgTriggerSignMint{}, &MsgSignWithdraw{}, &MsgTriggerSignWithdraw{}, &MsgClaimFee{}

func NewMsgSignMint(sender string) *MsgSignMint {
	return &MsgSignMint{
		Sender: sender,
	}
}

func (msg *MsgSignMint) Route() string {
	return RouterKey
}

func (msg *MsgSignMint) Type() string {
	return "SignMint"
}

func (msg *MsgSignMint) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSignMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgTriggerSignMint(sender string) *MsgTriggerSignMint {
	return &MsgTriggerSignMint{
		Sender: sender,
	}
}

func (msg *MsgTriggerSignMint) Route() string {
	return RouterKey
}

func (msg *MsgTriggerSignMint) Type() string {
	return "TriggerSignMint"
}

func (msg *MsgTriggerSignMint) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgTriggerSignMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerSignMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgSignWithdraw(sender string) *MsgSignWithdraw {
	return &MsgSignWithdraw{
		Sender: sender,
	}
}

func (msg *MsgSignWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgSignWithdraw) Type() string {
	return "SignWithdraw"
}

func (msg *MsgSignWithdraw) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSignWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgTriggerSignWithdraw(sender string) *MsgTriggerSignWithdraw {
	return &MsgTriggerSignWithdraw{
		Sender: sender,
	}
}

func (msg *MsgTriggerSignWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgTriggerSignWithdraw) Type() string {
	return "TriggerSignWithdraw"
}

func (msg *MsgTriggerSignWithdraw) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgTriggerSignWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerSignWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgClaimFee(
	delAddress string, chainId uint64, tokenAddress string, nonce uint64, sender string) *MsgClaimFee {
	return &MsgClaimFee{
		DelegatorAddress: delAddress,
		ChainId:          chainId,
		TokenAddress:     tokenAddress,
		Nonce:            nonce,
		Sender:           sender,
	}
}

func (msg *MsgClaimFee) Route() string {
	return RouterKey
}

func (msg *MsgClaimFee) Type() string {
	return "ClaimFee"
}

func (msg *MsgClaimFee) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgClaimFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimFee) ValidateBasic() error {
	if msg.DelegatorAddress == "" {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

func NewMsgClaimRefund(sender string) *MsgClaimRefund {
	return &MsgClaimRefund{
		Sender: sender,
	}
}

func (msg *MsgClaimRefund) Route() string {
	return RouterKey
}

func (msg *MsgClaimRefund) Type() string {
	return "ClaimRefund"
}

func (msg *MsgClaimRefund) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgClaimRefund) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimRefund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
