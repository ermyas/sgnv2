package keeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ApplyEvent(ctx sdk.Context, data []byte) (bool, error) {
	applyEvent := new(cbrtypes.OnChainEvent)
	err := applyEvent.Unmarshal(data)
	if err != nil {
		return false, err
	}
	log.Debugf("x/message applied:%+v", applyEvent.String())
	switch applyEvent.GetEvtype() {
	case types.MsgEventMessage:
		return k.applyMessageNoTransfer(ctx, applyEvent)
	case types.MsgEventMessageWithTransfer:
		return k.applyMessageWithTransfer(ctx, applyEvent)
	case types.MsgEventExecuted:
		return k.applyMessageExecuted(ctx, applyEvent)
	}
	return true, nil
}

func (k Keeper) applyMessageNoTransfer(ctx sdk.Context, applyEvent *cbrtypes.OnChainEvent) (bool, error) {
	execCtx, err := buildExecutionContextForMessageNoTransfer(ctx, applyEvent)
	if err != nil {
		return false, err
	}
	msg := execCtx.Message
	// The messageId for a message without a transfer
	messageId := eth.Bytes2Hash(execCtx.MessageId)
	if k.HasMessage(ctx, messageId) {
		log.Infof("skip already applied message (id %s)", messageId.String())
		return false, nil
	}
	log.Debugf("message applied, msgId:%s", messageId)
	k.SetActiveMessageId(ctx, msg.GetDstChainId(), eth.Hex2Addr(msg.GetReceiver()), messageId)
	k.SetMessage(ctx, messageId, msg)
	fee, success := new(big.Int).SetString(msg.GetFee(), 10)
	if !success {
		return false, errors.New("invalid fee")
	}
	err = k.MintFee(ctx, msg.SrcChainId, fee)
	if err != nil {
		return false, err
	}
	log.Debugln("emitting tendermint event with messageId", messageId.Hex())
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeMessageToSign,
		sdk.NewAttribute(types.AttributeKeyMessageId, messageId.Hex()),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
	return true, nil
}

func (k Keeper) applyMessageWithTransfer(ctx sdk.Context, applyEvent *cbrtypes.OnChainEvent) (bool, error) {
	ev, err := parseMessageWithTransfer(applyEvent)
	if err != nil {
		return false, err
	}
	srcXferId := eth.Bytes2Hash(ev.SrcTransferId[:])
	srcChainId := applyEvent.Chainid
	transferType := k.getTransferType(ctx, ev.Bridge, srcChainId)

	// Query the respective x modules to check if the transfer is in refund status
	switch transferType {
	case types.TRANSFER_TYPE_LIQUIDITY_SEND:
		xferStatus := k.cbridgeKeeper.QueryXferStatus(ctx, srcXferId)
		if shouldRefund(xferStatus) {
			return k.applyTransferRefund(ctx, ev)
		}
	case types.TRANSFER_TYPE_PEG_MINT:
		// TODO
	case types.TRANSFER_TYPE_PEG_WITHDRAW:
		// TODO
	default:
		return false, fmt.Errorf("cannot determine refund status: msg (xferId %s) transfer type (%s) not supported", ev.SrcTransferId, transferType)
	}

	// process successful transfer case
	dstToken, dstAmt, dstBridge, err := k.getTransferInfo(ctx, srcXferId, transferType, ev.DstChainId.Uint64())
	if err != nil {
		return false, fmt.Errorf("getTransferByTransferId: cannot get transfer:%s now, err:%s", srcXferId, err)
	}
	execCtx := types.NewExecutionContext(ev, srcChainId, eth.Bytes2Addr(dstToken), dstAmt, dstBridge, transferType)
	messageId := eth.Bytes2Hash(execCtx.MessageId)
	if k.HasMessage(ctx, messageId) {
		log.Infof("skip already applied message with transfer (id %s)", messageId.String())
		return false, nil
	}
	msg := execCtx.Message
	// messageId with transfer: keccak256(abi.encodePacked(MsgType.MessageWithTransfer, dstBridgeAddr, dstTransferId))
	log.Debugf("message applied, msgId:%s, srcTransferId:%s", messageId, eth.Bytes2Hash(execCtx.Transfer.GetRefId()))
	k.SetActiveMessageId(ctx, msg.GetDstChainId(), eth.Hex2Addr(msg.GetReceiver()), messageId)
	k.SetMessage(ctx, messageId, msg)
	k.SetTransfer(ctx, messageId, *execCtx.Transfer)
	fee, success := new(big.Int).SetString(msg.GetFee(), 10)
	if !success {
		return false, errors.New("invalid fee")
	}
	err = k.MintFee(ctx, msg.SrcChainId, fee)
	if err != nil {
		return false, err
	}
	log.Debugln("emitting tendermint event with messageId", messageId.Hex())
	emitMessageToSign(ctx, messageId.Hex())
	return true, nil
}

// DeleteActiveMessageId and set message completed after executed
func (k Keeper) applyMessageExecuted(ctx sdk.Context, applyEvent *cbrtypes.OnChainEvent) (bool, error) {
	evlog := new(ethtypes.Log)
	err := json.Unmarshal(applyEvent.GetElog(), evlog)
	if err != nil {
		return false, err
	}
	msgContracts, err := eth.NewMessageBusFilterer(eth.ZeroAddr, nil)
	if err != nil {
		return false, err
	}
	ev, err := msgContracts.ParseExecuted(*evlog)
	if err != nil {
		log.Errorln("getMessageId: cannot parse event:", err)
		return false, err
	}
	if !k.HasMessage(ctx, ev.Id) {
		log.Infof("skip already applied Executed event (id %x)", ev.Id)
		return false, nil
	}
	log.Debugf("applying MessageBus Executed event: %+v", ev)
	// ev.Id is messageId with transfer: keccak256(abi.encodePacked(MsgType.MessageWithTransfer, dstBridgeAddr, dstTransferId))
	msg, found := k.GetMessage(ctx, ev.Id)
	if !found {
		return false, fmt.Errorf("msg not found for ev.Id: %x", ev.Id)
	}
	k.DeleteActiveMessageId(ctx, msg.GetDstChainId(), eth.Hex2Addr(msg.GetReceiver()), ev.Id)
	status := types.ExecutionStatus(ev.Status)
	if status == types.EXECUTION_STATUS_PENDING {
		return false, fmt.Errorf("error pending status for ev.Id:%s", ev.Id)
	}
	msg.ExecutionStatus = status
	k.SetMessage(ctx, ev.Id, msg)
	return true, nil
}

func buildExecutionContextForMessageNoTransfer(ctx sdk.Context, event *cbrtypes.OnChainEvent) (*types.ExecutionContext, error) {
	chainId := event.Chainid
	evlog := new(ethtypes.Log)
	err := json.Unmarshal(event.GetElog(), evlog)
	if err != nil {
		return nil, err
	}
	msgContracts, err := eth.NewMessageBusFilterer(eth.ZeroAddr, nil)
	if err != nil {
		return nil, err
	}
	ev, err := msgContracts.ParseMessage(*evlog)
	if err != nil {
		return nil, err
	}

	message := types.Message{
		SrcChainId:      chainId,
		Sender:          ev.Sender.String(),
		DstChainId:      ev.DstChainId.Uint64(),
		Receiver:        ev.Receiver.String(),
		Data:            ev.Message,
		Fee:             "0", // ev.Fee.String(),
		ExecutionStatus: types.EXECUTION_STATUS_PENDING,
	}

	execCtx := &types.ExecutionContext{
		Message: message,
	}
	execCtx.MessageId = message.ComputeMessageIdNoTransfer()
	return execCtx, nil
}

func (k Keeper) getTransferType(ctx sdk.Context, bridgeAddr eth.Addr, srcChainId uint64) types.TransferType {
	vault, vaultFound := k.pegbridgeKeeper.GetOriginalTokenVault(ctx, srcChainId)
	pegbr, pegbrFound := k.pegbridgeKeeper.GetPeggedTokenBridge(ctx, srcChainId)
	cbrContractAddr, cbrFound := k.cbridgeKeeper.GetCbrContractAddr(ctx, srcChainId)

	if cbrFound && cbrContractAddr == bridgeAddr {
		return types.TRANSFER_TYPE_LIQUIDITY_SEND
	} else if vaultFound && eth.Hex2Addr(vault.GetAddress()) == bridgeAddr {
		return types.TRANSFER_TYPE_PEG_MINT
	} else if pegbrFound && eth.Hex2Addr(pegbr.GetAddress()) == bridgeAddr {
		return types.TRANSFER_TYPE_PEG_WITHDRAW
	} else {
		return types.TRANSFER_TYPE_NULL
	}
}

func (k Keeper) getTransferInfo(
	c sdk.Context, srcTransferId eth.Hash, transferType types.TransferType, dstChainId uint64) (token []byte, amt string, dstBridgeAddr eth.Addr, err error) {
	ctx := sdk.WrapSDKContext(c)
	switch transferType {
	case types.TRANSFER_TYPE_LIQUIDITY_SEND:
		var relay *cbrtypes.QueryRelayResponse
		relay, err = k.cbridgeKeeper.QueryRelay(ctx, &cbrtypes.QueryRelayRequest{XrefId: srcTransferId.Bytes()})
		relayOnChain := new(cbrtypes.RelayOnChain)
		if err == nil {
			err = relayOnChain.Unmarshal(relay.GetXferRelay().GetRelay())
		}
		if err == nil {
			token = relayOnChain.GetToken()
			amt = new(big.Int).SetBytes(relayOnChain.GetAmount()).String()
		}
		found := false
		dstBridgeAddr, found = k.cbridgeKeeper.GetCbrContractAddr(c, dstChainId)
		if !found {
			err = fmt.Errorf("bridge addr not found for relay. dstChainId %d", dstChainId)
		}
	case types.TRANSFER_TYPE_PEG_MINT:
		var deposit *pegbrtypes.QueryDepositInfoResponse
		deposit, err = k.pegbridgeKeeper.DepositInfo(ctx, &pegbrtypes.QueryDepositInfoRequest{DepositId: srcTransferId.String()})
		if err == nil {
			var mint *pegbrtypes.QueryMintInfoResponse
			mint, err = k.pegbridgeKeeper.MintInfo(ctx, &pegbrtypes.QueryMintInfoRequest{MintId: eth.Bytes2Hash(deposit.DepositInfo.GetMintId()).String()})
			if err == nil {
				mintOnChain := new(pegbrtypes.MintOnChain)
				err = mintOnChain.Unmarshal(mint.MintInfo.GetMintProtoBytes())
				if err != nil {
					log.Errorf("Unmarshal mintInfo.MintProtoBytes err %s", err)
					return
				}
				amt = new(big.Int).SetBytes(mintOnChain.GetAmount()).String()
				token = mintOnChain.GetToken()
				vault, found := k.pegbridgeKeeper.GetPeggedTokenBridge(c, dstChainId)
				if !found {
					err = fmt.Errorf("bridge addr not found for mint. dstChainId %d", dstChainId)
				} else {
					dstBridgeAddr = eth.Hex2Addr(vault.GetAddress())
				}
			}
		}
	case types.TRANSFER_TYPE_PEG_WITHDRAW:
		var burn *pegbrtypes.QueryBurnInfoResponse
		burn, err = k.pegbridgeKeeper.BurnInfo(ctx, &pegbrtypes.QueryBurnInfoRequest{BurnId: srcTransferId.String()})
		if err == nil {
			var withdraw *pegbrtypes.QueryWithdrawInfoResponse
			withdraw, err = k.pegbridgeKeeper.WithdrawInfo(ctx, &pegbrtypes.QueryWithdrawInfoRequest{WithdrawId: eth.Bytes2Hash(burn.BurnInfo.GetWithdrawId()).String()})
			if err == nil {
				withdrawOnChain := new(pegbrtypes.WithdrawOnChain)
				err = withdrawOnChain.Unmarshal(withdraw.WithdrawInfo.GetWithdrawProtoBytes())
				if err != nil {
					log.Errorf("Unmarshal withdrawInfo.WithdrawProtoBytes err %s", err)
					return
				}
				amt = new(big.Int).SetBytes(withdrawOnChain.GetAmount()).String()
				token = withdrawOnChain.GetToken()
				pegBridge, found := k.pegbridgeKeeper.GetOriginalTokenVault(c, dstChainId)
				if !found {
					err = fmt.Errorf("bridge addr not found for withdraw. dstChainId %d", dstChainId)
				} else {
					dstBridgeAddr = eth.Hex2Addr(pegBridge.GetAddress())
				}
			}
		}
	default:
		err = fmt.Errorf("unknown transfer type")
	}
	return
}

func parseMessageWithTransfer(event *cbrtypes.OnChainEvent) (*eth.MessageBusMessageWithTransfer, error) {
	evlog := new(ethtypes.Log)
	err := json.Unmarshal(event.GetElog(), evlog)
	if err != nil {
		return nil, err
	}
	msgContracts, err := eth.NewMessageBusFilterer(eth.ZeroAddr, nil)
	if err != nil {
		return nil, err
	}
	msgXfer, err := msgContracts.ParseMessageWithTransfer(*evlog)
	return msgXfer, err
}

func emitMessageToSign(ctx sdk.Context, id string) {
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeMessageToSign,
		sdk.NewAttribute(types.AttributeKeyMessageId, id),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
	))
}

func shouldRefund(s cbrtypes.XferStatus) (shouldRefund bool) {
	return s == cbrtypes.XferStatus_BAD_LIQUIDITY ||
		s == cbrtypes.XferStatus_BAD_SLIPPAGE ||
		s == cbrtypes.XferStatus_BAD_XFER_DISABLED ||
		s == cbrtypes.XferStatus_BAD_DEST_CHAIN ||
		s == cbrtypes.XferStatus_EXCEED_MAX_OUT_AMOUNT
}
