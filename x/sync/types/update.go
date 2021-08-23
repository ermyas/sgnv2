package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func NewUpdate(
	id uint64, updateType UpdateType, data []byte, ethBlock uint64,
	proposer string, proposeTs, closingTs uint64) *Update {

	return &Update{
		Id:        id,
		Type:      updateType,
		Data:      data,
		EthBlock:  ethBlock,
		Proposer:  proposer,
		ProposeTs: proposeTs,
		ClosingTs: closingTs,
	}
}

func MustMarshalUpdate(cdc codec.BinaryCodec, update *Update) []byte {
	return cdc.MustMarshal(update)
}

func MustUnmarshalUpdate(cdc codec.BinaryCodec, value []byte) Update {
	update, err := UnmarshalUpdate(cdc, value)
	if err != nil {
		panic(err)
	}

	return update
}

func UnmarshalUpdate(cdc codec.BinaryCodec, value []byte) (u Update, err error) {
	err = cdc.Unmarshal(value, &u)
	return u, err
}
