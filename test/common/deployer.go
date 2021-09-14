package common

import (
	"context"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func DeployERC20Contract() (*types.Transaction, eth.Addr, *eth.Erc20) {
	initAmt := NewBigInt(1, 28)
	erc20Addr, tx, erc20, err := eth.DeployErc20(EtherBaseAuth, EthClient, "Celer", "CELR", initAmt, 18)
	ChkErr(err, "failed to deploy ERC20")

	log.Infoln("Erc20 address:", erc20Addr.String())
	return tx, erc20Addr, erc20
}

func DeployCelrContract() {
	var tx *types.Transaction
	tx, CelrAddr, CelrContract = DeployERC20Contract()
	WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "DeployERC20")
}

func DeploySgnStakingContracts(contractParams *ContractParams) *types.Transaction {
	Contracts = &eth.Contracts{}
	stakingContractAddr, _, staking, err := eth.DeployStaking(
		EtherBaseAuth,
		EthClient,
		contractParams.CelrAddr,
		contractParams.ProposalDeposit,
		contractParams.VotePeriod,
		contractParams.UnbondingPeriod,
		contractParams.MaxBondedValidators,
		contractParams.MinValidatorTokens,
		contractParams.MinSelfDelegation,
		contractParams.AdvanceNoticePeriod,
		contractParams.ValidatorBondInterval,
		contractParams.MaxSlashFactor)
	ChkErr(err, "failed to deploy Staking contract")
	Contracts.Staking, err = eth.NewStakingContract(stakingContractAddr, EthClient)
	ChkErr(err, "failed to set staking contract")

	sgnContractAddr, _, _, err := eth.DeploySGN(EtherBaseAuth, EthClient, stakingContractAddr)
	ChkErr(err, "failed to deploy sgn contract")
	Contracts.Sgn, err = eth.NewSgnContract(sgnContractAddr, EthClient)
	ChkErr(err, "failed to set sgn contract")

	rewardContractAddr, _, _, err := eth.DeployReward(
		EtherBaseAuth, EthClient, stakingContractAddr, contractParams.CelrAddr)
	ChkErr(err, "failed to deploy reward contract")
	Contracts.Reward, err = eth.NewRewardContract(rewardContractAddr, EthClient)
	ChkErr(err, "failed to set reward contract")

	viewerContractAddr, _, _, err := eth.DeployViewer(EtherBaseAuth, EthClient, stakingContractAddr)
	ChkErr(err, "failed to deploy viewer contract")
	Contracts.Viewer, err = eth.NewViewerContract(viewerContractAddr, EthClient)
	ChkErr(err, "failed to set viewer contract")

	governContractAddr, _, _, err := eth.DeployGovern(
		EtherBaseAuth, EthClient, stakingContractAddr, contractParams.CelrAddr, rewardContractAddr)
	ChkErr(err, "failed to deploy govern contract")
	Contracts.Govern, err = eth.NewGovernContract(governContractAddr, EthClient)
	ChkErr(err, "failed to set govern contract")

	EtherBaseAuth.GasLimit = 8000000
	_, err = staking.SetGovContract(EtherBaseAuth, governContractAddr)
	ChkErr(err, "failed to set gov contract")
	tx, err := staking.SetRewardContract(EtherBaseAuth, rewardContractAddr)
	ChkErr(err, "failed to set gov contract")
	EtherBaseAuth.GasLimit = 0

	log.Infoln("Staking address:", stakingContractAddr.String())
	log.Infoln("SGN address:", sgnContractAddr.String())
	log.Infoln("Reward address:", rewardContractAddr.String())
	log.Infoln("Viewer address:", viewerContractAddr.String())
	log.Infoln("Govern address:", governContractAddr.String())

	return tx
}

func DeployCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy contracts",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			configFileViper := viper.New()
			configDir := configFileViper.GetString(flags.FlagHome)
			configPath := filepath.Join(configDir, "config")
			configFileViper.SetConfigType("toml")
			configFileViper.SetConfigName("sgn")
			configFileViper.AddConfigPath(configPath)
			err = configFileViper.ReadInConfig()
			if err != nil {
				return err
			}
			ethurl := configFileViper.GetString(common.FlagEthGateway)
			var rpcClient *rpc.Client
			rpcClient, err = rpc.Dial(ethurl)
			if err != nil {
				return err
			}
			EthClient = ethclient.NewClient(rpcClient)

			var ksBytes []byte
			ksBytes, err = ioutil.ReadFile(configFileViper.GetString(common.FlagEthSignerKeystore))
			if err != nil {
				return err
			}
			EtherBaseAuth, err = bind.NewTransactorWithChainID(
				strings.NewReader(string(ksBytes)), configFileViper.GetString(common.FlagEthSignerPassphrase), big.NewInt(viper.GetInt64(common.FlagEthChainId)))
			if err != nil {
				return err
			}

			if ethurl == LocalGeth {
				SetEthBaseKs("./docker-volumes/geth-env")
				err = FundAddrsETH(
					[]eth.Addr{
						ValEthAddrs[0],
						ValSignerAddrs[0],
						ClientEthAddrs[0],
						ClientEthAddrs[1],
					}, NewBigInt(1, 20))
				ChkErr(err, "fund ETH to validator and clients")
			}

			_, celrAddr, erc20 := DeployERC20Contract()
			// NOTE: values below are for local tests
			contractParams := &ContractParams{
				CelrAddr:              celrAddr,
				ProposalDeposit:       big.NewInt(1e18), // 1 CELR
				VotePeriod:            big.NewInt(90),
				UnbondingPeriod:       big.NewInt(15),
				MaxBondedValidators:   big.NewInt(5),
				MinValidatorTokens:    big.NewInt(1e18),
				MinSelfDelegation:     big.NewInt(1e18),
				AdvanceNoticePeriod:   big.NewInt(30),
				ValidatorBondInterval: big.NewInt(0),
				MaxSlashFactor:        big.NewInt(1e5),
			}
			tx := DeploySgnStakingContracts(contractParams)
			WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "DeployStakingContracts")

			configFileViper.Set(common.FlagEthContractCelr, celrAddr.Hex())
			configFileViper.Set(common.FlagEthContractStaking, Contracts.Staking.Address.Hex())
			configFileViper.Set(common.FlagEthContractSgn, Contracts.Sgn.Address.Hex())
			configFileViper.Set(common.FlagEthContractReward, Contracts.Reward.Address.Hex())
			configFileViper.Set(common.FlagEthContractViewer, Contracts.Viewer.Address.Hex())
			configFileViper.Set(common.FlagEthContractGovern, Contracts.Govern.Address.Hex())
			err = configFileViper.WriteConfig()
			ChkErr(err, "failed to write config")

			if ethurl == LocalGeth {
				amt := NewBigInt(1, 20)
				tx, err := erc20.Approve(EtherBaseAuth, Contracts.Reward.Address, amt)
				ChkErr(err, "failed to approve erc20")
				WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "approve erc20")
				_, err = Contracts.Reward.ContributeToRewardPool(EtherBaseAuth, amt)
				ChkErr(err, "failed to call ContributeToMiningPool of Staking contract")
				err = FundAddrsErc20(celrAddr,
					[]eth.Addr{
						ClientEthAddrs[0],
						ClientEthAddrs[1],
					},
					NewBigInt(1, 20),
				)
				ChkErr(err, "fund test CELR to clients")
			}

			return nil
		},
	}

	return cmd
}
