package relayer

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

func (c *CbrOneChain) setCurss(ss []*cbrtypes.Signer) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs, c.curss.powers = cbrtypes.SignersToEthArrays(ss)
}

func (c *CbrOneChain) setCurssByEvent(e *eth.BridgeSignersUpdated) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.curss.addrs = make([]eth.Addr, len(e.Signers))
	c.curss.powers = make([]*big.Int, len(e.Powers))
	for i, addr := range e.Signers {
		c.curss.addrs[i] = addr
		c.curss.powers[i] = e.Powers[i]
	}
}

func (c *CbrOneChain) getCurss() currentSigners {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.curss
}

// each event's key is name-blkNum-index, value is json marshaled elog
func (c *CbrOneChain) saveEvent(name string, elog ethtypes.Log) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	key := fmt.Sprintf("%s-%d-%d-%x", name, elog.BlockNumber, elog.Index, elog.TxHash)
	val, _ := json.Marshal(elog)
	return c.db.Set([]byte(key), val)
}

func (c *CbrOneChain) getEventCheckInterval(name string) uint64 {
	eventNameInConfig := strcase.ToSnake(name)

	var defaultInterval uint64
	m := viper.GetStringMap(common.FlagBridgeDefaultCheckInterval)
	if m[eventNameInConfig] != nil {
		defaultInterval = uint64(m[eventNameInConfig].(int64))
	}

	if c.checkIntervals[eventNameInConfig] != 0 {
		return c.checkIntervals[eventNameInConfig]
	}
	return defaultInterval
}
