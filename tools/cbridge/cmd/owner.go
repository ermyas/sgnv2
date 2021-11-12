/*
Copyright © 2021 Celer Network

*/
package cmd

import (
	"log"
	"math/big"
	"strconv"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ownerCmd)
	// owner sub cmds
	ownerCmd.AddCommand(
		resetSignerCmd,
		notifyResetSignerCmd,
		wrapCmd,
		minSendCmd,
		minSlipCmd,
	)
}

var cbrContract *eth.Bridge // set in PersistentPreRun

// ownerCmd represents the owner command
var ownerCmd = &cobra.Command{
	Use:   "owner",
	Short: "call owner only onchain funcs",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cbrContract, _ = eth.NewBridge(eth.Hex2Addr(cfg.Cbridge), ec)
	},
}

var resetSignerCmd = &cobra.Command{
	Use:   "resetSigner",
	Short: "call resetSigners, args are sorted by eth signer1Eth,power signer2Eth,power",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var addrs []eth.Addr
		var powers []*big.Int
		for _, arg := range args {
			ethamt := strings.Split(arg, ",")
			amt, _ := new(big.Int).SetString(ethamt[1], 10)
			addrs = append(addrs, eth.Hex2Addr(ethamt[0]))
			powers = append(powers, amt)
		}
		tx, err := cbrContract.ResetSigners(auth, addrs, powers)
		chkTxErr(tx, err)
	},
}

var notifyResetSignerCmd = &cobra.Command{
	Use:   "notifyResetSigner",
	Short: "call NotifyResetSigners",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tx, err := cbrContract.NotifyResetSigners(auth)
		chkTxErr(tx, err)
	},
}

var wrapCmd = &cobra.Command{
	Use:   "warp",
	Short: "call setWrap, arg is weth address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wethAddr := eth.Hex2Addr(args[0])
		tx, err := cbrContract.SetWrap(auth, wethAddr)
		chkTxErr(tx, err)
	},
}

var minSendCmd = &cobra.Command{
	Use:   "minSend",
	Short: "call setMinSend, tokenaddr,minsend token2addr,minsend",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var addrList []eth.Addr
		var minSendList []*big.Int
		for _, arg := range args {
			addramt := strings.Split(arg, ",")
			addrList = append(addrList, eth.Hex2Addr(addramt[0]))
			amt, _ := new(big.Int).SetString(addramt[1], 10)
			minSendList = append(minSendList, amt)
		}
		tx, err := cbrContract.SetMinSend(auth, addrList, minSendList)
		chkTxErr(tx, err)
	},
}

var minSlipCmd = &cobra.Command{
	Use:   "minSlip",
	Short: "call setMinimalMaxSlippage x, x is slippage *1e6",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		minSlip, _ := strconv.Atoi(args[0])
		tx, err := cbrContract.SetMinimalMaxSlippage(auth, uint32(minSlip))
		chkTxErr(tx, err)
	},
}

func chkTxErr(tx *types.Transaction, err error) {
	if err != nil {
		log.Fatalln("tx err:", err)
	}
	log.Println("tx:", tx.Hash().String())
}
