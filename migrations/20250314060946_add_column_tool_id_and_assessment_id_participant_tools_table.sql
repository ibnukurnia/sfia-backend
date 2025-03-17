-- +goose Up
-- +goose StatementBegin
ALTER TABLE participant_tools
ADD COLUMN IF NOT EXISTS tool_id UUID;

ALTER TABLE participant_tools
ADD COLUMN IF NOT EXISTS assessment_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE participant_tools
DROP COLUMN IF EXISTS tool_id;

ALTER TABLE participant_tools
DROP COLUMN IF EXISTS assessment_id;
-- +goose StatementEnd
