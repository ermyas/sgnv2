package cli

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagTransactorOp = "op"
	flagMoniker      = "moniker"
	flagIdentity     = "identity"
	flagWebsite      = "website"
	flagContact      = "contact"
	flagDetails      = "details"
)

func GetTxCmd() *cobra.Command {
	validatorTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Validator transaction subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	validatorTxCmd.AddCommand(common.PostCommands(
		GetCmdSetTransactors(),
		GetCmdEditDescription(),
	)...)

	return validatorTxCmd
}

// GetCmdSetTransactors is the CLI command for sending a SetTransactors transaction
func GetCmdSetTransactors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-transactors",
		Short: "set transactors based on transactors in config",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			/*
				var op types.SetTransactorsOp
				opstr, err := cmd.Flags().GetString(flagTransactorOp)
				if err != nil {
					return err
				}
				if opstr == "overwrite" {
					op = types.SetTransactorsOp_Overwrite
				} else if opstr == "add" {
					op = types.SetTransactorsOp_Add
				} else if opstr == "remove" {
					op = types.SetTransactorsOp_Remove
				} else {
					return fmt.Errorf("invalid op, should be one of overwrite | add | remove ")
				}
			*/
			transactors := viper.GetStringSlice(common.FlagSgnTransactors)
			home, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}
			txr, err := transactor.NewCliTransactor(
				home, clientCtx.LegacyAmino, clientCtx.Codec, clientCtx.InterfaceRegistry)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetTransactors(types.SetTransactorsOp_Overwrite, transactors, txr.Key.GetAddress().String())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			txr.CliSendTxMsgWaitMined(&msg)
			return nil
		},
	}
	//cmd.Flags().String(flagTransactorOp, "overwrite", "operation (overwrite | add | remove)")

	return cmd
}

func GetCmdEditDescription() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-description",
		Short: "Edit validator description",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			moniker, err := cmd.Flags().GetString(flagMoniker)
			if err != nil {
				return err
			}
			identity, err := cmd.Flags().GetString(flagIdentity)
			if err != nil {
				return err
			}
			website, err := cmd.Flags().GetString(flagWebsite)
			if err != nil {
				return err
			}
			contact, err := cmd.Flags().GetString(flagContact)
			if err != nil {
				return err
			}
			details, err := cmd.Flags().GetString(flagDetails)
			if err != nil {
				return err
			}
			description := types.NewDescription(moniker, identity, website, contact, details)
			home, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}
			txr, err := transactor.NewCliTransactor(
				home, clientCtx.LegacyAmino, clientCtx.Codec, clientCtx.InterfaceRegistry)
			if err != nil {
				return err
			}

			msg := types.NewMsgEditDescription(description, txr.Key.GetAddress().String())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			txr.CliSendTxMsgWaitMined(&msg)

			return nil
		},
	}

	cmd.Flags().String(flagMoniker, types.DoNotModifyDesc, "The validator's name")
	cmd.Flags().String(flagIdentity, types.DoNotModifyDesc, "The validator's identity")
	cmd.Flags().String(flagWebsite, types.DoNotModifyDesc, "The validator's website")
	cmd.Flags().String(flagContact, types.DoNotModifyDesc, "The validator's contact email")
	cmd.Flags().String(flagDetails, types.DoNotModifyDesc, "The validator's details")

	return cmd
}
