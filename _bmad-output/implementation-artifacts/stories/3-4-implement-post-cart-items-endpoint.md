# Story 3.4: Implement POST /cart/items endpoint

**Epic:** Shopping Cart API

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
