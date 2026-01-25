# Epic 7: Pipeline Infrastructure

This epic covers setting up the local development environment for the coupon ingestion pipeline using Docker, including Kafka, Zookeeper, and Redis.

---

### Story 7.1: Define docker-compose.pipeline.yaml

**Description:**
As a developer, I want to create a `docker-compose.pipeline.yaml` file that sets up Kafka, Zookeeper, and Redis services, so that I can easily run the coupon ingestion pipeline's infrastructure locally.

**Acceptance Criteria:**
- A `docker-compose.pipeline.yaml` file is created in the project root (or a dedicated `pipeline` directory).
- The file defines services for `zookeeper`, `kafka`, and `redis`.
- `zookeeper` and `kafka` services are configured to communicate correctly.
- `redis` service is accessible to other services.
- All services are configured with appropriate ports and network settings for inter-service communication.
- Persistent volumes are configured for Kafka and Redis where necessary to retain data across restarts.

---

### Story 7.2: Configure Kafka topics on startup

**Description:**
As a developer, I want the Kafka service in `docker-compose.pipeline.yaml` to automatically create the necessary topics for the coupon ingestion pipeline on startup, so that the producer and consumer services have the correct channels to communicate.

**Acceptance Criteria:**
- The `docker-compose.pipeline.yaml` is updated to include configuration for Kafka to create the following topics upon initialization:
  - `couponbase1.gz`
  - `couponbase2.gz`
  - `couponbase3.gz`
- This configuration uses environment variables or a specific Kafka image feature for topic creation.
- The created topics have appropriate replication factors and partitions for a local development setup.

---

### Story 7.3: Write Pipeline README

**Description:**
As a developer, I want to create a `README.md` file specifically for the coupon ingestion pipeline, so that other developers can easily understand how to set up, run, and interact with the pipeline's infrastructure locally.

**Acceptance Criteria:**
- A `README.md` file is created within the pipeline's dedicated directory (e.g., `pipeline/README.md`).
- The README includes instructions on how to use `docker-compose -f docker-compose.pipeline.yaml up` to start the pipeline infrastructure.
- It details how to stop the services.
- It provides commands or instructions on how to check the status of Kafka topics (e.g., list topics, view messages) and Redis (e.g., connect to CLI, check keys).
- It explains how to access logs for each service within the `docker-compose` setup.
- It mentions any prerequisites for running the pipeline (e.g., Docker).
