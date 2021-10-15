package cli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group farming queries under a subcommand
	farmQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	farmQueryCmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryPool(),
		GetCmdQueryPools(),
		GetCmdQueryToken(),
		GetCmdQueryTokens(),
		GetCmdQueryNumPools(),
		GetCmdQueryStakeInfo(),
		GetCmdQueryEarnings(),
		GetCmdQueryAccountInfo(),
		GetCmdQueryAccountsStakedIn(),
		GetCmdQueryRewardClaimInfo(),
	)

	return farmQueryCmd
}

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query farming params",
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(&res.Params)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryPool gets the pool query command.
func GetCmdQueryPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool [pool-name]",
		Short: "query a pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about the reward token, the staked balance and the reward amount per block.

Example:
$ %s query farming pool cbridge-DAI/1
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.Pool(
				cmd.Context(),
				&types.QueryPoolRequest{PoolName: args[0]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(&res.Pool)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryPools gets the pools query command.
func GetCmdQueryPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pools",
		Short: "query for all pools",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about all pools.

Example:
$ %s query farming pools
`,
				version.AppName,
			),
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.Pools(
				cmd.Context(),
				&types.QueryPoolsRequest{},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryToken gets the token query command.
func GetCmdQueryToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token [chain-id] [symbol]",
		Short: "query a token",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the info about a token.

Example:
$ %s query farming token 1 DAI
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			chainId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.Token(
				cmd.Context(),
				&types.QueryTokenRequest{ChainId: chainId, Symbol: args[1]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(&res.Token)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryTokens gets the tokens query command.
func GetCmdQueryTokens() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tokens",
		Short: "query all tokens",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the info about all tokens.

Example:
$ %s query farming tokens
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.Tokens(
				cmd.Context(),
				&types.QueryTokensRequest{},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryEarnings gets the earnings query command.
func GetCmdQueryEarnings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "earnings [pool-name] [address]",
		Short: "query the current rewards of an account",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query available rewards for an address.

Example:
$ %s query farming earnings cbridge-DAI/1 0xab5801a7d398351b8be11c439e05c5b3259aec9b
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.Earnings(
				cmd.Context(),
				&types.QueryEarningsRequest{PoolName: args[0], Address: args[1]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(&res.Earnings)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAccountInfo gets the account info query command.
func GetCmdQueryAccountInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-info [address]",
		Short: "query the info of a farming account",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the info of a farming account.

Example:
$ %s query farming account-info 0xab5801a7d398351b8be11c439e05c5b3259aec9b
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.AccountInfo(
				cmd.Context(),
				&types.QueryAccountInfoRequest{Address: args[0]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryPoolNum gets the pool number query command.
func GetCmdQueryNumPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "num-pools",
		Short: "query the number of pools",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the number of pools that already exist.

Example:
$ %s query farming num-pools
`,
				version.AppName,
			),
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.NumPools(
				cmd.Context(),
				&types.QueryNumPoolsRequest{},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryAccountsStakedIn gets all addresses of accounts that staked tokens in a specific pool
func GetCmdQueryAccountsStakedIn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accounts-staked-in [pool-name]",
		Short: "query the addresses of accounts staked in a pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all the addresses of accounts that have staked tokens in a specific pool.

Example:
$ %s query farming accounts-staked-in cbridge-DAI/1
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.AccountsStakedIn(
				cmd.Context(),
				&types.QueryAccountsStakedInRequest{PoolName: args[0]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryStakeInfo gets the stake info of an account in a specific pool
func GetCmdQueryStakeInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-info [pool-name] [address]",
		Short: "query the stake info of an account on a pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the stake info of an account in a specific pool.

Example:
$ %s query farming stake-info cbridge-DAI/1 0xab5801a7d398351b8be11c439e05c5b3259aec9b
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.StakeInfo(
				cmd.Context(),
				&types.QueryStakeInfoRequest{PoolName: args[0], Address: args[1]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(&res.StakeInfo)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryRewardClaimInfo gets the reward claim info of an account
func GetCmdQueryRewardClaimInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reward-claim-info [address]",
		Short: "query the stake info of an account on a pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the reward claim info of an account.

Example:
$ %s query farming reward-claim-info 0xab5801a7d398351b8be11c439e05c5b3259aec9b
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(cliCtx)
			res, err := queryClient.RewardClaimInfo(
				cmd.Context(),
				&types.QueryRewardClaimInfoRequest{Address: args[0]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(&res.RewardClaimInfo)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func QueryRewardClaimInfo(goCtx context.Context, cliCtx client.Context, addr string) (*types.RewardClaimInfo, error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.RewardClaimInfo(
		goCtx,
		&types.QueryRewardClaimInfoRequest{Address: addr},
	)
	if err != nil {
		return nil, err
	}
	return &res.RewardClaimInfo, nil
}
