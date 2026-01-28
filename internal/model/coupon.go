package model

import (
	"time"

	"kart-challenge/internal/meta"

	"github.com/shopspring/decimal"
)

const (
	CouponTableName = "coupon"
)

// Coupon represents a discount couponLib
type Coupon struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	Code        string          `json:"code" gorm:"unique;not null"`
	Type        meta.CouponType `json:"type" gorm:"not null"`  // e.g., "percentage", "buy_one_get_one", "fixed_amount"
	Value       decimal.Decimal `json:"value" gorm:"not null"` // e.g., 0.18 for 18% off, or 5.00 for $5 off
	Description string          `json:"description"`
	ExpiresAt   *time.Time      `json:"expires_at"` // Nullable timestamp
	IsActive    bool            `json:"is_active" gorm:"not null;default:true"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

func (Coupon) TableName() string {
	return CouponTableName
}
