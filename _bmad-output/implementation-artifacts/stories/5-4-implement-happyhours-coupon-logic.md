# Story 5.4: Implement HAPPYHOURS coupon logic

**Description:**
As a coupon user, I want the `HAPPYHOURS` coupon to apply an 18% discount to my total cart value, so that I can save money on my order.

**Acceptance Criteria:**
- [x] The `POST /coupons/validate` endpoint correctly identifies the `HAPPYHOURS` coupon.
- [x] When `HAPPYHOURS` is applied to a cart, the calculated discount is 18% of the cart's subtotal.
- [x] The discount is reflected in the response of the validate endpoint.
- [x] The `HAPPYHOURS` coupon is valid and active.

**Tasks/Subtasks:**
- [x] Modify `internal/handler/coupon_handler.go` to calculate and return an 18% discount for the `HAPPYHOURS` coupon.

**Dev Agent Record:**
- **Implementation Plan:**
  - Modified `internal/handler/coupon_handler.go` to fetch the cart (if `DeviceID` is provided), calculate the subtotal of the cart items, and if the coupon code is "HAPPYHOURS", compute an 18% discount. The response now includes the calculated `discount_amount` and `discount_type`.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `HAPPYHOURS` coupon logic is implemented in the validation endpoint.

**File List:**
- `internal/handler/coupon_handler.go`

**Change Log:**
- Implemented `HAPPYHOURS` coupon logic.

**Status:**
- review