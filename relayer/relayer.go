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
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
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
	sgnAcct         sdk.AccAddress
	bonded          bool
	bootstrapped    bool // SGN is bootstrapped with at least one bonded validator on the eth contract
	startEthBlock   *big.Int
	syncer          Syncer
	lock            sync.RWMutex
	cbrMgr          CbrMgr
	cbrSsUpdating   bool
}

var CurRelayerInstance *Relayer

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
		operator.EthClient.Contracts.Staking.GetValidatorStatus(&bind.CallOpts{}, operator.ValAddr)
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
		bonded:          validatorStatus == eth.Bonded,
		bootstrapped:    bondedValNum.Uint64() > 0,
		startEthBlock:   startEthBlock,
	}

	CurRelayerInstance = &r

	r.sgnAcct, err = sdk.AccAddressFromBech32(viper.GetString(common.FlagSgnValidatorAccount))
	if err != nil {
		log.Fatalln("sgn acct error")
	}

	r.monitorEthValidatorNotice()
	r.monitorEthValidatorStatusUpdate()
	r.monitorEthDelegationUpdate()

	go r.monitorSgnSlash()
	go r.monitorSgnFarmingClaimAllEvent()
	go r.monitorSgnDistributionClaimAllStakingRewardEvent()

	go r.processPullerQueue()
	go r.processSlashQueue()
	go r.verifyPendingUpdates()

	r.cbrMgr = NewCbridgeMgr(db, r.Transactor.CliCtx) // do we need to save mgr somewhere?
	go r.monitorSgnCbrDataToSign()                    // cbr monitor set after cbrMgr initialization

	go r.doCbridgeSync(r.cbrMgr)
	r.doCbridgeOnchain(r.cbrMgr) // internal use goroutine
	go r.pullPriceChange()

	go r.checkSyncer()
}

type Syncer struct {
	isSyncer   bool
	updateTime time.Time
	lock       sync.RWMutex
}

func (r *Relayer) checkSyncer() {
	sgnBlkTime := viper.GetDuration(common.FlagConsensusTimeoutCommit)
	log.Infof("check syncer every %s", sgnBlkTime)
	for {
		time.Sleep(sgnBlkTime)
		syncer, err := stakingcli.QuerySyncer(r.Transactor.CliCtx)
		if err != nil {
			log.Errorln("Get syncer err", err)
		}
		isSyncerPrev := r.isSyncer()
		if eth.Hex2Addr(syncer.EthAddress) == r.Operator.ValAddr {
			// is current syncer
			if !isSyncerPrev {
				// just became syncer
				r.setSyncer(true)
			}
		} else {
			// is not current syncer
			if isSyncerPrev {
				// no longer a syncer
				r.setSyncer(false)
			}
		}
	}
}

func (r *Relayer) isSyncer() bool {
	r.syncer.lock.RLock()
	defer r.syncer.lock.RUnlock()
	return r.syncer.isSyncer
}

func (r *Relayer) getSyncer() (bool, time.Time) {
	r.syncer.lock.RLock()
	defer r.syncer.lock.RUnlock()
	return r.syncer.isSyncer, r.syncer.updateTime
}

func (r *Relayer) setSyncer(syncer bool) {
	r.syncer.lock.Lock()
	defer r.syncer.lock.Unlock()
	r.syncer.updateTime = time.Now()
	if syncer {
		r.syncer.isSyncer = true
		log.Debug("become a syncer")
	} else {
		r.syncer.isSyncer = false
		log.Debug("no longer a syncer")
	}
}
