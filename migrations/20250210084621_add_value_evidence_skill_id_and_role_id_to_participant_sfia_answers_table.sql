-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_sfia_answers
ADD COLUMN IF NOT EXISTS role_id UUID NOT NULL,
ADD COLUMN IF NOT EXISTS skill_id UUID NOT NULL,
ADD COLUMN IF NOT EXISTS value SMALLINT NOT NULL,
ADD COLUMN IF NOT EXISTS evidence TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_sfia_answers
DROP COLUMN IF EXISTS role_id,
DROP COLUMN IF EXISTS skill_id,
DROP COLUMN IF EXISTS value,
DROP COLUMN IF EXISTS evidence;
-- +goose StatementEnd

