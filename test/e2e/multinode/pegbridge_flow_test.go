package multinode

import (
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	flowutils "github.com/celer-network/cbridge-flow/utils"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/transactor"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

func TestFlowPegbridge(t *testing.T) {
	t.Run("e2e-flow-pegbridge", func(t *testing.T) {
		t.Run("pegbridgeFlowTest", pegbridgeFlowTest) // comment this out when commit, as it duplicates TestBridge
	})
}

// Test pegbridge
func pegbridgeFlowTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test flow pegbridge ===========================")
	setupBridgeTest()

	pbrFlowTest(t, prepareValidators(t))
}

func pbrFlowTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("======================== Test SafeBox deposit -> mint ===========================")
	log.Infof("start deposit on flow")
	nonce := uint64(time.Now().Second())
	_, err := flowutils.Deposit(context.Background(), tc.FlowUserAccountClient, 884,
		uint64(time.Now().Second()), "100.0", tc.CbrChain1.Users[0].Address.String(), tc.FlowContractAddr.String(), exampleTokenVault)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("send deposit on flow")
	rawDpIdData := fmt.Sprintf("0x%sA.%s.%s.Vault%s%d", tc.FlowUserAddr.String(), tc.FlowContractAddr.String(), exampleTokenName, "100.00000000", nonce)
	depositId := fmt.Sprintf("%x", sha3.Sum256([]byte(rawDpIdData)))
	log.Infof("SafeBox DepositId:%s, raw:%s", depositId, rawDpIdData)
	depositInfo := tc.WaitPbrDeposit(transactor, depositId)
	mintInfo := tc.CheckPbrMint(transactor, eth.Bytes2Hex(depositInfo.MintId))
	log.Infof("SafeBox Deposit related MintInfo:%+v", mintInfo)

	log.Infoln("======================== Test SafeBox burn -> withdraw ===========================")
	burnFlowAmt := new(big.Int).Mul(big.NewInt(5), big.NewInt(1e18))
	burnFlowId, err := tc.CbrChain2.PbrBurnWithUser(0, burnFlowAmt, rand.Uint64(), common.BytesToAddress(tc.FlowUserAddr.Bytes()))
	tc.ChkErr(err, "u0 chain2 burn")
	burnInfo := tc.WaitPbrBurn(transactor, burnFlowId[:])
	log.Infof("burnInfo: %+v", burnInfo)
	wdInfo := tc.CheckPbrWithdraw(transactor, eth.Bytes2Hex(burnInfo.WithdrawId))
	log.Infof("wdInfo: %+v", wdInfo)

	log.Infoln("======================== Test PegBridge Deposit -> Mint ===========================")
	depositAmt := new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18))
	depositId, err = tc.CbrChain1.PbrDepositWithMintAccount(0, depositAmt, 12340003, common.HexToAddress(tc.FlowUserAddr.String()), rand.Uint64())
	tc.ChkErr(err, "u0 chain1 deposit")
	log.Infof("depositId:%s", depositId)
	depositInfo = tc.WaitPbrDeposit(transactor, depositId)
	log.Infof("deposit info:%+v", depositInfo)
	mintInfo = tc.CheckPbrMint(transactor, eth.Bytes2Hex(depositInfo.MintId))
	log.Infof("PegBridge MintInfo:%+v", mintInfo)

	log.Infoln("======================== Test PegBridge Burn -> Withdraw ===========================")
	log.Infof("start burn on flow")
	nonce = uint64(time.Now().Second())
	_, err = flowutils.Burn(context.Background(), tc.FlowUserAccountClient, 883,
		nonce, "1.0", tc.CbrChain1.Users[0].Address.String(), tc.FlowContractAddr.String(), testPegTokenVault)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("send burn on flow")
	// 0x179b6b1cb6755e31A.01cf0e2f2f715450.PegToken.Vault5.0000000019
	rawBurnIdData := fmt.Sprintf("0x%sA.%s.%s.Vault%s%d", tc.FlowUserAddr.String(), tc.FlowContractAddr.String(), testPegTokenName, "1.00000000", nonce)
	burnId := fmt.Sprintf("%x", sha3.Sum256([]byte(rawBurnIdData)))
	log.Infof("PegBridge BurnId:%s, raw:%s", burnId, rawBurnIdData)
	burnInfo = tc.WaitPbrBurn(transactor, burnId)
	wdInfo = tc.CheckPbrWithdraw(transactor, eth.Bytes2Hex(burnInfo.WithdrawId))
	log.Infof("PegBridge Burn related WdInfo:%+v", wdInfo)
}
