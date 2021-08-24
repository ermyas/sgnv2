package monitor

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	tmcfg "github.com/tendermint/tendermint/config"
	pvm "github.com/tendermint/tendermint/privval"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
)

type Operator struct {
	EthClient  *eth.EthClient
	Transactor *transactor.Transactor

	PubKeyAny *codectypes.Any
}

func NewOperator(cdc codec.Codec, cliHome string, tmCfg *tmcfg.Config) (operator *Operator, err error) {
	ethClient, err := common.NewEthClientFromConfig()
	if err != nil {
		return
	}

	txr, err := transactor.NewTransactor(
		cliHome,
		viper.GetString(common.FlagSgnChainId),
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

	privValidator := pvm.LoadFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile())
	tmPubKey, err := privValidator.GetPubKey()
	if err != nil {
		return nil, err
	}
	sdkPubKey, err := cryptocodec.FromTmPubKeyInterface(tmPubKey)
	if err != nil {
		return nil, err
	}
	pubKeyAny, err := codectypes.NewAnyWithValue(sdkPubKey)
	if err != nil {
		return nil, err
	}

	return &Operator{
		EthClient:  ethClient,
		Transactor: txr,
		PubKeyAny:  pubKeyAny,
	}, nil
}

// return true if already updated or no need for retry
func (o *Operator) SyncValidator(valEthAddr eth.Addr) bool {

	// TODO
	// candidate, err := validator.CLIQueryCandidate(o.Transactor.CliCtx, validator.RouterKey, valEthAddr.Hex())
	// if err != nil {
	// 	log.Errorln("sidechain query candidate err:", err)
	// 	return false
	// }

	valInfo, err := o.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, valEthAddr)
	if err != nil {
		log.Errorln("Failed to query validator info:", err)
		return false
	}

	// commission, err := common.NewCommission(o.EthClient, valInfo.CommissionRate)
	// if err != nil {
	// 	log.Errorln("Failed to create new commission:", err)
	// 	return false
	// }

	newVal := sdk_staking.Validator{
		Description: sdk_staking.Description{
			Identity: eth.Addr2Hex(valEthAddr), // TODO: Use a dedicated field instead of Identity
		},
		Tokens: sdk.NewIntFromBigInt(valInfo.Tokens), // not QuoRaw(common.TokenDec) yet
		Status: eth.ParseStatus(valInfo.Status),
		// Commission: commission,
	}

	if o.EthClient.Address == valEthAddr {
		newVal.ConsensusPubkey = o.PubKeyAny
	}
	return false
}

func (o *Operator) SyncDelegator(candidatorAddr, delegatorAddr eth.Addr) {
	// TODO
}
