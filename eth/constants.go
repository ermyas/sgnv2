package eth

const (
	// CandidateStatus
	Unbonded  = 1
	Unbonding = 2
	Bonded    = 3

	// ParamNames
	ProposalDeposit       = 0
	VotePeriod            = 1
	UnbondingPeriod       = 2
	MaxBondedValidators   = 3
	MinValidatorTokens    = 4
	MinSelfDelegation     = 5
	AdvanceNoticePeriod   = 6
	ValidatorBondInterval = 7
	MaxSlashFactor        = 8

	Unvoted     = 0
	VoteYes     = 1
	VoteAbstain = 2
	VoteNo      = 3

	ProposalStatusUninitiated = 0
	ProposalStatusVoting      = 1
	ProposalStatusClosed      = 2

	TxFailure = 0
	TxSuccess = 1
)
