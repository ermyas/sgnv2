package types

import (
	comtypes "github.com/celer-network/sgn-v2/common/types"
)

type ContractConfig struct {
	ChainId uint64 `mapstructure:"chain_id"`
	Address string `mapstructure:"address"`
	// the payable value to add when calling executeMessage or executeMessageWithTransfer
	PayableValue string `mapstructure:"add_payable_value_for_execution"`
}

func (c *ContractConfig) toContractInfo() *comtypes.ContractInfo {
	return &comtypes.ContractInfo{
		ChainId: c.ChainId,
		Address: c.Address,
	}
}

func MapToContractInfos(configs []*ContractConfig) []*comtypes.ContractInfo {
	infos := []*comtypes.ContractInfo{}
	for _, config := range configs {
		infos = append(infos, config.toContractInfo())
	}
	return infos
}

func GetContractConfig(configs []*ContractConfig, addr string) (*ContractConfig, bool) {
	for _, config := range configs {
		if config.Address == addr {
			return config, true
		}
	}
	return nil, false
}
