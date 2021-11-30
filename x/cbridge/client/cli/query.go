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
		GetCmdQueryParams(),
		GetCmdChainTokensConfig(),
		GetCmdQueryTransfer(),
		GetCmdQueryWithdraw(),
		GetCmdQueryChainSigners(),
		GetCmdQueryLatestSigners(),
		GetCmdQueryChkLiqSum(),
		qDebugAnyCmd,
	)
	// this line is used by starport scaffolding # 1

	return cmd
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

			params, err := QueryConfig(cliCtx)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(&params)
		},
	}
}

func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current cbridge parameters information",
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

func GetCmdQueryTransfer() *cobra.Command {
	return &cobra.Command{
		Use:   "transfer [src-transfer-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query transfer info",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, _ := client.GetClientQueryContext(cmd)

			req := &types.QueryTransferStatusRequest{
				TransferId: []string{args[0]},
			}
			res, err := QueryTransferStatus(cliCtx, req)
			if err != nil {
				return err
			}
			fmt.Println(res)

			xferId := eth.Hex2Bytes(args[0])
			resp, err := QueryRelay(cliCtx, xferId)
			if err != nil {
				return err
			}
			relayOnChain := new(types.RelayOnChain)
			err = relayOnChain.Unmarshal(resp.Relay)
			if err != nil {
				return err
			}
			fmt.Printf("relay message: %s, %s, fee base %s perc %s, last req time %s \n",
				relayOnChain.String(), resp.SignersStr(),
				big.NewInt(0).SetBytes(resp.BaseFee), big.NewInt(0).SetBytes(resp.PercFee),
				common.TsSecToTime(uint64(resp.LastReqTime)))
			return nil
		},
	}
}

func GetCmdQueryWithdraw() *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw [eth-addr] [seq-num]",
		Args:  cobra.ExactArgs(2),
		Short: "Query withdraw info",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			seqNum, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			req := &types.QueryWithdrawLiquidityStatusRequest{
				SeqNum:  uint64(seqNum),
				UsrAddr: args[0],
			}
			resp, err := QueryWithdrawLiquidityStatus(cliCtx, req)
			if err != nil {
				return err
			}
			withdrawOnChain := new(types.WithdrawOnchain)
			err = withdrawOnChain.Unmarshal(resp.Detail.WdOnchain)
			if err != nil {
				return err
			}
			fmt.Printf("status: %s\n", resp.Status)
			fmt.Printf("withdraw message: %s, %s \n", withdrawOnChain.String(), resp.Detail.SignersStr())
			return nil
		},
	}
}

func GetCmdQueryChainSigners() *cobra.Command {
	return &cobra.Command{
		Use:   "chain-signers [chain-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query signers for chainid",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			chid, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

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

func GetCmdQueryChkLiqSum() *cobra.Command {
	return &cobra.Command{
		Use:   "liqsum [chainid] [token]", // requird, keep square bracket for consistency w/ other cmds
		Args:  cobra.ExactArgs(2),
		Short: "Query liq sum of chain,token, return both liqsum and itersum over lm- keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			chainid, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("%s not valid chainid, err:%w", args[0], err)
			}
			resp, err := QueryChkLiqSum(cliCtx, &types.CheckLiqSumReq{
				ChainId:   uint64(chainid),
				TokenAddr: args[1],
			})
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println("liqsum:", resp.Liqsum)
			fmt.Println("sumitr:", resp.Sumiter)
			return nil
		},
	}
}

// it's by design this doesn't have pkg level func so it can only be called via cmd line
var qDebugAnyCmd = &cobra.Command{
	Use:   "getany [internal-db-key]",
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
		case pre(key, "lm"), pre(key, "liqsum"), pre(key, "evliqadd"), pre(key, "wdDetail"), pre(key, "lpfee"), pre(key, "sgnfee"), pre(key, "cfg-feeperc"):
			return cliCtx.PrintString(new(big.Int).SetBytes(res).String())
		case pre(key, "evsend"):
			return cliCtx.PrintString(types.XferStatus(res[0]).String())
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

func QueryChkLiqSum(cliCtx client.Context, req *types.CheckLiqSumReq) (resp *types.CheckLiqSumResp, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(req)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryChkLiqSum)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.CheckLiqSumResp)
	err = cliCtx.Codec.Unmarshal(res, resp)
	return
}

// Query config info
func QueryConfig(cliCtx client.Context) (config types.CbrConfig, err error) {
	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryConfig)
	res, err := common.RobustQuery(cliCtx, route)
	if err != nil {
		return
	}

	err = cliCtx.LegacyAmino.UnmarshalJSON(res, &config)
	return
}

// Query params info
func QueryParams(cliCtx client.Context) (params types.Params, err error) {
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

func QueryTotalLiquidity(cliCtx client.Context, request *types.QueryTotalLiquidityRequest) (resp *types.QueryTotalLiquidityResponse, err error) {
	data, err := cliCtx.LegacyAmino.MarshalJSON(request)
	if err != nil {
		return
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryTotalLiquidity)
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	resp = new(types.QueryTotalLiquidityResponse)
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
	params := types.NewQueryChainParams(chainId)
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
