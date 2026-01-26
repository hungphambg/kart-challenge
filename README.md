# Kart Challenge API

This project implements a Go API for a shopping cart and order management system, built as part of the Kart Challenge. It features product listing, cart management, order placement, and coupon validation.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Local Development with Docker Compose](#local-development-with-docker-compose)
  - [Running Database Migrations](#running-database-migrations)
  - [Building and Running without Docker](#building-and-running-without-docker)
- [API Endpoints](#api-endpoints)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)
- [Testing](#testing)

## Getting Started

Follow these instructions to set up, build, and run the application.

### Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version 1.22 or higher ([Download Go](https://go.dev/dl/))
- **Docker & Docker Compose**: ([Install Docker Engine](https://docs.docker.com/engine/install/) and [Install Docker Compose](https://docs.docker.com/compose/install/))

### Local Development with Docker Compose

The easiest way to get the API and its MySQL database running locally is using Docker Compose.

1.  **Build and Start Services:**
    ```bash
    docker-compose up --build
    ```
    This command will:
    - Build the Go API Docker image using the `Dockerfile`.
    - Start a MySQL database service.
    - Start the Go API service, connected to the MySQL database.

2.  **Access the API:**
    The API will be accessible at `http://localhost:8080`.

3.  **Stop Services:**
    ```bash
    docker-compose down
    ```

### Running Database Migrations

The project uses GORM's `AutoMigrate` feature, which runs automatically when the API service starts. If you need to manually manage migrations or seed data, you would typically use a separate migration tool or script. For this project, simply starting the API will ensure the database schema is up-to-date.

### Building and Running without Docker

For Go development without Docker:

1.  **Clone the repository:**
    ```bash
    git clone [your-repository-url]
    cd kart-challenge
    ```
2.  **Download Dependencies:**
    ```bash
    go mod tidy
    ```
3.  **Build the application:**
    ```bash
    go build -o kart-api ./cmd/api
    ```
4.  **Run the application:**
    You will need a running MySQL database. Configure its connection string via environment variables or a `.env` file.
    Create a `.env` file in the project root:
    ```
    SERVER_PORT=8080
    DATABASE_URL="user:password@tcp(127.0.0.1:3306)/database_name?parseTime=true"
    ```
    Then run:
    ```bash
    ./kart-api
    ```

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
-   `migrations`: SQL migration scripts (not used directly with GORM AutoMigrate for schema creation, but can be used for initial data seeding if needed).
-   `Dockerfile`: Docker build instructions for the API.
-   `docker-compose.yaml`: Local development setup with API and MySQL.

### Testing

*(Instructions on how to run tests will be added here once tests are implemented.)*