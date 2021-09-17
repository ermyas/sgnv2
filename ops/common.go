package ops

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

const (
	FlagKeystore   = "keystore"
	FlagPassphrase = "passphrase"
	FlagValidator  = "validator"
	FlagAmount     = "amount"
)

func newEthClient() (*eth.EthClient, error) {
	return eth.NewEthClient(
		viper.GetString(common.FlagEthGateway),
		viper.GetString(FlagKeystore),
		viper.GetString(FlagPassphrase),
		&eth.TransactorConfig{
			BlockDelay:           viper.GetUint64(common.FlagEthBlockDelay),
			BlockPollingInterval: viper.GetUint64(common.FlagEthPollInterval),
			ChainId:              big.NewInt(viper.GetInt64(common.FlagEthChainId)),
			AddGasPriceGwei:      viper.GetUint64(common.FlagEthAddGasPriceGwei),
			MinGasPriceGwei:      viper.GetUint64(common.FlagEthMinGasPriceGwei),
		},
		viper.GetString(common.FlagEthContractStaking),
		viper.GetString(common.FlagEthContractSgn),
		viper.GetString(common.FlagEthContractReward),
		viper.GetString(common.FlagEthContractViewer),
		viper.GetString(common.FlagEthContractGovern),
	)
}

func calcRawAmount(amount string) *big.Int {
	rawAmount := new(big.Int)
	rawAmount.SetString(amount, 10)
	return new(big.Int).Mul(rawAmount, big.NewInt(common.TokenDec))
}

func approveCelr(ethClient *eth.EthClient, spender eth.Addr, amount *big.Int) error {
	celrContract, err := eth.NewErc20(
		eth.Hex2Addr(viper.GetString(common.FlagEthContractCelr)),
		ethClient.Client,
	)
	if err != nil {
		return err
	}
	allowance, err := celrContract.Allowance(&bind.CallOpts{}, ethClient.Transactor.Address(), spender)
	if err != nil {
		return err
	}
	if allowance.Cmp(amount) < 0 {
		log.Infof("Approving %s CELR to Staking contract", amount)
		_, approveErr := ethClient.Transactor.TransactWaitMined(
			"Approve",
			func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
				return celrContract.Approve(opts, spender, amount)
			},
		)
		if approveErr != nil {
			return approveErr
		}
	}
	return nil
}
