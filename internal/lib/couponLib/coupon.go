package couponLib

import (
	"fmt"

	"kart-challenge/internal/model"

	"github.com/shopspring/decimal"
)

type CouponData struct {
	Code  string
	Value decimal.Decimal
}

type Coupon interface {
	Init(data CouponData)
	CalculateDiscount(cart model.Cart) CouponResult
	GetInfo() string
}

type CouponResult struct {
	DiscountAmount decimal.Decimal `json:"discount_amount"`
	DiscountType   string          `json:"discount_type"`
	Message        string          `json:"message"`
}

// just for allow unknow couponLib
type DefaultCoupon struct {
	Code string
}

func (h *DefaultCoupon) Init(data CouponData) {
	h.Code = data.Code
}

func (h *DefaultCoupon) CalculateDiscount(_ model.Cart) CouponResult {
	return CouponResult{
		DiscountType: "unknown",
		Message:      fmt.Sprintf("%s applied!", h.Code),
	}
}

func (h *DefaultCoupon) GetInfo() string {
	return fmt.Sprintf("%s infomation!", h.Code)
}
