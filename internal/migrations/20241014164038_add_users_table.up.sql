CREATE TABLE users (
    id INT  PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOL NOT NULL DEFAULT false,
    created_at DATETIME DEFAULT (now()),
    updated_at DATETIME,
    UNIQUE (email)
);