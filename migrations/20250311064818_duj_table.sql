-- +goose Up
-- +goose StatementBegin

UPDATE skills SET uuid = uuid_generate_v4() WHERE uuid IS NULL;

ALTER TABLE skills ADD CONSTRAINT skills_uuid_unique UNIQUE (uuid);

ALTER TABLE skills DROP CONSTRAINT IF EXISTS skills_pkey;
ALTER TABLE skills ADD PRIMARY KEY (uuid);

CREATE TABLE duj (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    job_description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ NULL
);

CREATE TABLE duj_roles (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    duj_id UUID NOT NULL,
    role_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT fk_duj_roles_duj FOREIGN KEY (duj_id) REFERENCES duj(uuid) ON DELETE CASCADE,
    CONSTRAINT fk_duj_roles_role FOREIGN KEY (role_id) REFERENCES roles(uuid) ON DELETE CASCADE
);

CREATE TABLE duj_skills (
    uuid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    duj_id UUID NOT NULL,
    skill_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT fk_duj_skills_duj FOREIGN KEY (duj_id) REFERENCES duj(uuid) ON DELETE CASCADE,
    CONSTRAINT fk_duj_skills_skill FOREIGN KEY (skill_id) REFERENCES skills(uuid) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE duj_skills DROP CONSTRAINT IF EXISTS fk_duj_skills_skill;

DROP TABLE IF EXISTS duj_roles;
DROP TABLE IF EXISTS duj_skills;
DROP TABLE IF EXISTS duj;

ALTER TABLE skills DROP CONSTRAINT IF EXISTS skills_pkey;
ALTER TABLE skills DROP CONSTRAINT IF EXISTS skills_uuid_unique;

-- +goose StatementEnd
