# Story 3.1: Define Cart and Cart Item data models

**Description:**
As a developer, I want to define clear data models for a "Cart" and "Cart Item" in Go structs, so that we can accurately represent a user's shopping cart and its contents, linked to their device ID and products.

**Acceptance Criteria:**
- [x] A `model` package exists within the `internal` directory.
- [x] `cart.go` and `cart_item.go` files are created within the `model` package.
- [x] The `Cart` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `DeviceID` (string, foreign key to device identification)
  - `CreatedAt`, `UpdatedAt` (timestamps)
- [x] The `CartItem` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `CartID` (foreign key to `Cart.ID`)
  - `ProductID` (foreign key to `Product.ID`)
  - `Quantity` (integer)
  - `PriceAtAddition` (e.g., float or suitable monetary type, to capture price at time of addition)
- [x] Structs include appropriate JSON tags for serialization/deserialization.

**Tasks/Subtasks:**
- [x] Create `internal/model/cart.go` with the `Cart` struct.
- [x] Create `internal/model/cart_item.go` with the `CartItem` struct.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `internal/model/cart.go` and defined the `Cart` struct.
  - Created `internal/model/cart_item.go` and defined the `CartItem` struct.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `Cart` and `CartItem` data models are defined.

**File List:**
- `internal/model/cart.go`
- `internal/model/cart_item.go`

**Change Log:**
- Defined the `Cart` and `CartItem` data models.

**Status:**
- review
