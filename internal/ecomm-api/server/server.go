package server

import "github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/storer"

// Server is
type Server struct {
	storer *storer.MySQLStorer
}

// NewServer is
func NewServer(storer *storer.MySQLStorer) *Server {
	return &Server{storer: storer}
}
