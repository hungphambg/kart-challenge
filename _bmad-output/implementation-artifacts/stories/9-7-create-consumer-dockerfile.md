# Story 9.7: Create consumer Dockerfile

**Epic:** Coupon Consumer Service

**Description:**
As a developer, I want to create a `Dockerfile` for the Go coupon consumer service, so that it can be containerized and easily run within the `docker-compose.pipeline.yaml` environment.

**Acceptance Criteria:**
- A `Dockerfile` is created within the consumer service's directory (`pipeline/consumer`).
- The Dockerfile builds a self-contained image for the Go consumer application.
- It uses a multi-stage build to minimize the final image size.
- The consumer application runs as a non-root user inside the container.
- The Dockerfile includes any necessary environment variables or configurations for the consumer (e.g., Kafka broker address, Redis address).
