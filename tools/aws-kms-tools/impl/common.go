package impl

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
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
	FlagCfgHome      = "cfg"

	awskmsPre = "awskms"
)

var (
	DefaultCfgHome string

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
	chainId := new(big.Int).SetUint64(viper.GetUint64(FlagChainId))
	ksfile, passphrase := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	if strings.HasPrefix(ksfile, awskmsPre) {
		log.Debugln("use cfg kms ks", ksfile)
		kmskeyinfo := strings.SplitN(ksfile, ":", 3)
		if len(kmskeyinfo) != 3 {
			return nil, fmt.Errorf("%s has wrong format", ksfile)
		}
		awskeysec := []string{"", ""}
		if passphrase != "" {
			awskeysec = strings.SplitN(passphrase, ":", 2)
			if len(awskeysec) != 2 {
				return nil, fmt.Errorf("%s has wrong format", passphrase)
			}
		}
		return eth.NewKmsSigner(kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1], chainId)
	}
	return eth.NewKmsSigner(
		viper.GetString(FlagRegion),
		"alias/"+viper.GetString(FlagAlias),
		viper.GetString(FlagAwsKey),
		viper.GetString(FlagAwsSec),
		chainId,
	)
}

func getCfgRpc(chainId uint64) string {
	var mcc []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		log.Warnln("fail to load multichain configs err:", err)
		return ""
	}
	for _, cfg := range mcc {
		if cfg.ChainID == chainId {
			return cfg.Gateway
		}
	}
	return ""
}
