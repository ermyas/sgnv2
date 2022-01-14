package ops

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
			BlockDelay:               viper.GetUint64(common.FlagEthBlockDelay),
			BlockPollingInterval:     viper.GetUint64(common.FlagEthPollInterval),
			ChainId:                  big.NewInt(viper.GetInt64(common.FlagEthChainId)),
			AddGasPriceGwei:          viper.GetUint64(common.FlagEthAddGasPriceGwei),
			MinGasPriceGwei:          viper.GetUint64(common.FlagEthMinGasPriceGwei),
			MaxGasPriceGwei:          viper.GetUint64(common.FlagEthMaxGasPriceGwei),
			MaxFeePerGasGwei:         viper.GetUint64(common.FlagEthMaxFeePerGasGwei),
			MaxPriorityFeePerGasGwei: viper.GetUint64(common.FlagEthMaxPriorityFeePerGasGwei),
		},
		viper.GetString(common.FlagEthContractStaking),
		viper.GetString(common.FlagEthContractSgn),
		viper.GetString(common.FlagEthContractStakingReward),
		viper.GetString(common.FlagEthContractFarmingRewards),
		viper.GetString(common.FlagEthContractViewer),
		viper.GetString(common.FlagEthContractGovern),
	)
}

func calcRawAmount(amount string) *big.Int {
	rawAmount := new(big.Int)
	rawAmount.SetString(amount, 10)
	return new(big.Int).Mul(rawAmount, big.NewInt(common.CelrPrecision))
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

type cbrContract struct {
	*eth.Bridge
	Address eth.Addr
}

type CbrOneChain struct {
	*ethclient.Client
	*ethutils.Transactor
	contract *cbrContract
}

func newOneChain(chainId uint64) (*CbrOneChain, error) {
	var mcc []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		log.Fatalln("fail to load multichain configs err:", err)
	}
	signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	for _, cfg := range mcc {
		if cfg.ChainID == chainId {

			// Temporary hack for local manual test
			if chainId == 883 {
				cfg.Gateway = "http://127.0.0.1:8545"
			} else if chainId == 884 {
				cfg.Gateway = "http://127.0.0.1:8547"
			}

			ec, err := ethclient.Dial(cfg.Gateway)
			if err != nil {
				log.Fatalln("dial", cfg.Gateway, "err:", err)
			}
			chid, err := ec.ChainID(context.Background())
			if err != nil {
				log.Fatalf("get chainid %d err: %s", cfg.ChainID, err)
			}
			if chid.Uint64() != cfg.ChainID {
				log.Fatalf("chainid mismatch! cfg has %d but onchain has %d", cfg.ChainID, chid.Uint64())
			}
			cbr, err := eth.NewBridge(eth.Hex2Addr(cfg.CBridge), ec)
			if err != nil {
				log.Fatalln("cbridge contract at", cfg.CBridge, "err:", err)
			}
			signer, addr, err := eth.CreateSigner(signerKey, signerPass, chid)
			if err != nil {
				log.Fatalln("CreateSigner err:", err)
			}
			transactor := ethutils.NewTransactorByExternalSigner(
				addr,
				signer,
				ec,
				big.NewInt(int64(cfg.ChainID)),
				ethutils.WithBlockDelay(cfg.BlkDelay),
				ethutils.WithPollingInterval(time.Duration(cfg.BlkInterval)*time.Second),
				ethutils.WithAddGasEstimateRatio(cfg.AddGasEstimateRatio),
				ethutils.WithAddGasGwei(cfg.AddGasGwei),
				ethutils.WithMaxFeePerGasGwei(cfg.MaxFeePerGasGwei),
			)
			if err != nil {
				log.Fatalln("NewTransactor err:", err)
			}
			c := &CbrOneChain{
				Client:     ec,
				Transactor: transactor,
				contract: &cbrContract{
					Bridge:  cbr,
					Address: eth.Hex2Addr(cfg.CBridge),
				},
			}
			return c, nil
		}
	}

	return nil, fmt.Errorf("chainId %d not exist", chainId)
}
