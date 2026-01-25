# Story 4.1: Define Order and Order Item data models

**Epic:** Order Management API

**Description:**
As a developer, I want to define clear data models for an "Order" and "Order Item" in Go structs, so that we can accurately represent a placed order, its details, and link it to a device ID.

**Acceptance Criteria:**
- A `model` package exists within the `internal` directory.
- `order.go` and `order_item.go` files are created within the `model` package.
- The `Order` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `DeviceID` (string, foreign key to device identification)
  - `Status` (string, e.g., "pending", "completed", "cancelled")
  - `TotalAmount` (e.g., float or suitable monetary type)
  - `CreatedAt`, `UpdatedAt` (timestamps)
- The `OrderItem` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `OrderID` (foreign key to `Order.ID`)
  - `ProductID` (foreign key to `Product.ID`)
  - `Quantity` (integer)
  - `UnitPrice` (e.g., float or suitable monetary type, price at time of order)
- Structs include appropriate JSON tags for serialization/deserialization.
