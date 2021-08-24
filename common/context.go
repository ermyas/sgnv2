package common

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/viper"
)

func NewQueryCLIContext(cdc *codec.Codec) client.Context {
	ctx := client.Context{}.
		WithCodec(*cdc).
		WithNodeURI(viper.GetString(FlagSgnNodeURI)).
		WithChainID(viper.GetString(FlagSgnChainId))
	return ctx
}
