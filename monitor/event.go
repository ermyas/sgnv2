package monitor

import (
	"encoding/json"

	"github.com/celer-network/sgn-v2/contracts"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventName string

const (
	ValidatorParamsUpdate EventName = "ValidatorParamsUpdate"
	ValidatorStatusUpdate EventName = "ValidatorStatusUpdate"
	DelegationUpdate      EventName = "DelegationUpdate"
	SgnAddrUpdate         EventName = "SgnAddrUpdate"
)

// Wrapper for ethereum Event
type EventWrapper struct {
	Name EventName `json:"name"`
	Log  types.Log `json:"log"`
}

func NewEvent(name EventName, l types.Log) *EventWrapper {
	return &EventWrapper{
		Name: name,
		Log:  l,
	}
}

func NewEventFromBytes(input []byte) *EventWrapper {
	event := &EventWrapper{}
	event.MustUnMarshal(input)
	return event
}

// Marshal event into json bytes
func (e *EventWrapper) MustMarshal() []byte {
	res, err := json.Marshal(&e)
	if err != nil {
		panic(err)
	}

	return res
}

// Unmarshal json bytes to event
func (e *EventWrapper) MustUnMarshal(input []byte) {
	err := json.Unmarshal(input, e)
	if err != nil {
		panic(err)
	}
}

func (e *EventWrapper) ParseEvent(ethClient *contracts.EthClient) interface{} {
	var res interface{}
	var err error
	switch e.Name {
	case ValidatorParamsUpdate:
		res, err = ethClient.Staking.ParseValidatorParamsUpdate(e.Log)
	case ValidatorStatusUpdate:
		res, err = ethClient.Staking.ParseValidatorStatusUpdate(e.Log)
	case DelegationUpdate:
		res, err = ethClient.Staking.ParseDelegationUpdate(e.Log)
	case SgnAddrUpdate:
		res, err = ethClient.SGN.ParseSgnAddrUpdate(e.Log)
	default:
		panic("Unsupported event")
	}

	if err != nil {
		panic(err)
	}
	return res
}
