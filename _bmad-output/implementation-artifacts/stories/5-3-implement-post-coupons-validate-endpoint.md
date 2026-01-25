# Story 5.3: Implement POST /coupons/validate endpoint

**Epic:** Coupon & Discount API

**Description:**
As a client, I want to validate a coupon code against my current cart, so that I can see the potential discount before placing an order.

**Acceptance Criteria:**
- The API has a `POST /coupons/validate` endpoint.
- The request body includes the `coupon_code` and optionally the `cart_id` or `DeviceID` to calculate the discount.
- If the `coupon_code` is valid and active, the endpoint returns a JSON object with the coupon's details and the calculated discount amount (or effect, e.g., "free item") for the given cart.
- If the `coupon_code` is invalid, expired, or not applicable to the cart, the endpoint returns a `400 Bad Request` or `404 Not Found` with an appropriate error message.
- The endpoint interacts with the MySQL database to fetch coupon data.
