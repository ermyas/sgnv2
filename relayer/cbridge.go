package relayer

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/celer-network/endpoint-proxy/endpointproxy"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
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

// ethclient etc
type CbrOneChain struct {
	*ethclient.Client
	*ethutils.Transactor
	mon      *monitor.Service
	contract *cbrContract
	db       *dbm.PrefixDB // cbr-xxx xxx is chainid
	curss    currentSigners
	lock     sync.RWMutex

	// chainid and blkdelay and forwardblkdelay for verify/easy logging
	chainid, blkDelay, forwardBlkDelay, blkInterval uint64
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

func newOneChain(cfg *common.OneChainConfig, wdal *watcherDAL, cbrDb *dbm.PrefixDB, cliCtx client.Context) *CbrOneChain {
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
	wsvc := watcher.NewWatchService(ec, wdal, cfg.BlkInterval, cfg.MaxBlkDelta)
	mon := monitor.NewService(wsvc, cfg.BlkDelay, true)
	mon.Init()
	cbr, err := eth.NewBridge(eth.Hex2Addr(cfg.CBridge), ec)
	if err != nil {
		log.Fatalln("cbridge contract at", cfg.CBridge, "err:", err)
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
		ethutils.WithPollingInterval(time.Duration(cfg.BlkInterval)*time.Second),
		ethutils.WithAddGasEstimateRatio(cfg.AddGasEstimateRatio),
		ethutils.WithAddGasGwei(cfg.AddGasGwei),
		ethutils.WithMaxFeePerGasGwei(cfg.MaxFeePerGasGwei),
	)
	ret := &CbrOneChain{
		Client:     ec,
		Transactor: transactor,
		mon:        mon,
		contract: &cbrContract{
			Bridge:  cbr,
			Address: eth.Hex2Addr(cfg.CBridge),
		},
		db:              dbm.NewPrefixDB(cbrDb, []byte(fmt.Sprintf("%d", cfg.ChainID))),
		chainid:         cfg.ChainID,
		blkDelay:        cfg.BlkDelay,
		forwardBlkDelay: cfg.ForwardBlkDelay,
		blkInterval:     cfg.BlkInterval,
	}
	chainSigners, err := cbrcli.QueryChainSigners(cliCtx, cfg.ChainID)
	if err != nil {
		log.Warnf("failed to get chain %d signers: %s", cfg.ChainID, err)
	} else {
		log.Infof("Set chain %d signers %s:", cfg.ChainID, chainSigners.String())
		ret.setCurss(chainSigners.GetSortedSigners())
	}
	ret.startMon()
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
