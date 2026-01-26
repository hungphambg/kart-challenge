# Story 8.8: Implement producer scheduler

**Description:**
As an operator, I want to implement a scheduling mechanism to automatically trigger the execution of the coupon producer service for each `couponbaseX.gz` file on a regular basis, so that the coupon data in Kafka and Redis is kept up-to-date.

**Acceptance Criteria:**
- [x] A scheduling mechanism (e.g., a simple shell script triggered by host `cron`, or a dedicated scheduler service in `docker-compose`) is implemented.
- [x] The scheduler triggers the producer service for `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz`.
- [x] Each producer run processes its assigned `.gz` file.
- [x] The scheduler runs on a recurring basis (e.g., daily, as per `coupon_ingestion_pipeline.md`).
- [ ] The scheduling configuration is documented in the pipeline's README. (This will be done in the next step, as part of the overall pipeline README documentation)

**Tasks/Subtasks:**
- [x] Create `pipeline/run_producers.sh` script to execute producers for each `.gz` file.
- [x] Make `pipeline/run_producers.sh` executable.
- [x] Add `scheduler` service to `docker-compose.pipeline.yaml`.
- [x] Modify `pipeline/run_producers.sh` to correctly execute the producer binary within the producer service container.

**Dev Agent Record:**
- **Implementation Plan:**
  - Created `pipeline/run_producers.sh` to iterate through `couponbaseX.gz` files and call the producer for each.
  - Made the script executable using `chmod +x`.
  - Added a `scheduler` service to `docker-compose.pipeline.yaml`. This service builds from the producer Dockerfile, mounts `run_producers.sh` and the `.gz` files, and runs the script daily using a `while true` loop.
  - Adjusted `run_producers.sh` to directly execute the producer binary (`/app/producer/main`) with the correct file paths within the container, instead of using `docker-compose run`.
- **Debug Log:**
  - No issues.
- **Completion Notes:**
  - A scheduling mechanism is implemented using a shell script and a `docker-compose` service to automate the execution of the coupon producers.

**File List:**
- `pipeline/run_producers.sh`
- `docker-compose.pipeline.yaml`

**Change Log:**
- Implemented producer scheduler.

**Status:**
- review
