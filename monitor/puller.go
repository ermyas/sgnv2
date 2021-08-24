package monitor

import (
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/contracts"
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

	validators := make(map[contracts.Addr]bool)
	delegators := make(map[string]bool)
	for i, key := range keys {
		event := NewEventFromBytes(vals[i])
		logmsg := fmt.Sprintf("Process puller event %s at mainchain block %d", event.Name, event.Log.BlockNumber)
		err = m.dbDelete(key)
		if err != nil {
			log.Errorf("%s. db Delete err: %s", logmsg, err)
			continue
		}

		switch event.ParseEvent(m.EthClient).(type) {
		case *contracts.StakingValidatorParamsUpdate:
			//TODO

		case *contracts.StakingValidatorStatusUpdate:
			//TODO

		case *contracts.StakingDelegationUpdate:
			//TODO

		case *contracts.SGNSgnAddrUpdate:
			//TODO
		}
	}

	if m.isBootstrapped() {
		for validatorAddr := range validators {
			m.SyncValidator(validatorAddr)
		}
	}
	for delegatorKey := range delegators {
		candidatorAddr := contracts.Hex2Addr(strings.Split(delegatorKey, ":")[0])
		delegatorAddr := contracts.Hex2Addr(strings.Split(delegatorKey, ":")[1])
		m.SyncDelegator(candidatorAddr, delegatorAddr)
	}
}
