-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_skills
ADD COLUMN IF NOT EXISTS assessment_id uuid NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_skills
DROP COLUMN IF EXISTS assessment_id;
-- +goose StatementEnd
