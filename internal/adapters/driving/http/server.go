package http

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(handler *OrderHandler) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", handler.CreateOrder)
	return &Server{mux: mux}
}

func (s *Server) Start(addr string) {
	log.Println("Starting server on", addr)
	log.Fatal(http.ListenAndServe(addr, s.mux))
}