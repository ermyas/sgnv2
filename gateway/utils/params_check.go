package utils

import (
	"regexp"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

func CheckMarkTransferParams(transferId, txHash, addr string, sendInfo, receivedInfo *webapi.TransferInfo, txType webapi.TransferType) bool {
	isValidIndex := isValidHash(transferId) &&
		isValidHash(txHash)
	if txType == webapi.TransferType_TRANSFER_TYPE_REFUND {
		return isValidIndex
	} else {
		return isValidIndex &&
			isValidAddr(addr) &&
			isValidTxInfo(sendInfo) &&
			isValidTxInfo(receivedInfo) &&
			sendInfo.GetChain().GetId() != receivedInfo.GetChain().GetId()
	}
}

func CheckMarkLiquidityParams(lpType webapi.LPType, chainId uint32, amt, lpAddr, tokenAddr string) bool {
	return lpType != webapi.LPType_LP_TYPE_UNKNOWN &&
		IsvalidAmt(amt) &&
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
	return IsvalidAmt(info.GetAmount()) &&
		info.GetChain().GetId() > 0 &&
		info.GetToken().GetSymbol() != "" &&
		isValidAddr(info.GetToken().GetAddress())
}

func IsvalidAmt(amt string) bool {
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(amt) {
		// isNumeric
		return false
	}
	if len(amt) > 30 {
		// large than billion
		return false
	}
	if amt == "" || amt == "0" {
		// invalid value
		return false
	}
	return true
}
