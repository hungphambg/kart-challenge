# Epic 3: Shopping Cart API

This epic covers all functionality related to creating, managing, and viewing a user's shopping cart.

---

### Story 3.1: Define Cart and Cart Item data models

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

---

### Story 3.2: Create carts and cart_items database schemas

**Description:**
As a developer, I want to create the database schemas for the `carts` and `cart_items` tables in MySQL, so that we can persist and manage shopping cart data. This includes creating migration scripts.

**Acceptance Criteria:**
- SQL migration scripts are created to define the `carts` and `cart_items` tables.
- The `carts` table schema includes columns for `ID`, `DeviceID`, `CreatedAt`, `UpdatedAt`.
- The `cart_items` table schema includes columns for `ID`, `CartID`, `ProductID`, `Quantity`, `PriceAtAddition`.
- Foreign key constraints are defined: `CartID` in `cart_items` references `ID` in `carts`, and `ProductID` in `cart_items` references `ID` in `products`.
- Appropriate data types are used for MySQL (e.g., `VARCHAR`, `DATETIME`, `DECIMAL`, `INTEGER`).
- Primary keys are defined for both tables.

---

### Story 3.3: Implement GET /cart endpoint

**Description:**
As a client, I want to retrieve my current shopping cart. If I don't have one, I want the system to create a new empty cart for me, so that I can start adding items.

**Acceptance Criteria:**
- The API has a `GET /cart` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- If a cart exists for the given `DeviceID`, it is returned as a JSON object, including its `CartItems`.
- If no cart exists for the `DeviceID`, a new empty cart is created in the MySQL database and returned.
- The response includes the cart's unique identifier and its current list of items.
- The endpoint handles cases where the `DeviceID` is missing, returning a `400 Bad Request` with an appropriate error message.

---

### Story 3.4: Implement POST /cart/items endpoint

**Description:**
As a client, I want to add a specific product with a chosen quantity to my shopping cart, so that I can prepare my order.

**Acceptance Criteria:**
- The API has a `POST /cart/items` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- The request body includes `product_id` and `quantity`.
- If the product exists and is in stock, the item is added to the user's cart (or its quantity is increased if already present).
- The `PriceAtAddition` for the `CartItem` is recorded based on the current product price.
- The endpoint returns the updated cart as a JSON object with a `200 OK` or `201 Created` status.
- If the `product_id` is invalid or out of stock, appropriate error responses are returned (e.g., `404 Not Found`, `409 Conflict` or `400 Bad Request`).
- The endpoint interacts with the MySQL database to update cart and fetch product data.

---

### Story 3.5: Implement PUT /cart/items/{product_id} endpoint

**Description:**
As a client, I want to be able to change the quantity of a specific item in my shopping cart, so that I can adjust my order before checkout.

**Acceptance Criteria:**
- The API has a `PUT /cart/items/{product_id}` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- The request body includes the `quantity` (the new total quantity for the item).
- If the `product_id` is valid, the item is in the cart, and the new `quantity` is valid (e.g., non-negative, not exceeding stock), the item's quantity is updated in the MySQL database.
- If the `quantity` is `0`, the item is effectively removed from the cart.
- The endpoint returns the updated cart as a JSON object with a `200 OK` status.
- Appropriate error responses are returned for invalid `product_id`, item not in cart, or invalid `quantity` (e.g., `404 Not Found`, `400 Bad Request`).

---

### Story 3.6: Implement DELETE /cart/items/{product_id} endpoint

**Description:**
As a client, I want to remove a specific item completely from my shopping cart, so that I can refine my selection before placing an order.

**Acceptance Criteria:**
- The API has a `DELETE /cart/items/{product_id}` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- If the `product_id` is valid and the item exists in the user's cart, it is completely removed from the cart in the MySQL database.
- The endpoint returns the updated cart as a JSON object with a `200 OK` status.
- If the `product_id` is invalid or the item is not found in the cart, appropriate error responses are returned (e.g., `404 Not Found`).

---

### Story 3.7: Implement DELETE /cart endpoint

**Description:**
As a client, I want to be able to empty my entire shopping cart, so that I can start a new order or completely discard my current selection.

**Acceptance Criteria:**
- The API has a `DELETE /cart` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- The endpoint removes all items associated with the user's cart in the MySQL database.
- The endpoint returns a `200 OK` status with an empty cart representation or a success message.
- If no cart exists for the `DeviceID`, it returns a `200 OK` (idempotent operation).
