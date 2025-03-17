-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_duj_answers
DROP COLUMN IF EXISTS detail;

ALTER TABLE participant_duj_answers
DROP COLUMN IF EXISTS job;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
