package common

import (
	"context"
	"io/ioutil"
	"math/big"
	"strings"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type TestEthClient struct {
	Address eth.Addr
	Auth    *bind.TransactOpts
	Signer  ethutils.Signer
}

var (
	etherBaseKs = EnvDir + "/keystore/etherbase.json"
	ChainID     = 883

	EthClient     *ethclient.Client
	EtherBaseAuth *bind.TransactOpts

	Client0 *TestEthClient
	Client1 *TestEthClient

	Contracts    *eth.Contracts
	CelrAddr     eth.Addr
	CelrContract *eth.Erc20
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

	_, EtherBaseAuth, _ = GetAuth(etherBaseKs)
	Client0, err = SetupTestEthClient(ClientEthKs[0])
	if err != nil {
		log.Fatal(err)
	}
	Client1, err = SetupTestEthClient(ClientEthKs[1])
	if err != nil {
		log.Fatal(err)
	}
}

func SetupTestEthClient(ksfile string) (*TestEthClient, error) {
	addr, auth, err := GetAuth(ksfile)
	if err != nil {
		return nil, err
	}
	testClient := &TestEthClient{
		Address: addr,
		Auth:    auth,
	}
	ksBytes, _ := ioutil.ReadFile(ksfile)
	testClient.Signer, _ = ethutils.NewSignerFromKeystore(string(ksBytes), "", nil)
	return testClient, nil
}

func FundAddrsETH(amt string, recipients []eth.Addr) error {
	conn, auth, ctx, senderAddr, connErr := prepareEtherBaseClient()
	if connErr != nil {
		return connErr
	}
	value := big.NewInt(0)
	value.SetString(amt, 10)
	auth.Value = value
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
			log.Infof("Sending ETH %s to %x from %x", amt, addr, senderAddr)
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

func FundAddrsErc20(erc20Addr eth.Addr, addrs []eth.Addr, amount string) error {
	erc20Contract, err := eth.NewErc20(erc20Addr, EthClient)
	if err != nil {
		return err
	}
	tokenAmt := new(big.Int)
	tokenAmt.SetString(amount, 10)
	var lastTx *types.Transaction
	for _, addr := range addrs {
		tx, transferErr := erc20Contract.Transfer(EtherBaseAuth, addr, tokenAmt)
		if transferErr != nil {
			return transferErr
		}
		lastTx = tx
		log.Infof("Sending ERC20 %s to %x from %x", amount, addr, EtherBaseAuth.From)
	}
	_, err = ethutils.WaitMined(context.Background(), EthClient, lastTx, ethutils.WithBlockDelay(BlockDelay), ethutils.WithPollingInterval(PollingInterval))
	return err
}

func InitializeValidator(auth *bind.TransactOpts, sgnAddr sdk.AccAddress, minSelfDelegation *big.Int, commissionRate uint64) error {
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

	tx, err = Contracts.Staking.InitializeValidator(auth, auth.From, minSelfDelegation, commissionRate)
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
	WaitMinedWithChk(ctx, EthClient, tx, BlockDelay, PollingInterval, "Delegate to validator")
	return nil
}

func prepareEtherBaseClient() (
	*ethclient.Client, *bind.TransactOpts, context.Context, eth.Addr, error) {
	conn, err := ethclient.Dial(LocalGeth)
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
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(string(etherBaseKsBytes)), "", big.NewInt(int64(ChainID))) // Private Mainchain Testnet
	if err != nil {
		return nil, nil, nil, eth.Addr{}, err
	}
	return conn, auth, context.Background(), etherBaseAddr, nil
}
