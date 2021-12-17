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
	cdc.RegisterConcrete(&MsgSignMint{}, "sgn-v2/MsgSignMint", nil)
	cdc.RegisterConcrete(&MsgSignWithdraw{}, "sgn-v2/MsgSignWithdraw", nil)
	cdc.RegisterConcrete(&MsgTriggerSignMint{}, "sgn-v2/MsgTriggerSignMint", nil)
	cdc.RegisterConcrete(&MsgTriggerSignWithdraw{}, "sgn-v2/MsgTriggerSignWithdraw", nil)
	cdc.RegisterConcrete(&PegProposal{}, "sgn-v2/PegProposal", nil)
	cdc.RegisterConcrete(&PairDeleteProposal{}, "sgn-v2/PairDeleteProposal", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSignMint{},
		&MsgSignWithdraw{},
		&MsgTriggerSignMint{},
		&MsgTriggerSignWithdraw{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&PegProposal{},
		&PairDeleteProposal{},
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
