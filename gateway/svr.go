package webapi

import (
	"context"
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/sgn-v2/gateway/webapi"
)

type DAL struct {
	*sqldb.Db
}

func NewDAL(driver, info string, poolSize int) (*DAL, error) {
	db, err := sqldb.NewDb(driver, info, poolSize)
	if err != nil {
		log.Errorf("fail with db init:%s, %s, %d, err:%+v", driver, info, poolSize, err)
		return nil, err
	}

	dal := &DAL{
		db,
	}
	return dal, nil
}

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

func (gs *GatewayService) GetLPInfoList(ctx context.Context, request *webapi.GetLPInfoListRequest) (*webapi.GetLPInfoListResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) GetTotalLiquidity(ctx context.Context, request *webapi.GetTotalLiquidityRequest) (*webapi.GetTotalLiquidityResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) WithdrawLiquidity(ctx context.Context, request *webapi.WithdrawLiquidityRequest) (*webapi.WithdrawLiquidityResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) TransferHistory(ctx context.Context, request *webapi.TransferHistoryRequest) (*webapi.TransferHistoryResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) LPHistory(ctx context.Context, request *webapi.LPHistoryRequest) (*webapi.LPHistoryResponse, error) {
	panic("implement me")
}

func NewGatewayService() (*GatewayService, error) {
	// Make a private config copy.
	dal, err := NewDAL("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", "dburl"), 10)
	if err != nil {
		return nil, err
	}

	return &GatewayService{
		dal: dal,
	}, nil
}
