package eth

const (
	// ValidatorStatus
	Unbonded  = 1
	Unbonding = 2
	Bonded    = 3

	// ParamNames
	ParamProposalDeposit       = 0
	ParamVotingPeriod          = 1
	ParamUnbondingPeriod       = 2
	ParamMaxBondedValidators   = 3
	ParamMinValidatorTokens    = 4
	ParamMinSelfDelegation     = 5
	ParamAdvanceNoticePeriod   = 6
	ParamValidatorBondInterval = 7
	ParamMaxSlashFactor        = 8

	VoteNull    = 0
	VoteYes     = 1
	VoteAbstain = 2
	VoteNo      = 3

	ProposalStatusUninitiated = 0
	ProposalStatusVoting      = 1
	ProposalStatusClosed      = 2

	TxFailure = 0
	TxSuccess = 1

	CommissionRateBase = 10000
)
