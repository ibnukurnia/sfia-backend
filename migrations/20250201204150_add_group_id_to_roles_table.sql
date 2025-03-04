-- +goose Up
-- +goose StatementBegin
ALTER TABLE roles ADD COLUMN IF NOT EXISTS group_id uuid NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE roles DROP COLUMN IF EXISTS group_id;
-- +goose StatementEnd
