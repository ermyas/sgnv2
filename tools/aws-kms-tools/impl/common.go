package impl

import (
	"math/big"

	"github.com/celer-network/goutils/eth"
	"github.com/spf13/viper"
)

const (
	FlagRegion       = "region"
	FlagRegionShort  = "r"
	FlagAlias        = "alias"
	FlagAliasShort   = "a"
	FlagChainId      = "chainid"
	FlagChainIdShort = "c"
	FlagAwsKey       = "awskey"
	FlagAwsSec       = "awssec"
)

func getKmsSigner() (*eth.KmsSigner, error) {
	return eth.NewKmsSigner(
		viper.GetString(FlagRegion),
		"alias/"+viper.GetString(FlagAlias),
		viper.GetString(FlagAwsKey),
		viper.GetString(FlagAwsSec),
		new(big.Int).SetUint64(viper.GetUint64(FlagChainId)),
	)
}
