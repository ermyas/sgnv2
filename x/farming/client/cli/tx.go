package cli

import (
	"fmt"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/celer-network/sgn-v2/x/farming/types"
	govtypes "github.com/celer-network/sgn-v2/x/gov/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
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
		GetCmdSubmitAddTokensProposal(),
		GetCmdSubmitAdjustRewardProposal(),
		GetCmdSubmitSetRewardContractsProposal(),
		GetCmdClaimAllRewards(),
		// TODO: Support ClaimRewards for a single pool?
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
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func ClaimAllRewards(t *transactor.Transactor, req *types.MsgClaimAllRewards) (resp *types.MsgClaimAllRewardsResponse, err error) {
	txResponse, err := t.SendTxMsgsWaitMined([]sdk.Msg{req})
	if err != nil {
		return resp, err
	}
	for _, log := range txResponse.Logs {
		for _, e := range log.Events {
			if e.Type == types.EventTypeClaimAll {
				resp = new(types.MsgClaimAllRewardsResponse)
				break
			}
		}
		if resp != nil {
			break
		}
	}
	return resp, err
}

// GetCmdSubmitAddPoolProposal implements a command handler for submitting an AddPoolProposal
func GetCmdSubmitAddPoolProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "farming-add-pool [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an AddPoolProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit an AddPoolProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal farming-add-pool <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "Add cbridge-DAI/1 pool",
  "description": "add a CBridge farming pool for DAI on Ethereum",
  "pool_name": "cbridge-DAI/1",
  "stake_token": {
    "chain_id": 1,
    "symbol": "CB-DAI",
    "address": "0x6b175474e89094c44da98b954eedeac495271d0f",
    "decimals": 18
  },
  "reward_tokens": [
    {
      "chain_id": 1,
      "symbol": "CELR",
      "address": "0x4f9254c83eb525f9fcf346490bbb3ed28a81c667",
      "decimals": 18
    }
  ],
  "initial_reward_inputs": [
    {
      "add_amount": {
        "denom": "CELR/1",
        "amount": "100000000000000000000000"
      },
      "reward_start_block_delay": 0,
      "new_reward_amount_per_block": "100000000000000000"
    }
  ],
  "deposit": "10000000%s"
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

// GetCmdSubmitBatchAddPoolProposal implements a command handler for submitting a BatchAddPoolProposal
func GetCmdSubmitBatchAddPoolProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "farming-batch-add-pool [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a BatchAddPoolProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a BatchAddPoolProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal farming-batch-add-pool <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:
{
  "title": "Batch Add cbridge-DAI/1 pool",
  "description": "batch add a CBridge farming pool for DAI on Ethereum",
  "add_pool_infos": [
    {
      "pool_name": "cbridge-DAI/1",
      "stake_token": {
        "chain_id": 1,
        "symbol": "CB-DAI",
        "address": "0x6b175474e89094c44da98b954eedeac495271d0f",
        "decimals": 18
      },
      "reward_tokens": [
        {
          "chain_id": 1,
          "symbol": "CELR",
          "address": "0x4f9254c83eb525f9fcf346490bbb3ed28a81c667",
          "decimals": 18
        }
      ],
      "initial_reward_inputs": [
        {
          "add_amount": {
            "denom": "CELR/1",
            "amount": "100000000000000000000000"
          },
          "reward_start_block_delay": 0,
          "new_reward_amount_per_block": "100000000000000000"
        }
      ]
    }
  ],
  "deposit": "10000000%s"
}
`, version.AppName, stakingtypes.StakeDenom,
			)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseBatchAddPoolProposalWithDeposit(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			content :=
				types.NewBatchAddPoolProposal(
					proposal.Title,
					proposal.Description,
					proposal.AddPoolInfos,
				)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit[0].Amount, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdSubmitAddTokensProposal implements a command handler for submitting an AddTokensProposal
func GetCmdSubmitAddTokensProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "farming-add-tokens [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an AddTokensProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit an AddTokensProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal farming-add-tokens <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "Add USDT/1 and USDC/5",
  "description": "Add USDT on Ethereum and USDC on Goerli",
  "tokens": [
    {
      "chain_id": 1,
      "symbol": "USDT",
      "address": "0xdac17f958d2ee523a2206206994597c13d831ec7",
      "decimals": 6
    },
    {
      "chain_id": 5,
      "symbol": "USDC",
      "address": "0xCbE56b00d173A26a5978cE90Db2E33622fD95A28",
      "decimals": 6
    }
  ],
  "deposit": "10000000%s"
}
`, version.AppName, stakingtypes.StakeDenom,
			)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseAddTokensProposalWithDeposit(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			content :=
				types.NewAddTokensProposal(
					proposal.Title,
					proposal.Description,
					proposal.Tokens,
				)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit[0].Amount, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdSubmitAdjustRewardProposal implements a command handler for submitting an AdjustRewardProposal
func GetCmdSubmitAdjustRewardProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "farming-adjust-reward [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an AdjustRewardProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit an AdjustRewardProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal farming-adjust-reward <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "cbridge-DAI/1 reward adjustment",
  "description": "Add DAI reward for cbridge-DAI/1 and adjust CELR reward",
  "pool_name": "cbridge-DAI/1",
  "reward_adjustment_inputs": [
    {
      "add_amount": {
        "denom": "CELR/1",
        "amount": "100000000000000000000000"
      },
      "reward_start_block_delay": 0,
      "new_reward_amount_per_block": "500000000000000000"
    },
    {
      "add_amount": {
        "denom": "USDT/1",
        "amount": "100000000000"
      },
      "reward_start_block_delay": 3,
      "new_reward_amount_per_block": "1000000"
    }
  ],
  "deposit": "10000000%s"
}
`, version.AppName, stakingtypes.StakeDenom,
			)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseAdjustRewardProposalWithDeposit(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			content :=
				types.NewAdjustRewardProposal(
					proposal.Title,
					proposal.Description,
					proposal.PoolName,
					proposal.RewardAdjustmentInputs,
					proposal.RemoveDuplicates,
				)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit[0].Amount, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdSubmitBatchAdjustRewardProposal implements a command handler for submitting a BatchAdjustRewardProposal
func GetCmdSubmitBatchAdjustRewardProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "farming-batch-adjust-reward [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an batchAdjustRewardProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit a BatchAdjustRewardProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal farming-batch-adjust-reward <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:
{
  "title": "cbridge-DAI/1 reward adjustment",
  "description": "Add DAI reward for cbridge-DAI/1 and adjust CELR reward",
  "adjust_reward_infos": [
    {
      "pool_name": "cbridge-DAI/1",
      "reward_adjustment_inputs": [
        {
          "add_amount": {
            "denom": "CELR/1",
            "amount": "100000000000000000000000"
          },
          "reward_start_block_delay": 0,
          "new_reward_amount_per_block": "500000000000000000"
        },
        {
          "add_amount": {
            "denom": "USDT/1",
            "amount": "100000000000"
          },
          "reward_start_block_delay": 3,
          "new_reward_amount_per_block": "1000000"
        }
      ]
    }
  ],
  "deposit": "10000000%s"
}
`, version.AppName, stakingtypes.StakeDenom,
			)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseBatchAdjustRewardProposalWithDeposit(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			content :=
				types.NewBatchAdjustRewardProposal(
					proposal.Title,
					proposal.Description,
					proposal.AdjustRewardInfos,
				)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit[0].Amount, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}

// GetCmdSubmitSetRewardContractsProposal implements a command handler for submitting an SetRewardContractsProposal
func GetCmdSubmitSetRewardContractsProposal() *cobra.Command {
	return &cobra.Command{
		Use:   "farming-set-reward-contracts [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit an SetRewardContractsProposal",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Submit an SetRewardContractsProposal along with an initial deposit.
The proposal details must be supplied via a JSON file.

Example:
$ %s gov submit-proposal farming-set-reward-contracts <path/to/proposal.json> --from=<key_or_address>

Where proposal.json contains:

{
  "title": "Set FarmingRewardsContract for chain IDs 1 and 5",
  "description": "Set FarmingRewardsContract for Ethereum and Goerli",
  "reward_contracts": [
    {
      "chain_id": 1,
      "address": "0x0db9888166f26F0416977E6D4B41Cf8C266c942c",
    },
    {
      "chain_id": 5,
      "address": "0x0db9888166f26F0416977E6D4B41Cf8C266c942c",
    }
  ],
  "deposit": "10000000%s"
}
`, version.AppName, stakingtypes.StakeDenom,
			)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			proposal, err := ParseSetRewardContractsProposalWithDeposit(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}
			deposit, err := sdk.ParseCoinsNormalized(proposal.Deposit)
			if err != nil {
				return err
			}
			from := clientCtx.GetFromAddress()
			content :=
				types.NewSetRewardContractsProposal(
					proposal.Title,
					proposal.Description,
					proposal.RewardContracts,
				)

			msg, err := govtypes.NewMsgSubmitProposal(content, deposit[0].Amount, from)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
