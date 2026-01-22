# Story 1.3: Establish database connection and health check

**Epic:** Core API Infrastructure

**Description:**
As a developer, I want to establish a connection to the PostgreSQL database using the loaded configuration and expose a `/health` endpoint, so that I can verify that the API is running and the database connection is healthy.

**Acceptance Criteria:**
- A `database` package is created within the `internal` directory.
- The application connects to the PostgreSQL database using the credentials from the configuration module.
- A `/health` endpoint is available on the API.
- When the `/health` endpoint is called, it returns a `200 OK` status if the database connection is successful.
- If the database connection fails, the endpoint returns a `503 Service Unavailable` status.
