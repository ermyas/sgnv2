package relayer

import (
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	synctypes "github.com/celer-network/sgn-v2/x/sync/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

func (r *Relayer) processPullerQueue() {
	if !r.isSyncer() {
		return
	}
	var keys, vals [][]byte
	r.lock.RLock()
	iterator, err := r.db.Iterator(PullerKeyPrefix, storetypes.PrefixEndBytes(PullerKeyPrefix))
	if err != nil {
		log.Errorln("Create db iterator err", err)
		return
	}
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
		vals = append(vals, iterator.Value())
	}
	iterator.Close()
	r.lock.RUnlock()

	validators := make(map[eth.Addr]bool)
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
			log.Infof("%s. validator %x notice key %d", logmsg, e.ValAddr, e.Key)
			validators[e.ValAddr] = true

		case *eth.StakingValidatorStatusUpdate:
			log.Infof("%s. validator %x %s", logmsg, e.ValAddr, eth.ParseValStatus(e.Status))
			validators[e.ValAddr] = true

		case *eth.StakingDelegationUpdate:
			log.Infof("%s. delegation update validator %x tokens %s delta %s, delegator %x shares %s",
				logmsg, e.ValAddr, e.ValTokens, e.TokenDiff, e.DelAddr, e.DelShares)
			delegators[getDelegatorKey(e.ValAddr, e.DelAddr)] = true
		}
	}

	msgs := synctypes.MsgProposeUpdates{
		Updates:  make([]*synctypes.ProposeUpdate, 0),
		EthBlock: r.getCurrentBlockNumber().Uint64(),
		Sender:   string(r.Transactor.Key.GetAddress()),
	}

	if r.isBootstrapped() {
		for validatorAddr := range validators {
			updates := r.SyncValidatorMsgs(validatorAddr)
			if len(updates) > 0 {
				msgs.Updates = append(msgs.Updates, updates...)
			}
		}
	}
	for delegatorKey := range delegators {
		validatorAddr := eth.Hex2Addr(strings.Split(delegatorKey, ":")[0])
		delegatorAddr := eth.Hex2Addr(strings.Split(delegatorKey, ":")[1])
		update := r.SyncDelegatorMsg(validatorAddr, delegatorAddr)
		if update != nil {
			msgs.Updates = append(msgs.Updates, update)
		}
	}

	if len(msgs.Updates) > 0 {
		r.Transactor.AddTxMsg(&msgs)
	}
}

func (r *Relayer) syncBlkNum() {
	if !r.isSyncer() {
		return
	}

}
