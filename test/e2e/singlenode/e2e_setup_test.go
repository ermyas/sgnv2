// Setup eth mainchain and sgn sgnchain etc for e2e tests
package singlenode

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
)

var (
	CLIHome = os.ExpandEnv("$HOME/.sgncli")

	// root dir with ending / for all files, OutRootDirPrefix + epoch seconds
	// due to testframework etc in a different testing package, we have to define
	// same var in testframework.go and expose a set api
	outRootDir string
)

// TestMain handles common setup (start mainchain, deploy, start sgnchain etc)
// and teardown. Test specific setup should be done in TestXxx
func TestMain(m *testing.M) {
	flag.Parse()
	log.EnableColor()

	// mkdir out root
	outRootDir = fmt.Sprintf("%s%d/", tc.OutRootDirPrefix, time.Now().Unix())
	err := os.MkdirAll(outRootDir, os.ModePerm)
	tc.ChkErr(err, "creating root dir")
	log.Infoln("Using folder:", outRootDir)
	// set testing pkg level path
	// start geth, not waiting for it to be fully ready. also watch geth proc
	// if geth exits with non-zero, os.Exit(1)
	ethProc, err := startMainchain()
	tc.ChkErr(err, "starting mainchain")
	tc.SleepWithLog(2, "starting mainchain")

	// set up mainchain: deploy contracts and fund ethpool etc
	// first fund each account 100 ETH
	addrs := []eth.Addr{
		eth.Hex2Addr(tc.ValEthAddrs[0]),
		eth.Hex2Addr(tc.DelEthAddrs[0]),
		eth.Hex2Addr(tc.DelEthAddrs[1]),
		eth.Hex2Addr(tc.DelEthAddrs[2]),
		eth.Hex2Addr(tc.DelEthAddrs[3]),
		eth.Hex2Addr(tc.ClientEthAddrs[0]),
		eth.Hex2Addr(tc.ClientEthAddrs[1]),
	}
	err = tc.FundAddrsETH(addrs, tc.NewBigInt(1, 20))
	tc.ChkErr(err, "fund eth")
	tc.SetupEthClients()
	tc.DeployCelrContract()

	// fund CELR to each eth account
	log.Infoln("fund each addr 10 million CELR")
	err = tc.FundAddrsErc20(tc.CelrAddr, addrs, tc.NewBigInt(1, 25))
	tc.ChkErr(err, "fund each addr ERC20")

	// make install sgn and sgncli
	err = installSgn()
	tc.ChkErr(err, "installing sgn and sgncli")

	tc.SetupSgnchain()
	// run all e2e tests
	ret := m.Run()

	ethProc.Signal(syscall.SIGTERM)
	if ret == 0 {
		log.Infoln("All tests passed! 🎉🎉🎉")
		os.RemoveAll(outRootDir)
		os.Exit(0)
	} else {
		log.Errorln("Tests failed. 🔥🔥🔥")
		os.Exit(ret)
	}
}

// start process to handle eth rpc, and fund etherbase and server account
func startMainchain() (*os.Process, error) {
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
