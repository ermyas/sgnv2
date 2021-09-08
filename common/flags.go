package common

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

const (
	FlagEthGateway           = "eth.gateway"
	FlagEthContractCelr      = "eth.contract_address.celr"
	FlagEthContractStaking   = "eth.contract_address.staking"
	FlagEthContractSgn       = "eth.contract_address.sgn"
	FlagEthContractReward    = "eth.contract_address.reward"
	FlagEthContractViewer    = "eth.contract_address.viewer"
	FlagEthContractGovern    = "eth.contract_address.govern"
	FlagEthSignerKeystore    = "eth.signer_keystore"
	FlagEthSignerPassphrase  = "eth.signer_passphrase"
	FlagEthValidatorAddress  = "eth.validator_address"
	FlagEthPollInterval      = "eth.poll_interval"
	FlagEthSyncBlkInterval   = "eth.sync_blk_interval"
	FlagEthAcceptedBlkRange  = "eth.accepted_blk_range"
	FlagEthBlockDelay        = "eth.block_delay"
	FlagEthChainId           = "eth.chain_id"
	FlagEthCheckInterval     = "eth.check_interval"
	FlagEthMonitorStartBlock = "eth.monitor_start_block"
	FlagEthMinGasPriceGwei   = "eth.min_gas_price_gwei"
	FlagEthAddGasPriceGwei   = "eth.add_gas_price_gwei"
	FlagEthMaxBlockDelta     = "eth.max_block_delta"

	FlagSgnValidatorAccount = "sgnd.validator_account"
	FlagSgnTransactors      = "sgnd.transactors"
	FlagSgnPassphrase       = "sgnd.passphrase"
	FlagSgnChainId          = "sgnd.chain_id"
	FlagSgnNodeURI          = "sgnd.node_uri"
	FlagSgnBaseGasPrice     = "sgnd.base_gas_price"
	FlagSgnTimeoutCommit    = "sgnd.timeout_commit"
	FlagSgnKeyringBackend   = "sgnd.keyring_backend"
	FlagSgnGasAdjustment    = "sgnd.gas_adjustment"

	FlagSgnCheckIntervalSlashQueue = "sgnd.check_interval.slash_queue"

	FlagLogLevel = "log.level"
	FlagLogColor = "log.color"
)

const (
	DefaultSgnGasAdjustment = 1.5
	DefaultSgnGasLimit      = 300000
)

func PostCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		// c.Flags().Bool(flags.FlagIndentResponse, false, "Add indent to JSON response")
		// c.Flags().Bool(flags.FlagTrustNode, true, "Trust connected full node (don't verify proofs for responses)")

		// viper.BindPFlag(flags.FlagTrustNode, c.Flags().Lookup(flags.FlagTrustNode))

		c.SetErr(c.ErrOrStderr())
	}
	return cmds
}

func GetCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		// c.Flags().Bool(flags.FlagIndentResponse, false, "Add indent to JSON response")
		// c.Flags().Bool(flags.FlagTrustNode, false, "Trust connected full node (don't verify proofs for responses)")
		c.Flags().Int64(flags.FlagHeight, 0, "Use a specific height to query state at (this can error if the node is pruning state)")

		//viper.BindPFlag(flags.FlagTrustNode, c.Flags().Lookup(flags.FlagTrustNode))

		c.SetErr(c.ErrOrStderr())
	}
	return cmds
}
