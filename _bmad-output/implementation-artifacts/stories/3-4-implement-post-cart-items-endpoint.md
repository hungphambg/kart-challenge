# Story 3.4: Implement POST /cart/items endpoint

**Description:**
As a client, I want to add a specific product with a chosen quantity to my shopping cart, so that I can prepare my order.

**Acceptance Criteria:**
- [x] The API has a `POST /cart/items` endpoint.
- [x] The endpoint requires a `DeviceID` in the HTTP header.
- [x] The request body includes `product_id` and `quantity`.
- [x] If the product exists and is in stock, the item is added to the user's cart (or its quantity is increased if already present).
- [x] The `PriceAtAddition` for the `CartItem` is recorded based on the current product price.
- [x] The endpoint returns the updated cart as a JSON object with a `200 OK` or `201 Created` status.
- [x] If the `product_id` is invalid or out of stock, appropriate error responses are returned (e.g., `404 Not Found`, `409 Conflict` or `400 Bad Request`).
- [x] The endpoint interacts with the MySQL database to update cart and fetch product data.

**Tasks/Subtasks:**
- [x] Create `AddItemToCart` function in `internal/database/database.go`.
- [x] Create `AddItemToCart` handler in `internal/handler/cart_handler.go`.
- [x] Add `POST /cart/items` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `AddItemToCart` function to `DBClient` in `internal/database/database.go` to handle adding a new item or updating the quantity of an existing item.
  - Created `AddItemToCart` handler in `internal/handler/cart_handler.go` that validates the request, checks product stock, and calls the database function.
  - In `cmd/api/main.go`, added a `POST /cart/items` route.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `POST /cart/items` endpoint is implemented to add items to a cart.

**File List:**
- `internal/database/database.go`
- `internal/handler/cart_handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `POST /cart/items` endpoint.

**Status:**
- review
