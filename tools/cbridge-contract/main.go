package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func main() {
	flag.Parse()
	ec, err := ethclient.Dial(*gw)
	chkErr(err, "dial")
	chainid, err := ec.ChainID(context.Background())
	chkErr(err, "")
	log.Print("chainid: ", chainid)
	auth, err := kspath2auth(*ks, chainid)
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
