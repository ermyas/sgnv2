package types

import (
	"github.com/celer-network/sgn-v2/contracts"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdk_staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/gogo/protobuf/proto"
)

// valAccount is used for running validator node, and transactors is used for running gateway
type Validator struct {
	EthAddress     string                  `json:"eth_address"`
	EthSigner      string                  `json:"eth_signer"`
	SgnAddress     sdk.AccAddress          `json:"sgn_address"`
	Transactors    []sdk.AccAddress        `json:"transactors"`
	Tokens         sdk.Int                 `json:"tokens"`
	Shares         sdk.Int                 `json:"shares"`
	CommissionRate sdk.Dec                 `json:"commission_rate"`
	RequestCount   sdk.Int                 `json:"request_count"`
	Description    sdk_staking.Description `json:"description"`
}

func NewValidator(ethAddress, ethSigner string, sgnAddress sdk.AccAddress) Validator {
	return Validator{
		EthAddress: contracts.FormatAddrHex(ethAddress),
		EthSigner:  contracts.FormatAddrHex(ethSigner),
		SgnAddress: sgnAddress,
	}
}

func (v Validator) Reset()         { v = Validator{} }
func (v Validator) String() string { return proto.CompactTextString(v) }
func (Validator) ProtoMessage()    {}
