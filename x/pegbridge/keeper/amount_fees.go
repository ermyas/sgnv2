package keeper

import (
	"fmt"
	"math/big"
	"time"

	"github.com/celer-network/sgn-v2/common"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/relayer"
	distrtypes "github.com/celer-network/sgn-v2/x/distribution/types"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CalcAmountAndFees calculates the mint / withdraw amount after taking into account:
// 1. Potential difference in decimals
// 2. Base fee
// 3. Percentage fee, potentially different depending on mint / burn
func (k Keeper) CalcAmountAndFees(
	ctx sdk.Context, pair types.OrigPeggedPair, requestAmount *big.Int, isPeggedDest bool) (
	receiveAmount *big.Int, baseFeeInOrig *big.Int, percFeeInOrig *big.Int) {
	// Apply decimals diff
	var srcDecimals, destDecimals uint32
	if isPeggedDest {
		srcDecimals = pair.Orig.Decimals
		destDecimals = pair.Pegged.Decimals
	} else {
		srcDecimals = pair.Pegged.Decimals
		destDecimals = pair.Orig.Decimals
	}
	destAmt := new(big.Int).Set(requestAmount)
	if destDecimals > srcDecimals {
		destAmt.Mul(destAmt, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(destDecimals-srcDecimals)), nil))
	} else if destDecimals < srcDecimals {
		// NOTE: Rounds down as intended
		destAmt.Div(destAmt, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(srcDecimals-destDecimals)), nil))
	}
	// Base fee, denominated in ORIG decimals
	baseFeeInOrig = k.CalcBaseFee(ctx, pair, isPeggedDest)
	// NOTE: For amount calculation, we must scale base fee to DEST decimals
	baseFeeInDest := new(big.Int).Set(baseFeeInOrig)
	if isPeggedDest {
		if destDecimals > srcDecimals {
			baseFeeInDest.Div(baseFeeInDest, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(destDecimals-srcDecimals)), nil))
		} else if destDecimals < srcDecimals {
			baseFeeInDest.Mul(baseFeeInDest, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(srcDecimals-destDecimals)), nil))
		}
	}
	// Percentage fee, denominated in DEST decimals
	percFeeInDest := k.CalcPercFee(ctx, pair, destAmt, isPeggedDest)
	destAmtMinusBase := destAmt.Sub(destAmt, baseFeeInDest)
	destAmt = destAmtMinusBase.Sub(destAmtMinusBase, percFeeInDest)
	// NOTE: Now we must scale percentage fee to ORIG decimals
	percFeeInOrig = new(big.Int).Set(percFeeInDest)
	if isPeggedDest {
		if destDecimals > srcDecimals {
			percFeeInOrig.Div(percFeeInOrig, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(destDecimals-srcDecimals)), nil))
		} else if destDecimals < srcDecimals {
			percFeeInOrig.Mul(percFeeInOrig, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(srcDecimals-destDecimals)), nil))
		}
	}
	return destAmt, baseFeeInOrig, percFeeInOrig
}

// CalcBaseFee calculates the base fee. The base fee is a function of:
// 1. Asset price
// 2. Destination chain gas token price,
// 3. Destination chain gas price
// 4. Mint / withdraw gas cost
// NOTE: For simplicity, we approximate pegbridge mint / withdraw gas cost by directly using
// cbridge relay gas cost as the difference is small. Operation wise this requires that the original asset
// MUST BE part of the asset info in the cbridge module. Transfer can be disabled if we want to restrict the asset to
// mint / burn mode.
func (k Keeper) CalcBaseFee(ctx sdk.Context, pair types.OrigPeggedPair, isPeggedDest bool) (baseFee *big.Int) {
	var destChainId uint64
	if isPeggedDest {
		destChainId = pair.Pegged.ChainId
	} else {
		destChainId = pair.Orig.ChainId
	}
	// NOTE: Always use orig as the pegged symbol might be different across different
	// pegged versions of the same asset, and the fees are claimed in the form of orig tokens anyway.
	return k.cbridgeKeeper.CalcBaseFee(ctx, pair.Orig.Symbol, pair.Orig.ChainId, destChainId)
}

// CalcPercFee calculates the percentage fee based on the destination amount.
func (k Keeper) CalcPercFee(
	ctx sdk.Context, pair types.OrigPeggedPair, destAmount *big.Int, isPeggedDest bool) (percFee *big.Int) {
	var percFeePips uint32
	var maxPercFeeAmtStr string
	if isPeggedDest {
		percFeePips = pair.MintFeePips
		maxPercFeeAmtStr = pair.MaxMintFee
	} else {
		percFeePips = pair.BurnFeePips
		maxPercFeeAmtStr = pair.MaxBurnFee
	}
	if percFeePips == 0 {
		return new(big.Int)
	}
	percFee = new(big.Int).Mul(destAmount, big.NewInt(int64(percFeePips)))
	percFee.Div(percFee, big.NewInt(1e6))
	// Cap to max fee amount for destination chain
	maxPercFee, valid := new(big.Int).SetString(maxPercFeeAmtStr, 10)
	if !valid {
		return new(big.Int)
	}
	if percFee.Cmp(maxPercFee) > 0 {
		return maxPercFee
	}
	return percFee
}

// MintFeeAndSendToSyncer mints base fee to distribution module's fee collector account, then sends it to the active syncer directly.
func (k Keeper) MintFeeAndSendToSyncer(ctx sdk.Context, token commontypes.ERC20Token, amount *big.Int,
	srcChainId uint64, dstChainId uint64) error {
	coin, err := k.MintFee(ctx, token, amount)
	if err != nil {
		return err
	}
	syncer := k.stakingKeeper.GetSyncer(ctx)
	// Send coins from module to the active syncer address directly, bypassing distribution mechanism.
	syncerAddr := eth.Hex2Addr(syncer.GetEthAddress())
	derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(distrtypes.ModuleName, syncerAddr)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, k.feeCollectorName, derivedAccAddress, sdk.NewCoins(coin))
	if err != nil {
		return err
	}
	relayer.ReportBaseFeeDistribution(relayer.BridgeType_BRIDGE_TYPE_PEGGED, syncerAddr, time.Now(), amount, token.GetSymbol(), token.GetDecimals(), srcChainId, dstChainId)
	return nil
}

// MintFee mints fee to distribution module's fee collector account.
func (k Keeper) MintFee(ctx sdk.Context, token commontypes.ERC20Token, amount *big.Int) (coin sdk.Coin, err error) {
	denom := fmt.Sprintf("%s%s/%d", types.PegBridgeFeeDenomPrefix, token.Symbol, token.ChainId)
	coin = sdk.NewCoin(denom, sdk.NewIntFromBigInt(amount))
	if err = k.bankKeeper.MintCoins(ctx, k.feeCollectorName, sdk.NewCoins(coin)); err != nil {
		return coin, err
	}
	return coin, nil
}

// BurnFee burns coins from distribution module's fee collector account.
func (k Keeper) BurnFee(ctx sdk.Context, addr eth.Addr, coin sdk.Coin) error {
	// Send coins from address to module.
	derivedAccAddress := common.DeriveSdkAccAddressFromEthAddress(distrtypes.ModuleName, addr)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, derivedAccAddress, k.feeCollectorName, sdk.NewCoins(coin))
	if err != nil {
		return err
	}
	// Burn coins.
	err = k.bankKeeper.BurnCoins(ctx, k.feeCollectorName, sdk.NewCoins(coin))
	if err != nil {
		return err
	}
	return nil
}
