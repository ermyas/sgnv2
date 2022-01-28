package executor

import (
	"context"
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor/types"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	"google.golang.org/grpc"
)

type GatewayClient struct {
	conn *grpc.ClientConn
	cli  types.WebClient
}

func NewGatewayClient(gatewayUrl string) *GatewayClient {
	log.Infof("Dialing gateway grpc: %s", gatewayUrl)
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	context, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	conn, err := grpc.DialContext(context, gatewayUrl, opts...)
	defer cancel()
	if err != nil {
		log.Fatalln("failed to initialize gateway grpc connection", err)
	}
	return &GatewayClient{
		conn: conn,
		cli:  types.NewWebClient(conn),
	}
}

func (g *GatewayClient) GetExecutionContexts(filters []*commontypes.ContractInfo) ([]msgtypes.ExecutionContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), GatewayTimeout)
	defer cancel()
	req := &msgtypes.QueryExecutionContextsRequest{
		ContractInfos: filters,
	}
	res, err := g.cli.ExecutionContexts(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.ExecutionContexts, nil
}

func (g *GatewayClient) InitWithdraw(srcXferId []byte, nonce uint64) error {
	wdReq := &cbrtypes.WithdrawReq{
		XferId:       eth.Bytes2Hex(srcXferId),
		ReqId:        nonce,
		WithdrawType: cbrtypes.RefundTransfer,
	}
	wdReqBytes, err := wdReq.Marshal()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), GatewayTimeout)
	defer cancel()
	req := &types.InitWithdrawRequest{WithdrawReq: wdReqBytes}
	res, err := g.cli.InitWithdraw(ctx, req)
	if err != nil {
		return err
	}
	if res.Err.GetCode() != types.ErrCode_ERROR_CODE_UNDEFINED {
		return fmt.Errorf("initWithdraw err: %s", res.Err.Msg)
	}
	return nil
}

func (g *GatewayClient) InitPegRefund(refId []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), GatewayTimeout)
	defer cancel()
	req := &types.InitPegRefundRequest{RefId: refId}
	res, err := g.cli.InitPegRefund(ctx, req)
	if err != nil {
		return err
	}
	if res.Err.GetCode() != types.ErrCode_ERROR_CODE_UNDEFINED {
		return fmt.Errorf("InitPegRefund err: %s", res.Err.Msg)
	}
	return nil
}
