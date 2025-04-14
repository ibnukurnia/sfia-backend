-- +goose Up
-- +goose StatementBegin
DROP TYPE IF EXISTS role_skill_requirement;
CREATE TYPE role_skill_requirement AS ENUM ('basic', 'intermediate', 'advance');

ALTER TABLE role_skills
ADD COLUMN IF NOT EXISTS is_mandatory_for_junior BOOLEAN DEFAULT false,
ADD COLUMN IF NOT EXISTS is_mandatory_for_middle BOOLEAN DEFAULT false,
ADD COLUMN IF NOT EXISTS is_mandatory_for_senior BOOLEAN DEFAULT false,
ADD COLUMN IF NOT EXISTS requirement_for_junior role_skill_requirement DEFAULT 'basic',
ADD COLUMN IF NOT EXISTS requirement_for_middle role_skill_requirement DEFAULT 'intermediate',
ADD COLUMN IF NOT EXISTS requirement_for_senior role_skill_requirement DEFAULT 'advance';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE role_skills
DROP COLUMN IF EXISTS is_mandatory_for_junior,
DROP COLUMN IF EXISTS is_mandatory_for_middle,
DROP COLUMN IF EXISTS is_mandatory_for_senior,
DROP COLUMN IF EXISTS requirement_for_junior,
DROP COLUMN IF EXISTS requirement_for_middle,
DROP COLUMN IF EXISTS requirement_for_senior;
-- +goose StatementEnd
