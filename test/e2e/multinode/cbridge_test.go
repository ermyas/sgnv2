package multinode

import (
	"math/big"
	"testing"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/eth"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func setupCbridge() {
	log.Infoln("Set up another mainchain for bridge")
	SetupMainchain2ForBridge()

	log.Infoln("Set up new sgn env")
	p := &tc.ContractParams{
		CelrAddr:              tc.CelrAddr,
		ProposalDeposit:       big.NewInt(1e18),
		VotePeriod:            big.NewInt(5),
		UnbondingPeriod:       big.NewInt(5),
		MaxBondedValidators:   big.NewInt(3),
		MinValidatorTokens:    big.NewInt(1e18),
		MinSelfDelegation:     big.NewInt(1e18),
		AdvanceNoticePeriod:   big.NewInt(1),
		ValidatorBondInterval: big.NewInt(0),
		MaxSlashFactor:        big.NewInt(1e5),
	}
	SetupNewSgnEnv(p, false, true)
	tc.SleepWithLog(10, "sgn being ready")
}

func TestCbridge(t *testing.T) {
	t.Run("e2e-cbridge", func(t *testing.T) {
		t.Run("cbridgeTest", cbridgeTest)
	})
}

// Test cbridge
func cbridgeTest(t *testing.T) {
	log.Infoln("===================================================================")
	log.Infoln("======================== Test cbridge ===========================")
	setupCbridge()

	transactor := tc.NewTestTransactor(
		tc.SgnHomes[0],
		tc.SgnChainID,
		tc.SgnNodeURI,
		tc.ValSgnAddrStrs[0],
		tc.SgnPassphrase,
	)

	log.Infoln("================== Setup validators and bridge signers ======================")
	amts := []*big.Int{big.NewInt(3e18)}
	SetupValidators(transactor, amts)
	tc.CbrClient1.SetInitSigners(amts)
	tc.CbrClient2.SetInitSigners(amts)

	log.Infoln("======================== Add liquidity on chain 1 ===========================")
	addAmt := big.NewInt(5 * 1e10)
	err := tc.CbrClient1.Approve(addAmt)
	tc.ChkErr(err, "client1 approve")
	err = tc.CbrClient1.AddLiq(addAmt)
	tc.ChkErr(err, "client1 addliq")
	tc.CheckAddLiquidityStatus(transactor, tc.ChainID, 1)

	log.Infoln("======================== Add liquidity on chain 2 ===========================")
	err = tc.CbrClient2.Approve(addAmt)
	tc.ChkErr(err, "client2 approve")
	err = tc.CbrClient2.AddLiq(addAmt)
	tc.ChkErr(err, "client2 addliq")
	tc.CheckAddLiquidityStatus(transactor, tc.Geth2ChainID, 1)

	log.Infoln("======================== Xfer ===========================")
	xferAmt := big.NewInt(1e10)
	err = tc.CbrClient1.Approve(xferAmt)
	tc.ChkErr(err, "client1 approve")
	err = tc.CbrClient1.Send(xferAmt, tc.ValEthAddrs[0], tc.Geth2ChainID, 1)
	tc.ChkErr(err, "client1 send")
	// TODO: to check xferid generation rule
	// xferId := generateXferId(tc.CbrClient1.Auth.From, tc.ValEthAddrs[0], tc.CbrClient1.USDTAddr, xferAmt, int64(tc.Geth2ChainID), 1, int64(tc.ChainID))
	// tc.CheckXfer(transactor, xferId)
}

func generateXferId(sender, receiver, token eth.Addr, amt *big.Int, dstChainId, nonce, srcChainId int64) []byte {
	var b []byte
	b = append(b, sender[:]...)
	b = append(b, receiver[:]...)
	b = append(b, token[:]...)
	b = append(b, toPadBytes(amt)...)
	b = append(b, toPadBytes(dstChainId)...)
	b = append(b, toPadBytes(nonce)...)
	b = append(b, toPadBytes(srcChainId)...)

	return crypto.Keccak256(b)
}

// ToPadBytes return big-endian/network order bytes, left padded to specific length
// if v is uint32: 4 bytes, int64: 8 bytes, *big.Int: 32 bytes or rlen bytes if set
// return nil if type not supported
func toPadBytes(v interface{}, rlen ...int) []byte {
	var orig []byte
	var retlen int
	switch k := v.(type) {
	case uint32:
		retlen = 4
		orig = big.NewInt(int64(k)).Bytes()
	case int64:
		retlen = 8
		orig = big.NewInt(k).Bytes()
	case *big.Int:
		if len(rlen) == 1 {
			retlen = rlen[0]
		} else {
			retlen = 32
		}
		orig = k.Bytes()
	default:
		return nil
	}
	ret := make([]byte, retlen)
	copy(ret[retlen-len(orig):], orig)
	return ret
}
