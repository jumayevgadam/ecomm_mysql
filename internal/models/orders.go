package models

import "time"

// Order model is
type Order struct {
	ID            int64      `db:"id"`
	PaymentMethod string     `db:"payment_method"`
	TaxPrice      float64    `db:"tax_price"`
	ShippingPrice float64    `db:"shipping_price"`
	TotalPrice    float64    `db:"total_price"`
	CreatedAt     time.Time  `db:"created_at"`
	UpdatedAt     *time.Time `db:"updated_at"`
	Items         []OrderItem
}

// OrderItem model is
type OrderItem struct {
	ID        int64   `db:"id"`
	Name      string  `db:"name"`
	Quantity  int64   `db:"quantity"`
	Image     string  `db:"image"`
	Price     float64 `db:"price"`
	ProductID int64   `db:"product_id"`
	OrderID   int64   `db:"order_id"`
}
