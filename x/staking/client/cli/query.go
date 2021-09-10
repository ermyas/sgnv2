package cli

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	validatorQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the validator module",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	validatorQueryCmd.AddCommand(common.GetCommands(
		GetCmdValidator(),
		GetCmdValidators(),
		GetCmdDelegator(),
		GetCmdDelegators(),
		GetCmdSyncer(),
		GetCmdQueryParams(),
	)...)
	return validatorQueryCmd
}

// GetCmdSyncer queries syncer info
func GetCmdSyncer() *cobra.Command {
	return &cobra.Command{
		Use:   "syncer",
		Short: "query syncer info",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, _ := client.GetClientTxContext(cmd)
			cliCtx := common.NewQueryCLIContext(&clientCtx.Codec)
			syncer, err := QuerySyncer(cliCtx)
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(&syncer)
		},
	}
}

// GetCmdDelegator queries request info
func GetCmdDelegator() *cobra.Command {
	return &cobra.Command{
		Use:   "delegator [validator-eth-addr] [delegator-eth-addr]",
		Short: "query delegator info by validator and delegator ETH addresses",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, _ := client.GetClientTxContext(cmd)
			cliCtx := common.NewQueryCLIContext(&clientCtx.Codec)
			delegator, err := QueryDelegator(cliCtx, args[0], args[1])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(delegator)
		},
	}
}

// GetCmdValidator queries request info
func GetCmdValidator() *cobra.Command {
	return &cobra.Command{
		Use:   "validator [validator-eth-addr]",
		Short: "query validator info by validator ETH address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, _ := client.GetClientTxContext(cmd)
			cliCtx := common.NewQueryCLIContext(&clientCtx.Codec)
			validator, err := QueryValidator(cliCtx, args[0])
			if err != nil {
				log.Errorln("query error", err)
				return err
			}

			return cliCtx.PrintProto(validator)
		},
	}
}

func GetCmdValidators() *cobra.Command {
	return &cobra.Command{}
}

// GetCmdDelegators queries request info
// TODO: support pagination
func GetCmdDelegators() *cobra.Command {
	return &cobra.Command{}
}

// GetCmdQueryParams implements the params query command.
func GetCmdQueryParams() *cobra.Command {
	return &cobra.Command{}
}

// ----------------------- CLI print-friendly output --------------------
