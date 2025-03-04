-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS departments (
    uuid UUID DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL, 
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS departments;
-- +goose StatementEnd
