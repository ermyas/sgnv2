package gatewaysvc

import (
	"context"
	"fmt"
	"strconv"
	"time"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/spf13/viper"

	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/onchain"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	pegcli "github.com/celer-network/sgn-v2/x/pegbridge/client/cli"
	pegtypes "github.com/celer-network/sgn-v2/x/pegbridge/types"
	"github.com/lthibault/jitterbug"
)

// ================================= new gateway method below =====================================

// Close the database DAL.
func (gs *GatewayService) Close() {
	if dal.DB == nil {
		return
	}
	dal.DB.Close()
	dal.DB = nil
}

type ValidatorChainConnectivity struct {
	// validator addr -> current block number
	C map[eth.Addr]*webapi.CurrentBlockNumberReport
}

type IncentiveRewardsSigner struct {
	Signer *ethutils.Signer
	Addr   *eth.Addr
}

type GatewayService struct {
	F      *TokenPriceCache
	Chains onchain.ChainMgr
	S      *IncentiveRewardsSigner
	V      *ValidatorChainConnectivity
}

func NewGatewayService(db *dal.DAL) *GatewayService {
	// Make a private config copy.
	dal.DB = db
	gs := &GatewayService{}
	signerKey, signerPass := viper.GetString(common.FlagGatewayIncentiveRewardsKeystore), viper.GetString(common.FlagGatewayIncentiveRewardsPassphrase)
	signer, addr, err := eth.CreateSigner(signerKey, signerPass, nil)
	if err != nil {
		log.Fatalf("fail to CreateSigner in gateway server, err:%v", err)
	}
	gs.S = &IncentiveRewardsSigner{
		Signer: &signer,
		Addr:   &addr,
	}
	gs.V = &ValidatorChainConnectivity{
		C: make(map[eth.Addr]*webapi.CurrentBlockNumberReport),
	}
	return gs
}

// StartTokenPricePolling starts a loop with the given interval and 3s stdev for polling price
func (gs *GatewayService) StartChainTokenPolling(interval time.Duration) {
	gs.pollChainToken() // make sure run at least once before return
	go gs.makeSureHasChainToken()
	polledInside := false
	go func() {
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			if polledInside {
				gs.pollChainToken()
			}
			polledInside = true
		}
	}()
	log.Infof("chain token cached")
}
func (gs *GatewayService) makeSureHasChainToken() {
	list, _ := dal.DB.GetChainTokenList()
	pollingNeed := list == nil || len(list) == 0
	attempts := 10
	for pollingNeed {
		time.Sleep(3 * time.Second)
		gs.pollChainToken() // make sure run at least once before return
		list, _ = dal.DB.GetChainTokenList()
		attempts--
		pollingNeed = (list == nil || len(list) == 0) && attempts > 0
	}
}

// StartAvgLpFeeEarningPolling starts a loop with the given interval and 3s stdev for polling avg apy
func (gs *GatewayService) StartAvgLpFeeEarningPolling(interval time.Duration) {
	go func() {
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			gs.setAvgLpFeeEarningApy()
		}
	}()
}

// StartAbnormalBalanceCheckPolling starts a loop with the given interval and 3s stdev for polling avg apy
func (gs *GatewayService) StartAbnormalBalanceCheckPolling(interval time.Duration) {
	go func() {
		ticker := jitterbug.New(
			interval,
			&jitterbug.Norm{Stdev: 3 * time.Second},
		)
		defer ticker.Stop()
		for ; true; <-ticker.C {
			gs.AlertAbnormalBalance()
			gs.AlertAbnormalStatus()
		}
	}()
}

func (gs *GatewayService) pollChainToken() {
	tr := onchain.SGNTransactors.GetTransactor()
	resp, err := cbrcli.QueryChainTokensConfig(tr.CliCtx, &cbrtypes.ChainTokensConfigRequest{})
	if err != nil {
		log.Warnln("we will use cached chain tokens failed to load basic token info:", err)
	}
	chainTokens := resp.GetChainTokens()
	for chainIdStr, tokens := range chainTokens {
		chainId, convErr := strconv.Atoi(chainIdStr)
		if convErr != nil {
			log.Errorf("error chain id found:%s", chainIdStr)
			continue
		}
		for _, token := range tokens.Tokens {
			dbErr := dal.DB.UpsertTokenBaseInfo(token.GetSymbol(), eth.Hex2Addr(token.GetAddress()).String(), uint64(chainId), uint64(token.GetDecimal()), token.GetXferDisabled())
			if dbErr != nil {
				log.Errorf("failed to write token: %v", dbErr)
			}
		}
		blockDelay := tokens.GetBlockDelay()
		dbErr := dal.DB.UpsertChainBaseInfo(uint64(chainId), blockDelay, eth.Hex2Addr(tokens.GetContractAddr()).String())
		if dbErr != nil {
			log.Errorf("failed to write blockDelay: %v", dbErr)
		}
	}

	queryClient := farmingtypes.NewQueryClient(tr.CliCtx)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()

	farmingPools, err := queryClient.Pools(
		ctx,
		&farmingtypes.QueryPoolsRequest{},
	)
	if err != nil {
		log.Errorln("query pools err", err)
	}
	if farmingPools != nil {
		for _, pool := range farmingPools.GetPools() {
			for _, erc20Token := range pool.GetRewardTokens() {
				tokenSymbol := cbrtypes.GetSymbolFromStakeToken(erc20Token.GetSymbol())
				dbErr := dal.DB.UpsertRewardToken(
					tokenSymbol,
					eth.Hex2Addr(erc20Token.GetAddress()).String(),
					erc20Token.GetChainId(),
					uint64(erc20Token.GetDecimals()),
				)
				if dbErr != nil {
					log.Errorf("UpsertTokenBaseInfo error:%+v", dbErr)
				}
			}
		}
	}

	// pegged token
	log.Infof("start QueryAllPeggedPairs")
	pegPairs, err := pegcli.QueryOrigPeggedPairs(tr.CliCtx, &pegtypes.QueryOrigPeggedPairsRequest{})
	if err != nil {
		log.Warnln("query QueryAllPeggedPairs err", err)
	} else {
		for _, pair := range pegPairs {
			org := pair.GetOrig()
			pegged := pair.GetPegged()
			dbErr := dal.DB.InsertMintTokenBaseInfo(org.GetSymbol(), eth.Hex2Addr(org.GetAddress()).String(), org.GetChainId(), uint64(org.GetDecimals()))
			if dbErr != nil {
				log.Errorf("fail to save peg org token, dbErr:%s", dbErr.Error())
				continue
			}
			dbErr = dal.DB.InsertMintTokenBaseInfo(pegged.GetSymbol(), eth.Hex2Addr(pegged.GetAddress()).String(), pegged.GetChainId(), uint64(pegged.GetDecimals()))
			if dbErr != nil {
				log.Errorf("fail to save pegged token, dbErr:%s", dbErr.Error())
				continue
			}
			dbErr = dal.DB.InsertPeggedBaseInfo(&org, &pegged)
			if dbErr != nil {
				log.Errorf("failed to InsertPeggedBaseInfo: %v", dbErr)
				continue
			}
		}
	}
}

func (gs *GatewayService) GetAllValidPeggedPairs() ([]*webapi.PeggedPairConfig, error) {
	configs, dbErr := dal.DB.GetAllValidPeggedConfigList()
	if dbErr != nil {
		return nil, dbErr
	}
	log.Infof("GetAllValidPeggedPairs configs:%+v", configs)
	for _, c := range configs {
		srcChain, foundSrc := gs.Chains.GetOneChain(uint64(c.OrgChainId))
		peggedChain, foundPegged := gs.Chains.GetOneChain(uint64(c.PeggedChainId))
		if !foundSrc || !foundPegged || srcChain.GetOtvContract() == nil || peggedChain.GetPtbContract() == nil {
			log.Errorf("fail to find this pegged chain pair in onchain config: %+v", c)
			continue
		}
		c.PeggedDepositContractAddr = srcChain.GetOtvContract().GetAddr().String()
		c.PeggedBurnContractAddr = peggedChain.GetPtbContract().GetAddr().String()
	}
	log.Infof("GetAllValidPeggedPairs res:%+v", configs)
	return configs, nil
}

// ================================= common method below =====================================
func unknownChain(chainId uint32) *webapi.Chain {
	return &webapi.Chain{
		Id:   chainId,
		Name: fmt.Sprintf("Chain-%d", chainId),
		Icon: "https://cbridge.celer.network/ETH.png",
	}
}

func enrichChainUiInfo(chain *webapi.Chain) *webapi.Chain {
	if chain.GetName() == "" {
		if chain.GetId() > 0 {
			chain.Name = fmt.Sprintf("Chain-%d", chain.Id)
		} else {
			chain.Name = "New Added Chain"
		}
	}
	if chain.GetIcon() == "" {
		chain.Icon = "https://cbridge.celer.network/ETH.png"
	}
	return chain
}

func enrichUnknownToken(token *webapi.TokenInfo) {
	if token.GetName() == "" {
		token.Name = token.Token.GetSymbol()
	}
	if token.GetIcon() == "" {
		token.Icon = "https://get.celer.app/cbridge-icons/ETH.png"
	}
}
