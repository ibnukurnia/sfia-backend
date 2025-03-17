-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sfia_results(
    uuid UUID DEFAULT uuid_generate_v4(),
    skill_id UUID NOT NULL,
    assessment_id UUID NOT NULL,
    score INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sfia_results;
-- +goose StatementEnd
