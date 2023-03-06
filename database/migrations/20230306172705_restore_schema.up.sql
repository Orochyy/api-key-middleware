CREATE TABLE `auth` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `api-key` varchar(32) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `username` varchar(64) NOT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_profile` (
    `user_id` bigint(20) NOT NULL,
    `first_name` varchar(32) NOT NULL,
    `last_name` varchar(64) NOT NULL,
    `phone` varchar(64) NOT NULL,
    `address` varchar(64) NOT NULL,
    `city` varchar(64) NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_data` (
    `user_id` bigint(20) NOT NULL,
    `school` varchar(32) NOT NULL,
    PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
