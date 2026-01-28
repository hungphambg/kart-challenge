package handler

import (
	"log"
	"net/http"
	"strconv"

	"kart-challenge/internal/database"

	redis "github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
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

	// Placeholder for coupon code validation
	if req.CouponCode != "" {
		// Check Redis for coupon validity
		redisKey := "valid_coupon_codes"
		isMember, err := h.Redis.SIsMember(c.Request().Context(), redisKey, req.CouponCode).Result()
		if err != nil {
			log.Printf("Error checking Redis for coupon code %s: %v", req.CouponCode, err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to validate coupon")
		}

		if !isMember {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid or expired coupon code")
		}
		// If valid in Redis, retrieve coupon details from the database to apply discount logic (not implemented yet for orders)
		// For now, just proceed with order if coupon is valid in Redis.
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

	return c.JSON(http.StatusCreated, order)
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
