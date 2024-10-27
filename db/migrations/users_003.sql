-- db/migrations/users_nnn.sql

DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    registered_datetime DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    approved INT NOT NULL,
    human_or_service VARCHAR(7) NOT NULL
);

INSERT INTO users (email, password, approved, human_or_service) VALUES ('admin@localhost.com', '$2a$14$UczNlsxPY6cyPe0tQwJv6OVfjjqPdzqDy9s1s4H2cEnCdFgYtAMfS', 1, 'human');
INSERT INTO users (email, password, approved, human_or_service) VALUES ('Smakfullt Mat', '$2a$14$UczNlsxPY6cyPe0tQwJv6OVfjjqPdzqDy9s1s4H2cEnCdFgYtAMfS', 1, 'service');
INSERT INTO users (email, password, approved, human_or_service) VALUES ('Fjordens Delikatesser', '$2a$14$UczNlsxPY6cyPe0tQwJv6OVfjjqPdzqDy9s1s4H2cEnCdFgYtAMfS', 1, 'service');
INSERT INTO users (email, password, approved, human_or_service) VALUES ('Godt & Grønt', '$2a$14$UczNlsxPY6cyPe0tQwJv6OVfjjqPdzqDy9s1s4H2cEnCdFgYtAMfS', 1, 'service');


