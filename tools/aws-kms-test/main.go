package main

import (
	"flag"
	"log"
	"math/big"

	"github.com/celer-network/goutils/eth"
)

var (
	region = flag.String("r", "us-west-2", "AWS region that has the KMS key")
	alias  = flag.String("k", "", "alias for the key, no need to add alias/ prefix")
)

// test and verify aws kms apis
func main() {
	flag.Parse()
	k, err := eth.NewKmsSigner(*region, "alias/"+*alias, "", "", big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ETH addr:", k.Addr)
	_, err = k.SignEthMessage([]byte("123456"))
	if err != nil {
		log.Fatal("sign err: ", err)
	}
}
