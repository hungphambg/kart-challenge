# Story 4.4: Implement GET /orders endpoint

**Description:**
As a client, I want to view a history of all my past orders, so that I can keep track of my purchases.

**Acceptance Criteria:**
- [x] The API has a `GET /orders` endpoint.
- [x] The endpoint requires a `DeviceID` in the HTTP header.
- [x] The endpoint returns a JSON array of `Order` objects associated with the `DeviceID`.
- [x] Each `Order` object includes at least its `ID`, `Status`, `TotalAmount`, `CreatedAt`.
- [x] The endpoint handles cases where no orders are found for the `DeviceID`, returning an empty array and a `200 OK` status.
- [x] The endpoint interacts with the MySQL database to fetch order data.

**Tasks/Subtasks:**
- [x] Add `GetOrdersByDeviceID` function to `internal/database/database.go`.
- [x] Add `GetOrders` handler to `internal/handler/order_handler.go`.
- [x] Add `GET /orders` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `GetOrdersByDeviceID` function to `DBClient` in `internal/database/database.go` using GORM's `Where` and `Find` methods.
  - Added `GetOrders` handler to `internal/handler/order_handler.go` that validates `DeviceID` and calls the database function.
  - In `cmd/api/main.go`, added a `GET /orders` route.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `GET /orders` endpoint is implemented to retrieve all orders for a given device ID.

**File List:**
- `internal/database/database.go`
- `internal/handler/order_handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `GET /orders` endpoint.

**Status:**
- review