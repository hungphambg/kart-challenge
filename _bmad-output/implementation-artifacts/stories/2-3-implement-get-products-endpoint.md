# Story 2.3: Implement GET /products endpoint

**Epic:** Product Catalog

**Description:**
As a client, I want to retrieve a list of all available products, so that I can browse items to add to my shopping cart.

**Acceptance Criteria:**
- The API has a `GET /products` endpoint.
- Calling this endpoint returns a JSON array of `Product` objects.
- Each `Product` object includes ID, Name, Description, Price, ImageURL, and Stock.
- The endpoint handles cases where no products are found, returning an empty array and a `200 OK` status.
- The endpoint interacts with the MySQL database to fetch product data.
