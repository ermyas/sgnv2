package main

import (
	"context"
	"flag"

	"github.com/celer-network/goutils/log"

	nftbr "github.com/celer-network/sgn-v2/tools/nft-bridge"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/binding"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

// ops for setting map etc
var (
	fcfg  = flag.String("c", "nftbr_test.toml", "config toml file")
	fjson = flag.String("j", "", "if not empty, read local json file for bridge and nft configs")
	ffix  = flag.Bool("f", false, "if true, send onchain tx to fix mismatch map")
)

func main() {
	flag.Parse()
	viper.SetConfigFile(*fcfg)
	err := viper.ReadInConfig()
	chkErr(err, "viper ReadInConfig")

	jsonkey := "jsonurl"
	if *fjson != "" {
		jsonkey = *fjson // local file name
	}
	jsonCfg := nftbr.GetJsonCfg(jsonkey)
	chid2Bridge := make(map[uint64]common.Address)
	for _, onebr := range jsonCfg.Bridges {
		chid2Bridge[onebr.Chainid] = nftbr.Hex2addr(onebr.Addr)
	}
	// from sym to map[chid]address
	nftMap := make(map[string]map[uint64]string)
	// key is chainid, value is list of nft addresses whose orig is on this chain
	origNft := make(map[uint64][]string)
	for _, nft := range jsonCfg.Nfts {
		nftMap[nft.Symbol] = nft.Map()
		if nft.Orig != nil {
			origNft[nft.Orig.Chainid] = append(origNft[nft.Orig.Chainid], nft.Orig.Addr)
		}
	}

	mcc := nftbr.GetMcc("multichain")
	chainMap := make(map[uint64]*OpOneCh)
	kspath := viper.GetString("kspath")
	ksphrase := viper.GetString("ksphrase")
	for _, cfg := range mcc {
		ooc, err := NewOpOneCh(cfg.Gateway, kspath, ksphrase, chid2Bridge[cfg.ChainID])
		if err != nil {
			log.Errorln("new oponech err:", err)
			continue
		}
		ooc.CheckDstBridgeMap(chid2Bridge, *ffix)
		ooc.CheckNftMap(nftMap, *ffix)
		ooc.CheckOrigNFT(origNft[cfg.ChainID], *ffix)
		chainMap[cfg.ChainID] = ooc
	}
}

// no monitor, only ec and auth
type OpOneCh struct {
	chid  uint64
	ec    *ethclient.Client
	auth  *bind.TransactOpts // send tx onchain
	nftbr *binding.NFTBridge // set in
}

func NewOpOneCh(rpc, kspath, pass string, nftbrA common.Address) (*OpOneCh, error) {
	ret := new(OpOneCh)
	var err error
	ret.ec, err = ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	bgCtx := context.Background()
	chid, _ := ret.ec.ChainID(bgCtx)
	ret.chid = chid.Uint64()
	ret.auth, err = nftbr.Kspath2auth(kspath, pass, chid)
	if err != nil {
		return nil, err
	}
	ret.nftbr, _ = binding.NewNFTBridge(nftbrA, ret.ec)
	return ret, nil
}

// for this chain, query onchain for addr in origNfts and set if mismatch and sendTx is true
// can't delete entry!
func (o *OpOneCh) CheckOrigNFT(origNfts []string, sendTx bool) {
	for _, onft := range origNfts {
		nAddr := nftbr.Hex2addr(onft)
		hasSet, err := o.nftbr.OrigNFT(nil, nAddr)
		if err != nil {
			log.Warn("OrigNFT err: ", err)
			continue
		}
		if !hasSet {
			log.Infoln(o.chid, "orig", onft, "not set")
			if sendTx {
				tx, err := o.nftbr.SetOrigNFT(o.auth, nAddr)
				if err != nil {
					log.Error("SetOrigNFT err: ", err)
				} else {
					log.Info("SetOrigNFT tx: ", tx.Hash().Hex())
				}
			}
		}
	}
}

func (o *OpOneCh) CheckDstBridgeMap(allbrs map[uint64]common.Address, sendTx bool) {
	var tofixChids []uint64
	var tofixAddr []common.Address
	for chid, addr := range allbrs {
		if chid == o.chid {
			continue // no need for self
		}
		dstbr, err := o.nftbr.DestBridge(nil, chid)
		if err != nil {
			log.Warn("DestBridge err: ", err)
			continue
		}
		// eg. 0 address
		if dstbr != addr {
			log.Infoln(o.chid, "dstchid", chid, "nftbr got", dstbr, "expect:", addr)
			tofixChids = append(tofixChids, chid)
			tofixAddr = append(tofixAddr, addr)
		}
	}
	if sendTx && len(tofixChids) > 0 {
		log.Info("SetDestBridges chids: ", tofixChids)
		var tx *types.Transaction
		var err error
		if len(tofixChids) == 1 {
			tx, err = o.nftbr.SetDestBridge(o.auth, tofixChids[0], tofixAddr[0])
		} else {
			tx, err = o.nftbr.SetDestBridges(o.auth, tofixChids, tofixAddr)
		}
		if err != nil {
			log.Error("SetDestBridges err: ", err)
		} else {
			log.Info("SetDestBridges tx: ", tx.Hash().Hex())
		}
	}
}

func (o *OpOneCh) CheckNftMap(nftmap map[string]map[uint64]string, sendTx bool) {
	for sym, chaddr := range nftmap {
		// for each nft, find its address on this chain, then get all other chid and addr
		addrOnThisChain := chaddr[o.chid]
		if addrOnThisChain == "" {
			// nft not found for this chain, skip
			continue
		}
		var tofixChids []uint64
		var tofixAddr []common.Address
		// now check each chid,addr with contract
		for chid, addr := range chaddr {
			if chid == o.chid {
				continue // no need to check same chain
			}
			got, _ := o.nftbr.DestNFTAddr(nil, nftbr.Hex2addr(addrOnThisChain), chid)
			if got != nftbr.Hex2addr(addr) {
				log.Infoln(o.chid, sym, "nft mismatch got:", got, "exp:", addr)
				tofixChids = append(tofixChids, chid)
				tofixAddr = append(tofixAddr, nftbr.Hex2addr(addr))
			}
		}
		if sendTx && len(tofixChids) > 0 {
			var tx *types.Transaction
			var err error
			if len(tofixChids) == 1 {
				tx, err = o.nftbr.SetDestNFT(o.auth, nftbr.Hex2addr(addrOnThisChain), tofixChids[0], tofixAddr[0])
			} else {
				tx, err = o.nftbr.SetDestNFTs(o.auth, nftbr.Hex2addr(addrOnThisChain), tofixChids, tofixAddr)
			}
			if err != nil {
				log.Error("SetDestNFTs err: ", err)
			} else {
				log.Info("SetDestNFTs tx: ", tx.Hash().Hex())
			}
		}
	}
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
