-- +goose Up
-- +goose StatementBegin
ALTER TABLE assessments
ADD COLUMN IF NOT EXISTS year INTEGER NOT NULL; 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE assessments
DROP COLUMN IF EXISTS year; 
-- +goose StatementEnd
