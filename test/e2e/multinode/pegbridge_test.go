package multinode

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/transactor"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	proto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestPegbridge(t *testing.T) {
	t.Run("e2e-pegbridge", func(t *testing.T) {
		//t.Run("pegbridgeTest", pegbridgeTest)
	})
}

// Test pegbridge
func pegbridgeTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test pegbridge ===========================")
	setupBridgeTest()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	prepareValidators(t, transactor)

	pbrTest(t, transactor)
}

func pbrTest(t *testing.T, transactor *transactor.Transactor) {
	log.Infoln("======================== Deposit ===========================")
	supplyCap := tc.GetSupplyCap(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.UNIAddr)
	log.Infoln("Supply cap: ", supplyCap)
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.UNIAddr, "0")
	log.Infoln("total supply: 0")
	depositAmt := new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18))
	err := tc.CbrChain1.ApproveUNI(0, depositAmt)
	tc.ChkErr(err, "u0 chain1 approve")
	depositId, err := tc.CbrChain1.PbrDeposit(0, depositAmt, tc.CbrChain2.ChainId, rand.Uint64())
	tc.ChkErr(err, "u0 chain1 deposit")

	depositInfo := tc.WaitPbrDeposit(transactor, depositId)
	mintInfo := tc.CheckPbrMint(transactor, eth.Bytes2Hex(depositInfo.MintId))
	var mintOnChain pegbrtypes.MintOnChain
	err = proto.Unmarshal(mintInfo.MintProtoBytes, &mintOnChain)
	tc.ChkErr(err, "unmarshal MintOnChain")
	mintAmt := new(big.Int).SetBytes(mintOnChain.Amount)
	tc.CbrChain2.CheckPeggedUNIBalance(0, mintAmt)
	log.Infoln("depositAmt", depositAmt, "mintAmt", mintAmt)
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.UNIAddr, mintAmt.String())
	log.Infoln("total supply:", mintAmt)

	log.Infoln("======================== Deposit that would exceed supply cap===========================")
	err = tc.CbrChain1.ApproveUNI(0, depositAmt.Mul(depositAmt, big.NewInt(2)))
	tc.ChkErr(err, "u0 chain1 approve")
	depositId, err = tc.CbrChain1.PbrDeposit(0, depositAmt, tc.CbrChain2.ChainId, rand.Uint64())
	tc.ChkErr(err, "u0 chain1 deposit")
	err = tc.WaitPbrDepositWithEmptyMintId(transactor, depositId)
	tc.ChkErr(err, "wait pbr deposit with empty mintId")
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.UNIAddr, mintAmt.String())
	log.Infoln("total supply:", mintAmt)

	log.Infoln("======================== Claim refund for previous failed deposit===========================")
	balanceBefore, err := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, tc.CbrChain1.Users[0].Address)
	tc.ChkErr(err, "u0 balance before burn")
	log.Infoln("user0 uni balance is:", balanceBefore)
	err = tc.StartClaimPegbridgeRefund(transactor, depositId)
	tc.ChkErr(err, "user0 tried to claim pegbridge refund")
	log.Infoln("user0 claim refund initiated with success")
	withdrawId, withdrawInfo := tc.GetRefundWithdrawInfoWithSigs(transactor, depositId, 3)
	log.Infoln("user0 claim pegbridge refund, withdrawId:", withdrawId)
	curss, err := tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainPegVaultWithdraw(withdrawInfo, curss)
	tc.ChkErr(err, "chain1 onchain withdraw pegbridge refund")
	balanceAfter := new(big.Int).Add(balanceBefore, depositAmt)
	tc.CbrChain1.CheckUNIBalance(0, balanceAfter)
	log.Infoln("user0 claims refund with success, balance before:", balanceBefore, " balance after:", balanceAfter)
	err = tc.FakeStartClaimPegbridgeRefund(transactor, depositId)
	tc.ChkErr(err, "user0 tried to claim pegbridge refund again")
	log.Infoln("user0 failed to twice claim refund of the same invalid deposit")

	log.Infoln("======================== Burn ===========================")
	balance1, err := tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, tc.CbrChain1.Users[0].Address)
	tc.ChkErr(err, "u0 balance before burn")
	burnAmt := new(big.Int).Mul(big.NewInt(50), big.NewInt(1e18))
	burnId, err := tc.CbrChain2.PbrBurn(0, burnAmt, rand.Uint64())
	tc.ChkErr(err, "u0 chain2 burn")
	tc.RunAllAndWait(
		func() {
			burnInfo := tc.WaitPbrBurn(transactor, burnId[:])
			withdrawInfo = tc.CheckPbrWithdraw(transactor, eth.Bytes2Hex(burnInfo.WithdrawId))
			var withdrawOnChain pegbrtypes.WithdrawOnChain
			err = proto.Unmarshal(withdrawInfo.WithdrawProtoBytes, &withdrawOnChain)
			tc.ChkErr(err, "unmarshal BurnOnChain")
			withdrawAmt := new(big.Int).SetBytes(withdrawOnChain.Amount)
			tc.CbrChain1.CheckUNIBalance(0, new(big.Int).Add(balance1, withdrawAmt))
			log.Infoln("burnAmt", burnAmt, "withdrawAmt", withdrawAmt)
			expectedTotalSupply := new(big.Int).Sub(mintAmt, burnAmt)
			tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.UNIAddr, expectedTotalSupply.String())
			log.Infoln("total supply:", expectedTotalSupply)
		},
		func() {
			log.Infoln("======================== Burn that would be refunded ===========================")
			balance2, err := tc.CbrChain2.UNIContract.BalanceOf(&bind.CallOpts{}, tc.CbrChain2.Users[0].Address)
			tc.ChkErr(err, "u0 balance before burn")
			log.Infoln("user0 pegged uni balance is:", balance2)
			burnAmt := new(big.Int).Mul(big.NewInt(1), big.NewInt(1e15))
			burnId, err := tc.CbrChain2.PbrBurn(0, burnAmt, rand.Uint64())
			tc.ChkErr(err, "u0 chain2 burn")
			err = tc.WaitPbrBurnWithEmptyWithdrawId(transactor, burnId)
			tc.ChkErr(err, "wait pbr burn with empty withdrawId")
			tc.CbrChain2.CheckPeggedUNIBalance(0, new(big.Int).Sub(balance2, burnAmt))

			log.Infoln("======================== Claim refund for previous failed burn===========================")
			err = tc.StartClaimPegbridgeRefund(transactor, burnId)
			tc.ChkErr(err, "user0 tried to claim pegbridge refund")
			log.Infoln("user0 claim refund initiated with success")
			mintId, mintInfo := tc.GetRefundMintInfoWithSigs(transactor, burnId, 3)
			log.Infoln("user0 claim pegbridge refund, mintId:", mintId)
			curss2, err := tc.GetCurSortedSigners(transactor, tc.CbrChain2.ChainId)
			tc.ChkErr(err, "chain2 GetCurSortedSigners")
			err = tc.CbrChain2.OnchainPegBridgeMint(mintInfo, curss2)
			tc.ChkErr(err, "chain2 onchain mint pegbridge refund")
			tc.CbrChain2.CheckPeggedUNIBalance(0, balance2)
			log.Infoln("user0 claims refund with success, balance:", balance2)
		},
	)

	tc.RunAllAndWait(
		func() {
			log.Infoln("======================== Delegator 0 claim fee ===========================")
			feesInfo, err := tc.GetPegBridgeFeesInfo(transactor, 0)
			tc.ChkErr(err, "del0 get pegbridge fees info before claim")
			log.Infoln("feesInfo.ClaimableFeeAmounts before claim", feesInfo.ClaimableFeeAmounts)
			assert.Equal(t, 1, len(feesInfo.ClaimableFeeAmounts), "Should have 1 fee")
			fee0 := feesInfo.ClaimableFeeAmounts[0]
			assert.Equal(t, fmt.Sprintf("PBF-UNI/%d", tc.CbrChain1.ChainId), fee0.Denom)
			assert.True(t, fee0.Amount.GT(sdk.NewDec(1e14)))
			assert.True(t, fee0.Amount.LT(sdk.NewDec(1e15)))

			nonce := rand.Uint64()
			err = tc.CbrChain1.StartDelegatorClaimPegBridgeFee(transactor, 0, tc.CbrChain1.ChainId, tc.CbrChain1.UNIAddr, nonce)
			tc.ChkErr(err, "del0 chain1 start claim pegbridge fee")
			withdrawId, withdrawInfo = tc.GetPegBridgeFeeClaimWithdrawInfoWithSigs(
				transactor, tc.CbrChain1.Delegators[0].Address, nonce, 3)
			log.Infoln("del0 claim pegbridge fees withdrawId:", withdrawId)
			curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
			tc.ChkErr(err, "chain1 GetCurSortedSigners")
			err = tc.CbrChain1.OnchainPegVaultWithdraw(withdrawInfo, curss)
			tc.ChkErr(err, "chain1 onchain withdraw pegbridge fee")

			feesInfo, err = tc.GetPegBridgeFeesInfo(transactor, 0)
			tc.ChkErr(err, "del0 get pegbridge fees info after claim")
			log.Infoln("feesInfo.ClaimableFeeAmounts after claim", feesInfo.ClaimableFeeAmounts)
			assert.Equal(t, 0, len(feesInfo.ClaimableFeeAmounts), "Should have 0 fee")
		},
		func() {
			log.Infoln("======================== Validator 0 claim fee ===========================")
			nonce := rand.Uint64()
			err = tc.CbrChain1.StartValidatorClaimPegBridgeFee(transactor, 0, tc.CbrChain1.ChainId, tc.CbrChain1.UNIAddr, nonce)
			tc.ChkErr(err, "val0 chain1 start claim pegbridge fee")
			withdrawId, withdrawInfo = tc.GetPegBridgeFeeClaimWithdrawInfoWithSigs(
				transactor, tc.CbrChain1.Validators[0].Address, nonce, 3)
			log.Infoln("val0 claim pegbridge fees withdrawId:", withdrawId)
			curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
			tc.ChkErr(err, "chain1 GetCurSortedSigners")
			err = tc.CbrChain1.OnchainPegVaultWithdraw(withdrawInfo, curss)
			tc.ChkErr(err, "chain1 onchain withdraw pegbridge fee")
		},
		func() {
			log.Infoln("======================== Validator 1 claim fee by himself and without sig ===========================")
			nonce := rand.Uint64()
			err = tc.StartValidatorSelfClaimPegbrFee(1, tc.CbrChain1.ChainId, tc.CbrChain1.UNIAddr, nonce)
			tc.ChkErr(err, "val1 chain1 start claim pegbridge fee")
			withdrawId, withdrawInfo = tc.GetPegBridgeFeeClaimWithdrawInfoWithSigs(
				transactor, tc.CbrChain1.Validators[1].Address, nonce, 3)
			log.Infoln("val1 claim pegbridge fees withdrawId:", withdrawId)
			curss, err = tc.GetCurSortedSigners(transactor, tc.CbrChain1.ChainId)
			tc.ChkErr(err, "chain1 GetCurSortedSigners")
			err = tc.CbrChain1.OnchainPegVaultWithdraw(withdrawInfo, curss)
			tc.ChkErr(err, "chain1 onchain withdraw pegbridge fee")
		},
	)
	log.Infoln("======================== Finish PegBridge Test ===========================")
}
