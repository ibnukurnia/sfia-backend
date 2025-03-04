-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS participant_roles(
    uuid UUID DEFAULT uuid_generate_v4(),
    main_role_id UUID,
    secondary_role_id UUID NULL,
    interest_role_id UUID NULL,
    created_at TIMESTAMPTZ DEFAULT now(), 
    updated_at TIMESTAMPTZ DEFAULT now(), 
    deleted_at TIMESTAMPTZ NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participant_roles;
-- +goose StatementEnd
