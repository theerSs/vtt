-- +goose Up
CREATE TABLE rooms (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name text NOT NULL
);

-- +goose Down
DROP TABLE users;