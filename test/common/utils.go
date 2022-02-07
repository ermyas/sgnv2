package common

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAuth(ksfile string, chainId int64) (addr eth.Addr, auth *bind.TransactOpts, err error) {
	keystoreBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return
	}
	key, err := keystore.DecryptKey(keystoreBytes, "")
	if err != nil {
		return
	}
	addr = key.Address
	auth, err = bind.NewTransactorWithChainID(strings.NewReader(string(keystoreBytes)), "", big.NewInt(chainId))
	if err != nil {
		return
	}
	return
}

func GetEthPrivateKey(ksfile string) (*ecdsa.PrivateKey, error) {
	keystoreBytes, err := ioutil.ReadFile(ksfile)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(keystoreBytes, "")
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}

func WaitMinedWithChk(
	ctx context.Context,
	conn *ethclient.Client,
	tx *ethtypes.Transaction,
	blockDelay uint64,
	pollingInterval time.Duration,
	txname string,
) {
	ctx2, cancel := context.WithTimeout(ctx, WaitMinedTimeout)
	defer cancel()
	receipt, err := ethutils.WaitMined(ctx2, conn, tx, ethutils.WithBlockDelay(blockDelay), ethutils.WithPollingInterval(pollingInterval))
	ChkErr(err, "WaitMined error")
	if receipt.Status != ethtypes.ReceiptStatusSuccessful {
		log.Fatalln(txname, "tx failed")
	}
	log.Infoln(txname, "tx success")
}

func LogBlkNum(conn *ethclient.Client) {
	header, err := conn.HeaderByNumber(context.Background(), nil)
	ChkErr(err, "failed to get HeaderByNumber")
	log.Infoln("Latest block number on mainchain:", header.Number)
}

func Sleep(second time.Duration) {
	time.Sleep(second * time.Second)
}

func SleepWithLog(second time.Duration, waitFor string) {
	log.Infof("Sleep %d seconds for %s", second, waitFor)
	Sleep(second)
}

func SleepBlocksWithLog(count time.Duration, waitFor string) {
	SleepWithLog(count*SgnBlockInterval, waitFor)
}

func NewBigInt(nonZeroDigits, trailingZeros int) *big.Int {
	value := new(big.Int)
	value.SetString(strconv.Itoa(nonZeroDigits)+strings.Repeat("0", trailingZeros), 10)
	if value == nil {
		log.Fatalf("invalid NewBigInt input %d %d", nonZeroDigits, trailingZeros)
	}
	return value
}

func RunCmd(name string, args ...string) {
	argstr := strings.Join(args, " ")
	cmdstr := name + " " + argstr
	err := RunCmdNoChkErr(name, args...)
	ChkErr(err, "Failed to "+cmdstr)
}

func RunCmdNoChkErr(name string, args ...string) error {
	argstr := strings.Join(args, " ")
	log.Infoln(name, argstr)
	cmd := exec.Command(name, args...)
	cmd.Dir, _ = filepath.Abs("../../..")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func RunAllAndWait(funcs ...func()) {
	var wg sync.WaitGroup
	for _, run := range funcs {
		wg.Add(1)
		go func(run func()) {
			run()
			wg.Done()
		}(run)
	}
	wg.Wait()
}
