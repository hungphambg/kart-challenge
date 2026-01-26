# Story 8.5: Implement Kafka producer

**Description:**
As a developer, I want to implement a Kafka producer client in Go, so that identified coupon codes can be reliably published to their respective Kafka topics for further processing.

**Acceptance Criteria:**
- [x] The producer service includes a Kafka client library for Go (`github.com/segmentio/kafka-go`).
- [x] For each identified coupon code, a message is published to the Kafka topic named after the source file (e.g., `couponbase1.gz`).
- [x] Messages are published asynchronously to maximize throughput.
- [x] Error handling is implemented for Kafka connection issues and publishing failures (e.g., retries, logging).
- [x] The producer is configured to connect to the Kafka broker running in the `docker-compose.pipeline.yaml` setup.

**Tasks/Subtasks:**
- [x] Add Kafka client library dependency (`github.com/segmentio/kafka-go`).
- [x] Modify `pipeline/producer/main.go` to initialize Kafka producer, publish messages, and handle errors.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `github.com/segmentio/kafka-go` as a dependency to the producer module.
  - Modified `pipeline/producer/main.go` to initialize a `kafka.Writer` configured for asynchronous publishing to `kafka:9092`.
  - The topic name is dynamically derived from the `.gz` file name.
  - Each identified coupon code is published as a Kafka message, with basic error logging.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The Kafka producer is implemented, enabling identified coupon codes to be published to Kafka topics.

**File List:**
- `pipeline/producer/go.mod`
- `pipeline/producer/go.sum`
- `pipeline/producer/main.go`

**Change Log:**
- Implemented Kafka producer.

**Status:**
- review