package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"kart-challenge/internal/config"
	"kart-challenge/internal/database"
	"kart-challenge/internal/handler"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()

	dbClient, err := database.NewDBClient(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to initialize database client: %v", err)
	}
	defer func(dbClient *database.DBClient) {
		_ = dbClient.Close()
	}(dbClient)

	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Ping Redis to check connectivity
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	fmt.Println("Successfully connected to Redis!")

	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Handlers
	productHandler := &handler.ProductHandler{DB: dbClient}
	cartHandler := &handler.CartHandler{DB: dbClient}
	orderHandler := &handler.OrderHandler{DB: dbClient, Redis: redisClient}
	couponHandler := &handler.CouponHandler{DB: dbClient, Redis: redisClient} // Pass Redis client to couponLib handler
	//
	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(c echo.Context) error {
		if dbClient.Ping() != nil {
			return c.NoContent(http.StatusServiceUnavailable)
		}
		return c.NoContent(http.StatusOK)
	})

	e.GET("/api/product", productHandler.GetProducts)
	e.GET("/api/product/:id", productHandler.GetProduct)

	e.GET("/api/cart", cartHandler.GetCart)
	e.POST("/api/cart/create", cartHandler.CreateCart)
	e.POST("/api/cart/items", cartHandler.AddItemToCart)
	//e.PUT("/cart/items/:product_id", cartHandler.UpdateCartItem)
	e.DELETE("/api/cart/items/:product_id", cartHandler.RemoveCartItem)
	e.DELETE("/api/cart", cartHandler.ClearCart)

	e.POST("/api/orders", orderHandler.PlaceOrder)
	e.GET("/api/orders", orderHandler.GetOrders)
	e.GET("/api/orders/:order_id", orderHandler.GetOrder)

	e.POST("/api/coupons/validate", couponHandler.ValidateCoupon) // Added couponLib validation endpoint

	port := cfg.ServerPort
	if port == "" {
		port = "1323" // Default port if not set in config
	}

	e.Logger.Fatal(e.Start(":" + port))
}
