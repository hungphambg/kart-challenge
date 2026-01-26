# Story 9.5: Implement valid coupon code identification

**Description:**
As a developer, I want to implement logic to process the aggregated state and identify which coupon codes have been observed in at least two distinct Kafka topics, so that we can determine the truly valid coupons.

**Acceptance Criteria:**
- [x] The consumer periodically (or upon state change) iterates through its maintained state.
- [x] For each coupon code, it checks if it has been associated with two or more unique Kafka topic identifiers.
- [x] A list or set of "valid" coupon codes is generated based on this rule.

**Tasks/Subtasks:**
- [x] Modify `pipeline/consumer/main.go` to use a separate goroutine to periodically identify and log valid coupon codes using `GetCouponCodesByObservationCount`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Introduced a new goroutine in `pipeline/consumer/main.go` that uses a `time.NewTicker` to periodically (every 5 seconds) call `couponState.GetCouponCodesByObservationCount(2)`.
  - The identified valid coupons are logged to the console. This separates the logic for identifying valid coupons from the message consumption loop.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - Valid coupon codes, those observed in at least two distinct Kafka topics, are now periodically identified and logged.

**File List:**
- `pipeline/consumer/main.go`

**Change Log:**
- Implemented valid coupon code identification.

**Status:**
- review
