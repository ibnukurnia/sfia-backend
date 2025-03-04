-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant_sfia_answers(
    uuid UUID DEFAULT uuid_generate_v4(),
    participant_id UUID,
    question_id UUID,
    answer_id UUID,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participant_sfia_answers;
-- +goose StatementEnd
