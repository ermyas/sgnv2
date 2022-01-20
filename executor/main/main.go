package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/executor"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
)

var (
	home = flag.String("home", os.ExpandEnv("$HOME/.executor"), "home path")
)

func main() {
	flag.Parse()
	setupConfig()
	dal := executor.NewDAL()
	sgn := executor.NewSgnClient()
	chains := executor.NewChainMgr(dal)
	ex := executor.NewExecutor(dal, sgn, chains)
	ex.Start()
}

func setupConfig() {
	// sets account address prefix for transactors
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	sdkConfig.Seal()

	viper.SetDefault(flags.FlagHome, *home)
	readConfig("config/executor.toml")
	readConfig("config/cbridge.toml")
}

func readConfig(relativePath string) {
	path := filepath.Join(*home, relativePath)
	viper.SetConfigFile(path)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalln("failed to load", path, err)
	}
}
