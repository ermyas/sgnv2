package utils

import (
	"context"
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"google.golang.org/grpc"
	"math"
	"math/big"
)

// Implementation of account-svc API layer
type GatewayClient struct {
	conn   *grpc.ClientConn
	client webapi.WebClient
}

func NewGatewayAPI(gatewayUrl string) (*GatewayClient, error) {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(gatewayUrl, opts...)
	if err != nil {
		return nil, err
	}
	gateway := &GatewayClient{
		conn:   conn,
		client: webapi.NewWebClient(conn),
	}
	return gateway, nil
}

func (g *GatewayClient) GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error) {
	resp, err := g.client.GetTransferConfigs(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to get transfer status:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) GetTotalLiquidityProviderTokenBalance(ctx context.Context, request *webapi.GetTotalLiquidityProviderTokenBalanceRequest) (*webapi.GetTotalLiquidityProviderTokenBalanceResponse, error) {
	resp, err := g.client.GetTotalLiquidityProviderTokenBalance(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Err != nil {
		return nil, fmt.Errorf("fail to GetTotalLiquidityProviderTokenBalance:%s", resp.Err.Msg)
	}
	return resp, nil
}

func (g *GatewayClient) GetUpdatedLocalChainTokenInfoMap(context context.Context) (map[uint32]*webapi.ChainTokenInfo, error) {
	resp, err := g.GetTransferConfigs(context, &webapi.GetTransferConfigsRequest{})
	err = CheckErr(resp.GetErr(), err)
	if err != nil {
		return nil, err
	}
	return resp.GetChainToken(), nil
}

func (g *GatewayClient) GetOverallBalanceInfo(ctx context.Context) (map[string]float64, error) {
	getTransferConfigsResp, err := g.GetTransferConfigs(ctx, &webapi.GetTransferConfigsRequest{})
	if err != nil {
		return nil, err
	}
	chainTokenMap := getTransferConfigsResp.GetChainToken()
	balanceMap := make(map[string]float64)
	if len(chainTokenMap) > 0 {
		for chainId, chainTokenInfo := range chainTokenMap {
			if chainTokenInfo != nil {
				for _, tokenInfo := range chainTokenInfo.GetToken() {
					addedBalance, err := g.GetTokenFloat64Balance(ctx, tokenInfo.GetToken().GetSymbol(), uint32(tokenInfo.GetToken().GetDecimal()), chainId)
					if err != nil {
						log.Infof("fail to get GetTokenFloat64Balance, symbol:%s, chain:%d, err:%s", tokenInfo.GetToken().GetSymbol(), chainId, err.Error())
						return nil, err
					}
					mapKey := tokenInfo.GetToken().GetSymbol()
					log.Infof("chainId:%d, symbol:%s, amt:%f", chainId, tokenInfo.GetToken().GetSymbol(), addedBalance)
					balanceMap[mapKey] += addedBalance
				}
			}
		}
	}
	return balanceMap, nil
}

func (g *GatewayClient) GetTokenFloat64Balance(ctx context.Context, tokenSymbol string, tokenDecimal uint32, chainId uint32) (float64, error) {
	var chainIds = []uint32{chainId}
	resp, err := g.GetTotalLiquidityProviderTokenBalance(ctx,
		&webapi.GetTotalLiquidityProviderTokenBalanceRequest{
			ChainIds:    chainIds,
			TokenSymbol: tokenSymbol,
		})
	err = CheckErr(resp.GetErr(), err)
	if err != nil {
		log.Infof("fail to GetTotalLiquidityProviderTokenBalance, chainIds:%+v, symbol:%s, err:%s", chainIds, tokenSymbol, err.Error())
		return 0, err
	}
	balanceStringForToken := resp.TotalLiq[uint64(chainId)]
	balanceBigIntForToken, _ := new(big.Int).SetString(balanceStringForToken, 10)
	return GetTokenFloat64AmtFromWeiAmtAndDecimal(balanceBigIntForToken, tokenDecimal), nil
}

func GetTokenFloat64AmtFromWeiAmtAndDecimal(tokenAmt *big.Int, decimal uint32) float64 {
	if tokenAmt != nil {
		float64TokenAmount, _ := new(big.Float).Quo(new(big.Float).SetInt(tokenAmt), big.NewFloat(math.Pow(10, float64(decimal)))).Float64()
		return float64TokenAmount
	} else {
		return 0
	}
}

func CheckErr(errMsg *webapi.ErrMsg, rpcErr error) error {
	if rpcErr != nil {
		return rpcErr
	}
	if errMsg != nil {
		return fmt.Errorf("error code:%d, msg:%s", errMsg.GetCode(), errMsg.GetMsg())
	}
	return nil
}

func (g *GatewayClient) Close() {
	if err := g.conn.Close(); err != nil {
		log.Warnln("closeGatewayConn: error:", err)
	}
}
