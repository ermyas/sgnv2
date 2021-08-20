package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
)

var ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry()) //TODO

func init() {
	// RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	// cdc.RegisterConcrete(MsgSetTransactors{}, "validator/MsgSetTransactors", nil)
	// cdc.RegisterConcrete(MsgEditCandidateDescription{}, "validator/MsgEditCandidateDescription", nil)
	// cdc.RegisterConcrete(MsgClaimReward{}, "validator/MsgClaimReward", nil)
	// cdc.RegisterConcrete(MsgSignReward{}, "validator/MsgSignReward", nil)
}
