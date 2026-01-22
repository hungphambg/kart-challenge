# Story 2.1: Define the Product data model

**Epic:** Product Catalog

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
