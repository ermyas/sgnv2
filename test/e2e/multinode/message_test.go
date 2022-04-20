package multinode

import (
	"math/big"
	"testing"
	"time"

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
	SetupNewSgnEnv(p, &TestFlags{Bridge: true, Msg: true})
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
	prepareValidators(t)

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	tc.RunAllAndWait(
		func() {
			log.Infoln("======================= Message Only Test =====================")
			messageOnlyTest(transactor, msgtypes.EXECUTION_STATUS_SUCCESS)
		},
		func() {
			log.Infoln("======================= Batch Transfer Test =====================")
			prepareCbrLiquidity(transactor)
			batchTransferTest(transactor, big.NewInt(100*1e6), big.NewInt(20*1e6), msgtypes.EXECUTION_STATUS_SUCCESS)
		},
		func() {
			log.Infoln("======================= Batch Pegged Transfer Test =====================")
			batchPegTransferTest(transactor, tc.NewBigInt(5, 19), tc.NewBigInt(1, 19), msgtypes.EXECUTION_STATUS_SUCCESS)
		},
	)

	log.Infoln("======================= PegV2 Tests =====================")
	pegV2DepositTest(t, transactor)
	tc.RunAllAndWait(
		func() {
			pegV2BurnWithdrawTest(t, transactor)
		},
		func() {
			pegV2BurnMintTest(t, transactor)
		},
	)

	log.Infoln("======================= Refund Tests =====================")
	tc.RunAllAndWait(
		func() {
			refundTransferTest(t, transactor)
		},
		func() {
			refundPegDepositTest(t, transactor)
		},
		func() {
			refundPegBurnTest(t, transactor)
		},
		func() {
			refundPegV2DepositTest(t, transactor)
		},
		func() {
			refundPegV2BurnTest(t, transactor)
		},
	)

	log.Infoln("======================= Batch Transfer FallBack Test =====================")
	batchTransferTest(transactor, big.NewInt(10*1e6), big.NewInt(5*1e6), msgtypes.EXECUTION_STATUS_FALLBACK)
}

func messageOnlyTest(transactor *transactor.Transactor, expectedStatus msgtypes.ExecutionStatus) {
	messageId, err := tc.CbrChain1.SendMsg(
		0,
		tc.CbrChain2.MsgTestAddr,
		tc.CbrChain2.ChainId,
		new(big.Int).SetInt64(time.Now().UnixNano()).Bytes(),
	)
	tc.ChkErr(err, "message only test")
	tc.WaitForMessageOnlyExecuted(transactor, messageId, expectedStatus)

	_, err = msgcli.ClaimAllFees(transactor, &msgtypes.MsgClaimAllFees{
		DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[0]),
	})
	tc.ChkErr(err, "failed to claim all fees for delegator 0")

	claimInfo, err := tc.GetFeeClaimInfoWaitForSigs(transactor, tc.DelEthAddrs[0])
	tc.ChkErr(err, "get fee claim info wait for sigs")

	err = tc.WithdrawMsgFeesOnChain(transactor, claimInfo)
	tc.ChkErr(err, "WithdrawMsgFeesOnChain")
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
		tc.BrSendTypeLiquidity,
		[]eth.Addr{u1.Address, u2.Address},
		[]*big.Int{amtForEveryone, amtForEveryone},
	)
	tc.ChkErr(err, "u0 chain1 batch transfer")
	tc.CheckXfer(transactor, xferId[:])
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_LIQUIDITY, xferId, expectedStatus, msgtypes.TRANSFER_TYPE_LIQUIDITY_RELAY)

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
		tc.BrSendTypePegDeposit,
		[]eth.Addr{u1.Address, u2.Address},
		[]*big.Int{amtForEveryone, amtForEveryone},
	)
	tc.ChkErr(err, "u0 chain1 batch vault transfer")
	depositInfo := tc.WaitPbrDeposit(transactor, depositId.String())
	if len(depositInfo.MintId) == 0 {
		log.Fatalln("refunded deposit", nil)
	}
	mintInfo := tc.CheckPbrMint(transactor, eth.Bytes2Hex(depositInfo.GetMintId()))
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_VAULT, depositId, expectedStatus, msgtypes.TRANSFER_TYPE_PEG_MINT)

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
		tc.BrSendTypePegBurn,
		[]eth.Addr{u1.Address, u2.Address},
		[]*big.Int{amtForEveryone, amtForEveryone},
	)
	tc.ChkErr(err, "u1 chain2 batch peg transfer")
	burnInfo := tc.WaitPbrBurn(transactor, burnId.String())
	withdrawInfo := tc.CheckPbrWithdraw(transactor, eth.Bytes2Hex(burnInfo.WithdrawId))
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_BRIDGE, burnId, expectedStatus, msgtypes.TRANSFER_TYPE_PEG_WITHDRAW)

	var withdrawOnChain pegbrtypes.WithdrawOnChain
	err = proto.Unmarshal(withdrawInfo.WithdrawProtoBytes, &withdrawOnChain)
	tc.ChkErr(err, "unmarshal WithdrawOnChain")
	withdrawAmt := new(big.Int).SetBytes(withdrawOnChain.Amount)
	delta := new(big.Int).Sub(withdrawAmt, new(big.Int).Add(amtForEveryone, amtForEveryone))

	tc.CbrChain2.CheckPeggedUNIBalance(uint64(0), new(big.Int).Sub(originBalanceOfU02, sendAmt))
	tc.CbrChain1.CheckUNIBalance(uint64(0), new(big.Int).Add(originBalanceOfU01, delta))
	tc.CbrChain1.CheckUNIBalance(uint64(1), new(big.Int).Add(originBalanceOfU1, amtForEveryone))
	tc.CbrChain1.CheckUNIBalance(uint64(2), new(big.Int).Add(originBalanceOfU2, amtForEveryone))
}

func pegV2DepositTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Peg Deposit V2 Test ---------------------")
	uid := uint64(0)
	u := tc.CbrChain1.Users[uid]
	balBefore, err := tc.CbrChain2.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(100 * 1e6)
	err = tc.CbrChain1.ApproveBridgeTestToken(tc.CbrChain1.USDTContract, uid, amount, tc.CbrChain1.MsgTestAddr)
	tc.ChkErr(err, "approve USDT")
	depositId, err := tc.CbrChain1.SendMessageWithPegTransfer(
		uid, tc.CbrChain2.MsgTestAddr, tc.CbrChain1.USDTAddr, amount, tc.CbrChain2.ChainId, tc.BrSendTypePegV2Deposit)
	tc.ChkErr(err, "SendWithTransfer, peg deposit")
	depositInfo := tc.WaitPbrDeposit(transactor, depositId.String())
	if len(depositInfo.MintId) == 0 {
		log.Fatalln("refunded deposit", nil)
	}
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_VAULT, depositId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_PEG_V2_MINT)

	mintInfo := tc.CheckPbrMint(transactor, eth.Bytes2Hex(depositInfo.GetMintId()))
	var mintOnChain pegbrtypes.MintOnChain
	err = proto.Unmarshal(mintInfo.MintProtoBytes, &mintOnChain)
	tc.ChkErr(err, "unmarshal MintOnChain")
	mintAmt := new(big.Int).SetBytes(mintOnChain.Amount)
	balAfter, err := tc.CbrChain2.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	balExp := new(big.Int).Add(balBefore, mintAmt)
	log.Infof("bal after, %s", balAfter.String())
	if balAfter.Cmp(balExp) == 0 {
		return
	}
	log.Fatalf("peg deposit-mint balAfter (%s) not as expected (%s)", balAfter, balExp)
}

func pegV2BurnWithdrawTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Peg Deposit V2 Test ---------------------")
	u := tc.CbrChain1.Users[1]
	balBefore, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(8 * 1e6)
	err = tc.CbrChain2.ApproveBridgeTestToken(tc.CbrChain2.USDTContract, 1, amount, tc.CbrChain2.MsgTestAddr)
	tc.ChkErr(err, "approve USDT")
	burnId, err := tc.CbrChain2.SendMessageWithPegTransfer(
		1, tc.CbrChain1.MsgTestAddr, tc.CbrChain2.USDTAddr, amount, tc.CbrChain1.ChainId, tc.BrSendTypePegV2Burn)
	tc.ChkErr(err, "SendWithTransfer, peg burn-withdraw")
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_BRIDGE, burnId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_PEG_V2_WITHDRAW)

	burnInfo := tc.WaitPbrBurn(transactor, burnId.String())
	withdrawInfo := tc.CheckPbrWithdraw(transactor, eth.Bytes2Hex(burnInfo.WithdrawId))
	var withdrawOnChain pegbrtypes.WithdrawOnChain
	err = proto.Unmarshal(withdrawInfo.WithdrawProtoBytes, &withdrawOnChain)
	tc.ChkErr(err, "unmarshal WithdrawOnChain")
	withdrawAmt := new(big.Int).SetBytes(withdrawOnChain.Amount)
	balAfter, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	balExp := new(big.Int).Add(balBefore, withdrawAmt)
	log.Infof("bal after, %s", balAfter.String())
	if balAfter.Cmp(balExp) == 0 {
		return
	}
	log.Fatalf("peg burn-withdraw balAfter (%s) not as expected (%s)", balAfter, balExp)
}

func pegV2BurnMintTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Peg Deposit V2 Test ---------------------")
	u := tc.CbrChain1.Users[2]
	balBefore, err := tc.CbrChain3.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(5 * 1e6)
	err = tc.CbrChain2.ApproveBridgeTestToken(tc.CbrChain2.USDTContract, 2, amount, tc.CbrChain2.MsgTestAddr)
	tc.ChkErr(err, "approve USDT")
	burnId, err := tc.CbrChain2.SendMessageWithPegTransfer(
		2, tc.CbrChain3.MsgTestAddr, tc.CbrChain2.USDTAddr, amount, tc.CbrChain3.ChainId, tc.BrSendTypePegV2Burn)
	tc.ChkErr(err, "SendWithTransfer, peg burn-mint")
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_BRIDGE, burnId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_PEG_V2_MINT)

	burnInfo := tc.WaitPbrBurn(transactor, burnId.String())
	mintInfo := tc.CheckPbrMint(transactor, eth.Bytes2Hex(burnInfo.GetMintId()))
	var mintOnChain pegbrtypes.MintOnChain
	err = proto.Unmarshal(mintInfo.MintProtoBytes, &mintOnChain)
	tc.ChkErr(err, "unmarshal MintOnChain")
	mintAmt := new(big.Int).SetBytes(mintOnChain.Amount)
	balAfter, err := tc.CbrChain3.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	balExp := new(big.Int).Add(balBefore, mintAmt)
	log.Infof("bal after, %s", balAfter.String())
	if balAfter.Cmp(balExp) == 0 {
		return
	}
	log.Fatalf("peg burn-mint balAfter (%s) not as expected (%s)", balAfter, balExp)
}

func refundTransferTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund Test ---------------------")
	u := tc.CbrChain1.Users[0]
	balBefore, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(10000 * 1e6)
	err = tc.CbrChain1.ApproveBridgeTestToken(tc.CbrChain1.USDTContract, 0, amount, tc.CbrChain1.MsgTestAddr)
	tc.ChkErr(err, "approve USDT")
	xferId, err := tc.CbrChain1.SendMessageWithLiquidityTransfer(
		0, tc.CbrChain2.MsgTestAddr, tc.CbrChain1.USDTAddr, amount, tc.CbrChain2.ChainId, 1)
	tc.ChkErr(err, "SendWithTransfer")
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_LIQUIDITY, xferId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_LIQUIDITY_WITHDRAW)
	balAfter, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	log.Infof("bal after, %s", balAfter)
	if balAfter.Cmp(balBefore) == 0 {
		return
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}

func refundPegDepositTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund Peg Deposit Test ---------------------")
	u := tc.CbrChain1.Users[1]
	balBefore, err := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(10)
	err = tc.CbrChain1.ApproveBridgeTestToken(tc.CbrChain1.UNIContract, 1, amount, tc.CbrChain1.MsgTestAddr)
	tc.ChkErr(err, "approve UNI")
	xferId, err := tc.CbrChain1.SendMessageWithPegTransfer(
		1, tc.CbrChain2.MsgTestAddr, tc.CbrChain1.UNIAddr, amount, tc.CbrChain2.ChainId, tc.BrSendTypePegDeposit)
	tc.ChkErr(err, "SendWithTransfer, peg deposit")
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_VAULT, xferId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_PEG_WITHDRAW)
	balAfter, err := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, u.Address)
	log.Infof("bal after, %s", balAfter)
	if balAfter.Cmp(balBefore) == 0 {
		return
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}

func refundPegV2DepositTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund PegV2 Deposit Test ---------------------")
	u := tc.CbrChain1.Users[1]
	balBefore, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(10)
	err = tc.CbrChain1.ApproveBridgeTestToken(tc.CbrChain1.USDTContract, 1, amount, tc.CbrChain1.MsgTestAddr)
	tc.ChkErr(err, "approve USDT")
	xferId, err := tc.CbrChain1.SendMessageWithPegTransfer(
		1, tc.CbrChain2.MsgTestAddr, tc.CbrChain1.USDTAddr, amount, tc.CbrChain2.ChainId, tc.BrSendTypePegV2Deposit)
	tc.ChkErr(err, "SendWithTransfer, peg deposit")
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_VAULT, xferId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_PEG_V2_WITHDRAW)
	balAfter, err := tc.CbrChain1.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	log.Infof("bal after, %s", balAfter)
	if balAfter.Cmp(balBefore) == 0 {
		return
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}

func refundPegBurnTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund Peg Burn Test ---------------------")
	u := tc.CbrChain1.Users[2]
	balBefore, err := tc.CbrChain2.UNIContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(10)
	err = tc.CbrChain2.ApproveBridgeTestToken(tc.CbrChain2.UNIContract, 2, amount, tc.CbrChain2.MsgTestAddr)
	tc.ChkErr(err, "approve UNI")
	xferId, err := tc.CbrChain2.SendMessageWithPegTransfer(
		2, tc.CbrChain1.MsgTestAddr, tc.CbrChain2.UNIAddr, amount, tc.CbrChain1.ChainId, tc.BrSendTypePegBurn)
	tc.ChkErr(err, "SendWithTransfer, peg burn")
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_BRIDGE, xferId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_PEG_MINT)
	balAfter, err := tc.CbrChain2.UNIContract.BalanceOf(&bind.CallOpts{}, u.Address)
	log.Infof("bal after, %s", balAfter)
	if balAfter.Cmp(balBefore) == 0 {
		return
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}

func refundPegV2BurnTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("-------------------- Refund PegV2 Burn Test ---------------------")
	u := tc.CbrChain1.Users[2]
	balBefore, err := tc.CbrChain2.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	tc.ChkErr(err, "bal before")
	log.Infoln("bal before:", balBefore)
	amount := big.NewInt(10)
	err = tc.CbrChain2.ApproveBridgeTestToken(tc.CbrChain2.USDTContract, 2, amount, tc.CbrChain2.MsgTestAddr)
	tc.ChkErr(err, "approve USDT")
	xferId, err := tc.CbrChain2.SendMessageWithPegTransfer(
		2, tc.CbrChain1.MsgTestAddr, tc.CbrChain2.USDTAddr, amount, tc.CbrChain1.ChainId, tc.BrSendTypePegV2Burn)
	tc.ChkErr(err, "SendWithTransfer, peg burn")
	tc.WaitForMessageWithTransferExecuted(
		transactor, msgtypes.BRIDGE_TYPE_PEG_BRIDGE, xferId, msgtypes.EXECUTION_STATUS_SUCCESS, msgtypes.TRANSFER_TYPE_PEG_V2_MINT)
	balAfter, err := tc.CbrChain2.USDTContract.BalanceOf(&bind.CallOpts{}, u.Address)
	log.Infof("bal after, %s", balAfter)
	if balAfter.Cmp(balBefore) == 0 {
		return
	}
	log.Fatalf("balAfter (%s) not equal to balBefore (%s)", balAfter, balBefore)
}
