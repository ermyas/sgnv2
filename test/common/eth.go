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
	if err != nil {
		log.Fatal(err)
	}

	var auth *bind.TransactOpts
	for i := 0; i < len(ValEthKs); i++ {
		_, auth, err = GetAuth(ValEthKs[i])
		if err != nil {
			log.Fatal(err)
		}
		ValAuths = append(ValAuths, auth)
	}
	for i := 0; i < len(ValSignerKs); i++ {
		_, auth, err = GetAuth(ValSignerKs[i])
		if err != nil {
			log.Fatal(err)
		}
		SignerAuths = append(SignerAuths, auth)
	}
	for i := 0; i < len(DelEthKs); i++ {
		_, auth, err = GetAuth(DelEthKs[i])
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
	if err != nil {
		return nil, err
	}
	testClient.Signer, err = ethutils.NewSignerFromKeystore(string(ksBytes), "", big.NewInt(int64(ChainID)))
	if err != nil {
		return nil, err
	}
	return testClient, nil
}

func FundAddrsETH(recipients []eth.Addr, amount *big.Int) error {
	conn, auth, ctx, senderAddr, connErr := prepareEtherBaseClient()
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

func FundAddrsErc20(erc20Addr eth.Addr, recipients []eth.Addr, amount *big.Int) error {
	erc20Contract, err := eth.NewErc20(erc20Addr, EthClient)
	if err != nil {
		return err
	}
	var lastTx *types.Transaction
	for _, addr := range recipients {
		tx, transferErr := erc20Contract.Transfer(EtherBaseAuth, addr, amount)
		if transferErr != nil {
			return transferErr
		}
		lastTx = tx
		log.Infof("Sending ERC20 %s to %x from %x", amount, addr, EtherBaseAuth.From)
	}
	_, err = ethutils.WaitMined(context.Background(), EthClient, lastTx, ethutils.WithBlockDelay(BlockDelay), ethutils.WithPollingInterval(PollingInterval))
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
