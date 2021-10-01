/*
Copyright © 2021 Celer Network

*/
package cmd

import (
	"context"
	"io/ioutil"
	"log"
	"math/big"
	"reflect"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viper default delimiter for keys
const DELIM = "."

var (
	cfgFile, chainName, ksPath string

	// set in initConfig
	auth *bind.TransactOpts
	ec   *ethclient.Client
	cfg  *OneChainConfig // we could also use viper get x.y directly
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cbridge",
	Short: "onchain ops about cbridge contract",
}

type OneChainConfig struct {
	Chainid                      uint64
	Gateway, Cbridge, USDT, WETH string
}

// given string, return hex2addr, if not found, zero addr
func (c *OneChainConfig) GetTokenAddr(sym string) eth.Addr {
	rv := reflect.ValueOf(c).Elem().FieldByName(sym)
	if rv.IsValid() {
		return eth.Hex2Addr(rv.String())
	}
	// not found, return zero addr
	return eth.ZeroAddr
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "cfg", "cbridge.toml", "config file")
	rootCmd.PersistentFlags().StringVar(&chainName, "chain", "ropsten", "ropsten, goerli etc, must be defined i config")
	rootCmd.PersistentFlags().StringVar(&ksPath, "ks", "", "path to ks json")
}

// initConfig reads in config file and set var
func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.ReadInConfig()

	cfg = new(OneChainConfig)
	viper.UnmarshalKey(chainName, cfg)
	log.Printf("chain: %s, cfg: %+v", chainName, cfg)
	ec, _ = ethclient.Dial(cfg.Gateway)
	chid, _ := ec.ChainID(context.Background())
	if chid.Uint64() != cfg.Chainid {
		log.Fatalf("chainid mismatch! cfg has %d, rpc has %d", cfg.Chainid, chid.Uint64())
	}
	if ksPath != "" {
		auth, _ = kspath2auth(ksPath, big.NewInt(int64(cfg.Chainid)))
	}
}

func kspath2auth(kspath string, chainid *big.Int) (*bind.TransactOpts, error) {
	ksjson, err := ioutil.ReadFile(kspath)
	if err != nil {
		return nil, err
	}
	kss := string(ksjson)
	return bind.NewTransactorWithChainID(strings.NewReader(kss), "", chainid)
}
