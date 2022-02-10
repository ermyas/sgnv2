package cli

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
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

	distTxCmd.AddCommand(common.PostCommands(
		GetCmdClaimAllStakingReward(),
	)...)

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

func GetCmdClaimAllStakingReward() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-staking-reward [delegator-address]",
		Short: "Claim all staking reward of a delegator or a validator",
		Args:  cobra.ExactArgs(1),
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
				return err
			}
			msg := &types.MsgClaimAllStakingReward{
				DelegatorAddress: args[0],
				Sender:           txr.Key.GetAddress().String(),
			}
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			_, err = ClaimAllStakingReward(txr, msg)
			if err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
