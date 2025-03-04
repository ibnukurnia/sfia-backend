-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_sfia_answers
ADD COLUMN IF NOT EXISTS assessment_id UUID NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_sfia_answers
DROP COLUMN IF EXISTS assessment_id;
-- +goose StatementEnd
