CREATE DATABASE sql_test;

CREATE TABLE `user` (
                        `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
                        `name` VARCHAR(20) DEFAULT '',
                        `age` INT(11) DEFAULT '0',
                        PRIMARY KEY(`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

INSERT INTO user(name,age) values ("刘一",21);
INSERT INTO user(name,age) values ("陈二",22);