-- +goose Up
-- +goose StatementBegin
ALTER TABLE skills
ADD COLUMN description TEXT NOT NULL,
ADD COLUMN code TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
