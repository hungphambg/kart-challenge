# Epic 1: Core API Infrastructure

This epic covers setting up the basic Go API server, including routing, configuration, and database connectivity.

---

### Story 1.1: Set up initial Go project structure

**Description:**
As a developer, I want to create a well-organized directory structure for the Go project so that the codebase is easy to navigate, maintain, and scale.

**Acceptance Criteria:**
- A new Go module is initialized.
- The project has a root directory containing a `main.go` file.
- The following subdirectories are created:
  - `cmd/api`: Main application entry point.
  - `internal/config`: For configuration loading.
  - `internal/database`: For database connection and logic.
  - `internal/handler`: For HTTP handlers.
  - `internal/model`: For data models/structs.
  - `internal/service`: For business logic.
- The `main.go` file has a basic "hello world" server to confirm the setup.

---

### Story 1.2: Implement a configuration module

**Description:**
As a developer, I want to create a configuration module that loads settings from environment variables or a config file, so that I can manage application settings (like database credentials and server port) without hardcoding them.

**Acceptance Criteria:**
- A `config` package is created within the `internal` directory.
- The configuration module can read settings from environment variables.
- The application can be configured for different environments (e.g., development, production) using `.env` files.
- The loaded configuration is available as a struct throughout the application.

---

### Story 1.3: Establish database connection and health check

**Description:**
As a developer, I want to establish a connection to the MySQL database using the loaded configuration and expose a `/health` endpoint, so that I can verify that the API is running and the database connection is healthy.

**Acceptance Criteria:**
- A `database` package is created within the `internal` directory.
- The application connects to the MySQL database using the credentials from the configuration module.
- A `/health` endpoint is available on the API.
- When the `/health` endpoint is called, it returns a `200 OK` status if the database connection is successful.
- If the database connection fails, the endpoint returns a `503 Service Unavailable` status.

---

### Story 1.4: Set up HTTP server and routing

**Description:**
As a developer, I want to set up a basic HTTP server with a routing library, so that I can define API endpoints and handle incoming requests in a structured way. I also want to include basic middleware for logging and error handling.

**Acceptance Criteria:**
- The Echo framework is added to the project.
- The main application starts an HTTP server on the configured port.
- A `handler` package is created in the `internal` directory.
- Middleware for request logging (e.g., logging the method, path, and duration of each request) is implemented.
- Middleware for basic error handling (e.g., catching panics and returning a `500 Internal Server Error` response) is implemented.
