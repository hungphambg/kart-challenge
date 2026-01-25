# Story 10.3: Update POST /orders endpoint for Redis validation

**Epic:** Live API Coupon Validation

**Description:**
As a developer, I want to update the `POST /orders` endpoint to perform a real-time validity check of any provided coupon code against the Redis `Set` before processing the order, so that only genuinely valid coupons are applied to purchases.

**Acceptance Criteria:**
- When a `coupon_code` is provided in the `POST /orders` request, the endpoint first uses the Redis client to check its existence in the `Set` of valid codes.
- If the coupon code is not found in Redis, the order placement is rejected with an appropriate error (e.g., `400 Bad Request` or `404 Not Found`).
- If the coupon code is found in Redis, the order processing continues, and the relevant discount logic (from Epic 5) is applied.
- The `POST /orders` endpoint continues to perform all other necessary validations (e.g., stock checks) as previously defined.
