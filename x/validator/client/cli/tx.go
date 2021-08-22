package cli

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagMoniker = "moniker"
	flagWebsite = "website"
	flagContact = "contact"
	flagDetails = "details"
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
				transactors, err := common.ParseTransactorAddrs(viper.GetStringSlice(common.FlagSgnTransactors))
				if err != nil {
					return err
				}*/
			transactors := viper.GetStringSlice(common.FlagSgnTransactors)

			txr, err := transactor.NewCliTransactor(clientCtx.Codec, viper.GetString(flags.FlagHome))
			if err != nil {
				return err
			}

			msg := types.NewMsgSetTransactors(transactors, txr.Key.GetAddress().String())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			//TODO: txr.CliSendTxMsgWaitMined(msg)

			return nil
		},
	}

	return cmd
}
