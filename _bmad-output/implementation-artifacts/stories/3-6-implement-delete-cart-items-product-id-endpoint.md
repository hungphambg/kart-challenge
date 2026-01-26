# Story 3.6: Implement DELETE /cart/items/{product_id} endpoint

**Description:**
As a client, I want to remove a specific item completely from my shopping cart, so that I can refine my selection before placing an order.

**Acceptance Criteria:**
- [x] The API has a `DELETE /cart/items/{product_id}` endpoint.
- [x] The endpoint requires a `DeviceID` in the HTTP header.
- [x] If the `product_id` is valid and the item exists in the user's cart, it is completely removed from the cart in the MySQL database.
- [x] The endpoint returns the updated cart as a JSON object with a `200 OK` status.
- [x] If the `product_id` is invalid or the item is not found in the cart, appropriate error responses are returned (e.g., `404 Not Found`).

**Tasks/Subtasks:**
- [x] Create `RemoveCartItem` function in `internal/database/database.go`.
- [x] Create `RemoveCartItem` handler in `internal/handler/cart_handler.go`.
- [x] Add `DELETE /cart/items/:product_id` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `RemoveCartItem` function to `DBClient` in `internal/database/database.go`.
  - Created `RemoveCartItem` handler in `internal/handler/cart_handler.go`.
  - In `cmd/api/main.go`, added a `DELETE /cart/items/:product_id` route.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `DELETE /cart/items/{product_id}` endpoint is implemented to remove an item from a cart.

**File List:**
- `internal/database/database.go`
- `internal/handler/cart_handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `DELETE /cart/items/{product_id}` endpoint.

**Status:**
- review
