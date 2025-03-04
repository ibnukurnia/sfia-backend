-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_sfia_answers
DROP COLUMN IF EXISTS answer_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_sfia_answers 
ADD COLUMN IF NOT EXISTS answer_id UUID;
-- +goose StatementEnd
