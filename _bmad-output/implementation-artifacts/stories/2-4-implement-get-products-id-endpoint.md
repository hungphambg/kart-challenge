# Story 2.4: Implement GET /products/{id} endpoint

**Description:**
As a client, I want to retrieve detailed information for a specific product by its ID, so that I can view its individual details before adding it to my cart.

**Acceptance Criteria:**
- [x] The API has a `GET /products/{id}` endpoint, where `{id}` is the product's unique identifier.
- [x] Calling this endpoint with a valid ID returns a JSON object representing the `Product`.
- [x] If the product with the given ID is not found, the endpoint returns a `404 Not Found` status with a simple JSON error object (e.g., `{"code": "PRODUCT_NOT_FOUND", "message": "Product with ID {id} not found."}`).
- [x] The endpoint interacts with the MySQL database to fetch product data.

**Tasks/Subtasks:**
- [x] Create `GetProductByID` function in `internal/database/database.go` to retrieve a single product from the database by its ID.
- [x] Create `GetProduct` method in `ProductHandler` in `internal/handler/handler.go`.
- [x] Add `GET /products/:id` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added a `GetProductByID` function to the `DBClient` in `internal/database/database.go` that handles the database query and `sql.ErrNoRows`.
  - Added a `GetProduct` method to the `ProductHandler` in `internal/handler/handler.go` to parse the ID, call the database function, and handle the "not found" case by returning a 404.
  - In `cmd/api/main.go`, added a `GET /products/:id` route that uses the `GetProduct` handler method.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `GET /products/{id}` endpoint is implemented to retrieve a single product from the database by its ID.

**File List:**
- `internal/database/database.go`
- `internal/handler/handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `GET /products/{id}` endpoint.

**Status:**
- review
