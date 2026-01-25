# Story 4.5: Implement GET /orders/{order_id} endpoint

**Epic:** Order Management API

**Description:**
As a client, I want to view the detailed information of a specific past order, so that I can review its contents and status.

**Acceptance Criteria:**
- The API has a `GET /orders/{order_id}` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- If the `order_id` is valid and belongs to the `DeviceID`, it returns a JSON object representing the `Order`, including its `OrderItem` details.
- If the `order_id` is not found or does not belong to the `DeviceID`, the endpoint returns a `404 Not Found` status with an appropriate error message.
- The endpoint interacts with the MySQL database to fetch order data.
