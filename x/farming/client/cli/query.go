package cli

import (
	"fmt"
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
		GetCmdQueryPool(),
		GetCmdQueryPools(),
		GetCmdQueryNumPools(),
		GetCmdQueryStakeInfo(),
		GetCmdQueryEarnings(),
		GetCmdQueryStakedPools(),
		GetCmdQueryAccountsStakedIn(),
	)

	return farmQueryCmd
}

// GetCmdQueryPool gets the pool query command.
func GetCmdQueryPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pool [pool-name]",
		Short: "query a pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about the reward token, the staked balance and the reward amount per block.

Example:
$ %s query farming pool cbridge-1-DAI
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

// GetCmdQueryEarnings gets the earnings query command.
func GetCmdQueryEarnings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "earnings [pool-name] [address]",
		Short: "query the current rewards of an account",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query available rewards for an address.

Example:
$ %s query farming earnings cbridge-1-DAI 0xab5801a7d398351b8be11c439e05c5b3259aec9b
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

// GetCmdQueryStakedPools gets the staked pools query command.
func GetCmdQueryStakedPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staked-pools [address]",
		Short: "query the pools that an account has staked tokens in",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all the pools that an account has staked tokens in.

Example:
$ %s query farming staked-pools 0xab5801a7d398351b8be11c439e05c5b3259aec9b
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
			res, err := queryClient.StakedPools(
				cmd.Context(),
				&types.QueryStakedPoolsRequest{Address: args[0]},
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
$ %s query farming accounts-staked-in cbridge-1-DAI
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
$ %s query farming stake-info cbridge-1-DAI 0xab5801a7d398351b8be11c439e05c5b3259aec9b
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
