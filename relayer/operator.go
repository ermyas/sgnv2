package relayer

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	validatorcli "github.com/celer-network/sgn-v2/x/validator/client/cli"
	validatortypes "github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
	tmcfg "github.com/tendermint/tendermint/config"
	pvm "github.com/tendermint/tendermint/privval"
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

func (o *Operator) SyncValidator(valAddr eth.Addr, currentBlkNum *big.Int) bool {
	updates := o.SyncValidatorMsgs(valAddr, ValSyncFlag{true, true})
	if len(updates) > 0 {
		msgs := synctypes.MsgProposeUpdates{
			Updates:  make([]*synctypes.ProposeUpdate, 0),
			EthBlock: currentBlkNum.Uint64(),
			Sender:   string(o.Transactor.Key.GetAddress()),
		}
		msgs.Updates = append(msgs.Updates, updates...)
		if len(msgs.Updates) > 0 {
			o.Transactor.AddTxMsg(&msgs)
		}
		return true
	} else {
		return false
	}
}

func (o *Operator) SyncValidatorMsgs(valAddr eth.Addr, flag ValSyncFlag) []*synctypes.ProposeUpdate {
	ethVal, err := o.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, valAddr)
	if err != nil {
		log.Errorln("Failed to query validator info:", err)
		return nil
	}
	storeVal, _ := validatorcli.QueryValidator(o.Transactor.CliCtx, valAddr.Hex())

	updates := make([]*synctypes.ProposeUpdate, 0)

	if flag.params {
		sgnAddr, err := o.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, valAddr)
		if err != nil {
			log.Errorf("Failed to query sgn address err: %s", err)
			return nil
		}
		updateVal := validatortypes.Validator{
			EthAddress:      valAddr.Hex(),
			EthSigner:       ethVal.Signer.Hex(),
			SgnAddress:      sdk.AccAddress(sgnAddr).String(),
			ConsensusPubkey: storeVal.ConsensusPubkey,
			CommissionRate:  ethVal.CommissionRate,
		}
		var skip bool
		if storeVal != nil {
			if updateVal.EthSigner == storeVal.EthSigner && updateVal.SgnAddress == storeVal.SgnAddress &&
				updateVal.CommissionRate == storeVal.CommissionRate {
				log.Debugf("validator %x params already updated", valAddr)
				skip = true
			}
		}
		if !skip {
			update := &synctypes.ProposeUpdate{
				Type: synctypes.DataType_ValidatorParams,
				Data: o.Transactor.CliCtx.Codec.MustMarshal(&updateVal),
			}
			updates = append(updates, update)
		}
	}

	if flag.states {
		updateVal := validatortypes.Validator{
			EthAddress: valAddr.Hex(),
			Status:     validatortypes.ValidatorStatus(ethVal.Status),
			Tokens:     ethVal.Tokens.String(),
			Shares:     ethVal.Shares.String(),
		}
		var skip bool
		if storeVal != nil {
			if updateVal.Status == storeVal.Status && updateVal.Tokens == storeVal.Tokens && updateVal.Shares == storeVal.Shares {
				log.Debugf("validator %x states already updated", valAddr)
				skip = true
			}
		}
		if !skip {
			update := &synctypes.ProposeUpdate{
				Type: synctypes.DataType_ValidatorStates,
				Data: o.Transactor.CliCtx.Codec.MustMarshal(&updateVal),
			}
			updates = append(updates, update)
		}
	}

	return updates
}

func (o *Operator) SyncDelegatorMsg(valAddr, delAddr eth.Addr) *synctypes.ProposeUpdate {
	ethDel, err := o.EthClient.Contracts.Staking.GetDelegatorInfo(&bind.CallOpts{}, valAddr, delAddr)
	if err != nil {
		log.Errorf("failed to query delegator info err: %s", err)
		return nil
	}

	updateDel := validatortypes.Delegator{
		ValAddress: valAddr.Hex(),
		DelAddress: delAddr.Hex(),
		Shares:     ethDel.Shares.String(),
	}

	storeVal, _ := validatorcli.QueryDelegator(o.Transactor.CliCtx, valAddr.Hex(), delAddr.Hex())

	if storeVal != nil {
		if updateDel.Shares == ethDel.Shares.String() {
			log.Debugf("delegator %x - %x shares already updated", valAddr, delAddr)
			return nil
		}
	}

	return &synctypes.ProposeUpdate{
		Type: synctypes.DataType_DelegatorShares,
		Data: o.Transactor.CliCtx.Codec.MustMarshal(&updateDel),
	}

}
