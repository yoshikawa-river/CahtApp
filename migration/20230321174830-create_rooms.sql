
-- +migrate Up
CREATE TABLE IF NOT EXISTS rooms (
    id bigint UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(600) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS messages;