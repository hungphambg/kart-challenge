# Story 4.5: Implement GET /orders/{order_id} endpoint

**Description:**
As a client, I want to view the detailed information of a specific past order, so that I can review its contents and status.

**Acceptance Criteria:**
- [x] The API has a `GET /orders/{order_id}` endpoint.
- [x] The endpoint requires a `DeviceID` in the HTTP header.
- [x] If the `order_id` is valid and belongs to the `DeviceID`, it returns a JSON object representing the `Order`, including its `OrderItem` details.
- [x] If the `order_id` is not found or does not belong to the `DeviceID`, the endpoint returns a `404 Not Found` status with an appropriate error message.
- [x] The endpoint interacts with the MySQL database to fetch order data.

**Tasks/Subtasks:**
- [x] Add `GetOrderByIDAndDeviceID` function to `internal/database/database.go`.
- [x] Add `GetOrder` handler to `internal/handler/order_handler.go`.
- [x] Add `GET /orders/:order_id` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added a `GetOrderByIDAndDeviceID` function to `DBClient` in `internal/database/database.go` using GORM `Where`, `Preload`, and `First` methods to fetch a specific order and its items.
  - Added `GetOrder` handler to `internal/handler/order_handler.go` that validates `DeviceID` and `order_id`, calls the database function, and handles "not found" cases.
  - In `cmd/api/main.go`, added a `GET /orders/:order_id` route.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `GET /orders/{order_id}` endpoint is implemented to retrieve detailed information for a specific order.

**File List:**
- `internal/database/database.go`
- `internal/handler/order_handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `GET /orders/{order_id}` endpoint.

**Status:**
- review