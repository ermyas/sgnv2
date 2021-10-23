package cli

import (
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	FlagStatus = "status"
)

func GetQueryCmd() *cobra.Command {
	stakingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the staking module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	stakingQueryCmd.AddCommand(
		GetCmdQueryValidator(),
		GetCmdQueryValidators(),
		GetCmdQueryTransactors(),
		GetCmdQueryDelegation(),
		GetCmdQueryDelegatorDelegations(),
		GetCmdQueryValidatorDelegations(),
		GetCmdQuerySyncer(),
		GetCmdQueryParams(),
	)
	return stakingQueryCmd
}

// GetCmdQueryValidator implements the validator query command.
func GetCmdQueryValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator [validator-eth-addr]",
		Short: "Query a validator",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query details about an individual validator by Ethereum address.

Example:
$ %s query staking validator 0x00078b31fa8b29a76bce074b5ea0d515a6aeaee7
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryValidatorRequest{ValidatorAddr: args[0]}
			res, err := queryClient.Validator(cmd.Context(), params)
			if err != nil {
				return err
			}

			PrintValidator(clientCtx, res.Validator)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryValidators implements the query all validators command.
func GetCmdQueryValidators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validators",
		Short: "query all validators",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)
			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			var status string
			st := viper.GetString(FlagStatus)
			switch st {
			case "bonded":
				status = types.BondStatus_name[int32(types.Bonded)]
			case "unbonded":
				status = types.BondStatus_name[int32(types.Unbonded)]
			case "unbonding":
				status = types.BondStatus_name[int32(types.Unbonding)]
			}

			result, err := queryClient.Validators(cmd.Context(), &types.QueryValidatorsRequest{
				Status:     status,
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}
			validators := types.Validators(result.GetValidators())
			validators.Sort()

			if status == "" {
				st = "all"
			}
			fmt.Printf("Number of %s validators: %d\n\n", st, len(validators))
			for _, validator := range validators {
				PrintValidator(clientCtx, validator)
			}
			return nil
		},
	}
	cmd.Flags().String(FlagStatus, "", "Validator status (bonded | unbonded | unbonding)")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryTransactors implements the transactors query command
// TODO: Migrate to gRPC
func GetCmdQueryTransactors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transactors [validator-eth-addr]",
		Short: "query validator transactors by validator's Ethereum address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			transactors, err := QueryTransactors(cliCtx, args[0])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println(transactors.Transactors)
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryDelegation the query delegation command.
func GetCmdQueryDelegation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegation [delegator-eth-addr] [validator-eth-addr]",
		Short: "Query a delegation based on delegator and validator's Ethereum addresses",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query delegations for an individual delegator on an individual validator.

Example:
$ %s query staking delegation 0x00078b31fa8b29a76bce074b5ea0d515a6aeaee7 0xd0f2596d700c9bd4d605c938e586ec67b01c7364
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDelegationRequest{
				DelegatorAddr: args[0],
				ValidatorAddr: args[1],
			}

			res, err := queryClient.Delegation(cmd.Context(), params)
			if err != nil {
				return err
			}

			fmt.Println(res.DelegationResponse.YamlStr())
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryDelegations implements the command to query all the delegations
// made from one delegator.
func GetCmdQueryDelegatorDelegations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegator-delegations [delegator-eth-addr]",
		Short: "Query all delegations made by one delegator",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query delegations for an individual delegator on all validators.

Example:
$ %s query staking delegations 0xd0f2596d700c9bd4d605c938e586ec67b01c7364
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryDelegatorDelegationsRequest{
				DelegatorAddr: args[0],
				Pagination:    pageReq,
			}

			res, err := queryClient.DelegatorDelegations(cmd.Context(), params)
			if err != nil {
				return err
			}

			for _, d := range res.DelegationResponses {
				fmt.Println(d.YamlStr())
			}
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "delegations")

	return cmd
}

// GetCmdQueryValidatorDelegations implements the command to query all the
// delegations to a specific validator.
func GetCmdQueryValidatorDelegations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-delegations [validator-eth-addr]",
		Short: "Query all delegations made to one validator",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query delegations on an individual validator.

Example:
$ %s query staking delegations-to 0x00078b31fa8b29a76bce074b5ea0d515a6aeaee7
`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			params := &types.QueryValidatorDelegationsRequest{
				ValidatorAddr: args[0],
				Pagination:    pageReq,
			}

			res, err := queryClient.ValidatorDelegations(cmd.Context(), params)
			if err != nil {
				return err
			}

			for _, d := range res.DelegationResponses {
				fmt.Println(d.YamlStr())
			}
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "validator delegations")

	return cmd
}

// GetCmdQuerySyncer implements the syncer query command
// TODO: Migrate to gRPC
func GetCmdQuerySyncer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "syncer",
		Short: "query syncer info",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			syncer, err := QuerySyncer(cliCtx)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			return cliCtx.PrintProto(&syncer)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current staking parameters information",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Query values set as staking parameters.

Example:
$ %s query staking params
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
