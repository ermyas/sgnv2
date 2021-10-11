/*
Copyright © 2021 Celer Network

*/
package cmd

import (
	"log"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy cbridge contract",
	Run: func(cmd *cobra.Command, args []string) {
		addr, tx, _, err := eth.DeployBridge(auth, ec)
		if err != nil {
			log.Fatal("deploy err:", err)
		}
		log.Println("cbridge addr:", addr, "tx:", tx.Hash())
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
