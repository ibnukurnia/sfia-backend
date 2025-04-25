-- +goose Up
-- +goose StatementBegin

CREATE TYPE role_enum as ENUM ('user','admin','manager');

CREATE TABLE IF NOT EXISTS users(
    uuid UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    pn VARCHAR(255) NOT NULL,
    password VARCHAR(500) NOT NULL,
    role_access role_enum DEFAULT 'user',
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
