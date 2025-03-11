-- +goose Up
-- +goose StatementBegin
CREATE TYPE category_enum AS ENUM ('role', 'skill');

CREATE TABLE tresholds (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL, 
    category category_enum NOT NULL,
    treshold_from FLOAT NOT NULL DEFAULT 0,
    treshold_to FLOAT NOT NULL DEFAULT 0,
    color TEXT NOT NULL DEFAULT '#000000',
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tresholds;
DROP TYPE IF EXISTS category_enum;
-- +goose StatementEnd
