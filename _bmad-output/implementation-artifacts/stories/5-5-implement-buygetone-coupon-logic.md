# Story 5.5: Implement BUYGETONE coupon logic

**Epic:** Coupon & Discount API

**Description:**
As a coupon user, I want the `BUYGETONE` coupon to make one lowest-priced item in my cart free, so that I can get a free product.

**Acceptance Criteria:**
- The `POST /coupons/validate` endpoint correctly identifies the `BUYGETONE` coupon.
- When `BUYGETONE` is applied to a cart, one instance of the lowest-priced item in the cart is made free.
- If multiple items share the same lowest price, only one of them (the first encountered or chosen by a tie-breaking rule) is discounted.
- The discount is reflected in the response of the validate endpoint.
- The `BUYGETONE` coupon is valid and active.
