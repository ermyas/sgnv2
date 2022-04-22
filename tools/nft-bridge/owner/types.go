package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	gobig "github.com/celer-network/goutils/big"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	nftbr "github.com/celer-network/sgn-v2/tools/nft-bridge"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/binding"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const defaultPollSec = 30

var ZeroAddr nftbr.Addr

// nfts.toml file
type NftList []NftCfg

// one nft entry in nfts.toml
type NftCfg struct {
	// Startblk only effective if db is empty
	// Pollintv can be small if in catch up mode, then change to normal
	Chainid, Startblk, Pollintv uint64
	// nft address
	Contract string
}

// setup monitor per chain, no send tx
type OneChain struct {
	cfg *nftbr.OneChainConfig
	ec  *ethclient.Client
	mon *mon2.Monitor
	db  *dal.DAL

	nft *binding.OrigNFTFilterer
}

func NewOneChain(cfg nftbr.OneChainConfig, dal *dal.DAL) (*OneChain, error) {
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
	ret.mon, err = mon2.NewMonitor(ret.ec, dal, mon2.PerChainCfg{
		BlkIntv:         time.Duration(cfg.BlkInterval) * time.Second,
		BlkDelay:        cfg.BlkDelay,
		MaxBlkDelta:     cfg.MaxBlkDelta,
		ForwardBlkDelay: cfg.ForwardBlkDelay,
	})
	if err != nil {
		return nil, err
	}
	ret.nft, _ = binding.NewOrigNFTFilterer(ZeroAddr, nil)
	return ret, nil
}

func (c *OneChain) Close() {
	c.mon.Close()
}

// startblk is only set if db is empty
func (c *OneChain) MonNft(n NftCfg) {
	nftAddr := nftbr.Hex2addr(n.Contract)
	key := fmt.Sprintf("%d-%x", c.cfg.ChainID, nftAddr)
	_, _, found, _ := c.db.GetMonitorBlock(key)
	pollsec := n.Pollintv
	if pollsec == 0 {
		pollsec = defaultPollSec
	}
	perAddrCfg := mon2.PerAddrCfg{
		Addr:    nftAddr,
		ChkIntv: time.Second * time.Duration(pollsec),
		AbiStr:  binding.OrigNFTABI,
	}
	if !found {
		// not in db, use startblk instead of current head
		perAddrCfg.FromBlk = n.Startblk
	}
	go c.mon.MonAddr(perAddrCfg, c.evCallback)
}

func (c *OneChain) evCallback(evname string, elog types.Log) {
	switch evname {
	case "Transfer":
		ev, err := c.nft.ParseTransfer(elog)
		if err != nil {
			log.Error("parse sent err:", err, elog)
		} else {
			c.handleTransfer(elog.Address, ev)
		}
	default:
		// ignore other events for now
		return
	}
}

func (c *OneChain) handleTransfer(nft nftbr.Addr, ev *binding.OrigNFTTransfer) {
	chid := c.cfg.ChainID
	nftStr, fromStr, toStr := nftbr.A2hex(nft), nftbr.A2hex(ev.From), nftbr.A2hex(ev.To)
	tokIdStr := ev.TokenId.String()
	log.Infoln(chid, ev.Raw.BlockNumber, nftStr, fromStr, toStr, tokIdStr)
	var errCtx string
	doTxErr := c.db.DoTx(func(tx *sql.Tx) error {
		dtx := dal.New(tx)
		err := dtx.AllEvsAdd(context.Background(), dal.AllEvsAddParams{
			Chid:     chid,
			Nft:      nftStr,
			FromAddr: fromStr,
			ToAddr:   toStr,
			TokID:    *gobig.New(ev.TokenId),
		})
		if err != nil {
			errCtx = "AllEvsAdd"
			return err
		}
		// now update usrnft table
		if ev.From != ZeroAddr {
			// remove tokid
			has, err := dtx.UsrGetNfts(context.Background(), dal.UsrGetNftsParams{
				Chid: chid,
				Nft:  nftStr,
				Usr:  fromStr,
			})
			if err != nil {
				errCtx = "Get ev.From Nfts"
				return err
			}
			found := -1
			for i, tokid := range has {
				if tokid == tokIdStr {
					found = i
				}
			}
			if found == -1 {
				return fmt.Errorf("%s not found in from %v", tokIdStr, has)
			} else {
				// delete from has
				has = append(has[:found], has[found+1:]...)
				err = dtx.UsrSetNfts(context.Background(), dal.UsrSetNftsParams{
					Chid:   chid,
					Nft:    nftStr,
					Usr:    fromStr,
					Tokens: has,
				})
				if err != nil {
					errCtx = "Set ev.From Nfts"
					return err
				}
			}
		}

		if ev.To != ZeroAddr {
			// add tokid
			var alreadyHas bool
			// has could be empty if (chid, nft, usr) not found
			has, err := dtx.UsrGetNfts(context.Background(), dal.UsrGetNftsParams{
				Chid: chid,
				Nft:  nftStr,
				Usr:  toStr,
			})
			_, err2 := dal.ChkQueryRow(err) // if not found, don't consider as error
			if err2 != nil {
				errCtx = "Get ev.To Nfts"
				return err2
			}
			// has could be empty
			for _, tokid := range has {
				if tokid == tokIdStr {
					alreadyHas = true
				}
			}
			if !alreadyHas {
				has = append(has, tokIdStr)
				err = dtx.UsrSetNfts(context.Background(), dal.UsrSetNftsParams{
					Chid:   chid,
					Nft:    nftStr,
					Usr:    toStr,
					Tokens: has,
				})
				if err != nil {
					errCtx = "Set ev.To Nfts"
					return err
				}
			}
		}

		return nil
	})
	if doTxErr != nil {
		log.Error(errCtx, doTxErr)
	}
}
