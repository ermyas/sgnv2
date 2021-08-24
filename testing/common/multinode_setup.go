package common

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/core/types"
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
	SleepWithLog(5, "geth start")

	// set up mainchain: deploy contracts, fund addrs, etc
	addrs := []eth.Addr{
		eth.Hex2Addr(ValEthAddrs[0]),
		eth.Hex2Addr(ValEthAddrs[1]),
		eth.Hex2Addr(ValEthAddrs[2]),
		eth.Hex2Addr(ValEthAddrs[3]),
		eth.Hex2Addr(DelEthAddrs[0]),
		eth.Hex2Addr(DelEthAddrs[1]),
		eth.Hex2Addr(DelEthAddrs[2]),
		eth.Hex2Addr(DelEthAddrs[3]),
		eth.Hex2Addr(ClientEthAddrs[0]),
		eth.Hex2Addr(ClientEthAddrs[1]),
	}
	log.Infoln("fund each test addr 100 ETH")
	err := FundAddrsETH("1"+strings.Repeat("0", 20), addrs)
	ChkErr(err, "fund each test addr 100 ETH")

	log.Infoln("set up mainchain")
	SetupEthClients()
	SetupE2eProfile()

	// fund CELR to each eth account
	log.Infoln("fund each test addr 10 million CELR")
	err = FundAddrsErc20(E2eProfile.CelrAddr, addrs, "1"+strings.Repeat("0", 25))
	ChkErr(err, "fund each test addr 10 million CELR")
}

func SetupNewSGNEnv(sgnParams *SGNParams, manual bool) {
	log.Infoln("Deploy Staking and SGN contracts")
	if sgnParams == nil {
		sgnParams = &SGNParams{
			CelrAddr:               E2eProfile.CelrAddr,
			GovernProposalDeposit:  big.NewInt(1),
			GovernVoteTimeout:      big.NewInt(5),
			SlashTimeout:           big.NewInt(50),
			MaxBondedValidators:    big.NewInt(7),
			MinValidatorTokens:     big.NewInt(1e18),
			MinSelfDelegation:      big.NewInt(1e18),
			AdvanceNoticePeriod:    big.NewInt(1),
			ValidatorBondInterval:  big.NewInt(24 * 3600),
			SidechainGoLiveTimeout: big.NewInt(0),
		}
	}
	var tx *types.Transaction
	tx, E2eProfile.StakingContractAddr, E2eProfile.SgnContractAddr = DeployStakingSGNContracts(sgnParams)
	WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "DeployStakingSGNContracts")

	log.Infoln("make localnet-down-nodes")
	cmd := exec.Command("make", "localnet-down-nodes")
	repoRoot, _ := filepath.Abs("../../..")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	ChkErr(err, "Failed to make localnet-down-nodes")

	log.Infoln("make prepare-sgn-data")
	cmd = exec.Command("make", "prepare-sgn-data")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	ChkErr(err, "Failed to make prepare-sgn-data")

	log.Infoln("Updating config files of SGN nodes")
	for i := 0; i < len(ValEthKs); i++ {
		configPath := fmt.Sprintf("../../../docker-volumes/node%d/sgncli/config/sgn.toml", i)
		configFileViper := viper.New()
		configFileViper.SetConfigFile(configPath)
		err = configFileViper.ReadInConfig()
		ChkErr(err, "Failed to read config")
		configFileViper.Set(common.FlagEthContractCelr, E2eProfile.CelrAddr.Hex())
		configFileViper.Set(common.FlagEthContractStaking, E2eProfile.StakingContractAddr.Hex())
		configFileViper.Set(common.FlagEthContractSgn, E2eProfile.SgnContractAddr.Hex())
		err = configFileViper.WriteConfig()
		ChkErr(err, "Failed to write config")

		if manual {
			genesisPath := fmt.Sprintf("../../../docker-volumes/node%d/sgnd/config/genesis.json", i)
			genesisViper := viper.New()
			genesisViper.SetConfigFile(genesisPath)
			err = genesisViper.ReadInConfig()
			ChkErr(err, "Failed to read genesis")
			genesisViper.Set("app_state.govern.voting_params.voting_period", "120000000000")
			err = genesisViper.WriteConfig()
			ChkErr(err, "Failed to write genesis")
		}
	}

	// Update global viper
	node0ConfigPath := "../../../docker-volumes/node0/sgncli/config/sgn.toml"
	viper.SetConfigFile(node0ConfigPath)
	err = viper.ReadInConfig()
	ChkErr(err, "Failed to read config")
	viper.Set(common.FlagEthContractCelr, E2eProfile.CelrAddr.Hex())
	viper.Set(common.FlagEthContractStaking, E2eProfile.StakingContractAddr.Hex())
	viper.Set(common.FlagEthContractSgn, E2eProfile.SgnContractAddr.Hex())

	err = SetContracts(E2eProfile.StakingContractAddr, E2eProfile.SgnContractAddr)
	ChkErr(err, "Failed to SetContracts")

	log.Infoln("make localnet-up-nodes")
	cmd = exec.Command("make", "localnet-up-nodes")
	cmd.Dir = repoRoot
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	ChkErr(err, "Failed to make localnet-up-nodes")
}
