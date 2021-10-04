package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/test/e2e/multinode"
	"github.com/celer-network/sgn-v2/transactor"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	"github.com/spf13/viper"
)

var (
	start   = flag.Bool("start", false, "start local testnet")
	cbr     = flag.Bool("cbr", false, "start with cbridge")
	auto    = flag.Bool("auto", false, "auto-add all validators")
	down    = flag.Bool("down", false, "shutdown local testnet")
	up      = flag.Int("up", -1, "start a testnet node")
	stop    = flag.Int("stop", -1, "stop a testnet node")
	upall   = flag.Bool("upall", false, "start all nodes")
	stopall = flag.Bool("stopall", false, "stop all nodes")
	rebuild = flag.Bool("rebuild", false, "rebuild sgn node docker image")
)

func main() {
	flag.Parse()
	repoRoot, _ := filepath.Abs("../../..")
	if *start {
		multinode.SetupMainchain()
		if *cbr {
			multinode.SetupMainchain2ForBridge()
		}
		tc.SetupSgnchain()

		p := &tc.ContractParams{
			CelrAddr:              tc.CelrAddr,
			ProposalDeposit:       big.NewInt(1e17),
			VotePeriod:            big.NewInt(30),
			UnbondingPeriod:       big.NewInt(5),
			MaxBondedValidators:   big.NewInt(5),
			MinValidatorTokens:    big.NewInt(5e18),
			MinSelfDelegation:     big.NewInt(2e18),
			AdvanceNoticePeriod:   big.NewInt(1),
			ValidatorBondInterval: big.NewInt(0),
			MaxSlashFactor:        big.NewInt(1e5),
		}
		multinode.SetupNewSgnEnv(p, true, *cbr)
		if *cbr {
			amts := []*big.Int{big.NewInt(1e18)}
			tc.CbrClient1.SetInitSigners(amts)
			tc.CbrClient2.SetInitSigners(amts)
		}

		log.Infoln("install sgnd in host machine")
		cmd := exec.Command("make", "install")
		cmd.Dir = repoRoot
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "WITH_CLEVELDB=yes")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}

		log.Infoln("copy config files")
		cmd2 := exec.Command("make", "copy-manual-test-data")
		cmd2.Dir = repoRoot
		if err := cmd2.Run(); err != nil {
			log.Error(err)
		}
		log.Infoln("update config files")
		for i := 0; i < len(tc.ValEthKs); i++ {
			configPath := fmt.Sprintf("./data/node%d/sgnd/config/sgn.toml", i)
			configFileViper := viper.New()
			configFileViper.SetConfigFile(configPath)
			if err := configFileViper.ReadInConfig(); err != nil {
				log.Error(err)
			}
			ksPath, _ := filepath.Abs(fmt.Sprintf("./data/node%d/keys/vsigner%d.json", i, i))
			configFileViper.Set(common.FlagEthSignerKeystore, ksPath)
			configFileViper.Set(common.FlagEthGateway, tc.LocalGeth)
			configFileViper.Set(common.FlagSgnNodeURI, tc.SgnNodeURIs[i])
			if err := configFileViper.WriteConfig(); err != nil {
				log.Error(err)
			}
		}
		if *auto {
			time.Sleep(10 * time.Second)
			addValidators()
		}
	} else if *down {
		log.Infoln("Tearing down all containers...")
		cmd := exec.Command("make", "localnet-down")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}
		os.Exit(0)
	} else if *up != -1 {
		log.Infoln("Start node", *up)
		cmd := exec.Command("docker-compose", "up", "-d", fmt.Sprintf("sgnnode%d", *up))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}
	} else if *stop != -1 {
		log.Infoln("Stop node", *stop)
		cmd := exec.Command("docker-compose", "stop", fmt.Sprintf("sgnnode%d", *stop))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}
	} else if *upall {
		log.Infoln("Start all nodes ...")
		cmd := exec.Command("make", "localnet-up-nodes")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}
		os.Exit(0)
	} else if *stopall {
		log.Infoln("Stop all nodes ...")
		cmd := exec.Command("make", "localnet-down-nodes")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}
		os.Exit(0)
	} else if *rebuild {
		log.Infoln("Rebuild sgn node docker image ...")
		cmd := exec.Command("make", "build-node")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}
		os.Exit(0)
	}
}

func addValidators() {
	encodingConfig := app.MakeEncodingConfig()
	txr, err := transactor.NewTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.SgnValAcct,
		tc.SgnPassphrase,
		encodingConfig.Amino,
		encodingConfig.Codec,
		encodingConfig.InterfaceRegistry,
	)
	if err != nil {
		log.Error(err)
		return
	}
	txr.Run()

	amts := []*big.Int{
		new(big.Int).Mul(big.NewInt(10000), big.NewInt(common.TokenDec)),
		new(big.Int).Mul(big.NewInt(20000), big.NewInt(common.TokenDec)),
		new(big.Int).Mul(big.NewInt(15000), big.NewInt(common.TokenDec)),
		new(big.Int).Mul(big.NewInt(18000), big.NewInt(common.TokenDec)),
	}
	commissions := []uint64{eth.CommissionRate(0.15), eth.CommissionRate(0.2), eth.CommissionRate(0.12), eth.CommissionRate(0.1)}

	for i := 0; i < 4; i++ {
		log.Infoln("Adding validator ", i, tc.ValEthAddrs[i].Hex())
		err := tc.InitializeValidator(tc.ValAuths[i], tc.ValSignerAddrs[i], tc.ValSgnAddrs[i], amts[i], commissions[i])
		if err != nil {
			log.Errorln("failed to initialize validator: ", err)
		}

		for retry := 0; retry < tc.RetryLimit; retry++ {
			v, err := stakingcli.QueryValidator(txr.CliCtx, eth.Addr2Hex(tc.ValEthAddrs[i]))
			if err == nil {
				log.Infof("query validator success: %s", v)
				break
			}
			time.Sleep(tc.RetryPeriod)
		}
	}
}
