package model

import (
	"time"

	"kart-challenge/internal/meta"
)

const (
	CartTableName     = "cart"
	CartColNameStatus = "status"
)

// Cart represents a shopping cart
type Cart struct {
	ID        uint64          `json:"id" gorm:"column:id;primaryKey"`
	DeviceID  string          `json:"device_id" gorm:"column:device_id"`
	Status    meta.CartStatus `json:"status" gorm:"column:status"`
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	Items     []CartItem      `json:"items" gorm:"foreignKey:CartID"`
}

func (Cart) TableName() string {
	return CartTableName
}
