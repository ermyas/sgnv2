package monitor

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	valcli "github.com/celer-network/sgn-v2/x/validator/client/cli"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

func (m *Monitor) isSyncer() bool {
	syncer, err := valcli.QuerySyncer(m.Transactor.CliCtx)
	if err != nil {
		log.Errorln("Get syncer err", err)
		return false
	}

	validator, err := valcli.QueryValidator(m.Transactor.CliCtx, m.Transactor.Key.GetAddress().String())
	if err != nil {
		log.Errorln("Get validator err", err)
		return false
	}

	return syncer.SgnAddress == validator.SgnAddress
}

func (m *Monitor) getCurrentBlockNumber() *big.Int {
	return m.ethMonitor.GetCurrentBlockNumber()
}

func (m *Monitor) dbGet(key []byte) ([]byte, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.db.Get(key)
}

func (m *Monitor) dbSet(key, val []byte) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.db.Set(key, val)
}

func (m *Monitor) dbDelete(key []byte) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.db.Delete(key)
}

func (m *Monitor) isBonded() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.bonded
}

func (m *Monitor) setBonded() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.bonded = true
}

func (m *Monitor) clearBonded() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.bonded = false
}

func (m *Monitor) isBootstrapped() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.bootstrapped
}

func (m *Monitor) setBootstrapped() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.bootstrapped = true
}

func getEventCheckInterval(name string) uint64 {
	m := viper.GetStringMap(common.FlagEthCheckInterval)
	eventNameInConfig := strcase.ToSnake(string(name))
	if m[eventNameInConfig] != nil {
		return uint64(m[eventNameInConfig].(int64))
	}
	// If not specified, use the default value of 0
	return 0
}

func getDelegatorKey(validator, delegator eth.Addr) string {
	return eth.Addr2Hex(validator) + ":" + eth.Addr2Hex(delegator)
}
