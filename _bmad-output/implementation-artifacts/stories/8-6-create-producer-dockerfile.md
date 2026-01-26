# Story 8.6: Create producer Dockerfile

**Description:**
As a developer, I want to create a `Dockerfile` for the Go coupon producer service, so that it can be containerized and easily run within the `docker-compose.pipeline.yaml` environment.

**Acceptance Criteria:**
- [x] A `Dockerfile` is created within the producer service's directory (`pipeline/producer`).
- [x] The Dockerfile builds a self-contained image for the Go producer application.
- [x] It uses a multi-stage build to minimize the final image size.
- [x] The producer application runs as a non-root user inside the container.
- [x] The Dockerfile includes any necessary environment variables or configurations for the producer (e.g., Kafka broker address).

**Tasks/Subtasks:**
- [x] Create `pipeline/producer/Dockerfile` with multi-stage build, non-root user, and necessary build steps.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `pipeline/producer/Dockerfile` with a multi-stage build process.
  - The builder stage compiles the Go application, and the runner stage uses `alpine:latest` for a minimal image, runs as `appuser`, and copies the built binary.
  - Set `ENTRYPOINT` to `./main`, expecting the filename as an argument.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `Dockerfile` for the coupon producer service is created, enabling its containerization.

**File List:**
- `pipeline/producer/Dockerfile`

**Change Log:**
- Created producer `Dockerfile`.

**Status:**
- review