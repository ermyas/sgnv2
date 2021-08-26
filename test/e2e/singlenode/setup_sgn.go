package singlenode

import (
	"context"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"
)

func setupNewSGNEnv(sgnParams *tc.SGNParams, testName string) []tc.Killable {
	if sgnParams == nil {
		sgnParams = &tc.SGNParams{
			CelrAddr:              tc.E2eProfile.CelrAddr,
			GovernProposalDeposit: big.NewInt(1),
			GovernVoteTimeout:     big.NewInt(1),
			SlashTimeout:          big.NewInt(50),
			MaxBondedValidators:   big.NewInt(11),
			MinValidatorTokens:    big.NewInt(1e18),
			MinSelfDelegation:     big.NewInt(1e18),
			AdvanceNoticePeriod:   big.NewInt(1),
			ValidatorBondInterval: big.NewInt(0),
		}
	}
	var tx *types.Transaction
	tx, tc.E2eProfile.StakingContractAddr, tc.E2eProfile.SgnContractAddr = tc.DeployStakingSGNContracts(sgnParams)
	tc.WaitMinedWithChk(context.Background(), tc.EthClient, tx, tc.BlockDelay, tc.PollingInterval, "DeployStakingSGNContracts")

	updateSGNConfig()

	sgnProc, err := startSgnchain("", testName)
	tc.ChkErr(err, "start sgnchain")
	tc.SetContracts(tc.E2eProfile.StakingContractAddr, tc.E2eProfile.SgnContractAddr)

	killable := []tc.Killable{sgnProc}
	if sgnParams.StartGateway {
		gatewayProc, err := StartGateway("", testName)
		tc.ChkErr(err, "start gateway")
		killable = append(killable, gatewayProc)
	}

	return killable
}

func updateSGNConfig() {
	log.Infoln("Updating sgn.toml")

	configFilePath := "../../data/.sgncli/config/sgn.toml"
	configFileViper := viper.New()
	configFileViper.SetConfigFile(configFilePath)
	err := configFileViper.ReadInConfig()
	tc.ChkErr(err, "failed to read config")

	keystore, err := filepath.Abs("../../keys/vethks0.json")
	tc.ChkErr(err, "get keystore path")

	configFileViper.Set(common.FlagEthGateway, tc.LocalGeth)
	configFileViper.Set(common.FlagEthContractCelr, tc.E2eProfile.CelrAddr.Hex())
	configFileViper.Set(common.FlagEthContractStaking, tc.E2eProfile.StakingContractAddr.Hex())
	configFileViper.Set(common.FlagEthContractSgn, tc.E2eProfile.SgnContractAddr.Hex())
	configFileViper.Set(common.FlagEthSignerKeystore, keystore)
	err = configFileViper.WriteConfig()
	tc.ChkErr(err, "failed to write config")
	// Update global viper
	viper.SetConfigFile(configFilePath)
	err = viper.ReadInConfig()
	tc.ChkErr(err, "failed to read config")
}

func installSgn() error {
	cmd := exec.Command("make", "install")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "WITH_CLEVELDB=yes")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// set cmd.Dir under repo root path
	cmd.Dir, _ = filepath.Abs("../../..")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("cp", "./test/data/.sgncli/config/sgn_template.toml", "./test/data/.sgncli/config/sgn.toml")
	// set cmd.Dir under repo root path
	cmd.Dir, _ = filepath.Abs("../../..")
	return cmd.Run()
}

// startSgnchain starts sgn sgnchain with the data in test/data
func startSgnchain(rootDir, testName string) (*os.Process, error) {
	cmd := exec.Command("make", "update-test-data")
	// set cmd.Dir under repo root path
	cmd.Dir, _ = filepath.Abs("../../..")
	if err := cmd.Run(); err != nil {
		log.Errorln("Failed to run \"make update-test-data\": ", err)
		return nil, err
	}

	cmd = exec.Command("sgnd", "start")
	cmd.Dir, _ = filepath.Abs("../../..")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Errorln("Failed to run \"sgnd start\": ", err)
		return nil, err
	}

	log.Infoln("sgn pid:", cmd.Process.Pid)
	return cmd.Process, nil
}

func StartGateway(rootDir, testName string) (*os.Process, error) {
	cmd := exec.Command("sgncli", "gateway")
	cmd.Dir, _ = filepath.Abs("../../..")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	log.Infoln("gateway pid:", cmd.Process.Pid)
	return cmd.Process, nil
}
