# Story 7.1: Define docker-compose.pipeline.yaml

**Epic:** Pipeline Infrastructure

**Description:**
As a developer, I want to create a `docker-compose.pipeline.yaml` file that sets up Kafka, Zookeeper, and Redis services, so that I can easily run the coupon ingestion pipeline's infrastructure locally.

**Acceptance Criteria:**
- A `docker-compose.pipeline.yaml` file is created in the project root (or a dedicated `pipeline` directory).
- The file defines services for `zookeeper`, `kafka`, and `redis`.
- `zookeeper` and `kafka` services are configured to communicate correctly.
- `redis` service is accessible to other services.
- All services are configured with appropriate ports and network settings for inter-service communication.
- Persistent volumes are configured for Kafka and Redis where necessary to retain data across restarts.
