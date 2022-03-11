package types

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	commontypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/eth"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/sha3"
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

func CalcMintIdV2(
	account eth.Addr, token eth.Addr, amount *big.Int, depositor eth.Addr, refChainId uint64, refId eth.Hash, bridgeAddr eth.Addr) eth.Hash {
	hash := solsha3.SoliditySHA3(
		[]string{"address", "address", "uint256", "address", "uint64", "bytes32", "address"},
		[]interface{}{account, token, amount, depositor, refChainId, refId, bridgeAddr},
	)
	return eth.Bytes2Hash(hash)
}

func CalcBurnIdV2(
	account eth.Addr, token eth.Addr, amount *big.Int, toChainId uint64, toAccount eth.Addr,
	nonce, chainId uint64, bridgeAddr eth.Addr) eth.Hash {
	hash := solsha3.SoliditySHA3(
		[]string{"address", "address", "uint256", "uint64", "address", "uint64", "uint64", "address"},
		[]interface{}{account, token, amount, toChainId, toAccount, nonce, chainId, bridgeAddr},
	)
	return eth.Bytes2Hash(hash)
}

func CalcDepositIdV2(
	depositor eth.Addr, token eth.Addr, amount *big.Int, mintChainId uint64, mintAccount eth.Addr,
	nonce, chainId uint64, vaultAddr eth.Addr) eth.Hash {
	hash := solsha3.SoliditySHA3(
		[]string{"address", "address", "uint256", "uint64", "address", "uint64", "uint64", "address"},
		[]interface{}{depositor, token, amount, mintChainId, mintAccount, nonce, chainId, vaultAddr},
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

// withdraw id on flow chain is different due to standard SHA3 and token is string
// must match SafeBox.cdc
func CalcFlowWithdrawId(wdmsg []byte) eth.Hash {
	return eth.Bytes2Hash(StdSHA3Hash(wdmsg))
}

// withdraw id on flow chain is different due to standard SHA3 and token is string
// must match PegBridge.cdc
func CalcFlowMintId(mintmsg []byte) eth.Hash {
	return eth.Bytes2Hash(StdSHA3Hash(mintmsg))
}

// Standard SHA3 hash, used by Flow
func StdSHA3Hash(data []byte) []byte {
	h := sha3.New256()
	h.Write(data)
	return h.Sum(nil)
}

func CalcWithdrawIdV2(
	receiver eth.Addr, token eth.Addr, amount *big.Int, burnAccount eth.Addr, refChainId uint64, refId eth.Hash, bridgeAddr eth.Addr) eth.Hash {
	hash := solsha3.SoliditySHA3(
		[]string{"address", "address", "uint256", "address", "uint64", "bytes32", "address"},
		[]interface{}{receiver, token, amount, burnAccount, refChainId, refId, bridgeAddr},
	)
	return eth.Bytes2Hash(hash)
}

// EncodeDataToSign generates the message to sign for a mint
// domain = keccak256(abi.encodePacked(block.chainid, address(this), "Mint"));
// data = abi.encodePacked(domain, _request)
func (i *MintInfo) EncodeDataToSign(verifyingContractAddress eth.Addr) []byte {
	if commontypes.IsFlowChain(i.ChainId) {
		// uint64 chainid must left pad 00 to full 8 bytes to be consisten w/ cadence logic
		// flow address only has 8 bytes and its string form will left pad 00
		domain := append(eth.ToPadBytes(i.ChainId), []byte(fmt.Sprintf("A.%x.PegBridgeMint", verifyingContractAddress[12:]))...)
		return append(domain, i.MintProtoBytes...)
	}
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
// for flow chain, domain calculation is different. must same as cdc contract
// concat of (chainid.toBigEndianBytes, A.xxxx.Vault.utf8, Withdraw)
func (i *WithdrawInfo) EncodeDataToSign(verifyingContractAddress eth.Addr) []byte {
	if commontypes.IsFlowChain(i.ChainId) {
		domain := append(eth.ToPadBytes(i.ChainId), []byte(fmt.Sprintf("A.%x.SafeBoxWithdraw", verifyingContractAddress[12:]))...)
		return append(domain, i.WithdrawProtoBytes...)
	}
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
