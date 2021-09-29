package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	gw = flag.String("gw", "https://goerli.infura.io/v3/c58bb2a63c1a466abb6d7b5e0b7b6828", "json rpc server")
	ks = flag.String("ks", "", "keystore json")
	// after deploy, setup initSigner
	cbr = flag.String("cbr", "", "cbridge hex address")
	// set weth addr
	weth = flag.String("weth", "", "weth hex address")
)

const (
	// from genesis.json staking validators
	valEth = "0xa500023551388763b720808c0b0cdf00a752b69f"
	valAmt = "1000000000000"
)

func main() {
	flag.Parse()
	ec, err := ethclient.Dial(*gw)
	chkErr(err, "dial")
	chainid, err := ec.ChainID(context.Background())
	chkErr(err, "")
	log.Print("chainid: ", chainid)
	auth, err := kspath2auth(*ks, chainid)
	if *cbr != "" {
		cbridge, _ := eth.NewBridge(eth.Hex2Addr(*cbr), ec)
		var tx *types.Transaction
		if *weth != "" {
			// set weth
			wethAddr := eth.Hex2Addr(*weth)
			tx, err = cbridge.SetWrap(auth, wethAddr)
		} else {
			// init signer
			amt, _ := new(big.Int).SetString(valAmt, 10)
			signers := new(cbrtypes.SortedSigners)
			signers.Signers = append(signers.Signers, &cbrtypes.AddrAmt{
				Addr: eth.Hex2Bytes(valEth),
				Amt:  amt.Bytes(),
			})
			raw, _ := signers.Marshal()
			tx, err = cbridge.SetInitSigners(auth, raw)
		}
		log.Print(tx.Hash().String(), err)
		return
	}
	addr, _, _, _ := eth.DeployBridge(auth, ec, nil)
	log.Print("cbridge addr: ", addr.String())
}

func kspath2auth(kspath string, chainid *big.Int) (*bind.TransactOpts, error) {
	ksjson, err := ioutil.ReadFile(kspath)
	if err != nil {
		return nil, err
	}
	kss := string(ksjson)
	return bind.NewTransactorWithChainID(strings.NewReader(kss), "", chainid)
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
