package cli

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/slash/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	slashTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Slash transaction subcommands",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	slashTxCmd.AddCommand(common.PostCommands()...)

	return slashTxCmd
}
