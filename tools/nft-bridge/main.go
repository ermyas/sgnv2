package main

import (
	"flag"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	"github.com/spf13/viper"
)

var (
	fcfg = flag.String("c", "nftbr_test.toml", "config toml file")
)

func main() {
	flag.Parse()
	viper.SetConfigFile(*fcfg)
	err := viper.ReadInConfig()
	chkErr(err, "viper ReadInConfig")

	dal, err := dal.NewDAL(viper.GetString("db"))
	chkErr(err, "new dal")

	mcc := GetMcc("multichain")
	jsonCfg := GetJsonCfg("jsonurl")

	// chainid to *OneChain
	chainMap := make(map[uint64]*OneChain)
	kspath := viper.GetString("kspath")
	ksphrase := viper.GetString("ksphrase")
	for _, cfg := range mcc {
		oc, err := NewOneChain(*cfg, kspath, ksphrase, dal)
		if err != nil {
			log.Errorln("newOneChain err:", err, "cfg:", *cfg)
			continue
		}
		chainMap[cfg.ChainID] = oc
		defer oc.Close()
	}

	// now per chain mon nft bridge
	for _, nftbr := range jsonCfg.Bridges {
		chainMap[nftbr.Chainid].MonNftBridge(nftbr.Addr)
	}
	// poll sgn for available msg to send
	go PollSgn(time.Minute, jsonCfg.Bridges, chainMap)
	// block serving http
	NewServer(dal).Run(viper.GetInt("port"))
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

/*
	nftAddr := common.HexToAddress("0x7A46219950d8a9bf2186549552DA35Bf6fb85b1F")
	receiver := common.HexToAddress("0x51D36E18E3D32d121A3CfE2F3E5771A6FD53274E")
	ec, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545")
	chkErr(err, "dial")
	chid, _ := ec.ChainID(context.Background())
	chkErr(err, "chainid")
	auth, err := kspath2auth("/Users/junda/.ssh/ks-9f6b.json", chid)
	chkErr(err, "auth")

	oNft, err := NewOrigNFT(nftAddr, ec)
	chkErr(err, "NewOrigNFT")
	for id := 1; id <= 10; id++ {
		tx, err := oNft.Mint(auth, receiver, big.NewInt(int64(id)), fmt.Sprintf("https://celerx-avatars.s3.us-west-2.amazonaws.com/default/%d.png", id+10))
		chkErr(err, fmt.Sprintf("Mint %d", id))
		log.Println(tx.Hash())
	}
*/
