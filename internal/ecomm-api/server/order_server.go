package server

import (
	"context"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateOrder is
func (s *Server) CreateOrder(ctx context.Context, o *models.Order) (*models.Order, error) {
	return s.storer.CreateOrder(ctx, o)
}

// GetOrder is
func (s *Server) GetOrder(ctx context.Context, id int64) (*models.Order, error) {
	return s.storer.GetOrder(ctx, id)
}

// ListOrders is
func (s *Server) ListOrders(ctx context.Context) ([]models.Order, error) {
	return s.storer.ListOrders(ctx)
}

// DeleteOrder is
func (s *Server) DeleteOrder(ctx context.Context, id int64) error {
	return s.storer.DeleteOrder(ctx, id)
}
