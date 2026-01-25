# Story 5.2: Create coupons database schema

**Epic:** Coupon & Discount API

**Description:**
As a developer, I want to create the database schema for the `coupons` table in MySQL, so that we can persist and manage coupon codes and their properties. This includes creating migration scripts.

**Acceptance Criteria:**
- SQL migration scripts are created to define the `coupons` table.
- The `coupons` table schema includes columns for `ID`, `Code`, `Type`, `Value`, `Description`, `ExpiresAt`, `IsActive`.
- The `Code` column is unique.
- Appropriate data types are used for MySQL (e.g., `VARCHAR`, `TEXT`, `DECIMAL`, `DATETIME`, `BOOLEAN`).
- A primary key is defined for the table.
