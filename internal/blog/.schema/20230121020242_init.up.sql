CREATE TABLE `articles` (
    `id` VARCHAR(128) CHARACTER SET ascii NOT NULL,
    `version` INT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255),
    `date` DATE,
    `published_at` TIMESTAMP,
    `created_at` TIMESTAMP,
    `updated_at` TIMESTAMP,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`, `version`)
);

CREATE TABLE `mapping_articles_source_articles` (
    `article_id` VARCHAR(128) CHARACTER SET ascii NOT NULL,
    `article_version` INT NOT NULL,
    `article_source_id` VARCHAR(255) NOT NULL,
    `article_source_version` VARCHAR(255) CHARACTER SET ascii NOT NULL,
    `meta` TEXT NOT NULL,
    CONSTRAINT FOREIGN KEY `fk_mapping_articles_source_articles_article_id` (`article_id`, `article_version`) REFERENCES `articles` (`id`, `version`),
    PRIMARY KEY (
        `article_id`,
        `article_version`
    )
);

CREATE TABLE `tags` (
    `id` VARCHAR(255) NOT NULL PRIMARY KEY
);

CREATE TABLE `mapping_articles_tags` (
    `article_id` VARCHAR(128) CHARACTER SET ascii NOT NULL,
    `article_version` INT NOT NULL,
    `tag_id` VARCHAR(255) NOT NULL,
    CONSTRAINT FOREIGN KEY `fk_mapping_articles_tags_article_id` (`article_id`, `article_version`) REFERENCES `articles` (`id`, `version`),
    CONSTRAINT FOREIGN KEY `fk_mapping_articles_tags_tag_id` (`tag_id`) REFERENCES `tags` (`id`),
    PRIMARY KEY (
        `article_id`,
        `article_version`,
        `tag_id`
    )
);

CREATE TABLE `articles_search_index` (
    `article_id` VARCHAR(128) CHARACTER SET ascii NOT NULL,
    `current_article_version` INT NOT NULL,
    `tags` TEXT,
    `date` DATE,
    FULLTEXT (`tags`) WITH PARSER ngram,
    CONSTRAINT FOREIGN KEY `fk_articles_search_index_article_id` (`article_id`) REFERENCES `articles` (`id`),
    PRIMARY KEY (
        `article_id`
    )
);