package model

import (
	"time"

	"kart-challenge/internal/meta"

	"github.com/shopspring/decimal"
)

const (
	OrderTableName = "order"
)

// Order represents a customer order
type Order struct {
	ID          uint64           `json:"id" gorm:"primaryKey"`
	DeviceID    string           `json:"device_id"`
	Status      meta.OrderStatus `json:"status"` // e.g., "pending", "completed", "cancelled"
	TotalAmount decimal.Decimal  `json:"total_amount"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	Items       []OrderItem      `json:"items" gorm:"foreignKey:OrderID"`
}

func (Order) TableName() string {
	return OrderTableName
}
