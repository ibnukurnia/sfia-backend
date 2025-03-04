-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_skills
ADD COLUMN participant_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_skills
DROP COLUMN participant_id;
-- +goose StatementEnd
