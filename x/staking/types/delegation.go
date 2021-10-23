package types

import (
	"encoding/json"
	fmt "fmt"
	"sort"
	"strings"

	"github.com/celer-network/sgn-v2/eth"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"
)

// Implements Delegation interface
var _ DelegationI = Delegation{}

func NewDelegation(
	delegatorAddress eth.Addr,
	validatorAddress eth.Addr,
	shares sdk.Int,
) Delegation {
	return Delegation{
		DelegatorAddress: eth.Addr2Hex(delegatorAddress),
		ValidatorAddress: eth.Addr2Hex(validatorAddress),
		Shares:           shares,
	}
}

func MustMarshalDelegation(cdc codec.BinaryCodec, delegation Delegation) []byte {
	return cdc.MustMarshal(&delegation)
}

func MustUnmarshalDelegation(cdc codec.BinaryCodec, value []byte) Delegation {
	delegator, err := UnmarshalDelegation(cdc, value)
	if err != nil {
		panic(err)
	}

	return delegator
}

func UnmarshalDelegation(cdc codec.BinaryCodec, value []byte) (d Delegation, err error) {
	err = cdc.Unmarshal(value, &d)
	return d, err
}

func (d Delegation) GetDelegatorAddr() eth.Addr {
	delAddr := eth.Hex2Addr(d.DelegatorAddress)
	return delAddr
}

func (d Delegation) GetValidatorAddr() eth.Addr {
	valAddr := eth.Hex2Addr(d.ValidatorAddress)
	return valAddr
}

func (d Delegation) GetShares() sdk.Int { return d.Shares }

// String returns a human readable string representation of a Delegation.
func (d Delegation) String() string {
	return proto.CompactTextString(&d)
}

type Delegations []Delegation

// Sort Validators sorts validator array in ascending operator address order
func (d Delegations) Sort() {
	sort.Sort(d)
}

// Implements sort interface
func (d Delegations) Len() int {
	return len(d)
}

// Implements sort interface
func (d Delegations) Less(i, j int) bool {
	return d[i].Shares.GT(d[j].Shares)
}

// Implements sort interface
func (d Delegations) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// ----------------------------------------------------------------------------
// Client Types

// NewDelegationResp creates a new DelegationResponse instance
func NewDelegationResp(
	delegatorAddr eth.Addr, validatorAddr eth.Addr, shares sdk.Int, balance sdk.Coin,
) DelegationResponse {
	return DelegationResponse{
		Delegation: NewDelegation(delegatorAddr, validatorAddr, shares),
		Balance:    balance,
	}
}

// String implements the Stringer interface for DelegationResponse.
func (d DelegationResponse) String() string {
	return fmt.Sprintf("%s\n  Balance:   %s", d.Delegation.String(), d.Balance)
}

type delegationRespAlias DelegationResponse

// MarshalJSON implements the json.Marshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (d DelegationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal((delegationRespAlias)(d))
}

// UnmarshalJSON implements the json.Unmarshaler interface. This is so we can
// achieve a flattened structure while embedding other types.
func (d *DelegationResponse) UnmarshalJSON(bz []byte) error {
	return json.Unmarshal(bz, (*delegationRespAlias)(d))
}

// DelegationResponses is a collection of DelegationResp
type DelegationResponses []DelegationResponse

// String implements the Stringer interface for DelegationResponses.
func (d DelegationResponses) String() (out string) {
	for _, del := range d {
		out += del.String() + "\n"
	}

	return strings.TrimSpace(out)
}

// ----------------------- CLI print-friendly output --------------------

type DelegationOutput struct {
	DelegatorAddress string  `json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress string  `json:"validator_address,omitempty" yaml:"validator_address"`
	Shares           sdk.Dec `json:"shares" yaml:"shares"`
	Tokens           sdk.Dec `json:"tokens" yaml:"tokens"`
}

func newDelegationOutput(d *DelegationResponse) *DelegationOutput {
	return &DelegationOutput{
		DelegatorAddress: d.Delegation.DelegatorAddress,
		ValidatorAddress: d.Delegation.ValidatorAddress,
		Shares:           sdk.NewDecFromIntWithPrec(d.Delegation.Shares, 18),
		Tokens:           sdk.NewDecFromIntWithPrec(d.Balance.Amount, 18),
	}
}

func (d *DelegationResponse) YamlStr() string {
	output := newDelegationOutput(d)
	out, _ := yaml.Marshal(output)
	return string(out)
}
