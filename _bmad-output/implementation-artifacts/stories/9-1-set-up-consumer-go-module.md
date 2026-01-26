# Story 9.1: Set up consumer Go module

**Description:**
As a developer, I want to create a new, independent Go module for the coupon consumer (aggregator) service, so that it can be developed, built, and deployed as a separate microservice.

**Acceptance Criteria:**
- [x] A new directory for the consumer service (e.g., `pipeline/consumer`) is created.
- [x] A `go.mod` file is initialized within this directory.
- [x] The basic `main.go` entry point is set up, ready to integrate Kafka consumption, aggregation, and Redis publishing logic.

**Tasks/Subtasks:**
- [x] Create `pipeline/consumer` directory.
- [x] Initialize Go module (`go mod init kart-challenge/consumer`).
- [x] Create basic `main.go` file.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created the `pipeline/consumer` directory.
  - Initialized a new Go module named `kart-challenge/consumer` within the directory.
  - Created a `main.go` file with a simple "Service started" message.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The consumer Go module is set up and ready for further development.

**File List:**
- `pipeline/consumer/go.mod`
- `pipeline/consumer/main.go`

**Change Log:**
- Set up consumer Go module.

**Status:**
- review