# Story 5.1: Define Coupon data model

**Description:**
As a developer, I want to define a clear data model for a "Coupon" in a Go struct, so that we can accurately represent various types of discount codes, their values, and validity.

**Acceptance Criteria:**
- [x] A `model` package exists within the `internal` directory.
- [x] `coupon.go` file is created within the `model` package.
- [x] The `Coupon` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `Code` (string, unique, the actual coupon code like "HAPPYHOURS")
  - `Type` (string, e.g., "percentage", "buy_one_get_one", "fixed_amount")
  - `Value` (e.g., float or suitable monetary type, for percentage or fixed amount)
  - `Description` (string, human-readable description of the discount)
  - `ExpiresAt` (timestamp, optional)
  - `IsActive` (boolean)
- [x] Struct includes appropriate JSON tags for serialization/deserialization.

**Tasks/Subtasks:**
- [x] Create `internal/model/coupon.go` with the `Coupon` struct.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `internal/model/coupon.go` and defined the `Coupon` struct with specified fields, JSON tags, and GORM compatibility tags (e.g., `primaryKey`, `unique`, `not null`).
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `Coupon` data model is defined and is GORM compatible.

**File List:**
- `internal/model/coupon.go`

**Change Log:**
- Defined the `Coupon` data model.

**Status:**
- review