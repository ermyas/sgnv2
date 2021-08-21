package cli

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/validator/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	flagMoniker = "moniker"
	flagWebsite = "website"
	flagContact = "contact"
	flagDetails = "details"
)

func GetTxCmd(storeKey string, cdc codec.Codec) *cobra.Command {
	validatorTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Validator transaction subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	validatorTxCmd.AddCommand(common.PostCommands(
		GetCmdSetTransactors(cdc),
	)...)

	return validatorTxCmd
}

// GetCmdSetTransactors is the CLI command for sending a SetTransactors transaction
func GetCmdSetTransactors(cdc codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-transactors",
		Short: "set transactors based on transactors in config",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			/*
				transactors, err := common.ParseTransactorAddrs(viper.GetStringSlice(common.FlagSgnTransactors))
				if err != nil {
					return err
				}*/
			transactors := viper.GetStringSlice(common.FlagSgnTransactors)

			txr, err := transactor.NewCliTransactor(cdc, viper.GetString(flags.FlagHome))
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
