package multinode

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func setupPegbridge() {
	log.Infoln("Set up another mainchain for bridge")
	SetupMainchain2ForBridge()

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
	SetupNewSgnEnv(p, true, false, false)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestPegbridge(t *testing.T) {
	t.Run("e2e-pegbridge", func(t *testing.T) {
		t.Run("pegbridgeTest", pegbridgeTest)
	})
}

// Test pegbridge
func pegbridgeTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test pegbridge ===========================")
	setupPegbridge()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	log.Infoln("================== Setup validators and bridge signers ======================")
	// Make the stake amounts more realistic to test precision handling when distributing fee share
	vAmts := []*big.Int{
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(2e8), big.NewInt(1e18)),
	}
	vAddrs := []eth.Addr{
		tc.ValEthAddrs[0], tc.ValEthAddrs[1], tc.ValEthAddrs[2],
	}
	err := tc.FundAddrsErc20(tc.CelrAddr, vAddrs, vAmts[0], tc.EthClient, tc.EtherBaseAuth)
	tc.ChkErr(err, "fund validator accounts")
	numVals := len(vAmts)
	tc.SetupValidators(t, transactor, vAmts)
	tc.CbrChain1.SetInitSigners(vAmts)
	tc.CbrChain2.SetInitSigners(vAmts)
	expSigners := genSortedSigners([]eth.Addr{tc.ValSignerAddrs[0], tc.ValSignerAddrs[1], tc.ValSignerAddrs[2]}, vAmts)
	tc.CheckChainSigners(t, transactor, tc.CbrChain1.ChainId, expSigners)
	tc.CheckChainSigners(t, transactor, tc.CbrChain2.ChainId, expSigners)

	log.Infoln("================== Delegate from delegator 0 to all validators ======================")
	dAmts := []*big.Int{
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
		new(big.Int).Mul(big.NewInt(1e6), big.NewInt(1e18)),
	}
	dAddrs := []eth.Addr{tc.DelEthAddrs[0]}
	err = tc.FundAddrsErc20(tc.CelrAddr, dAddrs, new(big.Int).Mul(big.NewInt(3), dAmts[0]), tc.EthClient, tc.EtherBaseAuth)
	tc.ChkErr(err, "fund delegator account")
	for i := 0; i < numVals; i++ {
		tc.Delegate(tc.DelAuths[0], tc.ValEthAddrs[i], dAmts[i])
	}
	for i := 0; i < numVals; i++ {
		expDel := &stakingtypes.Delegation{
			DelegatorAddress: eth.Addr2Hex(tc.DelEthAddrs[0]),
			ValidatorAddress: eth.Addr2Hex(tc.ValEthAddrs[i]),
			Shares:           sdk.NewIntFromBigInt(dAmts[i]),
		}
		tc.CheckDelegation(t, transactor, expDel)
	}

	log.Infoln("======================== Deposit ===========================")
	supplyCap := tc.GetSupplyCap(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.PeggedUNIAddr)
	log.Infoln("Supply cap: ", supplyCap)
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.PeggedUNIAddr, "0")
	log.Infoln("total supply: 0")
	depositAmt := new(big.Int).Mul(big.NewInt(100), big.NewInt(1e18))
	err = tc.CbrChain1.ApproveUNI(0, depositAmt)
	tc.ChkErr(err, "u0 chain1 approve")
	randBytes := crypto.Keccak256Hash([]byte(fmt.Sprintf("%d-%d-%d-%x", tc.CbrChain1.ChainId, 1337, tc.CbrChain2.ChainId, tc.CbrChain1.Users[0].Address.Bytes())))
	randSeed := new(big.Int).SetBytes(randBytes.Bytes()).Int64()
	r := rand.New(rand.NewSource(randSeed))
	nonce := r.Uint64()
	depositId, err := tc.CbrChain1.PbrDeposit(0, depositAmt, tc.CbrChain2.ChainId, nonce)
	tc.ChkErr(err, "u0 chain1 deposit")

	depositInfo := tc.WaitPbrDeposit(transactor, depositId)
	mintInfo := tc.CheckPbrMint(transactor, eth.Bytes2Hex(depositInfo.MintId))
	var mintOnChain pegbrtypes.MintOnChain
	err = proto.Unmarshal(mintInfo.MintProtoBytes, &mintOnChain)
	tc.ChkErr(err, "unmarshal MintOnChain")
	mintAmt := new(big.Int).SetBytes(mintOnChain.Amount)
	tc.CbrChain2.CheckPeggedUNIBalance(0, mintAmt)
	log.Infoln("depositAmt", depositAmt, "mintAmt", mintAmt)
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.PeggedUNIAddr, mintAmt.String())
	log.Infoln("total supply:", mintAmt)

	log.Infoln("======================== Deposit that would exceed supply cap===========================")
	err = tc.CbrChain1.ApproveUNI(0, depositAmt.Mul(depositAmt, big.NewInt(2)))
	tc.ChkErr(err, "u0 chain1 approve")
	randBytes = crypto.Keccak256Hash([]byte(fmt.Sprintf("%d-%d-%d-%x", tc.CbrChain1.ChainId, 1666, tc.CbrChain2.ChainId, tc.CbrChain1.Users[0].Address.Bytes())))
	randSeed = new(big.Int).SetBytes(randBytes.Bytes()).Int64()
	r = rand.New(rand.NewSource(randSeed))
	nonce = r.Uint64()
	depositId, err = tc.CbrChain1.PbrDeposit(0, depositAmt, tc.CbrChain2.ChainId, nonce)
	tc.ChkErr(err, "u0 chain1 deposit")
	err = tc.WaitPbrDepositWithEmptyMintId(transactor, depositId)
	tc.ChkErr(err, "wait pbr deposit with empty mintId")
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.PeggedUNIAddr, mintAmt.String())
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
	balanceBefore, err = tc.CbrChain1.UNIContract.BalanceOf(&bind.CallOpts{}, tc.CbrChain1.Users[0].Address)
	tc.ChkErr(err, "u0 balance before burn")
	burnAmt := new(big.Int).Mul(big.NewInt(50), big.NewInt(1e18))
	// burn action does not need approve.
	//err = tc.CbrChain2.ApprovePeggedUNI(0, burnAmt)
	//tc.ChkErr(err, "u0 chain2 approve")
	burnId, err := tc.CbrChain2.PbrBurn(0, burnAmt, 1234)
	tc.ChkErr(err, "u0 chain2 burn")

	burnInfo := tc.WaitPbrBurn(transactor, burnId[:])
	withdrawInfo = tc.CheckPbrWithdraw(transactor, eth.Bytes2Hex(burnInfo.WithdrawId))
	var withdrawOnChain pegbrtypes.WithdrawOnChain
	err = proto.Unmarshal(withdrawInfo.WithdrawProtoBytes, &withdrawOnChain)
	tc.ChkErr(err, "unmarshal BurnOnChain")
	withdrawAmt := new(big.Int).SetBytes(withdrawOnChain.Amount)
	tc.CbrChain1.CheckUNIBalance(0, balanceBefore.Add(balanceBefore, withdrawAmt))
	log.Infoln("burnAmt", burnAmt, "withdrawAmt", withdrawAmt)
	expectedTotalSupply := new(big.Int).Sub(mintAmt, burnAmt)
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.PeggedUNIAddr, expectedTotalSupply.String())
	log.Infoln("total supply:", expectedTotalSupply)

	log.Infoln("======================== Burn that would be refunded ===========================")
	balanceBefore, err = tc.CbrChain2.PeggedUNIContract.BalanceOf(&bind.CallOpts{}, tc.CbrChain2.Users[0].Address)
	tc.ChkErr(err, "u0 balance before burn")
	log.Infoln("user0 pegged uni balance is:", balanceBefore)
	burnAmt = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e15))
	burnId, err = tc.CbrChain2.PbrBurn(0, burnAmt, 4321)
	tc.ChkErr(err, "u0 chain2 burn")
	err = tc.WaitPbrBurnWithEmptyWithdrawId(transactor, burnId)
	tc.ChkErr(err, "wait pbr burn with empty withdrawId")
	expectedTotalSupply.Sub(expectedTotalSupply, burnAmt)
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.PeggedUNIAddr, expectedTotalSupply.String())
	tc.CbrChain2.CheckPeggedUNIBalance(0, new(big.Int).Sub(balanceBefore, burnAmt))

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
	tc.CbrChain2.CheckPeggedUNIBalance(0, balanceBefore)
	log.Infoln("user0 claims refund with success, balance before:", balanceBefore, " balance after:", balanceBefore)
	expectedTotalSupply.Add(expectedTotalSupply, burnAmt)
	tc.CheckTotalSupply(transactor, tc.CbrChain2.ChainId, tc.CbrChain2.PeggedUNIAddr, expectedTotalSupply.String())

	log.Infoln("======================== Delegator 0 claim fee ===========================")
	feesInfo, err := tc.GetPegBridgeFeesInfo(transactor, 0)
	tc.ChkErr(err, "del0 get pegbridge fees info before claim")
	log.Infoln("feesInfo.ClaimableFeeAmounts before claim", feesInfo.ClaimableFeeAmounts)
	assert.Equal(t, 1, len(feesInfo.ClaimableFeeAmounts), "Should have 1 fee")
	fee0 := feesInfo.ClaimableFeeAmounts[0]
	assert.Equal(t, fmt.Sprintf("PBF-UNI/%d", tc.CbrChain1.ChainId), fee0.Denom)
	assert.True(t, fee0.Amount.GT(sdk.NewDec(1e14)))
	assert.True(t, fee0.Amount.LT(sdk.NewDec(1e15)))

	nonce = uint64(time.Now().Unix())
	err = tc.CbrChain1.StartClaimPegBridgeFee(transactor, 0, tc.CbrChain1.ChainId, tc.CbrChain1.UNIAddr, nonce)
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
}
