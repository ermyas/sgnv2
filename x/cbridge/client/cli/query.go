package cli

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	distrcli "github.com/celer-network/sgn-v2/x/distribution/client/cli"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

const (
	flagWdList = "wdlist"
	flagMinUSD = "min-usd"
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
		GetCmdQueryLPOrigin(),
		GetCmdQueryChainSigners(),
		GetCmdQueryLatestSigners(),
		GetCmdQueryChkLiqSum(),
		GetCmdQueryFeeShareInfo(),
		GetCmdQueryAssetPrice(),
		GetCmdQueryLpBalance(),
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

			config, err := QueryConfig(cliCtx)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(config)
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
			return cliCtx.PrintProto(params)
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
			fmt.Printf("withdraw message: %s, %s, last req time %s \n",
				withdrawOnChain.String(), resp.Detail.SignersStr(), common.TsSecToTime(uint64(resp.Detail.LastReqTime)))
			return nil
		},
	}
}

func GetCmdQueryLPOrigin() *cobra.Command {
	return &cobra.Command{
		Use:   "lp-origin [eth-addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query lp original chain id",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			req := &types.QueryLPOriginRequest{
				UsrAddr: args[0],
			}
			resp, err := QueryLPOrigin(cliCtx, req)
			if err != nil {
				return err
			}
			fmt.Printf("Original chain id of lp:%s is %d\n", args[0], resp.ChainId)
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
			resp, err := QueryChkLiqSum(cliCtx, &types.CheckLiqSumRequest{
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

func GetCmdQueryFeeShareInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fee-share [validator/delegator-eth-addr]",
		Args:  cobra.ExactArgs(1),
		Short: "Query fee share of a sgn validator or delgator",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			genWdList, err := cmd.Flags().GetBool(flagWdList)
			if err != nil {
				return err
			}
			minUsd, err := cmd.Flags().GetUint32(flagMinUSD)
			if err != nil {
				return err
			}

			wdList, err := GenerateClaimFeeWdList(cliCtx, args[0], minUsd, types.CBridgeFeeDenomPrefix, false)
			if err != nil {
				return fmt.Errorf("GenerateClaimFeeWdList err: %s", err)
			}
			if genWdList {
				fmt.Printf("\nvalidator withdraw fee inputs:\n")
				for _, wd := range wdList {
					fmt.Println(wd)
				}
			}
			return nil
		},
	}
	cmd.Flags().Bool(flagWdList, false, "generate withdraw file content")
	cmd.Flags().Uint32(flagMinUSD, 0, "minimal USD value to generate withraw request")

	return cmd
}

func GetCmdQueryAssetPrice() *cobra.Command {
	return &cobra.Command{
		Use:   "asset-price [symbol]",
		Args:  cobra.ExactArgs(1),
		Short: "Query asset price",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			price, extraPower10, err := QueryAssetPrice(cliCtx, args[0])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Printf("%s price: %f USD (%d %d)\n", args[0], float64(price)/math.Pow10(4+int(extraPower10)), price, extraPower10)
			return nil
		},
	}
}

func GetCmdQueryLpBalance() *cobra.Command {
	return &cobra.Command{
		Use:   "lp-balance [lp-addr] [symbol]",
		Args:  cobra.ExactArgs(2),
		Short: "Query lp balance for a certain token",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			config, err := QueryChainTokensConfig(cliCtx, &types.ChainTokensConfigRequest{})
			if err != nil {
				return err
			}

			decimals := make(map[uint64]int32)
			request := &types.LiquidityDetailListRequest{LpAddr: args[0]}
			for cidstr, tokens := range config.GetChainTokens() {
				chainId, err := strconv.ParseUint(cidstr, 10, 64)
				if err != nil {
					return err
				}
				for _, token := range tokens.GetTokens() {
					if token.Symbol == args[1] {
						request.ChainToken = append(
							request.ChainToken,
							&types.ChainTokenAddrPair{
								ChainId:   chainId,
								TokenAddr: token.Address,
							},
						)
						decimals[chainId] = token.GetDecimal()
						break
					}
				}
			}
			total := big.NewFloat(0)
			resp, err := QueryLiquidityDetailList(cliCtx, request)
			for _, detail := range resp.GetLiquidityDetail() {
				if detail.UsrLiquidity == "0" {
					continue
				}
				f, _ := new(big.Float).SetString(detail.UsrLiquidity)
				f = new(big.Float).Quo(f, new(big.Float).SetFloat64(math.Pow10(int(decimals[detail.ChainId]))))
				total.Add(total, f)

				fmt.Println("chainId:", detail.ChainId)
				fmt.Println("token:", detail.Token.Address)
				fmt.Printf("balance: %f (%s)\n\n", f, detail.UsrLiquidity)
			}
			fmt.Println("total balance:", total)

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
		key := args[0]
		queryClient := types.NewQueryClient(cliCtx)
		resp, err := queryClient.QueryDebugAny(context.Background(), &types.QueryDebugAnyRequest{Key: []byte(key)})
		if resp == nil || err != nil {
			log.Errorln("query err:", err)
			return err
		}
		res := resp.GetData()
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

func QueryChkLiqSum(cliCtx client.Context, req *types.CheckLiqSumRequest) (resp *types.CheckLiqSumResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryChkLiqSum(context.Background(), req)
	return
}

// Query config info
func QueryConfig(cliCtx client.Context) (config *types.CbrConfig, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryConfig(context.Background(), &types.EmptyRequest{})
	if resp != nil {
		config = resp.GetCbrConfig()
	}
	return
}

// Query params info
func QueryParams(cliCtx client.Context) (params *types.Params, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryParams(context.Background(), &types.EmptyRequest{})
	if resp != nil {
		params = resp.GetParams()
	}
	return
}

func QueryRelay(cliCtx client.Context, xrefId []byte) (relay *types.XferRelay, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryRelay(context.Background(), &types.QueryRelayRequest{XrefId: xrefId})
	if resp != nil {
		relay = resp.GetXferRelay()
	}
	return
}

func QueryChainTokensConfig(cliCtx client.Context, request *types.ChainTokensConfigRequest) (resp *types.ChainTokensConfigResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.ChainTokensConfig(context.Background(), request)
	return
}

func QueryFee(cliCtx client.Context, request *types.GetFeeRequest) (resp *types.GetFeeResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.GetFee(context.Background(), request)
	return
}

func QueryFeePerc(cliCtx client.Context, request *types.GetFeePercentageRequest) (resp *types.GetFeePercentageResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.GetFeePercentage(context.Background(), request)
	return
}

func QueryCheckChainTokenValid(cliCtx client.Context, request *types.CheckChainTokenValidRequest) (resp *types.CheckChainTokenValidResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryCheckChainTokenValid(context.Background(), request)
	return
}

func QueryTransferStatus(cliCtx client.Context, request *types.QueryTransferStatusRequest) (resp *types.QueryTransferStatusResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryTransferStatus(context.Background(), request)
	return
}

func QueryLiquidityDetailList(cliCtx client.Context, request *types.LiquidityDetailListRequest) (resp *types.LiquidityDetailListResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.LiquidityDetailList(context.Background(), request)
	return
}

func QueryTotalLiquidity(cliCtx client.Context, request *types.QueryTotalLiquidityRequest) (resp *types.QueryTotalLiquidityResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryTotalLiquidity(context.Background(), request)
	return
}

func QueryAddLiquidityStatus(cliCtx client.Context, request *types.QueryAddLiquidityStatusRequest) (resp *types.QueryLiquidityStatusResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryAddLiquidityStatus(context.Background(), request)
	return
}

func QueryWithdrawLiquidityStatus(cliCtx client.Context, request *types.QueryWithdrawLiquidityStatusRequest) (resp *types.QueryLiquidityStatusResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryWithdrawLiquidityStatus(context.Background(), request)
	return
}

func QueryLPOrigin(cliCtx client.Context, request *types.QueryLPOriginRequest) (resp *types.QueryLPOriginResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryLPOrigin(context.Background(), request)
	return
}

func QueryLiquidity(cliCtx client.Context, request *types.QueryLiquidityRequest) (resp *types.QueryLiquidityResponse, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err = queryClient.QueryLiquidity(context.Background(), request)
	return
}

func QueryChainSigners(cliCtx client.Context, chainId uint64) (chainSigners *types.ChainSigners, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryChainSigners(context.Background(), &types.QueryChainSignersRequest{ChainId: chainId})
	if resp != nil {
		chainSigners = resp.GetChainSigners()
	}
	return
}

func QueryLatestSigners(cliCtx client.Context) (latestSigners *types.LatestSigners, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryLatestSigners(context.Background(), &types.EmptyRequest{})
	if resp != nil {
		latestSigners = resp.GetLatestSigners()
	}
	return
}

func QueryAssets(cliCtx client.Context) (assets []*types.ChainAsset, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryAssets(context.Background(), &types.EmptyRequest{})
	if resp != nil {
		assets = resp.GetAssets()
	}
	return
}

func QueryAssetPrice(cliCtx client.Context, symbol string) (price, extraPower10 uint32, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryAssetPrice(context.Background(), &types.QueryAssetPriceRequest{Symbol: symbol})
	price = resp.GetPrice()
	extraPower10 = resp.GetExtraPower10()
	return
}

func GenerateClaimFeeWdList(cliCtx client.Context, delAddr string, minUsd uint32, denomPrefix string, isPegbr bool) ([]string, error) {
	var feeInfo *distrtypes.ClaimableFeesInfo
	var err error
	if isPegbr {
		feeInfo, err = distrcli.QueryPegBridgeFeesInfo(context.Background(), cliCtx, delAddr)
	} else {
		feeInfo, err = distrcli.QueryCBridgeFeeShareInfo(context.Background(), cliCtx, delAddr)
	}
	if err != nil {
		return []string{}, err
	}

	type AssetPrice struct {
		Price        uint32
		ExtraPower10 uint32
	}
	assetsPrice := make(map[string]*AssetPrice)

	assets := make(map[string]map[uint64]*types.ChainAsset) // symbol -> (chainId -> asset)
	assetsList, err := QueryAssets(cliCtx)
	for _, asset := range assetsList {
		_, ok := assets[asset.Symbol]
		if !ok {
			assets[asset.Symbol] = make(map[uint64]*types.ChainAsset)
			price, extraPower10, err2 := QueryAssetPrice(cliCtx, asset.Symbol)
			if err2 != nil {
				return []string{}, err2
			}
			assetsPrice[asset.Symbol] = &AssetPrice{Price: price, ExtraPower10: extraPower10}

		}
		assets[asset.Symbol][asset.ChainId] = asset
	}

	ts := time.Now().Unix()

	var totalValue float64
	var wdList []string
	fmt.Printf("claimable fee amounts:\n\n")
	for _, coin := range feeInfo.ClaimableFeeAmounts {
		amount := coin.Amount
		denom := coin.Denom
		symch := strings.TrimPrefix(denom, denomPrefix)
		symbol := strings.Split(symch, "/")[0]
		chainId, err := strconv.Atoi(strings.Split(symch, "/")[1])
		if err != nil {
			return []string{}, err
		}
		asset := assets[symbol][uint64(chainId)]
		fmt.Println("token:", symbol, "-", chainId, "-", asset.Addr)
		fmt.Println("amount: ", amount)

		famt, err := amount.QuoInt(
			sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(asset.Decimal)), nil))).Float64()
		if err != nil {
			return []string{}, err
		}
		value := famt * float64(assetsPrice[symbol].Price) / math.Pow10(4+int(assetsPrice[symbol].ExtraPower10))
		totalValue += value
		fmt.Printf("usd value: %0.2f\n\n", value)

		if value >= float64(minUsd) {
			wdList = append(wdList, fmt.Sprintf("%d %d %s", ts, chainId, asset.Addr))
			ts += 1
		}
	}
	fmt.Printf("total usd value: %0.2f\n", totalValue)
	return wdList, nil
}

func QueryAssetsSymbols(cliCtx client.Context, chainTokens []*types.ChainTokenAddrPair) (symbols []string, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryAssetsSymbols(context.Background(), &types.QueryAssetsSymbolsRequest{ChainTokens: chainTokens})
	if resp != nil {
		symbols = resp.Symbols
	}
	return
}

func QueryAssetsInfos(cliCtx client.Context, symbols []string, chainIds []uint64) (assets []*types.ChainAsset, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	resp, err := queryClient.QueryAssetsInfos(context.Background(), &types.QueryAssetsInfosRequest{Symbols: symbols, ChainIds: chainIds})
	if resp != nil {
		assets = resp.Assets
	}
	return
}
