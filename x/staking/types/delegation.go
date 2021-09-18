package types

import (
	"sort"

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

// ----------------------- CLI print-friendly output --------------------

type DelegationOutput struct {
	DelegatorAddress string  `json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress string  `json:"validator_address,omitempty" yaml:"validator_address"`
	Shares           sdk.Dec `json:"shares" yaml:"shares"`
}

func newDelegationOutput(d *Delegation) *DelegationOutput {
	return &DelegationOutput{
		DelegatorAddress: d.DelegatorAddress,
		ValidatorAddress: d.ValidatorAddress,
		Shares:           sdk.NewDecFromIntWithPrec(d.Shares, 18),
	}
}

func (d *Delegation) YamlStr() string {
	output := newDelegationOutput(d)
	out, _ := yaml.Marshal(output)
	return string(out)
}
