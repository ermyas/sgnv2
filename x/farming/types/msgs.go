package types

import (
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// farming message types
const (
	MaxPoolNameLength = 128

	TypeMsgClaimRewards    = "claim_rewards"
	TypeMsgClaimAllRewards = "claim_all_rewards"
	TypeMsgSignRewards     = "sign_rewards"
)

// Verify interface at compile time
var _, _, _ sdk.Msg = &MsgClaimRewards{}, &MsgClaimAllRewards{}, &MsgSignRewards{}

func NewMsgClaimRewards(poolName string, address eth.Addr, sender sdk.AccAddress) *MsgClaimRewards {
	return &MsgClaimRewards{
		PoolName: poolName,
		Address:  eth.Addr2Hex(address),
		Sender:   sender.String(),
	}
}

func (msg MsgClaimRewards) Route() string {
	return RouterKey
}

func (msg MsgClaimRewards) Type() string {
	return TypeMsgClaimRewards
}

func (msg MsgClaimRewards) ValidateBasic() error {
	if msg.PoolName == "" || len(msg.PoolName) > MaxPoolNameLength {
		return WrapErrInvalidInput(msg.PoolName)
	}
	if msg.Address == "" {
		return WrapErrInvalidAddress("")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return WrapErrInvalidAddress(msg.Sender)
	}
	return nil
}

func (msg MsgClaimRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgClaimRewards) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{senderAddr}
}

func NewMsgClaimAllRewards(address eth.Addr, sender sdk.AccAddress) *MsgClaimAllRewards {
	return &MsgClaimAllRewards{
		Address: eth.Addr2Hex(address),
		Sender:  sender.String(),
	}
}

func (msg MsgClaimAllRewards) Route() string {
	return RouterKey
}

func (msg MsgClaimAllRewards) Type() string {
	return TypeMsgClaimAllRewards
}

func (msg MsgClaimAllRewards) ValidateBasic() error {
	if msg.Address == "" {
		return WrapErrInvalidAddress("")
	}
	if msg.Sender == "" {
		return WrapErrInvalidAddress("")
	}
	return nil
}

func (msg MsgClaimAllRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgClaimAllRewards) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{senderAddr}
}

func NewMsgSignRewards(
	address eth.Addr, sender sdk.AccAddress, signatureDetailsList []SignatureDetails) *MsgSignRewards {
	return &MsgSignRewards{
		Address:              eth.Addr2Hex(address),
		Sender:               sender.String(),
		SignatureDetailsList: signatureDetailsList,
	}
}

func (msg MsgSignRewards) Route() string {
	return RouterKey
}

func (msg MsgSignRewards) Type() string {
	return TypeMsgSignRewards
}

func (msg MsgSignRewards) ValidateBasic() error {
	if msg.Address == "" {
		return WrapErrInvalidAddress("")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return WrapErrInvalidAddress(msg.Sender)
	}
	if len(msg.SignatureDetailsList) == 0 {
		return WrapErrInvalidSig("")
	}
	return nil
}

func (msg MsgSignRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSignRewards) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{senderAddr}
}
