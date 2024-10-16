package storer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateOrder method is
func (ms *MySQLStorer) CreateOrder(ctx context.Context, o *models.Order) (*models.Order, error) {
	// start a transaction
	if err := ms.ExecTx(ctx, func(tx *sqlx.Tx) error {
		// insert into orders
		order, err := createOrder(ctx, tx, o)
		if err != nil {
			return fmt.Errorf("error creating order")
		}

		for _, oi := range order.Items {
			oi.OrderID = order.ID
			// insert into order items
			err = createOrderItem(ctx, tx, oi)
			if err != nil {
				return fmt.Errorf("error creating order item")
			}
		}

		return nil
	}); err != nil {
		return nil, nil
	}

	// commit transaction
	// rollback transaction if error occur
	return nil, nil
}

// createOrder method is
func createOrder(ctx context.Context, tx *sqlx.Tx, o *models.Order) (*models.Order, error) {
	res, err := tx.NamedExecContext(
		ctx,
		"INSERT INTO orders (payment_method, tax_price, shipping_price, total_price) VALUES (:payment_method, :tax_price, :shipping_price, :total_price)",
		o,
	)
	if err != nil {
		return nil, fmt.Errorf("error insert order: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %w", err)
	}
	o.ID = id

	return o, nil
}

// createOrderItem method is
func createOrderItem(ctx context.Context, tx *sqlx.Tx, oi models.OrderItem) error {
	res, err := tx.NamedExecContext(
		ctx,
		"INSERT INTO order_items (name, quantity, image, price, product_id, order_id) VALUES (:name, :quantity, :image, :price, :product_id, :order_id)",
		oi,
	)
	if err != nil {
		return fmt.Errorf("error insert order item: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert id: %w", err)
	}
	oi.ID = id

	return nil
}
