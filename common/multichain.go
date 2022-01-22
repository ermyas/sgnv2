package common

// represent one chain in config file, include necessary info like chainid, gateway, cbridge address etc
type OneChainConfig struct {
	ChainID                                             uint64
	Name, Gateway                                       string
	BlkInterval, BlkDelay, MaxBlkDelta, ForwardBlkDelay uint64
	AddGasEstimateRatio                                 float64
	// Legacy gas price flag
	AddGasGwei uint64
	// EIP-1559 gas price flag
	MaxFeePerGasGwei uint64
	// cbridge contract address
	CBridge string
	// OriginalTokenVault contract address
	OTVault string
	// PeggedTokenBridge contract address
	PTBridge string
	// MsgBus contract address
	MsgBus string
	// if ProxyPort > 0, a proxy with this port will be created to support some special chain such as harmony, celo.
	// chainID will be used to determined which type proxy to create, so make sure the chainID is supported in the "endpoint-proxy"
	// create a proxy to the Gateway, and eth-client will be created to "127.0.0.1:ProxyPort"
	// more detail, https://github.com/celer-network/endpoint-proxy
	ProxyPort     int
	CheckInterval map[string]interface{}
}

type MultiChainConfig []*OneChainConfig

// return config if chainid is found, otherwise return nil
func (m MultiChainConfig) GetConfig(chainid uint64) *OneChainConfig {
	for _, c := range m {
		if c.ChainID == chainid {
			return c
		}
	}
	return nil
}
