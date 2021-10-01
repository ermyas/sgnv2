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

var waitOpts = []ethutil.TxOption{ethutil.WithPollingInterval(time.Second * 5), ethutil.WithBlockDelay(1)}

const maxSlip = 50000 // todo: use flag

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send token, eg. send usdt 3 10000000, 3 is dst chainid",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if auth == nil {
			log.Fatal("must set -ks flag")
		}
		sym := strings.ToUpper(args[0])
		tokenAddr := cfg.GetTokenAddr(sym)
		dstChid, _ := strconv.Atoi(args[1])
		amt, _ := new(big.Int).SetString(args[2], 10)
		log.Println("send", sym, "from", chainName, "addr:", tokenAddr, "to", dstChid, "amt:", amt)
		cbrAddr := eth.Hex2Addr(cfg.Cbridge)
		erc20, _ := eth.NewErc20(tokenAddr, ec)
		bal, _ := erc20.BalanceOf(nil, auth.From)
		if bal.Cmp(amt) == -1 {
			log.Fatalln("balance", bal, "<", amt)
		}
		allowed, _ := erc20.Allowance(nil, auth.From, cbrAddr)
		if allowed.Cmp(amt) == -1 {
			log.Println(allowed, "<", amt, "need to approve first")
			tx, err := erc20.Approve(auth, cbrAddr, amt)
			if err != nil {
				log.Fatalln("approve err:", err)
			}
			// wait mine so send won't err
			ethutil.WaitMined(context.Background(), ec, tx, waitOpts...)
		}
		cbr, _ := eth.NewBridge(cbrAddr, ec)
		log.Println("calling onchain send")
		tx, err := cbr.Send(auth, auth.From, tokenAddr, amt, uint64(dstChid), uint64(time.Now().Unix()), maxSlip)
		chkTxErr(tx, err)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	// sendCmd.PersistentFlags().Uint32() max slippage flag
}
