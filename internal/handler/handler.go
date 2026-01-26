package handler

import (
	"net/http"
	"strconv"

	"kart-challenge/internal/database"

	"github.com/labstack/echo/v4"
)

// ProductHandler handles product related requests
type ProductHandler struct {
	DB *database.DBClient
}

// GetProducts handles the GET /products endpoint
func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.DB.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}

// GetProduct handles the GET /products/{id} endpoint
func (h *ProductHandler) GetProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid product ID")
	}

	product, err := h.DB.GetProductByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if product == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Product not found")
	}

	return c.JSON(http.StatusOK, product)
}
