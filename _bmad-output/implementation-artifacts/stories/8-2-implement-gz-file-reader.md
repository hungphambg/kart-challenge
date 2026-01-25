# Story 8.2: Implement GZ file reader

**Epic:** Coupon Producer Service

**Description:**
As a developer, I want to implement a Go application that can open and read the content of a `.gz` file from the local filesystem, so that we can access the compressed coupon data.

**Acceptance Criteria:**
- The producer application can take a file path to a `.gz` file as an argument or configuration.
- The application can successfully open the specified `.gz` file.
- Error handling is in place for scenarios like file not found or permission issues.
