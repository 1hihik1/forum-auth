CREATE TABLE threads
(
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    title     TEXT     NOT NULL,
    content   TEXT     NOT NULL,
    create_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id   INTEGER  NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_threads_user_id ON threads (user_id);
CREATE INDEX idx_threads_create_at ON threads (create_at DESC);