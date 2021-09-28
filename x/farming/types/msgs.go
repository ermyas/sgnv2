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
var _, _ sdk.Msg = &MsgClaimRewards{}, &MsgClaimAllRewards{}

func NewMsgClaimRewards(poolName string, address eth.Addr, sender sdk.AccAddress) *MsgClaimRewards {
	return &MsgClaimRewards{
		PoolName: poolName,
		Address:  eth.Addr2Hex(address),
		Sender:   sender.String(),
	}
}

func (m MsgClaimRewards) Route() string {
	return RouterKey
}

func (m MsgClaimRewards) Type() string {
	return TypeMsgClaimRewards
}

func (m MsgClaimRewards) ValidateBasic() error {
	if m.PoolName == "" || len(m.PoolName) > MaxPoolNameLength {
		return WrapErrInvalidInput(m.PoolName)
	}
	if m.Address == "" {
		return WrapErrInvalidAddress("")
	}
	if m.Sender == "" {
		return WrapErrInvalidAddress("")
	}
	return nil
}

func (m MsgClaimRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgClaimRewards) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(m.Sender)
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

func (m MsgClaimAllRewards) Route() string {
	return RouterKey
}

func (m MsgClaimAllRewards) Type() string {
	return TypeMsgClaimAllRewards
}

func (m MsgClaimAllRewards) ValidateBasic() error {
	if m.Address == "" {
		return WrapErrInvalidAddress("")
	}
	if m.Sender == "" {
		return WrapErrInvalidAddress("")
	}
	return nil
}

func (m MsgClaimAllRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgClaimAllRewards) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(m.Sender)
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

func (m MsgSignRewards) Route() string {
	return RouterKey
}

func (m MsgSignRewards) Type() string {
	return TypeMsgSignRewards
}

func (m MsgSignRewards) ValidateBasic() error {
	if m.Address == "" {
		return WrapErrInvalidAddress("")
	}
	if m.Sender == "" {
		return WrapErrInvalidAddress("")
	}
	if len(m.SignatureDetailsList) == 0 {
		return WrapErrInvalidSig("")
	}
	return nil
}

func (m MsgSignRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

func (m MsgSignRewards) GetSigners() []sdk.AccAddress {
	senderAddr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{senderAddr}
}
