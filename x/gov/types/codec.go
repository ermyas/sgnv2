package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

// RegisterLegacyAminoCodec registers the necessary x/validator interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitProposal{}, "sgn-v2/MsgSubmitProposal", nil)
	cdc.RegisterConcrete(&MsgDeposit{}, "sgn-v2/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgVote{}, "sgn-v2/MsgVote", nil)
	cdc.RegisterConcrete(&TextProposal{}, "sgn-v2/TextProposal", nil)
	cdc.RegisterConcrete(&ParameterProposal{}, "sgn-v2/ParameterProposal", nil)
	cdc.RegisterConcrete(&UpgradeProposal{}, "sgn-v2/UpgradeProposal", nil)
	cdc.RegisterInterface((*Content)(nil), nil)
}

// RegisterInterfaces registers the x/staking interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitProposal{},
		&MsgDeposit{},
		&MsgVote{},
	)
	registry.RegisterInterface(
		"sgn.gov.v1.Content",
		(*Content)(nil),
		&TextProposal{},
		&ParameterProposal{},
		&UpgradeProposal{},
	)
}

// RegisterProposalTypeCodec registers an external proposal content type defined
// in another module for the internal ModuleCdc. This allows the MsgSubmitProposal
// to be correctly Amino encoded and decoded.
func RegisterProposalTypeCodec(o interface{}, name string) {
	ModuleCdc.RegisterConcrete(o, name, nil)
}

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
