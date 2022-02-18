package relayer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PegContracts struct {
	bridge  *eth.PegBridgeContract
	vault   *eth.PegVaultContract
	bridge2 *eth.PegBridgeV2Contract
	vault2  *eth.PegVaultV2Contract
}

func (pc *PegContracts) GetPegVaultContract() *eth.PegVaultContract {
	if pc == nil {
		return nil
	}
	return pc.vault
}

func (pc *PegContracts) GetPegBridgeContract() *eth.PegBridgeContract {
	if pc == nil {
		return nil
	}
	return pc.bridge
}

func (pc *PegContracts) GetPegVaultV2Contract() *eth.PegVaultV2Contract {
	if pc == nil {
		return nil
	}
	return pc.vault2
}

func (pc *PegContracts) GetPegBridgeV2Contract() *eth.PegBridgeV2Contract {
	if pc == nil {
		return nil
	}
	return pc.bridge2
}

func NewPegContracts(cfg *common.OneChainConfig, client *ethclient.Client) (*PegContracts, error) {
	pegContracts := &PegContracts{}
	var err error
	if cfg.OTVault != "" {
		pegContracts.vault, err = eth.NewPegVaultContract(eth.Hex2Addr(cfg.OTVault), client)
		if err != nil {
			return nil, fmt.Errorf("OriginalTokenVault contract at %s, err %w", cfg.OTVault, err)
		}
	}
	if cfg.PTBridge != "" {
		pegContracts.bridge, err = eth.NewPegBridgeContract(eth.Hex2Addr(cfg.PTBridge), client)
		if err != nil {
			return nil, fmt.Errorf("PeggedTokenBridge contract at %s, err %w", cfg.OTVault, err)
		}
	}
	if cfg.OTVault2 != "" {
		pegContracts.vault2, err = eth.NewPegVaultV2Contract(eth.Hex2Addr(cfg.OTVault2), client)
		if err != nil {
			return nil, fmt.Errorf("OriginalTokenVaultV2 contract at %s, err %w", cfg.OTVault, err)
		}
	}
	if cfg.PTBridge2 != "" {
		pegContracts.bridge2, err = eth.NewPegBridgeV2Contract(eth.Hex2Addr(cfg.PTBridge2), client)
		if err != nil {
			return nil, fmt.Errorf("PeggedTokenBridgeV2 contract at %s, err %w", cfg.OTVault, err)
		}
	}
	return pegContracts, nil
}

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
