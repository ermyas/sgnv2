package types

import (
	"math/big"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

// CalcMintId calculates mint ID with the following
// keccak256(abi.encodePacked(request.account, request.token, request.amount, request.refChainId, request.refId))
func CalcMintId(
	account eth.Addr, token eth.Addr, amount *big.Int, depositor eth.Addr, refChainId uint64, refId eth.Hash) eth.Hash {
	hash := solsha3.SoliditySHA3(
		[]string{"address", "address", "uint256", "address", "uint64", "bytes32"},
		[]interface{}{account, token, amount, depositor, refChainId, refId},
	)
	return eth.Bytes2Hash(hash)
}

// CalcWithdrawId calculates withdraw ID with the following
// keccak256(abi.encodePacked(request.receiver, request.token, request.amount, request.refChainId, request.refId)
func CalcWithdrawId(receiver eth.Addr, token eth.Addr, amount *big.Int, burnAccount eth.Addr, refChainId uint64, refId eth.Hash) eth.Hash {
	hash := solsha3.SoliditySHA3(
		[]string{"address", "address", "uint256", "address", "uint64", "bytes32"},
		[]interface{}{receiver, token, amount, burnAccount, refChainId, refId},
	)
	return eth.Bytes2Hash(hash)
}

// EncodeDataToSign generates the message to sign for a mint
// domain = keccak256(abi.encodePacked(block.chainid, address(this), "Mint"));
// data = abi.encodePacked(domain, _request)
func (i *MintInfo) EncodeDataToSign(verifyingContractAddress eth.Addr) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(i.ChainId), verifyingContractAddress, "Mint"},
	)
	// NOTE: Manual concatenation as solsha3 DOES NOT SUPPORT dynamic "bytes"
	return append(domain, i.MintProtoBytes...)
}

// EncodeDataToSign generates the message to sign for a withdraw
// domain = keccak256(abi.encodePacked(block.chainid, address(this), "Withdraw"));
// data = abi.encodePacked(domain, _request)
func (i *WithdrawInfo) EncodeDataToSign(verifyingContractAddress eth.Addr) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(i.ChainId), verifyingContractAddress, "Withdraw"},
	)
	// NOTE: Manual concatenation as solsha3 DOES NOT SUPPORT dynamic "bytes"
	return append(domain, i.WithdrawProtoBytes...)
}

// AddSig adds a signature to a mint info
func (i *MintInfo) AddSig(msgToSign []byte, sig []byte, expectedSigner string) error {
	sigs, err := commontypes.AddSig(i.Signatures, msgToSign, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}
	i.Signatures = sigs
	return nil
}

// AddSig adds a signature to a withdraw info
func (i *WithdrawInfo) AddSig(msgToSign []byte, sig []byte, expectedSigner string) error {
	sigs, err := commontypes.AddSig(i.Signatures, msgToSign, sig, expectedSigner)
	if err != nil {
		log.Error(err)
		return err
	}
	i.Signatures = sigs
	return nil
}
