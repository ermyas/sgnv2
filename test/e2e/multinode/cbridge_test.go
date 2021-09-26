package multinode

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/celer-network/goutils/log"
	tc "github.com/celer-network/sgn-v2/test/common"
	"github.com/celer-network/sgn-v2/transactor"
	bridgecli "github.com/celer-network/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
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

	log.Infoln("======================== Add liquidity on chain 1 ===========================")
	ctx, cancel := context.WithTimeout(context.Background(), tc.DefaultTimeout)
	defer cancel()

	tx, err := tc.Usdt1Contract.Approve(tc.ValAuths[0], tc.Cbr1Contract.Address, big.NewInt(5*1e6))
	tc.ChkErr(err, "failed to approve allowance")
	tc.WaitMinedWithChk(ctx, tc.EthClient, tx, tc.BlockDelay, tc.PollingInterval, "Approve")

	tx, err = tc.Cbr1Contract.Bridge.AddLiquidity(tc.ValAuths[0], tc.Usdt1Addr, big.NewInt(5*1e6))
	tc.ChkErr(err, "failed to add liquidity")
	tc.WaitMinedWithChk(ctx, tc.EthClient, tx, tc.BlockDelay, tc.PollingInterval, "AddLiquidity")

	checkAddLiquidityStatus(transactor, 883, 0)
}

func checkAddLiquidityStatus(transactor *transactor.Transactor, chainId, seqNum uint64) {
	var resp *cbrtypes.QueryLiquidityStatusResponse
	var err error
	for retry := 0; retry < tc.RetryLimit; retry++ {
		resp, err = bridgecli.QueryAddLiquidityStatus(transactor.CliCtx, &cbrtypes.QueryAddLiquidityStatusRequest{
			ChainId: 883,
			SeqNum:  0,
		})
		if err != nil {
			log.Debugln("retry due to err:", err)
		}
		if err == nil && resp.Status == cbrtypes.LPHistoryStatus_LP_COMPLETED {
			break
		}
		time.Sleep(tc.RetryPeriod)
	}
	tc.ChkErr(err, "failed to QueryAddLiquidityStatus")
	// TODO: status check
	// if resp.Status != cbrtypes.LPHistoryStatus_LP_COMPLETED {
	// 	log.Fatalln("incorrect status")
	// }
}
