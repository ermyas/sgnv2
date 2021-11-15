package cli

import (
	"fmt"
	"strings"

	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	"github.com/celer-network/sgn-v2/x/mint/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

// NewTxCmd returns a root CLI command handler for all x/mint transaction commands.
func NewTxCmd() *cobra.Command {
	mintTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	mintTxCmd.AddCommand(
		GetCmdSubmitAdjustProvisionsProposal(),
	)
	return mintTxCmd
}

// GetCmdSubmitAdjustProvisionsProposal implements a command handler for submitting an AddTokensProposal
func GetCmdSubmitAdjustProvisionsProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "mint-adjust-provisions [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an AdjustProvisionsProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit an AdjustProvisionsProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal mint-adjust-provisions <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "Adjust CELR/reward provisions",
  "description": "Adjust CELR staking reward provisions",
  "new_annual_provisions": "30000000000000000000000000",
  "deposit": "10000000%s"
}
`, version.AppName, stakingtypes.StakeDenom,
			)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseAdjustProvisionsProposalWithDeposit(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			content :=
				types.NewAdjustProvisionsProposal(
					proposal.Title,
					proposal.Description,
					proposal.NewAnnualProvisions,
				)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit[0].Amount, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
