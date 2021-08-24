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

func DeployERC20Contract() (*types.Transaction, eth.Addr, *eth.Erc20) {
	initAmt := new(big.Int)
	initAmt.SetString("1"+strings.Repeat("0", 28), 10)
	erc20Addr, tx, erc20, err := eth.DeployErc20(EtherBaseAuth, EthClient, "Celer", "CELR", initAmt, 18)
	ChkErr(err, "failed to deploy ERC20")

	log.Infoln("Erc20 address:", erc20Addr.String())
	return tx, erc20Addr, erc20
}

func DeployStakingSGNContracts(sgnParams *SGNParams) (*types.Transaction, eth.Addr, eth.Addr) {
	stakingAddr, _, _, err := eth.DeployStaking(
		EtherBaseAuth,
		EthClient,
		sgnParams.CelrAddr,
		sgnParams.GovernProposalDeposit,
		sgnParams.GovernVoteTimeout,
		sgnParams.SlashTimeout,
		sgnParams.MaxBondedValidators,
		sgnParams.MinValidatorTokens,
		sgnParams.MinSelfDelegation,
		sgnParams.AdvanceNoticePeriod,
		sgnParams.ValidatorBondInterval,
		sgnParams.MaxSlashFactor)

	ChkErr(err, "failed to deploy Staking contract")

	sgnAddr, _, _, err := eth.DeploySGN(EtherBaseAuth, EthClient, stakingAddr)
	ChkErr(err, "failed to deploy SGN contract")

	// TODO: register SGN address on Staking contract
	// staking, err := eth.NewStaking(stakingAddr, EthClient)
	// ChkErr(err, "failed to new Staking instance")
	// EtherBaseAuth.GasLimit = 8000000
	// tx, err := staking.RegisterSidechain(EtherBaseAuth, sgnAddr)
	// EtherBaseAuth.GasLimit = 0
	ChkErr(err, "failed to register SGN address on Staking contract")

	log.Infoln("Staking address:", stakingAddr.String())
	log.Infoln("SGN address:", sgnAddr.String())

	// TODO
	return nil, stakingAddr, sgnAddr
}

func DeployCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy contracts",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			configFileViper := viper.New()
			configFileViper.SetConfigFile(viper.GetString(common.FlagConfig))
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
			ksBytes, err = ioutil.ReadFile(configFileViper.GetString(common.FlagEthKeystore))
			if err != nil {
				return err
			}
			EtherBaseAuth, err = bind.NewTransactorWithChainID(
				strings.NewReader(string(ksBytes)), configFileViper.GetString(common.FlagEthPassphrase), big.NewInt(viper.GetInt64(common.FlagEthChainID)))
			if err != nil {
				return err
			}

			if ethurl == LocalGeth {
				SetEthBaseKs("./docker-volumes/geth-env")
				err = FundAddrsETH("1"+strings.Repeat("0", 20),
					[]eth.Addr{
						eth.Hex2Addr(ValEthAddrs[0]),
						eth.Hex2Addr(ClientEthAddrs[0]),
						eth.Hex2Addr(ClientEthAddrs[1]),
					})
				ChkErr(err, "fund ETH to validator and clients")
			}

			_, erc20Addr, erc20 := DeployERC20Contract()
			// NOTE: values below are for local tests
			sgnParams := &SGNParams{
				CelrAddr:               erc20Addr,
				GovernProposalDeposit:  big.NewInt(1000000000000000000), // 1 CELR
				GovernVoteTimeout:      big.NewInt(90),
				SlashTimeout:           big.NewInt(15),
				MaxBondedValidators:    big.NewInt(5),
				MinValidatorTokens:     big.NewInt(1000000000000000000),
				MinSelfDelegation:      big.NewInt(1e18),
				AdvanceNoticePeriod:    big.NewInt(30),
				ValidatorBondInterval:  big.NewInt(24 * 3600),
				SidechainGoLiveTimeout: big.NewInt(0),
			}
			tx, stakingAddr, sgnAddr := DeployStakingSGNContracts(sgnParams)
			WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "DeployStakingContracts")

			configFileViper.Set(common.FlagEthCelrAddress, erc20Addr.Hex())
			configFileViper.Set(common.FlagEthStakingAddress, stakingAddr.Hex())
			configFileViper.Set(common.FlagEthSGNAddress, sgnAddr.Hex())
			err = configFileViper.WriteConfig()
			ChkErr(err, "failed to write config")

			if ethurl == LocalGeth {
				amt := new(big.Int)
				amt.SetString("1"+strings.Repeat("0", 20), 10)
				tx, err := erc20.Approve(EtherBaseAuth, stakingAddr, amt)
				ChkErr(err, "failed to approve erc20")
				WaitMinedWithChk(context.Background(), EthClient, tx, BlockDelay, PollingInterval, "approve erc20")
				StakingContract, _ = eth.NewStaking(stakingAddr, EthClient)
				_, err = StakingContract.ContributeToRewardPool(EtherBaseAuth, amt)
				ChkErr(err, "failed to call ContributeToMiningPool of Staking contract")
				err = FundAddrsErc20(erc20Addr,
					[]eth.Addr{
						eth.Hex2Addr(ClientEthAddrs[0]),
						eth.Hex2Addr(ClientEthAddrs[1]),
					},
					"1"+strings.Repeat("0", 20),
				)
				ChkErr(err, "fund test CELR to clients")

			}

			return nil
		},
	}

	return cmd
}
