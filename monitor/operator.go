package monitor

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/contracts"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/viper"
)

type Operator struct {
	EthClient  *contracts.EthClient
	Transactor *transactor.Transactor
}

func NewOperator(cdc codec.Codec, cliHome string) (operator *Operator, err error) {
	ethClient, err := common.NewEthClientFromConfig()
	if err != nil {
		return
	}

	txr, err := transactor.NewTransactor(
		cliHome,
		viper.GetString(common.FlagSgnChainID),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		cdc,
		transactor.NewGasPriceEstimator(viper.GetString(common.FlagSgnNodeURI)),
	)
	if err != nil {
		return
	}
	txr.Run()

	return &Operator{
		EthClient:  ethClient,
		Transactor: txr,
	}, nil
}

// return true if already updated or no need for retry
func (o *Operator) SyncValidator(candidateAddr contracts.Addr) bool {

	// TODO
	return false
}

func (o *Operator) SyncDelegator(candidatorAddr, delegatorAddr contracts.Addr) {
	// TODO
}
