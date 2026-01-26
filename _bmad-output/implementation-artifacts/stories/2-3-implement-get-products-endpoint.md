# Story 2.3: Implement GET /products endpoint

**Description:**
As a client, I want to retrieve a list of all available products, so that I can browse items to add to my shopping cart.

**Acceptance Criteria:**
- [x] The API has a `GET /products` endpoint.
- [x] Calling this endpoint returns a JSON array of `Product` objects.
- [x] Each `Product` object includes ID, Name, Description, Price, ImageURL, and Stock.
- [x] The endpoint handles cases where no products are found, returning an empty array and a `200 OK` status.
- [x] The endpoint interacts with the MySQL database to fetch product data.

**Tasks/Subtasks:**
- [x] Create `GetProducts` function in `internal/database/database.go` to retrieve all products from the database.
- [x] Create `ProductHandler` with `GetProducts` method in `internal/handler/handler.go`.
- [x] Add `GET /products` endpoint in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added a `GetProducts` function to the `DBClient` in `internal/database/database.go`.
  - Created a `ProductHandler` in `internal/handler/handler.go` with a `GetProducts` method that calls the database function and returns a JSON response.
  - In `cmd/api/main.go`, instantiated the `ProductHandler` and added a `GET /products` route that uses the handler.
- **Debug Log:**
  - No code-related issues. Verification with `curl` continues to be a problem in this environment.
- **Completion Notes:**
  - The `GET /products` endpoint is implemented to retrieve all products from the database.

**File List:**
- `internal/database/database.go`
- `internal/handler/handler.go`
- `cmd/api/main.go`

**Change Log:**
- Implemented `GET /products` endpoint.

**Status:**
- review
