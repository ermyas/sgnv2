package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/celer-network/endpoint-proxy/endpointproxy"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	dbm "github.com/tendermint/tm-db"
)

// NOTE: to keep cbridge related as independent as possible, we create another client for eth mainnet
// and only use it for cbridge related monitoring

// multichain support for cbridge, including eth client,
// monitor, transactor etc for each chain.

const (
	cbrDbPrefix = "cbr-"
)

// ethclient etc
type CbrOneChain struct {
	*ethclient.Client
	*ethutils.Transactor
	mon          *mon2.Monitor
	cbrContract  *eth.BridgeContract
	pegContracts *PegContracts
	wdiContract  *eth.WdInboxContract
	msgContract  *eth.MsgBusContract
	db           *dbm.PrefixDB // cbr-xxx xxx is chainid
	curss        currentSigners
	lock         sync.RWMutex

	// chainid and blkdelay and forwardblkdelay for verify/easy logging
	chainid, blkDelay, forwardBlkDelay, blkInterval uint64

	checkIntervals map[string]uint64

	// all in one helper for flow chain, only set when IsFlowChain(chainid) is true
	// above eth related fields are all nil
	*FlowClient
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
	watcherDal := newWatcherDAL(cbrDb)
	ret := make(CbrMgr)
	for _, onecfg := range mcc {
		log.Infof("Add cbridge chain: %+v", onecfg)
		ret[onecfg.ChainID] = newOneChain(onecfg, watcherDal, cbrDb, cliCtx)
	}
	CbrMgrInstance = ret
	return ret
}

// return CbrOneChain for flow, no eth stuff
func newOneChainForFlow(cfg *common.OneChainConfig, wdal *watcherDAL, cbrDb *dbm.PrefixDB) *CbrOneChain {
	db := dbm.NewPrefixDB(cbrDb, []byte(fmt.Sprintf("%d", cfg.ChainID)))
	return &CbrOneChain{
		chainid:         cfg.ChainID,
		blkDelay:        cfg.BlkDelay,
		forwardBlkDelay: cfg.ForwardBlkDelay,
		blkInterval:     cfg.BlkInterval,
		FlowClient:      NewFlowClient(cfg, wdal, db),
		db:              db, // do we need to set db here for flow?
	}
}

func newOneChain(cfg *common.OneChainConfig, wdal *watcherDAL, cbrDb *dbm.PrefixDB, cliCtx client.Context) *CbrOneChain {
	if commontypes.IsFlowChain(cfg.ChainID) {
		return newOneChainForFlow(cfg, wdal, cbrDb) // cbrDb to save events in monitor callback
	}
	var ec *ethclient.Client
	var err error
	if cfg.ProxyPort > 0 {
		if err = endpointproxy.StartProxy(cfg.Gateway, cfg.ChainID, cfg.ProxyPort); err != nil {
			log.Fatalln("can not start proxy for chain:", cfg.ChainID, "gateway:", cfg.Gateway, "port:", cfg.ProxyPort, "err:", err)
		}
		ec, err = ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d", cfg.ProxyPort))
		if err != nil {
			log.Fatalln("dial", cfg.Gateway, "err:", err)
		}
	} else {
		ec, err = ethclient.Dial(cfg.Gateway)
		if err != nil {
			log.Fatalln("dial", cfg.Gateway, "err:", err)
		}
	}
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainid %d err: %s", cfg.ChainID, err)
	}
	if chid.Uint64() != cfg.ChainID {
		log.Fatalf("chainid mismatch! cfg has %d but onchain has %d", cfg.ChainID, chid.Uint64())
	}
	cbr, err := eth.NewBridgeContract(eth.Hex2Addr(cfg.CBridge), ec)
	if err != nil {
		log.Fatalln("cbridge contract at", cfg.CBridge, "err:", err)
	}
	pegContracts, err := NewPegContracts(cfg, ec)
	if err != nil {
		log.Fatalln(err)
	}
	wdi, err := eth.NewWdInboxContract(eth.Hex2Addr(cfg.WdInbox), ec)
	if err != nil {
		log.Fatalln("WithdrawInbox contract at", cfg.WdInbox, "err:", err)
	}
	msg, err := eth.NewMsgBusContract(eth.Hex2Addr(cfg.MsgBus), ec)
	if err != nil {
		log.Fatalln("MessageBus contract at", cfg.MsgBus, "err:", err)
	}
	signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	signer, addr, err := eth.CreateSigner(signerKey, signerPass, chid)
	if err != nil {
		log.Fatalln("CreateSigner err:", err)
	}

	transactor := ethutils.NewTransactorByExternalSigner(
		addr,
		signer,
		ec,
		big.NewInt(int64(cfg.ChainID)),
		ethutils.WithBlockDelay(cfg.BlkDelay),
		ethutils.WithPollingInterval(time.Duration(cfg.BlkInterval)*time.Second*4),
		ethutils.WithAddGasEstimateRatio(cfg.AddGasEstimateRatio),
		ethutils.WithGasLimit(cfg.GasLimit),
		ethutils.WithAddGasGwei(cfg.AddGasGwei),
		ethutils.WithMaxGasGwei(cfg.MaxGasGwei),
		ethutils.WithMinGasGwei(cfg.MinGasGwei),
		ethutils.WithMaxFeePerGasGwei(cfg.MaxFeePerGasGwei),
		ethutils.WithMaxPriorityFeePerGasGwei(cfg.MaxPriorityFeePerGasGwei),
	)

	checkIntervals := make(map[string]uint64)
	for name, interval := range cfg.CheckInterval {
		checkIntervals[name] = uint64(interval.(int64))
	}

	ret := &CbrOneChain{
		Client:          ec,
		Transactor:      transactor,
		cbrContract:     cbr,
		pegContracts:    pegContracts,
		wdiContract:     wdi,
		msgContract:     msg,
		db:              dbm.NewPrefixDB(cbrDb, []byte(fmt.Sprintf("%d", cfg.ChainID))),
		chainid:         cfg.ChainID,
		blkDelay:        cfg.BlkDelay,
		forwardBlkDelay: cfg.ForwardBlkDelay,
		blkInterval:     cfg.BlkInterval,
		checkIntervals:  checkIntervals,
	}
	ret.mon, err = mon2.NewMonitor(ec, wdal, mon2.PerChainCfg{
		BlkIntv:         time.Duration(cfg.BlkInterval) * time.Second,
		BlkDelay:        cfg.BlkDelay,
		MaxBlkDelta:     cfg.MaxBlkDelta,
		ForwardBlkDelay: cfg.ForwardBlkDelay,
	})
	if err != nil {
		log.Fatalln("failed to create monitor, err:", err)
	}
	chainSigners, err := cbrcli.QueryChainSigners(cliCtx, cfg.ChainID)
	if err != nil {
		errmsg := fmt.Sprintf("failed to get chain %d signers: %s", cfg.ChainID, err)
		if strings.Contains(err.Error(), "key not found") {
			log.Warn(errmsg)
		} else {
			log.Error(errmsg)
		}
	} else {
		log.Infof("Set chain %d signers %s:", cfg.ChainID, chainSigners.String())
		ret.setCurss(chainSigners.GetSortedSigners())
	}
	return ret
}

type RelayRequest struct {
	XferId     []byte    `json:"xfer_id"` // src transfer id
	RetryCount uint64    `json:"retry_count"`
	DstChainId uint64    `json:"dst_chain_id"`
	CreateTime time.Time `json:"create_time"`
}

func NewRelayRequest(xferId []byte, dstChainId uint64) RelayRequest {
	return RelayRequest{
		XferId:     xferId,
		RetryCount: 0,
		DstChainId: dstChainId,
		CreateTime: time.Now(),
	}
}

func NewRelayRequestFromBytes(input []byte) RelayRequest {
	relay := RelayRequest{}
	relay.MustUnMarshal(input)
	return relay
}

// Marshal RelayRequest into json bytes
func (r RelayRequest) MustMarshal() []byte {
	res, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}

	return res
}

// Unmarshal json bytes to RelayRequest
func (r *RelayRequest) MustUnMarshal(input []byte) {
	err := json.Unmarshal(input, r)
	if err != nil {
		panic(err)
	}
}

type currentSigners struct {
	addrs  []eth.Addr
	powers []*big.Int
}

func (s currentSigners) String() string {
	var out string
	for i, addr := range s.addrs {
		out += fmt.Sprintf("<addr %x power %s> ", addr, s.powers[i])
	}
	return fmt.Sprintf("< %s>", out)
}

func (m CbrMgr) ForEach(run func(*CbrOneChain)) {
	for _, onech := range m {
		run(onech)
	}
}
