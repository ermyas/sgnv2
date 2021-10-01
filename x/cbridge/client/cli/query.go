package cli

import (
	"fmt"
	"strconv"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group cbridge queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryConfig(),
		GetCmdChainTokensConfig(),
		qRelayCmd,
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

// relay and sigs about this xfer
var qRelayCmd = &cobra.Command{
	Use:   "relay",
	Args:  cobra.ExactArgs(1),
	Short: "Query relay for xfer id",
	RunE: func(cmd *cobra.Command, args []string) error {
		cliCtx, _ := client.GetClientQueryContext(cmd)
		xfid := eth.Hex2Bytes(args[0])

		resp, err := QueryRelay(cliCtx, xfid)
		if err != nil {
			log.Errorln("query error", err)
			return err
		}
		return cliCtx.PrintObjectLegacy(resp)
	},
}

var qSignersCmd = &cobra.Command{
	Use:   "signers",
	Args:  cobra.ExactArgs(1),
	Short: "Query signers for chainid",
	RunE: func(cmd *cobra.Command, args []string) error {
		cliCtx, _ := client.GetClientQueryContext(cmd)
		chid, _ := strconv.Atoi(args[0])

		resp, err := QueryChainSigners(cliCtx, uint64(chid))
		if err != nil {
			log.Errorln("query error", err)
			return err
		}
		return cliCtx.PrintObjectLegacy(resp)
	},
}

// GetCmdQueryConfig implements the params query command.
func GetCmdQueryConfig() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Args:  cobra.NoArgs,
		Short: "Query the current cbridge config",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			params, err := QueryParams(cliCtx)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintObjectLegacy(params)
		},
	}
}

func GetCmdChainTokensConfig() *cobra.Command {
	return &cobra.Command{
		Use:   "chaintokens",
		Args:  cobra.NoArgs,
		Short: "Query the chain tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			resp, err := QueryChainTokensConfig(cliCtx, &types.ChainTokensConfigRequest{})
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintObjectLegacy(resp)
		},
	}
}

// Query params info
func QueryParams(cliCtx client.Context) (params types.CbrConfig, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &params)
	return
}

func QueryRelay(cliCtx client.Context, xrefId []byte) (relay types.XferRelay, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(types.NewQueryRelayParams(xrefId))
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryRelay)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &relay)
	return
}

func QueryChainTokensConfig(cliCtx client.Context, request *types.ChainTokensConfigRequest) (resp *types.ChainTokensConfigResponse, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryChainTokensConfig)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	resp = new(types.ChainTokensConfigResponse)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

func QueryFee(cliCtx client.Context, request *types.GetFeeRequest) (resp *types.GetFeeResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFee)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.GetFeeResponse)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

func QueryTransferStatus(cliCtx client.Context, request *types.QueryTransferStatusRequest) (resp *types.QueryTransferStatusResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTransferStatus)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.QueryTransferStatusResponse)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

func QueryLiquidityDetailList(cliCtx client.Context, request *types.LiquidityDetailListRequest) (resp *types.LiquidityDetailListResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryLiquidityDetailList)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.LiquidityDetailListResponse)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

func QueryAddLiquidityStatus(cliCtx client.Context, request *types.QueryAddLiquidityStatusRequest) (resp *types.QueryLiquidityStatusResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryAddLiquidityStatus)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.QueryLiquidityStatusResponse)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

func QueryWithdrawLiquidityStatus(cliCtx client.Context, request *types.QueryWithdrawLiquidityStatusRequest) (resp *types.QueryLiquidityStatusResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryWithdrawLiquidityStatus)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.QueryLiquidityStatusResponse)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

func QueryChainSigners(cliCtx client.Context, chainId uint64) (chainSigners *types.ChainSigners, err error) {
	params := types.NewQueryChainSignersParams(chainId)
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return
	}
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryChainSigners)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}
	chainSigners = new(types.ChainSigners)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, chainSigners)
	return
}

func QueryLatestSigners(cliCtx client.Context) (latestSigners *types.LatestSigners, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryLatestSigners)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}
	latestSigners = new(types.LatestSigners)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, latestSigners)
	return
}
