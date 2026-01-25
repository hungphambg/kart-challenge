# Story 4.3: Implement POST /orders endpoint

**Epic:** Order Management API

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
