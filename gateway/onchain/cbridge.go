package onchain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/endpoint-proxy/endpointproxy"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

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

type OneChain struct {
	*ethclient.Client
	*ethutils.Transactor
	mon      *monitor.Service
	contract *cbrContract

	// chainid and blkdelay and forwardblkdelay for verify/easy logging
	chainid, blkDelay, forwardBlkDelay uint64
}

// key is chainid
type ChainMgr map[uint64]*OneChain

var Chains ChainMgr

func InitChainMgr(db *dal.DAL) ChainMgr {
	var mcc []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		log.Fatalln("fail to load multichain configs err:", err)
	}
	// watcherDal is shared because monitor adds chainID automatically
	watcherDal := NewWatcherDAL(db)
	ret := make(ChainMgr)
	for _, onecfg := range mcc {
		log.Infof("Add cbridge chain: %+v", onecfg)
		ret[onecfg.ChainID] = newOneChain(onecfg, watcherDal)
	}
	Chains = ret
	return ret
}

func (chains ChainMgr) GetEthClient(chid uint64) *ethclient.Client {
	chain, ok := chains[chid]
	if !ok {
		return nil
	}
	return chain.Client
}

func newOneChain(chainConf *common.OneChainConfig, wdal *watcherDAL) *OneChain {
	log.Infoln("Dialing eth client at", chainConf.Gateway)
	var ec *ethclient.Client
	var err error
	if chainConf.ProxyPort > 0 {
		if err = endpointproxy.StartProxy(chainConf.Gateway, chainConf.ChainID, chainConf.ProxyPort); err != nil {
			log.Fatalln("can not start proxy for chain:", chainConf.ChainID, "gateway:", chainConf.Gateway, "port:", chainConf.ProxyPort, "err:", err)
		}
		ec, err = ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d", chainConf.ProxyPort))
		if err != nil {
			log.Fatalln("dial", chainConf.Gateway, "err:", err)
		}
	} else {
		ec, err = ethclient.Dial(chainConf.Gateway)
		if err != nil {
			log.Fatalln("dial", chainConf.Gateway, "err:", err)
		}
	}
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainid %d err: %s", chainConf.ChainID, err)
	}
	if chid.Uint64() != chainConf.ChainID {
		log.Fatalf("chainid mismatch! chainConf has %d but onchain has %d", chainConf.ChainID, chid.Uint64())
	}
	wsvc := watcher.NewWatchService(ec, wdal, chainConf.BlkInterval, chainConf.MaxBlkDelta)
	mon := monitor.NewService(wsvc, chainConf.BlkDelay, true)
	mon.Init()
	cbr, err := eth.NewBridge(eth.Hex2Addr(chainConf.CBridge), ec)
	if err != nil {
		log.Fatalln("cbridge contract at", chainConf.CBridge, "err:", err)
	}
	ret := &OneChain{
		mon:    mon,
		Client: ec,
		contract: &cbrContract{
			Address: eth.Hex2Addr(chainConf.CBridge),
			Bridge:  cbr,
		},
		chainid:         chainConf.ChainID,
		blkDelay:        chainConf.BlkDelay,
		forwardBlkDelay: chainConf.ForwardBlkDelay,
	}
	go ret.startMon()
	return ret
}

// funcs for monitor cbridge events
func (c *OneChain) startMon() {
	smallDelay := func() {
		time.Sleep(100 * time.Millisecond)
	}
	blkNum := c.mon.GetCurrentBlockNumber()
	c.monSend(blkNum)
	smallDelay()
	c.monRelay(blkNum)
	smallDelay()
	c.monLiqAdd(blkNum)
	smallDelay()
	c.monWithdraw(blkNum)
	smallDelay()
	c.monDelayXferAdd(blkNum)
	smallDelay()
	c.monDelayXferExec(blkNum)
}

func (c *OneChain) monSend(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventSend,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseSend(eLog)
		if err != nil {
			log.Errorln("monSend: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = GatewayOnSend(eth.Hash(ev.TransferId).String(), ev.Sender.String(), ev.Token.String(), ev.Amount.String(), eLog.TxHash.String(), c.chainid, ev.DstChainId)
		if err != nil {
			log.Warnf("GatewayOnSend err: %s, txId %x, txHash %x, chainId %d", err, ev.TransferId, eLog.TxHash, c.chainid)
			return true
		}
		return false
	})
}

func (c *OneChain) monRelay(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventRelay,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseRelay(eLog)
		if err != nil {
			log.Errorln("monRelay: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = GatewayOnRelay(c.Client, eth.Hash(ev.SrcTransferId).String(), eLog.TxHash.String(), eth.Hash(ev.TransferId).String(), ev.Amount.String())
		if err != nil {
			log.Warnf("UpdateTransfer err: %s, srcId %x, dstId %x, txHash %x, chainId %d", err, ev.SrcTransferId, ev.TransferId, eLog.TxHash, c.chainid)
		}
		return false
	})
}

func (c *OneChain) monDelayXferAdd(blk *big.Int) {
	blkDelay := c.blkDelay / 2
	if blkDelay < 1 {
		blkDelay = 1
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventDelayXferAdd,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
		// much lower than the blk delay in cfg because the gateway service needs "Relay"
		// to be preceded by "DelayedTransferAdded".
		// this is ok because the the result action of monitoring "DelayedTransferAdded"
		// only changes the status for display use and is not related to fund safety.
		BlockDelay: blkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseDelayedTransferAdded(eLog)
		if err != nil {
			log.Errorln("monRelay: cannot parse event:", err)
			return false
		}
		idstr := eth.Hash(ev.Id).String()
		log.Infof("MonEv: DelayedTransferAdded chainId: %d, tx: %s, id %s", c.chainid, eLog.TxHash.String(), idstr)
		err = GatewayOnDelayXferAdd(idstr, eLog.TxHash.String())
		if err != nil {
			log.Errorln("GatewayOnDelayXferAdd err:", err)
		}
		return false
	})
}

func (c *OneChain) monDelayXferExec(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventDelayXferExec,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseDelayedTransferExecuted(eLog)
		if err != nil {
			log.Errorln("monRelay: cannot parse event:", err)
			return false
		}
		idstr := eth.Hash(ev.Id).String()
		log.Infof("MonEv: DelayedTransferExecuted chainId: %d, tx: %s, id %s", c.chainid, eLog.TxHash.String(), idstr)
		GatewayOnDelayXferExec(idstr, eLog.TxHash.String())
		if err != nil {
			log.Errorln("DelayedTransferExecuted err:", err)
		}
		return false
	})
}

func (c *OneChain) monLiqAdd(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventLiqAdd,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseLiquidityAdded(eLog)
		if err != nil {
			log.Errorln("monLiqAdd: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		nonce := uint64(0)
		tx, _, err := c.TransactionByHash(context.Background(), eLog.TxHash)
		if tx != nil && err == nil {
			nonce = tx.Nonce()
		} else {
			log.Warnf("get nonce failed, use ts:%d instead, TxHash:%s, err: %s", nonce, eLog.TxHash.String(), err)
		}
		err = GatewayOnLiqAdd(ev.Provider.String(), ev.Token.String(), ev.Amount.String(), eLog.TxHash.String(), c.chainid, ev.Seqnum, nonce)
		if err != nil {
			log.Warnf("UpsertLP err: %s, seqNum %d, amt %s, txHash %x, chainId %d", err, ev.Seqnum, ev.Amount.String(), eLog.TxHash, c.chainid)
			return false
		}
		return false
	})
}

func (c *OneChain) monWithdraw(blk *big.Int) {
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventWithdraw,
		Contract:     c.contract,
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.contract.ParseWithdrawDone(eLog)
		if err != nil {
			log.Errorln("monWithdraw: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		idstr := eth.Hash(ev.WithdrawId).String()
		GatewayOnLiqWithdraw(idstr, eLog.TxHash.String(), c.chainid, ev.Seqnum, ev.Receiver.String())
		return false
	})
}
