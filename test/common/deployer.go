package common

import (
	"context"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/test/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *CbrChain) DeployMessageContracts() {
	c.MessageBusAddr, c.MessageBusContract =
		DeployMessageBusContract(c.Ec, c.Auth, c.CbrAddr, c.PegBridgeAddr, c.PegVaultAddr, c.PegBridgeV2Addr, c.PegVaultV2Addr)
	c.BatchTransferAddr, c.BatchTransferContract =
		DeployBatchTransferContract(c.Ec, c.Auth, c.MessageBusAddr)
	c.MsgTestAddr, c.MsgTestContract =
		DeployMsgTestContract(c.Ec, c.Auth, c.MessageBusAddr)
}

func (c *CbrChain) DeployPegVaultContracts() {
	log.Infof("Deploying peg vault and test token")
	c.PegVaultAddr, c.PegVaultContract = DeployPegVaultContract(c.Ec, c.Auth, c.CbrAddr)
	c.PegVaultV2Addr, c.PegVaultV2Contract = DeployPegVaultV2Contract(c.Ec, c.Auth, c.CbrAddr)
	c.UNIAddr, c.UNIContract = DeployBridgeTestTokenContract(c.Ec, c.Auth, "UNI", "UNI", 18)
}

func (c *CbrChain) DeployPegBridgeContracts() {
	log.Infof("Deploying peg bridge and test token")
	c.PegBridgeAddr, c.PegBridgeContract = DeployPegBridgeContract(c.Ec, c.Auth, c.CbrAddr)
	c.PegBridgeV2Addr, c.PegBridgeV2Contract = DeployPegBridgeV2Contract(c.Ec, c.Auth, c.CbrAddr)
	c.UNIAddr, c.UNIContract = DeployBridgeTestTokenContract(c.Ec, c.Auth, "UNI", "UNI", 18)
	c.FETAddr, c.FETContract = DeployBridgeTestTokenContract(c.Ec, c.Auth, "FET", "FET", 18)
}

func DeployERC20Contract(ethClient *ethclient.Client, auth *bind.TransactOpts, name, symbol string, decimal uint8) (eth.Addr, *eth.Erc20) {
	initAmt := NewBigInt(1, 28) // 10 billion in 18 decimal
	erc20Addr, tx, erc20, err := eth.DeployErc20(auth, ethClient, name, symbol, initAmt, decimal)
	ChkErr(err, "failed to deploy ERC20")
	log.Infoln("Erc20 address:", erc20Addr.String())
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployERC20")
	return erc20Addr, erc20
}

func DeployBridgeTestTokenContract(
	ethClient *ethclient.Client, auth *bind.TransactOpts, name, symbol string, decimals uint8) (eth.Addr, *contracts.BridgeTestToken) {
	tokenAddr, tx, token, err := contracts.DeployBridgeTestToken(auth, ethClient, name, symbol, decimals)
	ChkErr(err, "failed to deploy BridgeTestToken")
	log.Infoln("PeggedToken address:", tokenAddr.String())
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployBridgeTestToken")

	tx, err = token.UpdateBridgeSupplyCap(auth, auth.From, NewBigInt(1, 28))
	ChkErr(err, "failed to update supply cap")
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "UpdateBridgeSupplyCap")
	tx, err = token.Mint(auth, auth.From, NewBigInt(1, 28))
	ChkErr(err, "failed to mint init tokens")
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "Mint")

	return tokenAddr, token
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

func DeployPegBridgeContract(
	ethClient *ethclient.Client, auth *bind.TransactOpts, sigsVerifier eth.Addr) (ptbAddr eth.Addr, ptbContract *eth.PegBridgeContract) {
	ptbAddr, tx, _, err := eth.DeployPeggedTokenBridge(auth, ethClient, sigsVerifier)
	ChkErr(err, "failed to deploy PeggedTokenBridge contract")
	ptbContract, err = eth.NewPegBridgeContract(ptbAddr, ethClient)
	ChkErr(err, "failed to set PeggedTokenBridge contract")
	log.Infoln("ptb address:", ptbAddr.String())
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployPegBridgeContract")
	return
}

func DeployPegBridgeV2Contract(
	ethClient *ethclient.Client, auth *bind.TransactOpts, sigsVerifier eth.Addr) (ptbAddr eth.Addr, ptbContract *eth.PegBridgeV2Contract) {
	ptbAddr, tx, _, err := eth.DeployPeggedTokenBridgeV2(auth, ethClient, sigsVerifier)
	ChkErr(err, "failed to deploy PeggedTokenBridgeV2 contract")
	ptbContract, err = eth.NewPegBridgeV2Contract(ptbAddr, ethClient)
	ChkErr(err, "failed to set PeggedTokenBridgeV2 contract")
	log.Infoln("ptb v2 address:", ptbAddr.String())
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployPegBridgeContractV2")
	return
}

func DeployPegVaultContract(
	ethClient *ethclient.Client, auth *bind.TransactOpts, sigsVerifier eth.Addr) (otvAddr eth.Addr, otvContract *eth.PegVaultContract) {
	otvAddr, tx, _, err := eth.DeployOriginalTokenVault(auth, ethClient, sigsVerifier)
	ChkErr(err, "failed to deploy OriginalTokenVault contract")
	otvContract, err = eth.NewPegVaultContract(otvAddr, ethClient)
	ChkErr(err, "failed to set OriginalTokenVault contract")
	log.Infoln("otv address:", otvAddr.String())
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployPegVaultContract")
	return
}

func DeployPegVaultV2Contract(
	ethClient *ethclient.Client, auth *bind.TransactOpts, sigsVerifier eth.Addr) (otvAddr eth.Addr, otvContract *eth.PegVaultV2Contract) {
	otvAddr, tx, _, err := eth.DeployOriginalTokenVaultV2(auth, ethClient, sigsVerifier)
	ChkErr(err, "failed to deploy OriginalTokenVaultV2 contract")
	otvContract, err = eth.NewPegVaultV2Contract(otvAddr, ethClient)
	ChkErr(err, "failed to set OriginalTokenVaultV2 contract")
	log.Infoln("otv v2 address:", otvAddr.String())
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployPegVaultContractV2")
	return
}

func DeployWithdrawInboxContract(ethClient *ethclient.Client, auth *bind.TransactOpts) (eth.Addr, *eth.WdInboxContract) {
	addr, tx, contract, err := eth.DeployWithdrawInbox(auth, ethClient)
	ChkErr(err, "failed to deploy WithdrawInbox")
	log.Infoln("WithdrawInbox address", addr)
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployWithdrawInboxContract")
	return addr, &eth.WdInboxContract{WithdrawInbox: contract, Address: addr}
}

func DeployContractAsLPContract(ethClient *ethclient.Client, auth *bind.TransactOpts, bridge, wdInbox eth.Addr) (eth.Addr, *eth.CLPContract) {
	addr, tx, contract, err := eth.DeployContractAsLP(auth, ethClient, bridge, wdInbox)
	ChkErr(err, "failed to deploy ContractAsLP")
	log.Infoln("ContractAsLP address", addr)
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployContractAsLPContract")
	return addr, &eth.CLPContract{ContractAsLP: contract, Address: addr}
}

func DeployMessageBusContract(ethClient *ethclient.Client, auth *bind.TransactOpts, bridge, pegBridge, pegVault, pegBridgeV2, pegVaultV2 eth.Addr) (eth.Addr, *eth.MessageBus) {
	addr, tx, contract, err := eth.DeployMessageBus(auth, ethClient, bridge, bridge, pegBridge, pegVault, pegBridgeV2, pegVaultV2)
	ChkErr(err, "failed to deploy MessageBus")
	log.Infoln("MessageBus address", addr)
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployMessageBusContract")
	// Only set fee base in tests
	tx, err = contract.SetFeeBase(auth, MsgFeeBase)
	ChkErr(err, "failed to set message fee base")
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "SetFeeBase")
	return addr, contract
}

func DeployBatchTransferContract(ethClient *ethclient.Client, auth *bind.TransactOpts, bus eth.Addr) (eth.Addr, *contracts.BatchTransfer) {
	addr, tx, contract, err := contracts.DeployBatchTransfer(auth, ethClient, bus)
	ChkErr(err, "failed to deploy BatchTransfer")
	log.Infoln("BatchTransfer address", addr)
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployBatchTransferContract")
	return addr, contract
}

func DeployMsgTestContract(ethClient *ethclient.Client, auth *bind.TransactOpts, bus eth.Addr) (eth.Addr, *contracts.MsgTest) {
	addr, tx, contract, err := contracts.DeployMsgTest(auth, ethClient, bus)
	ChkErr(err, "failed to deploy BatchTransfer")
	log.Infoln("BatchTransfer address", addr)
	WaitMinedWithChk(context.Background(), ethClient, tx, BlockDelay, PollingInterval, "DeployBatchTransferContract")
	return addr, contract
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
