package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewPendingUpdate(
	id uint64, dataType DataType, data []byte, chainId, chainBlock uint64,
	proposer string, proposeTs, closingTs uint64) *PendingUpdate {

	return &PendingUpdate{
		Id:         id,
		Type:       dataType,
		Data:       data,
		ChainId:    chainId,
		ChainBlock: chainBlock,
		Proposer:   proposer,
		ProposeTs:  proposeTs,
		ClosingTs:  closingTs,
	}
}

func MustMarshalPendingUpdate(cdc codec.BinaryCodec, update *PendingUpdate) []byte {
	return cdc.MustMarshal(update)
}

func MustUnmarshalPendingUpdate(cdc codec.BinaryCodec, value []byte) PendingUpdate {
	update, err := UnmarshalPendingUpdate(cdc, value)
	if err != nil {
		panic(err)
	}

	return update
}

func UnmarshalPendingUpdate(cdc codec.BinaryCodec, value []byte) (u PendingUpdate, err error) {
	err = cdc.Unmarshal(value, &u)
	return u, err
}
