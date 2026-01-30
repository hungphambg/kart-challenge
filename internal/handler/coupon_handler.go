package handler

import (
	"fmt"
	"hash/crc32"
	"net/http"

	"kart-challenge/internal/database"
	"kart-challenge/internal/lib/couponLib"
	"kart-challenge/internal/meta"

	redis "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

// CouponHandler handles couponLib related requests
type CouponHandler struct {
	DB    *database.DBClient
	Redis *redis.Client
}

func ShardOf(promo string) int {
	return int(crc32.ChecksumIEEE([]byte(promo)) % 64)
}

func getRedisKeyByCouponCode(code string) (shard int, key string) {
	shard = ShardOf(code)
	key = fmt.Sprintf("promo:%d:%s", shard, code)
	return
}

// ValidateCoupon handles the POST /coupons/validate endpoint
func (h *CouponHandler) ValidateCoupon(c echo.Context) error {
	var req struct {
		CouponCode string `json:"coupon_code"`
		DeviceID   string `json:"device_id"` // Optional, for cart-specific validation
	}
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if req.CouponCode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Coupon code is required")
	}
	// 1. Check Redis for quick validity lookup
	_, key := getRedisKeyByCouponCode(req.CouponCode)

	cnt, err := h.Redis.BitCount(ctx, key, &redis.BitCount{}).Result()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cnt < 2 {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid or expired coupon code")
	}

	couponDb, err := h.DB.GetCouponByCode(req.CouponCode)
	if err != nil {
		return err
	}
	var (
		couponData = couponLib.CouponData{
			Code:  req.CouponCode,
			Value: decimal.Zero,
		}

		couponType = meta.CouponTypeUnknown
	)
	if couponDb != nil {
		couponType = couponDb.Type
		couponData.Value = couponDb.Value
	}

	coupon := couponLib.CreateCoupon(couponType, couponData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "valid",
		"info":   coupon.GetInfo(),
	})
}
