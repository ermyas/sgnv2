package common

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

const (
	FlagEthGateway                = "eth.gateway"
	FlagEthContractCelr           = "eth.contract_address.celr"
	FlagEthContractStaking        = "eth.contract_address.staking"
	FlagEthContractSgn            = "eth.contract_address.sgn"
	FlagEthContractStakingReward  = "eth.contract_address.staking_reward"
	FlagEthContractFarmingRewards = "eth.contract_address.farming_rewards"
	FlagEthContractViewer         = "eth.contract_address.viewer"
	FlagEthContractGovern         = "eth.contract_address.govern"
	FlagEthSignerKeystore         = "eth.signer_keystore"
	FlagEthSignerPassphrase       = "eth.signer_passphrase"
	FlagEthValidatorAddress       = "eth.validator_address"
	FlagEthPollInterval           = "eth.poll_interval"
	FlagEthBlockDelay             = "eth.block_delay"
	FlagEthChainId                = "eth.chain_id"
	FlagEthCheckInterval          = "eth.check_interval"
	FlagEthMonitorStartBlock      = "eth.monitor_start_block"
	FlagEthMinGasPriceGwei        = "eth.min_gas_price_gwei"
	FlagEthAddGasPriceGwei        = "eth.add_gas_price_gwei"
	FlagEthMaxBlockDelta          = "eth.max_block_delta"

	FlagSgnValidatorAccount = "sgnd.validator_account"
	FlagSgnTransactors      = "sgnd.transactors"
	FlagSgnPassphrase       = "sgnd.passphrase"
	FlagSgnChainId          = "sgnd.chain_id"
	FlagSgnNodeURI          = "sgnd.node_uri"
	FlagSgnBaseGasPrice     = "sgnd.base_gas_price"
	FlagSgnPriceUpdateUrl   = "sgnd.price_update_url"
	FlagSgnKeyringBackend   = "sgnd.keyring_backend"
	FlagSgnGasAdjustment    = "sgnd.gas_adjustment"

	FlagSgnCheckIntervalSlash    = "sgnd.check_interval.slash"
	FlagSgnCheckIntervalCbridge  = "sgnd.check_interval.cbridge"
	FlagSgnCheckIntervalVerifier = "sgnd.check_interval.verifier"

	FlagConsensusTimeoutCommit = "consensus.timeout_commit"

	FlagLogLevel = "log.level"
	FlagLogColor = "log.color"

	FlagToStartGateway     = "gateway.start_gateway"
	FlagGatewayDbUrl       = "gateway.db_url"
	FlagGatewayAwsS3Region = "gateway.aws.s3.region"
	FlagGatewayAwsS3Bucket = "gateway.aws.s3.bucket"
	FlagGatewayAwsKey      = "gateway.aws.key"
	FlagGatewayAwsSecret   = "gateway.aws.secret"
	FlagBlockNativeApiKey  = "gateway.block_native.apikey"

	FlagMultiChain = "multichain" // array of toml tables, each table represents one chain, see common/multichain.go for details
)

const (
	DefaultSgnGasAdjustment = 1.5
	DefaultSgnGasLimit      = 300000
)

func PostCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		c.SetErr(c.ErrOrStderr())
	}
	return cmds
}

func GetCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		c.Flags().Int64(flags.FlagHeight, 0, "Use a specific height to query state at (this can error if the node is pruning state)")

		c.SetErr(c.ErrOrStderr())
	}
	return cmds
}
