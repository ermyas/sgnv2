package relayer

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	valcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

func (r *Relayer) isSyncer() bool {
	syncer, err := valcli.QuerySyncer(r.Transactor.CliCtx)
	if err != nil {
		log.Errorln("Get syncer err", err)
		return false
	}
	return eth.Hex2Addr(syncer.EthAddress) == r.Operator.ValAddr
}

func (r *Relayer) getCurrentBlockNumber() *big.Int {
	return r.ethMonitor.GetCurrentBlockNumber()
}

func (r *Relayer) dbGet(key []byte) ([]byte, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.db.Get(key)
}

func (r *Relayer) dbSet(key, val []byte) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.db.Set(key, val)
}

func (r *Relayer) dbDelete(key []byte) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.db.Delete(key)
}

func (r *Relayer) isBonded() bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.bonded
}

func (r *Relayer) setBonded() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.bonded = true
}

func (r *Relayer) clearBonded() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.bonded = false
}

func (r *Relayer) isBootstrapped() bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.bootstrapped
}

func (r *Relayer) setBootstrapped() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.bootstrapped = true
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
