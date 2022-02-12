package sgn

import (
	"context"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor/types"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
)

func (c *SgnClient) GetExecutionContexts(filters []*commontypes.ContractInfo) ([]msgtypes.ExecutionContext, error) {
	qc := msgtypes.NewQueryClient(c.txrs.GetTransactor().CliCtx)
	req := &msgtypes.QueryExecutionContextsRequest{
		ContractInfos: filters,
	}
	res, err := qc.ExecutionContexts(context.Background(), req)
	if err != nil {
		log.Errorln("failed to query messages from sgn", err)
		return nil, err
	}
	return res.GetExecutionContexts(), nil
}

func (c *SgnClient) GetXferWithdrawStatus(
	addr string, nonce uint64, chainId uint64) (*cbrtypes.WithdrawDetail, cbrtypes.WithdrawStatus, error) {

	req := &cbrtypes.QueryWithdrawLiquidityStatusRequest{
		SeqNum:  nonce,
		UsrAddr: addr,
	}
	qc := cbrtypes.NewQueryClient(c.grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), types.GatewayTimeout)
	defer cancel()
	res, err := qc.QueryWithdrawLiquidityStatus(ctx, req)
	if err != nil {
		return nil, cbrtypes.WithdrawStatus_WD_UNKNOWN, err
	}
	log.Debugf("withdraw status %v", res.GetStatus())

	return res.GetDetail(), res.GetStatus(), nil
}

func (c *SgnClient) GetPegRefundClaimId(depositId []byte) (string, error) {
	qc := pegbrtypes.NewQueryClient(c.grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), types.GatewayTimeout)
	defer cancel()
	req := &pegbrtypes.QueryRefundClaimInfoRequest{DepositId: eth.Bytes2Hex(depositId)}
	res, err := qc.RefundClaimInfo(ctx, req)
	if err != nil {
		return "", err
	}
	return res.WithdrawId, nil
}

func (c *SgnClient) GetPegWithdrawInfo(withdrawId string) (*pegbrtypes.WithdrawInfo, error) {
	qc := pegbrtypes.NewQueryClient(c.grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), types.GatewayTimeout)
	defer cancel()
	req := &pegbrtypes.QueryWithdrawInfoRequest{WithdrawId: withdrawId}
	res, err := qc.WithdrawInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return &res.WithdrawInfo, nil
}

func (c *SgnClient) GetPegMintInfo(burnId string) (*pegbrtypes.MintInfo, error) {
	qc := pegbrtypes.NewQueryClient(c.grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), types.GatewayTimeout)
	defer cancel()
	req := &pegbrtypes.QueryMintInfoRequest{MintId: burnId}
	res, err := qc.MintInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return &res.MintInfo, nil
}

func (c *SgnClient) GetChainSigners(chainId uint64) (*cbrtypes.ChainSigners, error) {
	qc := cbrtypes.NewQueryClient(c.grpcConn)
	ctx, cancel := context.WithTimeout(context.Background(), types.GatewayTimeout)
	defer cancel()
	req := &cbrtypes.QueryChainSignersRequest{ChainId: chainId}
	res, err := qc.QueryChainSigners(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.GetChainSigners(), nil
}