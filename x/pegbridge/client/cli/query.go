package cli

import (
	"context"
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group pegbridge queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryConfig(),
		GetCmdQueryDeposit(),
		GetCmdQueryWithdraw(),
		GetCmdQueryMint(),
		GetCmdQueryBurn(),
	)

	return cmd
}

func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current pegbridge param",
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

func GetCmdQueryConfig() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Args:  cobra.NoArgs,
		Short: "Query the current pegbridge config",
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

func GetCmdQueryDeposit() *cobra.Command {
	return &cobra.Command{
		Use:   "deposit [deposit-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query pegbridge deposit info",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			depositId := args[0]
			deposit, err := QueryDepositInfo(cliCtx, depositId)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println(deposit.String())
			return nil
		},
	}
}

func GetCmdQueryWithdraw() *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw [withdraw-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query pegbridge withdraw info",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			withdrawId := args[0]
			withdraw, err := QueryWithdrawInfo(cliCtx, withdrawId)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println(withdraw.String())
			return nil
		},
	}
}

func GetCmdQueryMint() *cobra.Command {
	return &cobra.Command{
		Use:   "mint [mint-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query pegbridge mint info",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			mintId := args[0]
			mint, err := QueryMintInfo(cliCtx, mintId)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println(mint.String())
			return nil
		},
	}
}

func GetCmdQueryBurn() *cobra.Command {
	return &cobra.Command{
		Use:   "burn [burn-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query pegbridge burn info",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			burnId := args[0]
			burn, err := QueryBurnInfo(cliCtx, burnId)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println(burn.String())
			return nil
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

func QueryConfig(cliCtx client.Context) (config *types.PegConfig, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	config, err = queryClient.Config(context.Background(), &types.QueryConfigRequest{})
	return
}

func QueryDepositInfo(cliCtx client.Context, depositId string) (deposit types.DepositInfo, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.DepositInfo(context.Background(), &types.QueryDepositInfoRequest{
		DepositId: depositId,
	})
	if err != nil {
		return
	}
	deposit = res.DepositInfo
	return
}

func QueryWithdrawInfo(cliCtx client.Context, wdId string) (withdraw types.WithdrawInfo, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.WithdrawInfo(context.Background(), &types.QueryWithdrawInfoRequest{
		WithdrawId: wdId,
	})
	if err != nil {
		return
	}
	withdraw = res.WithdrawInfo
	return
}

func QueryMintInfo(cliCtx client.Context, mintId string) (mint types.MintInfo, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.MintInfo(context.Background(), &types.QueryMintInfoRequest{
		MintId: mintId,
	})
	if err != nil {
		return
	}
	mint = res.MintInfo
	return
}

func QueryBurnInfo(cliCtx client.Context, burnId string) (burn types.BurnInfo, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.BurnInfo(context.Background(), &types.QueryBurnInfoRequest{
		BurnId: burnId,
	})
	if err != nil {
		return
	}
	burn = res.BurnInfo
	return
}

func QueryFeeClaimInfo(cliCtx client.Context, delAddr eth.Addr, nonce uint64) (feeClaim types.FeeClaimInfo, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.FeeClaimInfo(context.Background(), &types.QueryFeeClaimInfoRequest{
		Address: eth.Addr2Hex(delAddr),
		Nonce:   nonce,
	})
	if err != nil {
		return
	}
	feeClaim = res.FeeClaimInfo
	return
}

func QueryOrigPeggedPairs(cliCtx client.Context, params *types.QueryOrigPeggedPairsRequest) (resp []types.OrigPeggedPair, err error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.OrigPeggedPairs(context.Background(), params)
	if err != nil {
		return
	}

	resp = res.Pairs
	return
}
