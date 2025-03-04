-- +goose Up
-- +goose StatementBegin
CREATE TABLE role_skills (
    uuid UUID DEFAULT uuid_generate_v4(),
    role_id UUID, 
    skill_id UUID, 
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_skills;
-- +goose StatementEnd
