package handler

import (
	"net/http"
	"strconv"
	"time"

	"kart-challenge/internal/database"
	"kart-challenge/internal/model"

	"github.com/labstack/echo/v4"
)

// CartHandler handles cart related requests
type CartHandler struct {
	DB *database.DBClient
}

// GetCart handles the GET /cart endpoint
func (h *CartHandler) GetCart(c echo.Context) error {
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	cart, err := h.DB.GetCartByDeviceID(deviceID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cart == nil {
		cart, err = h.DB.CreateCart(deviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, cart)
}

// CreateCart handles the POST /cart/create endpoint
func (h *CartHandler) CreateCart(c echo.Context) error {
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	cart, err := h.DB.CreateCart(deviceID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, cart)
}

// AddItemToCart handles the POST /cart/items endpoint
func (h *CartHandler) AddItemToCart(c echo.Context) error {
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	var req struct {
		ID        uint64 `json:"id" validate:"required,min=1"`
		ProductID uint64 `json:"product_id"`
		Quantity  uint32 `json:"quantity"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	cart, err := h.DB.GetCartByID(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cart == nil {
		cart, err = h.DB.CreateCart(deviceID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	product, err := h.DB.GetProductByID(req.ProductID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if product == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	if product.Stock < req.Quantity {
		return echo.NewHTTPError(http.StatusConflict, "Not enough stock")
	}

	err = h.DB.AddItemToCart(cart.ID, req.ProductID, req.Quantity, product.Price)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedCart, err := h.DB.GetCartByDeviceID(deviceID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedCart)
}

// ClearCart handles the DELETE /cart endpoint
func (h *CartHandler) ClearCart(c echo.Context) error {
	var req struct {
		ID uint64 `json:"id" validate:"required,min=1"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	cart, err := h.DB.GetCartByID(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if cart != nil {
		err = h.DB.ClearCart(cart.ID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	// Return empty cart representation
	emptyCart := &model.Cart{
		DeviceID: deviceID,
		Items:    []model.CartItem{},
	}
	if cart != nil {
		emptyCart.ID = cart.ID
		emptyCart.CreatedAt = cart.CreatedAt
		emptyCart.UpdatedAt = time.Now()
	}

	return c.JSON(http.StatusOK, emptyCart)
}

// RemoveCartItem handles the DELETE /cart/items/{product_id} endpoint

func (h *CartHandler) RemoveCartItem(c echo.Context) error {
	var req struct {
		ID uint64 `json:"id" validate:"required,min=1"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	deviceID := c.Request().Header.Get("DeviceID")
	if deviceID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "DeviceID header is required")
	}

	productID, err := strconv.ParseInt(c.Param("product_id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}

	cart, err := h.DB.GetCartByID(req.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if cart == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Cart not found")
	}

	err = h.DB.RemoveCartItem(cart.ID, productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedCart, err := h.DB.GetCartByDeviceID(deviceID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedCart)
}
