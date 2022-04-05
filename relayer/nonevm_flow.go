package relayer

// struct/funcs for interacting with flow chain, only support peg original vault(SafeBox) for now

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	flowSigner "github.com/celer-network/cbridge-flow/signer"
	flowtypes "github.com/celer-network/cbridge-flow/types"
	flowutils "github.com/celer-network/cbridge-flow/utils"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	pbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/spf13/viper"
	dbm "github.com/tendermint/tm-db"
)

// all in one helper to do everything about flow chain
type FlowClient struct {
	fcc  *flowutils.FlowCbrClient
	Db   *dbm.PrefixDB // save event and monitor
	lock sync.RWMutex  // serialize db write

	ChainID      uint64
	ContractAddr string // needed for EncodeDataToSign

	txLock sync.Mutex // sgn send flow mint and withdraw serially to prevent nonce conflict
}

// flow chain id to monitor polling interval
var chid2Intv = map[uint64]time.Duration{
	12340001: time.Minute,
	12340002: time.Minute,
	12340003: time.Second, // emulator
}

// wdal is for persist monitor block, must impl mon2.DAL funcs. db is to save event in monitor callback and later used by puller
func NewFlowClient(cfg *common.OneChainConfig, wdal *watcherDAL, db *dbm.PrefixDB) *FlowClient {
	// check basic config correctness
	if !commontypes.IsFlowChain(cfg.ChainID) {
		log.Fatalln("invalid flow chainId:", cfg.ChainID)
	}
	if cfg.CBridge != cfg.PTBridge || cfg.CBridge != cfg.OTVault {
		log.Fatalln("mismatch contract addr. all flow contracts must be under same account.", cfg.CBridge, cfg.PTBridge, cfg.OTVault)
	}
	var sender *flowutils.FlowSender
	var err error
	// for partner, no need to send transactions on flow, then set nil sender
	if viper.GetString(common.FlagFlowAccount) != "" {
		sender, err = buildFlowSender()
		if err != nil {
			log.Fatalln("init flow signer err:", err)
		}
	} else {
		log.Warnf("this node will not send tx on flow, flow sender is nil")
	}
	// now build return obj
	ret := &FlowClient{
		ChainID:      cfg.ChainID,
		Db:           db,
		ContractAddr: cfg.CBridge,
	}
	ret.fcc, err = flowutils.NewFlowCbrClient(cfg.ChainID, cfg.Gateway, cfg.CBridge, sender, wdal, mon2.PerChainCfg{
		BlkIntv:     time.Duration(cfg.BlkInterval) * time.Second,
		MaxBlkDelta: cfg.MaxBlkDelta,
		// other fields don't apply to flow chain
	})
	if err != nil {
		log.Fatalf("init flow transactor err: %s", err.Error())
	}
	return ret
}

// parse viper flags and return sender for NewFlowCbrClient
func buildFlowSender() (*flowutils.FlowSender, error) {
	// build sender
	sender := &flowutils.FlowSender{
		SenderHex: viper.GetString(common.FlagFlowAccount),
		KeyIdx:    viper.GetInt(common.FlagFlowPubkeyIndex),
	}
	// set up sender.Signer
	var err error
	signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	region, kayalias := eth.ParseAwsKms(signerKey)
	if region != "" {
		sender.Signer, err = flowSigner.NewFlowKmsSigner(region, kayalias)
	} else {
		sender.Signer, err = flowSigner.NewFlowSigner(signerKey, signerPass)
	}
	return sender, err
}

// polling interval is done automatically by chainid, emulator: 1sec, test/main: 1min
func (f *FlowClient) monitorFlow() {
	intv := chid2Intv[f.ChainID]
	// must async call
	go f.fcc.Monitor(f.genEvCallback(pbrtypes.PegbrEventDeposited), flowutils.SafeBoxDepositedIdFmt, intv)
	go f.fcc.Monitor(f.genEvCallback(pbrtypes.PegbrEventWithdrawn), flowutils.SafeBoxWithdrawnIdFmt, intv)
	go f.fcc.Monitor(f.genEvCallback(pbrtypes.PegbrEventMint), flowutils.PegBridgeMintIdFmt, intv)
	go f.fcc.Monitor(f.genEvCallback(pbrtypes.PegbrEventBurn), flowutils.PegBridgeBurnIdFmt, intv)
}

// generate per event handler, evname is solidity event for consistency.
// we could split flow event type and use last part, but it's not explicit and prone to future error
func (f *FlowClient) genEvCallback(evname string) func(*flowtypes.FlowMonitorLog) {
	return func(ev *flowtypes.FlowMonitorLog) {
		log.Infoln("Mon Flow ev", ev.Type, string(ev.Event))
		key := fmt.Sprintf("%s-%d-%d-%d", evname, ev.Height, ev.TransactionIndex, ev.EventIndex)
		raw, _ := json.Marshal(ev)
		f.lock.Lock()
		defer f.lock.Unlock()
		f.Db.Set([]byte(key), raw)
	}
}

func (f *FlowClient) sendWithdraw(logmsg string, msg []byte, tokeId string, pubKeySig [][]byte) {
	f.txLock.Lock()
	defer f.txLock.Unlock()
	exist, err := f.fcc.QuerySafeBoxRecordExist(context.Background(), fmt.Sprintf("%x", pbrtypes.StdSHA3Hash(msg)))
	if err != nil {
		log.Warnf("query flow wd fail, err:%v", err)
		// still continue to send the tx, no break
	} else if exist {
		log.Infof("flow wd %s exist, skip", logmsg)
		return
	}
	txHash, err := f.fcc.Withdraw(context.Background(), msg, tokeId, pubKeySig)
	if err != nil {
		log.Errorf("send flow wd fail, msg:%s, err: %v", logmsg, err)
		return
	}
	withTxErr, err := f.fcc.WaitTxSealed(context.Background(), txHash, 30*time.Second)
	if err != nil {
		if withTxErr && !strings.Contains(err.Error(), "wdId already exists") {
			log.Errorf("send flow wd fail with tx err, msg:%s, err:%v", logmsg, err)
		} else {
			log.Warnf("wait flow wd fail, msg:%s, err:%v", logmsg, err)
		}
		return
	}
	log.Infof("send flow wd success, msg:%s", logmsg)
	return
}

func (f *FlowClient) sendMint(logmsg string, msg []byte, tokeId string, pubKeySig [][]byte) {
	f.txLock.Lock()
	defer f.txLock.Unlock()
	exist, err := f.fcc.QueryPegBridgeRecordExist(context.Background(), fmt.Sprintf("%x", pbrtypes.StdSHA3Hash(msg)))
	if err != nil {
		log.Warnf("query flow mint fail, err:%v", err)
		// still continue to send the tx, no break
	} else if exist {
		log.Infof("flow mint %s exist, skip", logmsg)
		return
	}
	txHash, err := f.fcc.Mint(context.Background(), msg, tokeId, pubKeySig)
	if err != nil {
		log.Errorf("send flow mint fail, msg:%s, err: %v", logmsg, err)
		return
	}
	withTxErr, err := f.fcc.WaitTxSealed(context.Background(), txHash, 30*time.Second)
	if err != nil {
		if withTxErr && !strings.Contains(err.Error(), "mintId already exists") {
			log.Errorf("send flow mint fail with tx err, but fail, msg:%s, err:%v", logmsg, err)
		} else {
			log.Warnf("wait flow mint fail, msg:%s, err:%v", logmsg, err)
		}
		return
	}
	log.Infof("Send Flow mint success, msg:%s", logmsg)
	return
}
