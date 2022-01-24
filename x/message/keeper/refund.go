package keeper

import (
	"encoding/binary"

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
