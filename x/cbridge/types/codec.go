package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/cbridge interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInitWithdraw{}, "cbridge/InitWithdraw", nil)
	cdc.RegisterConcrete(&MsgSendMySig{}, "cbridge/SendMySig", nil)
	cdc.RegisterConcrete(&MsgSignAgain{}, "cbridge/SignAgain", nil)
	cdc.RegisterConcrete(&CbrProposal{}, "cbridge/CbrProposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgInitWithdraw{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSendMySig{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSignAgain{})
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&CbrProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
