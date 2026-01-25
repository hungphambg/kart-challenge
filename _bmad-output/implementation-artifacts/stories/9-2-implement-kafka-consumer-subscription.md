# Story 9.2: Implement Kafka consumer subscription

**Epic:** Coupon Consumer Service

**Description:**
As a developer, I want to implement a Kafka consumer client in Go that subscribes to all three coupon-related topics, so that it can receive all potential coupon codes published by the producer services.

**Acceptance Criteria:**
- The consumer service includes a Kafka client library for Go (e.g., `confluentinc/confluent-kafka-go` or `segmentio/kafka-go`).
- The consumer is configured to subscribe to `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz` topics.
- The consumer is part of a single consumer group to ensure proper message distribution and offset management.
- The consumer is configured to connect to the Kafka broker running in the `docker-compose.pipeline.yaml` setup.
- Basic message consumption and error handling (e.g., logging consumed messages, handling connection errors) are implemented.
