package eth

import (
	"encoding/hex"
	"io/ioutil"
	"math/big"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)

type EthClient struct {
	// init by NewEthClient
	Client     *ethclient.Client
	Transactor *ethutils.Transactor
	Signer     ethutils.Signer
	Address    Addr
	// init by SetContracts
	Contracts *Contracts
}

type Contracts struct {
	Staking *StakingContract
	Sgn     *SgnContract
	Reward  *RewardContract
	Viewer  *ViewerContract
	Govern  *GovernContract
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
	stakingContract, sgnContract, rewardContract, viewerContract, governContract string) (*EthClient, error) {
	ethClient := &EthClient{
		Contracts: &Contracts{},
	}

	rpcClient, err := ethrpc.Dial(ethurl)
	if err != nil {
		return nil, err
	}

	ethClient.Client = ethclient.NewClient(rpcClient)
	err = ethClient.setContracts(stakingContract, sgnContract, rewardContract, viewerContract, governContract)
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
	ethClient.Signer, err = ethutils.NewSigner(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), tconfig.ChainId)
	if err != nil {
		return err
	}

	ethClient.Transactor, err = ethutils.NewTransactor(
		string(ksBytes),
		passphrase,
		ethClient.Client,
		tconfig.ChainId,
		ethutils.WithBlockDelay(tconfig.BlockDelay),
		ethutils.WithPollingInterval(time.Duration(tconfig.BlockPollingInterval)*time.Second),
		ethutils.WithAddGasGwei(tconfig.AddGasPriceGwei),
		ethutils.WithMinGasGwei(tconfig.MinGasPriceGwei),
	)

	return err
}

func (ethClient *EthClient) setContracts(
	stakingContract, sgnContract, rewardContract, viewerContract, governContract string) error {
	var err error
	ethClient.Contracts.Staking, err = NewStakingContract(Hex2Addr(stakingContract), ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Contracts.Sgn, err = NewSgnContract(Hex2Addr(sgnContract), ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Contracts.Reward, err = NewRewardContract(Hex2Addr(rewardContract), ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Contracts.Viewer, err = NewViewerContract(Hex2Addr(viewerContract), ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Contracts.Govern, err = NewGovernContract(Hex2Addr(governContract), ethClient.Client)
	if err != nil {
		return err
	}

	return nil
}

func (ethClient *EthClient) SignEthMessage(data []byte) ([]byte, error) {
	return ethClient.Signer.SignEthMessage(data)
}
