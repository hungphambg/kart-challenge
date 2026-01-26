# Story 5.2: Create coupons database schema

**Description:**
As a developer, I want to create the database schema for the `coupons` table in MySQL, so that we can persist and manage coupon codes and their properties. This includes creating migration scripts.

**Acceptance Criteria:**
- [x] SQL migration scripts are created to define the `coupons` table. (Implemented via GORM AutoMigrate)
- [x] The `coupons` table schema includes columns for `ID`, `Code`, `Type`, `Value`, `Description`, `ExpiresAt`, `IsActive`.
- [x] The `Code` column is unique.
- [x] Appropriate data types are used for MySQL (e.g., `VARCHAR`, `TEXT`, `DECIMAL`, `DATETIME`, `BOOLEAN`).
- [x] A primary key is defined for the table.

**Tasks/Subtasks:**
- [x] Add `model.Coupon{}` to the `AutoMigrate` calls in `cmd/api/main.go`.
- [x] Rely on GORM's `AutoMigrate` feature for schema creation.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `model.Coupon{}` to the `AutoMigrate` calls in `cmd/api/main.go`.
  - Leveraged GORM's `AutoMigrate` functionality to create the `coupons` table based on the `model.Coupon` struct.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `coupons` database schema is managed through GORM's `AutoMigrate` based on the defined `Coupon` model.

**File List:**
- `cmd/api/main.go` (modified)
- `internal/model/coupon.go` (created in Story 5.1)

**Change Log:**
- Implemented `coupons` database schema using GORM AutoMigrate.

**Status:**
- review