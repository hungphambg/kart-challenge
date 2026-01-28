package model

import "github.com/shopspring/decimal"

const OrderItemTableName = "order_item"

// OrderItem represents an item within an order
type OrderItem struct {
	ID        uint64          `json:"id" gorm:"primaryKey"`
	OrderID   uint64          `json:"order_id"`
	ProductID uint64          `json:"product_id"`
	Quantity  uint32          `json:"quantity"`
	UnitPrice decimal.Decimal `json:"unit_price"` // Price at the time of order
}

func (OrderItem) TableName() string {
	return OrderItemTableName
}
