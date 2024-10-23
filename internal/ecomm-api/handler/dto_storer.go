package handler

import (
	"time"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// toStorerProduct is
func toStorerProduct(p ProductReq) *models.Product {
	return &models.Product{
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}
}

// toProductRes is
func toProductRes(p *models.Product) ProductRes {
	return ProductRes{
		ID:           p.ID,
		Name:         p.Name,
		Image:        p.Image,
		Category:     p.Category,
		Description:  p.Description,
		Rating:       p.Rating,
		NumReviews:   p.NumReviews,
		Price:        p.Price,
		CountInStock: p.CountInStock,
	}
}

// patchProductReq is
func patchProductReq(product *models.Product, p ProductReq) {
	if p.Name != "" {
		product.Name = p.Name
	}

	if p.Image != "" {
		product.Image = p.Image
	}

	if p.Category != "" {
		product.Category = p.Category
	}

	if p.Description != "" {
		product.Description = p.Description
	}

	if p.Rating != 0 {
		product.Rating = p.Rating
	}

	if p.NumReviews != 0 {
		product.NumReviews = p.NumReviews
	}

	if p.Price != 0 {
		product.Price = p.Price
	}

	if p.CountInStock != 0 {
		product.CountInStock = p.CountInStock
	}
	product.UpdatedAt = toTimePtr(time.Now())
}

// toTimePtr is
func toTimePtr(t time.Time) *time.Time {
	return &t
}

// toStorerOrder is
func toStorerOrder(o OrderReq) *models.Order {
	return &models.Order{
		PaymentMethod: o.PaymentMethod,
		TaxPrice:      o.TaxPrice,
		ShippingPrice: o.ShippingPrice,
		TotalPrice:    o.TotalPrice,
		Items:         toStorerOrderItems(o.Items),
	}
}

// toStorerOrderItems is
func toStorerOrderItems(items []OrderItem) []models.OrderItem {
	var res []models.OrderItem
	for _, i := range items {
		res = append(res, models.OrderItem{
			Name:      i.Name,
			Quantity:  i.Quantity,
			Image:     i.Image,
			Price:     i.Price,
			ProductID: i.ProductID,
		})
	}

	return res
}

// toOrderRes is
func toOrderRes(o *models.Order) OrderRes {
	return OrderRes{
		ID:            o.ID,
		Items:         toOrderItems(o.Items),
		PaymentMethod: o.PaymentMethod,
		TaxPrice:      o.TaxPrice,
		ShippingPrice: o.ShippingPrice,
		TotalPrice:    o.TotalPrice,
		CreatedAt:     o.CreatedAt,
		UpdatedAt:     o.UpdatedAt,
	}
}

// toOrderItems is
func toOrderItems(items []models.OrderItem) []OrderItem {
	var res []OrderItem
	for _, i := range items {
		res = append(res, OrderItem{
			Name:      i.Name,
			Quantity:  i.Quantity,
			Image:     i.Image,
			Price:     i.Price,
			ProductID: i.ProductID,
		})
	}

	return res
}
