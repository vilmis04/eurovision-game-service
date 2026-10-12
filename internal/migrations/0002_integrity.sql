-- +goose Up
-- Data integrity: unique keys, indexes and a real text[] for group.members.
-- Duplicates that the old code could create are collapsed first, otherwise the
-- unique indexes below could not be built. Take a backup before running this
-- on a database that holds data you want to keep.

-- group.members must be a real text[]. Groups and members are reset every
-- season, so a table that still has the old column type is recreated rather
-- than converted. An existing text[] column is left alone.
-- +goose StatementBegin
DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_schema = current_schema()
          AND table_name = 'group' AND column_name = 'members'
          AND data_type <> 'ARRAY'
    ) THEN
        DROP TABLE "group";
        CREATE TABLE "group" (
            id          SERIAL PRIMARY KEY,
            name        VARCHAR(20) NOT NULL,
            owner       VARCHAR(50) NOT NULL,
            members     TEXT[] NOT NULL,
            datecreated TIMESTAMP NOT NULL
        );
    END IF;
END
$$;
-- +goose StatementEnd

-- score: one row per user, country and year. Keep the row with the most progress.
DELETE FROM score s
USING (
    SELECT id,
           ROW_NUMBER() OVER (
               PARTITION BY "user", country, year
               ORDER BY infinal DESC, position DESC, id ASC
           ) AS rn
    FROM score
) d
WHERE s.id = d.id AND d.rn > 1;

CREATE UNIQUE INDEX IF NOT EXISTS score_user_country_year_key ON score ("user", country, year);
CREATE INDEX IF NOT EXISTS score_user_year_idx ON score ("user", year);

-- country: one row per year and name (the code and the seed migrations rely on it).
DELETE FROM country c
USING (
    SELECT id,
           ROW_NUMBER() OVER (
               PARTITION BY year, name
               ORDER BY isinfinal DESC, score DESC, id ASC
           ) AS rn
    FROM country
) d
WHERE c.id = d.id AND d.rn > 1;

CREATE UNIQUE INDEX IF NOT EXISTS country_year_name_key ON country (year, name);

-- group: a name is unique per owner. Later duplicates get the id appended.
UPDATE "group" g
SET name = LEFT(g.name, 12) || '-' || g.id
FROM (
    SELECT id, ROW_NUMBER() OVER (PARTITION BY owner, name ORDER BY id) AS rn
    FROM "group"
) d
WHERE g.id = d.id AND d.rn > 1;

CREATE UNIQUE INDEX IF NOT EXISTS group_owner_name_key ON "group" (owner, name);

-- admin_config: the code only ever reads row 1.
-- +goose StatementBegin
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'admin_config_single_row') THEN
        ALTER TABLE admin_config ADD CONSTRAINT admin_config_single_row CHECK (id = 1) NOT VALID;
    END IF;
END
$$;
-- +goose StatementEnd

-- +goose Down
ALTER TABLE admin_config DROP CONSTRAINT IF EXISTS admin_config_single_row;
DROP INDEX IF EXISTS group_owner_name_key;
DROP INDEX IF EXISTS country_year_name_key;
DROP INDEX IF EXISTS score_user_year_idx;
DROP INDEX IF EXISTS score_user_country_year_key;
