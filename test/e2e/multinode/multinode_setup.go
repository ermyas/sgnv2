package multinode

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/farming/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/viper"
)

func SetupMainchain() {
	repoRoot, _ := filepath.Abs("../../..")
	log.Infoln("make localnet-down")
	cmd := exec.Command("make", "localnet-down")
	cmd.Dir = repoRoot
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	log.Infoln("build dockers, get geth, build sgn binary")
	cmd = exec.Command("make", "prepare-docker-env")
	cmd.Dir = repoRoot
	if err := cmd.Run(); err != nil {
		log.Error(err)
	}

	log.Infoln("start geth container")
	cmd = exec.Command("make", "localnet-start-geth")
	cmd.Dir = repoRoot
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	tc.SleepWithLog(5, "geth start")

	// set up mainchain: deploy contracts, fund addrs, etc
	addrs := []eth.Addr{
		tc.ValEthAddrs[0],
		tc.ValEthAddrs[1],
		tc.ValEthAddrs[2],
		tc.ValEthAddrs[3],
		tc.ValSignerAddrs[0],
		tc.ValSignerAddrs[1],
		tc.ValSignerAddrs[2],
		tc.ValSignerAddrs[3],
		tc.DelEthAddrs[0],
		tc.DelEthAddrs[1],
		tc.DelEthAddrs[2],
		tc.DelEthAddrs[3],
		tc.ClientEthAddrs[0],
		tc.ClientEthAddrs[1],
	}
	log.Infoln("fund each test addr 100 ETH")
	err := tc.FundAddrsETH(addrs, tc.NewBigInt(1, 20), tc.LocalGeth, int64(tc.ChainID))
	tc.ChkErr(err, "fund each test addr 100 ETH")

	log.Infoln("set up mainchain")
	tc.SetupEthClients()
	tc.CelrAddr, tc.CelrContract = tc.DeployERC20Contract(tc.EthClient, tc.EtherBaseAuth, "Celer", "CELR", 18)

	// fund CELR to each eth account
	log.Infoln("fund each test addr 10 million CELR")
	err = tc.FundAddrsErc20(tc.CelrAddr, addrs, tc.NewBigInt(1, 25), tc.EthClient, tc.EtherBaseAuth)
	tc.ChkErr(err, "fund each test addr 10 million CELR")
}

// should be invoked after mainchain 1 setup
func SetupMainchain2ForBridge() {
	repoRoot, _ := filepath.Abs("../../..")
	log.Infoln("prepare geth2 env")
	cmd := exec.Command("make", "prepare-geth2-env")
	cmd.Dir = repoRoot
	if err := cmd.Run(); err != nil {
		log.Error(err)
	}

	log.Infoln("start geth2 container")
	cmd = exec.Command("make", "localnet-start-geth2")
	cmd.Dir = repoRoot
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	tc.SleepWithLog(5, "geth2 start")

	// set up mainchain: deploy contracts, fund addrs, etc
	addrs := []eth.Addr{
		tc.ValEthAddrs[0],
		tc.ValEthAddrs[1],
		tc.ValEthAddrs[2],
		tc.ValEthAddrs[3],
		tc.ValSignerAddrs[0],
		tc.ValSignerAddrs[1],
		tc.ValSignerAddrs[2],
		tc.ValSignerAddrs[3],
		tc.DelEthAddrs[0],
		tc.DelEthAddrs[1],
		tc.DelEthAddrs[2],
		tc.DelEthAddrs[3],
		tc.ClientEthAddrs[0],
		tc.ClientEthAddrs[1],
	}
	log.Infoln("fund each test addr 100 ETH")
	err := tc.FundAddrsETH(addrs, tc.NewBigInt(1, 20), tc.LocalGeth2, int64(tc.Geth2ChainID))
	tc.ChkErr(err, "fund each test addr 100 ETH")

	log.Infoln("set up mainchain2")
	tc.SetupEthClient2()
}

func SetupNewSgnEnv(contractParams *tc.ContractParams, manual bool, cbridgeTest bool) {
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

	log.Infoln("make localnet-down-nodes")
	cmd := exec.Command("make", "localnet-down-nodes")
	repoRoot, _ := filepath.Abs("../../..")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	tc.ChkErr(err, "Failed to make localnet-down-nodes")

	log.Infoln("make prepare-sgn-data")
	cmd = exec.Command("make", "prepare-sgn-data")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		cmd = exec.Command("make", "prepare-sgn-data-sudo")
		cmd.Dir = repoRoot
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
	}
	tc.ChkErr(err, "Failed to make prepare-sgn-data")

	log.Infoln("Updating config files of SGN nodes")
	for i := 0; i < len(tc.ValEthKs); i++ {
		configPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/sgn.toml", i)
		configFileViper := viper.New()
		configFileViper.SetConfigFile(configPath)
		err = configFileViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		configFileViper.Set(common.FlagEthContractCelr, tc.CelrAddr.Hex())
		configFileViper.Set(common.FlagEthContractStaking, tc.Contracts.Staking.Address.Hex())
		configFileViper.Set(common.FlagEthContractSgn, tc.Contracts.Sgn.Address.Hex())
		configFileViper.Set(common.FlagEthContractStakingReward, tc.Contracts.StakingReward.Address.Hex())
		configFileViper.Set(common.FlagEthContractFarmingRewards, tc.Contracts.FarmingRewards.Address.Hex())
		configFileViper.Set(common.FlagEthContractViewer, tc.Contracts.Viewer.Address.Hex())
		configFileViper.Set(common.FlagEthContractGovern, tc.Contracts.Govern.Address.Hex())
		configFileViper.Set(common.FlagEthValidatorAddress, tc.ValEthAddrs[i].Hex())
		err = configFileViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err = genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		if manual {
			genesisViper.Set("app_state.gov.voting_params.voting_period", "120s")
		} else {
			genesisViper.Set("app_state.gov.voting_params.voting_period", "10s")
		}
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}

	if cbridgeTest {
		DeployUsdtForBridge()
		DeployBridgeContract()
		CreateFarmingPools()
	}

	// Update global viper
	node0ConfigPath := "../../../docker-volumes/node0/sgnd/config/sgn.toml"
	viper.SetConfigFile(node0ConfigPath)
	err = viper.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")
	viper.Set(common.FlagEthContractCelr, tc.CelrAddr.Hex())
	viper.Set(common.FlagEthContractStaking, tc.Contracts.Staking.Address.Hex())
	viper.Set(common.FlagEthContractSgn, tc.Contracts.Sgn.Address.Hex())
	viper.Set(common.FlagEthContractStakingReward, tc.Contracts.StakingReward.Address.Hex())
	viper.Set(common.FlagEthContractFarmingRewards, tc.Contracts.FarmingRewards.Address.Hex())
	viper.Set(common.FlagEthContractViewer, tc.Contracts.Viewer.Address.Hex())
	viper.Set(common.FlagEthContractGovern, tc.Contracts.Govern.Address.Hex())

	log.Infoln("make localnet-up-nodes")
	cmd = exec.Command("make", "localnet-up-nodes")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	tc.ChkErr(err, "Failed to make localnet-up-nodes")
}

func DeployUsdtForBridge() {
	addrs := []eth.Addr{
		tc.ValEthAddrs[0],
		tc.ValEthAddrs[1],
		tc.ValEthAddrs[2],
		tc.ValEthAddrs[3],
		tc.ValSignerAddrs[0],
		tc.ValSignerAddrs[1],
		tc.ValSignerAddrs[2],
		tc.ValSignerAddrs[3],
		tc.DelEthAddrs[0],
		tc.DelEthAddrs[1],
		tc.DelEthAddrs[2],
		tc.DelEthAddrs[3],
		tc.ClientEthAddrs[0],
		tc.ClientEthAddrs[1],
	}

	tc.CbrClient1.USDTAddr, tc.CbrClient1.USDTContract = tc.DeployERC20Contract(tc.CbrClient1.Ec, tc.CbrClient1.Auth, "USDT", "USDT", 6)
	tc.CbrClient2.USDTAddr, tc.CbrClient2.USDTContract = tc.DeployERC20Contract(tc.CbrClient2.Ec, tc.CbrClient2.Auth, "USDT", "USDT", 6)

	// fund usdt to each eth account
	log.Infoln("fund each test addr 10 million usdt on each chain")
	err := tc.FundAddrsErc20(tc.CbrClient1.USDTAddr, addrs, tc.NewBigInt(1, 13), tc.CbrClient1.Ec, tc.CbrClient1.Auth)
	tc.ChkErr(err, "fund each test addr 10 million usdt on chain 1")
	err = tc.FundAddrsErc20(tc.CbrClient2.USDTAddr, addrs, tc.NewBigInt(1, 13), tc.CbrClient2.Ec, tc.CbrClient2.Auth)
	tc.ChkErr(err, "fund each test addr 10 million usdt on chain 2")

	log.Infoln("Updating config files of SGN nodes")
	for i := 0; i < len(tc.ValEthKs); i++ {
		genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
		genesisViper := viper.New()
		genesisViper.SetConfigFile(genesisPath)
		err = genesisViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read genesis")
		cbrConfig := new(cbrtypes.CbrConfig)
		jsonByte, _ := json.Marshal(genesisViper.Get("app_state.cbridge.config"))
		json.Unmarshal(jsonByte, cbrConfig)
		cbrConfig.Assets[0].Addr = eth.Addr2Hex(tc.CbrClient1.USDTAddr)
		cbrConfig.Assets[1].Addr = eth.Addr2Hex(tc.CbrClient2.USDTAddr)
		genesisViper.Set("app_state.cbridge.config", cbrConfig)
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
	}
}

func DeployBridgeContract() {
	tc.CbrClient1.CbrAddr, tc.CbrClient1.CbrContract = tc.DeployBridgeContract(tc.CbrClient1.Ec, tc.CbrClient1.Auth, nil)
	tc.CbrClient2.CbrAddr, tc.CbrClient2.CbrContract = tc.DeployBridgeContract(tc.CbrClient2.Ec, tc.CbrClient2.Auth, nil)

	for i := 0; i < len(tc.ValEthKs); i++ {
		configPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/sgn.toml", i)
		configFileViper := viper.New()
		configFileViper.SetConfigFile(configPath)
		err := configFileViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		multichains := configFileViper.Get("multichain").([]interface{})
		multichains[0].(map[string]interface{})["cbridge"] = eth.Addr2Hex(tc.CbrClient1.CbrAddr)
		multichains[1].(map[string]interface{})["cbridge"] = eth.Addr2Hex(tc.CbrClient2.CbrAddr)
		configFileViper.Set("multichain", multichains)
		err = configFileViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")
	}
}

func CreateFarmingPools() {
	tc.CbrClient1.FarmingRewardsContract = tc.Contracts.FarmingRewards
	tc.CbrClient2.FarmingRewardsContract = tc.Contracts.FarmingRewards

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

		// Add a pool
		poolName := "cbridge-USDT/883"
		var pools farmingtypes.FarmingPools
		pool := farmingtypes.NewFarmingPool(
			poolName,
			farmingtypes.ERC20Token{
				ChainId: 883,
				Symbol:  "CB-USDT",
				Address: eth.Addr2Hex(tc.CbrClient1.USDTAddr),
			},
			[]farmingtypes.ERC20Token{
				{
					ChainId: 883,
					Symbol:  "CELR",
					Address: eth.Addr2Hex(tc.CelrAddr),
				},
			},
			sdk.NewDecCoin("CB-USDT/883", sdk.ZeroInt()),
			[]farmingtypes.RewardTokenInfo{
				{
					RemainingAmount:        sdk.NewDecCoin("CELR/883", sdk.NewInt(10000).Mul(sdk.NewInt(1e18))),
					RewardStartBlockHeight: 1,
					RewardAmountPerBlock:   sdk.NewDec(10),
				},
			},
			sdk.NewDecCoins(sdk.NewDecCoin("CELR/883", sdk.ZeroInt())),
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
			Rewards:  types.NewPoolCurrentRewards(0, 1, sdk.DecCoins{}),
		}
		genesisViper.Set(
			"app_state.farming.pool_historical_rewards",
			[]farmingtypes.PoolHistoricalRewardsRecord{poolHistoricalRewardsRecord})
		genesisViper.Set(
			"app_state.farming.pool_current_rewards",
			[]farmingtypes.PoolCurrentRewardsRecord{poolCurrentRewardsRecord})

		// Fund reward module account
		rewardCoins := sdk.NewCoins(sdk.NewCoin("CELR/883", sdk.NewInt(10000).Mul(sdk.NewInt(1e18))))
		var balances []banktypes.Balance
		jsonByte, _ := json.Marshal(genesisViper.Get("app_state.bank.balances"))
		json.Unmarshal(jsonByte, &balances)
		balances = append(balances, banktypes.Balance{
			Address: authtypes.NewModuleAddress(farmingtypes.RewardModuleAccountName).String(),
			Coins:   rewardCoins,
		})
		genesisViper.Set("app_state.bank.balances", balances)

		// Change genesis supply
		var supply sdk.Coins
		jsonByte, _ = json.Marshal(genesisViper.Get("app_state.bank.supply"))
		json.Unmarshal(jsonByte, &supply)
		supply = supply.Add(rewardCoins...)
		genesisViper.Set("app_state.bank.supply", supply)

		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")
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
