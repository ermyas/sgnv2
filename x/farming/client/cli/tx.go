package cli

import (
	"fmt"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/farming/types"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

// NewTxCmd returns a root CLI command handler for all x/farming transaction commands.
func NewTxCmd() *cobra.Command {
	farmingTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	farmingTxCmd.AddCommand(
		GetCmdSubmitAddPoolProposal(),
		// TODO: Support the rest of the proposal types
		GetCmdClaimAllRewards(),
		// TODO: Support ClaimRewards for a single pool
	)
	return farmingTxCmd
}

func GetCmdClaimAllRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-all [address]",
		Short: "claim farming rewards from all staked pools",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Claim farming rewards.

Example:
$ %s tx farming claim-all 0xab5801a7d398351b8be11c439e05c5b3259aec9b --from mykey
`, version.AppName),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			senderAddr := clientCtx.GetFromAddress()
			addr := eth.Hex2Addr(args[0])

			msg := types.NewMsgClaimAllRewards(addr, senderAddr)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	return cmd
}

// GetCmdSubmitAddPoolProposal implements a command handler for submitting an AddPoolProposal
func GetCmdSubmitAddPoolProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "add-pool [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an AddPoolProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit an AddPoolProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal add-pool <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
 "title": "add a farming pool",
 "description": "add a CBridge farming pool for DAI on Ethereum",
 "pool_name": "cbridge-CB-DAI/1",
 "stake_token": {
   "chain_id": 1,
   "symbol": "CB-DAI",
   "address": "0x6b175474e89094c44da98b954eedeac495271d0f",
 },
 "reward_tokens": [
   {
     "chain_id": 1,
     "symbol": "CELR",
     "address": "0x4f9254c83eb525f9fcf346490bbb3ed28a81c667",
   }
 ],
 "initial_reward_inputs": [
   {
     "add_amount": {
       "denom": "CELR/1",
       "amount": "100000000000000000000000"
	 },
     "reward_start_block_delay": 8640,
     "new_reward_amount_per_block": "100000000000000000"
   }
 ],
 "deposit": [
   {
     "denom": "%s",
     "amount": "100"
   }
 ]
}
`, version.AppName, stakingtypes.StakeDenom,
			)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseAddPoolProposalWithDeposit(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			content :=
				types.NewAddPoolProposal(
					proposal.Title,
					proposal.Description,
					proposal.PoolName,
					proposal.StakeToken,
					proposal.RewardTokens,
					proposal.InitialRewardInputs,
				)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit[0].Amount, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
