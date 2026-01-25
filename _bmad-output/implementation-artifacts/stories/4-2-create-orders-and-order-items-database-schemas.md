# Story 4.2: Create orders and order_items database schemas

**Epic:** Order Management API

**Description:**
As a developer, I want to create the database schemas for the `orders` and `order_items` tables in MySQL, so that we can persist and manage order data. This includes creating migration scripts.

**Acceptance Criteria:**
- SQL migration scripts are created to define the `orders` and `order_items` tables.
- The `orders` table schema includes columns for `ID`, `DeviceID`, `Status`, `TotalAmount`, `CreatedAt`, `UpdatedAt`.
- The `order_items` table schema includes columns for `ID`, `OrderID`, `ProductID`, `Quantity`, `UnitPrice`.
- Foreign key constraints are defined: `OrderID` in `order_items` references `ID` in `orders`, and `ProductID` in `order_items` references `ID` in `products`.
- Appropriate data types are used for MySQL (e.g., `VARCHAR`, `DATETIME`, `DECIMAL`, `INTEGER`).
- Primary keys are defined for both tables.
