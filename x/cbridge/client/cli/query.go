package cli

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

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
		GetCmdQueryRelay(),
		GetCmdQueryChainSigners(),
		GetCmdQueryLatestSigners(),
		qDebugAnyCmd,
	)
	// this line is used by starport scaffolding # 1

	return cmd
}

// relay and sigs about this xfer
func GetCmdQueryRelay() *cobra.Command {
	return &cobra.Command{
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
			relayOnChain := new(types.RelayOnChain)
			err = relayOnChain.Unmarshal(resp.Relay)
			if err != nil {
				log.Errorln("unmarshal relay error", err)
				return err
			}
			fmt.Printf("Relay: %s, %s", relayOnChain.String(), resp.SignersStr())
			return nil
		},
	}
}

func GetCmdQueryChainSigners() *cobra.Command {
	return &cobra.Command{
		Use:   "chain-signers",
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
			fmt.Println(resp.String())
			return nil
		},
	}
}

func GetCmdQueryLatestSigners() *cobra.Command {
	return &cobra.Command{
		Use:   "latest-signers",
		Args:  cobra.NoArgs,
		Short: "Query the latest signers",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			resp, err := QueryLatestSigners(cliCtx)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			fmt.Println(resp.String())
			return nil
		},
	}
}

// it's by design this doesn't have pkg level func so it can only be called via cmd line
var qDebugAnyCmd = &cobra.Command{
	Use:   "getany",
	Args:  cobra.ExactArgs(1),
	Short: "Query any kv value for given full key",
	RunE: func(cmd *cobra.Command, args []string) error {
		cliCtx, _ := client.GetClientQueryContext(cmd)
		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDebugAny)
		key := args[0]
		res, err := common.RobustQueryWithData(cliCtx, route, []byte(key))
		if err != nil {
			log.Errorln("query err:", err)
			return err
		}
		if len(res) == 0 {
			return cliCtx.PrintString("nil value")
		}
		// now switch on key prefix to decode and print properly
		switch {
		// big.Int Bytes
		case pre(key, "lm"), pre(key, "evliqadd"), pre(key, "withdrawSeqNum"), pre(key, "lpfee"), pre(key, "sgnfee"), pre(key, "cfg-feeperc"):
			return cliCtx.PrintString(new(big.Int).SetBytes(res).String())
		case pre(key, "evsend"):
			return cliCtx.PrintString(types.XferStatus(res[0]).String())
		case pre(key, "evrelay"):
			// val is src transfer id
			return cliCtx.PrintString(eth.Bytes2Hex(res))
		case pre(key, "xferRelay"):
			pbmsg := new(types.XferRelay)
			pbmsg.Unmarshal(res)
			return cliCtx.PrintString(common.PbToJSONHexBytes(pbmsg))
		case pre(key, "xferRefund"):
			pbmsg := new(types.WithdrawOnchain)
			pbmsg.Unmarshal(res)
			return cliCtx.PrintString(common.PbToJSONHexBytes(pbmsg))
		case pre(key, "wdDetail"):
			pbmsg := new(types.WithdrawDetail)
			pbmsg.Unmarshal(res)
			return cliCtx.PrintString(common.PbToJSONHexBytes(pbmsg))
		case pre(key, "cfg-ch2sym"):
			return cliCtx.PrintString(string(res))
		case pre(key, "cfg-sym2info"):
			pbmsg := new(types.ChainAsset)
			pbmsg.Unmarshal(res)
			return cliCtx.PrintString(common.PbToJSONHexBytes(pbmsg))
		case pre(key, "cfg-chpair"):
			pbmsg := new(types.ChainPair)
			pbmsg.Unmarshal(res)
			return cliCtx.PrintString(common.PbToJSONHexBytes(pbmsg))
		default:
			return cliCtx.PrintString(eth.Bytes2Hex(res))
		}
	},
}

func pre(a, pre string) bool {
	return strings.HasPrefix(a, pre)
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

			return cliCtx.PrintProto(&params)
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

			return cliCtx.PrintProto(resp)
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

func QueryFeePerc(cliCtx client.Context, request *types.GetFeePercentageRequest) (resp *types.GetFeePercentageResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryFeePerc)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.GetFeePercentageResponse)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

func QueryCheckChainTokenValid(cliCtx client.Context, request *types.CheckChainTokenValidRequest) (resp *types.CheckChainTokenValidResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryCheckChainTokenValid)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.CheckChainTokenValidResponse)
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
