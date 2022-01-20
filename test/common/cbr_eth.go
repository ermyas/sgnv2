package common

import (
	"context"
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func InitCbrChainConfigs() {
	CbrChain1 = &CbrChain{
		ChainId: ChainID,
		Ec:      EthClient,
		Auth:    EtherBaseAuth,
	}
	CbrChain1.SetupTestEthClients()

	rpcClient, err := rpc.Dial(LocalGeth2)
	if err != nil {
		log.Fatal(err)
	}
	_, etherBaseAuth, err := GetAuth(etherBaseKs, int64(Geth2ChainID))
	if err != nil {
		log.Fatal(err)
	}

	CbrChain2 = &CbrChain{
		ChainId: Geth2ChainID,
		Ec:      ethclient.NewClient(rpcClient),
		Auth:    etherBaseAuth,
	}
	CbrChain2.SetupTestEthClients()
}

func (c *CbrChain) SetupTestEthClients() {
	users := []*TestEthClient{}
	for _, clientKs := range ClientEthKs {
		u, err := SetupTestEthClient(clientKs, c.ChainId)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	c.Users = users

	vals := []*TestEthClient{}
	for _, valKs := range ValEthKs {
		val, err := SetupTestEthClient(valKs, c.ChainId)
		if err != nil {
			log.Fatal(err)
		}
		vals = append(vals, val)
	}
	c.Validators = vals

	valSigners := []*TestEthClient{}
	for _, valSignerKs := range ValSignerKs {
		valSigner, err := SetupTestEthClient(valSignerKs, c.ChainId)
		if err != nil {
			log.Fatal(err)
		}
		valSigners = append(valSigners, valSigner)
	}
	c.ValidatorSigners = valSigners

	dels := []*TestEthClient{}
	for _, delKs := range DelEthKs {
		del, err := SetupTestEthClient(delKs, c.ChainId)
		if err != nil {
			log.Fatal(err)
		}
		dels = append(dels, del)
	}
	c.Delegators = dels
}

func (c *CbrChain) ApproveUSDT(uid uint64, amt *big.Int) error {
	return c.ApproveErc20(c.USDTContract, uid, amt, c.CbrAddr)
}

func (c *CbrChain) ApproveUNI(uid uint64, amt *big.Int) error {
	return c.ApproveErc20(c.UNIContract, uid, amt, c.PegVaultAddr)
}

func (c *CbrChain) ApproveUNIForBatchTransfer(uid uint64, amt *big.Int) error {
	return c.ApproveErc20(c.UNIContract, uid, amt, c.BatchTransferAddr)
}

func (c *CbrChain) ApproveErc20(erc20 *eth.Erc20, uid uint64, amt *big.Int, spender eth.Addr) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	tx, err := erc20.Approve(c.Users[uid].Auth, spender, amt)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "Approve")
	return nil
}

func (c *CbrChain) ApprovePeggedUNI(uid uint64, amt *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	tx, err := c.PeggedUNIContract.Approve(c.Users[uid].Auth, c.PegBridgeAddr, amt)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "Approve")
	return nil
}

func (c *CbrChain) ApprovePeggedUNIForBatchTransfer(uid uint64, amt *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	tx, err := c.PeggedUNIContract.Approve(c.Users[uid].Auth, c.BatchTransferAddr, amt)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "Approve")
	return nil
}

func (c *CbrChain) AddLiq(uid uint64, amt *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	tx, err := c.CbrContract.AddLiquidity(c.Users[uid].Auth, c.USDTAddr, amt)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "AddLiquidity")
	return nil
}

// only used for test
func (c *CbrChain) Send(uid uint64, amt *big.Int, dstChainId, nonce uint64) (eth.Hash, error) {
	return c.SendAny(uid, uid, amt, dstChainId, nonce, 100000) // 10% slippage
}

func (c *CbrChain) SendAny(fromUid, toUid uint64, amt *big.Int, dstChainId, nonce uint64, maxSlippage uint32) (eth.Hash, error) {
	tx, err := c.CbrContract.Send(
		c.Users[fromUid].Auth, c.Users[toUid].Address, c.USDTAddr, amt, dstChainId, nonce, maxSlippage)
	if err != nil {
		return eth.ZeroHash, err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return eth.ZeroHash, err
	}
	sendLog := receipt.Logs[len(receipt.Logs)-1] // last log is Send event (NOTE Polygon breaks this assumption)
	sendEv, err := c.CbrContract.ParseSend(*sendLog)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("parse log %+v err: %w", sendLog, err)
	}
	return sendEv.TransferId, nil
}

func (c *CbrChain) OnchainCbrWithdraw(wdDetail *cbrtypes.WithdrawDetail, signers []*cbrtypes.Signer) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	addrs, powers := cbrtypes.SignersToEthArrays(signers)
	tx, err := c.CbrContract.Withdraw(c.Auth, wdDetail.WdOnchain, wdDetail.GetSortedSigsBytes(), addrs, powers)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "OnchainCbrWithdraw")
	return nil
}

func (c *CbrChain) SetInitSigners(amts []*big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	var addrs []eth.Addr
	for i := range amts {
		addrs = append(addrs, ValSignerAddrs[i])
	}
	tx, err := c.CbrContract.ResetSigners(c.Auth, addrs, amts)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "SetInitSigners")
	return nil
}

func OnchainClaimFarmingRewards(details *farmingtypes.RewardClaimDetails) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	var sigs [][]byte
	for _, signature := range details.Signatures {
		sigs = append(sigs, signature.SigBytes)
	}
	tx, err := Contracts.FarmingRewards.ClaimRewards(EtherBaseAuth, details.RewardProtoBytes, sigs, nil, nil)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "OnchainClaimFarmingRewards")
	return nil
}

func OnchainClaimStakingReward(claimInfo *distrtypes.StakingRewardClaimInfo) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	var sigs [][]byte
	for _, signature := range claimInfo.Signatures {
		sigs = append(sigs, signature.SigBytes)
	}
	tx, err := Contracts.StakingReward.ClaimReward(EtherBaseAuth, claimInfo.RewardProtoBytes, sigs)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "OnchainClaimStakingReward")
	return nil
}

func (c *CbrChain) PbrDeposit(fromUid uint64, amt *big.Int, mintChainId uint64, nonce uint64) (string, error) {
	tx, err := c.PegVaultContract.Deposit(c.Users[fromUid].Auth, c.UNIAddr, amt, mintChainId, c.Users[fromUid].Address, nonce)
	if err != nil {
		return "", err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return "", err
	}
	// last log is Deposit event (NOTE: test only)
	depositLog := receipt.Logs[len(receipt.Logs)-1]
	depositEv, err := c.PegVaultContract.ParseDeposited(*depositLog)
	if err != nil {
		return "", fmt.Errorf("parse log %+v err: %w", depositEv, err)
	}
	log.Infof("Deposit tx success, depositId: %x", depositEv.DepositId)
	return eth.Hash(depositEv.DepositId).Hex(), nil
}

func (c *CbrChain) PbrBurn(fromUid uint64, amt *big.Int, nonce uint64) (string, error) {
	tx, err := c.PegBridgeContract.Burn(c.Users[fromUid].Auth, c.PeggedUNIAddr, amt, c.Users[fromUid].Address, nonce)
	if err != nil {
		return "", err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return "", err
	}
	// last log is Deposit event (NOTE: test only)
	burnLog := receipt.Logs[len(receipt.Logs)-1]
	burnEv, err := c.PegBridgeContract.ParseBurn(*burnLog)
	if err != nil {
		return "", fmt.Errorf("parse log %+v err: %w", burnEv, err)
	}
	log.Infof("Burn tx success, burnId: %x", burnEv.BurnId)
	return eth.Hash(burnEv.BurnId).Hex(), nil
}

func (c *CbrChain) OnchainPegVaultWithdraw(info *pegbrtypes.WithdrawInfo, signers []*cbrtypes.Signer) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	addrs, powers := cbrtypes.SignersToEthArrays(signers)
	tx, err := c.PegVaultContract.Withdraw(c.Auth, info.WithdrawProtoBytes, info.GetSortedSigsBytes(), addrs, powers)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "OnchainPegVaultWithdraw")
	return nil
}

func (c *CbrChain) OnchainPegBridgeMint(info *pegbrtypes.MintInfo, signers []*cbrtypes.Signer) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	addrs, powers := cbrtypes.SignersToEthArrays(signers)
	tx, err := c.PegBridgeContract.Mint(c.Auth, info.MintProtoBytes, info.GetSortedSigsBytes(), addrs, powers)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "OnchainPegBridgeMint")
	return nil
}

func (c *CbrChain) CheckUSDTBalance(uid uint64, expectedAmt *big.Int) {
	var err error
	var expected bool
	balanceStr := ""
	for retry := 0; retry < RetryLimit*2; retry++ {
		balance, err := c.USDTContract.BalanceOf(&bind.CallOpts{}, c.Users[uid].Address)
		balanceStr = balance.String()
		if err == nil && balance.Cmp(expectedAmt) == 0 {
			expected = true
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to CheckUSDTBalance")
	if !expected {
		log.Fatalf("CheckUSDTBalance failed,now:%s, expect:%s", balanceStr, expectedAmt)
	}
}

func (c *CbrChain) CheckUNIBalance(uid uint64, expectedAmt *big.Int) {
	var err error
	var expected bool
	balanceStr := ""
	for retry := 0; retry < RetryLimit*2; retry++ {
		balance, err := c.UNIContract.BalanceOf(&bind.CallOpts{}, c.Users[uid].Address)
		balanceStr = balance.String()
		if err == nil && balance.Cmp(expectedAmt) == 0 {
			expected = true
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to CheckUNIBalance")
	if !expected {
		log.Fatalf("CheckUNIBalance failed,now:%s, expect:%s", balanceStr, expectedAmt)
	}
}

func (c *CbrChain) CheckPeggedUNIBalance(uid uint64, expectedAmt *big.Int) {
	var err error
	var expected bool
	balanceStr := ""
	for retry := 0; retry < RetryLimit*2; retry++ {
		balance, err := c.PeggedUNIContract.BalanceOf(&bind.CallOpts{}, c.Users[uid].Address)
		balanceStr = balance.String()
		if err == nil && balance.Cmp(expectedAmt) == 0 {
			expected = true
			break
		}
		time.Sleep(RetryPeriod)
	}
	ChkErr(err, "failed to CheckPeggedUNIBalance")
	if !expected {
		log.Fatalf("CheckPeggedUNIBalance failed, now:%s, expect:%s", balanceStr, expectedAmt)
	}
}

func (c *CbrChain) TransferMsg(uid uint64, receiver eth.Addr, dstChainId uint64, message []byte) error {
	tx, err := c.TransferMessageContract.TransferMessage(
		c.Users[uid].Auth,
		receiver,
		dstChainId,
		message,
	)
	if err != nil {
		return err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return err
	}
	// last log is MessageWithTransfer event (NOTE: test only)
	msgLog := receipt.Logs[len(receipt.Logs)-1]
	msgEv, err := c.MessageBusContract.ParseMessage(*msgLog)
	if err != nil {
		return fmt.Errorf("parse log %+v err: %w", msgEv, err)
	}
	log.Infof("SendMessage tx success, message: %x", msgEv.Message)
	return nil
}

func (c *CbrChain) BatchTransfer(
	uid uint64, receiver eth.Addr, token eth.Addr, amount *big.Int, dstChainId uint64, maxSlippage uint32,
	bridgeType uint8, accounts []eth.Addr, amounts []*big.Int) (xferId eth.Hash, err error) {

	tx, err := c.BatchTransferContract.BatchTransfer(
		c.Users[uid].Auth,
		receiver,
		token,
		amount,
		dstChainId,
		maxSlippage,
		bridgeType,
		accounts,
		amounts,
	)
	if err != nil {
		return eth.ZeroHash, err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return eth.ZeroHash, err
	}
	// last log is MessageWithTransfer event (NOTE: test only)
	msgLog := receipt.Logs[len(receipt.Logs)-1]
	msgEv, err := c.MessageBusContract.ParseMessageWithTransfer(*msgLog)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("parse log %+v err: %w", msgEv, err)
	}
	log.Infof("SendMessageWithTransfer tx success, srcTransferId: %x", msgEv.SrcTransferId)
	return msgEv.SrcTransferId, nil
}

func (c *CbrChain) BatchTransferForDeposit(
	uid uint64, receiver eth.Addr, token eth.Addr, amount *big.Int, dstChainId uint64, maxSlippage uint32,
	bridgeType uint8, accounts []eth.Addr, amounts []*big.Int) (xferId eth.Hash, err error) {

	tx, err := c.BatchTransferContract.BatchTransfer(
		c.Users[uid].Auth,
		receiver,
		token,
		amount,
		dstChainId,
		maxSlippage,
		bridgeType,
		accounts,
		amounts,
	)
	if err != nil {
		return eth.ZeroHash, err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return eth.ZeroHash, err
	}
	// last log is MessageWithTransfer event (NOTE: test only)
	msgLog := receipt.Logs[len(receipt.Logs)-1]
	msgEv, err := c.MessageBusContract.ParseMessageWithTransfer(*msgLog)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("parse log %+v err: %w", msgEv, err)
	}
	log.Infof("SendMessageWithTransfer tx success, srcTransferId: %x", msgEv.SrcTransferId)
	return msgEv.SrcTransferId, nil
}

func (c *CbrChain) ApproveUSDTForBatchTransfer(uid uint64, amt *big.Int) error {
	return c.ApproveErc20(c.USDTContract, uid, amt, c.BatchTransferAddr)
}

func (c *CbrChain) SendMessageWithTransfer(fromUid uint64, amt *big.Int, mintChainId uint64, nonce uint64) error {

	return nil
}
