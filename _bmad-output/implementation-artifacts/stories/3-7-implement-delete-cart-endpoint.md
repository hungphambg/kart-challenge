# Story 3.7: Implement DELETE /cart endpoint

**Epic:** Shopping Cart API

**Description:**
As a client, I want to be able to empty my entire shopping cart, so that I can start a new order or completely discard my current selection.

**Acceptance Criteria:**
- The API has a `DELETE /cart` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- The endpoint removes all items associated with the user's cart in the MySQL database.
- The endpoint returns a `200 OK` status with an empty cart representation or a success message.
- If no cart exists for the `DeviceID`, it returns a `200 OK` (idempotent operation) or a `404 Not Found` if strictly interpreted. I will lean towards `200 OK` for idempotency.
