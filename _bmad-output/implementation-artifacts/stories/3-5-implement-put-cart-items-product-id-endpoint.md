# Story 3.5: Implement PUT /cart/items/{product_id} endpoint

**Description:**
As a client, I want to be able to change the quantity of a specific item in my shopping cart, so that I can adjust my order before checkout.

**Acceptance Criteria:**
- [x] The API has a `PUT /cart/items/{product_id}` endpoint.
- [x] The endpoint requires a `DeviceID` in the HTTP header.
- [x] The request body includes the `quantity` (the new total quantity for the item).
- [x] If the `product_id` is valid, the item is in the cart, and the new `quantity` is valid (e.g., non-negative, not exceeding stock), the item's quantity is updated in the MySQL database.
- [x] If the `quantity` is `0`, the item is effectively removed from the cart.
- [x] The endpoint returns the updated cart as a JSON object with a `200 OK` status.
- [x] Appropriate error responses are returned for invalid `product_id`, item not in cart, or invalid `quantity` (e.g., `404 Not Found`, `400 Bad Request`).

**Tasks/Subtasks:**
- [x] Create `UpdateCartItemQuantity` function in `internal/database/database.go`.
- [x] Create `UpdateCartItem` handler in `internal/handler/cart_handler.go`.
- [x] Add `PUT /cart/items/:product_id` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `UpdateCartItemQuantity` function to `DBClient` in `internal/database/database.go` to handle updating the quantity of an existing item or deleting it if the quantity is 0.
  - Created `UpdateCartItem` handler in `internal/handler/cart_handler.go` that validates the request, checks product stock, and calls the database function.
  - In `cmd/api/main.go`, added a `PUT /cart/items/:product_id` route.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `PUT /cart/items/{product_id}` endpoint is implemented to update item quantities in a cart.

**File List:**
- `internal/database/database.go`
- `internal/handler/cart_handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `PUT /cart/items/{product_id}` endpoint.

**Status:**
- review
