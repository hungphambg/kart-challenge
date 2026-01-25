# Story 7.2: Configure Kafka topics on startup

**Epic:** Pipeline Infrastructure

**Description:**
As a developer, I want the Kafka service in `docker-compose.pipeline.yaml` to automatically create the necessary topics for the coupon ingestion pipeline on startup, so that the producer and consumer services have the correct channels to communicate.

**Acceptance Criteria:**
- The `docker-compose.pipeline.yaml` is updated to include configuration for Kafka to create the following topics upon initialization:
  - `couponbase1.gz`
  - `couponbase2.gz`
  - `couponbase3.gz`
- This configuration uses environment variables or a specific Kafka image feature for topic creation.
- The created topics have appropriate replication factors and partitions for a local development setup.
```

I will now create this file.