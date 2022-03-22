package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/sgn-v2/tools/nft-bridge/dal"
	"github.com/julienschmidt/httprouter"
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

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Error("listen err: ", err)
	}
}

func (s *Server) History(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usr := ps.ByName("usr")
	usr = a2hex(hex2addr(usr)) // clean up user addr
	xfers, _ := s.db.NftGetBySender(context.Background(), usr)
	if len(xfers) == 0 {
		w.Write([]byte("[]"))
		return
	}
	raw, _ := json.Marshal(xfers)
	w.Write(raw)
}
