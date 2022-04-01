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
		1:          "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
		10:         "https://mainnet.optimism.io",
		56:         "https://bsc-dataseed.binance.org",
		57:         "https://rpc.syscoin.org",
		66:         "https://exchainrpc.okex.org",
		100:        "https://rpc.gnosischain.com",
		128:        "https://http-mainnet-node.huobichain.com",
		137:        "https://polygon-rpc.com",
		250:        "https://rpc.ftm.tools",
		288:        "https://mainnet.boba.network/",
		336:        "https://evm.shiden.astar.network",
		592:        "https://rpc.astar.network:8545",
		1024:       "https://api-para.clover.finance",
		1030:       "https://evm.confluxrpc.com",
		1088:       "https://andromeda.metis.io/?owner=1088",
		1284:       "https://rpc.api.moonbeam.network",
		1285:       "https://rpc.api.moonriver.moonbeam.network",
		2001:       "https://rpc-mainnet-cardano-evm.c1.milkomeda.com",
		47805:      "https://rpc.rei.network",
		42161:      "https://arb1.arbitrum.io/rpc",
		42220:      "https://forno.celo.org",
		42262:      "https://emerald.oasis.dev",
		43114:      "https://api.avax.network/ext/bc/C/rpc",
		1313161554: "https://mainnet.aurora.dev",
		1666600000: "https://harmony-0-rpc.gateway.pokt.network",
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
