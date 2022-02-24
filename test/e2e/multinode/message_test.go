package multinode

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgcli "github.com/celer-network/sgn-v2/x/message/client/cli"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gogo/protobuf/proto"
)

func setupMessage() {
	log.Infoln("Set up new sgn env")
	p := &tc.ContractParams{
		CelrAddr:              tc.CelrAddr,
		ProposalDeposit:       big.NewInt(1e18),
		VotePeriod:            big.NewInt(5),
		UnbondingPeriod:       big.NewInt(5),
		MaxBondedValidators:   big.NewInt(3),
		MinValidatorTokens:    big.NewInt(1e18),
		MinSelfDelegation:     big.NewInt(1e18),
		AdvanceNoticePeriod:   big.NewInt(1),
		ValidatorBondInterval: big.NewInt(0),
		MaxSlashFactor:        big.NewInt(1e5),
	}
	SetupNewSgnEnv(p, true, true, false, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestMessage(t *testing.T) {
	t.Run("e2e-message", func(t *testing.T) {
		t.Run("messageTest", messageTest)
	})
}

func messageTest(t *testing.T) {
	log.Infoln("******************************************************************")
	log.Infoln("======================== Test message ============================")
	setupMessage()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	prepareValidators(t)
	prepareCbrLiquidity(transactor)

	log.Infoln("======================= Message Only Test =====================")
	messageOnlyTest(transactor, msgtypes.EXECUTION_STATUS_SUCCESS)

	log.Infoln("======================= Batch Transfer Test =====================")
	batchTransferTest(transactor, big.NewInt(100*1000000), big.NewInt(5*1000000), msgtypes.EXECUTION_STATUS_SUCCESS)

	log.Infoln("======================= Batch Pegged Transfer Test =====================")
	batchPegTransferTest(transactor, new(big.Int).Mul(big.NewInt(50), big.NewInt(1e18)), new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)), msgtypes.EXECUTION_STATUS_SUCCESS)

	log.Infoln("======================= Refund Tests =====================")
	refundTransferTest(t, transactor)
	refundPegDepositTest(t, transactor)
	// refundPegBurnTest(t, transactor)

	log.Infoln("======================= Batch Transfer FallBack Test =====================")
	batchTransferTest(transactor, big.NewInt(10*1000000), big.NewInt(5*1000000), msgtypes.EXECUTION_STATUS_FALLBACK)
}

func batchPegTransferTest(transactor *transactor.Transactor, sendAmt *big.Int, amtForEveryone *big.Int, expectedStatus msgtypes.ExecutionStatus) {
	log.Infoln("----------------------- Batch Pegged Deposit Transfer Test -----------------------")
	batchPegDepositTest(transactor, sendAmt, amtForEveryone, expectedStatus)

	log.Infoln("----------------------- Batch Pegged burn Transfer Test -----------------------")
	// burn test using pegged token from deposit test
	burnAmt := new(big.Int).Div(amtForEveryone, new(big.Int).SetUint64(2))
	withdrawAmtForEveryone := new(big.Int).Div(burnAmt, new(big.Int).SetUint64(3))
	batchPegBurnTest(transactor, burnAmt, withdrawAmtForEveryone, expectedStatus)
}

func batchPegDepositTest(transactor *transactor.Transactor, sendAmt *big.Int, amtForEveryone *big.Int, expectedStatus msgtypes.ExecutionStatus) {
	err := tc.CbrChain1.ApproveUNIForBatchTransfer(0, sendAmt)
	tc.ChkErr(err, "ApproveUNIForBatchTransfer")

	u0 := tc.CbrChain1.Users[0]
	u1 := tc.CbrChain1.Users[1]
	u2 := tc.CbrChain1.Users[2]
	originBalanceOfU01, _ := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, u0.Address)
	originBalanceOfU02, _ := tc.CbrChain2.UNIContract.BalanceOf(&bind.CallOpts{}, u0.Address)
	originBalanceOfU1, _ := tc.CbrChain2.UNIContract.BalanceOf(&bind.CallOpts{}, u1.Address)
	originBalanceOfU2, _ := tc.CbrChain2.UNIContract.BalanceOf(&bind.CallOpts{}, u2.Address)

	// chain1 -> chain2 deposit -> mint
	depositId, err := tc.CbrChain1.BatchTransfer(
		0,
		tc.CbrChain2.BatchTransferAddr,
		tc.CbrChain1.UNIAddr,
		sendAmt,
		tc.CbrChain2.ChainId,
		100000,
		uint8(2),
		[]eth.Addr{u1.Address, u2.Address},
		[]*big.Int{amtForEveryone, amtForEveryone},
	)
	tc.ChkErr(err, "u0 chain1 batch vault transfer")
	depositInfo := tc.WaitPbrDeposit(transactor, depositId.String())
	if len(depositInfo.MintId) == 0 {
		log.Fatalln("refunded deposit", nil)
	}
	mintInfo := tc.CheckPbrMint(transactor, eth.Bytes2Hex(depositInfo.GetMintId()))
	tc.WaitForMessageExecuted(transactor, expectedStatus)

	var mintOnChain pegbrtypes.MintOnChain
	err = proto.Unmarshal(mintInfo.MintProtoBytes, &mintOnChain)
	tc.ChkErr(err, "unmarshal MintOnChain")
	mintAmt := new(big.Int).SetBytes(mintOnChain.Amount)
	delta := new(big.Int).Sub(mintAmt, new(big.Int).Add(amtForEveryone, amtForEveryone))
	tc.CbrChain1.CheckUNIBalance(uint64(0), new(big.Int).Sub(originBalanceOfU01, sendAmt))
	tc.CbrChain2.CheckPeggedUNIBalance(uint64(0), new(big.Int).Add(originBalanceOfU02, delta))
	tc.CbrChain2.CheckPeggedUNIBalance(uint64(1), new(big.Int).Add(originBalanceOfU1, amtForEveryone))
	tc.CbrChain2.CheckPeggedUNIBalance(uint64(2), new(big.Int).Add(originBalanceOfU2, amtForEveryone))

}

func batchPegBurnTest(transactor *transactor.Transactor, sendAmt *big.Int, amtForEveryone *big.Int, expectedStatus msgtypes.ExecutionStatus) {
	err := tc.CbrChain2.ApprovePeggedUNIForBatchTransfer(0, sendAmt)
	tc.ChkErr(err, "ApprovePeggedUNIForBatchTransfer")

	u0 := tc.CbrChain2.Users[0]
	u1 := tc.CbrChain2.Users[1]
	u2 := tc.CbrChain2.Users[2]
	originBalanceOfU02, _ := tc.CbrChain2.UNIContract.BalanceOf(&bind.CallOpts{}, u0.Address)
	originBalanceOfU01, _ := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, u0.Address)
	originBalanceOfU1, _ := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, u1.Address)
	originBalanceOfU2, _ := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, u2.Address)

	// chain2 -> chain1 burn -> withdraw
	burnId, err := tc.CbrChain2.BatchTransfer(
		0,
		tc.CbrChain1.BatchTransferAddr,
		tc.CbrChain2.UNIAddr,
		sendAmt,
		tc.CbrChain1.ChainId,
		100000,
		uint8(3),
		[]eth.Addr{u1.Address, u2.Address},
		[]*big.Int{amtForEveryone, amtForEveryone},
	)
	tc.ChkErr(err, "u1 chain2 batch peg transfer")
	burnInfo := tc.WaitPbrBurn(transactor, burnId.String())
	withdrawInfo := tc.CheckPbrWithdraw(transactor, eth.Bytes2Hex(burnInfo.WithdrawId))
	tc.WaitForMessageExecuted(transactor, expectedStatus)

	var withdrawOnChain pegbrtypes.WithdrawOnChain
	err = proto.Unmarshal(withdrawInfo.WithdrawProtoBytes, &withdrawOnChain)
	tc.ChkErr(err, "unmarshal BurnOnChain")
	withdrawAmt := new(big.Int).SetBytes(withdrawOnChain.Amount)
	delta := new(big.Int).Sub(withdrawAmt, new(big.Int).Add(amtForEveryone, amtForEveryone))

	tc.CbrChain2.CheckPeggedUNIBalance(uint64(0), new(big.Int).Sub(originBalanceOfU02, sendAmt))
	tc.CbrChain1.CheckUNIBalance(uint64(0), new(big.Int).Add(originBalanceOfU01, delta))
	tc.CbrChain1.CheckUNIBalance(uint64(1), new(big.Int).Add(originBalanceOfU1, amtForEveryone))
	tc.CbrChain1.CheckUNIBalance(uint64(2), new(big.Int).Add(originBalanceOfU2, amtForEveryone))
}

func batchTransferTest(transactor *transactor.Transactor, sendAmt *big.Int, amtForEveryone *big.Int, expectedStatus msgtypes.ExecutionStatus) {
	log.Infoln("------------------------ batchTransferTest ------------------------")

	u0 := tc.CbrChain1.Users[0]
	u1 := tc.CbrChain1.Users[1]
	u2 := tc.CbrChain1.Users[2]
	err := tc.CbrChain1.ApproveUSDTForBatchTransfer(0, sendAmt)
	tc.ChkErr(err, "u0 chain1 approve")
	originBalanceOfU01, _ := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u0.Address)
	originBalanceOfU02, _ := tc.CbrChain2.USDTContract.BalanceOf(&bind.CallOpts{}, u0.Address)
	originBalanceOfU1, _ := tc.CbrChain2.USDTContract.BalanceOf(&bind.CallOpts{}, u1.Address)
	originBalanceOfU2, _ := tc.CbrChain2.USDTContract.BalanceOf(&bind.CallOpts{}, u2.Address)
	xferId, err := tc.CbrChain1.BatchTransfer(
		0,
		tc.CbrChain2.BatchTransferAddr,
		tc.CbrChain1.USDTAddr,
		sendAmt,
		tc.CbrChain2.ChainId,
		100000,
		uint8(1),
		[]eth.Addr{u1.Address, u2.Address},
		[]*big.Int{amtForEveryone, amtForEveryone},
	)
	tc.ChkErr(err, "u0 chain1 batch transfer")
	tc.CheckXfer(transactor, xferId[:])
	tc.WaitForMessageExecuted(transactor, expectedStatus)

	// check balance
	if expectedStatus == msgtypes.EXECUTION_STATUS_SUCCESS {
		relay, err := cbrcli.QueryRelay(transactor.CliCtx, xferId[:])
		tc.ChkErr(err, "u0 chain1 batch transfer not found after executed")
		relayOnChain := new(cbrtypes.RelayOnChain)
		err = relayOnChain.Unmarshal(relay.Relay)
		tc.ChkErr(err, "unmarshal RelayOnChain")
		relayAmt := new(big.Int).SetBytes(relayOnChain.Amount)
		delta := new(big.Int).Sub(relayAmt, new(big.Int).Add(amtForEveryone, amtForEveryone))
		tc.CbrChain1.CheckUSDTBalance(uint64(0), new(big.Int).Sub(originBalanceOfU01, sendAmt))
		tc.CbrChain2.CheckUSDTBalance(uint64(0), new(big.Int).Add(originBalanceOfU02, delta))
		tc.CbrChain2.CheckUSDTBalance(uint64(1), new(big.Int).Add(originBalanceOfU1, amtForEveryone))
		tc.CbrChain2.CheckUSDTBalance(uint64(2), new(big.Int).Add(originBalanceOfU2, amtForEveryone))
	} else {
		log.Infof("skip balance check for expected status:%s", expectedStatus)
	}
}

func messageOnlyTest(transactor *transactor.Transactor, expectedStatus msgtypes.ExecutionStatus) {
	err := tc.CbrChain1.TransferMsg(
		0,
		tc.CbrChain2.TransferMessageAddr,
		tc.CbrChain2.ChainId,
		new(big.Int).SetInt64(time.Now().UnixNano()).Bytes(),
	)
	tc.ChkErr(err, "message only test")
	tc.WaitForMessageExecuted(transactor, expectedStatus)

	_, err = msgcli.ClaimAllFees(transactor, &msgtypes.MsgClaimAllFees{
		DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[0]),
	})
	tc.ChkErr(err, "failed to claim all fees for delegator 0")

	claimInfo, err := tc.GetFeeClaimInfoWaitForSigs(transactor, tc.DelEthAddrs[0])
	tc.ChkErr(err, "get fee claim info wait for sigs")

	err = tc.WithdrawMsgFeesOnChain(transactor, claimInfo)
	tc.ChkErr(err, "WithdrawMsgFeesOnChain")
}

func refundTransferTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund Test ---------------------")
	u := tc.CbrChain1.Users[0]
	balBefore, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infof("bal before, %s", balBefore.String())
	token := tc.CbrChain1.USDTAddr
	amount := big.NewInt(300000 * 1e6)
	err = tc.CbrChain1.ApproveBridgeTestToken(tc.CbrChain1.USDTContract, 0, amount, tc.CbrChain1.TestRefundAddr)
	tc.ChkErr(err, "approve USDT")
	u.Auth.Value = tc.MsgFeeBase
	tx, err := tc.CbrChain1.TestRefundContract.SendWithTransfer(u.Auth, u.Address, token, amount, tc.CbrChain2.ChainId, 1, 1, uint8(1))
	tc.ChkErr(err, "SendWithTransfer")
	_, err = ethutils.WaitMined(context.Background(), tc.CbrChain1.Ec, tx, ethutils.WithPollingInterval(time.Second))
	tc.ChkErr(err, "SendWithTransfer WaitMined")
	balAfter := new(big.Int)
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		balAfter, err = tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
		tc.ChkErr(err, "bal after")
		log.Infof("%d bal after, %s", i, balAfter.String())
		if balAfter.Cmp(balBefore) == 0 {
			return
		}
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}

func refundPegDepositTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund Peg Deposit Test ---------------------")
	u := tc.CbrChain1.Users[0]
	balBefore, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infof("bal before, %s", balBefore.String())
	token := tc.CbrChain1.USDTAddr
	amount := big.NewInt(10)
	err = tc.CbrChain1.ApproveBridgeTestToken(tc.CbrChain1.USDTContract, 0, amount, tc.CbrChain1.TestRefundAddr)
	tc.ChkErr(err, "approve USDT")
	tx, err := tc.CbrChain1.TestRefundContract.SendWithTransfer(u.Auth, u.Address, token, amount, tc.CbrChain2.ChainId, 1, 1, uint8(2))
	tc.ChkErr(err, "SendWithTransfer")
	_, err = ethutils.WaitMined(context.Background(), tc.CbrChain1.Ec, tx, ethutils.WithPollingInterval(time.Second))
	tc.ChkErr(err, "SendWithTransfer WaitMined")
	balAfter := new(big.Int)
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		balAfter, err = tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
		tc.ChkErr(err, "bal after")
		log.Infof("%d bal after, %s", i, balAfter.String())
		if balAfter.Cmp(balBefore) == 0 {
			return
		}
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}

func refundPegBurnTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund Peg Burn Test ---------------------")
	u := tc.CbrChain1.Users[0]
	balBefore, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infof("bal before, %s", balBefore.String())
	token := tc.CbrChain1.USDTAddr
	amount := big.NewInt(10)
	err = tc.CbrChain1.ApproveBridgeTestToken(tc.CbrChain1.USDTContract, 0, amount, tc.CbrChain1.TestRefundAddr)
	tc.ChkErr(err, "approve USDT")
	tx, err := tc.CbrChain1.TestRefundContract.SendWithTransfer(u.Auth, u.Address, token, amount, tc.CbrChain2.ChainId, 1, 1, uint8(2))
	tc.ChkErr(err, "SendWithTransfer")
	_, err = ethutils.WaitMined(context.Background(), tc.CbrChain1.Ec, tx, ethutils.WithPollingInterval(time.Second))
	tc.ChkErr(err, "SendWithTransfer WaitMined")
	balAfter := new(big.Int)
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Second)
		balAfter, err = tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
		tc.ChkErr(err, "bal after")
		log.Infof("%d bal after, %s", i, balAfter.String())
		if balAfter.Cmp(balBefore) == 0 {
			return
		}
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}

func prepareCbrLiquidity(transactor *transactor.Transactor) {
	log.Infoln("------------------------ Add liquidity on chain 1 ------------------------")
	addAmt := big.NewInt(500000 * 1e6)
	var i uint64
	var err error
	for i = 0; i < 2; i++ {
		err = tc.CbrChain1.ApproveUSDT(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 approve", i))
		err = tc.CbrChain1.AddLiq(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 addliq", i))
		tc.CheckAddLiquidityStatus(transactor, tc.CbrChain1.ChainId, i+1)
		liq, err := tc.QueryTotalLiquidity(transactor, tc.CbrChain1.ChainId, tc.CbrChain1.USDTAddr)
		tc.ChkErr(err, "check liq")
		log.Infoln("chain 1 total liq", liq.String())
	}

	log.Infoln("------------------------ Add liquidity on chain 2 ------------------------")
	for i = 0; i < 2; i++ {
		err = tc.CbrChain2.ApproveUSDT(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain2 approve", i))
		err = tc.CbrChain2.AddLiq(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain2 addliq", i))
		tc.CheckAddLiquidityStatus(transactor, tc.CbrChain2.ChainId, i+1)
		liq, err := tc.QueryTotalLiquidity(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.USDTAddr)
		tc.ChkErr(err, "check liq")
		log.Infoln("chain 2 total liq", liq.String())
	}
}
