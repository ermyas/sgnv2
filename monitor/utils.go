package monitor

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

func getEventCheckInterval(name eth.EventName) uint64 {
	m := viper.GetStringMap(common.FlagEthCheckInterval)
	eventNameInConfig := strcase.ToSnake(string(name))
	if m[eventNameInConfig] != nil {
		return uint64(m[eventNameInConfig].(int64))
	}
	// If not specified, use the default value of 0
	return 0
}
