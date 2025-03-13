-- +goose Up
-- +goose StatementBegin

CREATE TABLE training_masters (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    code VARCHAR(50) NOT NULL,
    name TEXT NOT NULL,
    jenjang VARCHAR(50) NOT NULL,
    level VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL,
    mode VARCHAR(50) NOT NULL,
    provider TEXT NOT NULL,
    silabus TEXT NOT NULL,
    skills_id UUID,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT fk_training_master_skills FOREIGN KEY (skills_id) REFERENCES skills(uuid) ON DELETE SET NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS training_masters;

-- +goose StatementEnd
