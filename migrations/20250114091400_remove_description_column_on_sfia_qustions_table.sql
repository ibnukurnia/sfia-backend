-- +goose Up
-- +goose StatementBegin
ALTER TABLE sfia_questions
DROP COLUMN description;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
