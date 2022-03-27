package common

import (
	"time"

	flowSigner "github.com/celer-network/cbridge-flow/signer"
	flowutils "github.com/celer-network/cbridge-flow/utils"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
)

type LogEventID struct {
	BlkNum uint64 // Number of the block containing the event
	Index  int64  // Index of the event within the block, use int to support -1 in fast forward case
}
type MockDAL map[string]LogEventID

func (d MockDAL) GetMonitorBlock(key string) (uint64, int64, bool, error) {
	le, ok := d[key]
	return le.BlkNum, le.Index, ok, nil
}

func (d MockDAL) SetMonitorBlock(key string, blockNum uint64, blockIdx int64) error {
	d[key] = LogEventID{BlkNum: blockNum, Index: blockIdx}
	return nil
}

// SetupFlowServiceAccountClient sets Client part (Client) and Auth part of serviceAccount
// bridge and safe box part is set after deploying contracts
func SetupFlowServiceAccountClient() error {
	var err error
	FlowServiceAccountSigner, err = flowSigner.NewFlowSigner(flowBaseKs, "")
	if err != nil {
		log.Fatal(err)
	}
	// no monitor
	FlowServiceAccountClient, err = flowutils.NewFlowCbrClient(LocalFlowChainId, LocalFlow, "", &flowutils.FlowSender{
		Signer:    FlowServiceAccountSigner,
		SenderHex: LocalFlowServiceAccount,
	}, make(MockDAL), mon2.PerChainCfg{}) // 0 blkintv disables update
	if err != nil {
		return err
	}
	return nil
}

func SetupContractFlowClient(bridgeAddr, safeBoxAddr, pegBridgeAddr string) {
	signer, err := flowSigner.NewFlowSigner(flowBaseKs, "")
	if err != nil {
		log.Fatal(err)
	}
	FlowContractAccountClient, err = flowutils.NewFlowCbrClient(LocalFlowChainId, LocalFlow, FlowContractAddr.Hex(), &flowutils.FlowSender{
		Signer:    signer,
		SenderHex: FlowContractAddr.Hex(),
	}, make(MockDAL), mon2.PerChainCfg{
		BlkIntv:     time.Second,
		MaxBlkDelta: 240,
	})
}

func SetupUserFlowClient(bridgeAddr, safeBoxAddr, pegBridgeAddr string) {
	signer, err := flowSigner.NewFlowSigner(flowBaseKs, "")
	if err != nil {
		log.Fatal(err)
	}
	FlowUserAccountClient, err = flowutils.NewFlowCbrClient(LocalFlowChainId, LocalFlow, FlowContractAddr.Hex(), &flowutils.FlowSender{
		Signer:    signer,
		SenderHex: FlowUserAddr.Hex(),
	}, make(MockDAL), mon2.PerChainCfg{
		BlkIntv:     time.Second,
		MaxBlkDelta: 240,
	})
}
