package gateway

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/relayer"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"io"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

const (
	stSvr    = "localhost:3333"
	stWebSvr = "localhost:9099"
	stDir    = "/tmp/crdbtest"
	stSchema = "dal/schema.sql"
)

func setGlobal() {
	rootDir = os.ExpandEnv("$HOME/.sgnd")
	legacyAmino = codec.NewLegacyAmino()
	interfaceRegistry = codectypes.NewInterfaceRegistry()
	cdc = codec.NewProtoCodec(interfaceRegistry)
	selfStart = true
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
	if err != nil {
		t.Errorf("invalid error, it must be nil: %v", err)
	}
}

func errMsgIsNil(t *testing.T, err *webapi.ErrMsg) {
	if err != nil {
		t.Errorf("invalid error in response, it must be nil: %v", err)
	}
}

func checkTransferStatus(t *testing.T, status types.TransferHistoryStatus, dest types.TransferHistoryStatus) {
	if status != dest {
		t.Errorf("invalid status, current is:%d,  expect: %d", status, dest)
	}
}

func checkLpStatus(t *testing.T, status types.LPHistoryStatus, dest types.LPHistoryStatus) {
	if status != dest {
		t.Errorf("invalid status, current is:%d,  expect: %d", status, dest)
	}
}

func newTestSvc(t *testing.T) *GatewayService {
	gs, err := NewGatewayService(stSvr)
	err = gs.initTransactor()
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

func mockChian() {
	dal.DB.UpsertChainInfo(883, "chain1", "test1", "url1")
	dal.DB.UpsertChainInfo(884, "chain2", "test2", "url2")
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

	tlrsResp, err := svc.SetAdvancedInfo(nil, &webapi.SetAdvancedInfoRequest{
		Addr:              "0x25846D545a60A029E5C83f0FB96e41b408528e9E",
		SlippageTolerance: 200,
	})
	errIsNil(t, err)
	errMsgIsNil(t, tlrsResp.Err)
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
	fee, _ := strconv.Atoi(estimateAmt.GetFee())
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

	err = relayer.GatewayOnSend(transferId)
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

	tlrsResp, err := svc.SetAdvancedInfo(nil, &webapi.SetAdvancedInfoRequest{
		Addr:              "0x25846D545a60A029E5C83f0FB96e41b408528e9E",
		SlippageTolerance: 200,
	})
	errIsNil(t, err)
	errMsgIsNil(t, tlrsResp.Err)

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
		Addr:           usrAddr,
		SrcTxHash:      srcTxHash,
		Type:           webapi.TransferType_TRANSFER_TYPE_REFUND,
		WithdrawSeqNum: seqNum,
	})
	errIsNil(t, err)
	errMsgIsNil(t, markTransferResponse.Err)

	statusRsp, err = svc.GetTransferStatus(nil, &webapi.GetTransferStatusRequest{TransferId: transferId})
	errIsNil(t, err)
	errMsgIsNil(t, statusRsp.Err)
	checkTransferStatus(t, statusRsp.GetStatus(), types.TransferHistoryStatus_TRANSFER_CONFIRMING_YOUR_REFUND)

	relayer.GatewayOnLiqWithdraw(seqNum)
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
	checkLpStatus(t, lpHistory.History[0].Status, types.LPHistoryStatus_LP_SUBMITTING)

	// onchain status
	relayer.GatewayOnLiqAdd(addr, token.Token.Symbol, tokenAddr, amt, txHash, uint64(chainId), seqNum)
	liquidityStatus, err := svc.QueryLiquidityStatus(nil, &webapi.QueryLiquidityStatusRequest{
		SeqNum:  seqNum,
		LpAddr:  addr,
		ChainId: uint32(chainId),
		Type:    webapi.LPType_LP_TYPE_ADD,
	}) //polling

	errIsNil(t, err)
	checkLpStatus(t, liquidityStatus.Status, types.LPHistoryStatus_LP_WAITING_FOR_SGN)

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
	tokenAddr := common.Hex2Addr(token.Token.Address).String()
	chainId := 883
	txHash := "111"

	withdrawLiquidityResponse, err := svc.WithdrawLiquidity(nil, &webapi.WithdrawLiquidityRequest{
		TransferId:   "",
		ReceiverAddr: addr,
		Amount:       amt,
		TokenAddr:    tokenAddr,
		ChainId:      uint32(chainId),
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
	checkLpStatus(t, lpHistory.History[0].Status, types.LPHistoryStatus_LP_WAITING_FOR_SGN)

	////polling can not used for testing
	//var status types.LPHistoryStatus
	//for i := 1; i < 10 && status != types.LPHistoryStatus_LP_WAITING_FOR_LP; i++ {
	//	t.Log("polling ", i)
	//	liquidityStatus, err := svc.QueryLiquidityStatus(nil, &webapi.QueryLiquidityStatusRequest{SeqNum: seqNum}) //polling
	//	errIsNil(t, err)
	//	status = liquidityStatus.Status
	//	time.Sleep(1 * time.Second)
	//}
	//checkLpStatus(t, status, types.LPHistoryStatus_LP_WAITING_FOR_LP)

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
	checkLpStatus(t, lpHistory.History[0].Status, types.LPHistoryStatus_LP_SUBMITTING)

	// onchain status
	relayer.GatewayOnLiqWithdraw(seqNum)
	lpHistory, err = svc.LPHistory(nil, &webapi.LPHistoryRequest{
		NextPageToken: "",
		PageSize:      10,
		Addr:          addr,
	})
	errIsNil(t, err)
	errMsgIsNil(t, lpHistory.Err)
	checkLpStatus(t, lpHistory.History[0].Status, types.LPHistoryStatus_LP_COMPLETED)
}

func TestLPList(t *testing.T) {
	svc := newTestSvc(t)
	if svc == nil {
		t.Errorf("fail to init service")
		return
	}
	mockChian()
	addr := "0x25846D545a60A029E5C83f0FB96e41b408528e9E"
	list, err := svc.GetLPInfoList(nil, &webapi.GetLPInfoListRequest{Addr: addr})
	errIsNil(t, err)
	errMsgIsNil(t, list.Err)
	t.Log(list)
}
