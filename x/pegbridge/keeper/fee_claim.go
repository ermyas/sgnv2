package keeper

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) claimFee(
	ctx sdk.Context, delAddr eth.Addr, chainId uint64, token eth.Addr, nonce uint64, sender string) (
	amount *big.Int, withdrawOnChain *types.WithdrawOnChain, err error) {
	// 1. Claim pegbridge fees in distribution module
	err = k.distrKeeper.ClaimPegBridgeFees(ctx, delAddr)
	if err != nil {
		return nil, nil, err
	}
	logmsg := fmt.Sprintf("claimFee:%x nonce:%d", delAddr, nonce)
	// 2. Take the fee balance in the distribution module and generate a WithdrawOnchain
	var symbol string
	k.IterateOrigPeggedPairsByOrig(ctx, chainId, token, func(pair types.OrigPeggedPair) bool {
		symbol = pair.Orig.Symbol
		return true
	})
	denom := fmt.Sprintf("%s%s/%d", types.PegBridgeFeeDenomPrefix, symbol, chainId)
	coin := k.distrKeeper.GetWithdrawableBalance(ctx, delAddr, denom)
	err = k.BurnFee(ctx, delAddr, coin)
	if err != nil {
		return nil, nil, err
	}
	amount = coin.Amount.BigInt()
	logmsg = fmt.Sprintf("%s fee_token:%x, amt:%s", logmsg, token, amount)
	log.Infof("x/pegbridge handle fee claim: %s sender:%s", logmsg, sender)

	return amount, &types.WithdrawOnChain{
		Token:       token.Bytes(),
		Receiver:    delAddr.Bytes(),
		Amount:      amount.Bytes(),
		BurnAccount: eth.ZeroAddr[:], // N/A
		RefChainId:  0,               // A ref_chain_id of 0 indicates a fee claim
		RefId:       new(big.Int).SetUint64(nonce).FillBytes(make([]byte, 32)),
	}, nil
}
