# Epic 4: Order Management API

This epic covers all functionality related to placing and viewing orders.

---

### Story 4.1: Define Order and Order Item data models

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

---

### Story 4.2: Create orders and order_items database schemas

**Description:**
As a developer, I want to create the database schemas for the `orders` and `order_items` tables in MySQL, so that we can persist and manage order data. This includes creating migration scripts.

**Acceptance Criteria:**
- SQL migration scripts are created to define the `orders` and `order_items` tables.
- The `orders` table schema includes columns for `ID`, `DeviceID`, `Status`, `TotalAmount`, `CreatedAt`, `UpdatedAt`.
- The `order_items` table schema includes columns for `ID`, `OrderID`, `ProductID`, `Quantity`, `UnitPrice`.
- Foreign key constraints are defined: `OrderID` in `order_items` references `ID` in `orders`, and `ProductID` in `order_items` references `ID` in `products`.
- Appropriate data types are used for MySQL (e.g., `VARCHAR`, `DATETIME`, `DECIMAL`, `INTEGER`).
- Primary keys are defined for both tables.

---

### Story 4.3: Implement POST /orders endpoint

**Description:**
As a client, I want to place an order using the items in my current shopping cart, so that I can finalize my purchase. The system should validate stock, apply any applicable coupons, and then clear my cart.

**Acceptance Criteria:**
- The API has a `POST /orders` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- The system checks the availability (stock) of all items in the user's cart.
- If any item is out of stock, the entire order is rejected, and an appropriate error is returned (e.g., `409 Conflict` with details on out-of-stock items).
- If a coupon code is provided in the request body, it is validated and applied to the cart total.
- A new `Order` record is created in the MySQL database, along with corresponding `OrderItem` records.
- The user's shopping cart is cleared upon successful order placement.
- The endpoint returns the created `Order` object with a `201 Created` status.
- The `system_context.md` specifies: "The shopping cart associated with a device will be cleared automatically and immediately upon a successful order placement."

---

### Story 4.4: Implement GET /orders endpoint

**Description:**
As a client, I want to view a history of all my past orders, so that I can keep track of my purchases.

**Acceptance Criteria:**
- The API has a `GET /orders` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- The endpoint returns a JSON array of `Order` objects associated with the `DeviceID`.
- Each `Order` object includes at least its `ID`, `Status`, `TotalAmount`, `CreatedAt`.
- The endpoint handles cases where no orders are found for the `DeviceID`, returning an empty array and a `200 OK` status.
- The endpoint interacts with the MySQL database to fetch order data.

---

### Story 4.5: Implement GET /orders/{order_id} endpoint

**Description:**
As a client, I want to view the detailed information of a specific past order, so that I can review its contents and status.

**Acceptance Criteria:**
- The API has a `GET /orders/{order_id}` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- If the `order_id` is valid and belongs to the `DeviceID`, it returns a JSON object representing the `Order`, including its `OrderItem` details.
- If the `order_id` is not found or does not belong to the `DeviceID`, the endpoint returns a `404 Not Found` status with an appropriate error message.
- The endpoint interacts with the MySQL database to fetch order data.
