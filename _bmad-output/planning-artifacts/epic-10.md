# Epic 10: Live API Coupon Validation

This epic covers modifying the main API server to use the Redis set populated by the coupon pipeline for fast, real-time coupon validation.

---

### Story 10.1: Implement API Redis client

**Description:**
As a developer, I want to implement a Redis client within the main API server, so that it can connect to the Redis instance populated by the coupon ingestion pipeline and perform fast lookups for coupon validation.

**Acceptance Criteria:**
- The main API server's codebase includes a Redis client library for Go.
- The API server is configured to connect to the Redis service, using connection details from its configuration.
- A health check or similar mechanism confirms the API server can successfully connect to Redis on startup.
- The Redis client is made available to relevant service layers for coupon validation.

---

### Story 10.2: Modify POST /coupons/validate endpoint

**Description:**
As a developer, I want to modify the `POST /coupons/validate` endpoint to use the Redis `Set` for quickly checking the existence and validity of a coupon code, so that validation is highly performant.

**Acceptance Criteria:**
- The `POST /coupons/validate` endpoint utilizes the Redis client to perform an `SISMEMBER` operation against the `Set` of valid coupon codes.
- If Redis indicates the coupon code is not valid, the endpoint immediately returns an appropriate error (e.g., `404 Not Found` or `400 Bad Request`).
- The existing coupon-specific logic for `HAPPYHOURS` and `BUYGETONE` (if applicable for validation) is integrated with the Redis check.
- Performance of the validation endpoint is improved due to the O(1) Redis lookup.

---

### Story 10.3: Update POST /orders endpoint for Redis validation

**Description:**
As a developer, I want to update the `POST /orders` endpoint to perform a real-time validity check of any provided coupon code against the Redis `Set` before processing the order, so that only genuinely valid coupons are applied to purchases.

**Acceptance Criteria:**
- When a `coupon_code` is provided in the `POST /orders` request, the endpoint first uses the Redis client to check its existence in the `Set` of valid codes.
- If the coupon code is not found in Redis, the order placement is rejected with an appropriate error (e.g., `400 Bad Request`).
- If the coupon code is found in Redis, the order processing continues, and the relevant discount logic is applied.
- The `POST /orders` endpoint continues to perform all other necessary validations (e.g., stock checks).
