package handler

import (
	"context"

	"github.com/jumayevgadam/ecomm_mysql/internal/ecomm-api/server"
	"github.com/jumayevgadam/ecomm_mysql/internal/middleware/token"
)

// Handler struct is
type Handler struct {
	ctx        context.Context
	server     *server.Server
	tokenMaker *token.JWTMaker
}

// NewHandler is
func NewHandler(server *server.Server, secretKey string) *Handler {
	return &Handler{
		ctx:        context.Background(),
		server:     server,
		tokenMaker: token.NewJWTMaker(secretKey),
	}
}
