# Epic 6: Deployment & Documentation

This epic covers the necessary work to containerize the application, set up a local development environment, and provide comprehensive documentation.

---

### Story 6.1: Create Dockerfile for Go API

**Description:**
As a developer, I want to create a Dockerfile for the Go API application, so that it can be easily containerized, distributed, and deployed consistently across different environments.

**Acceptance Criteria:**
- A `Dockerfile` is created in the project root.
- The Dockerfile builds a self-contained image for the Go API.
- It uses a multi-stage build to minimize the final image size.
- The API application runs as a non-root user inside the container.
- The Dockerfile exposes the configured port for the API (e.g., 8080).
- The Dockerfile includes necessary dependencies for the Go application.

---

### Story 6.2: Create docker-compose for local development

**Description:**
As a developer, I want to create a `docker-compose.yaml` file, so that I can easily set up and run the Go API and its MySQL database locally with a single command.

**Acceptance Criteria:**
- A `docker-compose.yaml` file is created in the project root.
- The `docker-compose.yaml` defines two services: one for the Go API and one for the MySQL database.
- The Go API service uses the `Dockerfile` created in Story 6.1.
- The MySQL service uses an official MySQL Docker image.
- Environment variables for the API (e.g., database connection string) are configured to connect to the MySQL service within the Docker network.
- Volumes are used for persistent MySQL data.
- The setup allows for easy local development and testing.

---

### Story 6.3: Update README with instructions

**Description:**
As a developer, I want to update the `README.md` file with clear instructions on how to set up, build, run, and test the application, so that other developers (and reviewers) can easily get started with the project.

**Acceptance Criteria:**
- The `README.md` file in the project root is updated.
- It includes a "Getting Started" section with prerequisites (Go, Docker).
- It provides instructions on how to use `docker-compose` to start the API and database.
- It provides instructions on how to run database migrations.
- It explains how to build and run the Go application directly (without Docker, for development convenience).
- It details how to run any implemented tests.
- It includes a brief overview of the API endpoints and how to interact with them.

---

### Story 6.4: Ensure API documentation

**Description:**
As a developer, I want to ensure that the API documentation, based on the provided OpenAPI specification, is accessible and up-to-date, so that consumers of the API can easily understand how to interact with it.

**Acceptance Criteria:**
- The `openapi.yaml` file (`api/openapi.yaml`) is kept up-to-date with the implemented API endpoints.
- Instructions are provided in the `README.md` (or a dedicated `docs` section) on how to view the interactive API documentation (e.g., using Swagger UI or Redoc).
- The documentation accurately reflects the implemented endpoints, request/response formats, and error structures.
