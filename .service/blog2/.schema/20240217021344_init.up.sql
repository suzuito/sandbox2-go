CREATE TABLE articles (
    `id` VARCHAR(128) PRIMARY KEY NOT NULL,
    `title` VARCHAR(128) NOT NULL, -- TODO: Rethink MAX length
    `published` BOOLEAN NOT NULL DEFAULT FALSE,
    `published_at` TIMESTAMP DEFAULT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
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


CREATE TABLE `articles_search_index` (
    `article_id` VARCHAR(128) NOT NULL,
    `tags` TEXT,
    `published` BOOLEAN NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `published_at` TIMESTAMP DEFAULT NULL,
    FULLTEXT (`tags`) WITH PARSER ngram,
    FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT,
    PRIMARY KEY (
        `article_id`
    )
);
