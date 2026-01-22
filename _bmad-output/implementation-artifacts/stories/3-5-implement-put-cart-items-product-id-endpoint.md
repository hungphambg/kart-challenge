# Story 3.5: Implement PUT /cart/items/{product_id} endpoint

**Epic:** Shopping Cart API

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
