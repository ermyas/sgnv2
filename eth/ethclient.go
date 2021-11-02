package eth

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
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
	Staking        *StakingContract
	Sgn            *SgnContract
	StakingReward  *StakingRewardContract
	FarmingRewards *FarmingRewardsContract
	Viewer         *ViewerContract
	Govern         *GovernContract
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
	stakingContract, sgnContract,
	stakingRewardContract, farmingRewardsContract,
	viewerContract, governContract string) (*EthClient, error) {
	ethClient := &EthClient{
		Contracts: &Contracts{},
	}

	rpcClient, err := ethrpc.Dial(ethurl)
	if err != nil {
		return nil, err
	}

	ethClient.Client = ethclient.NewClient(rpcClient)
	err = ethClient.setContracts(
		stakingContract, sgnContract,
		stakingRewardContract, farmingRewardsContract,
		viewerContract, governContract)
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
	var err error
	ethClient.Signer, ethClient.Address, err = CreateSigner(ksfile, passphrase, tconfig.ChainId)
	if err != nil {
		return err
	}
	ethClient.Transactor = ethutils.NewTransactorByExternalSigner(
		ethClient.Address,
		ethClient.Signer,
		ethClient.Client,
		ethutils.WithBlockDelay(tconfig.BlockDelay),
		ethutils.WithPollingInterval(time.Duration(tconfig.BlockPollingInterval)*time.Second),
		ethutils.WithAddGasGwei(tconfig.AddGasPriceGwei),
		ethutils.WithMinGasGwei(tconfig.MinGasPriceGwei),
	)
	return nil
}

func (ethClient *EthClient) setContracts(
	stakingContract, sgnContract,
	stakingRewardContract, farmingRewardsContract,
	viewerContract, governContract string) error {
	var err error
	ethClient.Contracts.Staking, err = NewStakingContract(Hex2Addr(stakingContract), ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Contracts.Sgn, err = NewSgnContract(Hex2Addr(sgnContract), ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Contracts.StakingReward, err = NewStakingRewardContract(Hex2Addr(stakingRewardContract), ethClient.Client)
	if err != nil {
		return err
	}

	ethClient.Contracts.FarmingRewards, err = NewFarmingRewardsContract(Hex2Addr(farmingRewardsContract), ethClient.Client)
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
	sig, err := ethClient.Signer.SignEthMessage(data)
	if err != nil {
		return nil, err
	}
	if sig[64] <= 1 {
		// Use 27/28 for v to be compatible with openzeppelin ECDSA lib
		sig[64] = sig[64] + 27
	}
	return sig, nil
}

// if ksfile is like awskms:us-west-2:alias/mytestkey, use KmsSigner
// passphrase will be awsKey:awsSec or if empty, will use aws auto search env variable etc
// otherwise normal ks json file based signer
const awskmsPre = "awskms"

// return signer, address
func CreateSigner(ksfile, passphrase string, chainid *big.Int) (ethutils.Signer, Addr, error) {
	if strings.HasPrefix(ksfile, awskmsPre) {
		kmskeyinfo := strings.SplitN(ksfile, ":", 3)
		if len(kmskeyinfo) != 3 {
			return nil, ZeroAddr, fmt.Errorf("%s has wrong format", ksfile)
		}
		awskeysec := []string{"", ""}
		if passphrase != "" {
			awskeysec = strings.SplitN(passphrase, ":", 2)
			if len(awskeysec) != 2 {
				return nil, ZeroAddr, fmt.Errorf("%s has wrong format", passphrase)
			}
		}
		kmsSigner, err := ethutils.NewKmsSigner(kmskeyinfo[1], kmskeyinfo[2], awskeysec[0], awskeysec[1], chainid)
		if err != nil {
			return nil, ZeroAddr, err
		}
		return kmsSigner, kmsSigner.Addr, nil
	}
	ksBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return nil, ZeroAddr, err
	}
	key, err := keystore.DecryptKey(ksBytes, passphrase)
	if err != nil {
		return nil, ZeroAddr, err
	}
	signer, err := ethutils.NewSigner(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)), chainid)
	return signer, key.Address, err
}
