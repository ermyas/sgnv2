package onchain

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/sgn-v2/relayer"

	"github.com/celer-network/endpoint-proxy/endpointproxy"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	pegtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
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
	mon          *monitor.Service
	contract     *cbrContract
	pegContracts *relayer.PegContracts
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

func (chains ChainMgr) GetOneChain(chid uint64) (*OneChain, bool) {
	chain, found := chains[chid]
	if found {
		return chain, true
	}
	return nil, false
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
		pegContracts:    &relayer.PegContracts{},
		chainid:         chainConf.ChainID,
		blkDelay:        chainConf.BlkDelay,
		forwardBlkDelay: chainConf.ForwardBlkDelay,
	}

	if chainConf.OTVault != "" {
		vault, err := eth.NewPegVaultContract(eth.Hex2Addr(chainConf.OTVault), ec)
		if err != nil {
			log.Fatalln("OriginalTokenVaults contract at", chainConf.OTVault, "err:", err)
		}
		ret.pegContracts.SetPegVaultContract(vault)
	}

	if chainConf.PTBridge != "" {
		pegBridge, err := eth.NewPegBridgeContract(eth.Hex2Addr(chainConf.PTBridge), ec)
		if err != nil {
			log.Fatalln("PeggedTokenBridge contract at", chainConf.PTBridge, "err:", err)
		}
		ret.pegContracts.SetegBridgeContract(pegBridge)
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
	smallDelay()

	// pegged event monitor
	if c.GetOtvContract() != nil {
		c.monPegbrDeposited(blkNum)
		smallDelay()
		c.monPegbrWithdrawn(blkNum)
		smallDelay()
		c.monPegbrDelayWithdrawAdd(blkNum)
		smallDelay()
		c.monPegbrDelayWithdrawExec(blkNum)
		smallDelay()
	}

	if c.GetPtbContract() != nil {
		c.monPegbrMint(blkNum)
		smallDelay()
		c.monPegbrBurn(blkNum)
		smallDelay()
		c.monPegbrDelayMintAdd(blkNum)
		smallDelay()
		c.monPegbrDelayMintExec(blkNum)
		smallDelay()
	}
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

		err = GatewayOnRelay(c.Client, eth.Hash(ev.SrcTransferId).String(), eLog.TxHash.String(), eth.Hash(ev.TransferId).String(), ev.Amount.String(), ev.Receiver.String(), ev.Token.String(), ev.SrcChainId, c.chainid)
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

func (c *OneChain) monPegbrDeposited(blk *big.Int) {
	if c.GetOtvContract().Address == eth.ZeroAddr {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventDeposited,
		Contract:     c.GetOtvContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetOtvContract().ParseDeposited(eLog)
		if err != nil {
			log.Errorln("monPegbrDeposited: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		err = GatewayOnSend(eth.Hash(ev.DepositId).String(), ev.Depositor.String(), ev.Token.String(), ev.Amount.String(), eLog.TxHash.String(), c.chainid, ev.MintChainId)
		if err != nil {
			log.Warnf("GatewayOnSend err: %s, txId %x, txHash %x, chainId %d", err, eth.Hash(ev.DepositId).String(), eLog.TxHash, c.chainid)
			return true
		}
		return false
	})
}

func (c *OneChain) monPegbrWithdrawn(blk *big.Int) {
	if c.GetOtvContract().Address == eth.ZeroAddr {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventWithdrawn,
		Contract:     c.GetOtvContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetOtvContract().ParseWithdrawn(eLog)
		if err != nil {
			log.Errorln("monPegbrWithdrawn: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())

		err = GatewayOnRelay(c.Client, eth.Hash(ev.RefId).String(), eLog.TxHash.String(), eth.Hash(ev.WithdrawId).String(), ev.Amount.String(), ev.Receiver.String(), ev.Token.String(), ev.RefChainId, c.chainid)
		if err != nil {
			log.Warnf("UpdateTransfer pegged withdraw err: %s, srcId %x, dstId %x, txHash %x, chainId %d", err, ev.RefId, ev.WithdrawId, eLog.TxHash, c.chainid)
		}
		return false
	})
}

func (c *OneChain) monPegbrDelayWithdrawAdd(blk *big.Int) {
	if c.GetOtvContract().Address == eth.ZeroAddr {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventDelayXferAdd,
		Contract:     c.GetOtvContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetOtvContract().ParseDelayedTransferAdded(eLog)
		if err != nil {
			log.Errorln("monPegbrDelayWithdrawn: cannot parse event:", err)
			return false
		}
		log.Infof("MonEv: delayed pegged withdraw added:%x, tx:%x", ev.Id, eLog.TxHash)
		idstr := eth.Hash(ev.Id).String()
		err = GatewayOnDelayXferAdd(idstr, eLog.TxHash.String())
		if err != nil {
			log.Errorln("GatewayOnDelayXferAdd err:", err)
		}
		return false
	})
}

func (c *OneChain) monPegbrDelayWithdrawExec(blk *big.Int) {
	if c.GetOtvContract().Address == eth.ZeroAddr {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventDelayXferExec,
		Contract:     c.GetOtvContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetOtvContract().ParseDelayedTransferExecuted(eLog)
		if err != nil {
			log.Errorln("monPegbrDelayWithdrawExec: cannot parse event:", err)
			return false
		}
		log.Infof("MonEv: delayed pegged withdraw exec:%x, tx:%x", ev.Id, eLog.TxHash)
		idstr := eth.Hash(ev.Id).String()
		err = GatewayOnDelayXferExec(idstr, eLog.TxHash.String())
		if err != nil {
			log.Errorln("GatewayOnDelayPeggedWithdrawExec err:", err)
		}
		return false
	})
}

func (c *OneChain) monPegbrMint(blk *big.Int) {
	if c.GetPtbContract() != nil {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventMint,
		Contract:     c.GetPtbContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetPtbContract().ParseMint(eLog)
		if err != nil {
			log.Errorln("monPegbrMint: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		err = GatewayOnRelay(c.Client, eth.Hash(ev.RefId).String(), eLog.TxHash.String(), eth.Hash(ev.MintId).String(), ev.Amount.String(), ev.Account.String(), ev.Token.String(), ev.RefChainId, c.chainid)
		if err != nil {
			log.Warnf("UpdateTransfer mint err: %s, srcId %x, dstId %x, txHash %x, chainId %d", err, ev.RefId, ev.MintId, eLog.TxHash, c.chainid)
		}
		return false
	})
}

func (c *OneChain) monPegbrBurn(blk *big.Int) {
	if c.GetPtbContract().Address == eth.ZeroAddr {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    pegtypes.PegbrEventBurn,
		Contract:     c.GetPtbContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetPtbContract().ParseBurn(eLog)
		if err != nil {
			log.Errorln("monPegbrActionBurn: cannot parse event:", err)
			return false
		}
		log.Infoln("MonEv:", ev.PrettyLog(c.chainid), "tx:", eLog.TxHash.String())
		withdrawChainId, foundWithdrawChainId, dbErr := dal.DB.GetWithdrawChainIdByBurnChainIdAndTokenAddr(c.chainid, ev.Token)
		if dbErr != nil {
			log.Errorf("fail to GetWithdrawChainIdByBurnChainIdAndTokenAddr, dbErr:%x", dbErr.Error())
			return true
		}
		if !foundWithdrawChainId {
			// use 0 as default
			log.Errorf("fail to find this withdraw chain, burnChainId:%d, token:%x", c.chainid, ev.Token)
			withdrawChainId = 0
		}

		err = GatewayOnSend(eth.Hash(ev.BurnId).String(), ev.Account.String(), ev.Token.String(), ev.Amount.String(), eLog.TxHash.String(), c.chainid, withdrawChainId)
		if err != nil {
			log.Warnf("GatewayOnSend err: %s, txId %x, txHash %x, chainId %d", err, eth.Hash(ev.BurnId).String(), eLog.TxHash, c.chainid)
			return true
		}
		return false
	})
}

func (c *OneChain) monPegbrDelayMintAdd(blk *big.Int) {
	if c.GetPtbContract().Address == eth.ZeroAddr {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventDelayXferAdd,
		Contract:     c.GetPtbContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetPtbContract().ParseDelayedTransferAdded(eLog)
		if err != nil {
			log.Errorln("monPegbrDelayMint: cannot parse event:", err)
			return false
		}
		log.Infof("MonEv: delayed pegged mint added:%x, tx:%x", ev.Id, eLog.TxHash)
		idstr := eth.Hash(ev.Id).String()
		err = GatewayOnDelayXferAdd(idstr, eLog.TxHash.String())
		if err != nil {
			log.Errorln("GatewayOnDelayXferAdd err:", err)
		}
		return false
	})
}

func (c *OneChain) monPegbrDelayMintExec(blk *big.Int) {
	if c.GetPtbContract().Address == eth.ZeroAddr {
		return
	}
	cfg := &monitor.Config{
		ChainId:      c.chainid,
		EventName:    types.CbrEventDelayXferExec,
		Contract:     c.GetPtbContract(),
		StartBlock:   blk,
		ForwardDelay: c.forwardBlkDelay,
	}
	c.mon.Monitor(cfg, func(id monitor.CallbackID, eLog ethtypes.Log) (recreate bool) {
		ev, err := c.GetPtbContract().ParseDelayedTransferExecuted(eLog)
		if err != nil {
			log.Errorln("monPegbrDelayMintExec: cannot parse event:", err)
			return false
		}
		log.Infof("MonEv: delayed pegged mint exec:%x, tx:%x", ev.Id, eLog.TxHash)
		idstr := eth.Hash(ev.Id).String()
		err = GatewayOnDelayXferExec(idstr, eLog.TxHash.String())
		if err != nil {
			log.Errorln("GatewayOnDelayMintExec err:", err)
		}
		return false
	})
}

func (c *OneChain) GetPtbContract() *eth.PegBridgeContract {
	if c.pegContracts == nil {
		return nil
	}
	return c.pegContracts.GetPegBridgeContract()
}

func (c *OneChain) GetOtvContract() *eth.PegVaultContract {
	if c.pegContracts == nil {
		return nil
	}
	return c.pegContracts.GetPegVaultContract()
}
