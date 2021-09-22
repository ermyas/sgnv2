package webapi

import (
	"context"
	"fmt"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

// Close the database DAL.
func (gs *GatewayService) Close() {
	if gs.dal == nil {
		return
	}
	gs.dal.Close()
	gs.dal = nil
}

type GatewayConfig struct {
}

type GatewayService struct {
	dal *DAL
}

func (gs *GatewayService) SetAdvancedInfo(ctx context.Context, request *webapi.SetAdvancedInfoRequest) (*webapi.SetAdvancedInfoResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) EstimateAmt(ctx context.Context, request *webapi.EstimateAmtRequest) (*webapi.EstimateAmtResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) MarkTransferBegin(ctx context.Context, request *webapi.MarkTransferBeginRequest) (*webapi.MarkTransferBeginResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) GetLPInfoList(ctx context.Context, request *webapi.GetLPInfoListRequest) (*webapi.GetLPInfoListResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) MarkLiquidityAdd(ctx context.Context, request *webapi.MarkLiquidityAddRequest) (*webapi.MarkLiquidityAddResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) WithdrawLiquidity(ctx context.Context, request *webapi.WithdrawLiquidityRequest) (*webapi.WithdrawLiquidityResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) QueryLiquidityStatus(ctx context.Context, request *webapi.QueryLiquidityStatusRequest) (*types.QueryLiquidityStatusResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) TransferHistory(ctx context.Context, request *webapi.TransferHistoryRequest) (*webapi.TransferHistoryResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) LPHistory(ctx context.Context, request *webapi.LPHistoryRequest) (*webapi.LPHistoryResponse, error) {
	panic("implement me")
}

func NewGatewayService(dbUrl string) (*GatewayService, error) {
	// Make a private config copy.
	dal, err := NewDAL("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", dbUrl), 10)
	if err != nil {
		return nil, err
	}

	return &GatewayService{
		dal: dal,
	}, nil
}
