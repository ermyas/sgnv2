package validator

import (
	"github.com/celer-network/sgn-v2/x/validator/client/cli"
	"github.com/celer-network/sgn-v2/x/validator/types"
)

const (
	ModuleName             = types.ModuleName
	RouterKey              = types.RouterKey
	StoreKey               = types.StoreKey
	QuerySyncer            = types.QuerySyncer
	QueryDelegator         = types.QueryDelegator
	QueryValidator         = types.QueryValidator
	QueryValidators        = types.QueryValidators
	QueryDelegators        = types.QueryDelegators
	QueryParameters        = types.QueryParameters
	AttributeKeyEthAddress = types.AttributeKeyEthAddress
	ActionInitiateWithdraw = types.ActionInitiateWithdraw
)

var (
	NewMsgSetTransactors           = types.NewMsgSetTransactors
	NewMsgEditValidatorDescription = types.NewMsgEditValidatorDescription
	ModuleCdc                      = types.ModuleCdc
	RegisterCodec                  = types.RegisterCodec
	SyncerKey                      = types.SyncerKey
	ValidatorKeyPrefix             = types.ValidatorKeyPrefix
	GetDelegatorKey                = types.GetDelegatorKey
	GetDelegatorsKey               = types.GetDelegatorsKey
	GetValidatorKey                = types.GetValidatorKey
	NewSyncer                      = types.NewSyncer
	NewDelegator                   = types.NewDelegator
	NewValidator                   = types.NewValidator
	CLIQuerySyncer                 = cli.QuerySyncer
	CLIQueryValidator              = cli.QueryValidator
	CLIQueryValidators             = cli.QueryValidators
	CLIQueryDelegators             = cli.QueryDelegators
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
	QueryDelegatorParams        = types.QueryDelegatorParams
	QueryValidatorParams        = types.QueryValidatorParams
	MsgSetTransactors           = types.MsgSetTransactors
	MsgEditValidatorDescription = types.MsgEditValidatorDescription
	GenesisState                = types.GenesisState
)
