-- +goose Up
-- +goose StatementBegin
ALTER TABLE assessments
ADD COLUMN IF NOT EXISTS status VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE assessments
DROP COLUMN IF EXISTS status;
-- +goose StatementEnd
