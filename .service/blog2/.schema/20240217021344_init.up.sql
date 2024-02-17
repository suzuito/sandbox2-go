CREATE TABLE users (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    -- user_name は英数字からなる文字列。
    -- ユーザーを一意に特定するために、インターネット上へ公開される。
    `user_name` VARCHAR(64) NOT NULL UNIQUE,
    -- display_name は任意の文字セットからなる文字列。
    -- ユーザーの名前。日本人であれば日本語かもしれない。
    `display_name` VARCHAR(64) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP DEFAULT NULL
);

CREATE TABLE articles (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `user_id` VARCHAR(128) NOT NULL,
    `title` VARCHAR(128) NOT NULL, -- TODO: Rethink MAX length
    `summary` TEXT NOT NULL,
    `current_version` INT DEFAULT NULL, -- 外部キーにはできない。不整合の発生をアプリケーション側で防止する
    `published` BOOLEAN NOT NULL DEFAULT FALSE,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE article_versions(
    `article_id` VARCHAR(128) NOT NULL,
    `version` INT NOT NULL,
    `published` BOOLEAN NOT NULL DEFAULT FALSE,
    `published_at` TIMESTAMP DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    PRIMARY KEY (`article_id`, `version`)
);

CREATE TABLE tags (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `name` VARCHAR(64) NOT NULL UNIQUE,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mapping_articles_tags (
    `article_id` VARCHAR(128) NOT NULL,
    `tag_id` VARCHAR(128) NOT NULL,
    FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
    PRIMARY KEY (`article_id`, `tag_id`)
);
