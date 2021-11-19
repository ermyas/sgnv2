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
	"strconv"
	"testing"
	"time"

	"github.com/spf13/viper"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/fee"
	gatewaysvc "github.com/celer-network/sgn-v2/gateway/svc"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
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
	gatewaysvc.RootDir = os.ExpandEnv("$HOME/.sgnd")
	gatewaysvc.LegacyAmino = codec.NewLegacyAmino()
	gatewaysvc.InterfaceRegistry = codectypes.NewInterfaceRegistry()
	gatewaysvc.Cdc = codec.NewProtoCodec(gatewaysvc.InterfaceRegistry)
	gatewaysvc.SelfStart = true
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
	gs, err := gatewaysvc.NewGatewayService(stSvr)
	require.NoError(t, err, "failed to initialize gateway service", err)
	err = gs.InitTransactors()
	require.NoError(t, err, "failed to initialize gateway transactors", err)
	gs.StartChainTokenPolling(1 * time.Hour)
	gs.StartUpdateTokenPricePolling(time.Duration(viper.GetInt32(common.FlagSgnCheckIntervalCbrPrice)) * time.Second)
	gs.F = fee.NewTokenPriceCache(gs.TP.GetTransactor())
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
	_db, err := dal.NewDAL("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", stSvr), 10)
	errIsNil(t, err)

	dal.DB = _db
	usrAddr := "0x25846D545a60A029E5C83f0FB96e41b408528e9E"

	err = dal.DB.InsertClaimWithdrawRewardLog(usrAddr)

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
func TestTransfer(t *testing.T) {
	svc := newTestSvc(t)
	if svc == nil {
		t.Errorf("fail to init service")
		return
	}
	mockChian()

	configs, err := svc.GetTransferConfigs(nil, nil)
	errIsNil(t, err)
	errMsgIsNil(t, configs.Err)
	chainTokens := configs.GetChainToken()
	chains := configs.GetChains()
	chain1 := chains[0].GetId()
	chain2 := chains[1].GetId()
	chainToken1 := chainTokens[chain1]
	chainToken2 := chainTokens[chain2]

	srcAmt := "10000"
	usrAddr := "0x25846D545a60A029E5C83f0FB96e41b408528e9E"
	srcTxHash := "111111111"
	transferId := "1"

	estimateAmt, err := svc.EstimateAmt(nil, &webapi.EstimateAmtRequest{
		SrcChainId:  chain1,
		DstChainId:  chain2,
		TokenSymbol: chainToken1.Token[0].Token.Symbol,
		Amt:         srcAmt,
		UsrAddr:     usrAddr,
	})
	errIsNil(t, err)
	errMsgIsNil(t, estimateAmt.Err)
	t.Log("estimate amt:", estimateAmt)
	dstAmt, _ := strconv.Atoi(estimateAmt.EqValueTokenAmt)
	fee, _ := strconv.Atoi(estimateAmt.GetPercFee())
	dstAmt = int(float64(dstAmt)*(1-float64(estimateAmt.SlippageTolerance)/10000.0)) - fee
	t.Log("min received amt:", dstAmt)

	markTransferResponse, err := svc.MarkTransfer(nil, &webapi.MarkTransferRequest{
		TransferId: transferId,
		SrcSendInfo: &webapi.TransferInfo{
			Chain:  chains[0],
			Token:  chainToken1.GetToken()[0].Token,
			Amount: srcAmt,
		},
		DstMinReceivedInfo: &webapi.TransferInfo{
			Chain:  chains[1],
			Token:  chainToken2.GetToken()[0].Token,
			Amount: fmt.Sprint(dstAmt),
		},
		Addr:      usrAddr,
		SrcTxHash: srcTxHash,
		Type:      webapi.TransferType_TRANSFER_TYPE_SEND,
	})
	errIsNil(t, err)
	errMsgIsNil(t, markTransferResponse.Err)

	history, err := svc.TransferHistory(nil, &webapi.TransferHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          usrAddr,
	})
	errIsNil(t, err)
	errMsgIsNil(t, history.Err)
	checkTransferStatus(t, history.History[0].GetStatus(), types.TransferHistoryStatus_TRANSFER_SUBMITTING)

	err = relayer.GatewayOnSend(transferId, usrAddr, chainToken1.Token[0].Token.Symbol, srcAmt, srcTxHash, uint64(chain1), uint64(chain2))
	errIsNil(t, err)
	history, err = svc.TransferHistory(nil, &webapi.TransferHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          usrAddr,
	})
	errIsNil(t, err)
	checkTransferStatus(t, history.History[0].GetStatus(), types.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE)
	err = relayer.GatewayOnRelay(transferId, srcTxHash, "2", string(rune(dstAmt)))
	errIsNil(t, err)
	history, err = svc.TransferHistory(nil, &webapi.TransferHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          usrAddr,
	})
	errIsNil(t, err)
	checkTransferStatus(t, history.History[0].GetStatus(), types.TransferHistoryStatus_TRANSFER_COMPLETED)
}
func TestTransferRefund(t *testing.T) {
	svc := newTestSvc(t)
	if svc == nil {
		t.Errorf("fail to init service")
		return
	}
	mockChian()

	configs, err := svc.GetTransferConfigs(nil, nil)
	errIsNil(t, err)
	errMsgIsNil(t, configs.Err)
	chainTokens := configs.GetChainToken()
	chains := configs.GetChains()
	chain1 := chains[0].GetId()
	chain2 := chains[1].GetId()
	chainToken1 := chainTokens[chain1]
	chainToken2 := chainTokens[chain2]

	srcAmt := "10000"
	usrAddr := "0x25846D545a60A029E5C83f0FB96e41b408528e9E"
	srcTxHash := "111111111"
	transferId := "1"

	seqNum := uint64(1)

	markTransferResponse, err := svc.MarkTransfer(nil, &webapi.MarkTransferRequest{
		TransferId: transferId,
		SrcSendInfo: &webapi.TransferInfo{
			Chain:  chains[0],
			Token:  chainToken1.GetToken()[0].Token,
			Amount: srcAmt,
		},
		DstMinReceivedInfo: &webapi.TransferInfo{
			Chain:  chains[1],
			Token:  chainToken2.GetToken()[0].Token,
			Amount: srcAmt,
		},
		Addr:      usrAddr,
		SrcTxHash: srcTxHash,
		Type:      webapi.TransferType_TRANSFER_TYPE_SEND,
	})
	errIsNil(t, err)
	errMsgIsNil(t, markTransferResponse.Err)

	//withdraw directly first
	err = dal.DB.MarkTransferRequestingRefund(transferId, seqNum)
	errIsNil(t, err)
	statusRsp, err := svc.GetTransferStatus(nil, &webapi.GetTransferStatusRequest{TransferId: transferId})
	errIsNil(t, err)
	errMsgIsNil(t, statusRsp.Err)
	checkTransferStatus(t, statusRsp.GetStatus(), types.TransferHistoryStatus_TRANSFER_REQUESTING_REFUND)

	/**
	...
	update history: get types.TransferHistoryStatus_TRANSFER_REFUND_TO_BE_CONFIRMED from sgn
	...
	*/

	markTransferResponse, err = svc.MarkTransfer(nil, &webapi.MarkTransferRequest{
		TransferId: transferId,
		SrcSendInfo: &webapi.TransferInfo{
			Chain:  chains[0],
			Token:  chainToken1.GetToken()[0].Token,
			Amount: srcAmt,
		},
		DstMinReceivedInfo: &webapi.TransferInfo{
			Chain:  chains[1],
			Token:  chainToken2.GetToken()[0].Token,
			Amount: srcAmt,
		},
		Addr:      usrAddr,
		SrcTxHash: srcTxHash,
		Type:      webapi.TransferType_TRANSFER_TYPE_REFUND,
	})
	errIsNil(t, err)
	errMsgIsNil(t, markTransferResponse.Err)

	statusRsp, err = svc.GetTransferStatus(nil, &webapi.GetTransferStatusRequest{TransferId: transferId})
	errIsNil(t, err)
	errMsgIsNil(t, statusRsp.Err)
	checkTransferStatus(t, statusRsp.GetStatus(), types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND)

	relayer.GatewayOnLiqWithdraw(uint64(chain1), seqNum, usrAddr)
	statusRsp, err = svc.GetTransferStatus(nil, &webapi.GetTransferStatusRequest{TransferId: transferId})
	errIsNil(t, err)
	errMsgIsNil(t, statusRsp.Err)
	checkTransferStatus(t, statusRsp.GetStatus(), types.TransferHistoryStatus_TRANSFER_REFUNDED)
}

func TestLPAdd(t *testing.T) {
	svc := newTestSvc(t)
	if svc == nil {
		t.Errorf("fail to init service")
		return
	}
	mockChian()
	configs, err := svc.GetTransferConfigs(nil, nil)
	errIsNil(t, err)
	errMsgIsNil(t, configs.Err)
	token := configs.GetChainToken()[883].Token[0]
	// add
	addr := "0x25846D545a60A029E5C83f0FB96e41b408528e9E"
	amt := "1000"
	tokenAddr := common.Hex2Addr(token.Token.Address).String()
	chainId := 883
	txHash := "111"
	seqNum := uint64(1)
	markLiquidityResponse, err := svc.MarkLiquidity(nil, &webapi.MarkLiquidityRequest{
		LpAddr:    addr,
		Amt:       amt,
		TokenAddr: tokenAddr,
		ChainId:   uint32(chainId),
		SeqNum:    seqNum,
		TxHash:    txHash,
		Type:      webapi.LPType_LP_TYPE_ADD,
	})
	errIsNil(t, err)
	errMsgIsNil(t, markLiquidityResponse.Err)
	lpHistory, err := svc.LPHistory(nil, &webapi.LPHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          addr,
	})
	errIsNil(t, err)
	errMsgIsNil(t, lpHistory.Err)
	checkLpStatus(t, lpHistory.History[0].Status, types.WithdrawStatus_WD_SUBMITTING)

	// onchain status
	relayer.GatewayOnLiqAdd(addr, token.Token.Symbol, tokenAddr, amt, txHash, uint64(chainId), seqNum)
	liquidityStatus, err := svc.QueryLiquidityStatus(nil, &webapi.QueryLiquidityStatusRequest{
		SeqNum:  seqNum,
		LpAddr:  addr,
		ChainId: uint32(chainId),
		Type:    webapi.LPType_LP_TYPE_ADD,
	}) //polling

	errIsNil(t, err)
	checkLpStatus(t, liquidityStatus.Status, types.WithdrawStatus_WD_WAITING_FOR_SGN)

	// skip complete check without on chain sgn running
}

func TestLPWithdraw(t *testing.T) {
	svc := newTestSvc(t)
	if svc == nil {
		t.Errorf("fail to init service")
		return
	}
	mockChian()
	configs, err := svc.GetTransferConfigs(nil, nil)
	errIsNil(t, err)
	errMsgIsNil(t, configs.Err)
	token := configs.GetChainToken()[883].Token[0]
	// add
	addr := "0x25846D545a60A029E5C83f0FB96e41b408528e9E"
	amt := "1000"
	ratio := 2000 // 20%
	tokenAddr := common.Hex2Addr(token.Token.Address).String()
	chainId := 883
	txHash := "111"

	withdrawLq := &types.WithdrawLq{
		FromChainId: uint64(chainId),
		TokenAddr:   tokenAddr,
		Ratio:       uint32(ratio),
	}
	withdrawReq := &types.WithdrawReq{
		Withdraws:   []*types.WithdrawLq{withdrawLq},
		ExitChainId: uint64(chainId),
		ReqId:       1,
	}
	wdBytes, _ := withdrawReq.Marshal()

	withdrawLiquidityResponse, err := svc.WithdrawLiquidity(nil, &webapi.WithdrawLiquidityRequest{
		WithdrawReq: wdBytes,
		// TODO: add sig
	})
	errIsNil(t, err)
	errMsgIsNil(t, withdrawLiquidityResponse.Err)
	seqNum := uint64(1)
	lpHistory, err := svc.LPHistory(nil, &webapi.LPHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          addr,
	})
	errIsNil(t, err)
	errMsgIsNil(t, lpHistory.Err)
	checkLpStatus(t, lpHistory.History[0].Status, types.WithdrawStatus_WD_WAITING_FOR_SGN)

	////polling can not used for testing
	//var status types.WithdrawStatus
	//for i := 1; i < 10 && status != types.WithdrawStatus_WD_WAITING_FOR_LP; i++ {
	//	t.Log("polling ", i)
	//	liquidityStatus, err := svc.QueryLiquidityStatus(nil, &webapi.QueryLiquidityStatusRequest{SeqNum: seqNum}) //polling
	//	errIsNil(t, err)
	//	status = liquidityStatus.Status
	//	time.Sleep(1 * time.Second)
	//}
	//checkLpStatus(t, status, types.WithdrawStatus_WD_WAITING_FOR_LP)

	markLiquidityResponse, err := svc.MarkLiquidity(nil, &webapi.MarkLiquidityRequest{
		LpAddr:    addr,
		Amt:       amt,
		TokenAddr: tokenAddr,
		ChainId:   uint32(chainId),
		SeqNum:    seqNum,
		TxHash:    txHash,
		Type:      webapi.LPType_LP_TYPE_REMOVE,
	})
	errIsNil(t, err)
	errMsgIsNil(t, markLiquidityResponse.Err)
	lpHistory, err = svc.LPHistory(nil, &webapi.LPHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          addr,
	})
	errIsNil(t, err)
	errMsgIsNil(t, lpHistory.Err)
	checkLpStatus(t, lpHistory.History[0].Status, types.WithdrawStatus_WD_SUBMITTING)

	// onchain status
	relayer.GatewayOnLiqWithdraw(uint64(chainId), seqNum, addr)
	lpHistory, err = svc.LPHistory(nil, &webapi.LPHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          addr,
	})
	errIsNil(t, err)
	errMsgIsNil(t, lpHistory.Err)
	checkLpStatus(t, lpHistory.History[0].Status, types.WithdrawStatus_WD_COMPLETED)
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
	utils.SendBalanceAlert(alerts, "local")
}
