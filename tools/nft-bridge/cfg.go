package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/spf13/viper"
)

func GetMcc(key string) MultiChainConfig {
	var m MultiChainConfig
	viper.UnmarshalKey(key, &m)
	return m
}

type MultiChainConfig []*OneChainConfig

type OneChainConfig struct {
	ChainID       uint64
	Name, Gateway string
	// blk related for monitor
	BlkInterval, BlkDelay        uint64
	MaxBlkDelta, ForwardBlkDelay uint64
	// gas related for send tx
	GasLimit            uint64
	AddGasEstimateRatio float64
	// Legacy gas price flags
	AddGasGwei uint64
	MinGasGwei uint64
	MaxGasGwei uint64
	// EIP-1559 gas price flags
	MaxFeePerGasGwei         uint64
	MaxPriorityFeePerGasGwei uint64

	MsgBus string // hex addr of msg bus contract on this chain
}

// json from url

type ChidAddr struct {
	Chainid uint64
	Addr    string
}

type OneNft struct {
	Name, Symbol, Url string
	// addr per chain
	Orig *ChidAddr
	Pegs []*ChidAddr
}

type JsonCfg struct {
	Bridges []*ChidAddr
	Nfts    []*OneNft
}

// fetch cfg key's value via http, parse into JsonCfg
func GetJsonCfg(key string) *JsonCfg {
	resp, err := http.Get(viper.GetString(key))
	chkErr(err, "get json")
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	chkErr(err, "read resp.Body")
	jsonCfg := new(JsonCfg)
	err = json.Unmarshal(raw, jsonCfg)
	chkErr(err, "unmarshal")
	return jsonCfg
}
