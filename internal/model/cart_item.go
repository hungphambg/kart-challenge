package model

import "github.com/shopspring/decimal"

const CartItemTableName = "cart_item"

// CartItem represents an item in a shopping cart
type CartItem struct {
	ID              uint64          `json:"id" gorm:"primaryKey"`
	CartID          uint64          `json:"cart_id"`
	ProductID       uint64          `json:"product_id"`
	Quantity        uint32          `json:"quantity"`
	PriceAtAddition decimal.Decimal `json:"price_at_addition"`
}

func (CartItem) TableName() string {
	return CartItemTableName
}
