# Story 8.6: Create producer Dockerfile

**Epic:** Coupon Producer Service

**Description:**
As a developer, I want to create a `Dockerfile` for the Go coupon producer service, so that it can be containerized and easily run within the `docker-compose.pipeline.yaml` environment.

**Acceptance Criteria:**
- A `Dockerfile` is created within the producer service's directory (`pipeline/producer`).
- The Dockerfile builds a self-contained image for the Go producer application.
- It uses a multi-stage build to minimize the final image size.
- The producer application runs as a non-root user inside the container.
- The Dockerfile includes any necessary environment variables or configurations for the producer (e.g., Kafka broker address).
