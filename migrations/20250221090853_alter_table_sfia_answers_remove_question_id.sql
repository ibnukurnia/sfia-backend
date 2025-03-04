-- +goose Up
-- +goose StatementBegin
ALTER TABLE sfia_answers
DROP COLUMN IF EXISTS question_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
