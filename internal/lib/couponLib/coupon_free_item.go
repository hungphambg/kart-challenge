package couponLib

import (
	"fmt"

	"kart-challenge/internal/model"

	"github.com/shopspring/decimal"
)

// var defaultDiscountPercentage = decimal.NewFromInt32(18).Shift(-2)
type FreeItemCoupon struct {
	Code string
}

func (h *FreeItemCoupon) Init(data CouponData) {
	h.Code = data.Code
}

func (h *FreeItemCoupon) CalculateDiscount(cart model.Cart) CouponResult {
	lowestPrice := decimal.Zero
	lowestPriceItemIndex := -1

	// Find the lowest priced item with available quantity
	for i, item := range cart.Items {
		if item.Quantity > 0 { // Only consider items that are actually in the cart
			if lowestPriceItemIndex == -1 || item.PriceAtAddition.LessThan(lowestPrice) {
				lowestPrice = item.PriceAtAddition
				lowestPriceItemIndex = i
			}
		}
	}

	var result = CouponResult{
		DiscountType: "free_item",
	}

	if lowestPriceItemIndex != -1 {
		result.DiscountAmount = lowestPrice // Make one instance of the lowest priced item free
		result.Message = fmt.Sprintf("%s applied! One item of price %v is free.", h.Code, lowestPrice)
	} else {
		result.Message = fmt.Sprintf("%s not applicable: No items in cart with quantity.", h.Code)
	}

	return result
}
