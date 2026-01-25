# Story 4.4: Implement GET /orders endpoint

**Epic:** Order Management API

**Description:**
As a client, I want to view a history of all my past orders, so that I can keep track of my purchases.

**Acceptance Criteria:**
- The API has a `GET /orders` endpoint.
- The endpoint requires a `DeviceID` in the HTTP header.
- The endpoint returns a JSON array of `Order` objects associated with the `DeviceID`.
- Each `Order` object includes at least its `ID`, `Status`, `TotalAmount`, `CreatedAt`.
- The endpoint handles cases where no orders are found for the `DeviceID`, returning an empty array and a `200 OK` status.
- The endpoint interacts with the MySQL database to fetch order data.
