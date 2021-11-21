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
