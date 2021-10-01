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
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ownerCmd)
	// owner sub cmds
	ownerCmd.AddCommand(
		initSignerCmd,
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

var initSignerCmd = &cobra.Command{
	Use:   "initSigner",
	Short: "call setInitSigners, args are sorted by eth signer1Eth,power signer2Eth,power",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		signers := new(cbrtypes.SortedSigners)
		for _, arg := range args {
			ethamt := strings.Split(arg, ",")
			amt, _ := new(big.Int).SetString(ethamt[1], 10)
			signers.Signers = append(signers.Signers, &cbrtypes.AddrAmt{
				Addr: eth.Hex2Bytes(ethamt[0]),
				Amt:  amt.Bytes(),
			})
		}
		raw, _ := signers.Marshal()
		tx, err := cbrContract.SetInitSigners(auth, raw)
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
	Short: "call setMinSlippage x, x is slippage *1e6",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		minSlip, _ := strconv.Atoi(args[0])
		tx, err := cbrContract.SetMinSlippage(auth, uint32(minSlip))
		chkTxErr(tx, err)
	},
}

func chkTxErr(tx *types.Transaction, err error) {
	if err != nil {
		log.Fatalln("tx err:", err)
	}
	log.Println("tx:", tx.Hash().String())
}
