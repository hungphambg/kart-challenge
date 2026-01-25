# Epic 5: Coupon & Discount API

This epic covers the initial, simple implementation of coupon validation and application, before the more complex ingestion pipeline is introduced.

---

### Story 5.1: Define Coupon data model

**Description:**
As a developer, I want to define a clear data model for a "Coupon" in a Go struct, so that we can accurately represent various types of discount codes, their values, and validity.

**Acceptance Criteria:**
- A `model` package exists within the `internal` directory.
- `coupon.go` file is created within the `model` package.
- The `Coupon` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `Code` (string, unique, the actual coupon code like "HAPPYHOURS")
  - `Type` (string, e.g., "percentage", "buy_one_get_one", "fixed_amount")
  - `Value` (e.g., float or suitable monetary type, for percentage or fixed amount)
  - `Description` (string, human-readable description of the discount)
  - `ExpiresAt` (timestamp, optional)
  - `IsActive` (boolean)
- Struct includes appropriate JSON tags for serialization/deserialization.

---

### Story 5.2: Create coupons database schema

**Description:**
As a developer, I want to create the database schema for the `coupons` table in MySQL, so that we can persist and manage coupon codes and their properties. This includes creating migration scripts.

**Acceptance Criteria:**
- SQL migration scripts are created to define the `coupons` table.
- The `coupons` table schema includes columns for `ID`, `Code`, `Type`, `Value`, `Description`, `ExpiresAt`, `IsActive`.
- The `Code` column is unique.
- Appropriate data types are used for MySQL (e.g., `VARCHAR`, `TEXT`, `DECIMAL`, `DATETIME`, `BOOLEAN`).
- A primary key is defined for the table.

---

### Story 5.3: Implement POST /coupons/validate endpoint

**Description:**
As a client, I want to validate a coupon code against my current cart, so that I can see the potential discount before placing an order.

**Acceptance Criteria:**
- The API has a `POST /coupons/validate` endpoint.
- The request body includes the `coupon_code` and optionally the `cart_id` or `DeviceID` to calculate the discount.
- If the `coupon_code` is valid and active, the endpoint returns a JSON object with the coupon's details and the calculated discount amount (or effect, e.g., "free item") for the given cart.
- If the `coupon_code` is invalid, expired, or not applicable to the cart, the endpoint returns a `400 Bad Request` or `404 Not Found` with an appropriate error message.
- The endpoint interacts with the MySQL database to fetch coupon data.

---

### Story 5.4: Implement HAPPYHOURS coupon logic

**Description:**
As a coupon user, I want the `HAPPYHOURS` coupon to apply an 18% discount to my total cart value, so that I can save money on my order.

**Acceptance Criteria:**
- The `POST /coupons/validate` endpoint correctly identifies the `HAPPYHOURS` coupon.
- When `HAPPYHOURS` is applied to a cart, the calculated discount is 18% of the cart's subtotal.
- The discount is reflected in the response of the validate endpoint.
- The `HAPPYHOURS` coupon is valid and active.

---

### Story 5.5: Implement BUYGETONE coupon logic

**Description:**
As a coupon user, I want the `BUYGETONE` coupon to make one lowest-priced item in my cart free, so that I can get a free product.

**Acceptance Criteria:**
- The `POST /coupons/validate` endpoint correctly identifies the `BUYGETONE` coupon.
- When `BUYGETONE` is applied to a cart, one instance of the lowest-priced item in the cart is made free.
- If multiple items share the same lowest price, only one of them (the first encountered or chosen by a tie-breaking rule) is discounted.
- The discount is reflected in the response of the validate endpoint.
- The `BUYGETONE` coupon is valid and active.
