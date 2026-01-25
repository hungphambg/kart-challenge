# Story 9.8: Integrate consumer into docker-compose.pipeline.yaml

**Epic:** Coupon Consumer Service

**Description:**
As a developer, I want to integrate the coupon consumer service into the `docker-compose.pipeline.yaml` file, so that it can be launched and managed as part of the local pipeline infrastructure alongside Kafka and Redis.

**Acceptance Criteria:**
- The `docker-compose.pipeline.yaml` file is updated to include a service for the coupon consumer.
- The consumer service's Docker image is built using the Dockerfile created in Story 9.7.
- The consumer service is configured to connect to the Kafka broker and Redis services defined in the same `docker-compose` file.
- The consumer service can be successfully started and stopped via `docker-compose`.
- Logging output from the consumer indicates successful connection to Kafka and Redis, and consumption of messages.
