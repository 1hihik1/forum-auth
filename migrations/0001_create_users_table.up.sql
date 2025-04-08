CREATE TABLE users
(
    id       INTEGER PRIMARY KEY AUTOINCREMENT,
    name     TEXT NOT NULL,
    email    TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    active   BOOLEAN DEFAULT 1
);

CREATE INDEX idx_users_email ON users (email);