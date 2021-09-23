package multinode

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/transactor"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
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
	tc.CelrAddr, tc.CelrContract = tc.DeployCelrContract(tc.EthClient, tc.EtherBaseAuth)

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

	log.Infoln("set up mainchain")
	tc.SetupEthClient2()
	tc.CelrAddr2, tc.CelrContract2 = tc.DeployCelrContract(tc.EthClient2, tc.EtherBaseAuth2)

	// fund CELR to each eth account
	log.Infoln("fund each test addr 10 million CELR")
	err = tc.FundAddrsErc20(tc.CelrAddr, addrs, tc.NewBigInt(1, 25), tc.EthClient2, tc.EtherBaseAuth2)
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
		configFileViper.Set(common.FlagEthContractReward, tc.Contracts.Reward.Address.Hex())
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
			genesisViper.Set("app_state.govern.voting_params.voting_period", "120s")
		} else {
			genesisViper.Set("app_state.govern.voting_params.voting_period", "10s")
		}
		err = genesisViper.WriteConfig()
		tc.ChkErr(err, "Failed to write genesis")

	}

	// Update global viper
	node0ConfigPath := "../../../docker-volumes/node0/sgnd/config/sgn.toml"
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

func SetupValidators(t *testing.T, transactor *transactor.Transactor, amts []*big.Int) {
	var expVals stakingtypes.Validators
	log.Infoln("---------- It should add bonded validators successfully ----------")
	for i := 0; i < len(amts); i++ {
		log.Infoln("Adding validator", i, tc.ValEthAddrs[i].Hex())
		err := tc.InitializeValidator(tc.ValAuths[i], tc.ValSignerAddrs[i], tc.ValSgnAddrs[i], amts[i], eth.CommissionRate(0.02))
		require.NoError(t, err, "failed to initialize validator")
		tc.Sleep(5)
		expVal := stakingtypes.Validator{
			EthAddress:      eth.Addr2Hex(tc.ValEthAddrs[i]),
			EthSigner:       eth.Addr2Hex(tc.ValSignerAddrs[i]),
			Status:          eth.Bonded,
			SgnAddress:      tc.ValSgnAddrs[i].String(),
			Tokens:          sdk.NewIntFromBigInt(amts[i]),
			DelegatorShares: sdk.NewIntFromBigInt(amts[i]),
			CommissionRate:  sdk.NewDecWithPrec(2, 2),
		}
		expVals = append(expVals, expVal)
		tc.CheckValidators(t, transactor, expVals)
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
