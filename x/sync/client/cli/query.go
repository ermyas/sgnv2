package cli

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/sync/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func GetQueryCmd() *cobra.Command {
	syncQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the sync module",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	syncQueryCmd.AddCommand(common.GetCommands()...)
	return syncQueryCmd
}
