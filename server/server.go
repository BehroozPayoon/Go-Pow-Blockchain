package main

import (
	"io"
	"log"
	"net/http"
	"pow-blockchain/core"
	"pow-blockchain/wallet"
	"strconv"
)

var cache map[string]*core.Blockchain = make(map[string]*core.Blockchain)

type Server struct {
	port uint16
}

func NewServer(port uint16) *Server {
	return &Server{port}
}

func (s *Server) Port() uint16 {
	return s.port
}

func (s *Server) GetBlockchain() *core.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = core.NewBlockchain(minersWallet.Address(), s.Port())
		cache["blockchain"] = bc
		log.Printf("private_key %v", minersWallet.PrivateKeyStr())
		log.Printf("public_key %v", minersWallet.PublickKetStr())
		log.Printf("blockchain_address %v", minersWallet.Address())
	}
	return bc
}

func (s *Server) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := s.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (s *Server) Run() {
	http.HandleFunc("/", s.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(s.Port())), nil))
}
