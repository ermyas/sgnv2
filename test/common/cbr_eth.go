package common

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogo/protobuf/proto"
)

func InitCbrChainConfigs() {
	CbrChain1 = &CbrChain{
		ChainId: ChainID,
		Ec:      EthClient,
		Auth:    EtherBaseAuth,
	}
	CbrChain1.SetUsers()

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
	CbrChain2.SetUsers()
}

func (c *CbrChain) SetUsers() {
	u0, err := SetupTestEthClient(ClientEthKs[0], c.ChainId)
	if err != nil {
		log.Fatal(err)
	}
	u1, err := SetupTestEthClient(ClientEthKs[1], c.ChainId)
	if err != nil {
		log.Fatal(err)
	}
	u2, err := SetupTestEthClient(ClientEthKs[2], c.ChainId)
	if err != nil {
		log.Fatal(err)
	}
	c.Users = []*TestEthClient{u0, u1, u2}
}

func (c *CbrChain) Approve(uid uint64, amt *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	tx, err := c.USDTContract.Approve(c.Users[uid].Auth, c.CbrAddr, amt)
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

func (c *CbrChain) Send(uid uint64, amt *big.Int, dstChainId, nonce uint64) ([32]byte, error) {
	return c.SendAny(uid, uid, amt, dstChainId, nonce, 10000) //1% slippage
}

func (c *CbrChain) SendAny(fromUid, toUid uint64, amt *big.Int, dstChainId, nonce uint64, maxSlippage uint32) ([32]byte, error) {
	tx, err := c.CbrContract.Send(
		c.Users[fromUid].Auth, c.Users[toUid].Address, c.USDTAddr, amt, dstChainId, nonce, maxSlippage)
	if err != nil {
		return eth.ZeroCid, err
	}
	receipt, err := ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	if err != nil {
		return eth.ZeroCid, err
	}
	sendLog := receipt.Logs[len(receipt.Logs)-1] // last log is Send event
	sendEv, err := c.CbrContract.ParseSend(*sendLog)
	if err != nil {
		return eth.ZeroCid, fmt.Errorf("parse log %+v err: %w", sendLog, err)
	}
	return sendEv.TransferId, nil
}

func (c *CbrChain) OnchainWithdraw(wdDetail *cbrtypes.WithdrawDetail, curss []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	tx, err := c.CbrContract.Withdraw(c.Auth, wdDetail.WdOnchain, curss, wdDetail.GetSortedSigsBytes())
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "OnchainWithdraw")
	return nil
}

func (c *CbrChain) SetInitSigners(amts []*big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	var signers []*cbrtypes.AddrAmt
	for i, amt := range amts {
		signers = append(signers, &cbrtypes.AddrAmt{
			Addr: ValSignerAddrs[i].Bytes(),
			Amt:  amt.Bytes(),
		})
	}
	ss, err := proto.Marshal(&cbrtypes.SortedSigners{
		Signers: signers,
	})
	if err != nil {
		return err
	}
	tx, err := c.CbrContract.SetInitSigners(c.Auth, ss)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "SetInitSigners")
	return nil
}

func OnchainClaimRewards(details *farmingtypes.RewardClaimDetails) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	sort.Slice(details.Signatures, func(i int, j int) bool {
		return details.Signatures[i].Signer < details.Signatures[j].Signer
	})
	var sigs [][]byte
	for _, signature := range details.Signatures {
		sigs = append(sigs, signature.SigBytes)
	}
	tx, err := Contracts.FarmingRewards.ClaimRewards(EtherBaseAuth, details.RewardProtoBytes, nil, sigs)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "OnchainClaimRewards")
	return nil
}
