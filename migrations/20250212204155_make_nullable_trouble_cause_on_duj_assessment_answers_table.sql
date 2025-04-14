-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_duj_answers 
ALTER COLUMN trouble_cause DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM participant_duj_answers;
ALTER TABLE participant_duj_answers 
ALTER COLUMN trouble_cause SET NOT NULL;
-- +goose StatementEnd
