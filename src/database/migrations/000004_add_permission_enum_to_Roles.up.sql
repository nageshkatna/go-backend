CREATE TYPE permission_enum AS ENUM ('create', 'read', 'update', 'delete');
ALTER TABLE roles DROP COLUMN IF EXISTS permissions;

ALTER TABLE roles
  ADD COLUMN permissions permission_enum[] DEFAULT '{}' NOT NULL;