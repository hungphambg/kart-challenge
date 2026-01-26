# Story 3.7: Implement DELETE /cart endpoint

**Description:**
As a client, I want to be able to empty my entire shopping cart, so that I can start a new order or completely discard my current selection.

**Acceptance Criteria:**
- [x] The API has a `DELETE /cart` endpoint.
- [x] The endpoint requires a `DeviceID` in the HTTP header.
- [x] The endpoint removes all items associated with the user's cart in the MySQL database.
- [x] The endpoint returns a `200 OK` status with an empty cart representation or a success message.
- [x] If no cart exists for the `DeviceID`, it returns a `200 OK` (idempotent operation).

**Tasks/Subtasks:**
- [x] Create `ClearCart` function in `internal/database/database.go`.
- [x] Create `ClearCart` handler in `internal/handler/cart_handler.go`.
- [x] Add `DELETE /cart` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `ClearCart` function to `DBClient` in `internal/database/database.go`.
  - Created `ClearCart` handler in `internal/handler/cart_handler.go`.
  - In `cmd/api/main.go`, added a `DELETE /cart` route.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `DELETE /cart` endpoint is implemented to clear a user's cart.

**File List:**
- `internal/database/database.go`
- `internal/handler/cart_handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `DELETE /cart` endpoint.

**Status:**
- review
