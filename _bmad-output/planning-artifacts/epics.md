---
stepsCompleted:
  - step-01-validate-prerequisites
  - step-02-design-epics
  - step-03-create-stories
inputDocuments:
  - request_email.txt
  - README.md
  - system_context.md
  - _bmad-output/planning-artifacts/backend-stories-breakdown.md
---

# kart-challenge - Epic Breakdown

## Overview

This document provides the complete epic and story breakdown for kart-challenge, decomposing the requirements from the PRD, UX Design if it exists, and Architecture requirements into implementable stories.

## Requirements Inventory

### Functional Requirements

FR1: The system shall display products with images.
FR2: The system shall allow users to add items to the cart.
FR3: The system shall allow users to remove items from the cart.
FR4: The system shall show the correct order total.
FR5: The system shall allow users to increase or decrease the item count in the cart.
FR6: The system shall show an order confirmation after placing an order.
FR7: The system shall have interactive hover and focus states for elements.
FR8: The system shall provide an API server compliant with the provided OpenAPI specification.

### NonFunctional Requirements

NFR1: The solution must be scalable to handle a large number of users.
NFR2: Coupon validation must be robust and secure.
NFR3: The system should be designed for extensibility.
NFR4: The project must include a detailed README with instructions on how to run and test the solution.
NFR5: The project must include a Docker setup for easy environment provisioning.
NFR6: The solution should favor an indigenous approach over complete reliance on AI generation.
NFR7: The application should have a responsive design that adapts to different screen sizes.

### Additional Requirements

- A "user" is identified by a unique device ID, not traditional user accounts.
- Carts are persistent and stored on the server, linked to the device ID.
- The shopping cart is cleared immediately upon successful order placement.
- Coupon validation occurs via a dedicated service before an order is confirmed.
- The coupon service responds with detailed information about the discount.
- The entire order is rejected if any item is out of stock during placement.
- A clear error message is provided to the client specifying out-of-stock items.
- Completed order records are stored indefinitely in the primary database.
- The `HAPPYHOURS` discount code applies an 18% discount to the order total.
- The `BUYGETONE` discount code gives the lowest priced item for free.

### FR Coverage Map

FR1: Epic 1 - The system shall display products with images.
FR2: Epic 2 - The system shall allow users to add items to the cart.
FR3: Epic 2 - The system shall allow users to remove items from the cart.
FR4: Epic 3 - The system shall show the correct order total.
FR5: Epic 2 - The system shall allow users to increase or decrease the item count in the cart.
FR6: Epic 3 - The system shall show an order confirmation after placing an order.
FR7: Epic 1 - The system shall have interactive hover and focus states for elements. (This is more of a frontend concern, but for the backend it means providing the necessary data).
FR8: Epic 1 - The system shall provide an API server compliant with the provided OpenAPI specification.

## Epic List



### Epic 1: Core API Infrastructure & Product Catalog
**Goal:** Establish the foundational API structure and enable users to browse the product catalog.
**FRs covered:** FR1, FR7, FR8

### Epic 2: Shopping Cart Management API
**Goal:** Allow users to manage items in their shopping cart.
**FRs covered:** FR2, FR3, FR5

### Epic 3: Order Management & Discount API
**Goal:** Enable users to apply discounts, place orders, and view order confirmations.
**FRs covered:** FR4, FR6

### Epic 4: API Robustness & Development Experience
**Goal:** Ensure the API is robust, scalable, and provides a smooth developer experience.
**FRs covered:** (This epic is mostly about NFRs)

## Epic 1: Core API Infrastructure & Product Catalog

Establish the foundational API structure and enable users to browse the product catalog.

### Story 1.1: Setup Go Echo Project

As a backend developer,
I want to set up a new Go project with the Echo framework and implement graceful shutdown,
So that I can begin implementing API endpoints and ensure application stability.

**Acceptance Criteria:**
*   **Given** a new project directory
*   **When** I initialize the Go module
*   **Then** a `go.mod` file is created with the correct module path.
*   **And** the Echo framework is added as a dependency.
*   **And** `main.go` contains a basic Echo server setup with a root route.
*   **And** the server starts successfully.
*   **And** the server gracefully shuts down on interrupt signals.

### Story 1.2: Database Integration for Product Catalog (MySQL)

As a backend developer,
I want to integrate with a MySQL database and define the product schema,
So that product information can be persistently stored and accessed.

**Acceptance Criteria:**
*   **Given** a running MySQL instance
*   **When** the application starts
*   **Then** a connection to the MySQL database is successfully established.
*   **And** a `products` table is created with `ID`, `MerchantCode`, `Name`, `Description`, `Price`, and `ImageURL` fields.
*   **And** a `device_merchants` table is created with `DeviceID` and `MerchantCode` fields to map devices to merchants.
*   **And** a data access layer (repository pattern) is implemented for `products`.
*   **And** unit tests confirm successful database connection and basic CRUD operations (create, read, update, delete) for products.

### Story 1.3: Product Catalog Service API

As a client application,
I want to retrieve product information via API endpoints,
So that I can display products to users.

**Acceptance Criteria:**
*   **Given** a running API server with database integration for products
*   **When** I send a GET request to `/products` with `merchant_code` in context
*   **Then** the API responds with a list of all available products for that `merchant_code`, conforming to the OpenAPI specification.
*   **And** when I send a GET request to `/products/{id}` with a valid product ID and `merchant_code` in context
*   **Then** the API responds with details of the specified product for that `merchant_code`, conforming to the OpenAPI specification.
*   **And** unit tests verify the functionality of both `/products` and `/products/{id}` endpoints.

### Story 1.4: Device ID Handling Middleware

As a client application,
I want the API to consistently identify my device,
So that my shopping cart and orders can be associated correctly.

**Acceptance Criteria:**
*   **Given** an API request with a valid `Device-ID` in the header
*   **When** the request passes through the middleware
*   **Then** the `Device-ID` is extracted.
*   **And** the middleware looks up the corresponding `merchant_code` from the `device_merchants` table.
*   **And** both `Device-ID` and `merchant_code` are made available in the request context for downstream handlers.
*   **And** unit tests verify that requests without a `Device-ID` are rejected with an appropriate error.
*   **And** unit tests verify that requests with an invalid `Device-ID` format are rejected with an appropriate error.

## Epic 2: Shopping Cart Management API

**Goal:** Allow users to manage items in their shopping cart.
**FRs covered:** FR2, FR3, FR5
**Additional Requirements:** Carts persistent on server, linked to device ID.

This epic focuses on the core functionality of a shopping cart, enabling users to add, remove, and update items before placing an order. It also ensures the cart's state is persistent across sessions for a given device.

### Story 2.1: Shopping Cart Database Schema and DAL
As a backend developer,
I want to define the database schema for shopping carts and cart items and implement a data access layer,
So that shopping cart data can be persistently stored and managed.

**Acceptance Criteria:**
*   **Given** an existing MySQL database
*   **When** the application starts
*   **Then** a `shopping_carts` table is created with `CartID`, `MerchantCode`, `DeviceID`, `CreatedAt`, and `UpdatedAt` fields.
*   **And** a `cart_items` table is created with `CartItemID`, `CartID`, `ProductID`, `Quantity`, and `PriceAtTimeOfAdd` fields.
*   **And** a data access layer (repository pattern) is implemented for `shopping_carts` and `cart_items`.
*   **And** unit tests confirm basic CRUD operations for shopping carts and cart items.

### Story 2.2: Add Item to Cart API
As a client application,
I want to add a product to my shopping cart,
So that I can purchase it later.

**Acceptance Criteria:**
*   **Given** a valid `Device-ID`, `merchant_code` in context, and a valid product ID and quantity
*   **When** I send a POST request to `/cart/items`
*   **Then** the product is added to the cart associated with the `Device-ID` and `merchant_code`.
*   **And** if the item already exists in the cart, its quantity is increased.
*   **And** the API responds with the updated cart details.
*   **And** unit tests verify adding new items and increasing quantities of existing items.

### Story 2.3: Remove Item from Cart API
As a client application,
I want to remove a specific product from my shopping cart,
So that I can adjust my purchase.

**Acceptance Criteria:**
*   **Given** a valid `Device-ID`, `merchant_code` in context, and a product in the cart
*   **When** I send a DELETE request to `/cart/items/{product_id}`
*   **Then** the specified product is removed from the cart.
*   **And** the API responds with the updated cart details.
*   **And** unit tests verify successful removal of an item.

### Story 2.4: Update Cart Item Quantity API
As a client application,
I want to modify the quantity of a product in my shopping cart,
So that I can fine-tune my order.

**Acceptance Criteria:**
*   **Given** a valid `Device-ID`, `merchant_code` in context, a product in the cart, and a new quantity
*   **When** I send a PUT request to `/cart/items/{product_id}` with the new quantity
*   **Then** the quantity of the specified product in the cart is updated.
*   **And** if the new quantity is zero, the item is removed from the cart.
*   **And** the API responds with the updated cart details.
*   **And** unit tests verify increasing, decreasing, and setting quantity to zero.

### Story 2.5: Get Cart Contents API
As a client application,
I want to view the current contents of my shopping cart,
So that I can review my selections and total.

**Acceptance Criteria:**
*   **Given** a valid `Device-ID`, `merchant_code` in context
*   **When** I send a GET request to `/cart`
*   **Then** the API responds with a list of all products in the cart, their quantities, and the calculated current total.
*   **And** unit tests verify correct cart contents and total calculation for various scenarios.

## Epic 3: Order Management & Discount API

**Goal:** Enable users to apply discounts, place orders, and view order confirmations.
**FRs covered:** FR4 (Show order total), FR6 (Show order confirmation).
**NFRs covered:** NFR2 (Coupon validation robust).
**Additional Requirements:** Coupon validation before order, order rejection if out of stock, `HAPPYHOURS`, `BUYGETONE` discounts, order records stored indefinitely.

This epic focuses on the critical functionalities of applying discounts, placing orders, ensuring order integrity, and managing order history. All operations within this epic will implicitly leverage the `merchant_code` from the request context, ensuring multi-merchant support.

### Story 3.1: Promo Code Validation Service
As a client application,
I want to validate promo codes against a predefined list,
So that I can check if a code is legitimate before proceeding with an order.

**Acceptance Criteria:**
*   **Given** the application starts up
*   **When** the application initializes
*   **Then** the `couponbase1.gz`, `couponbase2.gz`, and `couponbase3.gz` files are downloaded.
*   **And** a Map-Reduce process is executed to identify promo codes that meet the criteria (8-10 characters long, found in at least two files).
*   **And** these valid promo codes are loaded into an efficient in-memory data structure for quick lookup.
*   **Given** a valid `Device-ID`, `merchant_code` in context, and a candidate promo code
*   **When** I send a POST request to `/cart/validate-promo-code` (or similar new endpoint)
*   **Then** the API responds indicating if the promo code is valid based on the in-memory list.
*   **And** unit tests verify the Map-Reduce processing of the files, including edge cases for promo code length and occurrence count.
*   **And** unit tests verify the promo code validation API endpoint.

### Story 3.2: Order Database Schema and DAL
As a backend developer,
I want to define the database schema for orders and order items and implement a data access layer,
So that all placed orders and their contents can be persistently stored for a specific merchant.

**Acceptance Criteria:**
*   **Given** an existing MySQL database
*   **When** the application starts
*   **Then** an `orders` table is created with `OrderID`, `MerchantCode`, `DeviceID`, `TotalAmount`, `DiscountApplied`, `Status`, and `CreatedAt` fields.
*   **And** an `order_items` table is created with `OrderItemID`, `OrderID`, `ProductID`, `Quantity`, and `PriceAtPurchase` fields.
*   **And** a data access layer (repository pattern) is implemented for `orders` and `order_items`.
*   **And** unit tests confirm basic CRUD operations for orders and order items.

### Story 3.3: Place Order API
As a client application,
I want to place an order from my current shopping cart,
So that my purchase is finalized.

**Acceptance Criteria:**
*   **Given** a valid `Device-ID`, `merchant_code` in context, and an active cart with items
*   **When** I send a POST request to `/orders`
*   **Then** an order integrity check is performed to verify all cart items are in stock for the `merchant_code`.
*   **And** if all items are in stock, cart items are transferred to new order items.
*   **And** the shopping cart associated with the `Device-ID` and `merchant_code` is cleared.
*   **And** the API responds with an order confirmation, including the `OrderID`.
*   **And** if any item is out of stock, the entire order is rejected, and a clear error message specifying the out-of-stock items is returned.
*   **And** unit tests verify successful order placement and handling of out-of-stock scenarios.

### Story 3.4: Get Order Confirmation API
As a client application,
I want to retrieve the details of a specific placed order,
So that I can review my purchase confirmation.

**Acceptance Criteria:**
*   **Given** a valid `Device-ID`, `merchant_code` in context, and a valid `OrderID`
*   **When** I send a GET request to `/orders/{order_id}`
*   **Then** the API responds with the complete details of the specified order, including order items and total amount.
*   **And** unit tests verify successful retrieval of an order.

### Story 3.5: Order History API
As a client application,
I want to view all my past orders,
So that I can keep track of my purchase history.

**Acceptance Criteria:**
*   **Given** a valid `Device-ID`, `merchant_code` in context
*   **When** I send a GET request to `/orders`
*   **Then** the API responds with a list of all orders placed by that `Device-ID` for that `merchant_code`.
*   **And** unit tests verify successful retrieval of the order history.

## Epic 4: API Robustness & Development Experience

**Goal:** Ensure the API is robust, scalable, and provides a smooth developer experience.
**NFRs covered:** NFR1 (Scalability), NFR6 (Indigenous approach), NFR4 (Detailed README), NFR5 (Docker setup).
**Additional Requirements:** Concurrency handling, rate limiting.

This epic addresses crucial non-functional requirements to ensure the API is performant, reliable, and easy to develop and deploy. These aspects are particularly important for a scalable multi-merchant system.

### Story 4.1: Concurrency Handling
As a backend developer,
I want to implement proper concurrency handling for critical operations,
So that race conditions are prevented and data integrity is maintained, especially under high load.

**Acceptance Criteria:**
*   **Given** concurrent requests to critical sections (e.g., cart updates, order placement)
*   **When** these operations are performed
*   **Then** appropriate concurrency primitives (e.g., mutexes, channels) are used to prevent race conditions.
*   **And** database transaction isolation levels are configured to ensure order integrity.
*   **And** unit tests simulate concurrent access scenarios and demonstrate correct behavior without data corruption.

### Story 4.2: Rate Limiting Middleware
As a system administrator,
I want to implement API rate limiting based on `Device-ID`,
So that the service is protected from abuse and overload.

**Acceptance Criteria:**
*   **Given** a running API server
*   **When** a `Device-ID` sends requests exceeding a configured threshold
*   **Then** the Echo middleware effectively limits the rate of requests from that `Device-ID`.
*   **And** configurable rate limits are applied to different API endpoints.
*   **And** unit tests verify the rate limiting functionality, including blocking excessive requests and allowing valid ones after the cooldown period.

### Story 4.3: Detailed README & Docker Setup
As a developer,
I want a comprehensive `README.md` and Docker setup,
So that I can easily understand, run, and test the API locally.

**Acceptance Criteria:**
*   **Given** a developer new to the project
*   **When** they read the `README.md`
*   **Then** it clearly explains the API's purpose, how to run it locally (including Docker), and how to run tests.
*   **And** a `Dockerfile` is provided for building the Go API server.
*   **And** a `docker-compose.yml` is provided to orchestrate the API server and a MySQL database for local development.

### Story 4.4: OpenAPI Specification Adherence
As an API consumer,
I want the API to consistently match the provided OpenAPI specification,
So that I can integrate with it reliably.

**Acceptance Criteria:**
*   **Given** the provided `openapi.yaml` specification
*   **When** the API endpoints are implemented
*   **Then** all implemented API endpoints (paths, methods, request/response bodies) strictly adhere to the `openapi.yaml` specification.
*   **And** if available, a Go OpenAPI code generator or validator is integrated into the development process to ensure adherence.
*   **And** unit/integration tests verify that API responses conform to the defined schemas.