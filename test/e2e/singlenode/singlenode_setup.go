package singlenode

import (
	"context"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

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
			ProposalDeposit:       big.NewInt(1e18),
			VotePeriod:            big.NewInt(10),
			UnbondingPeriod:       big.NewInt(10),
			MaxBondedValidators:   big.NewInt(11),
			MinValidatorTokens:    big.NewInt(2e18),
			MinSelfDelegation:     big.NewInt(1e18),
			AdvanceNoticePeriod:   big.NewInt(10),
			ValidatorBondInterval: big.NewInt(0),
			MaxSlashFactor:        big.NewInt(1e5),
		}
	}
	tx := tc.DeploySgnStakingContracts(contractParams)
	tc.WaitMinedWithChk(context.Background(), tc.EthClient, tx, tc.BlockDelay, tc.PollingInterval, "DeploySgnStakingContracts")

	updateSgnConfig()

	sgnProc, err := startSgnChain("", testName)
	tc.ChkErr(err, "start sgnchain")

	killable := []tc.Killable{sgnProc}
	// if contractParams.StartGateway {
	// 	gatewayProc, err := StartGateway("", testName)
	// 	tc.ChkErr(err, "start gateway")
	// 	killable = append(killable, gatewayProc)
	// }

	return killable
}

func updateSgnConfig() {
	log.Infoln("Updating sgn.toml")

	configFilePath := "../../data/.sgnd/config/sgn.toml"
	configFileViper := viper.New()
	configFileViper.SetConfigFile(configFilePath)
	err := configFileViper.ReadInConfig()
	tc.ChkErr(err, "failed to read config")

	keystore, err := filepath.Abs("../../keys/vsigner0.json")
	tc.ChkErr(err, "get keystore path")

	configFileViper.Set(common.FlagEthGateway, tc.LocalGeth)
	configFileViper.Set(common.FlagEthContractCelr, tc.CelrAddr.Hex())
	configFileViper.Set(common.FlagEthContractStaking, tc.Contracts.Staking.Address.Hex())
	configFileViper.Set(common.FlagEthContractSgn, tc.Contracts.Sgn.Address.Hex())
	configFileViper.Set(common.FlagEthContractReward, tc.Contracts.Reward.Address.Hex())
	configFileViper.Set(common.FlagEthContractViewer, tc.Contracts.Viewer.Address.Hex())
	configFileViper.Set(common.FlagEthContractGovern, tc.Contracts.Govern.Address.Hex())
	configFileViper.Set(common.FlagEthSignerKeystore, keystore)
	configFileViper.Set(common.FlagEthValidatorAddress, eth.Addr2Hex(tc.ValEthAddrs[0]))
	err = configFileViper.WriteConfig()
	tc.ChkErr(err, "failed to write config")
	// Update global viper
	viper.SetConfigFile(configFilePath)
	err = viper.ReadInConfig()
	tc.ChkErr(err, "failed to read config")
}

func installSgnd() error {
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

	cmd = exec.Command("cp", "./test/data/.sgnd/config/sgn_template.toml", "./test/data/.sgnd/config/sgn.toml")
	// set cmd.Dir under repo root path
	cmd.Dir, _ = filepath.Abs("../../..")
	return cmd.Run()
}

// startSgnChain starts SGN chain with the data in test/data
func startSgnChain(rootDir, testName string) (*os.Process, error) {
	cmd := exec.Command("make", "update-test-data")
	// set cmd.Dir under repo root path
	cmd.Dir, _ = filepath.Abs("../../..")
	if err := cmd.Run(); err != nil {
		log.Errorln("Failed to run \"make update-test-data\": ", err)
		return nil, err
	}

	genesisPath := os.ExpandEnv("$HOME/.sgnd/config/genesis.json")
	genesisViper := viper.New()
	genesisViper.SetConfigFile(genesisPath)
	err := genesisViper.ReadInConfig()
	tc.ChkErr(err, "Failed to read genesis")
	genesisViper.Set("app_state.govern.voting_params.voting_period", "10s")
	err = genesisViper.WriteConfig()
	tc.ChkErr(err, "Failed to write genesis")

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

// func StartGateway(rootDir, testName string) (*os.Process, error) {
// 	cmd := exec.Command("sgnd", "gateway")
// 	cmd.Dir, _ = filepath.Abs("../../..")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Start(); err != nil {
// 		return nil, err
// 	}

// 	log.Infoln("gateway pid:", cmd.Process.Pid)
// 	return cmd.Process, nil
// }

// start process to handle eth rpc, and fund etherbase and server account
func startMainchain(outRootDir string) (*os.Process, error) {
	log.Infoln("outRootDir", outRootDir, "envDir", tc.EnvDir)
	chainDataDir := outRootDir + "mainchaindata"
	logFname := outRootDir + "mainchain.log"
	if err := os.MkdirAll(chainDataDir, os.ModePerm); err != nil {
		return nil, err
	}

	// geth init
	cmdInit := exec.Command("geth", "--datadir", chainDataDir, "init", "mainchain_genesis.json")
	// set cmd.Dir because relative files are under testing/env
	cmdInit.Dir, _ = filepath.Abs(tc.EnvDir)
	if err := cmdInit.Run(); err != nil {
		return nil, err
	}

	// actually run geth, blocking. set syncmode full to avoid bloom mem cache by fast sync
	cmd := exec.Command("geth", "--networkid", strconv.Itoa(tc.ChainID), "--cache", "256", "--nousb", "--syncmode", "full", "--nodiscover", "--maxpeers", "0",
		"--netrestrict", "127.0.0.1/8", "--datadir", chainDataDir, "--keystore", "keystore", "--miner.gastarget", "8000000",
		"--ws", "--ws.addr", "localhost", "--ws.port", "8546", "--ws.api", "admin,debug,eth,miner,net,personal,shh,txpool,web3",
		"--mine", "--allow-insecure-unlock", "--unlock", "0xb5BB8b7f6f1883e0c01ffb8697024532e6F3238C", "--password", "empty_password.txt",
		"--http", "--http.corsdomain", "*", "--http.addr", "localhost", "--http.port", "8545", "--http.api",
		"admin,debug,eth,miner,net,personal,shh,txpool,web3")
	cmd.Dir = cmdInit.Dir

	logF, _ := os.Create(logFname)
	cmd.Stderr = logF
	cmd.Stdout = logF
	log.Infoln("ready to start cmd")
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	log.Infoln("geth pid:", cmd.Process.Pid)
	// in case geth exits with non-zero, exit test early
	// if geth is killed by ethProc.Signal, it exits w/ 0
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Errorln("geth process failed:", err)
			os.Exit(1)
		}
	}()
	return cmd.Process, nil
}
