package ops

import (
	"github.com/celer-network/sgn-v2/common"
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
		ValidatorCommands(),
		DelegatorCommands(),
		SyncCommands(),
		BridgeCommands(),
		EthViewerCommand(),
	)

	return cmd
}

func ValidatorCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "validator",
		Short:                      "Validator subcommands (init, bond, collect fees, etc.)",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(common.PostCommands(
		InitValidatorCommand(),
		BondValidatorCommand(),
		ClaimValidatorStakingRewardCmd(),
		WithdrawValidatorCbrFeeCmd(),
		WithdrawValidatorPegbrFeeCmd(),
		ValidatorAddressCommand(),
	)...)

	return cmd
}

func DelegatorCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "delegator",
		Short:                      "Delegator subcommands (delegate, undelegate, etc.)",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(common.PostCommands(
		DelegateCommand(),
		UndelegateCommand(),
		CompleteUndelegateCommand(),
	)...)

	return cmd
}

func BridgeCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "bridge",
		Short:                      "Bridge subcommands (submit relay, sync farming, etc.)",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(common.PostCommands(
		SubmitRelayCmd(),
		SyncFarmingCommand(),
		TriggerSetRefundCommand(),
	)...)

	return cmd
}

func SyncCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "sync",
		Short:                      "Sync onchain (staking, bridge) states to sgnchain",
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(common.PostCommands(
		GetSyncSigners(),
		GetSyncCbrEvent(),
		GetSyncStaking(),
	)...)

	return cmd
}
