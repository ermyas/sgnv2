package singlenode

import (
	"context"
	"io/ioutil"
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

func setupNewSgnEnv(contractParams *tc.ContractParams, testName string) []tc.Killable {
	if contractParams == nil {
		contractParams = &tc.ContractParams{
			CelrAddr:              tc.CelrAddr,
			ProposalDeposit:       big.NewInt(1),
			VotePeriod:            big.NewInt(1),
			UnbondingPeriod:       big.NewInt(50),
			MaxBondedValidators:   big.NewInt(11),
			MinValidatorTokens:    big.NewInt(1e18),
			MinSelfDelegation:     big.NewInt(1e18),
			AdvanceNoticePeriod:   big.NewInt(1),
			ValidatorBondInterval: big.NewInt(0),
		}
	}
	tx := tc.DeploySgnStakingContracts(contractParams)
	tc.WaitMinedWithChk(context.Background(), tc.EthClient, tx, tc.BlockDelay, tc.PollingInterval, "DeploySgnStakingContracts")

	updateSgnConfig()

	sgnProc, err := startSgnchain("", testName)
	tc.ChkErr(err, "start sgnchain")

	killable := []tc.Killable{sgnProc}
	if contractParams.StartGateway {
		gatewayProc, err := StartGateway("", testName)
		tc.ChkErr(err, "start gateway")
		killable = append(killable, gatewayProc)
	}

	return killable
}

func updateSgnConfig() {
	log.Infoln("Updating sgn.toml")

	configFilePath := "../../data/.sgncli/config/sgn.toml"
	configFileViper := viper.New()
	configFileViper.SetConfigFile(configFilePath)
	err := configFileViper.ReadInConfig()
	tc.ChkErr(err, "failed to read config")

	keystore, err := filepath.Abs("../../keys/vethks0.json")
	tc.ChkErr(err, "get keystore path")

	configFileViper.Set(common.FlagEthGateway, tc.LocalGeth)
	configFileViper.Set(common.FlagEthContractCelr, tc.CelrAddr.Hex())
	configFileViper.Set(common.FlagEthContractStaking, tc.Contracts.Staking.Address.Hex())
	configFileViper.Set(common.FlagEthContractSgn, tc.Contracts.Sgn.Address.Hex())
	configFileViper.Set(common.FlagEthContractReward, tc.Contracts.Reward.Address.Hex())
	configFileViper.Set(common.FlagEthContractViewer, tc.Contracts.Viewer.Address.Hex())
	configFileViper.Set(common.FlagEthContractGovern, tc.Contracts.Govern.Address.Hex())
	configFileViper.Set(common.FlagEthSignerKeystore, keystore)
	// TODO: different config for validator and signer
	ksbytes, err := ioutil.ReadFile(keystore)
	tc.ChkErr(err, "failed to read keystore config")
	ksAddrStr, err := eth.GetAddressFromKeystore(ksbytes)
	tc.ChkErr(err, "failed get addr from keystore")

	configFileViper.Set(common.FlagEthValidatorAddress, ksAddrStr)
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
