package model

import "github.com/shopspring/decimal"

const ProductTableName = "product"

// Product represents a product in the system
type Product struct {
	ID           uint64          `json:"id" gorm:"id"`
	MerchantCode string          `json:"-" gorm:"merchant_code"`
	Name         string          `json:"name" gorm:"name"`
	Description  string          `json:"description" gorm:"description"`
	Price        decimal.Decimal `json:"price" gorm:"price"`
	ImageURL     string          `json:"image_url" gorm:"image_url"`
	Stock        uint32          `json:"stock" gorm:"stock"`
}

func (Product) TableName() string {
	return ProductTableName
}
