package keeper

import (
	"context"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/message/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) ExecutionContexts(
	c context.Context, req *types.QueryExecutionContextsRequest) (*types.QueryExecutionContextsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	// 1. Process filters
	var messageIds []eth.Hash
	// No filters, return all
	if len(req.ContractInfos) == 0 {
		k.IterateAllActiveMessageIds(ctx, func(messageId eth.Hash) bool {
			messageIds = append(messageIds, messageId)
			return false
		})
	} else {
		// Go through filters
		var dstChainIds []uint64
		dstContracts := make(map[uint64]map[eth.Addr]bool)
		for _, info := range req.ContractInfos {
			currChainContracts, exists := dstContracts[info.ChainId]
			if info.Address != "" {
				if !exists {
					currChainContracts = make(map[eth.Addr]bool)
				}
				currChainContracts[eth.Hex2Addr(info.Address)] = true
				dstContracts[info.ChainId] = currChainContracts
			} else {
				if exists {
					// Get all active messages for the chainID instead
					delete(dstContracts, info.ChainId)
				}
				dstChainIds = append(dstChainIds, info.ChainId)
			}
		}
		for _, dstChainId := range dstChainIds {
			currChainMessageIds, found := k.GetActiveMessageIdsByDstChainId(ctx, dstChainId)
			if found {
				messageIds = append(messageIds, currChainMessageIds...)
			}
		}
		for chainId, contracts := range dstContracts {
			for address := range contracts {
				currInfoMessageIds, found := k.GetActiveMessageIdsByChainIdTarget(ctx, chainId, address)
				if found {
					messageIds = append(messageIds, currInfoMessageIds...)
				}
			}
		}
	}

	// 2. Get Messages
	messages := make(map[eth.Hash]types.Message)
	for _, id := range messageIds {
		message, found := k.GetMessage(ctx, id)
		if !found {
			return nil, status.Error(codes.Internal, "message not found")
		}
		// check message sig quorum
		curss, found := k.cbridgeKeeper.GetChainSigners(ctx, message.DstChainId)
		if !found {
			log.Errorf("cannot find current signers for chain %d", message.DstChainId)
			continue
		}
		// only return messages with enough sigs
		pass, _ := cbrtypes.ValidateSignatureQuorum(message.Signatures, curss.GetSortedSigners())
		if pass {
			messages[id] = message
		}
	}

	// 3. Populate ExecutionContexts
	// 3.1. Get powers
	chainIds := make(map[uint64]bool)
	for _, message := range messages {
		chainIds[message.DstChainId] = true
	}
	chainPowers := make(map[uint64][]string)
	for chainId := range chainIds {
		chainSigners, found := k.cbridgeKeeper.GetChainSigners(ctx, chainId)
		if !found {
			return nil, status.Error(codes.Internal, "powers not found")
		}
		var currChainPowers []string
		for _, signer := range chainSigners.GetSortedSigners() {
			currChainPowers = append(currChainPowers, new(big.Int).SetBytes(signer.GetPower()).String())
		}
		chainPowers[chainId] = currChainPowers
	}

	// 3.2. Add Transfer if applicable
	var execCtxs []types.ExecutionContext
	for id, message := range messages {
		execCtx := types.ExecutionContext{
			MessageId: id.Bytes(),
			Message:   message,
			Powers:    chainPowers[message.DstChainId],
		}
		if message.TransferType != types.TRANSFER_TYPE_NULL {
			transfer, found := k.GetTransfer(ctx, eth.Bytes2Hash(execCtx.MessageId))
			if !found {
				return nil, status.Error(codes.Internal, "transfer not found")
			}
			execCtx.Transfer = &transfer
		}
		execCtxs = append(execCtxs, execCtx)
	}
	return &types.QueryExecutionContextsResponse{ExecutionContexts: execCtxs}, nil
}

func (k Keeper) MessageExists(c context.Context, request *types.QueryMessageExistsRequest) (*types.QueryMessageExistsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryMessageExistsResponse{Exists: k.HasMessage(ctx, eth.Hex2Hash(request.GetMessageId()))}, nil
}

func (k Keeper) IsMessageActive(c context.Context, request *types.IsMessageActiveRequest) (*types.IsMessageActiveResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	message, found := k.GetMessage(ctx, eth.Hex2Hash(request.GetMessageId()))
	if !found {
		return &types.IsMessageActiveResponse{Exists: false}, nil
	}
	return &types.IsMessageActiveResponse{Exists: k.HasActiveMessageId(ctx, message.GetDstChainId(), eth.Hex2Addr(message.GetReceiver()), eth.Hex2Hash(request.GetMessageId()))}, nil
}
func (k Keeper) RefundExists(c context.Context, request *types.QueryRefundExistsRequest) (*types.QueryRefundExistsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return &types.QueryRefundExistsResponse{Exists: k.HasRefund(ctx, eth.Hex2Hash(request.GetSrcTransferId()))}, nil
}

func (k Keeper) Message(c context.Context, req *types.QueryMessageRequest) (*types.QueryMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	messageId := eth.Hex2Hash(req.MessageId)
	msg, found := k.GetMessage(ctx, messageId)
	if !found {
		return nil, types.WrapErrNoMessageFound(messageId)
	}
	return &types.QueryMessageResponse{Message: msg}, nil
}

func (k Keeper) Transfer(c context.Context, req *types.QueryTransferRequest) (*types.QueryTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	messageId := eth.Hex2Hash(req.MessageId)
	xfer, found := k.GetTransfer(ctx, messageId)
	if !found {
		return nil, types.WrapErrNoTransferFound(messageId)
	}
	return &types.QueryTransferResponse{Transfer: xfer}, nil
}

func (k Keeper) MessageBus(c context.Context, req *types.QueryMessageBusRequest) (*types.QueryMessageBusResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	chainId := req.ChainId
	msgBus, found := k.GetMessageBus(ctx, chainId)
	if !found {
		return nil, types.WrapErrNoMessageBusFound(chainId)
	}
	return &types.QueryMessageBusResponse{MessageBus: msgBus}, nil
}

func (k Keeper) FeeClaimInfo(c context.Context, req *types.QueryFeeClaimInfoRequest) (*types.QueryFeeClaimInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	info, found := k.GetFeeClaimInfo(ctx, eth.Hex2Addr(req.Address))
	if !found {
		return nil, types.WrapErrNoClaimInfoFound(req.Address)
	}
	return &types.QueryFeeClaimInfoResponse{FeeClaimInfo: info}, nil
}
