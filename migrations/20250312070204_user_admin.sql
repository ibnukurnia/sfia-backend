-- +goose Up
-- +goose StatementBegin

CREATE TABLE corporations (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

ALTER TABLE participants ADD COLUMN corporation_id UUID NULL;

ALTER TABLE participants 
ADD CONSTRAINT fk_participants_corporation FOREIGN KEY (corporation_id) 
REFERENCES corporations(uuid) ON DELETE SET NULL;

CREATE TYPE role_access_enum AS ENUM ('user', 'admin');

ALTER TABLE participants ADD COLUMN role_access role_access_enum NOT NULL DEFAULT 'user';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

ALTER TABLE participants DROP COLUMN IF EXISTS role_access;

DROP TYPE IF EXISTS role_access_enum;

ALTER TABLE participants DROP CONSTRAINT IF EXISTS fk_participants_corporation;

ALTER TABLE participants DROP COLUMN IF EXISTS corporation_id;

DROP TABLE IF EXISTS corporations;

-- +goose StatementEnd
