-- +goose Up
-- +goose StatementBegin

CREATE TABLE parameter_scores (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL, 
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS parameter_scores;
-- +goose StatementEnd
