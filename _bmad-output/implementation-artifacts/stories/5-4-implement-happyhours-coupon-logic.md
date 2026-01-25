# Story 5.4: Implement HAPPYHOURS coupon logic

**Epic:** Coupon & Discount API

**Description:**
As a coupon user, I want the `HAPPYHOURS` coupon to apply an 18% discount to my total cart value, so that I can save money on my order.

**Acceptance Criteria:**
- The `POST /coupons/validate` endpoint correctly identifies the `HAPPYHOURS` coupon.
- When `HAPPYHOURS` is applied to a cart, the calculated discount is 18% of the cart's subtotal.
- The discount is reflected in the response of the validate endpoint.
- The `HAPPYHOURS` coupon is valid and active.
