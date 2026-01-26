# Story 6.4: Ensure API documentation

**Description:**
As a developer, I want to ensure that the API documentation, based on the provided OpenAPI specification, is accessible and up-to-date, so that consumers of the API can easily understand how to interact with it.

**Acceptance Criteria:**
- [x] The `openapi.yaml` file (`api/openapi.yaml`) is kept up-to-date with the implemented API endpoints.
- [x] Instructions are provided in the `README.md` (or a dedicated `docs` section) on how to view the interactive API documentation (e.g., using Swagger UI or Redoc).
- [x] The documentation accurately reflects the implemented endpoints, request/response formats, and error structures.

**Tasks/Subtasks:**
- [x] Create/update `api/openapi.yaml` to reflect all implemented API endpoints and schemas.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created a `api/openapi.yaml` file that outlines all the currently implemented API endpoints, including their parameters, request bodies, and responses.
  - The `README.md` was already updated in Story 6.3 to include instructions on how to view this documentation.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - The API documentation is updated and accessible via `api/openapi.yaml`, reflecting the current state of the implemented endpoints.

**File List:**
- `api/openapi.yaml`

**Change Log:**
- Updated API documentation.

**Status:**
- review