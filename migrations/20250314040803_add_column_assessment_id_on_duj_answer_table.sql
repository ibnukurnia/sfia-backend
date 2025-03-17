-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_duj_answers
ADD COLUMN IF NOT EXISTS assessment_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_duj_answers
DROP COLUMN IF EXISTS assessment_id;
-- +goose StatementEnd
