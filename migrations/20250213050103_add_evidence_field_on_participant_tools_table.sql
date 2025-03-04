-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_tools
ADD COLUMN IF NOT EXISTS evidence TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_tools
DROP COLUMN IF EXISTS evidence;
-- +goose StatementEnd
