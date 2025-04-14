-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS skill_level_thresholds(
    uuid UUID DEFAULT uuid_generate_v4(),
    basic FLOAT NOT NULL DEFAULT 3,
    intermediate FLOAT NOT NULL DEFAULT 5,
    advance FLOAT NOT NULL DEFAULT 7,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS skill_level_thresholds;
-- +goose StatementEnd
