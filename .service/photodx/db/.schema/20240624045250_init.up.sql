-- Admin auth
CREATE TABLE `photo_studios` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `name` VARCHAR(128) NOT NULL, -- TODO: Rethink MAX length
    `active` BOOLEAN DEFAULT false,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `photo_studio_members` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `photo_studio_id` VARCHAR(128) NOT NULL,
    `email` VARCHAR(128) NOT NULL, -- TODO: Rethink MAX length
    `name` VARCHAR(128) NOT NULL, -- TODO: Rethink MAX length
    `active` BOOLEAN DEFAULT false,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY (`photo_studio_id`, `email`),
    FOREIGN KEY (`photo_studio_id`) REFERENCES `photo_studios` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE `photo_studio_member_roles` (
    `photo_studio_member_id` VARCHAR(128) NOT NULL,
    `role_id` VARCHAR(128) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`photo_studio_member_id`, `role_id`),
    FOREIGN KEY (`photo_studio_member_id`) REFERENCES `photo_studio_members` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
);

CREATE TABLE `photo_studio_members_password_change_tokens` (
    `value` VARCHAR(512) PRIMARY KEY NOT NULL,
    `photo_studio_member_id` VARCHAR(128) NOT NULL,
    `ttl_seconds` INT NOT NULL DEFAULT 600,
    FOREIGN KEY (`photo_studio_member_id`) REFERENCES `photo_studio_members` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
);

CREATE TABLE `photo_studio_member_password_hash_values` (
    `photo_studio_member_id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `value` VARCHAR(512) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`photo_studio_member_id`) REFERENCES `photo_studio_members` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
);

-- Admin
CREATE TABLE `line_link_infos` (
    `photo_studio_id` VARCHAR(128) PRIMARY KEY NOT NULL,
    -- コンソール > チャネル基本設定 > チャネルシークレット
    `messaging_api_channel_secret` VARCHAR(128),
    -- コンソール > Messaging API設定 > チャネルアクセストークン
    `long_access_token` VARCHAR(256),
    `active` BOOLEAN DEFAULT false,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `photo_studio_users` (
    `photo_studio_id` VARCHAR(128) NOT NULL,
    `user_id` VARCHAR(128) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`photo_studio_id`, `user_id`)
);

-- User auth
CREATE TABLE `oauth2_loginflow_states` (
    -- Not master data table
    `code` VARCHAR(128) PRIMARY KEY NOT NULL,
    `provider_id` VARCHAR(128) NOT NULL,
    `redirect_url` VARCHAR(512) NOT NULL,
    `expires_at` TIMESTAMP NOT NULL
);

CREATE TABLE `users` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `name` VARCHAR(128),
    `profile_image_url` VARCHAR(512) NOT NULL,
    `initialized_by_user` BOOLEAN NOT NULL,
    `active` BOOLEAN NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE `provider_resource_owners_users_mappings` (
    `provider_id` VARCHAR(128) NOT NULL,
    `resource_owner_id` VARCHAR(128) NOT NULL,
    `user_id` VARCHAR(128) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`provider_id`, `resource_owner_id`),
    UNIQUE KEY (`user_id`, `provider_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE `users_web_push_subscriptions` (
    -- Not master data table
    `endpoint` VARCHAR(512) PRIMARY KEY NOT NULL,
    `user_id` VARCHAR(128) NOT NULL,
    `expiration_time` TIMESTAMP,
    `value` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
);

-- Unused tables
CREATE TABLE `photo_studio_messages` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `photo_studio_id` VARCHAR(128) NOT NULL,
    `poster` VARCHAR(128) NOT NULL,
    `poster_type` VARCHAR(128) NOT NULL,
    `body` TEXT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`photo_studio_id`) REFERENCES `photo_studios` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE `customers` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `name` VARCHAR(128) NOT NULL, -- TODO: Rethink MAX length
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `customers_photostudio_mappings` (
    `customer_id` VARCHAR(128) NOT NULL,
    `photo_studio_id` VARCHAR(128) NOT NULL,
    PRIMARY KEY (`customer_id`, `photo_studio_id`),
    FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    FOREIGN KEY (`photo_studio_id`) REFERENCES `photo_studios` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE `photo_studio_schedules` (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `photo_studio_id` VARCHAR(128) NOT NULL,
    `photo_studio_member_id` VARCHAR(128) NOT NULL,
    `start` TIMESTAMP NOT NULL,
    `end` TIMESTAMP NOT NULL,
    FOREIGN KEY (`photo_studio_id`) REFERENCES `photo_studios` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    FOREIGN KEY (`photo_studio_member_id`) REFERENCES `photo_studio_members` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);