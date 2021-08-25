package eth

import "github.com/ethereum/go-ethereum/ethclient"

type StakingContract struct {
	*Staking
	Address Addr
}

func NewStakingContract(address Addr, client *ethclient.Client) (*StakingContract, error) {
	staking, err := NewStaking(address, client)
	if err != nil {
		return nil, err
	}
	return &StakingContract{
		Staking: staking,
		Address: address,
	}, nil
}

func (c *StakingContract) GetAddr() Addr {
	return c.Address
}

func (c *StakingContract) GetABI() string {
	return StakingABI
}

type SgnContract struct {
	*SGN
	Address Addr
}

func NewSgnContract(address Addr, client *ethclient.Client) (*SgnContract, error) {
	sgn, err := NewSGN(address, client)
	if err != nil {
		return nil, err
	}
	return &SgnContract{
		SGN:     sgn,
		Address: address,
	}, nil
}

func (c *SgnContract) GetAddr() Addr {
	return c.Address
}

func (c *SgnContract) GetABI() string {
	return SGNABI
}
