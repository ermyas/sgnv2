package keeper

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetRefund(ctx sdk.Context, srcXferId eth.Hash, execCtx *types.ExecutionContext) {
	ctx.KVStore(k.storeKey).Set(types.GetMessageRefundKey(srcXferId), execCtx.MustMarshal())
}

func (k Keeper) HasRefund(ctx sdk.Context, srcXferId eth.Hash) bool {
	return ctx.KVStore(k.storeKey).Has(types.GetMessageRefundKey(srcXferId))
}

func (k Keeper) applyTransferRefund(ctx sdk.Context, ev *eth.MessageBusMessageWithTransfer) (bool, error) {
	if k.HasRefund(ctx, ev.SrcTransferId) {
		log.Infof("skip already applied message (srcXferId %x) with transfer refund", ev.SrcTransferId)
		return false, nil
	}
	log.Debugf("applying transfer refund for sender app\n%s", ev.PrettyLog(0))
	nonce := k.incrRefundNonce(ctx)
	wdOnchain := k.cbridgeKeeper.QueryXferRefund(ctx, ev.SrcTransferId)
	if wdOnchain == nil {
		return false, fmt.Errorf("wdOnchain not found for srcXferId %x", ev.SrcTransferId)
	}
	log.Debugf("found WdOnchain: %+v", wdOnchain)
	transfer := &types.Transfer{
		Amount: new(big.Int).SetBytes(wdOnchain.Amount).String(),
		Token:  wdOnchain.Token,
		SeqNum: nonce,
		RefId:  ev.SrcTransferId[:],
	}
	message := &types.Message{
		SrcChainId:      wdOnchain.Chainid,
		DstChainId:      wdOnchain.Chainid,
		Sender:          ev.Sender.String(),
		Receiver:        ev.Sender.String(),
		Data:            ev.Message,
		TransferType:    types.TRANSFER_TYPE_LIQUIDITY_WITHDRAW,
		ExecutionStatus: types.EXECUTION_STATUS_PENDING,
	}
	execCtx := &types.ExecutionContext{
		Message:  *message,
		Transfer: transfer,
	}
	bridge, found := k.cbridgeKeeper.GetCbrContractAddr(ctx, wdOnchain.Chainid)
	if !found {
		return false, fmt.Errorf("bridge addr not found for chainId %d", wdOnchain.Chainid)
	}
	messageId := execCtx.PopulateMessageId(bridge)
	k.SetActiveMessageId(ctx, message.DstChainId, eth.Hex2Addr(message.Receiver), messageId)
	k.SetMessage(ctx, messageId, *message)
	k.SetTransfer(ctx, messageId, *transfer)
	k.SetRefund(ctx, eth.Bytes2Hash(ev.SrcTransferId[:]), execCtx)
	emitMessageToSign(ctx, messageId.Hex())
	return true, nil
}

func (k Keeper) incrRefundNonce(ctx sdk.Context) uint64 {
	kv := ctx.KVStore(k.storeKey)
	key := types.GetRefundNonceKey()
	oldNonceBytes := kv.Get(key)
	var oldNonce uint64
	if oldNonceBytes == nil {
		oldNonce = 0
	} else {
		oldNonce = uint64(binary.LittleEndian.Uint64(oldNonceBytes))
	}
	newNonce := oldNonce + 1
	newNonceBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(newNonceBytes, newNonce)
	kv.Set(key, newNonceBytes)
	log.Debugf("incremented msg refund nonce: old %d, new %d", oldNonce, newNonce)
	return newNonce
}
