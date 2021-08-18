package common

import (
	"context"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type TestEthClient struct {
	Address contracts.Addr
	Auth    *bind.TransactOpts
	Signer  eth.Signer
}

var (
	etherBaseKs = EnvDir + "/keystore/etherbase.json"

	EthClient     *ethclient.Client
	EtherBaseAuth *bind.TransactOpts
	DposContract  *contracts.DPoS
	SgnContract   *contracts.SGN

	Client0 *TestEthClient
	Client1 *TestEthClient
)

func SetEthBaseKs(prefix string) {
	etherBaseKs = prefix + "/keystore/etherbase.json"
}

// SetupEthClients sets Client part (Client) and Auth part (PrivateKey, Address, Auth)
// Contracts part (DPoSAddress, DPoS, SGNAddress, SGN, LedgerAddress, Ledger) is set after deploying DPoS and SGN contracts in SetupNewSGNEnv()
func SetupEthClients() {
	rpcClient, err := rpc.Dial(LocalGeth)
	if err != nil {
		log.Fatal(err)
	}
	EthClient = ethclient.NewClient(rpcClient)

	_, EtherBaseAuth, err = GetAuth(etherBaseKs)
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
	ksBytes, err := ioutil.ReadFile(ksfile)
	testClient.Signer, err = eth.NewSignerFromKeystore(string(ksBytes), "", nil)
	return testClient, nil
}

func SetContracts(dposAddr, sgnAddr contracts.Addr) error {
	log.Infof("set contracts dpos %x sgn %x", dposAddr, sgnAddr)
	var err error
	DposContract, err = contracts.NewDPoS(dposAddr, EthClient)
	if err != nil {
		return err
	}
	SgnContract, err = contracts.NewSGN(sgnAddr, EthClient)
	if err != nil {
		return err
	}
	return nil
}

func SetupE2eProfile() {
	// Deploy sample ERC20 contract (CELR)
	tx, erc20Addr, erc20 := DeployERC20Contract()
	WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "DeployERC20")

	E2eProfile = &TestProfile{
		// deployed addresses
		CelrAddr:     erc20Addr,
		CelrContract: erc20,
	}
}

func FundAddrsETH(amt string, recipients []contracts.Addr) error {
	conn, auth, ctx, senderAddr, connErr := prepareEtherBaseClient()
	if connErr != nil {
		return connErr
	}
	value := big.NewInt(0)
	value.SetString(amt, 10)
	auth.Value = value
	// chainID := big.NewInt(883) // Private Mainchain Testnet
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
		tx := types.NewTransaction(nonce, addr, auth.Value, gasLimit, gasPrice, nil)
		// TODO: tx, err = auth.Signer(types.NewEIP155Signer(chainID), senderAddr, tx)
		if err != nil {
			return err
		}
		if addr == contracts.ZeroAddr {
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
	receipt, err := eth.WaitMined(ctx2, conn, lastTx, eth.WithBlockDelay(BlockDelay), eth.WithPollingInterval(PollingInterval))
	if err != nil {
		log.Error(err)
	}
	if receipt.Status != 1 {
		log.Errorf("last tx failed. tx hash: %x", receipt.TxHash)
	} else {
		for _, addr := range recipients {
			if addr == contracts.ZeroAddr {
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

func FundAddrsErc20(erc20Addr contracts.Addr, addrs []contracts.Addr, amount string) error {
	erc20Contract, err := contracts.NewErc20(erc20Addr, EthClient)
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
	_, err = eth.WaitMined(context.Background(), EthClient, lastTx, eth.WithBlockDelay(BlockDelay), eth.WithPollingInterval(PollingInterval))
	return err
}

func DelegateStake(fromAuth *bind.TransactOpts, toEthAddress contracts.Addr, amt *big.Int) error {
	conn := EthClient
	dposContract := DposContract
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	log.Info("Call delegate on dpos contract to delegate stake to the validator eth address...")
	_, err := E2eProfile.CelrContract.Approve(fromAuth, E2eProfile.DPoSAddr, amt)
	if err != nil {
		return err
	}

	fromAuth.GasLimit = 8000000
	tx, err := dposContract.Delegate(fromAuth, toEthAddress, amt)
	fromAuth.GasLimit = 0
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, conn, tx, BlockDelay, PollingInterval, "Delegate to validator")
	return nil
}

func prepareEtherBaseClient() (
	*ethclient.Client, *bind.TransactOpts, context.Context, contracts.Addr, error) {
	conn, err := ethclient.Dial(LocalGeth)
	if err != nil {
		return nil, nil, nil, contracts.Addr{}, err
	}
	log.Infoln("EtherBaseKs: ", etherBaseKs)
	etherBaseKsBytes, err := ioutil.ReadFile(etherBaseKs)
	if err != nil {
		return nil, nil, nil, contracts.Addr{}, err
	}
	etherBaseAddrStr, err := contracts.GetAddressFromKeystore(etherBaseKsBytes)
	if err != nil {
		return nil, nil, nil, contracts.Addr{}, err
	}
	etherBaseAddr := contracts.Hex2Addr(etherBaseAddrStr)
	auth, err := bind.NewTransactor(strings.NewReader(string(etherBaseKsBytes)), "")
	if err != nil {
		return nil, nil, nil, contracts.Addr{}, err
	}
	return conn, auth, context.Background(), etherBaseAddr, nil
}
