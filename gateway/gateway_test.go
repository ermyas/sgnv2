package gateway

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"io"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"
)

const (
	stSvr    = "localhost:3333"
	stWebSvr = "localhost:9099"
	stDir    = "/tmp/crdbtest"
	stSchema = "dal/schema.sql"
)

// TestMain is used to setup/teardown a temporary CockroachDB instance
// and run all the unit tests in between.
func TestMain(m *testing.M) {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	if err := setup(); err != nil {
		fmt.Println("cannot setup DB:", err)
		os.Exit(1)
	}

	exitCode := m.Run() // run all unittests

	teardown()
	os.Exit(exitCode)
}

func setup() error {
	// Start the DB.
	err := os.RemoveAll(stDir)
	if err != nil {
		return fmt.Errorf("cannot remove old DB directory: %s: %v", stDir, err)
	}

	schema, err := os.Open(stSchema)
	if err != nil {
		return fmt.Errorf("cannot open schema file: %s: %v", stSchema, err)
	}
	defer schema.Close()

	cmd := exec.Command("cockroach", "start-single-node", "--insecure",
		"--listen-addr="+stSvr, "--http-addr="+stWebSvr,
		"--store=path="+stDir)
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("cannot start DB: %s", err)
	}

	time.Sleep(time.Second)

	// Setup the DB schema.
	cmd = exec.Command("cockroach", "sql", "--insecure", "--host="+stSvr)
	pipe, err := cmd.StdinPipe()
	if err != nil {
		teardown()
		return fmt.Errorf("cannot get stdin of DB command: %v", err)
	}

	go func() {
		defer pipe.Close()
		io.Copy(pipe, schema)
	}()

	if err = cmd.Run(); err != nil {
		teardown()
		return fmt.Errorf("cannot setup DB schema: %v", err)
	}

	return nil
}

func teardown() {
	cmd := exec.Command("cockroach", "quit", "--insecure", "--host="+stSvr)
	if err := cmd.Run(); err != nil {
		fmt.Printf("WARNING: cannot terminate DB: %v", err)
	}

	time.Sleep(time.Second)
	os.RemoveAll(stDir)
}

func errIsNil(t *testing.T, err error) {
	if err != nil {
		t.Errorf("invalid error, it must be nil: %v", err)
	}
}

func checkStatus(t *testing.T, status int, dest int) {
	if status != dest {
		t.Errorf("invalid status, current is:%d,  expect: %d", status, dest)
	}
}

func newTestSvc(t *testing.T) *GatewayService {
	gs, err := NewGatewayService(stSvr)
	sgnRootDir := os.ExpandEnv("$HOME/.sgnd")
	err = gs.initTransactor(sgnRootDir)
	if err != nil {
		t.Errorf("fail to init transactor in gateway server, err:%v", err)
		return nil
	}
	gs.StartChainTokenPolling(10 * time.Second)
	gs.f = fee.NewTokenPriceCache(gs.tr)
	return gs
}

func TestToken(t *testing.T) {
	svc := newTestSvc(t)
	if svc == nil {
		t.Errorf("fail to init service")
		return
	}
	configs, err := svc.GetTransferConfigs(nil, nil)
	errIsNil(t, err)
	t.Log(json.Marshal(configs))
	//transferId := "123"
	//err = dal.DB.InsertTransfer(transferId, "0x0000000", "USDT", 1, 2)
	//errIsNil(t, err)
	//addr, token, srcChainId, dstChainId, status, found, err := dal.DB.GetTransfer(transferId)
	//errIsNil(t, err)
	//if !found {
	//	t.Error("transfer not found")
	//}
	//log.Infof("transfer info: addr:%s, token:%s, src_chain_id:%d, dst_chain_id:%d, status:%d", addr, token, srcChainId, dstChainId, status)
}

func TestTokenAndFee(t *testing.T) {
	svc := newTestSvc(t)
	if svc == nil {
		t.Errorf("fail to init service")
		return
	}
	token := &types.Token{
		Symbol:  "USDT",
		Address: "3efc487eef37187483d8f7dbe5f8781f2af4b5c5",
		Decimal: 6,
	}
	tokenUsdPrice := svc.f.GetUsdVolume(token, big.NewInt(2500))
	t.Log("DAI eth prize:", tokenUsdPrice)
	configs, err := svc.GetTransferConfigs(nil, nil)
	errIsNil(t, err)
	t.Logf("configs:%s", configs)
}
