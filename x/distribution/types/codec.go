package types

import (
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/distribution interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgWithdrawDelegatorReward{}, "sgn-v2/MsgWithdrawDelegationReward", nil)
	cdc.RegisterConcrete(&MsgWithdrawValidatorCommission{}, "sgn-v2/MsgWithdrawValidatorCommission", nil)
	cdc.RegisterConcrete(&MsgSetWithdrawAddress{}, "sgn-v2/MsgModifyWithdrawAddress", nil)
	cdc.RegisterConcrete(&MsgFundCommunityPool{}, "sgn-v2/MsgFundCommunityPool", nil)
	cdc.RegisterConcrete(&MsgClaimAllStakingReward{}, "sgn-v2/MsgClaimAllStakingReward", nil)
	cdc.RegisterConcrete(&MsgSignStakingReward{}, "sgn-v2/MsgSignStakingReward", nil)

	cdc.RegisterConcrete(&CommunityPoolSpendProposal{}, "sgn-v2/CommunityPoolSpendProposal", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgWithdrawDelegatorReward{},
		&MsgWithdrawValidatorCommission{},
		&MsgSetWithdrawAddress{},
		&MsgFundCommunityPool{},
		&MsgClaimAllStakingReward{},
		&MsgSignStakingReward{},
	)
	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&CommunityPoolSpendProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/distribution module codec. Note, the codec
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
