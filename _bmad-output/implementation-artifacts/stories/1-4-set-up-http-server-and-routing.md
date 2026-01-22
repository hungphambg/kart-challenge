# Story 1.4: Set up HTTP server and routing

**Epic:** Core API Infrastructure

**Description:**
As a developer, I want to set up a basic HTTP server with a routing library, so that I can define API endpoints and handle incoming requests in a structured way. I also want to include basic middleware for logging and error handling.

**Acceptance Criteria:**
- A routing library (e.g., `gorilla/mux` or `chi`) is added to the project.
- The main application starts an HTTP server on the configured port.
- A `handler` package is created in the `internal` directory.
- Middleware for request logging (e.g., logging the method, path, and duration of each request) is implemented.
- Middleware for basic error handling (e.g., catching panics and returning a `500 Internal Server Error` response) is implemented.
