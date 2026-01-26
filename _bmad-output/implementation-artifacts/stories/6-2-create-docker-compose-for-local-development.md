# Story 6.2: Create docker-compose for local development

**Description:**
As a developer, I want to create a `docker-compose.yaml` file, so that I can easily set up and run the Go API and its MySQL database locally with a single command.

**Acceptance Criteria:**
- [x] A `docker-compose.yaml` file is created in the project root.
- [x] The `docker-compose.yaml` defines two services: one for the Go API and one for the MySQL database.
- [x] The Go API service uses the `Dockerfile` created in Story 6.1.
- [x] The MySQL service uses an official MySQL Docker image.
- [x] Environment variables for the API (e.g., database connection string) are configured to connect to the MySQL service within the Docker network.
- [x] Volumes are used for persistent MySQL data.
- [x] The setup allows for easy local development and testing.

**Tasks/Subtasks:**
- [x] Create `docker-compose.yaml` file with `db` and `api` services.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `docker-compose.yaml` defining a `db` service (MySQL 8.0 with persistent volume) and an `api` service (builds from local Dockerfile, connects to `db` service using environment variables).
  - Mapped API container port 1323 to host port 8080.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - `docker-compose.yaml` is created, enabling easy local development setup for the API and MySQL database.

**File List:**
- `docker-compose.yaml`

**Change Log:**
- Created `docker-compose.yaml`.

**Status:**
- review
