-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_trainings
RENAME COLUMN need_sertification TO need_certification;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_trainings
RENAME COLUMN need_certification TO need_sertification;
-- +goose StatementEnd
