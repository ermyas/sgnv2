/*
Copyright © 2021 Celer Network

*/
package cmd

import (
	"log"
	"math/big"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/spf13/cobra"
)

var (
	signerks, signerpw string // path to signer ks json
)

// updateSigners can be called by anyone w/ correct msg
var updateSignersCmd = &cobra.Command{
	Use:   "updateSigners",
	Short: "call updateSigners will sign using signer ks only support one signer for now. signer1Eth,power for both curss and newss",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ethamt := strings.Split(args[0], ",")
		amt, _ := new(big.Int).SetString(ethamt[1], 10)
		onesigner := &cbrtypes.AddrAmt{
			Addr: eth.Hex2Bytes(ethamt[0]),
			Amt:  amt.Bytes(),
		}
		ss := cbrtypes.SortedSigners{
			Signers: []*cbrtypes.AddrAmt{onesigner},
		}
		raw, _ := ss.Marshal()
		signer := kspath2signer(signerks, signerpw)
		sig := signer.SignData(raw)
		log.Printf("raw: %x\nsig: %x", raw, sig)
		// now try to submit onchain
		cbrContract, _ = eth.NewBridge(eth.Hex2Addr(cfg.Cbridge), ec)
		tx, err := cbrContract.UpdateSigners(auth, raw, raw, [][]byte{sig})
		chkTxErr(tx, err)
	},
}

func init() {
	rootCmd.AddCommand(updateSignersCmd)
	updateSignersCmd.Flags().StringVarP(&signerks, "signer", "s", "", "path to signer ks json")
	updateSignersCmd.Flags().StringVarP(&signerpw, "signerpw", "pw", "", "password to signer ks json")
}
