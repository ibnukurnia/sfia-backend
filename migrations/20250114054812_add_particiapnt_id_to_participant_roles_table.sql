-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_roles
ADD COLUMN participant_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
