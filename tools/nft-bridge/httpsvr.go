package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	"github.com/julienschmidt/httprouter"
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
	// add cors support
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow")) // header["Allow"] is set by httprouter
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "*")
			header.Set("Access-Control-Max-Age", "7200") // 2hrs
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Error("listen err: ", err)
	}
}

type HistResp struct {
	History  []dal.Nftxfer `json:"history"`
	NextPage int64         `json:"nextPageToken"`
	PageSize int           `json:"pageSize"`
}

func (s *Server) History(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usr := a2hex(hex2addr(ps.ByName("usr"))) // clean up user addr
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
