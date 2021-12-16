/*
Copyright © 2021 Celer Network

*/
package cmd

import (
	"context"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	ethutil "github.com/celer-network/goutils/eth"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/cobra"
)

var (
	waitOpts = []ethutil.TxOption{ethutil.WithPollingInterval(time.Second * 5), ethutil.WithBlockDelay(1)}
	// set max_uint = 2**256-1
	// "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	MaxUint256 = new(big.Int).SetBytes(eth.Hex2Bytes(strings.Repeat("ff", 32)))
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send token, eg. send usdt 3 10000000 50000, 3 is dst chainid, 50000 is max slip, means (50000/10000)% = 5%",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		if auth == nil {
			log.Fatal("must set -ks flag")
		}
		sym := args[0]
		tokenAddr := cfg.GetTokenAddr(sym)
		dstChid, _ := strconv.Atoi(args[1])
		amt, _ := new(big.Int).SetString(args[2], 10)
		maxSlip, _ := new(big.Int).SetString(args[3], 10)
		log.Println("send", sym, "from", chainName, "addr:", tokenAddr, "to", dstChid, "amt:", amt, "max slip:", maxSlip)
		cbrAddr := eth.Hex2Addr(cfg.Cbridge)
		erc20, _ := eth.NewErc20(tokenAddr, ec)
		bal, _ := erc20.BalanceOf(nil, auth.From)
		if bal.Cmp(amt) == -1 {
			log.Fatalln("balance", bal, "<", amt)
		}
		allowed, _ := erc20.Allowance(nil, auth.From, cbrAddr)
		if allowed.Cmp(amt) == -1 {
			log.Println(allowed, "<", amt, "need to approve first")
			tx, err := erc20.Approve(auth, cbrAddr, MaxUint256)
			if err != nil {
				log.Fatalln("approve err:", err)
			}
			// wait mine so send won't err
			ethutil.WaitMined(context.Background(), ec, tx, waitOpts...)
		}
		cbr, _ := eth.NewBridge(cbrAddr, ec)
		log.Println("calling onchain send")
		tx, err := cbr.Send(auth, auth.From, tokenAddr, amt, uint64(dstChid), uint64(time.Now().Unix()), uint32(maxSlip.Uint64()))
		chkTxErr(tx, err)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	// sendCmd.PersistentFlags().Uint32() max slippage flag
}
