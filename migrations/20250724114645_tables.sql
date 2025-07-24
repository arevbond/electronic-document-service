-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT NOT NULL UNIQUE,
    hashed_password TEXT NOT NULL,
    created_at   TIMESTAMP WITH TIMEZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS documents (
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id  UUID REFERENCES users(id) ON DELETE CASCADE,
    name      TEXT NOT NULL,
    mime      TEXT NOT NULL,
    file      BOOLEAN NOT NULL,
    public    BOOLEAN NOT NULL DEFAULT false,
    json_data JSONB,
    created_at   TIMESTAMP WITH TIMEZONE DEFAULT now()
);

CREATE TABLE IF NOT EXISTS document_access (
    doc_id   UUID REFERENCES documents(id) ON DELETE CASCADE,
    user_id  UUID REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (doc_id, user_id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS document_access;
DROP TABLE IF EXISTS documents;
DROP TABLE IF EXISTS users;