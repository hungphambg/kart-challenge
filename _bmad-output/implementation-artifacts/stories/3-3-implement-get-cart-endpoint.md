# Story 3.3: Implement GET /cart endpoint

**Description:**
As a client, I want to retrieve my current shopping cart. If I don't have one, I want the system to create a new empty cart for me, so that I can start adding items.

**Acceptance Criteria:**
- [x] The API has a `GET /cart` endpoint.
- [x] The endpoint requires a `DeviceID` in the HTTP header.
- [x] If a cart exists for the given `DeviceID`, it is returned as a JSON object, including its `CartItems`.
- [x] If no cart exists for the `DeviceID`, a new empty cart is created in the MySQL database and returned.
- [x] The response includes the cart's unique identifier and its current list of items.
- [x] The endpoint handles cases where the `DeviceID` is missing, returning a `400 Bad Request` with an appropriate error message.

**Tasks/Subtasks:**
- [x] Create `GetCartByDeviceID` and `CreateCart` functions in `internal/database/database.go`.
- [x] Create `CartHandler` with `GetCart` method in `internal/handler/cart_handler.go`.
- [x] Add `GET /cart` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `GetCartByDeviceID` and `CreateCart` functions to the `DBClient` in `internal/database/database.go`.
  - Created a `CartHandler` in `internal/handler/cart_handler.go` with a `GetCart` method that checks for `DeviceID` header, gets or creates a cart, and returns it.
  - In `cmd/api/main.go`, instantiated the `CartHandler` and added a `GET /cart` route.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `GET /cart` endpoint is implemented to retrieve or create a cart for a given device ID.

**File List:**
- `internal/database/database.go`
- `internal/handler/cart_handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `GET /cart` endpoint.

**Status:**
- review
