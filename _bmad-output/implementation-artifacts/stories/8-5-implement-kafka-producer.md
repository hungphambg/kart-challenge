# Story 8.5: Implement Kafka producer

**Epic:** Coupon Producer Service

**Description:**
As a developer, I want to implement a Kafka producer client in Go, so that identified coupon codes can be reliably published to their respective Kafka topics for further processing.

**Acceptance Criteria:**
- The producer service includes a Kafka client library for Go (e.g., `confluentinc/confluent-kafka-go` or `segmentio/kafka-go`).
- For each identified coupon code, a message is published to the Kafka topic named after the source file (e.g., `couponbase1.gz`).
- Messages are published asynchronously to maximize throughput.
- Error handling is implemented for Kafka connection issues and publishing failures (e.g., retries, logging).
- The producer is configured to connect to the Kafka broker running in the `docker-compose.pipeline.yaml` setup.
