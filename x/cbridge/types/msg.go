package types

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _, _, _ sdk.Msg = &MsgInitWithdraw{}, &MsgSendMySig{}, &MsgSignAgain{}

func NewMsgInitWithdraw(creator string) *MsgInitWithdraw {
	return &MsgInitWithdraw{
		Creator: creator,
	}
}

func (msg *MsgInitWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgInitWithdraw) Type() string {
	return "InitWithdraw"
}

func (msg *MsgInitWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgInitWithdraw) String() string {
	if msg == nil {
		return "nil"
	}
	return fmt.Sprintf("xferId %x, chainId %d, lpAddr %x, token %x, amount %s, reqId %d, creator %s",
		msg.XferId, msg.Chainid, msg.LpAddr, msg.Token, big.NewInt(0).SetBytes(msg.Amount), msg.ReqId, msg.Creator)
}

func NewMsgSendMySig(creator string, dataType SignDataType, data []byte, mySig []byte) *MsgSendMySig {
	return &MsgSendMySig{
		Creator:  creator,
		Datatype: dataType,
		Data:     data,
		MySig:    mySig,
	}
}

func (msg *MsgSendMySig) Route() string {
	return RouterKey
}

func (msg *MsgSendMySig) Type() string {
	return "SendMySig"
}

func (msg *MsgSendMySig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendMySig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendMySig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgSignAgain(creator string) *MsgSignAgain {
	return &MsgSignAgain{
		Creator: creator,
	}
}

func (msg *MsgSignAgain) Route() string {
	return RouterKey
}

func (msg *MsgSignAgain) Type() string {
	return "SignAgain"
}

func (msg *MsgSignAgain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSignAgain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSignAgain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgInternalTransfer(transfer, sig []byte, creator string) *MsgInternalTransfer {
	return &MsgInternalTransfer{
		Transfer: transfer,
		Sig:      sig,
		Creator:  creator,
	}
}

func (msg *MsgInternalTransfer) Route() string {
	return RouterKey
}

func (msg *MsgInternalTransfer) Type() string {
	return "InternalTransfer"
}

func (msg *MsgInternalTransfer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInternalTransfer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInternalTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Transfer) == 0 {
		return fmt.Errorf("empty transfer")
	}
	if len(msg.Sig) == 0 {
		return fmt.Errorf("empty sig")
	}
	return nil
}
