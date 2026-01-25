# Story 10.1: Implement API Redis client

**Epic:** Live API Coupon Validation

**Description:**
As a developer, I want to implement a Redis client within the main API server, so that it can connect to the Redis instance populated by the coupon ingestion pipeline and perform fast lookups for coupon validation.

**Acceptance Criteria:**
- The main API server's codebase includes a Redis client library for Go (e.g., `go-redis/redis`).
- The API server is configured to connect to the Redis service, using connection details from its configuration.
- A health check or similar mechanism confirms the API server can successfully connect to Redis on startup.
- The Redis client is made available to relevant service layers for coupon validation.
