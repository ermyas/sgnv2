package ops

import (
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func OpsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "ops",
		Short:                      "Operation subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		InitValidatorCommand(),
		BondValidatorCommand(),
		DelegateCommand(),
		UndelegateCommand(),
		CompleteUndelegateCommand(),
		tc.DeployCommand(),
		GetSyncCmd(),
		EthViewerCommand(),
	)

	return cmd
}
