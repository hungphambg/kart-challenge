# Story 10.1: Implement API Redis client

**Description:**
As a developer, I want to implement a Redis client within the main API server, so that it can connect to the Redis instance populated by the coupon ingestion pipeline and perform fast lookups for coupon validation.

**Acceptance Criteria:**
- [x] The main API server's codebase includes a Redis client library for Go (`github.com/go-redis/redis/v8`).
- [x] The API server is configured to connect to the Redis service, using connection details from its configuration.
- [x] A health check or similar mechanism confirms the API server can successfully connect to Redis on startup.
- [x] The Redis client is made available to relevant service layers for coupon validation.

**Tasks/Subtasks:**
- [x] Add Redis client library dependency (`github.com/go-redis/redis/v8`).
- [x] Update `internal/config/config.go` with `RedisAddr` field.
- [x] Modify `cmd/api/main.go` to initialize Redis client, perform ping on startup, and pass client to `CouponHandler`.
- [x] Update `internal/handler/coupon_handler.go` to include Redis client in `CouponHandler` struct.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `github.com/go-redis/redis/v8` dependency to the main module.
  - Added `RedisAddr` field to `internal/config/config.go` to load Redis address from environment variables.
  - Modified `cmd/api/main.go` to initialize a `redis.Client`, perform a `Ping()` to ensure connectivity, and pass this client to the `CouponHandler`.
  - Updated the `CouponHandler` struct in `internal/handler/coupon_handler.go` to include a `Redis` client field.
- **Debug Log:**
  - No code-related issues. Verification for Redis connectivity on startup is in place but requires a running Redis instance to fully confirm.
- **Completion Notes:**
  - The API server now includes a Redis client, is configured to connect to Redis, and checks connectivity on startup.

**File List:**
- `go.mod`
- `go.sum`
- `internal/config/config.go`
- `cmd/api/main.go`
- `internal/handler/coupon_handler.go`

**Change Log:**
- Implemented API Redis client.

**Status:**
- review