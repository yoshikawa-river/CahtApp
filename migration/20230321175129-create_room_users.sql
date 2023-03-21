
-- +migrate Up
CREATE TABLE IF NOT EXISTS room_users (
    id bigint UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    room_id bigint UNSIGNED NOT NULL,
    user_id bigint UNSIGNED NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    FOREIGN KEY (room_id) REFERENCES rooms (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS room_users;