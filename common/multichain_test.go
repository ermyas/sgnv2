package common

import (
	"bytes"
	"testing"

	"github.com/spf13/viper"
)

const mcc = `
[[multichain]]
chainID = 123
name = "test1"
gateway = "wss://chain123.net/ws"
blkinterval = 15
blkdelay = 5
maxblkdelta = 1000
cbridge = "0x123123123"

[[multichain]]
chainID = 456
name = "test2"
gateway = "wss://chain456.net/ws"
blkinterval = 15
blkdelay = 5
maxblkdelta = 5000
cbridge = "0x456456"
`

func TestMultiChainCfg(t *testing.T) {
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer([]byte(mcc)))
	var m MultiChainConfig
	viper.UnmarshalKey(FlagMultiChain, &m)
	t.Error(m.GetConfig(123), m.GetConfig(456))
}
