package model

import "time"

// Cart represents a shopping cart
type Cart struct {
	ID        uint64     `json:"id" gorm:"primaryKey"`
	DeviceID  string     `json:"device_id" gorm:"device_id"`
	CreatedAt time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"updated_at"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`
}
