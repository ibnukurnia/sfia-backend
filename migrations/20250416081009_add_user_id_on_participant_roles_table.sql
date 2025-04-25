-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_roles
ADD COLUMN IF NOT EXISTS user_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_roles
DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd
