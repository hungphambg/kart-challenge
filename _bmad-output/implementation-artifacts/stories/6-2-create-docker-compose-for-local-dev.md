# Story 6.2: Create docker-compose for local development

**Epic:** Deployment & Documentation

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
