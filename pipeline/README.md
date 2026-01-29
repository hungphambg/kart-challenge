# Coupon Ingestion Pipeline Local Setup

This README provides instructions for setting up and interacting with the local Kafka and Redis infrastructure required for the coupon ingestion pipeline.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Starting the Pipeline Infrastructure](#starting-the-pipeline-infrastructure)
- [Stopping the Pipeline Infrastructure](#stopping-the-pipeline-infrastructure)
- [Kafka Operations](#kafka-operations)
  - [List Topics](#list-topics)
  - [View Topic Messages](#view-topic-messages)
- [Redis Operations](#redis-operations)
  - [Connect to Redis CLI](#connect-to-redis-cli)
  - [Check Redis Keys](#check-redis-keys)
- [Accessing Service Logs](#accessing-service-logs)

## Prerequisites

- **Docker & Docker Compose**: ([Install Docker Engine](https://docs.docker.com/engine/install/) and [Install Docker Compose](https://docs.docker.com/compose/install/))

## Starting the Pipeline Infrastructure

To start Mysql Database and Redis services:

```bash
docker-compose -f docker-compose.infra.yaml up -d
```
The `-d` flag runs the services in detached mode (in the background).

## Stopping the Pipeline Infrastructure

To stop and remove the services:

```bash
docker-compose -f docker-compose.infra.yaml down
```

## Start coupon producer to reading files  set coupon to redis
Make sure that 3 files stored in ./pipeline folder
```bash
docker-compose -f docker-compose.services-producer.yaml up -d
```

## Start api service for make order

```bash
docker-compose -f docker-compose.services-producer.yaml up -d
```

[//]: # ()
[//]: # (## Kafka Operations)

[//]: # ()
[//]: # (The Kafka broker is accessible on `localhost:9093`.)

[//]: # ()
[//]: # (### List Topics)

[//]: # ()
[//]: # (To list all Kafka topics:)

[//]: # ()
[//]: # (```bash)

[//]: # (docker-compose -f docker-compose.pipeline.yaml exec kafka kafka-topics --bootstrap-server kafka:9092 --list)

[//]: # (```)

[//]: # (&#40;Note: `kafka:9092` is the internal Docker network address for the Kafka broker&#41;)

[//]: # ()
[//]: # (### View Topic Messages)

[//]: # ()
[//]: # (To view messages in a specific topic &#40;e.g., `couponbase1.gz`&#41;:)

[//]: # ()
[//]: # (```bash)

[//]: # (docker-compose -f docker-compose.pipeline.yaml exec kafka kafka-console-consumer --bootstrap-server kafka:9092 --topic couponbase1.gz --from-beginning)

[//]: # (```)

[//]: # (Press `Ctrl+C` to exit the consumer.)

[//]: # (## Redis Operations)

[//]: # ()
[//]: # (The Redis service is accessible on `localhost:6379`.)

[//]: # ()
[//]: # (### Connect to Redis CLI)

[//]: # ()
[//]: # (To connect to the Redis command-line interface:)

[//]: # ()
[//]: # (```bash)

[//]: # (docker-compose -f docker-compose.pipeline.yaml exec redis redis-cli)

[//]: # (```)

[//]: # ()
[//]: # (### Check Redis Keys)

[//]: # ()
[//]: # (Once connected to the Redis CLI, you can use commands like:)

[//]: # ()
[//]: # (- `KEYS *`: List all keys.)

[//]: # (- `GET <key_name>`: Get the value of a specific key.)

[//]: # (- `HGETALL <hash_key>`: Get all fields and values in a hash.)

[//]: # ()
[//]: # (## Accessing Service Logs)

[//]: # ()
[//]: # (To view the logs of a specific service &#40;e.g., `kafka`&#41;:)

[//]: # ()
[//]: # (```bash)

[//]: # (docker-compose -f docker-compose.pipeline.yaml logs -f kafka)

[//]: # (```)

[//]: # (Replace `kafka` with `zookeeper` or `redis` to view their respective logs.)

[//]: # (Press `Ctrl+C` to stop tailing the logs.)
