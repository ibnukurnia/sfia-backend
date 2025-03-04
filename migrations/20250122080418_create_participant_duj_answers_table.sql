-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant_duj_answers(
    uuid UUID DEFAULT uuid_generate_v4(),
    participant_id UUID,
    job TEXT NOT NULL,
    detail TEXT NOT NULL,
    current_job BOOLEAN,
    have_trouble BOOLEAN,
    trouble_cause TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participant_duj_answers;
-- +goose StatementEnd
