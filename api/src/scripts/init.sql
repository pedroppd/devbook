DROP DATABASE IF EXISTS;

CREATE DATABASE devbook;

USE devbook;

DROP TABLE IF EXISTS;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    nick VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,    
    userpassword VARCHAR(255) NOT NULL UNIQUE,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;