package gatewaysvc

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	farmingtypes "github.com/celer-network/sgn-v2/x/farming/types"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/fee"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/transactor"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/lthibault/jitterbug"
	"github.com/spf13/viper"
)

var (
	SelfStart         bool
	RootDir           string
	LegacyAmino       *codec.LegacyAmino
	Cdc               codec.Codec
	InterfaceRegistry codectypes.InterfaceRegistry
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

type GatewayConfig struct {
}

type GatewayService struct {
	F  *fee.TokenPriceCache
	TP *transactor.TransactorPool
	EC map[uint64]*ethclient.Client
}

func NewGatewayService(dbUrl string) (*GatewayService, error) {
	if SelfStart {
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
		config.SetBech32PrefixForValidator(common.Bech32PrefixValAddr, common.Bech32PrefixValPub)
		config.SetBech32PrefixForConsensusNode(common.Bech32PrefixConsAddr, common.Bech32PrefixConsPub)
		config.Seal()
	}
	// Make a private config copy.
	_db, err := dal.NewDAL("postgres", fmt.Sprintf("postgresql://root@%s/gateway?sslmode=disable", dbUrl), 10)
	if err != nil {
		return nil, err
	}

	dal.DB = _db
	gateway := &GatewayService{}

	return gateway, nil
}

// StartTokenPricePolling starts a loop with the given interval and 3s stdev for polling price
func (gs *GatewayService) StartChainTokenPolling(interval time.Duration) {
	gs.pollChainToken() // make sure run at least once before return
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
		}
	}()
}

func (gs *GatewayService) pollChainToken() {
	tr := gs.TP.GetTransactor()
	resp, err := cbrcli.QueryChainTokensConfig(tr.CliCtx, &cbrtypes.ChainTokensConfigRequest{})
	if err != nil {
		log.Errorln("we will use mocked chain tokens failed to load basic token info:", err)
	}
	chainTokens := resp.GetChainTokens()
	for chainIdStr, tokens := range chainTokens {
		chainId, convErr := strconv.Atoi(chainIdStr)
		if convErr != nil {
			log.Errorf("error chain id found:%s", chainIdStr)
			continue
		}
		for _, token := range tokens.Tokens {
			dbErr := dal.DB.UpsertTokenBaseInfo(token.GetSymbol(), common.Hex2Addr(token.GetAddress()).String(), uint64(chainId), uint64(token.GetDecimal()), token.GetXferDisabled())
			if dbErr != nil {
				log.Errorf("failed to write token: %v", dbErr)
			}
		}
		blockDelay := tokens.GetBlockDelay()
		dbErr := dal.DB.UpsertChainBaseInfo(uint64(chainId), blockDelay, common.Hex2Addr(tokens.GetContractAddr()).String())
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
					common.Hex2Addr(erc20Token.GetAddress()).String(),
					erc20Token.GetChainId(),
					uint64(erc20Token.GetDecimals()),
				)
				if dbErr != nil {
					log.Errorf("UpsertTokenBaseInfo error:%+v", dbErr)
				}
			}
		}
	}
}

func (gs *GatewayService) InitTransactors() error {
	if SelfStart {
		cbrCfgFile := filepath.Join(RootDir, "config", "cbridge.toml")
		viper.SetConfigFile(cbrCfgFile)
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read in cbridge configuration: %w", err)
		}
		configFilePath := filepath.Join(RootDir, "config", "sgn.toml")
		viper.SetConfigFile(configFilePath)
		if err := viper.MergeInConfig(); err != nil {
			return fmt.Errorf("failed to read in SGN configuration: %w", err)
		}
	}

	tp := transactor.NewTransactorPool(RootDir, viper.GetString(common.FlagSgnChainId), LegacyAmino, Cdc, InterfaceRegistry)
	err := tp.AddTransactors(
		viper.GetString(common.FlagSgnNodeURI), viper.GetString(common.FlagSgnPassphrase), viper.GetStringSlice(common.FlagSgnTransactors))
	if err != nil {
		return fmt.Errorf("failed to add transactors: %w", err)
	}
	gs.TP = tp

	var mcc []*common.OneChainConfig
	err = viper.UnmarshalKey(common.FlagMultiChain, &mcc)
	if err != nil {
		return fmt.Errorf("failed to new mcc: %w", err)
	}
	e := make(map[uint64]*ethclient.Client)
	for _, m := range mcc {
		ec, ecErr := ethclient.Dial(m.Gateway)
		if ecErr == nil {
			e[m.ChainID] = ec
		}
	}
	gs.EC = e
	return nil
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
