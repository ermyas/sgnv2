package types

import (
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func (msg *MsgClaimFee) EncodeDataToSignByDelegator() []byte {
	// keccak256(abi.encodePacked("PegClaimFee", delegator_address, chain_id, token_address, nonce))
	return solsha3.SoliditySHA3(
		[]string{"string", "address", "uint64", "address", "uint64"},
		[]interface{}{"PegClaimFee", msg.DelegatorAddress, msg.ChainId, msg.TokenAddress, msg.Nonce},
	)
}
