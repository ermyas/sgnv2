package cli

import (
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/transactor"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/spf13/cobra"
)

const (
	// TimeFormat specifies ISO UTC format for submitting the time for a new upgrade proposal
	TimeFormat = "2006-01-02T15:04:05Z"

	FlagUpgradeHeight = "upgrade-height"
	FlagUpgradeTime   = "time"
	FlagUpgradeInfo   = "info"
)

func parseArgsToContent(cmd *cobra.Command, name string) (govtypes.Content, error) {
	title, err := cmd.Flags().GetString(FlagTitle)
	if err != nil {
		return nil, err
	}

	description, err := cmd.Flags().GetString(FlagDescription)
	if err != nil {
		return nil, err
	}

	height, err := cmd.Flags().GetInt64(FlagUpgradeHeight)
	if err != nil {
		return nil, err
	}

	timeStr, err := cmd.Flags().GetString(FlagUpgradeTime)
	if err != nil {
		return nil, err
	}

	if height != 0 && len(timeStr) != 0 {
		return nil, fmt.Errorf("only one of --upgrade-time or --upgrade-height should be specified")
	}

	var upgradeTime time.Time
	if len(timeStr) != 0 {
		upgradeTime, err = time.Parse(TimeFormat, timeStr)
		if err != nil {
			return nil, err
		}
	}

	info, err := cmd.Flags().GetString(FlagUpgradeInfo)
	if err != nil {
		return nil, err
	}

	plan := upgradetypes.Plan{Name: name, Time: upgradeTime, Height: height, Info: info}
	content := govtypes.NewUpgradeProposal(title, description, plan)
	return content, nil
}

// GetCmdSubmitUpgradeProposal implements a command handler for submitting a software upgrade proposal transaction.
func GetCmdSubmitUpgradeProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "software-upgrade [name] (--upgrade-height [height] | --upgrade-time [time]) (--upgrade-info [info]) [flags]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a software upgrade proposal",
		Long: "Submit a software upgrade along with an initial deposit.\n" +
			"Please specify a unique name and height OR time for the upgrade to take effect.\n" +
			"You may include info to reference a binary download link, in a format compatible with: https://github.com/regen-network/cosmosd",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			content, err := parseArgsToContent(cmd, name)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			home, err := cmd.Flags().GetString(flags.FlagHome)
			if err != nil {
				return err
			}
			txr, err := transactor.NewCliTransactor(home, clientCtx.LegacyAmino, clientCtx.Codec, clientCtx.InterfaceRegistry)
			if err != nil {
				log.Error(err)
				return err
			}

			deposit, err := cmd.Flags().GetUint64(FlagDeposit)
			if err != nil {
				return err
			}

			msg, _ := govtypes.NewMsgSubmitProposal(content, sdk.NewIntFromUint64(deposit), txr.Key.GetAddress())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			txr.CliSendTxMsgWaitMined(msg)

			return nil
		},
	}

	cmd.Flags().String(FlagTitle, "", "title of proposal")
	cmd.Flags().String(FlagDescription, "", "description of proposal")
	cmd.Flags().Uint64(FlagDeposit, 0, "deposit of proposal")
	cmd.Flags().Int64(FlagUpgradeHeight, 0, "The height at which the upgrade must happen (not to be used together with --upgrade-time)")
	cmd.Flags().String(FlagUpgradeTime, "", fmt.Sprintf("The time at which the upgrade must happen (ex. %s) (not to be used together with --upgrade-height)", TimeFormat))
	cmd.Flags().String(FlagUpgradeInfo, "", "Optional info for the planned upgrade such as commit hash, etc.")

	return cmd
}
