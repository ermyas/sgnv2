package common

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"sort"
	"strings"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogo/protobuf/proto"
)

func SetEthBaseKs(prefix string) {
	etherBaseKs = prefix + "/keystore/etherbase.json"
}

// SetupEthClients sets Client part (Client) and Auth part (PrivateKey, Address, Auth)
// Contracts part is set after deploying Staking and SGN contracts in SetupNewSGNEnv()
func SetupEthClients() {
	rpcClient, err := rpc.Dial(LocalGeth)
	if err != nil {
		log.Fatal(err)
	}
	EthClient = ethclient.NewClient(rpcClient)

	_, EtherBaseAuth, err = GetAuth(etherBaseKs, int64(ChainID))
	if err != nil {
		log.Fatal(err)
	}

	var auth *bind.TransactOpts
	for i := 0; i < len(ValEthKs); i++ {
		_, auth, err = GetAuth(ValEthKs[i], int64(ChainID))
		if err != nil {
			log.Fatal(err)
		}
		ValAuths = append(ValAuths, auth)
	}
	for i := 0; i < len(ValSignerKs); i++ {
		_, auth, err = GetAuth(ValSignerKs[i], int64(ChainID))
		if err != nil {
			log.Fatal(err)
		}
		SignerAuths = append(SignerAuths, auth)
	}
	for i := 0; i < len(DelEthKs); i++ {
		_, auth, err = GetAuth(DelEthKs[i], int64(ChainID))
		if err != nil {
			log.Fatal(err)
		}
		DelAuths = append(DelAuths, auth)
	}

	Client0, err = SetupTestEthClient(ClientEthKs[0])
	if err != nil {
		log.Fatal(err)
	}
	Client1, err = SetupTestEthClient(ClientEthKs[1])
	if err != nil {
		log.Fatal(err)
	}

	CbrClient1 = &CbrClient{
		Ec:   EthClient,
		Auth: EtherBaseAuth,
	}
}

func SetupEthClient2() {
	rpcClient, err := rpc.Dial(LocalGeth2)
	if err != nil {
		log.Fatal(err)
	}
	EthClient2 = ethclient.NewClient(rpcClient)
	_, EtherBaseAuth2, err = GetAuth(etherBaseKs, int64(Geth2ChainID))
	if err != nil {
		log.Fatal(err)
	}

	CbrClient2 = &CbrClient{
		Ec:   EthClient2,
		Auth: EtherBaseAuth2,
	}
}

func SetupTestEthClient(ksfile string) (*TestEthClient, error) {
	addr, auth, err := GetAuth(ksfile, int64(ChainID))
	if err != nil {
		return nil, err
	}
	testClient := &TestEthClient{
		Address: addr,
		Auth:    auth,
	}
	ksBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return nil, err
	}
	testClient.Signer, err = ethutils.NewSignerFromKeystore(string(ksBytes), "", big.NewInt(int64(ChainID)))
	if err != nil {
		return nil, err
	}
	return testClient, nil
}

func FundAddrsETH(recipients []eth.Addr, amount *big.Int, gatewayAddr string, chainId int64) error {
	conn, auth, ctx, senderAddr, connErr := prepareEtherBaseClient(gatewayAddr, chainId)
	if connErr != nil {
		return connErr
	}
	auth.Value = amount
	var gasLimit uint64 = 21000
	var lastTx *types.Transaction
	for _, addr := range recipients {
		nonce, err := conn.PendingNonceAt(ctx, senderAddr)
		if err != nil {
			return err
		}
		gasPrice, err := conn.SuggestGasPrice(ctx)
		if err != nil {
			return err
		}
		txData := &types.DynamicFeeTx{
			Nonce:     nonce,
			GasTipCap: big.NewInt(0),
			GasFeeCap: gasPrice,
			Gas:       gasLimit,
			To:        &addr,
			Value:     auth.Value,
			Data:      nil,
		}
		tx := types.NewTx(txData)
		tx, err = auth.Signer(senderAddr, tx)
		if err != nil {
			return err
		}
		if addr == eth.ZeroAddr {
			log.Info("Advancing block")
		} else {
			log.Infof("Sending ETH %s to %x from %x", amount, addr, senderAddr)
		}

		err = conn.SendTransaction(ctx, tx)
		if err != nil {
			return err
		}
		lastTx = tx
	}
	ctx2, cancel := context.WithTimeout(ctx, waitMinedTimeout)
	defer cancel()
	receipt, err := ethutils.WaitMined(ctx2, conn, lastTx, ethutils.WithBlockDelay(BlockDelay), ethutils.WithPollingInterval(PollingInterval))
	if err != nil {
		log.Error(err)
	}
	if receipt.Status != 1 {
		log.Errorf("last tx failed. tx hash: %x", receipt.TxHash)
	} else {
		for _, addr := range recipients {
			if addr == eth.ZeroAddr {
				head, _ := conn.HeaderByNumber(ctx, nil)
				log.Infoln("Current block number:", head.Number.String())
			} else {
				bal, _ := conn.BalanceAt(ctx, addr, nil)
				log.Infoln("Funded.", addr.String(), "bal:", bal.String())
			}
		}
	}
	return nil
}

func FundAddrsErc20(erc20Addr eth.Addr, recipients []eth.Addr, amount *big.Int, ethClient *ethclient.Client, auth *bind.TransactOpts) error {
	erc20Contract, err := eth.NewErc20(erc20Addr, ethClient)
	if err != nil {
		return err
	}
	var lastTx *types.Transaction
	for _, addr := range recipients {
		tx, transferErr := erc20Contract.Transfer(auth, addr, amount)
		if transferErr != nil {
			return transferErr
		}
		lastTx = tx
		log.Infof("Sending ERC20 %s to %x from %x", amount, addr, auth.From)
	}
	_, err = ethutils.WaitMined(context.Background(), ethClient, lastTx, ethutils.WithBlockDelay(BlockDelay), ethutils.WithPollingInterval(PollingInterval))
	return err
}

func InitializeValidator(auth *bind.TransactOpts, signerAddr eth.Addr, sgnAddr sdk.AccAddress, minSelfDelegation *big.Int, commissionRate uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	log.Infof("%x calls staking contract to initialize validator minSelfDelegation: %s, commissionRate: %d",
		auth.From, minSelfDelegation, commissionRate)

	tx, err := CelrContract.Approve(auth, Contracts.Staking.Address, minSelfDelegation)
	if err != nil {
		log.Error(err)
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "Approve")

	tx, err = Contracts.Staking.InitializeValidator(auth, signerAddr, minSelfDelegation, commissionRate)
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "InitializeValidator")

	log.Infof("%x calls sgn contract to update sgnAddr %s", auth.From, sgnAddr)
	tx, err = Contracts.Sgn.UpdateSgnAddr(auth, sgnAddr.Bytes())
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "UpdateSgnAddr")
	return nil
}

func Delegate(auth *bind.TransactOpts, valAddr eth.Addr, amt *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	log.Infof("%x calls staking contract to delegate to the validator %x...", auth.From, valAddr)
	tx, err := CelrContract.Approve(auth, Contracts.Staking.Address, amt)
	if err != nil {
		log.Error(err)
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "Approve")

	auth.GasLimit = 8000000
	tx, err = Contracts.Staking.Delegate(auth, valAddr, amt)
	auth.GasLimit = 0
	if err != nil {
		log.Error(err)
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "Delegate")
	return nil
}

func Undelegate(auth *bind.TransactOpts, valAddr eth.Addr, amt *big.Int) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	tx, err := Contracts.Staking.Undelegate(auth, valAddr, amt)
	if err != nil {
		log.Error(err)
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "Undelegate")
	return nil
}

func ConfirmUnbondedValidator(auth *bind.TransactOpts, valAddr eth.Addr) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	tx, err := Contracts.Staking.ConfirmUnbondedValidator(auth, valAddr)
	if err != nil {
		log.Error(err)
		return err
	}
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "ConfirmUnbondedValidator")
	return nil
}

func prepareEtherBaseClient(gatewayAddr string, chainId int64) (
	*ethclient.Client, *bind.TransactOpts, context.Context, eth.Addr, error) {
	conn, err := ethclient.Dial(gatewayAddr)
	if err != nil {
		return nil, nil, nil, eth.Addr{}, err
	}
	log.Infoln("EtherBaseKs: ", etherBaseKs)
	etherBaseKsBytes, err := ioutil.ReadFile(etherBaseKs)
	if err != nil {
		return nil, nil, nil, eth.Addr{}, err
	}
	etherBaseAddrStr, err := eth.GetAddressFromKeystore(etherBaseKsBytes)
	if err != nil {
		return nil, nil, nil, eth.Addr{}, err
	}
	etherBaseAddr := eth.Hex2Addr(etherBaseAddrStr)
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(etherBaseKsBytes)), "", big.NewInt(int64(chainId))) // Private Mainchain Testnet
	if err != nil {
		return nil, nil, nil, eth.Addr{}, err
	}
	return conn, auth, context.Background(), etherBaseAddr, nil
}

// call usdt contract approve for cbr addr
func (c *CbrClient) Approve(amt *big.Int) error {
	tx, err := c.USDTContract.Approve(c.Auth, c.CbrAddr, amt)
	if err != nil {
		return err
	}
	_, err = ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	return err
}

func (c *CbrClient) AddLiq(amt *big.Int) error {
	tx, err := c.CbrContract.AddLiquidity(c.Auth, c.USDTAddr, amt)
	if err != nil {
		return err
	}
	_, err = ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	return err
}

func (c *CbrClient) Send(amt *big.Int, receiver eth.Addr, dstChainId, nonce uint64) ([32]byte, error) {
	tx, err := c.CbrContract.Send(c.Auth, receiver, c.USDTAddr, amt, dstChainId, nonce, 10000) //1% slippage
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

func (c *CbrClient) SetInitSigners(amts []*big.Int) error {
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
	_, err = ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	return err
}

func (c *CbrClient) OnchainWithdraw(wdDetail *cbrtypes.WithdrawDetail, curss []byte) error {
	tx, err := c.CbrContract.Withdraw(c.Auth, wdDetail.WdOnchain, curss, wdDetail.GetSortedSigsBytes())
	if err != nil {
		return err
	}
	_, err = ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	return err
}

func (c *CbrClient) OnchainClaimRewards(details *farmingtypes.RewardClaimDetails) error {
	sort.Slice(details.Signatures, func(i int, j int) bool {
		return details.Signatures[i].Signer < details.Signatures[j].Signer
	})
	var sigs [][]byte
	for _, signature := range details.Signatures {
		sigs = append(sigs, signature.SigBytes)
	}
	tx, err := c.FarmingRewardsContract.ClaimRewards(c.Auth, details.RewardProtoBytes, nil, sigs)
	if err != nil {
		return err
	}
	_, err = ethutils.WaitMined(context.Background(), c.Ec, tx, ethutils.WithPollingInterval(time.Second))
	return err
}
