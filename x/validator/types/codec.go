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
	// cdc.RegisterConcrete(MsgEditValidatorDescription{}, "validator/MsgEditValidatorDescription", nil)
}
