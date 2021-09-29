package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	gw     = flag.String("gw", "https://goerli.infura.io/v3/c58bb2a63c1a466abb6d7b5e0b7b6828", "json rpc server")
	celr   = flag.String("celr", "0x5D3c0F4cA5EE99f8E8F59Ff9A5fAb04F6a7e007f", "celer addr")
	ks     = flag.String("ks", "", "keystore json")
	kspass = flag.String("kspass", "", "keystore password")
)

func main() {
	flag.Parse()
	ec, err := ethclient.Dial(*gw)
	chkErr(err, "dial")
	chainid, err := ec.ChainID(context.Background())
	chkErr(err, "")
	log.Print("chainid: ", chainid)
	auth, _ := kspath2auth(*ks, *kspass, chainid)

	celrAddr := eth.Hex2Addr(*celr)
	stakingContractAddr, _, staking, _ := eth.DeployStaking(
		auth,
		ec,
		celrAddr,
		big.NewInt(1e18),
		big.NewInt(90),
		big.NewInt(15),
		big.NewInt(5),
		big.NewInt(1e18),
		big.NewInt(1e18),
		big.NewInt(30),
		big.NewInt(0),
		big.NewInt(1e5))
	log.Print("Staking address:", stakingContractAddr.String())

	sgnContractAddr, _, _, _ := eth.DeploySGN(auth, ec, stakingContractAddr)
	log.Print("SGN address:", sgnContractAddr.String())

	stakingRewardContractAddr, _, _, _ := eth.DeployStakingReward(
		auth, ec, stakingContractAddr)
	log.Print("StakingReward address:", stakingRewardContractAddr.String())

	farmingRewardsContractAddr, _, _, _ := eth.DeployFarmingRewards(
		auth, ec, stakingContractAddr)
	log.Print("FarmingRewards address:", farmingRewardsContractAddr.String())

	viewerContractAddr, _, _, _ := eth.DeployViewer(auth, ec, stakingContractAddr)
	log.Print("Viewer address:", viewerContractAddr.String())

	governContractAddr, _, _, _ := eth.DeployGovern(
		auth, ec, stakingContractAddr, celrAddr, stakingRewardContractAddr)
	log.Print("Govern address:", governContractAddr.String())

	auth.GasLimit = 8000000
	_, err = staking.SetGovContract(auth, governContractAddr)
	if err != nil {
		log.Print("Error: failed to set gov contract!")
	}

	_, err = staking.SetRewardContract(auth, stakingRewardContractAddr)
	if err != nil {
		log.Print("Error: failed to set gov contract!")
	}
}

func kspath2auth(kspath string, kspass string, chainid *big.Int) (*bind.TransactOpts, error) {
	ksjson, err := ioutil.ReadFile(kspath)
	if err != nil {
		return nil, err
	}
	kss := string(ksjson)
	return bind.NewTransactorWithChainID(strings.NewReader(kss), kspass, chainid)
}

func chkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
