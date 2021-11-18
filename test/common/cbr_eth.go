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
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func InitCbrChainConfigs() {
	CbrChain1 = &CbrChain{
		ChainId: ChainID,
		Ec:      EthClient,
		Auth:    EtherBaseAuth,
	}
	CbrChain1.SetUsersAndDelegators()

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
	CbrChain2.SetUsersAndDelegators()
}

func (c *CbrChain) SetUsersAndDelegators() {
	users := []*TestEthClient{}
	for _, clientKs := range ClientEthKs {
		u, err := SetupTestEthClient(clientKs, c.ChainId)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	c.Users = users
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

// only used for test
func (c *CbrChain) Send(uid uint64, amt *big.Int, dstChainId, nonce uint64) ([32]byte, error) {
	return c.SendAny(uid, uid, amt, dstChainId, nonce, 100000) //10% slippage
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
	sendLog := receipt.Logs[len(receipt.Logs)-1] // last log is Send event (NOTE Polygon breaks this assumption)
	sendEv, err := c.CbrContract.ParseSend(*sendLog)
	if err != nil {
		return eth.ZeroCid, fmt.Errorf("parse log %+v err: %w", sendLog, err)
	}
	return sendEv.TransferId, nil
}

func (c *CbrChain) OnchainWithdraw(wdDetail *cbrtypes.WithdrawDetail, signers []*cbrtypes.Signer) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	addrs, powers := cbrtypes.SignersToEthArrays(signers)
	tx, err := c.CbrContract.Withdraw(c.Auth, wdDetail.WdOnchain, wdDetail.GetSortedSigsBytes(), addrs, powers)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, c.Ec, tx, BlockDelay, PollingInterval, "OnchainWithdraw")
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
	sort.Slice(details.Signatures, func(i int, j int) bool {
		return details.Signatures[i].Signer < details.Signatures[j].Signer
	})
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
	sort.Slice(claimInfo.Signatures, func(i int, j int) bool {
		return claimInfo.Signatures[i].Signer < claimInfo.Signatures[j].Signer
	})
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
