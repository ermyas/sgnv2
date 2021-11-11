package common

import (
	"context"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func DeployERC20Contract(ethClient *ethclient.Client, auth *bind.TransactOpts, name, symbol string, decimal uint8) (eth.Addr, *eth.Erc20) {
	initAmt := NewBigInt(1, 28)
	erc20Addr, tx, erc20, err := eth.DeployErc20(auth, ethClient, name, symbol, initAmt, decimal)
	ChkErr(err, "failed to deploy ERC20")

	log.Infoln("Erc20 address:", erc20Addr.String())
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployERC20")

	return erc20Addr, erc20
}

func DeployBridgeContract(ethClient *ethclient.Client, auth *bind.TransactOpts) (cbrAddr eth.Addr, cbrContract *eth.BridgeContract) {
	cbrAddr, tx, _, err := eth.DeployBridge(auth, ethClient)
	ChkErr(err, "failed to deploy bridge contract")
	cbrContract, err = eth.NewBridgeContract(cbrAddr, ethClient)
	ChkErr(err, "failed to set bridge contract")

	log.Infoln("bridge address:", cbrAddr.String())

	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployBridgeContract")

	return
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

	stakingRewardContractAddr, _, _, err := eth.DeployStakingReward(
		EtherBaseAuth, EthClient, stakingContractAddr)
	ChkErr(err, "failed to deploy staking reward contract")
	Contracts.StakingReward, err = eth.NewStakingRewardContract(stakingRewardContractAddr, EthClient)
	ChkErr(err, "failed to set reward contract")

	farmingRewardsContractAddr, _, _, err := eth.DeployFarmingRewards(
		EtherBaseAuth, EthClient, stakingContractAddr)
	ChkErr(err, "failed to deploy farming rewards contract")
	Contracts.FarmingRewards, err = eth.NewFarmingRewardsContract(farmingRewardsContractAddr, EthClient)
	ChkErr(err, "failed to set reward contract")

	viewerContractAddr, _, _, err := eth.DeployViewer(EtherBaseAuth, EthClient, stakingContractAddr)
	ChkErr(err, "failed to deploy viewer contract")
	Contracts.Viewer, err = eth.NewViewerContract(viewerContractAddr, EthClient)
	ChkErr(err, "failed to set viewer contract")

	governContractAddr, _, _, err := eth.DeployGovern(
		EtherBaseAuth, EthClient, stakingContractAddr, contractParams.CelrAddr, stakingRewardContractAddr)
	ChkErr(err, "failed to deploy govern contract")
	Contracts.Govern, err = eth.NewGovernContract(governContractAddr, EthClient)
	ChkErr(err, "failed to set govern contract")

	EtherBaseAuth.GasLimit = 8000000
	_, err = staking.SetGovContract(EtherBaseAuth, governContractAddr)
	ChkErr(err, "failed to set gov contract")
	tx, err := staking.SetRewardContract(EtherBaseAuth, stakingRewardContractAddr)
	ChkErr(err, "failed to set gov contract")
	EtherBaseAuth.GasLimit = 0

	// Contribute to reward pools
	amt := NewBigInt(1, 25)
	celrContract, err := eth.NewErc20(contractParams.CelrAddr, EthClient)
	ChkErr(err, "failed to instantiate CELR contract")
	approveTx, err := celrContract.Approve(EtherBaseAuth, Contracts.StakingReward.Address, amt)
	ChkErr(err, "failed to approve CELR to StakingReward")
	WaitMinedWithChk(context.Background(), EthClient, approveTx, BlockDelay, PollingInterval, "approve CELR")
	allowance, _ := celrContract.Allowance(&bind.CallOpts{}, EtherBaseAuth.From, Contracts.StakingReward.Address)
	log.Infoln("allowance to StakingReward", allowance.String())
	_, err = Contracts.StakingReward.ContributeToRewardPool(EtherBaseAuth, amt)
	ChkErr(err, "failed to call StakingReward.ContributeToRewardPool")

	approveTx, err = celrContract.Approve(EtherBaseAuth, Contracts.FarmingRewards.Address, amt)
	ChkErr(err, "failed to approve CELR to FarmingRewards")
	WaitMinedWithChk(context.Background(), EthClient, approveTx, BlockDelay, PollingInterval, "approve CELR")
	allowance, _ = celrContract.Allowance(&bind.CallOpts{}, EtherBaseAuth.From, Contracts.FarmingRewards.Address)
	log.Infoln("CELR allowance to FarmingRewards", allowance.String())
	_, err = Contracts.FarmingRewards.ContributeToRewardPool(EtherBaseAuth, contractParams.CelrAddr, amt)
	ChkErr(err, "failed to contribute CELR to FarmingRewards")

	log.Infoln("Staking address:", stakingContractAddr.String())
	log.Infoln("SGN address:", sgnContractAddr.String())
	log.Infoln("StakingReward address:", stakingRewardContractAddr.String())
	log.Infoln("FarmingRewards address:", farmingRewardsContractAddr.String())
	log.Infoln("Viewer address:", viewerContractAddr.String())
	log.Infoln("Govern address:", governContractAddr.String())

	return tx
}

func DeployCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy contracts",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			ethurl := viper.GetString(common.FlagEthGateway)
			var rpcClient *rpc.Client
			rpcClient, err = rpc.Dial(ethurl)
			if err != nil {
				return err
			}
			EthClient = ethclient.NewClient(rpcClient)

			var ksBytes []byte
			ksBytes, err = ioutil.ReadFile(viper.GetString(common.FlagEthSignerKeystore))
			if err != nil {
				return err
			}
			EtherBaseAuth, err = bind.NewTransactorWithChainID(
				strings.NewReader(string(ksBytes)), viper.GetString(common.FlagEthSignerPassphrase), big.NewInt(viper.GetInt64(common.FlagEthChainId)))
			if err != nil {
				return err
			}

			if ethurl == LocalGeth {
				SetEthBaseKs("./docker-volumes/geth-env")
				err = FundAddrsETH(
					[]eth.Addr{
						ValEthAddrs[0],
						ValSignerAddrs[0],
						DelEthAddrs[0],
						ClientEthAddrs[0],
						ClientEthAddrs[1],
					}, NewBigInt(1, 20), LocalGeth, int64(ChainID))
				ChkErr(err, "fund test ETH")
			}

			celrAddr, erc20 := DeployERC20Contract(EthClient, EtherBaseAuth, "Celer", "CELR", 18)
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

			viper.Set(common.FlagEthContractCelr, celrAddr.Hex())
			viper.Set(common.FlagEthContractStaking, Contracts.Staking.Address.Hex())
			viper.Set(common.FlagEthContractSgn, Contracts.Sgn.Address.Hex())
			viper.Set(common.FlagEthContractStakingReward, Contracts.StakingReward.Address.Hex())
			viper.Set(common.FlagEthContractFarmingRewards, Contracts.FarmingRewards.Address.Hex())
			viper.Set(common.FlagEthContractViewer, Contracts.Viewer.Address.Hex())
			viper.Set(common.FlagEthContractGovern, Contracts.Govern.Address.Hex())
			err = viper.WriteConfig()
			ChkErr(err, "failed to write config")

			if ethurl == LocalGeth {
				amt := NewBigInt(1, 25)
				tx, err := erc20.Approve(EtherBaseAuth, Contracts.StakingReward.Address, amt)
				ChkErr(err, "failed to approve erc20 to StakingReward")
				WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "approve erc20")
				_, err = Contracts.StakingReward.ContributeToRewardPool(EtherBaseAuth, amt)
				ChkErr(err, "failed to call StakingReward.ContributeToRewardPool")

				tx, err = erc20.Approve(EtherBaseAuth, Contracts.FarmingRewards.Address, amt)
				ChkErr(err, "failed to approve erc20 to FarmingRewards")
				WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "approve erc20")
				_, err = Contracts.FarmingRewards.ContributeToRewardPool(EtherBaseAuth, celrAddr, amt)
				ChkErr(err, "failed to call FarmingRewards.ContributeToRewardPool")

				err = FundAddrsErc20(celrAddr,
					[]eth.Addr{
						ValEthAddrs[0],
						DelEthAddrs[0],
					},
					amt, EthClient, EtherBaseAuth,
				)
				ChkErr(err, "fund test CELR")
			}

			return nil
		},
	}

	return cmd
}
