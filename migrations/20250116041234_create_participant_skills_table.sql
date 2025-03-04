-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant_skills(
    uuid UUID DEFAULT uuid_generate_v4(),
    skill_id UUID,
    is_mastered BOOLEAN DEFAULT false,
    used_for SMALLINT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participant_skills;
-- +goose StatementEnd
