package handler

import "time"

// OrderReq is
type OrderReq struct {
	Items         []OrderItem `json:"items"`
	PaymentMethod string      `json:"payment_method"`
	TaxPrice      float64     `json:"tax_price"`
	ShippingPrice float64     `json:"shipping_price"`
	TotalPrice    float64     `json:"total_price"`
}

// OrderItem is
type OrderItem struct {
	Name      string  `json:"name"`
	Quantity  int64   `json:"quantity"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
	ProductID int64   `json:"product_id"`
}

// OrderRes is
type OrderRes struct {
	ID            int64       `json:"id"`
	Items         []OrderItem `json:"items"`
	PaymentMethod string      `json:"payment_method"`
	TaxPrice      float64     `json:"tax_price"`
	ShippingPrice float64     `json:"shipping_price"`
	TotalPrice    float64     `json:"total_price"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     *time.Time  `json:"updated_at"`
}
