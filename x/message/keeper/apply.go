package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/message/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (k Keeper) ApplyEvent(ctx sdk.Context, data []byte) (bool, error) {
	applyEvent := new(cbrtypes.OnChainEvent)
	err := applyEvent.Unmarshal(data)
	if err != nil {
		return false, err
	}
	switch applyEvent.GetEvtype() {
	case types.MsgEventMessage:
		return k.applyMessage(ctx, applyEvent)
	case types.MsgEventMessageWithTransfer:
		return k.applyMessageWithTransfer(ctx, applyEvent)
	case types.MsgEventExecuted:
		return k.applyMessageExecuted(ctx, applyEvent)
	}
	return true, nil
}

func (k Keeper) applyMessage(ctx sdk.Context, applyEvent *cbrtypes.OnChainEvent) (bool, error) {
	ev, err := parseEventMessage(applyEvent)
	if err != nil {
		return false, err
	}
	msgId, msg := types.NewMessage(ev, applyEvent.Chainid)
	if k.HasMessage(ctx, msgId) {
		log.Debugf("skip already applied message (id %x)", msgId)
		return false, nil
	}
	log.Infof("x/msg applied Message, msgId:%x sender:%s receiver:%s chainId:%d->%d srcTx:%s",
		msgId, msg.Sender, msg.Receiver, msg.SrcChainId, msg.DstChainId, msg.SrcTxHash)
	k.SetActiveMessageId(ctx, msg.GetDstChainId(), eth.Hex2Addr(msg.GetReceiver()), msgId)
	k.SetMessage(ctx, msgId, msg)
	fee, success := new(big.Int).SetString(msg.GetFee(), 10)
	if !success {
		return false, fmt.Errorf("invalid fee")
	}
	err = k.MintFee(ctx, msg.SrcChainId, fee)
	if err != nil {
		return false, err
	}
	emitMessageToSign(ctx, msgId.Hex())
	return true, nil
}

func (k Keeper) applyMessageWithTransfer(ctx sdk.Context, applyEvent *cbrtypes.OnChainEvent) (bool, error) {
	ev, err := parseEventMessageWithTransfer(applyEvent)
	if err != nil {
		return false, err
	}
	srcXferId := eth.Bytes2Hash(ev.SrcTransferId[:])
	srcChainId := applyEvent.Chainid
	srcBridgeType := k.getSrcBridgeType(ctx, ev.Bridge, srcChainId)
	dstChainId := ev.DstChainId.Uint64()

	if k.HasSrcTransferMessageId(ctx, srcBridgeType, srcXferId) {
		log.Debugf("skip already applied message with srcTransferId %x, srcBridgeType %s", srcXferId, srcBridgeType)
		return false, nil
	}

	errMsg := fmt.Sprintf(
		"cannot apply message with transfer (srcXferId %x, srcChainId %d, srcBridgeType %v): ",
		srcXferId, srcChainId, srcBridgeType)

	switch srcBridgeType {
	case types.BRIDGE_TYPE_LIQUIDITY:
		xferStatus := k.cbridgeKeeper.QueryXferStatus(ctx, srcXferId)
		if shouldRefundXfer(xferStatus) {
			return k.applyLqTransferRefund(ctx, ev)
		}
		relay, found := k.cbridgeKeeper.GetXferRelay(ctx, srcXferId)
		if !found {
			return false, fmt.Errorf(errMsg + "relay not found")
		}
		relayOnChain := new(cbrtypes.RelayOnChain)
		err = relayOnChain.Unmarshal(relay.GetRelay())
		if err != nil {
			return false, fmt.Errorf(errMsg+"failed to unmarshal relay %x", relay.GetRelay())
		}
		dstToken := relayOnChain.GetToken()
		dstAmt := new(big.Int).SetBytes(relayOnChain.GetAmount()).String()
		dstBridge, found := k.cbridgeKeeper.GetCbrContractAddr(ctx, dstChainId)
		if !found {
			return false, fmt.Errorf(errMsg+"bridge addr not found for relay. dstChainId %d", dstChainId)
		}
		execCtx := types.NewMsgXferExecutionContext(
			ev, srcChainId, eth.Bytes2Addr(dstToken), dstAmt, dstBridge, types.TRANSFER_TYPE_LIQUIDITY_RELAY)
		return k.processMessageWithTransfer(ctx, execCtx, srcBridgeType)

	case types.BRIDGE_TYPE_PEG_VAULT:
		deposit, found := k.pegbridgeKeeper.GetDepositInfo(ctx, srcXferId)
		if !found {
			return false, fmt.Errorf(errMsg + "deposit info not found")
		}
		if len(deposit.GetMintId()) == 0 {
			return k.applyPegDepositRefund(ctx, ev, deposit.VaultVersion)
		}
		mint, found := k.pegbridgeKeeper.GetMintInfo(ctx, eth.Bytes2Hash(deposit.GetMintId()))
		if !found {
			return false, fmt.Errorf(errMsg+"mint info (id %x) not found", deposit.GetMintId())
		}
		mintOnChain := new(pegbrtypes.MintOnChain)
		err = mintOnChain.Unmarshal(mint.GetMintProtoBytes())
		if err != nil {
			return false, fmt.Errorf(errMsg+"unable to unmarshal mintOnchain %v", mint.GetMintProtoBytes())
		}
		dstAmt := new(big.Int).SetBytes(mintOnChain.GetAmount()).String()
		dstToken := mintOnChain.GetToken()
		dstBridge, found := k.pegbridgeKeeper.GetPeggedBridge(ctx, dstChainId, mint.BridgeVersion)
		if !found {
			return false, fmt.Errorf(errMsg+"pegged token bridge not found for dstChainId %d", dstChainId)
		}
		xferType := types.TRANSFER_TYPE_PEG_MINT
		switch mint.BridgeVersion {
		case 0:
			xferType = types.TRANSFER_TYPE_PEG_MINT
		case 2:
			xferType = types.TRANSFER_TYPE_PEG_V2_MINT
		default:
			return false, fmt.Errorf("invalid bridge version %d", mint.BridgeVersion)
		}
		execCtx := types.NewMsgXferExecutionContext(ev, srcChainId, eth.Bytes2Addr(dstToken), dstAmt, dstBridge, xferType)
		return k.processMessageWithTransfer(ctx, execCtx, srcBridgeType)

	case types.BRIDGE_TYPE_PEG_BRIDGE:
		burn, found := k.pegbridgeKeeper.GetBurnInfo(ctx, srcXferId)
		if !found {
			return false, fmt.Errorf(errMsg + "burn info not found")
		}
		if len(burn.GetWithdrawId()) == 0 && len(burn.GetMintId()) == 0 {
			return k.applyPegBurnRefund(ctx, ev, burn.BridgeVersion)
		}
		var dstToken, dstBridge eth.Addr
		var xferType types.TransferType
		var dstAmt string
		if len(burn.GetWithdrawId()) > 0 {
			withdraw, found := k.pegbridgeKeeper.GetWithdrawInfo(ctx, eth.Bytes2Hash(burn.GetWithdrawId()))
			if !found {
				return false, fmt.Errorf(errMsg+"withdraw info (id %x) not found", burn.GetWithdrawId())
			}
			withdrawOnChain := new(pegbrtypes.WithdrawOnChain)
			err = withdrawOnChain.Unmarshal(withdraw.GetWithdrawProtoBytes())
			if err != nil {
				return false, fmt.Errorf(errMsg+"unable to unmarshal withdrawOnchain %v", withdraw.GetWithdrawProtoBytes())
			}
			dstAmt = new(big.Int).SetBytes(withdrawOnChain.GetAmount()).String()
			dstToken = eth.Bytes2Addr(withdrawOnChain.GetToken())
			dstBridge, found = k.pegbridgeKeeper.GetOriginalVault(ctx, dstChainId, withdraw.VaultVersion)
			if !found {
				return false, fmt.Errorf(errMsg+"pegged token vault not found for dstChainId %d", dstChainId)
			}
			switch withdraw.VaultVersion {
			case 0:
				xferType = types.TRANSFER_TYPE_PEG_WITHDRAW
			case 2:
				xferType = types.TRANSFER_TYPE_PEG_V2_WITHDRAW
			default:
				return false, fmt.Errorf("invalid vault version %d", withdraw.VaultVersion)
			}
		} else { //mint found instead of withdraw
			mint, found := k.pegbridgeKeeper.GetMintInfo(ctx, eth.Bytes2Hash(burn.GetMintId()))
			if !found {
				return false, fmt.Errorf(errMsg+"mint info (id %x) not found", burn.GetMintId())
			}
			mintOnChain := new(pegbrtypes.MintOnChain)
			err = mintOnChain.Unmarshal(mint.GetMintProtoBytes())
			if err != nil {
				return false, fmt.Errorf(errMsg+"unable to unmarshal mintOnchain %v", mint.GetMintProtoBytes())
			}
			dstAmt = new(big.Int).SetBytes(mintOnChain.GetAmount()).String()
			dstToken = eth.Bytes2Addr(mintOnChain.GetToken())
			dstBridge, found = k.pegbridgeKeeper.GetPeggedBridge(ctx, dstChainId, mint.BridgeVersion)
			if !found {
				return false, fmt.Errorf(errMsg+"pegged token bridge not found for dstChainId %d", dstChainId)
			}
			switch mint.BridgeVersion {
			case 0:
				xferType = types.TRANSFER_TYPE_PEG_MINT
			case 2:
				xferType = types.TRANSFER_TYPE_PEG_V2_MINT
			default:
				return false, fmt.Errorf("invalid bridge version %d", mint.BridgeVersion)
			}
		}
		execCtx := types.NewMsgXferExecutionContext(ev, srcChainId, dstToken, dstAmt, dstBridge, xferType)
		return k.processMessageWithTransfer(ctx, execCtx, srcBridgeType)
	}
	return false, fmt.Errorf(errMsg + "src bridge type not supported")
}

func (k Keeper) applyLqTransferRefund(ctx sdk.Context, ev *eth.MessageBusMessageWithTransfer) (bool, error) {
	wdOnchain := k.cbridgeKeeper.QueryXferRefund(ctx, ev.SrcTransferId)
	if wdOnchain == nil {
		return false, fmt.Errorf("wdOnchain not found for srcXferId %x", ev.SrcTransferId)
	}
	nonce := wdOnchain.GetSeqnum()
	if nonce == 0 {
		nonce = k.genLqBridgeRefundNonce(ctx)
		err := k.cbridgeKeeper.Refund(ctx, ev.SrcTransferId, nonce)
		if err != nil {
			return false, fmt.Errorf("refund err: %w", err)
		}
	}

	bridge, found := k.cbridgeKeeper.GetCbrContractAddr(ctx, wdOnchain.Chainid)
	if !found {
		return false, fmt.Errorf("bridge addr not found for chainId %d", wdOnchain.Chainid)
	}
	execCtx := types.NewMsgXferRefundExecutionContext(ev, wdOnchain, nonce, bridge)
	return k.processMessageWithTransfer(ctx, execCtx, types.BRIDGE_TYPE_LIQUIDITY)
}

func (k Keeper) applyPegDepositRefund(ctx sdk.Context, ev *eth.MessageBusMessageWithTransfer, vaultVersion uint32) (bool, error) {
	depositId := ev.SrcTransferId
	wdOnChain, found := k.pegbridgeKeeper.GetDepositRefund(ctx, depositId)
	if !found {
		return false, fmt.Errorf("wdOnChain not found for depositId %x", depositId)
	}
	bridge, found := k.pegbridgeKeeper.GetOriginalVault(ctx, wdOnChain.RefChainId, vaultVersion)
	if !found {
		return false, fmt.Errorf("otvault addr not found for chainId %d", wdOnChain.RefChainId)
	}
	execCtx := types.NewMsgPegDepositRefundExecutionContext(ev, wdOnChain, bridge, vaultVersion)
	return k.processMessageWithTransfer(ctx, execCtx, types.BRIDGE_TYPE_PEG_VAULT)
}

func (k Keeper) applyPegBurnRefund(ctx sdk.Context, ev *eth.MessageBusMessageWithTransfer, bridgeVersion uint32) (bool, error) {
	burnId := ev.SrcTransferId
	mintOnChain, found := k.pegbridgeKeeper.GetBurnRefund(ctx, burnId)
	if !found {
		return false, fmt.Errorf("mintOnChain not found for burnId %x", burnId)
	}
	bridge, found := k.pegbridgeKeeper.GetPeggedBridge(ctx, mintOnChain.RefChainId, bridgeVersion)
	if !found {
		return false, fmt.Errorf("ptbridge addr not found for chainId %d", mintOnChain.RefChainId)
	}
	execCtx := types.NewMsgPegBurnRefundExecutionContext(ev, mintOnChain, bridge, bridgeVersion)
	return k.processMessageWithTransfer(ctx, execCtx, types.BRIDGE_TYPE_PEG_BRIDGE)
}

func (k Keeper) processMessageWithTransfer(
	ctx sdk.Context, execCtx *types.ExecutionContext, srcBridgeType types.BridgeType) (bool, error) {

	messageId := eth.Bytes2Hash(execCtx.MessageId)
	if k.HasMessage(ctx, messageId) {
		return false, fmt.Errorf("message with transfer already applied. %x", messageId)
	}

	msg := execCtx.Message
	log.Infof("x/msg applied MessageWithTransfer, msgId:%x sender:%s receiver:%s chainId:%d->%d srcBridge:%s transferType:%s transferRefId:%x srcTx:%s",
		messageId, msg.Sender, msg.Receiver, msg.SrcChainId, msg.DstChainId, srcBridgeType, msg.TransferType, msg.TransferRefId, msg.SrcTxHash)

	k.SetActiveMessageId(ctx, msg.GetDstChainId(), eth.Hex2Addr(msg.GetReceiver()), messageId)
	k.SetMessage(ctx, messageId, &msg)
	k.SetSrcTransferMessageId(ctx, srcBridgeType, eth.Bytes2Hash(msg.GetTransferRefId()), messageId)
	fee, success := new(big.Int).SetString(msg.GetFee(), 10)
	if !success {
		return false, fmt.Errorf("invalid fee")
	}
	err := k.MintFee(ctx, msg.GetSrcChainId(), fee)
	if err != nil {
		return false, err
	}
	emitMessageToSign(ctx, messageId.Hex())
	return true, nil
}

// DeleteActiveMessageId and set message completed after executed
func (k Keeper) applyMessageExecuted(ctx sdk.Context, applyEvent *cbrtypes.OnChainEvent) (bool, error) {
	ev, err := parseEventExecuted(applyEvent)
	if err != nil {
		log.Errorln("getMessageId: cannot parse event:", err)
		return false, err
	}
	msg, found := k.GetMessage(ctx, ev.MsgId)
	if !found {
		return false, fmt.Errorf("msg not found for msgId: %x", ev.MsgId)
	}
	if msg.ExecutionStatus != types.EXECUTION_STATUS_NULL {
		log.Debugf("skip already applied message (id %x)", ev.MsgId)
		return false, nil
	}
	log.Infof("x/msg applied Executed: chain %d, %s", msg.DstChainId, ev)

	// remove the active message record
	k.DeleteActiveMessageId(ctx, msg.GetDstChainId(), eth.Hex2Addr(msg.GetReceiver()), ev.MsgId)
	status := types.ExecutionStatus(ev.Status)
	if status == types.EXECUTION_STATUS_NULL {
		return false, fmt.Errorf("error pending status for msgId:%x status%d", ev.MsgId, ev.Status)
	}

	// update msg status
	msg.ExecutionStatus = status
	k.SetMessage(ctx, ev.MsgId, &msg)
	return true, nil
}

func (k Keeper) getSrcBridgeType(ctx sdk.Context, bridgeAddr eth.Addr, srcChainId uint64) types.BridgeType {
	cbrContractAddr, cbrFound := k.cbridgeKeeper.GetCbrContractAddr(ctx, srcChainId)
	if cbrFound && cbrContractAddr == bridgeAddr {
		return types.BRIDGE_TYPE_LIQUIDITY
	}
	vault, vaultFound := k.pegbridgeKeeper.GetOriginalVault(ctx, srcChainId, 0)
	if vaultFound && vault == bridgeAddr {
		return types.BRIDGE_TYPE_PEG_VAULT
	}
	pegbr, pegbrFound := k.pegbridgeKeeper.GetPeggedBridge(ctx, srcChainId, 0)
	if pegbrFound && pegbr == bridgeAddr {
		return types.BRIDGE_TYPE_PEG_BRIDGE
	}
	vault, vaultFound = k.pegbridgeKeeper.GetOriginalVault(ctx, srcChainId, 2)
	if vaultFound && vault == bridgeAddr {
		return types.BRIDGE_TYPE_PEG_VAULT
	}
	pegbr, pegbrFound = k.pegbridgeKeeper.GetPeggedBridge(ctx, srcChainId, 2)
	if pegbrFound && pegbr == bridgeAddr {
		return types.BRIDGE_TYPE_PEG_BRIDGE
	}
	return types.BRIDGE_TYPE_NULL
}

func parseEventMessageWithTransfer(event *cbrtypes.OnChainEvent) (*eth.MessageBusMessageWithTransfer, error) {
	evlog := new(ethtypes.Log)
	err := json.Unmarshal(event.GetElog(), evlog)
	if err != nil {
		return nil, err
	}
	msgContract, err := eth.NewMessageBusFilterer(eth.ZeroAddr, nil)
	if err != nil {
		return nil, err
	}
	msgXfer, err := msgContract.ParseMessageWithTransfer(*evlog)
	return msgXfer, err
}

func parseEventMessage(event *cbrtypes.OnChainEvent) (*eth.MessageBusMessage, error) {
	evlog := new(ethtypes.Log)
	err := json.Unmarshal(event.GetElog(), evlog)
	if err != nil {
		return nil, err
	}
	msgContract, err := eth.NewMessageBusFilterer(eth.ZeroAddr, nil)
	if err != nil {
		return nil, err
	}
	msg, err := msgContract.ParseMessage(*evlog)
	return msg, err
}

func parseEventExecuted(event *cbrtypes.OnChainEvent) (*eth.MessageBusExecuted, error) {
	evlog := new(ethtypes.Log)
	err := json.Unmarshal(event.GetElog(), evlog)
	if err != nil {
		return nil, err
	}
	msgContract, err := eth.NewMessageBusFilterer(eth.ZeroAddr, nil)
	if err != nil {
		return nil, err
	}
	executed, err := msgContract.ParseExecuted(*evlog)
	if err != nil {
		log.Errorln("getMessageId: cannot parse event:", err)
		return nil, err
	}
	return executed, nil
}

func shouldRefundXfer(s cbrtypes.XferStatus) (shouldRefund bool) {
	return s == cbrtypes.XferStatus_BAD_LIQUIDITY ||
		s == cbrtypes.XferStatus_BAD_SLIPPAGE ||
		s == cbrtypes.XferStatus_BAD_XFER_DISABLED ||
		s == cbrtypes.XferStatus_BAD_DEST_CHAIN ||
		s == cbrtypes.XferStatus_EXCEED_MAX_OUT_AMOUNT
}

func emitMessageToSign(ctx sdk.Context, id string) {
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeMessageToSign,
		sdk.NewAttribute(types.AttributeKeyMessageId, id),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
}
