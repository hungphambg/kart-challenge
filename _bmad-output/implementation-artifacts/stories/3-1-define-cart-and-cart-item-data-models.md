# Story 3.1: Define Cart and Cart Item data models

**Epic:** Shopping Cart API

**Description:**
As a developer, I want to define clear data models for a "Cart" and "Cart Item" in Go structs, so that we can accurately represent a user's shopping cart and its contents, linked to their device ID and products.

**Acceptance Criteria:**
- A `model` package exists within the `internal` directory.
- `cart.go` and `cart_item.go` files are created within the `model` package.
- The `Cart` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `DeviceID` (string, foreign key to device identification)
  - `CreatedAt`, `UpdatedAt` (timestamps)
- The `CartItem` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `CartID` (foreign key to `Cart.ID`)
  - `ProductID` (foreign key to `Product.ID`)
  - `Quantity` (integer)
  - `PriceAtAddition` (e.g., float or suitable monetary type, to capture price at time of addition)
- Structs include appropriate JSON tags for serialization/deserialization.
