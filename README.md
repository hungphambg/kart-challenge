# Kart Challenge API

This project implements a Go API for a shopping cart and order management system, built as part of the Kart Challenge. It features product listing, cart management, order placement, and coupon validation.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Running Database Migrations](#running-database-migrations)
- [API Endpoints](#api-endpoints)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Testing](#testing)

## Getting Started

Follow these instructions to set up, build, and run the application.

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.25 or higher ([Download Go](https://go.dev/dl/))
- **Docker & Docker Compose**: ([Install Docker Engine](https://docs.docker.com/engine/install/) and [Install Docker Compose](https://docs.docker.com/compose/install/))

## Starting the Infrastructure

To start Mysql Database and Redis services:

```bash
docker-compose -f docker-compose.infra.yaml up -d
```
The `-d` flag runs the services in detached mode (in the background).

## Stopping the Infrastructure

To stop and remove the services:

```bash
docker-compose -f docker-compose.infra.yaml down
```

## Start coupon producer to reading files set coupon to redis
Make sure that 3 files stored in ./pipeline folder
```bash
docker-compose -f docker-compose.services-producer.yaml up -d
```

## Start api service for make order
```bash
docker-compose -f docker-compose.services-producer.yaml up -d
```


### Running Database Migrations

Running mysql script in ./migrations

### API Endpoints

Here's a brief overview of the main API endpoints:

-   `GET /`: Hello World endpoint.
-   `GET /health`: Checks database connectivity (returns 200 OK if connected, 503 Service Unavailable otherwise).
-   `GET /products`: Retrieve a list of all products.
-   `GET /products/{id}`: Retrieve details of a specific product by ID.
-   `GET /cart`: Retrieve user's current cart; creates one if it doesn't exist. Requires `DeviceID` header.
-   `POST /cart/items`: Add or update an item in the cart. Requires `DeviceID` header and `product_id`, `quantity` in body.
-   `PUT /cart/items/{product_id}`: Update quantity of a specific item in the cart. Requires `DeviceID` header and `quantity` in body.
-   `DELETE /cart/items/{product_id}`: Remove a specific item from the cart. Requires `DeviceID` header.
-   `DELETE /cart`: Clear all items from the user's cart. Requires `DeviceID` header.
-   `POST /orders`: Place an order from the current cart. Requires `DeviceID` header.
-   `GET /orders`: Retrieve a list of all past orders for a device. Requires `DeviceID` header.
-   `GET /orders/{order_id}`: Retrieve details of a specific order by ID. Requires `DeviceID` header.
-   `POST /coupons/validate`: Validate a coupon code and calculate potential discount. Requires `coupon_code` in body, `DeviceID` optional.

### API Documentation

The OpenAPI specification for this API is available at `api/openapi.yaml`. You can use tools like [Swagger UI](https://swagger.io/tools/swagger-ui/) or [Redoc](https://redoc.ly/redoc/) to view interactive documentation by serving this file.

### Project Structure

-   `cmd/api`: Main application entry point.
-   `internal/config`: Configuration loading.
-   `internal/database`: Database connection and GORM interactions.
-   `internal/handler`: HTTP handlers for API endpoints.
-   `internal/model`: Go structs for data models (Product, Cart, Order, Coupon, etc.).
-   `migrations`: SQL migration scripts.
-   `Dockerfile`: Docker build instructions for the API.
-   `docker-compose.*.yaml`: Local development setup with API and MySQL, redis.

### Testing

*(Instructions on how to run tests will be added here once tests are implemented.)*