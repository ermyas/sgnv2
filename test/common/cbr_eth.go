package common

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/transactor"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	pegbrtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func InitCbrChainConfigs() {
	for i := 1; i < 10; i++ {
		if EtherBaseAuth != nil {
			break
		}
		Sleep(2)
	}
	CbrChain1 = &CbrChain{
		ChainId:    Geth1ChainID,
		Ec:         EthClient,
		Auth:       EtherBaseAuth,
		Transactor: GetEtherBaseTransactor(Geth1ChainID),
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
		ChainId:    Geth2ChainID,
		Ec:         ethclient.NewClient(rpcClient),
		Auth:       etherBaseAuth,
		Transactor: GetEtherBaseTransactor(Geth2ChainID),
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

func (c *CbrChain) ApproveBridgeTestToken(token *eth.BridgeTestToken, uid uint64, amt *big.Int, spender eth.Addr) error {
	receipt, err := c.Users[uid].Transactor.TransactWaitMined(
		"ApproveBridgeTestToken",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return token.Approve(opts, spender, amt)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

func (c *CbrChain) ApproveUSDT(uid uint64, amt *big.Int) error {
	return c.ApproveBridgeTestToken(c.USDTContract, uid, amt, c.CbrAddr)
}

func (c *CbrChain) ApproveUNI(uid uint64, amt *big.Int) error {
	return c.ApproveBridgeTestToken(c.UNIContract, uid, amt, c.PegVaultAddr)
}

func (c *CbrChain) ApproveUSDTForContractAsLP(uid uint64, amt *big.Int) error {
	return c.ApproveBridgeTestToken(c.USDTContract, uid, amt, c.CLPAddr)
}

func (c *CbrChain) ApproveUNIForBatchTransfer(uid uint64, amt *big.Int) error {
	return c.ApproveBridgeTestToken(c.UNIContract, uid, amt, c.BatchTransferAddr)
}

func (c *CbrChain) ApprovePeggedUNI(uid uint64, amt *big.Int) error {
	return c.ApproveBridgeTestToken(c.UNIContract, uid, amt, c.PegBridgeAddr)
}

func (c *CbrChain) ApprovePeggedUNIForBatchTransfer(uid uint64, amt *big.Int) error {
	return c.ApproveBridgeTestToken(c.UNIContract, uid, amt, c.BatchTransferAddr)
}

func (c *CbrChain) AddLiq(uid uint64, amt *big.Int) error {
	receipt, err := c.Users[uid].Transactor.TransactWaitMined(
		"AddLiq",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.CbrContract.AddLiquidity(opts, c.USDTAddr, amt)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

func (c *CbrChain) DepositToContractAsLP(uid uint64, amt *big.Int) error {
	receipt, err := c.Users[uid].Transactor.TransactWaitMined(
		"DepositToContractAsLP",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.CLPContract.Deposit(opts, c.USDTAddr, amt)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

func (c *CbrChain) AddLiqByContractAsLP(uid uint64, amt *big.Int) error {
	receipt, err := c.Users[uid].Transactor.TransactWaitMined(
		"AddLiqByContractAsLP",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.CLPContract.AddLiquidity(opts, c.USDTAddr, amt)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

func (c *CbrChain) SendWithdrawRequest(uid, wdSeq, toUid, toChain uint64, fromChains []uint64, tokens []eth.Addr, ratios, slippages []uint32) error {
	receipt, err := c.Users[uid].Transactor.TransactWaitMined(
		"SendWithdrawRequest",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.CLPContract.Withdraw(opts, wdSeq, c.Users[toUid].Address, toChain, fromChains, tokens, ratios, slippages)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

// only used for test
func (c *CbrChain) Send(uid uint64, amt *big.Int, dstChainId, nonce uint64) (eth.Hash, error) {
	return c.SendAny(uid, uid, amt, dstChainId, nonce, 100000) // 10% slippage
}

func (c *CbrChain) SendAny(fromUid, toUid uint64, amt *big.Int, dstChainId, nonce uint64, maxSlippage uint32) (eth.Hash, error) {
	receipt, err := c.Users[fromUid].Transactor.TransactWaitMined(
		"SendAny",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.CbrContract.Send(opts, c.Users[toUid].Address, c.USDTAddr, amt, dstChainId, nonce, maxSlippage)
		},
	)
	if err != nil {
		return eth.ZeroHash, err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return eth.ZeroHash, fmt.Errorf("tx failed")
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return eth.ZeroHash, errors.New("send tx failed")
	}
	sendLog := receipt.Logs[len(receipt.Logs)-1] // last log is Send event (NOTE Polygon breaks this assumption)
	sendEv, err := c.CbrContract.ParseSend(*sendLog)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("parse log %+v err: %w", sendLog, err)
	}
	return sendEv.TransferId, nil
}

func (c *CbrChain) OnchainCbrWithdraw(wdDetail *cbrtypes.WithdrawDetail, signers []*cbrtypes.Signer) error {
	addrs, powers := cbrtypes.SignersToEthArrays(signers)
	receipt, err := c.Transactor.TransactWaitMined(
		"OnchainCbrWithdraw",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.CbrContract.Withdraw(opts, wdDetail.WdOnchain, wdDetail.GetSortedSigsBytes(), addrs, powers)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

func (c *CbrChain) SetInitSigners(amts []*big.Int) error {
	var addrs []eth.Addr
	for i := range amts {
		addrs = append(addrs, ValSignerAddrs[i])
	}
	receipt, err := c.Transactor.TransactWaitMined(
		"TransactWaitMined",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.CbrContract.ResetSigners(opts, addrs, amts)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

func OnchainClaimFarmingRewards(details *farmingtypes.RewardClaimDetails) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	var sigs [][]byte
	for _, signature := range details.Signatures {
		sigs = append(sigs, signature.SigBytes)
	}
	// use valAuth[0] instead of etherbase to reduce nonce conflict chance
	tx, err := Contracts.FarmingRewards.ClaimRewards(ValAuths[0], details.RewardProtoBytes, sigs, nil, nil)
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
	// use valAuth[0] instead of etherbase to reduce nonce conflict chance
	tx, err := Contracts.StakingReward.ClaimReward(ValAuths[0], claimInfo.RewardProtoBytes, sigs)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "OnchainClaimStakingReward")
	return nil
}

func (c *CbrChain) PbrDeposit(fromUid uint64, amt *big.Int, mintChainId uint64, nonce uint64) (string, error) {
	receipt, err := c.Users[fromUid].Transactor.TransactWaitMined(
		"PbrDeposit",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.PegVaultContract.Deposit(opts, c.UNIAddr, amt, mintChainId, c.Users[fromUid].Address, nonce)
		},
	)
	if err != nil {
		return "", err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return "", fmt.Errorf("tx failed")
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return "", errors.New("deposit tx failed")
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
	receipt, err := c.Users[fromUid].Transactor.TransactWaitMined(
		"PbrBurn",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.PegBridgeContract.Burn(opts, c.UNIAddr, amt, c.Users[fromUid].Address, nonce)
		},
	)
	if err != nil {
		return "", err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return "", fmt.Errorf("tx failed")
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return "", errors.New("burn tx failed")
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

func WithdrawMsgFeesOnChain(txr *transactor.Transactor, claimInfo *msgtypes.FeeClaimInfo) error {
	for _, detail := range claimInfo.FeeClaimDetailsList {
		chain := GetChain(detail.ChainId)
		if chain == nil {
			log.Fatalf("chain not found for chainid (%d)", detail.ChainId)
		}

		log.Infoln("withdraw msg fee on-chain")
		curss, err := GetCurSortedSigners(txr, chain.ChainId)
		ChkErr(err, "unable to query chain signers")
		pass, sigsBytes := cbrtypes.ValidateSignatureQuorum(detail.Signatures, curss)
		if !pass {
			return fmt.Errorf("not enough sigs")
		}
		receipt, err := chain.Transactor.TransactWaitMined(
			"WithdrawMsgFeeOnChain",
			func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
				signers, powers := cbrtypes.SignersToEthArrays(curss)
				return chain.MessageBusContract.WithdrawFee(
					opts, eth.Hex2Addr(claimInfo.Recipient), detail.CumulativeFeeAmount.Amount.RoundInt().BigInt(), sigsBytes, signers, powers)
			},
		)
		if err != nil {
			return err
		}
		if receipt.Status != ethtypes.ReceiptStatusSuccessful {
			return fmt.Errorf("tx failed")
		}
	}
	return nil
}

func (c *CbrChain) OnchainPegVaultWithdraw(info *pegbrtypes.WithdrawInfo, signers []*cbrtypes.Signer) error {
	addrs, powers := cbrtypes.SignersToEthArrays(signers)
	receipt, err := c.Transactor.TransactWaitMined(
		"OnchainPegVaultWithdraw",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.PegVaultContract.Withdraw(opts, info.WithdrawProtoBytes, info.GetSortedSigsBytes(), addrs, powers)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
	return nil
}

func (c *CbrChain) OnchainPegBridgeMint(info *pegbrtypes.MintInfo, signers []*cbrtypes.Signer) error {
	addrs, powers := cbrtypes.SignersToEthArrays(signers)
	receipt, err := c.Transactor.TransactWaitMined(
		"OnchainPegBridgeMint",
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*ethtypes.Transaction, error) {
			return c.PegBridgeContract.Mint(opts, info.MintProtoBytes, info.GetSortedSigsBytes(), addrs, powers)
		},
	)
	if err != nil {
		return err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return fmt.Errorf("tx failed")
	}
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
		balance, err := c.UNIContract.BalanceOf(&bind.CallOpts{}, c.Users[uid].Address)
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
	auth := c.Users[uid].Auth
	auth.Value = MsgFeeBase
	tx, err := c.TransferMessageContract.TransferMessage(
		auth,
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
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return errors.New("transferMessage tx failed")
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

	auth := c.Users[uid].Auth
	auth.Value = MsgFeeBase
	tx, err := c.BatchTransferContract.BatchTransfer(
		auth,
		receiver,
		token,
		amount,
		dstChainId,
		maxSlippage,
		bridgeType,
		accounts,
		amounts,
	)
	auth.Value = big.NewInt(0)
	if err != nil {
		return eth.ZeroHash, err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return eth.ZeroHash, err
	}
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return eth.ZeroHash, errors.New("batchTransfer tx failed")
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
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		return eth.ZeroHash, errors.New("batchTransfer tx failed")
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
	return c.ApproveBridgeTestToken(c.USDTContract, uid, amt, c.BatchTransferAddr)
}

func (c *CbrChain) SendMessageWithTransfer(fromUid uint64, amt *big.Int, mintChainId uint64, nonce uint64) error {

	return nil
}
