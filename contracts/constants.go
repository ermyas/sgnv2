package contracts

const (
	// ValidatorChangeType
	AddValidator    = 0
	RemoveValidator = 1

	// CandidateStatus
	Unbonded  = 0
	Bonded    = 1
	Unbonding = 2

	// ParamNames
	ProposalDeposit     = 0
	GovernVoteTimeout   = 1
	SlashTimeout        = 2
	MinValidatorNum     = 3
	MaxValidatorNum     = 4
	MinStakeInPool      = 5
	AdvanceNoticePeriod = 6
	MigrationTime       = 7

	Unvoted     = 0
	VoteYes     = 1
	VoteNo      = 2
	VoteAbstain = 3

	ProposalStatusUninitiated = 0
	ProposalStatusVoting      = 1
	ProposalStatusClosed      = 2

	TxFailure = 0
	TxSuccess = 1
)
