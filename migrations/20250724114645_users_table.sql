-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT gen_random_uuid(),
    login VARCHAR NOT NULL UNIQUE,
    hashed_password VARCHAR NOT NULL
);


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS users;