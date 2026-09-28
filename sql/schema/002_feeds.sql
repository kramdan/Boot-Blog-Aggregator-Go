-- +goose Up
CREATE TABLE feeds (
    id UUID NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    name text NOT NULL,
    url text NOT NULL UNIQUE,
    user_id UUID NOT NULL,
    PRIMARY KEY(id),
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;