# Story 3.3: Implement GET /cart endpoint

**Epic:** Shopping Cart API

**Description:**
As a client, I want to retrieve my current shopping cart. If I don't have one, I want the system to create a new empty cart for me, so that I can start adding items.

**Acceptance Criteria:**
- The API has a `GET /cart` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- If a cart exists for the given `DeviceID`, it is returned as a JSON object, including its `CartItems`.
- If no cart exists for the `DeviceID`, a new empty cart is created in the MySQL database and returned.
- The response includes the cart's unique identifier and its current list of items.
- The endpoint handles cases where the `DeviceID` is missing, returning a `400 Bad Request` with an appropriate error message.
