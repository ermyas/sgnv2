package relayer

import (
	"encoding/json"
	"time"
)

type MintRequest struct {
	MintId         []byte    `json:"mint_id"`
	MintChainId    uint64    `json:"mint_chain_id"`
	DepositChainId uint64    `json:"deposit_chain_id"`
	DepositId      []byte    `json:"deposit_id"`
	RetryCount     uint64    `json:"retry_count"`
	CreateTime     time.Time `json:"create_time"`
}

func NewMintRequest(mintId []byte, mintChainId uint64, depositChainId uint64, depositId []byte) MintRequest {
	return MintRequest{
		MintId:         mintId,
		MintChainId:    mintChainId,
		DepositChainId: depositChainId,
		DepositId:      depositId,
		RetryCount:     0,
		CreateTime:     time.Now(),
	}
}

func NewMintRequestFromBytes(input []byte) MintRequest {
	mint := MintRequest{}
	mint.MustUnMarshal(input)
	return mint
}

// Marshal MintRequest into json bytes
func (r MintRequest) MustMarshal() []byte {
	res, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}

	return res
}

// Unmarshal json bytes to MintRequest
func (r *MintRequest) MustUnMarshal(input []byte) {
	err := json.Unmarshal(input, r)
	if err != nil {
		panic(err)
	}
}

type WithdrawRequest struct {
	WithdrawId      []byte    `json:"withdraw_id"`
	WithdrawChainId uint64    `json:"withdraw_chain_id"`
	BurnChainId     uint64    `json:"burn_chain_id"`
	BurnId          []byte    `json:"burn_id"`
	RetryCount      uint64    `json:"retry_count"`
	CreateTime      time.Time `json:"create_time"`
}

func NewWithdrawRequest(wdId []byte, wdChid uint64, burnChid uint64, burnId []byte) WithdrawRequest {
	return WithdrawRequest{
		WithdrawId:      wdId,
		WithdrawChainId: wdChid,
		BurnChainId:     burnChid,
		BurnId:          burnId,
		RetryCount:      0,
		CreateTime:      time.Now(),
	}
}

func NewWithdrawRequestFromBytes(input []byte) WithdrawRequest {
	wd := WithdrawRequest{}
	wd.MustUnMarshal(input)
	return wd
}

// Marshal WithdrawRequest into json bytes
func (r WithdrawRequest) MustMarshal() []byte {
	res, err := json.Marshal(&r)
	if err != nil {
		panic(err)
	}

	return res
}

// Unmarshal json bytes to WithdrawRequest
func (r *WithdrawRequest) MustUnMarshal(input []byte) {
	err := json.Unmarshal(input, r)
	if err != nil {
		panic(err)
	}
}
