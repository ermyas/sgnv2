package common

import (
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/viper"
)

func NewEthClientFromConfig() (*eth.EthClient, error) {
	return eth.NewEthClient(
		viper.GetString(FlagEthGateway),
		viper.GetString(FlagEthSignerKeystore),
		viper.GetString(FlagEthSignerPassphrase),
		&eth.TransactorConfig{
			BlockDelay:               viper.GetUint64(FlagEthBlockDelay),
			BlockPollingInterval:     viper.GetUint64(FlagEthPollInterval),
			ChainId:                  big.NewInt(viper.GetInt64(FlagEthChainId)),
			AddGasPriceGwei:          viper.GetUint64(FlagEthAddGasPriceGwei),
			MinGasPriceGwei:          viper.GetUint64(FlagEthMinGasPriceGwei),
			MaxGasPriceGwei:          viper.GetUint64(FlagEthMaxGasPriceGwei),
			MaxFeePerGasGwei:         viper.GetUint64(FlagEthMaxFeePerGasGwei),
			MaxPriorityFeePerGasGwei: viper.GetUint64(FlagEthMaxPriorityFeePerGasGwei),
		},
		viper.GetString(FlagEthContractStaking),
		viper.GetString(FlagEthContractSgn),
		viper.GetString(FlagEthContractStakingReward),
		viper.GetString(FlagEthContractFarmingRewards),
		viper.GetString(FlagEthContractViewer),
		viper.GetString(FlagEthContractGovern),
	)
}
