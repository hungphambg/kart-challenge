INSERT INTO `product` (`merchant_code`, `name`, `description`, `price`, `image_url`, `stock`) VALUES
('MERCH001', 'Laptop Pro', 'Powerful laptop for professionals', 1200.00, 'http://example.com/laptop-pro.jpg', 50),
('MERCH001', 'Mechanical Keyboard', 'RGB mechanical keyboard with brown switches', 80.00, 'http://example.com/keyboard.jpg', 150),
('MERCH002', 'Wireless Mouse', 'Ergonomic wireless mouse with long battery life', 25.00, 'http://example.com/mouse.jpg', 200),
('MERCH003', 'Monitor 4K', '27 inch 4K UHD monitor with HDR', 350.00, 'http://example.com/monitor.jpg', 30),
('MERCH004', 'Webcam HD', 'Full HD webcam for video conferencing', 45.00, 'http://example.com/webcam.jpg', 100);

-- Insert sample carts (assuming current time for created_at and updated_at)
INSERT INTO `cart` (`device_id`, `created_at`, `updated_at`) VALUES
('DEVICE001', NOW(), NOW()),
('DEVICE002', NOW(), NOW());

-- Insert sample cart items (referencing product and cart IDs)
-- For DEVICE001: Laptop Pro (ID 1) quantity 1, Wireless Mouse (ID 3) quantity 2
INSERT INTO `cart_item` (`cart_id`, `product_id`, `quantity`, `price_at_addition`) VALUES
(1, 1, 1, 1200.00),
(1, 3, 2, 25.00);

-- For DEVICE002: Mechanical Keyboard (ID 2) quantity 1, Webcam HD (ID 5) quantity 1
INSERT INTO `cart_item` (`cart_id`, `product_id`, `quantity`, `price_at_addition`) VALUES
(2, 2, 1, 80.00),
(2, 5, 1, 45.00);

INSERT INTO kart_challenge_db.coupon
(id, code, `type`, value, description, expires_at, is_active, created_at, updated_at) VALUES
(1, 'HAPPYHOURS', 'percentage', 0.18, '18% off for all products', '2026-12-31 23:59:59', 1, NULL, NULL),
(2, 'SALE18', 'percentage', 0.18, '18% off for all products', '2026-12-31 23:59:59', 1, NULL, NULL),
(3, 'BOGO2026', 'buy_one_get_one', 1.00, 'Buy one get one free', '2026-06-30 23:59:59', 1, NULL, NULL),
(4, 'SAVE5', 'fixed_amount', 5.00, 'Save $5 on your order', '2026-03-31 23:59:59', 1, NULL, NULL);