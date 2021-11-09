package gatewaysvc

import (
	"context"

	"github.com/celer-network/sgn-v2/gateway/webapi"
)

func (gs *GatewayService) GetInfoByTxHash(ctx context.Context, request *webapi.GetInfoByTxHashRequest) (*webapi.GetInfoByTxHashResponse, error) {
	return &webapi.GetInfoByTxHashResponse{
		Operation: webapi.CSOperation_CA_UNKNOWN,
		Memo:      "unimplemented",
	}, nil
}

func (gs *GatewayService) FixEventMiss(ctx context.Context, request *webapi.FixEventMissRequest) (*webapi.FixEventMissResponse, error) {
	return &webapi.FixEventMissResponse{}, nil
}
