/*
Copyright © 2021 Celer Network

*/
package cmd

import (
	"context"
	"log"
	"math/big"
	"strings"

	ethutil "github.com/celer-network/goutils/eth"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/cobra"
)

// addLiqCmd represents the addLiq command
var addLiqCmd = &cobra.Command{
	Use:   "addLiq",
	Short: "addLiq symbol amount",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		sym := strings.ToUpper(args[0])
		tokenAddr := cfg.GetTokenAddr(sym)
		amt, _ := new(big.Int).SetString(args[1], 10)
		log.Println("addliq for", chainName, sym, "addr:", tokenAddr, "amt:", amt)
		erc20, _ := eth.NewErc20(tokenAddr, ec)
		bal, _ := erc20.BalanceOf(nil, auth.From)
		cbrAddr := eth.Hex2Addr(cfg.Cbridge)
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
		cbrContract, _ = eth.NewBridge(cbrAddr, ec)
		log.Println("calling onchain addLiquidity")
		tx, err := cbrContract.AddLiquidity(auth, tokenAddr, amt)
		chkTxErr(tx, err)
	},
}

func init() {
	rootCmd.AddCommand(addLiqCmd)
}
