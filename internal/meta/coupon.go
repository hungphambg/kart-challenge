package meta

type CouponType string

const (
	CouponTypePercentage CouponType = "percentage"
	CouponTypeFreeItem   CouponType = "free-item"

	CouponTypeUnknown CouponType = "unknown"
)
