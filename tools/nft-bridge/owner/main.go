package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/celer-network/goutils/log"
	nftbr "github.com/celer-network/sgn-v2/tools/nft-bridge"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	"github.com/spf13/viper"
)

var (
	fchain = flag.String("c", "nftbr_test.toml", "multi chain config toml file")
	fnfts  = flag.String("n", "nfts.toml", "which nfts to monitor for owner change")
)

func main() {
	flag.Parse()
	viper.SetConfigFile(*fchain)
	err := viper.ReadInConfig()
	chkErr(err, "viper chain config "+*fchain)
	mcc := nftbr.GetMcc("multichain")
	dal, err := dal.NewDAL(viper.GetString("db"))
	chkErr(err, "new dal")
	chainMap := make(map[uint64]*OneChain)
	for _, cfg := range mcc {
		oc, err := NewOneChain(*cfg, dal)
		if err != nil {
			log.Errorln("newOneChain err:", err, "cfg:", *cfg)
			continue
		}
		chainMap[cfg.ChainID] = oc
		defer oc.Close()
	}
	// now go over nfts.toml and MonNft
	viper2 := viper.New()
	viper2.SetConfigFile(*fnfts)
	err = viper2.ReadInConfig()
	chkErr(err, "viper nft config "+*fnfts)
	var nftList NftList
	viper2.UnmarshalKey("nft", &nftList)
	for _, n := range nftList {
		chainMap[n.Chainid].MonNft(n)
	}
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM)
	<-sigch
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
