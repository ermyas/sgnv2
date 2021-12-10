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
	"github.com/celer-network/sgn-v2/gateway/onchain"
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

func (gs *GatewayService) IsWithdrawNormal(addr, amt, tokenSymbol string, decimal int) bool {
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
	cmpDp := new(big.Float).Mul(d, new(big.Float).SetFloat64(1.05))
	if cmpWd.Cmp(cmpDp) > 0 {
		//Gateway should raise alert and block any withdrawal request that will make the total withdrawal more than 120% of the total deposit
		// alert
		wd, _ := w.Float64()
		dp, _ := d.Float64()
		dt, _ := withdrawAmt.Float64()
		utils.SendWithdrawAlert(addr, fmt.Sprintf("%.6f", wd), fmt.Sprintf("%.6f", dp), fmt.Sprintf("%.6f", dt), tokenSymbol)
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
			cmpIO := new(big.Float).Mul(dw.deposit, new(big.Float).SetFloat64(1.05))
			cmpBlc := new(big.Float).Add(balance, dw.withdraw)
			if cmpIO.Cmp(cmpBlc) < 0 {
				// total_deposit * 1.05 < (total_withdrawal + current_lp_balance)
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
		utils.SendBalanceAlert(alerts)
	}
}

func (gs *GatewayService) AlertAbnormalStatus() {
	endTime := time.Now().Add(-2 * time.Hour)
	startTime := endTime.Add(-2 * time.Hour)
	sendNoSync, err := dal.DB.GetTransfersWithStatus(cbrtypes.TransferHistoryStatus_TRANSFER_WAITING_FOR_SGN_CONFIRMATION, startTime, endTime)
	if err != nil {
		log.Warnf("GetTransfersWithStatus failed, from:%s, to:%s, err:%+v", startTime, endTime, err)
	}
	sendNoSync = gs.filterAlertByTransferStatus(sendNoSync, cbrtypes.TransferHistoryStatus_TRANSFER_WAITING_FOR_SGN_CONFIRMATION)
	if sendNoSync != nil && len(sendNoSync) > 0 {
		utils.SendStatusAlert(sendNoSync, "send events not synced to sgn")
	}
	fundNoRelease, err := dal.DB.GetTransfersWithStatus(cbrtypes.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE, startTime, endTime)
	if err != nil {
		log.Warnf("GetTransfersWithStatus failed, from:%s, to:%s, err:%+v", startTime, endTime, err)
	}
	fundNoRelease = gs.filterAlertByTransferStatus(fundNoRelease, cbrtypes.TransferHistoryStatus_TRANSFER_WAITING_FOR_FUND_RELEASE)
	if fundNoRelease != nil && len(fundNoRelease) > 0 {
		utils.SendStatusAlert(fundNoRelease, "relays that have been waiting for fund release")
	}
	addLiqNoSync, err := dal.DB.GetLPWithStatus(cbrtypes.WithdrawStatus_WD_WAITING_FOR_SGN, startTime, endTime)
	if err != nil {
		log.Warnf("GetLPWithStatus failed, from:%s, to:%s, err:%+v", startTime, endTime, err)
	}
	addLiqNoSync = gs.filterAlertByLPStatus(addLiqNoSync, cbrtypes.WithdrawStatus_WD_WAITING_FOR_SGN)
	if addLiqNoSync != nil && len(addLiqNoSync) > 0 {
		utils.SendStatusAlert(addLiqNoSync, "addLiq events not synced to sgn")
	}
}

func (gs *GatewayService) filterAlertByTransferStatus(alerts []*utils.StatusAlertInfo, status cbrtypes.TransferHistoryStatus) []*utils.StatusAlertInfo {
	var filteredAlerts []*utils.StatusAlertInfo
	if alerts == nil || len(alerts) == 0 {
		return filteredAlerts
	}
	for _, alert := range alerts {
		srcTxHash := alert.TxHash
		chainId := uint32(alert.ChainId)
		tx, found, err := dal.DB.GetTransferBySrcTxHash(srcTxHash, chainId)
		if tx == nil || !found || err != nil {
			log.Warnf("alert, tx not found for chainId:%d, srcTxHash:%s", chainId, srcTxHash)
			filteredAlerts = append(filteredAlerts, alert)
		} else {
			var transfers []*dal.Transfer
			transfers = append(transfers, tx)
			gs.updateTransferStatusInHistory(context.Background(), transfers)
			tx, _, _ = dal.DB.GetTransferBySrcTxHash(srcTxHash, chainId)
			if tx.Status == status {
				filteredAlerts = append(filteredAlerts, alert)
			}
		}

	}
	return filteredAlerts
}

func (gs *GatewayService) filterAlertByLPStatus(alerts []*utils.StatusAlertInfo, status cbrtypes.WithdrawStatus) []*utils.StatusAlertInfo {
	var filteredAlerts []*utils.StatusAlertInfo
	if alerts == nil || len(alerts) == 0 {
		return filteredAlerts
	}
	for _, alert := range alerts {
		srcTxHash := alert.TxHash
		chainId := alert.ChainId
		lp, found, err := dal.DB.GetOneLPInfoByHash(chainId, srcTxHash)
		if lp == nil || !found || err != nil {
			log.Warnf("alert, lp not found for chainId:%d, srcTxHash:%s", chainId, srcTxHash)
			filteredAlerts = append(filteredAlerts, alert)
		} else {
			var lps []*dal.LP
			lps = append(lps, lp)
			gs.updateLpStatusInHistory(lps)
			lp, _, _ = dal.DB.GetOneLPInfoByHash(chainId, srcTxHash)
			if lp.Status == status {
				filteredAlerts = append(filteredAlerts, alert)
				log.Infof("alert, update status from sgn failed, status:%s, chainId:%d, srcTxHash:%s，lp info:%+v", status, chainId, srcTxHash, lp)
			} else {
				log.Infof("alert, update status from sgn success, from status:%s, to status:%s, chainId:%d, srcTxHash:%s", status, lp.Status, chainId, srcTxHash)
			}
		}
	}
	return filteredAlerts
}

func (gs *GatewayService) getUsrBalance(usrAddr string, chainTokens []*cbrtypes.ChainTokenAddrPair, chainTokenAddrMap map[uint64]map[string]*webapi.TokenInfo) map[string]*big.Float {
	balanceMap := make(map[string]*big.Float)
	tr := onchain.SGNTransactors.GetTransactor()
	detailList, detailErr := cbrcli.QueryLiquidityDetailList(tr.CliCtx, &cbrtypes.LiquidityDetailListRequest{
		LpAddr:     usrAddr,
		ChainToken: chainTokens,
	})
	if detailList == nil || detailErr != nil || detailList.LiquidityDetail == nil {
		return balanceMap
	}

	for _, liq := range detailList.LiquidityDetail {
		token := chainTokenAddrMap[liq.GetChainId()][liq.GetToken().GetAddress()]
		if token == nil {
			continue
		}
		balance := balanceMap[token.GetToken().GetSymbol()]
		if balance == nil {
			balance = new(big.Float).SetInt64(0)
		}
		balance = new(big.Float).Add(balance, rmAmtDecimal(liq.GetUsrLiquidity(), int(token.GetToken().GetDecimal())))
		balanceMap[token.GetToken().GetSymbol()] = balance
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
	usrIOMap := make(map[string]map[string]*TotalIO)
	lps, err := dal.DB.AllLpAmtForBalance()
	if err != nil {
		log.Errorf("AllLpAmtForBalance db error, err:%+v", err)
		return usrIOMap
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
	return usrIOMap
}

// return withdraw and deposit
func getUsrWithdrawAndDeposit(addr string) *TotalIO {
	lpHistory, err := dal.DB.GetAllLpHistoryForBalance(addr)
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
