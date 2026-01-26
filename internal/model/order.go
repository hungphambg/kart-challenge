package model

import (
	"time"

	"github.com/shopspring/decimal"
)

// Order represents a customer order
type Order struct {
	ID          uint64          `json:"id" gorm:"primaryKey"`
	DeviceID    string          `json:"device_id"`
	Status      string          `json:"status"` // e.g., "pending", "completed", "cancelled"
	TotalAmount decimal.Decimal `json:"total_amount"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Items       []OrderItem     `json:"items" gorm:"foreignKey:OrderID"`
}
