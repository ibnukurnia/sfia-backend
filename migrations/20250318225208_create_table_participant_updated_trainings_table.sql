-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant_updated_trainings(
    uuid UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    has_certification BOOLEAN DEFAULT false,
    get_certification BOOLEAN DEFAULT false,
    implementation VARCHAR(50) NOT NULL DEFAULT 'SYNC',
    location VARCHAR(50) NOT NULL DEFAULT 'OFFLINE',
    provider VARCHAR(255) NOT NULL,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    assessment_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participant_updated_trainings;
-- +goose StatementEnd
