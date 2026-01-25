# Story 5.1: Define Coupon data model

**Epic:** Coupon & Discount API

**Description:**
As a developer, I want to define a clear data model for a "Coupon" in a Go struct, so that we can accurately represent various types of discount codes, their values, and validity.

**Acceptance Criteria:**
- A `model` package exists within the `internal` directory.
- `coupon.go` file is created within the `model` package.
- The `Coupon` struct is defined with at least:
  - `ID` (e.g., UUID or integer, primary key)
  - `Code` (string, unique, the actual coupon code like "HAPPYHOURS")
  - `Type` (string, e.g., "percentage", "buy_one_get_one", "fixed_amount")
  - `Value` (e.g., float or suitable monetary type, for percentage or fixed amount)
  - `Description` (string, human-readable description of the discount)
  - `ExpiresAt` (timestamp, optional)
  - `IsActive` (boolean)
- Struct includes appropriate JSON tags for serialization/deserialization.
