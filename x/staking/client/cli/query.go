package cli

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	stakingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the staking module",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	stakingQueryCmd.AddCommand(common.GetCommands(
		GetCmdValidator(),
		GetCmdValidators(),
		GetCmdTransactors(),
		GetCmdDelegation(),
		GetCmdDelegations(),
		GetCmdSyncer(),
		GetCmdQueryParams(),
	)...)
	return stakingQueryCmd
}

func GetCmdValidator() *cobra.Command {
	return &cobra.Command{
		Use:   "validator [validator-eth-addr]",
		Short: "query validator info by validator ETH address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			validator, err := QueryValidator(cliCtx, args[0])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println(validator.YamlStr())
			return nil
		},
	}
}

func GetCmdValidators() *cobra.Command {
	return &cobra.Command{
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

			result, err := queryClient.Validators(cmd.Context(), &types.QueryValidatorsRequest{
				// Leaving status empty on purpose to query all validators.
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(result)
		},
	}
}

func GetCmdTransactors() *cobra.Command {
	return &cobra.Command{
		Use:   "transactors [validator-eth-addr]",
		Short: "query validator transactors validator ETH address",
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
}

func GetCmdDelegation() *cobra.Command {
	return &cobra.Command{
		Use:   "delegation [validator-eth-addr] [delegator-eth-addr]",
		Short: "query delegator info by validator and delegator ETH addresses",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			delegation, err := QueryDelegation(cliCtx, args[0], args[1])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			fmt.Println(delegation.YamlStr())
			return nil
		},
	}
}

// GetCmdDelegators queries request info
// TODO: support pagination
func GetCmdDelegations() *cobra.Command {
	return &cobra.Command{
		Use:   "delegations [validator-eth-addr]",
		Short: "query validator info by validator ETH address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			delegations, err := QueryDelegations(cliCtx, args[0])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}
			delegations.Sort()
			for _, delegation := range delegations {
				fmt.Println(delegation.YamlStr())
			}
			return nil
		},
	}
}

func GetCmdSyncer() *cobra.Command {
	return &cobra.Command{
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
}

func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Args:  cobra.NoArgs,
		Short: "Query the current staking parameters information",
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
