package relayer

import (
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mapset "github.com/deckarep/golang-set"
)

func (r *Relayer) getCurrentBlockNumber() *big.Int {
	return big.NewInt(int64(r.mon.GetBlkNum()))
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

func (r *Relayer) setCbrSsUpdating() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.cbrSsUpdating = true
}

func (r *Relayer) setCbrSsUpdated() {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.cbrSsUpdating = false
}

func (r *Relayer) isCbrSsUpdating() bool {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.cbrSsUpdating
}

func (r *Relayer) isEthAddrBlocked(addrs ...eth.Addr) bool {
	r.cfgLock.RLock()
	defer r.cfgLock.RUnlock()
	for _, addr := range addrs {
		if r.blockedEthAddrs[addr] {
			return true
		}
	}
	return false
}

func (r *Relayer) validateSigs(signedValidators mapset.Set) (pass bool) {
	validators, err := stakingcli.QueryValidators(r.Transactor.CliCtx)
	if err != nil {
		log.Errorln("QueryValidators err", err)
		return false
	}

	totalStake := sdk.ZeroInt()
	votingStake := sdk.ZeroInt()
	for _, v := range validators {
		totalStake = totalStake.Add(v.BondedTokens())

		if signedValidators.Contains(v.EthSigner) {
			votingStake = votingStake.Add(v.BondedTokens())
		}
	}
	quorumStake := totalStake.MulRaw(2).QuoRaw(3)
	return votingStake.GT(quorumStake)
}

func RunWithInterval(run func(), checkIntervalSeconds uint64) {
	interval := time.Duration(checkIntervalSeconds) * time.Second
	for {
		time.Sleep(interval)
		run()
	}
}
