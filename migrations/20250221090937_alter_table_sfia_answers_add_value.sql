-- +goose Up
-- +goose StatementBegin
ALTER TABLE sfia_answers
ADD COLUMN IF NOT EXISTS value INTEGER NOT NULL DEFAULT 1;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sfia_answers
DROP COLUMN IF EXISTS value;
-- +goose StatementEnd
