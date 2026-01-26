# Story 3.2: Create carts and cart_items database schemas

**Description:**
As a developer, I want to create the database schemas for the `carts` and `cart_items` tables in MySQL, so that we can persist and manage shopping cart data. This includes creating migration scripts.

**Acceptance Criteria:**
- [x] SQL migration scripts are created to define the `carts` and `cart_items` tables.
- [x] The `carts` table schema includes columns for `ID`, `DeviceID`, `CreatedAt`, `UpdatedAt`.
- [x] The `cart_items` table schema includes columns for `ID`, `CartID`, `ProductID`, `Quantity`, `PriceAtAddition`.
- [x] Foreign key constraints are defined: `CartID` in `cart_items` references `ID` in `carts`, and `ProductID` in `cart_items` references `ID` in `products`.
- [x] Appropriate data types are used for MySQL (e.g., `VARCHAR`, `DATETIME`, `DECIMAL`, `INTEGER`).
- [x] Primary keys are defined for both tables.

**Tasks/Subtasks:**
- [x] Create a SQL migration script `002_create_carts_and_cart_items_tables.up.sql` with the `carts` and `cart_items` table schemas.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created a SQL migration script `migrations/002_create_carts_and_cart_items_tables.up.sql` that defines the `carts` and `cart_items` tables with appropriate columns, data types, and foreign key constraints.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The database migration script for the `carts` and `cart_items` tables is created.

**File List:**
- `migrations/002_create_carts_and_cart_items_tables.up.sql`

**Change Log:**
- Created the `carts` and `cart_items` tables migration script.

**Status:**
- review
