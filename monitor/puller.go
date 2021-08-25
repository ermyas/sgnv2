package monitor

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

func (m *Monitor) processPullerQueue() {
	if !m.isSyncer() {
		return
	}
	var keys, vals [][]byte
	m.lock.RLock()
	iterator, err := m.db.Iterator(PullerKeyPrefix, storetypes.PrefixEndBytes(PullerKeyPrefix))
	if err != nil {
		log.Errorln("Create db iterator err", err)
		return
	}
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, iterator.Key())
		vals = append(vals, iterator.Value())
	}
	iterator.Close()
	m.lock.RUnlock()

	validators := make(map[eth.Addr]bool)
	delegators := make(map[string]bool)
	for i, key := range keys {
		event := eth.NewEventFromBytes(vals[i])
		logmsg := fmt.Sprintf("Process puller event %s at block %d", event.Name, event.Log.BlockNumber)
		err = m.dbDelete(key)
		if err != nil {
			log.Errorf("%s. db Delete err: %s", logmsg, err)
			continue
		}

		switch e := event.ParseEvent(m.EthClient).(type) {
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
}

func (m *Monitor) syncBlkNum() {
	if !m.isSyncer() {
		return
	}

}
