package multinode

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/spf13/viper"
)

var mainchain2Started bool

func BuildDockers() {
	tc.RunCmd("make", "localnet-down")
	tc.RunCmd("make", "build-node")
	tc.RunCmd("make", "build-linux")
	tc.RunCmd("make", "build-geth")
}

func SetupMainchain() {
	tc.RunCmd("make", "prepare-geth-data")
	tc.RunCmd("make", "localnet-start-geth")
	tc.SleepWithLog(5, "geth start")

	// set up mainchain: deploy contracts, fund addrs, etc
	log.Infoln("fund each test addr 100 ETH")
	err := tc.FundAddrsETH(tc.Addrs, tc.NewBigInt(1, 20), tc.LocalGeth, int64(tc.ChainID))
	tc.ChkErr(err, "fund each test addr 100 ETH")

	log.Infoln("set up mainchain")
	tc.SetupEthClients()
	tc.CelrAddr, tc.CelrContract = tc.DeployERC20Contract(tc.EthClient, tc.EtherBaseAuth, "Celer", "CELR", 18)

	// fund CELR to each eth account
	log.Infoln("fund each validator and delegator addr 1 billion CELR")
	err = tc.FundAddrsErc20(tc.CelrAddr, tc.ValDelAddrs, tc.NewBigInt(1, 27), tc.EthClient, tc.EtherBaseAuth)
	tc.ChkErr(err, "fund each validator and delegator addr 1 billion CELR")
}

// should be invoked after mainchain 1 setup
func SetupMainchain2ForBridge() {
	if mainchain2Started {
		log.Infoln("mainchain2 already started")
		return
	}
	tc.RunCmd("make", "prepare-geth2-env")
	tc.RunCmd("make", "localnet-start-geth2")
	tc.SleepWithLog(5, "geth2 start")

	// set up mainchain: deploy contracts, fund addrs, etc
	log.Infoln("fund each test addr 100 ETH")
	err := tc.FundAddrsETH(tc.Addrs2, tc.NewBigInt(1, 20), tc.LocalGeth2, int64(tc.Geth2ChainID))
	tc.ChkErr(err, "fund each test addr 100 ETH")

	log.Infoln("set up mainchain2")
	tc.InitCbrChainConfigs()
	mainchain2Started = true
}

func SetupNewSgnEnv(contractParams *tc.ContractParams, cbridge, msg, manual, report bool) {
	tc.RunAllAndWait(
		SetupMainchain2ForBridge,
		func() {
			deployContractsAndPrepareSgnData(contractParams, cbridge, msg, manual, report)
		},
	)

	if cbridge {
		DeployUsdtForBridge()
		DeployBridgeContract()
		DeployPegBridgeContract()
		CreateFarmingPools()
		FundUsdtFarmingReward()
	}
	if msg {
		DeployBatchTransferAndMessageTransferAndMessageBusContracts()
	}

	// Update global viper
	node0ConfigPath := "../../../docker-volumes/node0/sgnd/config/sgn.toml"
	viper.SetConfigFile(node0ConfigPath)
	err := viper.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")

	tc.RunCmd("make", "localnet-up-nodes")
	if msg {
		PrepareExecutor()
	}
}

func deployContractsAndPrepareSgnData(contractParams *tc.ContractParams, cbridge, msg, manual, report bool) {
	log.Infoln("Deploy Staking and SGN contracts")
	if contractParams == nil {
		contractParams = &tc.ContractParams{
			CelrAddr:              tc.CelrAddr,
			ProposalDeposit:       big.NewInt(1e18),
			VotePeriod:            big.NewInt(10),
			UnbondingPeriod:       big.NewInt(50),
			MaxBondedValidators:   big.NewInt(7),
			MinValidatorTokens:    big.NewInt(2e18),
			MinSelfDelegation:     big.NewInt(1e18),
			AdvanceNoticePeriod:   big.NewInt(1),
			ValidatorBondInterval: big.NewInt(0),
			MaxSlashFactor:        big.NewInt(1e5),
		}
	}
	tx := tc.DeploySgnStakingContracts(contractParams)
	tc.WaitMinedWithChk(context.Background(), tc.EthClient, tx, tc.BlockDelay, tc.PollingInterval, "DeploySgnStakingContracts")

	tc.RunCmd("make", "localnet-down-nodes")
	err := tc.RunCmdNoChkErr("make", "prepare-sgn-data")
	if err != nil {
		tc.RunCmd("make", "prepare-sgn-data-sudo")
	}

	log.Infoln("Updating config files of SGN nodes")
	for i := 0; i < len(tc.ValEthKs); i++ {
		configPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/sgn.toml", i)
		configFileViper := viper.New()
		configFileViper.SetConfigFile(configPath)
		err := configFileViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		configFileViper.Set(common.FlagEthContractCelr, tc.CelrAddr.Hex())
		configFileViper.Set(common.FlagEthContractStaking, tc.Contracts.Staking.Address.Hex())
		configFileViper.Set(common.FlagEthContractSgn, tc.Contracts.Sgn.Address.Hex())
		configFileViper.Set(common.FlagEthContractStakingReward, tc.Contracts.StakingReward.Address.Hex())
		configFileViper.Set(common.FlagEthContractFarmingRewards, tc.Contracts.FarmingRewards.Address.Hex())
		configFileViper.Set(common.FlagEthContractViewer, tc.Contracts.Viewer.Address.Hex())
		configFileViper.Set(common.FlagEthContractGovern, tc.Contracts.Govern.Address.Hex())
		configFileViper.Set(common.FlagEthValidatorAddress, tc.ValEthAddrs[i].Hex())
		configFileViper.BindEnv(common.FlagGatewayAwsS3Region, "GATEWAY_AWS_S3_REGION")
		configFileViper.BindEnv(common.FlagGatewayAwsS3Bucket, "GATEWAY_AWS_S3_BUCKET")
		configFileViper.BindEnv(common.FlagGatewayAwsKey, "GATEWAY_AWS_KEY")
		configFileViper.BindEnv(common.FlagGatewayAwsSecret, "GATEWAY_AWS_SECRET")
		if !report {
			configFileViper.Set(common.FlagSgnCheckIntervalCbrPrice, 0)
			configFileViper.Set(common.FlagSgnLivenessReportEndpoint, "")
		}

		err = configFileViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		if !cbridge {
			cbrCfgPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/cbridge.toml", i)
			cbrViper := viper.New()
			cbrViper.SetConfigFile(cbrCfgPath)
			err = cbrViper.ReadInConfig()
			tc.ChkErr(err, "Failed to read config")
			cbrViper.Set(common.FlagMultiChain, []string{})
			err = cbrViper.WriteConfig()
			tc.ChkErr(err, "Failed to write config")
		}

		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err = genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		genesisViper.Set("app_state.slashing.params.staking_contract.address", eth.Addr2Hex(tc.Contracts.Staking.Address))
		genesisViper.Set("app_state.distribution.params.reward_contract.address", eth.Addr2Hex(tc.Contracts.StakingReward.Address))

		var farmingRewardContracts []commontypes.ContractInfo
		farmingRewardContracts = append(farmingRewardContracts,
			commontypes.ContractInfo{
				ChainId: 883,
				Address: eth.Addr2Hex(tc.Contracts.FarmingRewards.Address),
			},
		)
		genesisViper.Set("app_state.farming.reward_contracts", farmingRewardContracts)

		if manual {
			genesisViper.Set("app_state.gov.voting_params.voting_period", "120s")
		} else {
			genesisViper.Set("app_state.gov.voting_params.voting_period", "10s")
		}
		if !cbridge {
			genesisViper.Set("app_state.cbridge.config.assets", []string{})
			genesisViper.Set("app_state.cbridge.config.chain_pairs", []string{})
		}
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func DeployUsdtForBridge() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.USDTAddr, tc.CbrChain1.USDTContract =
				tc.DeployBridgeTestTokenContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, "USDT", "USDT", 6)
		},
		func() {
			tc.CbrChain2.USDTAddr, tc.CbrChain2.USDTContract =
				tc.DeployBridgeTestTokenContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, "USDT", "USDT", 6)
		},
	)

	// fund usdt to each user
	addrs := []eth.Addr{
		tc.ClientEthAddrs[0],
		tc.ClientEthAddrs[1],
		tc.ClientEthAddrs[2],
		tc.ClientEthAddrs[3],
	}
	log.Infoln("fund each test addr 10 million usdt on each chain")
	tc.RunAllAndWait(
		func() {
			err := tc.FundAddrsErc20(tc.CbrChain1.USDTAddr, addrs, tc.NewBigInt(1, 13), tc.CbrChain1.Ec, tc.CbrChain1.Auth)
			tc.ChkErr(err, "fund each test addr 10 million usdt on chain 1")
		},
		func() {
			err := tc.FundAddrsErc20(tc.CbrChain2.USDTAddr, addrs, tc.NewBigInt(1, 13), tc.CbrChain2.Ec, tc.CbrChain2.Auth)
			tc.ChkErr(err, "fund each test addr 10 million usdt on chain 2")
		},
	)

	log.Infoln("Updating config files of SGN nodes")
	for i := 0; i < len(tc.ValEthKs); i++ {
		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err := genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		cbrConfig := new(cbrtypes.CbrConfig)
		jsonByte, _ := json.Marshal(genesisViper.Get("app_state.cbridge.config"))
		json.Unmarshal(jsonByte, cbrConfig)
		cbrConfig.Assets[0].Addr = eth.Addr2Hex(tc.CbrChain1.USDTAddr)
		cbrConfig.Assets[1].Addr = eth.Addr2Hex(tc.CbrChain2.USDTAddr)
		genesisViper.Set("app_state.cbridge.config", cbrConfig)
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func DeployBridgeContract() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.CbrAddr, tc.CbrChain1.CbrContract = tc.DeployBridgeContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth)
		},
		func() {
			tc.CbrChain2.CbrAddr, tc.CbrChain2.CbrContract = tc.DeployBridgeContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth)
		},
	)

	for i := 0; i < len(tc.ValEthKs); i++ {
		cbrCfgPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/cbridge.toml", i)
		cbrViper := viper.New()
		cbrViper.SetConfigFile(cbrCfgPath)
		err := cbrViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		multichains := cbrViper.Get("multichain").([]interface{})
		multichains[0].(map[string]interface{})["cbridge"] = tc.CbrChain1.CbrAddr.Hex()
		multichains[1].(map[string]interface{})["cbridge"] = tc.CbrChain2.CbrAddr.Hex()
		cbrViper.Set("multichain", multichains)
		err = cbrViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err = genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		cbrConfig := new(cbrtypes.CbrConfig)
		jsonByte, _ := json.Marshal(genesisViper.Get("app_state.cbridge.config"))
		json.Unmarshal(jsonByte, cbrConfig)
		cbrConfig.CbrContracts[0].Address = eth.Addr2Hex(tc.CbrChain1.CbrAddr)
		cbrConfig.CbrContracts[1].Address = eth.Addr2Hex(tc.CbrChain2.CbrAddr)
		genesisViper.Set("app_state.cbridge.config", cbrConfig)
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func DeployPegBridgeContract() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.PegVaultAddr, tc.CbrChain1.PegVaultContract =
				tc.DeployPegVaultContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, tc.CbrChain1.CbrAddr)
			tc.CbrChain1.UNIAddr, tc.CbrChain1.UNIContract =
				tc.DeployBridgeTestTokenContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, "UNI", "UNI", 18)
			// fund UNI to each user on chan 1
			addrs := []eth.Addr{
				tc.ClientEthAddrs[0],
				tc.ClientEthAddrs[1],
				tc.ClientEthAddrs[2],
				tc.ClientEthAddrs[3],
			}
			log.Infoln("fund each test addr 10 million UNI on chain 1")
			err := tc.FundAddrsErc20(tc.CbrChain1.UNIAddr, addrs, tc.NewBigInt(1, 25), tc.CbrChain1.Ec, tc.CbrChain1.Auth)
			tc.ChkErr(err, "fund each test addr 10 million UNI on chain 1")
		},
		func() {
			tc.CbrChain2.PegBridgeAddr, tc.CbrChain2.PegBridgeContract =
				tc.DeployPegBridgeContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, tc.CbrChain2.CbrAddr)
			tc.CbrChain2.UNIAddr, tc.CbrChain2.UNIContract =
				tc.DeployBridgeTestTokenContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, "UNI", "UNI", 18)

			tx, err := tc.CbrChain2.UNIContract.UpdateBridgeSupplyCap(
				tc.CbrChain2.Auth, tc.CbrChain2.PegBridgeAddr, tc.NewBigInt(1, 28))
			tc.ChkErr(err, "failed to update bridge supply cap")
			tc.WaitMinedWithChk(context.Background(), tc.CbrChain2.Ec, tx, tc.BlockDelay, tc.PollingInterval, "UpdateBridgeSupplyCap")
		},
	)

	for i := 0; i < len(tc.ValEthKs); i++ {
		cbrCfgPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/cbridge.toml", i)
		cbrViper := viper.New()
		cbrViper.SetConfigFile(cbrCfgPath)
		err := cbrViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		multichains := cbrViper.Get("multichain").([]interface{})
		multichains[0].(map[string]interface{})["otvault"] = tc.CbrChain1.PegVaultAddr.Hex()
		multichains[1].(map[string]interface{})["ptbridge"] = tc.CbrChain2.PegBridgeAddr.Hex()
		cbrViper.Set("multichain", multichains)
		err = cbrViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		// Modify genesis to include pegbridge info
		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err = genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		peggedTokenBridges := []commontypes.ContractInfo{{
			ChainId: tc.CbrChain2.ChainId,
			Address: eth.Addr2Hex(tc.CbrChain2.PegBridgeAddr),
		}}
		originalTokenVaults := []commontypes.ContractInfo{{
			ChainId: tc.CbrChain1.ChainId,
			Address: eth.Addr2Hex(tc.CbrChain1.PegVaultAddr),
		}}
		origPeggedPairs := []pegbrtypes.OrigPeggedPair{{
			Orig: commontypes.ERC20Token{
				Symbol:   "UNI",
				ChainId:  tc.CbrChain1.ChainId,
				Address:  eth.Addr2Hex(tc.CbrChain1.UNIAddr),
				Decimals: 18,
			},
			Pegged: commontypes.ERC20Token{
				Symbol:   "UNI",
				ChainId:  tc.CbrChain2.ChainId,
				Address:  eth.Addr2Hex(tc.CbrChain2.UNIAddr),
				Decimals: 18,
			},
			MintFeePips: 100,
			BurnFeePips: 500,
			MaxMintFee:  "1000000000000000000",
			MaxBurnFee:  "1000000000000000000",
			SupplyCap:   "100000000000000000000",
		}}
		config := pegbrtypes.PegConfig{
			PeggedTokenBridges:  peggedTokenBridges,
			OriginalTokenVaults: originalTokenVaults,
			OrigPeggedPairs:     origPeggedPairs,
		}
		genesisViper.Set("app_state.pegbridge.config", config)
		genesisViper.Set("app_state.pegbridge.params.trigger_sign_cooldown", "10s")

		// Also update cbr config to add original UNI and pegged UNI, required by base fee calculation
		cbrConfig := new(cbrtypes.CbrConfig)
		jsonByte, _ := json.Marshal(genesisViper.Get("app_state.cbridge.config"))
		json.Unmarshal(jsonByte, cbrConfig)
		cbrConfig.Assets[2].Addr = eth.Addr2Hex(tc.CbrChain1.UNIAddr)
		cbrConfig.Assets[3].Addr = eth.Addr2Hex(tc.CbrChain2.UNIAddr)
		genesisViper.Set("app_state.cbridge.config", cbrConfig)

		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func DeployBatchTransferAndMessageTransferAndMessageBusContracts() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.MessageBusAddr, tc.CbrChain1.MessageBusContract =
				tc.DeployMessageBusContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, tc.CbrChain1.CbrAddr, tc.CbrChain1.PegBridgeAddr, tc.CbrChain1.PegVaultAddr)
			tc.CbrChain1.BatchTransferAddr, tc.CbrChain1.BatchTransferContract =
				tc.DeployBatchTransferContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, tc.CbrChain1.MessageBusAddr, tc.CbrChain1.CbrAddr)
			tc.CbrChain1.TransferMessageAddr, tc.CbrChain1.TransferMessageContract =
				tc.DeployTransferMessageContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, tc.CbrChain1.MessageBusAddr)
			tc.CbrChain1.TestRefundAddr, tc.CbrChain1.TestRefundContract =
				tc.DeployTestRefundContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, tc.CbrChain1.MessageBusAddr)
		}, func() {
			tc.CbrChain2.MessageBusAddr, tc.CbrChain2.MessageBusContract =
				tc.DeployMessageBusContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, tc.CbrChain2.CbrAddr, tc.CbrChain2.PegBridgeAddr, tc.CbrChain2.PegVaultAddr)
			tc.CbrChain2.BatchTransferAddr, tc.CbrChain2.BatchTransferContract =
				tc.DeployBatchTransferContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, tc.CbrChain2.MessageBusAddr, tc.CbrChain2.CbrAddr)
			tc.CbrChain2.TransferMessageAddr, tc.CbrChain2.TransferMessageContract =
				tc.DeployTransferMessageContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, tc.CbrChain2.MessageBusAddr)
			tc.CbrChain2.TestRefundAddr, tc.CbrChain2.TestRefundContract =
				tc.DeployTestRefundContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, tc.CbrChain2.MessageBusAddr)
		})

	messageBuses := make([]msgtypes.MessageBusInfo, 0)
	bus1 := msgtypes.MessageBusInfo{
		ContractInfo: &commontypes.ContractInfo{
			ChainId: tc.CbrChain1.ChainId,
			Address: eth.Addr2Hex(tc.CbrChain1.MessageBusAddr),
		},
		FeeTokenSymbol: "ETH",
	}
	bus2 := msgtypes.MessageBusInfo{
		ContractInfo: &commontypes.ContractInfo{
			ChainId: tc.CbrChain2.ChainId,
			Address: eth.Addr2Hex(tc.CbrChain2.MessageBusAddr),
		},
		FeeTokenSymbol: "ETH",
	}
	messageBuses = append(messageBuses, bus1, bus2)
	for i := 0; i < len(tc.ValEthKs); i++ {
		cbrCfgPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/cbridge.toml", i)
		cbrViper := viper.New()
		cbrViper.SetConfigFile(cbrCfgPath)
		err := cbrViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		multichains := cbrViper.Get("multichain").([]interface{})
		multichains[0].(map[string]interface{})["msgbus"] = bus1.ContractInfo.Address
		multichains[1].(map[string]interface{})["msgbus"] = bus2.ContractInfo.Address
		cbrViper.Set("multichain", multichains)
		err = cbrViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err = genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		genesisViper.Set("app_state.message.message_buses", messageBuses)
		genesisViper.Set("app_state.message.params.trigger_sign_cooldown", "10s")
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func PrepareExecutor() {
	tc.RunAllAndWait(
		func() {
			tc.RunCmd("make", "localnet-start-crdb")
		},
		func() {
			SetupExecutorConfig()
		},
		func() {
			tc.RunCmd("make", "build-executor")
		},
	)
	tc.RunCmd("docker-compose", "up", "-d", "executor")
}

func SetupExecutorConfig() {
	tc.RunCmd("make", "prepare-executor-data")
	// setup cbridge.toml contract addresses
	msgViper := viper.New()
	msgViper.SetConfigFile("../../../docker-volumes/executor/config/cbridge.toml")
	err := msgViper.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")
	multichains := msgViper.Get("multichain").([]interface{})
	multichains[0].(map[string]interface{})["cbridge"] = tc.CbrChain1.CbrAddr.Hex()
	multichains[1].(map[string]interface{})["cbridge"] = tc.CbrChain2.CbrAddr.Hex()
	multichains[0].(map[string]interface{})["msgbus"] = tc.CbrChain1.MessageBusAddr.Hex()
	multichains[1].(map[string]interface{})["msgbus"] = tc.CbrChain2.MessageBusAddr.Hex()
	multichains[0].(map[string]interface{})["otvault"] = tc.CbrChain1.PegVaultAddr.Hex()
	multichains[1].(map[string]interface{})["otvault"] = tc.CbrChain2.PegVaultAddr.Hex()
	multichains[0].(map[string]interface{})["ptbridge"] = tc.CbrChain1.PegBridgeAddr.Hex()
	multichains[1].(map[string]interface{})["ptbridge"] = tc.CbrChain2.PegBridgeAddr.Hex()
	msgViper.Set("multichain", multichains)
	err = msgViper.WriteConfig()
	tc.ChkErr(err, "Failed to write config")

	msgViper.SetConfigFile("../../../docker-volumes/executor/config/executor.toml")
	err = msgViper.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")

	var contracts []interface{}
	contractInfo1 := make(map[string]interface{})
	contractInfo1["address"] = tc.CbrChain1.BatchTransferAddr.Hex()
	contractInfo1["chainid"] = tc.CbrChain1.ChainId
	contracts = append(contracts, contractInfo1)
	contractInfo2 := make(map[string]interface{})
	contractInfo2["address"] = tc.CbrChain1.TestRefundAddr.Hex()
	contractInfo2["chainid"] = tc.CbrChain1.ChainId
	contracts = append(contracts, contractInfo2)
	contractInfo3 := make(map[string]interface{})
	contractInfo3["address"] = tc.CbrChain1.TransferMessageAddr.Hex()
	contractInfo3["chainid"] = tc.CbrChain1.ChainId
	contracts = append(contracts, contractInfo3)
	contractInfo4 := make(map[string]interface{})
	contractInfo4["address"] = tc.CbrChain2.BatchTransferAddr.Hex()
	contractInfo4["chainid"] = tc.CbrChain2.ChainId
	contracts = append(contracts, contractInfo4)
	contractInfo5 := make(map[string]interface{})
	contractInfo5["address"] = tc.CbrChain2.TestRefundAddr.Hex()
	contractInfo5["chainid"] = tc.CbrChain2.ChainId
	contracts = append(contracts, contractInfo5)
	contractInfo6 := make(map[string]interface{})
	contractInfo6["address"] = tc.CbrChain2.TransferMessageAddr.Hex()
	contractInfo6["chainid"] = tc.CbrChain2.ChainId
	contracts = append(contracts, contractInfo6)
	msgViper.Set("executor.contracts", contracts)
	err = msgViper.WriteConfig()
	tc.ChkErr(err, "Failed to write config")
}

func CreateFarmingPools() {
	log.Infoln("Creating farming pools in genesis")
	for i := 0; i < len(tc.ValEthKs); i++ {
		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err := genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")

		// TODO: Extract constants
		// Set claim_cooldown
		genesisViper.Set("app_state.farming.params.claim_cooldown", "1s")

		// Add a farming pool with two reward tokens
		poolName := "cbridge-USDT/883"
		var pools farmingtypes.FarmingPools
		pool := farmingtypes.NewFarmingPool(
			poolName,
			commontypes.ERC20Token{
				ChainId: 883,
				Symbol:  "CB-USDT",
				Address: eth.Addr2Hex(tc.CbrChain1.USDTAddr),
			},
			[]commontypes.ERC20Token{
				{
					ChainId: 883,
					Symbol:  "CELR",
					Address: eth.Addr2Hex(tc.CelrAddr),
				},
				{
					ChainId: 883,
					Symbol:  "USDT",
					Address: eth.Addr2Hex(tc.CbrChain1.USDTAddr),
				},
			},
			sdk.NewDecCoin("CB-USDT/883", sdk.ZeroInt()),
			[]farmingtypes.RewardTokenInfo{
				{
					RemainingAmount:        sdk.NewDecCoin("CELR/883", sdk.NewInt(1000000).Mul(sdk.NewInt(1e18))),
					RewardStartBlockHeight: 1,
					RewardAmountPerBlock:   sdk.NewDec(1e18),
				},
				{
					RemainingAmount:        sdk.NewDecCoin("USDT/883", sdk.NewInt(1000000).Mul(sdk.NewInt(1e6))),
					RewardStartBlockHeight: 1,
					RewardAmountPerBlock:   sdk.NewDec(1e6),
				},
			},
			sdk.NewDecCoins(
				sdk.NewDecCoin("CELR/883", sdk.ZeroInt()),
				sdk.NewDecCoin("USDT/883", sdk.ZeroInt()),
			),
		)
		pools = append(pools, pool)
		genesisViper.Set("app_state.farming.pools", pools)

		// Set initial reward records
		poolHistoricalRewardsRecord := farmingtypes.PoolHistoricalRewardsRecord{
			PoolName: poolName,
			Period:   0,
			Rewards:  farmingtypes.NewPoolHistoricalRewards(sdk.DecCoins{}, 1),
		}
		poolCurrentRewardsRecord := farmingtypes.PoolCurrentRewardsRecord{
			PoolName: poolName,
			Rewards:  farmingtypes.NewPoolCurrentRewards(0, 1, sdk.DecCoins{}),
		}
		genesisViper.Set(
			"app_state.farming.pool_historical_rewards",
			[]farmingtypes.PoolHistoricalRewardsRecord{poolHistoricalRewardsRecord})
		genesisViper.Set(
			"app_state.farming.pool_current_rewards",
			[]farmingtypes.PoolCurrentRewardsRecord{poolCurrentRewardsRecord})

		// Ensure reward module account balances
		rewardCoins := sdk.NewCoins(
			sdk.NewCoin("CELR/883", sdk.NewInt(1000000).Mul(sdk.NewInt(1e18))),
			sdk.NewCoin("USDT/883", sdk.NewInt(1000000).Mul(sdk.NewInt(1e6))),
		).Sort()
		var balances []banktypes.Balance
		jsonByte, _ := json.Marshal(genesisViper.Get("app_state.bank.balances"))
		json.Unmarshal(jsonByte, &balances)
		rewardModuleAccountAddress := authtypes.NewModuleAddress(farmingtypes.RewardModuleAccountName).String()
		hasBalance := false
		for _, balance := range balances {
			if balance.Address == rewardModuleAccountAddress {
				hasBalance = true
				balance.Coins = rewardCoins
				break
			}
		}
		if !hasBalance {
			balances = append(balances, banktypes.Balance{
				Address: rewardModuleAccountAddress,
				Coins:   rewardCoins,
			})
		}
		genesisViper.Set("app_state.bank.balances", balances)

		// Ensure genesis supply
		var supply sdk.Coins
		jsonByte, _ = json.Marshal(genesisViper.Get("app_state.bank.supply"))
		json.Unmarshal(jsonByte, &supply)
		supply.Sort()
		for _, reward := range rewardCoins {
			existingAmount := supply.AmountOf(reward.Denom)
			existingCoin := sdk.NewCoin(reward.Denom, existingAmount)
			supply = supply.Sub(sdk.NewCoins(existingCoin)).Add(sdk.NewCoins(reward)...)
		}
		genesisViper.Set("app_state.bank.supply", supply)

		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func FundUsdtFarmingReward() {
	amt := tc.NewBigInt(1, 13)
	usdtContract := tc.CbrChain1.USDTContract
	approveTx, err := usdtContract.Approve(tc.EtherBaseAuth, tc.Contracts.FarmingRewards.Address, amt)
	tc.ChkErr(err, "failed to approve USDT to FarmingRewards")
	tc.WaitMinedWithChk(context.Background(), tc.EthClient, approveTx, tc.BlockDelay, tc.PollingInterval, "approve USDT")
	allowance, _ := usdtContract.Allowance(&bind.CallOpts{}, tc.EtherBaseAuth.From, tc.Contracts.FarmingRewards.Address)
	log.Infoln("allowance to FarmingRewards", allowance.String())
	_, err = tc.Contracts.FarmingRewards.ContributeToRewardPool(tc.EtherBaseAuth, tc.CbrChain1.USDTAddr, amt)
	tc.ChkErr(err, "failed to contribute USDT to FarmingRewards")
}

func ShutdownNode(node uint) {
	log.Infoln("Shutdown node", node)
	cmd := exec.Command("docker-compose", "stop", fmt.Sprintf("sgnnode%d", node))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Error(err)
	}
}

func BringupNode(node uint) {
	log.Infoln("Shutdown node", node)
	cmd := exec.Command("docker-compose", "up", "-d", fmt.Sprintf("sgnnode%d", node))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Error(err)
	}
}
