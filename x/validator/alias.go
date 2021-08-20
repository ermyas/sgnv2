package validator

import (
	"github.com/celer-network/sgn-v2/x/validator/client/cli"
	"github.com/celer-network/sgn-v2/x/validator/types"
)

const (
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	QuerySyncer              = types.QuerySyncer
	QueryDelegator           = types.QueryDelegator
	QueryValidator           = types.QueryValidator
	QueryValidators          = types.QueryValidators
	QueryValidatorDelegators = types.QueryValidatorDelegators
	QueryReward              = types.QueryReward
	QueryRewardEpoch         = types.QueryRewardEpoch
	QueryRewardStats         = types.QueryRewardStats
	QueryParameters          = types.QueryParameters
	TypeMsgClaimReward       = types.TypeMsgClaimReward
	AttributeKeyEthAddress   = types.AttributeKeyEthAddress
	ActionInitiateWithdraw   = types.ActionInitiateWithdraw
)

var (
	NewMsgSetTransactors           = types.NewMsgSetTransactors
	NewMsgClaimReward              = types.NewMsgClaimReward
	NewMsgEditValidatorDescription = types.NewMsgEditValidatorDescription
	NewMsgSignReward               = types.NewMsgSignReward
	NewQueryRewardParams           = types.NewQueryRewardParams
	ModuleCdc                      = types.ModuleCdc
	RegisterCodec                  = types.RegisterCodec
	SyncerKey                      = types.SyncerKey
	ValidatorKeyPrefix             = types.ValidatorKeyPrefix
	GetDelegatorKey                = types.GetDelegatorKey
	GetDelegatorsKey               = types.GetDelegatorsKey
	GetValidatorKey                = types.GetValidatorKey
	RewardKeyPrefix                = types.RewardKeyPrefix
	GetRewardKey                   = types.GetRewardKey
	PendingRewardKeyPrefix         = types.PendingRewardKeyPrefix
	GetPendingRewardKey            = types.GetPendingRewardKey
	RewardEpochKey                 = types.RewardEpochKey
	NewSyncer                      = types.NewSyncer
	NewDelegator                   = types.NewDelegator
	NewValidator                   = types.NewValidator
	NewReward                      = types.NewReward
	CLIQuerySyncer                 = cli.QuerySyncer
	CLIQueryValidator              = cli.QueryValidator
	CLIQueryValidators             = cli.QueryValidators
	CLIQueryValidatorDelegators    = cli.QueryValidatorDelegators
	CLIQueryReward                 = cli.QueryReward
	CLIQueryDelegator              = cli.QueryDelegator
	CLIQuerySgnValidator           = cli.QuerySgnValidator
	CLIQuerySgnValidators          = cli.QuerySgnValidators
	CLIQueryBondedSgnValidators    = cli.QueryBondedSgnValidators
	CLIQueryParams                 = cli.QueryParams
	DefaultParams                  = types.DefaultParams
)

type (
	Syncer                      = types.Syncer
	Params                      = types.Params
	Delegator                   = types.Delegator
	Validator                   = types.Validator
	Reward                      = types.Reward
	QueryDelegatorParams        = types.QueryDelegatorParams
	QueryValidatorParams        = types.QueryValidatorParams
	QueryRewardParams           = types.QueryRewardParams
	MsgSetTransactors           = types.MsgSetTransactors
	MsgEditValidatorDescription = types.MsgEditValidatorDescription
	MsgClaimReward              = types.MsgClaimReward
	MsgSignReward               = types.MsgSignReward
)
