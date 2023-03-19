
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id bigint UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name varchar(30) NOT NULL,
    email varchar(319) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS users;