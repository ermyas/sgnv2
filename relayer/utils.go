package relayer

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mapset "github.com/deckarep/golang-set"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

func (r *Relayer) isSyncer() bool {
	syncer, err := stakingcli.QuerySyncer(r.Transactor.CliCtx)
	if err != nil {
		log.Errorln("Get syncer err", err)
		return false
	}
	return eth.Hex2Addr(syncer.EthAddress) == r.Operator.ValAddr
}

func (r *Relayer) getCurrentBlockNumber() *big.Int {
	return r.ethMonitor.GetCurrentBlockNumber()
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

func (r *Relayer) validateCbrSigs(sigs []*cbrtypes.AddrSig, curss *cbrtypes.SortedSigners) bool {
	if len(curss.GetSigners()) == 0 {
		return false
	}
	totalPower := big.NewInt(0)
	cursMap := make(map[eth.Addr]*cbrtypes.AddrAmt)
	for _, s := range curss.GetSigners() {
		power := big.NewInt(0).SetBytes(s.Amt)
		totalPower.Add(totalPower, power)
		cursMap[eth.Bytes2Addr(s.Addr)] = s
	}

	signedPower := big.NewInt(0)
	i := 0
	for _, s := range sigs {
		if addrAmt, ok := cursMap[eth.Bytes2Addr(s.Addr)]; ok {
			power := big.NewInt(0).SetBytes(addrAmt.Amt)
			signedPower.Add(signedPower, power)
			sigs[i] = s
			i++
		}
	}
	// truncate sigs not in the current signers set
	for j := i; j < len(sigs); j++ {
		sigs[j] = nil
	}
	sigs = sigs[:i]

	quorumStake := big.NewInt(0).Mul(totalPower, big.NewInt(2))
	quorumStake = quorumStake.Quo(quorumStake, big.NewInt(3))

	if signedPower.Cmp(quorumStake) > 0 {
		return true
	}

	return false
}
