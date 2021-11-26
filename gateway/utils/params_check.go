package utils

import (
	"regexp"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
)

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

func IsValidAmt(amt string) bool {
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
