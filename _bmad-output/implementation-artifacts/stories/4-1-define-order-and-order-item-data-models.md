# Story 4.1: Define Order and Order Item data models

**Description:**
As a developer, I want to define clear data models for an "Order" and "Order Item" in Go structs, so that we can accurately represent a placed order, its details, and link it to a device ID.

**Acceptance Criteria:**
- [x] A `model` package exists within the `internal` directory.
- [x] `order.go` and `order_item.go` files are created within the `model` package.
- [x] The `Order` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `DeviceID` (string, foreign key to device identification)
  - `Status` (string, e.g., "pending", "completed", "cancelled")
  - `TotalAmount` (e.g., float or suitable monetary type)
  - `CreatedAt`, `UpdatedAt` (timestamps)
- [x] The `OrderItem` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `OrderID` (foreign key to `Order.ID`)
  - `ProductID` (foreign key to `Product.ID`)
  - `Quantity` (integer)
  - `UnitPrice` (e.g., float or suitable monetary type, price at time of order)
- [x] Structs include appropriate JSON tags for serialization/deserialization.

**Tasks/Subtasks:**
- [x] Create `internal/model/order.go` with the `Order` struct.
- [x] Create `internal/model/order_item.go` with the `OrderItem` struct.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `internal/model/order.go` and defined the `Order` struct with GORM primary key and foreign key tags.
  - Created `internal/model/order_item.go` and defined the `OrderItem` struct with GORM primary key tag.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `Order` and `OrderItem` data models are defined and are GORM compatible.

**File List:**
- `internal/model/order.go`
- `internal/model/order_item.go`

**Change Log:**
- Defined the `Order` and `OrderItem` data models.

**Status:**
- review