package ropstentest

import (
	"context"
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"google.golang.org/grpc"
)

type GatewayAPI interface {
	Close()
	MarkTransfer(ctx context.Context, request *webapi.MarkTransferRequest) (*webapi.MarkTransferResponse, error)
	GetTransferStatus(ctx context.Context, request *webapi.GetTransferStatusRequest) (*webapi.GetTransferStatusResponse, error)
	EstimateAmt(ctx context.Context, request *webapi.EstimateAmtRequest) (*webapi.EstimateAmtResponse, error)
	WithdrawLiquidity(ctx context.Context, request *webapi.WithdrawLiquidityRequest) (*webapi.WithdrawLiquidityResponse, error)
	SetAdvancedInfo(ctx context.Context, request *webapi.SetAdvancedInfoRequest) (*webapi.SetAdvancedInfoResponse, error)
	GetAdvancedInfo(ctx context.Context, request *webapi.GetAdvancedInfoRequest) (*webapi.GetAdvancedInfoResponse, error)
	GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error)
	GetLPInfoList(ctx context.Context, request *webapi.GetLPInfoListRequest) (*webapi.GetLPInfoListResponse, error)
	MarkLiquidity(ctx context.Context, request *webapi.MarkLiquidityRequest) (*webapi.MarkLiquidityResponse, error)
	QueryLiquidityStatus(ctx context.Context, request *webapi.QueryLiquidityStatusRequest) (*webapi.QueryLiquidityStatusResponse, error)
	ClaimWithdrawReward(ctx context.Context, request *webapi.ClaimWithdrawRewardRequest) (*webapi.ClaimWithdrawRewardResponse, error)
	ClaimRewardDetails(ctx context.Context, request *webapi.ClaimRewardDetailsRequest) (*webapi.ClaimRewardDetailsResponse, error)
	TransferHistory(ctx context.Context, request *webapi.TransferHistoryRequest) (*webapi.TransferHistoryResponse, error)
	LPHistory(ctx context.Context, request *webapi.LPHistoryRequest) (*webapi.LPHistoryResponse, error)
}

// Implementation of account-svc API layer
type GatewayClient struct {
	gatewayConn *grpc.ClientConn
}

func NewGatewayAPI(gatewayUrl string) (*GatewayClient, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(gatewayUrl, opts...)
	if err != nil {
		return nil, err
	}
	gateway := &GatewayClient{
		gatewayConn: conn,
	}
	return gateway, nil
}

func (g *GatewayClient) MarkTransfer(ctx context.Context, request *webapi.MarkTransferRequest) (*webapi.MarkTransferResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).MarkTransfer(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to mark transfer:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) GetTransferStatus(ctx context.Context, request *webapi.GetTransferStatusRequest) (*webapi.GetTransferStatusResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).GetTransferStatus(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) EstimateAmt(ctx context.Context, request *webapi.EstimateAmtRequest) (*webapi.EstimateAmtResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).EstimateAmt(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) WithdrawLiquidity(ctx context.Context, request *webapi.WithdrawLiquidityRequest) (*webapi.WithdrawLiquidityResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).WithdrawLiquidity(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) SetAdvancedInfo(ctx context.Context, request *webapi.SetAdvancedInfoRequest) (*webapi.SetAdvancedInfoResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).SetAdvancedInfo(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) GetAdvancedInfo(ctx context.Context, request *webapi.GetAdvancedInfoRequest) (*webapi.GetAdvancedInfoResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).GetAdvancedInfo(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).GetTransferConfigs(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) GetLPInfoList(ctx context.Context, request *webapi.GetLPInfoListRequest) (*webapi.GetLPInfoListResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).GetLPInfoList(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) MarkLiquidity(ctx context.Context, request *webapi.MarkLiquidityRequest) (*webapi.MarkLiquidityResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).MarkLiquidity(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) QueryLiquidityStatus(ctx context.Context, request *webapi.QueryLiquidityStatusRequest) (*webapi.QueryLiquidityStatusResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).QueryLiquidityStatus(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) ClaimWithdrawReward(ctx context.Context, request *webapi.ClaimWithdrawRewardRequest) (*webapi.ClaimWithdrawRewardResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).ClaimWithdrawReward(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) ClaimRewardDetails(ctx context.Context, request *webapi.ClaimRewardDetailsRequest) (*webapi.ClaimRewardDetailsResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).ClaimRewardDetails(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) TransferHistory(ctx context.Context, request *webapi.TransferHistoryRequest) (*webapi.TransferHistoryResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).TransferHistory(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) LPHistory(ctx context.Context, request *webapi.LPHistoryRequest) (*webapi.LPHistoryResponse, error) {
	resp, err := webapi.NewWebClient(g.gatewayConn).LPHistory(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) Close() {
	if err := g.gatewayConn.Close(); err != nil {
		log.Warnln("closeGatewayConn: error:", err)
	}
}
