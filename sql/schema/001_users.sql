-- +goose Up
CREATE TABLE users (
    id UUID NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text NOT NULL UNIQUE,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;