CREATE TABLE `USERS` (
    `user_id` INT NOT NULL AUTO_INCREMENT,
    `hashed_password` varchar(128) NOT NULL,
    PRIMARY KEY (`user_id`)
);

CREATE TABLE `COMM` (
    `comm_id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `type_id` INT NOT NULL,
    `value` varchar(50) NOT NULL,
    PRIMARY KEY (`comm_id`)
);

CREATE TABLE `COMM_TYPE` (
    `type_id` INT NOT NULL AUTO_INCREMENT,
    `value` varchar(128) NOT NULL,
    PRIMARY KEY (`type_id`)
);

CREATE TABLE `ALIAS` (
    `alias_id` INT NOT NULL AUTO_INCREMENT,
    `value` varchar(50) NOT NULL,
    `is_user_defined` BOOLEAN NOT NULL,
    PRIMARY KEY (`alias_id`)
);

CREATE TABLE `CHANNEL_USER` (
    `channel_user_id` INT NOT NULL AUTO_INCREMENT,
    `channel_id` INT NOT NULL,
    `comm_id` INT NOT NULL,
    `alias_id` INT NOT NULL,
    `is_owner` BOOLEAN NOT NULL,
    PRIMARY KEY (`channel_user_id`)
);

CREATE TABLE `CHANNEL` (
    `channel_id` INT NOT NULL AUTO_INCREMENT,
    `preference_id` INT NOT NULL,
    `name` varchar(50) NOT NULL,
    PRIMARY KEY (`channel_id`)
);

CREATE TABLE `PREFERENCE` (
    `preference_id` INT NOT NULL AUTO_INCREMENT,
    `start_date` DATETIME NOT NULL,
    `duration_days` INT NOT NULL,
    `user_id` INT NOT NULL,
    `chanel_id` INT NOT NULL,
    PRIMARY KEY (`preference_id`)
);

CREATE TABLE `RESTRICTION` (
    `restriction_id` INT NOT NULL AUTO_INCREMENT,
    `preference_id` INT NOT NULL,
    `type_id` INT NOT NULL,
    `value` varchar(50) NOT NULL,
    PRIMARY KEY (`restriction_id`)
);

CREATE TABLE `CONTROL_TYPE` (
    `control_type_id` INT NOT NULL AUTO_INCREMENT,
    `value` varchar(50) NOT NULL,
    PRIMARY KEY (`control_type_id`)
);

CREATE TABLE `RESTRICTION_TYPE` (
    `restriction_type_id` INT NOT NULL AUTO_INCREMENT,
    `control_type_id` INT NOT NULL,
    `type_value` varchar(50) NOT NULL,
    PRIMARY KEY (`restriction_type_id`)
);

CREATE TABLE `INTERVAL` (
    `interval_id` INT NOT NULL AUTO_INCREMENT,
    `preference_id` INT NOT NULL,
    `start_time_in_minutes` INT NOT NULL,
    `duration` INT NOT NULL,
    PRIMARY KEY (`interval_id`)
);

CREATE TABLE `RESOURCE` (
    `resource_id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(50) NOT NULL,
    `value` varchar(50) NOT NULL,
    `language` varchar(50) NOT NULL,
    PRIMARY KEY (`resource_id`)
);

ALTER TABLE `COMM` ADD CONSTRAINT `COMM_fk0` FOREIGN KEY (`user_id`) REFERENCES `USERS`(`user_id`);

ALTER TABLE `COMM` ADD CONSTRAINT `COMM_fk1` FOREIGN KEY (`type_id`) REFERENCES `COMM_TYPE`(`type_id`);

ALTER TABLE `CHANNEL_USER` ADD CONSTRAINT `CHANNEL_USER_fk0` FOREIGN KEY (`channel_id`) REFERENCES `CHANNEL`(`channel_id`);

ALTER TABLE `CHANNEL_USER` ADD CONSTRAINT `CHANNEL_USER_fk1` FOREIGN KEY (`comm_id`) REFERENCES `COMM`(`comm_id`);

ALTER TABLE `CHANNEL_USER` ADD CONSTRAINT `CHANNEL_USER_fk2` FOREIGN KEY (`alias_id`) REFERENCES `ALIAS`(`alias_id`);

ALTER TABLE `CHANNEL` ADD CONSTRAINT `CHANNEL_fk0` FOREIGN KEY (`preference_id`) REFERENCES `PREFERENCE`(`preference_id`);

ALTER TABLE `PREFERENCE` ADD CONSTRAINT `PREFERENCE_fk0` FOREIGN KEY (`user_id`) REFERENCES `USERS`(`user_id`);

ALTER TABLE `PREFERENCE` ADD CONSTRAINT `PREFERENCE_fk1` FOREIGN KEY (`chanel_id`) REFERENCES `CHANNEL`(`channel_id`);

ALTER TABLE `RESTRICTION` ADD CONSTRAINT `RESTRICTION_fk0` FOREIGN KEY (`preference_id`) REFERENCES `PREFERENCE`(`preference_id`);

ALTER TABLE `RESTRICTION` ADD CONSTRAINT `RESTRICTION_fk1` FOREIGN KEY (`type_id`) REFERENCES `COMM_TYPE`(`type_id`);

ALTER TABLE `RESTRICTION_TYPE` ADD CONSTRAINT `RESTRICTION_TYPE_fk0` FOREIGN KEY (`control_type_id`) REFERENCES `CONTROL_TYPE`(`control_type_id`);

ALTER TABLE `INTERVAL` ADD CONSTRAINT `INTERVAL_fk0` FOREIGN KEY (`preference_id`) REFERENCES `PREFERENCE`(`preference_id`);

