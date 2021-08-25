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

func (o *Operator) SyncValidator(valEthAddr eth.Addr, currentBlkNum *big.Int) bool {
	updates := o.SyncValidatorMsgs(valEthAddr)
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

func (o *Operator) SyncValidatorMsgs(valEthAddr eth.Addr) []*synctypes.ProposeUpdate {
	valInfo, err := o.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, valEthAddr)
	if err != nil {
		log.Errorln("Failed to query validator info:", err)
		return nil
	}

	sgnAddr, err := o.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, valEthAddr)
	if err != nil {
		log.Errorf("Failed to query sgn address err: %s", err)
		return nil
	}

	newVal := validatortypes.Validator{
		EthAddress: valEthAddr.Hex(),
		EthSigner:  valInfo.Signer.Hex(),
		Status:     validatortypes.ValidatorStatus(valInfo.Status),
		SgnAddress: string(sgnAddr),
		Tokens:     valInfo.Tokens.String(),
		Shares:     valInfo.Shares.String(),
		Description: &validatortypes.Description{
			Identity: eth.Addr2Hex(valEthAddr),
		},
	}

	storedVal, err := validatorcli.QueryValidator(o.Transactor.CliCtx, valEthAddr.Hex())
	if err != nil {
		log.Errorln("sgn query validator err:", err)
		return nil
	}

	updates := make([]*synctypes.ProposeUpdate, 0)
	if newVal.SgnAddress != storedVal.SgnAddress {
		update := &synctypes.ProposeUpdate{
			Type: synctypes.DataType_ValidatorAddrs,
			Data: o.Transactor.CliCtx.Codec.MustMarshal(&newVal),
		}
		updates = append(updates, update)
	}
	if newVal.Status != storedVal.Status {
		update := &synctypes.ProposeUpdate{
			Type: synctypes.DataType_ValidatorStates,
			Data: o.Transactor.CliCtx.Codec.MustMarshal(&newVal),
		}
		updates = append(updates, update)
	}

	return updates
}

func (o *Operator) SyncDelegatorMsg(valEthAddr, delEthAddr eth.Addr) *synctypes.ProposeUpdate {
	dInfo, err := o.EthClient.Contracts.Staking.GetDelegatorInfo(&bind.CallOpts{}, valEthAddr, delEthAddr)
	if err != nil {
		log.Errorf("failed to query delegator info err: %s", err)
		return nil
	}

	newVal := validatortypes.Delegator{
		ValAddress: valEthAddr.Hex(),
		DelAddress: delEthAddr.Hex(),
		Shares:     dInfo.Shares.String(),
	}

	storeVal, err := validatorcli.QueryDelegator(o.Transactor.CliCtx, valEthAddr.Hex(), delEthAddr.Hex())
	if err != nil {
		return nil
	}

	if storeVal.Shares != newVal.Shares {
		return &synctypes.ProposeUpdate{
			Type: synctypes.DataType_DelegatorShares,
			Data: o.Transactor.CliCtx.Codec.MustMarshal(&newVal),
		}
	} else {
		return nil
	}
}
