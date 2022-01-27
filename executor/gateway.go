package executor

import (
	"context"
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor/types"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
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
	req := &types.InitWithdrawRequest{WithdrawReq: wdReqBytes}
	res, err := g.cli.InitWithdraw(context.Background(), req)
	if err != nil {
		return err
	}
	if res.Err.Code != types.ErrCode_ERROR_CODE_UNDEFINED {
		return fmt.Errorf("initWithdraw err: %s", res.Err.Msg)
	}
	return nil
}
