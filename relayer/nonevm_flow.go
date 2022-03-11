package relayer

// struct/funcs for interacting with flow chain, only support peg original vault(SafeBox) for now

import (
	"fmt"
	"sync"

	flowSigner "github.com/celer-network/cbridge-flow/signer"
	flowtypes "github.com/celer-network/cbridge-flow/types"
	flowutils "github.com/celer-network/cbridge-flow/utils"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/spf13/viper"
	dbm "github.com/tendermint/tm-db"
)

// all in one helper to do everything about flow chain
type FlowClient struct {
	ChainID   uint64
	Account   string
	PubkeyIdx int // to be compatible w/ flow sdk
	Db        *dbm.PrefixDB
	fcc       *flowutils.FlowCbrClient
	lock      sync.RWMutex

	BridgeAddr    string
	SafeBoxAddr   string // string from config, just flow addr hex
	PegBridgeAddr string // string from config, just flow addr hex
}

func NewFlowClient(cfg *common.OneChainConfig, cbrDb dbm.DB) *FlowClient {
	ret := &FlowClient{
		ChainID:   cfg.ChainID,
		Account:   viper.GetString(common.FlagFlowAccount),
		PubkeyIdx: viper.GetInt(common.FlagFlowPubkeyIndex),
		Db:        dbm.NewPrefixDB(cbrDb, []byte(fmt.Sprintf("%d", cfg.ChainID))),
		// save string address to be used when sign msg
		BridgeAddr:    cfg.CBridge, // TODO, case cbrtypes.SignDataType_SIGNERS need it?
		SafeBoxAddr:   cfg.OTVault,
		PegBridgeAddr: cfg.PTBridge,
	}
	if !commontypes.IsFlowChain(cfg.ChainID) {
		log.Fatalf("find invalid flow chainId:%d", cfg.ChainID)
	}
	// todo: support awskms
	signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	signer, err := flowSigner.NewFlowSigner(signerKey, signerPass)
	if err != nil {
		log.Fatalf("init flow signer err: %s", err.Error())
	}
	// net is defined enum name like FLOW_MAINNET so flowutils can replace eg. FungibleToken address correctly
	net := commontypes.NonEvmChainID(cfg.ChainID).String()
	// todo: pubkey idx
	ret.fcc, err = flowutils.NewFlowCbrClient(signer, cfg.Gateway, ret.Account, cfg.CBridge, cfg.OTVault, cfg.PTBridge, net, cfg.MaxBlkDelta)
	if err != nil {
		log.Fatalf("init flow transactor err: %s", err.Error())
	}
	return ret
}

func (f *FlowClient) monDeposited(interval uint64) {
	if f.SafeBoxAddr == "" {
		return
	}
	log.Infof("start monitor flow deposited, maxDelta:%d, interval:%v", f.fcc.MonMaxDelta, interval)
	err := f.fcc.Monitor(func(eLog *flowtypes.FlowMonitorLog) {
		log.Infof("Mon Flow deposit:%+v", eLog)
		serr := f.saveEvent(pegbrtypes.PegbrEventDeposited, eLog)
		if serr != nil {
			log.Errorf("saveFlowDepositedEvent err: %s", serr.Error())
			return
		}
		return
	}, f.fcc.DepositedEventId, interval)
	if err != nil {
		log.Fatalf("fail mon flow deposited, err:%s", err.Error())
	}
}

func (f *FlowClient) monWithdrawn(interval uint64) {
	if f.SafeBoxAddr == "" {
		return
	}
	log.Infof("start monitor flow withdrawn, maxDelta:%d, interval:%v", f.fcc.MonMaxDelta, interval)
	err := f.fcc.Monitor(func(eLog *flowtypes.FlowMonitorLog) {
		log.Infof("MonFlowWithdrawn: %x", eLog.TxHash)
		wdEv, perr := flowtypes.FlowSafeBoxWithdrawnUnmarshal(eLog.Event)
		if perr != nil {
			log.Errorf("file to parse flow withdrawn event, err:%s", perr.Error())
			return
		}
		log.Infof("MonFlowWithdrawn: %+v", wdEv)
		if CurRelayerInstance == nil {
			log.Errorln("CurRelayerInstance not initialized")
		} else {
			CurRelayerInstance.dbDelete(GetPegbrWdKey(f.ChainID, wdEv.RefChainId, wdEv.RefId[:]))
		}

		serr := f.saveEvent(pegbrtypes.PegbrEventWithdrawn, eLog)
		if serr != nil {
			log.Errorf("saveFlowWithdrawnEvent err: %s", serr.Error())
			return
		}
		return
	}, f.fcc.WithdrawnEventId, interval)
	if err != nil {
		log.Fatalf("fail mon flow withdrawn, err:%s", err.Error())
	}
}

func (f *FlowClient) monBurn(interval uint64) {
	if f.PegBridgeAddr == "" {
		return
	}
	log.Infof("start monitor flow burn, maxDelta:%d, interval:%v", f.fcc.MonMaxDelta, interval)
	err := f.fcc.Monitor(func(eLog *flowtypes.FlowMonitorLog) {
		log.Infof("Mon Flow burn:%x", eLog.TxHash)
		serr := f.saveEvent(pegbrtypes.PegbrEventBurn, eLog)
		if serr != nil {
			log.Errorf("saveFlowBurnEvent err: %s", serr.Error())
			return
		}
		return
	}, f.fcc.BurnEventId, interval)
	if err != nil {
		log.Fatalf("fail mon flow burn, err:%s", err.Error())
	}
}

func (f *FlowClient) monMint(interval uint64) {
	if f.PegBridgeAddr == "" {
		return
	}
	log.Infof("start monitor flow mint, maxDelta:%d, interval:%v", f.fcc.MonMaxDelta, interval)
	err := f.fcc.Monitor(func(eLog *flowtypes.FlowMonitorLog) {
		log.Infof("MonFlowMint: %x", eLog.TxHash)
		mintEv, perr := flowtypes.FlowFlowPegBridgeMintUnmarshal(eLog.Event)
		if perr != nil {
			log.Errorf("file to parse flow mint event, err:%s", perr.Error())
			return
		}
		log.Infof("MonFlowMint: %+v", mintEv)
		if CurRelayerInstance == nil {
			log.Errorln("CurRelayerInstance not initialized")
		} else {
			CurRelayerInstance.dbDelete(GetPegbrMintKey(f.ChainID, mintEv.RefChainId, mintEv.RefId[:]))
		}

		serr := f.saveEvent(pegbrtypes.PegbrEventMint, eLog)
		if serr != nil {
			log.Errorf("saveFlowMintEvent err: %s", serr.Error())
			return
		}
		return
	}, f.fcc.MintEventId, interval)
	if err != nil {
		log.Fatalf("fail mon flow mint, err:%s", err.Error())
	}
}

func (f *FlowClient) saveEvent(name string, eLog *flowtypes.FlowMonitorLog) error {
	f.lock.Lock()
	defer f.lock.Unlock()
	// TODO, different with evm, name-blocknum-transactionid-eventid enough for unique?
	key := fmt.Sprintf("%s-%d-%d-%d", name, eLog.Height, eLog.TransactionIndex, eLog.EventIndex)
	return f.Db.Set([]byte(key), eLog.Event)
}
