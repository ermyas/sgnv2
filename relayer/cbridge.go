package relayer

import (
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/log"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"

	dbm "github.com/tendermint/tm-db"
)

// NOTE: to keep cbridge related as independent as possible, we create another client for eth mainnet
// and only use it for cbridge related monitoring

// TODO: re-org code files

// multichain support for cbridge, including eth client,
// monitor, transactor etc for each chain.

const (
	cbrDbPrefix = "cbr-"

	// event names
	CbrEventSend  = "Send"
	CbrEventRelay = "Relay"
	// from pool.sol
	CbrEventLiqAdd   = "LiquidityAdded"
	CbrEventWithdraw = "WithdrawDone" // could be LP or user
	// from signers.sol
	CbrEventNewSigners = "SignersUpdated"
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
	*ethclient.Client
	mon      *monitor.Service
	contract *cbrContract
}

// key is chainid
type CbrMgr map[uint64]*CbrOneChain

// for each chain, dial gw, newprefixdb, newWatchDAL, monitor
func NewCbridgeMgr(db dbm.DB) CbrMgr {
	var mcc []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		log.Fatalln("fail to load multichain configs err:", err)
	}
	// setup db/dal, shared by all chains
	cbrDb := dbm.NewPrefixDB(db, []byte(cbrDbPrefix))
	watcherDal := newWatcherDAL(cbrDb) // TODO: watcherDAL concurrency?
	ret := make(CbrMgr)
	ethChainID := viper.GetUint64(common.FlagEthChainId)
	for _, onecfg := range mcc {
		fixCfg(onecfg, ethChainID) // if cfg.chainid equals ethchainid, uses eth.xxx
		log.Infof("Add cbridge chain: %+v", onecfg)
		ret[onecfg.ChainID] = newOneChain(onecfg, watcherDal)
	}
	return ret
}

func newOneChain(cfg *common.OneChainConfig, wdal *watcherDAL) *CbrOneChain {
	ec, err := ethclient.Dial(cfg.Gateway)
	if err != nil {
		log.Fatalln("dial", cfg.Gateway, "err:", err)
	}
	wsvc := watcher.NewWatchService(ec, wdal, cfg.BlkInterval, cfg.MaxBlkDelta)
	mon := monitor.NewService(wsvc, cfg.BlkDelay, true)
	mon.Init()
	cbr, err := eth.NewBridge(eth.Hex2Addr(cfg.CBridge), ec)
	if err != nil {
		log.Fatalln("cbridge contract at", cfg.CBridge, "err:", err)
	}
	ret := &CbrOneChain{
		Client: ec,
		mon:    mon,
		contract: &cbrContract{
			Bridge:  cbr,
			Address: eth.Hex2Addr(cfg.CBridge),
		},
	}
	ret.startMon()
	return ret
}

func fixCfg(cfg *common.OneChainConfig, ethchainid uint64) {
	if cfg.ChainID != ethchainid {
		return
	}
	if cfg.Gateway == "" {
		cfg.Gateway = viper.GetString(common.FlagEthGateway)
	}
	cfg.BlkDelay = viper.GetUint64(common.FlagEthBlockDelay)
	cfg.BlkInterval = viper.GetUint64(common.FlagEthPollInterval)
	cfg.MaxBlkDelta = viper.GetUint64(common.FlagEthMaxBlockDelta)
}
