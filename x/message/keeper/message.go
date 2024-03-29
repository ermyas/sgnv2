package keeper

import (
	"fmt"
	"math/big"

	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/message/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetMessage(ctx sdk.Context, messageId eth.Hash, message *types.Message) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetMessageKey(messageId), k.cdc.MustMarshal(message))
}

func (k Keeper) GetMessage(ctx sdk.Context, messageId eth.Hash) (message types.Message, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetMessageKey(messageId))
	if bz == nil {
		return message, false
	}
	k.cdc.MustUnmarshal(bz, &message)
	return message, true
}

func (k Keeper) HasMessage(ctx sdk.Context, messageId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetMessageKey(messageId))
}

func (k Keeper) DeleteMessage(ctx sdk.Context, messageId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetMessageKey(messageId))
}

func (k Keeper) IterateAllMessages(
	ctx sdk.Context, handler func(message types.Message) (stop bool),
) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.MessagePrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var message types.Message
		k.cdc.MustUnmarshal(iter.Value(), &message)
		if handler(message) {
			break
		}
	}
}

func (k Keeper) getMessageTransferInfo(ctx sdk.Context, message *types.Message) (*types.Transfer, error) {
	var transfer *types.Transfer
	srcTransferId := eth.Bytes2Hash(message.GetTransferRefId())
	switch message.GetTransferType() {
	case types.TRANSFER_TYPE_LIQUIDITY_RELAY:
		relay, found := k.cbridgeKeeper.GetXferRelay(ctx, srcTransferId)
		if !found {
			return nil, fmt.Errorf("relay not found for src transfer %x", srcTransferId)
		}
		relayOnChain := new(cbrtypes.RelayOnChain)
		err := relayOnChain.Unmarshal(relay.GetRelay())
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal relay %x", relay.GetRelay())
		}
		transfer = &types.Transfer{
			Token:  relayOnChain.GetToken(),
			Amount: new(big.Int).SetBytes(relayOnChain.GetAmount()).String(),
		}

	case types.TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		wdOnchain := k.cbridgeKeeper.QueryXferRefund(ctx, srcTransferId)
		if wdOnchain == nil {
			return nil, fmt.Errorf("refund withdrawal not found for src transfer %x", srcTransferId)
		}
		transfer = &types.Transfer{
			Token:    wdOnchain.GetToken(),
			Amount:   new(big.Int).SetBytes(wdOnchain.GetAmount()).String(),
			WdSeqNum: wdOnchain.GetSeqnum(),
		}

	case types.TRANSFER_TYPE_PEG_MINT, types.TRANSFER_TYPE_PEG_V2_MINT:
		var mintId eth.Hash
		if deposit, found := k.pegbridgeKeeper.GetDepositInfo(ctx, srcTransferId); found {
			mintId = eth.Bytes2Hash(deposit.GetMintId())
		} else if burn, found := k.pegbridgeKeeper.GetBurnInfo(ctx, srcTransferId); found && len(burn.GetMintId()) > 0 {
			mintId = eth.Bytes2Hash(burn.GetMintId())
		} else {
			// check if refund for burn
			mintOnChain, found := k.pegbridgeKeeper.GetBurnRefund(ctx, srcTransferId)
			if found {
				transfer = &types.Transfer{
					Token:  mintOnChain.GetToken(),
					Amount: new(big.Int).SetBytes(mintOnChain.GetAmount()).String(),
				}
				break
			}
			mintId, found = k.pegbridgeKeeper.GetRefundClaimInfo(ctx, srcTransferId)
			if !found {
				return nil, fmt.Errorf("deposit or burn-refund not found for src transfer %x", srcTransferId)
			}
		}
		mint, found := k.pegbridgeKeeper.GetMintInfo(ctx, mintId)
		if !found {
			return nil, fmt.Errorf("mint not found for src transfer %x, mintId %x", srcTransferId, mintId)
		}
		mintOnChain := new(pegbrtypes.MintOnChain)
		err := mintOnChain.Unmarshal(mint.GetMintProtoBytes())
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal mintOnchain %x", mint.GetMintProtoBytes())
		}
		transfer = &types.Transfer{
			Token:  mintOnChain.GetToken(),
			Amount: new(big.Int).SetBytes(mintOnChain.GetAmount()).String(),
		}

	case types.TRANSFER_TYPE_PEG_WITHDRAW, types.TRANSFER_TYPE_PEG_V2_WITHDRAW:
		var withdrawId eth.Hash
		burn, found := k.pegbridgeKeeper.GetBurnInfo(ctx, srcTransferId)
		if !found {
			// check if refund
			wdOnChain, found := k.pegbridgeKeeper.GetDepositRefund(ctx, srcTransferId)
			if found {
				transfer = &types.Transfer{
					Token:  wdOnChain.GetToken(),
					Amount: new(big.Int).SetBytes(wdOnChain.GetAmount()).String(),
				}
				break
			}
			withdrawId, found = k.pegbridgeKeeper.GetRefundClaimInfo(ctx, srcTransferId)
			if !found {
				return nil, fmt.Errorf("burn or deposit refund not found for src transfer %x", srcTransferId)
			}
		} else {
			withdrawId = eth.Bytes2Hash(burn.GetWithdrawId())
		}
		withdraw, found := k.pegbridgeKeeper.GetWithdrawInfo(ctx, withdrawId)
		if !found {
			return nil, fmt.Errorf("withdraw not found for src transfer %x", srcTransferId)
		}
		withdrawOnChain := new(pegbrtypes.WithdrawOnChain)
		err := withdrawOnChain.Unmarshal(withdraw.GetWithdrawProtoBytes())
		if err != nil {
			return nil, fmt.Errorf("unable to unmarshal withdrawOnChain %x", withdraw.GetWithdrawProtoBytes())
		}
		transfer = &types.Transfer{
			Token:  withdrawOnChain.GetToken(),
			Amount: new(big.Int).SetBytes(withdrawOnChain.GetAmount()).String(),
		}
	}
	return transfer, nil
}

func (k Keeper) SetSrcTransferMessageId(ctx sdk.Context, srcBridgeType types.BridgeType, srcTransferId eth.Hash, messageId eth.Hash) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetSrcTransferKey(srcBridgeType, srcTransferId), messageId.Bytes())
}

func (k Keeper) GetSrcTransferMessageId(
	ctx sdk.Context, srcBridgeType types.BridgeType, srcTransferId eth.Hash) (messageId eth.Hash, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetSrcTransferKey(srcBridgeType, srcTransferId))
	if bz == nil {
		return messageId, false
	}
	return eth.Bytes2Hash(bz), true
}

func (k Keeper) HasSrcTransferMessageId(ctx sdk.Context, srcBridgeType types.BridgeType, srcTransferId eth.Hash) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetSrcTransferKey(srcBridgeType, srcTransferId))
}
