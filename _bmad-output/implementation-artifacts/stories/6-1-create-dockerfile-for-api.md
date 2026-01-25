# Story 6.1: Create Dockerfile for Go API

**Epic:** Deployment & Documentation

**Description:**
As a developer, I want to create a Dockerfile for the Go API application, so that it can be easily containerized, distributed, and deployed consistently across different environments.

**Acceptance Criteria:**
- A `Dockerfile` is created in the project root.
- The Dockerfile builds a self-contained image for the Go API.
- It uses a multi-stage build to minimize the final image size.
- The API application runs as a non-root user inside the container.
- The Dockerfile exposes the configured port for the API (e.g., 8080).
- The Dockerfile includes necessary dependencies for the Go application.
