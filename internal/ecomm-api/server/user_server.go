package server

import (
	"context"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateUser method is
func (s *Server) CreateUser(ctx context.Context, u *models.User) (*models.User, error) {
	return s.storer.CreateUser(ctx, u)
}

// GetUser method is
func (s *Server) GetUser(ctx context.Context, email string) (*models.User, error) {
	return s.storer.GetUser(ctx, email)
}

// ListUsers method is
func (s *Server) ListUsers(ctx context.Context) ([]models.User, error) {
	return s.storer.ListUsers(ctx)
}

// UpdateUser method is
func (s *Server) UpdateUser(ctx context.Context, u *models.User) (*models.User, error) {
	return s.storer.UpdateUser(ctx, u)
}

// DeleteUser method is
func (s *Server) DeleteUser(ctx context.Context, id int64) error {
	return s.storer.DeleteUser(ctx, id)
}
