package main

import (
	"context"
	"database/sql"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"time"

	gobig "github.com/celer-network/goutils/big"
	"github.com/celer-network/goutils/log"
	comtypes "github.com/celer-network/sgn-v2/common/types"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	cbrtypes "github.com/celer-network/sgn-v2/x/cbridge/types"
	msgtypes "github.com/celer-network/sgn-v2/x/message/types"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// query sgn grpc to get what to send onchain
// msgtypes.NewQueryClient()

// block polling sgn
func PollSgn(intv time.Duration, ntfbrs []*ChidAddr, chainMap map[uint64]*OneChain) {
	conn, err := grpc.Dial(viper.GetString("sgn"), grpc.WithInsecure())
	chkErr(err, "dial sgn")
	qc := msgtypes.NewQueryClient(conn)
	req := &msgtypes.QueryExecutionContextsRequest{
		ContractInfos: convert(ntfbrs),
	}
	ticker := time.NewTicker(intv)
	defer ticker.Stop()
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigch:
			log.Infoln("receive term signal", sig, "exiting")
			return

		case <-ticker.C:
			resp, err := qc.ExecutionContexts(context.Background(), req)
			if err != nil {
				log.Error("query sgn err: ", err)
			}
			for _, exeCtx := range resp.ExecutionContexts {
				msg := exeCtx.Message
				// msg.PrettyLog()
				if onech, ok := chainMap[msg.DstChainId]; ok {
					nftMsg := DecodeNFTMsg(msg.Data)
					log.Infoln("from:", msg.SrcChainId, "to:", msg.DstChainId, "nftMsg:", nftMsg)
					srcTx, _ := onech.db.NftGetByDstInfo(context.Background(), dal.NftGetByDstInfoParams{
						SrcChid:  msg.SrcChainId,
						DstChid:  msg.DstChainId,
						Receiver: a2hex(nftMsg.User),
						DstNft:   a2hex(nftMsg.Nft),
						TokID:    *gobig.New(nftMsg.Id),
						Status:   int16(Status_WAITSGN),
					})
					if srcTx == "" {
						// not found could be missed event or it's already sent onchain
						log.Infoln("msg not found in db, miss event or pending onchain tx")
						continue
					}
					// query sgn to get signers/powers, todo: cache by chid?
					signers, powers := getSigners(conn, msg.DstChainId)
					// send onchain
					tx, err := onech.msgBus.ExecuteMessage(onech.auth, msg.Data, MsgDataTypesRouteInfo{
						Sender:     hex2addr(msg.Sender),
						Receiver:   hex2addr(msg.Receiver),
						SrcChainId: msg.SrcChainId,
						SrcTxHash:  hex2hash(msg.SrcTxHash),
					}, msg.GetSigBytes(), signers, powers)
					if err != nil {
						log.Error("onchain exe msg err: ", err)
					} else {
						// update db
						dstTx := tx.Hash().Hex()
						log.Infoln("chainid:", onech.cfg.ChainID, "Sent ExecuteMessage tx:", dstTx)
						err := onech.db.DoTx(func(tx *sql.Tx) error {
							return dal.New(tx).NftSetDstTx(context.Background(), dal.NftSetDstTxParams{
								SrcChid:  msg.SrcChainId,
								DstChid:  msg.DstChainId,
								Receiver: a2hex(nftMsg.User),
								DstNft:   a2hex(nftMsg.Nft),
								TokID:    *gobig.New(nftMsg.Id),
								DstTx:    dstTx,
							})
						})
						if err != nil {
							log.Error("update dstTx err: ", err)
						}
					}
				}
			}
		}
	}
}

func getSigners(cc *grpc.ClientConn, dstChid uint64) (_signers []Addr, _powers []*big.Int) {
	qc := cbrtypes.NewQueryClient(cc)
	resp, err := qc.QueryChainSigners(context.Background(), &cbrtypes.QueryChainSignersRequest{
		ChainId: dstChid,
	})
	if err != nil {
		log.Error("QueryChainSigners err: ", err)
		return
	}
	return resp.ChainSigners.GetAddrsPowers()
}

// convert config to sgn.common.v1.ContractInfo
func convert(nftbrs []*ChidAddr) (ret []*comtypes.ContractInfo) {
	for _, one := range nftbrs {
		ret = append(ret, &comtypes.ContractInfo{
			ChainId: one.Chainid,
			Address: one.Addr,
		})
	}
	return
}

/*
func TestSgn(t *testing.T) {
	conn, err := grpc.Dial("cbridge-v2-test.celer.network:9090", grpc.WithInsecure())
	chkErr(err, "dial sgn")
	msgid := "0x1994042e3dd06f0feaf1045e656a498776ca1926a261b9c2cb0f98fc506662e7"
	qc := msgtypes.NewQueryClient(conn)
	resp, err := qc.Message(context.Background(), &msgtypes.QueryMessageRequest{
		MessageId: msgid,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Errorf("%x", resp.Message.Data)
	}
}
*/
