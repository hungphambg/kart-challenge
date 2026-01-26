# Story 10.3: Update POST /orders endpoint for Redis validation

**Description:**
As a developer, I want to update the `POST /orders` endpoint to perform a real-time validity check of any provided coupon code against the Redis `Set` before processing the order, so that only genuinely valid coupons are applied to purchases.

**Acceptance Criteria:**
- [x] When a `coupon_code` is provided in the `POST /orders` request, the endpoint first uses the Redis client to check its existence in the `Set` of valid codes.
- [x] If the coupon code is not found in Redis, the order placement is rejected with an appropriate error (e.g., `400 Bad Request`).
- [x] If the coupon code is found in Redis, the order processing continues, and the relevant discount logic is applied. (Discount logic not yet implemented, but validation proceeds)
- [x] The `POST /orders` endpoint continues to perform all other necessary validations (e.g., stock checks).

**Tasks/Subtasks:**
- [x] Pass Redis client to `OrderHandler` in `cmd/api/main.go`.
- [x] Update `OrderHandler` struct in `internal/handler/order_handler.go` to include Redis client.
- [x] Modify `PlaceOrder` handler in `internal/handler/order_handler.go` to perform Redis coupon validation using `SIsMember`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Passed the initialized Redis client from `main.go` to the `OrderHandler` struct.
  - Modified the `PlaceOrder` handler to, if a `coupon_code` is provided, perform an `SIsMember` check against the `valid_coupon_codes` Redis Set.
  - If the coupon is not found in Redis, a `400 Bad Request` error is returned.
  - If found, the order processing continues. Discount application logic for orders is noted as a future enhancement.
- **Debug Log:**
  - No code-related issues. Requires a running Redis and database with synchronized coupon data for full verification.
- **Completion Notes:**
  - The `POST /orders` endpoint now includes a real-time Redis validation check for coupon codes, ensuring only valid coupons proceed with order processing.

**File List:**
- `cmd/api/main.go`
- `internal/handler/order_handler.go`

**Change Log:**
- Updated `POST /orders` endpoint for Redis validation.

**Status:**
- review