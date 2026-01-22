# System Context: Food Ordering API

This document outlines the high-level architectural decisions and system context for the food ordering API, based on our initial brainstorming session.

## 1. High-Level Architectural Domains

The system will be designed around a service-oriented approach, with the following distinct logical domains:

-   **Product Catalog Service**: Manages all product information, serving as the single source of truth for items that can be ordered.
-   **Shopping Cart Service**: Manages the state of a user's current shopping cart.
-   **Order Management Service**: Handles the transactional process of placing an order, including validation and persistence.
-   **Coupon & Discount Service**: A dedicated service, consulted by the Order Management Service, to handle all logic related to validating and applying discounts.

## 2. Core System Decisions

The following points summarize the key decisions made regarding the system's behavior and design:

### User Identity
-   **Model**: The system will not have traditional user accounts (username/password).
-   **Identifier**: A "user" is identified by a **unique device ID**. This ID will be the key for associating carts and orders with a specific client.
-   **Audience**: The system is intended to be used directly by **customers**.

### Shopping Cart Lifecycle
-   **Persistence**: Carts are persistent and will be stored on the **server**, linked to the user's device ID. This ensures that if the client application is closed and reopened, the cart's contents are retained.
-   **Clearing Event**: The shopping cart associated with a device will be **cleared automatically and immediately upon a successful order placement**.

### Coupon & Discount Application
-   **Interaction**: The user will apply a coupon via an interactive "Apply Coupon" button in the client application.
-   **Validation Flow**: The client will send the coupon code to the **Coupon & Discount Service** *before* the order is confirmed.
-   **Service Response**: The service will respond with detailed information, including the coupon's validity, a human-readable description of the discount, and the calculated discount amount or new total. This information will be displayed to the user for confirmation.

### Order Integrity & Stock Management
-   **Out-of-Stock Handling**: If a user attempts to place an order containing an item that has become unavailable, the system will **reject the entire order**.
-   **Error Feedback**: A clear error message will be sent to the client, specifying which item(s) are out of stock, requiring the user to modify their cart before resubmitting.

### Data Retention
-   **Order Records**: All completed order records will be stored in the primary database **indefinitely**. This allows for long-term access to order history for any given device ID.
