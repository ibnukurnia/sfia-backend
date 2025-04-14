-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_trainings
ADD COLUMN IF NOT EXISTS assessment_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_trainings
DROP COLUMN IF EXISTS assessment_id;
-- +goose StatementEnd
