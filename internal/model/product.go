package model

import "github.com/shopspring/decimal"

// Product represents a product in the system
type Product struct {
	ID          uint64          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       decimal.Decimal `json:"price"`
	ImageURL    string          `json:"image_url"`
	Stock       uint32          `json:"stock"`
}
