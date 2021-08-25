package monitor

import (
	"math/big"
	"sync"
	"time"

	"github.com/allegro/bigcache"
	"github.com/celer-network/goutils/eth/monitor"
	"github.com/celer-network/goutils/eth/watcher"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	vtypes "github.com/celer-network/sgn-v2/x/validator/types"
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
	verifiedUpdates *bigcache.BigCache
	sgnAcct         sdk.AccAddress
	bonded          bool
	bootstrapped    bool // SGN is bootstrapped with at least one bonded validator on the mainchain contract
	startEthBlock   *big.Int
	lock            sync.RWMutex
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
	ethMonitor := monitor.NewService(watchService, blkDelay, true)
	ethMonitor.Init()

	validatorStatus, err :=
		operator.EthClient.Contracts.Staking.GetValidatorStatus(&bind.CallOpts{}, operator.EthClient.Address)
	if err != nil {
		log.Fatalln("GetValidatorStatus err", err)
	}

	bondedValNum, err := operator.EthClient.Contracts.Staking.GetBondedValidatorNum(&bind.CallOpts{})
	if err != nil {
		log.Fatalln("GetValidatorNum err", err)
	}

	verifiedUpdates, err := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	if err != nil {
		log.Fatalln("NewBigCache err", err)
	}

	startEthBlock := big.NewInt(viper.GetInt64(common.FlagEthMonitorStartBlock))
	if startEthBlock.Sign() == 0 {
		startEthBlock = ethMonitor.GetCurrentBlockNumber()
	}

	m := Monitor{
		Operator:        operator,
		db:              db,
		ethMonitor:      ethMonitor,
		verifiedUpdates: verifiedUpdates,
		bonded:          validatorStatus == eth.Bonded,
		bootstrapped:    bondedValNum.Uint64() > 0,
		startEthBlock:   startEthBlock,
	}
	m.sgnAcct, err = vtypes.SdkAccAddrFromSgnBech32(viper.GetString(common.FlagSgnValidatorAccount))
	if err != nil {
		log.Fatalln("Sidechain acct error")
	}

	m.monitorEthValidatorParamsUpdate()
	m.monitorEthValidatorStatusUpdate()
	m.monitorEthDelegationUpdate()
	m.monitorEthSgnAddrUpdate()

	go m.monitorSgnchainCreateValidator()

	go m.processQueues()
}

func (m *Monitor) processQueues() {
	pullerInterval := time.Duration(viper.GetUint64(common.FlagEthPollInterval)) * time.Second
	syncBlkInterval := time.Duration(viper.GetUint64(common.FlagEthSyncBlkInterval)) * time.Second
	slashInterval := time.Duration(viper.GetUint64(common.FlagSgnCheckIntervalSlashQueue)) * time.Second
	log.Infof("Queue process interval: puller %s, slash %s", pullerInterval, slashInterval)

	pullerTicker := time.NewTicker(pullerInterval)
	syncBlkTicker := time.NewTicker(syncBlkInterval)
	slashTicker := time.NewTicker(slashInterval)
	defer func() {
		pullerTicker.Stop()
		syncBlkTicker.Stop()
		slashTicker.Stop()
	}()

	blkNum := m.getCurrentBlockNumber().Uint64()
	for {
		select {
		case <-pullerTicker.C:
			newblk := m.getCurrentBlockNumber().Uint64()
			if blkNum == newblk {
				continue
			}
			blkNum = newblk
			m.processPullerQueue()
			m.verifyPendingUpdates()

		case <-syncBlkTicker.C:
			m.syncBlkNum()

		case <-slashTicker.C:
			m.processSlashQueue()
		}
	}
}
