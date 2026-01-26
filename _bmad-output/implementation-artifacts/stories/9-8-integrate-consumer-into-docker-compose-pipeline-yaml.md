# Story 9.8: Integrate consumer into docker-compose.pipeline.yaml

**Description:**
As a developer, I want to integrate the coupon consumer service into the `docker-compose.pipeline.yaml` file, so that it can be launched and managed as part of the local pipeline infrastructure.

**Acceptance Criteria:**
- [x] The `docker-compose.pipeline.yaml` file is updated to include a service for the coupon consumer.
- [x] The consumer service's Docker image is built using the Dockerfile from Story 9.7.
- [x] The consumer service is configured to connect to the Kafka and Redis services.
- [x] The consumer service can be successfully started and stopped via `docker-compose`.

**Tasks/Subtasks:**
- [x] Add `consumer` service to `docker-compose.pipeline.yaml`.
- [x] Configure `consumer` service to build from `pipeline/consumer/Dockerfile`.
- [x] Set environment variables (`KAFKA_BROKERS`, `REDIS_ADDR`) for the consumer.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added a `consumer` service to `docker-compose.pipeline.yaml`.
  - Configured the `consumer` service to build from `pipeline/consumer/Dockerfile`.
  - Set the `KAFKA_BROKERS` environment variable to `kafka:9092` and `REDIS_ADDR` to `redis:6379`.
  - Made the `consumer` service depend on `kafka` and `redis`.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The coupon consumer service is integrated into the local pipeline infrastructure, enabling its seamless management with Docker Compose.

**File List:**
- `docker-compose.pipeline.yaml`

**Change Log:**
- Integrated consumer into `docker-compose.pipeline.yaml`.

**Status:**
- review
