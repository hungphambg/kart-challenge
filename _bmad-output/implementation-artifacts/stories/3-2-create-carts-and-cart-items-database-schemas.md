# Story 3.2: Create carts and cart_items database schemas

**Epic:** Shopping Cart API

**Description:**
As a developer, I want to create the database schemas for the `carts` and `cart_items` tables in MySQL, so that we can persist and manage shopping cart data. This includes creating migration scripts.

**Acceptance Criteria:**
- SQL migration scripts are created to define the `carts` and `cart_items` tables.
- The `carts` table schema includes columns for `ID`, `DeviceID`, `CreatedAt`, `UpdatedAt`.
- The `cart_items` table schema includes columns for `ID`, `CartID`, `ProductID`, `Quantity`, `PriceAtAddition`.
- Foreign key constraints are defined: `CartID` in `cart_items` references `ID` in `carts`, and `ProductID` in `cart_items` references `ID` in `products`.
- Appropriate data types are used for MySQL (e.g., `VARCHAR`, `DATETIME`, `DECIMAL`, `INTEGER`).
- Primary keys are defined for both tables.
