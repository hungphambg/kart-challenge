### **Epic 1: Core API Infrastructure & Product Catalog**

**Objective:** Establish the foundational API structure, integrate with the database, and implement the product listing functionality.

*   **Story 1.1: Setup Go Echo Project**
    *   Initialize a new Go module.
    *   Integrate Echo framework.
    *   Configure basic server setup (routes, middleware).
    *   Implement graceful shutdown.
*   **Story 1.2: Database Integration (MySQL)**
    *   Set up MySQL connection.
    *   Define database schema for products (ID, Name, Description, Price, ImageURL, etc.).
    *   Implement a data access layer (repository pattern) for products.
    *   Write unit tests for database connection and basic CRUD operations.
*   **Story 1.3: Product Catalog Service API**
    *   Implement API endpoint to list all products (GET `/products`).
    *   Implement API endpoint to get a single product by ID (GET `/products/{id}`).
    *   Integrate with the Product Catalog Service domain logic.
    *   Ensure API responses conform to OpenAPI spec (where applicable).
    *   Write unit tests for product API endpoints.
*   **Story 1.4: Device ID Handling Middleware**
    *   Implement Echo middleware to extract the `Device-ID` from request headers.
    *   Validate the presence and format of the `Device-ID`.
    *   Make the `Device-ID` available in the request context for downstream handlers.
    *   Write unit tests for the middleware.

### **Epic 2: Shopping Cart Management API**

**Objective:** Implement the core shopping cart functionalities, ensuring persistence and correct state management.

*   **Story 2.1: Shopping Cart Database Schema**
    *   Define database schema for shopping carts (CartID, DeviceID, CreatedAt, UpdatedAt).
    *   Define database schema for cart items (CartItemID, CartID, ProductID, Quantity, PriceAtTimeOfAdd).
    *   Implement data access layer for carts and cart items.
*   **Story 2.2: Add Item to Cart API**
    *   Implement API endpoint to add a product to the cart (POST `/cart/items`).
    *   Handle existing items (increase quantity) vs. new items.
    *   Associate cart with `Device-ID` from header.
    *   Write unit tests.
*   **Story 2.3: Remove Item from Cart API**
    *   Implement API endpoint to remove an item from the cart (DELETE `/cart/items/{product_id}`).
    *   Write unit tests.
*   **Story 2.4: Update Cart Item Quantity API**
    *   Implement API endpoint to increase/decrease item quantity (PUT `/cart/items/{product_id}`).
    *   Handle quantity of zero (remove item).
    *   Write unit tests.
*   **Story 2.5: Get Cart Contents API**
    *   Implement API endpoint to retrieve the current cart contents for a `Device-ID` (GET `/cart`).
    *   Calculate and return the current total.
    *   Write unit tests.
*   **Story 2.6: Cart Persistence Logic**
    *   Ensure cart state is saved to MySQL.
    *   Implement logic to retrieve existing cart for a `Device-ID` on subsequent requests.

### **Epic 3: Order Management & Discount API**

**Objective:** Implement the order placement process, including discount application and order integrity checks.

*   **Story 3.1: Coupon & Discount Service API**
    *   Define database schema for discount codes (Code, Type, Value, IsActive, etc.).
    *   Implement API endpoint to apply a discount code to a cart (POST `/cart/apply-discount`).
    *   Implement discount logic for `HAPPYHOURS` (18% off total).
    *   Implement discount logic for `BUYGETONE` (lowest priced item free).
    *   Return validity, description, and calculated discount/new total.
    *   Write unit tests for discount logic.
*   **Story 3.2: Order Database Schema**
    *   Define database schema for orders (OrderID, DeviceID, TotalAmount, DiscountApplied, Status, CreatedAt).
    *   Define database schema for order items (OrderItemID, OrderID, ProductID, Quantity, PriceAtPurchase).
    *   Implement data access layer for orders and order items.
*   **Story 3.3: Place Order API**
    *   Implement API endpoint to place an order from the current cart (POST `/orders`).
    *   **Order Integrity Check:** Before placing, verify all items in cart are in stock. If not, reject order and return specific out-of-stock items.
    *   Transfer cart items to order items.
    *   Clear the shopping cart upon successful order placement.
    *   Write unit tests for order placement and integrity checks.
*   **Story 3.4: Get Order Confirmation API**
    *   Implement API endpoint to retrieve details of a placed order (GET `/orders/{order_id}`).
    *   Write unit tests.
*   **Story 3.5: Order History API**
    *   Implement API endpoint to retrieve all orders for a given `Device-ID` (GET `/orders`).
    *   Write unit tests.

### **Epic 4: API Robustness & Development Experience**

**Objective:** Address the prioritized robustness concerns and ensure a smooth development and review process.

*   **Story 4.1: Concurrency Handling**
    *   Implement appropriate concurrency primitives (e.g., mutexes, channels) for critical sections (e.g., cart updates, order placement) to prevent race conditions.
    *   Consider database transaction isolation levels for order integrity.
    *   Write unit tests to simulate concurrent access.
*   **Story 4.2: Rate Limiting Middleware**
    *   Implement Echo middleware for API rate limiting based on `Device-ID` or IP address.
    *   Configure sensible rate limits for different endpoints.
    *   Write unit tests for rate limiting.
*   **Story 4.3: Detailed README & Docker Setup**
    *   Create a comprehensive `README.md` explaining the API, how to run it locally (including Docker), and how to run tests.
    *   Create a `Dockerfile` for the Go API server.
    *   Create a `docker-compose.yml` to orchestrate the API server and MySQL database.
*   **Story 4.4: OpenAPI Specification Adherence**
    *   Ensure all implemented API endpoints match the provided OpenAPI specification (`openapi.yaml`).
    *   Consider using a Go OpenAPI code generator or validator if available for Echo.
