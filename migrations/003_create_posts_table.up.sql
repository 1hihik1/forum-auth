CREATE TABLE posts
(
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    content   TEXT     NOT NULL,
    create_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    thread_id INTEGER  NOT NULL,
    user_id   INTEGER  NOT NULL,
    FOREIGN KEY (thread_id) REFERENCES threads (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_posts_thread_id ON posts (thread_id);
CREATE INDEX idx_posts_user_id ON posts (user_id);
CREATE INDEX idx_posts_create_at ON posts (create_at DESC);