BEGIN;

CREATE TABLE IF NOT EXISTS `products`
(
    `id` varchar(255) NOT NULL,
    `name` varchar(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (id)
);
CREATE UNIQUE INDEX `UK_products_name` ON `products`(`name`);

COMMIT;