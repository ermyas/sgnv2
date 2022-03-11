package common

import (
	flowSigner "github.com/celer-network/cbridge-flow/signer"
	flowutils "github.com/celer-network/cbridge-flow/utils"
	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
)

// SetupFlowServiceAccountClient sets Client part (Client) and Auth part of serviceAccount
// bridge and safe box part is set after deploying contracts
func SetupFlowServiceAccountClient() error {
	var err error
	net := commontypes.NonEvmChainID(LocalFlowChainId).String()
	FlowServiceAccountSigner, err = flowSigner.NewFlowSigner(flowBaseKs, "")
	if err != nil {
		log.Fatal(err)
	}
	FlowServiceAccountClient, err = flowutils.NewFlowCbrClient(FlowServiceAccountSigner, LocalFlow, LocalFlowServiceAccount, "", "", "", net, 240)
	if err != nil {
		return err
	}
	return nil
}

func SetupContractFlowClient(bridgeAddr, safeBoxAddr, pegBridgeAddr string) {
	net := commontypes.NonEvmChainID(LocalFlowChainId).String()
	signer, err := flowSigner.NewFlowSigner(flowBaseKs, "")
	if err != nil {
		log.Fatal(err)
	}
	FlowContractAccountClient, err = flowutils.NewFlowCbrClient(signer, LocalFlow, FlowContractAddr.String(), bridgeAddr, safeBoxAddr, pegBridgeAddr, net, 240)
}

func SetupUserFlowClient(bridgeAddr, safeBoxAddr, pegBridgeAddr string) {
	net := commontypes.NonEvmChainID(LocalFlowChainId).String()
	signer, err := flowSigner.NewFlowSigner(flowBaseKs, "")
	if err != nil {
		log.Fatal(err)
	}
	FlowUserAccountClient, err = flowutils.NewFlowCbrClient(signer, LocalFlow, FlowUserAddr.String(), bridgeAddr, safeBoxAddr, pegBridgeAddr, net, 240)
}
