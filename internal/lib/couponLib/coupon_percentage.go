package couponLib

import (
	"fmt"

	"kart-challenge/internal/model"

	"github.com/shopspring/decimal"
)

// var defaultDiscountPercentage = decimal.NewFromInt32(18).Shift(-2)
type PercentageCoupon struct {
	percentage decimal.Decimal
	Code       string
}

func (h *PercentageCoupon) Init(data CouponData) {
	h.percentage = data.Value
	h.Code = data.Code
}

func (h *PercentageCoupon) CalculateDiscount(cart model.Cart) CouponResult {
	var subtotal decimal.Decimal
	for _, item := range cart.Items {
		subtotal.Add(item.PriceAtAddition.Mul(decimal.NewFromUint64(uint64(item.Quantity))))
	}

	return CouponResult{
		DiscountAmount: subtotal.Mul(h.percentage),
		DiscountType:   "percentage",
		Message:        fmt.Sprintf("%s applied! %v off your cart total.", h.Code, h.percentage),
	}
}

func (h *PercentageCoupon) GetInfo() string {
	return fmt.Sprintf("coupon code: %s ! %v percentage off your cart total.", h.Code, h.percentage)
}
