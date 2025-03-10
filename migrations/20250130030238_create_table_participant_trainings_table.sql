-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant_trainings(
    uuid UUID DEFAULT uuid_generate_v4(),
    participant_id UUID,
    training_id UUID,
    need_sertification BOOLEAN,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participant_trainings;
-- +goose StatementEnd
