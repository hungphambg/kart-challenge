# Story 7.2: Configure Kafka topics on startup

**Description:**
As a developer, I want the Kafka service in `docker-compose.pipeline.yaml` to automatically create the necessary topics for the coupon ingestion pipeline on startup, so that the producer and consumer services have the correct channels to communicate.

**Acceptance Criteria:**
- [x] The `docker-compose.pipeline.yaml` is updated to include configuration for Kafka to create the following topics upon initialization:
  - `couponbase1.gz`
  - `couponbase2.gz`
  - `couponbase3.gz`
- [x] This configuration uses environment variables or a specific Kafka image feature for topic creation.
- [x] The created topics have appropriate replication factors and partitions for a local development setup.

**Tasks/Subtasks:**
- [x] Update `docker-compose.pipeline.yaml` to configure Kafka to create specified topics on startup.

**Dev Agent Record:**
- **Implementation Plan:**
  - Modified the `kafka` service definition in `docker-compose.pipeline.yaml` to include the `KAFKA_CREATE_TOPICS` environment variable.
  - Configured `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz` topics with 1 partition and 1 replica, suitable for local development.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - Kafka topics are now automatically created on startup, facilitating pipeline communication.

**File List:**
- `docker-compose.pipeline.yaml`

**Change Log:**
- Configured Kafka topic creation on startup.

**Status:**
- review
