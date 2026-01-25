# Story 8.7: Integrate producer into docker-compose.pipeline.yaml

**Epic:** Coupon Producer Service

**Description:**
As a developer, I want to integrate the coupon producer service into the `docker-compose.pipeline.yaml` file, so that it can be launched and managed as part of the local pipeline infrastructure alongside Kafka and Redis.

**Acceptance Criteria:**
- The `docker-compose.pipeline.yaml` file is updated to include a service for the coupon producer.
- The producer service's Docker image is built using the Dockerfile created in Story 8.6.
- The producer service is configured to connect to the Kafka broker service defined in the same `docker-compose` file.
- The producer service can be successfully started and stopped via `docker-compose`.
- The `couponbaseX.gz` files are mounted as volumes into the producer container, so it can read them.
