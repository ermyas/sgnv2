package types

import (
	"context"

	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	"github.com/celer-network/sgn-v2/x/pegbridge/types"
	stakingtypes "github.com/celer-network/sgn-v2/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

type StakingKeeper interface {
	GetValidatorBySgnAddr(sdk.Context, sdk.AccAddress) (stakingtypes.ValidatorI, bool)
	CheckSenderBondedValidator(ctx sdk.Context, sender string) (stakingtypes.ValidatorI, error)
}

type CbridgeKeeper interface {
	GetChainSigners(ctx sdk.Context, chainId uint64) (signers cbrtypes.ChainSigners, found bool)
	GetCbrContractAddr(ctx sdk.Context, chid uint64) (eth.Addr, bool)
	QueryRelay(c context.Context, request *cbrtypes.QueryRelayRequest) (*cbrtypes.QueryRelayResponse, error)
	QueryXferStatus(ctx sdk.Context, srcXferId eth.Hash) cbrtypes.XferStatus
	QueryXferRefund(ctx sdk.Context, srcXferId eth.Hash) *cbrtypes.WithdrawOnchain
}

type PegbridgeKeeper interface {
	GetOriginalTokenVault(ctx sdk.Context, chainId uint64) (vault commontypes.ContractInfo, found bool)
	GetPeggedTokenBridge(ctx sdk.Context, chainId uint64) (bridge commontypes.ContractInfo, found bool)
	DepositInfo(c context.Context, req *types.QueryDepositInfoRequest) (*types.QueryDepositInfoResponse, error)
	WithdrawInfo(c context.Context, req *types.QueryWithdrawInfoRequest) (*types.QueryWithdrawInfoResponse, error)
	MintInfo(c context.Context, req *types.QueryMintInfoRequest) (*types.QueryMintInfoResponse, error)
	BurnInfo(c context.Context, req *types.QueryBurnInfoRequest) (*types.QueryBurnInfoResponse, error)
}

type DistributionKeeper interface {
	ClaimMessageFees(ctx sdk.Context, delAddr eth.Addr) error
}
