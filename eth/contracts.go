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

type RewardContract struct {
	*Reward
	Address Addr
}

func NewRewardContract(address Addr, client *ethclient.Client) (*RewardContract, error) {
	Reward, err := NewReward(address, client)
	if err != nil {
		return nil, err
	}
	return &RewardContract{
		Reward:  Reward,
		Address: address,
	}, nil
}

func (c *RewardContract) GetAddr() Addr {
	return c.Address
}

func (c *RewardContract) GetABI() string {
	return RewardABI
}

type GovernContract struct {
	*Govern
	Address Addr
}

func NewGovernContract(address Addr, client *ethclient.Client) (*GovernContract, error) {
	govern, err := NewGovern(address, client)
	if err != nil {
		return nil, err
	}
	return &GovernContract{
		Govern:  govern,
		Address: address,
	}, nil
}

func (c *GovernContract) GetAddr() Addr {
	return c.Address
}

func (c *GovernContract) GetABI() string {
	return GovernABI
}

type ViewerContract struct {
	*Viewer
	Address Addr
}

func NewViewerContract(address Addr, client *ethclient.Client) (*ViewerContract, error) {
	Viewer, err := NewViewer(address, client)
	if err != nil {
		return nil, err
	}
	return &ViewerContract{
		Viewer:  Viewer,
		Address: address,
	}, nil
}

func (c *ViewerContract) GetAddr() Addr {
	return c.Address
}

func (c *ViewerContract) GetABI() string {
	return ViewerABI
}

type BridgeContract struct {
	*Bridge
	Address Addr
}

func NewBridgeContract(address Addr, client *ethclient.Client) (*BridgeContract, error) {
	bridge, err := NewBridge(address, client)
	if err != nil {
		return nil, err
	}
	return &BridgeContract{
		Bridge:  bridge,
		Address: address,
	}, nil
}

func (c *BridgeContract) GetAddr() Addr {
	return c.Address
}

func (c *BridgeContract) GetABI() string {
	return BridgeABI
}
