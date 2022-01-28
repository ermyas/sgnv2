<!--
order: 2
-->

# State

## Parameters and base types

`Parameters` define the rules according to which votes are run. There can only
be one active parameter set at any given time. If governance wants to change a
parameter set, either to modify a value or add/remove a parameter field, a new
parameter set has to be created and the previous one rendered inactive.

### DepositParams

### VotingParams

### TallyParams

Parameters are stored in a global `GlobalParams` KVStore.

Additionally, we introduce some basic types:

```go
type Vote struct {
	ProposalId uint64     `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"proposal_id,omitempty" yaml:"proposal_id"`
	Voter      string     `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	Option     VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=sgn.gov.v1.VoteOption" json:"option,omitempty"`
}

const (
    VoteNull    = 0
	VoteYes     = 1
	VoteAbstain = 2
	VoteNo      = 3
)

type Proposal struct {
	ProposalId       uint64                                 `protobuf:"varint,1,opt,name=proposal_id,json=proposalId,proto3" json:"id" yaml:"id"`
	Content          *types.Any                             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Status           ProposalStatus                         `protobuf:"varint,3,opt,name=status,proto3,enum=sgn.gov.v1.ProposalStatus" json:"status,omitempty" yaml:"proposal_status"`
	FinalTallyResult TallyResult                            `protobuf:"bytes,4,opt,name=final_tally_result,json=finalTallyResult,proto3" json:"final_tally_result" yaml:"final_tally_result"`
	SubmitTime       time.Time                              `protobuf:"bytes,5,opt,name=submit_time,json=submitTime,proto3,stdtime" json:"submit_time" yaml:"submit_time"`
	DepositEndTime   time.Time                              `protobuf:"bytes,6,opt,name=deposit_end_time,json=depositEndTime,proto3,stdtime" json:"deposit_end_time" yaml:"deposit_end_time"`
	TotalDeposit     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=total_deposit,json=totalDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_deposit"`
	VotingStartTime  time.Time                              `protobuf:"bytes,8,opt,name=voting_start_time,json=votingStartTime,proto3,stdtime" json:"voting_start_time" yaml:"voting_start_time"`
	VotingEndTime    time.Time                              `protobuf:"bytes,9,opt,name=voting_end_time,json=votingEndTime,proto3,stdtime" json:"voting_end_time" yaml:"voting_end_time"`
}

type ProposalStatus int32

const (
	// PROPOSAL_STATUS_UNSPECIFIED defines the default proposal status.
	StatusNil ProposalStatus = 0
	// PROPOSAL_STATUS_DEPOSIT_PERIOD defines a proposal status during the deposit
	// period.
	StatusDepositPeriod ProposalStatus = 1
	// PROPOSAL_STATUS_VOTING_PERIOD defines a proposal status during the voting
	// period.
	StatusVotingPeriod ProposalStatus = 2
	// PROPOSAL_STATUS_PASSED defines a proposal status of a proposal that has
	// passed.
	StatusPassed ProposalStatus = 3
	// PROPOSAL_STATUS_REJECTED defines a proposal status of a proposal that has
	// been rejected.
	StatusRejected ProposalStatus = 4
	// PROPOSAL_STATUS_FAILED defines a proposal status of a proposal that has
	// failed.
	StatusFailed ProposalStatus = 5
)
```

## Proposals

`Proposal` objects are used to account votes and generally track the proposal's state. They contain `Content` which denotes
what this proposal is about, and other fields, which are the mutable state of
the governance process.

+++ https://github.com/cosmos/cosmos-sdk/blob/v0.40.0/proto/cosmos/gov/v1beta1/gov.proto#L55-L77

```go
type Content interface {
	GetTitle() string
	GetDescription() string
	ProposalRoute() string
	ProposalType() string
	ValidateBasic() error
	String() string
}
```

The `Content` on a proposal is an interface which contains the information about
the `Proposal` such as the tile, description, and any notable changes. Also, this
`Content` type can by implemented by any module. The `Content`'s `ProposalRoute`
returns a string which must be used to route the `Content`'s `Handler` in the
governance keeper. This allows the governance keeper to execute proposal logic
implemented by any module. If a proposal passes, the handler is executed. Only
if the handler is successful does the state get persisted and the proposal finally
passes. Otherwise, the proposal is rejected.

```go
type Handler func(ctx sdk.Context, content Content) error
```

The `Handler` is responsible for actually executing the proposal and processing
any state changes specified by the proposal. It is executed only if a proposal
passes during `EndBlock`.

## Stores

```go
	ProposalsKeyPrefix          = []byte{0x00}
	ActiveProposalQueuePrefix   = []byte{0x01}
	InactiveProposalQueuePrefix = []byte{0x02}
	ProposalIDKey               = []byte{0x03}

	DepositsKeyPrefix   = []byte{0x10}
	DepositorsKeyPrefix = []byte{0x11}

	VotesKeyPrefix = []byte{0x20}
```

## Proposal Processing Queue

**Store:**

- `ActiveProposalQueue`: A queue `queue[proposalID]` containing all the
  `ProposalIDs` of proposals that reached `MinDeposit`. During each `EndBlock`,
  all the proposals that have reached the end of their voting period are processed.
  To process a finished proposal, the application tallies the votes, computes the
  votes of each validator and checks if every validator in the validator set has
  voted. If the proposal is accepted, deposits are refunded. Finally, the proposal
  content `Handler` is executed.
