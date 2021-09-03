package multinode

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
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
		eth.Hex2Addr(tc.ValEthAddrs[0]),
		eth.Hex2Addr(tc.ValEthAddrs[1]),
		eth.Hex2Addr(tc.ValEthAddrs[2]),
		eth.Hex2Addr(tc.ValEthAddrs[3]),
		eth.Hex2Addr(tc.DelEthAddrs[0]),
		eth.Hex2Addr(tc.DelEthAddrs[1]),
		eth.Hex2Addr(tc.DelEthAddrs[2]),
		eth.Hex2Addr(tc.DelEthAddrs[3]),
		eth.Hex2Addr(tc.ClientEthAddrs[0]),
		eth.Hex2Addr(tc.ClientEthAddrs[1]),
	}
	log.Infoln("fund each test addr 100 ETH")
	err := tc.FundAddrsETH(addrs, tc.NewBigInt(1, 20))
	tc.ChkErr(err, "fund each test addr 100 ETH")

	log.Infoln("set up mainchain")
	tc.SetupEthClients()
	tc.DeployCelrContract()

	// fund CELR to each eth account
	log.Infoln("fund each test addr 10 million CELR")
	err = tc.FundAddrsErc20(tc.CelrAddr, addrs, tc.NewBigInt(1, 25))
	tc.ChkErr(err, "fund each test addr 10 million CELR")
}

func SetupNewSgnEnv(contractParams *tc.ContractParams, manual bool) {
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
	tc.ChkErr(err, "Failed to make prepare-sgn-data")

	log.Infoln("Updating config files of SGN nodes")
	for i := 0; i < len(tc.ValEthKs); i++ {
		configPath := fmt.Sprintf("../../../docker-volumes/node%d/sgncli/config/sgn.toml", i)
		configFileViper := viper.New()
		configFileViper.SetConfigFile(configPath)
		err = configFileViper.ReadInConfig()
		tc.ChkErr(err, "Failed to read config")
		configFileViper.Set(common.FlagEthContractCelr, tc.CelrAddr.Hex())
		configFileViper.Set(common.FlagEthContractStaking, tc.Contracts.Staking.Address.Hex())
		configFileViper.Set(common.FlagEthContractSgn, tc.Contracts.Sgn.Address.Hex())
		configFileViper.Set(common.FlagEthContractReward, tc.Contracts.Reward.Address.Hex())
		configFileViper.Set(common.FlagEthContractViewer, tc.Contracts.Viewer.Address.Hex())
		configFileViper.Set(common.FlagEthContractGovern, tc.Contracts.Govern.Address.Hex())
		// TODO: different config for validator and signer
		configFileViper.Set(common.FlagEthValidatorAddress, tc.ValEthAddrs[i])
		err = configFileViper.WriteConfig()
		tc.ChkErr(err, "Failed to write config")

		if manual {
			genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
			genesisViper := viper.New()
			genesisViper.SetConfigFile(genesisPath)
			err = genesisViper.ReadInConfig()
			tc.ChkErr(err, "Failed to read genesis")
			genesisViper.Set("app_state.govern.voting_params.voting_period", "120000000000")
			err = genesisViper.WriteConfig()
			tc.ChkErr(err, "Failed to write genesis")
		}
	}

	// Update global viper
	node0ConfigPath := "../../../docker-volumes/node0/sgncli/config/sgn.toml"
	viper.SetConfigFile(node0ConfigPath)
	err = viper.ReadInConfig()
	tc.ChkErr(err, "Failed to read config")
	viper.Set(common.FlagEthContractCelr, tc.CelrAddr.Hex())
	viper.Set(common.FlagEthContractStaking, tc.Contracts.Staking.Address.Hex())
	viper.Set(common.FlagEthContractSgn, tc.Contracts.Sgn.Address.Hex())
	viper.Set(common.FlagEthContractReward, tc.Contracts.Reward.Address.Hex())
	viper.Set(common.FlagEthContractViewer, tc.Contracts.Viewer.Address.Hex())
	viper.Set(common.FlagEthContractGovern, tc.Contracts.Govern.Address.Hex())

	tc.ChkErr(err, "Failed to SetContracts")

	log.Infoln("make localnet-up-nodes")
	cmd = exec.Command("make", "localnet-up-nodes")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	tc.ChkErr(err, "Failed to make localnet-up-nodes")
}
