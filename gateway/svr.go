package gateway

import (
	"context"
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/lthibault/jitterbug"
	"os"
	"time"
)

// Close the database DAL.
func (gs *GatewayService) Close() {
	if dal.DB == nil {
		return
	}
	dal.DB.Close()
	dal.DB = nil
}

type GatewayConfig struct {
}

type GatewayService struct {
	f *fee.TokenPriceCache
}

func (gs *GatewayService) SetAdvancedInfo(ctx context.Context, request *webapi.SetAdvancedInfoRequest) (*webapi.SetAdvancedInfoResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) GetTransferConfigs(ctx context.Context, request *webapi.GetTransferConfigsRequest) (*webapi.GetTransferConfigsResponse, error) {
	chainTokenList, err := dal.DB.GetChainTokenList()
	if err != nil {
		return &webapi.GetTransferConfigsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get chain_token failed",
			},
			Chains:     nil,
			ChainToken: nil,
		}, nil
	}
	var chainIds []uint64
	for key := range chainTokenList {
		chainIds = append(chainIds, key)
	}
	chains, err := dal.DB.GetChainInfo(chainIds)
	if err != nil {
		return &webapi.GetTransferConfigsResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "get chain info failed",
			},
			Chains:     chains,
			ChainToken: chainTokenList,
		}, nil
	}
	return &webapi.GetTransferConfigsResponse{
		Err:        nil,
		Chains:     chains,
		ChainToken: chainTokenList,
	}, nil
}

func (gs *GatewayService) EstimateAmt(ctx context.Context, request *webapi.EstimateAmtRequest) (*webapi.EstimateAmtResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) MarkTransfer(ctx context.Context, request *webapi.MarkTransferRequest) (*webapi.MarkTransferResponse, error) {
	transferId := request.GetTransferId()
	addr := common.Hex2Addr(request.GetAddr())
	sendInfo := request.GetSrcSendInfo()
	receivedInfo := request.GetDstMinReceivedInfo()
	dstTransferId := request.GetDstTransferId()
	txHash := request.GetSrcTxHash()
	txType := request.GetType()
	withdrawSeqNum := request.GetWithdrawSeqNum()
	if txType == webapi.MarkTransferTypeRequest_TRANSFER_TYPE_SEND {
		err := dal.DB.MarkTransferSend(transferId, dstTransferId, addr.String(), sendInfo.GetToken().GetSymbol(),
			sendInfo.GetAmount(), receivedInfo.GetAmount(), txHash, sendInfo.GetChain().GetId(),
			receivedInfo.GetChain().GetId(), gs.f.GetUsdVolume(sendInfo.GetToken(), common.Str2BigInt(sendInfo.GetAmount())))
		if err != nil {
			return &webapi.MarkTransferResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "mark transfer refund failed",
				},
			}, nil
		}
	} else if txType == webapi.MarkTransferTypeRequest_TRANSFER_TYPE_REFUND {
		err := dal.DB.MarkTransferRefund(transferId, txHash, withdrawSeqNum)
		if err != nil {
			return &webapi.MarkTransferResponse{
				Err: &webapi.ErrMsg{
					Code: webapi.ErrCode_ERROR_CODE_COMMON,
					Msg:  "mark transfer refund failed",
				},
			}, nil
		}
	}
	return &webapi.MarkTransferResponse{
		Err: nil,
	}, nil
}

func (gs *GatewayService) GetLPInfoList(ctx context.Context, request *webapi.GetLPInfoListRequest) (*webapi.GetLPInfoListResponse, error) {
	panic("implement me")
}

func (gs *GatewayService) MarkLiquidity(ctx context.Context, request *webapi.MarkLiquidityRequest) (*webapi.MarkLiquidityResponse, error) {
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
	_db, err := dal.NewDAL("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", dbUrl), 10)
	if err != nil {
		return nil, err
	}
	dal.DB = _db
	gateway := &GatewayService{}

	gateway.f = fee.NewTokenPriceCache()
	return gateway, nil
}

// StartTokenPricePolling starts a loop with the given interval and 3s stdev for polling price
func (gs *GatewayService) StartChainTokenPolling(interval time.Duration) {
	go func() {
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			resp, err := cli.ChainTokensConfig(initClientCtx(), nil)
			if err != nil {
				log.Errorln("failed to load basic token info:", err)
			}
			chainTokens := resp.GetChainTokens()
			for chainId, assets := range chainTokens {
				for _, asset := range assets.Assets {
					token := asset.GetToken()
					_, found := gs.f.Prices[token.Symbol]
					if !found {
						gs.f.Prices[token.Symbol], err = gs.f.GetUsdPrice(token.Symbol)
						if err != nil {
							log.Error("get price error", err)
						}
					}
					dbErr := dal.DB.UpsertTokenBaseInfo(token.GetSymbol(), token.GetAddress(), asset.GetContractAddr(), asset.GetMaxAmt(), chainId, uint64(token.GetDecimal()))
					if dbErr != nil {
						log.Errorln("failed to write token:", err)
					}
				}
			}

		}
	}()
}

// todo add this in history query @aric
func UpdateTransferStatusInHistory(sender string, endTime time.Time, pageSize uint64) error {
	transferList, _, _, err := dal.DB.PaginateTransferList(sender, endTime, pageSize)

	var transferIds []string
	for _, transfer := range transferList {
		transferIds = append(transferIds, transfer.TransferId)
	}
	transferMap, err := cli.QueryTransferStatus(initClientCtx(), &types.QueryTransferStatusRequest{
		TransferId: transferIds,
	})
	if err != nil {
		return err
	}
	transferStatusMap := transferMap.Status

	for _, transfer := range transferList {
		transferId := transfer.TransferId
		status := transfer.Status
		srcChainId := transfer.SrcChainId
		txHash := transfer.SrcTxHash

		if status == types.TransferHistoryStatus_TRANSFER_SUBMITTING {
			var ctx context.Context
			receipt, recErr := relayer.CbrMgrInstance[srcChainId].TransactionReceipt(ctx, common.Bytes2Hash(common.Hex2Bytes(txHash)))
			if recErr == nil && receipt.Status != ethtypes.ReceiptStatusSuccessful {
				log.Warnf("find transfer failed, chain_id %d, hash:%s", srcChainId, txHash)
				dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(types.TransferHistoryStatus_TRANSFER_FAILED))
				if dbErr != nil {
					log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, txHash)
				}
			}
		}

		if err != nil {
			return err
		}
		if status == types.TransferHistoryStatus_TRANSFER_FAILED ||
			status == types.TransferHistoryStatus_TRANSFER_COMPLETED ||
			status == types.TransferHistoryStatus_TRANSFER_REFUNDED {
			continue
		}
		if transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_TO_BE_REFUNDED ||
			transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND ||
			transferStatusMap[transferId] == types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED {
			dbErr := dal.DB.UpdateTransferStatus(transferId, uint64(transferStatusMap[transferId]))
			if dbErr != nil {
				log.Warnf("UpdateTransferStatus failed, chain_id %d, hash:%s", srcChainId, txHash)
			}
		}
		return nil
	}
	return nil
}

func initClientCtx() client.Context {
	encodingConfig := app.MakeEncodingConfig()
	return client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithHomeDir(app.DefaultNodeHome).
		WithViper("SGN")
}
