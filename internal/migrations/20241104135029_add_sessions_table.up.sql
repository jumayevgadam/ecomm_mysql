CREATE TABLE sessions (
    id VARCHAR(255) PRIMARY KEY NOT NULL,
    user_email VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(512) NOT NULL,
    is_revoked BOOL NOT NULL DEFAULT false,
    created_at DATETIME DEFAULT (NOW()),
    expires_at DATETIME
);