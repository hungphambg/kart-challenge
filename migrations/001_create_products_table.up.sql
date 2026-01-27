-- kart.products definition

CREATE TABLE `products` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `merchant_code` varchar(255) NOT NULL,
    `name` varchar(255) NOT NULL,
    `description` text,
    `price` decimal(10,2) NOT NULL,
    `image_url` varchar(255) DEFAULT NULL,
    `stock` int unsigned NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_merchant_id` (`merchant_code`,`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;