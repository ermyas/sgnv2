package monitor

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/contracts"
	validatorcli "github.com/celer-network/sgn-v2/x/validator/client/cli"
)

type MonitorContractInfo struct {
	address contracts.Addr
	abi     string
}

func (info *MonitorContractInfo) GetAddr() contracts.Addr {
	return info.address
}

func (info *MonitorContractInfo) GetABI() string {
	return info.abi
}

func NewMonitorContractInfo(address contracts.Addr, abi string) *MonitorContractInfo {
	return &MonitorContractInfo{
		address: address,
		abi:     abi,
	}
}

func (m *Monitor) isSyncer() bool {
	syncer, err := validatorcli.QuerySyncer(m.Transactor.CliCtx)
	if err != nil {
		log.Errorln("Get syncer err", err)
		return false
	}

	validator, err := validatorcli.QueryValidator(m.Transactor.CliCtx, m.Transactor.Key.GetAddress().String())
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

func (m *Monitor) setUnbonded() {
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
