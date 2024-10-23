package server

import (
	"context"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateProduct is
func (s *Server) CreateProduct(ctx context.Context, p *models.Product) (*models.Product, error) {
	return s.storer.CreateProduct(ctx, p)
}

// GetProduct is
func (s *Server) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	return s.storer.GetProduct(ctx, id)
}

// ListProducts is
func (s *Server) ListProducts(ctx context.Context) ([]models.Product, error) {
	return s.storer.ListProducts(ctx)
}

// UpdateProduct is
func (s *Server) UpdateProduct(ctx context.Context, p *models.Product) (*models.Product, error) {
	return s.storer.UpdateProduct(ctx, p)
}

// DeleteProduct is
func (s *Server) DeleteProduct(ctx context.Context, id int64) error {
	return s.storer.DeleteProduct(ctx, id)
}
