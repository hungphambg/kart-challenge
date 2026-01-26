# Story 7.1: Define docker-compose.pipeline.yaml

**Description:**
As a developer, I want to create a `docker-compose.pipeline.yaml` file that sets up Kafka, Zookeeper, and Redis services, so that I can easily run the coupon ingestion pipeline's infrastructure locally.

**Acceptance Criteria:**
- [x] A `docker-compose.pipeline.yaml` file is created in the project root (or a dedicated `pipeline` directory).
- [x] The file defines services for `zookeeper`, `kafka`, and `redis`.
- [x] `zookeeper` and `kafka` services are configured to communicate correctly.
- [x] `redis` service is accessible to other services.
- [x] All services are configured with appropriate ports and network settings for inter-service communication.
- [x] Persistent volumes are configured for Kafka and Redis where necessary to retain data across restarts.

**Tasks/Subtasks:**
- [x] Create `docker-compose.pipeline.yaml` with Zookeeper, Kafka, and Redis services.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `docker-compose.pipeline.yaml` in the project root.
  - Defined `zookeeper` service using `confluentinc/cp-zookeeper`.
  - Defined `kafka` service using `confluentinc/cp-kafka`, configured to connect to Zookeeper, with advertised listeners for local access, and persistent volume.
  - Defined `redis` service using `redis:7.0.0-alpine`, with persistent volume.
  - Configured ports and environment variables for inter-service communication.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The `docker-compose.pipeline.yaml` is created, providing a local infrastructure setup for the coupon ingestion pipeline.

**File List:**
- `docker-compose.pipeline.yaml`

**Change Log:**
- Created `docker-compose.pipeline.yaml`.

**Status:**
- review