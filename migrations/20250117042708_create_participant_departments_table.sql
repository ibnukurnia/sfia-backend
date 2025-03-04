-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant_departments(
    uuid UUID DEFAULT uuid_generate_v4(),
    participant_id UUID,
    department_id UUID,
    department_team_id UUID,
    department_unit_id UUID,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participant_departments;
-- +goose StatementEnd
