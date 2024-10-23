package handler

import (
	"context"

	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/server"
)

// Handler struct is
type Handler struct {
	ctx    context.Context
	server *server.Server
}

// NewHandler is
func NewHandler(ctx context.Context, server *server.Server) *Handler {
	return &Handler{
		ctx:    ctx,
		server: server,
	}
}
