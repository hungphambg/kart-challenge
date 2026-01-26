# Story 9.2: Implement Kafka consumer subscription

**Description:**
As a developer, I want to implement a Kafka consumer client in Go that subscribes to all three coupon-related topics, so that it can receive all potential coupon codes published by the producer services.

**Acceptance Criteria:**
- [x] The consumer service includes a Kafka client library for Go (`github.com/segmentio/kafka-go`).
- [x] The consumer is configured to subscribe to `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz` topics.
- [x] The consumer is part of a single consumer group to ensure proper message distribution and offset management.
- [x] The consumer is configured to connect to the Kafka broker running in the `docker-compose.pipeline.yaml` setup.
- [x] Basic message consumption and error handling (e.g., logging consumed messages, handling connection errors) are implemented.

**Tasks/Subtasks:**
- [x] Add Kafka client library dependency (`github.com/segmentio/kafka-go`).
- [x] Modify `pipeline/consumer/main.go` to initialize Kafka consumer, subscribe to topics, and implement message consumption loop.

**Dev Agent Record:**
- **Implementation Plan:**
  - Added `github.com/segmentio/kafka-go` as a dependency to the consumer module.
  - Modified `pipeline/consumer/main.go` to initialize a `kafka.Reader` configured to connect to `kafka:9092`.
  - Configured the consumer to subscribe to `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz` topics as part of a single consumer group.
  - Implemented a loop to continuously read messages from Kafka, logging the consumed message details.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The Kafka consumer is implemented, subscribing to all coupon topics and performing basic message consumption.

**File List:**
- `pipeline/consumer/go.mod`
- `pipeline/consumer/go.sum`
- `pipeline/consumer/main.go`

**Change Log:**
- Implemented Kafka consumer subscription.

**Status:**
- review