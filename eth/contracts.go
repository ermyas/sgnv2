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

type StakingRewardContract struct {
	*StakingReward
	Address Addr
}

func NewStakingRewardContract(address Addr, client *ethclient.Client) (*StakingRewardContract, error) {
	stakingReward, err := NewStakingReward(address, client)
	if err != nil {
		return nil, err
	}
	return &StakingRewardContract{
		StakingReward: stakingReward,
		Address:       address,
	}, nil
}

func (c *StakingRewardContract) GetAddr() Addr {
	return c.Address
}

func (c *StakingRewardContract) GetABI() string {
	return StakingRewardABI
}

type FarmingRewardsContract struct {
	*FarmingRewards
	Address Addr
}

func NewFarmingRewardsContract(address Addr, client *ethclient.Client) (*FarmingRewardsContract, error) {
	farmingRewards, err := NewFarmingRewards(address, client)
	if err != nil {
		return nil, err
	}
	return &FarmingRewardsContract{
		FarmingRewards: farmingRewards,
		Address:        address,
	}, nil
}

func (c *FarmingRewardsContract) GetAddr() Addr {
	return c.Address
}

func (c *FarmingRewardsContract) GetABI() string {
	return FarmingRewardsABI
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

type PegBridgeContract struct {
	*PeggedTokenBridge
	Address Addr
}

func NewPegBridgeContract(address Addr, client *ethclient.Client) (*PegBridgeContract, error) {
	ptb, err := NewPeggedTokenBridge(address, client)
	if err != nil {
		return nil, err
	}
	return &PegBridgeContract{
		PeggedTokenBridge: ptb,
		Address:           address,
	}, nil
}

func (p *PegBridgeContract) GetAddr() Addr {
	return p.Address
}

func (p *PegBridgeContract) GetABI() string {
	return PeggedTokenBridgeABI
}

type PegVaultContract struct {
	*OriginalTokenVault
	Address Addr
}

func NewPegVaultContract(address Addr, client *ethclient.Client) (*PegVaultContract, error) {
	otv, err := NewOriginalTokenVault(address, client)
	if err != nil {
		return nil, err
	}
	return &PegVaultContract{
		OriginalTokenVault: otv,
		Address:            address,
	}, nil
}

func (o *PegVaultContract) GetAddr() Addr {
	return o.Address
}

func (o *PegVaultContract) GetABI() string {
	return OriginalTokenVaultABI
}

type MsgBusContract struct {
	*MessageBus
	Address Addr
}

func NewMsgBusContract(address Addr, client *ethclient.Client) (*MsgBusContract, error) {
	bus, err := NewMessageBus(address, client)
	if err != nil {
		return nil, err
	}
	return &MsgBusContract{
		MessageBus: bus,
		Address:    address,
	}, nil
}

func (c *MsgBusContract) GetAddr() Addr {
	return c.Address
}

func (c *MessageBus) GetABI() string {
	return MessageBusABI
}
