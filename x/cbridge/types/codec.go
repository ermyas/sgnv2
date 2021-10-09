package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgInitWithdraw{}, "cbridge/InitWithdraw", nil)
	cdc.RegisterConcrete(&MsgSendMySig{}, "cbridge/SendMySig", nil)
	cdc.RegisterConcrete(&MsgSignAgain{}, "cbridge/SignAgain", nil)
	cdc.RegisterConcrete(&CbrProposal{}, "cbridge/CbrProposal", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
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
