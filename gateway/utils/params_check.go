package utils

import (
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	"strconv"
)

func CheckMarkTransferParams(transferId, txHash, addr string, sendInfo, receivedInfo *webapi.TransferInfo) bool {
	return isValidHash(transferId) &&
		isValidHash(txHash) &&
		isValidAddr(addr) &&
		isValidTxInfo(sendInfo) &&
		isValidTxInfo(receivedInfo) &&
		sendInfo.GetChain().GetId() != receivedInfo.GetChain().GetId()
}

func CheckMarkLiquidityParams(lpType webapi.LPType, chainId uint32, amt, lpAddr, tokenAddr string) bool {
	amtNum, err := strconv.Atoi(amt)
	return lpType != webapi.LPType_LP_TYPE_UNKNOWN &&
		amtNum > 0 && err == nil &&
		chainId > 0 &&
		isValidAddr(lpAddr) &&
		isValidAddr(tokenAddr)
}

func CheckWithdrawLiquidityParams(req *types.WithdrawReq) bool {
	if req.GetXferId() != "" {
		return isValidHash(req.GetXferId())
	} else {
		return req.GetReqId() > 0 && len(req.GetWithdraws()) > 0 && req.GetExitChainId() > 0
	}
}

func CheckUnlockFarmingRewardParams(addr string) bool {
	return isValidAddr(addr)
}

func CheckUnlockStakingRewardParams(addr string) bool {
	return isValidAddr(addr)
}

func isValidAddr(addr string) bool {
	return common.IsHexAddress(addr)
}

func isValidHash(hash string) bool {
	return common.IsValidTxHash(hash)
}

func isValidTxInfo(info *webapi.TransferInfo) bool {
	amt, err := strconv.Atoi(info.GetAmount())
	return amt > 0 && err == nil &&
		info.GetChain().GetId() > 0 &&
		info.GetToken().GetSymbol() != "" &&
		isValidAddr(info.GetToken().GetAddress())
}
