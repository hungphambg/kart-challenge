# Epic 2: Product Catalog

This epic covers all aspects of managing and retrieving product information.

---

### Story 2.1: Define the Product data model

**Description:**
As a developer, I want to define a clear and comprehensive data model for a "Product" in a Go struct, so that we have a consistent representation of products throughout the application.

**Acceptance Criteria:**
- A `model` package exists within the `internal` directory.
- A `product.go` file is created within the `model` package.
- The `Product` struct is defined in `product.go` with the following fields:
  - `ID` (e.g., UUID or integer)
  - `Name` (string)
  - `Description` (string)
  - `Price` (e.g., float or a suitable monetary type)
  - `ImageURL` (string)
  - `Stock` (integer)
- The struct includes appropriate JSON tags for serialization/deserialization.

---

### Story 2.2: Create the products database schema

**Description:**
As a developer, I want to create a database schema for the `products` table in MySQL, so that we can persist and query product information. I also want to create a migration script to set up the table.

**Acceptance Criteria:**
- A `migrations` directory is created in the project root.
- A SQL migration script is created to define the `products` table.
- The `products` table schema includes columns that correspond to the `Product` data model (ID, name, description, price, image_url, stock).
- The data types for the columns are appropriate for MySQL (e.g., `VARCHAR`, `TEXT`, `DECIMAL`, `INT`).
- A primary key is defined for the `products` table.

---

### Story 2.3: Implement GET /products endpoint

**Description:**
As a client, I want to retrieve a list of all available products, so that I can browse items to add to my shopping cart.

**Acceptance Criteria:**
- The API has a `GET /products` endpoint.
- Calling this endpoint returns a JSON array of `Product` objects.
- Each `Product` object includes ID, Name, Description, Price, ImageURL, and Stock.
- The endpoint handles cases where no products are found, returning an empty array and a `200 OK` status.
- The endpoint interacts with the MySQL database to fetch product data.

---

### Story 2.4: Implement GET /products/{id} endpoint

**Description:**
As a client, I want to retrieve detailed information for a specific product by its ID, so that I can view its individual details before adding it to my cart.

**Acceptance Criteria:**
- The API has a `GET /products/{id}` endpoint, where `{id}` is the product's unique identifier.
- Calling this endpoint with a valid ID returns a JSON object representing the `Product`.
- If the product with the given ID is not found, the endpoint returns a `404 Not Found` status with a simple JSON error object (e.g., `{"code": "PRODUCT_NOT_FOUND", "message": "Product with ID {id} not found."}`).
- The endpoint interacts with the MySQL database to fetch product data.
