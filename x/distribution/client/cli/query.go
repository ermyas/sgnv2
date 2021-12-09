package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	distQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the distribution module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	distQueryCmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryValidatorOutstandingRewards(),
		GetCmdQueryValidatorCommission(),
		GetCmdQueryDelegatorRewards(),
		GetCmdQueryCommunityPool(),
		GetCmdQueryStakingRewardInfo(),
		GetCmdQueryStakingRewardClaimInfo(),
		GetCmdQueryCBridgeFeeShareInfo(),
	)

	return distQueryCmd
}

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query distribution params",
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

// GetCmdQueryValidatorOutstandingRewards implements the query validator
// outstanding rewards command.
func GetCmdQueryValidatorOutstandingRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-outstanding-rewards [validator]",
		Args:  cobra.ExactArgs(1),
		Short: "Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations.

Example:
$ %s query distribution validator-outstanding-rewards 0x00078b31fa8b29a76bce074b5ea0d515a6aeaee7
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorOutstandingRewards(
				cmd.Context(),
				&types.QueryValidatorOutstandingRewardsRequest{ValidatorAddress: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Rewards)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryValidatorCommission implements the query validator commission command.
func GetCmdQueryValidatorCommission() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commission [validator]",
		Args:  cobra.ExactArgs(1),
		Short: "Query distribution validator commission",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query validator commission rewards from delegators to that validator.

Example:
$ %s query distribution commission 0x00078b31fa8b29a76bce074b5ea0d515a6aeaee7
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ValidatorCommission(
				cmd.Context(),
				&types.QueryValidatorCommissionRequest{ValidatorAddress: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Commission)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryDelegatorRewards implements the query delegator rewards command.
func GetCmdQueryDelegatorRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rewards [delegator-addr] [validator-addr]",
		Args:  cobra.RangeArgs(1, 2),
		Short: "Query all distribution delegator rewards or rewards from a particular validator",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all rewards earned by a delegator, optionally restrict to rewards from a single validator.

Example:
$ %s query distribution rewards 0xd0f2596d700c9bd4d605c938e586ec67b01c7364
$ %s query distribution rewards 0xd0f2596d700c9bd4d605c938e586ec67b01c7364 0x00078b31fa8b29a76bce074b5ea0d515a6aeaee7
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// query for rewards from a particular delegation
			ctx := cmd.Context()
			if len(args) == 2 {
				res, err := queryClient.DelegationRewards(
					ctx,
					&types.QueryDelegationRewardsRequest{DelegatorAddress: args[0], ValidatorAddress: args[1]},
				)
				if err != nil {
					return err
				}

				return clientCtx.PrintProto(res)
			}

			res, err := queryClient.DelegationTotalRewards(
				ctx,
				&types.QueryDelegationTotalRewardsRequest{DelegatorAddress: args[0]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryCommunityPool returns the command for fetching community pool info.
func GetCmdQueryCommunityPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "community-pool",
		Args:  cobra.NoArgs,
		Short: "Query the amount of coins in the community pool",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query all coins in the community pool which is under Governance control.

Example:
$ %s query distribution community-pool
`,
				version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.CommunityPool(cmd.Context(), &types.QueryCommunityPoolRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryStakingRewardInfo gets the staking reward info of a delegator
func GetCmdQueryStakingRewardInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking-reward-info [delegator-address]",
		Short: "query the staking reward info of an account",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the staking reward info of an account.

Example:
$ %s query staking-reward-info 0xd0f2596d700c9bd4d605c938e586ec67b01c7364
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
			res, err := queryClient.StakingRewardInfo(
				cmd.Context(),
				&types.QueryStakingRewardInfoRequest{DelegatorAddress: args[0]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(&res.RewardInfo)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func QueryStakingRewardInfo(goCtx context.Context, cliCtx client.Context, addr string) (*types.StakingRewardInfo, error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.StakingRewardInfo(
		goCtx,
		&types.QueryStakingRewardInfoRequest{DelegatorAddress: addr},
	)
	if err != nil {
		return nil, err
	}
	return &res.RewardInfo, nil
}

// GetCmdQueryStakingRewardClaimInfo gets the staking reward claim info of a delegator
func GetCmdQueryStakingRewardClaimInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking-reward-claim-info [delegator-address]",
		Short: "query the staking reward claim info of an account",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the staking reward claim info of an account.

Example:
$ %s query staking-reward-claim-info 0xd0f2596d700c9bd4d605c938e586ec67b01c7364
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
			res, err := queryClient.StakingRewardClaimInfo(
				cmd.Context(),
				&types.QueryStakingRewardClaimInfoRequest{DelegatorAddress: args[0]},
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

func QueryStakingRewardClaimInfo(goCtx context.Context, cliCtx client.Context, addr string) (*types.StakingRewardClaimInfo, error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.StakingRewardClaimInfo(
		goCtx,
		&types.QueryStakingRewardClaimInfoRequest{DelegatorAddress: addr},
	)
	if err != nil {
		return nil, err
	}
	return &res.RewardClaimInfo, nil
}

// GetCmdQueryCBridgeFeeShareInfo gets the cBridge fee share info of a delegator
func GetCmdQueryCBridgeFeeShareInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cbridge-fee-share-info [delegator-address]",
		Short: "query the cBridge fee share info of an account",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query the cBridge fee share info of an account.

Example:
$ %s query cbridge-fee-share-info 0xd0f2596d700c9bd4d605c938e586ec67b01c7364
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
			res, err := queryClient.CBridgeFeeShareInfo(
				cmd.Context(),
				&types.QueryCBridgeFeeShareInfoRequest{DelegatorAddress: args[0]},
			)
			if err != nil {
				return err
			}
			return cliCtx.PrintProto(&res.FeeShareInfo)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func QueryCBridgeFeeShareInfo(
	goCtx context.Context, cliCtx client.Context, delAddr string) (*types.ClaimableFeesInfo, error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.CBridgeFeeShareInfo(
		goCtx,
		&types.QueryCBridgeFeeShareInfoRequest{DelegatorAddress: delAddr},
	)
	if err != nil {
		return nil, err
	}
	return &res.FeeShareInfo, nil
}

func QueryPegBridgeFeesInfo(
	goCtx context.Context, cliCtx client.Context, delAddr string) (*types.ClaimableFeesInfo, error) {
	queryClient := types.NewQueryClient(cliCtx)
	res, err := queryClient.PegBridgeFeesInfo(
		goCtx,
		&types.QueryPegBridgeFeesInfoRequest{DelegatorAddress: delAddr},
	)
	if err != nil {
		return nil, err
	}
	return &res.FeesInfo, nil
}
