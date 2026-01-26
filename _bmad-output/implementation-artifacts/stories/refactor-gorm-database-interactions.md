# Refactor: Implement GORM for Database Interactions

**Description:**
As a developer, I want to refactor the database access layer to use GORM, so that we can leverage an ORM for more structured and less error-prone database operations.

**Acceptance Criteria:**
- [x] GORM and its MySQL driver are added as project dependencies.
- [x] The `DBClient` in `internal/database/database.go` is updated to wrap a `gorm.DB` instance.
- [x] All database interaction functions in `internal/database/database.go` are refactored to use GORM's API.
- [x] All relevant data models (`Product`, `Cart`, `CartItem`) are updated to be compatible with GORM conventions (e.g., `gorm:"primaryKey"`, `gorm:"foreignKey"`).
- [x] `cmd/api/main.go` includes `AutoMigrate` calls for the GORM models to ensure schema creation/updates.

**Tasks/Subtasks:**
- [x] Add GORM dependency: `go get gorm.io/gorm`.
- [x] Add GORM MySQL driver dependency: `go get gorm.io/driver/mysql`.
- [x] Run `go mod tidy`.
- [x] Refactor `internal/database/database.go` to use `*gorm.DB` and GORM API for all functions.
- [x] Update `internal/model/product.go` for GORM compatibility.
- [x] Update `internal/model/cart.go` for GORM compatibility (including `Items` foreign key).
- [x] Update `internal/model/cart_item.go` for GORM compatibility.
- [x] Add `AutoMigrate` calls for `Product`, `Cart`, `CartItem` in `cmd/api/main.go`.
- [x] Import `model` package in `cmd/api/main.go`.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added GORM and its MySQL driver.
  - Rewrote `internal/database/database.go` to use GORM for all CRUD operations and database initialization.
  - Adjusted model structs (`Product`, `Cart`, `CartItem`) with GORM tags to define primary keys and relationships.
  - Integrated GORM `AutoMigrate` into `cmd/api/main.go` for automatic schema management.
- **Debug Log:**
  - No code errors during refactoring. Verification via `curl` for API endpoints and `AutoMigrate` for database changes was not performed due to persistent environment connectivity issues and lack of a live MySQL instance.
- **Completion Notes:**
  - The database layer has been successfully refactored to use GORM, replacing raw SQL queries with GORM's ORM capabilities. This enhances code readability, maintainability, and reduces boilerplate.

**File List:**
- `go.mod` (modified)
- `go.sum` (modified)
- `internal/database/database.go` (modified)
- `internal/model/product.go` (modified)
- `internal/model/cart.go` (modified)
- `internal/model/cart_item.go` (modified)
- `cmd/api/main.go` (modified)

**Change Log:**
- Refactored database interactions to use GORM.

**Status:**
- review
