package handler

import (
	"net/http"
	"strconv"

	"kart-challenge/internal/database"
	"kart-challenge/internal/lib/couponLib"
	"kart-challenge/internal/meta"

	redis "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

// OrderHandler handles order related requests
type OrderHandler struct {
	DB    *database.DBClient
	Redis *redis.Client
}

// PlaceOrder handles the POST /orders endpoint
func (h *OrderHandler) PlaceOrder(c echo.Context) error {
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	var req struct {
		CouponCode string `json:"coupon_code"`
		CartID     uint64 `json:"cart_id" validate:"required,min=1"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	var coupon couponLib.Coupon
	// Placeholder for coupon code validation
	if req.CouponCode != "" {
		_, key := getRedisKeyByCouponCode(req.CouponCode)
		cnt, err := h.Redis.BitCount(c.Request().Context(), key, &redis.BitCount{}).Result()
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

		coupon = couponLib.CreateCoupon(couponType, couponData)
	}

	cart, err := h.DB.GetCartByID(req.CartID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cart == nil || len(cart.Items) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Cart is empty")
	}

	order, err := h.DB.CreateOrder(cart.ID, deviceID)
	if err != nil {
		// Check for specific errors like out of stock
		if err.Error() == "cart is empty" { // This check should be redundant because of the above check
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	couponResult := couponLib.CouponResult{}
	if coupon != nil {
		couponResult = coupon.CalculateDiscount(*cart)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"order":         order,
		"coupon_result": couponResult,
	})
}

// GetOrders handles the GET /orders endpoint
func (h *OrderHandler) GetOrders(c echo.Context) error {
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	orders, err := h.DB.GetOrdersByDeviceID(deviceID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, orders)
}

// GetOrder handles the GET /orders/{order_id} endpoint
func (h *OrderHandler) GetOrder(c echo.Context) error {
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	orderID, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid order ID")
	}

	order, err := h.DB.GetOrderByIDAndDeviceID(orderID, deviceID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if order == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Order not found or does not belong to this device")
	}

	return c.JSON(http.StatusOK, order)
}
