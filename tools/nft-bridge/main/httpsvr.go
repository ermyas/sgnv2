package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	"github.com/ethereum/go-ethereum/common"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

const (
	paramNextPage = "nextPageToken"
	paramPageSize = "pageSize"
)

type Server struct {
	db *dal.DAL
}

func NewServer(dal *dal.DAL) *Server {
	return &Server{
		db: dal,
	}
}

// listen and block run
func (s *Server) Run(port int) {
	router := httprouter.New()
	router.GET("/nftbr/history/:usr", s.History)
	router.GET("/nftbr/own/:usr/:chid/:nft", s.Own)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), cors.AllowAll().Handler(router)); err != nil {
		log.Error("listen err: ", err)
	}
}

type HistResp struct {
	History  []dal.Nftxfer `json:"history"`
	NextPage int64         `json:"nextPageToken"`
	PageSize int           `json:"pageSize"`
}

func cleanup(addrStr string) string {
	addr := common.HexToAddress(addrStr)
	return hex.EncodeToString(addr[:])
}

func (s *Server) Own(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usr := cleanup(ps.ByName("usr")) // clean up user addr
	nft := cleanup(ps.ByName("nft"))
	chid, err := strconv.Atoi(ps.ByName("chid"))
	if err != nil {
		http.Error(w, "invalid chain id "+ps.ByName("chid"), http.StatusBadRequest)
		return
	}
	has, err := s.db.UsrGetNfts(context.Background(), dal.UsrGetNftsParams{
		Chid: uint64(chid),
		Nft:  nft,
		Usr:  usr,
	})
	// no rows is not an error
	_, err2 := dal.ChkQueryRow(err)
	if err2 != nil {
		http.Error(w, "internal db err: "+err2.Error(), http.StatusInternalServerError)
		return
	}
	if len(has) == 0 {
		w.Write([]byte("[]"))
	} else {
		raw, _ := json.Marshal(has)
		w.Write(raw)
	}
}

func (s *Server) History(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usr := cleanup(ps.ByName("usr")) // clean up user addr
	nextPage, _ := strconv.Atoi(r.URL.Query().Get(paramNextPage))
	if nextPage == 0 {
		nextPage = int(time.Now().Unix()) // query all records so far
	}
	pageSize, _ := strconv.Atoi(r.URL.Query().Get(paramPageSize))
	if pageSize == 0 {
		pageSize = 5
	}
	xfers, _ := s.db.NftGetBySender(context.Background(), dal.NftGetBySenderParams{
		Sender:    usr,
		CreatedAt: int64(nextPage),
		Limit:     int32(pageSize),
	})
	resp := HistResp{
		PageSize: pageSize,
	}
	if len(xfers) == 0 {
		resp.History = make([]dal.Nftxfer, 0) // so json has [] instead of null
		// next is default 0
	} else {
		resp.History = xfers
		resp.NextPage = xfers[len(xfers)-1].CreatedAt // last one is smallest because we order by desc
	}
	raw, _ := json.Marshal(resp)
	w.Write(raw)
}
