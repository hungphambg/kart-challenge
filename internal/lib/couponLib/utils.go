package couponLib

import (
	"kart-challenge/internal/meta"
)

func CreateCoupon(couponType meta.CouponType, data CouponData) Coupon {
	var coupon Coupon
	switch couponType {
	case meta.CouponTypePercentage:
		coupon = &PercentageCoupon{}
	case meta.CouponTypeFreeItem:
		coupon = &FreeItemCoupon{}
	default:
		coupon = &DefaultCoupon{}
	}
	coupon.Init(data)
	return coupon
}
