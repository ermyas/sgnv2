package relayer

import (
	"fmt"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/spf13/viper"
)

func (r *Relayer) processPullerQueue() {
	interval := time.Duration(viper.GetUint64(common.FlagEthPollInterval)) * time.Second
	log.Infoln("start process puller queue, interval:", interval)
	for {
		time.Sleep(interval)

		if !r.isSyncer() {
			continue
		}
		var keys, vals [][]byte
		r.lock.RLock()
		iterator, err := r.db.Iterator(PullerKeyPrefix, storetypes.PrefixEndBytes(PullerKeyPrefix))
		if err != nil {
			log.Errorln("Create db iterator err", err)
			r.lock.RUnlock()
			continue
		}
		for ; iterator.Valid(); iterator.Next() {
			keys = append(keys, iterator.Key())
			vals = append(vals, iterator.Value())
		}
		iterator.Close()
		r.lock.RUnlock()

		validators := make(map[eth.Addr]ValSyncOptions)
		delegators := make(map[string]bool)
		for i, key := range keys {
			event := eth.NewEventFromBytes(vals[i])
			logmsg := fmt.Sprintf("Process puller event %s at block %d", event.Name, event.Log.BlockNumber)
			err = r.dbDelete(key)
			if err != nil {
				log.Errorf("%s. db Delete err: %s", logmsg, err)
				continue
			}

			switch e := event.ParseEvent(r.EthClient).(type) {
			case *eth.StakingValidatorNotice:
				log.Infof("%s. validator %x notice key %s", logmsg, e.ValAddr, e.Key)
				if e.Key != "sgn-addr" {
					log.Errorf("puller only sync sgn-addr")
					continue
				}
				v := validators[e.ValAddr]
				v.sgnaddr = true
				validators[e.ValAddr] = v

			case *eth.StakingValidatorStatusUpdate:
				log.Infof("%s. validator %x %s", logmsg, e.ValAddr, eth.ParseValStatus(e.Status))
				v := validators[e.ValAddr]
				v.states = r.isBootstrapped()
				validators[e.ValAddr] = v

			case *eth.StakingDelegationUpdate:
				log.Infof("%s. delegation update validator %x tokens %s delta %s, delegator %x shares %s",
					logmsg, e.ValAddr, e.ValTokens, e.TokenDiff, e.DelAddr, e.DelShares)
				if e.DelAddr != eth.ZeroAddr { // zero address means slashing
					delegators[getDelegatorKey(e.ValAddr, e.DelAddr)] = true
				}
				v := validators[e.ValAddr]
				v.states = r.isBootstrapped()
				validators[e.ValAddr] = v
			}
		}

		msgs := synctypes.MsgProposeUpdates{
			Updates: make([]*synctypes.ProposeUpdate, 0),
			Sender:  r.Transactor.Key.GetAddress().String(),
		}

		blkNum := r.getCurrentBlockNumber().Uint64()
		for vaddr := range validators {
			updates, _ := r.SyncValidatorMsgs(vaddr, blkNum, validators[vaddr])
			msgs.Updates = append(msgs.Updates, updates...)
		}

		for delegatorKey := range delegators {
			validatorAddr := eth.Hex2Addr(strings.Split(delegatorKey, ":")[0])
			delegatorAddr := eth.Hex2Addr(strings.Split(delegatorKey, ":")[1])
			update := r.SyncDelegatorMsg(validatorAddr, delegatorAddr, blkNum)
			if update != nil {
				msgs.Updates = append(msgs.Updates, update)
			}
		}

		if len(msgs.Updates) > 0 {
			r.Transactor.AddTxMsg(&msgs)
		}
	}
}

func getDelegatorKey(validator, delegator eth.Addr) string {
	return eth.Addr2Hex(validator) + ":" + eth.Addr2Hex(delegator)
}
