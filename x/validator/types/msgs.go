package types

import (
	"github.com/celer-network/sgn-v2/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_errors "github.com/cosmos/cosmos-sdk/types/errors"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/proto"
)

const RouterKey = ModuleName // this was defined in your key.go file

const (
	TypeMsgSetTransactors           = "set_transactors"
	TypeMsgEditValidatorDescription = "edit_validator_description"
)

type MsgSetTransactors struct {
	Transactors []sdk.AccAddress `json:"transactors"`
	Sender      sdk.AccAddress   `json:"sender"`
}

// NewMsgSetTransactors is a constructor function for MsgSetTransactors
func NewMsgSetTransactors(transactors []sdk.AccAddress, sender sdk.AccAddress) MsgSetTransactors {
	return MsgSetTransactors{
		Transactors: transactors,
		Sender:      sender,
	}
}

func (m *MsgSetTransactors) Reset()         { *m = MsgSetTransactors{} }
func (m *MsgSetTransactors) String() string { return proto.CompactTextString(m) }
func (*MsgSetTransactors) ProtoMessage()    {}

// Route should return the name of the module
func (msg MsgSetTransactors) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetTransactors) Type() string { return TypeMsgSetTransactors }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetTransactors) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, msg.Sender.String())
	}

	for _, transactor := range msg.Transactors {
		if transactor.Empty() {
			return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, transactor.String())
		}

		err := sdk.VerifyAddressFormat(transactor)
		if err != nil {
			return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, err.Error())
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetTransactors) GetSignBytes() []byte {
	// TODO return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgSetTransactors) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

type MsgEditValidatorDescription struct {
	EthAddress  string                  `json:"eth_address"`
	Description sdk_staking.Description `json:"description"`
	Sender      sdk.AccAddress          `json:"sender"`
}

// TODO
func (m *MsgEditValidatorDescription) Reset()         { *m = MsgEditValidatorDescription{} }
func (m *MsgEditValidatorDescription) String() string { return proto.CompactTextString(m) }
func (*MsgEditValidatorDescription) ProtoMessage()    {}

// NewMsgEditValidatorDescription is a constructor function for MsgEditValidatorDescription
func NewMsgEditValidatorDescription(
	ethAddress string, description sdk_staking.Description, sender sdk.AccAddress) MsgEditValidatorDescription {

	return MsgEditValidatorDescription{
		EthAddress:  contracts.FormatAddrHex(ethAddress),
		Description: description,
		Sender:      sender,
	}
}

// Route should return the name of the module
func (msg MsgEditValidatorDescription) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditValidatorDescription) Type() string { return TypeMsgEditValidatorDescription }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditValidatorDescription) ValidateBasic() error {
	if msg.EthAddress == "" {
		return sdk_errors.Wrap(sdk_errors.ErrUnknownRequest, "EthAddress cannot be empty")
	}

	if msg.Description == (sdk_staking.Description{}) {
		return sdk_errors.Wrap(sdk_errors.ErrInvalidRequest, "empty description")
	}

	if msg.Sender.Empty() {
		return sdk_errors.Wrap(sdk_errors.ErrInvalidAddress, msg.Sender.String())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditValidatorDescription) GetSignBytes() []byte {
	// TODO return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgEditValidatorDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
