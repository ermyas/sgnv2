package multinode

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/celer-network/goutils/log"
	tc "github.com/celer-network/sgn-v2/test/common"
)

// TestMain handles common setup (start mainchain, deploy, start sgnchain etc)
// and teardown. Test specific setup should be done in TestXxx
func TestMain(m *testing.M) {
	flag.Parse()
	log.EnableColor()
	repoRoot, _ := filepath.Abs("../../..")

	BuildDockers()
	tc.RunAllAndWait(SetupMainchain, SetupBridgeChains)
	tc.SetupSgnchain()
	SetupFlowChain()

	log.Infoln("run all e2e tests")
	ret := m.Run()

	if ret == 0 {
		log.Infoln("All tests passed! 🎉🎉🎉")
		log.Infoln("Tearing down all containers...")
		cmd := exec.Command("make", "localnet-down")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Error(err)
		}
		os.Exit(0)
	} else {
		log.Errorln("Tests failed. 🚧🚧🚧 Geth and sgn containers are still running for debug. 🚧🚧🚧 Run \"make localnet-down\" to stop them")
		os.Exit(ret)
	}
}
