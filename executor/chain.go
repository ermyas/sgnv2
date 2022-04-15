package executor

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/celer-network/endpoint-proxy/endpointproxy"
	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/executor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

type Chain struct {
	ChainID     uint64
	Transactor  *ethutils.Transactor
	MsgBus      *eth.MsgBusContract
	LiqBridge   *eth.BridgeContract
	PegBridge   *eth.PegBridgeContract
	PegBridgeV2 *eth.PegBridgeV2Contract
	PegVault    *eth.PegVaultContract
	PegVaultV2  *eth.PegVaultV2Contract
	fwdBlkDelay uint64
	monitor2    *mon2.Monitor
	startBlk    *big.Int
	filterAddr  string
}

// key is chainid
type ChainMgr struct {
	chains map[uint64]*Chain
	lock   sync.RWMutex
	initWg sync.WaitGroup

	contractAddrs []eth.Addr
}

var Chains *ChainMgr

func NewChainMgr(dal *DAL) *ChainMgr {
	log.Infoln("Initializing chains")
	var configs []*common.OneChainConfig
	err := viper.UnmarshalKey(common.FlagMultiChain, &configs)
	if err != nil {
		log.Fatalln("failed to load multichain configs", err)
	}
	var chains = &ChainMgr{chains: make(map[uint64]*Chain)}
	for _, config := range configs {
		chains.initWg.Add(1)
		go chains.addChain(config, dal)
	}
	chains.initWg.Wait()
	// add filterAddr
	contracts := []*types.ContractConfig{}
	err = viper.UnmarshalKey(types.FlagExecutorContracts, &contracts)
	if err != nil {
		log.Fatalln("failed to initialize contract filters", err)
	}
	contractMap := make(map[uint64]string)
	for _, contract := range contracts {
		contractMap[contract.ChainId] = contract.Address
	}
	for _, chain := range chains.chains {
		chain.filterAddr = contractMap[chain.ChainID]
	}
	log.Infoln("Finished initializing all chains")
	Chains = chains
	return chains
}

func (m *ChainMgr) addChain(config *common.OneChainConfig, dal *DAL) {
	log.Infoln("Add chain", config)
	ec := newEthClient(config)
	transactor := newTransactor(config, ec)

	// init monitor
	chainConfig := mon2.PerChainCfg{
		BlkIntv:         time.Duration(config.BlkInterval) * time.Second,
		BlkDelay:        config.BlkDelay,
		MaxBlkDelta:     config.MaxBlkDelta,
		ForwardBlkDelay: config.ForwardBlkDelay,
	}
	mon, err := mon2.NewMonitor(ec, dal, chainConfig)
	if err != nil {
		log.Fatalln("failed to create monitor: ", err)
	}

	chain := &Chain{
		ChainID:    config.ChainID,
		Transactor: transactor,
		monitor2:   mon,
	}
	addrs := chain.initContracts(ec, config)
	m.contractAddrs = append(m.contractAddrs, addrs...)
	m.lock.Lock()
	defer func() {
		m.lock.Unlock()
		m.initWg.Done()
	}()
	m.chains[chain.ChainID] = chain
}

func (m *ChainMgr) GetChain(chid uint64) (*Chain, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	chain, ok := m.chains[chid]
	if !ok {
		err := fmt.Errorf("cannot get chain: chain id %d does not exist", chid)
		log.Error(err)
		return nil, err
	}
	return chain, nil
}

func (m *ChainMgr) GetAllContractAddrs() []eth.Addr {
	return m.contractAddrs
}

func (m *ChainMgr) StartMonitoring() {
	log.Info("Start monitoring on-chain events")
	for _, chain := range m.chains {
		go chain.startMonitoring()
	}
}

func (c *Chain) initContracts(ec *ethclient.Client, config *common.OneChainConfig) []eth.Addr {
	var err error
	c.LiqBridge, err = eth.NewBridgeContract(eth.Hex2Addr(config.CBridge), ec)
	check(err)
	c.PegBridge, err = eth.NewPegBridgeContract(eth.Hex2Addr(config.PTBridge), ec)
	check(err)
	c.PegVault, err = eth.NewPegVaultContract(eth.Hex2Addr(config.OTVault), ec)
	check(err)
	c.PegBridgeV2, err = eth.NewPegBridgeV2Contract(eth.Hex2Addr(config.PTBridge2), ec)
	check(err)
	c.PegVaultV2, err = eth.NewPegVaultV2Contract(eth.Hex2Addr(config.OTVault2), ec)
	check(err)
	c.MsgBus, err = eth.NewMsgBusContract(eth.Hex2Addr(config.MsgBus), ec)
	check(err)
	return []eth.Addr{c.LiqBridge.Address, c.PegBridge.Address, c.PegVault.Address, c.PegBridgeV2.Address, c.PegVaultV2.Address, c.MsgBus.Address}
}

func newEthClient(config *common.OneChainConfig) *ethclient.Client {
	// init eth client
	log.Infof("Dialing eth client for chain %d at %s", config.ChainID, config.Gateway)
	var ec *ethclient.Client
	var err error
	if config.ProxyPort > 0 {
		proxyPort := config.ProxyPort + 600
		if err = endpointproxy.StartProxy(config.Gateway, config.ChainID, proxyPort); err != nil {
			log.Fatalln("can not start proxy for chain:", config.ChainID, "gateway:", config.Gateway, "port:", proxyPort, "err:", err)
		}
		ec, err = ethclient.Dial(fmt.Sprintf("http://127.0.0.1:%d", proxyPort))
		if err != nil {
			log.Fatalln("dial", config.Gateway, "err:", err)
		}
	} else {
		ec, err = ethclient.Dial(config.Gateway)
		if err != nil {
			log.Fatalln("dial", config.Gateway, "err:", err)
		}
	}
	checkChainID(ec, config.ChainID)
	return ec
}

func checkChainID(ec *ethclient.Client, chainID uint64) {
	chid, err := ec.ChainID(context.Background())
	if err != nil {
		log.Fatalf("get chainid %d err: %s", chainID, err)
	}
	if chid.Uint64() != chainID {
		log.Fatalf("chainid mismatch! chainConf has %d but onchain has %d", chainID, chid.Uint64())
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func newTransactor(config *common.OneChainConfig, ec *ethclient.Client) *ethutils.Transactor {
	signerKey, signerPass := viper.GetString(common.FlagEthSignerKeystore), viper.GetString(common.FlagEthSignerPassphrase)
	signer, addr, err := eth.CreateSigner(signerKey, signerPass, new(big.Int).SetUint64(config.ChainID))
	if err != nil {
		log.Fatalln("CreateSigner err:", err)
	}
	return ethutils.NewTransactorByExternalSigner(
		addr,
		signer,
		ec,
		big.NewInt(int64(config.ChainID)),
		ethutils.WithBlockDelay(config.BlkDelay),
		ethutils.WithPollingInterval(time.Duration(config.BlkInterval)*time.Second),
		ethutils.WithAddGasEstimateRatio(config.AddGasEstimateRatio),
		ethutils.WithGasLimit(config.GasLimit),
		ethutils.WithAddGasGwei(config.AddGasGwei),
		ethutils.WithMaxGasGwei(config.MaxGasGwei),
		ethutils.WithMinGasGwei(config.MinGasGwei),
		ethutils.WithMaxFeePerGasGwei(config.MaxFeePerGasGwei),
		ethutils.WithMaxPriorityFeePerGasGwei(config.MaxPriorityFeePerGasGwei),
	)
}

func (c *Chain) NewExecuteRefundHandler(messageId []byte, execute types.RefundTxFunc) types.ExecuteRefund {
	// returns a handler function
	return func(req []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) error {
		log.Infof("executing refund init (messageId %x)", messageId)
		err := Dal.UpdateStatus(messageId, types.ExecutionStatus_Init_Refund_Executing)
		if err != nil {
			return err
		}
		tx, err := c.Transactor.Transact(&ethutils.TransactionStateHandler{
			OnMined: func(receipt *gethtypes.Receipt) {
				status := types.ExecutionStatus_Init_Refund_Failed
				if receipt.Status == gethtypes.ReceiptStatusSuccessful {
					log.Infof("Refund init (messageId %x) mined and succeeded: tx %x", messageId, receipt.TxHash)
					status = types.ExecutionStatus_Init_Refund_Executed
					// reset retry count to zero
					Dal.UpdateRetryCount(messageId, 0)
				} else {
					log.Errorf("Refund init (messageId %x) mined but failed: tx %x", messageId, receipt.TxHash)
				}
				Dal.UpdateStatus(messageId, status)
			},
			OnError: func(tx *gethtypes.Transaction, err error) {
				log.Errorf("execute refund init error: txhash %s, err %v", tx.Hash(), err)
				Dal.UpdateStatus(messageId, types.ExecutionStatus_Init_Refund_Failed)
			},
		}, func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
			return execute(opts, req, sortedSigs, signers, powers)
		})
		if err != nil {
			if strings.Contains(err.Error(), "transfer exists") ||
				strings.Contains(err.Error(), "record exists") {
				log.Errorf("refund transfer already executed (messageId %x)", messageId)
				Dal.UpdateStatus(messageId, types.ExecutionStatus_Init_Refund_Executed)
				return err
			}
			// increase retryCount
			retryCount := Dal.IncreaseRetryCount(messageId)
			// either revert its status or set it to failed due to hitting retry limit
			if retryCount > types.MaxExecuteRetry {
				e := Dal.UpdateStatus(messageId, types.ExecutionStatus_Init_Refund_Failed)
				if e != nil {
					log.Errorf("cannot update message (id %x) status: %s", messageId, e.Error())
				}
			} else {
				e := Dal.RevertStatus(messageId, types.ExecutionStatus_Unexecuted)
				if e != nil {
					log.Errorf("cannot revert message (id %x) status: %s", messageId, e.Error())
					return err
				}
			}
			return err
		}
		log.Infof("executed refund init (messageId %x): txhash %x", messageId, tx.Hash())
		return nil
	}
}

func (c *Chain) ExecuteLiqWithdraw(
	opts *bind.TransactOpts, wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return c.LiqBridge.Withdraw(opts, wdOnchain, sortedSigs, signers, powers)
}

func (c *Chain) ExecutePegWithdraw(
	opts *bind.TransactOpts, wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return c.PegVault.OriginalTokenVault.Withdraw(opts, wdOnchain, sortedSigs, signers, powers)
}

func (c *Chain) ExecutePegV2Withdraw(
	opts *bind.TransactOpts, wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return c.PegVaultV2.OriginalTokenVaultV2.Withdraw(opts, wdOnchain, sortedSigs, signers, powers)
}

func (c *Chain) ExecutePegMint(
	opts *bind.TransactOpts, mintOnChain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return c.PegBridge.PeggedTokenBridge.Mint(opts, mintOnChain, sortedSigs, signers, powers)
}

func (c *Chain) ExecutePegV2Mint(
	opts *bind.TransactOpts, mintOnChain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return c.PegBridgeV2.PeggedTokenBridgeV2.Mint(opts, mintOnChain, sortedSigs, signers, powers)
}

func (c *Chain) GetTransfer(xferId eth.Hash) (bool, error) {
	return c.LiqBridge.Transfers(&bind.CallOpts{}, xferId)
}
