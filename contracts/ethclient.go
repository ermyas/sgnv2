package contracts

import (
	"encoding/hex"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/celer-network/goutils/eth"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type EthClient struct {
	// init by NewEthClient
	Client     *ethclient.Client
	Transactor *eth.Transactor
	Signer     eth.Signer
	Address    Addr

	// init by SetContracts
	StakingAddress Addr
	Staking        *Staking
	SGNAddress     Addr
	SGN            *SGN
}

type TransactorConfig struct {
	BlockDelay           uint64
	BlockPollingInterval uint64
	ChainId              *big.Int
	AddGasPriceGwei      uint64
	MinGasPriceGwei      uint64
}

func NewEthClient(
	ethurl string,
	ksfile string,
	passphrase string,
	tconfig *TransactorConfig,
	stakingAddrStr string,
	sgnAddrStr string) (*EthClient, error) {
	ethClient := &EthClient{}

	rpcClient, err := ethrpc.Dial(ethurl)
	if err != nil {
		return nil, err
	}

	ethClient.Client = ethclient.NewClient(rpcClient)
	err = ethClient.setContracts(stakingAddrStr, sgnAddrStr)
	if err != nil {
		return nil, err
	}

	if ksfile != "" {
		err = ethClient.setTransactor(ksfile, passphrase, tconfig)
		if err != nil {
			return nil, err
		}
	}

	return ethClient, nil
}

func (ethClient *EthClient) setTransactor(ksfile string, passphrase string, tconfig *TransactorConfig) error {
	ksBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return err
	}

	key, err := keystore.DecryptKey(ksBytes, passphrase)
	if err != nil {
		return err
	}

	ethClient.Address = key.Address
	ethClient.Signer, err = eth.NewSigner(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), tconfig.ChainId)
	if err != nil {
		return err
	}

	ethClient.Transactor, err = eth.NewTransactor(
		string(ksBytes),
		passphrase,
		ethClient.Client,
		tconfig.ChainId,
		eth.WithBlockDelay(tconfig.BlockDelay),
		eth.WithPollingInterval(time.Duration(tconfig.BlockPollingInterval)*time.Second),
		eth.WithAddGasGwei(tconfig.AddGasPriceGwei),
		eth.WithMinGasGwei(tconfig.MinGasPriceGwei),
	)

	return err
}

func (ethClient *EthClient) setContracts(stakingAddrStr, sgnAddrStr string) error {
	ethClient.StakingAddress = Hex2Addr(stakingAddrStr)
	staking, err := NewStaking(ethClient.StakingAddress, ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.SGNAddress = Hex2Addr(sgnAddrStr)
	sgn, err := NewSGN(ethClient.SGNAddress, ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Staking = staking
	ethClient.SGN = sgn
	return nil
}

func (ethClient *EthClient) SignEthMessage(data []byte) ([]byte, error) {
	return ethClient.Signer.SignEthMessage(data)
}
