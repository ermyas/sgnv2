package monitor

import (
	"math/big"
	"time"

	"github.com/allegro/bigcache"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
	dbm "github.com/tendermint/tm-db"
)

const (
	prefixMonitor = "mon"
)

type Monitor struct {
	*Operator
	db              dbm.DB
	ethMonitor      *monitor.Service
	stakingContract monitor.Contract
	sgnContract     monitor.Contract
	verifiedChanges *bigcache.BigCache
	sidechainAcct   sdk.AccAddress
	bonded          bool
	executeSlash    bool
	bootstrapped    bool // SGN has bootstrapped with at least one bonded validator on the mainchain contract
	startBlock      *big.Int
	//lock            sync.RWMutex
}

func NewMonitor(operator *Operator, db dbm.DB) {
	monitorDb := dbm.NewPrefixDB(db, []byte(prefixMonitor))
	dal := newWatcherDAL(monitorDb)
	watchService := watcher.NewWatchService(operator.EthClient.Client, dal, viper.GetUint64(common.FlagEthPollInterval),
		viper.GetUint64(common.FlagEthMaxBlockDelta))
	if watchService == nil {
		log.Fatalln("Cannot create watch service")
	}

	blkDelay := viper.GetUint64(common.FlagEthBlockDelay)
	ethMonitor := monitor.NewService(watchService, blkDelay, true /* enabled */)
	ethMonitor.Init()

	stakingValidatorStatus, err := operator.EthClient.Staking.GetValidatorStatus(&bind.CallOpts{}, operator.EthClient.Address)
	if err != nil {
		log.Fatalln("GetValidatorStatus err", err)
	}

	valnum, err := operator.EthClient.Staking.GetValidatorNum(&bind.CallOpts{})
	if err != nil {
		log.Fatalln("GetValidatorNum err", err)
	}

	stakingContract := NewMonitorContractInfo(operator.EthClient.StakingAddress, contracts.StakingABI)
	sgnContract := NewMonitorContractInfo(operator.EthClient.SGNAddress, contracts.SGNABI)

	verifiedChanges, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Fatalln("NewBigCache err", err)
	}

	configuredStartBlock := viper.GetInt64(common.FlagEthMonitorStartBlock)
	var startBlock *big.Int
	if configuredStartBlock == 0 {
		startBlock = ethMonitor.GetCurrentBlockNumber()
	} else {
		startBlock = big.NewInt(viper.GetInt64(common.FlagEthMonitorStartBlock))
	}

	m := Monitor{
		Operator:        operator,
		db:              db,
		ethMonitor:      ethMonitor,
		stakingContract: stakingContract,
		sgnContract:     sgnContract,
		verifiedChanges: verifiedChanges,
		bonded:          contracts.IsBonded(stakingValidatorStatus),
		bootstrapped:    valnum.Uint64() > 0,
		executeSlash:    viper.GetBool(common.FlagSgnExecuteSlash),
		startBlock:      startBlock,
	}
	m.sidechainAcct, err = sdk.AccAddressFromBech32(viper.GetString(common.FlagSgnValidatorAccount))
	if err != nil {
		log.Fatalln("Sidechain acct error")
	}

	go m.monitorValidatorParamsUpdate()
	go m.monitorValidatorStatusUpdate()
	go m.monitorDelegationUpdate()
	go m.monitorSgnAddrUpdate()
}

func (m *Monitor) monitorValidatorParamsUpdate() {
}

func (m *Monitor) monitorValidatorStatusUpdate() {
}

func (m *Monitor) monitorDelegationUpdate() {
}

func (m *Monitor) monitorSgnAddrUpdate() {
}
