-- kart_challenge_db.carts definition

CREATE TABLE `cart` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT,
     `device_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
     `created_at` datetime(3) DEFAULT NULL,
     `updated_at` datetime(3) DEFAULT NULL,
     PRIMARY KEY (`id`),
     KEY `carts_device_id_IDX` (`device_id`,`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- kart_challenge_db.coupons definition

CREATE TABLE `coupon` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
   `code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
   `type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
   `value` decimal(20,2) NOT NULL,
   `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
   `expires_at` datetime(3) DEFAULT NULL,
   `is_active` tinyint(1) NOT NULL DEFAULT '1',
   `created_at` datetime(3) DEFAULT NULL,
   `updated_at` datetime(3) DEFAULT NULL,
   PRIMARY KEY (`id`),
   UNIQUE KEY `uni_coupons_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- kart_challenge_db.orders definition

CREATE TABLE `order` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `device_id` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `status` tinyint unsigned DEFAULT NULL,
  `total_amount` decimal(20,2) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- kart_challenge_db.products definition

CREATE TABLE `product` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `merchant_code` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
    `price` decimal(20,2) DEFAULT NULL,
    `image_url` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `stock` int unsigned DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- kart_challenge_db.cart_items definition

CREATE TABLE `cart_item` (
      `id` bigint unsigned NOT NULL AUTO_INCREMENT,
      `cart_id` bigint unsigned DEFAULT NULL,
      `product_id` bigint unsigned DEFAULT NULL,
      `quantity` int unsigned DEFAULT NULL,
      `price_at_addition` decimal(20,2) DEFAULT NULL,
      PRIMARY KEY (`id`),
      UNIQUE KEY `uidx_cart_product` (`cart_id`,`product_id`) USING BTREE,
      CONSTRAINT `fk_carts_items` FOREIGN KEY (`cart_id`) REFERENCES `carts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


-- kart_challenge_db.order_items definition

CREATE TABLE `order_item` (
       `id` bigint unsigned NOT NULL AUTO_INCREMENT,
       `order_id` bigint unsigned DEFAULT NULL,
       `product_id` bigint unsigned DEFAULT NULL,
       `quantity` int unsigned DEFAULT NULL,
       `unit_price` longtext,
       PRIMARY KEY (`id`),
       UNIQUE KEY `uidx_order_product` (`order_id`,`product_id`) USING BTREE,
       CONSTRAINT `fk_orders_items` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;