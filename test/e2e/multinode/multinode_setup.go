package multinode

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"time"

	flowSigner "github.com/celer-network/cbridge-flow/signer"
	flowutils "github.com/celer-network/cbridge-flow/utils"
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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

var bridgeChainStarted bool

type TestFlags struct {
	Bridge, Msg, Flow, Manual, Report bool
}

func BuildDockers() {
	tc.RunCmd("make", "localnet-down")
	tc.RunCmd("make", "build-node")
	tc.RunCmd("make", "build-linux")
	tc.RunCmd("make", "build-geth")
}

func SetupMainchain() {
	tc.RunCmd("make", "localnet-stop-geth")
	tc.RunCmd("make", "prepare-geth-data")
	tc.RunCmd("make", "localnet-start-geth")
	waitGethStart(tc.LocalGeth1)

	// set up mainchain: deploy contracts, fund addrs, etc
	log.Infoln("fund each test addr 100 ETH on chain 1")
	err := tc.FundAddrsETH(tc.Addrs, tc.NewBigInt(1, 20), tc.LocalGeth1, int64(tc.Geth1ChainID))
	tc.ChkErr(err, "fund each test addr 100 ETH")

	log.Infoln("set up chain 1")
	tc.SetupEthClients()
	tc.CelrAddr, tc.CelrContract = tc.DeployERC20Contract(tc.EthClient, tc.EtherBaseAuth, "Celer", "CELR", 18)

	// fund CELR to each eth account
	log.Infoln("fund each validator and delegator addr 1 billion CELR")
	err = tc.FundAddrsErc20(tc.CelrAddr, tc.ValDelAddrs, tc.NewBigInt(1, 27), tc.EthClient, tc.EtherBaseAuth)
	tc.ChkErr(err, "fund each validator and delegator addr 1 billion CELR")
}

func SetupBridgeChains() {
	if bridgeChainStarted {
		log.Info("bridge chains already started")
		return
	}

	tc.RunAllAndWait(
		func() {
			waitGethStart(tc.LocalGeth2)
			log.Infoln("fund each test addr 100 ETH on chain 2")
			err := tc.FundAddrsETH(tc.Addrs2, tc.NewBigInt(1, 20), tc.LocalGeth2, int64(tc.Geth2ChainID))
			tc.ChkErr(err, "fund each test addr 100 ETH on chain 2")
		},
		func() {
			waitGethStart(tc.LocalGeth3)
			log.Infoln("fund each test addr 100 ETH on chain 3")
			err := tc.FundAddrsETH(tc.Addrs2, tc.NewBigInt(1, 20), tc.LocalGeth3, int64(tc.Geth3ChainID))
			tc.ChkErr(err, "fund each test addr 100 ETH on chain 3")
		},
	)
	log.Infoln("set up bridge chains")
	tc.InitCbrChainConfigs()
	bridgeChainStarted = true
}

func SetupNewSgnEnv(contractParams *tc.ContractParams, tf *TestFlags) {
	if tf == nil {
		tf = &TestFlags{}
	}
	tc.RunAllAndWait(
		SetupBridgeChains,
		func() {
			deployContractsAndPrepareSgnData(contractParams, tf)
		},
	)

	if tf.Bridge {
		deployUsdtForBridge()
		deployBridgeContract()
		deployPegBridgeContract()
		createFarmingPools()
		fundUsdtFarmingReward()
		deployWdInboxContract()
		tc.ApproveTestTokenToBridges()
		if tf.Flow {
			setupFlowChain()
		} else {
			removeFlowCfg()
		}
	}
	if tf.Msg {
		deployMessageContracts()
	}

	// Update global viper
	node0ConfigPath := "../../../docker-volumes/node0/sgnd/config/sgn.toml"
	viper.SetConfigFile(node0ConfigPath)
	err := viper.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")
	if tf.Bridge && !tf.Manual {
		// set node3 account as node0 transactor
		viper.Set(common.FlagSgnTransactors, []string{tc.ValSgnAddrStrs[3]})
		err = viper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		node3ConfigPath := "../../../docker-volumes/node3/sgnd/config/sgn.toml"
		configFileViper := viper.New()
		configFileViper.SetConfigFile(node3ConfigPath)
		err := configFileViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read node3 config")
		configFileViper.Set(common.FlagSgnWitnessMode, true)
		err = configFileViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")
	}

	tc.RunCmd("make", "localnet-up-nodes")
	if tf.Msg {
		PrepareExecutor()
	}
}

func deployContractsAndPrepareSgnData(contractParams *tc.ContractParams, tf *TestFlags) {
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
		if !tf.Report {
			configFileViper.Set(common.FlagSgnCheckIntervalCbrPrice, 0)
			configFileViper.Set(common.FlagSgnLivenessReportEndpoint, "")
			configFileViper.Set(common.FlagSgnConsensusLogReportEndpoint, "")
		}

		err = configFileViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		if !tf.Bridge {
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

		if tf.Manual {
			genesisViper.Set("app_state.gov.voting_params.voting_period", "120s")
		} else {
			genesisViper.Set("app_state.gov.voting_params.voting_period", "10s")
		}
		if !tf.Bridge {
			genesisViper.Set("app_state.cbridge.config.assets", []string{})
			genesisViper.Set("app_state.cbridge.config.chain_pairs", []string{})
		}
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func deployUsdtForBridge() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.USDTAddr, tc.CbrChain1.USDTContract =
				tc.DeployBridgeTestTokenContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth, "USDT", "USDT", 6)
		},
		func() {
			tc.CbrChain2.USDTAddr, tc.CbrChain2.USDTContract =
				tc.DeployBridgeTestTokenContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth, "USDT", "USDT", 6)
		},
		func() {
			tc.CbrChain3.USDTAddr, tc.CbrChain3.USDTContract =
				tc.DeployBridgeTestTokenContract(tc.CbrChain3.Ec, tc.CbrChain3.Auth, "USDT", "USDT", 6)
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
		func() {
			err := tc.FundAddrsErc20(tc.CbrChain3.USDTAddr, addrs, tc.NewBigInt(1, 13), tc.CbrChain3.Ec, tc.CbrChain3.Auth)
			tc.ChkErr(err, "fund each test addr 10 million usdt on chain 3")
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
		cbrConfig.Assets[2].Addr = eth.Addr2Hex(tc.CbrChain3.USDTAddr)
		genesisViper.Set("app_state.cbridge.config", cbrConfig)
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func deployBridgeContract() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.CbrAddr, tc.CbrChain1.CbrContract = tc.DeployBridgeContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth)
		},
		func() {
			tc.CbrChain2.CbrAddr, tc.CbrChain2.CbrContract = tc.DeployBridgeContract(tc.CbrChain2.Ec, tc.CbrChain2.Auth)
		},
		func() {
			tc.CbrChain3.CbrAddr, tc.CbrChain3.CbrContract = tc.DeployBridgeContract(tc.CbrChain3.Ec, tc.CbrChain3.Auth)
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
		multichains[2].(map[string]interface{})["cbridge"] = tc.CbrChain3.CbrAddr.Hex()
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
		cbrConfig.CbrContracts[2].Address = eth.Addr2Hex(tc.CbrChain3.CbrAddr)
		genesisViper.Set("app_state.cbridge.config", cbrConfig)
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func deployWdInboxContract() {
	tc.CbrChain1.WdiAddr, tc.CbrChain1.WdInboxContract = tc.DeployWithdrawInboxContract(tc.CbrChain1.Ec, tc.CbrChain1.Auth)
	// let u0 be owner of ContractAsLP
	tc.CbrChain1.CLPAddr, tc.CbrChain1.CLPContract =
		tc.DeployContractAsLPContract(tc.CbrChain1.Ec, tc.CbrChain1.Users[0].Auth, tc.CbrChain1.CbrAddr, tc.CbrChain1.WdiAddr)

	for i := 0; i < len(tc.ValEthKs); i++ {
		cbrCfgPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/cbridge.toml", i)
		cbrViper := viper.New()
		cbrViper.SetConfigFile(cbrCfgPath)
		err := cbrViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		multichains := cbrViper.Get("multichain").([]interface{})
		multichains[0].(map[string]interface{})["wdinbox"] = tc.CbrChain1.WdiAddr.Hex()
		cbrViper.Set("multichain", multichains)
		err = cbrViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")
	}
}

func deployPegBridgeContract() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.DeployPegVaultContracts()
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
			tc.CbrChain2.DeployPegBridgeContracts()
			tx, err := tc.CbrChain2.UNIContract.UpdateBridgeSupplyCap(
				tc.CbrChain2.Auth, tc.CbrChain2.PegBridgeAddr, tc.NewBigInt(1, 28))
			tc.ChkErr(err, "failed to update bridge supply cap")
			tx, err = tc.CbrChain2.FETContract.UpdateBridgeSupplyCap(
				tc.CbrChain2.Auth, tc.CbrChain2.PegBridgeAddr, tc.NewBigInt(1, 28))
			tc.ChkErr(err, "failed to update bridge supply cap")
			tx, err = tc.CbrChain2.USDTContract.UpdateBridgeSupplyCap(
				tc.CbrChain2.Auth, tc.CbrChain2.PegBridgeV2Addr, tc.NewBigInt(1, 15))
			tc.ChkErr(err, "failed to update bridge supply cap")
			tc.WaitMinedWithChk(context.Background(), tc.CbrChain2.Ec, tx, tc.BlockDelay, tc.PollingInterval, "UpdateBridgeSupplyCap")
		},
		func() {
			tc.CbrChain3.PegBridgeV2Addr, tc.CbrChain3.PegBridgeV2Contract =
				tc.DeployPegBridgeV2Contract(tc.CbrChain3.Ec, tc.CbrChain3.Auth, tc.CbrChain3.CbrAddr)
			tx, err := tc.CbrChain3.USDTContract.UpdateBridgeSupplyCap(
				tc.CbrChain3.Auth, tc.CbrChain3.PegBridgeV2Addr, tc.NewBigInt(1, 15))
			tc.ChkErr(err, "failed to update bridge supply cap")
			tc.WaitMinedWithChk(context.Background(), tc.CbrChain3.Ec, tx, tc.BlockDelay, tc.PollingInterval, "UpdateBridgeSupplyCap")
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
		multichains[0].(map[string]interface{})["otvault2"] = tc.CbrChain1.PegVaultV2Addr.Hex()
		multichains[1].(map[string]interface{})["ptbridge"] = tc.CbrChain2.PegBridgeAddr.Hex()
		multichains[1].(map[string]interface{})["ptbridge2"] = tc.CbrChain2.PegBridgeV2Addr.Hex()
		multichains[2].(map[string]interface{})["ptbridge2"] = tc.CbrChain3.PegBridgeV2Addr.Hex()
		cbrViper.Set("multichain", multichains)
		err = cbrViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		// Modify genesis to include pegbridge info
		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err = genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		peggedTokenBridges := []pegbrtypes.ContractInfo{
			{
				Contract: commontypes.ContractInfo{
					ChainId: tc.CbrChain2.ChainId,
					Address: eth.Addr2Hex(tc.CbrChain2.PegBridgeAddr),
				},
			},
			{
				Contract: commontypes.ContractInfo{
					ChainId: tc.CbrChain2.ChainId,
					Address: eth.Addr2Hex(tc.CbrChain2.PegBridgeV2Addr),
				},
				Version: 2,
			},
			{
				Contract: commontypes.ContractInfo{
					ChainId: tc.CbrChain3.ChainId,
					Address: eth.Addr2Hex(tc.CbrChain3.PegBridgeV2Addr),
				},
				Version: 2,
			},
			{
				Contract: commontypes.ContractInfo{
					ChainId: 12340003,
					Address: "01cf0e2f2f715450",
				},
			},
		}
		originalTokenVaults := []pegbrtypes.ContractInfo{
			{
				Contract: commontypes.ContractInfo{
					ChainId: tc.CbrChain1.ChainId,
					Address: eth.Addr2Hex(tc.CbrChain1.PegVaultAddr),
				},
			},
			{
				Contract: commontypes.ContractInfo{
					ChainId: tc.CbrChain1.ChainId,
					Address: eth.Addr2Hex(tc.CbrChain1.PegVaultV2Addr),
				},
				Version: 2,
			},
			{
				Contract: commontypes.ContractInfo{
					ChainId: 12340003,
					Address: "01cf0e2f2f715450",
				},
			},
		}
		origPeggedPairs := []pegbrtypes.OrigPeggedPair{
			{
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
			},
			{
				Orig: commontypes.ERC20Token{
					Symbol:   "USDT",
					ChainId:  tc.CbrChain1.ChainId,
					Address:  eth.Addr2Hex(tc.CbrChain1.USDTAddr),
					Decimals: 6,
				},
				Pegged: commontypes.ERC20Token{
					Symbol:   "USDT",
					ChainId:  tc.CbrChain2.ChainId,
					Address:  eth.Addr2Hex(tc.CbrChain2.USDTAddr),
					Decimals: 6,
				},
				MintFeePips:   100,
				BurnFeePips:   500,
				MaxMintFee:    "100000000",
				MaxBurnFee:    "100000000",
				SupplyCap:     "1000000000000000",
				BridgeVersion: 2,
				VaultVersion:  2,
			},
			{
				Orig: commontypes.ERC20Token{
					Symbol:   "USDT",
					ChainId:  tc.CbrChain1.ChainId,
					Address:  eth.Addr2Hex(tc.CbrChain1.USDTAddr),
					Decimals: 6,
				},
				Pegged: commontypes.ERC20Token{
					Symbol:   "USDT",
					ChainId:  tc.CbrChain3.ChainId,
					Address:  eth.Addr2Hex(tc.CbrChain3.USDTAddr),
					Decimals: 6,
				},
				MintFeePips:   100,
				BurnFeePips:   500,
				MaxMintFee:    "100000000",
				MaxBurnFee:    "100000000",
				SupplyCap:     "1000000000000000",
				BridgeVersion: 2,
				VaultVersion:  2,
			},
			{
				// flow PegBridge
				Orig: commontypes.ERC20Token{
					Symbol:   "UNI",
					ChainId:  tc.CbrChain1.ChainId,
					Address:  eth.Addr2Hex(tc.CbrChain1.UNIAddr),
					Decimals: 18,
				},
				Pegged: commontypes.ERC20Token{
					Symbol:   "UNI",
					ChainId:  12340003,
					Address:  "A.01cf0e2f2f715450.PegToken.Vault",
					Decimals: 8,
				},
				MintFeePips: 100,
				BurnFeePips: 500,
				MaxMintFee:  "1000000000000000000",
				MaxBurnFee:  "1000000000000000000",
				SupplyCap:   "100000000000000000000",
			},
			{
				// flow SafeBox
				Orig: commontypes.ERC20Token{
					Symbol:   "FET",
					ChainId:  12340003,
					Address:  "A.01cf0e2f2f715450.ExampleToken.Vault",
					Decimals: 8,
				},
				Pegged: commontypes.ERC20Token{
					Symbol:   "FET",
					ChainId:  tc.CbrChain2.ChainId,
					Address:  eth.Addr2Hex(tc.CbrChain2.FETAddr),
					Decimals: 18,
				},
				MintFeePips: 0,
				BurnFeePips: 0,
				MaxMintFee:  "1000000000000000000000",
				MaxBurnFee:  "1000000000000000000000",
				SupplyCap:   "100000000000000000000000000",
			},
		}
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
		cbrConfig.Assets[3].Addr = eth.Addr2Hex(tc.CbrChain1.UNIAddr)
		cbrConfig.Assets[4].Addr = eth.Addr2Hex(tc.CbrChain2.UNIAddr)
		genesisViper.Set("app_state.cbridge.config", cbrConfig)

		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func deployMessageContracts() {
	tc.RunAllAndWait(
		func() {
			tc.CbrChain1.DeployMessageContracts()
		}, func() {
			tc.CbrChain2.DeployMessageContracts()
		}, func() {
			tc.CbrChain3.DeployMessageContracts()
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
	bus3 := msgtypes.MessageBusInfo{
		ContractInfo: &commontypes.ContractInfo{
			ChainId: tc.CbrChain3.ChainId,
			Address: eth.Addr2Hex(tc.CbrChain3.MessageBusAddr),
		},
		FeeTokenSymbol: "ETH",
	}
	messageBuses = append(messageBuses, bus1, bus2, bus3)
	for i := 0; i < len(tc.ValEthKs); i++ {
		cbrCfgPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/cbridge.toml", i)
		cbrViper := viper.New()
		cbrViper.SetConfigFile(cbrCfgPath)
		err := cbrViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		multichains := cbrViper.Get("multichain").([]interface{})
		multichains[0].(map[string]interface{})["msgbus"] = bus1.ContractInfo.Address
		multichains[1].(map[string]interface{})["msgbus"] = bus2.ContractInfo.Address
		multichains[2].(map[string]interface{})["msgbus"] = bus3.ContractInfo.Address
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
			time.Sleep(3 * time.Second)
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
	multichain := viper.New()
	multichain.SetConfigFile("../../../docker-volumes/executor/config/cbridge.toml")
	err := multichain.ReadInConfig()
	tc.ChkErr(err, "Failed to read configs")
	multichains := multichain.Get("multichain").([]interface{})
	multichains[0].(map[string]interface{})["cbridge"] = tc.CbrChain1.CbrAddr.Hex()
	multichains[1].(map[string]interface{})["cbridge"] = tc.CbrChain2.CbrAddr.Hex()
	multichains[2].(map[string]interface{})["cbridge"] = tc.CbrChain3.CbrAddr.Hex()

	multichains[0].(map[string]interface{})["msgbus"] = tc.CbrChain1.MessageBusAddr.Hex()
	multichains[1].(map[string]interface{})["msgbus"] = tc.CbrChain2.MessageBusAddr.Hex()
	multichains[2].(map[string]interface{})["msgbus"] = tc.CbrChain3.MessageBusAddr.Hex()

	multichains[0].(map[string]interface{})["otvault"] = tc.CbrChain1.PegVaultAddr.Hex()
	multichains[0].(map[string]interface{})["otvault2"] = tc.CbrChain1.PegVaultV2Addr.Hex()
	multichains[1].(map[string]interface{})["otvault"] = tc.CbrChain2.PegVaultAddr.Hex()
	multichains[2].(map[string]interface{})["otvault"] = tc.CbrChain3.PegVaultAddr.Hex()

	multichains[0].(map[string]interface{})["ptbridge"] = tc.CbrChain1.PegBridgeAddr.Hex()
	multichains[1].(map[string]interface{})["ptbridge"] = tc.CbrChain2.PegBridgeAddr.Hex()
	multichains[1].(map[string]interface{})["ptbridge2"] = tc.CbrChain2.PegBridgeV2Addr.Hex()
	multichains[2].(map[string]interface{})["ptbridge"] = tc.CbrChain3.PegBridgeAddr.Hex()
	multichains[2].(map[string]interface{})["ptbridge2"] = tc.CbrChain3.PegBridgeV2Addr.Hex()

	multichain.Set("multichain", multichains)
	err = multichain.WriteConfig()
	tc.ChkErr(err, "Failed to write config")

	executor := viper.New()
	executor.SetConfigFile("../../../docker-volumes/executor/config/executor.toml")
	err = executor.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")
	service := executor.Get("service").([]interface{})
	var contracts0 []interface{}
	var contracts1 []interface{}
	contractInfo1 := make(map[string]interface{})
	contractInfo1["address"] = tc.CbrChain1.BatchTransferAddr.Hex()
	contractInfo1["chain_id"] = tc.CbrChain1.ChainId
	contractInfo1["add_payable_value_for_execution"] = tc.MsgFeeBase
	contracts0 = append(contracts0, contractInfo1)
	contractInfo2 := make(map[string]interface{})
	contractInfo2["address"] = tc.CbrChain1.MsgTestAddr.Hex()
	contractInfo2["chain_id"] = tc.CbrChain1.ChainId
	contracts0 = append(contracts0, contractInfo2)

	contractInfo4 := make(map[string]interface{})
	contractInfo4["address"] = tc.CbrChain2.BatchTransferAddr.Hex()
	contractInfo4["chain_id"] = tc.CbrChain2.ChainId
	contractInfo4["add_payable_value_for_execution"] = tc.MsgFeeBase
	contracts1 = append(contracts1, contractInfo4)
	contractInfo5 := make(map[string]interface{})
	contractInfo5["address"] = tc.CbrChain2.MsgTestAddr.Hex()
	contractInfo5["chain_id"] = tc.CbrChain2.ChainId
	contracts1 = append(contracts1, contractInfo5)

	contractInfo7 := make(map[string]interface{})
	contractInfo7["address"] = tc.CbrChain3.BatchTransferAddr.Hex()
	contractInfo7["chain_id"] = tc.CbrChain3.ChainId
	contractInfo7["add_payable_value_for_execution"] = tc.MsgFeeBase
	contracts0 = append(contracts0, contractInfo7)
	contractInfo8 := make(map[string]interface{})
	contractInfo8["address"] = tc.CbrChain3.MsgTestAddr.Hex()
	contractInfo8["chain_id"] = tc.CbrChain3.ChainId
	contracts0 = append(contracts0, contractInfo8)
	service[0].(map[string]interface{})["contracts"] = contracts0
	service[1].(map[string]interface{})["contracts"] = contracts1

	executor.Set("service", service)
	err = executor.WriteConfig()
	tc.ChkErr(err, "Failed to write config")
}

func createFarmingPools() {
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

func fundUsdtFarmingReward() {
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

//Flow config
const (
	exampleTokenName     = "ExampleToken"
	exampleTokenVault    = "ExampleTokenVault"
	exampleTokenReceiver = "ExampleTokenReceiver"
	exampleTokenBalance  = "ExampleTokenBalance"

	testPegTokenName     = "PegToken"
	testPegTokenVault    = "PegTokenVault"
	testPegTokenReceiver = "PegTokenReceiver"
	testPegTokenBalance  = "PegTokenBalance"

	safeBoxAdmin   = "SafeBoxAdmin"
	pegBridgeAdmin = "PegBridgeAdmin"
)

func setupFlowChain() {
	log.Info("start setup flowchain")
	tc.RunCmd("make", "localnet-restart-flow")
	// create two account first
	err := tc.SetupFlowServiceAccountClient()
	tc.ChkErr(err, "SetupFlowServiceAccountClient")
	// create contract acc
	_, err = flowutils.CreateNewAccount(context.Background(), tc.FlowServiceAccountClient, tc.FlowServiceAccountSigner.FlowPubKey)
	if err != nil {
		log.Fatal(err)
	}
	flowutils.FundAccountInEmulator(tc.FlowServiceAccountClient, tc.FlowContractAddr, 10000.0)
	// create user acc
	_, err = flowutils.CreateNewAccount(context.Background(), tc.FlowServiceAccountClient, tc.FlowServiceAccountSigner.FlowPubKey)
	if err != nil {
		log.Fatal(err)
	}

	// reset signers of vault contract
	newSigners := make(map[string]*big.Int)
	flowutils.FundAccountInEmulator(tc.FlowServiceAccountClient, tc.FlowUserAddr, 10000.0)
	for _, s := range tc.ValSignerKs {
		vsigner, err := flowSigner.NewFlowSigner(s, "")
		if err != nil {
			log.Fatal(err)
		}
		_, err = flowutils.CreateNewAccount(context.Background(), tc.FlowServiceAccountClient, vsigner.FlowPubKey)
		if err != nil {
			log.Fatal(err)
		}
		newSigners[vsigner.PubKey] = new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18))
	}
	for _, addr := range tc.FlowSignerAddrs {
		flowutils.FundAccountInEmulator(tc.FlowServiceAccountClient, addr, 10000.0)
	}

	log.Infof("all account flow token added")
	tc.SetupContractFlowClient()
	tc.SetupUserFlowClient(tc.FlowContractAddr.String(), tc.FlowContractAddr.String(), tc.FlowContractAddr.String())

	flowutils.DeployAllContract(context.Background(), tc.FlowContractAccountClient, uint64(commontypes.NonEvmChainID_FLOW_EMULATOR))

	// add config in bridge vaultAddr, tokenVault, tokenReceiver, tokenName
	tokenCfg := flowutils.SafeBoxTokenConfig{
		TokenAddr:      tc.FlowContractAddr,
		MinDepo:        "0.0",
		MaxDepo:        "10000000.0",
		DelayThreshold: "10000000.0",
		Cap:            "0.0",
	}
	_, err = flowutils.AddNewTokenInSafeBox(context.Background(), tc.FlowContractAccountClient, tokenCfg,
		tc.FlowContractAddr.String(), exampleTokenVault, exampleTokenReceiver, exampleTokenName)
	if err != nil {
		log.Fatal(err)
	}

	pegTokenCfg := flowutils.PegBridgeTokenConfig{
		TokenAddr:      tc.FlowContractAddr,
		MinBurn:        "0.0",
		MaxBurn:        "10000000.0",
		DelayThreshold: "10000000.0",
		Cap:            "0.0",
	}
	_, err = flowutils.AddNewTokenInPegBridge(context.Background(), tc.FlowContractAccountClient, pegTokenCfg,
		tc.FlowContractAddr.String(), testPegTokenReceiver, testPegTokenName)
	if err != nil {
		log.Fatal(err)
	}

	_, err = flowutils.AddMinterAndBurnerInPegBridgeCdc(context.Background(), tc.FlowContractAccountClient, tc.FlowContractAddr, tc.FlowContractAddr.String(), tc.FlowContractAddr.String(), tc.FlowContractAddr.String(), testPegTokenName)
	if err != nil {
		log.Fatal(err)
	}

	// reset signers of vault contract
	_, err = flowutils.ResetBridgeSigners(context.Background(), tc.FlowContractAccountClient, tc.FlowContractAddr.String(), newSigners)
	if err != nil {
		log.Fatal(err)
	}

	// add token vault for acc2
	_, err = flowutils.SetupTokenVault(context.Background(), tc.FlowUserAccountClient, tc.FlowContractAddr.String(),
		exampleTokenName, exampleTokenVault, exampleTokenBalance, exampleTokenReceiver)
	if err != nil {
		log.Fatal(err)
	}

	_, err = flowutils.SetupTokenVault(context.Background(), tc.FlowUserAccountClient, tc.FlowContractAddr.String(),
		testPegTokenName, testPegTokenVault, testPegTokenBalance, testPegTokenReceiver)
	if err != nil {
		log.Fatal(err)
	}

	// test transfer
	_, err = flowutils.TransferToken(context.Background(), tc.FlowContractAccountClient, "1000000.0", tc.FlowUserAddr,
		tc.FlowContractAddr.String(), exampleTokenName, exampleTokenVault, exampleTokenReceiver)
	if err != nil {
		log.Fatal(err)
	}

	curuserFlowExampleBalance, err := tc.FlowContractAccountClient.QueryTokenBalance(context.Background(),
		eth.Hex2Addr(tc.FlowUserAddr.String()), exampleTokenBalance)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("curuserFlowExampleBalance:%s", curuserFlowExampleBalance)
}

func removeFlowCfg() {
	for i := 0; i < len(tc.ValEthKs); i++ {
		cbrCfgPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/cbridge.toml", i)
		cbrViper := viper.New()
		cbrViper.SetConfigFile(cbrCfgPath)
		err := cbrViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		multichains := cbrViper.Get("multichain").([]interface{})
		cbrViper.Set("multichain", multichains[:3])
		err = cbrViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")
	}
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

func waitGethStart(gethrpc string) {
	// wait for chain1 to start first
	conn, _ := ethclient.Dial(gethrpc)
	for i := 1; i < 10; i++ {
		head, err := conn.HeaderByNumber(context.Background(), nil)
		if err == nil && head.Number.Uint64() > 1 {
			break
		}
		tc.Sleep(2)
	}
}
