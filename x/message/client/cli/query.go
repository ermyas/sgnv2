package cli

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/celer-network/goutils/log"
	comtypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/message/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group message queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams(),
		CmdQueryExecutionContexts(),
		CmdQueryExecutionContextBySrcTransfer(),
		CmdQueryMessage(),
		CmdQueryMessageBus())
	return cmd
}

func CmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query the current message params",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			params, err := QueryParams(cliCtx)
			if err != nil {
				log.Errorln("query params error", err)
				return err
			}

			return cliCtx.PrintProto(params)
		},
	}
}

const flagInlineFilter = "filter"
const flagJsonFilter = "json-filter"
const flagAll = "all"

func CmdQueryExecutionContexts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exec-ctxs [request json]",
		Short: "Query the execution contexts of all messages",
		Long: `
inline filter format:
--filter '5:0x09E4534B11D400BFcd2026b69E399763CeAfB42D,97:0x570F9c2f224b002d75F287f5430Bc9598E850E13'
json file filter format:
--json-filter <path-to-json>
{
	"contract_infos": [
		{
			"chain_id": 5,
			"address": "0x09E4534B11D400BFcd2026b69E399763CeAfB42D"
		},
		{
			"chain_id": 97,
			"address": "0x570F9c2f224b002d75F287f5430Bc9598E850E13"
		}
	]
}
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			inlineFilterSlice, err := cmd.Flags().GetStringSlice(flagInlineFilter)
			if err != nil {
				return err
			}
			jsonFilterPath, err := cmd.Flags().GetString(flagJsonFilter)
			if err != nil {
				return err
			}
			all, err := cmd.Flags().GetBool(flagAll)
			if err != nil {
				return err
			}

			req := new(types.QueryExecutionContextsRequest)
			if all {
				req.All = true
			} else if len(inlineFilterSlice) != 0 {
				for _, filter := range inlineFilterSlice {
					params := strings.Split(filter, ":")
					if len(params) != 2 {
						return fmt.Errorf("malformatted filter")
					}
					chainIdStr := strings.Trim(params[0], " ")
					chainIdInt, err := strconv.Atoi(chainIdStr)
					if err != nil {
						return err
					}
					chainId := uint64(chainIdInt)
					addr := strings.Trim(params[1], " ")
					req.ContractInfos = append(req.ContractInfos, &comtypes.ContractInfo{
						ChainId: chainId,
						Address: addr,
					})
				}
			} else if len(jsonFilterPath) != 0 {
				filter, err := os.ReadFile(jsonFilterPath)
				if err != nil {
					return err
				}
				req.Unmarshal(filter)
			}

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			resp, err := QueryExecutionContexts(cliCtx, req)
			if err != nil {
				log.Errorln("query execution contexts error", err)
				return err
			}
			return cliCtx.PrintProto(resp)
		},
	}
	cmd.Flags().StringSlice(flagInlineFilter, []string{}, "contract filters")
	cmd.Flags().String(flagJsonFilter, "", "filter contracts with a json file")
	cmd.Flags().Bool(flagAll, false, "queries all pending execution contexts")
	return cmd
}

func CmdQueryExecutionContextBySrcTransfer() *cobra.Command {
	return &cobra.Command{
		Use:   "exe-ctx [bridge-type] [src-transfer-id]",
		Short: "Query execution context by bridge type (1:liquidity, 2:pegvault, 3:pegbridge) and src transferId",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			bridgeType, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			transferId := eth.Hex2Hash(args[1])
			resp, err := QueryExecutionContextBySrcTransfer(cliCtx, types.BridgeType(bridgeType), transferId)
			if err != nil {
				log.Errorln("query execution context error", err)
				return err
			}

			return cliCtx.PrintProto(resp)
		},
	}
}

func CmdQueryMessage() *cobra.Command {
	return &cobra.Command{
		Use:   "message [message-id]",
		Short: "Query message details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			messageId := args[0]
			msg, err := QueryMessage(cliCtx, messageId)
			if err != nil {
				log.Errorln("query message error", err)
				return err
			}

			return cliCtx.PrintProto(&msg)
		},
	}
}

func CmdQueryMessageBus() *cobra.Command {
	return &cobra.Command{
		Use:   "message-bus [chain-id]",
		Short: "Query message bus contract info",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			chainIdStr := args[0]
			chainId, err := strconv.Atoi(chainIdStr)
			if err != nil {
				log.Errorln("convert chainId string error")
				return err
			}

			msgBus, err := QueryMessageBus(cliCtx, uint64(chainId))
			if err != nil {
				log.Errorln("query message bus error", err)
				return err
			}

			return cliCtx.PrintProto(&msgBus)
		},
	}
}

func QueryParams(cliCtx client.Context) (params *types.Params, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
	if err != nil {
		return
	}
	params = &res.Params
	return
}

func QueryExecutionContexts(cliCtx client.Context, req *types.QueryExecutionContextsRequest) (resp *types.QueryExecutionContextsResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	return queryClient.ExecutionContexts(context.Background(), req)
}

func QueryExecutionContextBySrcTransfer(
	cliCtx client.Context, srcBridgeType types.BridgeType, srcTransferId eth.Hash) (*types.ExecutionContext, error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.ExecutionContextBySrcTransfer(context.Background(),
		&types.QueryExecutionContextBySrcTransferRequest{
			SrcTransferId: srcTransferId.Hex(),
			SrcBridgeType: srcBridgeType,
		})
	if err != nil {
		return nil, err
	}
	return res.ExecutionContext, nil
}

func QueryMessage(cliCtx client.Context, messageId string) (msg types.Message, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.Message(context.Background(), &types.QueryMessageRequest{MessageId: messageId})
	if err != nil {
		return
	}
	msg = res.Message
	return
}

func QueryMessageBus(cliCtx client.Context, chainId uint64) (msgBus types.MessageBusInfo, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.MessageBus(context.Background(), &types.QueryMessageBusRequest{ChainId: chainId})
	if err != nil {
		return
	}
	msgBus = res.MessageBus
	return
}

func QueryFeeClaimInfo(cliCtx client.Context, delAddr eth.Addr) (feeClaim types.FeeClaimInfo, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.FeeClaimInfo(context.Background(), &types.QueryFeeClaimInfoRequest{
		Address: eth.Addr2Hex(delAddr),
	})
	if err != nil {
		return
	}
	feeClaim = res.FeeClaimInfo
	return
}
