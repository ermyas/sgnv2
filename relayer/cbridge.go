package relayer

import "github.com/celer-network/sgn-v2/eth"

// NOTE: to keep cbridge related as independent as possible, we create another client for eth mainnet
// and only use it for cbridge related monitoring

// TODO: re-org code files

// multichain support for cbridge, including watcher DALs, eth client,
// monitor, transactor etc for each chain.

const (
	// each chain has its own prefix db. ie. cbr-123
	cbrDbPrefixFmt = "cbr-%d"
)

// just to satisfy monitor interface requirement
type cbrContract struct {
	*eth.Bridge
	Address eth.Addr
}

func (c *cbrContract) GetAddr() eth.Addr {
	return c.Address
}

func (c *cbrContract) GetABI() string {
	return eth.BridgeABI
}

// ethclient etc
type CbrOneChain struct {
}

// key is chainid
type CbrMgr map[uint64]*CbrOneChain

// for each chain, dial gw, newprefixdb, newWatchDAL, monitor
func NewCbridgeMgr() {}
