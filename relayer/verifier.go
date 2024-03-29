package relayer

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	validatorcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	synccli "github.com/celer-network/sgn-v2/x/sync/client/cli"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
)

func (r *Relayer) verifyPendingUpdates() {
	interval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalVerifier)) * time.Second
	log.Infoln("start verify pending updates, interval:", interval)
	for {
		time.Sleep(interval)
		v, err := validatorcli.QueryValidator(r.Transactor.CliCtx, r.Operator.ValAddr.Hex())
		if err != nil || v.Status != stakingtypes.Bonded {
			log.Traceln("skip verifying pending updates as I am not a bonded validator")
			continue
		}
		pendingUpdates, err := synccli.QueryPendingUpdates(r.Transactor.CliCtx)
		if err != nil {
			log.Errorln("Query pending updates error:", err)
			continue
		}

		puSize := len(pendingUpdates)
		batchSize := 100
		if puSize > 0 {
			var wg sync.WaitGroup
			for i := 0; i < puSize; i += batchSize {
				wg.Add(1)
				j := i
				go func() {
					defer wg.Done()
					end := j + batchSize
					if end > puSize {
						end = puSize
					}
					r.verifyUpdates(pendingUpdates[j:end])
				}()
			}
			wg.Wait()
		}
	}
}

func (r *Relayer) verifyUpdates(pendingUpdates []*synctypes.PendingUpdate) {
	msgs := synctypes.MsgVoteUpdates{
		Votes:  make([]*synctypes.VoteUpdate, 0),
		Sender: r.Transactor.Key.GetAddress().String(),
	}
	for _, update := range pendingUpdates {
		_, err := r.verifiedUpdates.Get(strconv.Itoa(int(update.Id)))
		if err == nil {
			continue
		}

		done, approve := r.verifyUpdate(update)
		if done {
			err = r.verifiedUpdates.Set(strconv.Itoa(int(update.Id)), []byte{})
			if err != nil {
				log.Errorln("verifiedUpdates Set err", err)
				continue
			}
			if approve {
				msgs.Votes = append(msgs.Votes, &synctypes.VoteUpdate{
					Id:     update.Id,
					Option: synctypes.VoteOption_Yes,
				})
			}
		}
	}

	if len(msgs.Votes) > 0 {
		r.Transactor.AddTxMsg(&msgs)
	}
}

func (r *Relayer) verifyUpdate(update *synctypes.PendingUpdate) (done, approve bool) {
	switch update.Type {
	case synctypes.DataType_ValidatorSgnAddr:
		return r.verifyValidatorSgnAddr(update)
	case synctypes.DataType_ValidatorParams:
		return r.verifyValidatorParams(update)
	case synctypes.DataType_ValidatorStates:
		return r.verifyValidatorStates(update)
	case synctypes.DataType_DelegatorShares:
		return r.verifyDelegatorShares(update)
	case synctypes.DataType_CbrOnchainEvent:
		return r.verifyCbrEventUpdate(update)
	case synctypes.DataType_CbrUpdateCbrPrice:
		return r.verifyUpdateCbrPrice(update)
	case synctypes.DataType_PegbrOnChainEvent:
		return r.verifyPegbrEventUpdate(update)
	case synctypes.DataType_MsgbrOnChainEvent:
		return r.verifyMsgbrEventUpdate(update)
	default:
		return false, false
	}
}

func (r *Relayer) verifyValidatorSgnAddr(update *synctypes.PendingUpdate) (done, approve bool) {
	updateVal, err := stakingtypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}
	logmsg := fmt.Sprintf("verify update id %d, sgnaddr for validator: %s", update.Id, updateVal.String())

	sgnAddr, err := r.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, eth.Hex2Addr(updateVal.EthAddress))
	if err != nil {
		log.Errorf("%s. query contract sgn address err: %s", logmsg, err)
		return false, false
	}

	exist, _ := validatorcli.QuerySgnAccount(r.Transactor.CliCtx, sdk.AccAddress(sgnAddr).String())
	if exist {
		log.Infof("%s. sgn account already updated", logmsg)
		return true, false
	}

	if updateVal.SgnAddress != sdk.AccAddress(sgnAddr).String() {
		values := fmt.Sprintf("sgnaddr %s", sdk.AccAddress(sgnAddr).String())
		if r.cmpBlkNum(update.ChainBlock) == 1 {
			log.Errorf("%s. validator params not match eth values: %s", logmsg, values)
			return true, false
		}
		log.Infof("%s. eth block not passed, values: %s", logmsg, values)
		return false, false
	}

	return true, true
}

func (r *Relayer) verifyValidatorParams(update *synctypes.PendingUpdate) (done, approve bool) {
	updateVal, err := stakingtypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}
	logmsg := fmt.Sprintf("verify update id %d, params for validator: %s", update.Id, updateVal.String())

	storeVal, err := validatorcli.QueryValidator(r.Transactor.CliCtx, updateVal.EthAddress)
	if err == nil {
		if sameValidatorParams(&updateVal, storeVal) {
			log.Infof("%s. validator already updated", logmsg)
			return true, false
		}
	}

	ethVal, err := r.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, eth.Hex2Addr(updateVal.EthAddress))
	if err != nil {
		log.Errorf("%s. query validator info err: %s", logmsg, err)
		return false, false
	}

	sgnAddr, err := r.EthClient.Contracts.Sgn.SgnAddrs(&bind.CallOpts{}, eth.Hex2Addr(updateVal.EthAddress))
	if err != nil {
		log.Errorf("%s. query sgn address err: %s", logmsg, err)
		return false, false
	}

	if updateVal.EthSigner != eth.Addr2Hex(ethVal.Signer) ||
		updateVal.SgnAddress != sdk.AccAddress(sgnAddr).String() ||
		!updateVal.CommissionRate.Equal(sdk.NewDec(int64(ethVal.CommissionRate)).QuoInt64(eth.CommissionRateBase)) {
		values := fmt.Sprintf("signer %x sgnaddr %s commission %d",
			ethVal.Signer, sdk.AccAddress(sgnAddr).String(), ethVal.CommissionRate)
		if r.cmpBlkNum(update.ChainBlock) == 1 {
			log.Errorf("%s. validator params not match eth values: %s", logmsg, values)
			return true, false
		}
		log.Infof("%s. eth block not passed, values: %s", logmsg, values)
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) verifyValidatorStates(update *synctypes.PendingUpdate) (done, approve bool) {
	updateVal, err := stakingtypes.UnmarshalValidator(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}
	logmsg := fmt.Sprintf("verify update id %d, states for validator: %s", update.Id, updateVal.String())

	storeVal, err := validatorcli.QueryValidator(r.Transactor.CliCtx, updateVal.EthAddress)
	if err != nil {
		log.Infof("%s. validator not found", logmsg)
		return true, false
	}
	if sameValidatorStates(&updateVal, storeVal) {
		log.Infof("%s. states already updated", logmsg)
		return true, false
	}
	ethVal, err := r.EthClient.Contracts.Staking.Validators(&bind.CallOpts{}, eth.Hex2Addr(updateVal.EthAddress))
	if err != nil {
		log.Errorf("%s. query validator info err: %s", logmsg, err)
		return false, false
	}

	if updateVal.Status != stakingtypes.BondStatus(ethVal.Status) ||
		updateVal.Tokens.BigInt().Cmp(ethVal.Tokens) != 0 ||
		!updateVal.DelegatorShares.Equal(sdk.NewIntFromBigInt(ethVal.Shares)) {
		states := fmt.Sprintf("status %s tokens %s shares %s",
			updateVal.Status, updateVal.Tokens.BigInt(), updateVal.DelegatorShares)
		values := fmt.Sprintf("status %s tokens %s shares %s",
			stakingtypes.BondStatus(ethVal.Status), ethVal.Tokens, sdk.NewIntFromBigInt(ethVal.Shares))
		if r.cmpBlkNum(update.ChainBlock) == 1 {
			log.Infof("%s. validator states not match, states: %s, eth values: %s", logmsg, states, values)
			return true, false
		}
		log.Infof("%s. eth block not passed, states: %s, eth values: %s", logmsg, states, values)
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) verifyDelegatorShares(update *synctypes.PendingUpdate) (done, approve bool) {
	updateDel, err := stakingtypes.UnmarshalDelegation(r.Transactor.CliCtx.Codec, update.Data)
	if err != nil {
		return true, false
	}
	if eth.Hex2Addr(updateDel.DelegatorAddress) == eth.ZeroAddr {
		return true, false
	}
	logmsg := fmt.Sprintf("verify update id %d, shares for delegator: %s", update.Id, updateDel.String())

	storeDel, err := validatorcli.QueryDelegation(r.Transactor.CliCtx, updateDel.ValidatorAddress, updateDel.DelegatorAddress)
	if err == nil {
		if updateDel.Shares == storeDel.Shares {
			log.Infof("%s. shares already updated", logmsg)
			return true, false
		}
	}

	ethDel, err := r.EthClient.Contracts.Staking.GetDelegatorInfo(
		&bind.CallOpts{},
		eth.Hex2Addr(updateDel.ValidatorAddress),
		eth.Hex2Addr(updateDel.DelegatorAddress))
	if err != nil {
		log.Errorf("%s. query delegator info err: %s", logmsg, err)
		return false, false
	}

	if !updateDel.Shares.Equal(sdk.NewIntFromBigInt(ethDel.Shares)) {
		if r.cmpBlkNum(update.ChainBlock) == 1 {
			log.Infof("%s. delegator shares not match eth value: %s", logmsg, ethDel.Shares)
			return true, false
		}
		log.Infof("%s. eth block not passed, eth value: %s", logmsg, ethDel.Shares)
		return false, false
	}

	log.Infof("%s, success", logmsg)
	return true, true
}

func (r *Relayer) cmpBlkNum(blkNum uint64) int8 {
	currentBlkNum := r.getCurrentBlockNumber().Uint64()
	if currentBlkNum > blkNum {
		return 1
	} else if currentBlkNum < blkNum {
		return -1
	}
	return 0
}

func sameValidatorParams(updateVal, storeVal *stakingtypes.Validator) bool {
	if updateVal.EthAddress == storeVal.EthAddress &&
		updateVal.EthSigner == storeVal.EthSigner &&
		updateVal.SgnAddress == storeVal.SgnAddress &&
		updateVal.CommissionRate.Equal(storeVal.CommissionRate) {
		return true
	}
	return false
}

func sameValidatorStates(updateVal, storeVal *stakingtypes.Validator) bool {
	if updateVal.EthAddress == storeVal.EthAddress &&
		updateVal.Status == storeVal.Status &&
		updateVal.Tokens.Equal(storeVal.Tokens) &&
		updateVal.DelegatorShares.Equal(storeVal.DelegatorShares) {
		return true
	}
	return false
}
