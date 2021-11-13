package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = ModuleName // this was defined in your key.go file

func NewMsgSignSlash(nonce uint64, sig []byte, sender sdk.AccAddress) MsgSignSlash {
	return MsgSignSlash{
		Nonce:  nonce,
		Sig:    sig,
		Sender: sender.String(),
	}
}

// Route should return the name of the module
func (msg MsgSignSlash) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSignSlash) Type() string { return "sign_penalty" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSignSlash) ValidateBasic() error {
	if len(msg.Sig) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Sig cannot be empty")
	}

	if msg.Sender == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Sender)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (m MsgSignSlash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners defines whose signature is required
func (m MsgSignSlash) GetSigners() []sdk.AccAddress {
	proposer, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{proposer}
}
