-- migrate:up
CREATE TABLE `league` ( 
    `id` CHAR(36) NOT NULL ,
    `clubname` VARCHAR(100) NOT NULL,
    `points` INT NOT NULL,
    `played` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP on update CURRENT_TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- migrate:down
DROP TABLE `league`
