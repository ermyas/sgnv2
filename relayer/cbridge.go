package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"sync"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
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

type sortedSigners struct {
	signers *cbrtypes.SortedSigners
	bytes   []byte
	updated bool
}

func (s *sortedSigners) setSigners(bytes []byte) {
	s.bytes = bytes
	s.signers = new(cbrtypes.SortedSigners)
	s.signers.Unmarshal(bytes)
}

// ethclient etc
type CbrOneChain struct {
	*ethclient.Client
	*ethutils.Transactor
	mon      *monitor.Service
	contract *cbrContract
	db       *dbm.PrefixDB // cbr-xxx xxx is chainid
	curss    *sortedSigners
	lock     sync.RWMutex

	// not required for flow but make log chain id easy
	chainid uint64
}

// key is chainid
type CbrMgr map[uint64]*CbrOneChain

var CbrMgrInstance CbrMgr

// for each chain, dial gw, newprefixdb, newWatchDAL, monitor
func NewCbridgeMgr(db dbm.DB, cliCtx client.Context) CbrMgr {
	var mcc []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		log.Fatalln("fail to load multichain configs err:", err)
	}
	// setup db/dal, shared by all chains
	cbrDb := dbm.NewPrefixDB(db, []byte(cbrDbPrefix))
	// watcherDal is shared because monitor adds chainID automatically
	watcherDal := newWatcherDAL(cbrDb) // TODO: watcherDAL concurrency?
	ret := make(CbrMgr)
	ethChainID := viper.GetUint64(common.FlagEthChainId)
	for _, onecfg := range mcc {
		fixCfg(onecfg, ethChainID) // if cfg.chainid equals ethchainid, uses eth.xxx
		log.Infof("Add cbridge chain: %+v", onecfg)
		ret[onecfg.ChainID] = newOneChain(onecfg, watcherDal, cbrDb, cliCtx)
	}
	CbrMgrInstance = ret
	return ret
}

func newOneChain(cfg *common.OneChainConfig, wdal *watcherDAL, cbrDb *dbm.PrefixDB, cliCtx client.Context) *CbrOneChain {
	ec, err := ethclient.Dial(cfg.Gateway)
	if err != nil {
		log.Fatalln("dial", cfg.Gateway, "err:", err)
	}
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainid %d err: %s", cfg.ChainID, err)
	}
	if chid.Uint64() != cfg.ChainID {
		log.Fatalf("chainid mismatch! cfg has %d but onchain has %d", cfg.ChainID, chid.Uint64())
	}
	wsvc := watcher.NewWatchService(ec, wdal, cfg.BlkInterval, cfg.MaxBlkDelta)
	mon := monitor.NewService(wsvc, cfg.BlkDelay, true)
	mon.Init()
	cbr, err := eth.NewBridge(eth.Hex2Addr(cfg.CBridge), ec)
	if err != nil {
		log.Fatalln("cbridge contract at", cfg.CBridge, "err:", err)
	}

	ksBytes, err := ioutil.ReadFile(viper.GetString(common.FlagEthSignerKeystore))
	if err != nil {
		log.Fatalln("ReadFile err:", err)
	}
	transactor, err := ethutils.NewTransactor(
		string(ksBytes),
		viper.GetString(common.FlagEthSignerPassphrase),
		ec,
		big.NewInt(int64(cfg.ChainID)),
		ethutils.WithBlockDelay(cfg.BlkDelay),
		ethutils.WithPollingInterval(time.Duration(cfg.BlkInterval)*time.Second),
	)
	if err != nil {
		log.Fatalln("NewTransactor err:", err)
	}
	ret := &CbrOneChain{
		Client:     ec,
		Transactor: transactor,
		mon:        mon,
		contract: &cbrContract{
			Bridge:  cbr,
			Address: eth.Hex2Addr(cfg.CBridge),
		},
		db: dbm.NewPrefixDB(cbrDb, []byte(fmt.Sprintf("%d", cfg.ChainID))),
		curss: &sortedSigners{
			signers: &cbrtypes.SortedSigners{},
		},
		chainid: cfg.ChainID,
	}
	chainSigners, err := cbrcli.QueryChainSigners(cliCtx, cfg.ChainID)
	if err != nil {
		log.Warnf("failed to get chain %d signers: %s", cfg.ChainID, err)
	} else {
		log.Infof("Set chain %d signers %s:", cfg.ChainID, chainSigners.String())
		ret.setCurss(chainSigners.GetSignersBytes())
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

type RelayEvent struct {
	XferId     []byte `json:"xfer_id"`
	RetryCount uint64 `json:"retry_count"`
}

func NewRelayEvent(xferId []byte) RelayEvent {
	return RelayEvent{
		XferId:     xferId,
		RetryCount: 0,
	}
}

func NewRelayEventFromBytes(input []byte) RelayEvent {
	event := RelayEvent{}
	event.MustUnMarshal(input)
	return event
}

// Marshal event into json bytes
func (e RelayEvent) MustMarshal() []byte {
	res, err := json.Marshal(&e)
	if err != nil {
		panic(err)
	}

	return res
}

// Unmarshal json bytes to relay event
func (e *RelayEvent) MustUnMarshal(input []byte) {
	err := json.Unmarshal(input, e)
	if err != nil {
		panic(err)
	}
}
