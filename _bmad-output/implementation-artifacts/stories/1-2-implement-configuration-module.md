# Story 1.2: Implement a configuration module

**Epic:** Core API Infrastructure

**Description:**
As a developer, I want to create a configuration module that loads settings from environment variables or a config file, so that I can manage application settings (like database credentials and server port) without hardcoding them.

**Acceptance Criteria:**
- A `config` package is created within the `internal` directory.
- The configuration module can read settings from environment variables.
- The application can be configured for different environments (e.g., development, production) using `.env` files.
- The loaded configuration is available as a struct throughout the application.
