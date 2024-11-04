package server

import (
	"context"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateSession method is
func (s *Server) CreateSession(ctx context.Context, se *models.Session) (*models.Session, error) {
	return s.storer.CreateSession(ctx, se)
}

// GetSession method is
func (s *Server) GetSession(ctx context.Context, id string) (*models.Session, error) {
	return s.storer.GetSession(ctx, id)
}

// RevokeSession method is
func (s *Server) RevokeSession(ctx context.Context, id string) error {
	return s.storer.RevokeSession(ctx, id)
}

// DeleteSession method is
func (s *Server) DeleteSession(ctx context.Context, id string) error {
	return s.storer.DeleteSession(ctx, id)
}
