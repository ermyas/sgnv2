/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"flag"
	"path/filepath"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/executor"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetStartCmd get the start command
func GetStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start executor service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ll, err := cmd.Flags().GetString("loglevel")
			if err != nil {
				return err
			}
			log.SetLevelByName(ll)
			test, err := cmd.InheritedFlags().GetBool("test")
			if err != nil {
				return err
			}
			home, err := cmd.InheritedFlags().GetString("home")
			if err != nil {
				return err
			}

			start(test, home)
			return nil
		},
	}
	cmd.Flags().String("loglevel", "info", "Log level")
	return cmd
}

func init() {
	rootCmd.AddCommand(GetStartCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func start(test bool, home string) {
	flag.Parse()
	if test {
		log.Infoln("Starting executor with test mode")
	} else {
		log.Infoln("Starting executor")
	}
	setupConfig(home)
	dal := executor.NewDAL()
	ex := executor.NewExecutor(dal, test)
	ex.Start()
}

func setupConfig(home string) {
	log.Infoln("Reading executor configs")
	// sets account address prefix for transactors
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	sdkConfig.Seal()

	viper.SetDefault(flags.FlagHome, home)
	readConfig(home, "config/executor.toml")
	readConfig(home, "config/cbridge.toml")
}

func readConfig(home, relativePath string) {
	path := filepath.Join(home, relativePath)
	viper.SetConfigFile(path)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalln("failed to load", path, err)
	}
}
