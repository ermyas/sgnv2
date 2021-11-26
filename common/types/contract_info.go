package types

import "fmt"

func NewContractInfo(chainId uint64, address string) ContractInfo {
	return ContractInfo{
		ChainId: chainId,
		Address: address,
	}
}

func (info *ContractInfo) FormatStr() string {
	return fmt.Sprintf("%d-%s", info.ChainId, info.Address)
}
