package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	gobig "github.com/celer-network/goutils/big"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Addr = common.Address

func hex2addr(addr string) Addr {
	return common.HexToAddress(addr)
}

// no 0x prefix, only hex, all lower case
func a2hex(addr Addr) string {
	return hex.EncodeToString(addr[:])
}

func hex2hash(hexstr string) common.Hash {
	return common.HexToHash(hexstr)
}

// status in db and return to web
type XferStatus int16

const (
	Status_NOTFOUND XferStatus = iota // havn't seen src event yet
	// valid status
	Status_WAITSGN // saw sent event, wait for sgn to get msg ready. set when insert new xfer into db
	Status_WAITDST // get msg from sgn, and submitted onchain. set
	Status_DONE    // after seen event from dest chain tx
)

// setup monitor per chain and send tx
type OneChain struct {
	cfg  *OneChainConfig
	ec   *ethclient.Client
	auth *bind.TransactOpts // send tx onchain
	mon  *mon2.Monitor
	db   *dal.DAL

	msgBus *MsgBusRecv        // contract binding for send onchain tx
	nftbr  *NFTBridgeFilterer // to parse event into object, set in MonNftBridge
}

// return err if dial fail or chainid mismatch
func NewOneChain(cfg OneChainConfig, kspath, pass string, dal *dal.DAL) (*OneChain, error) {
	ret := &OneChain{
		cfg: &cfg,
		db:  dal,
	}
	var err error
	ret.ec, err = ethclient.Dial(cfg.Gateway)
	if err != nil {
		return nil, err
	}
	bgCtx := context.Background()
	chid, _ := ret.ec.ChainID(bgCtx)
	if chid.Uint64() != cfg.ChainID {
		return nil, fmt.Errorf("mismatch chainid cfg: %d, rpc: %d", cfg.ChainID, chid.Uint64())
	}
	ret.auth, err = kspath2auth(kspath, pass, chid)
	if err != nil {
		return nil, err
	}
	ret.mon, err = mon2.NewMonitor(ret.ec, dal, mon2.PerChainCfg{
		BlkIntv:         time.Duration(cfg.BlkInterval) * time.Second,
		BlkDelay:        cfg.BlkDelay,
		MaxBlkDelta:     cfg.MaxBlkDelta,
		ForwardBlkDelay: cfg.ForwardBlkDelay,
	})
	if err != nil {
		return nil, err
	}
	ret.msgBus, _ = NewMsgBusRecv(hex2addr(cfg.MsgBus), ret.ec)
	return ret, nil
}

func (c *OneChain) Close() {
	c.mon.Close()
}

func (c *OneChain) MonNftBridge(addr string) {
	nftbrAddr := hex2addr(addr)
	c.nftbr, _ = NewNFTBridgeFilterer(nftbrAddr, nil)
	go c.mon.MonAddr(mon2.PerAddrCfg{
		Addr:    nftbrAddr,
		ChkIntv: time.Minute,  // getlog every minute
		AbiStr:  NFTBridgeABI, // to parse event name by topics[0]
	}, c.evCallback)
}

func (c *OneChain) evCallback(evname string, elog types.Log) {
	log.Infoln("event:", c.cfg.ChainID, evname)
	switch evname {
	case "Sent":
		evSent, err := c.nftbr.ParseSent(elog)
		if err != nil {
			log.Error("parse sent err:", err, elog)
		} else {
			c.handleSent(evSent)
		}
	case "Received":
		evRecv, err := c.nftbr.ParseReceived(elog)
		if err != nil {
			log.Error("parse recv err:", err, elog)
		} else {
			c.handleRecv(evRecv)
		}
	default:
		log.Error("unsupported evname: ", evname)
		return
	}
}

func (c *OneChain) handleSent(ev *NFTBridgeSent) {
	err := c.db.DoTx(func(tx *sql.Tx) error {
		dtx := dal.New(tx)
		return dtx.NftAddSend(context.Background(), dal.NftAddSendParams{
			CreatedAt: time.Now().Unix(),
			SrcChid:   c.cfg.ChainID,
			DstChid:   ev.DstChid,
			Sender:    a2hex(ev.Sender),
			Receiver:  a2hex(ev.Receiver),
			SrcNft:    a2hex(ev.SrcNft),
			DstNft:    a2hex(ev.DstNft),
			TokID:     *gobig.New(ev.Id),
			SrcTx:     ev.Raw.TxHash.Hex(),
		})
	})
	if err != nil {
		log.Error(ev, err)
	}
}

func (c *OneChain) handleRecv(ev *NFTBridgeReceived) {
	err := c.db.DoTx(func(tx *sql.Tx) error {
		return dal.New(tx).NftSetDoneByDstTx(context.Background(), ev.Raw.TxHash.Hex())
	})
	if err != nil {
		log.Error(ev, err)
	}
}

func kspath2auth(kspath, pass string, chainid *big.Int) (*bind.TransactOpts, error) {
	ksjson, err := os.ReadFile(kspath)
	if err != nil {
		return nil, err
	}
	return bind.NewTransactorWithChainID(strings.NewReader(string(ksjson)), pass, chainid)
}
