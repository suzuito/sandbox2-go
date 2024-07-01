CREATE TABLE `organizations` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `active` BOOLEAN DEFAULT false,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `principals` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `organization_id` VARCHAR(128) NOT NULL,
    `email` VARCHAR(128) NOT NULL, -- TODO: Rethink MAX length
    `active` BOOLEAN DEFAULT false,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY (`organization_id`, `email`),
    FOREIGN KEY (`organization_id`) REFERENCES `organizations` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);


CREATE TABLE `principal_roles` (
    `principal_id` VARCHAR(128) NOT NULL,
    `role_id` VARCHAR(128) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`principal_id`, `role_id`),
    FOREIGN KEY (`principal_id`) REFERENCES `principals` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
);

CREATE TABLE `principal_password_hashes` (
    `principal_id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `value` VARCHAR(512) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`principal_id`) REFERENCES `principals` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
);
