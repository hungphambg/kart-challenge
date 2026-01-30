package database

import (
	"errors"
	"fmt"
	"time"

	"kart-challenge/internal/meta"
	"kart-challenge/internal/model"

	"github.com/shopspring/decimal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBClient represents the database client
type DBClient struct {
	DB *gorm.DB
}

// NewDBClient initializes a new database client
func NewDBClient(databaseURL string) (*DBClient, error) {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting underlying sql.DB: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return &DBClient{DB: db}, nil
}

// Close closes the database connection
func (client *DBClient) Close() error {
	sqlDB, err := client.DB.DB()
	if err != nil {
		return fmt.Errorf("error getting underlying sql.DB: %w", err)
	}
	return sqlDB.Close()
}

// Ping checks the database connection
func (client *DBClient) Ping() error {
	sqlDB, err := client.DB.DB()
	if err != nil {
		return fmt.Errorf("error getting underlying sql.DB: %w", err)
	}
	return sqlDB.Ping()
}

// GetProducts retrieves all products from the database
func (client *DBClient) GetProducts() ([]model.Product, error) {
	var products []model.Product
	err := client.DB.Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("error querying products: %w", err)
	}
	return products, nil
}

// GetProductByID retrieves a single product from the database by its ID
func (client *DBClient) GetProductByID(id uint64) (*model.Product, error) {
	var product model.Product
	err := client.DB.First(&product, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No product found
		}
		return nil, fmt.Errorf("error querying product: %w", err)
	}
	return &product, nil
}

// GetCartByDeviceID retrieves a cart and its items by device ID
func (client *DBClient) GetCartByDeviceID(deviceID string) (*model.Cart, error) {
	var cart model.Cart
	err := client.DB.
		Where("device_id = ?", deviceID).
		Preload("Items").
		First(&cart).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No cart found
		}
		return nil, fmt.Errorf("error querying cart: %w", err)
	}
	return &cart, nil
}

// GetCartByID get cart by id
func (client *DBClient) GetCartByID(ID uint64) (*model.Cart, error) {
	var cart model.Cart
	err := client.DB.
		Where(&model.Cart{ID: ID}).
		Preload("Items").
		First(&cart).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No cart found
		}
		return nil, fmt.Errorf("error querying cart: %w", err)
	}
	return &cart, nil
}

// CreateCart creates a new cart for a given device ID
func (client *DBClient) CreateCart(deviceID string) (*model.Cart, error) {
	cart := model.Cart{
		DeviceID:  deviceID,
		Status:    meta.CartStatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Items:     []model.CartItem{},
	}
	err := client.DB.Create(&cart).Error
	if err != nil {
		return nil, fmt.Errorf("error creating cart: %w", err)
	}
	return &cart, nil
}

// AddItemToCart adds an item to a cart or updates its quantity if it already exists
func (client *DBClient) AddItemToCart(cartID uint64, productID uint64, quantity uint32, price decimal.Decimal) error {
	var item model.CartItem
	err := client.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Item does not exist, insert new item
			item = model.CartItem{
				CartID:          cartID,
				ProductID:       productID,
				Quantity:        quantity,
				PriceAtAddition: price,
			}
			err = client.DB.Create(&item).Error
			if err != nil {
				return fmt.Errorf("error inserting new cart item: %w", err)
			}
			return nil
		}
		return fmt.Errorf("error querying for existing cart item: %w", err)
	}

	// Item exists, update quantity
	item.Quantity += quantity
	err = client.DB.Save(&item).Error
	if err != nil {
		return fmt.Errorf("error updating cart item quantity: %w", err)
	}
	return nil
}

// UpdateCartItemQuantity updates the quantity of a cart item. If the quantity is 0, it removes the item.
func (client *DBClient) UpdateCartItemQuantity(cartID uint64, productID uint64, quantity uint32) error {
	if quantity == 0 {
		err := client.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&model.CartItem{}).Error
		if err != nil {
			return fmt.Errorf("error deleting cart item: %w", err)
		}
		return nil
	}

	var item model.CartItem
	err := client.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cart item not found")
		}
		return fmt.Errorf("error querying cart item: %w", err)
	}

	item.Quantity = quantity
	err = client.DB.Save(&item).Error
	if err != nil {
		return fmt.Errorf("error updating cart item quantity: %w", err)
	}
	return nil
}

// RemoveCartItem removes an item from a cart.
func (client *DBClient) RemoveCartItem(cartID uint64, productID int64) error {
	err := client.DB.Where("cart_id = ? AND product_id = ?", cartID, productID).Delete(&model.CartItem{}).Error
	if err != nil {
		return fmt.Errorf("error deleting cart item: %w", err)
	}
	return nil
}

// ClearCart removes all items from a cart.
func (client *DBClient) ClearCart(cartID uint64) error {
	err := client.DB.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error
	if err != nil {
		return fmt.Errorf("error clearing cart: %w", err)
	}
	return nil
}

//// GetValidCouponByCode retrieves a coupon by its code and checks its validity
//func (client *DBClient) GetValidCouponByCode(code string) (*model.Coupon, error) {
//	var coupon model.Coupon
//	now := time.Now()
//	err := client.DB.Where("code = ? AND is_active = ? AND (expires_at IS NULL OR expires_at > ?)", code, true, now).First(&coupon).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return nil, nil // No valid coupon found
//		}
//		return nil, fmt.Errorf("error querying coupon: %w", err)
//	}
//	return &coupon, nil
//}

// GetValidCouponByCode retrieves a coupon by its code and checks its validity
func (client *DBClient) GetCouponByCode(code string) (*model.Coupon, error) {
	var coupon = model.Coupon{
		Code: code,
	}
	err := client.DB.
		Model(&coupon).
		Where(&coupon).
		First(&coupon).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("error querying coupon: %w", err)
	}
	return &coupon, nil
}

// CreateOrder creates a new order from a cart, updates product stock, and clears the cart.
func (client *DBClient) CreateOrder(cartID uint64, deviceID string) (*model.Order, error) {
	tx := client.DB.Begin()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}
	defer tx.Rollback() // Rollback in case of panic or explicit return

	// 1. Get cart items
	var cartItems []model.CartItem
	err := tx.Where("cart_id = ?", cartID).Find(&cartItems).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get cart items: %w", err)
	}

	if len(cartItems) == 0 {
		return nil, fmt.Errorf("cart is empty")
	}

	var totalAmount decimal.Decimal
	var orderItems []model.OrderItem

	for _, cartItem := range cartItems {
		// 2. Check product stock
		var product model.Product
		err = tx.First(&product, cartItem.ProductID).Error
		if err != nil {
			return nil, fmt.Errorf("product %d not found: %w", cartItem.ProductID, err)
		}

		if product.Stock < cartItem.Quantity {
			return nil, fmt.Errorf("product %s (ID: %d) is out of stock. Available: %d, Requested: %d", product.Name, product.ID, product.Stock, cartItem.Quantity)
		}

		// 3. Prepare order item
		orderItem := model.OrderItem{
			ProductID: cartItem.ProductID,
			Quantity:  cartItem.Quantity,
			UnitPrice: product.Price, // Price at time of order
		}
		orderItems = append(orderItems, orderItem)
		totalAmount = totalAmount.Add(product.Price.Mul(decimal.NewFromUint64(uint64(cartItem.Quantity))))
	}

	// 4. Create new order
	now := time.Now()
	order := model.Order{
		DeviceID:    deviceID,
		Status:      meta.OrderStatusPending,
		TotalAmount: totalAmount,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err = tx.Create(&order).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	order.Items = orderItems
	// GORM should automatically save orderItems due to hasMany relationship, but let's be explicit
	for i := range order.Items {
		order.Items[i].OrderID = order.ID
		err = tx.Create(&order.Items[i]).Error
		if err != nil {
			return nil, fmt.Errorf("failed to create order item: %w", err)
		}
	}

	for _, cartItem := range cartItems {
		// 5. Update product stock
		err = tx.Model(&model.Product{}).Where("id = ?", cartItem.ProductID).
			Update("stock", gorm.Expr("stock - ?", cartItem.Quantity)).Error
		if err != nil {
			return nil, fmt.Errorf("failed to update stock for product %d: %w", cartItem.ProductID, err)
		}
	}

	err = tx.Model(&model.Cart{}).
		Where(&model.Cart{
			ID: cartID,
		}).
		Update(model.CartColNameStatus, meta.CartStatusDone).
		Error
	if err != nil {
		return nil, fmt.Errorf("failed to save cart item: %w", err)
	}

	//// 6. Clear the cart
	//err = tx.Where("cart_id = ?", cartID).Delete(&model.CartItem{}).Error
	//if err != nil {
	//	return nil, fmt.Errorf("failed to clear cart items: %w", err)
	//}
	//
	//err = tx.Where("id = ?", cartID).Delete(&model.Cart{}).Error
	//if err != nil {
	//	return nil, fmt.Errorf("failed to delete cart: %w", err)
	//}

	tx.Commit()
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", tx.Error)
	}

	return &order, nil

}

// GetOrdersByDeviceID retrieves all orders for a given device ID

func (client *DBClient) GetOrdersByDeviceID(deviceID string) ([]model.Order, error) {

	var orders []model.Order

	err := client.DB.Where("device_id = ?", deviceID).Find(&orders).Error

	if err != nil {

		return nil, fmt.Errorf("failed to get orders for device ID %s: %w", deviceID, err)

	}

	return orders, nil

}

// GetOrderByIDAndDeviceID retrieves a single order by its ID and DeviceID

func (client *DBClient) GetOrderByIDAndDeviceID(id int64, deviceID string) (*model.Order, error) {

	var order model.Order

	err := client.DB.Where("id = ? AND device_id = ?", id, deviceID).Preload("Items").First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No order found
		}
		return nil, fmt.Errorf("error querying order: %w", err)
	}

	return &order, nil

}
