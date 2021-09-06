package relayer

import (
	"fmt"
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

func NewOperator(cdc codec.Codec, cliHome string, tmCfg *tmcfg.Config, legacyAmino *codec.LegacyAmino) (operator *Operator, err error) {
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
		legacyAmino,
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

	consAddr := sdk.GetConsAddress(sdkPubKey)
	log.Infof("Validator %s consensus address %s", txr.Key.GetAddress(), consAddr)

	return &Operator{
		EthClient:  ethClient,
		Transactor: txr,
		PubKeyAny:  pubKeyAny,
	}, nil
}

func (o *Operator) SyncValidator(valAddr eth.Addr, currentBlkNum *big.Int, options ValSyncOptions) bool {
	updates, updated := o.SyncValidatorMsgs(valAddr, options)
	if len(updates) > 0 {
		msgs := synctypes.MsgProposeUpdates{
			Updates:  updates,
			EthBlock: currentBlkNum.Uint64(),
			Sender:   o.Transactor.Key.GetAddress().String(),
		}
		o.Transactor.AddTxMsg(&msgs)
	}
	return updated
}

func (o *Operator) SyncValidatorMsgs(valAddr eth.Addr, options ValSyncOptions) ([]*synctypes.ProposeUpdate, bool) {
	var updates []*synctypes.ProposeUpdate
	var update *synctypes.ProposeUpdate
	var updated ValSyncOptions
	if options.sgnaddr {
		update, updated.sgnaddr = o.SyncValidatorSgnAddrMsg(valAddr)
		if update != nil {
			updates = append(updates, update)
		}
	}
	if options.params {
		update, updated.params = o.SyncValidatorParamsMsg(valAddr)
		if update != nil {
			updates = append(updates, update)
		}
	}
	if options.states {
		update, updated.states = o.SyncValidatorStatesMsg(valAddr)
		if update != nil {
			updates = append(updates, update)
		}
	}
	return updates, updated == options
}

func (o *Operator) SyncValidatorSgnAddrMsg(valAddr eth.Addr) (*synctypes.ProposeUpdate, bool /*updated*/) {
	logmsg := fmt.Sprintf("Generate sync validator sgnaddr msg, val %x", valAddr)

	sgnAddr, err := o.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, valAddr)
	if err != nil {
		log.Errorf("%s. Failed to query contract sgn address err: %s", logmsg, err)
		return nil, false
	}
	exist, _ := validatorcli.QuerySgnAccount(o.Transactor.CliCtx, sdk.AccAddress(sgnAddr).String())
	if exist {
		log.Debugf("%s. already updated", logmsg)
		return nil, true
	}
	updateVal := &validatortypes.Validator{
		EthAddress: eth.Addr2Hex(valAddr),
		SgnAddress: sdk.AccAddress(sgnAddr).String(),
	}
	update := &synctypes.ProposeUpdate{
		Type: synctypes.DataType_ValidatorSgnAddr,
		Data: o.Transactor.CliCtx.Codec.MustMarshal(updateVal),
	}
	log.Debugf("%s. sgnaddr %s", logmsg, sdk.AccAddress(sgnAddr))
	return update, false
}

func (o *Operator) SyncValidatorParamsMsg(valAddr eth.Addr) (*synctypes.ProposeUpdate, bool /*updated*/) {
	logmsg := fmt.Sprintf("Generate sync validator params msg, val %x", valAddr)
	// TODO: separate signer and val addr
	if o.EthClient.Address != valAddr {
		log.Errorf("%s. Params sync can only be trigger by self validator", logmsg)
		return nil, false
	}
	ethVal, err := o.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, valAddr)
	if err != nil {
		log.Errorf("%s. Failed to query contract validator info:", logmsg, err)
		return nil, false
	}
	sgnAddrBytes, err := o.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, valAddr)
	if err != nil {
		log.Errorf("%s. Failed to query contract sgn address err: %s", logmsg, err)
		return nil, false
	}
	exist, err := validatorcli.QuerySgnAccount(o.Transactor.CliCtx, sdk.AccAddress(sgnAddrBytes).String())
	if !exist {
		log.Errorf("%s. Failed to query store sgn account err: %s", logmsg, err)
		return nil, false
	}

	updateVal := &validatortypes.Validator{
		EthAddress:      eth.Addr2Hex(valAddr),
		EthSigner:       eth.Addr2Hex(ethVal.Signer),
		SgnAddress:      sdk.AccAddress(sgnAddrBytes).String(),
		ConsensusPubkey: o.PubKeyAny,
		CommissionRate:  ethVal.CommissionRate,
		Description: &validatortypes.Description{
			Identity: eth.Addr2Hex(valAddr),
		},
	}
	storeVal, _ := validatorcli.QueryValidator(o.Transactor.CliCtx, valAddr.Hex())
	if storeVal != nil {
		if sameValidatorParams(updateVal, storeVal) {
			log.Debugf("%s, validator params already updated: %s", logmsg, updateVal)
			return nil, true
		}
	}
	update := &synctypes.ProposeUpdate{
		Type: synctypes.DataType_ValidatorParams,
		Data: o.Transactor.CliCtx.Codec.MustMarshal(updateVal),
	}
	log.Debugf("%s, updateVal: %s", logmsg, updateVal)
	return update, false
}

func (o *Operator) SyncValidatorStatesMsg(valAddr eth.Addr) (*synctypes.ProposeUpdate, bool /*updated*/) {
	logmsg := fmt.Sprintf("Generate sync validator states msg, val %x", valAddr)
	ethVal, err := o.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, valAddr)
	if err != nil {
		log.Errorf("%s. Failed to query contract validator info: %s", logmsg, err)
		return nil, false
	}
	storeVal, err := validatorcli.QueryValidator(o.Transactor.CliCtx, valAddr.Hex())
	if storeVal == nil {
		log.Debugf("%s. Failed to query store validator info: %s", logmsg, err)
		return nil, false
	}

	updateVal := &validatortypes.Validator{
		EthAddress: eth.Addr2Hex(valAddr),
		Status:     validatortypes.ValidatorStatus(ethVal.Status),
		Tokens:     ethVal.Tokens.String(),
		Shares:     ethVal.Shares.String(),
	}
	if sameValidatorStates(updateVal, storeVal) {
		log.Debugf("%s. Validator states already updated: %s", logmsg, updateVal)
		return nil, true
	}

	update := &synctypes.ProposeUpdate{
		Type: synctypes.DataType_ValidatorStates,
		Data: o.Transactor.CliCtx.Codec.MustMarshal(updateVal),
	}
	log.Debugf("%s, updateVal: %s", logmsg, updateVal)
	return update, false
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

type ValSyncOptions struct {
	sgnaddr bool
	params  bool
	states  bool
}

func (o ValSyncOptions) String() string {
	return fmt.Sprintf("ValSyncOptions={sgnaddr:%t, params:%t, states:%t}", o.sgnaddr, o.params, o.states)
}
