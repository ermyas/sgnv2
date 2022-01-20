package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSignMessage{}, "sgn-v2/MsgSignMessage", nil)
	cdc.RegisterConcrete(&MsgTriggerSignMessage{}, "sgn-v2/MsgTriggerSignMessage", nil)
	cdc.RegisterConcrete(&MsgClaimAllFees{}, "sgn-v2/MsgClaimAllFees", nil)
	cdc.RegisterConcrete(&MsgSignFees{}, "sgn-v2/MsgSignFees", nil)

	cdc.RegisterConcrete(&MsgProposal{}, "sgn-v2/MsgProposal", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSignMessage{},
		&MsgTriggerSignMessage{},
		&MsgClaimAllFees{},
		&MsgSignFees{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&MsgProposal{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
