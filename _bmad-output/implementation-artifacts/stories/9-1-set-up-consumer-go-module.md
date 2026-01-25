# Story 9.1: Set up consumer Go module

**Epic:** Coupon Consumer Service

**Description:**
As a developer, I want to create a new, independent Go module for the coupon consumer (aggregator) service, so that it can be developed, built, and deployed as a separate microservice.

**Acceptance Criteria:**
- A new directory for the consumer service (e.g., `pipeline/consumer`) is created.
- A `go.mod` file is initialized within this directory.
- The basic `main.go` entry point is set up, ready to integrate Kafka consumption, aggregation, and Redis publishing logic.
