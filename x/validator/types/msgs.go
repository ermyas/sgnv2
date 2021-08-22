package types

import (
	"github.com/celer-network/sgn-v2/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSetTransactors  = "set_transactors"
	TypeMsgEditDescription = "edit_validator_description"
)

// NewMsgSetTransactors is a constructor function for MsgSetTransactors
func NewMsgSetTransactors(transactors []string, sender string) MsgSetTransactors {
	return MsgSetTransactors{
		Transactors: transactors,
		Sender:      sender,
	}
}

// Route should return the name of the module
func (msg *MsgSetTransactors) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgSetTransactors) Type() string { return TypeMsgSetTransactors }

// ValidateBasic runs stateless checks on the message
func (msg *MsgSetTransactors) ValidateBasic() error {
	if msg.Sender == "" {
		return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, msg.Sender)
	}

	for _, transactor := range msg.Transactors {
		if transactor == "" {
			return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, transactor)
		}

		_, err := sdk.AccAddressFromBech32(transactor)
		if err != nil {
			return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, err.Error())
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgSetTransactors) GetSignBytes() []byte {
	// TODO return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
	return nil
}

// GetSigners defines whose signature is required
func (msg *MsgSetTransactors) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// NewMsgEditDescription is a constructor function for MsgEditDescription
func NewMsgEditDescription(
	ethAddress string, description *Description, sender string) MsgEditDescription {

	return MsgEditDescription{
		EthAddress:  contracts.FormatAddrHex(ethAddress),
		Description: description,
		Sender:      sender,
	}
}

// Route should return the name of the module
func (msg *MsgEditDescription) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgEditDescription) Type() string { return TypeMsgEditDescription }

// ValidateBasic runs stateless checks on the message
func (msg *MsgEditDescription) ValidateBasic() error {
	if msg.EthAddress == "" {
		return sdk_errors.Wrap(sdk_errors.ErrUnknownRequest, "EthAddress cannot be empty")
	}

	if msg.Sender == "" {
		return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, msg.Sender)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgEditDescription) GetSignBytes() []byte {
	// TODO return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
	return nil
}

// GetSigners defines whose signature is required
func (msg *MsgEditDescription) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
