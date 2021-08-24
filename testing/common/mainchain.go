package common

import (
	"context"
	"io/ioutil"
	"math/big"
	"strings"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
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

	EthClient       *ethclient.Client
	EtherBaseAuth   *bind.TransactOpts
	StakingContract *eth.Staking
	SgnContract     *eth.SGN

	Client0 *TestEthClient
	Client1 *TestEthClient
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
	testClient.Signer, err = ethutils.NewSignerFromKeystore(string(ksBytes), "", nil)
	return testClient, nil
}

func SetContracts(stakingContractAddr, sgnContractAddr eth.Addr) error {
	log.Infof("set contracts staking %x sgn %x", stakingContractAddr, sgnContractAddr)
	var err error
	StakingContract, err = eth.NewStaking(stakingContractAddr, EthClient)
	if err != nil {
		return err
	}
	SgnContract, err = eth.NewSGN(sgnContractAddr, EthClient)
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

func FundAddrsETH(amt string, recipients []eth.Addr) error {
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

func DelegateStake(fromAuth *bind.TransactOpts, toEthAddress eth.Addr, amt *big.Int) error {
	conn := EthClient
	stakingContract := StakingContract
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	log.Info("Call delegate on staking contract to delegate stake to the validator eth address...")
	_, err := E2eProfile.CelrContract.Approve(fromAuth, E2eProfile.StakingContractAddr, amt)
	if err != nil {
		return err
	}

	fromAuth.GasLimit = 8000000
	tx, err := stakingContract.Delegate(fromAuth, toEthAddress, amt)
	fromAuth.GasLimit = 0
	if err != nil {
		return err
	}
	WaitMinedWithChk(ctx, conn, tx, BlockDelay, PollingInterval, "Delegate to validator")
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
	auth, err := bind.NewTransactor(strings.NewReader(string(etherBaseKsBytes)), "")
	if err != nil {
		return nil, nil, nil, eth.Addr{}, err
	}
	return conn, auth, context.Background(), etherBaseAddr, nil
}
