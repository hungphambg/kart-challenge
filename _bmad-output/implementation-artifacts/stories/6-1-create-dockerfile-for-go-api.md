# Story 6.1: Create Dockerfile for Go API

**Description:**
As a developer, I want to create a Dockerfile for the Go API application, so that it can be easily containerized, distributed, and deployed consistently across different environments.

**Acceptance Criteria:**
- [x] A `Dockerfile` is created in the project root.
- [x] The Dockerfile builds a self-contained image for the Go API.
- [x] It uses a multi-stage build to minimize the final image size.
- [x] The API application runs as a non-root user inside the container.
- [x] The Dockerfile exposes the configured port for the API (e.g., 8080).
- [x] The Dockerfile includes necessary dependencies for the Go application.

**Tasks/Subtasks:**
- [x] Create `Dockerfile` in the project root with multi-stage build, non-root user, and exposed port.
- [x] Include `.env` file in the Dockerfile.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created a `Dockerfile` with a multi-stage build. The builder stage compiles the Go application, and the runner stage uses an `alpine:latest` base image for a minimal final image.
  - Configured the runner stage to run as a non-root user (`appuser`).
  - The Dockerfile exposes port `1323` (default application port).
  - Copied `go.mod`, `go.sum`, and project source, including `migrations` and `.env` file.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `Dockerfile` is created, enabling containerization of the Go API application.

**File List:**
- `Dockerfile`

**Change Log:**
- Created `Dockerfile` for Go API.

**Status:**
- review
