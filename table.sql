#mysql -u root
#CREATE DATABASE todos_db;
#use todos_db;

CREATE TABLE todos(
    id BIGINT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(100),
    priority INT,
    status VARCHAR(100)
);