# Story 9.7: Create consumer Dockerfile

**Description:**
As a developer, I want to create a `Dockerfile` for the Go coupon consumer service, so that it can be containerized and easily run within the `docker-compose.pipeline.yaml` environment.

**Acceptance Criteria:**
- [x] A `Dockerfile` is created within the consumer service's directory.
- [x] It uses a multi-stage build to minimize the final image size.
- [x] The consumer application runs as a non-root user inside the container.
- [x] The Dockerfile includes necessary environment variables or configurations.

**Tasks/Subtasks:**
- [x] Create `pipeline/consumer/Dockerfile` with multi-stage build, non-root user, and necessary build steps.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `pipeline/consumer/Dockerfile` with a multi-stage build process.
  - The builder stage compiles the Go application, and the runner stage uses `alpine:latest` for a minimal image, runs as `appuser`, and copies the built binary.
  - Set `ENTRYPOINT` to `./main`.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `Dockerfile` for the coupon consumer service is created, enabling its containerization.

**File List:**
- `pipeline/consumer/Dockerfile`

**Change Log:**
- Created consumer `Dockerfile`.

**Status:**
- review