-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_departments
ADD COLUMN IF NOT EXISTS assessment_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_departments
DROP COLUMN IF EXISTS assessment_id;
-- +goose StatementEnd