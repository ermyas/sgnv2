package eth

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/core/types"
)

const (
	EventValidatorNotice       = "ValidatorNotice"
	EventValidatorStatusUpdate = "ValidatorStatusUpdate"
	EventDelegationUpdate      = "DelegationUpdate"
)

// Wrapper for ethereum Event
type EventWrapper struct {
	Name string    `json:"name"`
	Log  types.Log `json:"log"`
}

func NewEvent(name string, l types.Log) *EventWrapper {
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

func (e *EventWrapper) ParseEvent(ethClient *EthClient) interface{} {
	var res interface{}
	var err error
	switch e.Name {
	case EventValidatorNotice:
		res, err = ethClient.Contracts.Staking.ParseValidatorNotice(e.Log)
	case EventValidatorStatusUpdate:
		res, err = ethClient.Contracts.Staking.ParseValidatorStatusUpdate(e.Log)
	case EventDelegationUpdate:
		res, err = ethClient.Contracts.Staking.ParseDelegationUpdate(e.Log)
	default:
		panic("Unsupported event")
	}

	if err != nil {
		panic(err)
	}
	return res
}
