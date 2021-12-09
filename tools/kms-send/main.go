package main

import (
	"context"
	"flag"
	"log"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// send aws kms signer's eth balance to dest addr
var (
	keyA = flag.String("a", "sgnv2-test-0", "kms alias like sgnv2-prod-0")
	dst  = flag.String("d", "", "dest addr hex")
	chid = flag.Int64("i", 0, "which chainid to send eth, if 0, do all known chains in chmap")
	minb = flag.Uint64("min", 1e17, "minimal balance in wei, if less than this, skip send")
	gas  = flag.Uint64("gas", 21000, "default gas limit, in case some chain requires more")

	bgCtx = context.Background()
	// chainid to rpc endpoint
	chmap = map[int64]string{
		1:     "https://eth-mainnet.alchemyapi.io/v2/imF7XWfmReAqqvyVafPOSJxDtY8U8_ZT",
		56:    "https://sparkling-wild-lake.bsc.quiknode.pro/2c5285fa561849e615c58fc36c1cb14a20f89a2b/",
		42161: "https://arbitrum-mainnet.infura.io/v3/029bfb981dc748b1affa9e77eb0e5477",
		137:   "https://bold-winter-surf.matic.quiknode.pro/9c226c540f9c66f92f4ef21af6e3dfa257c91d68/",
		43114: "https://api.avax.network/ext/bc/C/rpc",
		250:   "https://holy-floral-morning.fantom.quiknode.pro/867064df062739ca6f3805d7faafff8d4c605d1f/",
		10:    "https://optimism-mainnet.infura.io/v3/029bfb981dc748b1affa9e77eb0e5477",
	}
)

func main() {
	flag.Parse()
	if *chid > 0 {
		doOne(*chid, chmap[*chid])
		return
	}
	// note go map iter is random order
	for id, rpc := range chmap {
		doOne(id, rpc)
	}
}

func doOne(chid int64, rpc string) {
	ec, err := ethclient.Dial(rpc)
	chkErr(err, "dial "+rpc)
	kms, err := ethutils.NewKmsSigner("us-west-2", "alias/"+*keyA, "", "", big.NewInt(chid))
	chkErr(err, "newsigner")
	bal, err := ec.BalanceAt(bgCtx, kms.Addr, nil)
	chkErr(err, "get balance")
	log.Println("chid:", chid, "bal:", bal)
	minBal := new(big.Int).SetUint64(*minb)
	if bal.Cmp(minBal) <= 0 {
		log.Println("skip due to balance less than min")
		return
	}
	// now build tx and send
	sendETH(ec, kms.Addr, eth.Hex2Addr(*dst), bal, kms.SignerFn)
}

// send bal - gas, for eip1559 it's possible we still has some left b/c we only set cap
func sendETH(ec *ethclient.Client, from, to eth.Addr, bal *big.Int, signer bind.SignerFn) error {
	var rawTx *types.Transaction
	head, err := ec.HeaderByNumber(bgCtx, nil)
	chkErr(err, "HeaderByNumber")
	nonce, err := ec.PendingNonceAt(bgCtx, from)
	chkErr(err, "PendingNonceAt")
	gasPrice, err := ec.SuggestGasPrice(bgCtx)
	chkErr(err, "SuggestGasPrice")
	if head.BaseFee != nil {
		// eip 1559, new dynamic tx, per spec we should do
		// maxPriorityFeePerGas: eth_gasPrice - base_fee or just use the eth_maxPriorityFeePerGas rpc
		// maxFeePerGas: maxPriorityFeePerGas + 2 * base_fee = eth_gasPrice + base_fee
		// note if we calculate sendamt based on maxFeePerGas, it will leave one base_fee*gas residual
		// assume maxPriorityFee is way smaller than base fee, we could do following:
		// GasTipCap := eth_maxPriorityFeePerGas and GasFeeCap := eth_gasPrice + GasTipCap
		// but the risk is if eth becomes busy, our tx may pending for a long time. as here our gas is only 21K, we are ok w/ base_fee*gas residual
		gasFeeCap := new(big.Int).Add(gasPrice, head.BaseFee)
		gasCost := new(big.Int).Mul(gasFeeCap, big.NewInt(int64(*gas)))
		rawTx = types.NewTx(&types.DynamicFeeTx{
			Nonce:     nonce,
			To:        &to,
			Gas:       *gas,
			GasTipCap: new(big.Int).Sub(gasPrice, head.BaseFee),
			GasFeeCap: gasFeeCap,
			Value:     new(big.Int).Sub(bal, gasCost),
		})

	} else {
		rawTx = types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			To:       &to,
			Gas:      *gas,
			GasPrice: gasPrice,
			Value:    new(big.Int).Sub(bal, new(big.Int).Mul(gasPrice, big.NewInt(int64(*gas)))),
		})
	}
	signedTx, err := signer(from, rawTx)
	chkErr(err, "SignerFn")
	log.Println("tx:", signedTx.Hash())
	err = ec.SendTransaction(bgCtx, signedTx)
	chkErr(err, "send onchain")
	return nil
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
