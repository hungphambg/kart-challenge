# Coupon Ingestion Pipeline Architecture

This document outlines the detailed architecture for the coupon ingestion and validation pipeline, designed to be resilient, scalable, and performant.

## 1. Overview

The core challenge is to validate promotional codes against three large, multi-gigabyte text files, with a valid code being one that appears in at least two of the three files.

To solve this without impacting real-time API performance, we will use a decoupled, asynchronous ingestion pipeline based on a streaming architecture. This pipeline's sole purpose is to process the source files and populate a fast online datastore (Redis), which the live API will query for validation.

## 2. Architectural Diagram

The data flows through the following components:

```
+-------------------+      +-----------------+      +-----------------+      +--------------------+      +-------+
| Scheduled Trigger | ---> |    Producers    | ---> |  Apache Kafka   | ---> |      Consumer      | ---> | Redis |
|      (Cron)       |      | (File Readers)  |      | (Durable Buffer)|      | (Streaming Aggregator) |      |       |
+-------------------+      +-----------------+      +-----------------+      +--------------------+      +-------+
                                                                                                            ^
                                                                                                            |
                                                                                                    +-----------------+
                                                                                                    | Live API Server |
                                                                                                    +-----------------+
```

## 3. Component Breakdown

### 3.1. Scheduled Trigger
-   **Technology**: A system scheduler like `cron`.
-   **Role**: To initiate the entire ingestion process on a recurring basis (e.g., daily).
-   **Action**: Executes the Producer scripts.

### 3.2. Producers (File-Reading Scripts)
-   **Technology**: A standalone Go command-line application.
-   **Execution**: One instance of the script is run for each source file (e.g., `couponbase1.gz`). These can be run in parallel.
-   **Source**: The script reads its assigned `.gz` file from the **local server filesystem**.
-   **Logic**:
    1.  The script stream-reads the file, decompressing it on the fly to keep memory usage low.
    2.  It identifies potential coupon codes (e.g., via a regex for 8-10 character alphanumeric strings).
-   **Output**: For each potential code found, it publishes a message to a dedicated Apache Kafka topic, where the topic name corresponds to the source filename (e.g., `couponbase1.gz`).

### 3.3. Durable Buffer (Apache Kafka)
-   **Role**: To act as a high-durability, persistent buffer between the file-producers and the aggregator-consumer. This ensures that codes found by the producers are not lost if the consumer service fails, and allows for data replay.
-   **Topics**: Three topics will be used, one for each source file (e.g., `couponbase1.gz`, `couponbase2.gz`, `couponbase3.gz`).

### 3.4. Consumer (Streaming Aggregator Service)
-   **Technology**: A long-running, stateful Go service.
-   **Input**: It subscribes to all three Kafka topics as part of a single consumer group. This allows it to receive codes from all three source files.
-   **State Management**: It maintains a persistent, in-memory hash map.
    -   *Key*: The coupon code string.
    -   *Value*: A data structure (e.g., a bitmask or a set) that tracks the unique topics/files the code has been seen in.
-   **Processing Logic**: The service operates in a perpetual micro-batch loop:
    1.  It consumes messages from Kafka for a set duration or batch size (e.g., for 5 minutes, or a batch of 10,000 messages).
    2.  As messages arrive, it updates its in-memory state map.
    3.  At the end of each batch, it iterates through its internal map, identifies the full set of codes that have been seen in **two or more** topics.
-   **Output**: It writes the resulting set of valid codes to Redis. To ensure atomicity, this can be done by writing to a new temporary key in Redis and then atomically renaming it to the primary key used by the live API.

### 3.5. Online Datastore (Redis)
-   **Role**: To provide a near-instantaneous lookup mechanism for the live API server.
-   **Data Structure**: A single Redis `Set` containing all valid coupon codes.
-   **Interaction**: The live API server validates a user-submitted coupon by using the `SISMEMBER` command on this set, which is an extremely fast O(1) operation.
