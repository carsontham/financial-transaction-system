package bootstrap

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	address string
	router  chi.Router
}

func NewServer(address string) *Server {
	return &Server{
		address: address,
		router:  chi.NewRouter(),
	}
}

func (s *Server) RunServer() {
	fmt.Printf("Server running on port %s\n", s.address)
	log.Fatal(http.ListenAndServe(s.address, s.router))
}
