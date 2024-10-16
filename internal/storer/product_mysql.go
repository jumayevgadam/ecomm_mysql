package storer

import (
	"context"
	"fmt"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateProduct method is
func (ms *MySQLStorer) CreateProduct(ctx context.Context, p *models.Product) (*models.Product, error) {
	res, err := ms.DB.NamedExecContext(
		ctx,
		"INSERT INTO products (name, image, category, description, rating, num_reviews, price, count_in_stock) VALUES (:name, :image, :category, :description, :rating, :num_reviews, :price, :count_in_stock)",
		p,
	)
	if err != nil {
		return nil, fmt.Errorf("error inserting to product: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}
	p.ID = id

	return p, nil
}

func (ms *MySQLStorer) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	var p models.Product
	if err := ms.DB.GetContext(
		ctx,
		&p,
		"SELECT * FROM products WHERE id=?",
		id,
	); err != nil {
		return nil, fmt.Errorf("error getting product: %w", err)
	}

	return &p, nil
}

func (ms *MySQLStorer) ListProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	if err := ms.DB.SelectContext(
		ctx,
		&products,
		"SELECT * FROM products",
	); err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}

	return products, nil
}

func (ms *MySQLStorer) UpdateProduct(ctx context.Context, p *models.Product) (*models.Product, error) {
	_, err := ms.DB.NamedExecContext(
		ctx,
		"UPDATE products SET name=:name, image=:image, category=:category, description=:description, rating=:rating, num_reviews=:num_reviews WHERE id=:id",
		p,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}

	return p, nil
}

func (ms *MySQLStorer) DeleteProduct(ctx context.Context, id int64) error {
	_, err := ms.DB.ExecContext(
		ctx,
		"DELETE FROM products WHERE id=?",
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}
