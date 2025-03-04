-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS assessments(
    uuid UUID DEFAULT uuid_generate_v4(),
    participant_id UUID,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS assessments;
-- +goose StatementEnd
