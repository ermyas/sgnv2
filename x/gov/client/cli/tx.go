package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
)

// Proposal flags
const (
	FlagTitle        = "title"
	FlagDescription  = "description"
	FlagProposalType = "type"
	FlagDeposit      = "deposit"
	flagVoter        = "voter"
	flagDepositor    = "depositor"
	flagStatus       = "status"
	FlagProposal     = "proposal"
)

type proposal struct {
	Title       string
	Description string
	Type        string
	Deposit     string
}

// ProposalFlags defines the core required fields of a proposal. It is used to
// verify that these values are not provided in conjunction with a JSON proposal
// file.
var ProposalFlags = []string{
	FlagTitle,
	FlagDescription,
	FlagProposalType,
	FlagDeposit,
}

// NewTxCmd returns the transaction commands for this module
// governance ModuleClient is slightly different from other ModuleClients in that
// it contains a slice of "proposal" child commands. These commands are respective
// to proposal type handlers that are implemented in other modules but are mounted
// under the governance CLI (eg. parameter change proposals).
func NewTxCmd(propCmds []*cobra.Command) *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Governance transactions subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmdSubmitProp := NewCmdSubmitProposal()
	for _, propCmd := range propCmds {
		flags.AddTxFlagsToCmd(propCmd)
		cmdSubmitProp.AddCommand(propCmd)
	}

	govTxCmd.AddCommand(
		NewCmdDeposit(),
		NewCmdVote(),
		cmdSubmitProp,
	)

	return govTxCmd
}

// NewCmdSubmitProposal implements submitting a proposal transaction command.
func NewCmdSubmitProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal",
		Short: "Submit a proposal along with an initial deposit",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a proposal along with an initial deposit.
Proposal title, description, type and deposit can be given directly or through a proposal JSON file.

Example:
$ %s tx gov submit-proposal --proposal="path/to/proposal.json" --from mykey

Where proposal.json contains:

{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10test"
}

Which is equivalent to:

$ %s tx gov submit-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10test" --from mykey
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, err := parseSubmitProposalFlags(cmd.Flags())
			if err != nil {
				return fmt.Errorf("failed to parse proposal: %w", err)
			}

			amount, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			if len(amount) != 1 {
				return errors.New("only support a single amount")
			}

			content := types.ContentFromProposalType(proposal.Title, proposal.Description, proposal.Type)

			// TODO: Check amount
			msg, err := types.NewMsgSubmitProposal(content, amount[0].Amount, clientCtx.GetFromAddress())
			if err != nil {
				return fmt.Errorf("invalid message: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagTitle, "", "The proposal title")
	cmd.Flags().String(FlagDescription, "", "The proposal description")
	cmd.Flags().String(FlagProposalType, "", "The proposal Type")
	cmd.Flags().String(FlagDeposit, "", "The proposal deposit")
	cmd.Flags().String(FlagProposal, "", "Proposal file path (if this path is given, other proposal flags are ignored)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewCmdDeposit implements depositing tokens for an active proposal.
func NewCmdDeposit() *cobra.Command {
	return &cobra.Command{
		Use:   "deposit [proposal-id] [deposit]",
		Args:  cobra.ExactArgs(2),
		Short: "Deposit tokens for an active proposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a deposit for an active proposal. You can
find the proposal-id by running "%s query gov proposals".

Example:
$ %s tx gov deposit 1 10
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid uint, please input a valid proposal-id", args[0])
			}

			// Get amount of coins
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return err
			}

			msg := types.NewMsgDeposit(txr.Key.GetAddress(), proposalID, amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			txr.CliSendTxMsgWaitMined(msg)

			return nil
		},
	}
}

// NewCmdVote implements creating a new vote command.
func NewCmdVote() *cobra.Command {
	return &cobra.Command{
		Use:   "vote [proposal-id] [option]",
		Args:  cobra.ExactArgs(2),
		Short: "Vote for an active proposal, options: yes/no/no_with_veto/abstain",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a vote for an active proposal. You can
find the proposal-id by running "%s query gov proposals".


Example:
$ %s tx gov vote 1 yes
`,
				version.AppName, version.AppName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
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

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
			}

			// Find out which vote option user chose
			byteVoteOption, err := types.VoteOptionFromString(args[1])
			if err != nil {
				return err
			}

			// Build vote message and run basic validation
			msg := types.NewMsgVote(txr.Key.GetAddress(), proposalID, byteVoteOption)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			txr.CliSendTxMsgWaitMined(msg)

			return nil
		},
	}
}
