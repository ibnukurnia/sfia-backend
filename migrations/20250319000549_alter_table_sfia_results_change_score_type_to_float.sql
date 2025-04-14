-- +goose Up
-- +goose StatementBegin
ALTER TABLE sfia_results 
ALTER COLUMN score 
SET DATA TYPE FLOAT USING score::FLOAT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
