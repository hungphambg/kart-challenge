# Story 2.2: Create the products database schema

**Epic:** Product Catalog

**Description:**
As a developer, I want to create a database schema for the `products` table in PostgreSQL, so that we can persist and query product information. I also want to create a migration script to set up the table.

**Acceptance Criteria:**
- A `migrations` directory is created in the project root.
- A SQL migration script is created to define the `products` table.
- The `products` table schema includes columns that correspond to the `Product` data model (ID, name, description, price, image_url, stock).
- The data types for the columns are appropriate for PostgreSQL (e.g., `VARCHAR`, `TEXT`, `NUMERIC`, `INTEGER`).
- A primary key is defined for the `products` table.
