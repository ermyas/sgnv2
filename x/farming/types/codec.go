package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/farming interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaimRewards{}, "sgn-v2/MsgClaimRewards", nil)
	cdc.RegisterConcrete(&MsgClaimAllRewards{}, "sgn-v2/MsgClaimAllRewards", nil)

	cdc.RegisterConcrete(&AddPoolProposal{}, "sgn-v2/AddPoolProposal", nil)
	cdc.RegisterConcrete(&RemovePoolProposal{}, "sgn-v2/RemovePoolProposal", nil)
	cdc.RegisterConcrete(&AdjustRewardProposal{}, "sgn-v2/AdjustRewardProposal", nil)
	cdc.RegisterConcrete(&AddTokensProposal{}, "sgn-v2/AddTokensProposal", nil)
	cdc.RegisterConcrete(&SetRewardContractsProposal{}, "sgn-v2/SetRewardContractsProposal", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgClaimRewards{},
		&MsgClaimAllRewards{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&AddPoolProposal{},
		&RemovePoolProposal{},
		&AdjustRewardProposal{},
		&AddTokensProposal{},
		&SetRewardContractsProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/farming module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/distribution and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
