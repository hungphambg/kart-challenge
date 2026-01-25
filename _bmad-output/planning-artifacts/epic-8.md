# Epic 8: Coupon Producer Service

This epic covers the creation of the Go application that reads the large coupon files, extracts potential codes, and publishes them to Kafka topics.

---

### Story 8.1: Set up producer Go module - new

**Description:**
As a developer, I want to create a new, independent Go module for the coupon producer service, so that it can be developed, built, and deployed as a separate microservice.

**Acceptance Criteria:**
- A new directory for the producer service (e.g., `pipeline/producer`) is created.
- A `go.mod` file is initialized within this directory.
- The basic `main.go` entry point is set up, ready to integrate file reading and Kafka publishing logic.

---

### Story 8.2: Implement GZ file reader

**Description:**
As a developer, I want to implement a Go application that can open and read the content of a `.gz` file from the local filesystem, so that we can access the compressed coupon data.

**Acceptance Criteria:**
- The producer application can take a file path to a `.gz` file as an argument or configuration.
- The application can successfully open the specified `.gz` file.
- Error handling is in place for scenarios like file not found or permission issues.

---

### Story 8.3: Implement stream decompression logic

**Description:**
As a developer, I want to implement logic to decompress the `.gz` file on the fly and read its content line by line, so that we can process large files efficiently without loading the entire uncompressed content into memory.

**Acceptance Criteria:**
- The Go application utilizes a library (e.g., `compress/gzip` package) to decompress the file stream.
- The decompressed content is read line by line.
- The application's memory footprint remains low during processing of large files.
- Error handling for decompression issues is implemented.

---

### Story 8.4: Implement coupon code identification

**Description:**
As a developer, I want to implement logic to scan the streamed content and identify potential coupon codes, so that only relevant strings are passed to the Kafka topic.

**Acceptance Criteria:**
- A regular expression or equivalent string matching logic is used to identify sequences of 8 to 10 alphanumeric characters.
- Each identified sequence is considered a potential coupon code.
- The identification logic is applied to each line of the decompressed content.
- Leading/trailing whitespace is trimmed from potential coupon codes.

---

### Story 8.5: Implement Kafka producer

**Description:**
As a developer, I want to implement a Kafka producer client in Go, so that identified coupon codes can be reliably published to their respective Kafka topics for further processing.

**Acceptance Criteria:**
- The producer service includes a Kafka client library for Go (e.g., `confluentinc/confluent-kafka-go` or `segmentio/kafka-go`).
- For each identified coupon code, a message is published to the Kafka topic named after the source file (e.g., `couponbase1.gz`).
- Messages are published asynchronously to maximize throughput.
- Error handling is implemented for Kafka connection issues and publishing failures (e.g., retries, logging).
- The producer is configured to connect to the Kafka broker running in the `docker-compose.pipeline.yaml` setup.

---

### Story 8.6: Create producer Dockerfile

**Description:**
As a developer, I want to create a `Dockerfile` for the Go coupon producer service, so that it can be containerized and easily run within the `docker-compose.pipeline.yaml` environment.

**Acceptance Criteria:**
- A `Dockerfile` is created within the producer service's directory (`pipeline/producer`).
- The Dockerfile builds a self-contained image for the Go producer application.
- It uses a multi-stage build to minimize the final image size.
- The producer application runs as a non-root user inside the container.
- The Dockerfile includes any necessary environment variables or configurations for the producer (e.g., Kafka broker address).

---

### Story 8.7: Integrate producer into docker-compose.pipeline.yaml

**Description:**
As a developer, I want to integrate the coupon producer service into the `docker-compose.pipeline.yaml` file, so that it can be launched and managed as part of the local pipeline infrastructure alongside Kafka and Redis.

**Acceptance Criteria:**
- The `docker-compose.pipeline.yaml` file is updated to include a service for the coupon producer.
- The producer service's Docker image is built using the Dockerfile created in Story 8.6.
- The producer service is configured to connect to the Kafka broker service defined in the same `docker-compose` file.
- The producer service can be successfully started and stopped via `docker-compose`.
- The `couponbaseX.gz` files are mounted as volumes into the producer container, so it can read them.

---

### Story 8.8: Implement producer scheduler

**Description:**
As an operator, I want to implement a scheduling mechanism to automatically trigger the execution of the coupon producer service for each `couponbaseX.gz` file on a regular basis, so that the coupon data in Kafka and Redis is kept up-to-date.

**Acceptance Criteria:**
- A scheduling mechanism (e.g., a simple shell script triggered by host `cron`, or a dedicated scheduler service in `docker-compose`) is implemented.
- The scheduler triggers the producer service for `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz`.
- Each producer run processes its assigned `.gz` file.
- The scheduler runs on a recurring basis (e.g., daily, as per `coupon_ingestion_pipeline.md`).
- The scheduling configuration is documented in the pipeline's README.
