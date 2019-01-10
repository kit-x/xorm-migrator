CREATE TABLE users (
                     id BIGINT UNSIGNED NOT NULL,
                     nick VARCHAR(255) NOT NULL,
                     avatar VARCHAR(255) NOT NULL,
                     gender TINYINT NOT NULL,
                     age TINYINT UNSIGNED NOT NULL,
                     profession TINYINT NOT NULL,
                     interest_ids VARCHAR(255) DEFAULT '' NOT NULL,
                     PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
