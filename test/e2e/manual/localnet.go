package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/app"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/test/e2e/multinode"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	stakingcli "github.com/celer-network/sgn-v2/x/staking/client/cli"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	"github.com/spf13/viper"
)

var (
	start   = flag.Bool("start", false, "start local testnet")
	cbr     = flag.Bool("cbr", false, "start with cbridge")
	msg     = flag.Bool("msg", false, "start with message passing capabilities")
	op      = flag.Bool("op", false, "proceed with sample operations")
	report  = flag.Bool("report", false, "liveness report and price sync")
	full    = flag.Bool("full", false, "start with full stack setup")
	auto    = flag.Bool("auto", false, "auto-add all validators")
	down    = flag.Bool("down", false, "shutdown local testnet")
	up      = flag.Int("up", -1, "start a testnet node")
	stop    = flag.Int("stop", -1, "stop a testnet node")
	upall   = flag.Bool("upall", false, "start all nodes")
	stopall = flag.Bool("stopall", false, "stop all nodes")
	rebuild = flag.Bool("rebuild", false, "rebuild sgn node docker image")
	fund    = flag.String("fund", "", "fund test tokens to give address")
)

func main() {
	flag.Parse()
	repoRoot, _ := filepath.Abs("../../..")
	if *full {
		*cbr = true
		*msg = true
		*auto = true
	}
	if *start {
		multinode.BuildDockers()
		multinode.SetupMainchain()
		if *cbr {
			multinode.SetupBridgeChains()
		}
		tc.SetupSgnchain()

		p := &tc.ContractParams{
			CelrAddr:              tc.CelrAddr,
			ProposalDeposit:       big.NewInt(1e17),
			VotePeriod:            big.NewInt(30),
			UnbondingPeriod:       big.NewInt(5),
			MaxBondedValidators:   big.NewInt(5),
			MinValidatorTokens:    big.NewInt(5e18),
			MinSelfDelegation:     big.NewInt(2e18),
			AdvanceNoticePeriod:   big.NewInt(1),
			ValidatorBondInterval: big.NewInt(0),
			MaxSlashFactor:        big.NewInt(1e5),
		}
		multinode.SetupNewSgnEnv(p, *cbr, *msg, true, *report)
		if *cbr {
			amts := []*big.Int{big.NewInt(1e18)}
			tc.CbrChain1.SetInitSigners(amts)
			tc.CbrChain2.SetInitSigners(amts)
		}

		log.Infoln("install sgnd in host machine")
		cmd := exec.Command("make", "install")
		cmd.Dir = repoRoot
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "WITH_CLEVELDB=yes")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}

		log.Infoln("copy config files")
		cmd2 := exec.Command("make", "copy-manual-test-data")
		cmd2.Dir = repoRoot
		if err := cmd2.Run(); err != nil {
			log.Fatal(err)
		}
		log.Infoln("update config files")
		for i := 0; i < len(tc.ValEthKs); i++ {
			configPath := fmt.Sprintf("./data/node%d/sgnd/config/sgn.toml", i)
			configFileViper := viper.New()
			configFileViper.SetConfigFile(configPath)
			if err := configFileViper.ReadInConfig(); err != nil {
				log.Fatal(err)
			}
			ksPath, _ := filepath.Abs(fmt.Sprintf("./data/node%d/keys/vsigner%d.json", i, i))
			configFileViper.Set(common.FlagEthSignerKeystore, ksPath)
			configFileViper.Set(common.FlagEthGateway, tc.LocalGeth)
			configFileViper.Set(common.FlagSgnNodeURI, tc.SgnNodeURIs[i])
			if err := configFileViper.WriteConfig(); err != nil {
				log.Fatal(err)
			}
		}

		if *auto {
			time.Sleep(10 * time.Second)
			addValidators()

			if *cbr && *op {
				cbrOps()
			}
		}
	} else if *down {
		log.Infoln("Tearing down all containers...")
		cmd := exec.Command("make", "localnet-down")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	} else if *up != -1 {
		log.Infoln("Start node", *up)
		cmd := exec.Command("docker-compose", "up", "-d", fmt.Sprintf("sgnnode%d", *up))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else if *stop != -1 {
		log.Infoln("Stop node", *stop)
		cmd := exec.Command("docker-compose", "stop", fmt.Sprintf("sgnnode%d", *stop))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	} else if *upall {
		log.Infoln("Start all nodes ...")
		cmd := exec.Command("make", "localnet-up-nodes")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	} else if *stopall {
		log.Infoln("Stop all nodes ...")
		cmd := exec.Command("make", "localnet-down-nodes")
		cmd.Dir = repoRoot
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	} else if *rebuild {
		log.Infoln("Rebuild sgn node docker image ...")
		cmd := exec.Command("make", "build-node")
		cmd.Dir = repoRoot
		cmd.Env = os.Environ()
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		log.Infoln("install sgnd in host machine ...")
		cmd = exec.Command("make", "install")
		cmd.Dir = repoRoot
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "WITH_CLEVELDB=yes")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	} else if *fund != "" {
		err := fundAddr()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}
}

func addValidators() {
	encodingConfig := app.MakeEncodingConfig()
	txr, err := transactor.NewTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.SgnValAcct,
		tc.SgnPassphrase,
		encodingConfig.Amino,
		encodingConfig.Codec,
		encodingConfig.InterfaceRegistry,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	txr.Run(0)

	valAmts := []*big.Int{
		new(big.Int).Mul(big.NewInt(16000), big.NewInt(common.CelrPrecision)),
		new(big.Int).Mul(big.NewInt(20000), big.NewInt(common.CelrPrecision)),
		new(big.Int).Mul(big.NewInt(15000), big.NewInt(common.CelrPrecision)),
		new(big.Int).Mul(big.NewInt(18000), big.NewInt(common.CelrPrecision)),
	}
	commissions := []uint64{eth.CommissionRate(0.15), eth.CommissionRate(0.2), eth.CommissionRate(0.12), eth.CommissionRate(0.1)}

	for i := 0; i < 4; i++ {
		log.Infoln("Adding validator", i, tc.ValEthAddrs[i].Hex())
		err := tc.InitializeValidator(tc.ValAuths[i], tc.ValSignerAddrs[i], tc.ValSgnAddrs[i], valAmts[i], commissions[i])
		if err != nil {
			log.Fatalln("failed to initialize validator: ", err)
		}

		for retry := 0; retry < tc.RetryLimit; retry++ {
			v, err := stakingcli.QueryValidator(txr.CliCtx, eth.Addr2Hex(tc.ValEthAddrs[i]))
			if err == nil {
				log.Infof("query validator success: %s", v)
				break
			}
			time.Sleep(tc.RetryPeriod)
		}
		for j := 0; j <= i; j++ {
			err = tc.Delegate(tc.DelAuths[j], tc.ValEthAddrs[i], tc.NewBigInt(10+i+j, 19))
		}
		if i == 0 {
			configFileViper := viper.New()
			configFileViper.SetConfigFile("./data/node0/sgnd/config/sgn.toml")
			if err := configFileViper.ReadInConfig(); err != nil {
				log.Fatal(err)
			}
			transactors := configFileViper.GetStringSlice(common.FlagSgnTransactors)
			msg1 := stakingtypes.NewMsgSetTransactors(stakingtypes.SetTransactorsOp_Overwrite, transactors, txr.Key.GetAddress().String())
			err = msg1.ValidateBasic()
			if err != nil {
				log.Fatal(err)
				return
			}
			txr.AddTxMsg(&msg1)
			description := stakingtypes.NewDescription(
				"node0", stakingtypes.DoNotModifyDesc, "www.celer.network", "sgn-validator@celer.network", stakingtypes.DoNotModifyDesc)
			msg2 := stakingtypes.NewMsgEditDescription(description, txr.Key.GetAddress().String())
			err = msg2.ValidateBasic()
			if err != nil {
				log.Fatal(err)
				return
			}
			txr.AddTxMsg(&msg2)
		}
	}
}

func fundAddr() error {
	addrs := []eth.Addr{eth.Hex2Addr(*fund)}
	err := tc.FundAddrsETH(addrs, tc.NewBigInt(1, 20), tc.LocalGeth, int64(tc.ChainID))
	if err != nil {
		return err
	}
	tc.SetupEthClients()

	configFileViper := viper.New()
	configFileViper.SetConfigFile("./data/node0/sgnd/config/sgn.toml")
	if err := configFileViper.ReadInConfig(); err != nil {
		return err
	}
	celrAddr := eth.Hex2Addr(configFileViper.GetString(common.FlagEthContractCelr))
	err = tc.FundAddrsErc20(celrAddr, addrs, tc.NewBigInt(1, 25), tc.EthClient, tc.EtherBaseAuth)
	if err != nil {
		return err
	}
	if *cbr {
		tc.InitCbrChainConfigs()
		err = tc.FundAddrsETH(addrs, tc.NewBigInt(1, 20), tc.LocalGeth2, int64(tc.Geth2ChainID))
		if err != nil {
			return err
		}
		genesisViper := viper.New()
		genesisViper.SetConfigFile("./data/node0/sgnd/config/genesis.json")
		if err := genesisViper.ReadInConfig(); err != nil {
			return err
		}
		cbrConfig := new(cbrtypes.CbrConfig)
		jsonByte, _ := json.Marshal(genesisViper.Get("app_state.cbridge.config"))
		json.Unmarshal(jsonByte, cbrConfig)
		err = tc.FundAddrsErc20(eth.Hex2Addr(cbrConfig.Assets[0].Addr), addrs, tc.NewBigInt(1, 13), tc.CbrChain1.Ec, tc.CbrChain1.Auth)
		if err != nil {
			return err
		}
		err = tc.FundAddrsErc20(eth.Hex2Addr(cbrConfig.Assets[1].Addr), addrs, tc.NewBigInt(1, 13), tc.CbrChain2.Ec, tc.CbrChain2.Auth)
		if err != nil {
			return err
		}
	}
	return nil
}

func cbrOps() {
	encodingConfig := app.MakeEncodingConfig()
	txr, err := transactor.NewTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.SgnValAcct,
		tc.SgnPassphrase,
		encodingConfig.Amino,
		encodingConfig.Codec,
		encodingConfig.InterfaceRegistry,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infoln("======================== Add liquidity on chain 1 ===========================")
	addAmt := big.NewInt(5 * 1e10)
	var i uint64
	for i = 0; i < 2; i++ {
		err = tc.CbrChain1.ApproveUSDT(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 approve", i))
		err = tc.CbrChain1.AddLiq(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain1 addliq", i))
		tc.CheckAddLiquidityStatus(txr, tc.CbrChain1.ChainId, i+1)
	}
	log.Infoln("======================== Add liquidity on chain 2 ===========================")
	for i = 0; i < 2; i++ {
		err = tc.CbrChain2.ApproveUSDT(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain2 approve", i))
		err = tc.CbrChain2.AddLiq(i, addAmt)
		tc.ChkErr(err, fmt.Sprintf("u%d chain2 addliq", i))
		tc.CheckAddLiquidityStatus(txr, tc.CbrChain2.ChainId, i+1)
	}

	chainTokens := make([]*cbrtypes.ChainTokenAddrPair, 0)
	chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
		ChainId:   tc.CbrChain1.ChainId,
		TokenAddr: tc.CbrChain1.USDTAddr.Hex(),
	})
	chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
		ChainId:   tc.CbrChain2.ChainId,
		TokenAddr: tc.CbrChain2.USDTAddr.Hex(),
	})
	res, err := cbrcli.QueryLiquidityDetailList(txr.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     tc.ClientEthAddrs[0].Hex(),
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	log.Infoln("QueryLiquidityDetailList resp:", res.String())

	log.Infoln("======================== Xfer ===========================")
	xferAmt := big.NewInt(1e10)
	err = tc.CbrChain1.ApproveUSDT(0, xferAmt)
	tc.ChkErr(err, "u0 chain1 approve")
	xferId, err := tc.CbrChain1.Send(0, xferAmt, tc.CbrChain2.ChainId, 1)
	tc.ChkErr(err, "u0 chain1 send")
	tc.CheckXfer(txr, xferId[:])
	log.Infof("--- transfer Id %x", xferId)

	log.Infoln("======================== LP withdraw liquidity ===========================")
	reqid := uint64(time.Now().Unix())
	wdLq1 := tc.CbrChain1.GetWithdrawLq(20000000) // withdraw 20%
	wdLq2 := tc.CbrChain2.GetWithdrawLq(10000000) // withdraw 10%
	err = tc.CbrChain1.StartWithdrawRemoveLiquidity(txr, reqid, 0, wdLq1, wdLq2)
	tc.ChkErr(err, "u0 chain1 start withdraw")
	log.Infoln("--- withdraw reqid:", reqid)
	detail := tc.GetWithdrawDetailWithSigs(txr, tc.CbrChain1.Users[0].Address, reqid, 4)
	curss, err := tc.GetCurSortedSigners(txr, tc.CbrChain1.ChainId)
	tc.ChkErr(err, "chain1 GetCurSortedSigners")
	err = tc.CbrChain1.OnchainCbrWithdraw(detail, curss)
	tc.ChkErr(err, "chain1 onchain withdraw")

	res, err = cbrcli.QueryLiquidityDetailList(txr.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     tc.ClientEthAddrs[0].Hex(),
		ChainToken: chainTokens,
	})
	tc.ChkErr(err, "cli Query")
	log.Infoln("QueryLiquidityDetailList resp:", res.String())

	log.Infoln("======================== Xfer back ===========================")
	err = tc.CbrChain2.ApproveUSDT(0, xferAmt)
	tc.ChkErr(err, "u0 chain2 approve")
	xferId, err = tc.CbrChain2.Send(0, xferAmt, tc.CbrChain1.ChainId, 1)
	tc.ChkErr(err, "u0 chain2 send")
	tc.CheckXfer(txr, xferId[:])
	log.Infof("--- transfer Id %x", xferId)

	log.Infoln("======================== LP claim farming reward on-chain ===========================")
	err = tc.StartClaimFarmingRewards(txr, 0)
	tc.ChkErr(err, "u0 start claim all farming rewards")
	info := tc.GetFarmingRewardClaimInfoWithSigs(txr, 0, 4)
	err = tc.OnchainClaimFarmingRewards(&info.RewardClaimDetailsList[0])
	tc.ChkErr(err, "u0 onchain claim farming rewards")
}
