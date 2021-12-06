package gatewaysvc

import (
	"context"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/common"
	"github.com/celer-network/sgn-v2/gateway/dal"
	"github.com/celer-network/sgn-v2/gateway/webapi"
	"github.com/ethereum/go-ethereum/crypto"
)

func (gs *GatewayService) UpdateChain(ctx context.Context, request *webapi.UpdateChainRequest) (*webapi.UpdateChainResponse, error) {
	if !checkSigner(common.Hex2Addr(request.GetAddr()).Bytes(), request.Sig) {
		return &webapi.UpdateChainResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "invalid addr to update chain",
			},
		}, nil
	}
	chainInput := request.GetChain()
	chainId := uint64(chainInput.GetId())

	if chainId == 0 {
		return &webapi.UpdateChainResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "chainId is 0",
			},
		}, nil
	}
	var name, icon, gasTokenSymbol, exploreUrl, dropGasAmt string
	var suggestedBaseFee float64
	chainInDb, url, chainFound, err := dal.DB.GetChain(chainId)
	if chainInDb != nil && chainFound && err == nil {
		name = chainInDb.GetName()
		icon = chainInDb.GetIcon()
		gasTokenSymbol = chainInDb.GetGasTokenSymbol()
		exploreUrl = chainInDb.GetExploreUrl()
		dropGasAmt = chainInDb.GetDropGasAmt()
	}
	if chainInput.GetName() != "" {
		name = chainInput.GetName()
	}
	if chainInput.GetIcon() != "" {
		icon = chainInput.GetIcon()
	}
	if request.GetTxUrlPrefix() != "" {
		url = request.GetTxUrlPrefix()
	}
	if chainInput.GetGasTokenSymbol() != "" {
		gasTokenSymbol = chainInput.GetGasTokenSymbol()
	}
	if chainInput.GetExploreUrl() != "" {
		exploreUrl = chainInput.GetExploreUrl()
	}
	if chainInput.GetDropGasAmt() != "" {
		dropGasAmt = chainInput.GetDropGasAmt()
	}
	if chainInput.GetSuggestedBaseFee() != 0 {
		suggestedBaseFee = chainInput.GetSuggestedBaseFee()
	}

	dal.DB.UpsertChainUIInfo(chainId, name, icon, url, gasTokenSymbol, exploreUrl, "", dropGasAmt, suggestedBaseFee)
	chainInDb, url, _, _ = dal.DB.GetChain(chainId)
	return &webapi.UpdateChainResponse{
		Chain:       chainInDb,
		TxUrlPrefix: url,
	}, nil
}

func (gs *GatewayService) UpdateToken(ctx context.Context, request *webapi.UpdateTokenRequest) (*webapi.UpdateTokenResponse, error) {
	if !checkSigner(common.Hex2Addr(request.GetAddr()).Bytes(), request.Sig) {
		return &webapi.UpdateTokenResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "invalid addr to update token",
			},
		}, nil
	}
	chainId := uint64(request.GetChainId())
	tokenSymbol := request.GetTokenSymbol()
	if chainId == 0 || tokenSymbol == "" {
		return &webapi.UpdateTokenResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "invalid input, check chainId and tokenSymbol",
			},
		}, nil
	}
	var name, icon string
	tokenInDb, found, err := dal.DB.GetTokenBySymbol(tokenSymbol, chainId)
	if tokenInDb != nil && found && err == nil {
		name = tokenInDb.GetName()
		icon = tokenInDb.GetIcon()
	} else {
		return &webapi.UpdateTokenResponse{
			Err: &webapi.ErrMsg{
				Code: webapi.ErrCode_ERROR_CODE_COMMON,
				Msg:  "token not found, can not update UI info now",
			},
		}, nil
	}
	if request.GetTokenName() != "" {
		name = request.GetTokenName()
	}
	if request.GetTokenIcon() != "" {
		icon = request.GetTokenIcon()
	}
	dal.DB.UpdateTokenUIInfo(tokenSymbol, chainId, name, icon)
	tokenInDb, _, _ = dal.DB.GetTokenBySymbol(tokenSymbol, chainId)
	return &webapi.UpdateTokenResponse{
		Token: tokenInDb,
	}, nil
}

func checkSigner(data []byte, sig []byte) bool {
	if len(sig) == 65 { // we could return zeroAddr if len not 65
		if sig[64] == 27 || sig[64] == 28 {
			// SigToPub only expect v to be 0 or 1,
			// see https://github.com/ethereum/go-ethereum/blob/v1.8.23/internal/ethapi/api.go#L468.
			// we've been ok as our own code only has v 0 or 1, but using external signer may cause issue
			// we also fix v in celersdk.PublishSignedResult to be extra safe
			sig[64] -= 27
		}
	}
	pubKey, err := crypto.SigToPub(generatePrefixedHash(data), sig)
	if err != nil {
		log.Warnf("RecoverSigner err:%+v", err)
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	checked := dal.DB.IsAdminAddrValid(recoveredAddr.String())
	if !checked {
		log.Warnf("error addr:%s to use admin api:", recoveredAddr)
	}
	return checked
}

func generatePrefixedHash(data []byte) []byte {
	return crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n32"), crypto.Keccak256(data))
}
