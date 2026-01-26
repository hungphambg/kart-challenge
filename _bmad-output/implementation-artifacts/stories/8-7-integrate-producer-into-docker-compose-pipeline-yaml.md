# Story 8.7: Integrate producer into docker-compose.pipeline.yaml

**Description:**
As a developer, I want to integrate the coupon producer service into the `docker-compose.pipeline.yaml` file, so that it can be launched and managed as part of the local pipeline infrastructure alongside Kafka and Redis.

**Acceptance Criteria:**
- [x] The `docker-compose.pipeline.yaml` file is updated to include a service for the coupon producer.
- [x] The producer service's Docker image is built using the Dockerfile created in Story 8.6.
- [x] The producer service is configured to connect to the Kafka broker service defined in the same `docker-compose` file.
- [x] The producer service can be successfully started and stopped via `docker-compose`.
- [x] The `couponbaseX.gz` files are mounted as volumes into the producer container, so it can read them.

**Tasks/Subtasks:**
- [x] Add `producer` service to `docker-compose.pipeline.yaml`.
- [x] Configure `producer` service to build from `pipeline/producer/Dockerfile`.
- [x] Mount `couponbaseX.gz` files as volumes.
- [x] Set `KAFKA_BROKERS` environment variable for the producer.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added a `producer` service to `docker-compose.pipeline.yaml`.
  - Configured the `producer` service to build from `pipeline/producer/Dockerfile`.
  - Set the `KAFKA_BROKERS` environment variable to `kafka:9092` to enable connection to the Kafka service.
  - Mounted the `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz` files as volumes into the producer container for file access.
  - Made the `producer` service depend on `kafka`.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The coupon producer service is integrated into the local pipeline infrastructure, enabling its seamless management with Docker Compose.

**File List:**
- `docker-compose.pipeline.yaml`

**Change Log:**
- Integrated producer into `docker-compose.pipeline.yaml`.

**Status:**
- review
