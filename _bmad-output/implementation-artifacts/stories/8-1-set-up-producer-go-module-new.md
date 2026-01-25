# Story 8.1: Set up producer Go module

**Epic:** Coupon Producer Service

**Description:**
As a developer, I want to create a new, independent Go module for the coupon producer service, so that it can be developed, built, and deployed as a separate microservice.

**Acceptance Criteria:**
- A new directory for the producer service (e.g., `pipeline/producer`) is created.
- A `go.mod` file is initialized within this directory.
- The basic `main.go` entry point is set up, ready to integrate file reading and Kafka publishing logic.
