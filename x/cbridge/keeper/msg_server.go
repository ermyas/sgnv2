package keeper

import (
	"bytes"
	"context"
	"fmt"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	"github.com/celer-network/sgn-v2/x/cbridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) InitWithdraw(context.Context, *types.MsgInitWithdraw) (*types.MsgInitWithdrawResp, error) {
	return nil, nil
}

// user can request to sign a previous withdraw again
// to mitigate dos attack, we could be smart and re-use sigs if
// they are still valid. we should also deny if withdraw already
// completed
func (k msgServer) SignAgain(context.Context, *types.MsgSignAgain) (*types.MsgSignAgainResp, error) {
	return nil, nil
}

// send my sig for data, so it can be later submitted onchain
func (k msgServer) SendMySig(ctx context.Context, msg *types.MsgSendMySig) (*types.MsgSendMySigResp, error) {
	if msg == nil {
		return nil, fmt.Errorf("sendMySig could not be nil")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	kv := sdkCtx.KVStore(k.storeKey)
	res := &types.MsgSendMySigResp{}
	if msg.Datatype == types.SignDataType_RELAY {
		relay := new(types.RelayOnChain)
		err := relay.Unmarshal(msg.Data)
		if err != nil {
			return nil, err
		}

		senderAcct, _ := sdk.AccAddressFromBech32(msg.Creator)
		validator, found := k.stakingKeeper.GetValidatorBySgnAddr(sdkCtx, senderAcct)
		if !found {
			return nil, fmt.Errorf("sender is not a validator")
		}
		if !validator.IsBonded() {
			return nil, fmt.Errorf("validator is not bonded")
		}

		// validate sig
		tmpSig := make([]byte, len(msg.MySig))
		copy(tmpSig, msg.MySig)
		signer, err := ethutils.RecoverSigner(msg.Data, tmpSig)
		if err != nil {
			return nil, err
		}
		signerAddr := eth.Addr2Hex(signer)
		if signerAddr != eth.Addr2Hex(validator.GetSignerAddr()) {
			err = fmt.Errorf("invalid signer address %s %s", signerAddr, validator.GetSignerAddr())
			return nil, err
		}

		// add sig
		var xferId [32]byte
		copy(xferId[:], relay.SrcTransferId)
		xferRelay := GetXferRelay(kv, xferId, k.cdc)
		if xferRelay == nil {
			xferRelay = new(types.XferRelay)
			xferRelay.Relay = msg.Data
			xferRelay.SortedSigs = make([]*types.AddrSig, 1)
			xferRelay.SortedSigs[0] = &types.AddrSig{
				Addr: []byte(signerAddr),
				Sig:  msg.MySig,
			}
		} else {
			for i, s := range xferRelay.SortedSigs {
				if string(s.Addr) == signerAddr {
					if !bytes.Equal(s.Sig, msg.MySig) {
						log.Debugf("repeated signer %s overwite existing sig", signerAddr)
						xferRelay.SortedSigs[i] = &types.AddrSig{
							Addr: []byte(signerAddr),
							Sig:  msg.MySig,
						}
					}
					break
				}
				if string(s.Addr) > signerAddr {
					tmp := make([]*types.AddrSig, 0)
					if i == 0 {
						tmp = append(tmp, &types.AddrSig{
							Addr: []byte(signerAddr),
							Sig:  msg.MySig,
						})
						tmp = append(tmp, xferRelay.SortedSigs...)
					} else {
						tmp = append(tmp, xferRelay.SortedSigs[:i]...)
						tmp = append(tmp, &types.AddrSig{
							Addr: []byte(signerAddr),
							Sig:  msg.MySig,
						})
						tmp = append(tmp, xferRelay.SortedSigs[i:]...)
					}

					xferRelay.SortedSigs = tmp
					break
				}
			}
		}

		SetXferRelay(kv, xferId, xferRelay, k.cdc)
	}

	return res, nil
}
