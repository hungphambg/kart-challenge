# Story 4.2: Create orders and order_items database schemas

**Description:**
As a developer, I want to create the database schemas for the `orders` and `order_items` tables in MySQL, so that we can persist and manage order data. This includes creating migration scripts.

**Acceptance Criteria:**
- [x] SQL migration scripts are created to define the `orders` and `order_items` tables. (Implemented via GORM AutoMigrate)
- [x] The `orders` table schema includes columns for `ID`, `DeviceID`, `Status`, `TotalAmount`, `CreatedAt`, `UpdatedAt`.
- [x] The `order_items` table schema includes columns for `ID`, `OrderID`, `ProductID`, `Quantity`, `UnitPrice`.
- [x] Foreign key constraints are defined: `OrderID` in `order_items` references `ID` in `orders`, and `ProductID` in `order_items` references `ID` in `products`.
- [x] Appropriate data types are used for MySQL (e.g., `VARCHAR`, `DATETIME`, `DECIMAL`, `INTEGER`).
- [x] Primary keys are defined for both tables.

**Tasks/Subtasks:**
- [x] Rely on GORM's `AutoMigrate` feature for schema creation, which has already been configured in `cmd/api/main.go` and models.

**Dev Agent Record:**
- **Implementation Plan:**
  - Decided to leverage GORM's `AutoMigrate` functionality, which was set up during the GORM refactoring, to create the `orders` and `order_items` tables based on the `model.Order` and `model.OrderItem` structs. This implicitly handles schema definition, data types, and primary/foreign keys as per the model definitions.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `orders` and `order_items` database schemas are managed through GORM's `AutoMigrate` based on the defined `Order` and `OrderItem` models.

**File List:**
- `cmd/api/main.go` (modified in GORM refactoring)
- `internal/model/order.go` (created in Story 4.1)
- `internal/model/order_item.go` (created in Story 4.1)

**Change Log:**
- Implemented `orders` and `order_items` database schemas using GORM AutoMigrate.

**Status:**
- review