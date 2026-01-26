package handler

import (
	"log"
	"net/http"

	"kart-challenge/internal/database"
	couponLib "kart-challenge/internal/lib/couponLib"
	"kart-challenge/internal/meta"
	"kart-challenge/internal/model"

	redis "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

// CouponHandler handles couponLib related requests
type CouponHandler struct {
	DB    *database.DBClient
	Redis *redis.Client
}

// ValidateCoupon handles the POST /coupons/validate endpoint
func (h *CouponHandler) ValidateCoupon(c echo.Context) error {
	var req struct {
		CouponCode string `json:"coupon_code"`
		DeviceID   string `json:"device_id"` // Optional, for cart-specific validation
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if req.CouponCode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Coupon code is required")
	}

	// 1. Check Redis for quick validity lookup
	redisKey := "valid_coupon_codes" // This key should match what the consumer uses
	isMember, err := h.Redis.SIsMember(c.Request().Context(), redisKey, req.CouponCode).Result()
	if err != nil {
		log.Printf("Error checking Redis for couponLib code %s: %v", req.CouponCode, err)
		// Fallback to database or return internal server error if Redis is critical
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to check couponLib validity")
	}

	if !isMember {
		return echo.NewHTTPError(http.StatusNotFound, "Coupon code not found or is not valid")
	}

	var (
		couponData = couponLib.CouponData{
			Code: req.CouponCode,
		}
		couponType = meta.CouponTypeUnknown
	)
	// 2. If valid in Redis, retrieve couponLib details from the database (for type, value, etc.)
	couponDb, err := h.DB.GetValidCouponByCode(req.CouponCode)
	if err != nil {
		log.Printf("Error retrieving couponLib details from DB for code %s: %v", req.CouponCode, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve couponLib details")
	}
	// This case should ideally not happen if Redis is in sync with DB active coupons,
	// but as a safeguard, if DB doesn't find it, it's truly invalid.
	if couponDb != nil {
		couponData.Value = couponDb.Value
		couponType = couponDb.Type
	}

	var cart *model.Cart
	if req.DeviceID != "" {
		cart, err = h.DB.GetCartByDeviceID(req.DeviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	if cart == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Cart not found")
	}

	var coupon = couponLib.CreateCoupon(couponType, couponData)
	couponResult := coupon.CalculateDiscount(*cart)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"coupon":          couponDb,
		"status":          "valid",
		"cart_id":         cart.ID,
		"discount_amount": couponResult.DiscountAmount,
		"discount_type":   couponResult.DiscountType,
		"message":         couponResult.Message,
	})
}
