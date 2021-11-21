package gatewaysvc

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/utils"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	cbrcli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
)

func (gs *GatewayService) GetTotalLiquidityProviderTokenBalance(ctx context.Context, request *webapi.GetTotalLiquidityProviderTokenBalanceRequest) (*webapi.GetTotalLiquidityProviderTokenBalanceResponse, error) {
	tokenSymbol := request.GetTokenSymbol()
	chainIds := request.GetChainIds()
	ret := make(map[uint64]string)
	if len(chainIds) == 0 {
		// all chains
		chainTokenInfos, err := dal.DB.GetChainTokenList()
		if err == nil && len(chainTokenInfos) > 0 {
			for chainId, chainToken := range chainTokenInfos {
				for _, token := range chainToken.GetToken() {
					if token.GetToken().GetSymbol() == tokenSymbol {
						ret[uint64(chainId)] = gs.getLiquidityOnChainToken(uint64(chainId), token.GetToken().GetAddress())
						break
					}
				}
			}
		}
	} else {
		for _, chainId := range chainIds {
			token, tokenFound, dberr := dal.DB.GetTokenBySymbol(tokenSymbol, uint64(chainId))
			if tokenFound && dberr == nil {
				ret[uint64(chainId)] = gs.getLiquidityOnChainToken(uint64(chainId), token.GetToken().GetAddress())
			}
		}
	}
	return &webapi.GetTotalLiquidityProviderTokenBalanceResponse{
		TotalLiq: ret,
	}, nil
}

func (gs *GatewayService) IsWithdrawNormal(addr, amt string, decimal int) bool {
	usrWithdrawAndDeposit := getUsrWithdrawAndDeposit(addr)
	withdrawAmt := rmAmtDecimal(amt, decimal)

	if usrWithdrawAndDeposit == nil {
		usrWithdrawAndDeposit = &TotalIO{}
	}
	w := usrWithdrawAndDeposit.withdraw
	d := usrWithdrawAndDeposit.deposit
	if w == nil {
		w = new(big.Float).SetInt64(0)
	}
	if d == nil {
		d = new(big.Float).SetInt64(0)
	}
	// cmp with amt added and get bool result
	cmpWd := new(big.Float).Add(w, withdrawAmt)
	cmpDp := new(big.Float).Mul(d, new(big.Float).SetFloat64(1.2))
	if cmpWd.Cmp(cmpDp) > 0 {
		//Gateway should raise alert and block any withdrawal request that will make the total withdrawal more than 120% of the total deposit
		// alert
		wd, _ := w.Float64()
		dp, _ := d.Float64()
		dt, _ := withdrawAmt.Float64()
		utils.SendWithdrawAlert(addr, fmt.Sprintf("%.6f", wd), fmt.Sprintf("%.6f", dp), fmt.Sprintf("%.6f", dt), Env)
		return false
	}
	return true
}

func (gs *GatewayService) AlertAbnormalBalance() {
	allDepositAndWithdraw := getTotalWithdrawAndDeposit()
	// cli: get balance
	chainTokens, tokenMap, err := getChainTokens()
	if err != nil {
		log.Errorf("getChainTokens failed, err:%+v", err)
	}
	var alerts []*utils.BalanceAlert
	for usrAddr, dwTokenMap := range allDepositAndWithdraw {
		tokenBalance := gs.getUsrBalance(usrAddr, chainTokens, tokenMap)
		for tokenSymbol, dw := range dwTokenMap {
			// cmp and alert
			balance := tokenBalance[tokenSymbol]
			if balance == nil {
				balance = new(big.Float).SetInt64(0)
			}
			cmpIO := dw.deposit
			cmpBlc := new(big.Float).Mul(new(big.Float).Add(balance, dw.withdraw), new(big.Float).SetFloat64(0.95))
			if cmpIO.Cmp(cmpBlc) > 0 {
				// total_deposit > 0.95 * (total_withdrawal + current_lp_balance)
				wd, _ := dw.withdraw.Float64()
				dp, _ := dw.deposit.Float64()
				blc, _ := balance.Float64()
				alerts = append(alerts, &utils.BalanceAlert{
					Token:    tokenSymbol,
					Balance:  fmt.Sprintf("%.6f", blc),
					Addr:     usrAddr,
					Withdraw: fmt.Sprintf("%.6f", wd),
					Deposit:  fmt.Sprintf("%.6f", dp),
				})
			}
		}
	}
	if alerts != nil && len(alerts) > 0 {
		utils.SendBalanceAlert(alerts, Env)
	}
}

func (gs *GatewayService) getUsrBalance(usrAddr string, chainTokens []*cbrtypes.ChainTokenAddrPair, chainTokenAddrMap map[uint64]map[string]*webapi.TokenInfo) map[string]*big.Float {
	balanceMap := make(map[string]*big.Float)
	tr := gs.TP.GetTransactor()
	detailList, detailErr := cbrcli.QueryLiquidityDetailList(tr.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     usrAddr,
		ChainToken: chainTokens,
	})
	if detailList == nil || detailErr != nil || detailList.LiquidityDetail == nil {
		return balanceMap
	}

	for _, liq := range detailList.LiquidityDetail {
		balance := balanceMap[liq.GetToken().GetSymbol()]
		if balance == nil {
			balance = new(big.Float).SetInt64(0)
		}
		decimal := chainTokenAddrMap[liq.GetChainId()][liq.GetToken().GetAddress()].GetToken().GetDecimal()
		balance = new(big.Float).Add(balance, rmAmtDecimal(liq.GetUsrLiquidity(), int(decimal)))
		balanceMap[liq.GetToken().GetSymbol()] = balance
	}
	return balanceMap
}

func getChainTokens() ([]*cbrtypes.ChainTokenAddrPair, map[uint64]map[string]*webapi.TokenInfo, error) {
	chainTokenInfos, err := dal.DB.GetChainTokenList()
	if err != nil {
		return nil, nil, err
	}
	var chainTokens []*cbrtypes.ChainTokenAddrPair
	for chainId, tokens := range chainTokenInfos {
		for _, tokenInfo := range tokens.Token {
			chainTokens = append(chainTokens, &cbrtypes.ChainTokenAddrPair{
				ChainId:   uint64(chainId),
				TokenAddr: tokenInfo.GetToken().Address,
			})
		}
	}
	chainTokenAddrMap := make(map[uint64]map[string]*webapi.TokenInfo)
	for chainId, tokens := range chainTokenInfos {
		tokenAddrMap, chainFound := chainTokenAddrMap[uint64(chainId)]
		if !chainFound {
			tokenAddrMap = make(map[string]*webapi.TokenInfo)
		}
		for _, token := range tokens.Token {
			tokenAddrMap[token.GetToken().GetAddress()] = token
		}
		chainTokenAddrMap[uint64(chainId)] = tokenAddrMap
	}
	return chainTokens, chainTokenAddrMap, nil
}

type TotalIO struct {
	withdraw *big.Float
	deposit  *big.Float
}

// return map[addr][token_symbol]totalIO
func getTotalWithdrawAndDeposit() map[string]map[string]*TotalIO {
	pageSize := uint64(5000)
	end := time.Now()
	hasNextPage := true
	usrIOMap := make(map[string]map[string]*TotalIO)
	for hasNextPage {
		lps, size, nextTime, err := dal.DB.PaginateLpAmt(end, pageSize)
		if err != nil {
			log.Errorf("PaginateLpAmt error, end:%s, pageSize:%d, err:%+v", end.String(), pageSize, err)
			break
		}
		for _, entry := range lps {
			lpType := entry.LpType
			addr := entry.Addr
			tokenSymbol := entry.TokenSymbol
			if usrIOMap[addr] == nil {
				usrIOMap[addr] = make(map[string]*TotalIO)
			}
			if usrIOMap[addr][tokenSymbol] == nil {
				usrIOMap[addr][tokenSymbol] = &TotalIO{
					withdraw: new(big.Float).SetInt64(0),
					deposit:  new(big.Float).SetInt64(0),
				}
			}

			if lpType == webapi.LPType_LP_TYPE_REMOVE {
				usrIOMap[addr][entry.TokenSymbol].withdraw = new(big.Float).Add(usrIOMap[addr][entry.TokenSymbol].withdraw, getAmtFromLpHistory(entry))
			} else if lpType == webapi.LPType_LP_TYPE_ADD {
				usrIOMap[addr][entry.TokenSymbol].deposit = new(big.Float).Add(usrIOMap[addr][entry.TokenSymbol].deposit, getAmtFromLpHistory(entry))
			}
		}
		end = nextTime
		if uint64(size) != pageSize {
			hasNextPage = false
		}
	}
	return usrIOMap
}

// return withdraw and deposit
func getUsrWithdrawAndDeposit(addr string) *TotalIO {
	lpHistory, err := dal.DB.GetAllLpHistory(addr)
	if err != nil {
		log.Warnf("GetAllLpHistory err:%+v", err)
	}
	amtMap := make(map[int32]*big.Float)
	for _, entry := range lpHistory {
		lpType := entry.LpType
		preValue, found := amtMap[int32(lpType)]
		if !found {
			preValue = new(big.Float)
		}
		amtMap[int32(lpType)] = new(big.Float).Add(preValue, getAmtFromLpHistory(entry))
	}
	return &TotalIO{
		withdraw: amtMap[int32(webapi.LPType_LP_TYPE_REMOVE)],
		deposit:  amtMap[int32(webapi.LPType_LP_TYPE_ADD)],
	}
}

func getAmtFromLpHistory(entry *dal.LP) *big.Float {
	chainId := entry.ChainId
	tokenSymbol := entry.TokenSymbol
	amt := entry.Amt
	token, found, tokenErr := dal.DB.GetTokenBySymbol(tokenSymbol, chainId)
	if !found || tokenErr != nil {
		return new(big.Float).SetInt64(0)
	}
	return rmAmtDecimal(amt, int(token.GetToken().GetDecimal()))
}

func rmAmtDecimal(amt string, decimal int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(common.Str2BigInt(amt)), big.NewFloat(math.Pow10(decimal)))
}
