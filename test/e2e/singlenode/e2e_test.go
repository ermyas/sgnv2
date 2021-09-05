// Setup eth mainchain and sgn sgnchain etc for e2e tests
package singlenode

import (
	"flag"
	"fmt"
	"os"
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
	ethProc, err := startMainchain(outRootDir)
	tc.ChkErr(err, "starting mainchain")
	tc.SleepWithLog(2, "starting mainchain")

	// set up mainchain: deploy contracts and fund ethpool etc
	// first fund each account 100 ETH
	addrs := []eth.Addr{
		tc.ValEthAddrs[0],
		tc.DelEthAddrs[0],
		tc.DelEthAddrs[1],
		tc.DelEthAddrs[2],
		tc.DelEthAddrs[3],
		tc.ClientEthAddrs[0],
		tc.ClientEthAddrs[1],
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
