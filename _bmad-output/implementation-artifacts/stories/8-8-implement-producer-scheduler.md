# Story 8.8: Implement producer scheduler

**Epic:** Coupon Producer Service

**Description:**
As an operator, I want to implement a scheduling mechanism to automatically trigger the execution of the coupon producer service for each `couponbaseX.gz` file on a regular basis, so that the coupon data in Kafka and Redis is kept up-to-date.

**Acceptance Criteria:**
- A scheduling mechanism (e.g., a simple shell script triggered by host `cron`, or a dedicated scheduler service in `docker-compose`) is implemented.
- The scheduler triggers the producer service for `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz`.
- Each producer run processes its assigned `.gz` file.
- The scheduler runs on a recurring basis (e.g., daily, as per `coupon_ingestion_pipeline.md`).
- The scheduling configuration is documented in the pipeline's README.
```

I will now create this file.