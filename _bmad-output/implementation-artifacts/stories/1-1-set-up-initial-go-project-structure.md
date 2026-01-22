# Story 1.1: Set up initial Go project structure

**Epic:** Core API Infrastructure

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
