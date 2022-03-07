package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/celer-network/goutils/log"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/message/types"
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

func CmdQueryExecutionContexts() *cobra.Command {
	return &cobra.Command{
		Use:   "exec-ctxs [request json]",
		Short: "Query the execution contexts of all messages",
		Long: `
request json should be like below
{
	contract_infos:[
		{
			"address": "3ff73bab93c505809c68b0a8e4321a2713d9255c",
			"chain_id": 883
		},
		{
			"address": "58712219a4bdbb0e581dcaf6f5c4c2b2d2f42158",
			"chain_id": 884
		}
	]
}
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			req := new(types.QueryExecutionContextsRequest)
			if args[0] != "" {
				err := req.Unmarshal([]byte(args[0]))
				if err != nil {
					return err
				}
			}
			resp, err := QueryExecutionContexts(cliCtx, req)
			if err != nil {
				log.Errorln("query execution contexts error", err)
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
