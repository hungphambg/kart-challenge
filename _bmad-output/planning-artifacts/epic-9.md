# Epic 9: Coupon Consumer Service

This epic covers the creation of the Go stream-processing application that consumes from Kafka, aggregates the codes, identifies valid ones, and writes them to Redis.

---

### Story 9.1: Set up consumer Go module

**Description:**
As a developer, I want to create a new, independent Go module for the coupon consumer (aggregator) service, so that it can be developed, built, and deployed as a separate microservice.

**Acceptance Criteria:**
- A new directory for the consumer service (e.g., `pipeline/consumer`) is created.
- A `go.mod` file is initialized within this directory.
- The basic `main.go` entry point is set up, ready to integrate Kafka consumption, aggregation, and Redis publishing logic.

---

### Story 9.2: Implement Kafka consumer subscription

**Description:**
As a developer, I want to implement a Kafka consumer client in Go that subscribes to all three coupon-related topics, so that it can receive all potential coupon codes published by the producer services.

**Acceptance Criteria:**
- The consumer service includes a Kafka client library for Go (e.g., `confluentinc/confluent-kafka-go` or `segmentio/kafka-go`).
- The consumer is configured to subscribe to `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz` topics.
- The consumer is part of a single consumer group to ensure proper message distribution and offset management.
- The consumer is configured to connect to the Kafka broker running in the `docker-compose.pipeline.yaml` setup.
- Basic message consumption and error handling (e.g., logging consumed messages, handling connection errors) are implemented.

---

### Story 9.3: Integrate Go stream-processing library

**Description:**
As a developer, I want to integrate a Go stream-processing library (or implement a robust manual state management solution) to handle the continuous consumption of coupon codes from Kafka, enabling stateful aggregation and fault tolerance.

**Acceptance Criteria:**
- A suitable Go stream-processing library is chosen and integrated.
- The consumer logic leverages the chosen library's features for stateful processing and fault tolerance.
- The solution provides mechanisms to reprocess messages from Kafka in case of consumer restarts, ensuring no data loss in state aggregation.
- The stream processing handles potential out-of-order messages if necessary.

---

### Story 9.4: Implement consumer state management

**Description:**
As a developer, I want to implement (or leverage the stream-processing library's feature for) state management to track which Kafka topics/source files each unique coupon code has been observed in, so that we can determine its validity.

**Acceptance Criteria:**
- The consumer maintains a data structure (e.g., a hash map or a KTable) where the key is the coupon code and the value tracks the unique source files.
- This state is updated continuously as messages are consumed from Kafka.
- The state management mechanism is fault-tolerant, ensuring that the state is not lost upon consumer restarts.

---

### Story 9.5: Implement valid coupon code identification

**Description:**
As a developer, I want to implement logic to process the aggregated state and identify which coupon codes have been observed in at least two distinct Kafka topics, so that we can determine the truly valid coupons.

**Acceptance Criteria:**
- The consumer periodically (or upon state change) iterates through its maintained state.
- For each coupon code, it checks if it has been associated with two or more unique Kafka topic identifiers.
- A list or set of "valid" coupon codes is generated based on this rule.

---

### Story 9.6: Implement Redis client for valid codes

**Description:**
As a developer, I want to implement a Redis client in Go to store the identified valid coupon codes in a Redis `Set`, so that the live API can perform fast, O(1) lookups for coupon validation.

**Acceptance Criteria:**
- The consumer service includes a Redis client library for Go.
- The identified set of valid coupon codes is written to a Redis `Set`.
- The update operation to Redis is atomic (e.g., writing to a temporary set and then atomically renaming it).
- Error handling for Redis connection issues and write failures is implemented.
- The Redis client is configured to connect to the Redis service in the `docker-compose.pipeline.yaml` setup.

---

### Story 9.7: Create consumer Dockerfile

**Description:**
As a developer, I want to create a `Dockerfile` for the Go coupon consumer service, so that it can be containerized and easily run within the `docker-compose.pipeline.yaml` environment.

**Acceptance Criteria:**
- A `Dockerfile` is created within the consumer service's directory.
- It uses a multi-stage build to minimize the final image size.
- The consumer application runs as a non-root user inside the container.
- The Dockerfile includes necessary environment variables or configurations.

---

### Story 9.8: Integrate consumer into docker-compose.pipeline.yaml

**Description:**
As a developer, I want to integrate the coupon consumer service into the `docker-compose.pipeline.yaml` file, so that it can be launched and managed as part of the local pipeline infrastructure.

**Acceptance Criteria:**
- The `docker-compose.pipeline.yaml` file is updated to include a service for the coupon consumer.
- The consumer service's Docker image is built using the Dockerfile from Story 9.7.
- The consumer service is configured to connect to the Kafka and Redis services.
- The consumer service can be successfully started and stopped via `docker-compose`.
