package relayer

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
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
	dbm "github.com/tendermint/tm-db"
)

type Relayer struct {
	*Operator
	db              dbm.DB
	ethMonitor      *monitor.Service
	verifiedUpdates *bigcache.BigCache
	valAddr         eth.Addr
	sgnAcct         sdk.AccAddress
	bonded          bool
	bootstrapped    bool // SGN is bootstrapped with at least one bonded validator on the eth contract
	startEthBlock   *big.Int
	lock            sync.RWMutex
}

func NewRelayer(operator *Operator, db dbm.DB) {
	relayerDb := dbm.NewPrefixDB(db, RelayerDbPrefix)

	watchService := watcher.NewWatchService(
		operator.EthClient.Client, newWatcherDAL(relayerDb), viper.GetUint64(common.FlagEthPollInterval),
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

	r := Relayer{
		Operator:        operator,
		db:              db,
		ethMonitor:      ethMonitor,
		verifiedUpdates: verifiedUpdates,
		valAddr:         eth.Hex2Addr(viper.GetString(common.FlagEthValidatorAddress)),
		bonded:          validatorStatus == eth.Bonded,
		bootstrapped:    bondedValNum.Uint64() > 0,
		startEthBlock:   startEthBlock,
	}

	r.sgnAcct, err = sdk.AccAddressFromBech32(viper.GetString(common.FlagSgnValidatorAccount))
	if err != nil {
		log.Fatalln("sgn acct error")
	}

	r.monitorEthValidatorNotice()
	r.monitorEthValidatorStatusUpdate()
	r.monitorEthDelegationUpdate()

	//go r.monitorSgnchainCreateValidator()

	go r.processQueues()

	NewCbridgeMgr(db) // do we need to save mgr somewhere?
}

func (r *Relayer) processQueues() {
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

	blkNum := r.getCurrentBlockNumber().Uint64()
	for {
		select {
		case <-pullerTicker.C:
			newblk := r.getCurrentBlockNumber().Uint64()
			if blkNum == newblk {
				continue
			}
			blkNum = newblk
			r.processPullerQueue()
			r.verifyPendingUpdates()

		case <-syncBlkTicker.C:
			r.syncBlkNum()

		case <-slashTicker.C:
			r.processSlashQueue()
		}
	}
}
