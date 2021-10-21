package cli

import (
	"github.com/spf13/cobra"

	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Transaction flags for the x/distribution module
var (
	FlagCommission       = "commission"
	FlagMaxMessagesPerTx = "max-msgs"
)

const (
	MaxMessagesPerTxDefault = 0
)

// NewTxCmd returns a root CLI command handler for all x/distribution transaction commands.
func NewTxCmd() *cobra.Command {
	distTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Distribution transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	distTxCmd.AddCommand()

	return distTxCmd
}

func ClaimAllStakingReward(
	t *transactor.Transactor, req *types.MsgClaimAllStakingReward) (resp *types.MsgClaimAllStakingRewardResponse, err error) {
	txResponse, err := t.SendTxMsgsWaitMined([]sdk.Msg{req})
	if err != nil {
		return resp, err
	}
	for _, log := range txResponse.Logs {
		for _, e := range log.Events {
			if e.Type == types.EventTypeClaimAllStakingReward {
				resp = new(types.MsgClaimAllStakingRewardResponse)
				break
			}
		}
		if resp != nil {
			break
		}
	}
	return resp, err
}
