CREATE TABLE user (
    `id` INT NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(100) NOT NULL,
    `username` VARCHAR(100) NOT NULL,
    `password` VARCHAR(200) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE(`username`),
    UNIQUE(`email`)
);

CREATE TABLE todo (
    `id` INT NOT NULL AUTO_INCREMENT,
    `content` VARCHAR(300) NOT NULL,
    `user_id` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `completed_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
);
