package impl

import (
	"context"
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
	FlagRpc          = "rpc"
)

var (
	bgCtx = context.Background()

	// Common chain ID to JSON-RPC endpoint
	commonChainIdRpcs = map[uint64]string{
		1:     "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
		10:    "https://mainnet.optimism.io",
		56:    "https://bsc-dataseed.binance.org",
		137:   "https://polygon-rpc.com",
		250:   "https://rpc.ftm.tools",
		42161: "https://arb1.arbitrum.io/rpc",
		43114: "https://api.avax.network/ext/bc/C/rpc",
	}
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
