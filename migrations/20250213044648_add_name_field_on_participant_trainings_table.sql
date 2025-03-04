-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_trainings
ADD COLUMN IF NOT EXISTS name TEXT NULL,
ADD COLUMN IF NOT EXISTS is_needed BOOLEAN DEFAULT false,
ADD COLUMN IF NOT EXISTS priority SMALLINT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_trainings
DROP COLUMN IF EXISTS name,
DROP COLUMN IF EXISTS is_needed,
DROP COLUMN IF EXISTS priority;
-- +goose StatementEnd
