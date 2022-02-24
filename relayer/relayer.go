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
	db                 dbm.DB
	ethMonitor         *monitor.Service
	verifiedUpdates    *bigcache.BigCache
	sgnAcct            sdk.AccAddress
	bonded             bool
	bootstrapped       bool // SGN is bootstrapped with at least one bonded validator on the eth contract
	startEthBlock      *big.Int
	syncer             Syncer
	lock               sync.RWMutex
	cbrMgr             CbrMgr
	cbrSsUpdating      bool
	chainMonitorStatus ChainMonitorStatus
}

type ChainMonitorStatus int32

const (
	ChainMonitorStatusNull ChainMonitorStatus = iota // unknown
	ChainMonitorStatusYes                            // monitoring
	ChainMonitorStatusNo                             // unmonitoring
)

var CurRelayerInstance *Relayer

func NewRelayer(operator *Operator, db dbm.DB) {

	if viper.GetBool(common.FlagSgnWitnessMode) {
		log.Infoln("Entering witness mode")
		// TODO: report LpEarning and BaseFee distribution in witness mode
		return
	}

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

	bigCacheCfg := bigcache.DefaultConfig(20 * time.Minute)
	bigCacheCfg.CleanWindow = time.Hour
	verifiedUpdates, err := bigcache.NewBigCache(bigCacheCfg)
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
	go r.monitorSgnDistributionClaimMessageFeesEvent()

	r.cbrMgr = NewCbridgeMgr(db, r.Transactor.CliCtx) // cbrMgr should be initialized before verifyPendingUpdates
	go r.monitorSgnCbrDataToSign()                    // cbr monitor set after cbrMgr initialization
	go r.monitorSgnPegMintToSign()
	go r.monitorSgnPegWithdrawToSign()
	go r.monitorSgnMsgDataToSign()

	r.startReportSgnAnalytics()

	go r.processPullerQueue()
	go r.processSlashQueue()
	go r.verifyPendingUpdates()

	go r.doCbridgeSync(r.cbrMgr)
	r.doCbridgeOnchain(r.cbrMgr) // internal use goroutine
	go r.pullPriceChange()

	go r.doPegbrSync(r.cbrMgr)
	r.doPegbrOnchain(r.cbrMgr) // internal use goroutine

	go r.doMsgbrSync(r.cbrMgr)

	go r.checkSyncer()

	go r.monitorChain()
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
			continue
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

func (r *Relayer) monitorChain() {
	sgnBlkTime := viper.GetDuration(common.FlagConsensusTimeoutCommit)
	time.Sleep(sgnBlkTime)
	checkInterval := sgnBlkTime * 10
	log.Infof("check syncer candidates every %s", checkInterval)
	for {
		if r.chainMonitorStatus != ChainMonitorStatusNull {
			time.Sleep(checkInterval)
		}
		stakingParams, err := stakingcli.QueryParams(r.Transactor.CliCtx)
		if err != nil {
			log.Errorln("Get staking params err", err)
			continue
		}

		if len(stakingParams.SyncerCandidates) == 0 || containsAddr(stakingParams.SyncerCandidates, r.Operator.ValAddr) {
			if r.chainMonitorStatus != ChainMonitorStatusYes {
				log.Infoln("start bridge monitoring")
				for _, oc := range CbrMgrInstance {
					oc.startMon()
				}
				r.chainMonitorStatus = ChainMonitorStatusYes
			}
		} else {
			if r.chainMonitorStatus == ChainMonitorStatusYes {
				log.Infoln("close bridge monitoring")
				for _, oc := range CbrMgrInstance {
					oc.mon.Close()
				}
			}
			r.chainMonitorStatus = ChainMonitorStatusNo
			// if syncer candidates has been set and I am not the syncer candidate, stop the for loop. As in practice, after syncer candidates is set,
			// it's rare to change. In case it's changed and I am the syncer candidate again, workaround is to restart the sgnd.
			log.Infoln("not monitoring bridge chains")
			break
		}
	}
}

func containsAddr(addrs []string, addr eth.Addr) (found bool) {
	for _, a := range addrs {
		if eth.Hex2Addr(a) == addr {
			found = true
			break
		}
	}
	return
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
