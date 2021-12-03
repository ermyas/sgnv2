package gateway

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/viper"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/onchain"
	gatewaysvc "github.com/celer-network/sgn-v2/gateway/svc"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	stSvr    = "localhost:3333"
	stWebSvr = "localhost:9099"
	stDir    = "/tmp/crdbtest"
	stSchema = "dal/schema.sql"
)

func setGlobal() {
	*home = os.ExpandEnv("$HOME/.sgnd")
}

// TestMain is used to setup/teardown a temporary CockroachDB instance
// and run all the unit tests in between.
func TestMain(m *testing.M) {
	flag.Parse()
	rand.Seed(time.Now().Unix())
	setGlobal()

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
	require.NoError(t, err, "expect no error")
}

func errMsgIsNil(t *testing.T, errMsg *webapi.ErrMsg) {
	assert.Nil(t, errMsg, "expect nil ErrMsg")
}

func checkTransferStatus(t *testing.T, status types.TransferHistoryStatus, dest types.TransferHistoryStatus) {
	if status != dest {
		t.Errorf("invalid status, current is:%d,  expect: %d", status, dest)
	}
}

func checkLpStatus(t *testing.T, status types.WithdrawStatus, dest types.WithdrawStatus) {
	if status != dest {
		t.Errorf("invalid status, current is:%d,  expect: %d", status, dest)
	}
}

func newTestSvc(t *testing.T) *gatewaysvc.GatewayService {
	db := dal.NewDAL(viper.GetString(common.FlagGatewayDbUrl))
	gs := gatewaysvc.NewGatewayService(db)
	encoding := app.MakeEncodingConfig()
	onchain.InitSGNTransactors(*home, encoding)
	gs.StartChainTokenPolling(1 * time.Hour)
	gs.StartUpdateTokenPricePolling(time.Duration(viper.GetInt32(common.FlagSgnCheckIntervalCbrPrice)) * time.Second)
	gs.F = gatewaysvc.NewTokenPriceCache(onchain.SGNTransactors.GetTransactor())
	signerKey, signerPass := viper.GetString(common.FlagGatewayIncentiveRewardsKeystore), viper.GetString(common.FlagGatewayIncentiveRewardsPassphrase)
	signer, addr, err := eth.CreateSigner(signerKey, signerPass, nil)
	if err != nil {
		require.NoError(t, err, "fail to CreateSigner in gateway server, err:", err)
	}
	gs.S = &gatewaysvc.IncentiveRewardsSigner{
		Signer: &signer,
		Addr:   &addr,
	}
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
	tokenUsdPrice := svc.F.GetUsdVolume(token, big.NewInt(2500))
	t.Log("DAI eth prize:", tokenUsdPrice)
	configs, err := svc.GetTransferConfigs(nil, nil)
	errIsNil(t, err)
	t.Logf("configs:%s", configs)
}

func mockChian() {
	dal.DB.UpsertChainUIInfo(883, "chain1", "test1", "url1", "xxx", "url1", "url1")
	dal.DB.UpsertChainUIInfo(884, "chain2", "test2", "url2", "yyy", "url2", "url2")
}
func TestCampaign(t *testing.T) {
	_db := dal.NewDAL(stSvr)

	dal.DB = _db
	usrAddr := "0x25846D545a60A029E5C83f0FB96e41b408528e9E"

	err := dal.DB.InsertClaimWithdrawRewardLog(usrAddr)

	score, err := dal.DB.CalcCampaignScore(time.Now())
	errIsNil(t, err)
	if score[0].UsrAddr != usrAddr {
		t.Errorf("usrAddr wrong")
	}
	if score[0].Score != 1 {
		t.Errorf("Score wrong")
	}

	time.Sleep(2 * time.Second)

	err = dal.DB.InsertClaimWithdrawRewardLog(usrAddr)

	score, err = dal.DB.CalcCampaignScore(time.Now())
	errIsNil(t, err)
	if score[0].UsrAddr != usrAddr {
		t.Errorf("usrAddr wrong")
	}
	if score[0].Score != 2 {
		t.Errorf("Score wrong")
	}

	err = dal.DB.InsertClaimWithdrawRewardLog(usrAddr)
	err = dal.DB.InsertClaimWithdrawRewardLog(usrAddr)
	err = dal.DB.InsertClaimWithdrawRewardLog(usrAddr)
	err = dal.DB.InsertClaimWithdrawRewardLog(usrAddr)
	q := `INSERT INTO claim_withdraw_reward_log (usr_addr, create_time)
                VALUES ($1, $2)`
	dal.DB.Exec(q, usrAddr, time.Now().Add(24*time.Hour))
	q = `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	dal.DB.Exec(q, usrAddr, 33, "ddd", "ggg", "ggg", "ggg", time.Now(), time.Now(), 4, 1, 1)
	q = `INSERT INTO lp (usr_addr, chain_id, token_symbol, token_addr, amt, tx_hash, update_time, create_time, status, lp_type, seq_num)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	dal.DB.Exec(q, usrAddr, 33, "ddd", "ggg", "ggg", "ggg", time.Now(), time.Now(), 3, 1, 2)

	score, err = dal.DB.CalcCampaignScore(time.Now())
	errIsNil(t, err)
	if score[0].UsrAddr != usrAddr {
		t.Errorf("usrAddr wrong")
	}
	if score[0].Score != 7 {
		t.Errorf("Score wrong")
	}

}
func TestAlert(t *testing.T) {
	//utils.SendWithdrawAlert("0x2147F049De1D68bC8265B260760AbA6eda614367", "900", "800", "100")
	var alerts []*utils.BalanceAlert
	alerts = append(alerts, &utils.BalanceAlert{
		Token:    "test1",
		Balance:  "300",
		Addr:     "0x2147F049De1D68bC8265B260760AbA6eda614367",
		Withdraw: "1000",
		Deposit:  "800",
	})
	alerts = append(alerts, &utils.BalanceAlert{
		Token:    "test2",
		Balance:  "300",
		Addr:     "0x3147F049De1D68bC8265B260760AbA6eda614367",
		Withdraw: "1200",
		Deposit:  "700",
	})
	utils.SendBalanceAlert(alerts)
}

func TestRetentionRewards(t *testing.T) {
	_db := dal.NewDAL(stSvr)
	defer _db.Close()

	dal.DB = _db
	q := `insert into retention_rewards_log
          (usr_addr, event_id, group_level)
          values($1, $2, $3)`
	addr := "0x1111111111111111111111111111111111111111"
	dal.DB.Exec(q, addr, 1, 1)
	record, amt, claimTime, signature, found, err := dal.DB.GetRetentionRewardsRecord(addr, 1)
	if err != nil {
		t.Errorf("fail to GetRetentionRewardsRecord %v", err)
	}
	t.Log("", record, amt, claimTime, signature, found)
	return
}
