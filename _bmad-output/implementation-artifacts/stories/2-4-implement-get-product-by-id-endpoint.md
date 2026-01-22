# Story 2.4: Implement GET /products/{id} endpoint

**Epic:** Product Catalog

**Description:**
As a client, I want to retrieve detailed information for a specific product by its ID, so that I can view its individual details before adding it to my cart.

**Acceptance Criteria:**
- The API has a `GET /products/{id}` endpoint, where `{id}` is the product's unique identifier.
- Calling this endpoint with a valid ID returns a JSON object representing the `Product`.
- If the product with the given ID is not found, the endpoint returns a `404 Not Found` status with a simple JSON error object (e.g., `{"code": "PRODUCT_NOT_FOUND", "message": "Product with ID {id} not found."}`).
- The endpoint interacts with the MySQL database to fetch product data.
