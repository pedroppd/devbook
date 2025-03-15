CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    nick VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,    
    userpassword VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE followers (
    user_id INT NOT NULL, 
    follower_id INT NOT NULL,
    PRIMARY KEY (user_id, follower_id),
    FOREIGN KEY follower_id REFERENCES users(id) on DELETE CASCADE,
    FOREIGN KEY user_id REFERENCES users(id) on DELETE CASCADE
);
